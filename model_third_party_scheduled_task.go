/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"encoding/json"
	"fmt"
)

// checks if the ThirdPartyScheduledTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyScheduledTask{}

// ThirdPartyScheduledTask struct for ThirdPartyScheduledTask
type ThirdPartyScheduledTask struct {
	Type string `json:"type"`
	Cron string `json:"cron"`
	Host NullableString `json:"host,omitempty"`
	Status NullableString `json:"status,omitempty"`
	LastExecutionDate NullableInt64 `json:"lastExecutionDate,omitempty"`
	LastCompletionDate NullableInt64 `json:"lastCompletionDate,omitempty"`
	Detail NullableString `json:"detail,omitempty"`
	DryRun bool `json:"dryRun"`
	ExecutionId NullableString `json:"executionId,omitempty"`
	Enabled bool `json:"enabled"`
	Module string `json:"module"`
	Profile string `json:"profile"`
	Connector string `json:"connector"`
	Enroll bool `json:"enroll"`
	Revoke bool `json:"revoke"`
	Renew bool `json:"renew"`
	Results NullableThirdPartyConnectorSynchronizationResult `json:"results,omitempty"`
	Description NullableString `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThirdPartyScheduledTask ThirdPartyScheduledTask

// NewThirdPartyScheduledTask instantiates a new ThirdPartyScheduledTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyScheduledTask(type_ string, cron string, dryRun bool, enabled bool, module string, profile string, connector string, enroll bool, revoke bool, renew bool) *ThirdPartyScheduledTask {
	this := ThirdPartyScheduledTask{}
	this.Type = type_
	this.Cron = cron
	this.DryRun = dryRun
	this.Enabled = enabled
	this.Module = module
	this.Profile = profile
	this.Connector = connector
	this.Enroll = enroll
	this.Revoke = revoke
	this.Renew = renew
	return &this
}

// NewThirdPartyScheduledTaskWithDefaults instantiates a new ThirdPartyScheduledTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyScheduledTaskWithDefaults() *ThirdPartyScheduledTask {
	this := ThirdPartyScheduledTask{}
	return &this
}

// GetType returns the Type field value
func (o *ThirdPartyScheduledTask) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTask) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ThirdPartyScheduledTask) SetType(v string) {
	o.Type = v
}

// GetCron returns the Cron field value
func (o *ThirdPartyScheduledTask) GetCron() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cron
}

// GetCronOk returns a tuple with the Cron field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTask) GetCronOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cron, true
}

// SetCron sets field value
func (o *ThirdPartyScheduledTask) SetCron(v string) {
	o.Cron = v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTask) GetHost() string {
	if o == nil || IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTask) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTask) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *ThirdPartyScheduledTask) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *ThirdPartyScheduledTask) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *ThirdPartyScheduledTask) UnsetHost() {
	o.Host.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTask) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTask) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTask) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ThirdPartyScheduledTask) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ThirdPartyScheduledTask) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ThirdPartyScheduledTask) UnsetStatus() {
	o.Status.Unset()
}

// GetLastExecutionDate returns the LastExecutionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTask) GetLastExecutionDate() int64 {
	if o == nil || IsNil(o.LastExecutionDate.Get()) {
		var ret int64
		return ret
	}
	return *o.LastExecutionDate.Get()
}

// GetLastExecutionDateOk returns a tuple with the LastExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTask) GetLastExecutionDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExecutionDate.Get(), o.LastExecutionDate.IsSet()
}

// HasLastExecutionDate returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTask) HasLastExecutionDate() bool {
	if o != nil && o.LastExecutionDate.IsSet() {
		return true
	}

	return false
}

// SetLastExecutionDate gets a reference to the given NullableInt64 and assigns it to the LastExecutionDate field.
func (o *ThirdPartyScheduledTask) SetLastExecutionDate(v int64) {
	o.LastExecutionDate.Set(&v)
}
// SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil
func (o *ThirdPartyScheduledTask) SetLastExecutionDateNil() {
	o.LastExecutionDate.Set(nil)
}

// UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
func (o *ThirdPartyScheduledTask) UnsetLastExecutionDate() {
	o.LastExecutionDate.Unset()
}

// GetLastCompletionDate returns the LastCompletionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTask) GetLastCompletionDate() int64 {
	if o == nil || IsNil(o.LastCompletionDate.Get()) {
		var ret int64
		return ret
	}
	return *o.LastCompletionDate.Get()
}

// GetLastCompletionDateOk returns a tuple with the LastCompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTask) GetLastCompletionDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastCompletionDate.Get(), o.LastCompletionDate.IsSet()
}

// HasLastCompletionDate returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTask) HasLastCompletionDate() bool {
	if o != nil && o.LastCompletionDate.IsSet() {
		return true
	}

	return false
}

// SetLastCompletionDate gets a reference to the given NullableInt64 and assigns it to the LastCompletionDate field.
func (o *ThirdPartyScheduledTask) SetLastCompletionDate(v int64) {
	o.LastCompletionDate.Set(&v)
}
// SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil
func (o *ThirdPartyScheduledTask) SetLastCompletionDateNil() {
	o.LastCompletionDate.Set(nil)
}

// UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
func (o *ThirdPartyScheduledTask) UnsetLastCompletionDate() {
	o.LastCompletionDate.Unset()
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTask) GetDetail() string {
	if o == nil || IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTask) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTask) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *ThirdPartyScheduledTask) SetDetail(v string) {
	o.Detail.Set(&v)
}
// SetDetailNil sets the value for Detail to be an explicit nil
func (o *ThirdPartyScheduledTask) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *ThirdPartyScheduledTask) UnsetDetail() {
	o.Detail.Unset()
}

// GetDryRun returns the DryRun field value
func (o *ThirdPartyScheduledTask) GetDryRun() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTask) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DryRun, true
}

// SetDryRun sets field value
func (o *ThirdPartyScheduledTask) SetDryRun(v bool) {
	o.DryRun = v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTask) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionId.Get()
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTask) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionId.Get(), o.ExecutionId.IsSet()
}

// HasExecutionId returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTask) HasExecutionId() bool {
	if o != nil && o.ExecutionId.IsSet() {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given NullableString and assigns it to the ExecutionId field.
func (o *ThirdPartyScheduledTask) SetExecutionId(v string) {
	o.ExecutionId.Set(&v)
}
// SetExecutionIdNil sets the value for ExecutionId to be an explicit nil
func (o *ThirdPartyScheduledTask) SetExecutionIdNil() {
	o.ExecutionId.Set(nil)
}

// UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
func (o *ThirdPartyScheduledTask) UnsetExecutionId() {
	o.ExecutionId.Unset()
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyScheduledTask) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTask) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyScheduledTask) SetEnabled(v bool) {
	o.Enabled = v
}

// GetModule returns the Module field value
func (o *ThirdPartyScheduledTask) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTask) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *ThirdPartyScheduledTask) SetModule(v string) {
	o.Module = v
}

// GetProfile returns the Profile field value
func (o *ThirdPartyScheduledTask) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTask) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *ThirdPartyScheduledTask) SetProfile(v string) {
	o.Profile = v
}

// GetConnector returns the Connector field value
func (o *ThirdPartyScheduledTask) GetConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTask) GetConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connector, true
}

// SetConnector sets field value
func (o *ThirdPartyScheduledTask) SetConnector(v string) {
	o.Connector = v
}

// GetEnroll returns the Enroll field value
func (o *ThirdPartyScheduledTask) GetEnroll() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enroll
}

// GetEnrollOk returns a tuple with the Enroll field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTask) GetEnrollOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enroll, true
}

// SetEnroll sets field value
func (o *ThirdPartyScheduledTask) SetEnroll(v bool) {
	o.Enroll = v
}

// GetRevoke returns the Revoke field value
func (o *ThirdPartyScheduledTask) GetRevoke() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Revoke
}

// GetRevokeOk returns a tuple with the Revoke field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTask) GetRevokeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revoke, true
}

// SetRevoke sets field value
func (o *ThirdPartyScheduledTask) SetRevoke(v bool) {
	o.Revoke = v
}

// GetRenew returns the Renew field value
func (o *ThirdPartyScheduledTask) GetRenew() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Renew
}

// GetRenewOk returns a tuple with the Renew field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyScheduledTask) GetRenewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Renew, true
}

// SetRenew sets field value
func (o *ThirdPartyScheduledTask) SetRenew(v bool) {
	o.Renew = v
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTask) GetResults() ThirdPartyConnectorSynchronizationResult {
	if o == nil || IsNil(o.Results.Get()) {
		var ret ThirdPartyConnectorSynchronizationResult
		return ret
	}
	return *o.Results.Get()
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTask) GetResultsOk() (*ThirdPartyConnectorSynchronizationResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results.Get(), o.Results.IsSet()
}

// HasResults returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTask) HasResults() bool {
	if o != nil && o.Results.IsSet() {
		return true
	}

	return false
}

// SetResults gets a reference to the given NullableThirdPartyConnectorSynchronizationResult and assigns it to the Results field.
func (o *ThirdPartyScheduledTask) SetResults(v ThirdPartyConnectorSynchronizationResult) {
	o.Results.Set(&v)
}
// SetResultsNil sets the value for Results to be an explicit nil
func (o *ThirdPartyScheduledTask) SetResultsNil() {
	o.Results.Set(nil)
}

// UnsetResults ensures that no value is present for Results, not even an explicit nil
func (o *ThirdPartyScheduledTask) UnsetResults() {
	o.Results.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyScheduledTask) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyScheduledTask) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyScheduledTask) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ThirdPartyScheduledTask) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ThirdPartyScheduledTask) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ThirdPartyScheduledTask) UnsetDescription() {
	o.Description.Unset()
}

func (o ThirdPartyScheduledTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyScheduledTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["cron"] = o.Cron
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.LastExecutionDate.IsSet() {
		toSerialize["lastExecutionDate"] = o.LastExecutionDate.Get()
	}
	if o.LastCompletionDate.IsSet() {
		toSerialize["lastCompletionDate"] = o.LastCompletionDate.Get()
	}
	if o.Detail.IsSet() {
		toSerialize["detail"] = o.Detail.Get()
	}
	toSerialize["dryRun"] = o.DryRun
	if o.ExecutionId.IsSet() {
		toSerialize["executionId"] = o.ExecutionId.Get()
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["module"] = o.Module
	toSerialize["profile"] = o.Profile
	toSerialize["connector"] = o.Connector
	toSerialize["enroll"] = o.Enroll
	toSerialize["revoke"] = o.Revoke
	toSerialize["renew"] = o.Renew
	if o.Results.IsSet() {
		toSerialize["results"] = o.Results.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThirdPartyScheduledTask) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"cron",
		"dryRun",
		"enabled",
		"module",
		"profile",
		"connector",
		"enroll",
		"revoke",
		"renew",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varThirdPartyScheduledTask := _ThirdPartyScheduledTask{}

	err = json.Unmarshal(data, &varThirdPartyScheduledTask)

	if err != nil {
		return err
	}

	*o = ThirdPartyScheduledTask(varThirdPartyScheduledTask)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "cron")
		delete(additionalProperties, "host")
		delete(additionalProperties, "status")
		delete(additionalProperties, "lastExecutionDate")
		delete(additionalProperties, "lastCompletionDate")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "dryRun")
		delete(additionalProperties, "executionId")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "module")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "connector")
		delete(additionalProperties, "enroll")
		delete(additionalProperties, "revoke")
		delete(additionalProperties, "renew")
		delete(additionalProperties, "results")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThirdPartyScheduledTask struct {
	value *ThirdPartyScheduledTask
	isSet bool
}

func (v NullableThirdPartyScheduledTask) Get() *ThirdPartyScheduledTask {
	return v.value
}

func (v *NullableThirdPartyScheduledTask) Set(val *ThirdPartyScheduledTask) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyScheduledTask) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyScheduledTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyScheduledTask(val *ThirdPartyScheduledTask) *NullableThirdPartyScheduledTask {
	return &NullableThirdPartyScheduledTask{value: val, isSet: true}
}

func (v NullableThirdPartyScheduledTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyScheduledTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


