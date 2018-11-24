// Prints the count and text of lines that appear more than once
// in the input. Reads from stdin or from a list of named files
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	strings_in_files := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, strings_in_files)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, strings_in_files)
			f.Close()
		}
	}
	displayResults(counts, strings_in_files)
}

func countLines(f *os.File, counts map[string]int, strings_in_files map[string][]string) {
	input := bufio.NewScanner(f)
	line := ""
	file_name := f.Name()
	for input.Scan() {
		line = input.Text()
		counts[line]++
		if !sliceContains(strings_in_files[line], file_name) {
			strings_in_files[line] = append(strings_in_files[line], file_name)
		}
	}
}

func sliceContains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func displayResults(counts map[string]int, file_string_map map[string][]string) {
	fmt.Println("COUNT\tLINE\tFILES")
	for line, n := range counts {
		if n > 1 {
			file_names := file_string_map[line]
			file_name_str := strings.Join(file_names, ", ")
			fmt.Printf("%d\t%s\t%s\n", n, line, file_name_str)
		}
	}
}
