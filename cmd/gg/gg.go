package main

import (
	"os"
	"runtime"

	"github.com/blp1526/gg"
)

func main() {
	cli := &gg.CLI{
		OutStream: os.Stdout,
		ErrStream: os.Stderr,
		OS:        runtime.GOOS,
	}
	exitCode := cli.Run(os.Args[1:])
	os.Exit(exitCode)
}
