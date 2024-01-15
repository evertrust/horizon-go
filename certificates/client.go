// Package certificates provides utilities to interact with the Horizon api.certificate APIs.
package certificates

import (
	"encoding/json"
	"github.com/evertrust/horizon-go/http"
	"github.com/evertrust/horizon-go/types"
)

type Client struct {
	http *http.Client
}

func Init(http *http.Client) *Client {
	return &Client{http: http}
}

func (c *Client) Get(id string) (*types.Certificate, error) {
	response, err := c.http.Get("/api/v1/certificates/" + id)
	if err != nil {
		return nil, err
	}
	var certificate types.CertificateResponse
	err = response.Json().Decode(&certificate)
	if err != nil {
		return nil, err
	}

	return &certificate.Certificate, nil
}

// Search sends back paginated results
func (c *Client) Search(query types.CertificateSearchQuery) (*types.SearchResults[types.CertificateSearchResult], error) {
	jsonData, _ := json.Marshal(query)
	response, err := c.http.Post("/api/v1/certificates/search", jsonData)
	if err != nil {
		return nil, err
	}
	var resultPage types.SearchResults[types.CertificateSearchResult]
	err = response.Json().Decode(&resultPage)
	return &resultPage, err
}
