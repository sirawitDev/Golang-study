package method

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type IntSlice []int

func (s Student) FullName() string { // func (receiver Type) MethodName() ReturnType
	return s.FirstName + " " + s.LastName
}

func (s IntSlice) Sum() int {
	total := 0

	for _, v := range s {
		total += v
	}

	return total
}

func FullNameMethod() {
	student := Student{
		FirstName: "Thanapakorn",
		LastName:  "Seechompu",
	}

	fullName := student.FullName()

	fmt.Println("Fullname of the student : ", fullName)
}

func SumInt() {
	numbers := IntSlice{1, 2, 3, 4, 5}

	fmt.Println("Sum : ", numbers.Sum())
}
