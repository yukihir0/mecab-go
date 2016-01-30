package mecab

import "strings"

// Node represents MeCab node.
type Node struct {
	Surface string // 表層形
	Feature string // 形態素
	Pos     string // 品詞
	Pos1    string // 品詞細分類1
	Pos2    string // 品詞細分類2
	Pos3    string // 品詞細分類3
	Cform   string // 活用形
	Ctype   string // 活用型
	Base    string // 原形
	Read    string // 読み
	Pron    string // 発音
}

// NewNode return Mecab node.
func NewNode(line string) Node {
	n := Node{}

	// 表層形\t形態素
	l := strings.Split(line, "\t")
	n.Surface = l[0]
	n.Feature = l[1]

	// 品詞,品詞細分類1,品詞細分類2,品詞細分類3,活用形,活用型,原形,読み,発音
	feature := strings.Split(n.Feature, ",")
	n.Pos = feature[0]
	n.Pos1 = feature[1]
	n.Pos2 = feature[2]
	n.Pos3 = feature[3]
	n.Cform = feature[4]
	n.Ctype = feature[5]
	n.Base = feature[6]
	if len(feature) > 7 {
		n.Read = feature[7]
		n.Pron = feature[8]
	}

	return n
}
