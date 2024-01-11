# spaceape-carbon
A CLI tool used to estimate the carbon footprint of flights.

## How to run
The application can be ran with the following command:
```shell
API_KEY=<XXXXX> go run cmd/main.go -flight="LHR-JFK-e,JFK-CDG-e,CDG-LHR-p" -passengers=1 -units=km
```
Where you will use your own API key.

The format of the flight flag is in the following format: <departure>-<destination>-<cabin_class>.
The departure and destination should be the IATA 3 letter airport codes, e.g. LHR for london heathrow. The 3-letter codes can be found here: https://www.iata.org/en/publications/directories/code-search/.
The final character represents cabin class which can be e (economy) or p (premium).

The output from should like something like this where you can see the different values.
```json
{
 "data": {
  "id": "8478250c-0996-495e-8ca7-947d4e9a1a91",
  "type": "estimate",
  "attributes": {
   "passengers": 1,
   "legs": [
    {
     "departure_airport": "LHR",
     "destination_airport": "JFK",
     "cabin_class": "economy"
    },
    {
     "departure_airport": "JFK",
     "destination_airport": "CDG",
     "cabin_class": "economy"
    },
    {
     "departure_airport": "CDG",
     "destination_airport": "LHR",
     "cabin_class": "premium"
    }
   ],
   "distance_value": 12011.75,
   "distance_unit": "km",
   "estimated_at": "2024-01-11T15:28:08.653Z",
   "carbon_g": 1700149,
   "carbon_lb": 3748.19,
   "carbon_kg": 1700.15,
   "carbon_mt": 1.7
  }
 }
}
```

## Notes
### Testing
I would ideally like to have the GetEmissions() function unit tested further as this isn't tested at the moment. Testing some of the http functions would require a little more work around mocking, perhaps using a dependency injection pattern where an interface could be defined with methods that could be mocked in the tests.