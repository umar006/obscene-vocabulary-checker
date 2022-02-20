package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var inputFile, inputWord string
	fmt.Scan(&inputFile)
	fmt.Scan(&inputWord)

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	taboo_words := make(map[string]struct{})
	for scanner.Scan() {
		taboo_words[scanner.Text()] = struct{}{}
	}

	lower := strings.ToLower(inputWord)
	_, ok := taboo_words[lower]
	fmt.Println(ok)
}
