package vehicles

import "errors"

type Car struct {
	Time int
}

type Motorcycle struct {
	Time int
}

type Truck struct {
	Time int
}

type Vehicle interface {
	Distance() float64
}

const (
	CarVehicle        = "CAR"
	MotorcycleVehicle = "MOTORCYCLE"
	TruckVehicle      = "TRUCK"
)

// el & indica la dirección de memoria, lo trabajamos como puntero.
func New(v string, time int) (Vehicle, error) {
	switch v {
	case CarVehicle:
		return &Car{Time: time}, nil
	case MotorcycleVehicle:
		return &Motorcycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	}
	return nil, errors.New("Vehicle not found")
}

func (c *Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

func (c *Motorcycle) Distance() float64 {
	return 120 * (float64(c.Time) / 60)
}

func (c *Truck) Distance() float64 {
	return 70 * (float64(c.Time) / 60)
}
