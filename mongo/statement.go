package mongo

import (
	"encoding/json"
	"io"
	"os"
)

type Statement struct {
	// the switch to record the statement
	Switch bool
	// mongo native statement
	Statement string
	// the io.Writer to write the statement
	w io.WriteCloser
}

func newStatement(collName string) *Statement {
	return &Statement{Statement: "mongo-plus:\tdb." + collName + "."}
}

func (s *Statement) debugEnd(ope string, cond map[string]interface{}) {
	if !s.Switch {
		// if debug mode is not enabled, return directly
		return
	}

	s.debugJoin(ope, cond)
	s.Statement += "\n"
	// write to io.Writer
	if s.w == nil {
		s.w = os.Stdout
	}
	_, err := s.w.Write([]byte(s.Statement))
	if err != nil {
		panic(err)
	}
}

func (s *Statement) debugJoin(ope string, cond map[string]interface{}) {
	if !s.Switch {
		// if debug mode is not enabled, return directly
		return
	}

	byt, _ := json.Marshal(cond)
	s.Statement += ope + "(" + string(byt) + ")"
}
