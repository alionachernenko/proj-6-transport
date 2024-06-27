package transport

import "main/passenger"

type PublicTransport interface {
	TakePassenger(*passengers.Passenger)
	DisembarkPassenger(*passengers.Passenger)
}
