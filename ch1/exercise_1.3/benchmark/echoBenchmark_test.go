// Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses Strings.join.
package benchmark

import (
	"fmt"
	"strings"
	"testing"
)

//  26444 ns/op              48 B/op          3 allocs/op
func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1([]string{"command", "arg1", "arg2", "arg3"})
	}
}

// 25581 ns/op              48 B/op          3 allocs/op
func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2([]string{"command", "arg1", "arg2", "arg3"})
	}
}

// 25404 ns/op              32 B/op          2 allocs/op
func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3([]string{"command", "arg1", "arg2", "arg3"})
	}
}

func echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}
