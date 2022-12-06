package test

import (
	"context"
	"github.com/stretchr/testify/require"
	"study/mongo-plus/mongo"
	"testing"
)

type demo struct {
	collName string
}

// 实现mongo.Collection接口
func (d *demo) Collection() string {
	return d.collName
}

func TestFindOne(t *testing.T) {
	opts := []mongo.Option{
		mongo.WithDatabase("test"),
		mongo.WithMaxPoolSize(10),
		mongo.WithUsername("your username"),
		mongo.WithPassword("your password"),
		mongo.WithAddr("localhost:27017"),
	}
	conn, f := mongo.NewConn(opts...)
	defer f()

	var result struct {
		Value int    `json:"value"`
		Name  string `json:"name"`
	}
	err := conn.Collection(&demo{collName: "demo"}).FindOne(context.Background(), &result)
	require.Nil(t, err)
	require.Equal(t, "leslie", result.Name)
}
