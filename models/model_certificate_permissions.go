/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the CertificatePermissions type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CertificatePermissions{}

// CertificatePermissions struct for CertificatePermissions
type CertificatePermissions struct {
	// Whether the principal is authorized to re-enroll this certificate
	Enroll utils.NullableBool `json:"enroll,omitempty"`
	// Whether the principal is authorized to migrate this certificate
	Migrate bool `json:"migrate"`
	// Whether the principal is authorized to recover this certificate
	Recover utils.NullableBool `json:"recover,omitempty"`
	// Whether the principal is authorized to renew this certificate
	Renew utils.NullableBool `json:"renew,omitempty"`
	// Whether the principal is authorized to request re-enrollment of this certificate
	RequestEnroll utils.NullableBool `json:"requestEnroll,omitempty"`
	// Whether the principal is authorized to request migration of this certificate
	RequestMigrate bool `json:"requestMigrate"`
	// Whether the principal is authorized to request recovery of this certificate
	RequestRecover utils.NullableBool `json:"requestRecover,omitempty"`
	// Whether the principal is authorized to request renewal of this certificate
	RequestRenew utils.NullableBool `json:"requestRenew,omitempty"`
	// Whether the principal is authorized to request revocation of this certificate
	RequestRevoke bool `json:"requestRevoke"`
	// Whether the principal is authorized to request update of this certificate
	RequestUpdate bool `json:"requestUpdate"`
	// Whether the principal is authorized to revoke this certificate
	Revoke bool `json:"revoke"`
	// Whether the principal is authorized to update this certificate
	Update               bool `json:"update"`
	AdditionalProperties map[string]interface{}
}

type _CertificatePermissions CertificatePermissions

// NewCertificatePermissions instantiates a new CertificatePermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificatePermissions(migrate bool, requestMigrate bool, requestRevoke bool, requestUpdate bool, revoke bool, update bool) *CertificatePermissions {
	this := CertificatePermissions{}
	this.Migrate = migrate
	this.RequestMigrate = requestMigrate
	this.RequestRevoke = requestRevoke
	this.RequestUpdate = requestUpdate
	this.Revoke = revoke
	this.Update = update
	return &this
}

// NewCertificatePermissionsWithDefaults instantiates a new CertificatePermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificatePermissionsWithDefaults() *CertificatePermissions {
	this := CertificatePermissions{}
	return &this
}

// GetEnroll returns the Enroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificatePermissions) GetEnroll() bool {
	if o == nil || utils.IsNil(o.Enroll.Get()) {
		var ret bool
		return ret
	}
	return *o.Enroll.Get()
}

// GetEnrollOk returns a tuple with the Enroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificatePermissions) GetEnrollOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enroll.Get(), o.Enroll.IsSet()
}

// HasEnroll returns a boolean if a field has been set.
func (o *CertificatePermissions) HasEnroll() bool {
	if o != nil && o.Enroll.IsSet() {
		return true
	}

	return false
}

// SetEnroll gets a reference to the given NullableBool and assigns it to the Enroll field.
func (o *CertificatePermissions) SetEnroll(v bool) {
	o.Enroll.Set(&v)
}

// SetEnrollNil sets the value for Enroll to be an explicit nil
func (o *CertificatePermissions) SetEnrollNil() {
	o.Enroll.Set(nil)
}

// UnsetEnroll ensures that no value is present for Enroll, not even an explicit nil
func (o *CertificatePermissions) UnsetEnroll() {
	o.Enroll.Unset()
}

// GetMigrate returns the Migrate field value
func (o *CertificatePermissions) GetMigrate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Migrate
}

// GetMigrateOk returns a tuple with the Migrate field value
// and a boolean to check if the value has been set.
func (o *CertificatePermissions) GetMigrateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Migrate, true
}

// SetMigrate sets field value
func (o *CertificatePermissions) SetMigrate(v bool) {
	o.Migrate = v
}

// GetRecover returns the Recover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificatePermissions) GetRecover() bool {
	if o == nil || utils.IsNil(o.Recover.Get()) {
		var ret bool
		return ret
	}
	return *o.Recover.Get()
}

// GetRecoverOk returns a tuple with the Recover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificatePermissions) GetRecoverOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recover.Get(), o.Recover.IsSet()
}

// HasRecover returns a boolean if a field has been set.
func (o *CertificatePermissions) HasRecover() bool {
	if o != nil && o.Recover.IsSet() {
		return true
	}

	return false
}

// SetRecover gets a reference to the given NullableBool and assigns it to the Recover field.
func (o *CertificatePermissions) SetRecover(v bool) {
	o.Recover.Set(&v)
}

// SetRecoverNil sets the value for Recover to be an explicit nil
func (o *CertificatePermissions) SetRecoverNil() {
	o.Recover.Set(nil)
}

