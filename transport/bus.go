package transport

import (
	"fmt"
	"main/passenger"
)

type Bus struct {
	Passengers []passengers.Passenger
}

func NewBus() *Bus {
	return &Bus{}
}

func (b *Bus) TakePassenger(passenger *passengers.Passenger) {
	b.Passengers = append(b.Passengers, *passenger)
	fmt.Printf("%v is boarder on bus. Bon voyage!\n", passenger.Name)
}

func (b *Bus) DisembarkPassenger(passenger *passengers.Passenger) {
	for i, p := range b.Passengers {
		if p.Name == passenger.Name {
			b.Passengers = append(b.Passengers[:i], b.Passengers[i+1:]...)
		}
	}

	fmt.Printf("%v was disembarked from bus\n", passenger.Name)
}
