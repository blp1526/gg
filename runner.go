package gg

import (
	"fmt"
	"os/exec"
)

type Runner interface {
	CombinedOutput(c string, a ...string) ([]byte, error)
}

type RealRunner struct{}

func (rr *RealRunner) CombinedOutput(c string, a ...string) ([]byte, error) {
	return exec.Command(c, a...).CombinedOutput()
}

type MockRunner struct{}

func (mr *MockRunner) CombinedOutput(c string, a ...string) ([]byte, error) {
	if a[0] == "https://www.google.co.jp/search?q=err+case" {
		return nil, fmt.Errorf("err+case")
	}
	return nil, nil
}
