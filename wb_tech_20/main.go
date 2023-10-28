package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	phrase := scanner.Text()
	words := strings.Split(phrase, " ")
	return words, nil
}

func Reverse(word []string) {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
}

func main() {
	words, err := Input()
	if err != nil {
		panic(err)
	}
	Reverse(words)
	for _, val := range words {
		fmt.Printf("%s ", val)
	}
}
