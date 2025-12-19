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

// checks if the AutomationPolicy type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AutomationPolicy{}

// AutomationPolicy struct for AutomationPolicy
type AutomationPolicy struct {
	CompliancePolicy     NullableCompliancePolicy `json:"compliancePolicy,omitempty"`
	ExecutionPolicy      utils.NullableString     `json:"executionPolicy,omitempty"`
	Name                 string                   `json:"name"`
	Profile              string                   `json:"profile"`
	TrustChains          []string                 `json:"trustChains,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutomationPolicy AutomationPolicy

// NewAutomationPolicy instantiates a new AutomationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationPolicy(name string, profile string) *AutomationPolicy {
	this := AutomationPolicy{}
	this.Name = name
	this.Profile = profile
	return &this
}

// NewAutomationPolicyWithDefaults instantiates a new AutomationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationPolicyWithDefaults() *AutomationPolicy {
	this := AutomationPolicy{}
	return &this
}

// GetCompliancePolicy returns the CompliancePolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutomationPolicy) GetCompliancePolicy() CompliancePolicy {
	if o == nil || utils.IsNil(o.CompliancePolicy.Get()) {
		var ret CompliancePolicy
		return ret
	}
	return *o.CompliancePolicy.Get()
}

// GetCompliancePolicyOk returns a tuple with the CompliancePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutomationPolicy) GetCompliancePolicyOk() (*CompliancePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompliancePolicy.Get(), o.CompliancePolicy.IsSet()
}

// HasCompliancePolicy returns a boolean if a field has been set.
func (o *AutomationPolicy) HasCompliancePolicy() bool {
	if o != nil && o.CompliancePolicy.IsSet() {
		return true
	}

	return false
}

// SetCompliancePolicy gets a reference to the given NullableCompliancePolicy and assigns it to the CompliancePolicy field.
func (o *AutomationPolicy) SetCompliancePolicy(v CompliancePolicy) {
	o.CompliancePolicy.Set(&v)
}

// SetCompliancePolicyNil sets the value for CompliancePolicy to be an explicit nil
func (o *AutomationPolicy) SetCompliancePolicyNil() {
	o.CompliancePolicy.Set(nil)
}

// UnsetCompliancePolicy ensures that no value is present for CompliancePolicy, not even an explicit nil
func (o *AutomationPolicy) UnsetCompliancePolicy() {
	o.CompliancePolicy.Unset()
}

// GetExecutionPolicy returns the ExecutionPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutomationPolicy) GetExecutionPolicy() string {
	if o == nil || utils.IsNil(o.ExecutionPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionPolicy.Get()
}

// GetExecutionPolicyOk returns a tuple with the ExecutionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutomationPolicy) GetExecutionPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionPolicy.Get(), o.ExecutionPolicy.IsSet()
}

// HasExecutionPolicy returns a boolean if a field has been set.
func (o *AutomationPolicy) HasExecutionPolicy() bool {
	if o != nil && o.ExecutionPolicy.IsSet() {
		return true
	}

	return false
}

// SetExecutionPolicy gets a reference to the given NullableString and assigns it to the ExecutionPolicy field.
func (o *AutomationPolicy) SetExecutionPolicy(v string) {
	o.ExecutionPolicy.Set(&v)
}

// SetExecutionPolicyNil sets the value for ExecutionPolicy to be an explicit nil
func (o *AutomationPolicy) SetExecutionPolicyNil() {
	o.ExecutionPolicy.Set(nil)
}

// UnsetExecutionPolicy ensures that no value is present for ExecutionPolicy, not even an explicit nil
func (o *AutomationPolicy) UnsetExecutionPolicy() {
	o.ExecutionPolicy.Unset()
}

// GetName returns the Name field value
func (o *AutomationPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutomationPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutomationPolicy) SetName(v string) {
	o.Name = v
}

// GetProfile returns the Profile field value
func (o *AutomationPolicy) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *AutomationPolicy) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *AutomationPolicy) SetProfile(v string) {
	o.Profile = v
}

// GetTrustChains returns the TrustChains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutomationPolicy) GetTrustChains() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TrustChains
}

// GetTrustChainsOk returns a tuple with the TrustChains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutomationPolicy) GetTrustChainsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.TrustChains) {
		return nil, false
	}
	return o.TrustChains, true
}

// HasTrustChains returns a boolean if a field has been set.
func (o *AutomationPolicy) HasTrustChains() bool {
	if o != nil && !utils.IsNil(o.TrustChains) {
		return true
	}

	return false
}

// SetTrustChains gets a reference to the given []string and assigns it to the TrustChains field.
func (o *AutomationPolicy) SetTrustChains(v []string) {
	o.TrustChains = v
}

func (o AutomationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutomationPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CompliancePolicy.IsSet() {
		toSerialize["compliancePolicy"] = o.CompliancePolicy.Get()
	}
	if o.ExecutionPolicy.IsSet() {
		toSerialize["executionPolicy"] = o.ExecutionPolicy.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["profile"] = o.Profile
	if o.TrustChains != nil {
		toSerialize["trustChains"] = o.TrustChains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AutomationPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"profile",
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

	varAutomationPolicy := _AutomationPolicy{}

	err = json.Unmarshal(data, &varAutomationPolicy)

	if err != nil {
		return err
	}

	*o = AutomationPolicy(varAutomationPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "compliancePolicy")
		delete(additionalProperties, "executionPolicy")
		delete(additionalProperties, "name")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "trustChains")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutomationPolicy struct {
	value *AutomationPolicy
	isSet bool
}

func (v NullableAutomationPolicy) Get() *AutomationPolicy {
	return v.value
}

func (v *NullableAutomationPolicy) Set(val *AutomationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAutomationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAutomationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutomationPolicy(val *AutomationPolicy) *NullableAutomationPolicy {
	return &NullableAutomationPolicy{value: val, isSet: true}
}

func (v NullableAutomationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutomationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
