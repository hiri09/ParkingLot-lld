package vehicle

type Truck struct {
	BaseVehicle
}

func (truck Truck) GetVehicleType() string {
	return "Truck"
}
