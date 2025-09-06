package main

import "fmt"

func main() {

	m1 := make(map[int]string)
	m1[1] = "A"
	m1[2] = "B"
	m1[3] = "C"

	fmt.Println(m1)
	fmt.Println(m1[1])

	delete(m1, 2)
	fmt.Println(m1)

	m1[1] = "A2"
	fmt.Println(m1[1])

	m1[7] = ""
	fmt.Println(m1[7])
	fmt.Println(m1[99])

	v, ok := m1[7]
	fmt.Println("The value: ", v, "Present?", ok)

	_, ok = m1[99]
	fmt.Println("Present?", ok)

	m2 := map[int]string{
		4: "A",
		5: "C",
		6: "Z",
	}

	fmt.Println(m2)

}
