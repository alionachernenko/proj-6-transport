package transport

import (
	"fmt"
	"main/passenger"
)

type Airplane struct {
	Passengers []passengers.Passenger
}

func NewAirplane() *Airplane {
	return &Airplane{}
}

func (a *Airplane) TakePassenger(passenger *passengers.Passenger) {
	a.Passengers = append(a.Passengers, *passenger)

	fmt.Printf("%v is boarder on airplane. Bon voyage!\n", passenger.Name)
}

func (a *Airplane) DisembarkPassenger(passenger *passengers.Passenger) {
	for i, p := range a.Passengers {
		if p.Name == passenger.Name {
			a.Passengers = append(a.Passengers[:i], a.Passengers[i+1:]...)
		}
	}

	fmt.Printf("%v was disembarked from airplane\n", passenger.Name)
}
