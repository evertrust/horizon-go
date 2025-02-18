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

// checks if the CertificateProfileCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateProfileCapabilities{}

// CertificateProfileCapabilities struct for CertificateProfileCapabilities
type CertificateProfileCapabilities struct {
	// Centralized enrollment is enabled on this profile
	Centralized bool `json:"centralized"`
	// Decentralized enrollment is enabled on this profile
	Decentralized bool `json:"decentralized"`
	// Key type used when no keyType has been chosen
	DefaultKeyType NullableString `json:"defaultKeyType,omitempty" validate:"regexp=(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87)(\\\\\\\\+(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87))?"`
	// The list of key types that are authorized for enrollment. A null value means all keys are allowed
	AuthorizedKeyTypes []string `json:"authorizedKeyTypes,omitempty"`
	// The enrollment mode that should be prioritized when both are defined
	PreferredEnrollmentMode NullableString `json:"preferredEnrollmentMode,omitempty"`
	// The selected password policy for this profile. If none is defined and the password mode is `manual`, there is no constraint on the password. In `random` mode, the `Horizon-Default` policy is used
	PasswordPolicy NullablePasswordPolicy `json:"passwordPolicy,omitempty"`
	// A `manual` password mode means the password for the PKCS#12 must be set in the request. A `random` password will be generated on Horizon
	P12passwordMode NullableString `json:"p12passwordMode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateProfileCapabilities CertificateProfileCapabilities

// NewCertificateProfileCapabilities instantiates a new CertificateProfileCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateProfileCapabilities(centralized bool, decentralized bool) *CertificateProfileCapabilities {
	this := CertificateProfileCapabilities{}
	this.Centralized = centralized
	this.Decentralized = decentralized
	return &this
}

// NewCertificateProfileCapabilitiesWithDefaults instantiates a new CertificateProfileCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateProfileCapabilitiesWithDefaults() *CertificateProfileCapabilities {
	this := CertificateProfileCapabilities{}
	return &this
}

// GetCentralized returns the Centralized field value
func (o *CertificateProfileCapabilities) GetCentralized() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Centralized
}

// GetCentralizedOk returns a tuple with the Centralized field value
// and a boolean to check if the value has been set.
func (o *CertificateProfileCapabilities) GetCentralizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Centralized, true
}

// SetCentralized sets field value
func (o *CertificateProfileCapabilities) SetCentralized(v bool) {
	o.Centralized = v
}

// GetDecentralized returns the Decentralized field value
func (o *CertificateProfileCapabilities) GetDecentralized() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Decentralized
}

// GetDecentralizedOk returns a tuple with the Decentralized field value
// and a boolean to check if the value has been set.
func (o *CertificateProfileCapabilities) GetDecentralizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decentralized, true
}

// SetDecentralized sets field value
func (o *CertificateProfileCapabilities) SetDecentralized(v bool) {
	o.Decentralized = v
}

// GetDefaultKeyType returns the DefaultKeyType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCapabilities) GetDefaultKeyType() string {
	if o == nil || IsNil(o.DefaultKeyType.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultKeyType.Get()
}

// GetDefaultKeyTypeOk returns a tuple with the DefaultKeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCapabilities) GetDefaultKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultKeyType.Get(), o.DefaultKeyType.IsSet()
}

// HasDefaultKeyType returns a boolean if a field has been set.
func (o *CertificateProfileCapabilities) HasDefaultKeyType() bool {
	if o != nil && o.DefaultKeyType.IsSet() {
		return true
	}

	return false
}

// SetDefaultKeyType gets a reference to the given NullableString and assigns it to the DefaultKeyType field.
func (o *CertificateProfileCapabilities) SetDefaultKeyType(v string) {
	o.DefaultKeyType.Set(&v)
}
// SetDefaultKeyTypeNil sets the value for DefaultKeyType to be an explicit nil
func (o *CertificateProfileCapabilities) SetDefaultKeyTypeNil() {
	o.DefaultKeyType.Set(nil)
}

// UnsetDefaultKeyType ensures that no value is present for DefaultKeyType, not even an explicit nil
func (o *CertificateProfileCapabilities) UnsetDefaultKeyType() {
	o.DefaultKeyType.Unset()
}

// GetAuthorizedKeyTypes returns the AuthorizedKeyTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCapabilities) GetAuthorizedKeyTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AuthorizedKeyTypes
}

// GetAuthorizedKeyTypesOk returns a tuple with the AuthorizedKeyTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCapabilities) GetAuthorizedKeyTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthorizedKeyTypes) {
		return nil, false
	}
	return o.AuthorizedKeyTypes, true
}

// HasAuthorizedKeyTypes returns a boolean if a field has been set.
func (o *CertificateProfileCapabilities) HasAuthorizedKeyTypes() bool {
	if o != nil && !IsNil(o.AuthorizedKeyTypes) {
		return true
	}

	return false
}

// SetAuthorizedKeyTypes gets a reference to the given []string and assigns it to the AuthorizedKeyTypes field.
func (o *CertificateProfileCapabilities) SetAuthorizedKeyTypes(v []string) {
	o.AuthorizedKeyTypes = v
}

// GetPreferredEnrollmentMode returns the PreferredEnrollmentMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCapabilities) GetPreferredEnrollmentMode() string {
	if o == nil || IsNil(o.PreferredEnrollmentMode.Get()) {
		var ret string
		return ret
	}
	return *o.PreferredEnrollmentMode.Get()
}

// GetPreferredEnrollmentModeOk returns a tuple with the PreferredEnrollmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCapabilities) GetPreferredEnrollmentModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferredEnrollmentMode.Get(), o.PreferredEnrollmentMode.IsSet()
}

// HasPreferredEnrollmentMode returns a boolean if a field has been set.
func (o *CertificateProfileCapabilities) HasPreferredEnrollmentMode() bool {
	if o != nil && o.PreferredEnrollmentMode.IsSet() {
		return true
	}

	return false
}

// SetPreferredEnrollmentMode gets a reference to the given NullableString and assigns it to the PreferredEnrollmentMode field.
func (o *CertificateProfileCapabilities) SetPreferredEnrollmentMode(v string) {
	o.PreferredEnrollmentMode.Set(&v)
}
// SetPreferredEnrollmentModeNil sets the value for PreferredEnrollmentMode to be an explicit nil
func (o *CertificateProfileCapabilities) SetPreferredEnrollmentModeNil() {
	o.PreferredEnrollmentMode.Set(nil)
}

// UnsetPreferredEnrollmentMode ensures that no value is present for PreferredEnrollmentMode, not even an explicit nil
func (o *CertificateProfileCapabilities) UnsetPreferredEnrollmentMode() {
	o.PreferredEnrollmentMode.Unset()
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCapabilities) GetPasswordPolicy() PasswordPolicy {
	if o == nil || IsNil(o.PasswordPolicy.Get()) {
		var ret PasswordPolicy
		return ret
	}
	return *o.PasswordPolicy.Get()
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCapabilities) GetPasswordPolicyOk() (*PasswordPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordPolicy.Get(), o.PasswordPolicy.IsSet()
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *CertificateProfileCapabilities) HasPasswordPolicy() bool {
	if o != nil && o.PasswordPolicy.IsSet() {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given NullablePasswordPolicy and assigns it to the PasswordPolicy field.
func (o *CertificateProfileCapabilities) SetPasswordPolicy(v PasswordPolicy) {
	o.PasswordPolicy.Set(&v)
}
// SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil
func (o *CertificateProfileCapabilities) SetPasswordPolicyNil() {
	o.PasswordPolicy.Set(nil)
}

// UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
func (o *CertificateProfileCapabilities) UnsetPasswordPolicy() {
	o.PasswordPolicy.Unset()
}

// GetP12passwordMode returns the P12passwordMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCapabilities) GetP12passwordMode() string {
	if o == nil || IsNil(o.P12passwordMode.Get()) {
		var ret string
		return ret
	}
	return *o.P12passwordMode.Get()
}

// GetP12passwordModeOk returns a tuple with the P12passwordMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCapabilities) GetP12passwordModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.P12passwordMode.Get(), o.P12passwordMode.IsSet()
}

// HasP12passwordMode returns a boolean if a field has been set.
func (o *CertificateProfileCapabilities) HasP12passwordMode() bool {
	if o != nil && o.P12passwordMode.IsSet() {
		return true
	}

	return false
}

// SetP12passwordMode gets a reference to the given NullableString and assigns it to the P12passwordMode field.
func (o *CertificateProfileCapabilities) SetP12passwordMode(v string) {
	o.P12passwordMode.Set(&v)
}
// SetP12passwordModeNil sets the value for P12passwordMode to be an explicit nil
func (o *CertificateProfileCapabilities) SetP12passwordModeNil() {
	o.P12passwordMode.Set(nil)
}

// UnsetP12passwordMode ensures that no value is present for P12passwordMode, not even an explicit nil
func (o *CertificateProfileCapabilities) UnsetP12passwordMode() {
	o.P12passwordMode.Unset()
}

func (o CertificateProfileCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateProfileCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["centralized"] = o.Centralized
	toSerialize["decentralized"] = o.Decentralized
	if o.DefaultKeyType.IsSet() {
		toSerialize["defaultKeyType"] = o.DefaultKeyType.Get()
	}
	if o.AuthorizedKeyTypes != nil {
		toSerialize["authorizedKeyTypes"] = o.AuthorizedKeyTypes
	}
	if o.PreferredEnrollmentMode.IsSet() {
		toSerialize["preferredEnrollmentMode"] = o.PreferredEnrollmentMode.Get()
	}
	if o.PasswordPolicy.IsSet() {
		toSerialize["passwordPolicy"] = o.PasswordPolicy.Get()
	}
	if o.P12passwordMode.IsSet() {
		toSerialize["p12passwordMode"] = o.P12passwordMode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateProfileCapabilities) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"centralized",
		"decentralized",
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

	varCertificateProfileCapabilities := _CertificateProfileCapabilities{}

	err = json.Unmarshal(data, &varCertificateProfileCapabilities)

	if err != nil {
		return err
	}

	*o = CertificateProfileCapabilities(varCertificateProfileCapabilities)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "centralized")
		delete(additionalProperties, "decentralized")
		delete(additionalProperties, "defaultKeyType")
		delete(additionalProperties, "authorizedKeyTypes")
		delete(additionalProperties, "preferredEnrollmentMode")
		delete(additionalProperties, "passwordPolicy")
		delete(additionalProperties, "p12passwordMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateProfileCapabilities struct {
	value *CertificateProfileCapabilities
	isSet bool
}

func (v NullableCertificateProfileCapabilities) Get() *CertificateProfileCapabilities {
	return v.value
}

func (v *NullableCertificateProfileCapabilities) Set(val *CertificateProfileCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateProfileCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateProfileCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateProfileCapabilities(val *CertificateProfileCapabilities) *NullableCertificateProfileCapabilities {
	return &NullableCertificateProfileCapabilities{value: val, isSet: true}
}

func (v NullableCertificateProfileCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateProfileCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


