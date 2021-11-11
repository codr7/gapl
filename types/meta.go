package types

import (
	"github.com/codr7/gapl"
)

type Meta struct {
	Basic
}

func (self Meta) DumpVal(v gapl.Val) string {
	return v.Data().(gapl.Type).Name()
}
