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

	for {
		fmt.Scan(&inputWord)

		if inputWord == "exit" {
			fmt.Println("Bye!")
			break
		}

		lower := strings.ToLower(inputWord)
		if _, ok := taboo_words[lower]; ok {
			replacer := strings.NewReplacer(lower, strings.Repeat("*", len(lower)))
			fmt.Println(replacer.Replace(lower))
		} else {
			fmt.Println(inputWord)
		}
	}
}
