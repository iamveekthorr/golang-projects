// Package main provides TODO: add description
package main

import (
	"flag"
	"fmt"
	"os"
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

func main() {
	var arg string

	// Define usage for stdout
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: meera --path <path>\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}

	flag.StringVar(&arg, "path", "", "Search path")
	flag.StringVar(&arg, "help", "", "Find help docs")

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		flag.Usage()
		os.Exit(2)
	}

	// If nothing was provided as the value of an argument
	if len(strings.TrimSpace(arg)) == 0 {
		fmt.Printf("%v", arg)
		flag.Usage()
		os.Exit(2)
	}

	fmt.Printf("%v", arg)
	fmt.Println("searching...")

}
