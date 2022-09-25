/*
Copyright © 2022 Andre Kruger <andre@hyvemobile.co.za>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/AndreGKruger/k8generate/internal/generate/service"
	"github.com/spf13/cobra"
)

var (
	Appname   string
	Appenv    string
	Namespace string
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Generates a k8_service.yaml file",
	Long: `The service command generates a k8_service.yaml file.
The file is generated in the directory ./kubernetes/{Appenv}/k8_service.yaml 
For example: kubernetes/production/k8_service.yaml .`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generating kubernetes service file ...")
		s := service.New(strings.ToLower(Appname), strings.ToLower(Appenv), strings.ToLower(Namespace))
		err := s.Generate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("kubernetes service file generated at ./kubernetes/%s/k8_service.yaml", Appenv)
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)
	// Local flags
	serviceCmd.Flags().StringVarP(&Appname, "appname", "a", "", "name of the application")
	serviceCmd.MarkFlagRequired("appname")
	serviceCmd.Flags().StringVarP(&Appenv, "env", "e", "", "name of the environment IE:production, staging, development")
	serviceCmd.MarkFlagRequired("appenv")
	serviceCmd.Flags().StringVarP(&Namespace, "namespace", "n", "", "namespace of the application, defaults to appname-env IE myapp-production")
}
