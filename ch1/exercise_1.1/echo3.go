// Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Remove 1 from `os.Args[1:]`. Could also have changed to a 0.
	fmt.Println(strings.Join(os.Args[:], " "))
}