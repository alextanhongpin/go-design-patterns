package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.Build()

	if car.Wheels != 4 {
		t.Errorf("expected car to have 4 wheels, got %d", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("expected car structure to be equal 'Car', got %s", car.Structure)
	}

	if car.Seats != 5 {
		t.Errorf("expected car seats to be equal 5, got %d", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	bike := bikeBuilder.Build()

	if bike.Wheels != 2 {
		t.Errorf("expected bike to have 2 wheels, got %d", bike.Wheels)
	}

	if bike.Structure != "Motorbike" {
		t.Errorf("expected bike structure to be equal 'Motorbike', got %s", bike.Structure)
	}
}
