package floor

import (
	"fmt"
	"parkinglot-system-design/entities/parkingspot"
	"parkinglot-system-design/entities/vehicle"
)

type ParkingFloor struct {
	FloorNumber          int
	DisplayBoard         int
	TotalParkingSpots    map[parkingspot.ParkingSpotType][]*parkingspot.Parkingspot
	AvailableSpotsByType map[parkingspot.ParkingSpotType]int
}

func (floor *ParkingFloor) SetParkingFloor(floorNumber int) {
	floor.FloorNumber = floorNumber
}

func (floor *ParkingFloor) GetFreeParkingSpot(parkingType parkingspot.ParkingSpotType) *parkingspot.Parkingspot {
	for _, parkingspot := range floor.TotalParkingSpots[parkingType] {
		if parkingspot.IsSpotAvailable {
			floor.DisplayBoard--
			floor.AvailableSpotsByType[parkingType]--
			return parkingspot
		}
	}
	return nil
}

func (floor *ParkingFloor) AddNewParkingSpotOnFloor(parkingSpotType parkingspot.ParkingSpotType, parkingSpot *parkingspot.Parkingspot) {
	parkingSpots := floor.TotalParkingSpots[parkingSpotType]
	parkingSpots = append(parkingSpots, parkingSpot)
	floor.TotalParkingSpots[parkingSpotType] = parkingSpots
}

func (floor *ParkingFloor) CanParkVehicle(vehicle vehicle.IVehicle) (bool, parkingspot.ParkingSpotType) {
	if floor.DisplayBoard <= 0 {
		return false, parkingspot.NoType
	}

	vehicleTyp := vehicle.GetVehicleType()
	vehicleType := parkingspot.ParkingSpotType(vehicleTyp)
	for i := 1; i < 4; i++ {
		spotType := parkingspot.ParkingSpotType(i)
		if floor.AvailableSpotsByType[spotType] > 0 && (spotType >= vehicleType) {
			return true, spotType
		}
	}
	return false, parkingspot.NoType
}

func (floor *ParkingFloor) VacateVehicleFromParkingLot(vehicle vehicle.IVehicle) bool {
	vehicleTyp := vehicle.GetVehicleType()
	vehicleType := parkingspot.ParkingSpotType(vehicleTyp)

	for i := 1; i < 4; i++ {
		spotType := parkingspot.ParkingSpotType(i)
		if spotType >= vehicleType {
			for _, parkingSpot := range floor.TotalParkingSpots[spotType] {
				if parkingSpot.Vehicle == vehicle {
					// Update floor status
					parkingSpot.VacateVehicleFromSpot()
					floor.DisplayBoard++
					floor.AvailableSpotsByType[spotType]++

					fmt.Printf("Vehicle %s has successfully vacated the parking spot on Floor %d.\n", vehicle.GetVehicleNumberPlate(), floor.FloorNumber)
					return true
				}
			}
		}
	}

	// If the vehicle is not found
	fmt.Printf("Failed to vacate. Vehicle %s is not found on Floor %d.\n", vehicle.GetVehicleNumberPlate(), floor.FloorNumber)
	return false
}
