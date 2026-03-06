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

// checks if the DiscFeed004 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DiscFeed004{}

// DiscFeed004 Invalid Discovery Feed
type DiscFeed004 struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Title   string `json:"title"`
	// A human-readable explanation specific to this occurrence of the problem. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807)
	Detail utils.NullableString `json:"detail,omitempty"`
	// The http status code of the error. In compliance with [RFC7807](https://datatracker.ietf.org/doc/html/rfc7807)
	Status               int64 `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _DiscFeed004 DiscFeed004

// NewDiscFeed004 instantiates a new DiscFeed004 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscFeed004(error_ string, message string, title string, status int64) *DiscFeed004 {
	this := DiscFeed004{}
	this.Error = error_
	this.Message = message
	this.Status = status
	this.Title = title
	return &this
}

// NewDiscFeed004WithDefaults instantiates a new DiscFeed004 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscFeed004WithDefaults() *DiscFeed004 {
	this := DiscFeed004{}
	return &this
}

// GetError returns the Error field value
func (o *DiscFeed004) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *DiscFeed004) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *DiscFeed004) SetError(v string) {
	o.Error = v
}

// GetMessage returns the Message field value
func (o *DiscFeed004) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DiscFeed004) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *DiscFeed004) SetMessage(v string) {
	o.Message = v
}

// GetTitle returns the Title field value
func (o *DiscFeed004) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DiscFeed004) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DiscFeed004) SetTitle(v string) {
	o.Title = v
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscFeed004) GetDetail() string {
	if o == nil || utils.IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscFeed004) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *DiscFeed004) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *DiscFeed004) SetDetail(v string) {
	o.Detail.Set(&v)
}

// SetDetailNil sets the value for Detail to be an explicit nil
func (o *DiscFeed004) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *DiscFeed004) UnsetDetail() {
	o.Detail.Unset()
}

// GetStatus returns the Status field value
func (o *DiscFeed004) GetStatus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DiscFeed004) GetStatusOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DiscFeed004) SetStatus(v int64) {
	o.Status = v
}

func (o DiscFeed004) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscFeed004) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["message"] = o.Message
	toSerialize["title"] = o.Title
	if o.Detail.IsSet() {
		toSerialize["detail"] = o.Detail.Get()
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscFeed004) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
		"message",
		"title",
		"status",
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

	varDiscFeed004 := _DiscFeed004{}

	err = json.Unmarshal(data, &varDiscFeed004)

	if err != nil {
		return err
	}

	*o = DiscFeed004(varDiscFeed004)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		delete(additionalProperties, "message")
		delete(additionalProperties, "title")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscFeed004 struct {
	value *DiscFeed004
	isSet bool
}

func (v NullableDiscFeed004) Get() *DiscFeed004 {
	return v.value
}

func (v *NullableDiscFeed004) Set(val *DiscFeed004) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscFeed004) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscFeed004) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscFeed004(val *DiscFeed004) *NullableDiscFeed004 {
	return &NullableDiscFeed004{value: val, isSet: true}
}

func (v NullableDiscFeed004) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscFeed004) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
