package main

import "fmt"

func main() {

	var abs_val1 float64
	fmt.Println("Please type in your first abs val: ")
	fmt.Scan(&abs_val1)

	var unc1 float64
	fmt.Println("Please type in your first uncertainty: ")
	fmt.Scan(&unc1)

	relative_1 := unc1 / abs_val1

	var abs_val2 float64
	fmt.Println("Please type in your second abs val: ")
	fmt.Scan(&abs_val2)

	var unc2 float64
	fmt.Println("Please type in your seconds uncertainty: ")
	fmt.Scan(&unc2)

	relative_2 := unc2 / abs_val2

	complete_addition := relative_1 + relative_2

	result := fmt.Sprintf("The relative unc of your given values are: %f", complete_addition)

	fmt.Println(result)

}
