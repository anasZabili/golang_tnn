package main

import "fmt"

func main() {

	// array taille fixe
	var ages [3]int = [3]int{20, 22, 32}
	var ages2 = [3]int{20, 22, 32}
	names := [3]string{"mario", "warrio", "luigi"}
	names[1] = "anas"
	fmt.Println(ages)
	fmt.Println(ages2, len(ages2))
	fmt.Println(names, len(names))

	// slices array dynamique
	var scores = []int{1, 2, 3}
	scores[2] = 25
	scores2 := append(scores, 100)
	fmt.Println(scores)
	fmt.Println(scores2)

	// slice ranges

	rangeOne := names[1:3]
	rangeTwo := names[0:]
	rangeThree := names[:2]
	fmt.Println("a range", rangeOne)
	fmt.Println("a range", rangeTwo)
	fmt.Println("a range", rangeThree)
}
