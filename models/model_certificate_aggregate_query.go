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

// checks if the CertificateAggregateQuery type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CertificateAggregateQuery{}

// CertificateAggregateQuery struct for CertificateAggregateQuery
type CertificateAggregateQuery struct {
	// The field that the aggregation will take place on
	GroupBy []string `json:"groupBy,omitempty"`
	// A condition to apply to the result. Only the aggregates results with more than 5 certificates in them can be kept for example
	Having NullableHaving `json:"having,omitempty"`
	// In case of an aggregate sending a lot of different results, how many must be sent back
	Limit utils.NullableInt64 `json:"limit,omitempty"`
	// The HCQL query to use for the search, represents the way to filter certificates. If not specified, it will filter nothing
	Query     utils.NullableString `json:"query,omitempty"`
	SortOrder utils.NullableString `json:"sortOrder,omitempty"`
	// If set to `true`, the total count of certificates matching the HCQL query will be returned
	WithCount            utils.NullableBool `json:"withCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateAggregateQuery CertificateAggregateQuery

// NewCertificateAggregateQuery instantiates a new CertificateAggregateQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAggregateQuery() *CertificateAggregateQuery {
	this := CertificateAggregateQuery{}
	return &this
}

// NewCertificateAggregateQueryWithDefaults instantiates a new CertificateAggregateQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAggregateQueryWithDefaults() *CertificateAggregateQuery {
	this := CertificateAggregateQuery{}
	return &this
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAggregateQuery) GetGroupBy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAggregateQuery) GetGroupByOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.GroupBy) {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *CertificateAggregateQuery) HasGroupBy() bool {
	if o != nil && !utils.IsNil(o.GroupBy) {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *CertificateAggregateQuery) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetHaving returns the Having field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAggregateQuery) GetHaving() Having {
	if o == nil || utils.IsNil(o.Having.Get()) {
		var ret Having
		return ret
	}
	return *o.Having.Get()
}

// GetHavingOk returns a tuple with the Having field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAggregateQuery) GetHavingOk() (*Having, bool) {
	if o == nil {
		return nil, false
	}
	return o.Having.Get(), o.Having.IsSet()
}

// HasHaving returns a boolean if a field has been set.
func (o *CertificateAggregateQuery) HasHaving() bool {
	if o != nil && o.Having.IsSet() {
		return true
	}

	return false
}

// SetHaving gets a reference to the given NullableHaving and assigns it to the Having field.
func (o *CertificateAggregateQuery) SetHaving(v Having) {
	o.Having.Set(&v)
}

// SetHavingNil sets the value for Having to be an explicit nil
func (o *CertificateAggregateQuery) SetHavingNil() {
	o.Having.Set(nil)
}

// UnsetHaving ensures that no value is present for Having, not even an explicit nil
func (o *CertificateAggregateQuery) UnsetHaving() {
	o.Having.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAggregateQuery) GetLimit() int64 {
	if o == nil || utils.IsNil(o.Limit.Get()) {
		var ret int64
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAggregateQuery) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *CertificateAggregateQuery) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt64 and assigns it to the Limit field.
func (o *CertificateAggregateQuery) SetLimit(v int64) {
	o.Limit.Set(&v)
}

// SetLimitNil sets the value for Limit to be an explicit nil
func (o *CertificateAggregateQuery) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *CertificateAggregateQuery) UnsetLimit() {
	o.Limit.Unset()
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAggregateQuery) GetQuery() string {
	if o == nil || utils.IsNil(o.Query.Get()) {
		var ret string
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAggregateQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *CertificateAggregateQuery) HasQuery() bool {
	if o != nil && o.Query.IsSet() {
		return true
	}

	return false
}

// SetQuery gets a reference to the given NullableString and assigns it to the Query field.
func (o *CertificateAggregateQuery) SetQuery(v string) {
	o.Query.Set(&v)
}

// SetQueryNil sets the value for Query to be an explicit nil
func (o *CertificateAggregateQuery) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil
func (o *CertificateAggregateQuery) UnsetQuery() {
	o.Query.Unset()
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAggregateQuery) GetSortOrder() string {
	if o == nil || utils.IsNil(o.SortOrder.Get()) {
		var ret string
		return ret
	}
	return *o.SortOrder.Get()
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAggregateQuery) GetSortOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortOrder.Get(), o.SortOrder.IsSet()
}

// HasSortOrder returns a boolean if a field has been set.
func (o *CertificateAggregateQuery) HasSortOrder() bool {
	if o != nil && o.SortOrder.IsSet() {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given NullableString and assigns it to the SortOrder field.
func (o *CertificateAggregateQuery) SetSortOrder(v string) {
	o.SortOrder.Set(&v)
}

// SetSortOrderNil sets the value for SortOrder to be an explicit nil
func (o *CertificateAggregateQuery) SetSortOrderNil() {
	o.SortOrder.Set(nil)
}

// UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
func (o *CertificateAggregateQuery) UnsetSortOrder() {
	o.SortOrder.Unset()
}

// GetWithCount returns the WithCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAggregateQuery) GetWithCount() bool {
	if o == nil || utils.IsNil(o.WithCount.Get()) {
		var ret bool
		return ret
	}
	return *o.WithCount.Get()
}

// GetWithCountOk returns a tuple with the WithCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAggregateQuery) GetWithCountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithCount.Get(), o.WithCount.IsSet()
}

// HasWithCount returns a boolean if a field has been set.
func (o *CertificateAggregateQuery) HasWithCount() bool {
	if o != nil && o.WithCount.IsSet() {
		return true
	}

	return false
}

// SetWithCount gets a reference to the given NullableBool and assigns it to the WithCount field.
func (o *CertificateAggregateQuery) SetWithCount(v bool) {
	o.WithCount.Set(&v)
}

// SetWithCountNil sets the value for WithCount to be an explicit nil
func (o *CertificateAggregateQuery) SetWithCountNil() {
	o.WithCount.Set(nil)
}

// UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil
func (o *CertificateAggregateQuery) UnsetWithCount() {
	o.WithCount.Unset()
}

func (o CertificateAggregateQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateAggregateQuery) ToMap() (map[string]interface{}, error) {
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

func (o *CertificateAggregateQuery) UnmarshalJSON(data []byte) (err error) {
	varCertificateAggregateQuery := _CertificateAggregateQuery{}

	err = json.Unmarshal(data, &varCertificateAggregateQuery)

	if err != nil {
		return err
	}

	*o = CertificateAggregateQuery(varCertificateAggregateQuery)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "groupBy")
		delete(additionalProperties, "having")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "query")
		delete(additionalProperties, "sortOrder")
		delete(additionalProperties, "withCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateAggregateQuery struct {
	value *CertificateAggregateQuery
	isSet bool
}

func (v NullableCertificateAggregateQuery) Get() *CertificateAggregateQuery {
	return v.value
}

func (v *NullableCertificateAggregateQuery) Set(val *CertificateAggregateQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAggregateQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAggregateQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAggregateQuery(val *CertificateAggregateQuery) *NullableCertificateAggregateQuery {
	return &NullableCertificateAggregateQuery{value: val, isSet: true}
}

func (v NullableCertificateAggregateQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAggregateQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
