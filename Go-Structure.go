/*
Go-Array Allow user to define same kind to data-item but in structure user can save various type of data-items.
ex:- To Define Structure
type student struct {
	name string
	fname string
	rollNo int
	class int
	school string
}
}
*/
package main

import "fmt"

type students struct {
	name    string
	fname   string
	rolln   int
	subject string
	school  string
}

func main() {
	var student1 students // Declare student 1 of student
	var student2 students

	student1.name = "Bruce"
	student1.fname = "wayne"
	student1.rolln = 19
	student1.subject = "Science"
	student1.school = "DPS"

	student2.name = "Adam"
	student2.fname = "warlock"
	student2.rolln = 25
	student2.subject = "History"
	student2.school = "Navyug"

	printstu(student1)
	printstu(student2)

}
func printstu(student students) {
	fmt.Println("________________________________________________\n")
	fmt.Println("Name of The Student        :", student.name)
	fmt.Println("Father Name of The Student :", student.fname)
	fmt.Println("Roll Number of the Student :", student.rolln)
	fmt.Println("Subject of the student     :", student.subject)
	fmt.Println("School of the Student      :", student.school)
	fmt.Println("________________________________________________")

}
