package tools

import (
	"bufio"
	"fmt"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/procs"
	"io"
	"strings"
)

func Repl(in io.Reader, out io.Writer, vm *gapl.Vm) {
	fmt.Fprintf(out, "  ")
	var buf strings.Builder
	ins := bufio.NewScanner(in)
	
	for ins.Scan() {
		if line := ins.Text(); len(line) == 0 && buf.Len() > 0 {
			bin := bufio.NewReader(strings.NewReader(buf.String()))
			pos := gapl.NewPos("repl", 0, 0)
			var forms []gapl.Form
			
			for {
				if f, err := vm.ReadForm(bin, &pos); err == io.EOF {
					break
				} else if err != nil {
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
			
			for len(forms) != 0 {
				f := forms[0]
				var err error
				forms, err = f.Emit(forms[1:], vm)
				
				if err != nil {
					fmt.Fprintln(out, err)
					break
				}
			}

			if len(forms) == 0 && vm.Pc() != pc {
				procs.Fuse(pc, vm)
				vm.Emit(&gapl.STOP)
				
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
