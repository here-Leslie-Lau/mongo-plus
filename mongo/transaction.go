package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// OfficialSession returns a mongo.Session from the official mongo driver.
func (c *Conn) OfficialSession() (mongo.Session, error) {
	return c.cli.StartSession()
}

// Transaction executes a transaction with the given callback function.
//
// The callback function handles your business logic and should return an error
// if something goes wrong. If the callback function returns an error, the
// transaction will be aborted and rolled back.
// opts are the options for the transaction.
func (c *Conn) Transaction(ctx context.Context, fn func(ctx context.Context) error, opts ...*options.TransactionOptions) error {
	session, err := c.OfficialSession()
	if err != nil {
		return errors.Wrap(err, "failed to start session")
	}
	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		if err := fn(sessCtx); err != nil {
			return nil, errors.Wrap(err, "failed to execute callback meth")
		}
		return nil, nil
	}, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to execute Transaction")
	}

	return nil
}
