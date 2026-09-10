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

// checks if the IntuneConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IntuneConnectorResponse{}

// IntuneConnectorResponse struct for IntuneConnectorResponse
type IntuneConnectorResponse struct {
	// Object internal ID
	Id          string  `json:"_id"`
	AzureTenant *string `json:"azureTenant,omitempty"`
	// Name of the `password` [credentials](#tag/security.credentials) containing the App ID and Key to authenticate on Intune
	Credentials          string               `json:"credentials"`
	IntuneResourceUrl    utils.NullableString `json:"intuneResourceUrl,omitempty"`
	LegacyRevocationMode bool                 `json:"legacyRevocationMode"`
	Name                 string               `json:"name"`
	OsQueryString        utils.NullableString `json:"osQueryString,omitempty"`
	Proxy                utils.NullableString `json:"proxy,omitempty"`
	ThrottleDuration     string               `json:"throttleDuration" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	ThrottleParallelism  int64                `json:"throttleParallelism"`
	Timeout              utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string               `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _IntuneConnectorResponse IntuneConnectorResponse

// NewIntuneConnectorResponse instantiates a new IntuneConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntuneConnectorResponse(id string, credentials string, legacyRevocationMode bool, name string, throttleDuration string, throttleParallelism int64, type_ string) *IntuneConnectorResponse {
	this := IntuneConnectorResponse{}
	this.Id = id
	this.Credentials = credentials
	this.LegacyRevocationMode = legacyRevocationMode
	this.Name = name
	this.ThrottleDuration = throttleDuration
	this.ThrottleParallelism = throttleParallelism
	this.Type = type_
	return &this
}

// NewIntuneConnectorResponseWithDefaults instantiates a new IntuneConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntuneConnectorResponseWithDefaults() *IntuneConnectorResponse {
	this := IntuneConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *IntuneConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IntuneConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IntuneConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetAzureTenant returns the AzureTenant field value if set, zero value otherwise.
func (o *IntuneConnectorResponse) GetAzureTenant() string {
	if o == nil || utils.IsNil(o.AzureTenant) {
		var ret string
		return ret
	}
	return *o.AzureTenant
}

// GetAzureTenantOk returns a tuple with the AzureTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntuneConnectorResponse) GetAzureTenantOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AzureTenant) {
		return nil, false
	}
	return o.AzureTenant, true
}

// HasAzureTenant returns a boolean if a field has been set.
func (o *IntuneConnectorResponse) HasAzureTenant() bool {
	if o != nil && !utils.IsNil(o.AzureTenant) {
		return true
	}

	return false
}

// SetAzureTenant gets a reference to the given string and assigns it to the AzureTenant field.
func (o *IntuneConnectorResponse) SetAzureTenant(v string) {
	o.AzureTenant = &v
}

// GetCredentials returns the Credentials field value
func (o *IntuneConnectorResponse) GetCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *IntuneConnectorResponse) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *IntuneConnectorResponse) SetCredentials(v string) {
	o.Credentials = v
}

// GetIntuneResourceUrl returns the IntuneResourceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntuneConnectorResponse) GetIntuneResourceUrl() string {
	if o == nil || utils.IsNil(o.IntuneResourceUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IntuneResourceUrl.Get()
}

// GetIntuneResourceUrlOk returns a tuple with the IntuneResourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntuneConnectorResponse) GetIntuneResourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntuneResourceUrl.Get(), o.IntuneResourceUrl.IsSet()
}

// HasIntuneResourceUrl returns a boolean if a field has been set.
func (o *IntuneConnectorResponse) HasIntuneResourceUrl() bool {
	if o != nil && o.IntuneResourceUrl.IsSet() {
		return true
	}

	return false
}

// SetIntuneResourceUrl gets a reference to the given NullableString and assigns it to the IntuneResourceUrl field.
func (o *IntuneConnectorResponse) SetIntuneResourceUrl(v string) {
	o.IntuneResourceUrl.Set(&v)
}

// SetIntuneResourceUrlNil sets the value for IntuneResourceUrl to be an explicit nil
func (o *IntuneConnectorResponse) SetIntuneResourceUrlNil() {
	o.IntuneResourceUrl.Set(nil)
}

// UnsetIntuneResourceUrl ensures that no value is present for IntuneResourceUrl, not even an explicit nil
func (o *IntuneConnectorResponse) UnsetIntuneResourceUrl() {
	o.IntuneResourceUrl.Unset()
}

// GetLegacyRevocationMode returns the LegacyRevocationMode field value
func (o *IntuneConnectorResponse) GetLegacyRevocationMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LegacyRevocationMode
}

// GetLegacyRevocationModeOk returns a tuple with the LegacyRevocationMode field value
// and a boolean to check if the value has been set.
func (o *IntuneConnectorResponse) GetLegacyRevocationModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LegacyRevocationMode, true
}

