package flight

type EmissionsOptions struct {
	NoOfPassengers string
	DistanceUnit   string
	Legs           []Leg
}

type Leg struct {
	DepartureAirport string
	ArrivalAirport   string
	CabinClass       string
}
