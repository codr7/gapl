package readers

import (
	"bufio"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/forms"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func Int(in *bufio.Reader, pos *gapl.Pos, vm *gapl.Vm) (gapl.Form, error) {
	fpos := *pos
	var buf strings.Builder

	for {
		c, _, e := in.ReadRune()

		if e == io.EOF {
			break
		}
		
		if e != nil {
			return nil, e
		}

		if c != '-' && !unicode.IsDigit(c) {
			in.UnreadRune()
			break
		}

		buf.WriteRune(c)
		pos.Read(c)
	}

	if buf.Len() == 0 {
		return nil, nil
	}

	s := buf.String()

	if s == "-" {
		return forms.NewId(fpos, s), nil
	}
	
	n, e := strconv.ParseInt(s, 10, 64)

	if e != nil {
		return nil, gapl.NewERead(fpos, "Invalid Int: %v", s)
	}

	return forms.NewLiteral(fpos, vm.IntType, int(n)), nil
}
