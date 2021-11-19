package main

import (
	"flag"
	"fmt"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/setup"
	"os"
)

func main() {
	var vm gapl.Vm
	vm.NewScope()
	setup.InitVm(&vm)
	vm.NewState()

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Printf("g/>pl %v\n", gapl.VERSION)
		fmt.Println("press Return on empty line to Eval")
		fmt.Println("may the Source be with You\n")
		
		vm.AbcLib.Import(vm.Scope())
		vm.MathLib.Import(vm.Scope())
		vm.Repl(os.Stdin, os.Stdout)
	} else {
		vm.AbcLib.Import(vm.Scope(), "import")

		for _, a := range args {
			pc := vm.Pc()

			if err := vm.Include(a); err != nil {
				fmt.Println(err)
				break
			}

			if err := vm.Eval(pc); err != nil {
				fmt.Println(err)
				break
			}
		}
	}
}
