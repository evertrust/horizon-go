// Package webra provides utilities to use Horizon api protocol WebRA.
package requests

import (
	"encoding/json"
	"errors"
	"github.com/evertrust/horizon-go/http"
)

type Client struct {
	Http *http.Client
}

var InvalidTypeError = errors.New("invalid response type")

// WebRA Enroll

func (c *Client) GetEnrollTemplate(request WebRAEnrollTemplateParams) (*WebRAEnrollTemplate, error) {
	// Merge params in struct
	enrollRequest := WebRAEnrollRequest{
		Profile:        request.Profile,
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       &WebRAEnrollTemplate{Csr: request.Csr},
		Module:         WebRA,
		Workflow:       Enroll}
	response, err := c.GetTemplate(&enrollRequest)
	if err != nil {
		return nil, err
	}
	webRAEnrollRequest, ok := response.(*WebRAEnrollRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAEnrollRequest.Template, nil
}

func (c *Client) GetEnrollRequest(id string) (*WebRAEnrollRequest, error) {
	var webRAEnrollRequest WebRAEnrollRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRAEnrollRequest)
	if err != nil {
		return nil, err
	}
	return &webRAEnrollRequest, nil
}

func (c *Client) NewEnrollRequest(request WebRAEnrollRequestParams) (*WebRAEnrollRequest, error) {
	// Merge params in struct
	var password *Secret
	if request.Password != "" {
		password = new(Secret)
		password.Value = request.Password
	}
	enrollRequest := WebRAEnrollRequest{
		Profile:  request.Profile,
		Template: request.Template,
		Module:   WebRA,
		Workflow: Enroll,
		Password: password,
	}
	response, err := c.NewRequest(&enrollRequest)
	if err != nil {
		return nil, err
	}
	webRAEnrollRequest, ok := response.(*WebRAEnrollRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAEnrollRequest, nil
}

// SCEP Challenge

func (c *Client) GetScepChallengeTemplate(request ScepChallengeTemplateParams) (*ScepChallengeTemplate, error) {
	// Merge params in struct
	challengeRequest := ScepChallengeRequest{
		Profile:  request.Profile,
		Module:   Scep,
		Workflow: Enroll}
	response, err := c.GetTemplate(&challengeRequest)
	if err != nil {
		return nil, err
	}
	scepChallengeRequest, ok := response.(*ScepChallengeRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return scepChallengeRequest.Template, nil
}

func (c *Client) GetScepChallengeRequest(id string) (*ScepChallengeRequest, error) {
	var scepChallengeRequest ScepChallengeRequest
	// Merge params in struct
	err := c.GetRequest(id, &scepChallengeRequest)
	if err != nil {
		return nil, err
	}
	return &scepChallengeRequest, nil
}

func (c *Client) NewScepChallengeRequest(request ScepChallengeRequestParams) (*ScepChallengeRequest, error) {
	challengeRequest := ScepChallengeRequest{
		Profile:  request.Profile,
		Template: request.Template,
		Module:   Scep,
		Workflow: Enroll,
		Dn:       request.Dn,
	}
	response, err := c.NewRequest(&challengeRequest)
	if err != nil {
		return nil, err
	}
	scepChallengeRequest, ok := response.(*ScepChallengeRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return scepChallengeRequest, nil
}

// EST Challenge

func (c *Client) GetEstChallengeTemplate(request EstChallengeTemplateParams) (*EstChallengeTemplate, error) {
	// Merge params in struct
	challengeRequest := EstChallengeRequest{
		Profile:  request.Profile,
		Module:   Est,
		Workflow: Enroll}
	response, err := c.GetTemplate(&challengeRequest)
	if err != nil {
		return nil, err
	}
	estChallengeRequest, ok := response.(*EstChallengeRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return estChallengeRequest.Template, nil
}

func (c *Client) GetEstChallengeRequest(id string) (*EstChallengeRequest, error) {
	var estChallengeRequest EstChallengeRequest
	// Merge params in struct
	err := c.GetRequest(id, &estChallengeRequest)
	if err != nil {
		return nil, err
	}
	return &estChallengeRequest, nil
}

func (c *Client) NewEstChallengeRequest(request EstChallengeRequestParams) (*EstChallengeRequest, error) {
	challengeRequest := EstChallengeRequest{
		Profile:  request.Profile,
		Template: request.Template,
		Module:   Est,
		Workflow: Enroll,
		Dn:       request.Dn,
	}
	response, err := c.NewRequest(&challengeRequest)
	if err != nil {
		return nil, err
	}
	estChallengeRequest, ok := response.(*EstChallengeRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return estChallengeRequest, nil
}

// WebRA Renew

func (c *Client) GetRenewTemplate(request WebRARenewTemplateParams) (*WebRARenewTemplate, error) {
	// Merge params in struct
	renewRequest := WebRARenewRequest{Profile: request.Profile, CertificateId: request.CertificateId, CertificatePEM: request.CertificatePEM, Module: WebRA, Workflow: Renew}
	response, err := c.GetTemplate(&renewRequest)
	if err != nil {
		return nil, err
	}
	webRARenewRequest, ok := response.(*WebRARenewRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRARenewRequest.Template, nil
}

func (c *Client) GetRenewRequest(id string) (*WebRARenewRequest, error) {
	var webRARenewRequest WebRARenewRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRARenewRequest)
	if err != nil {
		return nil, err
	}
	return &webRARenewRequest, nil
}

func (c *Client) NewRenewRequest(request WebRARenewRequestParams) (*WebRARenewRequest, error) {
	// Merge params in struct
	var password *Secret
	if request.Password != "" {
		password = new(Secret)
		password.Value = request.Password
	}
	renewRequest := WebRARenewRequest{
		Profile:  request.Profile,
		Template: request.Template,
		Module:   WebRA,
		Workflow: Renew,
		Password: password,
	}
	response, err := c.NewRequest(&renewRequest)
	if err != nil {
		return nil, err
	}
	webRARenewRequest, ok := response.(*WebRARenewRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRARenewRequest, nil
}

// WebRA Revoke

func (c *Client) GetRevokeRequest(id string) (*WebRARevokeRequest, error) {
	var webRARevokeRequest WebRARevokeRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRARevokeRequest)
	if err != nil {
		return nil, err
	}
	return &webRARevokeRequest, nil
}

func (c *Client) NewRevokeRequest(request WebRARevokeRequestParams) (*WebRARevokeRequest, error) {
	// Merge params in struct
	revokeRequest := WebRARevokeRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       &WebRARevokeTemplate{RevocationReason: request.RevocationReason},
		Module:         WebRA,
		Workflow:       Revoke,
	}
	response, err := c.NewRequest(&revokeRequest)
	if err != nil {
		return nil, err
	}
	webRARevokeRequest, ok := response.(*WebRARevokeRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRARevokeRequest, nil
}

// WebRA Update

func (c *Client) GetUpdateTemplate(request WebRAUpdateTemplateParams) (*WebRAUpdateTemplate, error) {
	// Merge params in struct
	updateRequest := WebRAUpdateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Module:         WebRA,
		Workflow:       Update}
	response, err := c.GetTemplate(&updateRequest)
	if err != nil {
		return nil, err
	}
	webRAUpdateRequest, ok := response.(*WebRAUpdateRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAUpdateRequest.Template, nil
}

func (c *Client) GetUpdateRequest(id string) (*WebRAUpdateRequest, error) {
	var webRAUpdateRequest WebRAUpdateRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRAUpdateRequest)
	if err != nil {
		return nil, err
	}
	return &webRAUpdateRequest, nil
}

func (c *Client) NewUpdateRequest(request WebRAUpdateRequestParams) (*WebRAUpdateRequest, error) {
	updateRequest := WebRAUpdateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       request.Template,
		Module:         WebRA,
		Workflow:       Update,
	}
	response, err := c.NewRequest(&updateRequest)
	if err != nil {
		return nil, err
	}
	webRAUpdateRequest, ok := response.(*WebRAUpdateRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAUpdateRequest, nil
}

// WebRA Migrate

func (c *Client) GetMigrateTemplate(request WebRAMigrateTemplateParams) (*WebRAMigrateTemplate, error) {
	// Merge params in struct
	migrateRequest := WebRAMigrateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Profile:        request.Profile,
		Module:         WebRA,
		Workflow:       Migrate}
	response, err := c.GetTemplate(&migrateRequest)
	if err != nil {
		return nil, err
	}
	webRAMigrateRequest, ok := response.(*WebRAMigrateRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAMigrateRequest.Template, nil
}

func (c *Client) GetMigrateRequest(id string) (*WebRAMigrateRequest, error) {
	var webRAMigrateRequest WebRAMigrateRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRAMigrateRequest)
	if err != nil {
		return nil, err
	}
	return &webRAMigrateRequest, nil
}

func (c *Client) NewMigrateRequest(request WebRAMigrateRequestParams) (*WebRAMigrateRequest, error) {
	migrateRequest := WebRAMigrateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       request.Template,
		Profile:        request.Profile,
		Module:         WebRA,
		Workflow:       Migrate,
	}
	response, err := c.NewRequest(&migrateRequest)
	if err != nil {
		return nil, err
	}
	webRAMigrateRequest, ok := response.(*WebRAMigrateRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAMigrateRequest, nil
}

// WebRA Recover

func (c *Client) GetRecoverRequest(id string) (*WebRARecoverRequest, error) {
	var webRARecoverRequest WebRARecoverRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRARecoverRequest)
	if err != nil {
		return nil, err
	}
	return &webRARecoverRequest, nil
}

func (c *Client) NewRecoverRequest(request WebRARecoverRequestParams) (*WebRARecoverRequest, error) {
	var password *Secret
	if request.Password != "" {
		password = new(Secret)
		password.Value = request.Password
	}
	recoverRequest := WebRARecoverRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Contact:        request.Contact,
		Password:       password,
		Module:         WebRA,
		Workflow:       Recover,
	}
	response, err := c.NewRequest(&recoverRequest)
	if err != nil {
		return nil, err
	}
	webRARecoverRequest, ok := response.(*WebRARecoverRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRARecoverRequest, nil
}

// Low level functions

func (c *Client) NewRequest(request Request) (Request, error) {
	jsonData, _ := json.Marshal(request)
	response, err := c.Http.Post("/api/v1/requests/submit", jsonData)
	if err != nil {
		return nil, err
	}
	err = response.Json().Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, request.EnsureType()
}

func (c *Client) GetTemplate(request Request) (Request, error) {
	jsonData, _ := json.Marshal(request)
	response, err := c.Http.Post("/api/v1/requests/template", jsonData)
	if err != nil {
		return nil, err
	}
	err = response.Json().Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, request.EnsureType()
}

func (c *Client) GetRequest(id string, result Request) error {
	response, err := c.Http.Get("/api/v1/requests/" + id)
	if err != nil {
		return err
	}
	err = response.Json().Decode(&result)
	if err != nil {
		return err
	}
	// Ensuring that the type received is the expected one
	return result.EnsureType()
}
