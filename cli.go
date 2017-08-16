package gg

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os/exec"
	"strings"
)

var Version string = "0.0.1"

const (
	ExitCodeOK = iota
	ExitCodeNG
)

const CLIName = "gg"

const Help = `
usage:
  %s [option] [word word word...]
        search words by the default web browser

option:
`

type CLI struct {
	OutStream io.Writer
	ErrStream io.Writer
	OS        string
}

func (cli *CLI) Run(args []string) (exitCode int) {
	ggFlag := flag.NewFlagSet(CLIName, flag.ContinueOnError)
	ggFlag.SetOutput(cli.ErrStream)
	ggFlag.Usage = func() {
		fmt.Fprintf(cli.ErrStream, Help, CLIName)
		ggFlag.PrintDefaults()
	}

	var optVersion bool
	ggFlag.BoolVar(&optVersion, "version", false, "print version and exit")

	err := ggFlag.Parse(args)
	if err != nil {
		return ExitCodeNG
	}

	if optVersion {
		fmt.Fprintf(cli.OutStream, "%s version %s\n", CLIName, Version)
		return ExitCodeOK
	}

	opener := cli.Opener()
	if opener == "" {
		fmt.Fprintf(cli.ErrStream, "Unsupported OS")
		return ExitCodeNG
	}

	addr := cli.Addr(ggFlag.Args())
	fmt.Fprintf(cli.OutStream, "%s %s\n", opener, addr)
	cmd := exec.Command(opener, addr)
	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "%s\n", err)
		return ExitCodeNG
	}
	return ExitCodeOK
}

func (cli *CLI) Opener() (opener string) {
	switch cli.OS {
	case "linux":
		opener = "xdg-open"
	case "darwin":
		opener = "open"
	default:
		opener = ""
	}
	return opener
}

func (cli *CLI) Addr(words []string) (addr string) {
	params := url.Values{}
	params.Add("q", strings.Join(words, " "))
	addr = "https://www.google.co.jp/search?" + params.Encode()
	return addr
}
