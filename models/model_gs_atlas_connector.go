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

// checks if the GSAtlasConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GSAtlasConnector{}

// GSAtlasConnector struct for GSAtlasConnector
type GSAtlasConnector struct {
	Name string `json:"name"`
	Type string `json:"type"`
	// Name of the `password` [credentials](#tag/security.credentials) to use for technical account on the PKI
	LoginCredentials string               `json:"loginCredentials"`
	HashAlgorithm    utils.NullableString `json:"hashAlgorithm,omitempty"`
	CertificateUsage utils.NullableString `json:"certificateUsage,omitempty"`
	RetryInterval    utils.NullableString `json:"retryInterval,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Name of the `certificate` [credentials](#tag/security.credentials) to use to authenticate on the PKI
	AuthenticationCredentials string               `json:"authenticationCredentials"`
	Timeout                   utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Proxy                     utils.NullableString `json:"proxy,omitempty"`
	Queue                     utils.NullableString `json:"queue,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _GSAtlasConnector GSAtlasConnector

// NewGSAtlasConnector instantiates a new GSAtlasConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGSAtlasConnector(name string, type_ string, loginCredentials string, authenticationCredentials string) *GSAtlasConnector {
	this := GSAtlasConnector{}
	this.Name = name
	this.Type = type_
	this.LoginCredentials = loginCredentials
	this.AuthenticationCredentials = authenticationCredentials
	return &this
}

// NewGSAtlasConnectorWithDefaults instantiates a new GSAtlasConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGSAtlasConnectorWithDefaults() *GSAtlasConnector {
	this := GSAtlasConnector{}
	return &this
}

// GetName returns the Name field value
func (o *GSAtlasConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GSAtlasConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GSAtlasConnector) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *GSAtlasConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GSAtlasConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GSAtlasConnector) SetType(v string) {
	o.Type = v
}

// GetLoginCredentials returns the LoginCredentials field value
func (o *GSAtlasConnector) GetLoginCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoginCredentials
}

// GetLoginCredentialsOk returns a tuple with the LoginCredentials field value
// and a boolean to check if the value has been set.
func (o *GSAtlasConnector) GetLoginCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginCredentials, true
}

// SetLoginCredentials sets field value
func (o *GSAtlasConnector) SetLoginCredentials(v string) {
	o.LoginCredentials = v
}

// GetHashAlgorithm returns the HashAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSAtlasConnector) GetHashAlgorithm() string {
	if o == nil || utils.IsNil(o.HashAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.HashAlgorithm.Get()
}

// GetHashAlgorithmOk returns a tuple with the HashAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSAtlasConnector) GetHashAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HashAlgorithm.Get(), o.HashAlgorithm.IsSet()
}

// HasHashAlgorithm returns a boolean if a field has been set.
func (o *GSAtlasConnector) HasHashAlgorithm() bool {
	if o != nil && o.HashAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetHashAlgorithm gets a reference to the given NullableString and assigns it to the HashAlgorithm field.
func (o *GSAtlasConnector) SetHashAlgorithm(v string) {
	o.HashAlgorithm.Set(&v)
}

// SetHashAlgorithmNil sets the value for HashAlgorithm to be an explicit nil
func (o *GSAtlasConnector) SetHashAlgorithmNil() {
	o.HashAlgorithm.Set(nil)
}

// UnsetHashAlgorithm ensures that no value is present for HashAlgorithm, not even an explicit nil
func (o *GSAtlasConnector) UnsetHashAlgorithm() {
	o.HashAlgorithm.Unset()
}

// GetCertificateUsage returns the CertificateUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSAtlasConnector) GetCertificateUsage() string {
	if o == nil || utils.IsNil(o.CertificateUsage.Get()) {
		var ret string
		return ret
	}
	return *o.CertificateUsage.Get()
}

// GetCertificateUsageOk returns a tuple with the CertificateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSAtlasConnector) GetCertificateUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateUsage.Get(), o.CertificateUsage.IsSet()
}

// HasCertificateUsage returns a boolean if a field has been set.
func (o *GSAtlasConnector) HasCertificateUsage() bool {
	if o != nil && o.CertificateUsage.IsSet() {
		return true
	}

	return false
}

