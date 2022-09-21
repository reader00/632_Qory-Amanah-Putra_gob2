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
}
