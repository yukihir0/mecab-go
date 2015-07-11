package mecab

import (
	"testing"
)

func TestParse(t *testing.T) {
	input := "すもももももももものうち"
	result, err := Parse(input)

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

	for i, r := range result {
		if r.Surface != surface[i] {
			t.Errorf("expected %s, but got %s", surface[i], r.Surface)
		}

		if r.Feature != feature[i] {
			t.Errorf("expected %s, but got %s", feature[i], r.Feature)
		}

		if r.Pos != pos[i] {
			t.Errorf("expected %s, but got %s", pos[i], r.Pos)
		}

		if r.Pos1 != pos1[i] {
			t.Errorf("expected %s, but got %s", pos1[i], r.Pos1)
		}

		if r.Pos2 != pos2[i] {
			t.Errorf("expected %s, but got %s", pos2[i], r.Pos2)
		}

		if r.Cform != cform[i] {
			t.Errorf("expected %s, but got %s", cform[i], r.Cform)
		}

		if r.Ctype != ctype[i] {
			t.Errorf("expected %s, but got %s", ctype[i], r.Ctype)
		}

		if r.Base != base[i] {
			t.Errorf("expected %s, but got %s", base[i], r.Base)
		}

		if r.Read != read[i] {
			t.Errorf("expected %s, but got %s", read[i], r.Read)
		}

		if r.Pron != pron[i] {
			t.Errorf("expected %s, but got %s", pron[i], r.Pron)
		}
	}
}
