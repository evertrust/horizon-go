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

// checks if the EntrustConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EntrustConnector{}

// EntrustConnector struct for EntrustConnector
type EntrustConnector struct {
	// Name of the `certificate` [credentials](#tag/security.credentials) to use to authenticate on the PKI
	AuthenticationCredentials string               `json:"authenticationCredentials"`
	CertLifetime              utils.NullableString `json:"certLifetime,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	CertType                  string               `json:"certType"`
	ClientId                  utils.NullableInt64  `json:"clientId,omitempty"`
	// Name of the `password` [credentials](#tag/security.credentials) to use for technical account on the PKI
	LoginCredentials     string               `json:"loginCredentials"`
	Name                 string               `json:"name"`
	Proxy                utils.NullableString `json:"proxy,omitempty"`
	Queue                utils.NullableString `json:"queue,omitempty"`
	RequesterDefaultMail string               `json:"requesterDefaultMail"`
	RequesterName        utils.NullableString `json:"requesterName,omitempty"`
	RequesterPhone       utils.NullableString `json:"requesterPhone,omitempty"`
	Timeout              utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string               `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _EntrustConnector EntrustConnector

// NewEntrustConnector instantiates a new EntrustConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntrustConnector(authenticationCredentials string, certType string, loginCredentials string, name string, requesterDefaultMail string, type_ string) *EntrustConnector {
	this := EntrustConnector{}
	this.AuthenticationCredentials = authenticationCredentials
	this.CertType = certType
	this.LoginCredentials = loginCredentials
	this.Name = name
	this.RequesterDefaultMail = requesterDefaultMail
	this.Type = type_
	return &this
}

// NewEntrustConnectorWithDefaults instantiates a new EntrustConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntrustConnectorWithDefaults() *EntrustConnector {
	this := EntrustConnector{}
	return &this
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value
func (o *EntrustConnector) GetAuthenticationCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationCredentials
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value
// and a boolean to check if the value has been set.
func (o *EntrustConnector) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationCredentials, true
}

// SetAuthenticationCredentials sets field value
func (o *EntrustConnector) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials = v
}

// GetCertLifetime returns the CertLifetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntrustConnector) GetCertLifetime() string {
	if o == nil || utils.IsNil(o.CertLifetime.Get()) {
		var ret string
		return ret
	}
	return *o.CertLifetime.Get()
}

// GetCertLifetimeOk returns a tuple with the CertLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntrustConnector) GetCertLifetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertLifetime.Get(), o.CertLifetime.IsSet()
}

// HasCertLifetime returns a boolean if a field has been set.
func (o *EntrustConnector) HasCertLifetime() bool {
	if o != nil && o.CertLifetime.IsSet() {
		return true
	}

	return false
}

// SetCertLifetime gets a reference to the given NullableString and assigns it to the CertLifetime field.
func (o *EntrustConnector) SetCertLifetime(v string) {
	o.CertLifetime.Set(&v)
}

// SetCertLifetimeNil sets the value for CertLifetime to be an explicit nil
func (o *EntrustConnector) SetCertLifetimeNil() {
	o.CertLifetime.Set(nil)
}

// UnsetCertLifetime ensures that no value is present for CertLifetime, not even an explicit nil
func (o *EntrustConnector) UnsetCertLifetime() {
	o.CertLifetime.Unset()
}

// GetCertType returns the CertType field value
func (o *EntrustConnector) GetCertType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertType
}

// GetCertTypeOk returns a tuple with the CertType field value
// and a boolean to check if the value has been set.
func (o *EntrustConnector) GetCertTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertType, true
}

// SetCertType sets field value
func (o *EntrustConnector) SetCertType(v string) {
	o.CertType = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntrustConnector) GetClientId() int64 {
	if o == nil || utils.IsNil(o.ClientId.Get()) {
		var ret int64
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntrustConnector) GetClientIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *EntrustConnector) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableInt64 and assigns it to the ClientId field.
func (o *EntrustConnector) SetClientId(v int64) {
	o.ClientId.Set(&v)
}

// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *EntrustConnector) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *EntrustConnector) UnsetClientId() {
	o.ClientId.Unset()
}

// GetLoginCredentials returns the LoginCredentials field value
func (o *EntrustConnector) GetLoginCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoginCredentials
}

// GetLoginCredentialsOk returns a tuple with the LoginCredentials field value
// and a boolean to check if the value has been set.
func (o *EntrustConnector) GetLoginCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginCredentials, true
}

// SetLoginCredentials sets field value
func (o *EntrustConnector) SetLoginCredentials(v string) {
	o.LoginCredentials = v
}

// GetName returns the Name field value
func (o *EntrustConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntrustConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntrustConnector) SetName(v string) {
	o.Name = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntrustConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntrustConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *EntrustConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *EntrustConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *EntrustConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *EntrustConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntrustConnector) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntrustConnector) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *EntrustConnector) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *EntrustConnector) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *EntrustConnector) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *EntrustConnector) UnsetQueue() {
	o.Queue.Unset()
}

