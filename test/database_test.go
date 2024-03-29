package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/here-Leslie-Lau/mongo-plus/mongo"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestRunCommand(t *testing.T) {
	conn, f := newConn()
	defer f()

	cmd := bson.D{
		{
			Key:   "isMaster",
			Value: 1,
		},
	}

	res := make(map[string]interface{})
	err := conn.RunCommand(context.TODO(), cmd, &res)
	require.Nil(t, err)

	for k, v := range res {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func TestIsMaster(t *testing.T) {
	conn, f := newConn()
	defer f()

	res := make(map[string]interface{})
	err := conn.IsMaster(context.TODO(), &res)
	require.Nil(t, err)

	for k, v := range res {
		fmt.Printf("%s: %v\n", k, v)
	}

}

func TestPing(t *testing.T) {
	conn, f := newConn()
	defer f()

	err := conn.Ping(context.TODO())
	require.Nil(t, err)
}

func TestDbStats(t *testing.T) {
	conn, f := newConn()
	defer f()

	res := make(map[string]interface{})
	err := conn.DbStats(context.TODO(), &res)
	require.Nil(t, err)

	for k, v := range res {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func TestServerStatus(t *testing.T) {
	conn, f := newConn()
	defer f()

	res := make(map[string]interface{})
	err := conn.ServerStatus(context.TODO(), &res)
	require.Nil(t, err)

	for k, v := range res {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func TestCreateIndex(t *testing.T) {
	conn, f := newConn()
	defer f()

	d := &demo{collName: "demo"}
	indexName, err := conn.CreateIndex(context.TODO(), d, mongo.SortRule{Field: "name", Typ: mongo.SortTypeASC})
	require.Nil(t, err)

	fmt.Println("index:", indexName)
}
