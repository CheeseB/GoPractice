package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int	{
	wordmap := make(map[string]int)
	words := strings.Fields(s)
	for _,word := range words {
		wordmap[word] = wordmap[word] + 1
	}
	return wordmap
}

func main() {
	wc.Test(WordCount)
}
