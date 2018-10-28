// Benchmarks the time taken to echo args by iterating over them with range and by
// strings.Join
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	sep := " "
	N := 1000000  // number of times to perform the target action
	N_as_float := float64(N)
	
	// iterative method
	range_start := time.Now()
	for i := 0; i < N; i++ {
		range_str := ""
		for _, arg := range os.Args[1:] {
			range_str += arg + sep
		}
	}
	range_duration := time.Since(range_start).Seconds()
	fmt.Println("Total time using for loop and range: ", range_duration, " seconds")
	fmt.Println("Average time per iteration for loop and range: ", range_duration / N_as_float, " seconds")

	// using strings.Join
	join_start := time.Now()
	for i := 0; i < N; i++ {
		_ = strings.Join(os.Args[1:], sep)
	}
	join_duration := time.Since(join_start).Seconds()
	fmt.Println("Total time using strings.Join: ", join_duration, " seconds")
	fmt.Println("Average time per iteration for strings.Join: ", join_duration / N_as_float, " seconds")
}
