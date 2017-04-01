package report

import (
	"bufio"
	"fmt"
	"github.com/andrewchambers/cc/cpp"
	"os"
)

func ReportError(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintln(os.Stderr, "")
	errLoc, ok := err.(cpp.ErrorLoc)
	if !ok {
		return
	}
	pos := errLoc.Pos
	f, err := os.Open(pos.File)
	if err != nil {
		return
	}
	b := bufio.NewReader(f)
	lineno := 1
	for {
		done := false
		line, err := b.ReadString('\n')
		if err != nil {
			done = true
		}
		if lineno == pos.Line /* || lineno == pos.Line - 1 || lineno == pos.Line + 1 */ {
			for _, v := range line {
				switch v {
				case '\t':
					fmt.Fprintf(os.Stderr, "    ")
				default:
					fmt.Fprintf(os.Stderr, "%c", v)
				}
			}
		}
		if lineno == pos.Line {
			linelen := 0
			for _, v := range line {
				switch v {
				case '\t':
					linelen += 4
				case '\n':
					// nothing.
				default:
					linelen += 1
				}
			}
			for i := 0; i < linelen; i++ {
				if i+1 == pos.Col {
					fmt.Fprintf(os.Stderr, "%c", '^')
				} else {
					fmt.Fprintf(os.Stderr, "%c", ' ')
				}
			}
			fmt.Fprintln(os.Stderr, "")
		}
		lineno += 1
		if done {
			break
		}
	}
}
