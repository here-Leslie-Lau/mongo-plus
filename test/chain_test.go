package test

import (
	"context"
	"fmt"
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

	cnt, err := conn.Collection(&demo{collName: "demo"}).Where("name", "leslie").Count()
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
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Paginate(f, &list)
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

func TestProjection(t *testing.T) {
	conn, f := newConn()
	defer f()

	res := &result{}
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Projection("name").FindOne(res)
	require.Nil(t, err)
	require.Equal(t, 0, res.Value)
}

func TestChainLt(t *testing.T) {
	conn, f := newConn()
	defer f()

	var list []*result
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Lt("value", 23).Find(&list)
	require.Nil(t, err)

	for _, user := range list {
		if user.Value >= 23 {
			t.Fail()
		}
	}
}

func TestChainUpsertOne(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()

	content := map[string]interface{}{"value": 100}
	defaultContent := map[string]interface{}{"name": "leslie"}

	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Where("age", 18).UpsertOne(content, defaultContent)
	require.Nil(t, err)
}

func TestChainRegex(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()

	res := &result{}
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Regex("name", "le").FindOne(res)
	require.Nil(t, err)
	require.Equal(t, "leslie", res.Name)
}

func TestChainOr(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()

	var list []*result
	orMap := map[string]interface{}{
		"name":  "leslie",
		"value": 50,
	}
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Or(orMap).Find(&list)
	require.Nil(t, err)
	require.Equal(t, 3, len(list))
}

func TestChainOrs(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()

	var list []*result
	orMap1 := map[string]interface{}{
		"name":  "leslie",
		"value": 22,
	}
	orMap2 := map[string]interface{}{
		"name":  "skyle",
		"value": 78,
	}
	// 查询{"name": "leslie", "value": 22}或{"name": "skyle", "value": 78}的数据
	err := conn.Collection(&demo{collName: "demo"}).WithCtx(context.TODO()).Ors(orMap1, orMap2).Find(&list)
	require.Nil(t, err)
	for _, res := range list {
		fmt.Printf("%+v\n", res)
	}
}
