/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/AndreGKruger/k8generate/internal/generate/secret"
	"github.com/spf13/cobra"
)

// secretCmd represents the secret command
var secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Generates a k8_secret.yaml file",
	Long: `The secret command generates a k8_secrets.yaml file.
The file is generated in the directory ./kubernetes/{Appenv}/k8_secrets.yaml.
The command looks for a .env file in your applications root directory to build out the environment section
For example: kubernetes/production/k8_secrets.yaml .`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("secret called")
		s := secret.New(strings.ToLower(Appname), strings.ToLower(Appenv), strings.ToLower(Namespace))
		err := s.Generate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("kubernetes secrets file generated at ./kubernetes/%s/k8_secrets.yaml", Appenv)
	},
}

func init() {
	rootCmd.AddCommand(secretCmd)
	// Local flags
	secretCmd.Flags().StringVarP(&Appname, "appname", "a", "", "name of the application")
	secretCmd.MarkFlagRequired("appname")
	secretCmd.Flags().StringVarP(&Appenv, "env", "e", "", "name of the environment IE:production, staging, development")
	secretCmd.MarkFlagRequired("appenv")
	secretCmd.Flags().StringVarP(&Namespace, "namespace", "n", "", "namespace of the application, defaults to appname-env IE myapp-production")
}
