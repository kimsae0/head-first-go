// tocelsius converts a temperature...
package main

import (
	"fmt"
	"github.com/matthewbrahms/head-first-go/ch-04/keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}