package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestLog(t *testing.T) {
	opt := options.Logger().SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug)
	conn, f := newConnWithLog(opt)
	defer f()

	res := new(result)
	err := conn.Collection(&demo{"demo"}).Where("name", "leslie").FindOne(res)
	require.Nil(t, err)

	require.Equal(t, "leslie", res.Name)
}
