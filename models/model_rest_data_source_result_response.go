/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the RESTDataSourceResultResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RESTDataSourceResultResponse{}

// RESTDataSourceResultResponse struct for RESTDataSourceResultResponse
type RESTDataSourceResultResponse struct {
	// Headers that were sent
	ComputedHeaders []map[string]interface{} `json:"computedHeaders,omitempty"`
	// Url that was requested
	ComputedPayload utils.NullableString `json:"computedPayload,omitempty"`
	// Url that was requested
	ComputedUrl utils.NullableString `json:"computedUrl,omitempty"`
	// Received response body
	ResponseBody utils.NullableString `json:"responseBody,omitempty"`
	// Received response code
	ResponseCode utils.NullableInt64 `json:"responseCode,omitempty"`
	// Headers that were received
	ResponseHeaders []map[string]interface{} `json:"responseHeaders,omitempty"`
	Type            string                   `json:"type"`
	// Data fetched from the datasource
	Dictionary []MapEntry `json:"dictionary"`
	// If `status` is `failure`, the error message
	Error utils.NullableString `json:"error,omitempty"`
	// Name of the executed datasource
	Name string `json:"name"`
	// Status of the execution. `success` if the datasource data was fetched correctly, `failure` if an error occured, `not_found` if the datasource query returned no results and `ignored` if inputs were not all filled, resulting in no request being sent
	Status               string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _RESTDataSourceResultResponse RESTDataSourceResultResponse

// NewRESTDataSourceResultResponse instantiates a new RESTDataSourceResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRESTDataSourceResultResponse(type_ string, dictionary []MapEntry, name string, status string) *RESTDataSourceResultResponse {
	this := RESTDataSourceResultResponse{}
	this.Type = type_
	this.Dictionary = dictionary
	this.Name = name
	this.Status = status
	return &this
}

// NewRESTDataSourceResultResponseWithDefaults instantiates a new RESTDataSourceResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRESTDataSourceResultResponseWithDefaults() *RESTDataSourceResultResponse {
	this := RESTDataSourceResultResponse{}
	return &this
}

// GetComputedHeaders returns the ComputedHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDataSourceResultResponse) GetComputedHeaders() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.ComputedHeaders
}

// GetComputedHeadersOk returns a tuple with the ComputedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDataSourceResultResponse) GetComputedHeadersOk() ([]map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.ComputedHeaders) {
		return nil, false
	}
	return o.ComputedHeaders, true
}

// HasComputedHeaders returns a boolean if a field has been set.
func (o *RESTDataSourceResultResponse) HasComputedHeaders() bool {
	if o != nil && !utils.IsNil(o.ComputedHeaders) {
		return true
	}

	return false
}

// SetComputedHeaders gets a reference to the given []map[string]interface{} and assigns it to the ComputedHeaders field.
func (o *RESTDataSourceResultResponse) SetComputedHeaders(v []map[string]interface{}) {
	o.ComputedHeaders = v
}

// GetComputedPayload returns the ComputedPayload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDataSourceResultResponse) GetComputedPayload() string {
	if o == nil || utils.IsNil(o.ComputedPayload.Get()) {
		var ret string
		return ret
	}
	return *o.ComputedPayload.Get()
}

// GetComputedPayloadOk returns a tuple with the ComputedPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDataSourceResultResponse) GetComputedPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputedPayload.Get(), o.ComputedPayload.IsSet()
}

// HasComputedPayload returns a boolean if a field has been set.
func (o *RESTDataSourceResultResponse) HasComputedPayload() bool {
	if o != nil && o.ComputedPayload.IsSet() {
		return true
	}

	return false
}

// SetComputedPayload gets a reference to the given NullableString and assigns it to the ComputedPayload field.
func (o *RESTDataSourceResultResponse) SetComputedPayload(v string) {
	o.ComputedPayload.Set(&v)
}

// SetComputedPayloadNil sets the value for ComputedPayload to be an explicit nil
func (o *RESTDataSourceResultResponse) SetComputedPayloadNil() {
	o.ComputedPayload.Set(nil)
}

