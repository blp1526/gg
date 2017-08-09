package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Usage: gg [words]")
		os.Exit(0)
	}

	var opener string
	switch runtime.GOOS {
	case "linux":
		opener = "xdg-open"
	case "darwin":
		opener = "open"
	default:
		fmt.Println("Unsupported OS")
		os.Exit(1)
	}

	params := strings.Join(flag.Args(), "+")
	url := "https://www.google.co.jp/search?q=" + params
	fmt.Printf("%s %s\n", opener, url)
	cmd := exec.Command(opener, url)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
