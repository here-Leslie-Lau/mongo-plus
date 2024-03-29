/*
Copyright © 2023 here-Leslie-Lau

*/
package main

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

// The generated configuration file path
var path string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mongo-helper init|create|query [flags]",
	Short: "a tool for operating MongoDB databases or collections",
	Long: `"mongo-helper" is a tool for operating MongoDB databases or collections,
	capable of simplifying related operations.
	For more information, you can use "mongo-helper help" to query.
	For more information about subcommands, you can use "mongo-helper [command] --help"`,
}

func init() {
	// generated configuration file path
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	homeDir := user.HomeDir
	path = filepath.Join(homeDir, "mongo-plus.json")

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// config file not exist
		return
	}
	// load mongo conn from config file
	loadConn(path)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
