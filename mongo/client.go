package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	cfg *config
	cli *mongo.Client
	db  *mongo.Database
}

func NewClient(opts ...Option) (*Client, func()) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}

	ctx := context.Background()

	option := cfg.getOption()

	// connect
	cli := &Client{cfg: cfg}
	var err error
	cli.cli, err = mongo.Connect(ctx, option)
	if err != nil {
		panic(err)
	}
	// test ping
	if err := cli.cli.Ping(ctx, readpref.Primary()); err != nil {
		panic(errors.Wrap(err, "mongo ping fail"))
	}
	cli.db = cli.cli.Database(cfg.Database)

	return cli, func() {
		_ = cli.cli.Disconnect(context.Background())
	}
}
