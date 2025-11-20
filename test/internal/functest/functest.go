// internal package ไม่สามารถถูกเรียกใช้จากภายนอกแพ็กเกจ test ได้

package functest

import "fmt"

func Functest() {
	fmt.Println("Hello from functest package!")
}
