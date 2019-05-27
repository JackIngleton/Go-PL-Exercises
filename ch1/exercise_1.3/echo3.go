// Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses Strings.join.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	duration := time.Since(start).Seconds() * 1000
	fmt.Printf("Duration: %f ms", duration)

	//	Duration ~= 0.028-0.035 ms
}
