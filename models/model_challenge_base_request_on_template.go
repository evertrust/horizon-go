/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the ChallengeBaseRequestOnTemplate type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ChallengeBaseRequestOnTemplate{}

// ChallengeBaseRequestOnTemplate struct for ChallengeBaseRequestOnTemplate
type ChallengeBaseRequestOnTemplate struct {
	// The request module
	Module *string `json:"module,omitempty"`
	// The profile for which to return the template.
	Profile *string `json:"profile,omitempty"`
	// The workflow for which to return the template.
	Workflow             *string `json:"workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChallengeBaseRequestOnTemplate ChallengeBaseRequestOnTemplate

// NewChallengeBaseRequestOnTemplate instantiates a new ChallengeBaseRequestOnTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChallengeBaseRequestOnTemplate() *ChallengeBaseRequestOnTemplate {
	this := ChallengeBaseRequestOnTemplate{}
	return &this
}

// NewChallengeBaseRequestOnTemplateWithDefaults instantiates a new ChallengeBaseRequestOnTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChallengeBaseRequestOnTemplateWithDefaults() *ChallengeBaseRequestOnTemplate {
	this := ChallengeBaseRequestOnTemplate{}
	return &this
}

// GetModule returns the Module field value if set, zero value otherwise.
func (o *ChallengeBaseRequestOnTemplate) GetModule() string {
	if o == nil || utils.IsNil(o.Module) {
		var ret string
		return ret
	}
	return *o.Module
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeBaseRequestOnTemplate) GetModuleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Module) {
		return nil, false
	}
	return o.Module, true
}

// HasModule returns a boolean if a field has been set.
func (o *ChallengeBaseRequestOnTemplate) HasModule() bool {
	if o != nil && !utils.IsNil(o.Module) {
		return true
	}

	return false
}

// SetModule gets a reference to the given string and assigns it to the Module field.
func (o *ChallengeBaseRequestOnTemplate) SetModule(v string) {
	o.Module = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ChallengeBaseRequestOnTemplate) GetProfile() string {
	if o == nil || utils.IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeBaseRequestOnTemplate) GetProfileOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ChallengeBaseRequestOnTemplate) HasProfile() bool {
	if o != nil && !utils.IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *ChallengeBaseRequestOnTemplate) SetProfile(v string) {
	o.Profile = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *ChallengeBaseRequestOnTemplate) GetWorkflow() string {
	if o == nil || utils.IsNil(o.Workflow) {
		var ret string
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeBaseRequestOnTemplate) GetWorkflowOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Workflow) {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *ChallengeBaseRequestOnTemplate) HasWorkflow() bool {
	if o != nil && !utils.IsNil(o.Workflow) {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given string and assigns it to the Workflow field.
func (o *ChallengeBaseRequestOnTemplate) SetWorkflow(v string) {
	o.Workflow = &v
}

func (o ChallengeBaseRequestOnTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChallengeBaseRequestOnTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Module) {
		toSerialize["module"] = o.Module
	}
	if !utils.IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !utils.IsNil(o.Workflow) {
		toSerialize["workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChallengeBaseRequestOnTemplate) UnmarshalJSON(data []byte) (err error) {
	varChallengeBaseRequestOnTemplate := _ChallengeBaseRequestOnTemplate{}

	err = json.Unmarshal(data, &varChallengeBaseRequestOnTemplate)

	if err != nil {
		return err
	}

	*o = ChallengeBaseRequestOnTemplate(varChallengeBaseRequestOnTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "module")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "workflow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChallengeBaseRequestOnTemplate struct {
	value *ChallengeBaseRequestOnTemplate
	isSet bool
}

func (v NullableChallengeBaseRequestOnTemplate) Get() *ChallengeBaseRequestOnTemplate {
	return v.value
}

func (v *NullableChallengeBaseRequestOnTemplate) Set(val *ChallengeBaseRequestOnTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableChallengeBaseRequestOnTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableChallengeBaseRequestOnTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChallengeBaseRequestOnTemplate(val *ChallengeBaseRequestOnTemplate) *NullableChallengeBaseRequestOnTemplate {
	return &NullableChallengeBaseRequestOnTemplate{value: val, isSet: true}
}

func (v NullableChallengeBaseRequestOnTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChallengeBaseRequestOnTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
