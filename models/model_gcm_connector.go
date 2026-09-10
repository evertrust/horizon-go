/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the GCMConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GCMConnector{}

// GCMConnector struct for GCMConnector
type GCMConnector struct {
	// Name of the `raw` [credentials](#tag/security.credentials) containing User Account credentials.
	Credentials          string               `json:"credentials"`
	Location             string               `json:"location"`
	Name                 string               `json:"name"`
	Project              string               `json:"project"`
	Proxy                utils.NullableString `json:"proxy,omitempty"`
	TagKey               utils.NullableString `json:"tagKey,omitempty"`
	TagValue             utils.NullableString `json:"tagValue,omitempty"`
	ThrottleDuration     string               `json:"throttleDuration" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Timeout              utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string               `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _GCMConnector GCMConnector

// NewGCMConnector instantiates a new GCMConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCMConnector(credentials string, location string, name string, project string, throttleDuration string, type_ string) *GCMConnector {
	this := GCMConnector{}
	this.Credentials = credentials
	this.Location = location
	this.Name = name
	this.Project = project
	this.ThrottleDuration = throttleDuration
	this.Type = type_
	return &this
}

// NewGCMConnectorWithDefaults instantiates a new GCMConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCMConnectorWithDefaults() *GCMConnector {
	this := GCMConnector{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *GCMConnector) GetCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *GCMConnector) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *GCMConnector) SetCredentials(v string) {
	o.Credentials = v
}

// GetLocation returns the Location field value
func (o *GCMConnector) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *GCMConnector) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *GCMConnector) SetLocation(v string) {
	o.Location = v
}

// GetName returns the Name field value
func (o *GCMConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GCMConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GCMConnector) SetName(v string) {
	o.Name = v
}

// GetProject returns the Project field value
func (o *GCMConnector) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *GCMConnector) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *GCMConnector) SetProject(v string) {
	o.Project = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GCMConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GCMConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *GCMConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *GCMConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *GCMConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *GCMConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetTagKey returns the TagKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GCMConnector) GetTagKey() string {
	if o == nil || utils.IsNil(o.TagKey.Get()) {
		var ret string
		return ret
	}
	return *o.TagKey.Get()
}

// GetTagKeyOk returns a tuple with the TagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GCMConnector) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagKey.Get(), o.TagKey.IsSet()
}

// HasTagKey returns a boolean if a field has been set.
func (o *GCMConnector) HasTagKey() bool {
	if o != nil && o.TagKey.IsSet() {
		return true
	}

	return false
}

// SetTagKey gets a reference to the given NullableString and assigns it to the TagKey field.
func (o *GCMConnector) SetTagKey(v string) {
	o.TagKey.Set(&v)
}

// SetTagKeyNil sets the value for TagKey to be an explicit nil
func (o *GCMConnector) SetTagKeyNil() {
	o.TagKey.Set(nil)
}

// UnsetTagKey ensures that no value is present for TagKey, not even an explicit nil
func (o *GCMConnector) UnsetTagKey() {
	o.TagKey.Unset()
}

// GetTagValue returns the TagValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GCMConnector) GetTagValue() string {
	if o == nil || utils.IsNil(o.TagValue.Get()) {
		var ret string
		return ret
	}
	return *o.TagValue.Get()
}

// GetTagValueOk returns a tuple with the TagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GCMConnector) GetTagValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagValue.Get(), o.TagValue.IsSet()
}

// HasTagValue returns a boolean if a field has been set.
func (o *GCMConnector) HasTagValue() bool {
	if o != nil && o.TagValue.IsSet() {
		return true
	}

	return false
}

// SetTagValue gets a reference to the given NullableString and assigns it to the TagValue field.
func (o *GCMConnector) SetTagValue(v string) {
	o.TagValue.Set(&v)
}

// SetTagValueNil sets the value for TagValue to be an explicit nil
func (o *GCMConnector) SetTagValueNil() {
	o.TagValue.Set(nil)
}

// UnsetTagValue ensures that no value is present for TagValue, not even an explicit nil
func (o *GCMConnector) UnsetTagValue() {
	o.TagValue.Unset()
}

// GetThrottleDuration returns the ThrottleDuration field value
func (o *GCMConnector) GetThrottleDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThrottleDuration
}

// GetThrottleDurationOk returns a tuple with the ThrottleDuration field value
// and a boolean to check if the value has been set.
func (o *GCMConnector) GetThrottleDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleDuration, true
}

// SetThrottleDuration sets field value
func (o *GCMConnector) SetThrottleDuration(v string) {
	o.ThrottleDuration = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GCMConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GCMConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *GCMConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *GCMConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *GCMConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *GCMConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *GCMConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GCMConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GCMConnector) SetType(v string) {
	o.Type = v
}

func (o GCMConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GCMConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.Credentials
	toSerialize["location"] = o.Location
	toSerialize["name"] = o.Name
	toSerialize["project"] = o.Project
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.TagKey.IsSet() {
		toSerialize["tagKey"] = o.TagKey.Get()
	}
	if o.TagValue.IsSet() {
		toSerialize["tagValue"] = o.TagValue.Get()
	}
	toSerialize["throttleDuration"] = o.ThrottleDuration
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GCMConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"credentials",
		"location",
		"name",
		"project",
		"throttleDuration",
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

	varGCMConnector := _GCMConnector{}

	err = json.Unmarshal(data, &varGCMConnector)

	if err != nil {
		return err
	}

	*o = GCMConnector(varGCMConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "location")
		delete(additionalProperties, "name")
		delete(additionalProperties, "project")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "tagKey")
		delete(additionalProperties, "tagValue")
		delete(additionalProperties, "throttleDuration")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGCMConnector struct {
	value *GCMConnector
	isSet bool
}

func (v NullableGCMConnector) Get() *GCMConnector {
	return v.value
}

func (v *NullableGCMConnector) Set(val *GCMConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableGCMConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableGCMConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCMConnector(val *GCMConnector) *NullableGCMConnector {
	return &NullableGCMConnector{value: val, isSet: true}
}

func (v NullableGCMConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCMConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
