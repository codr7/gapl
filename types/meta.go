package types

import (
	"github.com/codr7/gapl"
)

type Meta struct {
	gapl.BasicType
}

func (self *Meta) DumpVal(v interface{}) string {
	return v.(gapl.Type).Name()
}

