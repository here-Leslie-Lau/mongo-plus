/*
Copyright © 2023 here-Leslie-Lau

*/
package main

import (
	"encoding/json"
	"os"

	"github.com/here-Leslie-Lau/mongo-plus/mongo"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mongo-cli create|query [flags]",
	Short: "a tool for operating MongoDB databases or collections",
	Long: `"mongo-cli" is a tool for operating MongoDB databases or collections,
	capable of simplifying related operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfgStr, err := cmd.Flags().GetString("cfg")
		if err != nil {
			_ = cmd.Usage()
			os.Exit(1)
		}

		// init mongo connection
		cancel := loadConn(cfgStr)
		defer cancel()
	},
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("cfg", "conf.json", "config file path")
}

// nolint
var conn *mongo.Conn

// load config file and connect to mongodb
func loadConn(cfgStr string) func() {
	content, err := os.ReadFile(cfgStr)
	if err != nil {
		panic(err)
	}
	var cfg struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
		Addr     string `json:"addr"`
	}
	if err := json.Unmarshal(content, &cfg); err != nil {
		panic(err)
	}

	// connect to mongodb
	opts := []mongo.Option{
		mongo.WithUsername(cfg.Username),
		mongo.WithPassword(cfg.Password),
		mongo.WithDatabase(cfg.Database),
		mongo.WithAddr(cfg.Addr),
		mongo.WithMaxPoolSize(2),
	}

	var f func()
	conn, f = mongo.NewConn(opts...)
	return f
}
