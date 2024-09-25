package vehicle

type Car struct {
	NumberPLate int
}

func (car *Car) GetVehicleNumberPlate() int {
	return car.NumberPLate
}

func (car *Car) SetVehicleNumberPlater(numberplate int) {
	car.NumberPLate = numberplate
}

func (car *Car) GetVehicleType() string {
	return "Car"
}
