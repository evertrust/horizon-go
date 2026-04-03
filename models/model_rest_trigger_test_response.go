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

// checks if the RestTriggerTestResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RestTriggerTestResponse{}

// RestTriggerTestResponse struct for RestTriggerTestResponse
type RestTriggerTestResponse struct {
	// A message describing the test
	Message string `json:"message"`
	// The body from this request
	RequestBody utils.NullableString `json:"requestBody,omitempty"`
	// The headers from this request
	RequestHeaders []RESTHeader `json:"requestHeaders,omitempty"`
	// The URL requested
	RequestURL string `json:"requestURL"`
	// The body from the response to this request
	ResponseBody utils.NullableString `json:"responseBody,omitempty"`
	// The response code to this request
	ResponseCode utils.NullableInt64 `json:"responseCode,omitempty"`
	// The headers from the response to this request
	ResponseHeaders []RESTHeader `json:"responseHeaders,omitempty"`
	// Status of the test
	Status string `json:"status"`
	// The trigger type that was executed
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _RestTriggerTestResponse RestTriggerTestResponse

// NewRestTriggerTestResponse instantiates a new RestTriggerTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestTriggerTestResponse(message string, requestURL string, status string, type_ string) *RestTriggerTestResponse {
	this := RestTriggerTestResponse{}
	this.Message = message
	this.RequestURL = requestURL
	this.Status = status
	this.Type = type_
	return &this
}

// NewRestTriggerTestResponseWithDefaults instantiates a new RestTriggerTestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestTriggerTestResponseWithDefaults() *RestTriggerTestResponse {
	this := RestTriggerTestResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *RestTriggerTestResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RestTriggerTestResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RestTriggerTestResponse) SetMessage(v string) {
	o.Message = v
}

// GetRequestBody returns the RequestBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestTriggerTestResponse) GetRequestBody() string {
	if o == nil || utils.IsNil(o.RequestBody.Get()) {
		var ret string
		return ret
	}
	return *o.RequestBody.Get()
}

// GetRequestBodyOk returns a tuple with the RequestBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestTriggerTestResponse) GetRequestBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestBody.Get(), o.RequestBody.IsSet()
}

// HasRequestBody returns a boolean if a field has been set.
func (o *RestTriggerTestResponse) HasRequestBody() bool {
	if o != nil && o.RequestBody.IsSet() {
		return true
	}

	return false
}

// SetRequestBody gets a reference to the given NullableString and assigns it to the RequestBody field.
func (o *RestTriggerTestResponse) SetRequestBody(v string) {
	o.RequestBody.Set(&v)
}

// SetRequestBodyNil sets the value for RequestBody to be an explicit nil
func (o *RestTriggerTestResponse) SetRequestBodyNil() {
	o.RequestBody.Set(nil)
}

// UnsetRequestBody ensures that no value is present for RequestBody, not even an explicit nil
func (o *RestTriggerTestResponse) UnsetRequestBody() {
	o.RequestBody.Unset()
}

// GetRequestHeaders returns the RequestHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestTriggerTestResponse) GetRequestHeaders() []RESTHeader {
	if o == nil {
		var ret []RESTHeader
		return ret
	}
	return o.RequestHeaders
}

// GetRequestHeadersOk returns a tuple with the RequestHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestTriggerTestResponse) GetRequestHeadersOk() ([]RESTHeader, bool) {
	if o == nil || utils.IsNil(o.RequestHeaders) {
		return nil, false
	}
	return o.RequestHeaders, true
}

// HasRequestHeaders returns a boolean if a field has been set.
func (o *RestTriggerTestResponse) HasRequestHeaders() bool {
	if o != nil && !utils.IsNil(o.RequestHeaders) {
		return true
	}

	return false
}

// SetRequestHeaders gets a reference to the given []RESTHeader and assigns it to the RequestHeaders field.
func (o *RestTriggerTestResponse) SetRequestHeaders(v []RESTHeader) {
	o.RequestHeaders = v
}

// GetRequestURL returns the RequestURL field value
func (o *RestTriggerTestResponse) GetRequestURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestURL
}

// GetRequestURLOk returns a tuple with the RequestURL field value
// and a boolean to check if the value has been set.
func (o *RestTriggerTestResponse) GetRequestURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestURL, true
}

// SetRequestURL sets field value
func (o *RestTriggerTestResponse) SetRequestURL(v string) {
	o.RequestURL = v
}

// GetResponseBody returns the ResponseBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestTriggerTestResponse) GetResponseBody() string {
	if o == nil || utils.IsNil(o.ResponseBody.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseBody.Get()
}

// GetResponseBodyOk returns a tuple with the ResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestTriggerTestResponse) GetResponseBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseBody.Get(), o.ResponseBody.IsSet()
}

