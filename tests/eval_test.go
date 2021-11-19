package tests

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/setup"
	"testing"
)

func TestEval(t *testing.T){
	var vm gapl.Vm
	vm.NewScope()
	setup.InitVm(&vm)
	vm.AbcLib.Import(vm.Scope())
	vm.MathLib.Import(vm.Scope())
	vm.NewState()
	pc := vm.Pc()

	if err := vm.Include("all.gapl"); err != nil {
		t.Error(err)
	}

	if err := vm.Eval(pc); err != nil {
		t.Error(err)
	}
}
