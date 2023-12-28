package mongo

import "encoding/json"

type Statement struct {
	// mongo native statement
	Statement  string
	Attributes json.RawMessage
}

func newStatement(collName string) *Statement {
	return &Statement{Statement: "db." + collName + "()."}
}