// UnsetRecover ensures that no value is present for Recover, not even an explicit nil
func (o *CertificatePermissions) UnsetRecover() {
	o.Recover.Unset()
}

// GetRenew returns the Renew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificatePermissions) GetRenew() bool {
	if o == nil || utils.IsNil(o.Renew.Get()) {
		var ret bool
		return ret
	}
	return *o.Renew.Get()
}

// GetRenewOk returns a tuple with the Renew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificatePermissions) GetRenewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Renew.Get(), o.Renew.IsSet()
}

// HasRenew returns a boolean if a field has been set.
func (o *CertificatePermissions) HasRenew() bool {
	if o != nil && o.Renew.IsSet() {
		return true
	}

	return false
}

// SetRenew gets a reference to the given NullableBool and assigns it to the Renew field.
func (o *CertificatePermissions) SetRenew(v bool) {
	o.Renew.Set(&v)
}

// SetRenewNil sets the value for Renew to be an explicit nil
func (o *CertificatePermissions) SetRenewNil() {
	o.Renew.Set(nil)
}

// UnsetRenew ensures that no value is present for Renew, not even an explicit nil
func (o *CertificatePermissions) UnsetRenew() {
	o.Renew.Unset()
}

// GetRequestEnroll returns the RequestEnroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificatePermissions) GetRequestEnroll() bool {
	if o == nil || utils.IsNil(o.RequestEnroll.Get()) {
		var ret bool
		return ret
	}
	return *o.RequestEnroll.Get()
}

// GetRequestEnrollOk returns a tuple with the RequestEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificatePermissions) GetRequestEnrollOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestEnroll.Get(), o.RequestEnroll.IsSet()
}

// HasRequestEnroll returns a boolean if a field has been set.
func (o *CertificatePermissions) HasRequestEnroll() bool {
	if o != nil && o.RequestEnroll.IsSet() {
		return true
	}

	return false
}

// SetRequestEnroll gets a reference to the given NullableBool and assigns it to the RequestEnroll field.
func (o *CertificatePermissions) SetRequestEnroll(v bool) {
	o.RequestEnroll.Set(&v)
}

// SetRequestEnrollNil sets the value for RequestEnroll to be an explicit nil
func (o *CertificatePermissions) SetRequestEnrollNil() {
	o.RequestEnroll.Set(nil)
}

// UnsetRequestEnroll ensures that no value is present for RequestEnroll, not even an explicit nil
func (o *CertificatePermissions) UnsetRequestEnroll() {
	o.RequestEnroll.Unset()
}

// GetRequestMigrate returns the RequestMigrate field value
func (o *CertificatePermissions) GetRequestMigrate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequestMigrate
}

// GetRequestMigrateOk returns a tuple with the RequestMigrate field value
// and a boolean to check if the value has been set.
func (o *CertificatePermissions) GetRequestMigrateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestMigrate, true
}

// SetRequestMigrate sets field value
func (o *CertificatePermissions) SetRequestMigrate(v bool) {
	o.RequestMigrate = v
}

// GetRequestRecover returns the RequestRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificatePermissions) GetRequestRecover() bool {
	if o == nil || utils.IsNil(o.RequestRecover.Get()) {
		var ret bool
		return ret
	}
	return *o.RequestRecover.Get()
}

// GetRequestRecoverOk returns a tuple with the RequestRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificatePermissions) GetRequestRecoverOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestRecover.Get(), o.RequestRecover.IsSet()
}

// HasRequestRecover returns a boolean if a field has been set.
func (o *CertificatePermissions) HasRequestRecover() bool {
	if o != nil && o.RequestRecover.IsSet() {
		return true
	}

	return false
}

// SetRequestRecover gets a reference to the given NullableBool and assigns it to the RequestRecover field.
func (o *CertificatePermissions) SetRequestRecover(v bool) {
	o.RequestRecover.Set(&v)
}

// SetRequestRecoverNil sets the value for RequestRecover to be an explicit nil
func (o *CertificatePermissions) SetRequestRecoverNil() {
	o.RequestRecover.Set(nil)
}

// UnsetRequestRecover ensures that no value is present for RequestRecover, not even an explicit nil
func (o *CertificatePermissions) UnsetRequestRecover() {
	o.RequestRecover.Unset()
}

// GetRequestRenew returns the RequestRenew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificatePermissions) GetRequestRenew() bool {
	if o == nil || utils.IsNil(o.RequestRenew.Get()) {
		var ret bool
		return ret
	}
	return *o.RequestRenew.Get()
}

// GetRequestRenewOk returns a tuple with the RequestRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificatePermissions) GetRequestRenewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestRenew.Get(), o.RequestRenew.IsSet()
}

// HasRequestRenew returns a boolean if a field has been set.
func (o *CertificatePermissions) HasRequestRenew() bool {
	if o != nil && o.RequestRenew.IsSet() {
		return true
	}

	return false
}

