package main

import (
	"fmt"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/forms"
	"github.com/codr7/gapl/ops"
	"github.com/codr7/gapl/readers"
	"github.com/codr7/gapl/tools"
	"github.com/codr7/gapl/types"
	"os"
)

func main() {
	var vm gapl.Vm
	vm.RegType = new(types.Reg)
	vm.RegType.Init("Reg")
	
	var abcLib gapl.Lib
	abcLib.Init("abc")

	var anyType types.Basic
	anyType.Init("Any")

	var metaType types.Meta
	metaType.Init("Meta", &anyType)
	abcLib.Bind("Any", &metaType, &anyType)
	abcLib.Bind("Meta", &metaType, &metaType)

	vm.BoolType = new(types.Bool)
	vm.BoolType.Init("Bool", &anyType)
	abcLib.Bind("Bool", &metaType, vm.BoolType)
	abcLib.Bind("T", vm.BoolType, true)
	abcLib.Bind("F", vm.BoolType, false)
	
	var funcType types.Func
	funcType.Init("Func", &anyType)
	abcLib.Bind("Func", &metaType, &funcType)
	
	vm.IntType = new(types.Int)
	vm.IntType.Init("Int")
	abcLib.Bind("Int", &metaType, vm.IntType)

	var macroType types.Macro
	macroType.Init("Macro", &anyType)
	abcLib.Bind("Macro", &metaType, &macroType)

	abcLib.Bind("_", &macroType, new(gapl.Macro).Init("_", 0, 
		func(self *gapl.Macro, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
			return in, nil
		}))

	abcLib.Bind("=", &macroType, new(gapl.Macro).Init("=", 2, 
		func(self *gapl.Macro, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
			var err error
			left := in[0]
			in = in[1:]
			var leftVal *gapl.Val
			
			if leftVal = left.Val(vm); leftVal == nil {
				if in, err = left.Emit(in, vm); err != nil {
					return in, err
				}
			}
			
			right := in[0]
			in = in[1:]
			var rightVal *gapl.Val
			
			if rightVal = right.Val(vm); rightVal == nil {
				if in, err = right.Emit(in, vm); err != nil {
					return in, err
				}
			}

			vm.Emit(ops.NewEqual(form, leftVal, rightVal))
			return in, nil
		}))

	abcLib.Bind("bench", &macroType, new(gapl.Macro).Init("bench", 2, 
		func(self *gapl.Macro, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
			var err error
			if in, err = in[0].Emit(in[1:], vm); err != nil {
				return in, err
			}

			op := vm.Emit(ops.NewBench(form, -1)).(*ops.Bench)
			
			if in, err = in[0].Emit(in[1:], vm); err != nil {
				return in, err
			}

			vm.Emit(&ops.STOP)
			op.EndPc = vm.Pc()
			return in, nil
		}))

	abcLib.Bind("call", &macroType, new(gapl.Macro).Init("call", 1, 
		func(self *gapl.Macro, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
			var err error
			
			if in, err = in[0].Emit(in[1:], vm); err != nil {
				return in, err
			}

			vm.Emit(ops.NewCall(form, nil, gapl.CallFlags{CheckArgs: true, CheckRets: true}))
			return in, nil
		}))
	
	abcLib.Bind("func", &macroType, new(gapl.Macro).Init("func", 3, 
		func(self *gapl.Macro, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
			f := in[0]
			var name string

			switch f := f.(type) {
			case *forms.Group:
				name = ""
			case *forms.Id:
				name = f.Name()
				in = in[1:]
			default:
				return in, gapl.NewEEmit(form.Pos(), "Expected identifier or group")
			}

			argsForm, retsForm := in[0], in[1]
			in = in[2:]
			
			argForms := argsForm.(*forms.Group).Members()
			scope := vm.Scope()
			var args gapl.Args
			
			for len(argForms) > 0 {
				idForm, typeForm := argForms[0], argForms[1]
				argForms = argForms[2:]
				_type := scope.Find(typeForm.(*forms.Id).Name())

				if _type == nil {
					return in, gapl.NewEEmit(typeForm.Pos(), "Unknown type: %v", typeForm)
				}

				args = args.Add(idForm.(*forms.Id).Name(), _type.Data().(gapl.Type))
			}
			
			retForms := retsForm.(*forms.Group).Members()
			var rets gapl.Rets
			
			for len(retForms) > 0 {
				idForm := retForms[0]
				retForms = retForms[1:]
				_type := scope.Find(idForm.(*forms.Id).Name())

				if _type == nil {
					return in, gapl.NewEEmit(idForm.Pos(), "Unknown type: %v", idForm)
				}

				rets = rets.Add(_type.Data().(gapl.Type))
			}

			_func := new(gapl.Func).Init(name, args, rets, nil)

			if name == "" {
				vm.Emit(ops.NewPush(form, &funcType, _func))
			} else {
				if v := scope.Find(name); v != nil {
					return in, gapl.NewEEmit(form.Pos(), "Duplicate binding: %v %v", name, v.Dump())
				}

				scope.Bind(name, &funcType, _func)
			}
			
			skip := vm.Emit(ops.NewJump(form, -1)).(*ops.Jump)
			startPc := vm.Pc()
			vm.NewScope()
			
			for i := 0; i < len(args); i++ {
				a := args[len(args)-i-1]
				vm.Emit(ops.NewStore(form, vm.BindReg(a.Name()), a.Type()))
			}
			
			var err error
			
			if in, err = in[0].Emit(in[1:], vm); err != nil {
				return in, err
			}
			
			vm.Emit(&ops.RET)
			skip.Pc = vm.Pc()
			_func.CompileBody(startPc)
			vm.EndScope()
			return in, nil
		}))
		
	abcLib.Bind("if", &macroType, new(gapl.Macro).Init("if", 3, 
		func(self *gapl.Macro, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
			var err error

			if in, err = in[0].Emit(in[1:], vm); err != nil {
				return in, err
			}
			
			op := vm.Emit(ops.NewBranch(form, -1)).(*ops.Branch)
			
			if in, err = in[0].Emit(in[1:], vm); err != nil {
				return in, err
			}
			
			skipRight := vm.Emit(ops.NewJump(form, -1)).(*ops.Jump)
			op.RightPc = vm.Pc()
			
			if in, err = in[0].Emit(in[1:], vm); err != nil {
				return in, err
			}

			skipRight.Pc = vm.Pc()
			return in, nil
		}))

	abcLib.Bind("let", &macroType, new(gapl.Macro).Init("let", 2, 
		func(self *gapl.Macro, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
			key := in[0].(*forms.Id).Name()
			valForm := in[1]
			val := valForm.Val(vm)
			in = in[2:]
			
			if val == nil {
				reg := vm.BindReg(key)
				var err error
				
				if in, err = valForm.Emit(in, vm); err != nil {
					return in, err
				}

				vm.Emit(ops.NewStore(form, reg, nil))
			} else {
				scope := vm.Scope()
				
				if v := scope.Find(key); v != nil {
					return in, gapl.NewEEmit(form.Pos(), "Duplicate binding: %v %v", key, v.Dump())
				}
				
				scope.Bind(key, val.Type(), val.Data())
			}
			
			return in, nil
		}))

	var mathLib gapl.Lib
	mathLib.Init("math")
	
	mathLib.Bind("+", &funcType, new(gapl.Func).Init("+",
		gapl.Args{}.Add("x", vm.IntType).Add("y", vm.IntType),
		gapl.Rets{}.Add(vm.IntType),
		func(self *gapl.Func, flags gapl.CallFlags, retPc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
			stack := vm.Stack()
			y := stack.Pop()
			x := stack.Peek()
			x.Set(x.Type(), x.Data().(int) + y.Data().(int))
			return retPc, nil
		}))

	mathLib.Bind("-", &funcType, new(gapl.Func).Init("-",
		gapl.Args{}.Add("x", vm.IntType).Add("y", vm.IntType),
		gapl.Rets{}.Add(vm.IntType),
		func(self *gapl.Func, flags gapl.CallFlags, retPc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
			stack := vm.Stack()
			y := stack.Pop()
			x := stack.Peek()
			x.Set(x.Type(), x.Data().(int) - y.Data().(int))
			return retPc, nil
		}))

	mathLib.Bind("<", &funcType, new(gapl.Func).Init("<",
		gapl.Args{}.Add("x", vm.IntType).Add("y", vm.IntType),
		gapl.Rets{}.Add(vm.BoolType),
		func(self *gapl.Func, flags gapl.CallFlags, retPc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
			stack := vm.Stack()
			y := stack.Pop()
			x := stack.Peek()
			x.Set(vm.BoolType, x.Data().(int) < y.Data().(int))
			return retPc, nil
		}))

	fmt.Printf("g/>pl %v\n", gapl.VERSION)
	fmt.Println("press Return on empty line to Eval")
	fmt.Println("may the Source be with You\n")

	vm.AddReader(readers.Ws, readers.Int, readers.Group, readers.Id)
	vm.NewScope()
	abcLib.Import(vm.Scope())
	mathLib.Import(vm.Scope())
	vm.NewState()
	
	tools.Repl(os.Stdin, os.Stdout, &vm)
}
