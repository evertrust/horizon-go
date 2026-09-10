/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the ThirdPartyScheduledTaskResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ThirdPartyScheduledTaskResponse{}

// ThirdPartyScheduledTaskResponse struct for ThirdPartyScheduledTaskResponse
type ThirdPartyScheduledTaskResponse struct {
	// Object internal ID
	Id                   string               `json:"_id"`
	Detail               *string              `json:"detail,omitempty"`
	Connector            string               `json:"connector"`
	Description          *string              `json:"description,omitempty"`
	DryRun               bool                 `json:"dryRun"`
	Enroll               bool                 `json:"enroll"`
	Module               string               `json:"module"`
	Profile              string               `json:"profile"`
	Renew                *bool                `json:"renew,omitempty"`
	Revoke               bool                 `json:"revoke"`
	Type                 string               `json:"type"`
	Cron                 string               `json:"cron"`
	Enabled              bool                 `json:"enabled"`
	ExecutionId          utils.NullableString `json:"executionId,omitempty"`
	Host                 utils.NullableString `json:"host,omitempty"`
	LastCompletionDate   utils.NullableInt64  `json:"lastCompletionDate,omitempty"`
	LastExecutionDate    utils.NullableInt64  `json:"lastExecutionDate,omitempty"`
	Name                 string               `json:"name"`
	Status               utils.NullableString `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThirdPartyScheduledTaskResponse ThirdPartyScheduledTaskResponse

// NewThirdPartyScheduledTaskResponse instantiates a new ThirdPartyScheduledTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyScheduledTaskResponse(id string, connector string, dryRun bool, enroll bool, module string, profile string, revoke bool, type_ string, cron string, enabled bool, name string) *ThirdPartyScheduledTaskResponse {
	this := ThirdPartyScheduledTaskResponse{}
	this.Id = id
	this.Connector = connector
	this.DryRun = dryRun
	this.Enroll = enroll
	this.Module = module
	this.Profile = profile
	this.Revoke = revoke
	this.Type = type_
	this.Cron = cron
	this.Enabled = enabled
	this.Name = name
	return &this
}

// NewThirdPartyScheduledTaskResponseWithDefaults instantiates a new ThirdPartyScheduledTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyScheduledTaskResponseWithDefaults() *ThirdPartyScheduledTaskResponse {
	this := ThirdPartyScheduledTaskResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ThirdPartyScheduledTaskResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyScheduledTaskResponse) SetId(v string) {
	o.Id = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ThirdPartyScheduledTaskResponse) GetDetail() string {
	if o == nil || utils.IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetDetailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTaskResponse) HasDetail() bool {
	if o != nil && !utils.IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ThirdPartyScheduledTaskResponse) SetDetail(v string) {
	o.Detail = &v
}

// GetConnector returns the Connector field value
func (o *ThirdPartyScheduledTaskResponse) GetConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connector, true
}

// SetConnector sets field value
func (o *ThirdPartyScheduledTaskResponse) SetConnector(v string) {
	o.Connector = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyScheduledTaskResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTaskResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyScheduledTaskResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value
func (o *ThirdPartyScheduledTaskResponse) GetDryRun() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DryRun, true
}

// SetDryRun sets field value
func (o *ThirdPartyScheduledTaskResponse) SetDryRun(v bool) {
	o.DryRun = v
}

// GetEnroll returns the Enroll field value
func (o *ThirdPartyScheduledTaskResponse) GetEnroll() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enroll
}

// GetEnrollOk returns a tuple with the Enroll field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetEnrollOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enroll, true
}

// SetEnroll sets field value
func (o *ThirdPartyScheduledTaskResponse) SetEnroll(v bool) {
	o.Enroll = v
}

// GetModule returns the Module field value
func (o *ThirdPartyScheduledTaskResponse) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *ThirdPartyScheduledTaskResponse) SetModule(v string) {
	o.Module = v
}

// GetProfile returns the Profile field value
func (o *ThirdPartyScheduledTaskResponse) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *ThirdPartyScheduledTaskResponse) SetProfile(v string) {
	o.Profile = v
}

// GetRenew returns the Renew field value if set, zero value otherwise.
func (o *ThirdPartyScheduledTaskResponse) GetRenew() bool {
	if o == nil || utils.IsNil(o.Renew) {
		var ret bool
		return ret
	}
	return *o.Renew
}

// GetRenewOk returns a tuple with the Renew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetRenewOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Renew) {
		return nil, false
	}
	return o.Renew, true
}

// HasRenew returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTaskResponse) HasRenew() bool {
	if o != nil && !utils.IsNil(o.Renew) {
		return true
	}

	return false
}

// SetRenew gets a reference to the given bool and assigns it to the Renew field.
func (o *ThirdPartyScheduledTaskResponse) SetRenew(v bool) {
	o.Renew = &v
}

// GetRevoke returns the Revoke field value
func (o *ThirdPartyScheduledTaskResponse) GetRevoke() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Revoke
}

// GetRevokeOk returns a tuple with the Revoke field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetRevokeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revoke, true
}

// SetRevoke sets field value
func (o *ThirdPartyScheduledTaskResponse) SetRevoke(v bool) {
	o.Revoke = v
}

// GetType returns the Type field value
func (o *ThirdPartyScheduledTaskResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ThirdPartyScheduledTaskResponse) SetType(v string) {
	o.Type = v
}

// GetCron returns the Cron field value
func (o *ThirdPartyScheduledTaskResponse) GetCron() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cron
}

// GetCronOk returns a tuple with the Cron field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetCronOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cron, true
}

// SetCron sets field value
func (o *ThirdPartyScheduledTaskResponse) SetCron(v string) {
	o.Cron = v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyScheduledTaskResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyScheduledTaskResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTaskResponse) GetExecutionId() string {
	if o == nil || utils.IsNil(o.ExecutionId.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionId.Get()
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTaskResponse) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionId.Get(), o.ExecutionId.IsSet()
}

// HasExecutionId returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTaskResponse) HasExecutionId() bool {
	if o != nil && o.ExecutionId.IsSet() {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given NullableString and assigns it to the ExecutionId field.
func (o *ThirdPartyScheduledTaskResponse) SetExecutionId(v string) {
	o.ExecutionId.Set(&v)
}

// SetExecutionIdNil sets the value for ExecutionId to be an explicit nil
func (o *ThirdPartyScheduledTaskResponse) SetExecutionIdNil() {
	o.ExecutionId.Set(nil)
}

// UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
func (o *ThirdPartyScheduledTaskResponse) UnsetExecutionId() {
	o.ExecutionId.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTaskResponse) GetHost() string {
	if o == nil || utils.IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTaskResponse) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTaskResponse) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *ThirdPartyScheduledTaskResponse) SetHost(v string) {
	o.Host.Set(&v)
}

// SetHostNil sets the value for Host to be an explicit nil
func (o *ThirdPartyScheduledTaskResponse) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *ThirdPartyScheduledTaskResponse) UnsetHost() {
	o.Host.Unset()
}

// GetLastCompletionDate returns the LastCompletionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTaskResponse) GetLastCompletionDate() int64 {
	if o == nil || utils.IsNil(o.LastCompletionDate.Get()) {
		var ret int64
		return ret
	}
	return *o.LastCompletionDate.Get()
}

// GetLastCompletionDateOk returns a tuple with the LastCompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTaskResponse) GetLastCompletionDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastCompletionDate.Get(), o.LastCompletionDate.IsSet()
}

// HasLastCompletionDate returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTaskResponse) HasLastCompletionDate() bool {
	if o != nil && o.LastCompletionDate.IsSet() {
		return true
	}

	return false
}

// SetLastCompletionDate gets a reference to the given NullableInt64 and assigns it to the LastCompletionDate field.
func (o *ThirdPartyScheduledTaskResponse) SetLastCompletionDate(v int64) {
	o.LastCompletionDate.Set(&v)
}

// SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil
func (o *ThirdPartyScheduledTaskResponse) SetLastCompletionDateNil() {
	o.LastCompletionDate.Set(nil)
}

// UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
func (o *ThirdPartyScheduledTaskResponse) UnsetLastCompletionDate() {
	o.LastCompletionDate.Unset()
}

// GetLastExecutionDate returns the LastExecutionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTaskResponse) GetLastExecutionDate() int64 {
	if o == nil || utils.IsNil(o.LastExecutionDate.Get()) {
		var ret int64
		return ret
	}
	return *o.LastExecutionDate.Get()
}

// GetLastExecutionDateOk returns a tuple with the LastExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTaskResponse) GetLastExecutionDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExecutionDate.Get(), o.LastExecutionDate.IsSet()
}

// HasLastExecutionDate returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTaskResponse) HasLastExecutionDate() bool {
	if o != nil && o.LastExecutionDate.IsSet() {
		return true
	}

	return false
}

// SetLastExecutionDate gets a reference to the given NullableInt64 and assigns it to the LastExecutionDate field.
func (o *ThirdPartyScheduledTaskResponse) SetLastExecutionDate(v int64) {
	o.LastExecutionDate.Set(&v)
}

// SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil
func (o *ThirdPartyScheduledTaskResponse) SetLastExecutionDateNil() {
	o.LastExecutionDate.Set(nil)
}

// UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
func (o *ThirdPartyScheduledTaskResponse) UnsetLastExecutionDate() {
	o.LastExecutionDate.Unset()
}

// GetName returns the Name field value
func (o *ThirdPartyScheduledTaskResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTaskResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ThirdPartyScheduledTaskResponse) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTaskResponse) GetStatus() string {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTaskResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTaskResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ThirdPartyScheduledTaskResponse) SetStatus(v string) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *ThirdPartyScheduledTaskResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ThirdPartyScheduledTaskResponse) UnsetStatus() {
	o.Status.Unset()
}

func (o ThirdPartyScheduledTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyScheduledTaskResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if !utils.IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	toSerialize["connector"] = o.Connector
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["dryRun"] = o.DryRun
	toSerialize["enroll"] = o.Enroll
	toSerialize["module"] = o.Module
	toSerialize["profile"] = o.Profile
	if !utils.IsNil(o.Renew) {
		toSerialize["renew"] = o.Renew
	}
	toSerialize["revoke"] = o.Revoke
	toSerialize["type"] = o.Type
	toSerialize["cron"] = o.Cron
	toSerialize["enabled"] = o.Enabled
	if o.ExecutionId.IsSet() {
		toSerialize["executionId"] = o.ExecutionId.Get()
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.LastCompletionDate.IsSet() {
		toSerialize["lastCompletionDate"] = o.LastCompletionDate.Get()
	}
	if o.LastExecutionDate.IsSet() {
		toSerialize["lastExecutionDate"] = o.LastExecutionDate.Get()
	}
	toSerialize["name"] = o.Name
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThirdPartyScheduledTaskResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"connector",
		"dryRun",
		"enroll",
		"module",
		"profile",
		"revoke",
		"type",
		"cron",
		"enabled",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varThirdPartyScheduledTaskResponse := _ThirdPartyScheduledTaskResponse{}

	err = json.Unmarshal(data, &varThirdPartyScheduledTaskResponse)

	if err != nil {
		return err
	}

	*o = ThirdPartyScheduledTaskResponse(varThirdPartyScheduledTaskResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "connector")
		delete(additionalProperties, "description")
		delete(additionalProperties, "dryRun")
		delete(additionalProperties, "enroll")
		delete(additionalProperties, "module")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "renew")
		delete(additionalProperties, "revoke")
		delete(additionalProperties, "type")
		delete(additionalProperties, "cron")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "executionId")
		delete(additionalProperties, "host")
		delete(additionalProperties, "lastCompletionDate")
		delete(additionalProperties, "lastExecutionDate")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThirdPartyScheduledTaskResponse struct {
	value *ThirdPartyScheduledTaskResponse
	isSet bool
}

func (v NullableThirdPartyScheduledTaskResponse) Get() *ThirdPartyScheduledTaskResponse {
	return v.value
}

func (v *NullableThirdPartyScheduledTaskResponse) Set(val *ThirdPartyScheduledTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyScheduledTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyScheduledTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyScheduledTaskResponse(val *ThirdPartyScheduledTaskResponse) *NullableThirdPartyScheduledTaskResponse {
	return &NullableThirdPartyScheduledTaskResponse{value: val, isSet: true}
}

func (v NullableThirdPartyScheduledTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyScheduledTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
