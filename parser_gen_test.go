package mecab

import (
	"testing"
)

func TestParserParse(t *testing.T) {
	args := NewArgs()
	parser := NewParser(args)
	defer parser.Release()

	input := "すもももももももものうち"
	nodes, err := parser.Parse(input)

	// expected
	surface := []string{"すもも", "も", "もも", "も", "もも", "の", "うち"}
	feature := []string{
		"名詞,一般,*,*,*,*,すもも,スモモ,スモモ",
		"助詞,係助詞,*,*,*,*,も,モ,モ",
		"名詞,一般,*,*,*,*,もも,モモ,モモ",
		"助詞,係助詞,*,*,*,*,も,モ,モ",
		"名詞,一般,*,*,*,*,もも,モモ,モモ",
		"助詞,連体化,*,*,*,*,の,ノ,ノ",
		"名詞,非自立,副詞可能,*,*,*,うち,ウチ,ウチ",
	}
	pos := []string{"名詞", "助詞", "名詞", "助詞", "名詞", "助詞", "名詞"}
	pos1 := []string{"一般", "係助詞", "一般", "係助詞", "一般", "連体化", "非自立"}
	pos2 := []string{"*", "*", "*", "*", "*", "*", "副詞可能"}
	cform := []string{"*", "*", "*", "*", "*", "*", "*"}
	ctype := []string{"*", "*", "*", "*", "*", "*", "*"}
	base := []string{"すもも", "も", "もも", "も", "もも", "の", "うち"}
	read := []string{"スモモ", "モ", "モモ", "モ", "モモ", "ノ", "ウチ"}
	pron := []string{"スモモ", "モ", "モモ", "モ", "モモ", "ノ", "ウチ"}

	if err != nil {
		t.Fail()
	}

	for i, node := range nodes {
		if node.Surface != surface[i] {
			t.Errorf("expected %s, but got %s", surface[i], node.Surface)
		}

		if node.Feature != feature[i] {
			t.Errorf("expected %s, but got %s", feature[i], node.Feature)
		}

		if node.Pos != pos[i] {
			t.Errorf("expected %s, but got %s", pos[i], node.Pos)
		}

		if node.Pos1 != pos1[i] {
			t.Errorf("expected %s, but got %s", pos1[i], node.Pos1)
		}

		if node.Pos2 != pos2[i] {
			t.Errorf("expected %s, but got %s", pos2[i], node.Pos2)
		}

		if node.Cform != cform[i] {
			t.Errorf("expected %s, but got %s", cform[i], node.Cform)
		}

		if node.Ctype != ctype[i] {
			t.Errorf("expected %s, but got %s", ctype[i], node.Ctype)
		}

		if node.Base != base[i] {
			t.Errorf("expected %s, but got %s", base[i], node.Base)
		}

		if node.Read != read[i] {
			t.Errorf("expected %s, but got %s", read[i], node.Read)
		}

		if node.Pron != pron[i] {
			t.Errorf("expected %s, but got %s", pron[i], node.Pron)
		}
	}
}
func BenchmarkParserParse(b *testing.B) {
	args := NewArgs()
	parser := NewParser(args)
	defer parser.Release()

	input := "すもももももももものうち"
	for n := 0; n < b.N; n++ {
		parser.Parse(input)
	}
}
