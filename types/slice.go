package types

import (
	"github.com/codr7/gapl"
)

type Slice struct {
	Basic
}

func (self *Slice) DumpVal(v gapl.Val) string {
	return v.Data().(gapl.Slice).String()
}

func (self *Slice) TrueVal(v gapl.Val) bool {
	return len(v.Data().(gapl.Slice)) != 0
}
