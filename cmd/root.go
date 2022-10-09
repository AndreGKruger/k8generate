/*
Copyright © 2022 Andre Kruger
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k8generate",
	Short: "Generate boilerplate Kubernetes yaml files",
	Long: `k8generate:
k8generate is a simple CLI tool to generate boilerplate Kubernetes yaml files
Generate boilerplate Kubernetes yaml files
Specifically written to generate boilerplate yaml files for internal projects.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&Appname, "appname", "a", "testapp", "name of the application")
	rootCmd.PersistentFlags().StringVarP(&Serviceport, "Serviceport", "s", "80", "port that the kubernetes service exposes")
	rootCmd.PersistentFlags().StringVarP(&Podport, "Podport", "p", "80", "port that the kubernetes pod uses in the deployment file")
	rootCmd.PersistentFlags().StringVarP(&Appenv, "env", "e", "staging", "name of the environment IE:production, staging, development")
	rootCmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "testapp-staging", "namespace of the application, defaults to appname-env IE myapp-production")
	rootCmd.PersistentFlags().StringVarP(&Repoendpoint, "repoendpoint", "r", "xyz.dkr.ecr.eu-west-1.amazonaws.com", "repository endpoint of the application, defaults to docker.io/{appname}")
	rootCmd.PersistentFlags().StringVarP(&Reponame, "reponame", "o", "myrepo/testapp", "repository name of the application, defaults to {appname}")
	rootCmd.PersistentFlags().StringVarP(&Repoversion, "repoversion", "v", "1.0.0", "version of the repository")
}
