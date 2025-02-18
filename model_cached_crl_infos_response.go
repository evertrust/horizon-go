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

// checks if the CachedCRLInfosResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CachedCRLInfosResponse{}

// CachedCRLInfosResponse struct for CachedCRLInfosResponse
type CachedCRLInfosResponse struct {
	Ca string `json:"ca"`
	CRLNumber NullableFloat32 `json:"cRLNumber,omitempty"`
	IssuerDn NullableString `json:"issuerDn,omitempty"`
	ThisUpdate NullableInt64 `json:"thisUpdate,omitempty"`
	NextUpdate NullableInt64 `json:"nextUpdate,omitempty"`
	LastRefresh NullableInt64 `json:"lastRefresh,omitempty"`
	NextRefresh NullableInt64 `json:"nextRefresh,omitempty"`
	CrlSize int64 `json:"crlSize"`
	Error NullableString `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CachedCRLInfosResponse CachedCRLInfosResponse

// NewCachedCRLInfosResponse instantiates a new CachedCRLInfosResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCachedCRLInfosResponse(ca string, crlSize int64) *CachedCRLInfosResponse {
	this := CachedCRLInfosResponse{}
	this.Ca = ca
	this.CrlSize = crlSize
	return &this
}

// NewCachedCRLInfosResponseWithDefaults instantiates a new CachedCRLInfosResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCachedCRLInfosResponseWithDefaults() *CachedCRLInfosResponse {
	this := CachedCRLInfosResponse{}
	return &this
}

// GetCa returns the Ca field value
func (o *CachedCRLInfosResponse) GetCa() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ca
}

// GetCaOk returns a tuple with the Ca field value
// and a boolean to check if the value has been set.
func (o *CachedCRLInfosResponse) GetCaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ca, true
}

// SetCa sets field value
func (o *CachedCRLInfosResponse) SetCa(v string) {
	o.Ca = v
}

// GetCRLNumber returns the CRLNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CachedCRLInfosResponse) GetCRLNumber() float32 {
	if o == nil || IsNil(o.CRLNumber.Get()) {
		var ret float32
		return ret
	}
	return *o.CRLNumber.Get()
}

// GetCRLNumberOk returns a tuple with the CRLNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CachedCRLInfosResponse) GetCRLNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CRLNumber.Get(), o.CRLNumber.IsSet()
}

// HasCRLNumber returns a boolean if a field has been set.
func (o *CachedCRLInfosResponse) HasCRLNumber() bool {
	if o != nil && o.CRLNumber.IsSet() {
		return true
	}

	return false
}

// SetCRLNumber gets a reference to the given NullableFloat32 and assigns it to the CRLNumber field.
func (o *CachedCRLInfosResponse) SetCRLNumber(v float32) {
	o.CRLNumber.Set(&v)
}
// SetCRLNumberNil sets the value for CRLNumber to be an explicit nil
func (o *CachedCRLInfosResponse) SetCRLNumberNil() {
	o.CRLNumber.Set(nil)
}

// UnsetCRLNumber ensures that no value is present for CRLNumber, not even an explicit nil
func (o *CachedCRLInfosResponse) UnsetCRLNumber() {
	o.CRLNumber.Unset()
}

// GetIssuerDn returns the IssuerDn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CachedCRLInfosResponse) GetIssuerDn() string {
	if o == nil || IsNil(o.IssuerDn.Get()) {
		var ret string
		return ret
	}
	return *o.IssuerDn.Get()
}

// GetIssuerDnOk returns a tuple with the IssuerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CachedCRLInfosResponse) GetIssuerDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuerDn.Get(), o.IssuerDn.IsSet()
}

// HasIssuerDn returns a boolean if a field has been set.
func (o *CachedCRLInfosResponse) HasIssuerDn() bool {
	if o != nil && o.IssuerDn.IsSet() {
		return true
	}

	return false
}

// SetIssuerDn gets a reference to the given NullableString and assigns it to the IssuerDn field.
func (o *CachedCRLInfosResponse) SetIssuerDn(v string) {
	o.IssuerDn.Set(&v)
}
// SetIssuerDnNil sets the value for IssuerDn to be an explicit nil
func (o *CachedCRLInfosResponse) SetIssuerDnNil() {
	o.IssuerDn.Set(nil)
}

// UnsetIssuerDn ensures that no value is present for IssuerDn, not even an explicit nil
func (o *CachedCRLInfosResponse) UnsetIssuerDn() {
	o.IssuerDn.Unset()
}

// GetThisUpdate returns the ThisUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CachedCRLInfosResponse) GetThisUpdate() int64 {
	if o == nil || IsNil(o.ThisUpdate.Get()) {
		var ret int64
		return ret
	}
	return *o.ThisUpdate.Get()
}

// GetThisUpdateOk returns a tuple with the ThisUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CachedCRLInfosResponse) GetThisUpdateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThisUpdate.Get(), o.ThisUpdate.IsSet()
}

// HasThisUpdate returns a boolean if a field has been set.
func (o *CachedCRLInfosResponse) HasThisUpdate() bool {
	if o != nil && o.ThisUpdate.IsSet() {
		return true
	}

	return false
}

// SetThisUpdate gets a reference to the given NullableInt64 and assigns it to the ThisUpdate field.
func (o *CachedCRLInfosResponse) SetThisUpdate(v int64) {
	o.ThisUpdate.Set(&v)
}
// SetThisUpdateNil sets the value for ThisUpdate to be an explicit nil
func (o *CachedCRLInfosResponse) SetThisUpdateNil() {
	o.ThisUpdate.Set(nil)
}

// UnsetThisUpdate ensures that no value is present for ThisUpdate, not even an explicit nil
func (o *CachedCRLInfosResponse) UnsetThisUpdate() {
	o.ThisUpdate.Unset()
}

// GetNextUpdate returns the NextUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CachedCRLInfosResponse) GetNextUpdate() int64 {
	if o == nil || IsNil(o.NextUpdate.Get()) {
		var ret int64
		return ret
	}
	return *o.NextUpdate.Get()
}

// GetNextUpdateOk returns a tuple with the NextUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CachedCRLInfosResponse) GetNextUpdateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextUpdate.Get(), o.NextUpdate.IsSet()
}

// HasNextUpdate returns a boolean if a field has been set.
func (o *CachedCRLInfosResponse) HasNextUpdate() bool {
	if o != nil && o.NextUpdate.IsSet() {
		return true
	}

	return false
}

// SetNextUpdate gets a reference to the given NullableInt64 and assigns it to the NextUpdate field.
func (o *CachedCRLInfosResponse) SetNextUpdate(v int64) {
	o.NextUpdate.Set(&v)
}
// SetNextUpdateNil sets the value for NextUpdate to be an explicit nil
func (o *CachedCRLInfosResponse) SetNextUpdateNil() {
	o.NextUpdate.Set(nil)
}

// UnsetNextUpdate ensures that no value is present for NextUpdate, not even an explicit nil
func (o *CachedCRLInfosResponse) UnsetNextUpdate() {
	o.NextUpdate.Unset()
}

// GetLastRefresh returns the LastRefresh field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CachedCRLInfosResponse) GetLastRefresh() int64 {
	if o == nil || IsNil(o.LastRefresh.Get()) {
		var ret int64
		return ret
	}
	return *o.LastRefresh.Get()
}

// GetLastRefreshOk returns a tuple with the LastRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CachedCRLInfosResponse) GetLastRefreshOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastRefresh.Get(), o.LastRefresh.IsSet()
}

// HasLastRefresh returns a boolean if a field has been set.
func (o *CachedCRLInfosResponse) HasLastRefresh() bool {
	if o != nil && o.LastRefresh.IsSet() {
		return true
	}

	return false
}

// SetLastRefresh gets a reference to the given NullableInt64 and assigns it to the LastRefresh field.
func (o *CachedCRLInfosResponse) SetLastRefresh(v int64) {
	o.LastRefresh.Set(&v)
}
// SetLastRefreshNil sets the value for LastRefresh to be an explicit nil
func (o *CachedCRLInfosResponse) SetLastRefreshNil() {
	o.LastRefresh.Set(nil)
}

// UnsetLastRefresh ensures that no value is present for LastRefresh, not even an explicit nil
func (o *CachedCRLInfosResponse) UnsetLastRefresh() {
	o.LastRefresh.Unset()
}

// GetNextRefresh returns the NextRefresh field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CachedCRLInfosResponse) GetNextRefresh() int64 {
	if o == nil || IsNil(o.NextRefresh.Get()) {
		var ret int64
		return ret
	}
	return *o.NextRefresh.Get()
}

// GetNextRefreshOk returns a tuple with the NextRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CachedCRLInfosResponse) GetNextRefreshOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextRefresh.Get(), o.NextRefresh.IsSet()
}

// HasNextRefresh returns a boolean if a field has been set.
func (o *CachedCRLInfosResponse) HasNextRefresh() bool {
	if o != nil && o.NextRefresh.IsSet() {
		return true
	}

	return false
}

// SetNextRefresh gets a reference to the given NullableInt64 and assigns it to the NextRefresh field.
func (o *CachedCRLInfosResponse) SetNextRefresh(v int64) {
	o.NextRefresh.Set(&v)
}
// SetNextRefreshNil sets the value for NextRefresh to be an explicit nil
func (o *CachedCRLInfosResponse) SetNextRefreshNil() {
	o.NextRefresh.Set(nil)
}

// UnsetNextRefresh ensures that no value is present for NextRefresh, not even an explicit nil
func (o *CachedCRLInfosResponse) UnsetNextRefresh() {
	o.NextRefresh.Unset()
}

// GetCrlSize returns the CrlSize field value
func (o *CachedCRLInfosResponse) GetCrlSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CrlSize
}

// GetCrlSizeOk returns a tuple with the CrlSize field value
// and a boolean to check if the value has been set.
func (o *CachedCRLInfosResponse) GetCrlSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CrlSize, true
}

// SetCrlSize sets field value
func (o *CachedCRLInfosResponse) SetCrlSize(v int64) {
	o.CrlSize = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CachedCRLInfosResponse) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CachedCRLInfosResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *CachedCRLInfosResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *CachedCRLInfosResponse) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *CachedCRLInfosResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *CachedCRLInfosResponse) UnsetError() {
	o.Error.Unset()
}

func (o CachedCRLInfosResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CachedCRLInfosResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ca"] = o.Ca
	if o.CRLNumber.IsSet() {
		toSerialize["cRLNumber"] = o.CRLNumber.Get()
	}
	if o.IssuerDn.IsSet() {
		toSerialize["issuerDn"] = o.IssuerDn.Get()
	}
	if o.ThisUpdate.IsSet() {
		toSerialize["thisUpdate"] = o.ThisUpdate.Get()
	}
	if o.NextUpdate.IsSet() {
		toSerialize["nextUpdate"] = o.NextUpdate.Get()
	}
	if o.LastRefresh.IsSet() {
		toSerialize["lastRefresh"] = o.LastRefresh.Get()
	}
	if o.NextRefresh.IsSet() {
		toSerialize["nextRefresh"] = o.NextRefresh.Get()
	}
	toSerialize["crlSize"] = o.CrlSize
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CachedCRLInfosResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ca",
		"crlSize",
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

	varCachedCRLInfosResponse := _CachedCRLInfosResponse{}

	err = json.Unmarshal(data, &varCachedCRLInfosResponse)

	if err != nil {
		return err
	}

	*o = CachedCRLInfosResponse(varCachedCRLInfosResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ca")
		delete(additionalProperties, "cRLNumber")
		delete(additionalProperties, "issuerDn")
		delete(additionalProperties, "thisUpdate")
		delete(additionalProperties, "nextUpdate")
		delete(additionalProperties, "lastRefresh")
		delete(additionalProperties, "nextRefresh")
		delete(additionalProperties, "crlSize")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCachedCRLInfosResponse struct {
	value *CachedCRLInfosResponse
	isSet bool
}

func (v NullableCachedCRLInfosResponse) Get() *CachedCRLInfosResponse {
	return v.value
}

func (v *NullableCachedCRLInfosResponse) Set(val *CachedCRLInfosResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCachedCRLInfosResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCachedCRLInfosResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCachedCRLInfosResponse(val *CachedCRLInfosResponse) *NullableCachedCRLInfosResponse {
	return &NullableCachedCRLInfosResponse{value: val, isSet: true}
}

func (v NullableCachedCRLInfosResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCachedCRLInfosResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


