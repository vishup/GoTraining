package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("Recovered after panic %v", r)
		}
	}()

	if _, err := strconv.Atoi(arg); err != nil {
		panic(err)
	}
}

//go build && ErrorHandling.exe 1
