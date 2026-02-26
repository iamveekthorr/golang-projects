package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		input, err := reader.ReadString(' ')

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if num, err := strconv.Atoi(input); err != nil {
			fmt.Println(err, "not a number")
		} else {
			sum += num
		}

		if err == io.EOF {
			break
		}

		// only if there's an error
		if err != nil {
			fmt.Println("Error reading standard input:", err)
		}

		fmt.Printf("Sum: %v\n", sum)
		println(input, "from input")
	}
}
