// Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses Strings.join.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	duration := time.Since(start).Seconds() * 1000
	fmt.Printf("Duration: %f ms", duration)

	//	Duration ~= 0.03-0.045 ms
}