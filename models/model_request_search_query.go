/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the RequestSearchQuery type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RequestSearchQuery{}

// RequestSearchQuery struct for RequestSearchQuery
type RequestSearchQuery struct {
	// The fields to be returned by the search. If this parameter is not specified, everything is returned by default. If this parameter is equal to an empty array, only the `_id` field is returned
	Fields []string `json:"fields,omitempty"`
	// The index of the page to retrieve
	PageIndex utils.NullableInt64 `json:"pageIndex,omitempty"`
	// The maximum number of items to retrieve for one page
	PageSize utils.NullableInt64 `json:"pageSize,omitempty"`
	// The HRQL query to use for the search, represents the way to filter requests. Filters nothing if not specified
	Query utils.NullableString `json:"query,omitempty"`
	// The scope of the search. `manage` only searches among requests that the currently logged in user has the rights to manage. `search` searches among all visible requests to the logged in user. `self` searches among requests that the currently logged in user or its team has issued
	Scope utils.NullableString `json:"scope,omitempty"`
	// The way to sort the search results
	SortedBy []SortElement `json:"sortedBy,omitempty"`
	// Whether to return the total count of requests matching the HRQL query
	WithCount            utils.NullableBool `json:"withCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestSearchQuery RequestSearchQuery

// NewRequestSearchQuery instantiates a new RequestSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSearchQuery() *RequestSearchQuery {
	this := RequestSearchQuery{}
	var pageIndex int64 = 1
	this.PageIndex = *utils.NewNullableInt64(&pageIndex)
	var pageSize int64 = 50
	this.PageSize = *utils.NewNullableInt64(&pageSize)
	var withCount bool = false
	this.WithCount = *utils.NewNullableBool(&withCount)
	return &this
}

// NewRequestSearchQueryWithDefaults instantiates a new RequestSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSearchQueryWithDefaults() *RequestSearchQuery {
	this := RequestSearchQuery{}
	var pageIndex int64 = 1
	this.PageIndex = *utils.NewNullableInt64(&pageIndex)
	var pageSize int64 = 50
	this.PageSize = *utils.NewNullableInt64(&pageSize)
	var withCount bool = false
	this.WithCount = *utils.NewNullableBool(&withCount)
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSearchQuery) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSearchQuery) GetFieldsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *RequestSearchQuery) HasFields() bool {
	if o != nil && !utils.IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *RequestSearchQuery) SetFields(v []string) {
	o.Fields = v
}

// GetPageIndex returns the PageIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSearchQuery) GetPageIndex() int64 {
	if o == nil || utils.IsNil(o.PageIndex.Get()) {
		var ret int64
		return ret
	}
	return *o.PageIndex.Get()
}

// GetPageIndexOk returns a tuple with the PageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSearchQuery) GetPageIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageIndex.Get(), o.PageIndex.IsSet()
}

// HasPageIndex returns a boolean if a field has been set.
func (o *RequestSearchQuery) HasPageIndex() bool {
	if o != nil && o.PageIndex.IsSet() {
		return true
	}

	return false
}

// SetPageIndex gets a reference to the given NullableInt64 and assigns it to the PageIndex field.
func (o *RequestSearchQuery) SetPageIndex(v int64) {
	o.PageIndex.Set(&v)
}

// SetPageIndexNil sets the value for PageIndex to be an explicit nil
func (o *RequestSearchQuery) SetPageIndexNil() {
	o.PageIndex.Set(nil)
}

// UnsetPageIndex ensures that no value is present for PageIndex, not even an explicit nil
func (o *RequestSearchQuery) UnsetPageIndex() {
	o.PageIndex.Unset()
}

// GetPageSize returns the PageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSearchQuery) GetPageSize() int64 {
	if o == nil || utils.IsNil(o.PageSize.Get()) {
		var ret int64
		return ret
	}
	return *o.PageSize.Get()
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSearchQuery) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageSize.Get(), o.PageSize.IsSet()
}

// HasPageSize returns a boolean if a field has been set.
func (o *RequestSearchQuery) HasPageSize() bool {
	if o != nil && o.PageSize.IsSet() {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given NullableInt64 and assigns it to the PageSize field.
func (o *RequestSearchQuery) SetPageSize(v int64) {
	o.PageSize.Set(&v)
}

// SetPageSizeNil sets the value for PageSize to be an explicit nil
func (o *RequestSearchQuery) SetPageSizeNil() {
	o.PageSize.Set(nil)
}

// UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
func (o *RequestSearchQuery) UnsetPageSize() {
	o.PageSize.Unset()
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSearchQuery) GetQuery() string {
	if o == nil || utils.IsNil(o.Query.Get()) {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSearchQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *RequestSearchQuery) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *RequestSearchQuery) SetQuery(v string) {
	o.Query.Set(&v)
}

// SetQueryNil sets the value for Query to be an explicit nil
func (o *RequestSearchQuery) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *RequestSearchQuery) UnsetQuery() {
	o.Query.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSearchQuery) GetScope() string {
	if o == nil || utils.IsNil(o.Scope.Get()) {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSearchQuery) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *RequestSearchQuery) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *RequestSearchQuery) SetScope(v string) {
	o.Scope.Set(&v)
}

// SetScopeNil sets the value for Scope to be an explicit nil
func (o *RequestSearchQuery) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *RequestSearchQuery) UnsetScope() {
	o.Scope.Unset()
}

// GetSortedBy returns the SortedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSearchQuery) GetSortedBy() []SortElement {
	if o == nil {
		var ret []SortElement
		return ret
	}
	return o.SortedBy
}

// GetSortedByOk returns a tuple with the SortedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSearchQuery) GetSortedByOk() ([]SortElement, bool) {
	if o == nil || utils.IsNil(o.SortedBy) {
		return nil, false
	}
	return o.SortedBy, true
}

// HasSortedBy returns a boolean if a field has been set.
func (o *RequestSearchQuery) HasSortedBy() bool {
	if o != nil && !utils.IsNil(o.SortedBy) {
		return true
	}

	return false
}

// SetSortedBy gets a reference to the given []SortElement and assigns it to the SortedBy field.
func (o *RequestSearchQuery) SetSortedBy(v []SortElement) {
	o.SortedBy = v
}

// GetWithCount returns the WithCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSearchQuery) GetWithCount() bool {
	if o == nil || utils.IsNil(o.WithCount.Get()) {
		var ret bool
		return ret
	}
	return *o.WithCount.Get()
}

// GetWithCountOk returns a tuple with the WithCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSearchQuery) GetWithCountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithCount.Get(), o.WithCount.IsSet()
}

// HasWithCount returns a boolean if a field has been set.
func (o *RequestSearchQuery) HasWithCount() bool {
	if o != nil && o.WithCount.IsSet() {
		return true
	}

	return false
}

// SetWithCount gets a reference to the given NullableBool and assigns it to the WithCount field.
func (o *RequestSearchQuery) SetWithCount(v bool) {
	o.WithCount.Set(&v)
}

// SetWithCountNil sets the value for WithCount to be an explicit nil
func (o *RequestSearchQuery) SetWithCountNil() {
	o.WithCount.Set(nil)
}

// UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil
func (o *RequestSearchQuery) UnsetWithCount() {
	o.WithCount.Unset()
}

func (o RequestSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestSearchQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.PageIndex.IsSet() {
		toSerialize["pageIndex"] = o.PageIndex.Get()
	}
	if o.PageSize.IsSet() {
		toSerialize["pageSize"] = o.PageSize.Get()
	}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	if o.SortedBy != nil {
		toSerialize["sortedBy"] = o.SortedBy
	}
	if o.WithCount.IsSet() {
		toSerialize["withCount"] = o.WithCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestSearchQuery) UnmarshalJSON(data []byte) (err error) {
	varRequestSearchQuery := _RequestSearchQuery{}

	err = json.Unmarshal(data, &varRequestSearchQuery)

	if err != nil {
		return err
	}

	*o = RequestSearchQuery(varRequestSearchQuery)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fields")
		delete(additionalProperties, "pageIndex")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "query")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "sortedBy")
		delete(additionalProperties, "withCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestSearchQuery struct {
	value *RequestSearchQuery
	isSet bool
}

func (v NullableRequestSearchQuery) Get() *RequestSearchQuery {
	return v.value
}

func (v *NullableRequestSearchQuery) Set(val *RequestSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSearchQuery(val *RequestSearchQuery) *NullableRequestSearchQuery {
	return &NullableRequestSearchQuery{value: val, isSet: true}
}

func (v NullableRequestSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
