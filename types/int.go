package types

import (
	"github.com/codr7/gapl"
	"strconv"
)

type Int struct {
	Basic
}

func (self Int) DumpVal(v gapl.Val) string {
	return strconv.Itoa(v.Data().(int))
}

func (self Int) TrueVal(v gapl.Val) bool {
	return v.Data().(int) != 0
}
