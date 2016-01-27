package mecab

// ParseResult represents MeCab parse result.
type ParseResult struct {
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
