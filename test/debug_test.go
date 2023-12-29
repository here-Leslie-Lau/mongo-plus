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
