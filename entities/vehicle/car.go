package vehicle

type Car struct {
	VehicleNumber int
}

func (car Car) GetVehicleNumberPlate() int {
	return car.VehicleNumber
}

func (car *Car) SetVehicleNumberPlate(plateNumber int) {
	car.VehicleNumber = plateNumber
}

func (car *Car) GetVehicleType() int {
	return 2
}
