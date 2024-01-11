package flags

import (
	"testing"

	"github.com/brickpeas/spaceape-carbon/flight"
	"github.com/stretchr/testify/assert"
)

func TestParseLegsFlag(t *testing.T) {
	t.Parallel()

	t.Run("leg has more than 3 values", func(t *testing.T) {
		leg, err := ParseLegsFlag("LHR-JFK-CDG-e")

		assert.Nil(t, leg)
		assert.Equal(t, "invalid leg: LHR-JFK-CDG-e", err.Error())
	})
	t.Run("IATA code is invalid", func(t *testing.T) {
		leg, err := ParseLegsFlag("LHRG-JFK-e")

		assert.Nil(t, leg)
		assert.Equal(t, "invalid IATA code: LHRG-JFK-e", err.Error())
	})
	t.Run("cabin class is invalid", func(t *testing.T) {
		leg, err := ParseLegsFlag("LHR-JFK-q")

		assert.Nil(t, leg)
		assert.Equal(t, "invalid cabin class: LHR-JFK-q", err.Error())
	})
	t.Run("leg is formatted correctly with one leg", func(t *testing.T) {

		leg, err := ParseLegsFlag("LHR-JFK-e")

		expectedLeg := []flight.LegOption{
			{
				DepartureAirport: "LHR",
				ArrivalAirport:   "JFK",
				CabinClass:       "economy",
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedLeg, leg)
	})
	t.Run("legs are formatted correctly with multiple leg", func(t *testing.T) {

		leg, err := ParseLegsFlag("LHR-JFK-e,JFK-CDG-e,CDG-LHR-p")

		expectedLegs := []flight.LegOption{
			{
				DepartureAirport: "LHR",
				ArrivalAirport:   "JFK",
				CabinClass:       "economy",
			},
			{
				DepartureAirport: "JFK",
				ArrivalAirport:   "CDG",
				CabinClass:       "economy",
			},
			{
				DepartureAirport: "CDG",
				ArrivalAirport:   "LHR",
				CabinClass:       "premium",
			},
		}

		assert.NoError(t, err)
		assert.Equal(t, expectedLegs, leg)
	})
}
