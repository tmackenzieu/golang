package main

import (
	"fmt"
)

func main() {

	var value int
	var myArrayVar []int

	for {
		fmt.Print("Ingresa un valor entero (0 para finalizar): ")
		fmt.Scanf("%d\n", &value)
		if value == 0 {
			break
		}
		myArrayVar = append(myArrayVar, value)
	}

	fmt.Println("Los valores del array son:", myArrayVar)
}
