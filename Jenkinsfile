node ('mesos') {
  stage 'checkout code'
  checkout scm
  env
  sh 'pwd'
  sh 'ls -lisahF'

  stage 'building libraries'
  sh 'docker run --rm -v $(pwd):/go/src/github.com/dawanda/go-mesos golang:1.6 go build github.com/dawanda/go-mesos/marathon'

  stage 'test'
  sh 'docker run --rm -v $(pwd):/go/src/github.com/dawanda/go-mesos golang:1.6 go build github.com/dawanda/go-mesos/marathon'
}
