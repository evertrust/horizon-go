/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the CertificateProfileSelfPermissions type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CertificateProfileSelfPermissions{}

// CertificateProfileSelfPermissions struct for CertificateProfileSelfPermissions
type CertificateProfileSelfPermissions struct {
	SelfPopRenew         utils.NullableBool `json:"selfPopRenew,omitempty"`
	SelfPopRevoke        utils.NullableBool `json:"selfPopRevoke,omitempty"`
	SelfPopUpdate        utils.NullableBool `json:"selfPopUpdate,omitempty"`
	SelfRecover          utils.NullableBool `json:"selfRecover,omitempty"`
	SelfRenew            utils.NullableBool `json:"selfRenew,omitempty"`
	SelfRevoke           utils.NullableBool `json:"selfRevoke,omitempty"`
	SelfUpdate           utils.NullableBool `json:"selfUpdate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateProfileSelfPermissions CertificateProfileSelfPermissions

// NewCertificateProfileSelfPermissions instantiates a new CertificateProfileSelfPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateProfileSelfPermissions() *CertificateProfileSelfPermissions {
	this := CertificateProfileSelfPermissions{}
	var selfPopRenew bool = false
	this.SelfPopRenew = *utils.NewNullableBool(&selfPopRenew)
	var selfPopRevoke bool = false
	this.SelfPopRevoke = *utils.NewNullableBool(&selfPopRevoke)
	var selfPopUpdate bool = false
	this.SelfPopUpdate = *utils.NewNullableBool(&selfPopUpdate)
	var selfRecover bool = false
	this.SelfRecover = *utils.NewNullableBool(&selfRecover)
	var selfRenew bool = false
	this.SelfRenew = *utils.NewNullableBool(&selfRenew)
	var selfRevoke bool = false
	this.SelfRevoke = *utils.NewNullableBool(&selfRevoke)
	var selfUpdate bool = false
	this.SelfUpdate = *utils.NewNullableBool(&selfUpdate)
	return &this
}

// NewCertificateProfileSelfPermissionsWithDefaults instantiates a new CertificateProfileSelfPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateProfileSelfPermissionsWithDefaults() *CertificateProfileSelfPermissions {
	this := CertificateProfileSelfPermissions{}
	var selfPopRenew bool = false
	this.SelfPopRenew = *utils.NewNullableBool(&selfPopRenew)
	var selfPopRevoke bool = false
	this.SelfPopRevoke = *utils.NewNullableBool(&selfPopRevoke)
	var selfPopUpdate bool = false
	this.SelfPopUpdate = *utils.NewNullableBool(&selfPopUpdate)
	var selfRecover bool = false
	this.SelfRecover = *utils.NewNullableBool(&selfRecover)
	var selfRenew bool = false
	this.SelfRenew = *utils.NewNullableBool(&selfRenew)
	var selfRevoke bool = false
	this.SelfRevoke = *utils.NewNullableBool(&selfRevoke)
	var selfUpdate bool = false
	this.SelfUpdate = *utils.NewNullableBool(&selfUpdate)
	return &this
}

// GetSelfPopRenew returns the SelfPopRenew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileSelfPermissions) GetSelfPopRenew() bool {
	if o == nil || utils.IsNil(o.SelfPopRenew.Get()) {
		var ret bool
		return ret
	}
	return *o.SelfPopRenew.Get()
}

// GetSelfPopRenewOk returns a tuple with the SelfPopRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileSelfPermissions) GetSelfPopRenewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelfPopRenew.Get(), o.SelfPopRenew.IsSet()
}

// HasSelfPopRenew returns a boolean if a field has been set.
func (o *CertificateProfileSelfPermissions) HasSelfPopRenew() bool {
	if o != nil && o.SelfPopRenew.IsSet() {
		return true
	}

	return false
}

// SetSelfPopRenew gets a reference to the given NullableBool and assigns it to the SelfPopRenew field.
func (o *CertificateProfileSelfPermissions) SetSelfPopRenew(v bool) {
	o.SelfPopRenew.Set(&v)
}

// SetSelfPopRenewNil sets the value for SelfPopRenew to be an explicit nil
func (o *CertificateProfileSelfPermissions) SetSelfPopRenewNil() {
	o.SelfPopRenew.Set(nil)
}

// UnsetSelfPopRenew ensures that no value is present for SelfPopRenew, not even an explicit nil
func (o *CertificateProfileSelfPermissions) UnsetSelfPopRenew() {
	o.SelfPopRenew.Unset()
}

// GetSelfPopRevoke returns the SelfPopRevoke field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileSelfPermissions) GetSelfPopRevoke() bool {
	if o == nil || utils.IsNil(o.SelfPopRevoke.Get()) {
		var ret bool
		return ret
	}
	return *o.SelfPopRevoke.Get()
}

// GetSelfPopRevokeOk returns a tuple with the SelfPopRevoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileSelfPermissions) GetSelfPopRevokeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelfPopRevoke.Get(), o.SelfPopRevoke.IsSet()
}

// HasSelfPopRevoke returns a boolean if a field has been set.
func (o *CertificateProfileSelfPermissions) HasSelfPopRevoke() bool {
	if o != nil && o.SelfPopRevoke.IsSet() {
		return true
	}

	return false
}

// SetSelfPopRevoke gets a reference to the given NullableBool and assigns it to the SelfPopRevoke field.
func (o *CertificateProfileSelfPermissions) SetSelfPopRevoke(v bool) {
	o.SelfPopRevoke.Set(&v)
}

// SetSelfPopRevokeNil sets the value for SelfPopRevoke to be an explicit nil
func (o *CertificateProfileSelfPermissions) SetSelfPopRevokeNil() {
	o.SelfPopRevoke.Set(nil)
}

// UnsetSelfPopRevoke ensures that no value is present for SelfPopRevoke, not even an explicit nil
func (o *CertificateProfileSelfPermissions) UnsetSelfPopRevoke() {
	o.SelfPopRevoke.Unset()
}

// GetSelfPopUpdate returns the SelfPopUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileSelfPermissions) GetSelfPopUpdate() bool {
	if o == nil || utils.IsNil(o.SelfPopUpdate.Get()) {
		var ret bool
		return ret
	}
	return *o.SelfPopUpdate.Get()
}

// GetSelfPopUpdateOk returns a tuple with the SelfPopUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileSelfPermissions) GetSelfPopUpdateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelfPopUpdate.Get(), o.SelfPopUpdate.IsSet()
}

// HasSelfPopUpdate returns a boolean if a field has been set.
func (o *CertificateProfileSelfPermissions) HasSelfPopUpdate() bool {
	if o != nil && o.SelfPopUpdate.IsSet() {
		return true
	}

	return false
}

// SetSelfPopUpdate gets a reference to the given NullableBool and assigns it to the SelfPopUpdate field.
func (o *CertificateProfileSelfPermissions) SetSelfPopUpdate(v bool) {
	o.SelfPopUpdate.Set(&v)
}

// SetSelfPopUpdateNil sets the value for SelfPopUpdate to be an explicit nil
func (o *CertificateProfileSelfPermissions) SetSelfPopUpdateNil() {
	o.SelfPopUpdate.Set(nil)
}

// UnsetSelfPopUpdate ensures that no value is present for SelfPopUpdate, not even an explicit nil
func (o *CertificateProfileSelfPermissions) UnsetSelfPopUpdate() {
	o.SelfPopUpdate.Unset()
}

// GetSelfRecover returns the SelfRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileSelfPermissions) GetSelfRecover() bool {
	if o == nil || utils.IsNil(o.SelfRecover.Get()) {
		var ret bool
		return ret
	}
	return *o.SelfRecover.Get()
}

// GetSelfRecoverOk returns a tuple with the SelfRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileSelfPermissions) GetSelfRecoverOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelfRecover.Get(), o.SelfRecover.IsSet()
}

// HasSelfRecover returns a boolean if a field has been set.
func (o *CertificateProfileSelfPermissions) HasSelfRecover() bool {
	if o != nil && o.SelfRecover.IsSet() {
		return true
	}

	return false
}

// SetSelfRecover gets a reference to the given NullableBool and assigns it to the SelfRecover field.
func (o *CertificateProfileSelfPermissions) SetSelfRecover(v bool) {
	o.SelfRecover.Set(&v)
}

// SetSelfRecoverNil sets the value for SelfRecover to be an explicit nil
func (o *CertificateProfileSelfPermissions) SetSelfRecoverNil() {
	o.SelfRecover.Set(nil)
}

// UnsetSelfRecover ensures that no value is present for SelfRecover, not even an explicit nil
func (o *CertificateProfileSelfPermissions) UnsetSelfRecover() {
	o.SelfRecover.Unset()
}

// GetSelfRenew returns the SelfRenew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileSelfPermissions) GetSelfRenew() bool {
	if o == nil || utils.IsNil(o.SelfRenew.Get()) {
		var ret bool
		return ret
	}
	return *o.SelfRenew.Get()
}

// GetSelfRenewOk returns a tuple with the SelfRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileSelfPermissions) GetSelfRenewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelfRenew.Get(), o.SelfRenew.IsSet()
}

// HasSelfRenew returns a boolean if a field has been set.
func (o *CertificateProfileSelfPermissions) HasSelfRenew() bool {
	if o != nil && o.SelfRenew.IsSet() {
		return true
	}

	return false
}

// SetSelfRenew gets a reference to the given NullableBool and assigns it to the SelfRenew field.
func (o *CertificateProfileSelfPermissions) SetSelfRenew(v bool) {
	o.SelfRenew.Set(&v)
}

// SetSelfRenewNil sets the value for SelfRenew to be an explicit nil
func (o *CertificateProfileSelfPermissions) SetSelfRenewNil() {
	o.SelfRenew.Set(nil)
}

// UnsetSelfRenew ensures that no value is present for SelfRenew, not even an explicit nil
func (o *CertificateProfileSelfPermissions) UnsetSelfRenew() {
	o.SelfRenew.Unset()
}

// GetSelfRevoke returns the SelfRevoke field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileSelfPermissions) GetSelfRevoke() bool {
	if o == nil || utils.IsNil(o.SelfRevoke.Get()) {
		var ret bool
		return ret
	}
	return *o.SelfRevoke.Get()
}

// GetSelfRevokeOk returns a tuple with the SelfRevoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileSelfPermissions) GetSelfRevokeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelfRevoke.Get(), o.SelfRevoke.IsSet()
}

// HasSelfRevoke returns a boolean if a field has been set.
func (o *CertificateProfileSelfPermissions) HasSelfRevoke() bool {
	if o != nil && o.SelfRevoke.IsSet() {
		return true
	}

	return false
}

// SetSelfRevoke gets a reference to the given NullableBool and assigns it to the SelfRevoke field.
func (o *CertificateProfileSelfPermissions) SetSelfRevoke(v bool) {
	o.SelfRevoke.Set(&v)
}

// SetSelfRevokeNil sets the value for SelfRevoke to be an explicit nil
func (o *CertificateProfileSelfPermissions) SetSelfRevokeNil() {
	o.SelfRevoke.Set(nil)
}

// UnsetSelfRevoke ensures that no value is present for SelfRevoke, not even an explicit nil
func (o *CertificateProfileSelfPermissions) UnsetSelfRevoke() {
	o.SelfRevoke.Unset()
}

// GetSelfUpdate returns the SelfUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileSelfPermissions) GetSelfUpdate() bool {
	if o == nil || utils.IsNil(o.SelfUpdate.Get()) {
		var ret bool
		return ret
	}
	return *o.SelfUpdate.Get()
}

// GetSelfUpdateOk returns a tuple with the SelfUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileSelfPermissions) GetSelfUpdateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelfUpdate.Get(), o.SelfUpdate.IsSet()
}

// HasSelfUpdate returns a boolean if a field has been set.
func (o *CertificateProfileSelfPermissions) HasSelfUpdate() bool {
	if o != nil && o.SelfUpdate.IsSet() {
		return true
	}

	return false
}

// SetSelfUpdate gets a reference to the given NullableBool and assigns it to the SelfUpdate field.
func (o *CertificateProfileSelfPermissions) SetSelfUpdate(v bool) {
	o.SelfUpdate.Set(&v)
}

// SetSelfUpdateNil sets the value for SelfUpdate to be an explicit nil
func (o *CertificateProfileSelfPermissions) SetSelfUpdateNil() {
	o.SelfUpdate.Set(nil)
}

// UnsetSelfUpdate ensures that no value is present for SelfUpdate, not even an explicit nil
func (o *CertificateProfileSelfPermissions) UnsetSelfUpdate() {
	o.SelfUpdate.Unset()
}

func (o CertificateProfileSelfPermissions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateProfileSelfPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SelfPopRenew.IsSet() {
		toSerialize["selfPopRenew"] = o.SelfPopRenew.Get()
	}
	if o.SelfPopRevoke.IsSet() {
		toSerialize["selfPopRevoke"] = o.SelfPopRevoke.Get()
	}
	if o.SelfPopUpdate.IsSet() {
		toSerialize["selfPopUpdate"] = o.SelfPopUpdate.Get()
	}
	if o.SelfRecover.IsSet() {
		toSerialize["selfRecover"] = o.SelfRecover.Get()
	}
	if o.SelfRenew.IsSet() {
		toSerialize["selfRenew"] = o.SelfRenew.Get()
	}
	if o.SelfRevoke.IsSet() {
		toSerialize["selfRevoke"] = o.SelfRevoke.Get()
	}
	if o.SelfUpdate.IsSet() {
		toSerialize["selfUpdate"] = o.SelfUpdate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateProfileSelfPermissions) UnmarshalJSON(data []byte) (err error) {
	varCertificateProfileSelfPermissions := _CertificateProfileSelfPermissions{}

	err = json.Unmarshal(data, &varCertificateProfileSelfPermissions)

	if err != nil {
		return err
	}

	*o = CertificateProfileSelfPermissions(varCertificateProfileSelfPermissions)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "selfPopRenew")
		delete(additionalProperties, "selfPopRevoke")
		delete(additionalProperties, "selfPopUpdate")
		delete(additionalProperties, "selfRecover")
		delete(additionalProperties, "selfRenew")
		delete(additionalProperties, "selfRevoke")
		delete(additionalProperties, "selfUpdate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateProfileSelfPermissions struct {
	value *CertificateProfileSelfPermissions
	isSet bool
}

func (v NullableCertificateProfileSelfPermissions) Get() *CertificateProfileSelfPermissions {
	return v.value
}

func (v *NullableCertificateProfileSelfPermissions) Set(val *CertificateProfileSelfPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateProfileSelfPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateProfileSelfPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateProfileSelfPermissions(val *CertificateProfileSelfPermissions) *NullableCertificateProfileSelfPermissions {
	return &NullableCertificateProfileSelfPermissions{value: val, isSet: true}
}

func (v NullableCertificateProfileSelfPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateProfileSelfPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
