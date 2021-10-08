package main

import (
	"fmt"
)

func main() {

	var molar_mass float64
	fmt.Println("Please type in the molar mass of the compound: ")
	fmt.Scan(&molar_mass)

	var moles float64
	fmt.Println("Please type in the number of moles of the compound: ")
	fmt.Scan(&moles)

	calculation := moles * molar_mass

	returned_val := fmt.Sprintf("The mass of your compound is approximately: %f", calculation)

	fmt.Print(returned_val)

}
