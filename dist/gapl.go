package main

import (
	"fmt"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/ops"
	"github.com/codr7/gapl/types"
	"log"
)

func main() {
	var vm gapl.VM
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
		func(self *gapl.Func, flags gapl.CallFlags, pc gapl.PC, vm *gapl.VM) (gapl.PC, error) {
			stack := &vm.State().Stack
			y := stack.Pop()
			x := stack.Peek()
			x.Set(x.Type(), x.Data().(int) + y.Data().(int))
			return pc, nil
		}))

	vm.NewScope()
	//vm.Import(abcLib)
	//vm.Import(mathLib)
	vm.Emit(ops.NewPush(nil, &intType, 35))
	vm.Emit(ops.NewPush(nil, &intType, 7))
	vm.Emit(ops.NewCall(nil, mathLib.Find("+").Data().(*gapl.Func), gapl.CallFlags{Check: true}))
	vm.Emit(&ops.STOP)
	vm.NewState()

	if err := vm.Eval(0); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%v\n", vm.State().Stack.Dump())
}
