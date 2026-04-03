/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the AcmeExternalInitParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AcmeExternalInitParameters{}

// AcmeExternalInitParameters struct for AcmeExternalInitParameters
type AcmeExternalInitParameters struct {
	// The ACME URL for the external ACME server.
	AcmeUrl *string `json:"acmeUrl,omitempty"`
	// The allowed authorization methods for ACME external.
	AllowedAuthorizationMethod []string `json:"allowedAuthorizationMethod,omitempty"`
	// The key type used for ACME external.
	KeyType *string `json:"keyType,omitempty"`
	// The module of the initialization parameters.
	Module string `json:"module"`
	// The profile used for ACME external.
	Profile string `json:"profile"`
	// Indicates whether EAB is required for ACME external.
	RequireEAB           *bool `json:"requireEAB,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcmeExternalInitParameters AcmeExternalInitParameters

// NewAcmeExternalInitParameters instantiates a new AcmeExternalInitParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcmeExternalInitParameters(module string, profile string) *AcmeExternalInitParameters {
	this := AcmeExternalInitParameters{}
	this.Module = module
	this.Profile = profile
	return &this
}

// NewAcmeExternalInitParametersWithDefaults instantiates a new AcmeExternalInitParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcmeExternalInitParametersWithDefaults() *AcmeExternalInitParameters {
	this := AcmeExternalInitParameters{}
	return &this
}

// GetAcmeUrl returns the AcmeUrl field value if set, zero value otherwise.
func (o *AcmeExternalInitParameters) GetAcmeUrl() string {
	if o == nil || utils.IsNil(o.AcmeUrl) {
		var ret string
		return ret
	}
	return *o.AcmeUrl
}

// GetAcmeUrlOk returns a tuple with the AcmeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeExternalInitParameters) GetAcmeUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AcmeUrl) {
		return nil, false
	}
	return o.AcmeUrl, true
}

// HasAcmeUrl returns a boolean if a field has been set.
func (o *AcmeExternalInitParameters) HasAcmeUrl() bool {
	if o != nil && !utils.IsNil(o.AcmeUrl) {
		return true
	}

	return false
}

// SetAcmeUrl gets a reference to the given string and assigns it to the AcmeUrl field.
func (o *AcmeExternalInitParameters) SetAcmeUrl(v string) {
	o.AcmeUrl = &v
}

// GetAllowedAuthorizationMethod returns the AllowedAuthorizationMethod field value if set, zero value otherwise.
func (o *AcmeExternalInitParameters) GetAllowedAuthorizationMethod() []string {
	if o == nil || utils.IsNil(o.AllowedAuthorizationMethod) {
		var ret []string
		return ret
	}
	return o.AllowedAuthorizationMethod
}

// GetAllowedAuthorizationMethodOk returns a tuple with the AllowedAuthorizationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeExternalInitParameters) GetAllowedAuthorizationMethodOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AllowedAuthorizationMethod) {
		return nil, false
	}
	return o.AllowedAuthorizationMethod, true
}

// HasAllowedAuthorizationMethod returns a boolean if a field has been set.
func (o *AcmeExternalInitParameters) HasAllowedAuthorizationMethod() bool {
	if o != nil && !utils.IsNil(o.AllowedAuthorizationMethod) {
		return true
	}

	return false
}

// SetAllowedAuthorizationMethod gets a reference to the given []string and assigns it to the AllowedAuthorizationMethod field.
func (o *AcmeExternalInitParameters) SetAllowedAuthorizationMethod(v []string) {
	o.AllowedAuthorizationMethod = v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *AcmeExternalInitParameters) GetKeyType() string {
	if o == nil || utils.IsNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeExternalInitParameters) GetKeyTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *AcmeExternalInitParameters) HasKeyType() bool {
	if o != nil && !utils.IsNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *AcmeExternalInitParameters) SetKeyType(v string) {
	o.KeyType = &v
}

// GetModule returns the Module field value
func (o *AcmeExternalInitParameters) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalInitParameters) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *AcmeExternalInitParameters) SetModule(v string) {
	o.Module = v
}

// GetProfile returns the Profile field value
func (o *AcmeExternalInitParameters) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalInitParameters) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *AcmeExternalInitParameters) SetProfile(v string) {
	o.Profile = v
}

// GetRequireEAB returns the RequireEAB field value if set, zero value otherwise.
func (o *AcmeExternalInitParameters) GetRequireEAB() bool {
	if o == nil || utils.IsNil(o.RequireEAB) {
		var ret bool
		return ret
	}
	return *o.RequireEAB
}

// GetRequireEABOk returns a tuple with the RequireEAB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeExternalInitParameters) GetRequireEABOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RequireEAB) {
		return nil, false
	}
	return o.RequireEAB, true
}

// HasRequireEAB returns a boolean if a field has been set.
func (o *AcmeExternalInitParameters) HasRequireEAB() bool {
	if o != nil && !utils.IsNil(o.RequireEAB) {
		return true
	}

	return false
}

// SetRequireEAB gets a reference to the given bool and assigns it to the RequireEAB field.
func (o *AcmeExternalInitParameters) SetRequireEAB(v bool) {
	o.RequireEAB = &v
}

func (o AcmeExternalInitParameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcmeExternalInitParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AcmeUrl) {
		toSerialize["acmeUrl"] = o.AcmeUrl
	}
	if !utils.IsNil(o.AllowedAuthorizationMethod) {
		toSerialize["allowedAuthorizationMethod"] = o.AllowedAuthorizationMethod
	}
	if !utils.IsNil(o.KeyType) {
		toSerialize["keyType"] = o.KeyType
	}
	toSerialize["module"] = o.Module
	toSerialize["profile"] = o.Profile
	if !utils.IsNil(o.RequireEAB) {
		toSerialize["requireEAB"] = o.RequireEAB
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcmeExternalInitParameters) UnmarshalJSON(data []byte) (err error) {
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

	varAcmeExternalInitParameters := _AcmeExternalInitParameters{}

	err = json.Unmarshal(data, &varAcmeExternalInitParameters)

	if err != nil {
		return err
	}

	*o = AcmeExternalInitParameters(varAcmeExternalInitParameters)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "acmeUrl")
		delete(additionalProperties, "allowedAuthorizationMethod")
		delete(additionalProperties, "keyType")
		delete(additionalProperties, "module")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "requireEAB")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcmeExternalInitParameters struct {
	value *AcmeExternalInitParameters
	isSet bool
}

func (v NullableAcmeExternalInitParameters) Get() *AcmeExternalInitParameters {
	return v.value
}

func (v *NullableAcmeExternalInitParameters) Set(val *AcmeExternalInitParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableAcmeExternalInitParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableAcmeExternalInitParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcmeExternalInitParameters(val *AcmeExternalInitParameters) *NullableAcmeExternalInitParameters {
	return &NullableAcmeExternalInitParameters{value: val, isSet: true}
}

func (v NullableAcmeExternalInitParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcmeExternalInitParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
