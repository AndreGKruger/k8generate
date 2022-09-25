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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
