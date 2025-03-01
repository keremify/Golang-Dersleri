package variables

import "fmt"

func Demo1() {

	var (
		Studentname = "John"
		grade       = 1
		isPassed    = false
	)
	fmt.Println("Hello!, My Name is", Studentname, "I am a", grade, ". grade student. Exam is passed?", isPassed)
}
