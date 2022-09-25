/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	cnfmap "github.com/AndreGKruger/k8generate/internal/generate/configmap"
	"github.com/spf13/cobra"
)

// configmapCmd represents the configmap command
var configmapCmd = &cobra.Command{
	Use:   "configmap",
	Short: "Generates a k8_configmap.yaml file",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configmapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configmapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	configmapCmd.Flags().StringVarP(&Appname, "appname", "a", "", "name of the application")
	configmapCmd.MarkFlagRequired("appname")
	configmapCmd.Flags().StringVarP(&Appenv, "env", "e", "", "name of the environment IE:production, staging, development")
	configmapCmd.MarkFlagRequired("appenv")
	configmapCmd.Flags().StringVarP(&Namespace, "namespace", "n", "", "namespace of the application, defaults to appname-env IE myapp-production")
}
