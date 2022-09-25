package secret

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/AndreGKruger/k8generate/internal/generate"
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

type envvar struct {
	Name  string
	Value string
}

type secretImpl struct {
	Appname    string
	Appenv     string
	Namespace  string
	Envvars    []envvar
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
	//Split the file into lines
	lines := strings.Split(string(envfile), "\n")
	//Loop through the lines and create a slice of envvars
	for _, line := range lines {
		if line != "" {
			if generate.ContainsSecrets(line) {
				env := strings.Split(line, "=")
				s.Envvars = append(s.Envvars, envvar{Name: strings.ToLower(env[0]), Value: base64.StdEncoding.EncodeToString([]byte(env[1]))})
			}
		}
	}

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
