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

// checks if the ScepEnrollRequestOnSubmit type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ScepEnrollRequestOnSubmit{}

// ScepEnrollRequestOnSubmit struct for ScepEnrollRequestOnSubmit
type ScepEnrollRequestOnSubmit struct {
	// Fill the DN if DN whitelist is enabled. Contains the DN of the challenge
	Dn interface{} `json:"dn,omitempty"`
	// The password of the challenge. Must be set if password mode is `manual`
	Password *SecretString `json:"password,omitempty"`
	// The SCEP profile name
	Profile interface{} `json:"profile"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment utils.NullableString `json:"requesterComment,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun utils.NullableBool `json:"dryRun,omitempty"`
	// The module that will be used to process this request. For a SCEP request, this is always `scep`
	Module string `json:"module"`
	// The user-data that will be used to generate the challenge
	Template *ScepEnrollRequestTemplate `json:"template,omitempty"`
	// What this request will do. For an enrollment request, this is always `enroll`
	Workflow             string `json:"workflow"`
	AdditionalProperties map[string]interface{}
}

type _ScepEnrollRequestOnSubmit ScepEnrollRequestOnSubmit

// NewScepEnrollRequestOnSubmit instantiates a new ScepEnrollRequestOnSubmit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScepEnrollRequestOnSubmit(profile interface{}, module string, workflow string) *ScepEnrollRequestOnSubmit {
	this := ScepEnrollRequestOnSubmit{}
	var dryRun bool = false
	this.DryRun = *utils.NewNullableBool(&dryRun)
	this.Module = module
	this.Workflow = workflow
	return &this
}

// NewScepEnrollRequestOnSubmitWithDefaults instantiates a new ScepEnrollRequestOnSubmit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScepEnrollRequestOnSubmitWithDefaults() *ScepEnrollRequestOnSubmit {
	this := ScepEnrollRequestOnSubmit{}
	var dryRun bool = false
	this.DryRun = *utils.NewNullableBool(&dryRun)
	return &this
}

// GetDn returns the Dn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScepEnrollRequestOnSubmit) GetDn() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScepEnrollRequestOnSubmit) GetDnOk() (*interface{}, bool) {
	if o == nil || utils.IsNil(o.Dn) {
		return nil, false
	}
	return &o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *ScepEnrollRequestOnSubmit) HasDn() bool {
	if o != nil && !utils.IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given interface{} and assigns it to the Dn field.
func (o *ScepEnrollRequestOnSubmit) SetDn(v interface{}) {
	o.Dn = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ScepEnrollRequestOnSubmit) GetPassword() SecretString {
	if o == nil || utils.IsNil(o.Password) {
		var ret SecretString
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScepEnrollRequestOnSubmit) GetPasswordOk() (*SecretString, bool) {
	if o == nil || utils.IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ScepEnrollRequestOnSubmit) HasPassword() bool {
	if o != nil && !utils.IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given SecretString and assigns it to the Password field.
func (o *ScepEnrollRequestOnSubmit) SetPassword(v SecretString) {
	o.Password = &v
}

// GetProfile returns the Profile field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ScepEnrollRequestOnSubmit) GetProfile() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScepEnrollRequestOnSubmit) GetProfileOk() (*interface{}, bool) {
	if o == nil || utils.IsNil(o.Profile) {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *ScepEnrollRequestOnSubmit) SetProfile(v interface{}) {
	o.Profile = v
}

// GetRequesterComment returns the RequesterComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScepEnrollRequestOnSubmit) GetRequesterComment() string {
	if o == nil || utils.IsNil(o.RequesterComment.Get()) {
		var ret string
		return ret
	}
	return *o.RequesterComment.Get()
}

// GetRequesterCommentOk returns a tuple with the RequesterComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScepEnrollRequestOnSubmit) GetRequesterCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterComment.Get(), o.RequesterComment.IsSet()
}

// HasRequesterComment returns a boolean if a field has been set.
func (o *ScepEnrollRequestOnSubmit) HasRequesterComment() bool {
	if o != nil && o.RequesterComment.IsSet() {
		return true
	}

	return false
}

// SetRequesterComment gets a reference to the given NullableString and assigns it to the RequesterComment field.
func (o *ScepEnrollRequestOnSubmit) SetRequesterComment(v string) {
	o.RequesterComment.Set(&v)
}

// SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil
func (o *ScepEnrollRequestOnSubmit) SetRequesterCommentNil() {
	o.RequesterComment.Set(nil)
}

// UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
func (o *ScepEnrollRequestOnSubmit) UnsetRequesterComment() {
	o.RequesterComment.Unset()
}

// GetDryRun returns the DryRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScepEnrollRequestOnSubmit) GetDryRun() bool {
	if o == nil || utils.IsNil(o.DryRun.Get()) {
		var ret bool
		return ret
	}
	return *o.DryRun.Get()
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScepEnrollRequestOnSubmit) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DryRun.Get(), o.DryRun.IsSet()
}

