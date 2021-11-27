package repin

import (
	"bytes"
	"io/fs"
	"strings"
	"testing"

	"github.com/josharian/txtarfs"
	"golang.org/x/tools/txtar"
)

func TestReplace(t *testing.T) {
	tests := []struct {
		txtar   string
		replace string
		start   string
		end     string
		nonl    bool
	}{
		{
			"testdata/replace/1.txtar",
			"$ echo hello world",
			"```",
			"```",
			false,
		},
		{
			"testdata/replace/2.txtar",
			"$ echo hello world",
			"```",
			"```",
			false,
		},
		{
			"testdata/replace/3.txtar",
			"Hello world!",
			"<h1>",
			"</h1>",
			false,
		},
		{
			"testdata/replace/4.txtar",
			"Hello world!",
			"<h1>",
			"</h1>",
			true,
		},
		{
			"testdata/replace/5.txtar",
			"$ echo hello world",
			"```",
			"```",
			false,
		},
	}
	for _, tt := range tests {
		ar, err := txtar.ParseFile(tt.txtar)
		if err != nil {
			t.Fatal(err)
		}
		fsys := txtarfs.As(ar)
		src, _ := fsys.Open("src")
		w, _ := fs.ReadFile(fsys, "want")
		want := string(w)

		out := new(bytes.Buffer)
		if err := Replace(src, strings.NewReader(tt.replace), tt.start, tt.end, tt.nonl, out); err != nil {
			t.Fatal(err)
		}
		got := out.String()

		if got != want {
			t.Errorf("\ngot  %#v\nwant %#v", got, want)
		}
	}
}