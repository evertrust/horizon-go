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

// checks if the CompliancePolicy type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CompliancePolicy{}

// CompliancePolicy struct for CompliancePolicy
type CompliancePolicy struct {
	AuthorizedSigningAlgorithms []string `json:"authorizedSigningAlgorithms,omitempty"`
	AuthorizedCas               []string `json:"authorizedCas,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _CompliancePolicy CompliancePolicy

// NewCompliancePolicy instantiates a new CompliancePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompliancePolicy() *CompliancePolicy {
	this := CompliancePolicy{}
	return &this
}

// NewCompliancePolicyWithDefaults instantiates a new CompliancePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompliancePolicyWithDefaults() *CompliancePolicy {
	this := CompliancePolicy{}
	return &this
}

// GetAuthorizedSigningAlgorithms returns the AuthorizedSigningAlgorithms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompliancePolicy) GetAuthorizedSigningAlgorithms() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AuthorizedSigningAlgorithms
}

// GetAuthorizedSigningAlgorithmsOk returns a tuple with the AuthorizedSigningAlgorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompliancePolicy) GetAuthorizedSigningAlgorithmsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AuthorizedSigningAlgorithms) {
		return nil, false
	}
	return o.AuthorizedSigningAlgorithms, true
}

// HasAuthorizedSigningAlgorithms returns a boolean if a field has been set.
func (o *CompliancePolicy) HasAuthorizedSigningAlgorithms() bool {
	if o != nil && !utils.IsNil(o.AuthorizedSigningAlgorithms) {
		return true
	}

	return false
}

// SetAuthorizedSigningAlgorithms gets a reference to the given []string and assigns it to the AuthorizedSigningAlgorithms field.
func (o *CompliancePolicy) SetAuthorizedSigningAlgorithms(v []string) {
	o.AuthorizedSigningAlgorithms = v
}

// GetAuthorizedCas returns the AuthorizedCas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompliancePolicy) GetAuthorizedCas() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AuthorizedCas
}

// GetAuthorizedCasOk returns a tuple with the AuthorizedCas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompliancePolicy) GetAuthorizedCasOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AuthorizedCas) {
		return nil, false
	}
	return o.AuthorizedCas, true
}

// HasAuthorizedCas returns a boolean if a field has been set.
func (o *CompliancePolicy) HasAuthorizedCas() bool {
	if o != nil && !utils.IsNil(o.AuthorizedCas) {
		return true
	}

	return false
}

// SetAuthorizedCas gets a reference to the given []string and assigns it to the AuthorizedCas field.
func (o *CompliancePolicy) SetAuthorizedCas(v []string) {
	o.AuthorizedCas = v
}

func (o CompliancePolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompliancePolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorizedSigningAlgorithms != nil {
		toSerialize["authorizedSigningAlgorithms"] = o.AuthorizedSigningAlgorithms
	}
	if o.AuthorizedCas != nil {
		toSerialize["authorizedCas"] = o.AuthorizedCas
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompliancePolicy) UnmarshalJSON(data []byte) (err error) {
	varCompliancePolicy := _CompliancePolicy{}

	err = json.Unmarshal(data, &varCompliancePolicy)

	if err != nil {
		return err
	}

	*o = CompliancePolicy(varCompliancePolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorizedSigningAlgorithms")
		delete(additionalProperties, "authorizedCas")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompliancePolicy struct {
	value *CompliancePolicy
	isSet bool
}

func (v NullableCompliancePolicy) Get() *CompliancePolicy {
	return v.value
}

func (v *NullableCompliancePolicy) Set(val *CompliancePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableCompliancePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableCompliancePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompliancePolicy(val *CompliancePolicy) *NullableCompliancePolicy {
	return &NullableCompliancePolicy{value: val, isSet: true}
}

func (v NullableCompliancePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompliancePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
