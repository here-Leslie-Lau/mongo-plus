package test

import (
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
