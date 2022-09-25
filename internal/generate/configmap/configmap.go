package configmap

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

const (
	TEMPLATE_LOCATION = "internal/generate/configmap/template.txt"
)

type Configmap interface {
	Generate() error
}

func New(Appname string, Appenv string, Namespace string) Configmap {
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
		return err
	}
	//Split the file into lines
	lines := strings.Split(string(envfile), "\n")
	//Loop through the lines and create a slice of envvars
	for _, line := range lines {
		if line != "" {
			if !containsSecrets(line) {
				env := strings.Split(line, "=")
				c.Envvars = append(c.Envvars, envvar{Name: strings.ToLower(env[0]), Value: env[1]})
			}
		}
	}

	//Setup the template
	tmpl, err := template.New("template.txt").ParseFiles(TEMPLATE_LOCATION)
	if err != nil {
		return err
	}

	//Check if the folder exists, create it if not
	if _, err := os.Stat(c.foldername); os.IsNotExist(err) {
		//Create the folder
		err = os.MkdirAll(c.foldername, 0755)
		if err != nil {
			fmt.Println("Error creating folder: ", err)
			return err
		}
	}

	//Create the file. This also truncates the file if it already exists
	f, err := os.Create(c.foldername + "/" + c.filename)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}
	defer f.Close()

	//Write the template file
	err = tmpl.Execute(f, c)
	if err != nil {
		return err
	}
	return nil
}

var listOfSecrets = []string{"secret", "api", "key", "pass", "user"}

func containsSecrets(envvar string) bool {
	//Check if the envvar contains any of the secrets
	for _, secret := range listOfSecrets {
		fmt.Printf("Checking if %s contains secret %s \n", envvar, secret)
		if strings.Contains(strings.ToLower(envvar), secret) {
			fmt.Println("Found secret")
			return true
		}
	}
	return false
}
