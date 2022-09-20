package main

import "fmt"

func main() {

	// Print
	fmt.Println("Hello world")
	fmt.Println("Sesi 1")
	fmt.Println("Qory Amanah Putra")

	// Variable with data type
	var name string = "Qory Amanah Putra"
	var age int = 21

	fmt.Println("My name is " + name)
	fmt.Println("I'm ", age, " years old")

	// Variable without data type
	var favFood = "Chocolate"
	var favFruit = "Avocado"
	fmt.Println("I like " + favFood + " and " + favFruit)

	// Short declaration variable
	favThing := "Cloud"
	fmt.Println("I wanna be " + favThing)

	// Multiple variable declaration
	var firstDay, secondDay, thirdDay string = "Sunday", "Monday", "Tuesday"
	var fourthDay, isThursday, day = "Wednesday", true, 6
	_ = firstDay
	_, _, _, _, _ = secondDay, thirdDay, fourthDay, isThursday, day

	// Tipe data
	var one uint8 = 32
	var two int8 = 45
	fmt.Printf("First data type: %T\n", one)
	fmt.Printf("second data type: %T\n", two)

	var height float32 = 168.5
	var status string = "single"
	fmt.Println("My height is ", height)
	fmt.Println("And I'm " + status)

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
