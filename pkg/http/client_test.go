package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	options := &ClientOptions{
		APIKey: "test-apikey",
	}

	t.Run("api key is missing", func(t *testing.T) {
		options.APIKey = ""

		client, err := NewClient(options)
		if err != nil {
			assert.Equal(t, "APIKey is required", err.Error())
		}
		assert.Nil(t, client)
	})
}

func TestGetEmissions(t *testing.T) {
	t.Skip("TODO")
}
