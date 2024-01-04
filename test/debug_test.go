package test

import (
	"context"
	"fmt"
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
		Name:  "skyle",
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

func TestDebugUpdateOne(t *testing.T) {
	conn, f := newConn()
	defer f()

	filter := map[string]interface{}{
		"name":  "leslie",
		"value": 22,
	}
	content := map[string]interface{}{
		"age": 101,
	}
	err := conn.Collection(&demo{collName: "demo"}).Debug(os.Stdout).WithCtx(context.TODO()).Filter(filter).UpdateOne(content)
	require.Nil(t, err)
}

func TestDebugUpdateMany(t *testing.T) {
	conn, f := newConn()
	defer f()

	filter := map[string]interface{}{
		"name":  "leslie",
		"value": 100,
	}
	content := map[string]interface{}{
		"age": 101,
	}

	err := conn.Collection(&demo{collName: "demo"}).Debug(os.Stdout).WithCtx(context.TODO()).Filter(filter).Update(content)
	require.Nil(t, err)
}

func TestDebugDeleteOne(t *testing.T) {
	conn, f := newConn()
	defer f()

	err := conn.Collection(&demo{collName: "demo"}).Debug(os.Stdout).WithCtx(context.TODO()).Where("name", "leslie").DeleteOne()
	require.Nil(t, err)
}

func TestDebugMany(t *testing.T) {
	conn, f := newConn()
	defer f()

	err := conn.Collection(&demo{collName: "demo"}).Debug(os.Stdout).WithCtx(context.TODO()).Where("name", "skyle").Delete()
	require.Nil(t, err)
}

func TestDebugCount(t *testing.T) {
	conn, f := newConn()
	defer f()

	cnt, err := conn.Collection(&demo{collName: "demo"}).Debug(os.Stdout).Where("name", "leslie").Count()
	require.Nil(t, err)

	fmt.Println("cnt:", cnt)
}

func TestDebugLimit(t *testing.T) {
	conn, f := newConn()
	defer f()

	var list []*result
	err := conn.Collection(&demo{collName: "demo"}).Debug(os.Stdout).Where("name", "leslie").Limit(2).Find(&list)
	require.Nil(t, err)
	require.Equal(t, 2, len(list))
}
