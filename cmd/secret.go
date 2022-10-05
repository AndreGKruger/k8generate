/*
Copyright © 2022 Andre Kruger
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/AndreGKruger/k8generate/internal/generate/secret"
	"github.com/spf13/cobra"
)

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
}
