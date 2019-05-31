// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type DuplicateLine struct {
	Count int
	Files map[string]int
}

func main() {
	counts := make(map[string]DuplicateLine)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.Count > 1 {
			fmt.Printf("Duplicate Line: %s\tOccurences:%d\n", line, n.Count)
			for name, count := range n.Files {
				fmt.Printf("\tFile name:%s\tOccurences:%d\n", name, count)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]DuplicateLine) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if duplicate, ok := counts[input.Text()]; ok {
			duplicate.Count++
			duplicate.Files[f.Name()]++
			counts[input.Text()] = duplicate
		} else {
			counts[input.Text()] = DuplicateLine{1, map[string]int{f.Name(): 1}}
		}
	}
	//	Note: ignoring potential errors from input.Err()
}
