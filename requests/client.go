// Package requests provides utilities to interact with the Horizon api.requests APIs.
package requests

import (
	"encoding/json"
	"errors"
	"github.com/evertrust/horizon-go/http"
	"github.com/evertrust/horizon-go/types"
)

type Client struct {
	http *http.Client
}

func Init(http *http.Client) *Client {
	return &Client{http: http}
}

var InvalidTypeError = errors.New("invalid response type")

// Search sends back paginated results
func (c *Client) Search(query types.RequestSearchQuery) (*types.SearchResults[types.RequestSearchResult], error) {
	jsonData, _ := json.Marshal(query)
	response, err := c.http.Post("/api/v1/requests/search", jsonData)
	if err != nil {
		return nil, err
	}
	var resultPage types.SearchResults[types.RequestSearchResult]
	err = response.Json().Decode(&resultPage)
	return &resultPage, err
}

// WebRA Enroll

func (c *Client) GetEnrollTemplate(request types.WebRAEnrollTemplateParams) (*types.WebRAEnrollTemplate, error) {
	// Merge params in struct
	enrollRequest := types.WebRAEnrollRequest{
		Profile:        request.Profile,
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       &types.WebRAEnrollTemplate{Csr: request.Csr},
		Module:         types.WebRA,
		Workflow:       types.Enroll}
	response, err := c.GetTemplate(&enrollRequest)
	if err != nil {
		return nil, err
	}
	webRAEnrollRequest, ok := response.(*types.WebRAEnrollRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAEnrollRequest.Template, nil
}

func (c *Client) GetEnrollRequest(id string) (*types.WebRAEnrollRequest, error) {
	var webRAEnrollRequest types.WebRAEnrollRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRAEnrollRequest)
	if err != nil {
		return nil, err
	}
	return &webRAEnrollRequest, nil
}

func (c *Client) CancelEnrollRequest(id string) (*types.WebRAEnrollRequest, error) {
	webRAEnrollRequest := types.WebRAEnrollRequest{
		Module:   types.WebRA,
		Workflow: types.Enroll,
		Id:       id,
	}
	err := c.CancelRequest(&webRAEnrollRequest)
	if err != nil {
		return nil, err
	}
	return &webRAEnrollRequest, nil
}

