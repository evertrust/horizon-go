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

// checks if the DNSDatasource type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DNSDatasource{}

// DNSDatasource struct for DNSDatasource
type DNSDatasource struct {
	// Type of datasource
	Type string `json:"type"`
	// Name of the datasource
	Name string `json:"name"`
	// The localized name of the datasource
	DisplayName []LocalizedString `json:"displayName,omitempty"`
	// Description of the datasource
	Description *string `json:"description,omitempty"`
	// Ip of the DNS server. If empty, Horizon Server DNS is used
	Host utils.NullableString `json:"host,omitempty"`
	// Port on which to join the DNS server
	Port utils.NullableInt64 `json:"port,omitempty"`
	// Timeout for the DNS request
	Timeout utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Type of DNS records to fetch. All available record types are fetched if null
	RecordTypes []string `json:"recordTypes,omitempty"`
	// Host to lookup
	Lookup               string `json:"lookup"`
	AdditionalProperties map[string]interface{}
}

type _DNSDatasource DNSDatasource

// NewDNSDatasource instantiates a new DNSDatasource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNSDatasource(type_ string, name string, lookup string) *DNSDatasource {
	this := DNSDatasource{}
	this.Type = type_
	this.Name = name
	var port int64 = 53
	this.Port = *utils.NewNullableInt64(&port)
	var timeout string = "10 seconds"
	this.Timeout = *utils.NewNullableString(&timeout)
	this.Lookup = lookup
	return &this
}

// NewDNSDatasourceWithDefaults instantiates a new DNSDatasource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNSDatasourceWithDefaults() *DNSDatasource {
	this := DNSDatasource{}
	var port int64 = 53
	this.Port = *utils.NewNullableInt64(&port)
	var timeout string = "10 seconds"
	this.Timeout = *utils.NewNullableString(&timeout)
	return &this
}

// GetType returns the Type field value
func (o *DNSDatasource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DNSDatasource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DNSDatasource) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *DNSDatasource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DNSDatasource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DNSDatasource) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DNSDatasource) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DNSDatasource) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DNSDatasource) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *DNSDatasource) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DNSDatasource) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DNSDatasource) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DNSDatasource) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DNSDatasource) SetDescription(v string) {
	o.Description = &v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DNSDatasource) GetHost() string {
	if o == nil || utils.IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DNSDatasource) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *DNSDatasource) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *DNSDatasource) SetHost(v string) {
	o.Host.Set(&v)
}

// SetHostNil sets the value for Host to be an explicit nil
func (o *DNSDatasource) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *DNSDatasource) UnsetHost() {
	o.Host.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DNSDatasource) GetPort() int64 {
	if o == nil || utils.IsNil(o.Port.Get()) {
		var ret int64
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DNSDatasource) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *DNSDatasource) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt64 and assigns it to the Port field.
func (o *DNSDatasource) SetPort(v int64) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil
func (o *DNSDatasource) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *DNSDatasource) UnsetPort() {
	o.Port.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DNSDatasource) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DNSDatasource) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *DNSDatasource) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *DNSDatasource) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *DNSDatasource) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *DNSDatasource) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetRecordTypes returns the RecordTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DNSDatasource) GetRecordTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RecordTypes
}

// GetRecordTypesOk returns a tuple with the RecordTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DNSDatasource) GetRecordTypesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.RecordTypes) {
		return nil, false
	}
	return o.RecordTypes, true
}

// HasRecordTypes returns a boolean if a field has been set.
func (o *DNSDatasource) HasRecordTypes() bool {
	if o != nil && !utils.IsNil(o.RecordTypes) {
		return true
	}

	return false
}

// SetRecordTypes gets a reference to the given []string and assigns it to the RecordTypes field.
func (o *DNSDatasource) SetRecordTypes(v []string) {
	o.RecordTypes = v
}

// GetLookup returns the Lookup field value
func (o *DNSDatasource) GetLookup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Lookup
}

// GetLookupOk returns a tuple with the Lookup field value
// and a boolean to check if the value has been set.
func (o *DNSDatasource) GetLookupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Lookup, true
}

// SetLookup sets field value
func (o *DNSDatasource) SetLookup(v string) {
	o.Lookup = v
}

func (o DNSDatasource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNSDatasource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	if o.RecordTypes != nil {
		toSerialize["recordTypes"] = o.RecordTypes
	}
	toSerialize["lookup"] = o.Lookup

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DNSDatasource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"name",
		"lookup",
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

	varDNSDatasource := _DNSDatasource{}

	err = json.Unmarshal(data, &varDNSDatasource)

	if err != nil {
		return err
	}

	*o = DNSDatasource(varDNSDatasource)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "host")
		delete(additionalProperties, "port")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "recordTypes")
		delete(additionalProperties, "lookup")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDNSDatasource struct {
	value *DNSDatasource
	isSet bool
}

func (v NullableDNSDatasource) Get() *DNSDatasource {
	return v.value
}

func (v *NullableDNSDatasource) Set(val *DNSDatasource) {
	v.value = val
	v.isSet = true
}

func (v NullableDNSDatasource) IsSet() bool {
	return v.isSet
}

func (v *NullableDNSDatasource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNSDatasource(val *DNSDatasource) *NullableDNSDatasource {
	return &NullableDNSDatasource{value: val, isSet: true}
}

func (v NullableDNSDatasource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNSDatasource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
