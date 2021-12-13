package repin

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// Replace strings between `start` and `end` of `src` to `replace` and write `out`.
func Replace(src, replace io.Reader, start, end string, nonl bool, out io.Writer) error {
	r, err := io.ReadAll(replace)
	if err != nil {
		return err
	}

	sl := len(start)
	el := len(end)
	flip := true
	c := 0
	sf := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if flip {
				if i >= sl {
					if string(data[i-sl:i]) == start {
						flip = !flip
						c += 1
						return i, data[:i], nil
					}
				}
			} else {
				if i >= el {
					if string(data[i-el:i]) == end {
						flip = !flip
						c += 1
						return i, data[:i], nil
					}
				}
			}
		}
		if atEOF {
			return 0, data, bufio.ErrFinalToken
		}
		return 0, nil, nil
	}

	scanner := bufio.NewScanner(src)
	scanner.Split(sf)
	in := false
	for scanner.Scan() {
		in = !in
		s := scanner.Text()
		b := []byte(s)
		if in {
			out.Write(b)
			if strings.Contains(s, start) {
				if !nonl && b[len(b)-1] != '\n' {
					out.Write([]byte("\n"))
				}
				out.Write(r)
			}
		} else {
			if !nonl {
				out.Write([]byte("\n" + end))
			} else {
				out.Write([]byte(end))
			}
		}
	}

	if c%2 != 0 {
		return errors.New("invalid keyword pair")
	}

	return nil
}

// Pick strings between `start` and `end` of `src` and write `out`.
func Pick(src io.Reader, start, end string, nonl bool, out io.Writer) error {
	sl := len(start)
	el := len(end)
	flip := true
	c := 0
	sf := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if flip {
				if i >= sl {
					if string(data[i-sl:i]) == start {
						flip = !flip
						c += 1
						return i, data[:i], nil
					}
				}
			} else {
				if i >= el {
					if string(data[i-el:i]) == end {
						flip = !flip
						c += 1
						return i, data[:i], nil
					}
				}
			}
		}
		if atEOF {
			return 0, data, bufio.ErrFinalToken
		}
		return 0, nil, nil
	}

	scanner := bufio.NewScanner(src)
	scanner.Split(sf)
	in := false
	for scanner.Scan() {
		in = !in
		s := scanner.Text()
		if !in {
			s = strings.TrimSuffix(s, end)
			if nonl {
				s = strings.TrimPrefix(strings.TrimSuffix(s, "\n"), "\n")
			}
			b := []byte(s)
			out.Write(b)
		}
	}

	if c%2 != 0 {
		return errors.New("invalid keyword pair")
	}

	return nil
}
