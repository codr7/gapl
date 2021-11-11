package types

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/ops"
)

type Func struct {
	Basic
}

func (self *Func) EmitVal(v gapl.Val, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	f := v.Data().(*gapl.Func)

	if f.Name() == "-" && f.Args()[0].Type() == vm.IntType && f.Args()[1].Type() == vm.IntType {
		x, y := in[0], in[1]
		xv, yv := x.Val(vm), y.Val(vm)

		if yv != nil && yv.Type() == vm.IntType && (xv == nil || xv.Type() == vm.IntType) {
			var err error

			if in, err = x.Emit(in, vm); err != nil {
				return in, err
			}	

			vm.Emit(ops.NewDec(form, yv.Data().(int)))
			return in[2:], nil
		}
	}

	litArgs := true
	
	for _, a := range f.Args() {
		if len(in) == 0 {
			return in, gapl.NewEEmit(form.Pos(), "Missing argument: %v %v", f, a.Name())
		}
		
		af := in[0]
		av := af.Val(vm)

		if av == nil {
			litArgs = false
		} else if !gapl.Isa(av.Type(), a.Type()) {
			return in, gapl.NewEEmit(af.Pos(), "Not applicable: %v %v", av.Type(), a.Type())
		}
		
		var err error

		if in, err = af.Emit(in[1:], vm); err != nil {
			return in, err
		}
	}

	vm.Emit(ops.NewCall(form, v.Data().(*gapl.Func), gapl.CallFlags{Check: !litArgs}))
	return in, nil
}

