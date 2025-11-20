package variable

import "fmt"

func PrintVariable() {
	var fullName string = "SirawitDev"
	var age int = 25
	about := "Software Developer"

	fmt.Println(fullName, " age : ", age, " years old.")

	fmt.Printf("About me: %s\n", about)

	about = "Golang Enthusiast"
	fmt.Printf("Updated About me: %s\n", about)
}
