package automation

import (
	"errors"
	"fmt"
	"github.com/evertrust/horizon-go"
	"github.com/evertrust/horizon-go/http"
)

type Client struct {
	http *http.Client
}

func Init(http *http.Client) *Client {
	return &Client{http: http}
}

func (c *Client) List() ([]horizon.Policy, error) {
	response, err := c.http.Get("/api/v1/automation/policies")
	if err != nil {
		return nil, err
	}
	var policies []horizon.Policy
	err = response.Json().Decode(&policies)
	if err != nil {
		return nil, err
	}

	return policies, nil
}

func (c *Client) Get(name string) (*horizon.Policy, error) {
	response, err := c.http.Get("/api/v1/automation/policies/" + name)
	if err != nil {
		return nil, err
	}
	var policy horizon.Policy
	err = response.Json().Decode(&policy)
	if err != nil {
		return nil, err
	}
	return &policy, nil
}

func (c *Client) GetParameters(policyName string) (horizon.InitParameters, error) {
	response, err := c.http.Get("/api/v1/automation/lifecycle/" + policyName)
	if err != nil {
		return nil, err
	}
	var policy horizon.InitParameters
	err = response.Json().Decode(&policy)
	switch policy.GetModule() {
	case horizon.Acme:
		var acmeParams horizon.AcmeInitParameters
		err = response.Json().Decode(&acmeParams)
		return &acmeParams, err
	case horizon.AcmeExternal:
		var acmeExternalParams horizon.AcmeExternalInitParameters
		err = response.Json().Decode(&acmeExternalParams)
		return &acmeExternalParams, err
	case horizon.Est:
		var estParams horizon.EstInitParameters
		err = response.Json().Decode(&estParams)
		return &estParams, err
	case horizon.Scep:
		var scepParams horizon.ScepInitParameters
		err = response.Json().Decode(&scepParams)
		return &scepParams, err
	}
	return nil, fmt.Errorf("unknown module '%s'", policy.GetModule())
}

// TODO: change this
// CheckCertificate checks the compliance of the certificate in the jwt against the automation policy
// It returns isCompliant, isRunnable (can be run now), enroll (if true, an enrollment can be performed), renew (if true, a renewal can be performed), and error
func (c *Client) Check(policyName string) (bool, bool, bool, bool, error) {
	response, err := c.http.Get("/api/v1/automation/lifecycle/" + policyName + "/verify")
	if err != nil {
		return false, false, false, false, err
	}
	switch response.HttpResponse.StatusCode {
	case 204:
		// Certificate is compliant, nothing to do
		return true, false, false, false, nil
	case 200:
		// Certificate is not compliant, check if runnable or not, and the reason
		var automationReport horizon.Report
		err = response.Json().Decode(&automationReport)
		if err != nil {
			return false, false, false, false, err
		}
		// Handle backwardsCompatibility
		if automationReport.IsRenewable == nil {
			// 2.4 version
			// Behavior expects to be enroll and renew
			return false, automationReport.IsRunnable, true, true, nil
		}
		return false, automationReport.IsRunnable, !*automationReport.IsRenewable, *automationReport.IsRenewable, nil
	default:
		// Horizon should not send error code - this should have been catched before
		return false, false, false, false, errors.New("unexpected error")
	}
}
