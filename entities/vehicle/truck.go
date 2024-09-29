package vehicle

type Truck struct {
	VehicleNumber int
}

func (truck Truck) GetVehicleNumberPlate() int {
	return truck.VehicleNumber
}

func (truck *Truck) SetVehicleNumberPlate(plateNumber int) {
	truck.VehicleNumber = plateNumber
}

func (truck *Truck) GetVehicleType() int {
	return 3
}
