// Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses Strings.join.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	duration := time.Since(start).Seconds() * 1000
	fmt.Printf("Duration: %f ms", duration)

	//	Duration ~= 0.03-0.045 ms
}
