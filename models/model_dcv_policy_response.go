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

// checks if the DCVPolicyResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DCVPolicyResponse{}

// DCVPolicyResponse struct for DCVPolicyResponse
type DCVPolicyResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Whether the DCV policy is enabled; disabled policies are not scheduled
	Enabled *bool `json:"enabled,omitempty"`
	// Maximum duration for a single DCV run
	ExecutionTimeout utils.NullableString `json:"executionTimeout" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Optional regex filter applied to domain hostnames
	Filter utils.NullableString `json:"filter,omitempty"`
	// Unique name of the DCV policy
	Name string `json:"name"`
	// Name of the DCV provider configuration to use
	Provider string `json:"provider"`
	// Name of the DCV provisioner configuration to use
	Provisioner string `json:"provisioner"`
	// Renewal scheduling policy; if absent no automatic renewal is triggered
	RenewalPolicy NullableDCVRenewalPolicy `json:"renewalPolicy,omitempty"`
	// Delay between retry attempts
	RetryDelay utils.NullableString `json:"retryDelay" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Optional trigger configuration for DCV lifecycle events
	Triggers             NullableDCVPolicyTriggers `json:"triggers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DCVPolicyResponse DCVPolicyResponse

// NewDCVPolicyResponse instantiates a new DCVPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDCVPolicyResponse(id string, executionTimeout utils.NullableString, name string, provider string, provisioner string, retryDelay utils.NullableString) *DCVPolicyResponse {
	this := DCVPolicyResponse{}
	this.Id = id
	this.ExecutionTimeout = executionTimeout
	this.Name = name
	this.Provider = provider
	this.Provisioner = provisioner
	this.RetryDelay = retryDelay
	return &this
}

// NewDCVPolicyResponseWithDefaults instantiates a new DCVPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDCVPolicyResponseWithDefaults() *DCVPolicyResponse {
	this := DCVPolicyResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DCVPolicyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DCVPolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DCVPolicyResponse) SetId(v string) {
	o.Id = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DCVPolicyResponse) GetEnabled() bool {
	if o == nil || utils.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DCVPolicyResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DCVPolicyResponse) HasEnabled() bool {
	if o != nil && !utils.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DCVPolicyResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExecutionTimeout returns the ExecutionTimeout field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DCVPolicyResponse) GetExecutionTimeout() string {
	if o == nil || o.ExecutionTimeout.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExecutionTimeout.Get()
}

// GetExecutionTimeoutOk returns a tuple with the ExecutionTimeout field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DCVPolicyResponse) GetExecutionTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionTimeout.Get(), o.ExecutionTimeout.IsSet()
}

// SetExecutionTimeout sets field value
func (o *DCVPolicyResponse) SetExecutionTimeout(v string) {
	o.ExecutionTimeout.Set(&v)
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DCVPolicyResponse) GetFilter() string {
	if o == nil || utils.IsNil(o.Filter.Get()) {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DCVPolicyResponse) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *DCVPolicyResponse) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableString and assigns it to the Filter field.
func (o *DCVPolicyResponse) SetFilter(v string) {
	o.Filter.Set(&v)
}

// SetFilterNil sets the value for Filter to be an explicit nil
func (o *DCVPolicyResponse) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *DCVPolicyResponse) UnsetFilter() {
	o.Filter.Unset()
}

// GetName returns the Name field value
func (o *DCVPolicyResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DCVPolicyResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DCVPolicyResponse) SetName(v string) {
	o.Name = v
}

// GetProvider returns the Provider field value
func (o *DCVPolicyResponse) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *DCVPolicyResponse) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *DCVPolicyResponse) SetProvider(v string) {
	o.Provider = v
}

// GetProvisioner returns the Provisioner field value
func (o *DCVPolicyResponse) GetProvisioner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provisioner
}

// GetProvisionerOk returns a tuple with the Provisioner field value
// and a boolean to check if the value has been set.
func (o *DCVPolicyResponse) GetProvisionerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provisioner, true
}

// SetProvisioner sets field value
func (o *DCVPolicyResponse) SetProvisioner(v string) {
	o.Provisioner = v
}

// GetRenewalPolicy returns the RenewalPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DCVPolicyResponse) GetRenewalPolicy() DCVRenewalPolicy {
	if o == nil || utils.IsNil(o.RenewalPolicy.Get()) {
		var ret DCVRenewalPolicy
		return ret
	}
	return *o.RenewalPolicy.Get()
}

// GetRenewalPolicyOk returns a tuple with the RenewalPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DCVPolicyResponse) GetRenewalPolicyOk() (*DCVRenewalPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPolicy.Get(), o.RenewalPolicy.IsSet()
}

// HasRenewalPolicy returns a boolean if a field has been set.
func (o *DCVPolicyResponse) HasRenewalPolicy() bool {
	if o != nil && o.RenewalPolicy.IsSet() {
		return true
	}

	return false
}

// SetRenewalPolicy gets a reference to the given NullableDCVRenewalPolicy and assigns it to the RenewalPolicy field.
func (o *DCVPolicyResponse) SetRenewalPolicy(v DCVRenewalPolicy) {
	o.RenewalPolicy.Set(&v)
}

// SetRenewalPolicyNil sets the value for RenewalPolicy to be an explicit nil
func (o *DCVPolicyResponse) SetRenewalPolicyNil() {
	o.RenewalPolicy.Set(nil)
}

// UnsetRenewalPolicy ensures that no value is present for RenewalPolicy, not even an explicit nil
func (o *DCVPolicyResponse) UnsetRenewalPolicy() {
	o.RenewalPolicy.Unset()
}

// GetRetryDelay returns the RetryDelay field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DCVPolicyResponse) GetRetryDelay() string {
	if o == nil || o.RetryDelay.Get() == nil {
		var ret string
		return ret
	}

	return *o.RetryDelay.Get()
}

// GetRetryDelayOk returns a tuple with the RetryDelay field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DCVPolicyResponse) GetRetryDelayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetryDelay.Get(), o.RetryDelay.IsSet()
}

// SetRetryDelay sets field value
func (o *DCVPolicyResponse) SetRetryDelay(v string) {
	o.RetryDelay.Set(&v)
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DCVPolicyResponse) GetTriggers() DCVPolicyTriggers {
	if o == nil || utils.IsNil(o.Triggers.Get()) {
		var ret DCVPolicyTriggers
		return ret
	}
	return *o.Triggers.Get()
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DCVPolicyResponse) GetTriggersOk() (*DCVPolicyTriggers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers.Get(), o.Triggers.IsSet()
}

// HasTriggers returns a boolean if a field has been set.
func (o *DCVPolicyResponse) HasTriggers() bool {
	if o != nil && o.Triggers.IsSet() {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given NullableDCVPolicyTriggers and assigns it to the Triggers field.
func (o *DCVPolicyResponse) SetTriggers(v DCVPolicyTriggers) {
	o.Triggers.Set(&v)
}

// SetTriggersNil sets the value for Triggers to be an explicit nil
func (o *DCVPolicyResponse) SetTriggersNil() {
	o.Triggers.Set(nil)
}

// UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
func (o *DCVPolicyResponse) UnsetTriggers() {
	o.Triggers.Unset()
}

func (o DCVPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DCVPolicyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if !utils.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["executionTimeout"] = o.ExecutionTimeout.Get()
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["provider"] = o.Provider
	toSerialize["provisioner"] = o.Provisioner
	if o.RenewalPolicy.IsSet() {
		toSerialize["renewalPolicy"] = o.RenewalPolicy.Get()
	}
	toSerialize["retryDelay"] = o.RetryDelay.Get()
	if o.Triggers.IsSet() {
		toSerialize["triggers"] = o.Triggers.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DCVPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"executionTimeout",
		"name",
		"provider",
		"provisioner",
		"retryDelay",
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

	varDCVPolicyResponse := _DCVPolicyResponse{}

	err = json.Unmarshal(data, &varDCVPolicyResponse)

	if err != nil {
		return err
	}

	*o = DCVPolicyResponse(varDCVPolicyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "executionTimeout")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "name")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "provisioner")
		delete(additionalProperties, "renewalPolicy")
		delete(additionalProperties, "retryDelay")
		delete(additionalProperties, "triggers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDCVPolicyResponse struct {
	value *DCVPolicyResponse
	isSet bool
}

func (v NullableDCVPolicyResponse) Get() *DCVPolicyResponse {
	return v.value
}

func (v *NullableDCVPolicyResponse) Set(val *DCVPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDCVPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDCVPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDCVPolicyResponse(val *DCVPolicyResponse) *NullableDCVPolicyResponse {
	return &NullableDCVPolicyResponse{value: val, isSet: true}
}

func (v NullableDCVPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDCVPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
