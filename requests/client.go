// Package requests provides utilities to interact with the Horizon api.requests APIs.
package requests

import (
	"encoding/json"
	"errors"
	"github.com/evertrust/horizon-go"
	"github.com/evertrust/horizon-go/http"
)

type Client struct {
	http *http.Client
}

func Init(http *http.Client) *Client {
	return &Client{http: http}
}

var InvalidTypeError = errors.New("invalid response type")

// Search sends back paginated results
func (c *Client) Search(query horizon.RequestSearchQuery) (*horizon.SearchResults[horizon.RequestSearchResult], error) {
	jsonData, _ := json.Marshal(query)
	response, err := c.http.Post("/api/v1/requests/search", jsonData)
	if err != nil {
		return nil, err
	}
	var resultPage horizon.SearchResults[horizon.RequestSearchResult]
	err = response.Json().Decode(&resultPage)
	return &resultPage, err
}

// WebRA Enroll

func (c *Client) GetEnrollTemplate(request horizon.WebRAEnrollTemplateParams) (*horizon.WebRAEnrollTemplate, error) {
	// Merge params in struct
	enrollRequest := horizon.WebRAEnrollRequest{
		Profile:        request.Profile,
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       &horizon.WebRAEnrollTemplate{Csr: request.Csr},
		Module:         horizon.WebRA,
		Workflow:       horizon.Enroll}
	err := c.GetTemplate(&enrollRequest)
	if err != nil {
		return nil, err
	}
	return enrollRequest.Template, nil
}

func (c *Client) GetEnrollRequest(id string) (*horizon.WebRAEnrollRequest, error) {
	var webRAEnrollRequest horizon.WebRAEnrollRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRAEnrollRequest)
	if err != nil {
		return nil, err
	}
	return &webRAEnrollRequest, nil
}

func (c *Client) CancelEnrollRequest(id string) (*horizon.WebRAEnrollRequest, error) {
	webRAEnrollRequest := horizon.WebRAEnrollRequest{
		Module:   horizon.WebRA,
		Workflow: horizon.Enroll,
		Id:       id,
	}
	err := c.CancelRequest(&webRAEnrollRequest)
	if err != nil {
		return nil, err
	}
	return &webRAEnrollRequest, nil
}

func (c *Client) NewEnrollRequest(request horizon.WebRAEnrollRequestParams) (*horizon.WebRAEnrollRequest, error) {
	// Merge params in struct
	var password *horizon.Secret
	if request.Password != "" {
		password = new(horizon.Secret)
		password.Value = request.Password
	}
	enrollRequest := horizon.WebRAEnrollRequest{
		Profile:  request.Profile,
		Template: request.Template,
		Module:   horizon.WebRA,
		Workflow: horizon.Enroll,
		Password: password,
	}
	err := c.NewRequest(&enrollRequest)
	if err != nil {
		return nil, err
	}
	return &enrollRequest, nil
}

// SCEP Challenge

func (c *Client) GetScepChallengeTemplate(request horizon.ScepChallengeTemplateParams) (*horizon.ScepChallengeTemplate, error) {
	// Merge params in struct
	challengeRequest := horizon.ScepChallengeRequest{
		Profile:  request.Profile,
		Module:   horizon.Scep,
		Workflow: horizon.Enroll}
	err := c.GetTemplate(&challengeRequest)
	if err != nil {
		return nil, err
	}
	return challengeRequest.Template, nil
}

func (c *Client) GetScepChallengeRequest(id string) (*horizon.ScepChallengeRequest, error) {
	var scepChallengeRequest horizon.ScepChallengeRequest
	// Merge params in struct
	err := c.GetRequest(id, &scepChallengeRequest)
	if err != nil {
		return nil, err
	}
	return &scepChallengeRequest, nil
}

func (c *Client) NewScepChallengeRequest(request horizon.ScepChallengeRequestParams) (*horizon.ScepChallengeRequest, error) {
	challengeRequest := horizon.ScepChallengeRequest{
		Profile:  request.Profile,
		Template: request.Template,
		Module:   horizon.Scep,
		Workflow: horizon.Enroll,
		Dn:       request.Dn,
	}
	err := c.NewRequest(&challengeRequest)
	if err != nil {
		return nil, err
	}
	return &challengeRequest, nil
}

// EST Challenge

func (c *Client) GetEstChallengeTemplate(request horizon.EstChallengeTemplateParams) (*horizon.EstChallengeTemplate, error) {
	// Merge params in struct
	challengeRequest := horizon.EstChallengeRequest{
		Profile:  request.Profile,
		Module:   horizon.Est,
		Workflow: horizon.Enroll}
	err := c.GetTemplate(&challengeRequest)
	if err != nil {
		return nil, err
	}
	return challengeRequest.Template, nil
}

