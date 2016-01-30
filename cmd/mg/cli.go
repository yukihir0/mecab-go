package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"text/template"
)

const (
	ExitCodeOK = iota
	ExitCodeParseError
	ExitMecabConfigParseError
	ExitMecabGenerateError
	ExitParserGenerateError
)

type CLI struct {
	outStream io.Writer
	errStream io.Writer
}

func (c *CLI) Run(args []string) int {
	var version bool
	flags := flag.NewFlagSet("mg", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&version, "version", false, "Print version information and quit")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseError
	}

	if version {
		fmt.Fprintf(c.errStream, "mg version %s\n", Version)
		return ExitCodeOK
	}

	config := MecabConfig{}
	if err := config.Parse(); err != nil {
		return ExitMecabConfigParseError
	}

	if err := GenerateMecab(config); err != nil {
		return ExitMecabGenerateError
	}

	if err := GenerateParser(config); err != nil {
		return ExitParserGenerateError
	}

	return ExitCodeOK
}

func GenerateMecab(config MecabConfig) error {
	t, err := template.ParseFiles("cmd/mg/mecab_gen.tpl")
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, config); err != nil {
		return err
	}

	if err := ioutil.WriteFile("mecab_gen.go", buf.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}

func GenerateParser(config MecabConfig) error {
	t, err := template.ParseFiles("cmd/mg/parser_gen.tpl")
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, config); err != nil {
		return err
	}

	if err := ioutil.WriteFile("parser_gen.go", buf.Bytes(), 0644); err != nil {
		return err
	}

	return nil
}
