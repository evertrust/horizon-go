/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the DiscoveryEventSearchQuery type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DiscoveryEventSearchQuery{}

// DiscoveryEventSearchQuery struct for DiscoveryEventSearchQuery
type DiscoveryEventSearchQuery struct {
	Query                utils.NullableString `json:"query,omitempty"`
	SortedBy             []SortElement        `json:"sortedBy,omitempty"`
	PageIndex            utils.NullableInt64  `json:"pageIndex,omitempty"`
	PageSize             utils.NullableInt64  `json:"pageSize,omitempty"`
	WithCount            utils.NullableBool   `json:"withCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveryEventSearchQuery DiscoveryEventSearchQuery

// NewDiscoveryEventSearchQuery instantiates a new DiscoveryEventSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryEventSearchQuery() *DiscoveryEventSearchQuery {
	this := DiscoveryEventSearchQuery{}
	return &this
}

// NewDiscoveryEventSearchQueryWithDefaults instantiates a new DiscoveryEventSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryEventSearchQueryWithDefaults() *DiscoveryEventSearchQuery {
	this := DiscoveryEventSearchQuery{}
	return &this
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventSearchQuery) GetQuery() string {
	if o == nil || utils.IsNil(o.Query.Get()) {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventSearchQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *DiscoveryEventSearchQuery) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *DiscoveryEventSearchQuery) SetQuery(v string) {
	o.Query.Set(&v)
}

// SetQueryNil sets the value for Query to be an explicit nil
func (o *DiscoveryEventSearchQuery) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *DiscoveryEventSearchQuery) UnsetQuery() {
	o.Query.Unset()
}

// GetSortedBy returns the SortedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventSearchQuery) GetSortedBy() []SortElement {
	if o == nil {
		var ret []SortElement
		return ret
	}
	return o.SortedBy
}

// GetSortedByOk returns a tuple with the SortedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventSearchQuery) GetSortedByOk() ([]SortElement, bool) {
	if o == nil || utils.IsNil(o.SortedBy) {
		return nil, false
	}
	return o.SortedBy, true
}

// HasSortedBy returns a boolean if a field has been set.
func (o *DiscoveryEventSearchQuery) HasSortedBy() bool {
	if o != nil && !utils.IsNil(o.SortedBy) {
		return true
	}

	return false
}

// SetSortedBy gets a reference to the given []SortElement and assigns it to the SortedBy field.
func (o *DiscoveryEventSearchQuery) SetSortedBy(v []SortElement) {
	o.SortedBy = v
}

// GetPageIndex returns the PageIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventSearchQuery) GetPageIndex() int64 {
	if o == nil || utils.IsNil(o.PageIndex.Get()) {
		var ret int64
		return ret
	}
	return *o.PageIndex.Get()
}

// GetPageIndexOk returns a tuple with the PageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventSearchQuery) GetPageIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageIndex.Get(), o.PageIndex.IsSet()
}

// HasPageIndex returns a boolean if a field has been set.
func (o *DiscoveryEventSearchQuery) HasPageIndex() bool {
	if o != nil && o.PageIndex.IsSet() {
		return true
	}

	return false
}

// SetPageIndex gets a reference to the given NullableInt64 and assigns it to the PageIndex field.
func (o *DiscoveryEventSearchQuery) SetPageIndex(v int64) {
	o.PageIndex.Set(&v)
}

// SetPageIndexNil sets the value for PageIndex to be an explicit nil
func (o *DiscoveryEventSearchQuery) SetPageIndexNil() {
	o.PageIndex.Set(nil)
}

// UnsetPageIndex ensures that no value is present for PageIndex, not even an explicit nil
func (o *DiscoveryEventSearchQuery) UnsetPageIndex() {
	o.PageIndex.Unset()
}

// GetPageSize returns the PageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventSearchQuery) GetPageSize() int64 {
	if o == nil || utils.IsNil(o.PageSize.Get()) {
		var ret int64
		return ret
	}
	return *o.PageSize.Get()
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventSearchQuery) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageSize.Get(), o.PageSize.IsSet()
}

// HasPageSize returns a boolean if a field has been set.
func (o *DiscoveryEventSearchQuery) HasPageSize() bool {
	if o != nil && o.PageSize.IsSet() {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given NullableInt64 and assigns it to the PageSize field.
func (o *DiscoveryEventSearchQuery) SetPageSize(v int64) {
	o.PageSize.Set(&v)
}

// SetPageSizeNil sets the value for PageSize to be an explicit nil
func (o *DiscoveryEventSearchQuery) SetPageSizeNil() {
	o.PageSize.Set(nil)
}

// UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
func (o *DiscoveryEventSearchQuery) UnsetPageSize() {
	o.PageSize.Unset()
}

// GetWithCount returns the WithCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventSearchQuery) GetWithCount() bool {
	if o == nil || utils.IsNil(o.WithCount.Get()) {
		var ret bool
		return ret
	}
	return *o.WithCount.Get()
}

// GetWithCountOk returns a tuple with the WithCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventSearchQuery) GetWithCountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithCount.Get(), o.WithCount.IsSet()
}

// HasWithCount returns a boolean if a field has been set.
func (o *DiscoveryEventSearchQuery) HasWithCount() bool {
	if o != nil && o.WithCount.IsSet() {
		return true
	}

	return false
}

// SetWithCount gets a reference to the given NullableBool and assigns it to the WithCount field.
func (o *DiscoveryEventSearchQuery) SetWithCount(v bool) {
	o.WithCount.Set(&v)
}

// SetWithCountNil sets the value for WithCount to be an explicit nil
func (o *DiscoveryEventSearchQuery) SetWithCountNil() {
	o.WithCount.Set(nil)
}

// UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil
func (o *DiscoveryEventSearchQuery) UnsetWithCount() {
	o.WithCount.Unset()
}

func (o DiscoveryEventSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryEventSearchQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.SortedBy != nil {
		toSerialize["sortedBy"] = o.SortedBy
	}
	if o.PageIndex.IsSet() {
		toSerialize["pageIndex"] = o.PageIndex.Get()
	}
	if o.PageSize.IsSet() {
		toSerialize["pageSize"] = o.PageSize.Get()
	}
	if o.WithCount.IsSet() {
		toSerialize["withCount"] = o.WithCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoveryEventSearchQuery) UnmarshalJSON(data []byte) (err error) {
	varDiscoveryEventSearchQuery := _DiscoveryEventSearchQuery{}

	err = json.Unmarshal(data, &varDiscoveryEventSearchQuery)

	if err != nil {
		return err
	}

	*o = DiscoveryEventSearchQuery(varDiscoveryEventSearchQuery)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "query")
		delete(additionalProperties, "sortedBy")
		delete(additionalProperties, "pageIndex")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "withCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveryEventSearchQuery struct {
	value *DiscoveryEventSearchQuery
	isSet bool
}

func (v NullableDiscoveryEventSearchQuery) Get() *DiscoveryEventSearchQuery {
	return v.value
}

func (v *NullableDiscoveryEventSearchQuery) Set(val *DiscoveryEventSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryEventSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryEventSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryEventSearchQuery(val *DiscoveryEventSearchQuery) *NullableDiscoveryEventSearchQuery {
	return &NullableDiscoveryEventSearchQuery{value: val, isSet: true}
}

func (v NullableDiscoveryEventSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryEventSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
