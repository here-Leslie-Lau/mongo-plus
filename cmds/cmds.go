package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/here-Leslie-Lau/mongo-plus/mongo"
)

var (
	username string
	pass     string
	databse  string
	addr     string
)

func init() {
	flag.StringVar(&username, "u", "", "mongodb username")
	flag.StringVar(&pass, "p", "", "mongodb password")
	flag.StringVar(&databse, "d", "", "mongodb database")
	flag.StringVar(&addr, "addr", "localhost:27017", "mongodb address and port")
}

func main() {
	flag.Parse()

	if username == "" || pass == "" || databse == "" {
		fmt.Println("u, p, d is required")
		flag.Usage()
		os.Exit(1)
	}

	// connect to mongodb
	opts := []mongo.Option{
		mongo.WithUsername(username),
		mongo.WithPassword(pass),
		mongo.WithDatabase(databse),
		mongo.WithAddr(addr),
		mongo.WithMaxPoolSize(2),
	}
	_, f := mongo.NewConn(opts...)
	defer f()
}
