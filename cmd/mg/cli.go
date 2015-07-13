package main

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK = iota
	ExitCodeParseError
	ExitMecabConfigParseError
	ExitMecabTemplateParseError
	ExitMecabSrcGenerateError
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

	template := MecabTemplate{}
	buf, err := template.Parse(config)
	if err != nil {
		return ExitMecabTemplateParseError
	}

	src := MecabSrc{}
	if err := src.Generate(buf); err != nil {
		return ExitMecabSrcGenerateError
	}

	return ExitCodeOK
}
