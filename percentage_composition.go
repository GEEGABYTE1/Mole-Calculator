package main

import "fmt"

func main() {

	var comp1 float64
	fmt.Println("Please type in the mass of one compound: ")
	fmt.Scan(&comp1)

	var comp2 float64
	fmt.Println("Please type in the mass of one compound: ")
	fmt.Scan(&comp2)

	var comp3 float64
	fmt.Println("Please type in the mass of one compound: ")
	fmt.Scan(&comp3)

	var comp4 float64
	fmt.Println("Please type in the mass of one compound: ")
	fmt.Scan(&comp4)

	var total_mass float64
	fmt.Println("Please type in the total mass of the compound: ")
	fmt.Scan(&total_mass)

	composition1 := (comp1 / total_mass) * 100
	composition2 := (comp2 / total_mass) * 100
	composition3 := (comp3 / total_mass) * 100
	composition4 := (comp4 / total_mass) * 100

	state1 := fmt.Sprintf("Percentage Composition of Compound 1: %f", composition1)
	state2 := fmt.Sprintf("Percentage Composition of Compound 2: %f", composition2)
	state3 := fmt.Sprintf("Percentage Composition of Compound 3: %f", composition3)
	state4 := fmt.Sprintf("Percentage Composition of Compound 4: %f", composition4)

	fmt.Println(state1)
	fmt.Println(state2)
	fmt.Println(state3)
	fmt.Println(state4)
}
