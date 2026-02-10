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

// checks if the EscrowedPrivateKey type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EscrowedPrivateKey{}

// EscrowedPrivateKey struct for EscrowedPrivateKey
type EscrowedPrivateKey struct {
	HorizonKey           string               `json:"horizonKey"`
	Value                utils.NullableString `json:"value,omitempty"`
	VaultKey             utils.NullableString `json:"vaultKey,omitempty"`
	Transient            utils.NullableBool   `json:"transient,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EscrowedPrivateKey EscrowedPrivateKey

// NewEscrowedPrivateKey instantiates a new EscrowedPrivateKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEscrowedPrivateKey(horizonKey string) *EscrowedPrivateKey {
	this := EscrowedPrivateKey{}
	this.HorizonKey = horizonKey
	return &this
}

// NewEscrowedPrivateKeyWithDefaults instantiates a new EscrowedPrivateKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEscrowedPrivateKeyWithDefaults() *EscrowedPrivateKey {
	this := EscrowedPrivateKey{}
	return &this
}

// GetHorizonKey returns the HorizonKey field value
func (o *EscrowedPrivateKey) GetHorizonKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HorizonKey
}

// GetHorizonKeyOk returns a tuple with the HorizonKey field value
// and a boolean to check if the value has been set.
func (o *EscrowedPrivateKey) GetHorizonKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HorizonKey, true
}

// SetHorizonKey sets field value
func (o *EscrowedPrivateKey) SetHorizonKey(v string) {
	o.HorizonKey = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EscrowedPrivateKey) GetValue() string {
	if o == nil || utils.IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EscrowedPrivateKey) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *EscrowedPrivateKey) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *EscrowedPrivateKey) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *EscrowedPrivateKey) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *EscrowedPrivateKey) UnsetValue() {
	o.Value.Unset()
}

// GetVaultKey returns the VaultKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EscrowedPrivateKey) GetVaultKey() string {
	if o == nil || utils.IsNil(o.VaultKey.Get()) {
		var ret string
		return ret
	}
	return *o.VaultKey.Get()
}

// GetVaultKeyOk returns a tuple with the VaultKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EscrowedPrivateKey) GetVaultKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VaultKey.Get(), o.VaultKey.IsSet()
}

// HasVaultKey returns a boolean if a field has been set.
func (o *EscrowedPrivateKey) HasVaultKey() bool {
	if o != nil && o.VaultKey.IsSet() {
		return true
	}

	return false
}

// SetVaultKey gets a reference to the given NullableString and assigns it to the VaultKey field.
func (o *EscrowedPrivateKey) SetVaultKey(v string) {
	o.VaultKey.Set(&v)
}

// SetVaultKeyNil sets the value for VaultKey to be an explicit nil
func (o *EscrowedPrivateKey) SetVaultKeyNil() {
	o.VaultKey.Set(nil)
}

// UnsetVaultKey ensures that no value is present for VaultKey, not even an explicit nil
func (o *EscrowedPrivateKey) UnsetVaultKey() {
	o.VaultKey.Unset()
}

// GetTransient returns the Transient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EscrowedPrivateKey) GetTransient() bool {
	if o == nil || utils.IsNil(o.Transient.Get()) {
		var ret bool
		return ret
	}
	return *o.Transient.Get()
}

// GetTransientOk returns a tuple with the Transient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EscrowedPrivateKey) GetTransientOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transient.Get(), o.Transient.IsSet()
}

// HasTransient returns a boolean if a field has been set.
func (o *EscrowedPrivateKey) HasTransient() bool {
	if o != nil && o.Transient.IsSet() {
		return true
	}

	return false
}

// SetTransient gets a reference to the given NullableBool and assigns it to the Transient field.
func (o *EscrowedPrivateKey) SetTransient(v bool) {
	o.Transient.Set(&v)
}

// SetTransientNil sets the value for Transient to be an explicit nil
func (o *EscrowedPrivateKey) SetTransientNil() {
	o.Transient.Set(nil)
}

// UnsetTransient ensures that no value is present for Transient, not even an explicit nil
func (o *EscrowedPrivateKey) UnsetTransient() {
	o.Transient.Unset()
}

func (o EscrowedPrivateKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EscrowedPrivateKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["horizonKey"] = o.HorizonKey
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.VaultKey.IsSet() {
		toSerialize["vaultKey"] = o.VaultKey.Get()
	}
	if o.Transient.IsSet() {
		toSerialize["transient"] = o.Transient.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EscrowedPrivateKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"horizonKey",
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

	varEscrowedPrivateKey := _EscrowedPrivateKey{}

	err = json.Unmarshal(data, &varEscrowedPrivateKey)

	if err != nil {
		return err
	}

	*o = EscrowedPrivateKey(varEscrowedPrivateKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "horizonKey")
		delete(additionalProperties, "value")
		delete(additionalProperties, "vaultKey")
		delete(additionalProperties, "transient")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEscrowedPrivateKey struct {
	value *EscrowedPrivateKey
	isSet bool
}

func (v NullableEscrowedPrivateKey) Get() *EscrowedPrivateKey {
	return v.value
}

func (v *NullableEscrowedPrivateKey) Set(val *EscrowedPrivateKey) {
	v.value = val
	v.isSet = true
}

func (v NullableEscrowedPrivateKey) IsSet() bool {
	return v.isSet
}

func (v *NullableEscrowedPrivateKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEscrowedPrivateKey(val *EscrowedPrivateKey) *NullableEscrowedPrivateKey {
	return &NullableEscrowedPrivateKey{value: val, isSet: true}
}

func (v NullableEscrowedPrivateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEscrowedPrivateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