func (c *Client) GetEstChallengeRequest(id string) (*horizon.EstChallengeRequest, error) {
	var estChallengeRequest horizon.EstChallengeRequest
	// Merge params in struct
	err := c.GetRequest(id, &estChallengeRequest)
	if err != nil {
		return nil, err
	}
	return &estChallengeRequest, nil
}

func (c *Client) NewEstChallengeRequest(request horizon.EstChallengeRequestParams) (*horizon.EstChallengeRequest, error) {
	challengeRequest := horizon.EstChallengeRequest{
		Profile:  request.Profile,
		Template: request.Template,
		Module:   horizon.Est,
		Workflow: horizon.Enroll,
		Dn:       request.Dn,
	}
	err := c.NewRequest(&challengeRequest)
	if err != nil {
		return nil, err
	}
	return &challengeRequest, nil
}

// WebRA Renew

func (c *Client) GetRenewTemplate(request horizon.WebRARenewTemplateParams) (*horizon.WebRARenewTemplate, error) {
	// Merge params in struct
	renewRequest := horizon.WebRARenewRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Module:         horizon.WebRA,
		Workflow:       horizon.Renew}
	err := c.GetTemplate(&renewRequest)
	if err != nil {
		return nil, err
	}
	return renewRequest.Template, nil
}

func (c *Client) GetRenewRequest(id string) (*horizon.WebRARenewRequest, error) {
	var webRARenewRequest horizon.WebRARenewRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRARenewRequest)
	if err != nil {
		return nil, err
	}
	return &webRARenewRequest, nil
}

func (c *Client) CancelRenewRequest(id string) (*horizon.WebRARenewRequest, error) {
	webRARenewRequest := horizon.WebRARenewRequest{
		Module:   horizon.WebRA,
		Workflow: horizon.Renew,
		Id:       id,
	}
	err := c.CancelRequest(&webRARenewRequest)
	if err != nil {
		return nil, err
	}
	return &webRARenewRequest, nil
}

func (c *Client) NewRenewRequest(request horizon.WebRARenewRequestParams) (*horizon.WebRARenewRequest, error) {
	// Merge params in struct
	var password *horizon.Secret
	if request.Password != "" {
		password = new(horizon.Secret)
		password.Value = request.Password
	}
	renewRequest := horizon.WebRARenewRequest{
		Template: request.Template,
		Module:   horizon.WebRA,
		Workflow: horizon.Renew,
		Password: password,
	}
	err := c.NewRequest(&renewRequest)
	if err != nil {
		return nil, err
	}
	return &renewRequest, nil
}

// WebRA Revoke

func (c *Client) GetRevokeRequest(id string) (*horizon.WebRARevokeRequest, error) {
	var webRARevokeRequest horizon.WebRARevokeRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRARevokeRequest)
	if err != nil {
		return nil, err
	}
	return &webRARevokeRequest, nil
}

func (c *Client) NewRevokeRequest(request horizon.WebRARevokeRequestParams) (*horizon.WebRARevokeRequest, error) {
	// Merge params in struct
	revokeRequest := horizon.WebRARevokeRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       &horizon.WebRARevokeTemplate{RevocationReason: request.RevocationReason},
		Module:         horizon.WebRA,
		Workflow:       horizon.Revoke,
	}
	err := c.NewRequest(&revokeRequest)
	if err != nil {
		return nil, err
	}
	return &revokeRequest, nil
}

// WebRA Update

func (c *Client) GetUpdateTemplate(request horizon.WebRAUpdateTemplateParams) (*horizon.WebRAUpdateTemplate, error) {
	// Merge params in struct
	updateRequest := horizon.WebRAUpdateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Module:         horizon.WebRA,
		Workflow:       horizon.Update}
	err := c.GetTemplate(&updateRequest)
	if err != nil {
		return nil, err
	}
	return updateRequest.Template, nil
}

func (c *Client) GetUpdateRequest(id string) (*horizon.WebRAUpdateRequest, error) {
	var webRAUpdateRequest horizon.WebRAUpdateRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRAUpdateRequest)
	if err != nil {
		return nil, err
	}
	return &webRAUpdateRequest, nil
}

func (c *Client) NewUpdateRequest(request horizon.WebRAUpdateRequestParams) (*horizon.WebRAUpdateRequest, error) {
	updateRequest := horizon.WebRAUpdateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       request.Template,
		Module:         horizon.WebRA,
		Workflow:       horizon.Update,
	}
	err := c.NewRequest(&updateRequest)
	if err != nil {
		return nil, err
	}
	return &updateRequest, nil
}

