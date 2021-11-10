package gapl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Repl(vm *VM, in io.Reader, out io.Writer) {
	fmt.Fprintf(out, "gapl %v\n", 1)
	fmt.Fprintf(out, "press Return on empty line to eval\n")
	fmt.Fprintf(out, "may the Source be with you\n\n")
	fmt.Fprintf(out, "  ")
	var buf strings.Builder
	ins := bufio.NewScanner(in)
	
	for ins.Scan() {
		if line := ins.Text(); len(line) == 0 && buf.Len() > 0 {
			pos := NewPos("repl", 0, 0)
			var forms []Form
			bin := bufio.NewReader(strings.NewReader(buf.String()))
			
			for {
				if f, err := ReadForm(bin, &pos, vm); err != nil {
					fmt.Fprintln(out, err)
					forms = nil
					break
				} else if f == nil {
					break
				} else {
					forms = append(forms, f)
				}
			}

			pc := vm.Pc()
			
			for len(forms) > 0 {
				f := forms[0]
				var err error
				forms, err = f.Emit(forms[1:], vm)
				
				if err != nil {
					fmt.Fprintln(out, err)
					break
				}
			}

			if len(forms) == 0 && vm.Pc() != pc {
				if err := vm.Eval(pc); err != nil {
					fmt.Fprintln(out, err)
				}
			}
			
			buf.Reset()
			fmt.Fprintf(out, "%v\n", *vm.Stack())
		} else {
			buf.WriteString(line)
			buf.WriteRune('\n')
		}

		fmt.Fprintf(out, "  ")
	}
}
