package test

import (
	"study/mongo-plus/mongo"
	"testing"
)

func TestNewConn(t *testing.T) {
	_, f := newConn()
	defer f()
}

func newConn() (*mongo.Conn, func()) {
	opts := []mongo.Option{
		mongo.WithDatabase("test"),
		mongo.WithMaxPoolSize(10),
		mongo.WithUsername("your username"),
		mongo.WithPassword("your password"),
		mongo.WithAddr("localhost:27017"),
	}
	return mongo.NewConn(opts...)
}
