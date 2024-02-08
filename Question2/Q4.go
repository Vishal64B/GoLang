// Q4
// You are tasked with creating a grade calculator program in GoLang. The program should prompt the user to input their scores in three subjects: Math, Science, and English. Based on these scores, calculate the average grade (assuming each subject is equally weighted) and display the corresponding letter grade (A, B, C, D, or F) using control flow.

package main

import "fmt"

func main() {
	fmt.Println("Grade Calculator: ")
	// User input
	fmt.Print("Enter Math score: ")
	var math float64
	fmt.Scanln(&math)

	fmt.Print("Enter Science score: ")
	var science float64
	fmt.Scanln(&science)

	fmt.Print("Enter English score: ")
	var english float64
	fmt.Scanln(&english)

	// Calculating average score
	averageScore := (math + science + english) / 3

	// Determining letter grade based on average score
	var letterGrade string

	switch {
	case averageScore >= 90:
		letterGrade = "A"
	case averageScore >= 80:
		letterGrade = "B"
	case averageScore >= 70:
		letterGrade = "C"
	case averageScore >= 60:
		letterGrade = "D"
	default:
		letterGrade = "F"
	}
	fmt.Printf("Average Score : %.2f", averageScore)
	fmt.Printf("\nGrade Obtained: %s\n", letterGrade)
}
