package mongo

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

// OfficialSession returns a mongo.Session from the official mongo driver.
func (c *Conn) OfficialSession() (mongo.Session, error) {
	return c.cli.StartSession()
}

func (c *Conn) Transaction(ctx context.Context, fn func() error) error {
	session, err := c.OfficialSession()
	if err != nil {
		return errors.Wrap(err, "failed to start session")
	}
	defer session.EndSession(ctx)

	return nil
}
