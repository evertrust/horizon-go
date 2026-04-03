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

// checks if the FCMSConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FCMSConnector{}

// FCMSConnector struct for FCMSConnector
type FCMSConnector struct {
	// Name of the `raw` [credentials](#tag/security.credentials) containing the API key to authenticate on the PKI
	ApiCredentials         string               `json:"apiCredentials"`
	AuthenticationDomainId int64                `json:"authenticationDomainId"`
	DefaultOwner           string               `json:"defaultOwner"`
	DeleteOnRevoke         bool                 `json:"deleteOnRevoke"`
	EndPoint               string               `json:"endPoint"`
	Name                   string               `json:"name"`
	OwnerGroups            utils.NullableString `json:"ownerGroups,omitempty"`
	Proxy                  utils.NullableString `json:"proxy,omitempty"`
	Queue                  utils.NullableString `json:"queue,omitempty"`
	TemplateId             int64                `json:"templateId"`
	Timeout                utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                   string               `json:"type"`
	AdditionalProperties   map[string]interface{}
}

type _FCMSConnector FCMSConnector

// NewFCMSConnector instantiates a new FCMSConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFCMSConnector(apiCredentials string, authenticationDomainId int64, defaultOwner string, deleteOnRevoke bool, endPoint string, name string, templateId int64, type_ string) *FCMSConnector {
	this := FCMSConnector{}
	this.ApiCredentials = apiCredentials
	this.AuthenticationDomainId = authenticationDomainId
	this.DefaultOwner = defaultOwner
	this.DeleteOnRevoke = deleteOnRevoke
	this.EndPoint = endPoint
	this.Name = name
	this.TemplateId = templateId
	this.Type = type_
	return &this
}

// NewFCMSConnectorWithDefaults instantiates a new FCMSConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFCMSConnectorWithDefaults() *FCMSConnector {
	this := FCMSConnector{}
	return &this
}

// GetApiCredentials returns the ApiCredentials field value
func (o *FCMSConnector) GetApiCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiCredentials
}

// GetApiCredentialsOk returns a tuple with the ApiCredentials field value
// and a boolean to check if the value has been set.
func (o *FCMSConnector) GetApiCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiCredentials, true
}

// SetApiCredentials sets field value
func (o *FCMSConnector) SetApiCredentials(v string) {
	o.ApiCredentials = v
}

// GetAuthenticationDomainId returns the AuthenticationDomainId field value
func (o *FCMSConnector) GetAuthenticationDomainId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AuthenticationDomainId
}

// GetAuthenticationDomainIdOk returns a tuple with the AuthenticationDomainId field value
// and a boolean to check if the value has been set.
func (o *FCMSConnector) GetAuthenticationDomainIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationDomainId, true
}

// SetAuthenticationDomainId sets field value
func (o *FCMSConnector) SetAuthenticationDomainId(v int64) {
	o.AuthenticationDomainId = v
}

// GetDefaultOwner returns the DefaultOwner field value
func (o *FCMSConnector) GetDefaultOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultOwner
}

// GetDefaultOwnerOk returns a tuple with the DefaultOwner field value
// and a boolean to check if the value has been set.
func (o *FCMSConnector) GetDefaultOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultOwner, true
}

// SetDefaultOwner sets field value
func (o *FCMSConnector) SetDefaultOwner(v string) {
	o.DefaultOwner = v
}

// GetDeleteOnRevoke returns the DeleteOnRevoke field value
func (o *FCMSConnector) GetDeleteOnRevoke() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DeleteOnRevoke
}

// GetDeleteOnRevokeOk returns a tuple with the DeleteOnRevoke field value
// and a boolean to check if the value has been set.
func (o *FCMSConnector) GetDeleteOnRevokeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeleteOnRevoke, true
}

// SetDeleteOnRevoke sets field value
func (o *FCMSConnector) SetDeleteOnRevoke(v bool) {
	o.DeleteOnRevoke = v
}

// GetEndPoint returns the EndPoint field value
func (o *FCMSConnector) GetEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value
// and a boolean to check if the value has been set.
func (o *FCMSConnector) GetEndPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPoint, true
}

