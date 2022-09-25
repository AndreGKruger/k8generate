package deployment

import (
	"fmt"
	"os"
	"strings"

	"github.com/AndreGKruger/k8generate/internal/generate"
)

const (
	TEMPLATE_LOCATION = "internal/generate/deployment/template.txt"
)

func New(Appname string, Appenv string, Namespace string, Repoendpoint string, Reponame string, Repoversion string) generate.Generate {
	if Namespace == "" {
		Namespace = Appname + "-" + Appenv
	}
	return &deploymentImpl{
		Appname:      Appname,
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

type deploymentImpl struct {
	Appname      string
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

func (s *deploymentImpl) Generate() error {
	//Open .env.example file in project root directory of the application
	envfile, err := os.ReadFile(".env")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No .env file found in project root directory. The deployment file will be generated without any environment variables.")
		} else {
			return err
		}
	}
	//Split the file into lines
	lines := strings.Split(string(envfile), "\n")
	//Loop through the lines and create a slice of envvars
	for _, line := range lines {
		if line != "" {
			if !generate.ContainsSecrets(line) {
				env := strings.Split(line, "=")
				s.Envvars = append(s.Envvars, envvar{Name: strings.ToLower(env[0]), CapsName: strings.ToUpper(env[0])})
			} else {
				env := strings.Split(line, "=")
				s.Secretvars = append(s.Secretvars, envvar{Name: strings.ToLower(env[0]), CapsName: strings.ToUpper(env[0])})
			}
		}
	}
	return generate.ProcessTemplate(TEMPLATE_LOCATION, s.foldername, s.filename, s)
}
