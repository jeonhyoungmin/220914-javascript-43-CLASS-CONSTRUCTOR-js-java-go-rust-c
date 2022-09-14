package main

import "fmt"

//struct definition
type Student struct {
	name string
	rollNumber int
}

// associate method to Student struct type
func (s Student) PrintDetails() {
	fmt.Println("Student Details\n---------------")
	fmt.Println("Name :", s.name)
	fmt.Println("Roll Number :", s.rollNumber)
}

func main() {
	var stud1 = Student{name: "Anu", rollNumber:14}
	stud1.PrintDetails()
}