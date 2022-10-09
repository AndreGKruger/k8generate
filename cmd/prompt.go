/*
Copyright © 2022 Andre Kruger
*/
package cmd

import (
	"fmt"

	"github.com/AndreGKruger/k8generate/internal/prompt"
	"github.com/spf13/cobra"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Interactive cli",
	Long:  `Interactive cli to help guide you through the process of creating k8 files`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("prompt called")
		prompt.Run()
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
}
