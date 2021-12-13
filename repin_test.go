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
		{
			"testdata/replace/6.txtar",
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

func TestLargeDataReplace(t *testing.T) {
	src := strings.NewReader(strings.Repeat("A", 10000))
	replace := strings.NewReader("hello")
	start := "```"
	end := "```"
	nonl := false
	out := new(bytes.Buffer)
	if err := Replace(src, replace, start, end, nonl, out); err != nil {
		t.Fatal(err)
	}
	got := len(out.String())

	if want := 10000; got != want {
		t.Errorf("\ngot  %#v\nwant %#v", got, want)
	}
}

func TestPick(t *testing.T) {
	tests := []struct {
		txtar string
		start string
		end   string
		nonl  bool
		want  string
	}{
		{
			"testdata/pick/1.txtar",
			"```",
			"```",
			false,
			"\n$ echo hello world\n",
		},
		{
			"testdata/pick/1.txtar",
			"```",
			"```",
			true,
			"$ echo hello world",
		},
		{
			"testdata/pick/2.txtar",
			"```",
			"```",
			false,
			"\n\n$ echo hello world\n\n",
		},
		{
			"testdata/pick/2.txtar",
			"```",
			"```",
			true,
			"\n$ echo hello world\n",
		},
		{
			"testdata/pick/3.txtar",
			"```",
			"```",
			false,
			"$ echo hello world\n",
		},
		{
			"testdata/pick/3.txtar",
			"```",
			"```",
			true,
			"$ echo hello world",
		},
		{
			"testdata/pick/4.txtar",
			"<h1>",
			"</h1>",
			false,
			"\nHello world!\n",
		},
		{
			"testdata/pick/4.txtar",
			"<h1>",
			"</h1>",
			true,
			"Hello world!",
		},
		{
			"testdata/pick/5.txtar",
			"<b>",
			"</b>",
			false,
			"Helloworld!",
		},
	}

	for _, tt := range tests {
		ar, err := txtar.ParseFile(tt.txtar)
		if err != nil {
			t.Fatal(err)
		}
		fsys := txtarfs.As(ar)
		src, _ := fsys.Open("src")
		out := new(bytes.Buffer)
		if err := Pick(src, tt.start, tt.end, tt.nonl, out); err != nil {
			t.Fatal(err)
		}
		got := out.String()

		if got != tt.want {
			t.Errorf("\ngot  %#v\nwant %#v", got, tt.want)
		}
	}

}
