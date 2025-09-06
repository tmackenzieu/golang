package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}

	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum++
	}
	fmt.Println(sum)

	// es como un while
	for {
		if sum > 1000 {
			break
		}
		sum++
	}

	fmt.Println(sum)

	arr := []int{1, 2, 3, 4, 5}

	fmt.Println()

	for i := range arr {
		fmt.Println("Index: ", i, "value: ", arr[i])
	}

	for i, v := range arr {
		fmt.Println("Index: ", i, "value: ", v)
	}

	for _, v := range arr {
		fmt.Println("value: ", v)
	}

	fmt.Println()

	map2 := map[string]float64{
		"A": 12.3,
		"B": 23.1,
		"C": 34,
	}

	for key, value := range map2 {
		fmt.Println("key: ", key, "value: ", value)
	}

	map3 := map[string][]int{
		"A": nil,
		"B": {2, 34, 1, 2, 4},
		"C": {4, 5, 3, 2, 1},
	}
	fmt.Println()

	for key, value := range map3 {
		fmt.Println("key", key)
		for _, v := range value {
			fmt.Println("value: ", v)
		}
		fmt.Println()
	}

}
