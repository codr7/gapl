package readers

import (
	"bufio"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/forms"
	"io"
	"strings"
)

func String(in *bufio.Reader, pos *gapl.Pos, vm *gapl.Vm) (gapl.Form, error) {
	fpos := *pos
	var c rune
	
	if c, _, _ = in.ReadRune(); c == '"' {
		pos.Read(c)
	} else {
		in.UnreadRune()
		return nil, nil
	}

	var buf strings.Builder

	for {
		var err error
		if c, _, err = in.ReadRune(); err == io.EOF || c == '"' {
			break
		} else if err != nil {
			return nil, err
		}

		buf.WriteRune(c)
		pos.Read(c)
	}

	if c != '"' {
		return nil, gapl.NewERead(fpos, "Open string")
	}

	pos.Read(c)
	return forms.NewLiteral(fpos, vm.StringType, buf.String()), nil
}
