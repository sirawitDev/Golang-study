package pointer

import "fmt"

func changeValue(val int) {
	val = 50
}

func PointerExample() {
	x := 20
	changeValue(x)
	fmt.Println(x) // Output: 20 (x is unchanged)
}
