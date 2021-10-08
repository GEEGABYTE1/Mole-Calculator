package main

import "fmt"

func main() {

	var total_mass float64
	fmt.Println("Please type the total mass of the compound: ")
	fmt.Scan(&total_mass)

	var experimental_mass float64
	fmt.Println("Please type the specific mass that you are looking for: ")
	fmt.Scan(&experimental_mass)

	calc := (experimental_mass / total_mass) * 100
	final_statement := fmt.Sprintf("Total mass percent: %f", calc)

	fmt.Print(final_statement)
}
