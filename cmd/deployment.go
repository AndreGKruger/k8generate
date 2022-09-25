/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Generates a k8_deployment.yaml file",
	Long: `The deployment command generates a k8_deployment.yaml file.
The file is generated in the directory ./kubernetes/{Appenv}/k8_deployment.yaml.
The command looks for a .env.example file in your applications root directory to build out the environment section
For example: kubernetes/production/k8_deployment.yaml .`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployment called")
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
}
