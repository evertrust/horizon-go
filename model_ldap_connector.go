/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"encoding/json"
	"fmt"
)

// checks if the LDAPConnector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LDAPConnector{}

// LDAPConnector struct for LDAPConnector
type LDAPConnector struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Hostname string `json:"hostname"`
	Port NullableInt64 `json:"port,omitempty"`
	BaseDn string `json:"baseDn"`
	Filter NullableString `json:"filter,omitempty"`
	CertAttr NullableString `json:"certAttr,omitempty"`
	FollowReferrals NullableBool `json:"followReferrals,omitempty"`
	UserIdentifierAttribute string `json:"userIdentifierAttribute"`
	CertificateAttribute string `json:"certificateAttribute"`
	ThrottleDuration string `json:"throttleDuration" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	ThrottleParallelism *int64 `json:"throttleParallelism,omitempty"`
	Timeout NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Proxy NullableString `json:"proxy,omitempty"`
	// Name of the `password` [credentials](#tag/api.security.credentials) containing login DN and password.
	Credentials string `json:"credentials"`
	MaxStoredCertificatePerHolder NullableInt64 `json:"maxStoredCertificatePerHolder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LDAPConnector LDAPConnector

// NewLDAPConnector instantiates a new LDAPConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPConnector(type_ string, name string, hostname string, baseDn string, userIdentifierAttribute string, certificateAttribute string, throttleDuration string, credentials string) *LDAPConnector {
	this := LDAPConnector{}
	this.Type = type_
	this.Name = name
	this.Hostname = hostname
	this.BaseDn = baseDn
	this.UserIdentifierAttribute = userIdentifierAttribute
	this.CertificateAttribute = certificateAttribute
	this.ThrottleDuration = throttleDuration
	this.Credentials = credentials
	return &this
}

// NewLDAPConnectorWithDefaults instantiates a new LDAPConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPConnectorWithDefaults() *LDAPConnector {
	this := LDAPConnector{}
	return &this
}

// GetType returns the Type field value
func (o *LDAPConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LDAPConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LDAPConnector) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *LDAPConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LDAPConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LDAPConnector) SetName(v string) {
	o.Name = v
}

// GetHostname returns the Hostname field value
func (o *LDAPConnector) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *LDAPConnector) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *LDAPConnector) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPConnector) GetPort() int64 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int64
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPConnector) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *LDAPConnector) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt64 and assigns it to the Port field.
func (o *LDAPConnector) SetPort(v int64) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *LDAPConnector) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *LDAPConnector) UnsetPort() {
	o.Port.Unset()
}

// GetBaseDn returns the BaseDn field value
func (o *LDAPConnector) GetBaseDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value
// and a boolean to check if the value has been set.
func (o *LDAPConnector) GetBaseDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseDn, true
}

// SetBaseDn sets field value
func (o *LDAPConnector) SetBaseDn(v string) {
	o.BaseDn = v
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPConnector) GetFilter() string {
	if o == nil || IsNil(o.Filter.Get()) {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPConnector) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *LDAPConnector) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableString and assigns it to the Filter field.
func (o *LDAPConnector) SetFilter(v string) {
	o.Filter.Set(&v)
}
// SetFilterNil sets the value for Filter to be an explicit nil
func (o *LDAPConnector) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *LDAPConnector) UnsetFilter() {
	o.Filter.Unset()
}

// GetCertAttr returns the CertAttr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPConnector) GetCertAttr() string {
	if o == nil || IsNil(o.CertAttr.Get()) {
		var ret string
		return ret
	}
	return *o.CertAttr.Get()
}

// GetCertAttrOk returns a tuple with the CertAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPConnector) GetCertAttrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertAttr.Get(), o.CertAttr.IsSet()
}

// HasCertAttr returns a boolean if a field has been set.
func (o *LDAPConnector) HasCertAttr() bool {
	if o != nil && o.CertAttr.IsSet() {
		return true
	}

	return false
}

