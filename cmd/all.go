/*
Copyright © 2022 Andre Kruger
*/
package cmd

import (
	"fmt"

	"github.com/AndreGKruger/k8generate/internal/generate/all"
	"github.com/spf13/cobra"
)

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
		a := all.New(Appname, Podport, Serviceport, Appenv, Namespace, Repoendpoint, Reponame, Repoversion)
		err := a.Generate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("all kubernetes file generated at ./kubernetes/%s/", Appenv)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
