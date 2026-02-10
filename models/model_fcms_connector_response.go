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

// checks if the FCMSConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FCMSConnectorResponse{}

// FCMSConnectorResponse struct for FCMSConnectorResponse
type FCMSConnectorResponse struct {
	// Object internal ID
	Id       string `json:"_id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	EndPoint string `json:"endPoint"`
	// Name of the `raw` [credentials](#tag/security.credentials) containing the API key to authenticate on the PKI
	ApiCredentials         string                     `json:"apiCredentials"`
	TemplateId             int64                      `json:"templateId"`
	DefaultOwner           string                     `json:"defaultOwner"`
	AuthenticationDomainId int64                      `json:"authenticationDomainId"`
	OwnerGroups            utils.NullableString       `json:"ownerGroups,omitempty"`
	DeleteOnRevoke         bool                       `json:"deleteOnRevoke"`
	Timeout                utils.NullableString       `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Proxy                  utils.NullableString       `json:"proxy,omitempty"`
	Queue                  utils.NullableString       `json:"queue,omitempty"`
	Status                 NullablePKIConnectorStatus `json:"status,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _FCMSConnectorResponse FCMSConnectorResponse

// NewFCMSConnectorResponse instantiates a new FCMSConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFCMSConnectorResponse(id string, name string, type_ string, endPoint string, apiCredentials string, templateId int64, defaultOwner string, authenticationDomainId int64, deleteOnRevoke bool) *FCMSConnectorResponse {
	this := FCMSConnectorResponse{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.EndPoint = endPoint
	this.ApiCredentials = apiCredentials
	this.TemplateId = templateId
	this.DefaultOwner = defaultOwner
	this.AuthenticationDomainId = authenticationDomainId
	this.DeleteOnRevoke = deleteOnRevoke
	return &this
}

// NewFCMSConnectorResponseWithDefaults instantiates a new FCMSConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFCMSConnectorResponseWithDefaults() *FCMSConnectorResponse {
	this := FCMSConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *FCMSConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FCMSConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FCMSConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FCMSConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FCMSConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FCMSConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *FCMSConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FCMSConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FCMSConnectorResponse) SetType(v string) {
	o.Type = v
}

// GetEndPoint returns the EndPoint field value
func (o *FCMSConnectorResponse) GetEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value
// and a boolean to check if the value has been set.
func (o *FCMSConnectorResponse) GetEndPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPoint, true
}

// SetEndPoint sets field value
func (o *FCMSConnectorResponse) SetEndPoint(v string) {
	o.EndPoint = v
}

// GetApiCredentials returns the ApiCredentials field value
func (o *FCMSConnectorResponse) GetApiCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiCredentials
}

// GetApiCredentialsOk returns a tuple with the ApiCredentials field value
// and a boolean to check if the value has been set.
func (o *FCMSConnectorResponse) GetApiCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiCredentials, true
}

// SetApiCredentials sets field value
func (o *FCMSConnectorResponse) SetApiCredentials(v string) {
	o.ApiCredentials = v
}

// GetTemplateId returns the TemplateId field value
func (o *FCMSConnectorResponse) GetTemplateId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *FCMSConnectorResponse) GetTemplateIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *FCMSConnectorResponse) SetTemplateId(v int64) {
	o.TemplateId = v
}

// GetDefaultOwner returns the DefaultOwner field value
func (o *FCMSConnectorResponse) GetDefaultOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultOwner
}

// GetDefaultOwnerOk returns a tuple with the DefaultOwner field value
// and a boolean to check if the value has been set.
func (o *FCMSConnectorResponse) GetDefaultOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultOwner, true
}

// SetDefaultOwner sets field value
func (o *FCMSConnectorResponse) SetDefaultOwner(v string) {
	o.DefaultOwner = v
}

// GetAuthenticationDomainId returns the AuthenticationDomainId field value
func (o *FCMSConnectorResponse) GetAuthenticationDomainId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AuthenticationDomainId
}

// GetAuthenticationDomainIdOk returns a tuple with the AuthenticationDomainId field value
// and a boolean to check if the value has been set.
func (o *FCMSConnectorResponse) GetAuthenticationDomainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationDomainId, true
}

// SetAuthenticationDomainId sets field value
func (o *FCMSConnectorResponse) SetAuthenticationDomainId(v int64) {
	o.AuthenticationDomainId = v
}

// GetOwnerGroups returns the OwnerGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FCMSConnectorResponse) GetOwnerGroups() string {
	if o == nil || utils.IsNil(o.OwnerGroups.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerGroups.Get()
}

// GetOwnerGroupsOk returns a tuple with the OwnerGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FCMSConnectorResponse) GetOwnerGroupsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerGroups.Get(), o.OwnerGroups.IsSet()
}

// HasOwnerGroups returns a boolean if a field has been set.
func (o *FCMSConnectorResponse) HasOwnerGroups() bool {
	if o != nil && o.OwnerGroups.IsSet() {
		return true
	}

	return false
}

// SetOwnerGroups gets a reference to the given NullableString and assigns it to the OwnerGroups field.
func (o *FCMSConnectorResponse) SetOwnerGroups(v string) {
	o.OwnerGroups.Set(&v)
}

// SetOwnerGroupsNil sets the value for OwnerGroups to be an explicit nil
func (o *FCMSConnectorResponse) SetOwnerGroupsNil() {
	o.OwnerGroups.Set(nil)
}

// UnsetOwnerGroups ensures that no value is present for OwnerGroups, not even an explicit nil
func (o *FCMSConnectorResponse) UnsetOwnerGroups() {
	o.OwnerGroups.Unset()
}

// GetDeleteOnRevoke returns the DeleteOnRevoke field value
func (o *FCMSConnectorResponse) GetDeleteOnRevoke() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DeleteOnRevoke
}

// GetDeleteOnRevokeOk returns a tuple with the DeleteOnRevoke field value
// and a boolean to check if the value has been set.
func (o *FCMSConnectorResponse) GetDeleteOnRevokeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeleteOnRevoke, true
}

// SetDeleteOnRevoke sets field value
func (o *FCMSConnectorResponse) SetDeleteOnRevoke(v bool) {
	o.DeleteOnRevoke = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FCMSConnectorResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FCMSConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *FCMSConnectorResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *FCMSConnectorResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *FCMSConnectorResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *FCMSConnectorResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FCMSConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FCMSConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *FCMSConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *FCMSConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *FCMSConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *FCMSConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FCMSConnectorResponse) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FCMSConnectorResponse) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *FCMSConnectorResponse) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *FCMSConnectorResponse) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *FCMSConnectorResponse) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *FCMSConnectorResponse) UnsetQueue() {
	o.Queue.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FCMSConnectorResponse) GetStatus() PKIConnectorStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret PKIConnectorStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FCMSConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *FCMSConnectorResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePKIConnectorStatus and assigns it to the Status field.
func (o *FCMSConnectorResponse) SetStatus(v PKIConnectorStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *FCMSConnectorResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *FCMSConnectorResponse) UnsetStatus() {
	o.Status.Unset()
}

func (o FCMSConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FCMSConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["endPoint"] = o.EndPoint
	toSerialize["apiCredentials"] = o.ApiCredentials
	toSerialize["templateId"] = o.TemplateId
	toSerialize["defaultOwner"] = o.DefaultOwner
	toSerialize["authenticationDomainId"] = o.AuthenticationDomainId
	if o.OwnerGroups.IsSet() {
		toSerialize["ownerGroups"] = o.OwnerGroups.Get()
	}
	toSerialize["deleteOnRevoke"] = o.DeleteOnRevoke
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FCMSConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"name",
		"type",
		"endPoint",
		"apiCredentials",
		"templateId",
		"defaultOwner",
		"authenticationDomainId",
		"deleteOnRevoke",
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

	varFCMSConnectorResponse := _FCMSConnectorResponse{}

	err = json.Unmarshal(data, &varFCMSConnectorResponse)

	if err != nil {
		return err
	}

	*o = FCMSConnectorResponse(varFCMSConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "endPoint")
		delete(additionalProperties, "apiCredentials")
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "defaultOwner")
		delete(additionalProperties, "authenticationDomainId")
		delete(additionalProperties, "ownerGroups")
		delete(additionalProperties, "deleteOnRevoke")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFCMSConnectorResponse struct {
	value *FCMSConnectorResponse
	isSet bool
}

func (v NullableFCMSConnectorResponse) Get() *FCMSConnectorResponse {
	return v.value
}

func (v *NullableFCMSConnectorResponse) Set(val *FCMSConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFCMSConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFCMSConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFCMSConnectorResponse(val *FCMSConnectorResponse) *NullableFCMSConnectorResponse {
	return &NullableFCMSConnectorResponse{value: val, isSet: true}
}

func (v NullableFCMSConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFCMSConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
