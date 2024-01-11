package flags

import (
	"fmt"
	"github.com/brickpeas/spaceape-carbon/flight"
	"strings"
)

func ParseLegsFlag(legsInput string) ([]flight.Leg, error) {
	var cabinClass = "economy"
	var legs []flight.Leg

	individualLegs := strings.Split(legsInput, ",")

	for _, leg := range individualLegs {
		legValues := strings.Split(leg, "-")
		if len(legValues) != 3 {
			return nil, fmt.Errorf("invalid leg: %v", leg)
		}

		if len(legValues[0]) != 3 || len(legValues[1]) != 3 {
			return nil, fmt.Errorf("invalid IATA code: %v", leg)
		}

		if legValues[2] != "e" && legValues[2] != "p" {
			return nil, fmt.Errorf("invalid cabin class: %v", leg)
		}

		if legValues[2] == "p" {
			cabinClass = "premium"
		}

		legs = append(legs, flight.Leg{
			DepartureAirport: legValues[0],
			ArrivalAirport:   legValues[1],
			CabinClass:       cabinClass,
		})
	}

	return legs, nil
}
