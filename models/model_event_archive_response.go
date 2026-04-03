/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the EventArchiveResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EventArchiveResponse{}

// EventArchiveResponse struct for EventArchiveResponse
type EventArchiveResponse struct {
	// Object internal ID
	Id        string               `json:"_id"`
	Count     utils.NullableInt64  `json:"count,omitempty"`
	CreatedAt utils.NullableInt64  `json:"createdAt,omitempty"`
	Error     utils.NullableString `json:"error,omitempty"`
	PurgeAt   utils.NullableInt64  `json:"purgeAt,omitempty"`
	Status    ArchiveStatus        `json:"status"`
	// Date before which all events will be archived
	Before               int64  `json:"before"`
	Filename             string `json:"filename"`
	Name                 string `json:"name"`
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _EventArchiveResponse EventArchiveResponse

// NewEventArchiveResponse instantiates a new EventArchiveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventArchiveResponse(id string, status ArchiveStatus, before int64, filename string, name string, type_ string) *EventArchiveResponse {
	this := EventArchiveResponse{}
	this.Before = before
	this.Filename = filename
	this.Name = name
	this.Type = type_
	return &this
}

// NewEventArchiveResponseWithDefaults instantiates a new EventArchiveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventArchiveResponseWithDefaults() *EventArchiveResponse {
	this := EventArchiveResponse{}
	return &this
}

// GetId returns the Id field value
func (o *EventArchiveResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventArchiveResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EventArchiveResponse) SetId(v string) {
	o.Id = v
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventArchiveResponse) GetCount() int64 {
	if o == nil || utils.IsNil(o.Count.Get()) {
		var ret int64
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventArchiveResponse) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *EventArchiveResponse) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt64 and assigns it to the Count field.
func (o *EventArchiveResponse) SetCount(v int64) {
	o.Count.Set(&v)
}

// SetCountNil sets the value for Count to be an explicit nil
func (o *EventArchiveResponse) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *EventArchiveResponse) UnsetCount() {
	o.Count.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventArchiveResponse) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventArchiveResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EventArchiveResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *EventArchiveResponse) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *EventArchiveResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *EventArchiveResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventArchiveResponse) GetError() string {
	if o == nil || utils.IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventArchiveResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *EventArchiveResponse) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *EventArchiveResponse) SetError(v string) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *EventArchiveResponse) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *EventArchiveResponse) UnsetError() {
	o.Error.Unset()
}

// GetPurgeAt returns the PurgeAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventArchiveResponse) GetPurgeAt() int64 {
	if o == nil || utils.IsNil(o.PurgeAt.Get()) {
		var ret int64
		return ret
	}
	return *o.PurgeAt.Get()
}

// GetPurgeAtOk returns a tuple with the PurgeAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventArchiveResponse) GetPurgeAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PurgeAt.Get(), o.PurgeAt.IsSet()
}

// HasPurgeAt returns a boolean if a field has been set.
func (o *EventArchiveResponse) HasPurgeAt() bool {
	if o != nil && o.PurgeAt.IsSet() {
		return true
	}

	return false
}

// SetPurgeAt gets a reference to the given NullableInt64 and assigns it to the PurgeAt field.
func (o *EventArchiveResponse) SetPurgeAt(v int64) {
	o.PurgeAt.Set(&v)
}

// SetPurgeAtNil sets the value for PurgeAt to be an explicit nil
func (o *EventArchiveResponse) SetPurgeAtNil() {
	o.PurgeAt.Set(nil)
}

// UnsetPurgeAt ensures that no value is present for PurgeAt, not even an explicit nil
func (o *EventArchiveResponse) UnsetPurgeAt() {
	o.PurgeAt.Unset()
}

// GetStatus returns the Status field value
func (o *EventArchiveResponse) GetStatus() ArchiveStatus {
	if o == nil {
		var ret ArchiveStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EventArchiveResponse) GetStatusOk() (*ArchiveStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EventArchiveResponse) SetStatus(v ArchiveStatus) {
	o.Status = v
}

// GetBefore returns the Before field value
func (o *EventArchiveResponse) GetBefore() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Before
}

// GetBeforeOk returns a tuple with the Before field value
// and a boolean to check if the value has been set.
func (o *EventArchiveResponse) GetBeforeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Before, true
}

// SetBefore sets field value
func (o *EventArchiveResponse) SetBefore(v int64) {
	o.Before = v
}

// GetFilename returns the Filename field value
func (o *EventArchiveResponse) GetFilename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value
// and a boolean to check if the value has been set.
func (o *EventArchiveResponse) GetFilenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filename, true
}

// SetFilename sets field value
func (o *EventArchiveResponse) SetFilename(v string) {
	o.Filename = v
}

// GetName returns the Name field value
func (o *EventArchiveResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventArchiveResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventArchiveResponse) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *EventArchiveResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventArchiveResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventArchiveResponse) SetType(v string) {
	o.Type = v
}

func (o EventArchiveResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventArchiveResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["createdAt"] = o.CreatedAt.Get()
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.PurgeAt.IsSet() {
		toSerialize["purgeAt"] = o.PurgeAt.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["before"] = o.Before
	toSerialize["filename"] = o.Filename
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EventArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"status",
		"before",
		"filename",
		"name",
		"type",
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

	varEventArchiveResponse := _EventArchiveResponse{}

	err = json.Unmarshal(data, &varEventArchiveResponse)

	if err != nil {
		return err
	}

	*o = EventArchiveResponse(varEventArchiveResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "count")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "error")
		delete(additionalProperties, "purgeAt")
		delete(additionalProperties, "status")
		delete(additionalProperties, "before")
		delete(additionalProperties, "filename")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEventArchiveResponse struct {
	value *EventArchiveResponse
	isSet bool
}

func (v NullableEventArchiveResponse) Get() *EventArchiveResponse {
	return v.value
}

func (v *NullableEventArchiveResponse) Set(val *EventArchiveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEventArchiveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEventArchiveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventArchiveResponse(val *EventArchiveResponse) *NullableEventArchiveResponse {
	return &NullableEventArchiveResponse{value: val, isSet: true}
}

func (v NullableEventArchiveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventArchiveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
