package controlstructure

import (
	"fmt"

	"github.com/sirawitDev/go-example/controlstructure/internal/dayofweek"
)

func IfElseExample() {
	var score int = 62

	if score >= 70 {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

func IfElseGradeExample() {
	var score int = 82
	var grade string

	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else {
		grade = "F"
	}

	fmt.Printf("Your grade is: %s\n", grade)
}

func DayOfWeekExample() {
	dayofweek.DayOfWeekExample()
}

func PreProcess() {
	num1 := 5
	num2 := 10

	sumNum := num1 + num2

	if sumNum >= 10 {
		fmt.Println("Sum is greater than or equal to 10")
	}

	if sumNum := num1 + num2; sumNum >= 10 { // PreDefine;
		fmt.Println("Sum (from pre-process) is greater than or equal to 10")
	}
}

func ForLoopExample() {
	for i := 1; i < 10; i++ {
		fmt.Println("Number :", i)
	}
}

func DoWhileExample() {
	i := 10
	for {
		fmt.Println("NumberDoWhile :", i)
		i--
		if i <= 5 {
			break
		}
	}
}
