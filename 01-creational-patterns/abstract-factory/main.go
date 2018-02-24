package abstract_factory

import "fmt"

type Vehicle interface {
	NumWheels() int
	NumSeats() int
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("factory with id %d not recognized", f)
	}
}
