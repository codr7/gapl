package main

import (
	"fmt"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/readers"
	"github.com/codr7/gapl/types"
	"os"
)

func main() {
	var vm gapl.Vm
	vm.RegType = new(types.Reg)
	
	var abcLib gapl.Lib
	abcLib.Init("abc")

	var anyType types.Basic
	anyType.Init("Any")

	var metaType types.Meta
	metaType.Init("Meta", &anyType)
	abcLib.Bind("Any", &metaType, &anyType)
	abcLib.Bind("Meta", &metaType, &metaType)

	var funcType types.Func
	funcType.Init("Func", &anyType)
	abcLib.Bind("Func", &metaType, &funcType)
	
	var intType types.Int
	intType.Init("Int", &anyType)
	abcLib.Bind("Int", &metaType, &intType)

	var mathLib gapl.Lib
	mathLib.Init("math")

	mathLib.Bind("+", &funcType, new(gapl.Func).Init("+",
		gapl.Args{}.Add("x", &intType).Add("y", &intType),
		gapl.Rets{}.Add(&intType),
		func(self *gapl.Func, flags gapl.CallFlags, pc gapl.Pc, vm *gapl.Vm) (gapl.Pc, error) {
			stack := &vm.State().Stack
			y := stack.Pop()
			x := stack.Peek()
			x.Set(x.Type(), x.Data().(int) + y.Data().(int))
			return pc, nil
		}))

	vm.NewScope()
	vm.NewState()

	fmt.Printf("gapl %v\n", gapl.VERSION)
	fmt.Println("press Return on empty line to Eval")
	fmt.Println("may the Source be with You\n")

	vm.AddReader(readers.Ws, readers.Id)
	gapl.Repl(os.Stdin, os.Stdout, &vm)
}
