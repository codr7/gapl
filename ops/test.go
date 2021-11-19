package ops

import (
	"fmt"
	"github.com/codr7/gapl"
)

type Test struct {
	form gapl.Form
	expected gapl.Slice
	EndPc gapl.Pc
}

func NewTest(form gapl.Form, expected gapl.Slice, endPc gapl.Pc) *Test {
	return &Test{form: form, expected: expected, EndPc: endPc}
}

func (self *Test) Eval(pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
	vm.NewState()
	defer vm.EndState()
	
	if err := vm.Eval(pc+1); err != nil {
		return -1, err
	}

	actual := vm.Stack().Items()

	if len(actual) != len(self.expected) {
		return -1, gapl.NewEEval(self.form.Pos(), "Test failed: %v %v", self.expected, actual)
	}

	for i, ev := range self.expected {
		if av := actual[i]; !ev.Equal(av) {
			return -1, gapl.NewEEval(self.form.Pos(), "Test failed: %v %v", self.expected[:i+1], actual[:i+1])
		}
	}
	
	return self.EndPc, nil
}

func (self *Test) String() string {
	return fmt.Sprintf("TEST %v", self.EndPc)
}
