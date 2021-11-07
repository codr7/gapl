package main

import (
	"fmt"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/types"
)

func main() {
	var vm gapl.VM
	
	var abcLib gapl.Lib
	abcLib.Init("abc")

	var metaType types.Meta
	metaType.Init("Meta")
	abcLib.BindNew("Meta", &metaType, &metaType)

	var funcType types.Func
	funcType.Init("Func")
	abcLib.BindNew("Func", &metaType, &funcType)
	
	var intType types.Int
	intType.Init("Int")
	abcLib.BindNew("Int", &metaType, &intType)

	var mathLib gapl.Lib
	mathLib.Init("math")

	mathLib.BindNew("+", &funcType, new(gapl.Func).Init("+",
		gapl.Args{}.Add("x", &intType).Add("y", &intType),
		gapl.Rets{}.Add(&intType),
		func(self *gapl.Func, pc gapl.PC, vm *gapl.VM) gapl.PC {
			stack := &vm.State().Stack
			y := stack.Pop()
			x := stack.Peek()
			x.Set(x.Type(), x.Data().(int) + y.Data().(int))
			return pc
		}))

	vm.BeginState()
	vm.PushNew(&intType, 35)
	vm.PushNew(&intType, 7)
	mathLib.Find("+").Data().(*gapl.Func).Eval(-1, &vm)
	fmt.Printf("%v\n", vm.State().Stack.Dump())
}
