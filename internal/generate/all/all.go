package all

import (
	"github.com/AndreGKruger/k8generate/internal/generate"
	"github.com/AndreGKruger/k8generate/internal/generate/configmap"
	"github.com/AndreGKruger/k8generate/internal/generate/deployment"
	"github.com/AndreGKruger/k8generate/internal/generate/secret"
	"github.com/AndreGKruger/k8generate/internal/generate/service"
)

func New(Appname string, Appport string, Appenv string, Namespace string, Repoendpoint string, Reponame string, Repoversion string) generate.Generate {
	if Namespace == "" {
		Namespace = Appname + "-" + Appenv
	}
	return &allImpl{
		Appname:      Appname,
		Appport:      Appport,
		Appenv:       Appenv,
		Namespace:    Namespace,
		Repoendpoint: Repoendpoint,
		Reponame:     Reponame,
		Repoversion:  Repoversion,
		filename:     "k8_deployment.yaml",
		foldername:   "./kubernetes/" + Appenv,
	}
}

type envvar struct {
	CapsName string
	Name     string
}

type allImpl struct {
	Appname      string
	Appport      string
	Appenv       string
	Namespace    string
	Repoendpoint string
	Reponame     string
	Repoversion  string
	Envvars      []envvar
	Secretvars   []envvar
	filename     string
	foldername   string
}

func (s *allImpl) Generate() error {
	c := configmap.New(s.Appname, s.Appenv, s.Namespace)
	err := c.Generate()
	if err != nil {
		return err
	}
	sc := secret.New(s.Appname, s.Appenv, s.Namespace)
	err = sc.Generate()
	if err != nil {
		return err
	}
	dep := deployment.New(s.Appname, s.Appenv, s.Namespace, s.Repoendpoint, s.Reponame, s.Repoversion)
	err = dep.Generate()
	if err != nil {
		return err
	}
	serv := service.New(s.Appname, s.Appenv, s.Namespace)
	err = serv.Generate()
	if err != nil {
		return err
	}
	return nil
}
