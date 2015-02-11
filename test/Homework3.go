package main

import "fmt"

type Example1 func(int) bool // define a function type of variable

func Odd(integer int) bool {

	if integer%2 == 0 {

		return false
	}

	return true
}

func Even(integer int) bool {

	if integer%2 == 0 {

		return true
	}
	return false
}

// pass the function `f` as an argument to another function

func filter(slice []int, f Example1) []int {

	var result []int

	for _, value := range slice {

		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {

	rem := []int{1, 2, 3, 4, 5, 7, 8, 9, 0}

	fmt.Println("organize = ", rem)
	odd := filter(rem, Odd) // use function as values
	fmt.Println("Odd numbers are: ", odd)
	even := filter(rem, Even)
	fmt.Println("Even numbers are: ", even)
}
