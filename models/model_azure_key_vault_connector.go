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

// checks if the AzureKeyVaultConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AzureKeyVaultConnector{}

// AzureKeyVaultConnector struct for AzureKeyVaultConnector
type AzureKeyVaultConnector struct {
	AzureTenant *string `json:"azureTenant,omitempty"`
	// Name of the `password` [credentials](#tag/security.credentials) containing the App ID and Key to authenticate on AKV
	Credentials          string               `json:"credentials"`
	Name                 string               `json:"name"`
	Prefix               utils.NullableString `json:"prefix,omitempty"`
	Proxy                utils.NullableString `json:"proxy,omitempty"`
	RenewalPeriod        utils.NullableString `json:"renewalPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	ThrottleDuration     string               `json:"throttleDuration" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	ThrottleParallelism  int64                `json:"throttleParallelism"`
	Timeout              utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string               `json:"type"`
	VaultBaseUrl         string               `json:"vaultBaseUrl"`
	AdditionalProperties map[string]interface{}
}

type _AzureKeyVaultConnector AzureKeyVaultConnector

// NewAzureKeyVaultConnector instantiates a new AzureKeyVaultConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureKeyVaultConnector(credentials string, name string, throttleDuration string, throttleParallelism int64, type_ string, vaultBaseUrl string) *AzureKeyVaultConnector {
	this := AzureKeyVaultConnector{}
	this.Credentials = credentials
	this.Name = name
	this.ThrottleDuration = throttleDuration
	this.ThrottleParallelism = throttleParallelism
	this.Type = type_
	this.VaultBaseUrl = vaultBaseUrl
	return &this
}

// NewAzureKeyVaultConnectorWithDefaults instantiates a new AzureKeyVaultConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureKeyVaultConnectorWithDefaults() *AzureKeyVaultConnector {
	this := AzureKeyVaultConnector{}
	return &this
}

// GetAzureTenant returns the AzureTenant field value if set, zero value otherwise.
func (o *AzureKeyVaultConnector) GetAzureTenant() string {
	if o == nil || utils.IsNil(o.AzureTenant) {
		var ret string
		return ret
	}
	return *o.AzureTenant
}

// GetAzureTenantOk returns a tuple with the AzureTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultConnector) GetAzureTenantOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AzureTenant) {
		return nil, false
	}
	return o.AzureTenant, true
}

// HasAzureTenant returns a boolean if a field has been set.
func (o *AzureKeyVaultConnector) HasAzureTenant() bool {
	if o != nil && !utils.IsNil(o.AzureTenant) {
		return true
	}

	return false
}

// SetAzureTenant gets a reference to the given string and assigns it to the AzureTenant field.
func (o *AzureKeyVaultConnector) SetAzureTenant(v string) {
	o.AzureTenant = &v
}

// GetCredentials returns the Credentials field value
func (o *AzureKeyVaultConnector) GetCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultConnector) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *AzureKeyVaultConnector) SetCredentials(v string) {
	o.Credentials = v
}

// GetName returns the Name field value
func (o *AzureKeyVaultConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AzureKeyVaultConnector) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultConnector) GetPrefix() string {
	if o == nil || utils.IsNil(o.Prefix.Get()) {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultConnector) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *AzureKeyVaultConnector) HasPrefix() bool {
	if o != nil && o.Prefix.IsSet() {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given NullableString and assigns it to the Prefix field.
func (o *AzureKeyVaultConnector) SetPrefix(v string) {
	o.Prefix.Set(&v)
}

// SetPrefixNil sets the value for Prefix to be an explicit nil
func (o *AzureKeyVaultConnector) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
func (o *AzureKeyVaultConnector) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *AzureKeyVaultConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *AzureKeyVaultConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *AzureKeyVaultConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *AzureKeyVaultConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultConnector) GetRenewalPeriod() string {
	if o == nil || utils.IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultConnector) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *AzureKeyVaultConnector) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *AzureKeyVaultConnector) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}

// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *AzureKeyVaultConnector) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *AzureKeyVaultConnector) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetThrottleDuration returns the ThrottleDuration field value
func (o *AzureKeyVaultConnector) GetThrottleDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThrottleDuration
}

// GetThrottleDurationOk returns a tuple with the ThrottleDuration field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultConnector) GetThrottleDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleDuration, true
}

// SetThrottleDuration sets field value
func (o *AzureKeyVaultConnector) SetThrottleDuration(v string) {
	o.ThrottleDuration = v
}

// GetThrottleParallelism returns the ThrottleParallelism field value
func (o *AzureKeyVaultConnector) GetThrottleParallelism() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ThrottleParallelism
}

// GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultConnector) GetThrottleParallelismOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleParallelism, true
}

// SetThrottleParallelism sets field value
func (o *AzureKeyVaultConnector) SetThrottleParallelism(v int64) {
	o.ThrottleParallelism = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzureKeyVaultConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzureKeyVaultConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *AzureKeyVaultConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *AzureKeyVaultConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *AzureKeyVaultConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *AzureKeyVaultConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *AzureKeyVaultConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AzureKeyVaultConnector) SetType(v string) {
	o.Type = v
}

// GetVaultBaseUrl returns the VaultBaseUrl field value
func (o *AzureKeyVaultConnector) GetVaultBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultBaseUrl
}

// GetVaultBaseUrlOk returns a tuple with the VaultBaseUrl field value
// and a boolean to check if the value has been set.
func (o *AzureKeyVaultConnector) GetVaultBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultBaseUrl, true
}

// SetVaultBaseUrl sets field value
func (o *AzureKeyVaultConnector) SetVaultBaseUrl(v string) {
	o.VaultBaseUrl = v
}

func (o AzureKeyVaultConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureKeyVaultConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AzureTenant) {
		toSerialize["azureTenant"] = o.AzureTenant
	}
	toSerialize["credentials"] = o.Credentials
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
	toSerialize["type"] = o.Type
	toSerialize["vaultBaseUrl"] = o.VaultBaseUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AzureKeyVaultConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"credentials",
		"name",
		"throttleDuration",
		"throttleParallelism",
		"type",
		"vaultBaseUrl",
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

	varAzureKeyVaultConnector := _AzureKeyVaultConnector{}

	err = json.Unmarshal(data, &varAzureKeyVaultConnector)

	if err != nil {
		return err
	}

	*o = AzureKeyVaultConnector(varAzureKeyVaultConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "azureTenant")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "name")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "renewalPeriod")
		delete(additionalProperties, "throttleDuration")
		delete(additionalProperties, "throttleParallelism")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		delete(additionalProperties, "vaultBaseUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAzureKeyVaultConnector struct {
	value *AzureKeyVaultConnector
	isSet bool
}

func (v NullableAzureKeyVaultConnector) Get() *AzureKeyVaultConnector {
	return v.value
}

func (v *NullableAzureKeyVaultConnector) Set(val *AzureKeyVaultConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureKeyVaultConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureKeyVaultConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureKeyVaultConnector(val *AzureKeyVaultConnector) *NullableAzureKeyVaultConnector {
	return &NullableAzureKeyVaultConnector{value: val, isSet: true}
}

func (v NullableAzureKeyVaultConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureKeyVaultConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
