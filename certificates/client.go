// Package certificates provides utilities to interact with the Horizon api.certificate APIs.
package certificates

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/http"
)

type certificateResponse struct {
	Certificate Certificate `json:"certificate"`
}

type searchResponse struct {
	Results   []Certificate `json:"results"`
	PageIndex int           `json:"pageIndex"`
	PageSize  int           `json:"pageSize"`
	Count     int           `json:"count"`
	HasMore   bool          `json:"hasMore"`
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

func (c *Client) Search(query string, pageIndex int, withCount bool) ([]Certificate, bool, int, error) {
	hrzRequest := HrzSearchQuery{
		Query:     query,
		PageIndex: pageIndex,
		PageSize:  10,
		Fields:    []string{"dn", "certificate", "serial", "profile"},
		WithCount: withCount,
	}
	jsonBody, err := json.Marshal(hrzRequest)
	if err != nil {
		return nil, false, 0, err
	}
	response, err := c.Http.Post("/api/v1/certificates/search", jsonBody)
	if err != nil {
		return nil, false, 0, err
	}
	var searchResponse searchResponse
	err = response.Json().Decode(&searchResponse)
	if err != nil {
		return nil, false, 0, err
	}

	return searchResponse.Results, searchResponse.HasMore, searchResponse.Count, nil
}
