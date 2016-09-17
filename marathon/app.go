package marathon

import (
	"fmt"
	"strings"
	"time"
)

type PortMapping struct {
	ContainerPort uint
	HostPort      uint
	ServicePort   uint
	Protocol      string
	Name          string
	Labels        map[string]string
}

type KeyValuePair struct {
	Key   string
	Value *string
}

const (
	NetworkBridged = "BRIDGE"
	NetworkHost    = "HOST"
)

type DockerContainer struct {
	Image          string
	Network        string
	PortMappings   []PortMapping
	Privileged     bool
	Parameters     []KeyValuePair
	ForcePullImage bool
}

type ReadinessCheck struct {
	Name                    string
	Protocol                string
	Path                    string
	PortName                string
	IntervalSeconds         uint
	TimeoutSeconds          uint
	HttpStatusCodesForReady []uint
	PreserveLastResponse    bool
}

type HealthCheck struct {
	Protocol               string
	Path                   string
	PortIndex              int
	GracePeriodSeconds     uint
	IntervalSeconds        uint
	TimeoutSeconds         uint
	MaxConsecutiveFailures uint
	IgnoreHttp1xx          bool
}

type HealthCheckResult struct {
	Alive               bool
	ConsecutiveFailures uint
	FirstSuccess        *time.Time
	LastFailure         *time.Time
	LastSuccess         *time.Time
}

type ContainerVolume struct {
	ContainerPath string
	HostPath      string
	Mode          string
}

type AppContainer struct {
	Type    string
	Volumes []ContainerVolume
	Docker  *DockerContainer
}

type UpgradeStrategy struct {
	MinimumHealthCapacity float64
	MaximumOverCapacity   float64
}

type Task struct {
	Id                 string
	Host               string
	Ports              []uint
	StartedAt          *time.Time
	StagedAt           *time.Time
	Version            time.Time
	SlaveId            string
	State              *TaskStatus
	AppId              string
	HealthCheckResults []HealthCheckResult
}

type FetchInfo struct {
	Uri        string
	Extract    bool
	Executable bool
	Cache      bool
}

type PortDefinition struct {
	Port     uint
	Protocol string
	Name     string
	Labels   map[string]string
}

type App struct {
	service               *Service
	Id                    string
	Cmd                   *string
	Args                  []string
	User                  *string
	Env                   map[string]string
	PortDefinitions       []PortDefinition
	Instances             uint
	Cpus                  float64
	Mem                   uint
	Disk                  uint
	Gpus                  uint
	Executor              string
	Constraints           [][]string
	Uris                  []string
	Fetch                 []FetchInfo
	StoreUrls             []string
	Ports                 []uint
	RequirePorts          bool
	BackoffSeconds        uint
	BackoffFactor         float64
	MaxLaunchDelaySeconds uint
	Container             AppContainer
	HealthChecks          []HealthCheck
	RedinessChecks        []ReadinessCheck
	Dependencies          *[]string
	UpgradeStrategy       UpgradeStrategy
	Labels                map[string]string
	Tasks                 []Task
	AcceptedResourceRoles *[]string
	IpAddress             *IpAddr
	Version               time.Time
	Residency             Residency
	// "Secrets": {},
	TaskKillGracePeriodSeconds *uint
	VersionInfo                VersionInfo
}

type Residency struct {
	TaskLostBehavior string
}

type VersionInfo struct {
	LastScalingAt      time.Time
	LastConfigChangeAt time.Time
}

func (app *App) GetTaskById(taskId string) *Task {
	for i := range app.Tasks {
		if app.Tasks[i].Id == taskId {
			return &app.Tasks[i]
		}
	}

	return nil
}

func (app *App) Scale(instance_count uint) error {
	// TODO
	_, err := app.service.HttpPost(
		"/v2/apps/TODO/scale",
		strings.NewReader(fmt.Sprintf("%v", instance_count)))

	return err
}

func (task *Task) IsAlive() bool {
	for _, healthCheckResult := range task.HealthCheckResults {
		if healthCheckResult.Alive == false {
			return false
		}
	}

	return true
}
