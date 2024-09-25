package vehicle

type BaseVehicle struct {
	VehicleNumber int
}

func (vehicle *BaseVehicle) GetVehicleNumberPlate() int {
	return vehicle.VehicleNumber
}

func (vehicle *BaseVehicle) SetVehicleNumberPlater(numberPlate int) {
	vehicle.VehicleNumber = numberPlate
}


