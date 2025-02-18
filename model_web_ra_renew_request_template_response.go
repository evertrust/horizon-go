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

// checks if the WebRARenewRequestTemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebRARenewRequestTemplateResponse{}

// WebRARenewRequestTemplateResponse struct for WebRARenewRequestTemplateResponse
type WebRARenewRequestTemplateResponse struct {
	// Describes how certificates will be enrolled on this profile
	Capabilities NullableCertificateProfileCryptoPolicy `json:"capabilities,omitempty"`
	// The password policy that will be used to generate the certificate's PKCS#12 password
	PasswordPolicy NullablePasswordPolicy `json:"passwordPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebRARenewRequestTemplateResponse WebRARenewRequestTemplateResponse

// NewWebRARenewRequestTemplateResponse instantiates a new WebRARenewRequestTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebRARenewRequestTemplateResponse() *WebRARenewRequestTemplateResponse {
	this := WebRARenewRequestTemplateResponse{}
	return &this
}

// NewWebRARenewRequestTemplateResponseWithDefaults instantiates a new WebRARenewRequestTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebRARenewRequestTemplateResponseWithDefaults() *WebRARenewRequestTemplateResponse {
	this := WebRARenewRequestTemplateResponse{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARenewRequestTemplateResponse) GetCapabilities() CertificateProfileCryptoPolicy {
	if o == nil || IsNil(o.Capabilities.Get()) {
		var ret CertificateProfileCryptoPolicy
		return ret
	}
	return *o.Capabilities.Get()
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARenewRequestTemplateResponse) GetCapabilitiesOk() (*CertificateProfileCryptoPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities.Get(), o.Capabilities.IsSet()
}

// HasCapabilities returns a boolean if a field has been set.
func (o *WebRARenewRequestTemplateResponse) HasCapabilities() bool {
	if o != nil && o.Capabilities.IsSet() {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given NullableCertificateProfileCryptoPolicy and assigns it to the Capabilities field.
func (o *WebRARenewRequestTemplateResponse) SetCapabilities(v CertificateProfileCryptoPolicy) {
	o.Capabilities.Set(&v)
}
// SetCapabilitiesNil sets the value for Capabilities to be an explicit nil
func (o *WebRARenewRequestTemplateResponse) SetCapabilitiesNil() {
	o.Capabilities.Set(nil)
}

// UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
func (o *WebRARenewRequestTemplateResponse) UnsetCapabilities() {
	o.Capabilities.Unset()
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARenewRequestTemplateResponse) GetPasswordPolicy() PasswordPolicy {
	if o == nil || IsNil(o.PasswordPolicy.Get()) {
		var ret PasswordPolicy
		return ret
	}
	return *o.PasswordPolicy.Get()
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARenewRequestTemplateResponse) GetPasswordPolicyOk() (*PasswordPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordPolicy.Get(), o.PasswordPolicy.IsSet()
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *WebRARenewRequestTemplateResponse) HasPasswordPolicy() bool {
	if o != nil && o.PasswordPolicy.IsSet() {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given NullablePasswordPolicy and assigns it to the PasswordPolicy field.
func (o *WebRARenewRequestTemplateResponse) SetPasswordPolicy(v PasswordPolicy) {
	o.PasswordPolicy.Set(&v)
}
// SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil
func (o *WebRARenewRequestTemplateResponse) SetPasswordPolicyNil() {
	o.PasswordPolicy.Set(nil)
}

// UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
func (o *WebRARenewRequestTemplateResponse) UnsetPasswordPolicy() {
	o.PasswordPolicy.Unset()
}

func (o WebRARenewRequestTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebRARenewRequestTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Capabilities.IsSet() {
		toSerialize["capabilities"] = o.Capabilities.Get()
	}
	if o.PasswordPolicy.IsSet() {
		toSerialize["passwordPolicy"] = o.PasswordPolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebRARenewRequestTemplateResponse) UnmarshalJSON(data []byte) (err error) {
	varWebRARenewRequestTemplateResponse := _WebRARenewRequestTemplateResponse{}

	err = json.Unmarshal(data, &varWebRARenewRequestTemplateResponse)

	if err != nil {
		return err
	}

	*o = WebRARenewRequestTemplateResponse(varWebRARenewRequestTemplateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilities")
		delete(additionalProperties, "passwordPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebRARenewRequestTemplateResponse struct {
	value *WebRARenewRequestTemplateResponse
	isSet bool
}

func (v NullableWebRARenewRequestTemplateResponse) Get() *WebRARenewRequestTemplateResponse {
	return v.value
}

func (v *NullableWebRARenewRequestTemplateResponse) Set(val *WebRARenewRequestTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebRARenewRequestTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebRARenewRequestTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebRARenewRequestTemplateResponse(val *WebRARenewRequestTemplateResponse) *NullableWebRARenewRequestTemplateResponse {
	return &NullableWebRARenewRequestTemplateResponse{value: val, isSet: true}
}

func (v NullableWebRARenewRequestTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebRARenewRequestTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


