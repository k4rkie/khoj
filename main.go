package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/denormal/go-gitignore"
)

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

	filePaths, err := FindFiles(*rootDir)

	if err != nil {
		log.Printf("Failed to find files in the provided directory. %v", err)
		os.Exit(1)
	}

	for _, file := range filePaths {
		if err := SearchInFile(file, *keyword); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %v\n", err)
			continue
		}
	}

}

func FindFiles(rootDir string) (filePaths []string, err error) {

	// create a matcher instance to compare the files in gitignore and current file
	matcher, err := gitignore.NewFromFile(".gitignore")
	if err != nil {
		return filePaths, fmt.Errorf("Could read gitignore: %v", err)
	}

	err = filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("Error tracking file %v: %v", path, err)
		}

		// check if the current file / dir matches entry in the gitignore
		match := matcher.Match(path)
		if match != nil && match.Ignore() {
			// skip if the dir is in gitignore
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		// if not a dir add the path to the file paths
		if !d.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return filePaths, fmt.Errorf("Error finding files int he provided directory %v: %v", rootDir, err)
	}
	return filePaths, nil
}

func SearchInFile(path string, keyword string) (err error) {
	file, err := os.Open(path)

	if err != nil {
		return fmt.Errorf("Could not open file %s: %w", path, err)
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
			fmt.Printf("%v:%v:%v: %s\n", path, lineNumber, keyword_index+1, scanner.Text()[keyword_index:])
		}
	}

	// if error occurs while reading then exit
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Could not read file %s: %w", path, err)
	}

	return nil
}
