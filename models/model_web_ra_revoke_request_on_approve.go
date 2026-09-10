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

// checks if the WebRARevokeRequestOnApprove type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebRARevokeRequestOnApprove{}

// WebRARevokeRequestOnApprove struct for WebRARevokeRequestOnApprove
type WebRARevokeRequestOnApprove struct {
	// Object internal ID
	Id string `json:"_id"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment utils.NullableString `json:"approverComment,omitempty"`
	// The module on which the revocation occurred
	Module string `json:"module"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun utils.NullableBool `json:"dryRun,omitempty"`
	// The user-data that will be used to revoke the certificate
	Template *WebRARevokeRequestTemplate `json:"template,omitempty"`
	// What this request will do. For a revocation request, this is always `revoke`
	Workflow             string `json:"workflow"`
	AdditionalProperties map[string]interface{}
}

type _WebRARevokeRequestOnApprove WebRARevokeRequestOnApprove

// NewWebRARevokeRequestOnApprove instantiates a new WebRARevokeRequestOnApprove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebRARevokeRequestOnApprove(id string, module string, workflow string) *WebRARevokeRequestOnApprove {
	this := WebRARevokeRequestOnApprove{}
	this.Id = id
	this.Module = module
	this.Workflow = workflow
	var dryRun bool = false
	this.DryRun = *utils.NewNullableBool(&dryRun)
	return &this
}

// NewWebRARevokeRequestOnApproveWithDefaults instantiates a new WebRARevokeRequestOnApprove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebRARevokeRequestOnApproveWithDefaults() *WebRARevokeRequestOnApprove {
	this := WebRARevokeRequestOnApprove{}
	var dryRun bool = false
	this.DryRun = *utils.NewNullableBool(&dryRun)
	return &this
}

// GetId returns the Id field value
func (o *WebRARevokeRequestOnApprove) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebRARevokeRequestOnApprove) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebRARevokeRequestOnApprove) SetId(v string) {
	o.Id = v
}

// GetApproverComment returns the ApproverComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARevokeRequestOnApprove) GetApproverComment() string {
	if o == nil || utils.IsNil(o.ApproverComment.Get()) {
		var ret string
		return ret
	}
	return *o.ApproverComment.Get()
}

// GetApproverCommentOk returns a tuple with the ApproverComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARevokeRequestOnApprove) GetApproverCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproverComment.Get(), o.ApproverComment.IsSet()
}

// HasApproverComment returns a boolean if a field has been set.
func (o *WebRARevokeRequestOnApprove) HasApproverComment() bool {
	if o != nil && o.ApproverComment.IsSet() {
		return true
	}

	return false
}

// SetApproverComment gets a reference to the given NullableString and assigns it to the ApproverComment field.
func (o *WebRARevokeRequestOnApprove) SetApproverComment(v string) {
	o.ApproverComment.Set(&v)
}

// SetApproverCommentNil sets the value for ApproverComment to be an explicit nil
func (o *WebRARevokeRequestOnApprove) SetApproverCommentNil() {
	o.ApproverComment.Set(nil)
}

// UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
func (o *WebRARevokeRequestOnApprove) UnsetApproverComment() {
	o.ApproverComment.Unset()
}

// GetModule returns the Module field value
func (o *WebRARevokeRequestOnApprove) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *WebRARevokeRequestOnApprove) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *WebRARevokeRequestOnApprove) SetModule(v string) {
	o.Module = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARevokeRequestOnApprove) GetDryRun() bool {
	if o == nil || utils.IsNil(o.DryRun.Get()) {
		var ret bool
		return ret
	}
	return *o.DryRun.Get()
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARevokeRequestOnApprove) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DryRun.Get(), o.DryRun.IsSet()
}

// HasDryRun returns a boolean if a field has been set.
func (o *WebRARevokeRequestOnApprove) HasDryRun() bool {
	if o != nil && o.DryRun.IsSet() {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given NullableBool and assigns it to the DryRun field.
func (o *WebRARevokeRequestOnApprove) SetDryRun(v bool) {
	o.DryRun.Set(&v)
}

// SetDryRunNil sets the value for DryRun to be an explicit nil
func (o *WebRARevokeRequestOnApprove) SetDryRunNil() {
	o.DryRun.Set(nil)
}

// UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
func (o *WebRARevokeRequestOnApprove) UnsetDryRun() {
	o.DryRun.Unset()
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *WebRARevokeRequestOnApprove) GetTemplate() WebRARevokeRequestTemplate {
	if o == nil || utils.IsNil(o.Template) {
		var ret WebRARevokeRequestTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRARevokeRequestOnApprove) GetTemplateOk() (*WebRARevokeRequestTemplate, bool) {
	if o == nil || utils.IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *WebRARevokeRequestOnApprove) HasTemplate() bool {
	if o != nil && !utils.IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given WebRARevokeRequestTemplate and assigns it to the Template field.
func (o *WebRARevokeRequestOnApprove) SetTemplate(v WebRARevokeRequestTemplate) {
	o.Template = &v
}

// GetWorkflow returns the Workflow field value
func (o *WebRARevokeRequestOnApprove) GetWorkflow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
func (o *WebRARevokeRequestOnApprove) GetWorkflowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workflow, true
}

// SetWorkflow sets field value
func (o *WebRARevokeRequestOnApprove) SetWorkflow(v string) {
	o.Workflow = v
}

func (o WebRARevokeRequestOnApprove) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebRARevokeRequestOnApprove) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.ApproverComment.IsSet() {
		toSerialize["approverComment"] = o.ApproverComment.Get()
	}
	toSerialize["module"] = o.Module
	if o.DryRun.IsSet() {
		toSerialize["dryRun"] = o.DryRun.Get()
	}
	if !utils.IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	toSerialize["workflow"] = o.Workflow

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebRARevokeRequestOnApprove) UnmarshalJSON(data []byte) (err error) {
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
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebRARevokeRequestOnApprove := _WebRARevokeRequestOnApprove{}

	err = json.Unmarshal(data, &varWebRARevokeRequestOnApprove)

	if err != nil {
		return err
	}

	*o = WebRARevokeRequestOnApprove(varWebRARevokeRequestOnApprove)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "approverComment")
		delete(additionalProperties, "module")
		delete(additionalProperties, "dryRun")
		delete(additionalProperties, "template")
		delete(additionalProperties, "workflow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebRARevokeRequestOnApprove struct {
	value *WebRARevokeRequestOnApprove
	isSet bool
}

func (v NullableWebRARevokeRequestOnApprove) Get() *WebRARevokeRequestOnApprove {
	return v.value
}

func (v *NullableWebRARevokeRequestOnApprove) Set(val *WebRARevokeRequestOnApprove) {
	v.value = val
	v.isSet = true
}

func (v NullableWebRARevokeRequestOnApprove) IsSet() bool {
	return v.isSet
}

func (v *NullableWebRARevokeRequestOnApprove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebRARevokeRequestOnApprove(val *WebRARevokeRequestOnApprove) *NullableWebRARevokeRequestOnApprove {
	return &NullableWebRARevokeRequestOnApprove{value: val, isSet: true}
}

func (v NullableWebRARevokeRequestOnApprove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebRARevokeRequestOnApprove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
