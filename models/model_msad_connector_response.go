/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the MSADConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MSADConnectorResponse{}

// MSADConnectorResponse struct for MSADConnectorResponse
type MSADConnectorResponse struct {
	// Object internal ID
	Id                            string               `json:"_id"`
	Type                          string               `json:"type"`
	Name                          string               `json:"name"`
	ThrottleDuration              string               `json:"throttleDuration" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	ThrottleParallelism           int64                `json:"throttleParallelism"`
	Timeout                       utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	MaxStoredCertificatePerHolder utils.NullableInt64  `json:"maxStoredCertificatePerHolder,omitempty"`
	Proxy                         utils.NullableString `json:"proxy,omitempty"`
	Hostname                      string               `json:"hostname"`
	Port                          utils.NullableInt64  `json:"port,omitempty"`
	// Name of the `password` [credentials](#tag/security.credentials) containing the DN and password to authenticate on Active Directory
	Credentials string               `json:"credentials"`
	BaseDn      string               `json:"baseDn"`
	Filter      utils.NullableString `json:"filter,omitempty"`
	// Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended.
	TlsInsecure          utils.NullableBool `json:"tlsInsecure,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MSADConnectorResponse MSADConnectorResponse

// NewMSADConnectorResponse instantiates a new MSADConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMSADConnectorResponse(id string, type_ string, name string, throttleDuration string, throttleParallelism int64, hostname string, credentials string, baseDn string) *MSADConnectorResponse {
	this := MSADConnectorResponse{}
	this.Id = id
	this.Type = type_
	this.Name = name
	this.ThrottleDuration = throttleDuration
	this.ThrottleParallelism = throttleParallelism
	this.Hostname = hostname
	this.Credentials = credentials
	this.BaseDn = baseDn
	var tlsInsecure bool = false
	this.TlsInsecure = *utils.NewNullableBool(&tlsInsecure)
	return &this
}

// NewMSADConnectorResponseWithDefaults instantiates a new MSADConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMSADConnectorResponseWithDefaults() *MSADConnectorResponse {
	this := MSADConnectorResponse{}
	var tlsInsecure bool = false
	this.TlsInsecure = *utils.NewNullableBool(&tlsInsecure)
	return &this
}

// GetId returns the Id field value
func (o *MSADConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MSADConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MSADConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *MSADConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MSADConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MSADConnectorResponse) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *MSADConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MSADConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MSADConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetThrottleDuration returns the ThrottleDuration field value
func (o *MSADConnectorResponse) GetThrottleDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThrottleDuration
}

// GetThrottleDurationOk returns a tuple with the ThrottleDuration field value
// and a boolean to check if the value has been set.
func (o *MSADConnectorResponse) GetThrottleDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleDuration, true
}

// SetThrottleDuration sets field value
func (o *MSADConnectorResponse) SetThrottleDuration(v string) {
	o.ThrottleDuration = v
}

// GetThrottleParallelism returns the ThrottleParallelism field value
func (o *MSADConnectorResponse) GetThrottleParallelism() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ThrottleParallelism
}

// GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field value
// and a boolean to check if the value has been set.
func (o *MSADConnectorResponse) GetThrottleParallelismOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleParallelism, true
}

// SetThrottleParallelism sets field value
func (o *MSADConnectorResponse) SetThrottleParallelism(v int64) {
	o.ThrottleParallelism = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSADConnectorResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSADConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *MSADConnectorResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *MSADConnectorResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *MSADConnectorResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *MSADConnectorResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSADConnectorResponse) GetMaxStoredCertificatePerHolder() int64 {
	if o == nil || utils.IsNil(o.MaxStoredCertificatePerHolder.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxStoredCertificatePerHolder.Get()
}

// GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSADConnectorResponse) GetMaxStoredCertificatePerHolderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStoredCertificatePerHolder.Get(), o.MaxStoredCertificatePerHolder.IsSet()
}

// HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.
func (o *MSADConnectorResponse) HasMaxStoredCertificatePerHolder() bool {
	if o != nil && o.MaxStoredCertificatePerHolder.IsSet() {
		return true
	}

	return false
}

// SetMaxStoredCertificatePerHolder gets a reference to the given NullableInt64 and assigns it to the MaxStoredCertificatePerHolder field.
func (o *MSADConnectorResponse) SetMaxStoredCertificatePerHolder(v int64) {
	o.MaxStoredCertificatePerHolder.Set(&v)
}

// SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil
func (o *MSADConnectorResponse) SetMaxStoredCertificatePerHolderNil() {
	o.MaxStoredCertificatePerHolder.Set(nil)
}

// UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
func (o *MSADConnectorResponse) UnsetMaxStoredCertificatePerHolder() {
	o.MaxStoredCertificatePerHolder.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSADConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSADConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *MSADConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *MSADConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *MSADConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *MSADConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetHostname returns the Hostname field value
func (o *MSADConnectorResponse) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *MSADConnectorResponse) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *MSADConnectorResponse) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSADConnectorResponse) GetPort() int64 {
	if o == nil || utils.IsNil(o.Port.Get()) {
		var ret int64
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSADConnectorResponse) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *MSADConnectorResponse) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt64 and assigns it to the Port field.
func (o *MSADConnectorResponse) SetPort(v int64) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil
func (o *MSADConnectorResponse) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *MSADConnectorResponse) UnsetPort() {
	o.Port.Unset()
}

// GetCredentials returns the Credentials field value
func (o *MSADConnectorResponse) GetCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *MSADConnectorResponse) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *MSADConnectorResponse) SetCredentials(v string) {
	o.Credentials = v
}

// GetBaseDn returns the BaseDn field value
func (o *MSADConnectorResponse) GetBaseDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value
// and a boolean to check if the value has been set.
func (o *MSADConnectorResponse) GetBaseDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseDn, true
}

// SetBaseDn sets field value
func (o *MSADConnectorResponse) SetBaseDn(v string) {
	o.BaseDn = v
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSADConnectorResponse) GetFilter() string {
	if o == nil || utils.IsNil(o.Filter.Get()) {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSADConnectorResponse) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *MSADConnectorResponse) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableString and assigns it to the Filter field.
func (o *MSADConnectorResponse) SetFilter(v string) {
	o.Filter.Set(&v)
}

// SetFilterNil sets the value for Filter to be an explicit nil
func (o *MSADConnectorResponse) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *MSADConnectorResponse) UnsetFilter() {
	o.Filter.Unset()
}

// GetTlsInsecure returns the TlsInsecure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MSADConnectorResponse) GetTlsInsecure() bool {
	if o == nil || utils.IsNil(o.TlsInsecure.Get()) {
		var ret bool
		return ret
	}
	return *o.TlsInsecure.Get()
}

// GetTlsInsecureOk returns a tuple with the TlsInsecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MSADConnectorResponse) GetTlsInsecureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsInsecure.Get(), o.TlsInsecure.IsSet()
}

// HasTlsInsecure returns a boolean if a field has been set.
func (o *MSADConnectorResponse) HasTlsInsecure() bool {
	if o != nil && o.TlsInsecure.IsSet() {
		return true
	}

	return false
}

// SetTlsInsecure gets a reference to the given NullableBool and assigns it to the TlsInsecure field.
func (o *MSADConnectorResponse) SetTlsInsecure(v bool) {
	o.TlsInsecure.Set(&v)
}

// SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil
func (o *MSADConnectorResponse) SetTlsInsecureNil() {
	o.TlsInsecure.Set(nil)
}

// UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
func (o *MSADConnectorResponse) UnsetTlsInsecure() {
	o.TlsInsecure.Unset()
}

func (o MSADConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MSADConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["throttleDuration"] = o.ThrottleDuration
	toSerialize["throttleParallelism"] = o.ThrottleParallelism
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	if o.MaxStoredCertificatePerHolder.IsSet() {
		toSerialize["maxStoredCertificatePerHolder"] = o.MaxStoredCertificatePerHolder.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	toSerialize["hostname"] = o.Hostname
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	toSerialize["credentials"] = o.Credentials
	toSerialize["baseDn"] = o.BaseDn
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	if o.TlsInsecure.IsSet() {
		toSerialize["tlsInsecure"] = o.TlsInsecure.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MSADConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"type",
		"name",
		"throttleDuration",
		"throttleParallelism",
		"hostname",
		"credentials",
		"baseDn",
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

	varMSADConnectorResponse := _MSADConnectorResponse{}

	err = json.Unmarshal(data, &varMSADConnectorResponse)

	if err != nil {
		return err
	}

	*o = MSADConnectorResponse(varMSADConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "throttleDuration")
		delete(additionalProperties, "throttleParallelism")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "maxStoredCertificatePerHolder")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "port")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "baseDn")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "tlsInsecure")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMSADConnectorResponse struct {
	value *MSADConnectorResponse
	isSet bool
}

func (v NullableMSADConnectorResponse) Get() *MSADConnectorResponse {
	return v.value
}

func (v *NullableMSADConnectorResponse) Set(val *MSADConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMSADConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMSADConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMSADConnectorResponse(val *MSADConnectorResponse) *NullableMSADConnectorResponse {
	return &NullableMSADConnectorResponse{value: val, isSet: true}
}

func (v NullableMSADConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMSADConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
