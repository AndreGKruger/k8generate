/*
Copyright © 2022 Andre Kruger
*/
package cmd

import (
	"fmt"

	"github.com/AndreGKruger/k8generate/internal/generate/deployment"
	"github.com/spf13/cobra"
)

var (
	Repoendpoint string
	Reponame     string
	Repoversion  string
	Podport      string
)

var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Generates a k8_deployment.yaml file",
	Long: `The deployment command generates a k8_deployment.yaml file.
The file is generated in the directory ./kubernetes/{Appenv}/k8_deployment.yaml.
The command looks for a .env file in your applications root directory to build out the environment section
For example: kubernetes/production/k8_deployment.yaml .`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployment called")
		d := deployment.New(Appname, Podport, Appenv, Namespace, Repoendpoint, Reponame, Repoversion)
		err := d.Generate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("kubernetes deployment file generated at ./kubernetes/%s/k8_deployment.yaml", Appenv)
	},
}

func init() {
	rootCmd.AddCommand(deploymentCmd)
}
