package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: [path to json file]")
		return
	}

	fmt.Println("Checking if {file} is valid json", args[0])
	f, err := os.ReadFile(args[0])

	if err != nil {
		fmt.Println("Unable to open file")
		panic(err)
	}

	result := isValidJson(string(f))

	if result {
		fmt.Println("File is a valid json file.")
	} else {
		fmt.Println("File is NOT a valid json file.")
	}
}
