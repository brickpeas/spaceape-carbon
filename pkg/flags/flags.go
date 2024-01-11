package flags

import (
	"fmt"
	"strings"

	"github.com/brickpeas/spaceape-carbon/flight"
)

// ParseLegsFlag parses a comma separated list from the legs flag and returns a slice of LegOption.
func ParseLegsFlag(legsInput string) ([]flight.LegOption, error) {
	// Default to economy.
	var cabinClass = "economy"
	var legs []flight.LegOption

	individualLegs := strings.Split(legsInput, ",")

	for _, leg := range individualLegs {
		legValues := strings.Split(leg, "-")

		// Only expect 3 values.
		if len(legValues) != 3 {
			return nil, fmt.Errorf("invalid leg: %v", leg)
		}

		// IATA codes should be 3 characters.
		if len(legValues[0]) != 3 || len(legValues[1]) != 3 {
			return nil, fmt.Errorf("invalid IATA code: %v", leg)
		}

		// Cabin class should be e or p.
		if legValues[2] != "e" && legValues[2] != "p" {
			return nil, fmt.Errorf("invalid cabin class: %v", leg)
		}

		// Override the default cabin class if premium.
		if legValues[2] == "p" {
			cabinClass = "premium"
		}

		legs = append(legs, flight.LegOption{
			DepartureAirport: legValues[0],
			ArrivalAirport:   legValues[1],
			CabinClass:       cabinClass,
		})
	}

	return legs, nil
}
