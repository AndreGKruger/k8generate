/*
Copyright © 2022 Andre Kruger <andre@hyvemobile.co.za>
*/
package service

import (
	"github.com/AndreGKruger/k8generate/internal/generate"
)

const (
	TEMPLATE_LOCATION = "internal/generate/service/template.txt"
)

func New(Appname string, Appenv string, Namespace string) generate.Generate {
	if Namespace == "" {
		Namespace = Appname + "-" + Appenv
	}
	return &serviceImpl{
		Appname:    Appname,
		Appenv:     Appenv,
		filename:   "k8_service.yaml",
		foldername: "./kubernetes/" + Appenv,
		Namespace:  Namespace,
	}
}

type serviceImpl struct {
	Appname    string
	Appenv     string
	Namespace  string
	filename   string
	foldername string
}

func (s *serviceImpl) Generate() error {
	return generate.ProcessTemplate(TEMPLATE_LOCATION, s.foldername, s.filename, s)
}
