package gg

import "testing"

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
