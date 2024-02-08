// Develop a scenario based on your domain that incorporates looping, control structures, variables, and constants:/Calculate the factorial of a number.

package main

import "fmt"

func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Print("Enter a number: ")
	var num int
	fmt.Scan(&num)

	fmt.Printf("Factorial of %d: %d \n", num, factorial(num))
}
