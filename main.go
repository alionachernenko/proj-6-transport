package main

import (
	"fmt"
	"main/passenger"
	"main/route"
	"main/transport"
)

func main() {
	passenger := passengers.NewPassenger("Steve")

	fmt.Printf("Hey, %v, welcome on this beatiful trip!\n", passenger.Name)

	bus := transport.NewBus()
	train := transport.NewTrain()
	airplane := transport.NewAirplane()

	route := route.NewRoute()

	bus.TakePassenger(passenger)
	route.AddVehicle(bus)
	bus.DisembarkPassenger(passenger)

	train.TakePassenger(passenger)
	route.AddVehicle(train)
	train.DisembarkPassenger(passenger)

	airplane.TakePassenger(passenger)
	route.AddVehicle(airplane)
	airplane.DisembarkPassenger(passenger)

	route.ShowVehicles()
}
