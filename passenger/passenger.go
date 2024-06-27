package passengers

type Passenger struct {
	Name string
}

func NewPassenger(name string)*Passenger {
	return &Passenger{
		Name: name,
	}
}
