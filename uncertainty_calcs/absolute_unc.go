package main

import "fmt"

func main() {

	var actual_val1 float64
	fmt.Println("Please print your first real value: ")
	fmt.Scan(&actual_val1)

	var unc1 float64
	fmt.Println("Please print your first uncertainty")
	fmt.Scan(&unc1)

	var actual_val2 float64
	fmt.Println("Please print your second real value: ")
	fmt.Scan(&actual_val2)

	var unc2 float64
	fmt.Println("Please print your second uncertainty")
	fmt.Scan(&unc2)

	unc_addition := unc1 + unc2

	result := fmt.Sprint(unc_addition)

	fmt.Println(result)

}
