package vehicle

type Bike struct {
	VehicleNumber int
}

func (bike Bike) GetVehicleNumberPlate() int {
	return bike.VehicleNumber
}

func (bike *Bike) SetVehicleNumberPlate(plateNumber int) {
	bike.VehicleNumber = plateNumber
}

func (bike *Bike) GetVehicleType() int {
	return 1
}
