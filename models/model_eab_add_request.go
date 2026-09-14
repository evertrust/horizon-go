/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the EabAddRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EabAddRequest{}

// EabAddRequest struct for EabAddRequest
type EabAddRequest struct {
	AllowedProfiles      []string                        `json:"allowedProfiles,omitempty"`
	Description          *string                         `json:"description,omitempty"`
	EabPolicy            string                          `json:"eabPolicy"`
	EabValidityDuration  utils.NullableString            `json:"eabValidityDuration,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	EmailConstraint      *string                         `json:"emailConstraint,omitempty"`
	IdentifierConstraint *string                         `json:"identifierConstraint,omitempty"`
	MacKeyAlgorithm      ExternalAccountBindingAlgorithm `json:"macKeyAlgorithm"`
	Name                 string                          `json:"name"`
	ValidationMethods    []AcmeAuthorizationType         `json:"validationMethods,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EabAddRequest EabAddRequest

// NewEabAddRequest instantiates a new EabAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEabAddRequest(eabPolicy string, macKeyAlgorithm ExternalAccountBindingAlgorithm, name string) *EabAddRequest {
	this := EabAddRequest{}
	this.EabPolicy = eabPolicy
	this.MacKeyAlgorithm = macKeyAlgorithm
	this.Name = name
	return &this
}

// NewEabAddRequestWithDefaults instantiates a new EabAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEabAddRequestWithDefaults() *EabAddRequest {
	this := EabAddRequest{}
	return &this
}

// GetAllowedProfiles returns the AllowedProfiles field value if set, zero value otherwise.
func (o *EabAddRequest) GetAllowedProfiles() []string {
	if o == nil || utils.IsNil(o.AllowedProfiles) {
		var ret []string
		return ret
	}
	return o.AllowedProfiles
}

// GetAllowedProfilesOk returns a tuple with the AllowedProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EabAddRequest) GetAllowedProfilesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AllowedProfiles) {
		return nil, false
	}
	return o.AllowedProfiles, true
}

// HasAllowedProfiles returns a boolean if a field has been set.
func (o *EabAddRequest) HasAllowedProfiles() bool {
	if o != nil && !utils.IsNil(o.AllowedProfiles) {
		return true
	}

	return false
}

// SetAllowedProfiles gets a reference to the given []string and assigns it to the AllowedProfiles field.
func (o *EabAddRequest) SetAllowedProfiles(v []string) {
	o.AllowedProfiles = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EabAddRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EabAddRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EabAddRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EabAddRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEabPolicy returns the EabPolicy field value
func (o *EabAddRequest) GetEabPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EabPolicy
}

// GetEabPolicyOk returns a tuple with the EabPolicy field value
// and a boolean to check if the value has been set.
func (o *EabAddRequest) GetEabPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EabPolicy, true
}

// SetEabPolicy sets field value
func (o *EabAddRequest) SetEabPolicy(v string) {
	o.EabPolicy = v
}

// GetEabValidityDuration returns the EabValidityDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EabAddRequest) GetEabValidityDuration() string {
	if o == nil || utils.IsNil(o.EabValidityDuration.Get()) {
		var ret string
		return ret
	}
	return *o.EabValidityDuration.Get()
}

// GetEabValidityDurationOk returns a tuple with the EabValidityDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EabAddRequest) GetEabValidityDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EabValidityDuration.Get(), o.EabValidityDuration.IsSet()
}

// HasEabValidityDuration returns a boolean if a field has been set.
func (o *EabAddRequest) HasEabValidityDuration() bool {
	if o != nil && o.EabValidityDuration.IsSet() {
		return true
	}

	return false
}

// SetEabValidityDuration gets a reference to the given NullableString and assigns it to the EabValidityDuration field.
func (o *EabAddRequest) SetEabValidityDuration(v string) {
	o.EabValidityDuration.Set(&v)
}

// SetEabValidityDurationNil sets the value for EabValidityDuration to be an explicit nil
func (o *EabAddRequest) SetEabValidityDurationNil() {
	o.EabValidityDuration.Set(nil)
}

// UnsetEabValidityDuration ensures that no value is present for EabValidityDuration, not even an explicit nil
func (o *EabAddRequest) UnsetEabValidityDuration() {
	o.EabValidityDuration.Unset()
}

// GetEmailConstraint returns the EmailConstraint field value if set, zero value otherwise.
func (o *EabAddRequest) GetEmailConstraint() string {
	if o == nil || utils.IsNil(o.EmailConstraint) {
		var ret string
		return ret
	}
	return *o.EmailConstraint
}

// GetEmailConstraintOk returns a tuple with the EmailConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EabAddRequest) GetEmailConstraintOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EmailConstraint) {
		return nil, false
	}
	return o.EmailConstraint, true
}

// HasEmailConstraint returns a boolean if a field has been set.
func (o *EabAddRequest) HasEmailConstraint() bool {
	if o != nil && !utils.IsNil(o.EmailConstraint) {
		return true
	}

	return false
}

// SetEmailConstraint gets a reference to the given string and assigns it to the EmailConstraint field.
func (o *EabAddRequest) SetEmailConstraint(v string) {
	o.EmailConstraint = &v
}

