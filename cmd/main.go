package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/brickpeas/spaceape-carbon/flight"

	client "github.com/brickpeas/spaceape-carbon/pkg/http"
)

func main() {
	// Read in api key and store this for use in the http client.
	apiKey := os.Getenv("API_KEY")

	//	Define our command line flags.
	var legsInput string
	flag.StringVar(&legsInput, "flight", "LHR-JFK-e", "Define your legs using IATA airport codes in the following format: LHR-JFK-e,JFK-CDG-p. The first code is the deparation airport, the second is the arrival airport and third represents cabin class, e=economy, p=premimum. Flight codes can be found here: https://www.iata.org/en/publications/directories/code-search/.")

	noOfPassengers := flag.String("passengers", "1", "Number of passengers")
	distanceUnit := flag.String("units", "km", "Distance unit")

	flag.Parse()

	parsedLegs, err := parseLegsFlag(legsInput)
	if err != nil {
		log.Fatalf("error parsing flags: %v", err)
	}

	// Create a new http client with options.
	options := client.ClientOptions{APIKey: apiKey}

	httpClient, err := client.NewClient(&options)
	if err != nil {
		log.Fatalf("error creating http client: %v", err)
	}

	emissionsOptions := flight.EmissionsOptions{
		NoOfPassengers: *noOfPassengers,
		Legs:           parsedLegs,
		DistanceUnit:   *distanceUnit,
	}

	emissionsOptions.FlightType = "flight"

	// Call the api to get the emissions for each flight.
	emissions, err := httpClient.GetEmissions(emissionsOptions)
	if err != nil {
		log.Fatalf("error getting emissions: %v", err)
	}

	// Print the emissions for each flight.
	for _, leg := range emissions.Data.Attributes.Legs {
		fmt.Printf("Flight from %s to %s:\n %+v", leg.DepartureAirport, leg.DestinationAirport, emissions.Data.Attributes)
	}

	// Print the total emissions for all flights.
	fmt.Printf("Total emissions: %d", emissions.Data.Attributes.CarbonG)
}

func parseLegsFlag(legsInput string) ([]flight.Leg, error) {
	var cabinClass = "economy"
	var legs []flight.Leg

	individualLegs := strings.Split(legsInput, ",")

	for _, leg := range individualLegs {
		legValues := strings.Split(leg, "-")
		if len(legValues) != 3 {
			return nil, fmt.Errorf("invalid leg: %s", leg)
		}

		if len(legValues[0]) != 3 || len(legValues[1]) != 3 {
			return nil, fmt.Errorf("invalid IATA code: %v", leg)
		}

		if legValues[2] != "e" && legValues[2] != "p" {
			return nil, fmt.Errorf("invalid cabin class: %v", leg[2])
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
