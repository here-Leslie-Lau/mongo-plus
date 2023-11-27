/*
Copyright © 2023 here-Leslie-Lau

*/
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mongo-cli",
	Short: "a tool for operating MongoDB databases or collections",
	Long: `"mongo-cli" is a tool for operating MongoDB databases or collections,
	capable of simplifying related operations.
	For mor information, run "mongo-cli --help" or "mongo-cli -h"`,
	Run: func(cmd *cobra.Command, args []string) {
		cfgStr, _ := cmd.Flags().GetString("cfg")
		fmt.Println(cfgStr)
	},
}

func main() {
	Execute()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("cfg", "conf.json", "config file path")
}
