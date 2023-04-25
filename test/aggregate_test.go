package test

import (
	"fmt"
	"testing"

	"github.com/here-Leslie-Lau/mongo-plus/mongo"
	"github.com/stretchr/testify/require"
)

func TestGetMatchStage(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()

	ch := conn.Collection(&demo{collName: "demo"})
	// 获取match的stage, 条件为name: leslie
	matchStage := ch.GetMatchStage("name", "leslie")

	var list []*result
	err := ch.Aggregate(&list, matchStage)
	require.Nil(t, err)

	for _, res := range list {
		require.Equal(t, "leslie", res.Name)
	}
}

func TestSumAndGroupStage(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()
	ch := conn.Collection(&demo{collName: "demo"})

	var list []*result
	sumStage := ch.GetSumStage("value", "value")
	groupStage := ch.GetGroupStage("name", sumStage)

	err := ch.Aggregate(&list, groupStage)
	require.Nil(t, err)
	for _, res := range list {
		fmt.Printf("%+v\n", res)
	}

}

type testStage struct {
	ID       string `json:"_id" bson:"_id"`
	TotalSum int64  `json:"total_sum" bson:"total_sum"`
	Avg      int64  `json:"avg" bson:"avg"`
}

func TestGroupStage(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()
	ch := conn.Collection(&demo{collName: "demo"})

	var list []*testStage
	avgStage := ch.GetAvgStage("avg", "value")
	totalSumStage := ch.GetSumStage("total_sum", "value")
	groupStage := ch.GetGroupStage("name", totalSumStage, avgStage)

	err := ch.Aggregate(&list, groupStage)
	require.Nil(t, err)
	for _, res := range list {
		fmt.Printf("%+v\n", res)
	}
}

func TestSortStage(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()
	ch := conn.Collection(&demo{collName: "demo"})

	var list []*result
	rules := []mongo.SortRule{
		{Typ: mongo.SortTypeDESC, Field: "value"},
	}

	sortStage := ch.GetSortStage(rules...)
	err := ch.Aggregate(&list, sortStage)
	require.Nil(t, err)

	for _, res := range list {
		fmt.Printf("%+v\n", res)
	}
}

func TestLimitAndSkipStage(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()
	ch := conn.Collection(&demo{collName: "demo"})

	var list []*result

	limitStage := ch.GetLimitStage(2)
	skipStage := ch.GetSkipStage(2)
	err := ch.Aggregate(&list, skipStage, limitStage)
	require.Nil(t, err)
	require.Equal(t, 2, len(list))

	for _, res := range list {
		fmt.Printf("%+v\n", res)
	}
}

func TestUnsetStage(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()
	ch := conn.Collection(&demo{collName: "demo"})

	var list []*result
	matchStage := ch.GetMatchStage("name", "leslie")
	unsetStage := ch.GetUnsetStage("_id", "value")
	err := ch.Aggregate(&list, matchStage, unsetStage)
	require.Nil(t, err)

	for _, res := range list {
		fmt.Printf("%+v\n", res)
	}
}
