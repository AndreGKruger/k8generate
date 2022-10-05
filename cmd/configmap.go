/*
Copyright © 2022 Andre Kruger
*/
package cmd

import (
	"fmt"
	"strings"

	cnfmap "github.com/AndreGKruger/k8generate/internal/generate/configmap"
	"github.com/spf13/cobra"
)

var configmapCmd = &cobra.Command{
	Use:   "configmap",
	Short: "Generates a k8_configmap.yaml file",
	Long: `The configmap command generates a k8_configmap.yaml file.
The file is generated in the directory ./kubernetes/{Appenv}/k8_configmap.yaml.
The command looks for a .env file in your applications root directory to build out the environment section
For example: kubernetes/production/k8_configmap.yaml .`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configmap called")
		c := cnfmap.New(strings.ToLower(Appname), strings.ToLower(Appenv), strings.ToLower(Namespace))
		err := c.Generate()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("kubernetes configmap file generated at ./kubernetes/%s/k8_configmap.yaml", Appenv)
	},
}

func init() {
	rootCmd.AddCommand(configmapCmd)
}
