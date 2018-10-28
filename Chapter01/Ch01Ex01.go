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

	fmt.Println("Program name: " + program_name)
	// It would be nice to only print "Args: " if args were actually
	// provided by the user, but for that we need if statements (not yet
	// covered in the textbook
	fmt.Println("Args: " + strings.Join(optional_args, " "))
}
