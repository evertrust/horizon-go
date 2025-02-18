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

// checks if the PasswordPolicyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyResponse{}

// PasswordPolicyResponse struct for PasswordPolicyResponse
type PasswordPolicyResponse struct {
	// The internal ID of the password policy
	Id string `json:"_id"`
	// The name of the password policy
	Name string `json:"name"`
	// The minimum number of characters of the password
	MinChar int64 `json:"minChar"`
	// The maximum number of characters of the password
	MaxChar NullableInt64 `json:"maxChar,omitempty"`
	// The minimum number of uppercase characters of the password
	MinUpChar NullableInt64 `json:"minUpChar,omitempty"`
	// The minimum number of lowercase characters of the password
	MinLoChar NullableInt64 `json:"minLoChar,omitempty"`
	// The minimum number of digits of the password
	MinDiChar NullableInt64 `json:"minDiChar,omitempty"`
	// The special characters of the password accepted by the password policy
	SpChar NullableString `json:"spChar,omitempty"`
	// The minimum number of special characters of the password
	MinSpChar NullableInt64 `json:"minSpChar,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PasswordPolicyResponse PasswordPolicyResponse

// NewPasswordPolicyResponse instantiates a new PasswordPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyResponse(id string, name string, minChar int64) *PasswordPolicyResponse {
	this := PasswordPolicyResponse{}
	this.Id = id
	this.Name = name
	this.MinChar = minChar
	return &this
}

// NewPasswordPolicyResponseWithDefaults instantiates a new PasswordPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyResponseWithDefaults() *PasswordPolicyResponse {
	this := PasswordPolicyResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PasswordPolicyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PasswordPolicyResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *PasswordPolicyResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicyResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PasswordPolicyResponse) SetName(v string) {
	o.Name = v
}

// GetMinChar returns the MinChar field value
func (o *PasswordPolicyResponse) GetMinChar() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinChar
}

// GetMinCharOk returns a tuple with the MinChar field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicyResponse) GetMinCharOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinChar, true
}

// SetMinChar sets field value
func (o *PasswordPolicyResponse) SetMinChar(v int64) {
	o.MinChar = v
}

// GetMaxChar returns the MaxChar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicyResponse) GetMaxChar() int64 {
	if o == nil || IsNil(o.MaxChar.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxChar.Get()
}

// GetMaxCharOk returns a tuple with the MaxChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicyResponse) GetMaxCharOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxChar.Get(), o.MaxChar.IsSet()
}

// HasMaxChar returns a boolean if a field has been set.
func (o *PasswordPolicyResponse) HasMaxChar() bool {
	if o != nil && o.MaxChar.IsSet() {
		return true
	}

	return false
}

// SetMaxChar gets a reference to the given NullableInt64 and assigns it to the MaxChar field.
func (o *PasswordPolicyResponse) SetMaxChar(v int64) {
	o.MaxChar.Set(&v)
}
// SetMaxCharNil sets the value for MaxChar to be an explicit nil
func (o *PasswordPolicyResponse) SetMaxCharNil() {
	o.MaxChar.Set(nil)
}

// UnsetMaxChar ensures that no value is present for MaxChar, not even an explicit nil
func (o *PasswordPolicyResponse) UnsetMaxChar() {
	o.MaxChar.Unset()
}

// GetMinUpChar returns the MinUpChar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicyResponse) GetMinUpChar() int64 {
	if o == nil || IsNil(o.MinUpChar.Get()) {
		var ret int64
		return ret
	}
	return *o.MinUpChar.Get()
}

// GetMinUpCharOk returns a tuple with the MinUpChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicyResponse) GetMinUpCharOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinUpChar.Get(), o.MinUpChar.IsSet()
}

// HasMinUpChar returns a boolean if a field has been set.
func (o *PasswordPolicyResponse) HasMinUpChar() bool {
	if o != nil && o.MinUpChar.IsSet() {
		return true
	}

	return false
}

// SetMinUpChar gets a reference to the given NullableInt64 and assigns it to the MinUpChar field.
func (o *PasswordPolicyResponse) SetMinUpChar(v int64) {
	o.MinUpChar.Set(&v)
}
// SetMinUpCharNil sets the value for MinUpChar to be an explicit nil
func (o *PasswordPolicyResponse) SetMinUpCharNil() {
	o.MinUpChar.Set(nil)
}

// UnsetMinUpChar ensures that no value is present for MinUpChar, not even an explicit nil
func (o *PasswordPolicyResponse) UnsetMinUpChar() {
	o.MinUpChar.Unset()
}

// GetMinLoChar returns the MinLoChar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicyResponse) GetMinLoChar() int64 {
	if o == nil || IsNil(o.MinLoChar.Get()) {
		var ret int64
		return ret
	}
	return *o.MinLoChar.Get()
}

// GetMinLoCharOk returns a tuple with the MinLoChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicyResponse) GetMinLoCharOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinLoChar.Get(), o.MinLoChar.IsSet()
}

// HasMinLoChar returns a boolean if a field has been set.
func (o *PasswordPolicyResponse) HasMinLoChar() bool {
	if o != nil && o.MinLoChar.IsSet() {
		return true
	}

	return false
}

// SetMinLoChar gets a reference to the given NullableInt64 and assigns it to the MinLoChar field.
func (o *PasswordPolicyResponse) SetMinLoChar(v int64) {
	o.MinLoChar.Set(&v)
}
// SetMinLoCharNil sets the value for MinLoChar to be an explicit nil
func (o *PasswordPolicyResponse) SetMinLoCharNil() {
	o.MinLoChar.Set(nil)
}

// UnsetMinLoChar ensures that no value is present for MinLoChar, not even an explicit nil
func (o *PasswordPolicyResponse) UnsetMinLoChar() {
	o.MinLoChar.Unset()
}

// GetMinDiChar returns the MinDiChar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicyResponse) GetMinDiChar() int64 {
	if o == nil || IsNil(o.MinDiChar.Get()) {
		var ret int64
		return ret
	}
	return *o.MinDiChar.Get()
}

// GetMinDiCharOk returns a tuple with the MinDiChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicyResponse) GetMinDiCharOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinDiChar.Get(), o.MinDiChar.IsSet()
}

// HasMinDiChar returns a boolean if a field has been set.
func (o *PasswordPolicyResponse) HasMinDiChar() bool {
	if o != nil && o.MinDiChar.IsSet() {
		return true
	}

	return false
}

// SetMinDiChar gets a reference to the given NullableInt64 and assigns it to the MinDiChar field.
func (o *PasswordPolicyResponse) SetMinDiChar(v int64) {
	o.MinDiChar.Set(&v)
}
// SetMinDiCharNil sets the value for MinDiChar to be an explicit nil
func (o *PasswordPolicyResponse) SetMinDiCharNil() {
	o.MinDiChar.Set(nil)
}

// UnsetMinDiChar ensures that no value is present for MinDiChar, not even an explicit nil
func (o *PasswordPolicyResponse) UnsetMinDiChar() {
	o.MinDiChar.Unset()
}

// GetSpChar returns the SpChar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicyResponse) GetSpChar() string {
	if o == nil || IsNil(o.SpChar.Get()) {
		var ret string
		return ret
	}
	return *o.SpChar.Get()
}

// GetSpCharOk returns a tuple with the SpChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicyResponse) GetSpCharOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpChar.Get(), o.SpChar.IsSet()
}

// HasSpChar returns a boolean if a field has been set.
func (o *PasswordPolicyResponse) HasSpChar() bool {
	if o != nil && o.SpChar.IsSet() {
		return true
	}

	return false
}

// SetSpChar gets a reference to the given NullableString and assigns it to the SpChar field.
func (o *PasswordPolicyResponse) SetSpChar(v string) {
	o.SpChar.Set(&v)
}
// SetSpCharNil sets the value for SpChar to be an explicit nil
func (o *PasswordPolicyResponse) SetSpCharNil() {
	o.SpChar.Set(nil)
}

// UnsetSpChar ensures that no value is present for SpChar, not even an explicit nil
func (o *PasswordPolicyResponse) UnsetSpChar() {
	o.SpChar.Unset()
}

// GetMinSpChar returns the MinSpChar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PasswordPolicyResponse) GetMinSpChar() int64 {
	if o == nil || IsNil(o.MinSpChar.Get()) {
		var ret int64
		return ret
	}
	return *o.MinSpChar.Get()
}

// GetMinSpCharOk returns a tuple with the MinSpChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PasswordPolicyResponse) GetMinSpCharOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinSpChar.Get(), o.MinSpChar.IsSet()
}

// HasMinSpChar returns a boolean if a field has been set.
func (o *PasswordPolicyResponse) HasMinSpChar() bool {
	if o != nil && o.MinSpChar.IsSet() {
		return true
	}

	return false
}

// SetMinSpChar gets a reference to the given NullableInt64 and assigns it to the MinSpChar field.
func (o *PasswordPolicyResponse) SetMinSpChar(v int64) {
	o.MinSpChar.Set(&v)
}
// SetMinSpCharNil sets the value for MinSpChar to be an explicit nil
func (o *PasswordPolicyResponse) SetMinSpCharNil() {
	o.MinSpChar.Set(nil)
}

// UnsetMinSpChar ensures that no value is present for MinSpChar, not even an explicit nil
func (o *PasswordPolicyResponse) UnsetMinSpChar() {
	o.MinSpChar.Unset()
}

func (o PasswordPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["minChar"] = o.MinChar
	if o.MaxChar.IsSet() {
		toSerialize["maxChar"] = o.MaxChar.Get()
	}
	if o.MinUpChar.IsSet() {
		toSerialize["minUpChar"] = o.MinUpChar.Get()
	}
	if o.MinLoChar.IsSet() {
		toSerialize["minLoChar"] = o.MinLoChar.Get()
	}
	if o.MinDiChar.IsSet() {
		toSerialize["minDiChar"] = o.MinDiChar.Get()
	}
	if o.SpChar.IsSet() {
		toSerialize["spChar"] = o.SpChar.Get()
	}
	if o.MinSpChar.IsSet() {
		toSerialize["minSpChar"] = o.MinSpChar.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PasswordPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"name",
		"minChar",
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

	varPasswordPolicyResponse := _PasswordPolicyResponse{}

	err = json.Unmarshal(data, &varPasswordPolicyResponse)

	if err != nil {
		return err
	}

	*o = PasswordPolicyResponse(varPasswordPolicyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "minChar")
		delete(additionalProperties, "maxChar")
		delete(additionalProperties, "minUpChar")
		delete(additionalProperties, "minLoChar")
		delete(additionalProperties, "minDiChar")
		delete(additionalProperties, "spChar")
		delete(additionalProperties, "minSpChar")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePasswordPolicyResponse struct {
	value *PasswordPolicyResponse
	isSet bool
}

func (v NullablePasswordPolicyResponse) Get() *PasswordPolicyResponse {
	return v.value
}

func (v *NullablePasswordPolicyResponse) Set(val *PasswordPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyResponse(val *PasswordPolicyResponse) *NullablePasswordPolicyResponse {
	return &NullablePasswordPolicyResponse{value: val, isSet: true}
}

func (v NullablePasswordPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


