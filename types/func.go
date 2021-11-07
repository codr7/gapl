package types

import (
	"github.com/codr7/gapl"
)

type Func struct {
	gapl.BasicType
}

func (self *Func) DumpVal(v interface{}) string {
	return v.(*gapl.Func).Dump()
}

