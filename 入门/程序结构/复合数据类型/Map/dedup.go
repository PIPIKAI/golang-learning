package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if !seen[input.Text()] {
			fmt.Printf("%q\n", input.Text())
			seen[input.Text()] = true
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "input %v\n", err)
		os.Exit(1)
	}
}
