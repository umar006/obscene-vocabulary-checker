package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func readFileToMap(file *os.File) map[string]struct{} {
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mapped := make(map[string]struct{})
	for scanner.Scan() {
		mapped[scanner.Text()] = struct{}{}
	}

	return mapped
}

func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(input)
}

func censorTabooWord(sentence string, taboo_words map[string]struct{}) {
	for k := range taboo_words {
		regex := regexp.MustCompile("(?i)" + k)
		censor := strings.Repeat("*", len(k))
		sentence = regex.ReplaceAllString(sentence, censor)
	}

	fmt.Println(sentence)
}

func main() {
	inputFile := userInput()
	openFile := openFile(inputFile)
	taboo_words := readFileToMap(openFile)

	for {
		sentence := userInput()

		if sentence == "exit\n" {
			fmt.Println("Bye!")
			break
		}

		censorTabooWord(sentence, taboo_words)
	}
}
