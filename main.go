package main

import (
	"fmt"
	"os"
)

func main() {
	arg1 := os.Args[1]
	fmt.Printf("Hello, %s!\n", arg1)
}
