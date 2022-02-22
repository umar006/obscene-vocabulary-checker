package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var inputFile string
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
		inputSentence := bufio.NewReader(os.Stdin)
		sentence, err := inputSentence.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if sentence == "exit\n" {
			fmt.Println("Bye!")
			break
		}

		for k := range taboo_words {
			regex := regexp.MustCompile("(?i)" + k)
			censor := strings.Repeat("*", len(k))
			sentence = regex.ReplaceAllString(sentence, censor)
		}
		fmt.Println(sentence)
	}
}
