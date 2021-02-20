package main

import "fmt"

// aliasing type
type Celsius float32
type Fahrenheit float32

// Function to convert celsius to fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {
	
	var temp Fahrenheit 
	temp = Fahrenheit((t * 9.0 / 5.0) + 32.0)
	return temp
}

func main() {
	fmt.Printf("Converting 10 C to F: %v\n", toFahrenheit(10))
}

