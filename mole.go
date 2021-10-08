package main

import (
	"fmt"
	"math"
)

func main() {

	constant := math.Pow(10, 23)
	avogrado := 6.02 * constant

	var mole float64
	fmt.Println("Please type in the number of molecules/particles")
	fmt.Scan(&mole)

	calculation := mole / avogrado

	fmt.Printf("Here is your number of particles in mol: %v", calculation)
}
