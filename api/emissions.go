package api

import "time"

type EmissionsResponse struct {
	Data struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Passengers int `json:"passengers"`
			Legs       []struct {
				DepartureAirport   string `json:"departure_airport"`
				DestinationAirport string `json:"destination_airport"`
				CabinClass         string `json:"cabin_class"`
			} `json:"legs"`
			DistanceValue float64   `json:"distance_value"`
			DistanceUnit  string    `json:"distance_unit"`
			EstimatedAt   time.Time `json:"estimated_at"`
			CarbonG       int       `json:"carbon_g"`
			CarbonLb      float64   `json:"carbon_lb"`
			CarbonKg      float64   `json:"carbon_kg"`
			CarbonMt      float64   `json:"carbon_mt"`
		} `json:"attributes"`
	} `json:"data"`
}
