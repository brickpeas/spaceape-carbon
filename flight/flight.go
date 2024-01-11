package flight

type EmissionsOptions struct {
	FlightType     string `json:"type,omitempty"`
	NoOfPassengers string `json:"passengers,omitempty"`
	DistanceUnit   string `json:"distance_unit,omitempty"`
	Legs           []Leg  `json:"legs,omitempty"`
}

type Leg struct {
	DepartureAirport string `json:"departure_airport,omitempty"`
	ArrivalAirport   string `json:"destination_airport,omitempty"`
	CabinClass       string `json:"cabin_class,omitempty"`
}
