// Q2
// You're tasked with building a student information system in GoLang for a school. Each student record needs to store details such as student ID, name, age, and grade. Define variables to store the information of a single student and assign values accordingly. Pay attention to selecting appropriate data types to represent each piece of information.

package main

import "fmt"

func main() {

	var noOfStudents int
	fmt.Printf("Student Information System: \n")
	fmt.Print("Enter the number of students: ")
	fmt.Scanln(&noOfStudents)

	var students []Student

	for i := 0; i < noOfStudents; i++ {
		var (
			name  string
			id    int
			age   int
			grade float64
		)
		fmt.Print("Enter Name of Student: ")
		fmt.Scan(&name)
		fmt.Print("Enter ID of Student: ")
		fmt.Scan(&id)
		fmt.Print("Enter Age of Student: ")
		fmt.Scan(&age)
		fmt.Print("Enter Grade of Student: ")
		fmt.Scan(&grade)

		student := Student{Name: name, StudentID: id, Age: age, Grade: grade}
		students = append(students, student)
	}

	fmt.Println("Student details are : ")
	for _, student := range students {
		// fmt.Printf("Student %v List : \n", i+1)
		fmt.Printf("{ Name : %v | ", student.Name)
		fmt.Printf("ID : %v | ", student.StudentID)
		fmt.Printf("Age : %v | ", student.Age)
		fmt.Printf("Grade : %v} ", student.Grade)
	}
}

type Student struct {
	StudentID int
	Name      string
	Age       int
	Grade     float64
}
