package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	frequency_map := map[string]int{}
	for _, word := range strings.Fields(s) {
		frequency_map[word] += 1
	}
	return frequency_map
}

func main() {
	wc.Test(WordCount)
}
