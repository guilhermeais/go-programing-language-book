package main

import (
	fetch "fetch/lib"
	"fmt"
	"os"
	"time"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		err := fetch.Fetch(url, os.Stdout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
	}

	fmt.Println("using concurrent fetch")
	start := time.Now()
	ch := make(chan string)
	for _, url := range urls {
		go fetch.FetchWithChannel(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elasped\n", time.Since(start).Seconds())
}
