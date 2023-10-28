package main

import "fmt"

func Reverse(word []rune) {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
}

func main() {

	var word string
	_, err := fmt.Scanf("%s", &word)
	if err != nil {
		panic(err)
	}
	runes := []rune(word)
	Reverse(runes)
	fmt.Printf("%s", string(runes))
}
