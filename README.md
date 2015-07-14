# mecab-go

MeCab binding for Go. [![Build Status](https://travis-ci.org/yukihir0/mecab-go.svg?branch=master)](https://travis-ci.org/yukihir0/mecab-go)

## Original
- http://qiita.com/rerofumi/items/2bb1e49b20f2175ecaac

## Environment
- os: mac os 10.10(yosemite)
- mecab: install by homebrew

## Install

```
go get github.com/yukihir0/mecab-go
cd cmd/mg
go install
```
## Generate mecab_gen.go

mecab_gen.go is a generated file that is setuped cflags and ldflag for cgo.
It is a goal that can support multiple operating systems.

```
go generate
```

## How to use

```
input := "すもももももももものうち"
result, err := mecab.Parse(input)

if err != nil {
	panic(err)
}

for _, r := range result {
	fmt.Println(r.Surface)
}
```

## License

Copyright &copy; 2015 yukihir0
