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

// checks if the ThirdPartyConnectorSynchronizationResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ThirdPartyConnectorSynchronizationResult{}

// ThirdPartyConnectorSynchronizationResult struct for ThirdPartyConnectorSynchronizationResult
type ThirdPartyConnectorSynchronizationResult struct {
	EnrollSuccess        utils.NullableInt64 `json:"enroll_success,omitempty"`
	EnrollFailure        utils.NullableInt64 `json:"enroll_failure,omitempty"`
	RenewSuccess         utils.NullableInt64 `json:"renew_success,omitempty"`
	RenewFailure         utils.NullableInt64 `json:"renew_failure,omitempty"`
	RevokeSuccess        utils.NullableInt64 `json:"revoke_success,omitempty"`
	RevokeFailure        utils.NullableInt64 `json:"revoke_failure,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThirdPartyConnectorSynchronizationResult ThirdPartyConnectorSynchronizationResult

// NewThirdPartyConnectorSynchronizationResult instantiates a new ThirdPartyConnectorSynchronizationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyConnectorSynchronizationResult() *ThirdPartyConnectorSynchronizationResult {
	this := ThirdPartyConnectorSynchronizationResult{}
	return &this
}

// NewThirdPartyConnectorSynchronizationResultWithDefaults instantiates a new ThirdPartyConnectorSynchronizationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyConnectorSynchronizationResultWithDefaults() *ThirdPartyConnectorSynchronizationResult {
	this := ThirdPartyConnectorSynchronizationResult{}
	return &this
}

// GetEnrollSuccess returns the EnrollSuccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConnectorSynchronizationResult) GetEnrollSuccess() int64 {
	if o == nil || utils.IsNil(o.EnrollSuccess.Get()) {
		var ret int64
		return ret
	}
	return *o.EnrollSuccess.Get()
}

// GetEnrollSuccessOk returns a tuple with the EnrollSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConnectorSynchronizationResult) GetEnrollSuccessOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollSuccess.Get(), o.EnrollSuccess.IsSet()
}

// HasEnrollSuccess returns a boolean if a field has been set.
func (o *ThirdPartyConnectorSynchronizationResult) HasEnrollSuccess() bool {
	if o != nil && o.EnrollSuccess.IsSet() {
		return true
	}

	return false
}

// SetEnrollSuccess gets a reference to the given NullableInt64 and assigns it to the EnrollSuccess field.
func (o *ThirdPartyConnectorSynchronizationResult) SetEnrollSuccess(v int64) {
	o.EnrollSuccess.Set(&v)
}

// SetEnrollSuccessNil sets the value for EnrollSuccess to be an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) SetEnrollSuccessNil() {
	o.EnrollSuccess.Set(nil)
}

// UnsetEnrollSuccess ensures that no value is present for EnrollSuccess, not even an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) UnsetEnrollSuccess() {
	o.EnrollSuccess.Unset()
}

// GetEnrollFailure returns the EnrollFailure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConnectorSynchronizationResult) GetEnrollFailure() int64 {
	if o == nil || utils.IsNil(o.EnrollFailure.Get()) {
		var ret int64
		return ret
	}
	return *o.EnrollFailure.Get()
}

// GetEnrollFailureOk returns a tuple with the EnrollFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConnectorSynchronizationResult) GetEnrollFailureOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollFailure.Get(), o.EnrollFailure.IsSet()
}

// HasEnrollFailure returns a boolean if a field has been set.
func (o *ThirdPartyConnectorSynchronizationResult) HasEnrollFailure() bool {
	if o != nil && o.EnrollFailure.IsSet() {
		return true
	}

	return false
}

// SetEnrollFailure gets a reference to the given NullableInt64 and assigns it to the EnrollFailure field.
func (o *ThirdPartyConnectorSynchronizationResult) SetEnrollFailure(v int64) {
	o.EnrollFailure.Set(&v)
}

// SetEnrollFailureNil sets the value for EnrollFailure to be an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) SetEnrollFailureNil() {
	o.EnrollFailure.Set(nil)
}

// UnsetEnrollFailure ensures that no value is present for EnrollFailure, not even an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) UnsetEnrollFailure() {
	o.EnrollFailure.Unset()
}

// GetRenewSuccess returns the RenewSuccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConnectorSynchronizationResult) GetRenewSuccess() int64 {
	if o == nil || utils.IsNil(o.RenewSuccess.Get()) {
		var ret int64
		return ret
	}
	return *o.RenewSuccess.Get()
}

// GetRenewSuccessOk returns a tuple with the RenewSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConnectorSynchronizationResult) GetRenewSuccessOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewSuccess.Get(), o.RenewSuccess.IsSet()
}

// HasRenewSuccess returns a boolean if a field has been set.
func (o *ThirdPartyConnectorSynchronizationResult) HasRenewSuccess() bool {
	if o != nil && o.RenewSuccess.IsSet() {
		return true
	}

	return false
}

// SetRenewSuccess gets a reference to the given NullableInt64 and assigns it to the RenewSuccess field.
func (o *ThirdPartyConnectorSynchronizationResult) SetRenewSuccess(v int64) {
	o.RenewSuccess.Set(&v)
}

// SetRenewSuccessNil sets the value for RenewSuccess to be an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) SetRenewSuccessNil() {
	o.RenewSuccess.Set(nil)
}

// UnsetRenewSuccess ensures that no value is present for RenewSuccess, not even an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) UnsetRenewSuccess() {
	o.RenewSuccess.Unset()
}

// GetRenewFailure returns the RenewFailure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConnectorSynchronizationResult) GetRenewFailure() int64 {
	if o == nil || utils.IsNil(o.RenewFailure.Get()) {
		var ret int64
		return ret
	}
	return *o.RenewFailure.Get()
}

// GetRenewFailureOk returns a tuple with the RenewFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConnectorSynchronizationResult) GetRenewFailureOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewFailure.Get(), o.RenewFailure.IsSet()
}

// HasRenewFailure returns a boolean if a field has been set.
func (o *ThirdPartyConnectorSynchronizationResult) HasRenewFailure() bool {
	if o != nil && o.RenewFailure.IsSet() {
		return true
	}

	return false
}

// SetRenewFailure gets a reference to the given NullableInt64 and assigns it to the RenewFailure field.
func (o *ThirdPartyConnectorSynchronizationResult) SetRenewFailure(v int64) {
	o.RenewFailure.Set(&v)
}

// SetRenewFailureNil sets the value for RenewFailure to be an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) SetRenewFailureNil() {
	o.RenewFailure.Set(nil)
}

// UnsetRenewFailure ensures that no value is present for RenewFailure, not even an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) UnsetRenewFailure() {
	o.RenewFailure.Unset()
}

// GetRevokeSuccess returns the RevokeSuccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConnectorSynchronizationResult) GetRevokeSuccess() int64 {
	if o == nil || utils.IsNil(o.RevokeSuccess.Get()) {
		var ret int64
		return ret
	}
	return *o.RevokeSuccess.Get()
}

// GetRevokeSuccessOk returns a tuple with the RevokeSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConnectorSynchronizationResult) GetRevokeSuccessOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevokeSuccess.Get(), o.RevokeSuccess.IsSet()
}

// HasRevokeSuccess returns a boolean if a field has been set.
func (o *ThirdPartyConnectorSynchronizationResult) HasRevokeSuccess() bool {
	if o != nil && o.RevokeSuccess.IsSet() {
		return true
	}

	return false
}

// SetRevokeSuccess gets a reference to the given NullableInt64 and assigns it to the RevokeSuccess field.
func (o *ThirdPartyConnectorSynchronizationResult) SetRevokeSuccess(v int64) {
	o.RevokeSuccess.Set(&v)
}

// SetRevokeSuccessNil sets the value for RevokeSuccess to be an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) SetRevokeSuccessNil() {
	o.RevokeSuccess.Set(nil)
}

// UnsetRevokeSuccess ensures that no value is present for RevokeSuccess, not even an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) UnsetRevokeSuccess() {
	o.RevokeSuccess.Unset()
}

// GetRevokeFailure returns the RevokeFailure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConnectorSynchronizationResult) GetRevokeFailure() int64 {
	if o == nil || utils.IsNil(o.RevokeFailure.Get()) {
		var ret int64
		return ret
	}
	return *o.RevokeFailure.Get()
}

// GetRevokeFailureOk returns a tuple with the RevokeFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConnectorSynchronizationResult) GetRevokeFailureOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevokeFailure.Get(), o.RevokeFailure.IsSet()
}

// HasRevokeFailure returns a boolean if a field has been set.
func (o *ThirdPartyConnectorSynchronizationResult) HasRevokeFailure() bool {
	if o != nil && o.RevokeFailure.IsSet() {
		return true
	}

	return false
}

// SetRevokeFailure gets a reference to the given NullableInt64 and assigns it to the RevokeFailure field.
func (o *ThirdPartyConnectorSynchronizationResult) SetRevokeFailure(v int64) {
	o.RevokeFailure.Set(&v)
}

// SetRevokeFailureNil sets the value for RevokeFailure to be an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) SetRevokeFailureNil() {
	o.RevokeFailure.Set(nil)
}

// UnsetRevokeFailure ensures that no value is present for RevokeFailure, not even an explicit nil
func (o *ThirdPartyConnectorSynchronizationResult) UnsetRevokeFailure() {
	o.RevokeFailure.Unset()
}

func (o ThirdPartyConnectorSynchronizationResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyConnectorSynchronizationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EnrollSuccess.IsSet() {
		toSerialize["enroll_success"] = o.EnrollSuccess.Get()
	}
	if o.EnrollFailure.IsSet() {
		toSerialize["enroll_failure"] = o.EnrollFailure.Get()
	}
	if o.RenewSuccess.IsSet() {
		toSerialize["renew_success"] = o.RenewSuccess.Get()
	}
	if o.RenewFailure.IsSet() {
		toSerialize["renew_failure"] = o.RenewFailure.Get()
	}
	if o.RevokeSuccess.IsSet() {
		toSerialize["revoke_success"] = o.RevokeSuccess.Get()
	}
	if o.RevokeFailure.IsSet() {
		toSerialize["revoke_failure"] = o.RevokeFailure.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThirdPartyConnectorSynchronizationResult) UnmarshalJSON(data []byte) (err error) {
	varThirdPartyConnectorSynchronizationResult := _ThirdPartyConnectorSynchronizationResult{}

	err = json.Unmarshal(data, &varThirdPartyConnectorSynchronizationResult)

	if err != nil {
		return err
	}

	*o = ThirdPartyConnectorSynchronizationResult(varThirdPartyConnectorSynchronizationResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enroll_success")
		delete(additionalProperties, "enroll_failure")
		delete(additionalProperties, "renew_success")
		delete(additionalProperties, "renew_failure")
		delete(additionalProperties, "revoke_success")
		delete(additionalProperties, "revoke_failure")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThirdPartyConnectorSynchronizationResult struct {
	value *ThirdPartyConnectorSynchronizationResult
	isSet bool
}

func (v NullableThirdPartyConnectorSynchronizationResult) Get() *ThirdPartyConnectorSynchronizationResult {
	return v.value
}

func (v *NullableThirdPartyConnectorSynchronizationResult) Set(val *ThirdPartyConnectorSynchronizationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyConnectorSynchronizationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyConnectorSynchronizationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyConnectorSynchronizationResult(val *ThirdPartyConnectorSynchronizationResult) *NullableThirdPartyConnectorSynchronizationResult {
	return &NullableThirdPartyConnectorSynchronizationResult{value: val, isSet: true}
}

func (v NullableThirdPartyConnectorSynchronizationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyConnectorSynchronizationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
