package main

import (
	"fmt"

	"github.com/yukihir0/mecab-go"
)

func main() {
	input := "すもももももももものうち"

	args := mecab.NewArgs()
	args.DicDir = "/usr/local/Cellar/mecab/0.996/lib/mecab/dic/mecab-ipadic-neologd"
	parser, err := mecab.InitializeParser(args)
	if err != nil {
		panic(err)
	}
	defer parser.Release()

	nodes, err := parser.Parse(input)
	if err != nil {
		panic(err)
	}

	for _, node := range nodes {
		if node.Pos == "名詞" && (node.Pos1 == "一般" || node.Pos1 == "固有名詞") {
			fmt.Println(node.Surface)
		}
	}
}
