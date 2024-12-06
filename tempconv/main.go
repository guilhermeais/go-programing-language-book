package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	tempconv "tempconv/lib"
)

func main() {
	for _, arg := range os.Args[1:] {
		temperatureHandler(arg)
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		temperatureHandler(input.Text())
	}
}

func temperatureHandler(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error converting the arg %v: %v\n", arg, err)
	}

	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)

	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}
