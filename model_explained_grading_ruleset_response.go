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

// checks if the ExplainedGradingRulesetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExplainedGradingRulesetResponse{}

// ExplainedGradingRulesetResponse struct for ExplainedGradingRulesetResponse
type ExplainedGradingRulesetResponse struct {
	Name string `json:"name"`
	Description []LocalizedString `json:"description,omitempty"`
	Certificate NullableString `json:"certificate,omitempty"`
	Scope NullableString `json:"scope,omitempty"`
	Apply bool `json:"apply"`
	Max NullableInt64 `json:"max,omitempty"`
	Obtained NullableInt64 `json:"obtained,omitempty"`
	Score NullableFloat32 `json:"score,omitempty"`
	Explained []ExplainedGradingRule `json:"explained,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExplainedGradingRulesetResponse ExplainedGradingRulesetResponse

// NewExplainedGradingRulesetResponse instantiates a new ExplainedGradingRulesetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExplainedGradingRulesetResponse(name string, apply bool) *ExplainedGradingRulesetResponse {
	this := ExplainedGradingRulesetResponse{}
	this.Name = name
	this.Apply = apply
	return &this
}

// NewExplainedGradingRulesetResponseWithDefaults instantiates a new ExplainedGradingRulesetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExplainedGradingRulesetResponseWithDefaults() *ExplainedGradingRulesetResponse {
	this := ExplainedGradingRulesetResponse{}
	return &this
}

// GetName returns the Name field value
func (o *ExplainedGradingRulesetResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExplainedGradingRulesetResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExplainedGradingRulesetResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExplainedGradingRulesetResponse) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExplainedGradingRulesetResponse) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExplainedGradingRulesetResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *ExplainedGradingRulesetResponse) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExplainedGradingRulesetResponse) GetCertificate() string {
	if o == nil || IsNil(o.Certificate.Get()) {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExplainedGradingRulesetResponse) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *ExplainedGradingRulesetResponse) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *ExplainedGradingRulesetResponse) SetCertificate(v string) {
	o.Certificate.Set(&v)
}
// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *ExplainedGradingRulesetResponse) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *ExplainedGradingRulesetResponse) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExplainedGradingRulesetResponse) GetScope() string {
	if o == nil || IsNil(o.Scope.Get()) {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExplainedGradingRulesetResponse) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *ExplainedGradingRulesetResponse) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *ExplainedGradingRulesetResponse) SetScope(v string) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *ExplainedGradingRulesetResponse) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *ExplainedGradingRulesetResponse) UnsetScope() {
	o.Scope.Unset()
}

// GetApply returns the Apply field value
func (o *ExplainedGradingRulesetResponse) GetApply() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Apply
}

// GetApplyOk returns a tuple with the Apply field value
// and a boolean to check if the value has been set.
func (o *ExplainedGradingRulesetResponse) GetApplyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Apply, true
}

// SetApply sets field value
func (o *ExplainedGradingRulesetResponse) SetApply(v bool) {
	o.Apply = v
}

// GetMax returns the Max field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExplainedGradingRulesetResponse) GetMax() int64 {
	if o == nil || IsNil(o.Max.Get()) {
		var ret int64
		return ret
	}
	return *o.Max.Get()
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExplainedGradingRulesetResponse) GetMaxOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Max.Get(), o.Max.IsSet()
}

// HasMax returns a boolean if a field has been set.
func (o *ExplainedGradingRulesetResponse) HasMax() bool {
	if o != nil && o.Max.IsSet() {
		return true
	}

	return false
}

// SetMax gets a reference to the given NullableInt64 and assigns it to the Max field.
func (o *ExplainedGradingRulesetResponse) SetMax(v int64) {
	o.Max.Set(&v)
}
// SetMaxNil sets the value for Max to be an explicit nil
func (o *ExplainedGradingRulesetResponse) SetMaxNil() {
	o.Max.Set(nil)
}

// UnsetMax ensures that no value is present for Max, not even an explicit nil
func (o *ExplainedGradingRulesetResponse) UnsetMax() {
	o.Max.Unset()
}

// GetObtained returns the Obtained field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExplainedGradingRulesetResponse) GetObtained() int64 {
	if o == nil || IsNil(o.Obtained.Get()) {
		var ret int64
		return ret
	}
	return *o.Obtained.Get()
}

// GetObtainedOk returns a tuple with the Obtained field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExplainedGradingRulesetResponse) GetObtainedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Obtained.Get(), o.Obtained.IsSet()
}

// HasObtained returns a boolean if a field has been set.
func (o *ExplainedGradingRulesetResponse) HasObtained() bool {
	if o != nil && o.Obtained.IsSet() {
		return true
	}

	return false
}

// SetObtained gets a reference to the given NullableInt64 and assigns it to the Obtained field.
func (o *ExplainedGradingRulesetResponse) SetObtained(v int64) {
	o.Obtained.Set(&v)
}
// SetObtainedNil sets the value for Obtained to be an explicit nil
func (o *ExplainedGradingRulesetResponse) SetObtainedNil() {
	o.Obtained.Set(nil)
}

// UnsetObtained ensures that no value is present for Obtained, not even an explicit nil
func (o *ExplainedGradingRulesetResponse) UnsetObtained() {
	o.Obtained.Unset()
}

// GetScore returns the Score field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExplainedGradingRulesetResponse) GetScore() float32 {
	if o == nil || IsNil(o.Score.Get()) {
		var ret float32
		return ret
	}
	return *o.Score.Get()
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExplainedGradingRulesetResponse) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Score.Get(), o.Score.IsSet()
}

// HasScore returns a boolean if a field has been set.
func (o *ExplainedGradingRulesetResponse) HasScore() bool {
	if o != nil && o.Score.IsSet() {
		return true
	}

	return false
}

// SetScore gets a reference to the given NullableFloat32 and assigns it to the Score field.
func (o *ExplainedGradingRulesetResponse) SetScore(v float32) {
	o.Score.Set(&v)
}
// SetScoreNil sets the value for Score to be an explicit nil
func (o *ExplainedGradingRulesetResponse) SetScoreNil() {
	o.Score.Set(nil)
}

// UnsetScore ensures that no value is present for Score, not even an explicit nil
func (o *ExplainedGradingRulesetResponse) UnsetScore() {
	o.Score.Unset()
}

// GetExplained returns the Explained field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExplainedGradingRulesetResponse) GetExplained() []ExplainedGradingRule {
	if o == nil {
		var ret []ExplainedGradingRule
		return ret
	}
	return o.Explained
}

// GetExplainedOk returns a tuple with the Explained field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExplainedGradingRulesetResponse) GetExplainedOk() ([]ExplainedGradingRule, bool) {
	if o == nil || IsNil(o.Explained) {
		return nil, false
	}
	return o.Explained, true
}

// HasExplained returns a boolean if a field has been set.
func (o *ExplainedGradingRulesetResponse) HasExplained() bool {
	if o != nil && !IsNil(o.Explained) {
		return true
	}

	return false
}

// SetExplained gets a reference to the given []ExplainedGradingRule and assigns it to the Explained field.
func (o *ExplainedGradingRulesetResponse) SetExplained(v []ExplainedGradingRule) {
	o.Explained = v
}

func (o ExplainedGradingRulesetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExplainedGradingRulesetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}
	toSerialize["apply"] = o.Apply
	if o.Max.IsSet() {
		toSerialize["max"] = o.Max.Get()
	}
	if o.Obtained.IsSet() {
		toSerialize["obtained"] = o.Obtained.Get()
	}
	if o.Score.IsSet() {
		toSerialize["score"] = o.Score.Get()
	}
	if o.Explained != nil {
		toSerialize["explained"] = o.Explained
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExplainedGradingRulesetResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"apply",
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

	varExplainedGradingRulesetResponse := _ExplainedGradingRulesetResponse{}

	err = json.Unmarshal(data, &varExplainedGradingRulesetResponse)

	if err != nil {
		return err
	}

	*o = ExplainedGradingRulesetResponse(varExplainedGradingRulesetResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "apply")
		delete(additionalProperties, "max")
		delete(additionalProperties, "obtained")
		delete(additionalProperties, "score")
		delete(additionalProperties, "explained")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExplainedGradingRulesetResponse struct {
	value *ExplainedGradingRulesetResponse
	isSet bool
}

func (v NullableExplainedGradingRulesetResponse) Get() *ExplainedGradingRulesetResponse {
	return v.value
}

func (v *NullableExplainedGradingRulesetResponse) Set(val *ExplainedGradingRulesetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExplainedGradingRulesetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExplainedGradingRulesetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExplainedGradingRulesetResponse(val *ExplainedGradingRulesetResponse) *NullableExplainedGradingRulesetResponse {
	return &NullableExplainedGradingRulesetResponse{value: val, isSet: true}
}

func (v NullableExplainedGradingRulesetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExplainedGradingRulesetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


