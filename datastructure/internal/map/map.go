package mapp

import "fmt"

func MapExample() {
	myMap := make(map[string]int) // make[ประเภทข้อมูลของ key] ประเภทข้อมูลของ value

	myMap["apple"] = 10
	myMap["banana"] = 20
	myMap["orange"] = 30

	fmt.Println("Map elements:", myMap)

	fmt.Println("Apples : ", myMap["apple"])

	//Update value
	myMap["apple"] = 15
	fmt.Println("Updated Apples : ", myMap["apple"])

	delete(myMap, "banana")

	//Iterate over map
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	fmt.Println("---------")

	//Checking if a key exists
	myMap["pear"] = 12
	val, ok := myMap["pear"]
	if ok {
		fmt.Println("Pear's value : ", val)
	} else {
		fmt.Println("Pear not found inmap")
	}
}
