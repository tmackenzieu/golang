package main

import (
	"fmt"
)

func main() {

	var value string
	var m1 = make(map[string]string) // la forma más común

	var myArrayVar []string

	m1["10"] = "notebook"
	m1["15"] = "tv"
	m1["21"] = "heladera"
	m1["27"] = "monitor"
	m1["35"] = "camara"

	for {
		fmt.Print("Ingresa un valor entero (0 para finalizar): ")

		fmt.Scanf("%s\n", &value)
		if value == "0" {
			break
		}

		// Buscar valor en el mapa
		if desc, ok := m1[value]; ok {
			myArrayVar = append(myArrayVar, desc)
		} else {
			myArrayVar = append(myArrayVar, "No encontrado")
		}
	}

	fmt.Println("Los valores del array son:", myArrayVar)
}
