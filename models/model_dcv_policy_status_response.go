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

// checks if the DCVPolicyStatusResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DCVPolicyStatusResponse{}

// DCVPolicyStatusResponse struct for DCVPolicyStatusResponse
type DCVPolicyStatusResponse struct {
	// Domain validation statuses, with an optional error if the provider could not be reached
	DomainsStatus DCVDomainsStatus `json:"domainsStatus"`
	// Whether the DCV policy is enabled
	Enabled bool `json:"enabled"`
	// Maximum duration allowed for a single DCV run
	ExecutionTimeout utils.NullableString `json:"executionTimeout" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Epoch milliseconds at which the current execution will be forcefully ended, only present when status is running
	ExecutionTimeoutAt utils.NullableInt64 `json:"executionTimeoutAt,omitempty"`
	// Unique name of the DCV policy
	Name string `json:"name"`
	// Epoch milliseconds of the next retry check, only present when status is running
	NextCheckAt utils.NullableInt64 `json:"nextCheckAt,omitempty"`
	// Duration before expiry at which renewal is triggered; absent if no renewal policy is set
	RenewalPeriod utils.NullableString `json:"renewalPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Delay between retry attempts
	RetryDelay utils.NullableString `json:"retryDelay" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Whether the policy is enabled and the principal has manage permission
	Runnable bool `json:"runnable"`
	// Epoch milliseconds when the current execution started, only present when status is running
	StartedAt utils.NullableInt64 `json:"startedAt,omitempty"`
	// Current status of the DCV policy
	Status               string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _DCVPolicyStatusResponse DCVPolicyStatusResponse

// NewDCVPolicyStatusResponse instantiates a new DCVPolicyStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDCVPolicyStatusResponse(domainsStatus DCVDomainsStatus, enabled bool, executionTimeout utils.NullableString, name string, retryDelay utils.NullableString, runnable bool, status string) *DCVPolicyStatusResponse {
	this := DCVPolicyStatusResponse{}
	this.DomainsStatus = domainsStatus
	this.Enabled = enabled
	this.ExecutionTimeout = executionTimeout
	this.Name = name
	this.RetryDelay = retryDelay
	this.Runnable = runnable
	this.Status = status
	return &this
}

// NewDCVPolicyStatusResponseWithDefaults instantiates a new DCVPolicyStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDCVPolicyStatusResponseWithDefaults() *DCVPolicyStatusResponse {
	this := DCVPolicyStatusResponse{}
	return &this
}

// GetDomainsStatus returns the DomainsStatus field value
func (o *DCVPolicyStatusResponse) GetDomainsStatus() DCVDomainsStatus {
	if o == nil {
		var ret DCVDomainsStatus
		return ret
	}

	return o.DomainsStatus
}

// GetDomainsStatusOk returns a tuple with the DomainsStatus field value
// and a boolean to check if the value has been set.
func (o *DCVPolicyStatusResponse) GetDomainsStatusOk() (*DCVDomainsStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainsStatus, true
}

// SetDomainsStatus sets field value
func (o *DCVPolicyStatusResponse) SetDomainsStatus(v DCVDomainsStatus) {
	o.DomainsStatus = v
}

// GetEnabled returns the Enabled field value
func (o *DCVPolicyStatusResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DCVPolicyStatusResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DCVPolicyStatusResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetExecutionTimeout returns the ExecutionTimeout field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DCVPolicyStatusResponse) GetExecutionTimeout() string {
	if o == nil || o.ExecutionTimeout.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExecutionTimeout.Get()
}

// GetExecutionTimeoutOk returns a tuple with the ExecutionTimeout field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DCVPolicyStatusResponse) GetExecutionTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionTimeout.Get(), o.ExecutionTimeout.IsSet()
}

// SetExecutionTimeout sets field value
func (o *DCVPolicyStatusResponse) SetExecutionTimeout(v string) {
	o.ExecutionTimeout.Set(&v)
}

// GetExecutionTimeoutAt returns the ExecutionTimeoutAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DCVPolicyStatusResponse) GetExecutionTimeoutAt() int64 {
	if o == nil || utils.IsNil(o.ExecutionTimeoutAt.Get()) {
		var ret int64
		return ret
	}
	return *o.ExecutionTimeoutAt.Get()
}

// GetExecutionTimeoutAtOk returns a tuple with the ExecutionTimeoutAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DCVPolicyStatusResponse) GetExecutionTimeoutAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionTimeoutAt.Get(), o.ExecutionTimeoutAt.IsSet()
}

// HasExecutionTimeoutAt returns a boolean if a field has been set.
func (o *DCVPolicyStatusResponse) HasExecutionTimeoutAt() bool {
	if o != nil && o.ExecutionTimeoutAt.IsSet() {
		return true
	}

	return false
}

// SetExecutionTimeoutAt gets a reference to the given NullableInt64 and assigns it to the ExecutionTimeoutAt field.
func (o *DCVPolicyStatusResponse) SetExecutionTimeoutAt(v int64) {
	o.ExecutionTimeoutAt.Set(&v)
}

// SetExecutionTimeoutAtNil sets the value for ExecutionTimeoutAt to be an explicit nil
func (o *DCVPolicyStatusResponse) SetExecutionTimeoutAtNil() {
	o.ExecutionTimeoutAt.Set(nil)
}

// UnsetExecutionTimeoutAt ensures that no value is present for ExecutionTimeoutAt, not even an explicit nil
func (o *DCVPolicyStatusResponse) UnsetExecutionTimeoutAt() {
	o.ExecutionTimeoutAt.Unset()
}

// GetName returns the Name field value
func (o *DCVPolicyStatusResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DCVPolicyStatusResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DCVPolicyStatusResponse) SetName(v string) {
	o.Name = v
}

// GetNextCheckAt returns the NextCheckAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DCVPolicyStatusResponse) GetNextCheckAt() int64 {
	if o == nil || utils.IsNil(o.NextCheckAt.Get()) {
		var ret int64
		return ret
	}
	return *o.NextCheckAt.Get()
}

// GetNextCheckAtOk returns a tuple with the NextCheckAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DCVPolicyStatusResponse) GetNextCheckAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCheckAt.Get(), o.NextCheckAt.IsSet()
}

// HasNextCheckAt returns a boolean if a field has been set.
func (o *DCVPolicyStatusResponse) HasNextCheckAt() bool {
	if o != nil && o.NextCheckAt.IsSet() {
		return true
	}

	return false
}

// SetNextCheckAt gets a reference to the given NullableInt64 and assigns it to the NextCheckAt field.
func (o *DCVPolicyStatusResponse) SetNextCheckAt(v int64) {
	o.NextCheckAt.Set(&v)
}

// SetNextCheckAtNil sets the value for NextCheckAt to be an explicit nil
func (o *DCVPolicyStatusResponse) SetNextCheckAtNil() {
	o.NextCheckAt.Set(nil)
}

// UnsetNextCheckAt ensures that no value is present for NextCheckAt, not even an explicit nil
func (o *DCVPolicyStatusResponse) UnsetNextCheckAt() {
	o.NextCheckAt.Unset()
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DCVPolicyStatusResponse) GetRenewalPeriod() string {
	if o == nil || utils.IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DCVPolicyStatusResponse) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *DCVPolicyStatusResponse) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *DCVPolicyStatusResponse) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}

// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *DCVPolicyStatusResponse) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *DCVPolicyStatusResponse) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetRetryDelay returns the RetryDelay field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DCVPolicyStatusResponse) GetRetryDelay() string {
	if o == nil || o.RetryDelay.Get() == nil {
		var ret string
		return ret
	}

	return *o.RetryDelay.Get()
}

// GetRetryDelayOk returns a tuple with the RetryDelay field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DCVPolicyStatusResponse) GetRetryDelayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetryDelay.Get(), o.RetryDelay.IsSet()
}

// SetRetryDelay sets field value
func (o *DCVPolicyStatusResponse) SetRetryDelay(v string) {
	o.RetryDelay.Set(&v)
}

// GetRunnable returns the Runnable field value
func (o *DCVPolicyStatusResponse) GetRunnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Runnable
}

// GetRunnableOk returns a tuple with the Runnable field value
// and a boolean to check if the value has been set.
func (o *DCVPolicyStatusResponse) GetRunnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Runnable, true
}

// SetRunnable sets field value
func (o *DCVPolicyStatusResponse) SetRunnable(v bool) {
	o.Runnable = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DCVPolicyStatusResponse) GetStartedAt() int64 {
	if o == nil || utils.IsNil(o.StartedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DCVPolicyStatusResponse) GetStartedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *DCVPolicyStatusResponse) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given NullableInt64 and assigns it to the StartedAt field.
func (o *DCVPolicyStatusResponse) SetStartedAt(v int64) {
	o.StartedAt.Set(&v)
}

// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *DCVPolicyStatusResponse) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *DCVPolicyStatusResponse) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetStatus returns the Status field value
func (o *DCVPolicyStatusResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DCVPolicyStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DCVPolicyStatusResponse) SetStatus(v string) {
	o.Status = v
}

func (o DCVPolicyStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DCVPolicyStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domainsStatus"] = o.DomainsStatus
	toSerialize["enabled"] = o.Enabled
	toSerialize["executionTimeout"] = o.ExecutionTimeout.Get()
	if o.ExecutionTimeoutAt.IsSet() {
		toSerialize["executionTimeoutAt"] = o.ExecutionTimeoutAt.Get()
	}
	toSerialize["name"] = o.Name
	if o.NextCheckAt.IsSet() {
		toSerialize["nextCheckAt"] = o.NextCheckAt.Get()
	}
	if o.RenewalPeriod.IsSet() {
		toSerialize["renewalPeriod"] = o.RenewalPeriod.Get()
	}
	toSerialize["retryDelay"] = o.RetryDelay.Get()
	toSerialize["runnable"] = o.Runnable
	if o.StartedAt.IsSet() {
		toSerialize["startedAt"] = o.StartedAt.Get()
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DCVPolicyStatusResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domainsStatus",
		"enabled",
		"executionTimeout",
		"name",
		"retryDelay",
		"runnable",
		"status",
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

	varDCVPolicyStatusResponse := _DCVPolicyStatusResponse{}

	err = json.Unmarshal(data, &varDCVPolicyStatusResponse)

	if err != nil {
		return err
	}

	*o = DCVPolicyStatusResponse(varDCVPolicyStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "domainsStatus")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "executionTimeout")
		delete(additionalProperties, "executionTimeoutAt")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nextCheckAt")
		delete(additionalProperties, "renewalPeriod")
		delete(additionalProperties, "retryDelay")
		delete(additionalProperties, "runnable")
		delete(additionalProperties, "startedAt")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDCVPolicyStatusResponse struct {
	value *DCVPolicyStatusResponse
	isSet bool
}

func (v NullableDCVPolicyStatusResponse) Get() *DCVPolicyStatusResponse {
	return v.value
}

func (v *NullableDCVPolicyStatusResponse) Set(val *DCVPolicyStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDCVPolicyStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDCVPolicyStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDCVPolicyStatusResponse(val *DCVPolicyStatusResponse) *NullableDCVPolicyStatusResponse {
	return &NullableDCVPolicyStatusResponse{value: val, isSet: true}
}

func (v NullableDCVPolicyStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDCVPolicyStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
