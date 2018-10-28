// Echos the name of the program and the list of command-line arguments
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	program_name := os.Args[0]
	optional_args := os.Args[1:]
	args_str := ""

	fmt.Println("Program name: " + program_name)
	if len(optional_args) != 0
	fmt.Println("Args: " + strings.Join(optional_args, " "))
}