// SetCertificateUsage gets a reference to the given NullableString and assigns it to the CertificateUsage field.
func (o *GSAtlasConnector) SetCertificateUsage(v string) {
	o.CertificateUsage.Set(&v)
}

// SetCertificateUsageNil sets the value for CertificateUsage to be an explicit nil
func (o *GSAtlasConnector) SetCertificateUsageNil() {
	o.CertificateUsage.Set(nil)
}

// UnsetCertificateUsage ensures that no value is present for CertificateUsage, not even an explicit nil
func (o *GSAtlasConnector) UnsetCertificateUsage() {
	o.CertificateUsage.Unset()
}

// GetRetryInterval returns the RetryInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSAtlasConnector) GetRetryInterval() string {
	if o == nil || utils.IsNil(o.RetryInterval.Get()) {
		var ret string
		return ret
	}
	return *o.RetryInterval.Get()
}

// GetRetryIntervalOk returns a tuple with the RetryInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSAtlasConnector) GetRetryIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetryInterval.Get(), o.RetryInterval.IsSet()
}

// HasRetryInterval returns a boolean if a field has been set.
func (o *GSAtlasConnector) HasRetryInterval() bool {
	if o != nil && o.RetryInterval.IsSet() {
		return true
	}

	return false
}

// SetRetryInterval gets a reference to the given NullableString and assigns it to the RetryInterval field.
func (o *GSAtlasConnector) SetRetryInterval(v string) {
	o.RetryInterval.Set(&v)
}

// SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil
func (o *GSAtlasConnector) SetRetryIntervalNil() {
	o.RetryInterval.Set(nil)
}

// UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
func (o *GSAtlasConnector) UnsetRetryInterval() {
	o.RetryInterval.Unset()
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value
func (o *GSAtlasConnector) GetAuthenticationCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationCredentials
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value
// and a boolean to check if the value has been set.
func (o *GSAtlasConnector) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationCredentials, true
}

// SetAuthenticationCredentials sets field value
func (o *GSAtlasConnector) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSAtlasConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSAtlasConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *GSAtlasConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *GSAtlasConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *GSAtlasConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *GSAtlasConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSAtlasConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSAtlasConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *GSAtlasConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *GSAtlasConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *GSAtlasConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *GSAtlasConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSAtlasConnector) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSAtlasConnector) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *GSAtlasConnector) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *GSAtlasConnector) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *GSAtlasConnector) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *GSAtlasConnector) UnsetQueue() {
	o.Queue.Unset()
}

func (o GSAtlasConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GSAtlasConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["loginCredentials"] = o.LoginCredentials
	if o.HashAlgorithm.IsSet() {
		toSerialize["hashAlgorithm"] = o.HashAlgorithm.Get()
	}
	if o.CertificateUsage.IsSet() {
		toSerialize["certificateUsage"] = o.CertificateUsage.Get()
	}
	if o.RetryInterval.IsSet() {
		toSerialize["retryInterval"] = o.RetryInterval.Get()
	}
	toSerialize["authenticationCredentials"] = o.AuthenticationCredentials
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GSAtlasConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"loginCredentials",
		"authenticationCredentials",
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

	varGSAtlasConnector := _GSAtlasConnector{}

	err = json.Unmarshal(data, &varGSAtlasConnector)

	if err != nil {
		return err
	}

	*o = GSAtlasConnector(varGSAtlasConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "loginCredentials")
		delete(additionalProperties, "hashAlgorithm")
		delete(additionalProperties, "certificateUsage")
		delete(additionalProperties, "retryInterval")
		delete(additionalProperties, "authenticationCredentials")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGSAtlasConnector struct {
	value *GSAtlasConnector
	isSet bool
}

func (v NullableGSAtlasConnector) Get() *GSAtlasConnector {
	return v.value
}

func (v *NullableGSAtlasConnector) Set(val *GSAtlasConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableGSAtlasConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableGSAtlasConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGSAtlasConnector(val *GSAtlasConnector) *NullableGSAtlasConnector {
	return &NullableGSAtlasConnector{value: val, isSet: true}
}

func (v NullableGSAtlasConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGSAtlasConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
