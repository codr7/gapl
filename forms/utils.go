package forms

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/ops"
)

func Eval(form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, []gapl.Val, error) {
	var err error
	skip := vm.Emit(ops.NewJump(form, -1)).(*ops.Jump)
	pc := vm.Pc()
	
	if in, err = form.Emit(in, vm); err != nil {
		return in, nil, err
	}

	vm.Emit(&gapl.STOP)
	skip.Pc = vm.Pc()
	vm.NewState()
	
	if err = vm.Eval(pc); err != nil {
		return in, nil, err
	}
	
	return in, vm.EndState().Stack().Items(), nil
}
