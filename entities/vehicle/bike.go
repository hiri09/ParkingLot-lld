package vehicle

type Bike struct {
	BaseVehicle
}

func (bike Bike) GetVehicleType() string {
	return "Bike"
}
