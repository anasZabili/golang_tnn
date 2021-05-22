package main

import "fmt"

func main() {
	// Print
	fmt.Print("hello world")

	// Println
	age := 22
	name := "anas"
	height := 1.77
	fmt.Println("hello")
	fmt.Println("my age is", age, "my name is", name)

	// Printf Formated string
	fmt.Printf("my name is %v my age is %v \n", name, age)
	fmt.Printf("my name is %q my age is %d \n", name, age)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("his height is %0.2f \n", height)

	// SprintF
	description := fmt.Sprintf("my age is %v and my name is %v", age, name)
	fmt.Printf("%v", description)
}
