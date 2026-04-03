/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the EverTrustADCSConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EverTrustADCSConnector{}

// EverTrustADCSConnector struct for EverTrustADCSConnector
type EverTrustADCSConnector struct {
	CaConfig string `json:"caConfig"`
	Domain   string `json:"domain"`
	EndPoint string `json:"endPoint"`
	// Name of the `certificate` [credentials](#tag/security.credentials) to use to enroll on the PKI
	EnrollmentCredentials string `json:"enrollmentCredentials"`
	// Name of the `password` [credentials](#tag/security.credentials) to use for technical account on the PKI
	LoginCredentials     string               `json:"loginCredentials"`
	Name                 string               `json:"name"`
	Profile              string               `json:"profile"`
	Proxy                utils.NullableString `json:"proxy,omitempty"`
	Queue                utils.NullableString `json:"queue,omitempty"`
	Timeout              utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string               `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _EverTrustADCSConnector EverTrustADCSConnector

// NewEverTrustADCSConnector instantiates a new EverTrustADCSConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEverTrustADCSConnector(caConfig string, domain string, endPoint string, enrollmentCredentials string, loginCredentials string, name string, profile string, type_ string) *EverTrustADCSConnector {
	this := EverTrustADCSConnector{}
	this.CaConfig = caConfig
	this.Domain = domain
	this.EndPoint = endPoint
	this.EnrollmentCredentials = enrollmentCredentials
	this.LoginCredentials = loginCredentials
	this.Name = name
	this.Profile = profile
	this.Type = type_
	return &this
}

// NewEverTrustADCSConnectorWithDefaults instantiates a new EverTrustADCSConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEverTrustADCSConnectorWithDefaults() *EverTrustADCSConnector {
	this := EverTrustADCSConnector{}
	return &this
}

// GetCaConfig returns the CaConfig field value
func (o *EverTrustADCSConnector) GetCaConfig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaConfig
}

// GetCaConfigOk returns a tuple with the CaConfig field value
// and a boolean to check if the value has been set.
func (o *EverTrustADCSConnector) GetCaConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaConfig, true
}

// SetCaConfig sets field value
func (o *EverTrustADCSConnector) SetCaConfig(v string) {
	o.CaConfig = v
}

// GetDomain returns the Domain field value
func (o *EverTrustADCSConnector) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *EverTrustADCSConnector) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *EverTrustADCSConnector) SetDomain(v string) {
	o.Domain = v
}

// GetEndPoint returns the EndPoint field value
func (o *EverTrustADCSConnector) GetEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value
// and a boolean to check if the value has been set.
func (o *EverTrustADCSConnector) GetEndPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPoint, true
}

// SetEndPoint sets field value
func (o *EverTrustADCSConnector) SetEndPoint(v string) {
	o.EndPoint = v
}

// GetEnrollmentCredentials returns the EnrollmentCredentials field value
func (o *EverTrustADCSConnector) GetEnrollmentCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnrollmentCredentials
}

// GetEnrollmentCredentialsOk returns a tuple with the EnrollmentCredentials field value
// and a boolean to check if the value has been set.
func (o *EverTrustADCSConnector) GetEnrollmentCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnrollmentCredentials, true
}

// SetEnrollmentCredentials sets field value
func (o *EverTrustADCSConnector) SetEnrollmentCredentials(v string) {
	o.EnrollmentCredentials = v
}

// GetLoginCredentials returns the LoginCredentials field value
func (o *EverTrustADCSConnector) GetLoginCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoginCredentials
}

// GetLoginCredentialsOk returns a tuple with the LoginCredentials field value
// and a boolean to check if the value has been set.
func (o *EverTrustADCSConnector) GetLoginCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginCredentials, true
}

// SetLoginCredentials sets field value
func (o *EverTrustADCSConnector) SetLoginCredentials(v string) {
	o.LoginCredentials = v
}

// GetName returns the Name field value
func (o *EverTrustADCSConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EverTrustADCSConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EverTrustADCSConnector) SetName(v string) {
	o.Name = v
}

// GetProfile returns the Profile field value
func (o *EverTrustADCSConnector) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *EverTrustADCSConnector) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *EverTrustADCSConnector) SetProfile(v string) {
	o.Profile = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EverTrustADCSConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EverTrustADCSConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *EverTrustADCSConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *EverTrustADCSConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *EverTrustADCSConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *EverTrustADCSConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EverTrustADCSConnector) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EverTrustADCSConnector) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *EverTrustADCSConnector) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *EverTrustADCSConnector) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *EverTrustADCSConnector) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *EverTrustADCSConnector) UnsetQueue() {
	o.Queue.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EverTrustADCSConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EverTrustADCSConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *EverTrustADCSConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *EverTrustADCSConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *EverTrustADCSConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *EverTrustADCSConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *EverTrustADCSConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EverTrustADCSConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EverTrustADCSConnector) SetType(v string) {
	o.Type = v
}

func (o EverTrustADCSConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EverTrustADCSConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["caConfig"] = o.CaConfig
	toSerialize["domain"] = o.Domain
	toSerialize["endPoint"] = o.EndPoint
	toSerialize["enrollmentCredentials"] = o.EnrollmentCredentials
	toSerialize["loginCredentials"] = o.LoginCredentials
	toSerialize["name"] = o.Name
	toSerialize["profile"] = o.Profile
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
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

func (o *EverTrustADCSConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"caConfig",
		"domain",
		"endPoint",
		"enrollmentCredentials",
		"loginCredentials",
		"name",
		"profile",
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

	varEverTrustADCSConnector := _EverTrustADCSConnector{}

	err = json.Unmarshal(data, &varEverTrustADCSConnector)

	if err != nil {
		return err
	}

	*o = EverTrustADCSConnector(varEverTrustADCSConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "caConfig")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "endPoint")
		delete(additionalProperties, "enrollmentCredentials")
		delete(additionalProperties, "loginCredentials")
		delete(additionalProperties, "name")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEverTrustADCSConnector struct {
	value *EverTrustADCSConnector
	isSet bool
}

func (v NullableEverTrustADCSConnector) Get() *EverTrustADCSConnector {
	return v.value
}

func (v *NullableEverTrustADCSConnector) Set(val *EverTrustADCSConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableEverTrustADCSConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableEverTrustADCSConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEverTrustADCSConnector(val *EverTrustADCSConnector) *NullableEverTrustADCSConnector {
	return &NullableEverTrustADCSConnector{value: val, isSet: true}
}

func (v NullableEverTrustADCSConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEverTrustADCSConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
