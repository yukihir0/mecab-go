package main

import (
	"os/exec"
	"strings"
)

type MecabConfig struct {
	CFlags  string
	LDFlags string
}

func (mc *MecabConfig) Parse() error {
	cflags, err := exec.Command("mecab-config", "--cflags").Output()
	if err != nil {
		return err
	}

	ldflags, err := exec.Command("mecab-config", "--libs").Output()
	if err != nil {
		return err
	}

	mc.CFlags = strings.TrimSpace(string(cflags))
	mc.LDFlags = strings.TrimSpace(string(ldflags))
	return nil
}
