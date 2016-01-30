package mecab

import "testing"

func TestNewNode(t *testing.T) {
	line := "すもも\t名詞,一般,*,*,*,*,すもも,スモモ,スモモ"
	expected := Node{
		Surface: "すもも",
		Feature: "名詞,一般,*,*,*,*,すもも,スモモ,スモモ",
		Pos:     "名詞",
		Pos1:    "一般",
		Pos2:    "*",
		Pos3:    "*",
		Cform:   "*",
		Ctype:   "*",
		Base:    "すもも",
		Read:    "スモモ",
		Pron:    "スモモ",
	}

	node := NewNode(line)

	if node.Surface != expected.Surface {
		t.Errorf("expected %s, but got %s", expected.Surface, node.Surface)
	}

	if node.Feature != expected.Feature {
		t.Errorf("expected %s, but got %s", expected.Feature, node.Feature)
	}

	if node.Pos != expected.Pos {
		t.Errorf("expected %s, but got %s", expected.Pos, node.Pos)
	}

	if node.Pos1 != expected.Pos1 {
		t.Errorf("expected %s, but got %s", expected.Pos1, node.Pos1)
	}

	if node.Pos2 != expected.Pos2 {
		t.Errorf("expected %s, but got %s", expected.Pos2, node.Pos2)
	}

	if node.Pos3 != expected.Pos3 {
		t.Errorf("expected %s, but got %s", expected.Pos3, node.Pos3)
	}

	if node.Cform != expected.Cform {
		t.Errorf("expected %s, but got %s", expected.Cform, node.Cform)
	}

	if node.Ctype != expected.Ctype {
		t.Errorf("expected %s, but got %s", expected.Ctype, node.Ctype)
	}

	if node.Base != expected.Base {
		t.Errorf("expected %s, but got %s", expected.Base, node.Base)
	}

	if node.Read != expected.Read {
		t.Errorf("expected %s, but got %s", expected.Read, node.Read)
	}

	if node.Pron != expected.Pron {
		t.Errorf("expected %s, but got %s", expected.Pron, node.Pron)
	}
}
