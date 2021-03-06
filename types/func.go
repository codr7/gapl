package types

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/forms"
	"github.com/codr7/gapl/ops"
)

type Func struct {
	Basic
}

func (self *Func) EmitVal(v gapl.Val, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
	f := v.Data().(*gapl.Func)

	if f.Name() == "+" && f.Args()[0].Type() == vm.IntType && f.Args()[1].Type() == vm.IntType {
		x, y := in[0], in[1]
		xv, yv := x.Val(vm), y.Val(vm)

		if yv != nil && yv.Type() == vm.IntType && (xv == nil || xv.Type() == vm.IntType) {
			var err error

			if in, err = x.Emit(in, vm); err != nil {
				return in, err
			}	

			vm.Emit(ops.NewInc(form, yv.Data().(int)))
			return in[2:], nil
		}
	}

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

	var flags gapl.CallFlags
	flags.CheckRets = true

	
	for len(in) > 0 {
		f, ok := in[0].(*forms.Id)

		if !ok || f.Name()[0] != '|' {
			break
		}

		switch f.Name()[1:] {
		case "d", "drop":
			flags.Drop = true
		case "t", "tco":
			flags.Tco = true
		case "u", "unsafe":
			flags.Unsafe = true
		default:
			return in, gapl.NewEEmit(f.Pos(), "Invalid call flag: %v", f)
		}

		in = in[1:]
	}
	
	for _, a := range f.Args() {
		if len(in) == 0 {
			return in, gapl.NewEEmit(form.Pos(), "Missing argument: %v %v", f, a.Name())
		}
		
		af := in[0]

		if av := af.Val(vm); av == nil {
			flags.CheckArgs = true
		} else if !gapl.Isa(av.Type(), a.Type()) {
			return in, gapl.NewEEmit(af.Pos(), "Not applicable: %v %v", av.Type(), a.Type())
		}
		
		var err error

		if in, err = af.Emit(in[1:], vm); err != nil {
			return in, err
		}
	}
	
	vm.Emit(ops.NewCall(form, v.Data().(*gapl.Func), flags))
	return in, nil
}

func (self *Func) LiteralVal(v gapl.Val) *gapl.Val {
	return nil
}