// UnsetComputedPayload ensures that no value is present for ComputedPayload, not even an explicit nil
func (o *RESTDataSourceResultResponse) UnsetComputedPayload() {
	o.ComputedPayload.Unset()
}

// GetComputedUrl returns the ComputedUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDataSourceResultResponse) GetComputedUrl() string {
	if o == nil || utils.IsNil(o.ComputedUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ComputedUrl.Get()
}

// GetComputedUrlOk returns a tuple with the ComputedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDataSourceResultResponse) GetComputedUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputedUrl.Get(), o.ComputedUrl.IsSet()
}

// HasComputedUrl returns a boolean if a field has been set.
func (o *RESTDataSourceResultResponse) HasComputedUrl() bool {
	if o != nil && o.ComputedUrl.IsSet() {
		return true
	}

	return false
}

// SetComputedUrl gets a reference to the given NullableString and assigns it to the ComputedUrl field.
func (o *RESTDataSourceResultResponse) SetComputedUrl(v string) {
	o.ComputedUrl.Set(&v)
}

// SetComputedUrlNil sets the value for ComputedUrl to be an explicit nil
func (o *RESTDataSourceResultResponse) SetComputedUrlNil() {
	o.ComputedUrl.Set(nil)
}

// UnsetComputedUrl ensures that no value is present for ComputedUrl, not even an explicit nil
func (o *RESTDataSourceResultResponse) UnsetComputedUrl() {
	o.ComputedUrl.Unset()
}

// GetResponseBody returns the ResponseBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDataSourceResultResponse) GetResponseBody() string {
	if o == nil || utils.IsNil(o.ResponseBody.Get()) {
		var ret string
		return ret
	}
	return *o.ResponseBody.Get()
}

// GetResponseBodyOk returns a tuple with the ResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDataSourceResultResponse) GetResponseBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseBody.Get(), o.ResponseBody.IsSet()
}

// HasResponseBody returns a boolean if a field has been set.
func (o *RESTDataSourceResultResponse) HasResponseBody() bool {
	if o != nil && o.ResponseBody.IsSet() {
		return true
	}

	return false
}

// SetResponseBody gets a reference to the given NullableString and assigns it to the ResponseBody field.
func (o *RESTDataSourceResultResponse) SetResponseBody(v string) {
	o.ResponseBody.Set(&v)
}

// SetResponseBodyNil sets the value for ResponseBody to be an explicit nil
func (o *RESTDataSourceResultResponse) SetResponseBodyNil() {
	o.ResponseBody.Set(nil)
}

// UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
func (o *RESTDataSourceResultResponse) UnsetResponseBody() {
	o.ResponseBody.Unset()
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDataSourceResultResponse) GetResponseCode() int64 {
	if o == nil || utils.IsNil(o.ResponseCode.Get()) {
		var ret int64
		return ret
	}
	return *o.ResponseCode.Get()
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDataSourceResultResponse) GetResponseCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponseCode.Get(), o.ResponseCode.IsSet()
}

// HasResponseCode returns a boolean if a field has been set.
func (o *RESTDataSourceResultResponse) HasResponseCode() bool {
	if o != nil && o.ResponseCode.IsSet() {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given NullableInt64 and assigns it to the ResponseCode field.
func (o *RESTDataSourceResultResponse) SetResponseCode(v int64) {
	o.ResponseCode.Set(&v)
}

// SetResponseCodeNil sets the value for ResponseCode to be an explicit nil
func (o *RESTDataSourceResultResponse) SetResponseCodeNil() {
	o.ResponseCode.Set(nil)
}

// UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
func (o *RESTDataSourceResultResponse) UnsetResponseCode() {
	o.ResponseCode.Unset()
}

// GetResponseHeaders returns the ResponseHeaders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDataSourceResultResponse) GetResponseHeaders() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.ResponseHeaders
}

// GetResponseHeadersOk returns a tuple with the ResponseHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDataSourceResultResponse) GetResponseHeadersOk() ([]map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.ResponseHeaders) {
		return nil, false
	}
	return o.ResponseHeaders, true
}

