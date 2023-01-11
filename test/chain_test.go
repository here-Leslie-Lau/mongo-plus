package test

import (
	"context"
	"testing"

	"github.com/here-Leslie-Lau/mongo-plus/mongo"

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
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Where("name", "leslie").FindOne(&result)
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
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Comparison("value", mongo.ComparisonLt, 25).FindOne(&result)
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
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).InInt64("value", []int64{1, 2, 22}).FindOne(&result)
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
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Exists("age", false).FindOne(&result)
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
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Type("value", mongo.Int32).FindOne(&result)
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

func TestChainLimit(t *testing.T) {
	conn, f := newConn()
	defer f()

	var list []interface{}
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Limit(2).Find(&list)
	require.Nil(t, err)
	require.Equal(t, 2, len(list))
}

type result struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}

func TestChainSort(t *testing.T) {
	conn, f := newConn()
	defer f()

	var list []result
	// 根据value值升序排
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Sort(mongo.SortRule{Typ: mongo.SortTypeASC, Field: "value"}).Find(&list)
	require.Nil(t, err)
	for i := 0; i < len(list)-1; i++ {
		if list[i].Value > list[i+1].Value {
			t.Fail()
		}
	}
}

func TestChainSkip(t *testing.T) {
	conn, f := newConn()
	defer f()

	var list []result
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Skip(2).Find(&list)
	require.Nil(t, err)
	require.Equal(t, 1, len(list))
}

func TestChainPaginate(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()

	f := &mongo.PageFilter{
		PageNum:  1,
		PageSize: 2,
	}

	var list []*result
	err := conn.Collection(&demo{collName: "demo"}).Paginate(context.Background(), f, &list)
	require.Nil(t, err)
	require.Equal(t, 2, len(list))
	require.Equal(t, "leslie", list[0].Name)
}

func TestChainInsert(t *testing.T) {
	conn, f := newConn()
	defer f()

	doc := &result{
		Value: 100,
		Name:  "leslie",
	}
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).InsertOne(doc)
	require.Nil(t, err)
}
