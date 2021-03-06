package main

import (
	"fmt"

	"github.com/yukihir0/mecab-go"
)

func main() {
	input := "すもももももももものうち"

	args := mecab.NewArgs()
	args.DicDir = "/usr/local/Cellar/mecab/0.996/lib/mecab/dic/mecab-ipadic-neologd"
	mecab.Initialize(args)

	nodes, err := mecab.Parse(input)
	if err != nil {
		panic(err)
	}

	for _, node := range nodes {
		if node.Pos == "名詞" && (node.Pos1 == "一般" || node.Pos1 == "固有名詞") {
			fmt.Println(node.Surface)
		}
	}
}
