/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the PanoramaConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PanoramaConnectorResponse{}

// PanoramaConnectorResponse struct for PanoramaConnectorResponse
type PanoramaConnectorResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Name of the `password` [credentials](#tag/security.credentials) containing the account to authenticate on Panorama
	Credentials string `json:"credentials"`
	// The hostname or URL of the Panorama appliance
	Hostname string `json:"hostname"`
	// Retry policy applied to the asynchronous deployment jobs run by this connector.
	JobRetryParameters RetryParameters `json:"jobRetryParameters"`
	Name               string          `json:"name"`
	// Certificate name prefix used when deploying certificates
	Prefix string               `json:"prefix"`
	Proxy  utils.NullableString `json:"proxy,omitempty"`
	// Synchronize (commit) the configuration to the devices managed by the Panorama after pushing the certificate
	SynchronizeDevices bool `json:"synchronizeDevices"`
	// Name of the Panorama template to push certificates to
	Template utils.NullableString `json:"template,omitempty"`
	// Name of the Panorama template stack to push changes to
	TemplateStack       utils.NullableString `json:"templateStack,omitempty"`
	ThrottleDuration    string               `json:"throttleDuration" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	ThrottleParallelism int64                `json:"throttleParallelism"`
	Timeout             string               `json:"timeout" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended.
	TlsInsecure utils.NullableBool `json:"tlsInsecure,omitempty"`
	Type        string             `json:"type"`
	// Virtual system name when targeting a specific VSYS within a template
	Vsys                 utils.NullableString `json:"vsys,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PanoramaConnectorResponse PanoramaConnectorResponse

// NewPanoramaConnectorResponse instantiates a new PanoramaConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPanoramaConnectorResponse(id string, credentials string, hostname string, jobRetryParameters RetryParameters, name string, prefix string, synchronizeDevices bool, throttleDuration string, throttleParallelism int64, timeout string, type_ string) *PanoramaConnectorResponse {
	this := PanoramaConnectorResponse{}
	this.Id = id
	this.Credentials = credentials
	this.Hostname = hostname
	this.JobRetryParameters = jobRetryParameters
	this.Name = name
	this.Prefix = prefix
	this.SynchronizeDevices = synchronizeDevices
	this.ThrottleDuration = throttleDuration
	this.ThrottleParallelism = throttleParallelism
	this.Timeout = timeout
	this.Type = type_
	var tlsInsecure bool = false
	this.TlsInsecure = *utils.NewNullableBool(&tlsInsecure)
	return &this
}

// NewPanoramaConnectorResponseWithDefaults instantiates a new PanoramaConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPanoramaConnectorResponseWithDefaults() *PanoramaConnectorResponse {
	this := PanoramaConnectorResponse{}
	var tlsInsecure bool = false
	this.TlsInsecure = *utils.NewNullableBool(&tlsInsecure)
	return &this
}

// GetId returns the Id field value
func (o *PanoramaConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PanoramaConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PanoramaConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetCredentials returns the Credentials field value
func (o *PanoramaConnectorResponse) GetCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *PanoramaConnectorResponse) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *PanoramaConnectorResponse) SetCredentials(v string) {
	o.Credentials = v
}

// GetHostname returns the Hostname field value
func (o *PanoramaConnectorResponse) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *PanoramaConnectorResponse) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *PanoramaConnectorResponse) SetHostname(v string) {
	o.Hostname = v
}

// GetJobRetryParameters returns the JobRetryParameters field value
func (o *PanoramaConnectorResponse) GetJobRetryParameters() RetryParameters {
	if o == nil {
		var ret RetryParameters
		return ret
	}

	return o.JobRetryParameters
}

// GetJobRetryParametersOk returns a tuple with the JobRetryParameters field value
// and a boolean to check if the value has been set.
func (o *PanoramaConnectorResponse) GetJobRetryParametersOk() (*RetryParameters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobRetryParameters, true
}

// SetJobRetryParameters sets field value
func (o *PanoramaConnectorResponse) SetJobRetryParameters(v RetryParameters) {
	o.JobRetryParameters = v
}

// GetName returns the Name field value
func (o *PanoramaConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PanoramaConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PanoramaConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value
func (o *PanoramaConnectorResponse) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *PanoramaConnectorResponse) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *PanoramaConnectorResponse) SetPrefix(v string) {
	o.Prefix = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PanoramaConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PanoramaConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *PanoramaConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *PanoramaConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *PanoramaConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *PanoramaConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetSynchronizeDevices returns the SynchronizeDevices field value
func (o *PanoramaConnectorResponse) GetSynchronizeDevices() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SynchronizeDevices
}

// GetSynchronizeDevicesOk returns a tuple with the SynchronizeDevices field value
// and a boolean to check if the value has been set.
func (o *PanoramaConnectorResponse) GetSynchronizeDevicesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SynchronizeDevices, true
}

// SetSynchronizeDevices sets field value
func (o *PanoramaConnectorResponse) SetSynchronizeDevices(v bool) {
	o.SynchronizeDevices = v
}

// GetTemplate returns the Template field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PanoramaConnectorResponse) GetTemplate() string {
	if o == nil || utils.IsNil(o.Template.Get()) {
		var ret string
		return ret
	}
	return *o.Template.Get()
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PanoramaConnectorResponse) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Template.Get(), o.Template.IsSet()
}

// HasTemplate returns a boolean if a field has been set.
func (o *PanoramaConnectorResponse) HasTemplate() bool {
	if o != nil && o.Template.IsSet() {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given NullableString and assigns it to the Template field.
func (o *PanoramaConnectorResponse) SetTemplate(v string) {
	o.Template.Set(&v)
}

// SetTemplateNil sets the value for Template to be an explicit nil
func (o *PanoramaConnectorResponse) SetTemplateNil() {
	o.Template.Set(nil)
}

// UnsetTemplate ensures that no value is present for Template, not even an explicit nil
func (o *PanoramaConnectorResponse) UnsetTemplate() {
	o.Template.Unset()
}

// GetTemplateStack returns the TemplateStack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PanoramaConnectorResponse) GetTemplateStack() string {
	if o == nil || utils.IsNil(o.TemplateStack.Get()) {
		var ret string
		return ret
	}
	return *o.TemplateStack.Get()
}

// GetTemplateStackOk returns a tuple with the TemplateStack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PanoramaConnectorResponse) GetTemplateStackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateStack.Get(), o.TemplateStack.IsSet()
}

// HasTemplateStack returns a boolean if a field has been set.
func (o *PanoramaConnectorResponse) HasTemplateStack() bool {
	if o != nil && o.TemplateStack.IsSet() {
		return true
	}

	return false
}

// SetTemplateStack gets a reference to the given NullableString and assigns it to the TemplateStack field.
func (o *PanoramaConnectorResponse) SetTemplateStack(v string) {
	o.TemplateStack.Set(&v)
}

// SetTemplateStackNil sets the value for TemplateStack to be an explicit nil
func (o *PanoramaConnectorResponse) SetTemplateStackNil() {
	o.TemplateStack.Set(nil)
}

// UnsetTemplateStack ensures that no value is present for TemplateStack, not even an explicit nil
func (o *PanoramaConnectorResponse) UnsetTemplateStack() {
	o.TemplateStack.Unset()
}

// GetThrottleDuration returns the ThrottleDuration field value
func (o *PanoramaConnectorResponse) GetThrottleDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThrottleDuration
}

// GetThrottleDurationOk returns a tuple with the ThrottleDuration field value
// and a boolean to check if the value has been set.
func (o *PanoramaConnectorResponse) GetThrottleDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleDuration, true
}

// SetThrottleDuration sets field value
func (o *PanoramaConnectorResponse) SetThrottleDuration(v string) {
	o.ThrottleDuration = v
}

// GetThrottleParallelism returns the ThrottleParallelism field value
func (o *PanoramaConnectorResponse) GetThrottleParallelism() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ThrottleParallelism
}

// GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field value
// and a boolean to check if the value has been set.
func (o *PanoramaConnectorResponse) GetThrottleParallelismOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleParallelism, true
}

// SetThrottleParallelism sets field value
func (o *PanoramaConnectorResponse) SetThrottleParallelism(v int64) {
	o.ThrottleParallelism = v
}

// GetTimeout returns the Timeout field value
func (o *PanoramaConnectorResponse) GetTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *PanoramaConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *PanoramaConnectorResponse) SetTimeout(v string) {
	o.Timeout = v
}

// GetTlsInsecure returns the TlsInsecure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PanoramaConnectorResponse) GetTlsInsecure() bool {
	if o == nil || utils.IsNil(o.TlsInsecure.Get()) {
		var ret bool
		return ret
	}
	return *o.TlsInsecure.Get()
}

// GetTlsInsecureOk returns a tuple with the TlsInsecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PanoramaConnectorResponse) GetTlsInsecureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsInsecure.Get(), o.TlsInsecure.IsSet()
}

// HasTlsInsecure returns a boolean if a field has been set.
func (o *PanoramaConnectorResponse) HasTlsInsecure() bool {
	if o != nil && o.TlsInsecure.IsSet() {
		return true
	}

	return false
}

// SetTlsInsecure gets a reference to the given NullableBool and assigns it to the TlsInsecure field.
func (o *PanoramaConnectorResponse) SetTlsInsecure(v bool) {
	o.TlsInsecure.Set(&v)
}

// SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil
func (o *PanoramaConnectorResponse) SetTlsInsecureNil() {
	o.TlsInsecure.Set(nil)
}

// UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
func (o *PanoramaConnectorResponse) UnsetTlsInsecure() {
	o.TlsInsecure.Unset()
}

// GetType returns the Type field value
func (o *PanoramaConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PanoramaConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PanoramaConnectorResponse) SetType(v string) {
	o.Type = v
}

// GetVsys returns the Vsys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PanoramaConnectorResponse) GetVsys() string {
	if o == nil || utils.IsNil(o.Vsys.Get()) {
		var ret string
		return ret
	}
	return *o.Vsys.Get()
}

// GetVsysOk returns a tuple with the Vsys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PanoramaConnectorResponse) GetVsysOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vsys.Get(), o.Vsys.IsSet()
}

// HasVsys returns a boolean if a field has been set.
func (o *PanoramaConnectorResponse) HasVsys() bool {
	if o != nil && o.Vsys.IsSet() {
		return true
	}

	return false
}

// SetVsys gets a reference to the given NullableString and assigns it to the Vsys field.
func (o *PanoramaConnectorResponse) SetVsys(v string) {
	o.Vsys.Set(&v)
}

// SetVsysNil sets the value for Vsys to be an explicit nil
func (o *PanoramaConnectorResponse) SetVsysNil() {
	o.Vsys.Set(nil)
}

// UnsetVsys ensures that no value is present for Vsys, not even an explicit nil
func (o *PanoramaConnectorResponse) UnsetVsys() {
	o.Vsys.Unset()
}

func (o PanoramaConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PanoramaConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["credentials"] = o.Credentials
	toSerialize["hostname"] = o.Hostname
	toSerialize["jobRetryParameters"] = o.JobRetryParameters
	toSerialize["name"] = o.Name
	toSerialize["prefix"] = o.Prefix
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	toSerialize["synchronizeDevices"] = o.SynchronizeDevices
	if o.Template.IsSet() {
		toSerialize["template"] = o.Template.Get()
	}
	if o.TemplateStack.IsSet() {
		toSerialize["templateStack"] = o.TemplateStack.Get()
	}
	toSerialize["throttleDuration"] = o.ThrottleDuration
	toSerialize["throttleParallelism"] = o.ThrottleParallelism
	toSerialize["timeout"] = o.Timeout
	if o.TlsInsecure.IsSet() {
		toSerialize["tlsInsecure"] = o.TlsInsecure.Get()
	}
	toSerialize["type"] = o.Type
	if o.Vsys.IsSet() {
		toSerialize["vsys"] = o.Vsys.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PanoramaConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"credentials",
		"hostname",
		"jobRetryParameters",
		"name",
		"prefix",
		"synchronizeDevices",
		"throttleDuration",
		"throttleParallelism",
		"timeout",
		"type",
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

	varPanoramaConnectorResponse := _PanoramaConnectorResponse{}

	err = json.Unmarshal(data, &varPanoramaConnectorResponse)

	if err != nil {
		return err
	}

	*o = PanoramaConnectorResponse(varPanoramaConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "jobRetryParameters")
		delete(additionalProperties, "name")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "synchronizeDevices")
		delete(additionalProperties, "template")
		delete(additionalProperties, "templateStack")
		delete(additionalProperties, "throttleDuration")
		delete(additionalProperties, "throttleParallelism")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "tlsInsecure")
		delete(additionalProperties, "type")
		delete(additionalProperties, "vsys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePanoramaConnectorResponse struct {
	value *PanoramaConnectorResponse
	isSet bool
}

func (v NullablePanoramaConnectorResponse) Get() *PanoramaConnectorResponse {
	return v.value
}

func (v *NullablePanoramaConnectorResponse) Set(val *PanoramaConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePanoramaConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePanoramaConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePanoramaConnectorResponse(val *PanoramaConnectorResponse) *NullablePanoramaConnectorResponse {
	return &NullablePanoramaConnectorResponse{value: val, isSet: true}
}

func (v NullablePanoramaConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePanoramaConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
