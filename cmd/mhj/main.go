package main

import (
	"fmt"
	"os"
)

const version = "0.1.0-bootstrap"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