// GetRequesterDefaultMail returns the RequesterDefaultMail field value
func (o *EntrustConnector) GetRequesterDefaultMail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequesterDefaultMail
}

// GetRequesterDefaultMailOk returns a tuple with the RequesterDefaultMail field value
// and a boolean to check if the value has been set.
func (o *EntrustConnector) GetRequesterDefaultMailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequesterDefaultMail, true
}

// SetRequesterDefaultMail sets field value
func (o *EntrustConnector) SetRequesterDefaultMail(v string) {
	o.RequesterDefaultMail = v
}

// GetRequesterName returns the RequesterName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntrustConnector) GetRequesterName() string {
	if o == nil || utils.IsNil(o.RequesterName.Get()) {
		var ret string
		return ret
	}
	return *o.RequesterName.Get()
}

// GetRequesterNameOk returns a tuple with the RequesterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntrustConnector) GetRequesterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterName.Get(), o.RequesterName.IsSet()
}

// HasRequesterName returns a boolean if a field has been set.
func (o *EntrustConnector) HasRequesterName() bool {
	if o != nil && o.RequesterName.IsSet() {
		return true
	}

	return false
}

// SetRequesterName gets a reference to the given NullableString and assigns it to the RequesterName field.
func (o *EntrustConnector) SetRequesterName(v string) {
	o.RequesterName.Set(&v)
}

// SetRequesterNameNil sets the value for RequesterName to be an explicit nil
func (o *EntrustConnector) SetRequesterNameNil() {
	o.RequesterName.Set(nil)
}

// UnsetRequesterName ensures that no value is present for RequesterName, not even an explicit nil
func (o *EntrustConnector) UnsetRequesterName() {
	o.RequesterName.Unset()
}

// GetRequesterPhone returns the RequesterPhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntrustConnector) GetRequesterPhone() string {
	if o == nil || utils.IsNil(o.RequesterPhone.Get()) {
		var ret string
		return ret
	}
	return *o.RequesterPhone.Get()
}

// GetRequesterPhoneOk returns a tuple with the RequesterPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntrustConnector) GetRequesterPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterPhone.Get(), o.RequesterPhone.IsSet()
}

// HasRequesterPhone returns a boolean if a field has been set.
func (o *EntrustConnector) HasRequesterPhone() bool {
	if o != nil && o.RequesterPhone.IsSet() {
		return true
	}

	return false
}

// SetRequesterPhone gets a reference to the given NullableString and assigns it to the RequesterPhone field.
func (o *EntrustConnector) SetRequesterPhone(v string) {
	o.RequesterPhone.Set(&v)
}

// SetRequesterPhoneNil sets the value for RequesterPhone to be an explicit nil
func (o *EntrustConnector) SetRequesterPhoneNil() {
	o.RequesterPhone.Set(nil)
}

// UnsetRequesterPhone ensures that no value is present for RequesterPhone, not even an explicit nil
func (o *EntrustConnector) UnsetRequesterPhone() {
	o.RequesterPhone.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntrustConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntrustConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *EntrustConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *EntrustConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *EntrustConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *EntrustConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *EntrustConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntrustConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EntrustConnector) SetType(v string) {
	o.Type = v
}

func (o EntrustConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntrustConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authenticationCredentials"] = o.AuthenticationCredentials
	if o.CertLifetime.IsSet() {
		toSerialize["certLifetime"] = o.CertLifetime.Get()
	}
	toSerialize["certType"] = o.CertType
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	toSerialize["loginCredentials"] = o.LoginCredentials
	toSerialize["name"] = o.Name
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	toSerialize["requesterDefaultMail"] = o.RequesterDefaultMail
	if o.RequesterName.IsSet() {
		toSerialize["requesterName"] = o.RequesterName.Get()
	}
	if o.RequesterPhone.IsSet() {
		toSerialize["requesterPhone"] = o.RequesterPhone.Get()
	}
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntrustConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authenticationCredentials",
		"certType",
		"loginCredentials",
		"name",
		"requesterDefaultMail",
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

	varEntrustConnector := _EntrustConnector{}

	err = json.Unmarshal(data, &varEntrustConnector)

	if err != nil {
		return err
	}

	*o = EntrustConnector(varEntrustConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticationCredentials")
		delete(additionalProperties, "certLifetime")
		delete(additionalProperties, "certType")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "loginCredentials")
		delete(additionalProperties, "name")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "requesterDefaultMail")
		delete(additionalProperties, "requesterName")
		delete(additionalProperties, "requesterPhone")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntrustConnector struct {
	value *EntrustConnector
	isSet bool
}

func (v NullableEntrustConnector) Get() *EntrustConnector {
	return v.value
}

func (v *NullableEntrustConnector) Set(val *EntrustConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableEntrustConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableEntrustConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntrustConnector(val *EntrustConnector) *NullableEntrustConnector {
	return &NullableEntrustConnector{value: val, isSet: true}
}

func (v NullableEntrustConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntrustConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
