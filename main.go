package main

import (
	"fmt"
	"math/rand"
	"parkinglot-system-design/entities/floor"
	"parkinglot-system-design/entities/parkinglot"
	"parkinglot-system-design/entities/parkingspot"
	"parkinglot-system-design/entities/vehicle"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Initialize floor1
	floor1 := &floor.ParkingFloor{
		FloorNumber:          1,
		DisplayBoard:         6,
		TotalParkingSpots:    make(map[parkingspot.ParkingSpotType][]*parkingspot.Parkingspot),
		AvailableSpotsByType: parkingspot.SetCountOfSpotByType(2),
	}

	for i := 0; i <= 5; i++ {
		if i <= 1 {
			floor1.AddNewParkingSpotOnFloor(parkingspot.CompactSpots, parkingspot.NewParkingSpot(700+i, parkingspot.CompactSpots))
		} else if i <= 3 {
			floor1.AddNewParkingSpotOnFloor(parkingspot.LargeSpots, parkingspot.NewParkingSpot(800+i-2, parkingspot.LargeSpots))
		} else {
			floor1.AddNewParkingSpotOnFloor(parkingspot.BikeSpots, parkingspot.NewParkingSpot(900+i-4, parkingspot.BikeSpots))
		}
	}

	// Initialize floor2
	floor2 := &floor.ParkingFloor{
		FloorNumber:          2,
		DisplayBoard:         9,
		TotalParkingSpots:    make(map[parkingspot.ParkingSpotType][]*parkingspot.Parkingspot),
		AvailableSpotsByType: parkingspot.SetCountOfSpotByType(3),
	}

	for i := 0; i <= 8; i++ {
		if i <= 2 {
			floor2.AddNewParkingSpotOnFloor(parkingspot.CompactSpots, parkingspot.NewParkingSpot(1000+i, parkingspot.CompactSpots))
		} else if i <= 5 {
			floor2.AddNewParkingSpotOnFloor(parkingspot.LargeSpots, parkingspot.NewParkingSpot(1100+i-3, parkingspot.LargeSpots))
		} else {
			floor2.AddNewParkingSpotOnFloor(parkingspot.BikeSpots, parkingspot.NewParkingSpot(1200+i-7, parkingspot.BikeSpots))
		}
	}

	// Initialize floor3
	floor3 := &floor.ParkingFloor{
		FloorNumber:          3,
		DisplayBoard:         9,
		TotalParkingSpots:    make(map[parkingspot.ParkingSpotType][]*parkingspot.Parkingspot),
		AvailableSpotsByType: parkingspot.SetCountOfSpotByType(3),
	}

	for i := 0; i <= 8; i++ {
		if i <= 2 {
			floor3.AddNewParkingSpotOnFloor(parkingspot.CompactSpots, parkingspot.NewParkingSpot(1300+i, parkingspot.CompactSpots))
		} else if i <= 5 {
			floor3.AddNewParkingSpotOnFloor(parkingspot.LargeSpots, parkingspot.NewParkingSpot(1400+i-4, parkingspot.LargeSpots))
		} else {
			floor3.AddNewParkingSpotOnFloor(parkingspot.LargeSpots, parkingspot.NewParkingSpot(1500+i-4, parkingspot.LargeSpots))
		}
	}

	// Create vehicles
	vehicles := []vehicle.IVehicle{
		&vehicle.Car{VehicleNumber: 1234},
		&vehicle.Car{VehicleNumber: 5678},
		&vehicle.Bike{VehicleNumber: 9101},
		&vehicle.Truck{VehicleNumber: 1121},
		&vehicle.Bike{VehicleNumber: 3141},
		&vehicle.Truck{VehicleNumber: 4151},
		&vehicle.Car{VehicleNumber: 5162},
		&vehicle.Bike{VehicleNumber: 6173},
		&vehicle.Car{VehicleNumber: 7184},
		&vehicle.Truck{VehicleNumber: 8195},
	}

	// Create a parking lot and add all floors
	parkinglot := parkinglot.NewParkingLot()
	parkinglot.AddFloor(floor1)
	parkinglot.AddFloor(floor2)
	parkinglot.AddFloor(floor3)

	// Step 1: First 6 vehicles park
	parkedVehicles := make(map[string]vehicle.IVehicle)
	fmt.Println("\n--- Step 1: Parking the first 6 vehicles ---")
	for i := 0; i < 6; i++ {
		vehicleToPark := vehicles[i]

		fmt.Printf("Attempting to park vehicle %d...\n", vehicleToPark.GetVehicleNumberPlate())
		isParked, entryTicket := parkinglot.EntryInparkingLot(vehicleToPark)
		if isParked {
			parkedVehicles[entryTicket] = vehicleToPark
			fmt.Printf("Vehicle %d parked. Ticket: %s.\n", vehicleToPark.GetVehicleNumberPlate(), entryTicket)
		} else {
			fmt.Println("Parking failed. Lot may be full.")
		}
		time.Sleep(1 * time.Second) // Pause between parking attempts
	}

	// Step 2: 3 vehicles exit
	fmt.Println("\n--- Step 2: Exiting 3 vehicles ---")
	exitCount := 3
	for i := 0; i < exitCount; i++ {
		for ticket, vehicleToExit := range parkedVehicles {
			fmt.Printf("Exiting vehicle %d with ticket %s...\n", vehicleToExit.GetVehicleNumberPlate(), ticket)
			parkinglot.ExitFromparkingLot(ticket, vehicleToExit)
			delete(parkedVehicles, ticket) // Remove exited vehicle from the list
			break
		}
		time.Sleep(1 * time.Second) // Pause between exits
	}

	// Step 3: Next 4 vehicles park
	fmt.Println("\n--- Step 3: Parking the next 4 vehicles ---")
	for i := 6; i < 10; i++ {
		vehicleToPark := vehicles[i]

		fmt.Printf("Attempting to park vehicle %d...\n", vehicleToPark.GetVehicleNumberPlate())
		isParked, entryTicket := parkinglot.EntryInparkingLot(vehicleToPark)
		if isParked {
			parkedVehicles[entryTicket] = vehicleToPark
			fmt.Printf("Vehicle %d parked. Ticket: %s.\n", vehicleToPark.GetVehicleNumberPlate(), entryTicket)
		} else {
			fmt.Println("Parking failed. Lot may be full.")
		}
		time.Sleep(1 * time.Second) // Pause between parking attempts
	}

	// Step 4: 4 vehicles exit
	fmt.Println("\n--- Step 4: Exiting 4 vehicles ---")
	exitCount = 4
	for i := 0; i < exitCount; i++ {
		for ticket, vehicleToExit := range parkedVehicles {
			fmt.Printf("Exiting vehicle %d with ticket %s...\n", vehicleToExit.GetVehicleNumberPlate(), ticket)
			parkinglot.ExitFromparkingLot(ticket, vehicleToExit)
			delete(parkedVehicles, ticket) // Remove exited vehicle from the list
			break
		}
		time.Sleep(1 * time.Second) // Pause between exits
	}

	fmt.Println("\nSimulation complete.")
}
