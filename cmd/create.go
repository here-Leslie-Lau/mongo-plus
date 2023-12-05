/*
Copyright © 2023 here-Leslie-Lau

*/
package main

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This command is used to create indexes, etc.",
	Long:  `This command is used to create indexes, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
