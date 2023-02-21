package automation

import (
	"github.com/evertrust/horizon-go/http"
)

type Client struct {
	Http *http.Client
}

func (c *Client) Get(policyName string) (Policy, error) {
	response, err := c.Http.Get("/api/v1/automation/lifecycle/" + policyName)
	if err != nil {
		return Policy{}, err
	}
	var policy Policy
	err = response.Json().Decode(&policy)
	if err != nil {
		return Policy{}, err
	}

	return policy, nil
}

func (c *Client) List() ([]Policy, error) {
	response, err := c.Http.Get("/api/v1/automation/policies")
	if err != nil {
		return nil, err
	}
	var policies []Policy
	err = response.Json().Decode(&policies)
	if err != nil {
		return nil, err
	}

	return policies, nil
}

func (c *Client) Check(jwt, policyName string) (bool, error) {
	response, err := c.Http.PostWithJwt("/api/v1/automation/lifecycle/"+policyName+"/check", jwt, nil)
	if err != nil {
		return false, err
	}

	if response.BaseResponse.StatusCode == 204 {
		return true, nil
	}
	return false, nil
}
