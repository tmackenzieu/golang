package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myIntVar int
	fmt.Println(myIntVar)
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	var myArrayVar [6]int
	myArrayVar[1] = 2
	myArrayVar[2] = 5
	myArrayVar[3] = 9
	fmt.Println(myArrayVar)
	fmt.Println("Size: ", len(myArrayVar))
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myArrayVar, unsafe.Sizeof(myArrayVar), unsafe.Sizeof(myArrayVar)*8)

	myArrayVarString := [3]string{"a", "b", "c"}
	fmt.Println(myArrayVarString)

}
