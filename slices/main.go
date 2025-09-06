package main

import (
	"fmt"
)

func main() {

	var myArrayVar [6]int
	myArrayVar[1] = 2
	myArrayVar[2] = 5
	myArrayVar[3] = 9

	var slice1 []int
	fmt.Printf("size: %d, value %v\n", len(slice1), slice1)

	slice1 = append(slice1, 10, 20, 30, 40, 50)
	fmt.Println(slice1)
	fmt.Printf("size: %d, value %v\n", len(slice1), slice1)

	mySlice := myArrayVar[2:4]
	fmt.Println(mySlice)
	fmt.Println("Size: ", len(mySlice))

	fmt.Println(&myArrayVar[2])
	fmt.Println(&mySlice[0])

	fmt.Println(myArrayVar)

	mySlice[0] = 80
	mySlice[1] = 100

	fmt.Println(myArrayVar)

	fmt.Println(myArrayVar[:4])
	fmt.Println(myArrayVar[2:])

	slice := make([]int, 3)
	fmt.Println(slice)
	fmt.Printf("size: %d, value %v\n", len(slice), slice)

}
