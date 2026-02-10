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

// checks if the ExecutionPolicyResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExecutionPolicyResponse{}

// ExecutionPolicyResponse struct for ExecutionPolicyResponse
type ExecutionPolicyResponse struct {
	// Object internal ID
	Id                   string               `json:"_id"`
	Name                 string               `json:"name"`
	Description          utils.NullableString `json:"description,omitempty"`
	AuthorizedPeriods    []ExecutionPeriod    `json:"authorizedPeriods,omitempty"`
	ForbiddenPeriods     []ExecutionPeriod    `json:"forbiddenPeriods,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionPolicyResponse ExecutionPolicyResponse

// NewExecutionPolicyResponse instantiates a new ExecutionPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionPolicyResponse(id string, name string) *ExecutionPolicyResponse {
	this := ExecutionPolicyResponse{}
	this.Id = id
	this.Name = name
	return &this
}

// NewExecutionPolicyResponseWithDefaults instantiates a new ExecutionPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionPolicyResponseWithDefaults() *ExecutionPolicyResponse {
	this := ExecutionPolicyResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ExecutionPolicyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExecutionPolicyResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ExecutionPolicyResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExecutionPolicyResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExecutionPolicyResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecutionPolicyResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExecutionPolicyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExecutionPolicyResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ExecutionPolicyResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ExecutionPolicyResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ExecutionPolicyResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetAuthorizedPeriods returns the AuthorizedPeriods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecutionPolicyResponse) GetAuthorizedPeriods() []ExecutionPeriod {
	if o == nil {
		var ret []ExecutionPeriod
		return ret
	}
	return o.AuthorizedPeriods
}

// GetAuthorizedPeriodsOk returns a tuple with the AuthorizedPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExecutionPolicyResponse) GetAuthorizedPeriodsOk() ([]ExecutionPeriod, bool) {
	if o == nil || utils.IsNil(o.AuthorizedPeriods) {
		return nil, false
	}
	return o.AuthorizedPeriods, true
}

// HasAuthorizedPeriods returns a boolean if a field has been set.
func (o *ExecutionPolicyResponse) HasAuthorizedPeriods() bool {
	if o != nil && !utils.IsNil(o.AuthorizedPeriods) {
		return true
	}

	return false
}

// SetAuthorizedPeriods gets a reference to the given []ExecutionPeriod and assigns it to the AuthorizedPeriods field.
func (o *ExecutionPolicyResponse) SetAuthorizedPeriods(v []ExecutionPeriod) {
	o.AuthorizedPeriods = v
}

// GetForbiddenPeriods returns the ForbiddenPeriods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExecutionPolicyResponse) GetForbiddenPeriods() []ExecutionPeriod {
	if o == nil {
		var ret []ExecutionPeriod
		return ret
	}
	return o.ForbiddenPeriods
}

// GetForbiddenPeriodsOk returns a tuple with the ForbiddenPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExecutionPolicyResponse) GetForbiddenPeriodsOk() ([]ExecutionPeriod, bool) {
	if o == nil || utils.IsNil(o.ForbiddenPeriods) {
		return nil, false
	}
	return o.ForbiddenPeriods, true
}

// HasForbiddenPeriods returns a boolean if a field has been set.
func (o *ExecutionPolicyResponse) HasForbiddenPeriods() bool {
	if o != nil && !utils.IsNil(o.ForbiddenPeriods) {
		return true
	}

	return false
}

// SetForbiddenPeriods gets a reference to the given []ExecutionPeriod and assigns it to the ForbiddenPeriods field.
func (o *ExecutionPolicyResponse) SetForbiddenPeriods(v []ExecutionPeriod) {
	o.ForbiddenPeriods = v
}

func (o ExecutionPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionPolicyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.AuthorizedPeriods != nil {
		toSerialize["authorizedPeriods"] = o.AuthorizedPeriods
	}
	if o.ForbiddenPeriods != nil {
		toSerialize["forbiddenPeriods"] = o.ForbiddenPeriods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExecutionPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"name",
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

	varExecutionPolicyResponse := _ExecutionPolicyResponse{}

	err = json.Unmarshal(data, &varExecutionPolicyResponse)

	if err != nil {
		return err
	}

	*o = ExecutionPolicyResponse(varExecutionPolicyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "authorizedPeriods")
		delete(additionalProperties, "forbiddenPeriods")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionPolicyResponse struct {
	value *ExecutionPolicyResponse
	isSet bool
}

func (v NullableExecutionPolicyResponse) Get() *ExecutionPolicyResponse {
	return v.value
}

func (v *NullableExecutionPolicyResponse) Set(val *ExecutionPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionPolicyResponse(val *ExecutionPolicyResponse) *NullableExecutionPolicyResponse {
	return &NullableExecutionPolicyResponse{value: val, isSet: true}
}

func (v NullableExecutionPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
