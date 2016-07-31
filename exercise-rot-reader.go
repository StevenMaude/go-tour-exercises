package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, e error) {
	n, e = r.r.Read(b)
	for i, v := range b {
		if v >= 65 && v <= 77 || v >= 97 && v <= 109 {
			b[i] = b[i] + 13
		} else if v >= 78 && v <= 90 || v >= 110 && v <= 122 {
			b[i] = b[i] - 13
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
