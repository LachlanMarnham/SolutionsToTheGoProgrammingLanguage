// Echos the name of the index and value of command line arguments 
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[1:] {
		fmt.Println(index, value)
	}
}
