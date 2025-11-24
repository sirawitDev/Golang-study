package structt

import "fmt"

// Define a struct type
type Student struct {
	Name   string
	Weight int
	Height int
	Grade  string
}

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	City    string
	Zipcode int
}

func StudentStruct() {
	var student1 Student
	student1.Name = "SirawitDev"
	student1.Weight = 70
	student1.Height = 170
	student1.Grade = "D"

	fmt.Println(student1)
}

func StudentStructArray() {
	var student [3]Student

	student[0] = Student{
		Name:   "Few",
		Weight: 60,
		Height: 180,
		Grade:  "A",
	}

	student[1] = Student{
		Name:   "Tar",
		Weight: 50,
		Height: 172,
		Grade:  "C",
	}

	student[2] = Student{
		Name:   "Poom",
		Weight: 64,
		Height: 176,
		Grade:  "A",
	}

	fmt.Println(student, "StructArray")
}

func StudentStructMap() {
	students := make(map[string]Student)

	students["st01"] = Student{Name: "Ohm", Weight: 55, Height: 165, Grade: "B"}
	students["st02"] = Student{Name: "Joy", Weight: 48, Height: 160, Grade: "B"}
	students["st03"] = Student{Name: "John", Weight: 85, Height: 177, Grade: "F"}

	fmt.Println("Students Map : ", students)

	fmt.Println("Student[st01] :", students["st01"])
}

func PersonStruct() {
	var person Person

	person.Name = "SirawitDev"
	person.Age = 22
	person.Address = Address{
		City:    "Nan",
		Zipcode: 55140,
	}

	// Alternative way to initialize a struct
	bob := Person{
		Name: "Bob",
		Age:  25,
		Address: Address{
			City:    "Metropolis",
			Zipcode: 5122,
		},
	}

	fmt.Println(person)
	fmt.Println(bob)
}
