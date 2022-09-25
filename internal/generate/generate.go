package generate

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Generate interface {
	Generate() error
}

var listOfSecrets = []string{"secret", "api", "key", "pass", "user", "token"}

func ContainsSecrets(envvar string) bool {
	//Check if the envvar contains any of the secrets
	for _, secret := range listOfSecrets {
		if strings.Contains(strings.ToLower(envvar), secret) {
			return true
		}
	}
	return false
}

func ProcessTemplate(templateLocation string, foldername string, filename string, data interface{}) error {
	//Setup the template
	tmpl, err := template.New("template.txt").ParseFiles(templateLocation)
	if err != nil {
		return err
	}

	//Check if the folder exists, create it if not
	if _, err := os.Stat(foldername); os.IsNotExist(err) {
		//Create the folder
		err = os.MkdirAll(foldername, 0755)
		if err != nil {
			fmt.Println("Error creating folder: ", err)
			return err
		}
	}

	//Create the file. This also truncates the file if it already exists
	f, err := os.Create(foldername + "/" + filename)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}
	defer f.Close()

	//Write the template file
	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}
	return nil
}
