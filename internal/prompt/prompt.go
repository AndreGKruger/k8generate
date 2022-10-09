package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func Run() {
	packageselection := struct {
		Package string
	}{}
	err := survey.Ask(introqs, &packageselection)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch packageselection.Package {
	case "all":
		runAll()
	case "deployment":
		runDeployment()
	case "service":
		runService()
	case "configmap":
		runConfigmap()
	case "secret":
		runSecret()
	}
}

// the questions to ask
var introqs = []*survey.Question{
	{
		Name: "Package",
		Prompt: &survey.Select{
			Message: "What files do you want to generate?",
			Help:    "Your files will be generated in the directory ./kubernetes/{Appenv}/",
			Options: []string{"all", "configmap", "deployment", "secret", "service"},
			Default: "all",
		},
	},
}
