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
		err  bool
	}{
		{
			v:    "v0.0.1-2-gcbaffea-dirty",
			want: "gg version 0.0.1, build gcbaffea",
			err:  false,
		},
		{
			v:    "fatal: No names found, cannot describe anything.",
			want: "",
			err:  true,
		},
	}

	for _, test := range tests {
		cli := &CLI{}
		got, err := cli.Version(test.v)
		if test.err && err == nil {
			t.Errorf("test.err: %s, err: %s", test.err, err)
		}
		if !test.err && err != nil {
			t.Errorf("test.err: %s, err: %s", test.err, err)
		}
		if got != test.want {
			t.Errorf("got: %s, test.want: %s", got, test.want)
		}
	}
}

func TestCLIRun(t *testing.T) {
	tests := []struct {
		version   string
		args      []string
		want      int
		outStream []byte
		errStream []byte
		os        string
	}{
		{
			version:   "foobarbaz",
			args:      []string{"--version"},
			want:      1,
			outStream: nil,
			errStream: []byte("\"foobarbaz\" is not expected string format.\n"),
			os:        "linux",
		},
		{
			version:   "v0.0.1-2-gcbaffea-dirty",
			args:      []string{"--version"},
			want:      0,
			outStream: []byte("gg version 0.0.1, build gcbaffea\n"),
			errStream: nil,
			os:        "linux",
		},
		{
			version:   "v0.0.1-2-gcbaffea-dirty",
			args:      []string{"foo", "bar", "baz"},
			want:      1,
			outStream: nil,
			errStream: []byte("Unsupported OS"),
			os:        "windows",
		},
		{
			version:   "v0.0.1-2-gcbaffea-dirty",
			args:      []string{"--dry-run", "foo", "bar", "baz"},
			want:      0,
			outStream: []byte("xdg-open https://www.google.co.jp/search?q=foo+bar+baz\n"),
			errStream: nil,
			os:        "linux",
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
		}
		got := cli.Run(tt.args)

		if got != tt.want {
			t.Errorf("got: %d, tt.want: %d", got, tt.want)
		}
		if bytes.Compare(outStream.Bytes(), tt.outStream) != 0 {
			t.Errorf("outStream: %s, tt.outStream: %s", string(outStream.Bytes()), string(tt.outStream))
		}
		if bytes.Compare(errStream.Bytes(), tt.errStream) != 0 {
			t.Errorf("errStream: %v, tt.errStream: %v", string(errStream.Bytes()), string(tt.errStream))
		}
	}
}
