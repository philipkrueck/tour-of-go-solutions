// Video explanation: https://www.youtube.com/watch?v=YXYL0pmbjBU&t=2s

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCounts := make(map[string]int)
	for _, word := range strings.Fields(s) {
		wordCounts[word]++
	}
	return wordCounts
}

func main() {
	wc.Test(WordCount)
}
