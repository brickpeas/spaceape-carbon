package flight

// EmissionsOptions are a set of overrides used when creating a new Client.
type EmissionsOptions struct {
	EmissionType   string      `json:"type,omitempty"`
	NoOfPassengers string      `json:"passengers,omitempty"`
	DistanceUnit   string      `json:"distance_unit,omitempty"`
	Legs           []LegOption `json:"legs,omitempty"`
}

// LegOption is a single leg of a flight.
type LegOption struct {
	DepartureAirport string `json:"departure_airport,omitempty"`
	ArrivalAirport   string `json:"destination_airport,omitempty"`
	CabinClass       string `json:"cabin_class,omitempty"`
}

// NewEmissionsOptions creates an emissionsOptions object with defaults.
func NewEmissionsOptions(opts EmissionsOptions) EmissionsOptions {
	if opts.EmissionType == "" {
		opts.EmissionType = "flight"
	}

	return opts
}
