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

// checks if the IntunePKCSConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IntunePKCSConnectorResponse{}

// IntunePKCSConnectorResponse struct for IntunePKCSConnectorResponse
type IntunePKCSConnectorResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Name of the `password` [credentials](#tag/security.credentials) containing the App ID and Key to authenticate on Intune PKCS
	Credentials                   string               `json:"credentials"`
	IntendedPurpose               utils.NullableString `json:"intendedPurpose,omitempty"`
	KeyName                       string               `json:"keyName"`
	MaxStoredCertificatePerHolder utils.NullableInt64  `json:"maxStoredCertificatePerHolder,omitempty"`
	Name                          string               `json:"name"`
	ProviderName                  utils.NullableString `json:"providerName,omitempty"`
	Proxy                         utils.NullableString `json:"proxy,omitempty"`
	PubKey                        string               `json:"pubKey"`
	RenewalPeriod                 utils.NullableString `json:"renewalPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	SearchFilter                  utils.NullableString `json:"searchFilter,omitempty"`
	Tenant                        string               `json:"tenant"`
	ThrottleDuration              string               `json:"throttleDuration" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	ThrottleParallelism           int64                `json:"throttleParallelism"`
	Timeout                       utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                          string               `json:"type"`
	AdditionalProperties          map[string]interface{}
}

type _IntunePKCSConnectorResponse IntunePKCSConnectorResponse

// NewIntunePKCSConnectorResponse instantiates a new IntunePKCSConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntunePKCSConnectorResponse(id string, credentials string, keyName string, name string, pubKey string, tenant string, throttleDuration string, throttleParallelism int64, type_ string) *IntunePKCSConnectorResponse {
	this := IntunePKCSConnectorResponse{}
	this.Id = id
	this.Credentials = credentials
	this.KeyName = keyName
	this.Name = name
	this.PubKey = pubKey
	this.Tenant = tenant
	this.ThrottleDuration = throttleDuration
	this.ThrottleParallelism = throttleParallelism
	this.Type = type_
	return &this
}

// NewIntunePKCSConnectorResponseWithDefaults instantiates a new IntunePKCSConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntunePKCSConnectorResponseWithDefaults() *IntunePKCSConnectorResponse {
	this := IntunePKCSConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *IntunePKCSConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IntunePKCSConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IntunePKCSConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetCredentials returns the Credentials field value
func (o *IntunePKCSConnectorResponse) GetCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *IntunePKCSConnectorResponse) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *IntunePKCSConnectorResponse) SetCredentials(v string) {
	o.Credentials = v
}

// GetIntendedPurpose returns the IntendedPurpose field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntunePKCSConnectorResponse) GetIntendedPurpose() string {
	if o == nil || utils.IsNil(o.IntendedPurpose.Get()) {
		var ret string
		return ret
	}
	return *o.IntendedPurpose.Get()
}

// GetIntendedPurposeOk returns a tuple with the IntendedPurpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntunePKCSConnectorResponse) GetIntendedPurposeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntendedPurpose.Get(), o.IntendedPurpose.IsSet()
}

// HasIntendedPurpose returns a boolean if a field has been set.
func (o *IntunePKCSConnectorResponse) HasIntendedPurpose() bool {
	if o != nil && o.IntendedPurpose.IsSet() {
		return true
	}

	return false
}

// SetIntendedPurpose gets a reference to the given NullableString and assigns it to the IntendedPurpose field.
func (o *IntunePKCSConnectorResponse) SetIntendedPurpose(v string) {
	o.IntendedPurpose.Set(&v)
}

// SetIntendedPurposeNil sets the value for IntendedPurpose to be an explicit nil
func (o *IntunePKCSConnectorResponse) SetIntendedPurposeNil() {
	o.IntendedPurpose.Set(nil)
}

// UnsetIntendedPurpose ensures that no value is present for IntendedPurpose, not even an explicit nil
func (o *IntunePKCSConnectorResponse) UnsetIntendedPurpose() {
	o.IntendedPurpose.Unset()
}

// GetKeyName returns the KeyName field value
func (o *IntunePKCSConnectorResponse) GetKeyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyName
}

// GetKeyNameOk returns a tuple with the KeyName field value
// and a boolean to check if the value has been set.
func (o *IntunePKCSConnectorResponse) GetKeyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyName, true
}

