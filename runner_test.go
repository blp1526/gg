package gg

import (
	"bytes"
	"testing"
)

func TestRealRunnerCombineOutput(t *testing.T) {
	tests := []struct {
		c    string
		a    []string
		want []byte
		err  bool
	}{
		{
			c:    "echo",
			a:    []string{"foo", "bar"},
			want: []byte("foo bar\n"),
			err:  false,
		},
	}

	for _, tt := range tests {
		cli := &CLI{Runner: &RealRunner{}}
		got, err := cli.Runner.CombinedOutput(tt.c, tt.a...)
		if tt.err && err == nil {
			t.Errorf("tt.err: %v, err: %v", tt.err, err)
		}
		if !tt.err && err != nil {
			t.Errorf("tt.err: %v, err: %v", tt.err, err)
		}
		if bytes.Compare(got, tt.want) != 0 {
			t.Errorf("got: %s, tt.want: %s", string(got), string(tt.want))
		}
	}
}

func TestMockRunnerCombineOutput(t *testing.T) {
	tests := []struct {
		c    string
		a    []string
		want []byte
		err  bool
	}{
		{
			c:    "xdg-open",
			a:    []string{"https://www.google.co.jp/search?q=err+case"},
			want: nil,
			err:  true,
		},
		{
			c:    "xdg-open",
			a:    []string{"https://www.google.co.jp/search?q=foo+bar+baz"},
			want: nil,
			err:  false,
		},
	}
	for _, tt := range tests {
		cli := &CLI{Runner: &MockRunner{}}
		got, err := cli.Runner.CombinedOutput(tt.c, tt.a...)
		if tt.err && err == nil {
			t.Errorf("tt.err: %v, err: %v", tt.err, err)
		}
		if !tt.err && err != nil {
			t.Errorf("tt.err: %v, err: %v", tt.err, err)
		}
		if bytes.Compare(got, tt.want) != 0 {
			t.Errorf("got: %s, tt.want: %s", string(got), string(tt.want))
		}
	}
}
