package test

import (
	"fmt"
	"testing"

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
	fmt.Println("+++", groupStage)

	err := ch.Aggregate(&list, groupStage)
	require.Nil(t, err)
	for _, res := range list {
		fmt.Printf("%+v\n", res)
	}

}
