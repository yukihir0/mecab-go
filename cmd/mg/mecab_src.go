package main

import (
	"bytes"
	"os"
)

const (
	mecabSrcFileName = "mecab_gen.go"
)

type MecabSrc struct{}

func (ms MecabSrc) Generate(buf *bytes.Buffer) error {
	f, err := os.Create(mecabSrcFileName)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(buf.Bytes())
	return nil
}
