package readers

import (
	"bufio"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/forms"
)

func Group(in *bufio.Reader, pos *gapl.Pos, vm *gapl.Vm) (gapl.Form, error) {
	fpos := *pos
	var c rune
	
	if c, _, _ = in.ReadRune(); c == '(' {
		pos.Read(c)
	} else {
		in.UnreadRune()
		return nil, nil
	}

	var members []gapl.Form

	for {
		if m, err := vm.ReadForm(in, pos); err != nil {
			return nil, err
		} else if m == nil {
			break
		} else {
			members = append(members, m)
		}
	}

	if c, _, _ = in.ReadRune(); c != ')' {
		return nil, gapl.NewERead(fpos, "Open group")
	}

	pos.Read(c)
	return forms.NewGroup(fpos, members...), nil
}
