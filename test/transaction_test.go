package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type errTrancsaction struct{}

func (t *errTrancsaction) Error() string {
	return "test transaction"
}

func TestTransaction(t *testing.T) {
	conn, cancel := newConn()
	defer cancel()

	fn := func(ctx context.Context) error {
		coll := conn.Collection(&demo{collName: "demo"}).WithCtx(ctx)
		if err := coll.InsertOne(&result{
			Value: 100,
			Name:  "leslieOne",
		}); err != nil {
			return err
		}

		// Create a situation to test the atomicity of the transaction.
		return new(errTrancsaction)
	}

	opt := options.Transaction().SetWriteConcern(writeconcern.Majority())

	ctx := context.TODO()
	err := conn.Transaction(ctx, fn, opt)
	if err != nil {
		fmt.Println("err:", err)
	}

	// check the data whether insert success.
	res := new(result)
	err = conn.Collection(&demo{collName: "demo"}).Where("name", "leslieOne").FindOne(res)
	require.Equal(t, mongo.ErrNoDocuments, err)
}
