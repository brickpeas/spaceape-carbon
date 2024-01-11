package api

import "time"

// EmissionsResponse is the response from the CarbonInterface estimates API.
type EmissionsResponse struct {
	Data struct {
		ID         string              `json:"id"`
		Type       string              `json:"type"`
		Attributes EmissionsAttributes `json:"attributes"`
	} `json:"data"`
}

// EmissionsAttributes are the attributes of the EmissionsResponse.
type EmissionsAttributes struct {
	Passengers    int            `json:"passengers"`
	Legs          []EmissionsLeg `json:"legs"`
	DistanceValue float64        `json:"distance_value"`
	DistanceUnit  string         `json:"distance_unit"`
	EstimatedAt   time.Time      `json:"estimated_at"`
	CarbonG       int            `json:"carbon_g"`
	CarbonLb      float64        `json:"carbon_lb"`
	CarbonKg      float64        `json:"carbon_kg"`
	CarbonMt      float64        `json:"carbon_mt"`
}

// EmissionsLeg is a single leg of a flight.
type EmissionsLeg struct {
	DepartureAirport   string `json:"departure_airport"`
	DestinationAirport string `json:"destination_airport"`
	CabinClass         string `json:"cabin_class"`
}
