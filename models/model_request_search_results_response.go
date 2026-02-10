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

// checks if the RequestSearchResultsResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RequestSearchResultsResponse{}

// RequestSearchResultsResponse struct for RequestSearchResultsResponse
type RequestSearchResultsResponse struct {
	// The total count of requests matching the HRQL query
	Count *int64 `json:"count,omitempty"`
	// Indicates whether the response represents the last page of results (if set to `false`) or not (if set to `true`)
	HasMore bool `json:"hasMore"`
	// The index of the results page
	PageIndex int64 `json:"pageIndex"`
	// The maximum number of items on this page
	PageSize int64 `json:"pageSize"`
	// The list of requests matching the HRQL query
	Results              []RequestSearchResult `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _RequestSearchResultsResponse RequestSearchResultsResponse

// NewRequestSearchResultsResponse instantiates a new RequestSearchResultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSearchResultsResponse(hasMore bool, pageIndex int64, pageSize int64, results []RequestSearchResult) *RequestSearchResultsResponse {
	this := RequestSearchResultsResponse{}
	this.HasMore = hasMore
	this.PageIndex = pageIndex
	this.PageSize = pageSize
	this.Results = results
	return &this
}

// NewRequestSearchResultsResponseWithDefaults instantiates a new RequestSearchResultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSearchResultsResponseWithDefaults() *RequestSearchResultsResponse {
	this := RequestSearchResultsResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RequestSearchResultsResponse) GetCount() int64 {
	if o == nil || utils.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSearchResultsResponse) GetCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RequestSearchResultsResponse) HasCount() bool {
	if o != nil && !utils.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *RequestSearchResultsResponse) SetCount(v int64) {
	o.Count = &v
}

// GetHasMore returns the HasMore field value
func (o *RequestSearchResultsResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *RequestSearchResultsResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *RequestSearchResultsResponse) SetHasMore(v bool) {
	o.HasMore = v
}

// GetPageIndex returns the PageIndex field value
func (o *RequestSearchResultsResponse) GetPageIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value
// and a boolean to check if the value has been set.
func (o *RequestSearchResultsResponse) GetPageIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageIndex, true
}

// SetPageIndex sets field value
func (o *RequestSearchResultsResponse) SetPageIndex(v int64) {
	o.PageIndex = v
}

// GetPageSize returns the PageSize field value
func (o *RequestSearchResultsResponse) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *RequestSearchResultsResponse) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *RequestSearchResultsResponse) SetPageSize(v int64) {
	o.PageSize = v
}

// GetResults returns the Results field value
func (o *RequestSearchResultsResponse) GetResults() []RequestSearchResult {
	if o == nil {
		var ret []RequestSearchResult
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *RequestSearchResultsResponse) GetResultsOk() ([]RequestSearchResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *RequestSearchResultsResponse) SetResults(v []RequestSearchResult) {
	o.Results = v
}

func (o RequestSearchResultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestSearchResultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	toSerialize["hasMore"] = o.HasMore
	toSerialize["pageIndex"] = o.PageIndex
	toSerialize["pageSize"] = o.PageSize
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestSearchResultsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hasMore",
		"pageIndex",
		"pageSize",
		"results",
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

	varRequestSearchResultsResponse := _RequestSearchResultsResponse{}

	err = json.Unmarshal(data, &varRequestSearchResultsResponse)

	if err != nil {
		return err
	}

	*o = RequestSearchResultsResponse(varRequestSearchResultsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "hasMore")
		delete(additionalProperties, "pageIndex")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestSearchResultsResponse struct {
	value *RequestSearchResultsResponse
	isSet bool
}

func (v NullableRequestSearchResultsResponse) Get() *RequestSearchResultsResponse {
	return v.value
}

func (v *NullableRequestSearchResultsResponse) Set(val *RequestSearchResultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSearchResultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSearchResultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSearchResultsResponse(val *RequestSearchResultsResponse) *NullableRequestSearchResultsResponse {
	return &NullableRequestSearchResultsResponse{value: val, isSet: true}
}

func (v NullableRequestSearchResultsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSearchResultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
