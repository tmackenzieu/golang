package main

import (
	"fmt"
)

func main() {
	yearsOld := 32
	fmt.Println("Operators")
	fmt.Println(yearsOld > 30)
	fmt.Println(yearsOld < 32)
	fmt.Println(yearsOld <= 32)
	fmt.Println(yearsOld >= 40)
	fmt.Println(yearsOld == 32)

	//OR
	fmt.Println()
	fmt.Println(yearsOld < 32 || yearsOld == 32)
	fmt.Println(yearsOld < 32 || yearsOld == 33)
	fmt.Println(yearsOld < 40 || yearsOld == 33)

	// AND
	fmt.Println()
	fmt.Println(yearsOld < 32 && yearsOld == 32)
	fmt.Println(yearsOld < 32 && yearsOld == 33)
	fmt.Println(yearsOld < 40 && yearsOld == 32)

	//NOT
	fmt.Println()
	fmt.Println(true)
	fmt.Println(!true)

	fmt.Println(yearsOld < 40)    //true
	fmt.Println(!(yearsOld < 40)) //false

	yearsOld = 15

	if yearsOld > 18 {
		fmt.Printf("%d is higher than 18\n", yearsOld)
	}

	boolVal := false

	if boolVal {
		fmt.Println("is true")
	} else {
		fmt.Println("is false")
	}

	//cuando necesitemos resultados de funciones, si necesitamos validar errores

	if value := true; value {
		fmt.Println("is true")
	}

	number := 3
	// else if
	if number == 1 {
		fmt.Println("one")
	} else if number == 2 {
		fmt.Println("two")
	} else if number == 3 {
		fmt.Println("three")
	}

	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	switch number := 1; number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	switch {
	case number > 0:
		fmt.Println("positive")
	case number < 0:
		fmt.Println("negative")
	case number == 0:
		fmt.Println("zero")
	default:
		fmt.Println("other")
	}

}
