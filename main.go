package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var ignore = [10]string{".json", "node_modules/", "README.md"}

func main() {

	// get the file name as an arg -f
	rootDir := flag.String("d", ".", "Directory to search in")
	// get the keyword to search for as an arg -k
	keyword := flag.String("k", "TODO", "Term to search for")
	// parse the arguments
	flag.Parse()

	if *rootDir == "" || *keyword == "" {
		fmt.Println("Usage: khoj -d <dir> -k <keyword>")
		os.Exit(1)
	}

	err := filepath.WalkDir(*rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			searchInFile(path, *keyword)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", *rootDir, err)
	}
}

func searchInFile(path string, keyword string) {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 0
	for scanner.Scan() {
		// starting index of keyword if exists
		keyword_index := strings.Index(scanner.Text(), keyword)

		lineNumber++
		// if keyword exists print it from its starting index
		if keyword_index != -1 {
			fmt.Printf("[File: %s] [Line: %d]. \033[1;32m%s\033[0m\n", path, lineNumber, scanner.Text()[keyword_index:])
		}
	}

	// if error occurs while reading then exit
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
