package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount counts the words, indicated by whitespace characters, in a string.
func WordCount(s string) map[string]int {
	split := strings.Fields(s)
	count := make(map[string]int)
	for _, w := range split {
		count[w]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
