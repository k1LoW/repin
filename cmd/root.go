/*
Copyright © 2021 Ken'ichiro Oyama <k1lowxb@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/k1LoW/repin"
	"github.com/k1LoW/repin/version"
	"github.com/spf13/cobra"
)

var (
	rf          string
	keywords    []string
	nonl        bool
	inp         bool
	rawKeywords bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "repin [FILE]",
	Short:        "repin is a tool to replace strings between keyword pair",
	Long:         `repin is a tool to replace strings between keyword pair.`,
	Version:      version.Version,
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			src        io.Reader
			replace    io.Reader
			out        io.Writer
			start, end string
		)

		keyRep := strings.NewReplacer("\\n", "\n", "\\t", "\t")
		if rawKeywords {
			keyRep = &strings.Replacer{}
		}

		switch len(keywords) {
		case 1:
			start = keyRep.Replace(keywords[0])
			end = keyRep.Replace(keywords[0])
		case 2:
			start = keyRep.Replace(keywords[0])
			end = keyRep.Replace(keywords[1])
		default:
			return errors.New("--keyword is required 1 or 2")
		}

		if rf != "" {
			f, err := os.Open(filepath.Clean(rf))
			if err == nil {
				defer f.Close() // #nosec
				replace = f
			} else {
				replace = strings.NewReader(rf)
			}
		} else {
			fi, err := os.Stdin.Stat()
			if err != nil {
				return err
			}
			if (fi.Mode() & os.ModeCharDevice) != 0 {
				return errors.New("--replace is not set")
			} else {
				replace = os.Stdin
			}
		}

		{
			f, err := os.Open(filepath.Clean(args[0]))
			if err != nil {
				return err
			}
			defer f.Close() // #nosec
			src = f
		}

		b := new(bytes.Buffer)
		if inp {
			out = b
		} else {
			out = os.Stdout
		}

		if _, err := repin.Replace(src, replace, start, end, nonl, out); err != nil {
			return err
		}

		if inp {
			if err := os.WriteFile(filepath.Clean(args[0]), b.Bytes(), fs.ModePerm); err != nil {
				return err
			}
		}

		return nil
	},
}

func Execute() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	log.SetOutput(io.Discard)
	if env := os.Getenv("DEBUG"); env != "" {
		log.SetOutput(os.Stderr)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&rf, "replace", "r", "", "replace file path or string")
	rootCmd.Flags().StringSliceVarP(&keywords, "keyword", "k", []string{}, "keywords to use as a delimiter. If 1 keyword is specified, it will be used as the start and end delimiters; if 2 keywords are specified, they will be used as the start and end delimiters, respectively.")
	rootCmd.Flags().BoolVarP(&nonl, "no-newline", "N", false, "disable appending newlines")
	rootCmd.Flags().BoolVarP(&inp, "in-place", "i", false, "edit file in place")
	rootCmd.Flags().BoolVarP(&rawKeywords, "raw-keywords", "", false, "do not convert \\n or \\t of the entered keywords")
}
