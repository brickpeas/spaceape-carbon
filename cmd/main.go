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
	flag.StringVar(&legsInput, "flight", "LHR-JFK-e,JFK-CDG-e,CDG-LHR-p", "Define your legs using IATA airport codes in the following format: LHR-JFK-e,JFK-CDG-p. The first code is the departure airport, the second is the arrival airport and third represents cabin class, e=economy, p=premimum. Flight codes can be found here: https://www.iata.org/en/publications/directories/code-search/.")

	noOfPassengers := flag.String("passengers", "1", "Number of passengers")
	distanceUnit := flag.String("units", "km", "Distance unit")

	flag.Parse()

	// Parse the legs flag.
	parsedLegs, err := flags.ParseLegsFlag(legsInput)
	if err != nil {
		log.Fatalf("error parsing flags: %v", err)
	}

	// Create a new client options object with the api key.
	options := client.ClientOptions{APIKey: apiKey}

	// Create a new http client with options.
	httpClient, err := client.NewClient(&options)
	if err != nil {
		log.Fatalf("error creating http client: %v", err)
	}

	// Create emissionsOptions object with desired values.
	emissionsOptions := flight.NewEmissionsOptions(flight.EmissionsOptions{
		NoOfPassengers: *noOfPassengers,
		Legs:           parsedLegs,
		DistanceUnit:   *distanceUnit,
	})

	// Call the api to get the emissions for each flight.
	emissions, err := httpClient.GetEmissions(emissionsOptions)
	if err != nil {
		log.Fatalf("error getting emissions: %v", err)
	}

	// Print the emissions data back to the user.
	str, _ := json.MarshalIndent(emissions, "", " ")
	fmt.Println(string(str))
}
