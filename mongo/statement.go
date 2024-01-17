package mongo

import (
	"encoding/json"
	"io"
	"os"
	"strconv"
	"unsafe"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Statement struct {
	// the switch to record the statement
	Switch bool
	// the buffer to store the statement
	buf []byte
	// the io.Writer to write the statement
	w io.WriteCloser
}

func newStatement(collName string) *Statement {
	// set the default buffer size to 32
	buf := make([]byte, 0, 32)
	s := &Statement{buf: buf}

	s.batchAppendBuf("mongo-plus:\tdb.", collName, ".")
	return s
}

func (s *Statement) appendBuf(str string) {
	buf := *(*[]byte)(unsafe.Pointer(&str))
	s.buf = append(s.buf, buf...)
}

func (s *Statement) batchAppendBuf(strs ...string) {
	for _, str := range strs {
		s.appendBuf(str)
	}
}

func (s *Statement) debugEnd(ope string, des interface{}, findOpt *options.FindOptions) {
	if !s.Switch {
		// if debug mode is not enabled, return directly
		return
	}

	s.debugJoin(ope, des)

	// specially handle the find operation
	if ope == "find" && findOpt != nil {
		// skip
		if findOpt.Skip != nil {
			s.batchAppendBuf(".skip(", strconv.FormatInt(*findOpt.Skip, 10), ")")
		}
		// limit
		if findOpt.Limit != nil && *findOpt.Limit > 0 {
			s.batchAppendBuf(".limit(", strconv.FormatInt(*findOpt.Limit, 10), ")")
		}
		// sort
		if findOpt.Sort != nil {
			sort := findOpt.Sort.(primitive.D)
			s.appendBuf(".sort({")
			for i, v := range sort {
				s.batchAppendBuf(v.Key, ":", *(*string)(unsafe.Pointer(&v.Value)))
				if i != len(sort)-1 {
					s.appendBuf(", ")
				}
			}
			s.appendBuf("})")
		}
	}

	s.appendBuf("\n")
	// write to io.Writer
	if s.w == nil {
		s.w = os.Stdout
	}
	_, err := s.w.Write(s.buf)
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

	// append to buf
	s.batchAppendBuf(ope, "(", *(*string)(unsafe.Pointer(&byt)), ")")
}

func (s *Statement) batchDebugEnd(ope string, list ...interface{}) {
	if !s.Switch {
		// if debug mode is not enabled, return directly
		return
	}

	s.batchAppendBuf(ope, "(")
	for index, ele := range list {
		byt, err := json.Marshal(ele)
		if err != nil {
			panic(err)
		}
		if index == len(list)-1 {
			s.appendBuf(*(*string)(unsafe.Pointer(&byt)))
		} else {
			s.batchAppendBuf(*(*string)(unsafe.Pointer(&byt)), ", ")
		}
	}
	s.appendBuf(")\n")

	// write to io.Writer
	if s.w == nil {
		s.w = os.Stdout
	}
	_, err := s.w.Write(s.buf)
	if err != nil {
		panic(err)
	}
}
