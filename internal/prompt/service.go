package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func runService() {
	// the answers will be written to this struct
	servqsanswers := struct {
		Appname     string
		Appenv      string
		Serviceport string
		Namespace   string
	}{}

	// perform the questions
	err := survey.Ask(servqs, &servqsanswers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Appname: %s\n", servqsanswers.Appname)
	fmt.Printf("Appenv: %s\n", servqsanswers.Appenv)
	fmt.Printf("Namespace: %s\n", servqsanswers.Namespace)
	fmt.Printf("Serviceport: %s\n", servqsanswers.Serviceport)
}

// the questions to ask
var servqs = []*survey.Question{
	{
		Name: "Appname",
		Prompt: &survey.Input{
			Message: "What is the name of your application?",
			Default: "testapp",
		},
		Validate:  survey.Required,
		Transform: survey.ToLower,
	},
	{
		Name: "Appenv",
		Prompt: &survey.Select{
			Message: "Choose the environment:",
			Help:    "Your files will be generated in the directory ./kubernetes/{Appenv}/",
			Options: []string{"development", "staging", "production"},
			Default: "staging",
		},
	},
	{
		Name: "Namespace",
		Prompt: &survey.Input{
			Message: "What namespace will this be deployed to on kubernetes?",
			Default: "testapp-staging",
		},
		Validate:  survey.Required,
		Transform: survey.ToLower,
	},
	{
		Name: "Serviceport",
		Prompt: &survey.Input{
			Message: "What port should be used by the service?",
			Help:    "This is the port that will be used in the k8_service.yaml",
			Default: "80",
		},
		Validate: survey.Required,
	},
}
