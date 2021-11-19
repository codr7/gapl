package readers

import (
	"fmt"
	"bufio"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/forms"
)

func Slice(in *bufio.Reader, pos *gapl.Pos, vm *gapl.Vm) (gapl.Form, error) {
	fpos := *pos
	var c rune
	
	if c, _, _ = in.ReadRune(); c == '[' {
		pos.Read(c)
	} else {
		in.UnreadRune()
		return nil, nil
	}

	var items []gapl.Form
	var vals gapl.Slice
	
	for {
		if it, err := vm.ReadForm(in, pos); err != nil {
			fmt.Printf("slice err %v\n", err)
			return nil, err
		} else if it == nil {
			break
		} else {
			items = append(items, it)

			if v := it.Val(vm); v != nil {
				vals = append(vals, *v)
			}
		}
	}

	if c, _, _ = in.ReadRune(); c != ']' {
		return nil, gapl.NewERead(fpos, "Open slice")
	}

	pos.Read(c)

	if len(vals) == len(items) {
		return forms.NewLiteral(fpos, vm.SliceType, vals), nil
	}
	
	return forms.NewSlice(fpos, items...), nil
}
