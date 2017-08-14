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
