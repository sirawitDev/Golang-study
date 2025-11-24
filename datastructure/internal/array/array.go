package array

import "fmt"

func ArrayExample() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	fmt.Println("---------")

	for i := 0; i < len(myArray); i++ {
		fmt.Printf("Element at index %d: %d\n", i, myArray[i])
	}

	fmt.Println("Array elements:", myArray)
}

func SliceExample() {
	mySlice := []int{10, 20, 30, 40, 50}

	fmt.Println("---------")

	fmt.Println("Slice elements:", mySlice)
	fmt.Println("Length of slice:", len(mySlice))
	fmt.Println("Capacity of slice:", cap(mySlice))

	fmt.Println("---------")

	// Slicing the slice
	subSlice := mySlice[1:3]
	fmt.Println("Sub-slice elements:", subSlice)
	fmt.Println("Length of sub-slice:", len(subSlice))
	fmt.Println("Capacity of sub-slice:", cap(subSlice))

	fmt.Println("---------")

	//Appending elements to the slice
	var mySliceAppend []int

	mySliceAppend = append(mySliceAppend, 100)
	mySliceAppend = append(mySliceAppend, 200, 300)

	fmt.Println("Appended slice elements:", mySliceAppend)

	fmt.Println("---------")

	// Converting array to slice
	var myArray [3]int

	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	mySliceFromArray := myArray[:]

	mySliceFromArray = append(mySliceFromArray, 40, 50)

	fmt.Println("Slice from array elements:", mySliceFromArray)

	fmt.Println("---------")
}
