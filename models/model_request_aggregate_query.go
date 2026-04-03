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

// checks if the RequestAggregateQuery type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RequestAggregateQuery{}

// RequestAggregateQuery struct for RequestAggregateQuery
type RequestAggregateQuery struct {
	// The field that the aggregation will take place on
	GroupBy []string `json:"groupBy,omitempty"`
	// A condition to apply to the result. Only the aggregates results with more than 5 requests in them can be kept for example
	Having NullableHaving `json:"having,omitempty"`
	// In case of an aggregate sending a lot of different results, how many must be sent back
	Limit utils.NullableInt64 `json:"limit,omitempty"`
	// The HRQL query to use for the search, represents the way to filter requests. If not specified, it will filter nothing
	Query utils.NullableString `json:"query,omitempty"`
	// The scope of the aggregate. `manage` only aggregates among requests that the currently logged in user has the rights to manage. `search` aggregates among all visible requests to the logged in user. `self` aggregates among requests that the currently logged in user or its team has issued
	Scope     utils.NullableString `json:"scope,omitempty"`
	SortOrder utils.NullableString `json:"sortOrder,omitempty"`
	// If set to `true`, the total count of requests matching the HRQL query will be returned
	WithCount            utils.NullableBool `json:"withCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestAggregateQuery RequestAggregateQuery

// NewRequestAggregateQuery instantiates a new RequestAggregateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestAggregateQuery() *RequestAggregateQuery {
	this := RequestAggregateQuery{}
	return &this
}

// NewRequestAggregateQueryWithDefaults instantiates a new RequestAggregateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestAggregateQueryWithDefaults() *RequestAggregateQuery {
	this := RequestAggregateQuery{}
	return &this
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestAggregateQuery) GetGroupBy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestAggregateQuery) GetGroupByOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.GroupBy) {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *RequestAggregateQuery) HasGroupBy() bool {
	if o != nil && !utils.IsNil(o.GroupBy) {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *RequestAggregateQuery) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetHaving returns the Having field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestAggregateQuery) GetHaving() Having {
	if o == nil || utils.IsNil(o.Having.Get()) {
		var ret Having
		return ret
	}
	return *o.Having.Get()
}

// GetHavingOk returns a tuple with the Having field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestAggregateQuery) GetHavingOk() (*Having, bool) {
	if o == nil {
		return nil, false
	}
	return o.Having.Get(), o.Having.IsSet()
}

// HasHaving returns a boolean if a field has been set.
func (o *RequestAggregateQuery) HasHaving() bool {
	if o != nil && o.Having.IsSet() {
		return true
	}

	return false
}

// SetHaving gets a reference to the given NullableHaving and assigns it to the Having field.
func (o *RequestAggregateQuery) SetHaving(v Having) {
	o.Having.Set(&v)
}

// SetHavingNil sets the value for Having to be an explicit nil
func (o *RequestAggregateQuery) SetHavingNil() {
	o.Having.Set(nil)
}

// UnsetHaving ensures that no value is present for Having, not even an explicit nil
func (o *RequestAggregateQuery) UnsetHaving() {
	o.Having.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestAggregateQuery) GetLimit() int64 {
	if o == nil || utils.IsNil(o.Limit.Get()) {
		var ret int64
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestAggregateQuery) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *RequestAggregateQuery) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt64 and assigns it to the Limit field.
func (o *RequestAggregateQuery) SetLimit(v int64) {
	o.Limit.Set(&v)
}

// SetLimitNil sets the value for Limit to be an explicit nil
func (o *RequestAggregateQuery) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *RequestAggregateQuery) UnsetLimit() {
	o.Limit.Unset()
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestAggregateQuery) GetQuery() string {
	if o == nil || utils.IsNil(o.Query.Get()) {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestAggregateQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *RequestAggregateQuery) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *RequestAggregateQuery) SetQuery(v string) {
	o.Query.Set(&v)
}

// SetQueryNil sets the value for Query to be an explicit nil
func (o *RequestAggregateQuery) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *RequestAggregateQuery) UnsetQuery() {
	o.Query.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestAggregateQuery) GetScope() string {
	if o == nil || utils.IsNil(o.Scope.Get()) {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestAggregateQuery) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *RequestAggregateQuery) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *RequestAggregateQuery) SetScope(v string) {
	o.Scope.Set(&v)
}

// SetScopeNil sets the value for Scope to be an explicit nil
func (o *RequestAggregateQuery) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *RequestAggregateQuery) UnsetScope() {
	o.Scope.Unset()
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestAggregateQuery) GetSortOrder() string {
	if o == nil || utils.IsNil(o.SortOrder.Get()) {
		var ret string
		return ret
	}
	return *o.SortOrder.Get()
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestAggregateQuery) GetSortOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortOrder.Get(), o.SortOrder.IsSet()
}

// HasSortOrder returns a boolean if a field has been set.
func (o *RequestAggregateQuery) HasSortOrder() bool {
	if o != nil && o.SortOrder.IsSet() {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given NullableString and assigns it to the SortOrder field.
func (o *RequestAggregateQuery) SetSortOrder(v string) {
	o.SortOrder.Set(&v)
}

// SetSortOrderNil sets the value for SortOrder to be an explicit nil
func (o *RequestAggregateQuery) SetSortOrderNil() {
	o.SortOrder.Set(nil)
}

// UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
func (o *RequestAggregateQuery) UnsetSortOrder() {
	o.SortOrder.Unset()
}

// GetWithCount returns the WithCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestAggregateQuery) GetWithCount() bool {
	if o == nil || utils.IsNil(o.WithCount.Get()) {
		var ret bool
		return ret
	}
	return *o.WithCount.Get()
}

// GetWithCountOk returns a tuple with the WithCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestAggregateQuery) GetWithCountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithCount.Get(), o.WithCount.IsSet()
}

// HasWithCount returns a boolean if a field has been set.
func (o *RequestAggregateQuery) HasWithCount() bool {
	if o != nil && o.WithCount.IsSet() {
		return true
	}

	return false
}

// SetWithCount gets a reference to the given NullableBool and assigns it to the WithCount field.
func (o *RequestAggregateQuery) SetWithCount(v bool) {
	o.WithCount.Set(&v)
}

// SetWithCountNil sets the value for WithCount to be an explicit nil
func (o *RequestAggregateQuery) SetWithCountNil() {
	o.WithCount.Set(nil)
}

// UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil
func (o *RequestAggregateQuery) UnsetWithCount() {
	o.WithCount.Unset()
}

func (o RequestAggregateQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestAggregateQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupBy != nil {
		toSerialize["groupBy"] = o.GroupBy
	}
	if o.Having.IsSet() {
		toSerialize["having"] = o.Having.Get()
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	if o.SortOrder.IsSet() {
		toSerialize["sortOrder"] = o.SortOrder.Get()
	}
	if o.WithCount.IsSet() {
		toSerialize["withCount"] = o.WithCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestAggregateQuery) UnmarshalJSON(data []byte) (err error) {
	varRequestAggregateQuery := _RequestAggregateQuery{}

	err = json.Unmarshal(data, &varRequestAggregateQuery)

	if err != nil {
		return err
	}

	*o = RequestAggregateQuery(varRequestAggregateQuery)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "groupBy")
		delete(additionalProperties, "having")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "query")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "sortOrder")
		delete(additionalProperties, "withCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestAggregateQuery struct {
	value *RequestAggregateQuery
	isSet bool
}

func (v NullableRequestAggregateQuery) Get() *RequestAggregateQuery {
	return v.value
}

func (v *NullableRequestAggregateQuery) Set(val *RequestAggregateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestAggregateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestAggregateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestAggregateQuery(val *RequestAggregateQuery) *NullableRequestAggregateQuery {
	return &NullableRequestAggregateQuery{value: val, isSet: true}
}

func (v NullableRequestAggregateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestAggregateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
