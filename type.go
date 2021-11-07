package gapl

import (
	"fmt"
)

type Type interface {
	Name() string

	DumpVal(v interface{}) string
}

type BasicType struct {
	name string
}

func (self *BasicType) Init(name string) *BasicType {
	self.name = name
	return self
}

func (self *BasicType) Name() string {
	return self.name
}

func (self *BasicType) DumpVal(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
