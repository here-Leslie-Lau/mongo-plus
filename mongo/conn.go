package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Conn struct {
	cfg *config
	cli *mongo.Client
	db  *mongo.Database
}

func NewConn(opts ...Option) (*Conn, func()) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}

	ctx := context.Background()

	option := cfg.getOption()

	// connect
	c := &Conn{cfg: cfg}
	var err error
	c.cli, err = mongo.Connect(ctx, option)
	if err != nil {
		panic(err)
	}
	// test ping
	if err := c.cli.Ping(ctx, readpref.Primary()); err != nil {
		panic(errors.Wrap(err, "mongo ping fail"))
	}
	c.db = c.cli.Database(cfg.Database)

	return c, func() {
		_ = c.cli.Disconnect(context.Background())
	}
}

// Collection 获取操作集合的对象
func (c *Conn) Collection(i Collection) *Chain {
	return &Chain{coll: c.db.Collection(i.Collection())}
}

// GetDB 获取go driver的database对象
func (c *Conn) GetDB() *mongo.Database {
	return c.db
}
