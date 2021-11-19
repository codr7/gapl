package tests

import (
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/setup"
	"testing"
)

func TestEval(t *testing.T){
	var vm gapl.Vm
	scope := vm.NewScope()
	setup.InitVm(&vm)
	vm.AbcLib.Import(scope)
	vm.MathLib.Import(scope)
	vm.NewState()
	pc := vm.Pc()

	if err := vm.Include("all.gapl"); err != nil {
		t.Error(err)
	}

	if err := vm.Eval(pc); err != nil {
		t.Error(err)
	}
}