// HasDryRun returns a boolean if a field has been set.
func (o *ScepEnrollRequestOnSubmit) HasDryRun() bool {
	if o != nil && o.DryRun.IsSet() {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given NullableBool and assigns it to the DryRun field.
func (o *ScepEnrollRequestOnSubmit) SetDryRun(v bool) {
	o.DryRun.Set(&v)
}

// SetDryRunNil sets the value for DryRun to be an explicit nil
func (o *ScepEnrollRequestOnSubmit) SetDryRunNil() {
	o.DryRun.Set(nil)
}

// UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
func (o *ScepEnrollRequestOnSubmit) UnsetDryRun() {
	o.DryRun.Unset()
}

// GetModule returns the Module field value
func (o *ScepEnrollRequestOnSubmit) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *ScepEnrollRequestOnSubmit) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *ScepEnrollRequestOnSubmit) SetModule(v string) {
	o.Module = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *ScepEnrollRequestOnSubmit) GetTemplate() ScepEnrollRequestTemplate {
	if o == nil || utils.IsNil(o.Template) {
		var ret ScepEnrollRequestTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScepEnrollRequestOnSubmit) GetTemplateOk() (*ScepEnrollRequestTemplate, bool) {
	if o == nil || utils.IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *ScepEnrollRequestOnSubmit) HasTemplate() bool {
	if o != nil && !utils.IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given ScepEnrollRequestTemplate and assigns it to the Template field.
func (o *ScepEnrollRequestOnSubmit) SetTemplate(v ScepEnrollRequestTemplate) {
	o.Template = &v
}

// GetWorkflow returns the Workflow field value
func (o *ScepEnrollRequestOnSubmit) GetWorkflow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
func (o *ScepEnrollRequestOnSubmit) GetWorkflowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workflow, true
}

// SetWorkflow sets field value
func (o *ScepEnrollRequestOnSubmit) SetWorkflow(v string) {
	o.Workflow = v
}

func (o ScepEnrollRequestOnSubmit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScepEnrollRequestOnSubmit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Dn != nil {
		toSerialize["dn"] = o.Dn
	}
	if !utils.IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.RequesterComment.IsSet() {
		toSerialize["requesterComment"] = o.RequesterComment.Get()
	}
	if o.DryRun.IsSet() {
		toSerialize["dryRun"] = o.DryRun.Get()
	}
	toSerialize["module"] = o.Module
	if !utils.IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	toSerialize["workflow"] = o.Workflow

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScepEnrollRequestOnSubmit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profile",
		"module",
		"workflow",
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

	varScepEnrollRequestOnSubmit := _ScepEnrollRequestOnSubmit{}

	err = json.Unmarshal(data, &varScepEnrollRequestOnSubmit)

	if err != nil {
		return err
	}

	*o = ScepEnrollRequestOnSubmit(varScepEnrollRequestOnSubmit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dn")
		delete(additionalProperties, "password")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "requesterComment")
		delete(additionalProperties, "dryRun")
		delete(additionalProperties, "module")
		delete(additionalProperties, "template")
		delete(additionalProperties, "workflow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScepEnrollRequestOnSubmit struct {
	value *ScepEnrollRequestOnSubmit
	isSet bool
}

func (v NullableScepEnrollRequestOnSubmit) Get() *ScepEnrollRequestOnSubmit {
	return v.value
}

func (v *NullableScepEnrollRequestOnSubmit) Set(val *ScepEnrollRequestOnSubmit) {
	v.value = val
	v.isSet = true
}

func (v NullableScepEnrollRequestOnSubmit) IsSet() bool {
	return v.isSet
}

func (v *NullableScepEnrollRequestOnSubmit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScepEnrollRequestOnSubmit(val *ScepEnrollRequestOnSubmit) *NullableScepEnrollRequestOnSubmit {
	return &NullableScepEnrollRequestOnSubmit{value: val, isSet: true}
}

func (v NullableScepEnrollRequestOnSubmit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScepEnrollRequestOnSubmit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
