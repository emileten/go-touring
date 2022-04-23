package main

import "golang.org/x/tour/reader"

type InfiniteAReader struct{}

func (e InfiniteAReader) Read(b []byte) (int, error) {
	b[0] = 65
	return 1, nil
}

func main() {
	reader.Validate(InfiniteAReader{})
}
