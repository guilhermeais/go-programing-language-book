package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("running...")

	files := os.Args[1:]

	if len(files) > 0 {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reding file: %v\n", err)
				continue
			}
			fmt.Printf("\nFilename %s:\n", f.Name())
			Dup(f)
			f.Close()
		}
	} else {
		Dup(os.Stdin)
	}
}

func Dup(r io.Reader) {
	counts := make(map[string]int)
	input := bufio.NewScanner(r)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
