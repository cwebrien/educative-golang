package main

import "fmt"

func main() {
	values := [15]int{}

	for i := 0; i < 15; i++ {
		values[i] = i
	}

	fmt.Printf("%v\n", values)
}

