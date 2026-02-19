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
	flag.StringVar(&arg, "path", "", "Search path")
	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			flag.Usage()
			os.Exit(1)
		}
		fmt.Println("Error:", err)
		os.Exit(2)
	}

	fmt.Printf("%v", arg)
	fmt.Println("searching...")

}
