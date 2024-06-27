package route

import (
	"fmt"
	"main/transport"
)

type Route struct {
	vehicles []transport.PublicTransport
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) AddVehicle(v transport.PublicTransport) {
	r.vehicles = append(r.vehicles, v)
}

func (r *Route) ShowVehicles() {
	for _, v := range r.vehicles {
		fmt.Printf("%T\n", v)
	}
}
