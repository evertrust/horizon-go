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

// checks if the ExplainedGradingPolicyResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExplainedGradingPolicyResponse{}

// ExplainedGradingPolicyResponse struct for ExplainedGradingPolicyResponse
type ExplainedGradingPolicyResponse struct {
	Certificate          string                    `json:"certificate"`
	Description          []LocalizedString         `json:"description,omitempty"`
	Explained            []ExplainedGradingRuleset `json:"explained,omitempty"`
	Grade                utils.NullableString      `json:"grade,omitempty"`
	Name                 string                    `json:"name"`
	Rulesets             []WeightedGradingRuleset  `json:"rulesets"`
	Score                utils.NullableFloat32     `json:"score,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExplainedGradingPolicyResponse ExplainedGradingPolicyResponse

// NewExplainedGradingPolicyResponse instantiates a new ExplainedGradingPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExplainedGradingPolicyResponse(certificate string, name string, rulesets []WeightedGradingRuleset) *ExplainedGradingPolicyResponse {
	this := ExplainedGradingPolicyResponse{}
	this.Certificate = certificate
	this.Name = name
	this.Rulesets = rulesets
	return &this
}

// NewExplainedGradingPolicyResponseWithDefaults instantiates a new ExplainedGradingPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExplainedGradingPolicyResponseWithDefaults() *ExplainedGradingPolicyResponse {
	this := ExplainedGradingPolicyResponse{}
	return &this
}

// GetCertificate returns the Certificate field value
func (o *ExplainedGradingPolicyResponse) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *ExplainedGradingPolicyResponse) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *ExplainedGradingPolicyResponse) SetCertificate(v string) {
	o.Certificate = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExplainedGradingPolicyResponse) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExplainedGradingPolicyResponse) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExplainedGradingPolicyResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *ExplainedGradingPolicyResponse) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetExplained returns the Explained field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExplainedGradingPolicyResponse) GetExplained() []ExplainedGradingRuleset {
	if o == nil {
		var ret []ExplainedGradingRuleset
		return ret
	}
	return o.Explained
}

// GetExplainedOk returns a tuple with the Explained field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExplainedGradingPolicyResponse) GetExplainedOk() ([]ExplainedGradingRuleset, bool) {
	if o == nil || utils.IsNil(o.Explained) {
		return nil, false
	}
	return o.Explained, true
}

// HasExplained returns a boolean if a field has been set.
func (o *ExplainedGradingPolicyResponse) HasExplained() bool {
	if o != nil && !utils.IsNil(o.Explained) {
		return true
	}

	return false
}

// SetExplained gets a reference to the given []ExplainedGradingRuleset and assigns it to the Explained field.
func (o *ExplainedGradingPolicyResponse) SetExplained(v []ExplainedGradingRuleset) {
	o.Explained = v
}

// GetGrade returns the Grade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExplainedGradingPolicyResponse) GetGrade() string {
	if o == nil || utils.IsNil(o.Grade.Get()) {
		var ret string
		return ret
	}
	return *o.Grade.Get()
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExplainedGradingPolicyResponse) GetGradeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Grade.Get(), o.Grade.IsSet()
}

// HasGrade returns a boolean if a field has been set.
func (o *ExplainedGradingPolicyResponse) HasGrade() bool {
	if o != nil && o.Grade.IsSet() {
		return true
	}

	return false
}

// SetGrade gets a reference to the given NullableString and assigns it to the Grade field.
func (o *ExplainedGradingPolicyResponse) SetGrade(v string) {
	o.Grade.Set(&v)
}

// SetGradeNil sets the value for Grade to be an explicit nil
func (o *ExplainedGradingPolicyResponse) SetGradeNil() {
	o.Grade.Set(nil)
}

// UnsetGrade ensures that no value is present for Grade, not even an explicit nil
func (o *ExplainedGradingPolicyResponse) UnsetGrade() {
	o.Grade.Unset()
}

// GetName returns the Name field value
func (o *ExplainedGradingPolicyResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExplainedGradingPolicyResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExplainedGradingPolicyResponse) SetName(v string) {
	o.Name = v
}

// GetRulesets returns the Rulesets field value
func (o *ExplainedGradingPolicyResponse) GetRulesets() []WeightedGradingRuleset {
	if o == nil {
		var ret []WeightedGradingRuleset
		return ret
	}

	return o.Rulesets
}

// GetRulesetsOk returns a tuple with the Rulesets field value
// and a boolean to check if the value has been set.
func (o *ExplainedGradingPolicyResponse) GetRulesetsOk() ([]WeightedGradingRuleset, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rulesets, true
}

// SetRulesets sets field value
func (o *ExplainedGradingPolicyResponse) SetRulesets(v []WeightedGradingRuleset) {
	o.Rulesets = v
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExplainedGradingPolicyResponse) GetScore() float32 {
	if o == nil || utils.IsNil(o.Score.Get()) {
		var ret float32
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExplainedGradingPolicyResponse) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// HasScore returns a boolean if a field has been set.
func (o *ExplainedGradingPolicyResponse) HasScore() bool {
	if o != nil && o.Score.IsSet() {
		return true
	}

	return false
}

// SetScore gets a reference to the given NullableFloat32 and assigns it to the Score field.
func (o *ExplainedGradingPolicyResponse) SetScore(v float32) {
	o.Score.Set(&v)
}

// SetScoreNil sets the value for Score to be an explicit nil
func (o *ExplainedGradingPolicyResponse) SetScoreNil() {
	o.Score.Set(nil)
}

// UnsetScore ensures that no value is present for Score, not even an explicit nil
func (o *ExplainedGradingPolicyResponse) UnsetScore() {
	o.Score.Unset()
}

func (o ExplainedGradingPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExplainedGradingPolicyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["certificate"] = o.Certificate
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Explained != nil {
		toSerialize["explained"] = o.Explained
	}
	if o.Grade.IsSet() {
		toSerialize["grade"] = o.Grade.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["rulesets"] = o.Rulesets
	if o.Score.IsSet() {
		toSerialize["score"] = o.Score.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExplainedGradingPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"certificate",
		"name",
		"rulesets",
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

	varExplainedGradingPolicyResponse := _ExplainedGradingPolicyResponse{}

	err = json.Unmarshal(data, &varExplainedGradingPolicyResponse)

	if err != nil {
		return err
	}

	*o = ExplainedGradingPolicyResponse(varExplainedGradingPolicyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "description")
		delete(additionalProperties, "explained")
		delete(additionalProperties, "grade")
		delete(additionalProperties, "name")
		delete(additionalProperties, "rulesets")
		delete(additionalProperties, "score")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExplainedGradingPolicyResponse struct {
	value *ExplainedGradingPolicyResponse
	isSet bool
}

func (v NullableExplainedGradingPolicyResponse) Get() *ExplainedGradingPolicyResponse {
	return v.value
}

func (v *NullableExplainedGradingPolicyResponse) Set(val *ExplainedGradingPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExplainedGradingPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExplainedGradingPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExplainedGradingPolicyResponse(val *ExplainedGradingPolicyResponse) *NullableExplainedGradingPolicyResponse {
	return &NullableExplainedGradingPolicyResponse{value: val, isSet: true}
}

func (v NullableExplainedGradingPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExplainedGradingPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
