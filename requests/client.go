// Package requests provides utilities to interact with the Horizon api.requests APIs.
package requests

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/evertrust/horizon-go/certificates"
	"github.com/evertrust/horizon-go/http"
	"github.com/evertrust/horizon-go/rfc5280"
)

type Client struct {
	Http *http.Client
}

func (c *Client) Submit(request HorizonRequest) (*HorizonRequest, error) {
	jsonData, _ := json.Marshal(request)
	response, err := c.Http.Post("/api/v1/requests/submit", jsonData)
	if err != nil {
		return nil, err
	}
	err = response.Json().Decode(&request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func (c *Client) Get(id string) (*HorizonRequest, error) {
	var request HorizonRequest
	response, err := c.Http.Get("/api/v1/requests/" + id)
	if err != nil {
		return nil, err
	}

	err = response.Json().Decode(&request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

// CentralizedEnroll is a wrapper method around the Requests API that generates a
// centralized enroll request given a profile, DN and SAN elements and a list of labels
func (c *Client) CentralizedEnroll(profile string, password string, subject []IndexedDNElement, sans []IndexedSANElement, labels []LabelElement, keyType string, owner *string, team *string) (*HorizonRequest, error) {
	template := WebRARequestTemplate{
		Subject:  subject,
		Sans:     sans,
		Labels:   labels,
		KeyTypes: []string{keyType},
	}

	if owner != nil {
		template.Owner = &CertificateOwner{
			Value:    *owner,
			Editable: false,
		}
	}

	if team != nil {
		template.Team = &CertificateTeam{
			Value:    *team,
			Editable: false,
		}
	}

	return c.Submit(HorizonRequest{
		Workflow: RequestWorkflowEnroll,
		Profile:  profile,
		Module:   "webra",
		Template: template,
		Password: P12Password{
			Value: password,
		},
	})
}

// DecentralizedEnroll is a wrapper method around the Requests API that generates a
// decentralized enroll request given a profile, a CSR and a list of labels
func (c *Client) DecentralizedEnroll(profile string, csr []byte, labels []LabelElement, owner *string, team *string) (*HorizonRequest, error) {
	rfcClient := rfc5280.Client{
		Http: c.Http,
	}
	parsedCsr, err := rfcClient.Pkcs10(csr)
	if err != nil {
		return nil, err
	}
	var typeCounts = make(map[string]int)

	// Translate the parsed certificate DN elements into the request elements
	var subject []IndexedDNElement
	for _, dnElement := range parsedCsr.DnElements {
		typeCounts[dnElement.Type]++
		subject = append(subject, IndexedDNElement{
			Element: fmt.Sprintf("%s.%d", strings.ToLower(dnElement.Type), typeCounts[dnElement.Type]),
			Type:    dnElement.Type,
			Value:   fmt.Sprintf("%v", dnElement.Value),
		})
	}
	// Translate the parsed certificate SAN elements into the request elements
	var sans []IndexedSANElement
	for _, sanElement := range parsedCsr.Sans {
		typeCounts[sanElement.SanType]++
		sans = append(sans, IndexedSANElement{
			Element: fmt.Sprintf("%s.%d", strings.ToLower(sanElement.SanType), typeCounts[sanElement.SanType]),
			Type:    sanElement.SanType,
			Value:   fmt.Sprintf("%v", sanElement.Value),
		})
	}
	template := WebRARequestTemplate{
		Csr:     parsedCsr.Pem,
		Subject: subject,
		Sans:    sans,
		Labels:  labels,
	}
	if owner != nil {
		template.Owner = &CertificateOwner{
			Value:    *owner,
			Editable: false,
		}
	}
	if team != nil {
		template.Team = &CertificateTeam{
			Value:    *team,
			Editable: false,
		}
	}
	return c.Submit(HorizonRequest{
		Workflow: RequestWorkflowEnroll,
		Profile:  profile,
		Module:   "webra",
		Template: template,
	})
}

// Revoke is a wrapper around the Requests API that generates a revocation request
// given a PEM-encoded certificate and a revocation reason.
func (c *Client) Revoke(certificatePem string, revocationReason certificates.RevocationReason) (*HorizonRequest, error) {
	return c.Submit(HorizonRequest{
		Workflow:       RequestWorkflowRevoke,
		Module:         "webra",
		CertificatePEM: certificatePem,
		Template:       WebRARevokeTemplate{RevocationReason: revocationReason},
	})
}

// GetTemplate returns the template information for a given WebRA profile
func (c *Client) GetTemplate(profile string) (*HorizonRequest, error) {
	var request HorizonRequest
	body := map[string]string{
		"module":   "webra",
		"profile":  profile,
		"workflow": "enroll",
	}
	marshalledBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	response, err := c.Http.Post("/api/v1/requests/template", marshalledBody)
	if err != nil {
		return nil, err
	}

	err = response.Json().Decode(&request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}
