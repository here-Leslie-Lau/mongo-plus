package test

import (
	"study/mongo-plus/mongo"
	"testing"
)

func TestNewConn(t *testing.T) {
	opts := []mongo.Option{
		mongo.WithDatabase("local"),
		mongo.WithMaxPoolSize(10),
		mongo.WithUsername("your username"),
		mongo.WithPassword("your password"),
		mongo.WithAddr("localhost:27017"),
	}
	_, f := mongo.NewConn(opts...)
	defer f()
}
