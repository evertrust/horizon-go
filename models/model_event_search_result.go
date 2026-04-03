/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the EventSearchResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EventSearchResult{}

// EventSearchResult struct for EventSearchResult
type EventSearchResult struct {
	Id                   utils.NullableString `json:"_id,omitempty"`
	Code                 NullableEventCode    `json:"code,omitempty"`
	Details              []EventDetail        `json:"details,omitempty"`
	Module               NullableEventModule  `json:"module,omitempty"`
	Node                 utils.NullableString `json:"node,omitempty"`
	RemoveAt             utils.NullableInt64  `json:"removeAt,omitempty"`
	Seal                 utils.NullableString `json:"seal,omitempty"`
	Status               utils.NullableString `json:"status,omitempty"`
	Timestamp            utils.NullableInt64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EventSearchResult EventSearchResult

// NewEventSearchResult instantiates a new EventSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSearchResult() *EventSearchResult {
	this := EventSearchResult{}
	return &this
}

// NewEventSearchResultWithDefaults instantiates a new EventSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSearchResultWithDefaults() *EventSearchResult {
	this := EventSearchResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchResult) GetId() string {
	if o == nil || utils.IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *EventSearchResult) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *EventSearchResult) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *EventSearchResult) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *EventSearchResult) UnsetId() {
	o.Id.Unset()
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchResult) GetCode() EventCode {
	if o == nil || utils.IsNil(o.Code.Get()) {
		var ret EventCode
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchResult) GetCodeOk() (*EventCode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *EventSearchResult) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableEventCode and assigns it to the Code field.
func (o *EventSearchResult) SetCode(v EventCode) {
	o.Code.Set(&v)
}

// SetCodeNil sets the value for Code to be an explicit nil
func (o *EventSearchResult) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *EventSearchResult) UnsetCode() {
	o.Code.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchResult) GetDetails() []EventDetail {
	if o == nil {
		var ret []EventDetail
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchResult) GetDetailsOk() ([]EventDetail, bool) {
	if o == nil || utils.IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *EventSearchResult) HasDetails() bool {
	if o != nil && !utils.IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []EventDetail and assigns it to the Details field.
func (o *EventSearchResult) SetDetails(v []EventDetail) {
	o.Details = v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchResult) GetModule() EventModule {
	if o == nil || utils.IsNil(o.Module.Get()) {
		var ret EventModule
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchResult) GetModuleOk() (*EventModule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *EventSearchResult) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableEventModule and assigns it to the Module field.
func (o *EventSearchResult) SetModule(v EventModule) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *EventSearchResult) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *EventSearchResult) UnsetModule() {
	o.Module.Unset()
}

// GetNode returns the Node field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchResult) GetNode() string {
	if o == nil || utils.IsNil(o.Node.Get()) {
		var ret string
		return ret
	}
	return *o.Node.Get()
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchResult) GetNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Node.Get(), o.Node.IsSet()
}

// HasNode returns a boolean if a field has been set.
func (o *EventSearchResult) HasNode() bool {
	if o != nil && o.Node.IsSet() {
		return true
	}

	return false
}

// SetNode gets a reference to the given NullableString and assigns it to the Node field.
func (o *EventSearchResult) SetNode(v string) {
	o.Node.Set(&v)
}

// SetNodeNil sets the value for Node to be an explicit nil
func (o *EventSearchResult) SetNodeNil() {
	o.Node.Set(nil)
}

// UnsetNode ensures that no value is present for Node, not even an explicit nil
func (o *EventSearchResult) UnsetNode() {
	o.Node.Unset()
}

// GetRemoveAt returns the RemoveAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchResult) GetRemoveAt() int64 {
	if o == nil || utils.IsNil(o.RemoveAt.Get()) {
		var ret int64
		return ret
	}
	return *o.RemoveAt.Get()
}

// GetRemoveAtOk returns a tuple with the RemoveAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchResult) GetRemoveAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveAt.Get(), o.RemoveAt.IsSet()
}

// HasRemoveAt returns a boolean if a field has been set.
func (o *EventSearchResult) HasRemoveAt() bool {
	if o != nil && o.RemoveAt.IsSet() {
		return true
	}

	return false
}

// SetRemoveAt gets a reference to the given NullableInt64 and assigns it to the RemoveAt field.
func (o *EventSearchResult) SetRemoveAt(v int64) {
	o.RemoveAt.Set(&v)
}

// SetRemoveAtNil sets the value for RemoveAt to be an explicit nil
func (o *EventSearchResult) SetRemoveAtNil() {
	o.RemoveAt.Set(nil)
}

// UnsetRemoveAt ensures that no value is present for RemoveAt, not even an explicit nil
func (o *EventSearchResult) UnsetRemoveAt() {
	o.RemoveAt.Unset()
}

// GetSeal returns the Seal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchResult) GetSeal() string {
	if o == nil || utils.IsNil(o.Seal.Get()) {
		var ret string
		return ret
	}
	return *o.Seal.Get()
}

// GetSealOk returns a tuple with the Seal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchResult) GetSealOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Seal.Get(), o.Seal.IsSet()
}

// HasSeal returns a boolean if a field has been set.
func (o *EventSearchResult) HasSeal() bool {
	if o != nil && o.Seal.IsSet() {
		return true
	}

	return false
}

// SetSeal gets a reference to the given NullableString and assigns it to the Seal field.
func (o *EventSearchResult) SetSeal(v string) {
	o.Seal.Set(&v)
}

// SetSealNil sets the value for Seal to be an explicit nil
func (o *EventSearchResult) SetSealNil() {
	o.Seal.Set(nil)
}

// UnsetSeal ensures that no value is present for Seal, not even an explicit nil
func (o *EventSearchResult) UnsetSeal() {
	o.Seal.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchResult) GetStatus() string {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *EventSearchResult) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *EventSearchResult) SetStatus(v string) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *EventSearchResult) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *EventSearchResult) UnsetStatus() {
	o.Status.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventSearchResult) GetTimestamp() int64 {
	if o == nil || utils.IsNil(o.Timestamp.Get()) {
		var ret int64
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventSearchResult) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *EventSearchResult) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableInt64 and assigns it to the Timestamp field.
func (o *EventSearchResult) SetTimestamp(v int64) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *EventSearchResult) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *EventSearchResult) UnsetTimestamp() {
	o.Timestamp.Unset()
}

func (o EventSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["_id"] = o.Id.Get()
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	if o.Node.IsSet() {
		toSerialize["node"] = o.Node.Get()
	}
	if o.RemoveAt.IsSet() {
		toSerialize["removeAt"] = o.RemoveAt.Get()
	}
	if o.Seal.IsSet() {
		toSerialize["seal"] = o.Seal.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventSearchResult) UnmarshalJSON(data []byte) (err error) {
	varEventSearchResult := _EventSearchResult{}

	err = json.Unmarshal(data, &varEventSearchResult)

	if err != nil {
		return err
	}

	*o = EventSearchResult(varEventSearchResult)

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

type NullableEventSearchResult struct {
	value *EventSearchResult
	isSet bool
}

func (v NullableEventSearchResult) Get() *EventSearchResult {
	return v.value
}

func (v *NullableEventSearchResult) Set(val *EventSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSearchResult(val *EventSearchResult) *NullableEventSearchResult {
	return &NullableEventSearchResult{value: val, isSet: true}
}

func (v NullableEventSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
