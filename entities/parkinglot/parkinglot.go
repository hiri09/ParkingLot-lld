package parkinglot

import (
	"fmt"
	"parkinglot-system-design/entities/floor"
	"parkinglot-system-design/entities/vehicle"

	"github.com/google/uuid"
)

var ticketToParkingFloorMap = make(map[string]*floor.ParkingFloor)

type ParkingLot struct {
	ParkingFloors []*floor.ParkingFloor
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{
		ParkingFloors: []*floor.ParkingFloor{},
	}
}

func (lot *ParkingLot) AddFloor(floor *floor.ParkingFloor) {
	lot.ParkingFloors = append(lot.ParkingFloors, floor)
}

func (parkinglot *ParkingLot) EntryInparkingLot(vehicle vehicle.IVehicle) (bool, string) {
	for _, floor := range parkinglot.ParkingFloors {
		isSpaceAvailabe, parkingspotType := floor.CanParkVehicle(vehicle)
		if isSpaceAvailabe {
			freeParkingSpot := floor.GetFreeParkingSpot(parkingspotType)
			freeParkingSpot.ParkVehicleOnSpot(vehicle)
			entryTicket := uuid.New()
			ticketToParkingFloorMap[entryTicket.String()] = floor
			fmt.Printf("Your vehicle with number %s has been successfully parked in spot number %d.\n", vehicle.GetVehicleNumberPlate(), freeParkingSpot.SpotId)
			fmt.Printf("Your ticket number is: %s\n", entryTicket.String())
			return true, entryTicket.String()
		}
	}
	fmt.Println("Sorry, We are at our full capacity")
	return false, ""
}

func (parkinglot *ParkingLot) ExitFromparkingLot(ticket string, vehicle vehicle.IVehicle) bool {
	vehiclesParkedOnFloor, isPresent := ticketToParkingFloorMap[ticket]
	if isPresent {
		vehiclesParkedOnFloor.VacateVehicleFromParkingLot(vehicle)
		delete(ticketToParkingFloorMap, ticket)
		return true
	}
	return false
}
