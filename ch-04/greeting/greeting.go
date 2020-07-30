// Package greeting exports functions for saying hello and hi
package greeting

import "fmt"

// Hello prints the word "Hello!"
func Hello() {
	fmt.Println("Hello!")
}

// Hi prints the work "Hi!"
func Hi() {
	fmt.Println("Hi!")
}

// AllGreetings prints both words
func AllGreetings() {
	Hello()
	Hi()
}