// SetKeyName sets field value
func (o *IntunePKCSConnectorResponse) SetKeyName(v string) {
	o.KeyName = v
}

// GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntunePKCSConnectorResponse) GetMaxStoredCertificatePerHolder() int64 {
	if o == nil || utils.IsNil(o.MaxStoredCertificatePerHolder.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxStoredCertificatePerHolder.Get()
}

// GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntunePKCSConnectorResponse) GetMaxStoredCertificatePerHolderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStoredCertificatePerHolder.Get(), o.MaxStoredCertificatePerHolder.IsSet()
}

// HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.
func (o *IntunePKCSConnectorResponse) HasMaxStoredCertificatePerHolder() bool {
	if o != nil && o.MaxStoredCertificatePerHolder.IsSet() {
		return true
	}

	return false
}

// SetMaxStoredCertificatePerHolder gets a reference to the given NullableInt64 and assigns it to the MaxStoredCertificatePerHolder field.
func (o *IntunePKCSConnectorResponse) SetMaxStoredCertificatePerHolder(v int64) {
	o.MaxStoredCertificatePerHolder.Set(&v)
}

// SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil
func (o *IntunePKCSConnectorResponse) SetMaxStoredCertificatePerHolderNil() {
	o.MaxStoredCertificatePerHolder.Set(nil)
}

// UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
func (o *IntunePKCSConnectorResponse) UnsetMaxStoredCertificatePerHolder() {
	o.MaxStoredCertificatePerHolder.Unset()
}

// GetName returns the Name field value
func (o *IntunePKCSConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntunePKCSConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntunePKCSConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntunePKCSConnectorResponse) GetProviderName() string {
	if o == nil || utils.IsNil(o.ProviderName.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderName.Get()
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntunePKCSConnectorResponse) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderName.Get(), o.ProviderName.IsSet()
}

// HasProviderName returns a boolean if a field has been set.
func (o *IntunePKCSConnectorResponse) HasProviderName() bool {
	if o != nil && o.ProviderName.IsSet() {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given NullableString and assigns it to the ProviderName field.
func (o *IntunePKCSConnectorResponse) SetProviderName(v string) {
	o.ProviderName.Set(&v)
}

// SetProviderNameNil sets the value for ProviderName to be an explicit nil
func (o *IntunePKCSConnectorResponse) SetProviderNameNil() {
	o.ProviderName.Set(nil)
}

// UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
func (o *IntunePKCSConnectorResponse) UnsetProviderName() {
	o.ProviderName.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntunePKCSConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntunePKCSConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *IntunePKCSConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *IntunePKCSConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *IntunePKCSConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *IntunePKCSConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetPubKey returns the PubKey field value
func (o *IntunePKCSConnectorResponse) GetPubKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PubKey
}

// GetPubKeyOk returns a tuple with the PubKey field value
// and a boolean to check if the value has been set.
func (o *IntunePKCSConnectorResponse) GetPubKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PubKey, true
}

// SetPubKey sets field value
func (o *IntunePKCSConnectorResponse) SetPubKey(v string) {
	o.PubKey = v
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntunePKCSConnectorResponse) GetRenewalPeriod() string {
	if o == nil || utils.IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntunePKCSConnectorResponse) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *IntunePKCSConnectorResponse) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *IntunePKCSConnectorResponse) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}

// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *IntunePKCSConnectorResponse) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *IntunePKCSConnectorResponse) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetSearchFilter returns the SearchFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntunePKCSConnectorResponse) GetSearchFilter() string {
	if o == nil || utils.IsNil(o.SearchFilter.Get()) {
		var ret string
		return ret
	}
	return *o.SearchFilter.Get()
}

// GetSearchFilterOk returns a tuple with the SearchFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntunePKCSConnectorResponse) GetSearchFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchFilter.Get(), o.SearchFilter.IsSet()
}

// HasSearchFilter returns a boolean if a field has been set.
func (o *IntunePKCSConnectorResponse) HasSearchFilter() bool {
	if o != nil && o.SearchFilter.IsSet() {
		return true
	}

	return false
}

// SetSearchFilter gets a reference to the given NullableString and assigns it to the SearchFilter field.
func (o *IntunePKCSConnectorResponse) SetSearchFilter(v string) {
	o.SearchFilter.Set(&v)
}

// SetSearchFilterNil sets the value for SearchFilter to be an explicit nil
func (o *IntunePKCSConnectorResponse) SetSearchFilterNil() {
	o.SearchFilter.Set(nil)
}

