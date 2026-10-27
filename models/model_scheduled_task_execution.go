/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in either the `X-API-TOKEN` header or the `Authorization` header using the `Bearer` scheme:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"Authorization: Bearer eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the ScheduledTaskExecution type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ScheduledTaskExecution{}

// ScheduledTaskExecution Fields written by the last run of the task. They are returned by the API and cannot be set by a caller.
type ScheduledTaskExecution struct {
	Detail               *string `json:"detail,omitempty"`
	ExecutionId          *string `json:"executionId,omitempty"`
	Host                 *string `json:"host,omitempty"`
	LastCompletionDate   *int64  `json:"lastCompletionDate,omitempty"`
	LastExecutionDate    *int64  `json:"lastExecutionDate,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScheduledTaskExecution ScheduledTaskExecution

// NewScheduledTaskExecution instantiates a new ScheduledTaskExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledTaskExecution() *ScheduledTaskExecution {
	this := ScheduledTaskExecution{}
	return &this
}

// NewScheduledTaskExecutionWithDefaults instantiates a new ScheduledTaskExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledTaskExecutionWithDefaults() *ScheduledTaskExecution {
	this := ScheduledTaskExecution{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ScheduledTaskExecution) GetDetail() string {
	if o == nil || utils.IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskExecution) GetDetailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ScheduledTaskExecution) HasDetail() bool {
	if o != nil && !utils.IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ScheduledTaskExecution) SetDetail(v string) {
	o.Detail = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *ScheduledTaskExecution) GetExecutionId() string {
	if o == nil || utils.IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskExecution) GetExecutionIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *ScheduledTaskExecution) HasExecutionId() bool {
	if o != nil && !utils.IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *ScheduledTaskExecution) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *ScheduledTaskExecution) GetHost() string {
	if o == nil || utils.IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskExecution) GetHostOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *ScheduledTaskExecution) HasHost() bool {
	if o != nil && !utils.IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *ScheduledTaskExecution) SetHost(v string) {
	o.Host = &v
}

// GetLastCompletionDate returns the LastCompletionDate field value if set, zero value otherwise.
func (o *ScheduledTaskExecution) GetLastCompletionDate() int64 {
	if o == nil || utils.IsNil(o.LastCompletionDate) {
		var ret int64
		return ret
	}
	return *o.LastCompletionDate
}

// GetLastCompletionDateOk returns a tuple with the LastCompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskExecution) GetLastCompletionDateOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.LastCompletionDate) {
		return nil, false
	}
	return o.LastCompletionDate, true
}

// HasLastCompletionDate returns a boolean if a field has been set.
func (o *ScheduledTaskExecution) HasLastCompletionDate() bool {
	if o != nil && !utils.IsNil(o.LastCompletionDate) {
		return true
	}

	return false
}

// SetLastCompletionDate gets a reference to the given int64 and assigns it to the LastCompletionDate field.
func (o *ScheduledTaskExecution) SetLastCompletionDate(v int64) {
	o.LastCompletionDate = &v
}

// GetLastExecutionDate returns the LastExecutionDate field value if set, zero value otherwise.
func (o *ScheduledTaskExecution) GetLastExecutionDate() int64 {
	if o == nil || utils.IsNil(o.LastExecutionDate) {
		var ret int64
		return ret
	}
	return *o.LastExecutionDate
}

// GetLastExecutionDateOk returns a tuple with the LastExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskExecution) GetLastExecutionDateOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.LastExecutionDate) {
		return nil, false
	}
	return o.LastExecutionDate, true
}

// HasLastExecutionDate returns a boolean if a field has been set.
func (o *ScheduledTaskExecution) HasLastExecutionDate() bool {
	if o != nil && !utils.IsNil(o.LastExecutionDate) {
		return true
	}

	return false
}

// SetLastExecutionDate gets a reference to the given int64 and assigns it to the LastExecutionDate field.
func (o *ScheduledTaskExecution) SetLastExecutionDate(v int64) {
	o.LastExecutionDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ScheduledTaskExecution) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskExecution) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ScheduledTaskExecution) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ScheduledTaskExecution) SetStatus(v string) {
	o.Status = &v
}

func (o ScheduledTaskExecution) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduledTaskExecution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !utils.IsNil(o.ExecutionId) {
		toSerialize["executionId"] = o.ExecutionId
	}
	if !utils.IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !utils.IsNil(o.LastCompletionDate) {
		toSerialize["lastCompletionDate"] = o.LastCompletionDate
	}
	if !utils.IsNil(o.LastExecutionDate) {
		toSerialize["lastExecutionDate"] = o.LastExecutionDate
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScheduledTaskExecution) UnmarshalJSON(data []byte) (err error) {
	varScheduledTaskExecution := _ScheduledTaskExecution{}

	err = json.Unmarshal(data, &varScheduledTaskExecution)

	if err != nil {
		return err
	}

	*o = ScheduledTaskExecution(varScheduledTaskExecution)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		delete(additionalProperties, "executionId")
		delete(additionalProperties, "host")
		delete(additionalProperties, "lastCompletionDate")
		delete(additionalProperties, "lastExecutionDate")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduledTaskExecution struct {
	value *ScheduledTaskExecution
	isSet bool
}

func (v NullableScheduledTaskExecution) Get() *ScheduledTaskExecution {
	return v.value
}

func (v *NullableScheduledTaskExecution) Set(val *ScheduledTaskExecution) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledTaskExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledTaskExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledTaskExecution(val *ScheduledTaskExecution) *NullableScheduledTaskExecution {
	return &NullableScheduledTaskExecution{value: val, isSet: true}
}

func (v NullableScheduledTaskExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledTaskExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
