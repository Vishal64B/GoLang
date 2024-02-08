// Q6
// Develop a scenario based on your domain that incorporates looping, control structures, variables, and constants:/Calculate the factorial of a number.

package main

import (
	"fmt"
)

const expenseCategories = 3

type Expense struct {
	Category string
	Amount   float64
}

var categories = [expenseCategories]string{"Rent", "Bills", "Food"}

func main() {
	var totalSpending float64
	var allExpenses []Expense
	var numUsers int

	fmt.Println("Monthly Budget Program")

	for {
		// Asking the user wants to continue
		if numUsers > 0 {
			var continueInput string
			fmt.Print("\nDo you want to enter expenses for another user? (yes/no): ")
			fmt.Scanln(&continueInput)
			if continueInput != "yes" {
				break
			}
		}

		// Increment the number of users
		numUsers++

		fmt.Printf("\nUser %d Expenses:\n", numUsers)

		// Loop through expense categories
		for i := 0; i < expenseCategories; i++ {
			var spending float64

			// User input for each category
			fmt.Printf("Enter spending for %s: Rs ", categories[i])
			_, err := fmt.Scanln(&spending)
			if err != nil || spending < 0 {
				fmt.Println("Invalid input. Enter a positive numeric value.")
				i--
				continue
			}

			// Accumulate total spending
			totalSpending += spending

			// Create an Expense struct and append it to the slice
			expense := Expense{
				Category: categories[i],
				Amount:   spending,
			}
			allExpenses = append(allExpenses, expense)
		}
	}

	// Display individual expenses
	fmt.Println("\nIndividual Expenses:")
	for _, expense := range allExpenses {
		fmt.Printf("%s: Rs %.2f\n", expense.Category, expense.Amount)
	}

	// Display total spending
	fmt.Printf("\nTotal Monthly Spending for %d Users: Rs %.2f\n", numUsers, totalSpending)

	// Calculate and display the factorial of the number of users
	factorial := 1
	for i := 1; i <= numUsers; i++ {
		factorial *= i
	}
	fmt.Printf("Factorial of %d Users: %d\n", numUsers, factorial)
}
