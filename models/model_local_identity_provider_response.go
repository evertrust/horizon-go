/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the LocalIdentityProviderResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LocalIdentityProviderResponse{}

// LocalIdentityProviderResponse struct for LocalIdentityProviderResponse
type LocalIdentityProviderResponse struct {
	// The internal ID of the Identity Provider
	Id string `json:"_id"`
	// The description of the local identity provider
	Description []LocalizedString `json:"description,omitempty"`
	// The display name of the local identity provider
	DisplayName []LocalizedString `json:"displayName,omitempty"`
	// The e-mail template to use for password recovery
	EmailTemplate NullableEmailTemplate `json:"emailTemplate,omitempty"`
	// Whether the local identity provider can be used to identify against Horizon
	Enabled bool `json:"enabled"`
	// Whether the local identity provider can be selected on login to the Horizon UI
	EnabledOnUI bool `json:"enabledOnUI"`
	// The internal name of the local identity provider
	Name string `json:"name"`
	// The password policy to enforce for user passwords on the local identity provider
	PasswordPolicy utils.NullableString `json:"passwordPolicy,omitempty"`
	// The type of identity provider
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _LocalIdentityProviderResponse LocalIdentityProviderResponse

// NewLocalIdentityProviderResponse instantiates a new LocalIdentityProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalIdentityProviderResponse(id string, enabled bool, enabledOnUI bool, name string, type_ string) *LocalIdentityProviderResponse {
	this := LocalIdentityProviderResponse{}
	this.Id = id
	this.Enabled = enabled
	this.EnabledOnUI = enabledOnUI
	this.Name = name
	this.Type = type_
	return &this
}

// NewLocalIdentityProviderResponseWithDefaults instantiates a new LocalIdentityProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalIdentityProviderResponseWithDefaults() *LocalIdentityProviderResponse {
	this := LocalIdentityProviderResponse{}
	return &this
}

// GetId returns the Id field value
func (o *LocalIdentityProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LocalIdentityProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LocalIdentityProviderResponse) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalIdentityProviderResponse) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalIdentityProviderResponse) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LocalIdentityProviderResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *LocalIdentityProviderResponse) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalIdentityProviderResponse) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalIdentityProviderResponse) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *LocalIdentityProviderResponse) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *LocalIdentityProviderResponse) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetEmailTemplate returns the EmailTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalIdentityProviderResponse) GetEmailTemplate() EmailTemplate {
	if o == nil || utils.IsNil(o.EmailTemplate.Get()) {
		var ret EmailTemplate
		return ret
	}
	return *o.EmailTemplate.Get()
}

// GetEmailTemplateOk returns a tuple with the EmailTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalIdentityProviderResponse) GetEmailTemplateOk() (*EmailTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailTemplate.Get(), o.EmailTemplate.IsSet()
}

// HasEmailTemplate returns a boolean if a field has been set.
func (o *LocalIdentityProviderResponse) HasEmailTemplate() bool {
	if o != nil && o.EmailTemplate.IsSet() {
		return true
	}

	return false
}

// SetEmailTemplate gets a reference to the given NullableEmailTemplate and assigns it to the EmailTemplate field.
func (o *LocalIdentityProviderResponse) SetEmailTemplate(v EmailTemplate) {
	o.EmailTemplate.Set(&v)
}

// SetEmailTemplateNil sets the value for EmailTemplate to be an explicit nil
func (o *LocalIdentityProviderResponse) SetEmailTemplateNil() {
	o.EmailTemplate.Set(nil)
}

// UnsetEmailTemplate ensures that no value is present for EmailTemplate, not even an explicit nil
func (o *LocalIdentityProviderResponse) UnsetEmailTemplate() {
	o.EmailTemplate.Unset()
}

// GetEnabled returns the Enabled field value
func (o *LocalIdentityProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LocalIdentityProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *LocalIdentityProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnabledOnUI returns the EnabledOnUI field value
func (o *LocalIdentityProviderResponse) GetEnabledOnUI() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnabledOnUI
}

// GetEnabledOnUIOk returns a tuple with the EnabledOnUI field value
// and a boolean to check if the value has been set.
func (o *LocalIdentityProviderResponse) GetEnabledOnUIOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnabledOnUI, true
}

// SetEnabledOnUI sets field value
func (o *LocalIdentityProviderResponse) SetEnabledOnUI(v bool) {
	o.EnabledOnUI = v
}

// GetName returns the Name field value
func (o *LocalIdentityProviderResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LocalIdentityProviderResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LocalIdentityProviderResponse) SetName(v string) {
	o.Name = v
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalIdentityProviderResponse) GetPasswordPolicy() string {
	if o == nil || utils.IsNil(o.PasswordPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordPolicy.Get()
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalIdentityProviderResponse) GetPasswordPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordPolicy.Get(), o.PasswordPolicy.IsSet()
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *LocalIdentityProviderResponse) HasPasswordPolicy() bool {
	if o != nil && o.PasswordPolicy.IsSet() {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given NullableString and assigns it to the PasswordPolicy field.
func (o *LocalIdentityProviderResponse) SetPasswordPolicy(v string) {
	o.PasswordPolicy.Set(&v)
}

// SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil
func (o *LocalIdentityProviderResponse) SetPasswordPolicyNil() {
	o.PasswordPolicy.Set(nil)
}

// UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
func (o *LocalIdentityProviderResponse) UnsetPasswordPolicy() {
	o.PasswordPolicy.Unset()
}

// GetType returns the Type field value
func (o *LocalIdentityProviderResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LocalIdentityProviderResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LocalIdentityProviderResponse) SetType(v string) {
	o.Type = v
}

func (o LocalIdentityProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalIdentityProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.EmailTemplate.IsSet() {
		toSerialize["emailTemplate"] = o.EmailTemplate.Get()
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["enabledOnUI"] = o.EnabledOnUI
	toSerialize["name"] = o.Name
	if o.PasswordPolicy.IsSet() {
		toSerialize["passwordPolicy"] = o.PasswordPolicy.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LocalIdentityProviderResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"enabled",
		"enabledOnUI",
		"name",
		"type",
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

	varLocalIdentityProviderResponse := _LocalIdentityProviderResponse{}

	err = json.Unmarshal(data, &varLocalIdentityProviderResponse)

	if err != nil {
		return err
	}

	*o = LocalIdentityProviderResponse(varLocalIdentityProviderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "emailTemplate")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "enabledOnUI")
		delete(additionalProperties, "name")
		delete(additionalProperties, "passwordPolicy")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLocalIdentityProviderResponse struct {
	value *LocalIdentityProviderResponse
	isSet bool
}

func (v NullableLocalIdentityProviderResponse) Get() *LocalIdentityProviderResponse {
	return v.value
}

func (v *NullableLocalIdentityProviderResponse) Set(val *LocalIdentityProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalIdentityProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalIdentityProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalIdentityProviderResponse(val *LocalIdentityProviderResponse) *NullableLocalIdentityProviderResponse {
	return &NullableLocalIdentityProviderResponse{value: val, isSet: true}
}

func (v NullableLocalIdentityProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalIdentityProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
