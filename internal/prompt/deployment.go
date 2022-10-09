package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AndreGKruger/k8generate/internal/generate/deployment"
)

func runDeployment() {
	// the answers will be written to this struct
	deplqsanswers := struct {
		Appname      string
		Appenv       string
		Podport      string
		Serviceport  string
		Namespace    string
		Repoendpoint string
		Reponame     string
		Repoversion  string
	}{}

	// perform the questions
	err := survey.Ask(deplqs, &deplqsanswers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Appname: %s\n", deplqsanswers.Appname)
	fmt.Printf("Appenv: %s\n", deplqsanswers.Appenv)
	fmt.Printf("Namespace: %s\n", deplqsanswers.Namespace)
	fmt.Printf("Podport: %s\n", deplqsanswers.Podport)
	fmt.Printf("Serviceport: %s\n", deplqsanswers.Serviceport)
	fmt.Printf("Repo: %s/%s:%s\n", deplqsanswers.Repoendpoint, deplqsanswers.Reponame, deplqsanswers.Repoversion)
	d := deployment.New(deplqsanswers.Appname,
		deplqsanswers.Podport,
		deplqsanswers.Appenv,
		deplqsanswers.Namespace,
		deplqsanswers.Repoendpoint,
		deplqsanswers.Reponame,
		deplqsanswers.Repoversion)
	err = d.Generate()
	if err != nil {
		fmt.Println(err)
	}
}

// the questions to ask
var deplqs = []*survey.Question{
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
		Name: "Podport",
		Prompt: &survey.Input{
			Message: "What port does your application / pod use?",
			Help:    "This is the port that will be used in the k8_deployment.yaml file under ports->containerPort",
			Default: "80",
		},
		Validate: survey.Required,
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
	{
		Name: "Repoendpoint",
		Prompt: &survey.Input{
			Message: "What is endpoint of your repository?",
			Help:    "This is the ECR repository endpoint that will be used in the k8_deployment.yaml file under image->repository",
		},
		Validate:  survey.Required,
		Transform: survey.ToLower,
	},
	{
		Name: "Reponame",
		Prompt: &survey.Input{
			Message: "What is the name of your repo that should be suffixed to the repo endpoint?",
			Help:    "if your repo url is amazonaws.com and you repo name is myrepo/testapp the result would be amazonaws.com/myrepo/testapp",
			Default: "myrepo/testapp",
		},
		Validate:  survey.Required,
		Transform: survey.ToLower,
	},
	{
		Name: "Repoversion",
		Prompt: &survey.Input{
			Message: "What is the version of your repo that should be suffixed to the repo endpoint?",
			Help:    "if your repo url is amazonaws.com and you repo name is myrepo/testapp and your version is 1.0.0 the result would be amazonaws.com/myrepo/testapp:1.0.0",
			Default: "1.0.0",
		},
		Validate: survey.Required,
	},
}
