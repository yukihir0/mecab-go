# mecab-go [![Build Status](https://travis-ci.org/yukihir0/mecab-go.svg?branch=master)](https://travis-ci.org/yukihir0/mecab-go) [![Coverage Status](https://coveralls.io/repos/yukihir0/mecab-go/badge.svg?branch=master&service=github)](https://coveralls.io/github/yukihir0/mecab-go?branch=master)

MeCab binding for Go.

## Original
- http://qiita.com/rerofumi/items/2bb1e49b20f2175ecaac

## Environment
- os: mac osx
- mecab: install by homebrew

## Install

```
go get github.com/yukihir0/mecab-go
```

## Generate *_gen.go

*_gen.go is a generated file that is setuped cflags and ldflag for cgo.  
It is a goal that can support multiple operating systems.  

```
./generate.sh
```

## How to use 1

```
input := "すもももももももものうち"

args := mecab.NewArgs()
args.DicDir = "/usr/local/Cellar/mecab/0.996/lib/mecab/dic/mecab-ipadic-neologd"
mecab.Initialize(args)

nodes, err := mecab.Parse(input)
if err != nil {
	panic(err)
}

for _, node := range nodes {
	fmt.Println(node.Surface)
}
```

## How to use 2

```
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
	fmt.Println(node.Surface)
}
```

## License

Copyright &copy; 2015 yukihir0
