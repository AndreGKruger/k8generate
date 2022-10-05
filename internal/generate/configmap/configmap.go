package configmap

import (
	"fmt"
	"os"
	"strings"

	"github.com/AndreGKruger/k8generate/internal/generate"
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

type envvar struct {
	Name  string
	Value string
}

type configmapImpl struct {
	Appname    string
	Appenv     string
	Namespace  string
	Envvars    []envvar
	filename   string
	foldername string
}

func (c *configmapImpl) Generate() error {
	//Open .env.example file in project root directory of the application
	envfile, err := os.ReadFile(".env")
	if err != nil {
		//check if the error is because the file does not exist
		if os.IsNotExist(err) {
			fmt.Println("No .env file found in project root directory. The configmap will be generated without any environment variables.")
		} else {
			return err
		}
	}
	//Split the file into lines
	lines := strings.Split(string(envfile), "\n")
	//Loop through the lines and create a slice of envvars
	for _, line := range lines {
		if line != "" && !strings.HasPrefix(line, "#") {
			if !generate.ContainsSecrets(line) {
				env := strings.Split(line, "=")
				c.Envvars = append(c.Envvars, envvar{Name: strings.ToLower(env[0]), Value: env[1]})
			}
		}
	}

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
