package docker

import "github.com/vendasta/mscli/pkg/spec"

type CloudbuildYamlTemplateData struct {
	OnJenkins    bool
	Microservice spec.MicroserviceConfig
}

// CloudbuildYamlTemplate template for cloudbuild.yaml
const CloudbuildYamlTemplate = `{{with .Microservice}}steps:
- name: 'gcr.io/cloud-builders/go'
  args: ['install', './server']
  env: ['PROJECT_ROOT={{.GoPackageName}}']
- name: 'gcr.io/cloud-builders/docker'
  args: ['build', '--tag=gcr.io/$PROJECT_ID/{{.Name}}:$_VERSION', '.']
images: ['gcr.io/$PROJECT_ID/{{.Name}}:$_VERSION']
tags:
- {{.Name}}
{{end}}
{{if .OnJenkins}}
options:
  machineType: N1_HIGHCPU_8
{{end}}
`

type DockerfileTemplateData struct {
	Name          string
	GoPackageName string
}

// DockerfileTemplate template for ci
const DockerfileTemplate = `FROM golang:1.9
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
RUN go get -u github.com/golang/lint/golint
RUN go get github.com/derekparker/delve/cmd/dlv
RUN go get github.com/pierrre/gotestcover
ENV GOPATH /go:$GOPATH
ENV PATH /go/bin:$PATH
COPY . /go/src/{{.GoPackageName }}
RUN go build -o /bin/{{.Name}} {{.GoPackageName }}/server/
CMD ["/bin/{{.Name}}"]
`

// CloudbuilderDockerfileTemplate dockerfile template for running cloudbuilider
const CloudbuilderDockerfileTemplate = `FROM alpine:3.5

RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*

COPY gopath/bin/server /server

ENTRYPOINT ["/server"]`

type DockerComposeYamlTemplateData struct {
	Name          string
	GRPCPort      string
	HTTPPort      string
	EndpointsPort string
}

// DockerComposeYamlTemplate template to run the µs locally
const DockerComposeYamlTemplate = `
version: "2"
services:
  {{.Name}}:
    image: "gcr.io/repcore-prod/{{.Name}}:0"
    ports:
      - "{{.GRPCPort}}:11000"
      - "{{.HTTPPort}}:11001"
    environment:
      GOOGLE_APPLICATION_CREDENTIALS: /creds/application_default_credentials.json
      ENVIRONMENT: local
    volumes:
      - ~/.config/gcloud:/creds
  {{.Name}}-endpoints:
    image: "gcr.io/endpoints-release/endpoints-runtime:1"
    ports:
      - "{{.EndpointsPort}}:11003"
    environment:
      GOOGLE_APPLICATION_CREDENTIALS: /creds/application_default_credentials.json
    command: ["-s{{.Name}}-api.vendasta-local.com", "-v[REPLACE THIS (INCLUDING THE []) WITH YOUR ENDPOINTS VERSION]", "-agrpc://{{.Name}}:11000", "-p11003", "-zhealthz", "-k/creds/local-service-account.json"]
    depends_on:
      - "{{.Name}}"
    volumes:
      - ./endpoints/local:/creds
`

type DockerDebugComposeYamlTemplateData struct {
	Name          string
	DelvePort     string
	GoPackageName string
}

const DockerDebugComposeYamlTemplate = `
version: "2"
services:
  {{.Name}}:
    ports:
      - "{{.DelvePort}}:11002"
    expose:
      - "{{.DelvePort}}"
    entrypoint: dlv debug ../{{.GoPackageName}}/server -l 0.0.0.0:{{.DelvePort}} --headless=true --log=true
    privileged: true
    security_opt:
      - seccomp:unconfined
`
