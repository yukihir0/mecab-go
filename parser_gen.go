// generated by mg; DO NOT EDIT

package mecab

// #cgo CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
// #cgo LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++
// #include <mecab.h>
// #include <stdio.h>
import "C"

import (
	"strings"
)

// Parser represents Mecab parser.
type Parser struct {
	model   *C.struct_mecab_model_t
	tagger  *C.struct_mecab_t
	lattice *C.struct_mecab_lattice_t
}

// NewParser returns Mecab parser.
func NewParser(a Args) Parser {
	p := Parser{}
	p.model = C.mecab_model_new2(C.CString(a.Build()))
	p.tagger = C.mecab_model_new_tagger(p.model)
	p.lattice = C.mecab_model_new_lattice(p.model)

	return p
}

// Release release Mecab parser.
func (p *Parser) Release() {
	C.mecab_destroy(p.tagger)
	C.mecab_lattice_destroy(p.lattice)
	C.mecab_model_destroy(p.model)
}

// Parse returns Mecab parse result.
func (p Parser) Parse(sentence string) ([]Node, error) {
	C.mecab_lattice_set_sentence(p.lattice, C.CString(sentence))
	C.mecab_parse_lattice(p.tagger, p.lattice)

	nodes := []Node{}
	lines := strings.Split(C.GoString(C.mecab_lattice_tostr(p.lattice)), "\n")
	for _, line := range lines {
		if strings.Index(line, "EOS") != 0 && len(line) > 1 {
			nodes = append(nodes, NewNode(line))
		}
	}

	return nodes, nil
}
