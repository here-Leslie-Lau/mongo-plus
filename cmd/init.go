/*
Copyright © 2023 here-Leslie-Lau

*/
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/here-Leslie-Lau/mongo-plus/mongo"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "The command is used to initialize the configuration file, which is used for connecting to MongoDB.",
	Long:  "The command is used to initialize the configuration file, which is used for connecting to MongoDB.",
	Run: func(cmd *cobra.Command, args []string) {
		addr, _ := cmd.Flags().GetString("addr")
		username, _ := cmd.Flags().GetString("u")
		passwd, _ := cmd.Flags().GetString("p")
		db, _ := cmd.Flags().GetString("db")

		// generated config file to path
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		cfg := &Cfg{
			Username: username,
			Password: passwd,
			Database: db,
			Addr:     addr,
		}
		byt, _ := json.Marshal(cfg)
		_, err = file.Write(byt)
		if err != nil {
			panic(err)
		}

		fmt.Printf("init %s success...\n", path)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().String("addr", "localhost:27017", "The address of MongoDB")
	initCmd.Flags().String("u", "root", "The username of MongoDB")
	initCmd.Flags().String("p", "root", "The password of MongoDB")
	initCmd.Flags().String("db", "test", "The database of MongoDB")
}

// nolint
var conn *mongo.Conn

type Cfg struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Addr     string `json:"addr"`
}

// nolint load config file and connect to mongodb
func loadConn(cfgStr string) func() {
	content, err := os.ReadFile(cfgStr)
	if err != nil {
		panic(err)
	}

	cfg := new(Cfg)
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
	fmt.Println("load mongodb success...")
	return f
}
