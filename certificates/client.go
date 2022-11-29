// Package certificates provides utilities to interact with the Horizon api.certificate APIs.
package certificates

import (
	"github.com/evertrust/horizon-go/http"
)

type certificateResponse struct {
	Certificate Certificate `json:"certificate"`
}

type Client struct {
	Http *http.Client
}

func (c *Client) Get(id string) (*Certificate, error) {
	response, err := c.Http.Get("/api/v1/certificates/" + id)
	if err != nil {
		return nil, err
	}
	var certificate certificateResponse
	err = response.Json().Decode(&certificate)
	if err != nil {
		return nil, err
	}

	return &certificate.Certificate, nil
}
