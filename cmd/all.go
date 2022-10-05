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
		a := all.New(Appname, Appport, Appenv, Namespace, Repoendpoint, Reponame, Repoversion)
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
	// allCmd.Flags().StringVarP(&Appname, "appname", "a", "testapp", "name of the application")
	// allCmd.Flags().StringVarP(&Appport, "app port", "o", "80", "port of the application")
	// allCmd.Flags().StringVarP(&Appenv, "env", "e", "staging", "name of the environment IE:production, staging, development")
	// allCmd.Flags().StringVarP(&Namespace, "namespace", "n", "testapp-staging", "namespace of the application, defaults to appname-env IE myapp-production")
	// allCmd.Flags().StringVarP(&Repoendpoint, "repo endpoint", "r", "xyz.dkr.ecr.eu-west-1.amazonaws.com", "endpoint of the repository IE: xyz.dkr.ecr.eu-west-1.amazonaws.com")
	// allCmd.Flags().StringVarP(&Reponame, "repo name", "p", "myrepo/myapp", "name of the repository")
	// allCmd.Flags().StringVarP(&Repoversion, "repo version", "v", "1.0.0", "version of the repository")
}
