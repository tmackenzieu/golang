package main

import (
	"encoding/json"
	"fmt"
	"go-initial/structsInterface/structs"
	"go-initial/structsInterface/vehicles"
)

func main() {

	var p1 structs.Product
	p1.ID = 12
	p1.Name = "Taby"
	fmt.Println(p1)

	p2 := structs.Product{}
	p2.ID = 13
	p2.Name = "Test"
	fmt.Println(p2)

	// No es la mejor práctica
	p3 := structs.Product{2, "test 3", structs.Type{}, 12.21, 1, nil}
	fmt.Println(p3)

	// la mejor práctica sí no quiero ingresar todos los campos

	p4 := structs.Product{
		ID:   12,
		Name: "Test4",
		Type: structs.Type{},
	}

	fmt.Println(p4)

	p5 := structs.Product{
		ID:    14,
		Name:  "Heladera marca sarasa",
		Price: 40000,
		Type: structs.Type{
			ID:          1,
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tags:  []string{"Heladera", "Freezer", "Refrigerador"},
		Count: 5,
	}

	p6 := structs.Product{
		ID:    14,
		Name:  "Naranja",
		Price: 20,
		Type: structs.Type{
			Code:        "B",
			Description: "Alimento",
		},
		Tags:  []string{"Alimento", "Verdura"},
		Count: 20,
	}

	p7 := structs.Product{
		ID:    17,
		Name:  "Cortina",
		Price: 2700,
		Type: structs.Type{
			Code:        "C",
			Description: "Hogar",
		},
		Tags:  []string{"hogar", "cortina"},
		Count: 3,
	}

	v, err := json.Marshal(p5)
	fmt.Println(err)
	fmt.Println(string(v))
	fmt.Println("El precio Total es:", p5.TotalPrice())
	fmt.Println(p5.Name)
	fmt.Println(p5.Tags)
	p5.SetName("other name")
	p5.AddTags("tag1", "tag2", "tag3")
	fmt.Println(p5.Name)
	fmt.Println(p5.Tags)

	car := structs.NewCar(11212)
	car.AddProducts(p5, p6, p7)

	fmt.Println("PRODUCTS CAR")
	fmt.Println("Total Products: ", len(car.Products))
	fmt.Println("Total Car:", car.Total())

	fmt.Println()
	fmt.Println("VEHICLES")

	car1 := vehicles.Car{Time: 120}
	fmt.Println(car1.Distance())

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "MOTORCYCLE", "TRUCK"}

	time := 400
	for _, v := range vArray {
		fmt.Println(v)
		vech, err := vehicles.New(v, time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		distance := vech.Distance()
		fmt.Println("Distance", distance)
	}
}
