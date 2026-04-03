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

// checks if the WebRAImportRequestOnSubmit type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebRAImportRequestOnSubmit{}

// WebRAImportRequestOnSubmit struct for WebRAImportRequestOnSubmit
type WebRAImportRequestOnSubmit struct {
	// The id of the certificate to import
	CertificateId utils.NullableString `json:"certificateId,omitempty"`
	// The PEM encoded certificate to import
	CertificatePem utils.NullableString `json:"certificatePem,omitempty"`
	// The profile name on which to import
	Profile utils.NullableString `json:"profile,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment utils.NullableString `json:"requesterComment,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun utils.NullableBool `json:"dryRun,omitempty"`
	// The module that will be used to process this request. For a WebRA request, this is always `webra`
	Module string `json:"module"`
	// The user-data that will be added on certificate import
	Template *WebRAImportRequestTemplate `json:"template,omitempty"`
	// What this request will do. For an import request, this is always `import`
	Workflow             string `json:"workflow"`
	AdditionalProperties map[string]interface{}
}

type _WebRAImportRequestOnSubmit WebRAImportRequestOnSubmit

// NewWebRAImportRequestOnSubmit instantiates a new WebRAImportRequestOnSubmit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebRAImportRequestOnSubmit(module string, workflow string) *WebRAImportRequestOnSubmit {
	this := WebRAImportRequestOnSubmit{}
	var dryRun bool = false
	this.DryRun = *utils.NewNullableBool(&dryRun)
	this.Module = module
	this.Workflow = workflow
	return &this
}

// NewWebRAImportRequestOnSubmitWithDefaults instantiates a new WebRAImportRequestOnSubmit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebRAImportRequestOnSubmitWithDefaults() *WebRAImportRequestOnSubmit {
	this := WebRAImportRequestOnSubmit{}
	var dryRun bool = false
	this.DryRun = *utils.NewNullableBool(&dryRun)
	return &this
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestOnSubmit) GetCertificateId() string {
	if o == nil || utils.IsNil(o.CertificateId.Get()) {
		var ret string
		return ret
	}
	return *o.CertificateId.Get()
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestOnSubmit) GetCertificateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateId.Get(), o.CertificateId.IsSet()
}

// HasCertificateId returns a boolean if a field has been set.
func (o *WebRAImportRequestOnSubmit) HasCertificateId() bool {
	if o != nil && o.CertificateId.IsSet() {
		return true
	}

	return false
}

// SetCertificateId gets a reference to the given NullableString and assigns it to the CertificateId field.
func (o *WebRAImportRequestOnSubmit) SetCertificateId(v string) {
	o.CertificateId.Set(&v)
}

// SetCertificateIdNil sets the value for CertificateId to be an explicit nil
func (o *WebRAImportRequestOnSubmit) SetCertificateIdNil() {
	o.CertificateId.Set(nil)
}

// UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
func (o *WebRAImportRequestOnSubmit) UnsetCertificateId() {
	o.CertificateId.Unset()
}

// GetCertificatePem returns the CertificatePem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestOnSubmit) GetCertificatePem() string {
	if o == nil || utils.IsNil(o.CertificatePem.Get()) {
		var ret string
		return ret
	}
	return *o.CertificatePem.Get()
}

// GetCertificatePemOk returns a tuple with the CertificatePem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestOnSubmit) GetCertificatePemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificatePem.Get(), o.CertificatePem.IsSet()
}

// HasCertificatePem returns a boolean if a field has been set.
func (o *WebRAImportRequestOnSubmit) HasCertificatePem() bool {
	if o != nil && o.CertificatePem.IsSet() {
		return true
	}

	return false
}

// SetCertificatePem gets a reference to the given NullableString and assigns it to the CertificatePem field.
func (o *WebRAImportRequestOnSubmit) SetCertificatePem(v string) {
	o.CertificatePem.Set(&v)
}

// SetCertificatePemNil sets the value for CertificatePem to be an explicit nil
func (o *WebRAImportRequestOnSubmit) SetCertificatePemNil() {
	o.CertificatePem.Set(nil)
}

// UnsetCertificatePem ensures that no value is present for CertificatePem, not even an explicit nil
func (o *WebRAImportRequestOnSubmit) UnsetCertificatePem() {
	o.CertificatePem.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestOnSubmit) GetProfile() string {
	if o == nil || utils.IsNil(o.Profile.Get()) {
		var ret string
		return ret
	}
	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestOnSubmit) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// HasProfile returns a boolean if a field has been set.
func (o *WebRAImportRequestOnSubmit) HasProfile() bool {
	if o != nil && o.Profile.IsSet() {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NullableString and assigns it to the Profile field.
func (o *WebRAImportRequestOnSubmit) SetProfile(v string) {
	o.Profile.Set(&v)
}

// SetProfileNil sets the value for Profile to be an explicit nil
func (o *WebRAImportRequestOnSubmit) SetProfileNil() {
	o.Profile.Set(nil)
}

// UnsetProfile ensures that no value is present for Profile, not even an explicit nil
func (o *WebRAImportRequestOnSubmit) UnsetProfile() {
	o.Profile.Unset()
}

// GetRequesterComment returns the RequesterComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestOnSubmit) GetRequesterComment() string {
	if o == nil || utils.IsNil(o.RequesterComment.Get()) {
		var ret string
		return ret
	}
	return *o.RequesterComment.Get()
}

// GetRequesterCommentOk returns a tuple with the RequesterComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestOnSubmit) GetRequesterCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterComment.Get(), o.RequesterComment.IsSet()
}

