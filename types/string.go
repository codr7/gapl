package types

import (
	"fmt"
	"github.com/codr7/gapl"
)

type String struct {
	Basic
}

func (self *String) DumpVal(v gapl.Val) string {
	return fmt.Sprintf("\"%v\"", v.Data().(string))
}

func (self *String) TrueVal(v gapl.Val) bool {
	return len(v.Data().(string)) != 0
}
