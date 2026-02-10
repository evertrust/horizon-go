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

// checks if the LDAPDataSourceResultResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LDAPDataSourceResultResponse{}

// LDAPDataSourceResultResponse struct for LDAPDataSourceResultResponse
type LDAPDataSourceResultResponse struct {
	Type string `json:"type"`
	// DN that was requested on the LDAP server
	ComputedDN utils.NullableString `json:"computedDN,omitempty"`
	// Filter that was requested on the LDAP server
	ComputedFilter utils.NullableString `json:"computedFilter,omitempty"`
	// Name of the executed datasource
	Name string `json:"name"`
	// Status of the execution. `success` if the datasource data was fetched correctly, `failure` if an error occured and `ignored` if inputs were not all filled, resulting in no request being sent
	Status string `json:"status"`
	// Data fetched from the datasource
	Dictionary []MapEntry `json:"dictionary"`
	// If `status` is `failure`, the error message
	Error                utils.NullableString `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LDAPDataSourceResultResponse LDAPDataSourceResultResponse

// NewLDAPDataSourceResultResponse instantiates a new LDAPDataSourceResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPDataSourceResultResponse(type_ string, name string, status string, dictionary []MapEntry) *LDAPDataSourceResultResponse {
	this := LDAPDataSourceResultResponse{}
	this.Name = name
	this.Type = type_
	this.Status = status
	this.Dictionary = dictionary
	return &this
}

// NewLDAPDataSourceResultResponseWithDefaults instantiates a new LDAPDataSourceResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPDataSourceResultResponseWithDefaults() *LDAPDataSourceResultResponse {
	this := LDAPDataSourceResultResponse{}
	return &this
}

// GetType returns the Type field value
func (o *LDAPDataSourceResultResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LDAPDataSourceResultResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LDAPDataSourceResultResponse) SetType(v string) {
	o.Type = v
}

// GetComputedDN returns the ComputedDN field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPDataSourceResultResponse) GetComputedDN() string {
	if o == nil || utils.IsNil(o.ComputedDN.Get()) {
		var ret string
		return ret
	}
	return *o.ComputedDN.Get()
}

// GetComputedDNOk returns a tuple with the ComputedDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPDataSourceResultResponse) GetComputedDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputedDN.Get(), o.ComputedDN.IsSet()
}

// HasComputedDN returns a boolean if a field has been set.
func (o *LDAPDataSourceResultResponse) HasComputedDN() bool {
	if o != nil && o.ComputedDN.IsSet() {
		return true
	}

	return false
}

// SetComputedDN gets a reference to the given NullableString and assigns it to the ComputedDN field.
func (o *LDAPDataSourceResultResponse) SetComputedDN(v string) {
	o.ComputedDN.Set(&v)
}

// SetComputedDNNil sets the value for ComputedDN to be an explicit nil
func (o *LDAPDataSourceResultResponse) SetComputedDNNil() {
	o.ComputedDN.Set(nil)
}

// UnsetComputedDN ensures that no value is present for ComputedDN, not even an explicit nil
func (o *LDAPDataSourceResultResponse) UnsetComputedDN() {
	o.ComputedDN.Unset()
}

// GetComputedFilter returns the ComputedFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPDataSourceResultResponse) GetComputedFilter() string {
	if o == nil || utils.IsNil(o.ComputedFilter.Get()) {
		var ret string
		return ret
	}
	return *o.ComputedFilter.Get()
}

// GetComputedFilterOk returns a tuple with the ComputedFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPDataSourceResultResponse) GetComputedFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputedFilter.Get(), o.ComputedFilter.IsSet()
}

// HasComputedFilter returns a boolean if a field has been set.
func (o *LDAPDataSourceResultResponse) HasComputedFilter() bool {
	if o != nil && o.ComputedFilter.IsSet() {
		return true
	}

	return false
}

// SetComputedFilter gets a reference to the given NullableString and assigns it to the ComputedFilter field.
func (o *LDAPDataSourceResultResponse) SetComputedFilter(v string) {
	o.ComputedFilter.Set(&v)
}

// SetComputedFilterNil sets the value for ComputedFilter to be an explicit nil
func (o *LDAPDataSourceResultResponse) SetComputedFilterNil() {
	o.ComputedFilter.Set(nil)
}

// UnsetComputedFilter ensures that no value is present for ComputedFilter, not even an explicit nil
func (o *LDAPDataSourceResultResponse) UnsetComputedFilter() {
	o.ComputedFilter.Unset()
}

// GetName returns the Name field value
func (o *LDAPDataSourceResultResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LDAPDataSourceResultResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LDAPDataSourceResultResponse) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *LDAPDataSourceResultResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LDAPDataSourceResultResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LDAPDataSourceResultResponse) SetStatus(v string) {
	o.Status = v
}

// GetDictionary returns the Dictionary field value
func (o *LDAPDataSourceResultResponse) GetDictionary() []MapEntry {
	if o == nil {
		var ret []MapEntry
		return ret
	}

	return o.Dictionary
}

// GetDictionaryOk returns a tuple with the Dictionary field value
// and a boolean to check if the value has been set.
func (o *LDAPDataSourceResultResponse) GetDictionaryOk() ([]MapEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dictionary, true
}

// SetDictionary sets field value
func (o *LDAPDataSourceResultResponse) SetDictionary(v []MapEntry) {
	o.Dictionary = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPDataSourceResultResponse) GetError() string {
	if o == nil || utils.IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPDataSourceResultResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *LDAPDataSourceResultResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *LDAPDataSourceResultResponse) SetError(v string) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *LDAPDataSourceResultResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *LDAPDataSourceResultResponse) UnsetError() {
	o.Error.Unset()
}

func (o LDAPDataSourceResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LDAPDataSourceResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.ComputedDN.IsSet() {
		toSerialize["computedDN"] = o.ComputedDN.Get()
	}
	if o.ComputedFilter.IsSet() {
		toSerialize["computedFilter"] = o.ComputedFilter.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["dictionary"] = o.Dictionary
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LDAPDataSourceResultResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"status",
		"dictionary",
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

	varLDAPDataSourceResultResponse := _LDAPDataSourceResultResponse{}

	err = json.Unmarshal(data, &varLDAPDataSourceResultResponse)

	if err != nil {
		return err
	}

	*o = LDAPDataSourceResultResponse(varLDAPDataSourceResultResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "computedDN")
		delete(additionalProperties, "computedFilter")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "dictionary")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLDAPDataSourceResultResponse struct {
	value *LDAPDataSourceResultResponse
	isSet bool
}

func (v NullableLDAPDataSourceResultResponse) Get() *LDAPDataSourceResultResponse {
	return v.value
}

func (v *NullableLDAPDataSourceResultResponse) Set(val *LDAPDataSourceResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPDataSourceResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPDataSourceResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPDataSourceResultResponse(val *LDAPDataSourceResultResponse) *NullableLDAPDataSourceResultResponse {
	return &NullableLDAPDataSourceResultResponse{value: val, isSet: true}
}

func (v NullableLDAPDataSourceResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPDataSourceResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
