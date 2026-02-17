// Package main provides TODO: add description
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%s", os.Args[1:])

	filePath := os.Args[1:][0]
	// 1. Get the data from the CLI
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Path to file does not exist: %s", filePath)
			return
		}
		fmt.Printf("Some error occurred: %s", err)
		return
	}
	fmt.Print(fileInfo)
	// 2. Check if the given arg is a path/file
	// 3. Use the given arg to determine if we need to get info on the video
}
