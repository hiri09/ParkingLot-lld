package parkingspot

import "parkinglot-system-design/entities/vehicle"

type Parkingspot struct {
	SpotId          int
	IsSpotAvailable bool
	Vehicle         vehicle.IVehicle
	ParkingSpotType ParkingSpotType
}

type ParkingSpotType int

const (
	BikeSpots ParkingSpotType = iota + 1
	CompactSpots
	LargeSpots
	NoType
)

func NewParkingSpot(spotId int, spotType ParkingSpotType) *Parkingspot {
	return &Parkingspot{
		SpotId:          spotId,
		IsSpotAvailable: true,
		ParkingSpotType: spotType,
		Vehicle:         nil,
	}
}

func (parkingspot *Parkingspot) ParkVehicleOnSpot(vehicle vehicle.IVehicle) {
	parkingspot.IsSpotAvailable = false
	parkingspot.Vehicle = vehicle
}

func (parkingspot *Parkingspot) VacateVehicleFromSpot() {
	parkingspot.IsSpotAvailable = true
	parkingspot.Vehicle = nil
}

func (parkingspot *Parkingspot) GetParkingSpotType() ParkingSpotType {
	return parkingspot.ParkingSpotType
}

func (parkingspot *Parkingspot) GetVehicle() vehicle.IVehicle {
	return parkingspot.Vehicle
}

func SetCountOfSpotByType(countOfSpot int) map[ParkingSpotType]int {
	spotsPerType := make(map[ParkingSpotType]int)
	spotsPerType[LargeSpots] = countOfSpot
	spotsPerType[CompactSpots] = countOfSpot
	spotsPerType[BikeSpots] = countOfSpot

	return spotsPerType
}