// SetEndPoint sets field value
func (o *FCMSConnector) SetEndPoint(v string) {
	o.EndPoint = v
}

// GetName returns the Name field value
func (o *FCMSConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FCMSConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FCMSConnector) SetName(v string) {
	o.Name = v
}

// GetOwnerGroups returns the OwnerGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FCMSConnector) GetOwnerGroups() string {
	if o == nil || utils.IsNil(o.OwnerGroups.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerGroups.Get()
}

// GetOwnerGroupsOk returns a tuple with the OwnerGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FCMSConnector) GetOwnerGroupsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerGroups.Get(), o.OwnerGroups.IsSet()
}

// HasOwnerGroups returns a boolean if a field has been set.
func (o *FCMSConnector) HasOwnerGroups() bool {
	if o != nil && o.OwnerGroups.IsSet() {
		return true
	}

	return false
}

// SetOwnerGroups gets a reference to the given NullableString and assigns it to the OwnerGroups field.
func (o *FCMSConnector) SetOwnerGroups(v string) {
	o.OwnerGroups.Set(&v)
}

// SetOwnerGroupsNil sets the value for OwnerGroups to be an explicit nil
func (o *FCMSConnector) SetOwnerGroupsNil() {
	o.OwnerGroups.Set(nil)
}

// UnsetOwnerGroups ensures that no value is present for OwnerGroups, not even an explicit nil
func (o *FCMSConnector) UnsetOwnerGroups() {
	o.OwnerGroups.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FCMSConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FCMSConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *FCMSConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *FCMSConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *FCMSConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *FCMSConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FCMSConnector) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FCMSConnector) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *FCMSConnector) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *FCMSConnector) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *FCMSConnector) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *FCMSConnector) UnsetQueue() {
	o.Queue.Unset()
}

// GetTemplateId returns the TemplateId field value
func (o *FCMSConnector) GetTemplateId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *FCMSConnector) GetTemplateIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *FCMSConnector) SetTemplateId(v int64) {
	o.TemplateId = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FCMSConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FCMSConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *FCMSConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *FCMSConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *FCMSConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *FCMSConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *FCMSConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FCMSConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FCMSConnector) SetType(v string) {
	o.Type = v
}

func (o FCMSConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FCMSConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiCredentials"] = o.ApiCredentials
	toSerialize["authenticationDomainId"] = o.AuthenticationDomainId
	toSerialize["defaultOwner"] = o.DefaultOwner
	toSerialize["deleteOnRevoke"] = o.DeleteOnRevoke
	toSerialize["endPoint"] = o.EndPoint
	toSerialize["name"] = o.Name
	if o.OwnerGroups.IsSet() {
		toSerialize["ownerGroups"] = o.OwnerGroups.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	toSerialize["templateId"] = o.TemplateId
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FCMSConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiCredentials",
		"authenticationDomainId",
		"defaultOwner",
		"deleteOnRevoke",
		"endPoint",
		"name",
		"templateId",
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

	varFCMSConnector := _FCMSConnector{}

	err = json.Unmarshal(data, &varFCMSConnector)

	if err != nil {
		return err
	}

	*o = FCMSConnector(varFCMSConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "apiCredentials")
		delete(additionalProperties, "authenticationDomainId")
		delete(additionalProperties, "defaultOwner")
		delete(additionalProperties, "deleteOnRevoke")
		delete(additionalProperties, "endPoint")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ownerGroups")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "templateId")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFCMSConnector struct {
	value *FCMSConnector
	isSet bool
}

func (v NullableFCMSConnector) Get() *FCMSConnector {
	return v.value
}

func (v *NullableFCMSConnector) Set(val *FCMSConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableFCMSConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableFCMSConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFCMSConnector(val *FCMSConnector) *NullableFCMSConnector {
	return &NullableFCMSConnector{value: val, isSet: true}
}

func (v NullableFCMSConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFCMSConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
