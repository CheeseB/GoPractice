package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readWords(ch chan string, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, isPrefix, err := reader.ReadLine()
		if isPrefix || err != nil {
			break
		}
		words := strings.Split(string(line), " ")
		for _, word := range words {
			ch <- word
		}
	}
}

func allWords(filename string) <-chan string {
	wordStream := make(chan string)

	go func() {
		defer close(wordStream)
		readWords(wordStream, filename)
	}()

	return wordStream
}

func main() {
	for w := range allWords("source/wordFile.txt") {
		fmt.Println(w)
	}
}
