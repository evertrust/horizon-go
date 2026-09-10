/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the WebraInitParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebraInitParameters{}

// WebraInitParameters struct for WebraInitParameters
type WebraInitParameters struct {
	// The authorization mode for WebRA.
	AuthorizationMode *string `json:"authorizationMode,omitempty"`
	// The enrollment mode for WebRA.
	EnrollmentMode *string `json:"enrollmentMode,omitempty"`
	// The key type used for WebRA.
	KeyType *string `json:"keyType,omitempty"`
	// The module of the initialization parameters.
	Module string `json:"module"`
	// The password policy for WebRA.
	PasswordPolicy *PasswordPolicy `json:"passwordPolicy,omitempty"`
	// The profile used for WebRA.
	Profile              string `json:"profile"`
	AdditionalProperties map[string]interface{}
}

type _WebraInitParameters WebraInitParameters

// NewWebraInitParameters instantiates a new WebraInitParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebraInitParameters(module string, profile string) *WebraInitParameters {
	this := WebraInitParameters{}
	this.Module = module
	this.Profile = profile
	return &this
}

// NewWebraInitParametersWithDefaults instantiates a new WebraInitParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebraInitParametersWithDefaults() *WebraInitParameters {
	this := WebraInitParameters{}
	return &this
}

// GetAuthorizationMode returns the AuthorizationMode field value if set, zero value otherwise.
func (o *WebraInitParameters) GetAuthorizationMode() string {
	if o == nil || utils.IsNil(o.AuthorizationMode) {
		var ret string
		return ret
	}
	return *o.AuthorizationMode
}

// GetAuthorizationModeOk returns a tuple with the AuthorizationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebraInitParameters) GetAuthorizationModeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AuthorizationMode) {
		return nil, false
	}
	return o.AuthorizationMode, true
}

// HasAuthorizationMode returns a boolean if a field has been set.
func (o *WebraInitParameters) HasAuthorizationMode() bool {
	if o != nil && !utils.IsNil(o.AuthorizationMode) {
		return true
	}

	return false
}

// SetAuthorizationMode gets a reference to the given string and assigns it to the AuthorizationMode field.
func (o *WebraInitParameters) SetAuthorizationMode(v string) {
	o.AuthorizationMode = &v
}

// GetEnrollmentMode returns the EnrollmentMode field value if set, zero value otherwise.
func (o *WebraInitParameters) GetEnrollmentMode() string {
	if o == nil || utils.IsNil(o.EnrollmentMode) {
		var ret string
		return ret
	}
	return *o.EnrollmentMode
}

// GetEnrollmentModeOk returns a tuple with the EnrollmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebraInitParameters) GetEnrollmentModeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EnrollmentMode) {
		return nil, false
	}
	return o.EnrollmentMode, true
}

// HasEnrollmentMode returns a boolean if a field has been set.
func (o *WebraInitParameters) HasEnrollmentMode() bool {
	if o != nil && !utils.IsNil(o.EnrollmentMode) {
		return true
	}

	return false
}

// SetEnrollmentMode gets a reference to the given string and assigns it to the EnrollmentMode field.
func (o *WebraInitParameters) SetEnrollmentMode(v string) {
	o.EnrollmentMode = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *WebraInitParameters) GetKeyType() string {
	if o == nil || utils.IsNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebraInitParameters) GetKeyTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *WebraInitParameters) HasKeyType() bool {
	if o != nil && !utils.IsNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *WebraInitParameters) SetKeyType(v string) {
	o.KeyType = &v
}

// GetModule returns the Module field value
func (o *WebraInitParameters) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *WebraInitParameters) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *WebraInitParameters) SetModule(v string) {
	o.Module = v
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise.
func (o *WebraInitParameters) GetPasswordPolicy() PasswordPolicy {
	if o == nil || utils.IsNil(o.PasswordPolicy) {
		var ret PasswordPolicy
		return ret
	}
	return *o.PasswordPolicy
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebraInitParameters) GetPasswordPolicyOk() (*PasswordPolicy, bool) {
	if o == nil || utils.IsNil(o.PasswordPolicy) {
		return nil, false
	}
	return o.PasswordPolicy, true
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *WebraInitParameters) HasPasswordPolicy() bool {
	if o != nil && !utils.IsNil(o.PasswordPolicy) {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given PasswordPolicy and assigns it to the PasswordPolicy field.
func (o *WebraInitParameters) SetPasswordPolicy(v PasswordPolicy) {
	o.PasswordPolicy = &v
}

// GetProfile returns the Profile field value
func (o *WebraInitParameters) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *WebraInitParameters) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *WebraInitParameters) SetProfile(v string) {
	o.Profile = v
}

func (o WebraInitParameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebraInitParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AuthorizationMode) {
		toSerialize["authorizationMode"] = o.AuthorizationMode
	}
	if !utils.IsNil(o.EnrollmentMode) {
		toSerialize["enrollmentMode"] = o.EnrollmentMode
	}
	if !utils.IsNil(o.KeyType) {
		toSerialize["keyType"] = o.KeyType
	}
	toSerialize["module"] = o.Module
	if !utils.IsNil(o.PasswordPolicy) {
		toSerialize["passwordPolicy"] = o.PasswordPolicy
	}
	toSerialize["profile"] = o.Profile

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebraInitParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"module",
		"profile",
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

	varWebraInitParameters := _WebraInitParameters{}

	err = json.Unmarshal(data, &varWebraInitParameters)

	if err != nil {
		return err
	}

	*o = WebraInitParameters(varWebraInitParameters)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorizationMode")
		delete(additionalProperties, "enrollmentMode")
		delete(additionalProperties, "keyType")
		delete(additionalProperties, "module")
		delete(additionalProperties, "passwordPolicy")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebraInitParameters struct {
	value *WebraInitParameters
	isSet bool
}

func (v NullableWebraInitParameters) Get() *WebraInitParameters {
	return v.value
}

func (v *NullableWebraInitParameters) Set(val *WebraInitParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableWebraInitParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableWebraInitParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebraInitParameters(val *WebraInitParameters) *NullableWebraInitParameters {
	return &NullableWebraInitParameters{value: val, isSet: true}
}

func (v NullableWebraInitParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebraInitParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
