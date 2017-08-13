package main

import (
	"os"

	"github.com/blp1526/gg"
)

func main() {
	cli := &gg.CLI{OutStream: os.Stdout, ErrStream: os.Stderr}
	exitCode := cli.Run(os.Args[1:])
	os.Exit(exitCode)
}
