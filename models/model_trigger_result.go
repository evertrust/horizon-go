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

// checks if the TriggerResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TriggerResult{}

// TriggerResult struct for TriggerResult
type TriggerResult struct {
	// Contains details on this trigger's execution
	Detail utils.NullableString `json:"detail,omitempty"`
	// The event that triggered the trigger
	Event string `json:"event"`
	// The last time this trigger was executed for this certificate and this event
	LastExecutionDate int64 `json:"lastExecutionDate"`
	// The name of the trigger that was executed
	Name string `json:"name"`
	// Time that will be waited between the next and the next+1 execution of this trigger
	NextDelay utils.NullableString `json:"nextDelay,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// The next scheduled execution time for this trigger
	NextExecutionDate utils.NullableInt64 `json:"nextExecutionDate,omitempty"`
	// The number of remaining tries before the trigger is abandoned
	Retries utils.NullableInt64 `json:"retries,omitempty"`
	// Is this trigger manually retryable (can be [run](#tag/certificate/operation/certificate.run))
	Retryable bool `json:"retryable"`
	// The status of the trigger after its execution
	Status string `json:"status"`
	// The type of the trigger
	TriggerType          string `json:"triggerType"`
	AdditionalProperties map[string]interface{}
}

type _TriggerResult TriggerResult

// NewTriggerResult instantiates a new TriggerResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerResult(event string, lastExecutionDate int64, name string, retryable bool, status string, triggerType string) *TriggerResult {
	this := TriggerResult{}
	this.Event = event
	this.LastExecutionDate = lastExecutionDate
	this.Name = name
	this.Retryable = retryable
	this.Status = status
	this.TriggerType = triggerType
	return &this
}

// NewTriggerResultWithDefaults instantiates a new TriggerResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerResultWithDefaults() *TriggerResult {
	this := TriggerResult{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerResult) GetDetail() string {
	if o == nil || utils.IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerResult) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *TriggerResult) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *TriggerResult) SetDetail(v string) {
	o.Detail.Set(&v)
}

// SetDetailNil sets the value for Detail to be an explicit nil
func (o *TriggerResult) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *TriggerResult) UnsetDetail() {
	o.Detail.Unset()
}

// GetEvent returns the Event field value
func (o *TriggerResult) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *TriggerResult) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *TriggerResult) SetEvent(v string) {
	o.Event = v
}

// GetLastExecutionDate returns the LastExecutionDate field value
func (o *TriggerResult) GetLastExecutionDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LastExecutionDate
}

// GetLastExecutionDateOk returns a tuple with the LastExecutionDate field value
// and a boolean to check if the value has been set.
func (o *TriggerResult) GetLastExecutionDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastExecutionDate, true
}

// SetLastExecutionDate sets field value
func (o *TriggerResult) SetLastExecutionDate(v int64) {
	o.LastExecutionDate = v
}

// GetName returns the Name field value
func (o *TriggerResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TriggerResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TriggerResult) SetName(v string) {
	o.Name = v
}

// GetNextDelay returns the NextDelay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerResult) GetNextDelay() string {
	if o == nil || utils.IsNil(o.NextDelay.Get()) {
		var ret string
		return ret
	}
	return *o.NextDelay.Get()
}

// GetNextDelayOk returns a tuple with the NextDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerResult) GetNextDelayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextDelay.Get(), o.NextDelay.IsSet()
}

// HasNextDelay returns a boolean if a field has been set.
func (o *TriggerResult) HasNextDelay() bool {
	if o != nil && o.NextDelay.IsSet() {
		return true
	}

	return false
}

// SetNextDelay gets a reference to the given NullableString and assigns it to the NextDelay field.
func (o *TriggerResult) SetNextDelay(v string) {
	o.NextDelay.Set(&v)
}

// SetNextDelayNil sets the value for NextDelay to be an explicit nil
func (o *TriggerResult) SetNextDelayNil() {
	o.NextDelay.Set(nil)
}

// UnsetNextDelay ensures that no value is present for NextDelay, not even an explicit nil
func (o *TriggerResult) UnsetNextDelay() {
	o.NextDelay.Unset()
}

// GetNextExecutionDate returns the NextExecutionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerResult) GetNextExecutionDate() int64 {
	if o == nil || utils.IsNil(o.NextExecutionDate.Get()) {
		var ret int64
		return ret
	}
	return *o.NextExecutionDate.Get()
}

// GetNextExecutionDateOk returns a tuple with the NextExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerResult) GetNextExecutionDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextExecutionDate.Get(), o.NextExecutionDate.IsSet()
}

// HasNextExecutionDate returns a boolean if a field has been set.
func (o *TriggerResult) HasNextExecutionDate() bool {
	if o != nil && o.NextExecutionDate.IsSet() {
		return true
	}

	return false
}

// SetNextExecutionDate gets a reference to the given NullableInt64 and assigns it to the NextExecutionDate field.
func (o *TriggerResult) SetNextExecutionDate(v int64) {
	o.NextExecutionDate.Set(&v)
}

// SetNextExecutionDateNil sets the value for NextExecutionDate to be an explicit nil
func (o *TriggerResult) SetNextExecutionDateNil() {
	o.NextExecutionDate.Set(nil)
}

// UnsetNextExecutionDate ensures that no value is present for NextExecutionDate, not even an explicit nil
func (o *TriggerResult) UnsetNextExecutionDate() {
	o.NextExecutionDate.Unset()
}

// GetRetries returns the Retries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TriggerResult) GetRetries() int64 {
	if o == nil || utils.IsNil(o.Retries.Get()) {
		var ret int64
		return ret
	}
	return *o.Retries.Get()
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TriggerResult) GetRetriesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Retries.Get(), o.Retries.IsSet()
}

// HasRetries returns a boolean if a field has been set.
func (o *TriggerResult) HasRetries() bool {
	if o != nil && o.Retries.IsSet() {
		return true
	}

	return false
}

// SetRetries gets a reference to the given NullableInt64 and assigns it to the Retries field.
func (o *TriggerResult) SetRetries(v int64) {
	o.Retries.Set(&v)
}

// SetRetriesNil sets the value for Retries to be an explicit nil
func (o *TriggerResult) SetRetriesNil() {
	o.Retries.Set(nil)
}

// UnsetRetries ensures that no value is present for Retries, not even an explicit nil
func (o *TriggerResult) UnsetRetries() {
	o.Retries.Unset()
}

// GetRetryable returns the Retryable field value
func (o *TriggerResult) GetRetryable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Retryable
}

// GetRetryableOk returns a tuple with the Retryable field value
// and a boolean to check if the value has been set.
func (o *TriggerResult) GetRetryableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Retryable, true
}

// SetRetryable sets field value
func (o *TriggerResult) SetRetryable(v bool) {
	o.Retryable = v
}

// GetStatus returns the Status field value
func (o *TriggerResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TriggerResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TriggerResult) SetStatus(v string) {
	o.Status = v
}

// GetTriggerType returns the TriggerType field value
func (o *TriggerResult) GetTriggerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerType
}

// GetTriggerTypeOk returns a tuple with the TriggerType field value
// and a boolean to check if the value has been set.
func (o *TriggerResult) GetTriggerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerType, true
}

// SetTriggerType sets field value
func (o *TriggerResult) SetTriggerType(v string) {
	o.TriggerType = v
}

func (o TriggerResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Detail.IsSet() {
		toSerialize["detail"] = o.Detail.Get()
	}
	toSerialize["event"] = o.Event
	toSerialize["lastExecutionDate"] = o.LastExecutionDate
	toSerialize["name"] = o.Name
	if o.NextDelay.IsSet() {
		toSerialize["nextDelay"] = o.NextDelay.Get()
	}
	if o.NextExecutionDate.IsSet() {
		toSerialize["nextExecutionDate"] = o.NextExecutionDate.Get()
	}
	if o.Retries.IsSet() {
		toSerialize["retries"] = o.Retries.Get()
	}
	toSerialize["retryable"] = o.Retryable
	toSerialize["status"] = o.Status
	toSerialize["triggerType"] = o.TriggerType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TriggerResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event",
		"lastExecutionDate",
		"name",
		"retryable",
		"status",
		"triggerType",
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

	varTriggerResult := _TriggerResult{}

	err = json.Unmarshal(data, &varTriggerResult)

	if err != nil {
		return err
	}

	*o = TriggerResult(varTriggerResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detail")
		delete(additionalProperties, "event")
		delete(additionalProperties, "lastExecutionDate")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nextDelay")
		delete(additionalProperties, "nextExecutionDate")
		delete(additionalProperties, "retries")
		delete(additionalProperties, "retryable")
		delete(additionalProperties, "status")
		delete(additionalProperties, "triggerType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerResult struct {
	value *TriggerResult
	isSet bool
}

func (v NullableTriggerResult) Get() *TriggerResult {
	return v.value
}

func (v *NullableTriggerResult) Set(val *TriggerResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerResult(val *TriggerResult) *NullableTriggerResult {
	return &NullableTriggerResult{value: val, isSet: true}
}

func (v NullableTriggerResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
