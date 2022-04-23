package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (cipherReader rot13Reader) Read(b []byte) (int, error) {
    n, err := cipherReader.r.Read(b)
	var lb, ub byte
	var decode bool
	if err == nil {
		for i, _ := range b {
			encoded := b[i]
			decode = true
			if encoded >= 65 && encoded <= 90 {
				lb = 65
				ub = 90
			} else if encoded >= 97 && encoded <= 122 {
				lb = 97
				ub = 122
			} else {
				decode = false
			}
			if decode {
				decoded := encoded + 13
				if decoded > ub {
					diff := decoded - ub
					decoded = lb + diff - 1
				}
				b[i] = decoded
			}
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
