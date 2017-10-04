package gg

import (
	"bytes"
	"testing"
)

func TestCLIOpener(t *testing.T) {
	tests := []struct {
		os   string
		want string
	}{
		{
			os:   "linux",
			want: "xdg-open",
		},
		{
			os:   "darwin",
			want: "open",
		},
		{
			os:   "windows",
			want: "",
		},
	}

	for _, test := range tests {
		cli := &CLI{OS: test.os}
		got := cli.Opener()
		if got != test.want {
			t.Errorf("os: %s, got: %s, test.want: %s",
				test.os, got, test.want)
		}
	}
}

func TestCLIAddr(t *testing.T) {
	tests := []struct {
		words []string
		want  string
	}{
		{
			words: []string{"foo", "bar"},
			want:  "https://www.google.co.jp/search?q=foo+bar",
		},
		{
			words: []string{"foo", "ほげ"},
			want:  "https://www.google.co.jp/search?q=foo+%E3%81%BB%E3%81%92",
		},
	}

	for _, test := range tests {
		cli := &CLI{}
		got := cli.Addr(test.words)
		if got != test.want {
			t.Errorf("\ngot: %s\ntest.want: %s", got, test.want)
		}
	}
}

func TestCLIVersion(t *testing.T) {
	tests := []struct {
		v    string
		want string
	}{
		{
			v:    "0.0.1",
			want: "gg version 0.0.1",
		},
	}

	for _, test := range tests {
		cli := &CLI{}
		got := cli.Version(test.v)
		if got != test.want {
			t.Errorf("got: %s, test.want: %s", got, test.want)
		}
	}
}

func TestCLIRun(t *testing.T) {
	tests := []struct {
		testDesc  string
		version   string
		args      []string
		want      int
		outStream []byte
		errStream []byte
		os        string
		runner    Runner
	}{
		{
			testDesc:  "Flag parse failure",
			version:   "foobarbaz",
			args:      []string{"---------"},
			want:      1,
			outStream: nil,
			errStream: []byte(
				`bad flag syntax: ---------

usage:
  gg [option] [word word word...]
        search words by the default web browser

option:
  -dry-run
    	print command line only and exit
  -version
    	print version and exit
`),
			os:     "linux",
			runner: &MockRunner{},
		},
		{
			testDesc:  "Version output success",
			version:   "0.0.1",
			args:      []string{"--version"},
			want:      0,
			outStream: []byte("gg version 0.0.1\n"),
			errStream: nil,
			os:        "linux",
			runner:    &MockRunner{},
		},
		{
			testDesc:  "Unsupported OS",
			version:   "0.0.1",
			args:      []string{"foo", "bar", "baz"},
			want:      1,
			outStream: nil,
			errStream: []byte("Unsupported OS"),
			os:        "windows",
			runner:    &MockRunner{},
		},
		{
			testDesc:  "Dry run",
			version:   "0.0.1",
			args:      []string{"--dry-run", "foo", "bar", "baz"},
			want:      0,
			outStream: []byte("xdg-open https://www.google.co.jp/search?q=foo+bar+baz\n"),
			errStream: nil,
			os:        "linux",
			runner:    &MockRunner{},
		},
		{
			testDesc:  "Command failure",
			version:   "0.0.1",
			args:      []string{"err", "case"},
			want:      1,
			outStream: []byte("xdg-open https://www.google.co.jp/search?q=err+case\n"),
			errStream: []byte("err+case\n"),
			os:        "linux",
			runner:    &MockRunner{},
		},
		{
			testDesc:  "Command success",
			version:   "0.0.1",
			args:      []string{"foo", "bar", "baz"},
			want:      0,
			outStream: []byte("xdg-open https://www.google.co.jp/search?q=foo+bar+baz\n"),
			errStream: nil,
			os:        "linux",
			runner:    &MockRunner{},
		},
	}

	for _, tt := range tests {
		Version = tt.version
		outStream := &bytes.Buffer{}
		errStream := &bytes.Buffer{}
		cli := &CLI{
			OutStream: outStream,
			ErrStream: errStream,
			OS:        tt.os,
			Runner:    tt.runner,
		}
		got := cli.Run(tt.args)

		if got != tt.want {
			t.Errorf("testDesc: %v, got: %d, tt.want: %d", tt.testDesc, got, tt.want)
		}
		if bytes.Compare(outStream.Bytes(), tt.outStream) != 0 {
			t.Errorf("testDesc: %v, outStream: %s, tt.outStream: %s",
				tt.testDesc, string(outStream.Bytes()), string(tt.outStream))
		}
		if bytes.Compare(errStream.Bytes(), tt.errStream) != 0 {
			t.Errorf("testDesc: %v, errStream: %v, tt.errStream: %v",
				tt.testDesc, string(errStream.Bytes()), string(tt.errStream))
		}
	}
}
