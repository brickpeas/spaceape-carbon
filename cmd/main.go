package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/brickpeas/spaceape-carbon/flight"

	flags "github.com/brickpeas/spaceape-carbon/pkg/flags"
	client "github.com/brickpeas/spaceape-carbon/pkg/http"
)

func main() {
	// Read in api key and store this for use in the http client.
	apiKey := os.Getenv("API_KEY")

	//	Define our command line flags.
	var legsInput string
	flag.StringVar(&legsInput, "flight", "LHR-JFK-e,JFK-CDG-e,CDG-LHR-p", "Define your legs using IATA airport codes in the following format: LHR-JFK-e,JFK-CDG-p. The first code is the deparation airport, the second is the arrival airport and third represents cabin class, e=economy, p=premimum. Flight codes can be found here: https://www.iata.org/en/publications/directories/code-search/.")

	noOfPassengers := flag.String("passengers", "1", "Number of passengers")
	distanceUnit := flag.String("units", "km", "Distance unit")

	flag.Parse()

	parsedLegs, err := flags.ParseLegsFlag(legsInput)
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

	str, _ := json.MarshalIndent(emissions, "", " ")
	fmt.Println(string(str))
}
