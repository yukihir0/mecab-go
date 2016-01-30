package mecab

import (
	"fmt"
	"testing"
)

func TestArgsBuild(t *testing.T) {
	dicDir := "/usr/local/Cellar/mecab/0.996/lib/mecab/dic/mecab-ipadic-neologd"
	expected := fmt.Sprintf("-d %s", dicDir)

	args := NewArgs()
	args.DicDir = dicDir
	a := args.Build()

	if a != expected {
		t.Errorf("expected %s, but got %s", expected, a)
	}
}
