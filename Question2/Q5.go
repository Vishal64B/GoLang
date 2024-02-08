// Q5
// You want to build a simple interest calculator in GoLang. The program should ask the user to input multiple sets of data, each containing the principal amount, the annual interest rate, and the number of years for which the interest is to be calculated. Use arrays to store the input data for each set, calculate the simple interest for each set using the formula: Simple Interest = (Principal Amount * Annual Interest Rate * Number of Years) / 100, and display the results.

package main

import "fmt"

func main() {
	var numSets int
	fmt.Print("Enter the number of sets of data: ")
	fmt.Scanln(&numSets)

	principal := make([]float64, numSets)
	annualRate := make([]float64, numSets)
	years := make([]int, numSets)

	// Input data for each set
	for i := 0; i < numSets; i++ {
		fmt.Printf("\nSet %d:\n", i+1)

		fmt.Print("Enter Principal Amount: ")
		fmt.Scanln(&principal[i])

		fmt.Print("Enter Annual Interest Rate: ")
		fmt.Scanln(&annualRate[i])

		fmt.Print("Enter Number of Years: ")
		fmt.Scanln(&years[i])
	}

	// Calculate and display simple interest for each set
	fmt.Println("\nResults:")
	for i := 0; i < numSets; i++ {
		simpleInterest := (principal[i] * annualRate[i] * float64(years[i])) / 100
		fmt.Printf("Set %d - Simple Interest: %.2f\n", i+1, simpleInterest)
	}
}
