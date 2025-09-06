package main

import "fmt"

type MyStruct struct {
	ID   int
	Name string
}

func main() {
	var i int
	var iP *int
	i = 34

	iP = &i
	fmt.Println(&i)
	fmt.Println(iP)

	fmt.Println(i)
	fmt.Println("&", &iP)
	*iP = 1
	fmt.Printf("val i:, %d, val pointer i: %d\n", i, *iP)
	fmt.Println()

	myVar := 30
	fmt.Printf("addrs:, %p, val pointer i: %v\n", &myVar, myVar)
	increment(myVar)
	fmt.Println(myVar)

	incrementP(&myVar)
	fmt.Println(myVar)

	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)

	fmt.Println("CREATE SLICE INITIAL")
	fmt.Printf("addrs:, %p, val pointer i: %v\n", &mySlice, mySlice)
	fmt.Printf("addrs 1:, %p, addrs 2:, %p, addrs 3:, %v\n", &mySlice[0], &mySlice[1], &mySlice[2])
	fmt.Println("CALL UPDATE SLICE")
	updateSlice(mySlice)
	fmt.Println(mySlice)
	fmt.Println()

	fmt.Println("STRUCT")
	myStruct := &MyStruct{ID: 1234, Name: "Test"}
	fmt.Printf("addrs:, %p\n", myStruct)
	fmt.Printf("id:, %d, name: %s\n", myStruct.ID, myStruct.Name)
	fmt.Println("UPDATE STRUCT")
	updateStruct(myStruct)
	fmt.Printf("id:, %d, name: %s\n", myStruct.ID, myStruct.Name)

}

func updateStruct(v *MyStruct) {
	fmt.Printf("addrs in function:, %p\n", v)
	v.ID = 999
	v.Name = "Update Struct"
}

func updateSlice(v []int) {
	fmt.Printf("addrs:, %p\n", v)
	fmt.Printf("addrs 1:, %p, addrs 2:, %p, addrs 3:, %p\n", &v[0], &v[1], &v[2])
	v[0] = 12
}

func increment(val int) {
	fmt.Println(&val)
	val++
}

func incrementP(val *int) {
	fmt.Println(val)
	*val++
}
