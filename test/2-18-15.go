package main

import "fmt"

func main() {

	fmt.Println("Numbers from 1 -25:")
	i := 1
	for {
		if i > 25 {
			break
		}
		if i%2 == 1 {
			i++
			continue
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	sliceOfNames := []string{"Her", "his", "they", "Them"}

	gradYearMap := map[string]int{
		"Nick": 2015,
		"Dan":  2016,
		"Kim":  2012,
		"Mike": 2009,
	}

	for _, item := range sliceOfNames {
		fmt.Println(item, "year grad", gradYearMap[item])
	}

}
