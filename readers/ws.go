package readers

import (
	"bufio"
	"github.com/codr7/gapl"
	"unicode"
)

func Ws(in *bufio.Reader, pos *gapl.Pos, vm *gapl.Vm) (gapl.Form, error) {
	for {
		if c, _, err := in.ReadRune(); err != nil {
			return nil, err
		} else if unicode.IsSpace(c) {
			pos.Read(c)
		} else {
			in.UnreadRune()
			break
		}
	}
	
	return nil, nil
}
