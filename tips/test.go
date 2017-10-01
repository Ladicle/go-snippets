package main

import (
	"bufio"
	"fmt"
	"os"
)

func maine() {
	var aa string

	fs.StringVar(&aa, "aaa", "OK", "this is aho value")
	if err := fs.Parse(args); err != nil {
		os.Exit(1)
	}
	fmt.Printf("Hello world")

	file, err := os.Open(file_path)
	if err != nil {
		os.Exit(0)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			error_handling
			break
		}

	}
}
