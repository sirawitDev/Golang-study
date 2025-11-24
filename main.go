package main

import (
	"fmt" //แสดงผลลัพธ์บนหน้าจอ Console

	"github.com/google/uuid" //สร้าง UUID แบบสุ่ม

	"github.com/sirawitDev/go-example/controlstructure" //นำเข้าแพ็กเกจ controlstructure ที่สร้างขึ้นเอง
	"github.com/sirawitDev/go-example/datastructure"    //นำเข้าแพ็กเกจ datastructure ที่สร้างขึ้นเอง
	"github.com/sirawitDev/go-example/functions"        //นำเข้าแพ็กเกจ functions ที่สร้างขึ้นเอง
	"github.com/sirawitDev/go-example/pointer"          //นำเข้าแพ็กเกจ variable ที่สร้างขึ้นเอง
	"github.com/sirawitDev/go-example/test"             //นำเข้าแพ็กเกจ test ที่สร้างขึ้นเอง
	"github.com/sirawitDev/go-example/variable"         //นำเข้าแพ็กเกจ variable ที่สร้างขึ้นเอง
)

func main() {
	id := uuid.New() //สร้าง UUID ใหม่
	fmt.Print("Hello World 3")

	fmt.Printf("\nGenerated UUID: %s\n", id.String()) //แสดง UUID ที่สร้างขึ้นแบบ printf
	fmt.Println("Generated UUID:", id.String())       //แสดง UUID ที่สร้างขึ้นแบบ println

	sayHello() //เรียกใช้ฟังก์ชัน sayHello

	test.SayHelloTest() //เรียกใช้ฟังก์ชัน SayHelloTest จากแพ็กเกจ test

	variable.PrintVariable() //เรียกใช้ฟังก์ชัน PrintVariable จากแพ็กเกจ variable

	controlstructure.IfElseExample()
	controlstructure.IfElseGradeExample()
	controlstructure.DayOfWeekExample()
	controlstructure.PreProcess()
	controlstructure.ForLoopExample()
	controlstructure.DoWhileExample()

	datastructure.ArrayExample()
	datastructure.SliceExample()
	datastructure.MapExample()
	datastructure.StructExample()

	functions.FunctionsExample()

	pointer.PointerExample()
}

func sayHello() {
	fmt.Println("Hello from sayHello function!")
}

//nodemon --exec go run main.go --signal SIGTERM (nodemon watching)
//go install github.com/air-verse/air@latest (Go watching)
