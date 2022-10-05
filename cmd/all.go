/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AndreGKruger/k8generate/internal/generate/all"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Generates all the k8 files (configmap, secret, deployment, service)",
	Long: `Generates all the k8 files (configmap, secret, deployment, service) at once.
The all command generates the following files k8_configmap.yaml, k8_secrets.yaml, k8_deployment.yaml and k8_service.yaml files.
The file is generated in the directory ./kubernetes/{Appenv}/.
The command looks for a .env file in your applications root directory to build out the environment section
For example: kubernetes/production/k8_deployment.yaml ....`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("all called")
		a := all.New(Appname, Appenv, Namespace, Repoendpoint, Reponame, Repoversion)
		err := a.Generate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("all kubernetes file generated at ./kubernetes/%s/", Appenv)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)

	// Local flags
	allCmd.Flags().StringVarP(&Appname, "appname", "a", "", "name of the application")
	allCmd.MarkFlagRequired("appname")
	allCmd.Flags().StringVarP(&Appenv, "env", "e", "", "name of the environment IE:production, staging, development")
	allCmd.MarkFlagRequired("appenv")
	allCmd.Flags().StringVarP(&Namespace, "namespace", "n", "", "namespace of the application, defaults to appname-env IE myapp-production")
	allCmd.Flags().StringVarP(&Repoendpoint, "repo endpoint", "r", "", "endpoint of the repository IE: xyz.dkr.ecr.eu-west-1.amazonaws.com")
	allCmd.MarkFlagRequired("repo endpoint")
	allCmd.Flags().StringVarP(&Reponame, "repo name", "p", "", "name of the repository IE: myrepo/myapp")
	allCmd.MarkFlagRequired("repo name")
	allCmd.Flags().StringVarP(&Repoversion, "repo version", "v", "", "version of the repository IE: 1.0.0")
	allCmd.MarkFlagRequired("repo version")
}