// SetRequestRenew gets a reference to the given NullableBool and assigns it to the RequestRenew field.
func (o *CertificatePermissions) SetRequestRenew(v bool) {
	o.RequestRenew.Set(&v)
}

// SetRequestRenewNil sets the value for RequestRenew to be an explicit nil
func (o *CertificatePermissions) SetRequestRenewNil() {
	o.RequestRenew.Set(nil)
}

// UnsetRequestRenew ensures that no value is present for RequestRenew, not even an explicit nil
func (o *CertificatePermissions) UnsetRequestRenew() {
	o.RequestRenew.Unset()
}

// GetRequestRevoke returns the RequestRevoke field value
func (o *CertificatePermissions) GetRequestRevoke() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequestRevoke
}

// GetRequestRevokeOk returns a tuple with the RequestRevoke field value
// and a boolean to check if the value has been set.
func (o *CertificatePermissions) GetRequestRevokeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestRevoke, true
}

// SetRequestRevoke sets field value
func (o *CertificatePermissions) SetRequestRevoke(v bool) {
	o.RequestRevoke = v
}

// GetRequestUpdate returns the RequestUpdate field value
func (o *CertificatePermissions) GetRequestUpdate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequestUpdate
}

// GetRequestUpdateOk returns a tuple with the RequestUpdate field value
// and a boolean to check if the value has been set.
func (o *CertificatePermissions) GetRequestUpdateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestUpdate, true
}

// SetRequestUpdate sets field value
func (o *CertificatePermissions) SetRequestUpdate(v bool) {
	o.RequestUpdate = v
}

// GetRevoke returns the Revoke field value
func (o *CertificatePermissions) GetRevoke() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Revoke
}

// GetRevokeOk returns a tuple with the Revoke field value
// and a boolean to check if the value has been set.
func (o *CertificatePermissions) GetRevokeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revoke, true
}

// SetRevoke sets field value
func (o *CertificatePermissions) SetRevoke(v bool) {
	o.Revoke = v
}

// GetUpdate returns the Update field value
func (o *CertificatePermissions) GetUpdate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Update
}

// GetUpdateOk returns a tuple with the Update field value
// and a boolean to check if the value has been set.
func (o *CertificatePermissions) GetUpdateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Update, true
}

// SetUpdate sets field value
func (o *CertificatePermissions) SetUpdate(v bool) {
	o.Update = v
}

func (o CertificatePermissions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificatePermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Enroll.IsSet() {
		toSerialize["enroll"] = o.Enroll.Get()
	}
	toSerialize["migrate"] = o.Migrate
	if o.Recover.IsSet() {
		toSerialize["recover"] = o.Recover.Get()
	}
	if o.Renew.IsSet() {
		toSerialize["renew"] = o.Renew.Get()
	}
	if o.RequestEnroll.IsSet() {
		toSerialize["requestEnroll"] = o.RequestEnroll.Get()
	}
	toSerialize["requestMigrate"] = o.RequestMigrate
	if o.RequestRecover.IsSet() {
		toSerialize["requestRecover"] = o.RequestRecover.Get()
	}
	if o.RequestRenew.IsSet() {
		toSerialize["requestRenew"] = o.RequestRenew.Get()
	}
	toSerialize["requestRevoke"] = o.RequestRevoke
	toSerialize["requestUpdate"] = o.RequestUpdate
	toSerialize["revoke"] = o.Revoke
	toSerialize["update"] = o.Update

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificatePermissions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"migrate",
		"requestMigrate",
		"requestRevoke",
		"requestUpdate",
		"revoke",
		"update",
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

	varCertificatePermissions := _CertificatePermissions{}

	err = json.Unmarshal(data, &varCertificatePermissions)

	if err != nil {
		return err
	}

	*o = CertificatePermissions(varCertificatePermissions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enroll")
		delete(additionalProperties, "migrate")
		delete(additionalProperties, "recover")
		delete(additionalProperties, "renew")
		delete(additionalProperties, "requestEnroll")
		delete(additionalProperties, "requestMigrate")
		delete(additionalProperties, "requestRecover")
		delete(additionalProperties, "requestRenew")
		delete(additionalProperties, "requestRevoke")
		delete(additionalProperties, "requestUpdate")
		delete(additionalProperties, "revoke")
		delete(additionalProperties, "update")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificatePermissions struct {
	value *CertificatePermissions
	isSet bool
}

func (v NullableCertificatePermissions) Get() *CertificatePermissions {
	return v.value
}

func (v *NullableCertificatePermissions) Set(val *CertificatePermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificatePermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificatePermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificatePermissions(val *CertificatePermissions) *NullableCertificatePermissions {
	return &NullableCertificatePermissions{value: val, isSet: true}
}

func (v NullableCertificatePermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificatePermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
