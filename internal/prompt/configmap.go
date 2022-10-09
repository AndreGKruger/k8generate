package prompt

import (
	"fmt"

	cnfmap "github.com/AndreGKruger/k8generate/internal/generate/configmap"

	"github.com/AlecAivazis/survey/v2"
)

func runConfigmap() {
	// the answers will be written to this struct
	configqsanswers := struct {
		Appname   string
		Appenv    string
		Namespace string
	}{}

	// perform the questions
	err := survey.Ask(configmpqs, &configqsanswers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Appname: %s\n", configqsanswers.Appname)
	fmt.Printf("Appenv: %s\n", configqsanswers.Appenv)
	fmt.Printf("Namespace: %s\n", configqsanswers.Namespace)
	c := cnfmap.New(configqsanswers.Appname, configqsanswers.Appenv, configqsanswers.Namespace)
	err = c.Generate()
	if err != nil {
		fmt.Println(err)
	}
}

// the questions to ask
var configmpqs = []*survey.Question{
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
