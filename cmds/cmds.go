package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/here-Leslie-Lau/mongo-plus/mongo"
)

func main() {
	// load json config
	content, err := os.ReadFile("cmds/conf.json")
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
	_, f := mongo.NewConn(opts...)
	defer f()

	if len(os.Args) < 2 {
		fmt.Println("Usage: mongoplus-cmd <subcommand> [options]")
		os.Exit(1)
	}

	// parse subcommand
	switch os.Args[1] {
	case "create":
		fmt.Println("TODO, create subcommand")
	default:
		fmt.Println("Unknown subcommand", os.Args[1])
	}
}
