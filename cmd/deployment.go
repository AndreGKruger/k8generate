/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
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
)

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Generates a k8_deployment.yaml file",
	Long: `The deployment command generates a k8_deployment.yaml file.
The file is generated in the directory ./kubernetes/{Appenv}/k8_deployment.yaml.
The command looks for a .env file in your applications root directory to build out the environment section
For example: kubernetes/production/k8_deployment.yaml .`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployment called")
		d := deployment.New(Appname, Appenv, Namespace, Repoendpoint, Reponame, Repoversion)
		err := d.Generate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("kubernetes deployment file generated at ./kubernetes/%s/k8_deployment.yaml", Appenv)
	},
}

func init() {
	rootCmd.AddCommand(deploymentCmd)
	// Local flags
	deploymentCmd.Flags().StringVarP(&Appname, "appname", "a", "", "name of the application")
	deploymentCmd.MarkFlagRequired("appname")
	deploymentCmd.Flags().StringVarP(&Appenv, "env", "e", "", "name of the environment IE:production, staging, development")
	deploymentCmd.MarkFlagRequired("appenv")
	deploymentCmd.Flags().StringVarP(&Namespace, "namespace", "n", "", "namespace of the application, defaults to appname-env IE myapp-production")
	deploymentCmd.Flags().StringVarP(&Repoendpoint, "repoendpoint", "r", "", "endpoint of the repository IE: xyz.dkr.ecr.eu-west-1.amazonaws.com")
	deploymentCmd.MarkFlagRequired("repoendpoint")
	deploymentCmd.Flags().StringVarP(&Reponame, "reponame", "p", "", "name of the repository IE: myrepo/myapp")
	deploymentCmd.MarkFlagRequired("reponame")
	deploymentCmd.Flags().StringVarP(&Repoversion, "repoversion", "v", "", "version of the repository IE: 1.0.0")
	deploymentCmd.MarkFlagRequired("repoversion")

}
