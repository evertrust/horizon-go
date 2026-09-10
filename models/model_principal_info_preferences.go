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

// checks if the PrincipalInfoPreferences type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PrincipalInfoPreferences{}

// PrincipalInfoPreferences struct for PrincipalInfoPreferences
type PrincipalInfoPreferences struct {
	// The user's preferred columns on certificate view
	CertificateFields []string `json:"certificateFields,omitempty"`
	// Dark Mode is enabled on UI for this user
	DarkMode utils.NullableBool `json:"darkMode,omitempty"`
	// Expert mode is enabled on UI for this user
	ExpertMode *bool `json:"expertMode,omitempty"`
	// The preferred language of the user
	Lang utils.NullableString `json:"lang,omitempty"`
	// The user's preferred columns on request view
	RequestFields        []string `json:"requestFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalInfoPreferences PrincipalInfoPreferences

// NewPrincipalInfoPreferences instantiates a new PrincipalInfoPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalInfoPreferences() *PrincipalInfoPreferences {
	this := PrincipalInfoPreferences{}
	var darkMode bool = false
	this.DarkMode = *utils.NewNullableBool(&darkMode)
	var expertMode bool = false
	this.ExpertMode = &expertMode
	return &this
}

// NewPrincipalInfoPreferencesWithDefaults instantiates a new PrincipalInfoPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalInfoPreferencesWithDefaults() *PrincipalInfoPreferences {
	this := PrincipalInfoPreferences{}
	var darkMode bool = false
	this.DarkMode = *utils.NewNullableBool(&darkMode)
	var expertMode bool = false
	this.ExpertMode = &expertMode
	return &this
}

// GetCertificateFields returns the CertificateFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoPreferences) GetCertificateFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CertificateFields
}

// GetCertificateFieldsOk returns a tuple with the CertificateFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoPreferences) GetCertificateFieldsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.CertificateFields) {
		return nil, false
	}
	return o.CertificateFields, true
}

// HasCertificateFields returns a boolean if a field has been set.
func (o *PrincipalInfoPreferences) HasCertificateFields() bool {
	if o != nil && !utils.IsNil(o.CertificateFields) {
		return true
	}

	return false
}

// SetCertificateFields gets a reference to the given []string and assigns it to the CertificateFields field.
func (o *PrincipalInfoPreferences) SetCertificateFields(v []string) {
	o.CertificateFields = v
}

// GetDarkMode returns the DarkMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoPreferences) GetDarkMode() bool {
	if o == nil || utils.IsNil(o.DarkMode.Get()) {
		var ret bool
		return ret
	}
	return *o.DarkMode.Get()
}

// GetDarkModeOk returns a tuple with the DarkMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoPreferences) GetDarkModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DarkMode.Get(), o.DarkMode.IsSet()
}

// HasDarkMode returns a boolean if a field has been set.
func (o *PrincipalInfoPreferences) HasDarkMode() bool {
	if o != nil && o.DarkMode.IsSet() {
		return true
	}

	return false
}

// SetDarkMode gets a reference to the given NullableBool and assigns it to the DarkMode field.
func (o *PrincipalInfoPreferences) SetDarkMode(v bool) {
	o.DarkMode.Set(&v)
}

// SetDarkModeNil sets the value for DarkMode to be an explicit nil
func (o *PrincipalInfoPreferences) SetDarkModeNil() {
	o.DarkMode.Set(nil)
}

// UnsetDarkMode ensures that no value is present for DarkMode, not even an explicit nil
func (o *PrincipalInfoPreferences) UnsetDarkMode() {
	o.DarkMode.Unset()
}

// GetExpertMode returns the ExpertMode field value if set, zero value otherwise.
func (o *PrincipalInfoPreferences) GetExpertMode() bool {
	if o == nil || utils.IsNil(o.ExpertMode) {
		var ret bool
		return ret
	}
	return *o.ExpertMode
}

// GetExpertModeOk returns a tuple with the ExpertMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalInfoPreferences) GetExpertModeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ExpertMode) {
		return nil, false
	}
	return o.ExpertMode, true
}

// HasExpertMode returns a boolean if a field has been set.
func (o *PrincipalInfoPreferences) HasExpertMode() bool {
	if o != nil && !utils.IsNil(o.ExpertMode) {
		return true
	}

	return false
}

// SetExpertMode gets a reference to the given bool and assigns it to the ExpertMode field.
func (o *PrincipalInfoPreferences) SetExpertMode(v bool) {
	o.ExpertMode = &v
}

// GetLang returns the Lang field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoPreferences) GetLang() string {
	if o == nil || utils.IsNil(o.Lang.Get()) {
		var ret string
		return ret
	}
	return *o.Lang.Get()
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoPreferences) GetLangOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lang.Get(), o.Lang.IsSet()
}

// HasLang returns a boolean if a field has been set.
func (o *PrincipalInfoPreferences) HasLang() bool {
	if o != nil && o.Lang.IsSet() {
		return true
	}

	return false
}

// SetLang gets a reference to the given NullableString and assigns it to the Lang field.
func (o *PrincipalInfoPreferences) SetLang(v string) {
	o.Lang.Set(&v)
}

// SetLangNil sets the value for Lang to be an explicit nil
func (o *PrincipalInfoPreferences) SetLangNil() {
	o.Lang.Set(nil)
}

// UnsetLang ensures that no value is present for Lang, not even an explicit nil
func (o *PrincipalInfoPreferences) UnsetLang() {
	o.Lang.Unset()
}

// GetRequestFields returns the RequestFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoPreferences) GetRequestFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RequestFields
}

// GetRequestFieldsOk returns a tuple with the RequestFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoPreferences) GetRequestFieldsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.RequestFields) {
		return nil, false
	}
	return o.RequestFields, true
}

// HasRequestFields returns a boolean if a field has been set.
func (o *PrincipalInfoPreferences) HasRequestFields() bool {
	if o != nil && !utils.IsNil(o.RequestFields) {
		return true
	}

	return false
}

// SetRequestFields gets a reference to the given []string and assigns it to the RequestFields field.
func (o *PrincipalInfoPreferences) SetRequestFields(v []string) {
	o.RequestFields = v
}

func (o PrincipalInfoPreferences) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalInfoPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CertificateFields != nil {
		toSerialize["certificateFields"] = o.CertificateFields
	}
	if o.DarkMode.IsSet() {
		toSerialize["darkMode"] = o.DarkMode.Get()
	}
	if !utils.IsNil(o.ExpertMode) {
		toSerialize["expertMode"] = o.ExpertMode
	}
	if o.Lang.IsSet() {
		toSerialize["lang"] = o.Lang.Get()
	}
	if o.RequestFields != nil {
		toSerialize["requestFields"] = o.RequestFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalInfoPreferences) UnmarshalJSON(data []byte) (err error) {
	varPrincipalInfoPreferences := _PrincipalInfoPreferences{}

	err = json.Unmarshal(data, &varPrincipalInfoPreferences)

	if err != nil {
		return err
	}

	*o = PrincipalInfoPreferences(varPrincipalInfoPreferences)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificateFields")
		delete(additionalProperties, "darkMode")
		delete(additionalProperties, "expertMode")
		delete(additionalProperties, "lang")
		delete(additionalProperties, "requestFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalInfoPreferences struct {
	value *PrincipalInfoPreferences
	isSet bool
}

func (v NullablePrincipalInfoPreferences) Get() *PrincipalInfoPreferences {
	return v.value
}

func (v *NullablePrincipalInfoPreferences) Set(val *PrincipalInfoPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalInfoPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalInfoPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalInfoPreferences(val *PrincipalInfoPreferences) *NullablePrincipalInfoPreferences {
	return &NullablePrincipalInfoPreferences{value: val, isSet: true}
}

func (v NullablePrincipalInfoPreferences) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalInfoPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
