package datastructure

import (
	"fmt"

	"github.com/sirawitDev/go-example/datastructure/internal/array"
	mapp "github.com/sirawitDev/go-example/datastructure/internal/map"
	structt "github.com/sirawitDev/go-example/datastructure/internal/struct"
)

func ArrayExample() {
	array.ArrayExample()
}

func SliceExample() {
	array.SliceExample()
}

func MapExample() {
	mapp.MapExample()
}

func StructExample() {

	fmt.Println("---------")

	fmt.Println("Struct")

	structt.StudentStruct()

	fmt.Println("StructArray")

	structt.StudentStructArray()

	fmt.Println("StructMap")

	structt.StudentStructMap()

	fmt.Println("PersonStruct")
	structt.PersonStruct()
}
