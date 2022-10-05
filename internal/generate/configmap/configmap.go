package configmap

import (
	"fmt"
	"os"

	"github.com/AndreGKruger/k8generate/internal/generate"
	e "github.com/AndreGKruger/k8generate/internal/generate/env_vars"
)

func New(Appname string, Appenv string, Namespace string) generate.Generate {
	if Namespace == "" {
		Namespace = Appname + "-" + Appenv
	}
	return &configmapImpl{
		Appname:    Appname,
		Appenv:     Appenv,
		filename:   "k8_configmap.yaml",
		foldername: "./kubernetes/" + Appenv,
		Namespace:  Namespace,
	}
}

type configmapImpl struct {
	Appname    string
	Appenv     string
	Namespace  string
	Envvars    []e.Envvar
	filename   string
	foldername string
}

func (c *configmapImpl) Generate() error {
	envfile, err := os.ReadFile(".env")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No .env file found in project root directory. The configmap will be generated without any environment variables.")
		} else {
			return err
		}
	}
	env := e.New()
	c.Envvars = env.GenerateVarsFromFileBytes(envfile, true)

	return generate.ProcessTemplate(template, c.foldername, c.filename, c)
}

var template = `apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Appname }}-cfg
  namespace: {{ .Namespace }}
  labels:
    app: {{ .Appname }}-cfg
data:
{{range .Envvars}}  {{.Name}}: "{{.Value}}"
{{end}}`
