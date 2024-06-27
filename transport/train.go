package transport

import (
	"fmt"
	"main/passenger"
)

type Train struct {
	Passengers []passengers.Passenger
}

func NewTrain() *Train {
	return &Train{}
}

func (t *Train) TakePassenger(passenger *passengers.Passenger) {
	t.Passengers = append(t.Passengers, *passenger)

	fmt.Printf("%v is boarder on train. Bon voyage!\n", passenger.Name)
}

func (t *Train) DisembarkPassenger(passenger *passengers.Passenger) {
	for i, p := range t.Passengers {
		if p.Name == passenger.Name {
			t.Passengers = append(t.Passengers[:i], t.Passengers[i+1:]...)
		}
	}

	fmt.Printf("%v was disembarked from train\n", passenger.Name)
}
