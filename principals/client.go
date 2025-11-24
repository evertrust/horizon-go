// Package principals provides utilities to interact with the Horizon api.principals APIs.
package principals

import (
	"github.com/evertrust/horizon-go"
	"github.com/evertrust/horizon-go/http"
)

type Client struct {
	http *http.Client
}

func Init(http *http.Client) *Client {
	return &Client{http: http}
}

func (c *Client) Self() (*horizon.Principal, error) {
	response, err := c.http.Get("/api/v1/security/principals/self")
	if err != nil {
		return nil, err
	}
	var result horizon.InternalPrincipal
	var out horizon.Principal
	if response.HttpResponse.StatusCode == 204 {
		return nil, nil
	}
	err = response.Json().Decode(&result)
	if err != nil {
		return nil, err
	}
	out.Identity = result.Identity
	out.Permissions = result.Permissions
	out.Roles = result.Roles
	out.Teams = result.Teams
	return &out, err
}
