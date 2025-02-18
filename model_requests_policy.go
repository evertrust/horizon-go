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

// checks if the RequestsPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestsPolicy{}

// RequestsPolicy struct for RequestsPolicy
type RequestsPolicy struct {
	Enroll NullableString `json:"enroll,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Revoke NullableString `json:"revoke,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Recover NullableString `json:"recover,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Update NullableString `json:"update,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Migrate NullableString `json:"migrate,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Renew NullableString `json:"renew,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	AdditionalProperties map[string]interface{}
}

type _RequestsPolicy RequestsPolicy

// NewRequestsPolicy instantiates a new RequestsPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsPolicy() *RequestsPolicy {
	this := RequestsPolicy{}
	return &this
}

// NewRequestsPolicyWithDefaults instantiates a new RequestsPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsPolicyWithDefaults() *RequestsPolicy {
	this := RequestsPolicy{}
	return &this
}

// GetEnroll returns the Enroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestsPolicy) GetEnroll() string {
	if o == nil || IsNil(o.Enroll.Get()) {
		var ret string
		return ret
	}
	return *o.Enroll.Get()
}

// GetEnrollOk returns a tuple with the Enroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestsPolicy) GetEnrollOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enroll.Get(), o.Enroll.IsSet()
}

// HasEnroll returns a boolean if a field has been set.
func (o *RequestsPolicy) HasEnroll() bool {
	if o != nil && o.Enroll.IsSet() {
		return true
	}

	return false
}

// SetEnroll gets a reference to the given NullableString and assigns it to the Enroll field.
func (o *RequestsPolicy) SetEnroll(v string) {
	o.Enroll.Set(&v)
}
// SetEnrollNil sets the value for Enroll to be an explicit nil
func (o *RequestsPolicy) SetEnrollNil() {
	o.Enroll.Set(nil)
}

// UnsetEnroll ensures that no value is present for Enroll, not even an explicit nil
func (o *RequestsPolicy) UnsetEnroll() {
	o.Enroll.Unset()
}

// GetRevoke returns the Revoke field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestsPolicy) GetRevoke() string {
	if o == nil || IsNil(o.Revoke.Get()) {
		var ret string
		return ret
	}
	return *o.Revoke.Get()
}

// GetRevokeOk returns a tuple with the Revoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestsPolicy) GetRevokeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revoke.Get(), o.Revoke.IsSet()
}

// HasRevoke returns a boolean if a field has been set.
func (o *RequestsPolicy) HasRevoke() bool {
	if o != nil && o.Revoke.IsSet() {
		return true
	}

	return false
}

// SetRevoke gets a reference to the given NullableString and assigns it to the Revoke field.
func (o *RequestsPolicy) SetRevoke(v string) {
	o.Revoke.Set(&v)
}
// SetRevokeNil sets the value for Revoke to be an explicit nil
func (o *RequestsPolicy) SetRevokeNil() {
	o.Revoke.Set(nil)
}

// UnsetRevoke ensures that no value is present for Revoke, not even an explicit nil
func (o *RequestsPolicy) UnsetRevoke() {
	o.Revoke.Unset()
}

// GetRecover returns the Recover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestsPolicy) GetRecover() string {
	if o == nil || IsNil(o.Recover.Get()) {
		var ret string
		return ret
	}
	return *o.Recover.Get()
}

// GetRecoverOk returns a tuple with the Recover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestsPolicy) GetRecoverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recover.Get(), o.Recover.IsSet()
}

// HasRecover returns a boolean if a field has been set.
func (o *RequestsPolicy) HasRecover() bool {
	if o != nil && o.Recover.IsSet() {
		return true
	}

	return false
}

// SetRecover gets a reference to the given NullableString and assigns it to the Recover field.
func (o *RequestsPolicy) SetRecover(v string) {
	o.Recover.Set(&v)
}
// SetRecoverNil sets the value for Recover to be an explicit nil
func (o *RequestsPolicy) SetRecoverNil() {
	o.Recover.Set(nil)
}

// UnsetRecover ensures that no value is present for Recover, not even an explicit nil
func (o *RequestsPolicy) UnsetRecover() {
	o.Recover.Unset()
}

// GetUpdate returns the Update field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestsPolicy) GetUpdate() string {
	if o == nil || IsNil(o.Update.Get()) {
		var ret string
		return ret
	}
	return *o.Update.Get()
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestsPolicy) GetUpdateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Update.Get(), o.Update.IsSet()
}

