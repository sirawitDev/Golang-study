package functions

import (
	"fmt"

	interface_test "github.com/sirawitDev/go-example/functions/internal/interface"
	"github.com/sirawitDev/go-example/functions/internal/method"
)

// sayHello function that prints a greeting multiple times
func sayHello(name string, round int) {
	for i := 0; i < round; i++ {
		fmt.Printf("Hello %s\n", name)
	}
}

// add function that returns the sum of two integers
func add(number1 int, number2 int) int {
	return number1 + number2
}

// FunctionsExample demonstrates the use of functions, methods, and interfaces
func FunctionsExample() {
	fmt.Println("---------")

	fmt.Println("Functions")

	sayHello("FewDev", 2)

	sumNum := add(10, 20)

	fmt.Println(sumNum)

	fmt.Println("FunctionsMethod")

	method.FullNameMethod()

	method.SumInt()

	fmt.Println("Interface")

	interface_test.InterfaceExample()

}
