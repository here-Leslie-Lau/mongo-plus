package mongo

import "io"

type Statement struct {
	// the switch to record the statement
	Switch bool
	// mongo native statement
	Statement string
	// the io.Writer to write the statement
	w io.WriteCloser
}

func newStatement(collName string) *Statement {
	return &Statement{Statement: "db." + collName + "."}
}
