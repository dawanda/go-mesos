node ('mesos') {
  stage 'checkout code'
  checkout scm

  stage 'building libraries'
  sh 'docker run --rm -ti -v $(pwd):/go/src/github.com/dawanda/go-mesos golang:1.6 go build github.com/dawanda/go-mesos/marathon'

  stage 'test'
  sh 'docker run --rm -ti -v $(pwd):/go/src/github.com/dawanda/go-mesos golang:1.6 go build github.com/dawanda/go-mesos/marathon'
}
