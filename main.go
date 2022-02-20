package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var input string
	fmt.Scan(&input)

	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
