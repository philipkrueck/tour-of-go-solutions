// Video solution: https://www.youtube.com/watch?v=ak6IilRFRQE&t=1s

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	for i := range n {
		b[i] = rot13(b[i])
	}
	return n, err
}

func rot13(b byte) byte {
	if (b >= 'a' && b <= 'm') || (b >= 'A' && b <= 'M') {
		return b + 13
	}
	if (b >= 'n' && b <= 'z') || (b >= 'N' && b <= 'Z') {
		return b - 13
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
