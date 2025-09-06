package main

import (
	"fmt"
)

func main() {
	license := true
	age := 15

	// tarea a realizar

	//si license es true y age es 19 debe mostrar el siguiente mensaje: Puede seguir avanzando

	switch {
	case license && age == 19:
		fmt.Println("Puede seguir avanzando")
	case !(license) && age == 19:
		fmt.Println("No puede seguir circulando")
	case license && age == 15:
		fmt.Println("No puede seguir circulando, es menor de edad")
	}

}