// SetCertAttr gets a reference to the given NullableString and assigns it to the CertAttr field.
func (o *LDAPConnector) SetCertAttr(v string) {
	o.CertAttr.Set(&v)
}
// SetCertAttrNil sets the value for CertAttr to be an explicit nil
func (o *LDAPConnector) SetCertAttrNil() {
	o.CertAttr.Set(nil)
}

// UnsetCertAttr ensures that no value is present for CertAttr, not even an explicit nil
func (o *LDAPConnector) UnsetCertAttr() {
	o.CertAttr.Unset()
}

// GetFollowReferrals returns the FollowReferrals field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPConnector) GetFollowReferrals() bool {
	if o == nil || IsNil(o.FollowReferrals.Get()) {
		var ret bool
		return ret
	}
	return *o.FollowReferrals.Get()
}

// GetFollowReferralsOk returns a tuple with the FollowReferrals field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPConnector) GetFollowReferralsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowReferrals.Get(), o.FollowReferrals.IsSet()
}

// HasFollowReferrals returns a boolean if a field has been set.
func (o *LDAPConnector) HasFollowReferrals() bool {
	if o != nil && o.FollowReferrals.IsSet() {
		return true
	}

	return false
}

// SetFollowReferrals gets a reference to the given NullableBool and assigns it to the FollowReferrals field.
func (o *LDAPConnector) SetFollowReferrals(v bool) {
	o.FollowReferrals.Set(&v)
}
// SetFollowReferralsNil sets the value for FollowReferrals to be an explicit nil
func (o *LDAPConnector) SetFollowReferralsNil() {
	o.FollowReferrals.Set(nil)
}

// UnsetFollowReferrals ensures that no value is present for FollowReferrals, not even an explicit nil
func (o *LDAPConnector) UnsetFollowReferrals() {
	o.FollowReferrals.Unset()
}

// GetUserIdentifierAttribute returns the UserIdentifierAttribute field value
func (o *LDAPConnector) GetUserIdentifierAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserIdentifierAttribute
}

// GetUserIdentifierAttributeOk returns a tuple with the UserIdentifierAttribute field value
// and a boolean to check if the value has been set.
func (o *LDAPConnector) GetUserIdentifierAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserIdentifierAttribute, true
}

// SetUserIdentifierAttribute sets field value
func (o *LDAPConnector) SetUserIdentifierAttribute(v string) {
	o.UserIdentifierAttribute = v
}

// GetCertificateAttribute returns the CertificateAttribute field value
func (o *LDAPConnector) GetCertificateAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateAttribute
}

// GetCertificateAttributeOk returns a tuple with the CertificateAttribute field value
// and a boolean to check if the value has been set.
func (o *LDAPConnector) GetCertificateAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateAttribute, true
}

// SetCertificateAttribute sets field value
func (o *LDAPConnector) SetCertificateAttribute(v string) {
	o.CertificateAttribute = v
}

// GetThrottleDuration returns the ThrottleDuration field value
func (o *LDAPConnector) GetThrottleDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThrottleDuration
}

// GetThrottleDurationOk returns a tuple with the ThrottleDuration field value
// and a boolean to check if the value has been set.
func (o *LDAPConnector) GetThrottleDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleDuration, true
}

// SetThrottleDuration sets field value
func (o *LDAPConnector) SetThrottleDuration(v string) {
	o.ThrottleDuration = v
}

// GetThrottleParallelism returns the ThrottleParallelism field value if set, zero value otherwise.
func (o *LDAPConnector) GetThrottleParallelism() int64 {
	if o == nil || IsNil(o.ThrottleParallelism) {
		var ret int64
		return ret
	}
	return *o.ThrottleParallelism
}

// GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPConnector) GetThrottleParallelismOk() (*int64, bool) {
	if o == nil || IsNil(o.ThrottleParallelism) {
		return nil, false
	}
	return o.ThrottleParallelism, true
}

// HasThrottleParallelism returns a boolean if a field has been set.
func (o *LDAPConnector) HasThrottleParallelism() bool {
	if o != nil && !IsNil(o.ThrottleParallelism) {
		return true
	}

	return false
}

