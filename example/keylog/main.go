package main

import (
	"fmt"
	ansi "github.com/solidiquis/ansigo"
)

func main() {
	stdin := make(chan string, 1)
	go ansi.GetChar(stdin)

	for {
		key := <-stdin
		fmt.Printf("You pressed: '%s'\n", key)
	}
}
