/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"encoding/json"
)

// checks if the PrincipalInfoSearchQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalInfoSearchQuery{}

// PrincipalInfoSearchQuery struct for PrincipalInfoSearchQuery
type PrincipalInfoSearchQuery struct {
	// The identifier of the principal
	Identifier NullableString `json:"identifier,omitempty"`
	// The contact e-mail of the principal
	Contact NullableString `json:"contact,omitempty"`
	// The role of the principal
	Role NullableString `json:"role,omitempty"`
	// The team of the principal
	Team NullableString `json:"team,omitempty"`
	// How to sort the results of the search
	SortedBy []SortElement `json:"sortedBy,omitempty"`
	// Which page result to display
	PageIndex NullableInt64 `json:"pageIndex,omitempty"`
	// How many results to display per page
	PageSize NullableInt64 `json:"pageSize,omitempty"`
	// Whether to include the total number of results in the response
	WithCount NullableBool `json:"withCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalInfoSearchQuery PrincipalInfoSearchQuery

// NewPrincipalInfoSearchQuery instantiates a new PrincipalInfoSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalInfoSearchQuery() *PrincipalInfoSearchQuery {
	this := PrincipalInfoSearchQuery{}
	return &this
}

// NewPrincipalInfoSearchQueryWithDefaults instantiates a new PrincipalInfoSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalInfoSearchQueryWithDefaults() *PrincipalInfoSearchQuery {
	this := PrincipalInfoSearchQuery{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoSearchQuery) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier.Get()) {
		var ret string
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoSearchQuery) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *PrincipalInfoSearchQuery) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableString and assigns it to the Identifier field.
func (o *PrincipalInfoSearchQuery) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}
// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *PrincipalInfoSearchQuery) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *PrincipalInfoSearchQuery) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoSearchQuery) GetContact() string {
	if o == nil || IsNil(o.Contact.Get()) {
		var ret string
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoSearchQuery) GetContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *PrincipalInfoSearchQuery) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableString and assigns it to the Contact field.
func (o *PrincipalInfoSearchQuery) SetContact(v string) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *PrincipalInfoSearchQuery) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *PrincipalInfoSearchQuery) UnsetContact() {
	o.Contact.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoSearchQuery) GetRole() string {
	if o == nil || IsNil(o.Role.Get()) {
		var ret string
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoSearchQuery) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PrincipalInfoSearchQuery) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableString and assigns it to the Role field.
func (o *PrincipalInfoSearchQuery) SetRole(v string) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *PrincipalInfoSearchQuery) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PrincipalInfoSearchQuery) UnsetRole() {
	o.Role.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoSearchQuery) GetTeam() string {
	if o == nil || IsNil(o.Team.Get()) {
		var ret string
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoSearchQuery) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *PrincipalInfoSearchQuery) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableString and assigns it to the Team field.
func (o *PrincipalInfoSearchQuery) SetTeam(v string) {
	o.Team.Set(&v)
}
// SetTeamNil sets the value for Team to be an explicit nil
func (o *PrincipalInfoSearchQuery) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *PrincipalInfoSearchQuery) UnsetTeam() {
	o.Team.Unset()
}

// GetSortedBy returns the SortedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoSearchQuery) GetSortedBy() []SortElement {
	if o == nil {
		var ret []SortElement
		return ret
	}
	return o.SortedBy
}

// GetSortedByOk returns a tuple with the SortedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoSearchQuery) GetSortedByOk() ([]SortElement, bool) {
	if o == nil || IsNil(o.SortedBy) {
		return nil, false
	}
	return o.SortedBy, true
}

// HasSortedBy returns a boolean if a field has been set.
func (o *PrincipalInfoSearchQuery) HasSortedBy() bool {
	if o != nil && !IsNil(o.SortedBy) {
		return true
	}

	return false
}

// SetSortedBy gets a reference to the given []SortElement and assigns it to the SortedBy field.
func (o *PrincipalInfoSearchQuery) SetSortedBy(v []SortElement) {
	o.SortedBy = v
}

// GetPageIndex returns the PageIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoSearchQuery) GetPageIndex() int64 {
	if o == nil || IsNil(o.PageIndex.Get()) {
		var ret int64
		return ret
	}
	return *o.PageIndex.Get()
}

// GetPageIndexOk returns a tuple with the PageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoSearchQuery) GetPageIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageIndex.Get(), o.PageIndex.IsSet()
}

// HasPageIndex returns a boolean if a field has been set.
func (o *PrincipalInfoSearchQuery) HasPageIndex() bool {
	if o != nil && o.PageIndex.IsSet() {
		return true
	}

	return false
}

// SetPageIndex gets a reference to the given NullableInt64 and assigns it to the PageIndex field.
func (o *PrincipalInfoSearchQuery) SetPageIndex(v int64) {
	o.PageIndex.Set(&v)
}
// SetPageIndexNil sets the value for PageIndex to be an explicit nil
func (o *PrincipalInfoSearchQuery) SetPageIndexNil() {
	o.PageIndex.Set(nil)
}

// UnsetPageIndex ensures that no value is present for PageIndex, not even an explicit nil
func (o *PrincipalInfoSearchQuery) UnsetPageIndex() {
	o.PageIndex.Unset()
}

// GetPageSize returns the PageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoSearchQuery) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize.Get()) {
		var ret int64
		return ret
	}
	return *o.PageSize.Get()
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoSearchQuery) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PageSize.Get(), o.PageSize.IsSet()
}

// HasPageSize returns a boolean if a field has been set.
func (o *PrincipalInfoSearchQuery) HasPageSize() bool {
	if o != nil && o.PageSize.IsSet() {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given NullableInt64 and assigns it to the PageSize field.
func (o *PrincipalInfoSearchQuery) SetPageSize(v int64) {
	o.PageSize.Set(&v)
}
// SetPageSizeNil sets the value for PageSize to be an explicit nil
func (o *PrincipalInfoSearchQuery) SetPageSizeNil() {
	o.PageSize.Set(nil)
}

// UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
func (o *PrincipalInfoSearchQuery) UnsetPageSize() {
	o.PageSize.Unset()
}

// GetWithCount returns the WithCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoSearchQuery) GetWithCount() bool {
	if o == nil || IsNil(o.WithCount.Get()) {
		var ret bool
		return ret
	}
	return *o.WithCount.Get()
}

// GetWithCountOk returns a tuple with the WithCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoSearchQuery) GetWithCountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WithCount.Get(), o.WithCount.IsSet()
}

// HasWithCount returns a boolean if a field has been set.
func (o *PrincipalInfoSearchQuery) HasWithCount() bool {
	if o != nil && o.WithCount.IsSet() {
		return true
	}

	return false
}

// SetWithCount gets a reference to the given NullableBool and assigns it to the WithCount field.
func (o *PrincipalInfoSearchQuery) SetWithCount(v bool) {
	o.WithCount.Set(&v)
}
// SetWithCountNil sets the value for WithCount to be an explicit nil
func (o *PrincipalInfoSearchQuery) SetWithCountNil() {
	o.WithCount.Set(nil)
}

// UnsetWithCount ensures that no value is present for WithCount, not even an explicit nil
func (o *PrincipalInfoSearchQuery) UnsetWithCount() {
	o.WithCount.Unset()
}

func (o PrincipalInfoSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalInfoSearchQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
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

func (o *PrincipalInfoSearchQuery) UnmarshalJSON(data []byte) (err error) {
	varPrincipalInfoSearchQuery := _PrincipalInfoSearchQuery{}

	err = json.Unmarshal(data, &varPrincipalInfoSearchQuery)

	if err != nil {
		return err
	}

	*o = PrincipalInfoSearchQuery(varPrincipalInfoSearchQuery)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "role")
		delete(additionalProperties, "team")
		delete(additionalProperties, "sortedBy")
		delete(additionalProperties, "pageIndex")
		delete(additionalProperties, "pageSize")
		delete(additionalProperties, "withCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalInfoSearchQuery struct {
	value *PrincipalInfoSearchQuery
	isSet bool
}

func (v NullablePrincipalInfoSearchQuery) Get() *PrincipalInfoSearchQuery {
	return v.value
}

func (v *NullablePrincipalInfoSearchQuery) Set(val *PrincipalInfoSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalInfoSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalInfoSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalInfoSearchQuery(val *PrincipalInfoSearchQuery) *NullablePrincipalInfoSearchQuery {
	return &NullablePrincipalInfoSearchQuery{value: val, isSet: true}
}

func (v NullablePrincipalInfoSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalInfoSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


