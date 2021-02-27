package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	value := args[0]

	fmt.Printf("Converting %s to integer\n", value)
	number, err := strconv.Atoi(value)

	if err != nil { // if it was an error, discontinue
		fmt.Printf("Value %s is not an integer - exiting with error\n", value)
		os.Exit(-1) 
	}
	fmt.Printf("The integer is %d\n", number)
	// rest of the code
}
