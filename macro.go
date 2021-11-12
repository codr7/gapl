package gapl

import (
	"fmt"
)

type MacroBody = func(self *Macro, form Form, in []Form, vm *Vm) ([]Form, error)

type Macro struct {
	name string
	argCount int
	body MacroBody
}

func (self *Macro) Init(name string, argCount int, body MacroBody) *Macro {
	self.name = name
	self.argCount = argCount
	self.body = body
	return self
}

func (self *Macro) Name() string {
	return self.name
}

func (self *Macro) String() string {
	return fmt.Sprintf("Macro(%v %v)", self.name, self.argCount)
}

func (self *Macro) Emit(form Form, in []Form, vm *Vm) ([]Form, error) {
	if len(in) < self.argCount {
		return in, NewEEmit(form.Pos(), "Not enough arguments: %v %v/%v", self.name, len(in), self.argCount)
	}
	
	return self.body(self, form, in, vm)
}
