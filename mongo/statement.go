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

func (s *Statement) debugEnd(ope string, des interface{}) {
	if !s.Switch {
		// if debug mode is not enabled, return directly
		return
	}

	s.debugJoin(ope, des)
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

func (s *Statement) debugJoin(ope string, des interface{}) {
	if !s.Switch {
		// if debug mode is not enabled, return directly
		return
	}

	byt, err := json.Marshal(des)
	if err != nil {
		panic(err)
	}
	s.Statement += ope + "(" + string(byt) + ")"
}

func (s *Statement) batchDebugEnd(ope string, list ...interface{}) {
	if !s.Switch {
		// if debug mode is not enabled, return directly
		return
	}

	s.Statement += ope + "("
	for index, ele := range list {
		byt, err := json.Marshal(ele)
		if err != nil {
			panic(err)
		}
		if index == len(list)-1 {
			s.Statement += string(byt)
		} else {
			s.Statement += string(byt) + ", "
		}
	}
	s.Statement += ")\n"

	// write to io.Writer
	if s.w == nil {
		s.w = os.Stdout
	}
	_, err := s.w.Write([]byte(s.Statement))
	if err != nil {
		panic(err)
	}
}