// WebRA Migrate

func (c *Client) GetMigrateTemplate(request horizon.WebRAMigrateTemplateParams) (*horizon.WebRAMigrateTemplate, error) {
	// Merge params in struct
	migrateRequest := horizon.WebRAMigrateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Profile:        request.Profile,
		Module:         horizon.WebRA,
		Workflow:       horizon.Migrate}
	err := c.GetTemplate(&migrateRequest)
	if err != nil {
		return nil, err
	}
	return migrateRequest.Template, nil
}

func (c *Client) GetMigrateRequest(id string) (*horizon.WebRAMigrateRequest, error) {
	var webRAMigrateRequest horizon.WebRAMigrateRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRAMigrateRequest)
	if err != nil {
		return nil, err
	}
	return &webRAMigrateRequest, nil
}

func (c *Client) NewMigrateRequest(request horizon.WebRAMigrateRequestParams) (*horizon.WebRAMigrateRequest, error) {
	migrateRequest := horizon.WebRAMigrateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       request.Template,
		Profile:        request.Profile,
		Module:         horizon.WebRA,
		Workflow:       horizon.Migrate,
	}
	err := c.NewRequest(&migrateRequest)
	if err != nil {
		return nil, err
	}
	return &migrateRequest, nil
}

// WebRA Import

func (c *Client) GetImportTemplate(request horizon.WebRAImportTemplateParams) (*horizon.WebRAImportTemplate, error) {
	// Merge params in struct
	importRequest := horizon.WebRAImportRequest{
		Profile:        request.Profile,
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Module:         horizon.WebRA,
		Workflow:       horizon.Import,
	}
	err := c.GetTemplate(&importRequest)
	if err != nil {
		return nil, err
	}
	return importRequest.Template, nil
}

func (c *Client) GetImportRequest(id string) (*horizon.WebRAImportRequest, error) {
	var webRAImportRequest horizon.WebRAImportRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRAImportRequest)
	if err != nil {
		return nil, err
	}
	return &webRAImportRequest, nil
}

func (c *Client) NewImportRequest(request horizon.WebRAImportRequestParams) (*horizon.WebRAImportRequest, error) {
	// Merge params in struct
	importRequest := horizon.WebRAImportRequest{
		Profile:        request.Profile,
		Template:       request.Template,
		CertificatePEM: request.CertificatePEM,
		CertificateId:  request.CertificateId,
		Module:         horizon.WebRA,
		Workflow:       horizon.Import,
	}
	err := c.NewRequest(&importRequest)
	if err != nil {
		return nil, err
	}
	return &importRequest, nil
}

// WebRA Recover

func (c *Client) GetRecoverRequest(id string) (*horizon.WebRARecoverRequest, error) {
	var webRARecoverRequest horizon.WebRARecoverRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRARecoverRequest)
	if err != nil {
		return nil, err
	}
	return &webRARecoverRequest, nil
}

func (c *Client) NewRecoverRequest(request horizon.WebRARecoverRequestParams) (*horizon.WebRARecoverRequest, error) {
	var password *horizon.Secret
	if request.Password != "" {
		password = new(horizon.Secret)
		password.Value = request.Password
	}
	recoverRequest := horizon.WebRARecoverRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Contact:        request.Contact,
		Password:       password,
		Module:         horizon.WebRA,
		Workflow:       horizon.Recover,
	}
	err := c.NewRequest(&recoverRequest)
	if err != nil {
		return nil, err
	}
	return &recoverRequest, nil
}

// Low level functions

func (c *Client) NewRequest(request horizon.Request) error {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response, err := c.http.Post("/api/v1/requests/submit", jsonData)
	if err != nil {
		return err
	}
	err = response.Json().Decode(&request)
	if err != nil {
		return err
	}
	return request.EnsureType()
}

func (c *Client) GetTemplate(request horizon.Request) error {
	jsonData, _ := json.Marshal(request)
	response, err := c.http.Post("/api/v1/requests/template", jsonData)
	if err != nil {
		return err
	}
	err = response.Json().Decode(&request)
	if err != nil {
		return err
	}
	return request.EnsureType()
}

func (c *Client) GetRequest(id string, result horizon.Request) error {
	response, err := c.http.Get("/api/v1/requests/" + id)
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

// Request operations
func (c *Client) CancelRequest(request horizon.Request) error {
	jsonData, _ := json.Marshal(request)
	response, err := c.http.Post("/api/v1/requests/cancel", jsonData)
	if err != nil {
		return err
	}
	err = response.Json().Decode(&request)
	if err != nil {
		return err
	}
	return request.EnsureType()
}
