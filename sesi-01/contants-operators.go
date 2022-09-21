package main

import "fmt"

func main() {
	// Constant
	const pi float32 = 3.14
	fmt.Println("PI value is ", pi)

	// Arithmetic Operator
	fmt.Println("2 + 2 = ", 1+2)
	fmt.Println("4 - 2 = ", 4-2)
	fmt.Println("2 * 2 = ", 2*2)
	fmt.Println("4 / 2 = ", 4/2)
	fmt.Println("4 % 2 = ", 4%2)

	// Relational Operator
	ten := 10
	aHundred := 100
	fmt.Println("Ten is equal to a hundred : ", ten == aHundred)
	fmt.Println("Ten is less than a hundred : ", ten < aHundred)
	fmt.Println("Ten is greater than hundred : ", ten > aHundred)
	fmt.Println("Ten is equal to or less than a hundred : ", ten <= aHundred)
	fmt.Println("Ten is equal to or greate than a hundred : ", ten >= aHundred)

	// Logical Operator
	var x = true
	var y = false
	fmt.Println("x && y : ", x && y)
	fmt.Println("x || y : ", x || y)
	fmt.Println("!x: ", !x)
	fmt.Println("!y : ", !y)
}