// GetIdentifierConstraint returns the IdentifierConstraint field value if set, zero value otherwise.
func (o *EabAddRequest) GetIdentifierConstraint() string {
	if o == nil || utils.IsNil(o.IdentifierConstraint) {
		var ret string
		return ret
	}
	return *o.IdentifierConstraint
}

// GetIdentifierConstraintOk returns a tuple with the IdentifierConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EabAddRequest) GetIdentifierConstraintOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IdentifierConstraint) {
		return nil, false
	}
	return o.IdentifierConstraint, true
}

// HasIdentifierConstraint returns a boolean if a field has been set.
func (o *EabAddRequest) HasIdentifierConstraint() bool {
	if o != nil && !utils.IsNil(o.IdentifierConstraint) {
		return true
	}

	return false
}

// SetIdentifierConstraint gets a reference to the given string and assigns it to the IdentifierConstraint field.
func (o *EabAddRequest) SetIdentifierConstraint(v string) {
	o.IdentifierConstraint = &v
}

// GetMacKeyAlgorithm returns the MacKeyAlgorithm field value
func (o *EabAddRequest) GetMacKeyAlgorithm() ExternalAccountBindingAlgorithm {
	if o == nil {
		var ret ExternalAccountBindingAlgorithm
		return ret
	}

	return o.MacKeyAlgorithm
}

// GetMacKeyAlgorithmOk returns a tuple with the MacKeyAlgorithm field value
// and a boolean to check if the value has been set.
func (o *EabAddRequest) GetMacKeyAlgorithmOk() (*ExternalAccountBindingAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacKeyAlgorithm, true
}

// SetMacKeyAlgorithm sets field value
func (o *EabAddRequest) SetMacKeyAlgorithm(v ExternalAccountBindingAlgorithm) {
	o.MacKeyAlgorithm = v
}

// GetName returns the Name field value
func (o *EabAddRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EabAddRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EabAddRequest) SetName(v string) {
	o.Name = v
}

// GetValidationMethods returns the ValidationMethods field value if set, zero value otherwise.
func (o *EabAddRequest) GetValidationMethods() []AcmeAuthorizationType {
	if o == nil || utils.IsNil(o.ValidationMethods) {
		var ret []AcmeAuthorizationType
		return ret
	}
	return o.ValidationMethods
}

// GetValidationMethodsOk returns a tuple with the ValidationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EabAddRequest) GetValidationMethodsOk() ([]AcmeAuthorizationType, bool) {
	if o == nil || utils.IsNil(o.ValidationMethods) {
		return nil, false
	}
	return o.ValidationMethods, true
}

// HasValidationMethods returns a boolean if a field has been set.
func (o *EabAddRequest) HasValidationMethods() bool {
	if o != nil && !utils.IsNil(o.ValidationMethods) {
		return true
	}

	return false
}

// SetValidationMethods gets a reference to the given []AcmeAuthorizationType and assigns it to the ValidationMethods field.
func (o *EabAddRequest) SetValidationMethods(v []AcmeAuthorizationType) {
	o.ValidationMethods = v
}

func (o EabAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EabAddRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AllowedProfiles) {
		toSerialize["allowedProfiles"] = o.AllowedProfiles
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["eabPolicy"] = o.EabPolicy
	if o.EabValidityDuration.IsSet() {
		toSerialize["eabValidityDuration"] = o.EabValidityDuration.Get()
	}
	if !utils.IsNil(o.EmailConstraint) {
		toSerialize["emailConstraint"] = o.EmailConstraint
	}
	if !utils.IsNil(o.IdentifierConstraint) {
		toSerialize["identifierConstraint"] = o.IdentifierConstraint
	}
	toSerialize["macKeyAlgorithm"] = o.MacKeyAlgorithm
	toSerialize["name"] = o.Name
	if !utils.IsNil(o.ValidationMethods) {
		toSerialize["validationMethods"] = o.ValidationMethods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EabAddRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eabPolicy",
		"macKeyAlgorithm",
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

	varEabAddRequest := _EabAddRequest{}

	err = json.Unmarshal(data, &varEabAddRequest)

	if err != nil {
		return err
	}

	*o = EabAddRequest(varEabAddRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allowedProfiles")
		delete(additionalProperties, "description")
		delete(additionalProperties, "eabPolicy")
		delete(additionalProperties, "eabValidityDuration")
		delete(additionalProperties, "emailConstraint")
		delete(additionalProperties, "identifierConstraint")
		delete(additionalProperties, "macKeyAlgorithm")
		delete(additionalProperties, "name")
		delete(additionalProperties, "validationMethods")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEabAddRequest struct {
	value *EabAddRequest
	isSet bool
}

func (v NullableEabAddRequest) Get() *EabAddRequest {
	return v.value
}

func (v *NullableEabAddRequest) Set(val *EabAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEabAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEabAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEabAddRequest(val *EabAddRequest) *NullableEabAddRequest {
	return &NullableEabAddRequest{value: val, isSet: true}
}

func (v NullableEabAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEabAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