// SetLegacyRevocationMode sets field value
func (o *IntuneConnectorResponse) SetLegacyRevocationMode(v bool) {
	o.LegacyRevocationMode = v
}

// GetName returns the Name field value
func (o *IntuneConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntuneConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntuneConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetOsQueryString returns the OsQueryString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntuneConnectorResponse) GetOsQueryString() string {
	if o == nil || utils.IsNil(o.OsQueryString.Get()) {
		var ret string
		return ret
	}
	return *o.OsQueryString.Get()
}

// GetOsQueryStringOk returns a tuple with the OsQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntuneConnectorResponse) GetOsQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OsQueryString.Get(), o.OsQueryString.IsSet()
}

// HasOsQueryString returns a boolean if a field has been set.
func (o *IntuneConnectorResponse) HasOsQueryString() bool {
	if o != nil && o.OsQueryString.IsSet() {
		return true
	}

	return false
}

// SetOsQueryString gets a reference to the given NullableString and assigns it to the OsQueryString field.
func (o *IntuneConnectorResponse) SetOsQueryString(v string) {
	o.OsQueryString.Set(&v)
}

// SetOsQueryStringNil sets the value for OsQueryString to be an explicit nil
func (o *IntuneConnectorResponse) SetOsQueryStringNil() {
	o.OsQueryString.Set(nil)
}

// UnsetOsQueryString ensures that no value is present for OsQueryString, not even an explicit nil
func (o *IntuneConnectorResponse) UnsetOsQueryString() {
	o.OsQueryString.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntuneConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntuneConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *IntuneConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *IntuneConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *IntuneConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *IntuneConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetThrottleDuration returns the ThrottleDuration field value
func (o *IntuneConnectorResponse) GetThrottleDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThrottleDuration
}

// GetThrottleDurationOk returns a tuple with the ThrottleDuration field value
// and a boolean to check if the value has been set.
func (o *IntuneConnectorResponse) GetThrottleDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleDuration, true
}

// SetThrottleDuration sets field value
func (o *IntuneConnectorResponse) SetThrottleDuration(v string) {
	o.ThrottleDuration = v
}

// GetThrottleParallelism returns the ThrottleParallelism field value
func (o *IntuneConnectorResponse) GetThrottleParallelism() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ThrottleParallelism
}

// GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field value
// and a boolean to check if the value has been set.
func (o *IntuneConnectorResponse) GetThrottleParallelismOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleParallelism, true
}

// SetThrottleParallelism sets field value
func (o *IntuneConnectorResponse) SetThrottleParallelism(v int64) {
	o.ThrottleParallelism = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntuneConnectorResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntuneConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *IntuneConnectorResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *IntuneConnectorResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *IntuneConnectorResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *IntuneConnectorResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *IntuneConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntuneConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IntuneConnectorResponse) SetType(v string) {
	o.Type = v
}

func (o IntuneConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntuneConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if !utils.IsNil(o.AzureTenant) {
		toSerialize["azureTenant"] = o.AzureTenant
	}
	toSerialize["credentials"] = o.Credentials
	if o.IntuneResourceUrl.IsSet() {
		toSerialize["intuneResourceUrl"] = o.IntuneResourceUrl.Get()
	}
	toSerialize["legacyRevocationMode"] = o.LegacyRevocationMode
	toSerialize["name"] = o.Name
	if o.OsQueryString.IsSet() {
		toSerialize["osQueryString"] = o.OsQueryString.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	toSerialize["throttleDuration"] = o.ThrottleDuration
	toSerialize["throttleParallelism"] = o.ThrottleParallelism
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IntuneConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"credentials",
		"legacyRevocationMode",
		"name",
		"throttleDuration",
		"throttleParallelism",
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

	varIntuneConnectorResponse := _IntuneConnectorResponse{}

	err = json.Unmarshal(data, &varIntuneConnectorResponse)

	if err != nil {
		return err
	}

	*o = IntuneConnectorResponse(varIntuneConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "azureTenant")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "intuneResourceUrl")
		delete(additionalProperties, "legacyRevocationMode")
		delete(additionalProperties, "name")
		delete(additionalProperties, "osQueryString")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "throttleDuration")
		delete(additionalProperties, "throttleParallelism")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntuneConnectorResponse struct {
	value *IntuneConnectorResponse
	isSet bool
}

func (v NullableIntuneConnectorResponse) Get() *IntuneConnectorResponse {
	return v.value
}

func (v *NullableIntuneConnectorResponse) Set(val *IntuneConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIntuneConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIntuneConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntuneConnectorResponse(val *IntuneConnectorResponse) *NullableIntuneConnectorResponse {
	return &NullableIntuneConnectorResponse{value: val, isSet: true}
}

func (v NullableIntuneConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntuneConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
