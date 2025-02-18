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

// checks if the EstEnrollRequestOnApprove type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstEnrollRequestOnApprove{}

// EstEnrollRequestOnApprove struct for EstEnrollRequestOnApprove
type EstEnrollRequestOnApprove struct {
	// Object internal ID
	Id string `json:"_id"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment NullableString `json:"approverComment,omitempty"`
	// The module that will be used to process this request. For an EST request, this is always `est`
	Module string `json:"module"`
	// What this request will do. For an enrollment request, this is always `enroll`
	Workflow string `json:"workflow"`
	// The user-data that will be used to generate the certificate
	Template *EstEnrollRequestTemplate `json:"template,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun NullableBool `json:"dryRun,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EstEnrollRequestOnApprove EstEnrollRequestOnApprove

// NewEstEnrollRequestOnApprove instantiates a new EstEnrollRequestOnApprove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstEnrollRequestOnApprove(id string, module string, workflow string) *EstEnrollRequestOnApprove {
	this := EstEnrollRequestOnApprove{}
	this.Module = module
	this.Workflow = workflow
	var dryRun bool = false
	this.DryRun = *NewNullableBool(&dryRun)
	return &this
}

// NewEstEnrollRequestOnApproveWithDefaults instantiates a new EstEnrollRequestOnApprove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstEnrollRequestOnApproveWithDefaults() *EstEnrollRequestOnApprove {
	this := EstEnrollRequestOnApprove{}
	var dryRun bool = false
	this.DryRun = *NewNullableBool(&dryRun)
	return &this
}

// GetId returns the Id field value
func (o *EstEnrollRequestOnApprove) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EstEnrollRequestOnApprove) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EstEnrollRequestOnApprove) SetId(v string) {
	o.Id = v
}

// GetApproverComment returns the ApproverComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstEnrollRequestOnApprove) GetApproverComment() string {
	if o == nil || IsNil(o.ApproverComment.Get()) {
		var ret string
		return ret
	}
	return *o.ApproverComment.Get()
}

// GetApproverCommentOk returns a tuple with the ApproverComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstEnrollRequestOnApprove) GetApproverCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproverComment.Get(), o.ApproverComment.IsSet()
}

// HasApproverComment returns a boolean if a field has been set.
func (o *EstEnrollRequestOnApprove) HasApproverComment() bool {
	if o != nil && o.ApproverComment.IsSet() {
		return true
	}

	return false
}

// SetApproverComment gets a reference to the given NullableString and assigns it to the ApproverComment field.
func (o *EstEnrollRequestOnApprove) SetApproverComment(v string) {
	o.ApproverComment.Set(&v)
}
// SetApproverCommentNil sets the value for ApproverComment to be an explicit nil
func (o *EstEnrollRequestOnApprove) SetApproverCommentNil() {
	o.ApproverComment.Set(nil)
}

// UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
func (o *EstEnrollRequestOnApprove) UnsetApproverComment() {
	o.ApproverComment.Unset()
}

// GetModule returns the Module field value
func (o *EstEnrollRequestOnApprove) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *EstEnrollRequestOnApprove) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *EstEnrollRequestOnApprove) SetModule(v string) {
	o.Module = v
}

// GetWorkflow returns the Workflow field value
func (o *EstEnrollRequestOnApprove) GetWorkflow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
func (o *EstEnrollRequestOnApprove) GetWorkflowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workflow, true
}

// SetWorkflow sets field value
func (o *EstEnrollRequestOnApprove) SetWorkflow(v string) {
	o.Workflow = v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *EstEnrollRequestOnApprove) GetTemplate() EstEnrollRequestTemplate {
	if o == nil || IsNil(o.Template) {
		var ret EstEnrollRequestTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstEnrollRequestOnApprove) GetTemplateOk() (*EstEnrollRequestTemplate, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *EstEnrollRequestOnApprove) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given EstEnrollRequestTemplate and assigns it to the Template field.
func (o *EstEnrollRequestOnApprove) SetTemplate(v EstEnrollRequestTemplate) {
	o.Template = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstEnrollRequestOnApprove) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun.Get()) {
		var ret bool
		return ret
	}
	return *o.DryRun.Get()
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstEnrollRequestOnApprove) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DryRun.Get(), o.DryRun.IsSet()
}

// HasDryRun returns a boolean if a field has been set.
func (o *EstEnrollRequestOnApprove) HasDryRun() bool {
	if o != nil && o.DryRun.IsSet() {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given NullableBool and assigns it to the DryRun field.
func (o *EstEnrollRequestOnApprove) SetDryRun(v bool) {
	o.DryRun.Set(&v)
}
// SetDryRunNil sets the value for DryRun to be an explicit nil
func (o *EstEnrollRequestOnApprove) SetDryRunNil() {
	o.DryRun.Set(nil)
}

// UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
func (o *EstEnrollRequestOnApprove) UnsetDryRun() {
	o.DryRun.Unset()
}

func (o EstEnrollRequestOnApprove) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstEnrollRequestOnApprove) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.ApproverComment.IsSet() {
		toSerialize["approverComment"] = o.ApproverComment.Get()
	}
	toSerialize["module"] = o.Module
	toSerialize["workflow"] = o.Workflow
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if o.DryRun.IsSet() {
		toSerialize["dryRun"] = o.DryRun.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EstEnrollRequestOnApprove) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"module",
		"workflow",
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

	varEstEnrollRequestOnApprove := _EstEnrollRequestOnApprove{}

	err = json.Unmarshal(data, &varEstEnrollRequestOnApprove)

	if err != nil {
		return err
	}

	*o = EstEnrollRequestOnApprove(varEstEnrollRequestOnApprove)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "approverComment")
		delete(additionalProperties, "module")
		delete(additionalProperties, "workflow")
		delete(additionalProperties, "template")
		delete(additionalProperties, "dryRun")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEstEnrollRequestOnApprove struct {
	value *EstEnrollRequestOnApprove
	isSet bool
}

func (v NullableEstEnrollRequestOnApprove) Get() *EstEnrollRequestOnApprove {
	return v.value
}

func (v *NullableEstEnrollRequestOnApprove) Set(val *EstEnrollRequestOnApprove) {
	v.value = val
	v.isSet = true
}

func (v NullableEstEnrollRequestOnApprove) IsSet() bool {
	return v.isSet
}

func (v *NullableEstEnrollRequestOnApprove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstEnrollRequestOnApprove(val *EstEnrollRequestOnApprove) *NullableEstEnrollRequestOnApprove {
	return &NullableEstEnrollRequestOnApprove{value: val, isSet: true}
}

func (v NullableEstEnrollRequestOnApprove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstEnrollRequestOnApprove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


