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

// checks if the DNElement type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DNElement{}

// DNElement struct for DNElement
type DNElement struct {
	// A computation rule that will dynamically generate a string value from the request's context
	ComputationRule      utils.NullableString `json:"computationRule,omitempty"`
	EditableByApprover   utils.NullableBool   `json:"editableByApprover,omitempty"`
	EditableByRequester  utils.NullableBool   `json:"editableByRequester,omitempty"`
	Mandatory            bool                 `json:"mandatory"`
	Regex                utils.NullableString `json:"regex,omitempty"`
	Type                 string               `json:"type"`
	Value                utils.NullableString `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DNElement DNElement

// NewDNElement instantiates a new DNElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDNElement(mandatory bool, type_ string) *DNElement {
	this := DNElement{}
	this.Mandatory = mandatory
	this.Type = type_
	return &this
}

// NewDNElementWithDefaults instantiates a new DNElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDNElementWithDefaults() *DNElement {
	this := DNElement{}
	return &this
}

// GetComputationRule returns the ComputationRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DNElement) GetComputationRule() string {
	if o == nil || utils.IsNil(o.ComputationRule.Get()) {
		var ret string
		return ret
	}
	return *o.ComputationRule.Get()
}

// GetComputationRuleOk returns a tuple with the ComputationRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DNElement) GetComputationRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputationRule.Get(), o.ComputationRule.IsSet()
}

// HasComputationRule returns a boolean if a field has been set.
func (o *DNElement) HasComputationRule() bool {
	if o != nil && o.ComputationRule.IsSet() {
		return true
	}

	return false
}

// SetComputationRule gets a reference to the given NullableString and assigns it to the ComputationRule field.
func (o *DNElement) SetComputationRule(v string) {
	o.ComputationRule.Set(&v)
}

// SetComputationRuleNil sets the value for ComputationRule to be an explicit nil
func (o *DNElement) SetComputationRuleNil() {
	o.ComputationRule.Set(nil)
}

// UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
func (o *DNElement) UnsetComputationRule() {
	o.ComputationRule.Unset()
}

// GetEditableByApprover returns the EditableByApprover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DNElement) GetEditableByApprover() bool {
	if o == nil || utils.IsNil(o.EditableByApprover.Get()) {
		var ret bool
		return ret
	}
	return *o.EditableByApprover.Get()
}

// GetEditableByApproverOk returns a tuple with the EditableByApprover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DNElement) GetEditableByApproverOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditableByApprover.Get(), o.EditableByApprover.IsSet()
}

// HasEditableByApprover returns a boolean if a field has been set.
func (o *DNElement) HasEditableByApprover() bool {
	if o != nil && o.EditableByApprover.IsSet() {
		return true
	}

	return false
}

// SetEditableByApprover gets a reference to the given NullableBool and assigns it to the EditableByApprover field.
func (o *DNElement) SetEditableByApprover(v bool) {
	o.EditableByApprover.Set(&v)
}

// SetEditableByApproverNil sets the value for EditableByApprover to be an explicit nil
func (o *DNElement) SetEditableByApproverNil() {
	o.EditableByApprover.Set(nil)
}

// UnsetEditableByApprover ensures that no value is present for EditableByApprover, not even an explicit nil
func (o *DNElement) UnsetEditableByApprover() {
	o.EditableByApprover.Unset()
}

// GetEditableByRequester returns the EditableByRequester field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DNElement) GetEditableByRequester() bool {
	if o == nil || utils.IsNil(o.EditableByRequester.Get()) {
		var ret bool
		return ret
	}
	return *o.EditableByRequester.Get()
}

// GetEditableByRequesterOk returns a tuple with the EditableByRequester field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DNElement) GetEditableByRequesterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditableByRequester.Get(), o.EditableByRequester.IsSet()
}

// HasEditableByRequester returns a boolean if a field has been set.
func (o *DNElement) HasEditableByRequester() bool {
	if o != nil && o.EditableByRequester.IsSet() {
		return true
	}

	return false
}

// SetEditableByRequester gets a reference to the given NullableBool and assigns it to the EditableByRequester field.
func (o *DNElement) SetEditableByRequester(v bool) {
	o.EditableByRequester.Set(&v)
}

// SetEditableByRequesterNil sets the value for EditableByRequester to be an explicit nil
func (o *DNElement) SetEditableByRequesterNil() {
	o.EditableByRequester.Set(nil)
}

// UnsetEditableByRequester ensures that no value is present for EditableByRequester, not even an explicit nil
func (o *DNElement) UnsetEditableByRequester() {
	o.EditableByRequester.Unset()
}

// GetMandatory returns the Mandatory field value
func (o *DNElement) GetMandatory() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Mandatory
}

// GetMandatoryOk returns a tuple with the Mandatory field value
// and a boolean to check if the value has been set.
func (o *DNElement) GetMandatoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mandatory, true
}

// SetMandatory sets field value
func (o *DNElement) SetMandatory(v bool) {
	o.Mandatory = v
}

// GetRegex returns the Regex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DNElement) GetRegex() string {
	if o == nil || utils.IsNil(o.Regex.Get()) {
		var ret string
		return ret
	}
	return *o.Regex.Get()
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DNElement) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regex.Get(), o.Regex.IsSet()
}

// HasRegex returns a boolean if a field has been set.
func (o *DNElement) HasRegex() bool {
	if o != nil && o.Regex.IsSet() {
		return true
	}

	return false
}

// SetRegex gets a reference to the given NullableString and assigns it to the Regex field.
func (o *DNElement) SetRegex(v string) {
	o.Regex.Set(&v)
}

// SetRegexNil sets the value for Regex to be an explicit nil
func (o *DNElement) SetRegexNil() {
	o.Regex.Set(nil)
}

// UnsetRegex ensures that no value is present for Regex, not even an explicit nil
func (o *DNElement) UnsetRegex() {
	o.Regex.Unset()
}

// GetType returns the Type field value
func (o *DNElement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DNElement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DNElement) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DNElement) GetValue() string {
	if o == nil || utils.IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DNElement) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *DNElement) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *DNElement) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *DNElement) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *DNElement) UnsetValue() {
	o.Value.Unset()
}

func (o DNElement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DNElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ComputationRule.IsSet() {
		toSerialize["computationRule"] = o.ComputationRule.Get()
	}
	if o.EditableByApprover.IsSet() {
		toSerialize["editableByApprover"] = o.EditableByApprover.Get()
	}
	if o.EditableByRequester.IsSet() {
		toSerialize["editableByRequester"] = o.EditableByRequester.Get()
	}
	toSerialize["mandatory"] = o.Mandatory
	if o.Regex.IsSet() {
		toSerialize["regex"] = o.Regex.Get()
	}
	toSerialize["type"] = o.Type
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DNElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mandatory",
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

	varDNElement := _DNElement{}

	err = json.Unmarshal(data, &varDNElement)

	if err != nil {
		return err
	}

	*o = DNElement(varDNElement)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "computationRule")
		delete(additionalProperties, "editableByApprover")
		delete(additionalProperties, "editableByRequester")
		delete(additionalProperties, "mandatory")
		delete(additionalProperties, "regex")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDNElement struct {
	value *DNElement
	isSet bool
}

func (v NullableDNElement) Get() *DNElement {
	return v.value
}

func (v *NullableDNElement) Set(val *DNElement) {
	v.value = val
	v.isSet = true
}

func (v NullableDNElement) IsSet() bool {
	return v.isSet
}

func (v *NullableDNElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDNElement(val *DNElement) *NullableDNElement {
	return &NullableDNElement{value: val, isSet: true}
}

func (v NullableDNElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDNElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