// SetThrottleParallelism gets a reference to the given int64 and assigns it to the ThrottleParallelism field.
func (o *LDAPConnector) SetThrottleParallelism(v int64) {
	o.ThrottleParallelism = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPConnector) GetTimeout() string {
	if o == nil || IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *LDAPConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *LDAPConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}
// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *LDAPConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *LDAPConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPConnector) GetProxy() string {
	if o == nil || IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *LDAPConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *LDAPConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}
// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *LDAPConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *LDAPConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetCredentials returns the Credentials field value
func (o *LDAPConnector) GetCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *LDAPConnector) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *LDAPConnector) SetCredentials(v string) {
	o.Credentials = v
}

// GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPConnector) GetMaxStoredCertificatePerHolder() int64 {
	if o == nil || IsNil(o.MaxStoredCertificatePerHolder.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxStoredCertificatePerHolder.Get()
}

// GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPConnector) GetMaxStoredCertificatePerHolderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStoredCertificatePerHolder.Get(), o.MaxStoredCertificatePerHolder.IsSet()
}

// HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.
func (o *LDAPConnector) HasMaxStoredCertificatePerHolder() bool {
	if o != nil && o.MaxStoredCertificatePerHolder.IsSet() {
		return true
	}

	return false
}

// SetMaxStoredCertificatePerHolder gets a reference to the given NullableInt64 and assigns it to the MaxStoredCertificatePerHolder field.
func (o *LDAPConnector) SetMaxStoredCertificatePerHolder(v int64) {
	o.MaxStoredCertificatePerHolder.Set(&v)
}
// SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil
func (o *LDAPConnector) SetMaxStoredCertificatePerHolderNil() {
	o.MaxStoredCertificatePerHolder.Set(nil)
}

// UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
func (o *LDAPConnector) UnsetMaxStoredCertificatePerHolder() {
	o.MaxStoredCertificatePerHolder.Unset()
}

func (o LDAPConnector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LDAPConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["hostname"] = o.Hostname
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	toSerialize["baseDn"] = o.BaseDn
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	if o.CertAttr.IsSet() {
		toSerialize["certAttr"] = o.CertAttr.Get()
	}
	if o.FollowReferrals.IsSet() {
		toSerialize["followReferrals"] = o.FollowReferrals.Get()
	}
	toSerialize["userIdentifierAttribute"] = o.UserIdentifierAttribute
	toSerialize["certificateAttribute"] = o.CertificateAttribute
	toSerialize["throttleDuration"] = o.ThrottleDuration
	if !IsNil(o.ThrottleParallelism) {
		toSerialize["throttleParallelism"] = o.ThrottleParallelism
	}
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	toSerialize["credentials"] = o.Credentials
	if o.MaxStoredCertificatePerHolder.IsSet() {
		toSerialize["maxStoredCertificatePerHolder"] = o.MaxStoredCertificatePerHolder.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LDAPConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"hostname",
		"baseDn",
		"userIdentifierAttribute",
		"certificateAttribute",
		"throttleDuration",
		"credentials",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLDAPConnector := _LDAPConnector{}

	err = json.Unmarshal(data, &varLDAPConnector)

	if err != nil {
		return err
	}

	*o = LDAPConnector(varLDAPConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "port")
		delete(additionalProperties, "baseDn")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "certAttr")
		delete(additionalProperties, "followReferrals")
		delete(additionalProperties, "userIdentifierAttribute")
		delete(additionalProperties, "certificateAttribute")
		delete(additionalProperties, "throttleDuration")
		delete(additionalProperties, "throttleParallelism")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "maxStoredCertificatePerHolder")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLDAPConnector struct {
	value *LDAPConnector
	isSet bool
}

func (v NullableLDAPConnector) Get() *LDAPConnector {
	return v.value
}

func (v *NullableLDAPConnector) Set(val *LDAPConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPConnector(val *LDAPConnector) *NullableLDAPConnector {
	return &NullableLDAPConnector{value: val, isSet: true}
}

func (v NullableLDAPConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


