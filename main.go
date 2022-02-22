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
	defer file.Close()

	return file
}

func readFileToMap(file *os.File) map[string]struct{} {
	scanner := bufio.NewScanner(file)

	mapped := make(map[string]struct{})
	for scanner.Scan() {
		mapped[scanner.Text()] = struct{}{}
	}

	return mapped
}

func main() {
	var inputFile string
	fmt.Scan(&inputFile)

	openFile := openFile(inputFile)

	taboo_words := readFileToMap(openFile)

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
