package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/brickpeas/spaceape-carbon/api"
	"github.com/brickpeas/spaceape-carbon/flight"
)

const (
	defaultBaseURL = "https://www.carboninterface.com/api/v1"
	defaultTimeout = time.Second * 10
)

// Client is the client for the CarbonInterface API.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// ClientOptions are a set of overrides used when creating a new Client.
type ClientOptions struct {
	// APIKey required. Generated at https://www.carboninterface.com/account/api_credentials.
	APIKey string
	// HTTPClient can be used to set a custom HTTP client on the CarbonInterface client.
	HTTPClient *http.Client
}

// NewClient creates a new CarbonInterface client with the provided options.
func NewClient(options *ClientOptions) (*Client, error) {
	if strings.TrimSpace(options.APIKey) == "" {
		return nil, fmt.Errorf("APIKey is required")
	}

	client := &Client{
		apiKey:  options.APIKey,
		baseURL: defaultBaseURL,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}

	return client, nil
}

// GetEmissions makes a request to the CarbonInterface estimates API.
func (c *Client) GetEmissions(emissionsOpts flight.EmissionsOptions) (*api.EmissionsResponse, error) {
	emissionsJSON, err := json.Marshal(emissionsOpts)
	if err != nil {
		return nil, fmt.Errorf("error marshalling emissions options: %w", err)
	}

	url := fmt.Sprintf("%s/%s", c.baseURL, "estimates")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(emissionsJSON))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set our API key and content type headers.
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	// Check for non 2XX response codes and return CarbonInterface API to user.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		errBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error reading response body: %w", err)
		}

		return nil, fmt.Errorf("non 2XX response code received: '%v' with error: '%s'", resp.StatusCode, string(errBody))
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Check for empty response body.
	if len(body) == 0 {
		return nil, fmt.Errorf("empty response body")
	}

	var emissionsResponse *api.EmissionsResponse
	if err := json.Unmarshal(body, &emissionsResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}

	// Check for nil response.
	if emissionsResponse == nil {
		return nil, fmt.Errorf("nil response after unmarshalling")
	}

	return emissionsResponse, nil
}
