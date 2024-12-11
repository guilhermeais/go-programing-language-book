package main

import (
	"fmt"
	"os"
	popcount "popcount/lib"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		n, err := strconv.ParseUint(arg, 10, 0)
		if err != nil {
			fmt.Printf("Error converting arg %s: %v", arg, err)
			return
		}

		fmt.Printf("Arg %s: %d\n", arg, popcount.PopCount(n))
	}
}
