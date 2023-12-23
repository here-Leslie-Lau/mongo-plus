package test

import (
	"testing"

	"github.com/here-Leslie-Lau/mongo-plus/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestNewConn(t *testing.T) {
	_, f := newConn()
	defer f()
}

func newConn() (*mongo.Conn, func()) {
	opts := []mongo.Option{
		mongo.WithDatabase("test"),
		mongo.WithMaxPoolSize(10),
		mongo.WithMinPoolSize(1),
		mongo.WithUsername("your username"),
		mongo.WithPassword("your password"),
		mongo.WithAddr("localhost:27017"),
	}
	return mongo.NewConn(opts...)
}

func newConnWithMonitor(monitor interface{}) (*mongo.Conn, func()) {
	opts := []mongo.Option{
		mongo.WithDatabase("test"),
		mongo.WithMaxPoolSize(10),
		mongo.WithUsername("your username"),
		mongo.WithPassword("your password"),
		mongo.WithAddr("localhost:27017"),
		mongo.WithMonitor(monitor),
	}
	return mongo.NewConn(opts...)
}

func newConnWithLog(opt *options.LoggerOptions) (*mongo.Conn, func()) {
	opts := []mongo.Option{
		mongo.WithDatabase("test"),
		mongo.WithMaxPoolSize(10),
		mongo.WithUsername("your username"),
		mongo.WithPassword("your password"),
		mongo.WithAddr("localhost:27017"),
		mongo.WithLoggerOptions(opt),
	}
	return mongo.NewConn(opts...)
}
