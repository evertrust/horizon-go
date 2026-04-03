/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the EventResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EventResponse{}

// EventResponse struct for EventResponse
type EventResponse struct {
	Id                   utils.NullableString `json:"_id,omitempty"`
	Code                 EventCode            `json:"code"`
	Details              []EventDetail        `json:"details,omitempty"`
	Module               EventModule          `json:"module"`
	Node                 string               `json:"node"`
	RemoveAt             utils.NullableInt64  `json:"removeAt,omitempty"`
	Seal                 utils.NullableString `json:"seal,omitempty"`
	Status               string               `json:"status"`
	Timestamp            int64                `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _EventResponse EventResponse

// NewEventResponse instantiates a new EventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventResponse(code EventCode, module EventModule, node string, status string, timestamp int64) *EventResponse {
	this := EventResponse{}
	this.Code = code
	this.Module = module
	this.Node = node
	this.Status = status
	this.Timestamp = timestamp
	return &this
}

// NewEventResponseWithDefaults instantiates a new EventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventResponseWithDefaults() *EventResponse {
	this := EventResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponse) GetId() string {
	if o == nil || utils.IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *EventResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *EventResponse) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *EventResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *EventResponse) UnsetId() {
	o.Id.Unset()
}

// GetCode returns the Code field value
func (o *EventResponse) GetCode() EventCode {
	if o == nil {
		var ret EventCode
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *EventResponse) GetCodeOk() (*EventCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *EventResponse) SetCode(v EventCode) {
	o.Code = v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponse) GetDetails() []EventDetail {
	if o == nil {
		var ret []EventDetail
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponse) GetDetailsOk() ([]EventDetail, bool) {
	if o == nil || utils.IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *EventResponse) HasDetails() bool {
	if o != nil && !utils.IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []EventDetail and assigns it to the Details field.
func (o *EventResponse) SetDetails(v []EventDetail) {
	o.Details = v
}

// GetModule returns the Module field value
func (o *EventResponse) GetModule() EventModule {
	if o == nil {
		var ret EventModule
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *EventResponse) GetModuleOk() (*EventModule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *EventResponse) SetModule(v EventModule) {
	o.Module = v
}

// GetNode returns the Node field value
func (o *EventResponse) GetNode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *EventResponse) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value
func (o *EventResponse) SetNode(v string) {
	o.Node = v
}

// GetRemoveAt returns the RemoveAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponse) GetRemoveAt() int64 {
	if o == nil || utils.IsNil(o.RemoveAt.Get()) {
		var ret int64
		return ret
	}
	return *o.RemoveAt.Get()
}

// GetRemoveAtOk returns a tuple with the RemoveAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponse) GetRemoveAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveAt.Get(), o.RemoveAt.IsSet()
}

// HasRemoveAt returns a boolean if a field has been set.
func (o *EventResponse) HasRemoveAt() bool {
	if o != nil && o.RemoveAt.IsSet() {
		return true
	}

	return false
}

// SetRemoveAt gets a reference to the given NullableInt64 and assigns it to the RemoveAt field.
func (o *EventResponse) SetRemoveAt(v int64) {
	o.RemoveAt.Set(&v)
}

// SetRemoveAtNil sets the value for RemoveAt to be an explicit nil
func (o *EventResponse) SetRemoveAtNil() {
	o.RemoveAt.Set(nil)
}

// UnsetRemoveAt ensures that no value is present for RemoveAt, not even an explicit nil
func (o *EventResponse) UnsetRemoveAt() {
	o.RemoveAt.Unset()
}

// GetSeal returns the Seal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventResponse) GetSeal() string {
	if o == nil || utils.IsNil(o.Seal.Get()) {
		var ret string
		return ret
	}
	return *o.Seal.Get()
}

// GetSealOk returns a tuple with the Seal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventResponse) GetSealOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Seal.Get(), o.Seal.IsSet()
}

// HasSeal returns a boolean if a field has been set.
func (o *EventResponse) HasSeal() bool {
	if o != nil && o.Seal.IsSet() {
		return true
	}

	return false
}

// SetSeal gets a reference to the given NullableString and assigns it to the Seal field.
func (o *EventResponse) SetSeal(v string) {
	o.Seal.Set(&v)
}

// SetSealNil sets the value for Seal to be an explicit nil
func (o *EventResponse) SetSealNil() {
	o.Seal.Set(nil)
}

// UnsetSeal ensures that no value is present for Seal, not even an explicit nil
func (o *EventResponse) UnsetSeal() {
	o.Seal.Unset()
}

// GetStatus returns the Status field value
func (o *EventResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EventResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EventResponse) SetStatus(v string) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value
func (o *EventResponse) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *EventResponse) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *EventResponse) SetTimestamp(v int64) {
	o.Timestamp = v
}

func (o EventResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["_id"] = o.Id.Get()
	}
	toSerialize["code"] = o.Code
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	toSerialize["module"] = o.Module
	toSerialize["node"] = o.Node
	if o.RemoveAt.IsSet() {
		toSerialize["removeAt"] = o.RemoveAt.Get()
	}
	if o.Seal.IsSet() {
		toSerialize["seal"] = o.Seal.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"module",
		"node",
		"status",
		"timestamp",
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

	varEventResponse := _EventResponse{}

	err = json.Unmarshal(data, &varEventResponse)

	if err != nil {
		return err
	}

	*o = EventResponse(varEventResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "code")
		delete(additionalProperties, "details")
		delete(additionalProperties, "module")
		delete(additionalProperties, "node")
		delete(additionalProperties, "removeAt")
		delete(additionalProperties, "seal")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventResponse struct {
	value *EventResponse
	isSet bool
}

func (v NullableEventResponse) Get() *EventResponse {
	return v.value
}

func (v *NullableEventResponse) Set(val *EventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventResponse(val *EventResponse) *NullableEventResponse {
	return &NullableEventResponse{value: val, isSet: true}
}

func (v NullableEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
