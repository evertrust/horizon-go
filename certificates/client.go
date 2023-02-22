// Package certificates provides utilities to interact with the Horizon api.certificate APIs.
package certificates

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/http"
)

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

func (c *Client) Search(query string, pageIndex int, withCount bool) ([]SearchResult, bool, int, error) {
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
	var searchResponse SearchResponse
	err = response.Json().Decode(&searchResponse)
	if err != nil {
		return nil, false, 0, err
	}

	return searchResponse.Results, searchResponse.HasMore, searchResponse.Count, nil
}

func (c *Client) UpdateMigrate(id, profile, owner, team, labelName, labelValue string) error {
	template := HorizonRequestTemplate{}
	if owner == "unset" {
		template.Owner = &HorizonRequestValue{}
	} else if owner != "" {
		template.Owner = &HorizonRequestValue{Value: owner}
	} else {
		template.Owner = nil
	}
	if team == "unset" {
		template.Team = &HorizonRequestValue{}
	} else if team != "" {
		template.Team = &HorizonRequestValue{Value: team}
	} else {
		template.Team = nil
	}
	if labelName != "" {
		template.Labels = []HorizonRequestValue{{Label: labelName, Value: labelValue}}
	}
	hrzRequest := HorizonRequest{
		CertificateId: id,
		Workflow:      "update",
		Template:      template,
	}
	if profile != "" {
		hrzRequest.Workflow = "migrate"
		hrzRequest.Profile = profile
	}
	jsonBody, err := json.Marshal(hrzRequest)
	if err != nil {
		return err
	}
	_, err = c.Http.Post("/api/v1/requests/submit", jsonBody)
	return err
}
