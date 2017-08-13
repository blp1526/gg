package gg

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
)

var Version string = "0.0.1"

const (
	ExitCodeOK = iota
	ExitCodeParseFlagError
)

type CLI struct {
	OutStream io.Writer
	ErrStream io.Writer
}

func (cli *CLI) Run(args []string) (exitCode int) {
	flags := flag.NewFlagSet("gg", flag.ContinueOnError)
	flags.SetOutput(cli.ErrStream)

	var optV bool
	var optVersion bool
	flags.BoolVar(&optV, "V", false, "")
	flags.BoolVar(&optVersion, "version", false, "")

	err := flags.Parse(args)
	if err != nil {
		return ExitCodeParseFlagError
	}

	if optV || optVersion {
		fmt.Fprintf(cli.OutStream, "gg version %s\n", Version)
		return ExitCodeOK
	}

	var opener string
	switch runtime.GOOS {
	case "linux":
		opener = "xdg-open"
	case "darwin":
		opener = "open"
	default:
		fmt.Fprintf(cli.OutStream, "Unsupported OS")
		return 1
	}

	params := url.Values{}
	params.Add("q", strings.Join(flags.Args(), " "))
	url := "https://www.google.co.jp/search?" + params.Encode()
	fmt.Printf("%s %s\n", opener, url)
	cmd := exec.Command(opener, url)
	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(cli.OutStream, "%s\n", err)
		return 1
	}
	return ExitCodeOK
}
