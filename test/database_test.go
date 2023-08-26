package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

func TestRunCommand(t *testing.T) {
	conn, f := newConn()
	defer f()

	cmd := bson.D{
		{
			Key:   "isMaster",
			Value: 1,
		},
	}

	res := make(map[string]interface{})
	err := conn.RunCommand(context.TODO(), cmd, &res)
	require.Nil(t, err)

	for k, v := range res {
		fmt.Printf("%s: %v\n", k, v)
	}
}