// HasUpdate returns a boolean if a field has been set.
func (o *RequestsPolicy) HasUpdate() bool {
	if o != nil && o.Update.IsSet() {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given NullableString and assigns it to the Update field.
func (o *RequestsPolicy) SetUpdate(v string) {
	o.Update.Set(&v)
}
// SetUpdateNil sets the value for Update to be an explicit nil
func (o *RequestsPolicy) SetUpdateNil() {
	o.Update.Set(nil)
}

// UnsetUpdate ensures that no value is present for Update, not even an explicit nil
func (o *RequestsPolicy) UnsetUpdate() {
	o.Update.Unset()
}

// GetMigrate returns the Migrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestsPolicy) GetMigrate() string {
	if o == nil || IsNil(o.Migrate.Get()) {
		var ret string
		return ret
	}
	return *o.Migrate.Get()
}

// GetMigrateOk returns a tuple with the Migrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestsPolicy) GetMigrateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Migrate.Get(), o.Migrate.IsSet()
}

// HasMigrate returns a boolean if a field has been set.
func (o *RequestsPolicy) HasMigrate() bool {
	if o != nil && o.Migrate.IsSet() {
		return true
	}

	return false
}

// SetMigrate gets a reference to the given NullableString and assigns it to the Migrate field.
func (o *RequestsPolicy) SetMigrate(v string) {
	o.Migrate.Set(&v)
}
// SetMigrateNil sets the value for Migrate to be an explicit nil
func (o *RequestsPolicy) SetMigrateNil() {
	o.Migrate.Set(nil)
}

// UnsetMigrate ensures that no value is present for Migrate, not even an explicit nil
func (o *RequestsPolicy) UnsetMigrate() {
	o.Migrate.Unset()
}

// GetRenew returns the Renew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestsPolicy) GetRenew() string {
	if o == nil || IsNil(o.Renew.Get()) {
		var ret string
		return ret
	}
	return *o.Renew.Get()
}

// GetRenewOk returns a tuple with the Renew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestsPolicy) GetRenewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Renew.Get(), o.Renew.IsSet()
}

// HasRenew returns a boolean if a field has been set.
func (o *RequestsPolicy) HasRenew() bool {
	if o != nil && o.Renew.IsSet() {
		return true
	}

	return false
}

// SetRenew gets a reference to the given NullableString and assigns it to the Renew field.
func (o *RequestsPolicy) SetRenew(v string) {
	o.Renew.Set(&v)
}
// SetRenewNil sets the value for Renew to be an explicit nil
func (o *RequestsPolicy) SetRenewNil() {
	o.Renew.Set(nil)
}

// UnsetRenew ensures that no value is present for Renew, not even an explicit nil
func (o *RequestsPolicy) UnsetRenew() {
	o.Renew.Unset()
}

func (o RequestsPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestsPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enroll.IsSet() {
		toSerialize["enroll"] = o.Enroll.Get()
	}
	if o.Revoke.IsSet() {
		toSerialize["revoke"] = o.Revoke.Get()
	}
	if o.Recover.IsSet() {
		toSerialize["recover"] = o.Recover.Get()
	}
	if o.Update.IsSet() {
		toSerialize["update"] = o.Update.Get()
	}
	if o.Migrate.IsSet() {
		toSerialize["migrate"] = o.Migrate.Get()
	}
	if o.Renew.IsSet() {
		toSerialize["renew"] = o.Renew.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestsPolicy) UnmarshalJSON(data []byte) (err error) {
	varRequestsPolicy := _RequestsPolicy{}

	err = json.Unmarshal(data, &varRequestsPolicy)

	if err != nil {
		return err
	}

	*o = RequestsPolicy(varRequestsPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enroll")
		delete(additionalProperties, "revoke")
		delete(additionalProperties, "recover")
		delete(additionalProperties, "update")
		delete(additionalProperties, "migrate")
		delete(additionalProperties, "renew")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestsPolicy struct {
	value *RequestsPolicy
	isSet bool
}

func (v NullableRequestsPolicy) Get() *RequestsPolicy {
	return v.value
}

func (v *NullableRequestsPolicy) Set(val *RequestsPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsPolicy(val *RequestsPolicy) *NullableRequestsPolicy {
	return &NullableRequestsPolicy{value: val, isSet: true}
}

func (v NullableRequestsPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


