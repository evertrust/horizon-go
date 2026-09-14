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

// checks if the RetryParameters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RetryParameters{}

// RetryParameters Parameters controlling how failed asynchronous jobs are retried, using an exponential backoff strategy.
type RetryParameters struct {
	// Maximum number of retry attempts before the job is considered failed.
	Attempts int64 `json:"attempts"`
	// Maximum delay between two retries, capping the exponential backoff.
	MaxBackoff string `json:"maxBackoff" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Minimum delay to wait before the first retry.
	MinBackoff string `json:"minBackoff" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Random jitter factor added to each backoff delay to avoid retry storms (e.g. `0.1` adds up to 10%).
	RandomFactor         float64 `json:"randomFactor"`
	AdditionalProperties map[string]interface{}
}

type _RetryParameters RetryParameters

// NewRetryParameters instantiates a new RetryParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetryParameters(attempts int64, maxBackoff string, minBackoff string, randomFactor float64) *RetryParameters {
	this := RetryParameters{}
	this.Attempts = attempts
	this.MaxBackoff = maxBackoff
	this.MinBackoff = minBackoff
	this.RandomFactor = randomFactor
	return &this
}

// NewRetryParametersWithDefaults instantiates a new RetryParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetryParametersWithDefaults() *RetryParameters {
	this := RetryParameters{}
	return &this
}

// GetAttempts returns the Attempts field value
func (o *RetryParameters) GetAttempts() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value
// and a boolean to check if the value has been set.
func (o *RetryParameters) GetAttemptsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attempts, true
}

// SetAttempts sets field value
func (o *RetryParameters) SetAttempts(v int64) {
	o.Attempts = v
}

// GetMaxBackoff returns the MaxBackoff field value
func (o *RetryParameters) GetMaxBackoff() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxBackoff
}

// GetMaxBackoffOk returns a tuple with the MaxBackoff field value
// and a boolean to check if the value has been set.
func (o *RetryParameters) GetMaxBackoffOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxBackoff, true
}

// SetMaxBackoff sets field value
func (o *RetryParameters) SetMaxBackoff(v string) {
	o.MaxBackoff = v
}

// GetMinBackoff returns the MinBackoff field value
func (o *RetryParameters) GetMinBackoff() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MinBackoff
}

// GetMinBackoffOk returns a tuple with the MinBackoff field value
// and a boolean to check if the value has been set.
func (o *RetryParameters) GetMinBackoffOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinBackoff, true
}

// SetMinBackoff sets field value
func (o *RetryParameters) SetMinBackoff(v string) {
	o.MinBackoff = v
}

// GetRandomFactor returns the RandomFactor field value
func (o *RetryParameters) GetRandomFactor() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.RandomFactor
}

// GetRandomFactorOk returns a tuple with the RandomFactor field value
// and a boolean to check if the value has been set.
func (o *RetryParameters) GetRandomFactorOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RandomFactor, true
}

// SetRandomFactor sets field value
func (o *RetryParameters) SetRandomFactor(v float64) {
	o.RandomFactor = v
}

func (o RetryParameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetryParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attempts"] = o.Attempts
	toSerialize["maxBackoff"] = o.MaxBackoff
	toSerialize["minBackoff"] = o.MinBackoff
	toSerialize["randomFactor"] = o.RandomFactor

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RetryParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attempts",
		"maxBackoff",
		"minBackoff",
		"randomFactor",
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

	varRetryParameters := _RetryParameters{}

	err = json.Unmarshal(data, &varRetryParameters)

	if err != nil {
		return err
	}

	*o = RetryParameters(varRetryParameters)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attempts")
		delete(additionalProperties, "maxBackoff")
		delete(additionalProperties, "minBackoff")
		delete(additionalProperties, "randomFactor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRetryParameters struct {
	value *RetryParameters
	isSet bool
}

func (v NullableRetryParameters) Get() *RetryParameters {
	return v.value
}

func (v *NullableRetryParameters) Set(val *RetryParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRetryParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRetryParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetryParameters(val *RetryParameters) *NullableRetryParameters {
	return &NullableRetryParameters{value: val, isSet: true}
}

func (v NullableRetryParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetryParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
