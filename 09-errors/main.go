package main

import (
	"errors"
	"fmt"
	"go-initial/09-errors/operator"
)

func main() {
	var err error
	err = errors.New("This is an error")
	fmt.Println(err)
	fmt.Println(err.Error())

	err2 := fmt.Errorf("This format err, string: %s, number: %d", "my string", 23)
	fmt.Println(err2)
	fmt.Println(err2.Error())

	defer func() {
		fmt.Println("In main defer")
		r := recover()
		if r != nil {
			fmt.Println("recovered in ", r)
		}
	}()

	z := operator.Div(4, 0)

	fmt.Println("Z is", z)

}
