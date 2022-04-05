// Package license provides utilities to interact with the Horizon api.license APIs.
package license

import (
	"github.com/evertrust/horizon-go/http"
)

type Client struct {
	Http *http.Client
}

func (c *Client) Get() (*LicenseInfo, error) {
	response, err := c.Http.Get("/api/v1/licenses")
	if err != nil {
		return nil, err
	}
	var license LicenseInfo
	err = response.Json().Decode(&license)
	if err != nil {
		return nil, err
	}

	return &license, nil
}
