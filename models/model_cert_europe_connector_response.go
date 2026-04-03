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

// checks if the CertEuropeConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CertEuropeConnectorResponse{}

// CertEuropeConnectorResponse struct for CertEuropeConnectorResponse
type CertEuropeConnectorResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Name of the `certificate` [credentials](#tag/security.credentials) to use to authenticate on the PKI
	AuthenticationCredentials string `json:"authenticationCredentials"`
	EndPoint                  string `json:"endPoint"`
	// Name of the `password` [credentials](#tag/security.credentials) to use for technical account on the PKI
	LoginCredentials     string                     `json:"loginCredentials"`
	Name                 string                     `json:"name"`
	OfferId              string                     `json:"offerId"`
	OrganizationId       string                     `json:"organizationId"`
	Proxy                utils.NullableString       `json:"proxy,omitempty"`
	Queue                utils.NullableString       `json:"queue,omitempty"`
	RetryInterval        utils.NullableString       `json:"retryInterval,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	RevReason            utils.NullableString       `json:"revReason,omitempty"`
	Status               NullablePKIConnectorStatus `json:"status,omitempty"`
	Timeout              utils.NullableString       `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string                     `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _CertEuropeConnectorResponse CertEuropeConnectorResponse

// NewCertEuropeConnectorResponse instantiates a new CertEuropeConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertEuropeConnectorResponse(id string, authenticationCredentials string, endPoint string, loginCredentials string, name string, offerId string, organizationId string, type_ string) *CertEuropeConnectorResponse {
	this := CertEuropeConnectorResponse{}
	this.Id = id
	this.AuthenticationCredentials = authenticationCredentials
	this.EndPoint = endPoint
	this.LoginCredentials = loginCredentials
	this.Name = name
	this.OfferId = offerId
	this.OrganizationId = organizationId
	this.Type = type_
	return &this
}

// NewCertEuropeConnectorResponseWithDefaults instantiates a new CertEuropeConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertEuropeConnectorResponseWithDefaults() *CertEuropeConnectorResponse {
	this := CertEuropeConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CertEuropeConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CertEuropeConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CertEuropeConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value
func (o *CertEuropeConnectorResponse) GetAuthenticationCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationCredentials
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value
// and a boolean to check if the value has been set.
func (o *CertEuropeConnectorResponse) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationCredentials, true
}

// SetAuthenticationCredentials sets field value
func (o *CertEuropeConnectorResponse) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials = v
}

// GetEndPoint returns the EndPoint field value
func (o *CertEuropeConnectorResponse) GetEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value
// and a boolean to check if the value has been set.
func (o *CertEuropeConnectorResponse) GetEndPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPoint, true
}

// SetEndPoint sets field value
func (o *CertEuropeConnectorResponse) SetEndPoint(v string) {
	o.EndPoint = v
}

// GetLoginCredentials returns the LoginCredentials field value
func (o *CertEuropeConnectorResponse) GetLoginCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoginCredentials
}

// GetLoginCredentialsOk returns a tuple with the LoginCredentials field value
// and a boolean to check if the value has been set.
func (o *CertEuropeConnectorResponse) GetLoginCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginCredentials, true
}

// SetLoginCredentials sets field value
func (o *CertEuropeConnectorResponse) SetLoginCredentials(v string) {
	o.LoginCredentials = v
}

// GetName returns the Name field value
func (o *CertEuropeConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CertEuropeConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CertEuropeConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetOfferId returns the OfferId field value
func (o *CertEuropeConnectorResponse) GetOfferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value
// and a boolean to check if the value has been set.
func (o *CertEuropeConnectorResponse) GetOfferIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferId, true
}

// SetOfferId sets field value
func (o *CertEuropeConnectorResponse) SetOfferId(v string) {
	o.OfferId = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *CertEuropeConnectorResponse) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *CertEuropeConnectorResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *CertEuropeConnectorResponse) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertEuropeConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertEuropeConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *CertEuropeConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *CertEuropeConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *CertEuropeConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *CertEuropeConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertEuropeConnectorResponse) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertEuropeConnectorResponse) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *CertEuropeConnectorResponse) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *CertEuropeConnectorResponse) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *CertEuropeConnectorResponse) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *CertEuropeConnectorResponse) UnsetQueue() {
	o.Queue.Unset()
}

// GetRetryInterval returns the RetryInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertEuropeConnectorResponse) GetRetryInterval() string {
	if o == nil || utils.IsNil(o.RetryInterval.Get()) {
		var ret string
		return ret
	}
	return *o.RetryInterval.Get()
}

// GetRetryIntervalOk returns a tuple with the RetryInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertEuropeConnectorResponse) GetRetryIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetryInterval.Get(), o.RetryInterval.IsSet()
}

// HasRetryInterval returns a boolean if a field has been set.
func (o *CertEuropeConnectorResponse) HasRetryInterval() bool {
	if o != nil && o.RetryInterval.IsSet() {
		return true
	}

	return false
}

// SetRetryInterval gets a reference to the given NullableString and assigns it to the RetryInterval field.
func (o *CertEuropeConnectorResponse) SetRetryInterval(v string) {
	o.RetryInterval.Set(&v)
}

// SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil
func (o *CertEuropeConnectorResponse) SetRetryIntervalNil() {
	o.RetryInterval.Set(nil)
}

// UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
func (o *CertEuropeConnectorResponse) UnsetRetryInterval() {
	o.RetryInterval.Unset()
}

// GetRevReason returns the RevReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertEuropeConnectorResponse) GetRevReason() string {
	if o == nil || utils.IsNil(o.RevReason.Get()) {
		var ret string
		return ret
	}
	return *o.RevReason.Get()
}

// GetRevReasonOk returns a tuple with the RevReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertEuropeConnectorResponse) GetRevReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevReason.Get(), o.RevReason.IsSet()
}

// HasRevReason returns a boolean if a field has been set.
func (o *CertEuropeConnectorResponse) HasRevReason() bool {
	if o != nil && o.RevReason.IsSet() {
		return true
	}

	return false
}

// SetRevReason gets a reference to the given NullableString and assigns it to the RevReason field.
func (o *CertEuropeConnectorResponse) SetRevReason(v string) {
	o.RevReason.Set(&v)
}

// SetRevReasonNil sets the value for RevReason to be an explicit nil
func (o *CertEuropeConnectorResponse) SetRevReasonNil() {
	o.RevReason.Set(nil)
}

// UnsetRevReason ensures that no value is present for RevReason, not even an explicit nil
func (o *CertEuropeConnectorResponse) UnsetRevReason() {
	o.RevReason.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertEuropeConnectorResponse) GetStatus() PKIConnectorStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret PKIConnectorStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertEuropeConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *CertEuropeConnectorResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePKIConnectorStatus and assigns it to the Status field.
func (o *CertEuropeConnectorResponse) SetStatus(v PKIConnectorStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *CertEuropeConnectorResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *CertEuropeConnectorResponse) UnsetStatus() {
	o.Status.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertEuropeConnectorResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertEuropeConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *CertEuropeConnectorResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *CertEuropeConnectorResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *CertEuropeConnectorResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *CertEuropeConnectorResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *CertEuropeConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CertEuropeConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CertEuropeConnectorResponse) SetType(v string) {
	o.Type = v
}

func (o CertEuropeConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertEuropeConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["authenticationCredentials"] = o.AuthenticationCredentials
	toSerialize["endPoint"] = o.EndPoint
	toSerialize["loginCredentials"] = o.LoginCredentials
	toSerialize["name"] = o.Name
	toSerialize["offerId"] = o.OfferId
	toSerialize["organizationId"] = o.OrganizationId
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.RetryInterval.IsSet() {
		toSerialize["retryInterval"] = o.RetryInterval.Get()
	}
	if o.RevReason.IsSet() {
		toSerialize["revReason"] = o.RevReason.Get()
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

func (o *CertEuropeConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"authenticationCredentials",
		"endPoint",
		"loginCredentials",
		"name",
		"offerId",
		"organizationId",
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

	varCertEuropeConnectorResponse := _CertEuropeConnectorResponse{}

	err = json.Unmarshal(data, &varCertEuropeConnectorResponse)

	if err != nil {
		return err
	}

	*o = CertEuropeConnectorResponse(varCertEuropeConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "authenticationCredentials")
		delete(additionalProperties, "endPoint")
		delete(additionalProperties, "loginCredentials")
		delete(additionalProperties, "name")
		delete(additionalProperties, "offerId")
		delete(additionalProperties, "organizationId")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "retryInterval")
		delete(additionalProperties, "revReason")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertEuropeConnectorResponse struct {
	value *CertEuropeConnectorResponse
	isSet bool
}

func (v NullableCertEuropeConnectorResponse) Get() *CertEuropeConnectorResponse {
	return v.value
}

func (v *NullableCertEuropeConnectorResponse) Set(val *CertEuropeConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertEuropeConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertEuropeConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertEuropeConnectorResponse(val *CertEuropeConnectorResponse) *NullableCertEuropeConnectorResponse {
	return &NullableCertEuropeConnectorResponse{value: val, isSet: true}
}

func (v NullableCertEuropeConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertEuropeConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
