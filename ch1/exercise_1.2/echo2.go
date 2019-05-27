// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	//  Replace blank identifier, print index and value as part of loop (increasing i by 1 to get index or args, rather than slice).
	for i, arg := range os.Args[1:] {
		fmt.Println(i+1, arg)
	}
}
