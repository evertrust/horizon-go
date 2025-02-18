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

// checks if the EventSearchResultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventSearchResultsResponse{}

// EventSearchResultsResponse struct for EventSearchResultsResponse
type EventSearchResultsResponse struct {
	Results []EventSearchResult `json:"results"`
	PageIndex int64 `json:"pageIndex"`
	PageSize int64 `json:"pageSize"`
	Count NullableInt64 `json:"count,omitempty"`
	HasMore bool `json:"hasMore"`
	AdditionalProperties map[string]interface{}
}

type _EventSearchResultsResponse EventSearchResultsResponse

// NewEventSearchResultsResponse instantiates a new EventSearchResultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSearchResultsResponse(results []EventSearchResult, pageIndex int64, pageSize int64, hasMore bool) *EventSearchResultsResponse {
	this := EventSearchResultsResponse{}
	this.Results = results
	this.PageIndex = pageIndex
	this.PageSize = pageSize
	this.HasMore = hasMore
	return &this
}

// NewEventSearchResultsResponseWithDefaults instantiates a new EventSearchResultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSearchResultsResponseWithDefaults() *EventSearchResultsResponse {
	this := EventSearchResultsResponse{}
	return &this
}

// GetResults returns the Results field value
func (o *EventSearchResultsResponse) GetResults() []EventSearchResult {
	if o == nil {
		var ret []EventSearchResult
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *EventSearchResultsResponse) GetResultsOk() ([]EventSearchResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *EventSearchResultsResponse) SetResults(v []EventSearchResult) {
	o.Results = v
}

// GetPageIndex returns the PageIndex field value
func (o *EventSearchResultsResponse) GetPageIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value
// and a boolean to check if the value has been set.
func (o *EventSearchResultsResponse) GetPageIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageIndex, true
}

// SetPageIndex sets field value
func (o *EventSearchResultsResponse) SetPageIndex(v int64) {
	o.PageIndex = v
}

// GetPageSize returns the PageSize field value
func (o *EventSearchResultsResponse) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *EventSearchResultsResponse) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *EventSearchResultsResponse) SetPageSize(v int64) {
	o.PageSize = v
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchResultsResponse) GetCount() int64 {
	if o == nil || IsNil(o.Count.Get()) {
		var ret int64
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchResultsResponse) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *EventSearchResultsResponse) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt64 and assigns it to the Count field.
func (o *EventSearchResultsResponse) SetCount(v int64) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *EventSearchResultsResponse) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *EventSearchResultsResponse) UnsetCount() {
	o.Count.Unset()
}

// GetHasMore returns the HasMore field value
func (o *EventSearchResultsResponse) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *EventSearchResultsResponse) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *EventSearchResultsResponse) SetHasMore(v bool) {
	o.HasMore = v
}

func (o EventSearchResultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSearchResultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	toSerialize["pageIndex"] = o.PageIndex
	toSerialize["pageSize"] = o.PageSize
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	toSerialize["hasMore"] = o.HasMore

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventSearchResultsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
		"pageIndex",
		"pageSize",
		"hasMore",
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

	varEventSearchResultsResponse := _EventSearchResultsResponse{}

	err = json.Unmarshal(data, &varEventSearchResultsResponse)

	if err != nil {
		return err
	}

	*o = EventSearchResultsResponse(varEventSearchResultsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		delete(additionalProperties, "pageIndex")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "count")
		delete(additionalProperties, "hasMore")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventSearchResultsResponse struct {
	value *EventSearchResultsResponse
	isSet bool
}

func (v NullableEventSearchResultsResponse) Get() *EventSearchResultsResponse {
	return v.value
}

func (v *NullableEventSearchResultsResponse) Set(val *EventSearchResultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSearchResultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSearchResultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSearchResultsResponse(val *EventSearchResultsResponse) *NullableEventSearchResultsResponse {
	return &NullableEventSearchResultsResponse{value: val, isSet: true}
}

func (v NullableEventSearchResultsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSearchResultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