// HasResponseBody returns a boolean if a field has been set.
func (o *RestTriggerTestResponse) HasResponseBody() bool {
	if o != nil && o.ResponseBody.IsSet() {
		return true
	}

	return false
}

// SetResponseBody gets a reference to the given NullableString and assigns it to the ResponseBody field.
func (o *RestTriggerTestResponse) SetResponseBody(v string) {
	o.ResponseBody.Set(&v)
}

// SetResponseBodyNil sets the value for ResponseBody to be an explicit nil
func (o *RestTriggerTestResponse) SetResponseBodyNil() {
	o.ResponseBody.Set(nil)
}

// UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
func (o *RestTriggerTestResponse) UnsetResponseBody() {
	o.ResponseBody.Unset()
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestTriggerTestResponse) GetResponseCode() int64 {
	if o == nil || utils.IsNil(o.ResponseCode.Get()) {
		var ret int64
		return ret
	}
	return *o.ResponseCode.Get()
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestTriggerTestResponse) GetResponseCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseCode.Get(), o.ResponseCode.IsSet()
}

// HasResponseCode returns a boolean if a field has been set.
func (o *RestTriggerTestResponse) HasResponseCode() bool {
	if o != nil && o.ResponseCode.IsSet() {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given NullableInt64 and assigns it to the ResponseCode field.
func (o *RestTriggerTestResponse) SetResponseCode(v int64) {
	o.ResponseCode.Set(&v)
}

// SetResponseCodeNil sets the value for ResponseCode to be an explicit nil
func (o *RestTriggerTestResponse) SetResponseCodeNil() {
	o.ResponseCode.Set(nil)
}

// UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
func (o *RestTriggerTestResponse) UnsetResponseCode() {
	o.ResponseCode.Unset()
}

// GetResponseHeaders returns the ResponseHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RestTriggerTestResponse) GetResponseHeaders() []RESTHeader {
	if o == nil {
		var ret []RESTHeader
		return ret
	}
	return o.ResponseHeaders
}

// GetResponseHeadersOk returns a tuple with the ResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RestTriggerTestResponse) GetResponseHeadersOk() ([]RESTHeader, bool) {
	if o == nil || utils.IsNil(o.ResponseHeaders) {
		return nil, false
	}
	return o.ResponseHeaders, true
}

// HasResponseHeaders returns a boolean if a field has been set.
func (o *RestTriggerTestResponse) HasResponseHeaders() bool {
	if o != nil && !utils.IsNil(o.ResponseHeaders) {
		return true
	}

	return false
}

// SetResponseHeaders gets a reference to the given []RESTHeader and assigns it to the ResponseHeaders field.
func (o *RestTriggerTestResponse) SetResponseHeaders(v []RESTHeader) {
	o.ResponseHeaders = v
}

// GetStatus returns the Status field value
func (o *RestTriggerTestResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RestTriggerTestResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RestTriggerTestResponse) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *RestTriggerTestResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RestTriggerTestResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RestTriggerTestResponse) SetType(v string) {
	o.Type = v
}

func (o RestTriggerTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RestTriggerTestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if o.RequestBody.IsSet() {
		toSerialize["requestBody"] = o.RequestBody.Get()
	}
	if o.RequestHeaders != nil {
		toSerialize["requestHeaders"] = o.RequestHeaders
	}
	toSerialize["requestURL"] = o.RequestURL
	if o.ResponseBody.IsSet() {
		toSerialize["responseBody"] = o.ResponseBody.Get()
	}
	if o.ResponseCode.IsSet() {
		toSerialize["responseCode"] = o.ResponseCode.Get()
	}
	if o.ResponseHeaders != nil {
		toSerialize["responseHeaders"] = o.ResponseHeaders
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RestTriggerTestResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"requestURL",
		"status",
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

	varRestTriggerTestResponse := _RestTriggerTestResponse{}

	err = json.Unmarshal(data, &varRestTriggerTestResponse)

	if err != nil {
		return err
	}

	*o = RestTriggerTestResponse(varRestTriggerTestResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "requestBody")
		delete(additionalProperties, "requestHeaders")
		delete(additionalProperties, "requestURL")
		delete(additionalProperties, "responseBody")
		delete(additionalProperties, "responseCode")
		delete(additionalProperties, "responseHeaders")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRestTriggerTestResponse struct {
	value *RestTriggerTestResponse
	isSet bool
}

func (v NullableRestTriggerTestResponse) Get() *RestTriggerTestResponse {
	return v.value
}

func (v *NullableRestTriggerTestResponse) Set(val *RestTriggerTestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRestTriggerTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRestTriggerTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestTriggerTestResponse(val *RestTriggerTestResponse) *NullableRestTriggerTestResponse {
	return &NullableRestTriggerTestResponse{value: val, isSet: true}
}

func (v NullableRestTriggerTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestTriggerTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
