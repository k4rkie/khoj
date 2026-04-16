package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("===Khoj===")

	// get the file name as an arg
	file_name := os.Args[1]

	// read data into a buffer
	data_buff, err := os.ReadFile(file_name)

	if err != nil {
		log.Fatal(err)
	}

	file_data_text := string(data_buff)
	fmt.Print(file_data_text)
}
