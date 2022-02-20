package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var input string
	fmt.Scan(&input)

	file, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(file))
}
