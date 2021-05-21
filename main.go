package main

import "fmt"

func main() {

	// String

	var name string = "gerad"
	var name2 = "didier"
	var name3 string
	fmt.Println(name, name2, name3)

	name3 = "brownie"
	fmt.Println(name, name2, name3)

	// utiliseable seulment a lal creation de variable
	name4 := "rosie"

	fmt.Println(name, name2, name3, name4)

	// Int
	var ageOne int = 22
	var ageTwo int = ageOne + 5
	ageThree := ageTwo + 7

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory

	var numOne int8 = 25
	var numTwo int8 = -123
	// uint = unsigned int
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree)

	// Float
	var scoreOne float32 = 5.3
	var scoreTwo float64 = 56378290382309.3
	scoreThree := 2.2
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
