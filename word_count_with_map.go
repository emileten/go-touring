// credits : https://go.dev/tour
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var fields []string = strings.Fields(s)
	var m map[string]int = make(map[string]int)
	for _, f := range fields {
		_, ok := m[f]
		if ok == true {
			m[f] = m[f] + 1
		} else {
			m[f] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
