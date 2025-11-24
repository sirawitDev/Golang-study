package interface_test

import (
	"fmt"
	"math"
)

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello"
}

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

func calculateArea(s Shape) float64 {
	return s.Area()
}

func InterfaceExample() {
	dog := Dog{Name: "Buddy"}
	person := Person{Name: "Alice"}
	makeSound(dog)
	makeSound(person)

	//Shape
	rectangle := Rectangle{20, 40}
	triangle := Triangle{10, 15}
	circle := Circle{7}

	fmt.Println(calculateArea(circle))
	fmt.Println(calculateArea(triangle))
	fmt.Println(calculateArea(rectangle))
}