// UnsetSearchFilter ensures that no value is present for SearchFilter, not even an explicit nil
func (o *IntunePKCSConnectorResponse) UnsetSearchFilter() {
	o.SearchFilter.Unset()
}

// GetTenant returns the Tenant field value
func (o *IntunePKCSConnectorResponse) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *IntunePKCSConnectorResponse) GetTenantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *IntunePKCSConnectorResponse) SetTenant(v string) {
	o.Tenant = v
}

// GetThrottleDuration returns the ThrottleDuration field value
func (o *IntunePKCSConnectorResponse) GetThrottleDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThrottleDuration
}

// GetThrottleDurationOk returns a tuple with the ThrottleDuration field value
// and a boolean to check if the value has been set.
func (o *IntunePKCSConnectorResponse) GetThrottleDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleDuration, true
}

// SetThrottleDuration sets field value
func (o *IntunePKCSConnectorResponse) SetThrottleDuration(v string) {
	o.ThrottleDuration = v
}

// GetThrottleParallelism returns the ThrottleParallelism field value
func (o *IntunePKCSConnectorResponse) GetThrottleParallelism() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ThrottleParallelism
}

// GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field value
// and a boolean to check if the value has been set.
func (o *IntunePKCSConnectorResponse) GetThrottleParallelismOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleParallelism, true
}

// SetThrottleParallelism sets field value
func (o *IntunePKCSConnectorResponse) SetThrottleParallelism(v int64) {
	o.ThrottleParallelism = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntunePKCSConnectorResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntunePKCSConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *IntunePKCSConnectorResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *IntunePKCSConnectorResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *IntunePKCSConnectorResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *IntunePKCSConnectorResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *IntunePKCSConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntunePKCSConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IntunePKCSConnectorResponse) SetType(v string) {
	o.Type = v
}

func (o IntunePKCSConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntunePKCSConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["credentials"] = o.Credentials
	if o.IntendedPurpose.IsSet() {
		toSerialize["intendedPurpose"] = o.IntendedPurpose.Get()
	}
	toSerialize["keyName"] = o.KeyName
	if o.MaxStoredCertificatePerHolder.IsSet() {
		toSerialize["maxStoredCertificatePerHolder"] = o.MaxStoredCertificatePerHolder.Get()
	}
	toSerialize["name"] = o.Name
	if o.ProviderName.IsSet() {
		toSerialize["providerName"] = o.ProviderName.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	toSerialize["pubKey"] = o.PubKey
	if o.RenewalPeriod.IsSet() {
		toSerialize["renewalPeriod"] = o.RenewalPeriod.Get()
	}
	if o.SearchFilter.IsSet() {
		toSerialize["searchFilter"] = o.SearchFilter.Get()
	}
	toSerialize["tenant"] = o.Tenant
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

func (o *IntunePKCSConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"credentials",
		"keyName",
		"name",
		"pubKey",
		"tenant",
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

	varIntunePKCSConnectorResponse := _IntunePKCSConnectorResponse{}

	err = json.Unmarshal(data, &varIntunePKCSConnectorResponse)

	if err != nil {
		return err
	}

	*o = IntunePKCSConnectorResponse(varIntunePKCSConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "intendedPurpose")
		delete(additionalProperties, "keyName")
		delete(additionalProperties, "maxStoredCertificatePerHolder")
		delete(additionalProperties, "name")
		delete(additionalProperties, "providerName")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "pubKey")
		delete(additionalProperties, "renewalPeriod")
		delete(additionalProperties, "searchFilter")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "throttleDuration")
		delete(additionalProperties, "throttleParallelism")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntunePKCSConnectorResponse struct {
	value *IntunePKCSConnectorResponse
	isSet bool
}

func (v NullableIntunePKCSConnectorResponse) Get() *IntunePKCSConnectorResponse {
	return v.value
}

func (v *NullableIntunePKCSConnectorResponse) Set(val *IntunePKCSConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIntunePKCSConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIntunePKCSConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntunePKCSConnectorResponse(val *IntunePKCSConnectorResponse) *NullableIntunePKCSConnectorResponse {
	return &NullableIntunePKCSConnectorResponse{value: val, isSet: true}
}

func (v NullableIntunePKCSConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntunePKCSConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
