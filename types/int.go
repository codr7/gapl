package types

import (
	"github.com/codr7/gapl"
	"strconv"
)

type Int struct {
	gapl.BasicType
}

func (self *Int) DumpVal(v interface{}) string {
	return strconv.Itoa(v.(int))
}

