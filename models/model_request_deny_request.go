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

// checks if the RequestDenyRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RequestDenyRequest{}

// RequestDenyRequest struct for RequestDenyRequest
type RequestDenyRequest struct {
	// The ID of the request to deny
	Id       string   `json:"_id"`
	Module   Module   `json:"module"`
	Workflow Workflow `json:"workflow"`
	// Free-text field editable by the approver to provider more context on the denial
	ApproverComment      *string `json:"approverComment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestDenyRequest RequestDenyRequest

// NewRequestDenyRequest instantiates a new RequestDenyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestDenyRequest(id string, module Module, workflow Workflow) *RequestDenyRequest {
	this := RequestDenyRequest{}
	this.Id = id
	this.Module = module
	this.Workflow = workflow
	return &this
}

// NewRequestDenyRequestWithDefaults instantiates a new RequestDenyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestDenyRequestWithDefaults() *RequestDenyRequest {
	this := RequestDenyRequest{}
	return &this
}

// GetId returns the Id field value
func (o *RequestDenyRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestDenyRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestDenyRequest) SetId(v string) {
	o.Id = v
}

// GetModule returns the Module field value
func (o *RequestDenyRequest) GetModule() Module {
	if o == nil {
		var ret Module
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *RequestDenyRequest) GetModuleOk() (*Module, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *RequestDenyRequest) SetModule(v Module) {
	o.Module = v
}

// GetWorkflow returns the Workflow field value
func (o *RequestDenyRequest) GetWorkflow() Workflow {
	if o == nil {
		var ret Workflow
		return ret
	}

	return o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
func (o *RequestDenyRequest) GetWorkflowOk() (*Workflow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workflow, true
}

// SetWorkflow sets field value
func (o *RequestDenyRequest) SetWorkflow(v Workflow) {
	o.Workflow = v
}

// GetApproverComment returns the ApproverComment field value if set, zero value otherwise.
func (o *RequestDenyRequest) GetApproverComment() string {
	if o == nil || utils.IsNil(o.ApproverComment) {
		var ret string
		return ret
	}
	return *o.ApproverComment
}

// GetApproverCommentOk returns a tuple with the ApproverComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestDenyRequest) GetApproverCommentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ApproverComment) {
		return nil, false
	}
	return o.ApproverComment, true
}

// HasApproverComment returns a boolean if a field has been set.
func (o *RequestDenyRequest) HasApproverComment() bool {
	if o != nil && !utils.IsNil(o.ApproverComment) {
		return true
	}

	return false
}

// SetApproverComment gets a reference to the given string and assigns it to the ApproverComment field.
func (o *RequestDenyRequest) SetApproverComment(v string) {
	o.ApproverComment = &v
}

func (o RequestDenyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestDenyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["module"] = o.Module
	toSerialize["workflow"] = o.Workflow
	if !utils.IsNil(o.ApproverComment) {
		toSerialize["approverComment"] = o.ApproverComment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestDenyRequest) UnmarshalJSON(data []byte) (err error) {
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

	varRequestDenyRequest := _RequestDenyRequest{}

	err = json.Unmarshal(data, &varRequestDenyRequest)

	if err != nil {
		return err
	}

	*o = RequestDenyRequest(varRequestDenyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "module")
		delete(additionalProperties, "workflow")
		delete(additionalProperties, "approverComment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestDenyRequest struct {
	value *RequestDenyRequest
	isSet bool
}

func (v NullableRequestDenyRequest) Get() *RequestDenyRequest {
	return v.value
}

func (v *NullableRequestDenyRequest) Set(val *RequestDenyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestDenyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestDenyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestDenyRequest(val *RequestDenyRequest) *NullableRequestDenyRequest {
	return &NullableRequestDenyRequest{value: val, isSet: true}
}

func (v NullableRequestDenyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestDenyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
