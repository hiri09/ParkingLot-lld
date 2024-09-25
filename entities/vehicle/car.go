package vehicle

type Car struct {
	BaseVehicle
}

func (car Car) GetVehicleType() string {
	return "Car"
}
