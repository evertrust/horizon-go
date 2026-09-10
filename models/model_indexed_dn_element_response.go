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

// checks if the IndexedDNElementResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IndexedDNElementResponse{}

// IndexedDNElementResponse An element of a certificate's Distinguished Name
type IndexedDNElementResponse struct {
	// Computation rule input will be evaluated and will override all other inputs
	ComputationRule utils.NullableString `json:"computationRule,omitempty"`
	// Whether the field is editable or not for the currently authenticated user
	Editable utils.NullableBool `json:"editable,omitempty"`
	// The element type and index. Indexes start at 1! Available elements are: `cn`, `e`, `ou`, `st`, `l`, `o`, `c`, `dc`, `uid`, `serialNumber`, `surname`, `givenName`, `unstructuredAddress`, `unstructuredName`, `organizationIdentifier`, `uniqueIdentifier`, `street`, `description`, `t`
	Element string `json:"element" validate:"regexp=(cn|e|ou|st|l|o|c|dc|uid|serialNumber|surname|givenName|unstructuredAddress|unstructuredName|organizationIdentifier|uniqueIdentifier|street|description|t)\\\\.[1-9]\\\\d*"`
	// Whether the field is mandatory or not
	Mandatory utils.NullableBool `json:"mandatory,omitempty"`
	// A regular expression that will be used to validate the element's value
	Regex utils.NullableString `json:"regex,omitempty"`
	// The formatted element type
	Type utils.NullableString `json:"type,omitempty"`
	// The element value
	Value                utils.NullableString `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndexedDNElementResponse IndexedDNElementResponse

// NewIndexedDNElementResponse instantiates a new IndexedDNElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexedDNElementResponse(element string) *IndexedDNElementResponse {
	this := IndexedDNElementResponse{}
	this.Element = element
	return &this
}

// NewIndexedDNElementResponseWithDefaults instantiates a new IndexedDNElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexedDNElementResponseWithDefaults() *IndexedDNElementResponse {
	this := IndexedDNElementResponse{}
	return &this
}

// GetComputationRule returns the ComputationRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexedDNElementResponse) GetComputationRule() string {
	if o == nil || utils.IsNil(o.ComputationRule.Get()) {
		var ret string
		return ret
	}
	return *o.ComputationRule.Get()
}

// GetComputationRuleOk returns a tuple with the ComputationRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexedDNElementResponse) GetComputationRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ComputationRule.Get(), o.ComputationRule.IsSet()
}

// HasComputationRule returns a boolean if a field has been set.
func (o *IndexedDNElementResponse) HasComputationRule() bool {
	if o != nil && o.ComputationRule.IsSet() {
		return true
	}

	return false
}

// SetComputationRule gets a reference to the given NullableString and assigns it to the ComputationRule field.
func (o *IndexedDNElementResponse) SetComputationRule(v string) {
	o.ComputationRule.Set(&v)
}

// SetComputationRuleNil sets the value for ComputationRule to be an explicit nil
func (o *IndexedDNElementResponse) SetComputationRuleNil() {
	o.ComputationRule.Set(nil)
}

// UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
func (o *IndexedDNElementResponse) UnsetComputationRule() {
	o.ComputationRule.Unset()
}

// GetEditable returns the Editable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexedDNElementResponse) GetEditable() bool {
	if o == nil || utils.IsNil(o.Editable.Get()) {
		var ret bool
		return ret
	}
	return *o.Editable.Get()
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexedDNElementResponse) GetEditableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Editable.Get(), o.Editable.IsSet()
}

// HasEditable returns a boolean if a field has been set.
func (o *IndexedDNElementResponse) HasEditable() bool {
	if o != nil && o.Editable.IsSet() {
		return true
	}

	return false
}

// SetEditable gets a reference to the given NullableBool and assigns it to the Editable field.
func (o *IndexedDNElementResponse) SetEditable(v bool) {
	o.Editable.Set(&v)
}

// SetEditableNil sets the value for Editable to be an explicit nil
func (o *IndexedDNElementResponse) SetEditableNil() {
	o.Editable.Set(nil)
}

// UnsetEditable ensures that no value is present for Editable, not even an explicit nil
func (o *IndexedDNElementResponse) UnsetEditable() {
	o.Editable.Unset()
}

// GetElement returns the Element field value
func (o *IndexedDNElementResponse) GetElement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Element
}

// GetElementOk returns a tuple with the Element field value
// and a boolean to check if the value has been set.
func (o *IndexedDNElementResponse) GetElementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Element, true
}

// SetElement sets field value
func (o *IndexedDNElementResponse) SetElement(v string) {
	o.Element = v
}

// GetMandatory returns the Mandatory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexedDNElementResponse) GetMandatory() bool {
	if o == nil || utils.IsNil(o.Mandatory.Get()) {
		var ret bool
		return ret
	}
	return *o.Mandatory.Get()
}

// GetMandatoryOk returns a tuple with the Mandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexedDNElementResponse) GetMandatoryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mandatory.Get(), o.Mandatory.IsSet()
}

// HasMandatory returns a boolean if a field has been set.
func (o *IndexedDNElementResponse) HasMandatory() bool {
	if o != nil && o.Mandatory.IsSet() {
		return true
	}

	return false
}

// SetMandatory gets a reference to the given NullableBool and assigns it to the Mandatory field.
func (o *IndexedDNElementResponse) SetMandatory(v bool) {
	o.Mandatory.Set(&v)
}

// SetMandatoryNil sets the value for Mandatory to be an explicit nil
func (o *IndexedDNElementResponse) SetMandatoryNil() {
	o.Mandatory.Set(nil)
}

// UnsetMandatory ensures that no value is present for Mandatory, not even an explicit nil
func (o *IndexedDNElementResponse) UnsetMandatory() {
	o.Mandatory.Unset()
}

// GetRegex returns the Regex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexedDNElementResponse) GetRegex() string {
	if o == nil || utils.IsNil(o.Regex.Get()) {
		var ret string
		return ret
	}
	return *o.Regex.Get()
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexedDNElementResponse) GetRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regex.Get(), o.Regex.IsSet()
}

// HasRegex returns a boolean if a field has been set.
func (o *IndexedDNElementResponse) HasRegex() bool {
	if o != nil && o.Regex.IsSet() {
		return true
	}

	return false
}

// SetRegex gets a reference to the given NullableString and assigns it to the Regex field.
func (o *IndexedDNElementResponse) SetRegex(v string) {
	o.Regex.Set(&v)
}

// SetRegexNil sets the value for Regex to be an explicit nil
func (o *IndexedDNElementResponse) SetRegexNil() {
	o.Regex.Set(nil)
}

// UnsetRegex ensures that no value is present for Regex, not even an explicit nil
func (o *IndexedDNElementResponse) UnsetRegex() {
	o.Regex.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexedDNElementResponse) GetType() string {
	if o == nil || utils.IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexedDNElementResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *IndexedDNElementResponse) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *IndexedDNElementResponse) SetType(v string) {
	o.Type.Set(&v)
}

// SetTypeNil sets the value for Type to be an explicit nil
func (o *IndexedDNElementResponse) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *IndexedDNElementResponse) UnsetType() {
	o.Type.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IndexedDNElementResponse) GetValue() string {
	if o == nil || utils.IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IndexedDNElementResponse) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *IndexedDNElementResponse) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *IndexedDNElementResponse) SetValue(v string) {
	o.Value.Set(&v)
}

// SetValueNil sets the value for Value to be an explicit nil
func (o *IndexedDNElementResponse) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *IndexedDNElementResponse) UnsetValue() {
	o.Value.Unset()
}

func (o IndexedDNElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexedDNElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ComputationRule.IsSet() {
		toSerialize["computationRule"] = o.ComputationRule.Get()
	}
	if o.Editable.IsSet() {
		toSerialize["editable"] = o.Editable.Get()
	}
	toSerialize["element"] = o.Element
	if o.Mandatory.IsSet() {
		toSerialize["mandatory"] = o.Mandatory.Get()
	}
	if o.Regex.IsSet() {
		toSerialize["regex"] = o.Regex.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IndexedDNElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"element",
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

	varIndexedDNElementResponse := _IndexedDNElementResponse{}

	err = json.Unmarshal(data, &varIndexedDNElementResponse)

	if err != nil {
		return err
	}

	*o = IndexedDNElementResponse(varIndexedDNElementResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "computationRule")
		delete(additionalProperties, "editable")
		delete(additionalProperties, "element")
		delete(additionalProperties, "mandatory")
		delete(additionalProperties, "regex")
		delete(additionalProperties, "type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndexedDNElementResponse struct {
	value *IndexedDNElementResponse
	isSet bool
}

func (v NullableIndexedDNElementResponse) Get() *IndexedDNElementResponse {
	return v.value
}

func (v *NullableIndexedDNElementResponse) Set(val *IndexedDNElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexedDNElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexedDNElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexedDNElementResponse(val *IndexedDNElementResponse) *NullableIndexedDNElementResponse {
	return &NullableIndexedDNElementResponse{value: val, isSet: true}
}

func (v NullableIndexedDNElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexedDNElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
