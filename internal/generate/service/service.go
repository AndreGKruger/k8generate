/*
Copyright © 2022 Andre Kruger
*/
package service

import (
	"github.com/AndreGKruger/k8generate/internal/generate"
)

func New(Appname string, Serviceport string, Appenv string, Namespace string) generate.Generate {
	if Namespace == "" {
		Namespace = Appname + "-" + Appenv
	}
	return &serviceImpl{
		Appname:     Appname,
		Serviceport: Serviceport,
		Appenv:      Appenv,
		filename:    "k8_service.yaml",
		foldername:  "./kubernetes/" + Appenv,
		Namespace:   Namespace,
	}
}

type serviceImpl struct {
	Appname     string
	Serviceport string
	Appenv      string
	Namespace   string
	filename    string
	foldername  string
}

func (s *serviceImpl) Generate() error {
	return generate.ProcessTemplate(template, s.foldername, s.filename, s)
}

var template = `apiVersion: v1
kind: Service
metadata:
  name: {{ .Appname }}-svc
  namespace: {{ .Namespace }}
  labels:
    app: {{ .Appname }}-svc
spec:
  type: NodePort
  ports:
    - port: {{ .Serviceport }}
  selector:
    app: {{ .Appname }}`
