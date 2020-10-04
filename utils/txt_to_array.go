package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./wordlist.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr []string
	for scanner.Scan() {
		arr = append(arr, fmt.Sprintf(" \"%s\",", scanner.Text()))
	}

	fmt.Println(arr)

}
