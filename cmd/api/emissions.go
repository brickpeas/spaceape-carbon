package api

import "time"

type EmissionsResponse struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes `json:"attributes"`
	} `json:"data"`
}

type Attributes struct {
	Passengers    int       `json:"passengers"`
	Legs          []Leg     `json:"legs"`
	EstimatedAt   time.Time `json:"estimated_at"`
	CarbonG       int       `json:"carbon_g"`
	CarbonLb      int       `json:"carbon_lb"`
	CarbonKg      int       `json:"carbon_kg"`
	CarbonMt      int       `json:"carbon_mt"`
	DistanceUnit  string    `json:"distance_unit"`
	DistanceValue float64   `json:"distance_value"`
}

type Leg struct {
	DepartureAirport   string `json:"departure_airport"`
	DestinationAirport string `json:"destination_airport"`
}
