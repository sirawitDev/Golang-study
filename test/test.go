package test

import (
	"fmt"

	"github.com/sirawitDev/go-example/test/internal/functest"
)

func SayHelloTest() {
	fmt.Println("Hello from test package!")

	functest.Functest()
}
