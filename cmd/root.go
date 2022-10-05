/*
Copyright © 2022 Andre Kruger <andre@hyvemobile.co.za>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k8generate",
	Short: "Generate boilerplate Kubernetes yaml files",
	Long: `k8generate:
k8generate is a simple CLI tool to generate boilerplate Kubernetes yaml files
Generate boilerplate Kubernetes yaml files
Specifically written to generate boilerplate yaml files for Hyve Mobile projects.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.k8generate.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&Appname, "appname", "a", "testapp", "name of the application")
	rootCmd.PersistentFlags().StringVarP(&Appport, "appport", "p", "80", "port of the application")
	rootCmd.PersistentFlags().StringVarP(&Appenv, "env", "e", "staging", "name of the environment IE:production, staging, development")
	rootCmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "testapp-staging", "namespace of the application, defaults to appname-env IE myapp-production")
	rootCmd.PersistentFlags().StringVarP(&Repoendpoint, "repoendpoint", "r", "xyz.dkr.ecr.eu-west-1.amazonaws.com", "repository endpoint of the application, defaults to docker.io/{appname}")
	rootCmd.PersistentFlags().StringVarP(&Reponame, "reponame", "o", "myrepo/testapp", "repository name of the application, defaults to {appname}")
	rootCmd.PersistentFlags().StringVarP(&Repoversion, "repoversion", "v", "1.0.0", "version of the repository")
}
