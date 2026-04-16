package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// get the file name as an arg -f
	filename := flag.String("f", "", "File to search")

	// get the keyword to search for as an arg -k
	keyword := flag.String("k", "", "Term to search for")

	// parse the arguments
	flag.Parse()

	if *filename == "" || *keyword == "" {
		fmt.Println("Usage: khoj -f <file> -k <keyword>")
		os.Exit(1)
	}

	// open the file to read
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	// Close on exit
	defer file.Close()

	// create a file scanner
	scanner := bufio.NewScanner(file)
	// track file number
	line_number := 0
	for scanner.Scan() {
		// starting index of keyword if exists
		keyword_index := strings.Index(scanner.Text(), *keyword)

		line_number++
		// if keyword exists print it from its starting index
		if keyword_index != -1 {
			fmt.Printf("[Line: %d]. \033[1;32m%s\033[0m\n", line_number, scanner.Text()[keyword_index:])
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
