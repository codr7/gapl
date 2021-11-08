package types

import (
	"github.com/codr7/gapl"
)

type Func struct {
	Basic
}

func (self *Func) DumpVal(v gapl.Val) string {
	return v.Data().(*gapl.Func).Dump()
}

