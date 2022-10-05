package deployment

import (
	"fmt"
	"os"

	"github.com/AndreGKruger/k8generate/internal/generate"
	e "github.com/AndreGKruger/k8generate/internal/generate/env_vars"
)

func New(Appname string, Podport string, Appenv string, Namespace string, Repoendpoint string, Reponame string, Repoversion string) generate.Generate {
	if Namespace == "" {
		Namespace = Appname + "-" + Appenv
	}
	return &deploymentImpl{
		Appname:      Appname,
		Appenv:       Appenv,
		Podport:      Podport,
		Namespace:    Namespace,
		Repoendpoint: Repoendpoint,
		Reponame:     Reponame,
		Repoversion:  Repoversion,
		filename:     "k8_deployment.yaml",
		foldername:   "./kubernetes/" + Appenv,
	}
}

type deploymentImpl struct {
	Appname      string
	Appenv       string
	Podport      string
	Namespace    string
	Repoendpoint string
	Reponame     string
	Repoversion  string
	Envvars      []e.Envvar
	Secretvars   []e.Envvar
	filename     string
	foldername   string
}

func (s *deploymentImpl) Generate() error {
	envfile, err := os.ReadFile(".env")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No .env file found in project root directory. The deployment file will be generated without any environment variables.")
		} else {
			return err
		}
	}
	env := e.New()
	s.Envvars = env.GenerateVarsFromFileBytes(envfile, false)
	s.Secretvars = env.GenerateSecretsFromFileBytes(envfile, false)
	return generate.ProcessTemplate(template, s.foldername, s.filename, s)
}

var template = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Appname }}
  namespace: {{ .Namespace }}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{ .Appname }}
  template:
    metadata:
      labels:
        app: {{ .Appname }}
    spec:
      containers:
        - name: {{ .Appname }}
          image: {{ .Repoendpoint }}/{{ .Reponame }}:{{ .Repoversion }}
          imagePullPolicy: IfNotPresent
          resources:
            requests:
              cpu: "10m"
              memory: "8Mi"
            limits:
              cpu: "50m"
              memory: "32Mi"
          livenessProbe:
            httpGet:
                path: /healthcheck
                port: {{ .Podport }}
                httpHeaders:
                  - name: Accept
                    value: application/json
            initialDelaySeconds: 10
            periodSeconds: 60
          ports:
            - name: http
              containerPort: {{ .Podport }}
              protocol: TCP
          env:
            {{range .Envvars}}- name: {{.Name}}
              valueFrom:
                configMapKeyRef:
                    name: {{ $.Appname }}-cfg
                    key: {{.Value}}
            {{end}}{{range .Secretvars}}- name: {{.Name}}
              valueFrom:
                secretKeyRef:
                    name: {{ $.Appname }}-sk
                    key: {{.Value}}
            {{end}}`
