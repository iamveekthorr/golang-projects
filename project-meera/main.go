// Package main provides TODO: add description
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	// "log"
	"os"
	// "path/filepath"
	"strings"
)

type Arg int

const (
	Invalid Arg = iota - 1
	Help
	Path
)

func (a Arg) String() string {
	return [...]string{"Help", "Path", "Invalid"}[a]
}

func ParseArgs(a string) (Arg, error) {
	switch strings.ToLower(a) {
	case "help":
		return Help, nil
	case "path":
		return Path, nil
	default:
		return Invalid, fmt.Errorf("invalid Argument provided: %q", a)
	}

}

func (a *Arg) Set(val string) error {
	parsed, err := ParseArgs(val)

	if err != nil {
		return err
	}

	*a = parsed

	return nil
}

type Result struct {
	FileName    string
	FileDetails fs.FileInfo
}

func traverse(filename string, result *Result) fs.WalkDirFunc {
	return func(path string, fd fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if fd.Name() == filename {
			info, err := fd.Info()
			if err != nil {
				return err
			}
			log.Printf("%v", filepath.Ext(filename))

			*result = Result{
				FileName:    path,
				FileDetails: info,
			}

			return filepath.SkipAll
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
	// flag.StringVar(&arg, "help", "", "Find help docs")

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

	ok, err := findFile(pathname, filename)
	if err != nil {
		log.Printf("%v", filename)
	}

	log.Printf("ok value -> %v", ok.FileDetails.Size())

}
