package operator

import "errors"

func Div(x, y float64) float64 {
	if y == 0 {
		panic(errors.New("division mustn't be zero"))
	}
	return x / y
}
