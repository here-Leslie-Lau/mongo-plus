package test

import (
	"context"
	"study/mongo-plus/mongo"
	"testing"

	"github.com/stretchr/testify/require"
)

type demo struct {
	collName string
}

// 实现mongo.Collection接口
func (d *demo) Collection() string {
	return d.collName
}

func TestChainFindOne(t *testing.T) {
	conn, f := newConn()
	defer f()

	var result struct {
		Value int    `json:"value"`
		Name  string `json:"name"`
	}
	err := conn.Collection(&demo{collName: "demo"}).Where("name", "leslie").FindOne(context.Background(), &result)
	require.Nil(t, err)
	require.Equal(t, "leslie", result.Name)
}

func TestChainComparison(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()

	var result struct {
		Value int    `json:"value"`
		Name  string `json:"name"`
	}
	err := conn.Collection(&demo{collName: "demo"}).Comparison("value", mongo.ComparisonLt, 25).FindOne(context.Background(), &result)
	require.Nil(t, err)
	require.Equal(t, "leslie", result.Name)
}

func TestChainInInt64(t *testing.T) {
	conn, f := newConn()
	defer f()

	var result struct {
		Value int    `json:"value"`
		Name  string `json:"name"`
	}
	err := conn.Collection(&demo{collName: "demo"}).InInt64("value", []int64{1, 2, 22}).FindOne(context.Background(), &result)
	require.Nil(t, err)
	require.Equal(t, "leslie", result.Name)
}

func TestChainExists(t *testing.T) {
	conn, f := newConn()
	defer f()

	var result struct {
		Value int    `json:"value"`
		Name  string `json:"name"`
	}
	err := conn.Collection(&demo{collName: "demo"}).Exists("age", false).FindOne(context.Background(), &result)
	require.Nil(t, err)
	require.Equal(t, "leslie", result.Name)
}

func TestChainType(t *testing.T) {
	conn, f := newConn()
	defer f()

	var result struct {
		Value int    `json:"value"`
		Name  string `json:"name"`
	}
	err := conn.Collection(&demo{collName: "demo"}).Type("value", mongo.Int32).FindOne(context.Background(), &result)
	require.Nil(t, err)
	require.Equal(t, "leslie", result.Name)
}

func TestChainCount(t *testing.T) {
	conn, f := newConn()
	defer f()

	cnt, err := conn.Collection(&demo{collName: "demo"}).Where("name", "leslie").Count(context.Background())
	require.Nil(t, err)
	require.Equal(t, int64(1), cnt)
}