// HasRequesterComment returns a boolean if a field has been set.
func (o *WebRAImportRequestOnSubmit) HasRequesterComment() bool {
	if o != nil && o.RequesterComment.IsSet() {
		return true
	}

	return false
}

// SetRequesterComment gets a reference to the given NullableString and assigns it to the RequesterComment field.
func (o *WebRAImportRequestOnSubmit) SetRequesterComment(v string) {
	o.RequesterComment.Set(&v)
}

// SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil
func (o *WebRAImportRequestOnSubmit) SetRequesterCommentNil() {
	o.RequesterComment.Set(nil)
}

// UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
func (o *WebRAImportRequestOnSubmit) UnsetRequesterComment() {
	o.RequesterComment.Unset()
}

// GetDryRun returns the DryRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestOnSubmit) GetDryRun() bool {
	if o == nil || utils.IsNil(o.DryRun.Get()) {
		var ret bool
		return ret
	}
	return *o.DryRun.Get()
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestOnSubmit) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DryRun.Get(), o.DryRun.IsSet()
}

// HasDryRun returns a boolean if a field has been set.
func (o *WebRAImportRequestOnSubmit) HasDryRun() bool {
	if o != nil && o.DryRun.IsSet() {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given NullableBool and assigns it to the DryRun field.
func (o *WebRAImportRequestOnSubmit) SetDryRun(v bool) {
	o.DryRun.Set(&v)
}

// SetDryRunNil sets the value for DryRun to be an explicit nil
func (o *WebRAImportRequestOnSubmit) SetDryRunNil() {
	o.DryRun.Set(nil)
}

// UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
func (o *WebRAImportRequestOnSubmit) UnsetDryRun() {
	o.DryRun.Unset()
}

// GetModule returns the Module field value
func (o *WebRAImportRequestOnSubmit) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *WebRAImportRequestOnSubmit) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *WebRAImportRequestOnSubmit) SetModule(v string) {
	o.Module = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *WebRAImportRequestOnSubmit) GetTemplate() WebRAImportRequestTemplate {
	if o == nil || utils.IsNil(o.Template) {
		var ret WebRAImportRequestTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRAImportRequestOnSubmit) GetTemplateOk() (*WebRAImportRequestTemplate, bool) {
	if o == nil || utils.IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *WebRAImportRequestOnSubmit) HasTemplate() bool {
	if o != nil && !utils.IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given WebRAImportRequestTemplate and assigns it to the Template field.
func (o *WebRAImportRequestOnSubmit) SetTemplate(v WebRAImportRequestTemplate) {
	o.Template = &v
}

// GetWorkflow returns the Workflow field value
func (o *WebRAImportRequestOnSubmit) GetWorkflow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
func (o *WebRAImportRequestOnSubmit) GetWorkflowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workflow, true
}

// SetWorkflow sets field value
func (o *WebRAImportRequestOnSubmit) SetWorkflow(v string) {
	o.Workflow = v
}

func (o WebRAImportRequestOnSubmit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebRAImportRequestOnSubmit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CertificateId.IsSet() {
		toSerialize["certificateId"] = o.CertificateId.Get()
	}
	if o.CertificatePem.IsSet() {
		toSerialize["certificatePem"] = o.CertificatePem.Get()
	}
	if o.Profile.IsSet() {
		toSerialize["profile"] = o.Profile.Get()
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

func (o *WebRAImportRequestOnSubmit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varWebRAImportRequestOnSubmit := _WebRAImportRequestOnSubmit{}

	err = json.Unmarshal(data, &varWebRAImportRequestOnSubmit)

	if err != nil {
		return err
	}

	*o = WebRAImportRequestOnSubmit(varWebRAImportRequestOnSubmit)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificateId")
		delete(additionalProperties, "certificatePem")
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

type NullableWebRAImportRequestOnSubmit struct {
	value *WebRAImportRequestOnSubmit
	isSet bool
}

func (v NullableWebRAImportRequestOnSubmit) Get() *WebRAImportRequestOnSubmit {
	return v.value
}

func (v *NullableWebRAImportRequestOnSubmit) Set(val *WebRAImportRequestOnSubmit) {
	v.value = val
	v.isSet = true
}

func (v NullableWebRAImportRequestOnSubmit) IsSet() bool {
	return v.isSet
}

func (v *NullableWebRAImportRequestOnSubmit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebRAImportRequestOnSubmit(val *WebRAImportRequestOnSubmit) *NullableWebRAImportRequestOnSubmit {
	return &NullableWebRAImportRequestOnSubmit{value: val, isSet: true}
}

func (v NullableWebRAImportRequestOnSubmit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebRAImportRequestOnSubmit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
