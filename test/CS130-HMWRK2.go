package main

import "fmt"

//THis is the Struct

type Car struct {
	model string
	price int
}

//Some Functions

func printCar(car Car) {

	fmt.Println("Discouted Model:\n")
	fmt.Println(car.model, car.price)
}

func buyCar(car Car) (string, int) {
	return car.model, car.price
}

//Some Constants we can use
const (
	BASEVAL = 500000
)

func main() {
	var car = Car{
		model: "Ferrari",
		price: 200000,
	}

	printCar(car)
	var model, _ = buyCar(car)
	fmt.Println(model)

	fmt.Println("\nbase model:")
	fmt.Println(BASEVAL)

}
