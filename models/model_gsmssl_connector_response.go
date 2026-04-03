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

// checks if the GSMSSLConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GSMSSLConnectorResponse{}

// GSMSSLConnectorResponse struct for GSMSSLConnectorResponse
type GSMSSLConnectorResponse struct {
	// Object internal ID
	Id                  string               `json:"_id"`
	CertificateValidity utils.NullableInt64  `json:"certificateValidity,omitempty"`
	DefaultEmail        utils.NullableString `json:"defaultEmail,omitempty"`
	DefaultPhone        utils.NullableString `json:"defaultPhone,omitempty"`
	DomainId            string               `json:"domainId"`
	EndpointType        string               `json:"endpointType"`
	// Name of the `password` [credentials](#tag/security.credentials) to use for technical account on the PKI
	LoginCredentials     string                     `json:"loginCredentials"`
	Name                 string                     `json:"name"`
	Profile              string                     `json:"profile"`
	Proxy                utils.NullableString       `json:"proxy,omitempty"`
	Queue                utils.NullableString       `json:"queue,omitempty"`
	RetryInterval        utils.NullableString       `json:"retryInterval,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Status               NullablePKIConnectorStatus `json:"status,omitempty"`
	Timeout              utils.NullableString       `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string                     `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _GSMSSLConnectorResponse GSMSSLConnectorResponse

// NewGSMSSLConnectorResponse instantiates a new GSMSSLConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGSMSSLConnectorResponse(id string, domainId string, endpointType string, loginCredentials string, name string, profile string, type_ string) *GSMSSLConnectorResponse {
	this := GSMSSLConnectorResponse{}
	this.Id = id
	this.DomainId = domainId
	this.EndpointType = endpointType
	this.LoginCredentials = loginCredentials
	this.Name = name
	this.Profile = profile
	this.Type = type_
	return &this
}

// NewGSMSSLConnectorResponseWithDefaults instantiates a new GSMSSLConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGSMSSLConnectorResponseWithDefaults() *GSMSSLConnectorResponse {
	this := GSMSSLConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GSMSSLConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GSMSSLConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GSMSSLConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetCertificateValidity returns the CertificateValidity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSMSSLConnectorResponse) GetCertificateValidity() int64 {
	if o == nil || utils.IsNil(o.CertificateValidity.Get()) {
		var ret int64
		return ret
	}
	return *o.CertificateValidity.Get()
}

// GetCertificateValidityOk returns a tuple with the CertificateValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSMSSLConnectorResponse) GetCertificateValidityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateValidity.Get(), o.CertificateValidity.IsSet()
}

// HasCertificateValidity returns a boolean if a field has been set.
func (o *GSMSSLConnectorResponse) HasCertificateValidity() bool {
	if o != nil && o.CertificateValidity.IsSet() {
		return true
	}

	return false
}

// SetCertificateValidity gets a reference to the given NullableInt64 and assigns it to the CertificateValidity field.
func (o *GSMSSLConnectorResponse) SetCertificateValidity(v int64) {
	o.CertificateValidity.Set(&v)
}

// SetCertificateValidityNil sets the value for CertificateValidity to be an explicit nil
func (o *GSMSSLConnectorResponse) SetCertificateValidityNil() {
	o.CertificateValidity.Set(nil)
}

// UnsetCertificateValidity ensures that no value is present for CertificateValidity, not even an explicit nil
func (o *GSMSSLConnectorResponse) UnsetCertificateValidity() {
	o.CertificateValidity.Unset()
}

// GetDefaultEmail returns the DefaultEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSMSSLConnectorResponse) GetDefaultEmail() string {
	if o == nil || utils.IsNil(o.DefaultEmail.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultEmail.Get()
}

// GetDefaultEmailOk returns a tuple with the DefaultEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSMSSLConnectorResponse) GetDefaultEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultEmail.Get(), o.DefaultEmail.IsSet()
}

// HasDefaultEmail returns a boolean if a field has been set.
func (o *GSMSSLConnectorResponse) HasDefaultEmail() bool {
	if o != nil && o.DefaultEmail.IsSet() {
		return true
	}

	return false
}

// SetDefaultEmail gets a reference to the given NullableString and assigns it to the DefaultEmail field.
func (o *GSMSSLConnectorResponse) SetDefaultEmail(v string) {
	o.DefaultEmail.Set(&v)
}

// SetDefaultEmailNil sets the value for DefaultEmail to be an explicit nil
func (o *GSMSSLConnectorResponse) SetDefaultEmailNil() {
	o.DefaultEmail.Set(nil)
}

// UnsetDefaultEmail ensures that no value is present for DefaultEmail, not even an explicit nil
func (o *GSMSSLConnectorResponse) UnsetDefaultEmail() {
	o.DefaultEmail.Unset()
}

// GetDefaultPhone returns the DefaultPhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSMSSLConnectorResponse) GetDefaultPhone() string {
	if o == nil || utils.IsNil(o.DefaultPhone.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultPhone.Get()
}

// GetDefaultPhoneOk returns a tuple with the DefaultPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSMSSLConnectorResponse) GetDefaultPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultPhone.Get(), o.DefaultPhone.IsSet()
}

// HasDefaultPhone returns a boolean if a field has been set.
func (o *GSMSSLConnectorResponse) HasDefaultPhone() bool {
	if o != nil && o.DefaultPhone.IsSet() {
		return true
	}

	return false
}

// SetDefaultPhone gets a reference to the given NullableString and assigns it to the DefaultPhone field.
func (o *GSMSSLConnectorResponse) SetDefaultPhone(v string) {
	o.DefaultPhone.Set(&v)
}

// SetDefaultPhoneNil sets the value for DefaultPhone to be an explicit nil
func (o *GSMSSLConnectorResponse) SetDefaultPhoneNil() {
	o.DefaultPhone.Set(nil)
}

// UnsetDefaultPhone ensures that no value is present for DefaultPhone, not even an explicit nil
func (o *GSMSSLConnectorResponse) UnsetDefaultPhone() {
	o.DefaultPhone.Unset()
}

// GetDomainId returns the DomainId field value
func (o *GSMSSLConnectorResponse) GetDomainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value
// and a boolean to check if the value has been set.
func (o *GSMSSLConnectorResponse) GetDomainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainId, true
}

// SetDomainId sets field value
func (o *GSMSSLConnectorResponse) SetDomainId(v string) {
	o.DomainId = v
}

// GetEndpointType returns the EndpointType field value
func (o *GSMSSLConnectorResponse) GetEndpointType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value
// and a boolean to check if the value has been set.
func (o *GSMSSLConnectorResponse) GetEndpointTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointType, true
}

// SetEndpointType sets field value
func (o *GSMSSLConnectorResponse) SetEndpointType(v string) {
	o.EndpointType = v
}

// GetLoginCredentials returns the LoginCredentials field value
func (o *GSMSSLConnectorResponse) GetLoginCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoginCredentials
}

// GetLoginCredentialsOk returns a tuple with the LoginCredentials field value
// and a boolean to check if the value has been set.
func (o *GSMSSLConnectorResponse) GetLoginCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginCredentials, true
}

// SetLoginCredentials sets field value
func (o *GSMSSLConnectorResponse) SetLoginCredentials(v string) {
	o.LoginCredentials = v
}

// GetName returns the Name field value
func (o *GSMSSLConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GSMSSLConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GSMSSLConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetProfile returns the Profile field value
func (o *GSMSSLConnectorResponse) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *GSMSSLConnectorResponse) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *GSMSSLConnectorResponse) SetProfile(v string) {
	o.Profile = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSMSSLConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSMSSLConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *GSMSSLConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *GSMSSLConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *GSMSSLConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *GSMSSLConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSMSSLConnectorResponse) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSMSSLConnectorResponse) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *GSMSSLConnectorResponse) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *GSMSSLConnectorResponse) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *GSMSSLConnectorResponse) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *GSMSSLConnectorResponse) UnsetQueue() {
	o.Queue.Unset()
}

// GetRetryInterval returns the RetryInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSMSSLConnectorResponse) GetRetryInterval() string {
	if o == nil || utils.IsNil(o.RetryInterval.Get()) {
		var ret string
		return ret
	}
	return *o.RetryInterval.Get()
}

// GetRetryIntervalOk returns a tuple with the RetryInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSMSSLConnectorResponse) GetRetryIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetryInterval.Get(), o.RetryInterval.IsSet()
}

// HasRetryInterval returns a boolean if a field has been set.
func (o *GSMSSLConnectorResponse) HasRetryInterval() bool {
	if o != nil && o.RetryInterval.IsSet() {
		return true
	}

	return false
}

// SetRetryInterval gets a reference to the given NullableString and assigns it to the RetryInterval field.
func (o *GSMSSLConnectorResponse) SetRetryInterval(v string) {
	o.RetryInterval.Set(&v)
}

// SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil
func (o *GSMSSLConnectorResponse) SetRetryIntervalNil() {
	o.RetryInterval.Set(nil)
}

// UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
func (o *GSMSSLConnectorResponse) UnsetRetryInterval() {
	o.RetryInterval.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSMSSLConnectorResponse) GetStatus() PKIConnectorStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret PKIConnectorStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSMSSLConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *GSMSSLConnectorResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePKIConnectorStatus and assigns it to the Status field.
func (o *GSMSSLConnectorResponse) SetStatus(v PKIConnectorStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *GSMSSLConnectorResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *GSMSSLConnectorResponse) UnsetStatus() {
	o.Status.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GSMSSLConnectorResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GSMSSLConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *GSMSSLConnectorResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *GSMSSLConnectorResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *GSMSSLConnectorResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *GSMSSLConnectorResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *GSMSSLConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GSMSSLConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GSMSSLConnectorResponse) SetType(v string) {
	o.Type = v
}

func (o GSMSSLConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GSMSSLConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.CertificateValidity.IsSet() {
		toSerialize["certificateValidity"] = o.CertificateValidity.Get()
	}
	if o.DefaultEmail.IsSet() {
		toSerialize["defaultEmail"] = o.DefaultEmail.Get()
	}
	if o.DefaultPhone.IsSet() {
		toSerialize["defaultPhone"] = o.DefaultPhone.Get()
	}
	toSerialize["domainId"] = o.DomainId
	toSerialize["endpointType"] = o.EndpointType
	toSerialize["loginCredentials"] = o.LoginCredentials
	toSerialize["name"] = o.Name
	toSerialize["profile"] = o.Profile
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.RetryInterval.IsSet() {
		toSerialize["retryInterval"] = o.RetryInterval.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
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

func (o *GSMSSLConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"domainId",
		"endpointType",
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

	varGSMSSLConnectorResponse := _GSMSSLConnectorResponse{}

	err = json.Unmarshal(data, &varGSMSSLConnectorResponse)

	if err != nil {
		return err
	}

	*o = GSMSSLConnectorResponse(varGSMSSLConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "certificateValidity")
		delete(additionalProperties, "defaultEmail")
		delete(additionalProperties, "defaultPhone")
		delete(additionalProperties, "domainId")
		delete(additionalProperties, "endpointType")
		delete(additionalProperties, "loginCredentials")
		delete(additionalProperties, "name")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "retryInterval")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGSMSSLConnectorResponse struct {
	value *GSMSSLConnectorResponse
	isSet bool
}

func (v NullableGSMSSLConnectorResponse) Get() *GSMSSLConnectorResponse {
	return v.value
}

func (v *NullableGSMSSLConnectorResponse) Set(val *GSMSSLConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGSMSSLConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGSMSSLConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGSMSSLConnectorResponse(val *GSMSSLConnectorResponse) *NullableGSMSSLConnectorResponse {
	return &NullableGSMSSLConnectorResponse{value: val, isSet: true}
}

func (v NullableGSMSSLConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGSMSSLConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
