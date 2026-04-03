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

// checks if the NetscalerConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NetscalerConnector{}

// NetscalerConnector struct for NetscalerConnector
type NetscalerConnector struct {
	CertificateStorePath string `json:"certificateStorePath"`
	// Name of the `password` [credentials](#tag/security.credentials) containing the account to authenticate on Netscaler
	Credentials                   string               `json:"credentials"`
	Hostname                      string               `json:"hostname"`
	MaxStoredCertificatePerHolder utils.NullableInt64  `json:"maxStoredCertificatePerHolder,omitempty"`
	Name                          string               `json:"name"`
	Prefix                        utils.NullableString `json:"prefix,omitempty"`
	Proxy                         utils.NullableString `json:"proxy,omitempty"`
	RenewalPeriod                 utils.NullableString `json:"renewalPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	ThrottleDuration              string               `json:"throttleDuration" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	ThrottleParallelism           int64                `json:"throttleParallelism"`
	Timeout                       utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Allow invalid server certificates when establishing the TLS connection. Use in production is *not* recommended.
	TlsInsecure          utils.NullableBool `json:"tlsInsecure,omitempty"`
	Type                 string             `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _NetscalerConnector NetscalerConnector

// NewNetscalerConnector instantiates a new NetscalerConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetscalerConnector(certificateStorePath string, credentials string, hostname string, name string, throttleDuration string, throttleParallelism int64, type_ string) *NetscalerConnector {
	this := NetscalerConnector{}
	this.CertificateStorePath = certificateStorePath
	this.Credentials = credentials
	this.Hostname = hostname
	this.Name = name
	this.ThrottleDuration = throttleDuration
	this.ThrottleParallelism = throttleParallelism
	var tlsInsecure bool = false
	this.TlsInsecure = *utils.NewNullableBool(&tlsInsecure)
	this.Type = type_
	return &this
}

// NewNetscalerConnectorWithDefaults instantiates a new NetscalerConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetscalerConnectorWithDefaults() *NetscalerConnector {
	this := NetscalerConnector{}
	var tlsInsecure bool = false
	this.TlsInsecure = *utils.NewNullableBool(&tlsInsecure)
	return &this
}

// GetCertificateStorePath returns the CertificateStorePath field value
func (o *NetscalerConnector) GetCertificateStorePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateStorePath
}

// GetCertificateStorePathOk returns a tuple with the CertificateStorePath field value
// and a boolean to check if the value has been set.
func (o *NetscalerConnector) GetCertificateStorePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateStorePath, true
}

// SetCertificateStorePath sets field value
func (o *NetscalerConnector) SetCertificateStorePath(v string) {
	o.CertificateStorePath = v
}

// GetCredentials returns the Credentials field value
func (o *NetscalerConnector) GetCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *NetscalerConnector) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *NetscalerConnector) SetCredentials(v string) {
	o.Credentials = v
}

// GetHostname returns the Hostname field value
func (o *NetscalerConnector) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *NetscalerConnector) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *NetscalerConnector) SetHostname(v string) {
	o.Hostname = v
}

// GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetscalerConnector) GetMaxStoredCertificatePerHolder() int64 {
	if o == nil || utils.IsNil(o.MaxStoredCertificatePerHolder.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxStoredCertificatePerHolder.Get()
}

// GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetscalerConnector) GetMaxStoredCertificatePerHolderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStoredCertificatePerHolder.Get(), o.MaxStoredCertificatePerHolder.IsSet()
}

// HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.
func (o *NetscalerConnector) HasMaxStoredCertificatePerHolder() bool {
	if o != nil && o.MaxStoredCertificatePerHolder.IsSet() {
		return true
	}

	return false
}

// SetMaxStoredCertificatePerHolder gets a reference to the given NullableInt64 and assigns it to the MaxStoredCertificatePerHolder field.
func (o *NetscalerConnector) SetMaxStoredCertificatePerHolder(v int64) {
	o.MaxStoredCertificatePerHolder.Set(&v)
}

// SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil
func (o *NetscalerConnector) SetMaxStoredCertificatePerHolderNil() {
	o.MaxStoredCertificatePerHolder.Set(nil)
}

// UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil
func (o *NetscalerConnector) UnsetMaxStoredCertificatePerHolder() {
	o.MaxStoredCertificatePerHolder.Unset()
}

// GetName returns the Name field value
func (o *NetscalerConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetscalerConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetscalerConnector) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetscalerConnector) GetPrefix() string {
	if o == nil || utils.IsNil(o.Prefix.Get()) {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetscalerConnector) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *NetscalerConnector) HasPrefix() bool {
	if o != nil && o.Prefix.IsSet() {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given NullableString and assigns it to the Prefix field.
func (o *NetscalerConnector) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// SetPrefixNil sets the value for Prefix to be an explicit nil
func (o *NetscalerConnector) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
func (o *NetscalerConnector) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetscalerConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetscalerConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *NetscalerConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *NetscalerConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *NetscalerConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *NetscalerConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetscalerConnector) GetRenewalPeriod() string {
	if o == nil || utils.IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetscalerConnector) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *NetscalerConnector) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *NetscalerConnector) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}

// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *NetscalerConnector) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *NetscalerConnector) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetThrottleDuration returns the ThrottleDuration field value
func (o *NetscalerConnector) GetThrottleDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThrottleDuration
}

// GetThrottleDurationOk returns a tuple with the ThrottleDuration field value
// and a boolean to check if the value has been set.
func (o *NetscalerConnector) GetThrottleDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleDuration, true
}

// SetThrottleDuration sets field value
func (o *NetscalerConnector) SetThrottleDuration(v string) {
	o.ThrottleDuration = v
}

// GetThrottleParallelism returns the ThrottleParallelism field value
func (o *NetscalerConnector) GetThrottleParallelism() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ThrottleParallelism
}

// GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field value
// and a boolean to check if the value has been set.
func (o *NetscalerConnector) GetThrottleParallelismOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleParallelism, true
}

// SetThrottleParallelism sets field value
func (o *NetscalerConnector) SetThrottleParallelism(v int64) {
	o.ThrottleParallelism = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetscalerConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetscalerConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *NetscalerConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *NetscalerConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *NetscalerConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *NetscalerConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetTlsInsecure returns the TlsInsecure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetscalerConnector) GetTlsInsecure() bool {
	if o == nil || utils.IsNil(o.TlsInsecure.Get()) {
		var ret bool
		return ret
	}
	return *o.TlsInsecure.Get()
}

// GetTlsInsecureOk returns a tuple with the TlsInsecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetscalerConnector) GetTlsInsecureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsInsecure.Get(), o.TlsInsecure.IsSet()
}

// HasTlsInsecure returns a boolean if a field has been set.
func (o *NetscalerConnector) HasTlsInsecure() bool {
	if o != nil && o.TlsInsecure.IsSet() {
		return true
	}

	return false
}

// SetTlsInsecure gets a reference to the given NullableBool and assigns it to the TlsInsecure field.
func (o *NetscalerConnector) SetTlsInsecure(v bool) {
	o.TlsInsecure.Set(&v)
}

// SetTlsInsecureNil sets the value for TlsInsecure to be an explicit nil
func (o *NetscalerConnector) SetTlsInsecureNil() {
	o.TlsInsecure.Set(nil)
}

// UnsetTlsInsecure ensures that no value is present for TlsInsecure, not even an explicit nil
func (o *NetscalerConnector) UnsetTlsInsecure() {
	o.TlsInsecure.Unset()
}

// GetType returns the Type field value
func (o *NetscalerConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetscalerConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetscalerConnector) SetType(v string) {
	o.Type = v
}

func (o NetscalerConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetscalerConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["certificateStorePath"] = o.CertificateStorePath
	toSerialize["credentials"] = o.Credentials
	toSerialize["hostname"] = o.Hostname
	if o.MaxStoredCertificatePerHolder.IsSet() {
		toSerialize["maxStoredCertificatePerHolder"] = o.MaxStoredCertificatePerHolder.Get()
	}
	toSerialize["name"] = o.Name
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.RenewalPeriod.IsSet() {
		toSerialize["renewalPeriod"] = o.RenewalPeriod.Get()
	}
	toSerialize["throttleDuration"] = o.ThrottleDuration
	toSerialize["throttleParallelism"] = o.ThrottleParallelism
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	if o.TlsInsecure.IsSet() {
		toSerialize["tlsInsecure"] = o.TlsInsecure.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetscalerConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"certificateStorePath",
		"credentials",
		"hostname",
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

	varNetscalerConnector := _NetscalerConnector{}

	err = json.Unmarshal(data, &varNetscalerConnector)

	if err != nil {
		return err
	}

	*o = NetscalerConnector(varNetscalerConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificateStorePath")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "maxStoredCertificatePerHolder")
		delete(additionalProperties, "name")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "renewalPeriod")
		delete(additionalProperties, "throttleDuration")
		delete(additionalProperties, "throttleParallelism")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "tlsInsecure")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetscalerConnector struct {
	value *NetscalerConnector
	isSet bool
}

func (v NullableNetscalerConnector) Get() *NetscalerConnector {
	return v.value
}

func (v *NullableNetscalerConnector) Set(val *NetscalerConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableNetscalerConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableNetscalerConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetscalerConnector(val *NetscalerConnector) *NullableNetscalerConnector {
	return &NullableNetscalerConnector{value: val, isSet: true}
}

func (v NullableNetscalerConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetscalerConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
