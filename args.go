package mecab

import (
	"fmt"
)

// Args represents Mecab command line args.
type Args struct {
	DicDir string
}

// NewArgs return Mecab commandline args.
func NewArgs() Args {
	return Args{}
}

// Build return Mecab command line args string.
func (a Args) Build() string {
	s := ""
	if a.DicDir != "" {
		s = s + fmt.Sprintf("-d %s", a.DicDir)
	}

	return s
}
