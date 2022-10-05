package secret

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
	return &secretImpl{
		Appname:    Appname,
		Appenv:     Appenv,
		filename:   "k8_secrets.yaml",
		foldername: "./kubernetes/" + Appenv,
		Namespace:  Namespace,
	}
}

type secretImpl struct {
	Appname    string
	Appenv     string
	Namespace  string
	Envvars    []e.Envvar
	filename   string
	foldername string
}

func (s *secretImpl) Generate() error {
	//Open .env.example file in project root directory of the application
	envfile, err := os.ReadFile(".env")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No .env file found in project root directory. The secrets file will be generated without any environment variables.")
		} else {
			return err
		}
	}
	env := e.New()
	s.Envvars = env.GenerateSecretsFromFileBytes(envfile, true)
	return generate.ProcessTemplate(template, s.foldername, s.filename, s)
}

var template = `apiVersion: v1
kind: Secret
metadata:
  name: {{ .Appname }}-sk
  namespace: {{ .Namespace }}
  labels:
    app: {{ .Appname }}-sk
data:
{{range .Envvars}}  {{.Name}}: {{.Value}}
{{end}}
type: Opaque`
