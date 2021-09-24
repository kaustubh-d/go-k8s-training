package main

import (
	"flag"
	"fmt"
)

func main() {
	messagePtr := flag.String("input", "Hello World!!!", "Enter message to be displayed.")

	flag.Parse()

	fmt.Println(*messagePtr)
}
