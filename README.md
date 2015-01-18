# mecab-go

MeCab binding for Go.

## Original
- http://qiita.com/rerofumi/items/2bb1e49b20f2175ecaac

## Environment
- os: mac os 10.10(yosemite)
- mecab: install by homebrew

## Install

```
go get github.com/yukihir0/mecab-go
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
