package marathon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

type Service struct {
	BaseURL string
}

func NewService(host net.IP, port uint) (*Service, error) {
	var url = fmt.Sprintf("http://%v:%v", host, port)
	var ms = &Service{BaseURL: url}

	return ms, nil
}

type errorMessage struct {
	message string
}

func (service *Service) HttpGet(path string) ([]byte, error) {
	url := service.BaseURL + path
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	output, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		contentType := response.Header.Get("Content-Type")
		// Marathon returns: Content-Type: "application/json; q=2"
		if strings.Contains(contentType, "application/json") {
			// return the HTTP response body as error message
			return nil, errors.New(string(output))
		}
		var em errorMessage
		err = json.Unmarshal(output, &em)
		if err != nil {
			return nil, errors.New(string(output))
		}
		return nil, errors.New(em.message)
	}

	return output, nil
}

func (service *Service) HttpPost(path string, body io.Reader) ([]byte, error) {
	response, err := http.Post(service.BaseURL+path, "application/json", body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	output, err := ioutil.ReadAll(response.Body)
	return output, err
}

func (service *Service) GetApp(path string) (*App, error) {
	jsonBlob, err := service.HttpGet("/v2/apps" + path + "?embed=apps.tasks")
	if err != nil {
		return nil, err
	}

	type jsonResponse struct {
		App *App
	}

	var v jsonResponse
	err = json.Unmarshal(jsonBlob, &v)

	return v.App, err
}

func (service *Service) GetApps() ([]*App, error) {
	jsonBlob, err := service.HttpGet("/v2/apps?embed=apps.tasks")
	if err != nil {
		return nil, fmt.Errorf("Failed to HTTP GET. %v", err)
	}

	type jsonResponse struct {
		Apps []*App
	}

	var v jsonResponse
	err = json.Unmarshal(jsonBlob, &v)
	if err != nil {
		return v.Apps, fmt.Errorf("Could not unmarshal JSON response. %v", err)
	} else {
		return v.Apps, nil
	}
}
