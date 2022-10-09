package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func runSecret() {
	// the answers will be written to this struct
	secrqsanswers := struct {
		Appname   string
		Appenv    string
		Namespace string
	}{}

	// perform the questions
	err := survey.Ask(secrqs, &secrqsanswers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Appname: %s\n", secrqsanswers.Appname)
	fmt.Printf("Appenv: %s\n", secrqsanswers.Appenv)
	fmt.Printf("Namespace: %s\n", secrqsanswers.Namespace)
}

// the questions to ask
var secrqs = []*survey.Question{
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
}
