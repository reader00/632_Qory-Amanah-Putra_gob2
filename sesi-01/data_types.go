package main

import "fmt"

func main() {

	// Data type

	var one uint8 = 32
	var two int8 = 45
	fmt.Printf("First data type: %T\n", one)
	fmt.Printf("second data type: %T\n", two)

	var height float32 = 168.5
	var status string = "single"
	fmt.Println("My height is ", height)
	fmt.Println("And I'm " + status)
}
