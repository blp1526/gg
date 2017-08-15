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

type CLI struct {
	OutStream io.Writer
	ErrStream io.Writer
	OS        string
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
		return ExitCodeNG
	}

	if optV || optVersion {
		fmt.Fprintf(cli.OutStream, "gg version %s\n", Version)
		return ExitCodeOK
	}

	opener := cli.Opener()
	if opener == "" {
		fmt.Fprintf(cli.OutStream, "Unsupported OS")
		return ExitCodeNG
	}

	params := url.Values{}
	params.Add("q", strings.Join(flags.Args(), " "))
	url := "https://www.google.co.jp/search?" + params.Encode()
	fmt.Printf("%s %s\n", opener, url)
	cmd := exec.Command(opener, url)
	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(cli.OutStream, "%s\n", err)
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
