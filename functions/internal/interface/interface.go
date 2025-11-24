package interface_test

import (
	"fmt"
	"math"
)

// Speaker interface
type Speaker interface {
	Speak() string
}

// Dog struct
type Dog struct {
	Name string
}

// Speak method for Dog
func (d Dog) Speak() string {
	return "Woof!"
}

// Person struct
type Person struct {
	Name string
}

// Speak method for Person
func (p Person) Speak() string {
	return "Hello"
}

// makeSound function that takes a Speaker interface
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

// Calculate area of different shapes using Shape interface

// Shap interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	length, width float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Triangle struct
type Triangle struct {
	base, height float64
}

// Area method for Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

// Circle struct
type Circle struct {
	radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// calculateArea function that takes a Shape interface
func calculateArea(s Shape) float64 {
	return s.Area()
}

// InterfaceExample demonstrates the use of interfaces
func InterfaceExample() {
	// Sound demonstration
	dog := Dog{Name: "Buddy"}
	person := Person{Name: "Alice"}
	makeSound(dog)
	makeSound(person)

	// Area calculation
	rectangle := Rectangle{20, 40}
	triangle := Triangle{10, 15}
	circle := Circle{7}

	fmt.Println(calculateArea(circle))
	fmt.Println(calculateArea(triangle))
	fmt.Println(calculateArea(rectangle))
}
