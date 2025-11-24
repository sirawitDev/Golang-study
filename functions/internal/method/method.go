package method

import "fmt"

// Student struct
type Student struct {
	FirstName string
	LastName  string
}

// IntSlice type
type IntSlice []int

// FullName method for Student
func (s Student) FullName() string { // func (receiver Type) MethodName() ReturnType
	return s.FirstName + " " + s.LastName
}

// Sum method for IntSlice
func (s IntSlice) Sum() int {
	total := 0

	for _, v := range s {
		total += v
	}

	return total
}

// FullNameMethod demonstrates the FullName method
func FullNameMethod() {
	student := Student{
		FirstName: "Thanapakorn",
		LastName:  "Seechompu",
	}

	fullName := student.FullName()

	fmt.Println("Fullname of the student : ", fullName)
}

// SumInt demonstrates the Sum method
func SumInt() {
	numbers := IntSlice{1, 2, 3, 4, 5}

	fmt.Println("Sum : ", numbers.Sum())
}
