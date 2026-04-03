/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the DigiCertConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DigiCertConnectorResponse{}

// DigiCertConnectorResponse struct for DigiCertConnectorResponse
type DigiCertConnectorResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Name of the `raw` [credentials](#tag/security.credentials) containing the API key to authenticate on the PKI
	ApiCredentials string `json:"apiCredentials"`
	// The base URL of the used digicert instance.
	BaseUrl                    string               `json:"baseUrl"`
	CaCertId                   utils.NullableString `json:"caCertId,omitempty"`
	CustomConnectorDataMapping map[string]string    `json:"customConnectorDataMapping,omitempty"`
	Name                       string               `json:"name"`
	OrganizationId             int64                `json:"organizationId"`
	// One of the DigiCert product identifier an exhaustive list can be found here: https://dev.digicert.com/en/certcentral-apis/services-api/glossary.html#product-identifiers
	ProductId            *string                    `json:"productId,omitempty"`
	Proxy                utils.NullableString       `json:"proxy,omitempty"`
	Queue                utils.NullableString       `json:"queue,omitempty"`
	RetryInterval        utils.NullableString       `json:"retryInterval,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	SkipApproval         utils.NullableBool         `json:"skipApproval,omitempty"`
	Status               NullablePKIConnectorStatus `json:"status,omitempty"`
	Timeout              utils.NullableString       `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string                     `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _DigiCertConnectorResponse DigiCertConnectorResponse

// NewDigiCertConnectorResponse instantiates a new DigiCertConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigiCertConnectorResponse(id string, apiCredentials string, baseUrl string, name string, organizationId int64, type_ string) *DigiCertConnectorResponse {
	this := DigiCertConnectorResponse{}
	this.Id = id
	this.ApiCredentials = apiCredentials
	this.BaseUrl = baseUrl
	this.Name = name
	this.OrganizationId = organizationId
	this.Type = type_
	return &this
}

// NewDigiCertConnectorResponseWithDefaults instantiates a new DigiCertConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigiCertConnectorResponseWithDefaults() *DigiCertConnectorResponse {
	this := DigiCertConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DigiCertConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DigiCertConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DigiCertConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetApiCredentials returns the ApiCredentials field value
func (o *DigiCertConnectorResponse) GetApiCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiCredentials
}

// GetApiCredentialsOk returns a tuple with the ApiCredentials field value
// and a boolean to check if the value has been set.
func (o *DigiCertConnectorResponse) GetApiCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiCredentials, true
}

// SetApiCredentials sets field value
func (o *DigiCertConnectorResponse) SetApiCredentials(v string) {
	o.ApiCredentials = v
}

// GetBaseUrl returns the BaseUrl field value
func (o *DigiCertConnectorResponse) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *DigiCertConnectorResponse) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *DigiCertConnectorResponse) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetCaCertId returns the CaCertId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigiCertConnectorResponse) GetCaCertId() string {
	if o == nil || utils.IsNil(o.CaCertId.Get()) {
		var ret string
		return ret
	}
	return *o.CaCertId.Get()
}

// GetCaCertIdOk returns a tuple with the CaCertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigiCertConnectorResponse) GetCaCertIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCertId.Get(), o.CaCertId.IsSet()
}

// HasCaCertId returns a boolean if a field has been set.
func (o *DigiCertConnectorResponse) HasCaCertId() bool {
	if o != nil && o.CaCertId.IsSet() {
		return true
	}

	return false
}

// SetCaCertId gets a reference to the given NullableString and assigns it to the CaCertId field.
func (o *DigiCertConnectorResponse) SetCaCertId(v string) {
	o.CaCertId.Set(&v)
}

// SetCaCertIdNil sets the value for CaCertId to be an explicit nil
func (o *DigiCertConnectorResponse) SetCaCertIdNil() {
	o.CaCertId.Set(nil)
}

// UnsetCaCertId ensures that no value is present for CaCertId, not even an explicit nil
func (o *DigiCertConnectorResponse) UnsetCaCertId() {
	o.CaCertId.Unset()
}

// GetCustomConnectorDataMapping returns the CustomConnectorDataMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigiCertConnectorResponse) GetCustomConnectorDataMapping() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.CustomConnectorDataMapping
}

// GetCustomConnectorDataMappingOk returns a tuple with the CustomConnectorDataMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigiCertConnectorResponse) GetCustomConnectorDataMappingOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.CustomConnectorDataMapping) {
		return nil, false
	}
	return &o.CustomConnectorDataMapping, true
}

// HasCustomConnectorDataMapping returns a boolean if a field has been set.
func (o *DigiCertConnectorResponse) HasCustomConnectorDataMapping() bool {
	if o != nil && !utils.IsNil(o.CustomConnectorDataMapping) {
		return true
	}

	return false
}

// SetCustomConnectorDataMapping gets a reference to the given map[string]string and assigns it to the CustomConnectorDataMapping field.
func (o *DigiCertConnectorResponse) SetCustomConnectorDataMapping(v map[string]string) {
	o.CustomConnectorDataMapping = v
}

// GetName returns the Name field value
func (o *DigiCertConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DigiCertConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DigiCertConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *DigiCertConnectorResponse) GetOrganizationId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *DigiCertConnectorResponse) GetOrganizationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *DigiCertConnectorResponse) SetOrganizationId(v int64) {
	o.OrganizationId = v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *DigiCertConnectorResponse) GetProductId() string {
	if o == nil || utils.IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigiCertConnectorResponse) GetProductIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *DigiCertConnectorResponse) HasProductId() bool {
	if o != nil && !utils.IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *DigiCertConnectorResponse) SetProductId(v string) {
	o.ProductId = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigiCertConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigiCertConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *DigiCertConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *DigiCertConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *DigiCertConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *DigiCertConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigiCertConnectorResponse) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigiCertConnectorResponse) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *DigiCertConnectorResponse) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *DigiCertConnectorResponse) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *DigiCertConnectorResponse) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *DigiCertConnectorResponse) UnsetQueue() {
	o.Queue.Unset()
}

// GetRetryInterval returns the RetryInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigiCertConnectorResponse) GetRetryInterval() string {
	if o == nil || utils.IsNil(o.RetryInterval.Get()) {
		var ret string
		return ret
	}
	return *o.RetryInterval.Get()
}

// GetRetryIntervalOk returns a tuple with the RetryInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigiCertConnectorResponse) GetRetryIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetryInterval.Get(), o.RetryInterval.IsSet()
}

// HasRetryInterval returns a boolean if a field has been set.
func (o *DigiCertConnectorResponse) HasRetryInterval() bool {
	if o != nil && o.RetryInterval.IsSet() {
		return true
	}

	return false
}

// SetRetryInterval gets a reference to the given NullableString and assigns it to the RetryInterval field.
func (o *DigiCertConnectorResponse) SetRetryInterval(v string) {
	o.RetryInterval.Set(&v)
}

// SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil
func (o *DigiCertConnectorResponse) SetRetryIntervalNil() {
	o.RetryInterval.Set(nil)
}

// UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
func (o *DigiCertConnectorResponse) UnsetRetryInterval() {
	o.RetryInterval.Unset()
}

// GetSkipApproval returns the SkipApproval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigiCertConnectorResponse) GetSkipApproval() bool {
	if o == nil || utils.IsNil(o.SkipApproval.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipApproval.Get()
}

// GetSkipApprovalOk returns a tuple with the SkipApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigiCertConnectorResponse) GetSkipApprovalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipApproval.Get(), o.SkipApproval.IsSet()
}

// HasSkipApproval returns a boolean if a field has been set.
func (o *DigiCertConnectorResponse) HasSkipApproval() bool {
	if o != nil && o.SkipApproval.IsSet() {
		return true
	}

	return false
}

// SetSkipApproval gets a reference to the given NullableBool and assigns it to the SkipApproval field.
func (o *DigiCertConnectorResponse) SetSkipApproval(v bool) {
	o.SkipApproval.Set(&v)
}

// SetSkipApprovalNil sets the value for SkipApproval to be an explicit nil
func (o *DigiCertConnectorResponse) SetSkipApprovalNil() {
	o.SkipApproval.Set(nil)
}

// UnsetSkipApproval ensures that no value is present for SkipApproval, not even an explicit nil
func (o *DigiCertConnectorResponse) UnsetSkipApproval() {
	o.SkipApproval.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigiCertConnectorResponse) GetStatus() PKIConnectorStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret PKIConnectorStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigiCertConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *DigiCertConnectorResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePKIConnectorStatus and assigns it to the Status field.
func (o *DigiCertConnectorResponse) SetStatus(v PKIConnectorStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *DigiCertConnectorResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *DigiCertConnectorResponse) UnsetStatus() {
	o.Status.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DigiCertConnectorResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DigiCertConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *DigiCertConnectorResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *DigiCertConnectorResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *DigiCertConnectorResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *DigiCertConnectorResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *DigiCertConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DigiCertConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DigiCertConnectorResponse) SetType(v string) {
	o.Type = v
}

func (o DigiCertConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigiCertConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["apiCredentials"] = o.ApiCredentials
	toSerialize["baseUrl"] = o.BaseUrl
	if o.CaCertId.IsSet() {
		toSerialize["caCertId"] = o.CaCertId.Get()
	}
	if o.CustomConnectorDataMapping != nil {
		toSerialize["customConnectorDataMapping"] = o.CustomConnectorDataMapping
	}
	toSerialize["name"] = o.Name
	toSerialize["organizationId"] = o.OrganizationId
	if !utils.IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.RetryInterval.IsSet() {
		toSerialize["retryInterval"] = o.RetryInterval.Get()
	}
	if o.SkipApproval.IsSet() {
		toSerialize["skipApproval"] = o.SkipApproval.Get()
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

func (o *DigiCertConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"apiCredentials",
		"baseUrl",
		"name",
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

	varDigiCertConnectorResponse := _DigiCertConnectorResponse{}

	err = json.Unmarshal(data, &varDigiCertConnectorResponse)

	if err != nil {
		return err
	}

	*o = DigiCertConnectorResponse(varDigiCertConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "apiCredentials")
		delete(additionalProperties, "baseUrl")
		delete(additionalProperties, "caCertId")
		delete(additionalProperties, "customConnectorDataMapping")
		delete(additionalProperties, "name")
		delete(additionalProperties, "organizationId")
		delete(additionalProperties, "productId")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "retryInterval")
		delete(additionalProperties, "skipApproval")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDigiCertConnectorResponse struct {
	value *DigiCertConnectorResponse
	isSet bool
}

func (v NullableDigiCertConnectorResponse) Get() *DigiCertConnectorResponse {
	return v.value
}

func (v *NullableDigiCertConnectorResponse) Set(val *DigiCertConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDigiCertConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDigiCertConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigiCertConnectorResponse(val *DigiCertConnectorResponse) *NullableDigiCertConnectorResponse {
	return &NullableDigiCertConnectorResponse{value: val, isSet: true}
}

func (v NullableDigiCertConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigiCertConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
