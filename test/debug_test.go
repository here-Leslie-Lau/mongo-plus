package test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDebugFindOne(t *testing.T) {
	conn, f := newConn()
	defer f()

	res := new(result)
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Debug(os.Stdout).Where("name", "leslie").FindOne(res)
	require.Nil(t, err)
	require.Equal(t, "leslie", res.Name)
}

func TestDebugFindAll(t *testing.T) {
	conn, f := newConn()
	defer f()

	var list []*result
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Debug(os.Stdout).Where("name", "leslie").Find(&list)
	require.Nil(t, err)
	require.Equal(t, true, len(list) > 0)
	require.Equal(t, "leslie", list[0].Name)
}

func TestDebugInsertOne(t *testing.T) {
	conn, f := newConn()
	defer f()

	doc := &result{
		Value: 101,
		Name:  "leslie",
	}
	err := conn.Collection(&demo{collName: "demo"}).Debug(os.Stdout).WithCtx(context.TODO()).InsertOne(doc)
	require.Nil(t, err)
}

func TestDebugInsertMany(t *testing.T) {
	conn, f := newConn()
	defer f()

	docs := []interface{}{
		&result{
			Value: 103,
			Name:  "leslie",
		},
		&result{
			Value: 104,
			Name:  "leslie",
		},
	}
	err := conn.Collection(&demo{collName: "demo"}).Debug(os.Stdout).WithCtx(context.TODO()).InsertMany(docs)
	require.Nil(t, err)
}
