package main

import (
	"fmt"
	"go-initial/06-functions/function"
)

func main() {

	fmt.Println(function.Add(3, 4))
	function.RepeatString("Te quiero", 2)

	v, err := function.Calc(function.SUM, 3, 6)
	fmt.Println("Value: ", v, "Error: ", err)

	v, err = function.Calc(function.DIV, 3, 0)
	fmt.Println("Value: ", v, "Error: ", err)

	num1, num2 := function.Split(20)
	fmt.Println("num1: ", num1, "num2: ", num2)

	v = function.MSum(23, 12, 32, 12, 3, 1, 2, 3, 2, 1, 23, 12, 1)
	fmt.Println("suma: ", v)

	v, err = function.MOperations(function.SUM, 2, 7, 1)
	fmt.Println("SUM: ", v, "Error: ", err)

	v, err = function.MOperations(function.MUL, 2, 1, 3, 2, 1)
	fmt.Println("MUL: ", v, "Error: ", err)

	v, err = function.MOperations(function.DIV, 2, 1, 0, 2, 1)
	fmt.Println("DIV: ", v, "Error: ", err)

	//Patron factory , mas dinamicas ...
	fn := function.FactoryOperation(function.SUB)
	v = fn(2, 3)
	fmt.Println("SUB: ", v, "Error: ", err)

	fn = function.FactoryOperation(function.MUL)
	v = fn(2, 3)
	fmt.Println("SUB: ", v, "Error: ", err)
}