func (c *Client) NewEnrollRequest(request types.WebRAEnrollRequestParams) (*types.WebRAEnrollRequest, error) {
	// Merge params in struct
	var password *types.Secret
	if request.Password != "" {
		password = new(types.Secret)
		password.Value = request.Password
	}
	enrollRequest := types.WebRAEnrollRequest{
		Profile:  request.Profile,
		Template: request.Template,
		Module:   types.WebRA,
		Workflow: types.Enroll,
		Password: password,
	}
	response, err := c.NewRequest(&enrollRequest)
	if err != nil {
		return nil, err
	}
	webRAEnrollRequest, ok := response.(*types.WebRAEnrollRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAEnrollRequest, nil
}

// SCEP Challenge

func (c *Client) GetScepChallengeTemplate(request types.ScepChallengeTemplateParams) (*types.ScepChallengeTemplate, error) {
	// Merge params in struct
	challengeRequest := types.ScepChallengeRequest{
		Profile:  request.Profile,
		Module:   types.Scep,
		Workflow: types.Enroll}
	response, err := c.GetTemplate(&challengeRequest)
	if err != nil {
		return nil, err
	}
	scepChallengeRequest, ok := response.(*types.ScepChallengeRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return scepChallengeRequest.Template, nil
}

func (c *Client) GetScepChallengeRequest(id string) (*types.ScepChallengeRequest, error) {
	var scepChallengeRequest types.ScepChallengeRequest
	// Merge params in struct
	err := c.GetRequest(id, &scepChallengeRequest)
	if err != nil {
		return nil, err
	}
	return &scepChallengeRequest, nil
}

func (c *Client) NewScepChallengeRequest(request types.ScepChallengeRequestParams) (*types.ScepChallengeRequest, error) {
	challengeRequest := types.ScepChallengeRequest{
		Profile:  request.Profile,
		Template: request.Template,
		Module:   types.Scep,
		Workflow: types.Enroll,
		Dn:       request.Dn,
	}
	response, err := c.NewRequest(&challengeRequest)
	if err != nil {
		return nil, err
	}
	scepChallengeRequest, ok := response.(*types.ScepChallengeRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return scepChallengeRequest, nil
}

// EST Challenge

func (c *Client) GetEstChallengeTemplate(request types.EstChallengeTemplateParams) (*types.EstChallengeTemplate, error) {
	// Merge params in struct
	challengeRequest := types.EstChallengeRequest{
		Profile:  request.Profile,
		Module:   types.Est,
		Workflow: types.Enroll}
	response, err := c.GetTemplate(&challengeRequest)
	if err != nil {
		return nil, err
	}
	estChallengeRequest, ok := response.(*types.EstChallengeRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return estChallengeRequest.Template, nil
}

func (c *Client) GetEstChallengeRequest(id string) (*types.EstChallengeRequest, error) {
	var estChallengeRequest types.EstChallengeRequest
	// Merge params in struct
	err := c.GetRequest(id, &estChallengeRequest)
	if err != nil {
		return nil, err
	}
	return &estChallengeRequest, nil
}

func (c *Client) NewEstChallengeRequest(request types.EstChallengeRequestParams) (*types.EstChallengeRequest, error) {
	challengeRequest := types.EstChallengeRequest{
		Profile:  request.Profile,
		Template: request.Template,
		Module:   types.Est,
		Workflow: types.Enroll,
		Dn:       request.Dn,
	}
	response, err := c.NewRequest(&challengeRequest)
	if err != nil {
		return nil, err
	}
	estChallengeRequest, ok := response.(*types.EstChallengeRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return estChallengeRequest, nil
}

// WebRA Renew

func (c *Client) GetRenewTemplate(request types.WebRARenewTemplateParams) (*types.WebRARenewTemplate, error) {
	// Merge params in struct
	renewRequest := types.WebRARenewRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Module:         types.WebRA,
		Workflow:       types.Renew}
	response, err := c.GetTemplate(&renewRequest)
	if err != nil {
		return nil, err
	}
	webRARenewRequest, ok := response.(*types.WebRARenewRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRARenewRequest.Template, nil
}

func (c *Client) GetRenewRequest(id string) (*types.WebRARenewRequest, error) {
	var webRARenewRequest types.WebRARenewRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRARenewRequest)
	if err != nil {
		return nil, err
	}
	return &webRARenewRequest, nil
}

func (c *Client) CancelRenewRequest(id string) (*types.WebRARenewRequest, error) {
	webRARenewRequest := types.WebRARenewRequest{
		Module:   types.WebRA,
		Workflow: types.Renew,
		Id:       id,
	}
	err := c.CancelRequest(&webRARenewRequest)
	if err != nil {
		return nil, err
	}
	return &webRARenewRequest, nil
}

func (c *Client) NewRenewRequest(request types.WebRARenewRequestParams) (*types.WebRARenewRequest, error) {
	// Merge params in struct
	var password *types.Secret
	if request.Password != "" {
		password = new(types.Secret)
		password.Value = request.Password
	}
	renewRequest := types.WebRARenewRequest{
		Template: request.Template,
		Module:   types.WebRA,
		Workflow: types.Renew,
		Password: password,
	}
	response, err := c.NewRequest(&renewRequest)
	if err != nil {
		return nil, err
	}
	webRARenewRequest, ok := response.(*types.WebRARenewRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRARenewRequest, nil
}

// WebRA Revoke

func (c *Client) GetRevokeRequest(id string) (*types.WebRARevokeRequest, error) {
	var webRARevokeRequest types.WebRARevokeRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRARevokeRequest)
	if err != nil {
		return nil, err
	}
	return &webRARevokeRequest, nil
}

func (c *Client) NewRevokeRequest(request types.WebRARevokeRequestParams) (*types.WebRARevokeRequest, error) {
	// Merge params in struct
	revokeRequest := types.WebRARevokeRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       &types.WebRARevokeTemplate{RevocationReason: request.RevocationReason},
		Module:         types.WebRA,
		Workflow:       types.Revoke,
	}
	response, err := c.NewRequest(&revokeRequest)
	if err != nil {
		return nil, err
	}
	webRARevokeRequest, ok := response.(*types.WebRARevokeRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRARevokeRequest, nil
}

// WebRA Update

func (c *Client) GetUpdateTemplate(request types.WebRAUpdateTemplateParams) (*types.WebRAUpdateTemplate, error) {
	// Merge params in struct
	updateRequest := types.WebRAUpdateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Module:         types.WebRA,
		Workflow:       types.Update}
	response, err := c.GetTemplate(&updateRequest)
	if err != nil {
		return nil, err
	}
	webRAUpdateRequest, ok := response.(*types.WebRAUpdateRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAUpdateRequest.Template, nil
}

func (c *Client) GetUpdateRequest(id string) (*types.WebRAUpdateRequest, error) {
	var webRAUpdateRequest types.WebRAUpdateRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRAUpdateRequest)
	if err != nil {
		return nil, err
	}
	return &webRAUpdateRequest, nil
}

func (c *Client) NewUpdateRequest(request types.WebRAUpdateRequestParams) (*types.WebRAUpdateRequest, error) {
	updateRequest := types.WebRAUpdateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       request.Template,
		Module:         types.WebRA,
		Workflow:       types.Update,
	}
	response, err := c.NewRequest(&updateRequest)
	if err != nil {
		return nil, err
	}
	webRAUpdateRequest, ok := response.(*types.WebRAUpdateRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAUpdateRequest, nil
}

// WebRA Migrate

func (c *Client) GetMigrateTemplate(request types.WebRAMigrateTemplateParams) (*types.WebRAMigrateTemplate, error) {
	// Merge params in struct
	migrateRequest := types.WebRAMigrateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Profile:        request.Profile,
		Module:         types.WebRA,
		Workflow:       types.Migrate}
	response, err := c.GetTemplate(&migrateRequest)
	if err != nil {
		return nil, err
	}
	webRAMigrateRequest, ok := response.(*types.WebRAMigrateRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAMigrateRequest.Template, nil
}

func (c *Client) GetMigrateRequest(id string) (*types.WebRAMigrateRequest, error) {
	var webRAMigrateRequest types.WebRAMigrateRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRAMigrateRequest)
	if err != nil {
		return nil, err
	}
	return &webRAMigrateRequest, nil
}

func (c *Client) NewMigrateRequest(request types.WebRAMigrateRequestParams) (*types.WebRAMigrateRequest, error) {
	migrateRequest := types.WebRAMigrateRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Template:       request.Template,
		Profile:        request.Profile,
		Module:         types.WebRA,
		Workflow:       types.Migrate,
	}
	response, err := c.NewRequest(&migrateRequest)
	if err != nil {
		return nil, err
	}
	webRAMigrateRequest, ok := response.(*types.WebRAMigrateRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAMigrateRequest, nil
}

// WebRA Import

func (c *Client) GetImportTemplate(request types.WebRAImportTemplateParams) (*types.WebRAImportTemplate, error) {
	// Merge params in struct
	importRequest := types.WebRAImportRequest{
		Profile:        request.Profile,
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Module:         types.WebRA,
		Workflow:       types.Import,
	}
	response, err := c.GetTemplate(&importRequest)
	if err != nil {
		return nil, err
	}
	webRAImportRequest, ok := response.(*types.WebRAImportRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAImportRequest.Template, nil
}

func (c *Client) GetImportRequest(id string) (*types.WebRAImportRequest, error) {
	var webRAImportRequest types.WebRAImportRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRAImportRequest)
	if err != nil {
		return nil, err
	}
	return &webRAImportRequest, nil
}

func (c *Client) NewImportRequest(request types.WebRAImportRequestParams) (*types.WebRAImportRequest, error) {
	// Merge params in struct
	importRequest := types.WebRAImportRequest{
		Profile:  request.Profile,
		Template: request.Template,
		Module:   types.WebRA,
		Workflow: types.Import,
	}
	response, err := c.NewRequest(&importRequest)
	if err != nil {
		return nil, err
	}
	webRAImportRequest, ok := response.(*types.WebRAImportRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRAImportRequest, nil
}

// WebRA Recover

func (c *Client) GetRecoverRequest(id string) (*types.WebRARecoverRequest, error) {
	var webRARecoverRequest types.WebRARecoverRequest
	// Merge params in struct
	err := c.GetRequest(id, &webRARecoverRequest)
	if err != nil {
		return nil, err
	}
	return &webRARecoverRequest, nil
}

func (c *Client) NewRecoverRequest(request types.WebRARecoverRequestParams) (*types.WebRARecoverRequest, error) {
	var password *types.Secret
	if request.Password != "" {
		password = new(types.Secret)
		password.Value = request.Password
	}
	recoverRequest := types.WebRARecoverRequest{
		CertificateId:  request.CertificateId,
		CertificatePEM: request.CertificatePEM,
		Contact:        request.Contact,
		Password:       password,
		Module:         types.WebRA,
		Workflow:       types.Recover,
	}
	response, err := c.NewRequest(&recoverRequest)
	if err != nil {
		return nil, err
	}
	webRARecoverRequest, ok := response.(*types.WebRARecoverRequest)
	if !ok {
		return nil, InvalidTypeError
	}
	return webRARecoverRequest, nil
}

// Low level functions

func (c *Client) NewRequest(request types.Request) (types.Request, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	response, err := c.http.Post("/api/v1/requests/submit", jsonData)
	if err != nil {
		return nil, err
	}
	err = response.Json().Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, request.EnsureType()
}

func (c *Client) GetTemplate(request types.Request) (types.Request, error) {
	jsonData, _ := json.Marshal(request)
	response, err := c.http.Post("/api/v1/requests/template", jsonData)
	if err != nil {
		return nil, err
	}
	err = response.Json().Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, request.EnsureType()
}

func (c *Client) GetRequest(id string, result types.Request) error {
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
func (c *Client) CancelRequest(request types.Request) error {
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
