package main

import (
	"fmt"
	"unsafe"
)

type foo struct {
	test1 int
	test2 string
}

func main() {

	someval := &foo{1, "Thing1"}

	pointer := unsafe.Pointer(someval)

	fmt.Printf("Addr: %v Val1 : %d Val2: %s\n",
		pointer,
		someval.test1,
		someval.test2)

	Switchval(someval)

	fmt.Printf("Addr: %v Val1 : %d Val2: %s\n",
		pointer,
		someval.test1,
		someval.test2)
}

func Switchval(someval *foo) {

	someval.test1 = 2
	someval.test2 = "Thing2"

	pointer := unsafe.Pointer(someval)

	fmt.Printf("Addr: %v Val1 : %d Val2: %s\n",
		pointer,
		someval.test1,
		someval.test2)
}
