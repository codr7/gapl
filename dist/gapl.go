package main

import (
	"fmt"
	"github.com/codr7/gapl"
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

	var boolType types.Bool
	boolType.Init("Bool", &anyType)
	abcLib.Bind("Bool", &metaType, &boolType)
	abcLib.Bind("T", &boolType, true)
	abcLib.Bind("F", &boolType, false)
	
	var funcType types.Func
	funcType.Init("Func", &anyType)
	abcLib.Bind("Func", &metaType, &funcType)
	
	vm.IntType = new(types.Int)
	vm.IntType.Init("Int")
	abcLib.Bind("Int", &metaType, vm.IntType)

	var macroType types.Macro
	macroType.Init("Macro", &anyType)
	abcLib.Bind("Macro", &metaType, &macroType)

	var mathLib gapl.Lib
	mathLib.Init("math")

	mathLib.Bind("if", &macroType, new(gapl.Macro).Init("if", 3, 
		func(self *gapl.Macro, form gapl.Form, in []gapl.Form, vm *gapl.Vm) ([]gapl.Form, error) {
			cond, left, right := in[0], in[1], in[2]
			var err error

			if in, err = cond.Emit(in[3:], vm); err != nil {
				return in, err
			}
			
			op := vm.Emit(ops.NewBranch(form, -1)).(*ops.Branch)
			
			if in, err = left.Emit(in, vm); err != nil {
				return in, err
			}
			
			skipRight := vm.Emit(ops.NewJump(form, -1)).(*ops.Jump)
			op.RightPc = vm.Pc()
			
			if in, err = right.Emit(in, vm); err != nil {
				return in, err
			}

			skipRight.Pc = vm.Pc()
			return in, nil
		}))
		
	mathLib.Bind("+", &funcType, new(gapl.Func).Init("+",
		gapl.Args{}.Add("x", vm.IntType).Add("y", vm.IntType),
		gapl.Rets{}.Add(vm.IntType),
		func(self *gapl.Func, flags gapl.CallFlags, pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
			stack := vm.Stack()
			y := stack.Pop()
			x := stack.Peek()
			x.Set(x.Type(), x.Data().(int) + y.Data().(int))
			return pc, nil
		}))

	mathLib.Bind("-", &funcType, new(gapl.Func).Init("-",
		gapl.Args{}.Add("x", vm.IntType).Add("y", vm.IntType),
		gapl.Rets{}.Add(vm.IntType),
		func(self *gapl.Func, flags gapl.CallFlags, pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
			stack := vm.Stack()
			y := stack.Pop()
			x := stack.Peek()
			x.Set(x.Type(), x.Data().(int) - y.Data().(int))
			return pc, nil
		}))

	mathLib.Bind("<", &funcType, new(gapl.Func).Init("<",
		gapl.Args{}.Add("x", vm.IntType).Add("y", vm.IntType),
		gapl.Rets{}.Add(&boolType),
		func(self *gapl.Func, flags gapl.CallFlags, pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
			stack := vm.Stack()
			y := stack.Pop()
			x := stack.Peek()
			x.Set(&boolType, x.Data().(int) < y.Data().(int))
			return pc, nil
		}))

	fmt.Printf("gapl %v\n", gapl.VERSION)
	fmt.Println("press Return on empty line to Eval")
	fmt.Println("may the Source be with You\n")

	vm.AddReader(readers.Ws, readers.Int, readers.Group, readers.Id)
	vm.NewScope()
	abcLib.Import(vm.Scope())
	mathLib.Import(vm.Scope())
	vm.NewState()
	
	tools.Repl(os.Stdin, os.Stdout, &vm)
}
