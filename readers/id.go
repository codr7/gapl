package readers

import (
	"bufio"
	"github.com/codr7/gapl"
	"github.com/codr7/gapl/forms"
	"io"
	"strings"
	"unicode"
)

func Id(in *bufio.Reader, pos *gapl.Pos, vm *gapl.Vm) (gapl.Form, error) {
	fpos := *pos
	var buf strings.Builder

	for {
		if c, _, err := in.ReadRune(); err == io.EOF {
			in.UnreadRune()
			break
		} else if err != nil {
			return nil, err
		} else if unicode.IsSpace(c) {
			in.UnreadRune()
			break
		} else {
			pos.Read(c)
			buf.WriteRune(c)
		}
	}

	if buf.Len() == 0 {
		return nil, nil
	}
	
	return forms.NewId(fpos, buf.String()), nil
}
