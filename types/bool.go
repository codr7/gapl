package types

import (
	"github.com/codr7/gapl"
)

type Bool struct {
	Basic
}

func (self Bool) DumpVal(v gapl.Val) string {
	if v.Data().(bool) {
		return "T"
	}

	return "F"
}

func (self Bool) TrueVal(v gapl.Val) bool {
	return v.Data().(bool)
}