// HasResponseHeaders returns a boolean if a field has been set.
func (o *RESTDataSourceResultResponse) HasResponseHeaders() bool {
	if o != nil && !utils.IsNil(o.ResponseHeaders) {
		return true
	}

	return false
}

// SetResponseHeaders gets a reference to the given []map[string]interface{} and assigns it to the ResponseHeaders field.
func (o *RESTDataSourceResultResponse) SetResponseHeaders(v []map[string]interface{}) {
	o.ResponseHeaders = v
}

// GetType returns the Type field value
func (o *RESTDataSourceResultResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RESTDataSourceResultResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RESTDataSourceResultResponse) SetType(v string) {
	o.Type = v
}

// GetDictionary returns the Dictionary field value
func (o *RESTDataSourceResultResponse) GetDictionary() []MapEntry {
	if o == nil {
		var ret []MapEntry
		return ret
	}

	return o.Dictionary
}

// GetDictionaryOk returns a tuple with the Dictionary field value
// and a boolean to check if the value has been set.
func (o *RESTDataSourceResultResponse) GetDictionaryOk() ([]MapEntry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dictionary, true
}

// SetDictionary sets field value
func (o *RESTDataSourceResultResponse) SetDictionary(v []MapEntry) {
	o.Dictionary = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDataSourceResultResponse) GetError() string {
	if o == nil || utils.IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDataSourceResultResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *RESTDataSourceResultResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *RESTDataSourceResultResponse) SetError(v string) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *RESTDataSourceResultResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *RESTDataSourceResultResponse) UnsetError() {
	o.Error.Unset()
}

// GetName returns the Name field value
func (o *RESTDataSourceResultResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RESTDataSourceResultResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RESTDataSourceResultResponse) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *RESTDataSourceResultResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RESTDataSourceResultResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RESTDataSourceResultResponse) SetStatus(v string) {
	o.Status = v
}

func (o RESTDataSourceResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RESTDataSourceResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ComputedHeaders != nil {
		toSerialize["computedHeaders"] = o.ComputedHeaders
	}
	if o.ComputedPayload.IsSet() {
		toSerialize["computedPayload"] = o.ComputedPayload.Get()
	}
	if o.ComputedUrl.IsSet() {
		toSerialize["computedUrl"] = o.ComputedUrl.Get()
	}
	if o.ResponseBody.IsSet() {
		toSerialize["responseBody"] = o.ResponseBody.Get()
	}
	if o.ResponseCode.IsSet() {
		toSerialize["responseCode"] = o.ResponseCode.Get()
	}
	if o.ResponseHeaders != nil {
		toSerialize["responseHeaders"] = o.ResponseHeaders
	}
	toSerialize["type"] = o.Type
	toSerialize["dictionary"] = o.Dictionary
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RESTDataSourceResultResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"dictionary",
		"name",
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

	varRESTDataSourceResultResponse := _RESTDataSourceResultResponse{}

	err = json.Unmarshal(data, &varRESTDataSourceResultResponse)

	if err != nil {
		return err
	}

	*o = RESTDataSourceResultResponse(varRESTDataSourceResultResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "computedHeaders")
		delete(additionalProperties, "computedPayload")
		delete(additionalProperties, "computedUrl")
		delete(additionalProperties, "responseBody")
		delete(additionalProperties, "responseCode")
		delete(additionalProperties, "responseHeaders")
		delete(additionalProperties, "type")
		delete(additionalProperties, "dictionary")
		delete(additionalProperties, "error")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRESTDataSourceResultResponse struct {
	value *RESTDataSourceResultResponse
	isSet bool
}

func (v NullableRESTDataSourceResultResponse) Get() *RESTDataSourceResultResponse {
	return v.value
}

func (v *NullableRESTDataSourceResultResponse) Set(val *RESTDataSourceResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRESTDataSourceResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRESTDataSourceResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRESTDataSourceResultResponse(val *RESTDataSourceResultResponse) *NullableRESTDataSourceResultResponse {
	return &NullableRESTDataSourceResultResponse{value: val, isSet: true}
}

func (v NullableRESTDataSourceResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRESTDataSourceResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
