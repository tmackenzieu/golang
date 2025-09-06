package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota // esto es similar a un enum
	SUB
	DIV
	MUL
)

// No se puede tener multiples valores con puntos suspensivos, solo una variable
// tiene que ir al final la variable con puntos suspensivos...
// en el caso de que necesitemos hacerlo pero no es lo mejor se puede recibir values ...interface{}
// y con reflection manejar los tipos de datos.
func MSum(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func MOperations(op Operation, values ...float64) (float64, error) {

	if len(values) == 0 {
		return 0, errors.New("there aren't values")
	}

	sum := values[0]

	for _, v := range values[1:] {
		switch op {
		case SUM:
			sum += v
		case SUB:
			sum -= v
		case MUL:
			sum *= v
		case DIV:
			if v == 0 {
				return 0, errors.New("division by zero")
			}
			sum /= v
		}
	}

	return sum, nil
}

func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x
	return
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case MUL:
		return x * y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	}
	return 0, errors.New("invalid operation")
}

// si empieza con mayuscula es publica
func Add(x int, y int) int {
	return x + y
}

func RepeatString(str string, increment int) {
	for i := 0; i < increment; i++ {
		fmt.Println(str)
	}
}
