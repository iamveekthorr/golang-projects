// Package main provides TODO: add description
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"

	"path/filepath"

	"os"
	"strings"
)

type Arg int

const (
	Invalid Arg = iota - 1
	Help
	Path
)

type Result struct {
	FileName    string
	FileDetails fs.FileInfo
}

func traverse(filename string, result *Result) fs.WalkDirFunc {
	return func(path string, fd fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		extensions := map[string]bool{
			".mp4":  true,
			".mkv":  true,
			".avi":  true,
			".mov":  true,
			".wmv":  true,
			".webm": true,
		}

		// Ignore case for each string.
		if strings.EqualFold(fd.Name(), filename) {
			info, err := fd.Info()
			if err != nil {
				return err
			}

			// extract the file extension
			ext := filepath.Ext(filename)
			if !extensions[ext] {
				log.Println("file extension in the error -> " + ext)
				return nil
			}

			if strings.Contains(fd.Name(), filename) {
				log.Println("File found -> " + filename)
				*result = Result{
					FileName:    path,
					FileDetails: info,
				}
				return filepath.SkipAll
			}
			return nil

		}

		return nil
	}
}

func findFile(root, filename string) (Result, error) {
	var result Result
	err := filepath.WalkDir(root, traverse(filename, &result))

	return result, err
}

func main() {
	var filename string
	var pathname string

	// Define usage for stdout
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: meera --path <path>\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}

	flag.StringVar(&pathname, "path", ".", "Search path") // Default path to current directory using (.) dot notation
	flag.StringVar(&filename, "filename", "", "name of the file to be found")

	// flag.Parse()
	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		flag.Usage()
		os.Exit(2)
	}

	// If nothing was provided as the value of filename
	if len(strings.TrimSpace(filename)) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	result, err := findFile(pathname, filename)
	if err != nil {
		log.Printf("%v", filename)
	}

	if result.FileName == "" {
		log.Printf("file not found: %s", filename)
		os.Exit(1)
	}

	log.Printf("ok value -> %s %s", result, pathname)

}
