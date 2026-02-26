package worker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result struct {
	Line       string
	Path       string
	LineNumber int
}

type Results struct {
	All []Result
}

func NewResult(line, path string, lineNumber int) Result {
	return Result{
		line,
		path,
		lineNumber}
}

func FindTextInFile(path string, searchString string) *Results {
	// open file
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	results := Results{make([]Result, 0)}

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchString) {
			result := NewResult(scanner.Text(), path, lineNumber)
			results.All = append(results.All, result)
		}

		lineNumber++
	}

	if len(results.All) == 0 {
		return nil
	}

	return &results // the ampersand operator returns a pointer to the Results struct

}
