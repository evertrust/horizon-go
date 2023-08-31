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
func (c *Client) CentralizedEnroll(profile string, password string, subject []IndexedDNElement, sans []ListSANElement, labels []LabelElement, keyType string, owner *string, team *string, contactEmail *string) (*HorizonRequest, error) {
	request, err := c.GetTemplate(profile)
	if err != nil {
		return nil, err
	}

	request.Template.KeyTypes = []string{keyType}
	request.Template.Subject = subject
	request.Template.Sans = sans
	request.Template.Labels = labels

	if contactEmail != nil {
		request.Template.ContactEmail = &HrzTemplateContactEmail{Value: *contactEmail}
	}
	if request.Template.Owner.Editable && owner != nil {
		request.Template.Owner.Value = owner
	}
	if request.Template.Team.Editable && team != nil {
		request.Template.Team.Value = team
	}

	request.Password.Value = password

	return c.Submit(*request)
}

// DecentralizedEnroll is a wrapper method around the Requests API that generates a
// decentralized enroll request given a profile, a CSR and a list of labels
func (c *Client) DecentralizedEnroll(profile string, csr []byte, labels []LabelElement, owner *string, team *string, contactEmail *string) (*HorizonRequest, error) {
	rfcClient := rfc5280.Client{
		Http: c.Http,
	}
	parsedCsr, err := rfcClient.Pkcs10(csr)
	if err != nil {
		return nil, err
	}

	request, err := c.GetTemplate(profile)
	if err != nil {
		return nil, err
	}

	// Translate the parsed certificate DN elements into the request elements
	var subject []IndexedDNElement
	var typeCounts = make(map[string]int)
	for _, dnElement := range parsedCsr.DnElements {
		typeCounts[dnElement.Type]++
		subject = append(subject, IndexedDNElement{
			Element: fmt.Sprintf("%s.%d", strings.ToLower(dnElement.Type), typeCounts[dnElement.Type]),
			Type:    dnElement.Type,
			Value:   fmt.Sprintf("%v", dnElement.Value),
		})
	}
	// Translate the parsed certificate SAN elements into the request elements
	var sans []ListSANElement
	for _, sanElement := range parsedCsr.Sans {
		isNew := true
		for _, indexedSan := range sans {
			if strings.ToUpper(indexedSan.Type) == strings.ToUpper(sanElement.SanType) {
				indexedSan.Value = append(indexedSan.Value, sanElement.Value)
				isNew = false
			}
		}
		if isNew {
			sans = append(sans, ListSANElement{
				Type:  strings.ToUpper(sanElement.SanType),
				Value: []string{sanElement.Value},
			})
		}
	}
	request.Template.Csr = parsedCsr.Pem
	request.Template.Subject = subject
	request.Template.Sans = sans
	request.Template.Labels = labels

	if contactEmail != nil {
		request.Template.ContactEmail = &HrzTemplateContactEmail{Value: *contactEmail}
	}
	if request.Template.Owner.Editable && owner != nil {
		request.Template.Owner.Value = owner
	}
	if request.Template.Team.Editable && team != nil {
		request.Template.Team.Value = team
	}

	return c.Submit(*request)
}

// DecentralizedRenew is a wrapper method around the Requests API that generates a
// renewal request given a profile and a certificate PEM
func (c *Client) DecentralizedRenew(csr []byte, lastCertificateId string) (*HorizonRequest, error) {
	request := HorizonRequest{
		Workflow:      RequestWorkflowRenew,
		Module:        "webra",
		CertificateId: lastCertificateId,
		Template:      CertificateTemplate{Csr: string(csr)},
	}

	return c.Submit(request)
}

// Revoke is a wrapper around the Requests API that generates a revocation request
// given a PEM-encoded certificate and a revocation reason.
func (c *Client) Revoke(certificatePem string, revocationReason certificates.RevocationReason) (*HorizonRequest, error) {
	return c.Submit(HorizonRequest{
		Workflow:       RequestWorkflowRevoke,
		Module:         "webra",
		CertificatePEM: certificatePem,
		Template:       CertificateTemplate{RevocationReason: revocationReason},
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
