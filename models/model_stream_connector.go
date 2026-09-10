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

// checks if the StreamConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &StreamConnector{}

// StreamConnector struct for StreamConnector
type StreamConnector struct {
	// Name of the `certificate` [credentials](#tag/security.credentials) to use to authenticate on the PKI
	AuthenticationCredentials utils.NullableString `json:"authenticationCredentials,omitempty"`
	// Stream's technical name of the CA on which to enroll
	Ca string `json:"ca"`
	// Stream's base endpoint
	EndPoint string `json:"endPoint"`
	// Name of the `password` [credentials](#tag/security.credentials) to use to authenticate on the PKI
	LoginCredentials utils.NullableString `json:"loginCredentials,omitempty"`
	Name             string               `json:"name"`
	Proxy            utils.NullableString `json:"proxy,omitempty"`
	Queue            utils.NullableString `json:"queue,omitempty"`
	// Stream's certificate template to use for enrollment
	Template             string               `json:"template"`
	Timeout              utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string               `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _StreamConnector StreamConnector

// NewStreamConnector instantiates a new StreamConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamConnector(ca string, endPoint string, name string, template string, type_ string) *StreamConnector {
	this := StreamConnector{}
	this.Ca = ca
	this.EndPoint = endPoint
	this.Name = name
	this.Template = template
	this.Type = type_
	return &this
}

// NewStreamConnectorWithDefaults instantiates a new StreamConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamConnectorWithDefaults() *StreamConnector {
	this := StreamConnector{}
	return &this
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StreamConnector) GetAuthenticationCredentials() string {
	if o == nil || utils.IsNil(o.AuthenticationCredentials.Get()) {
		var ret string
		return ret
	}
	return *o.AuthenticationCredentials.Get()
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StreamConnector) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationCredentials.Get(), o.AuthenticationCredentials.IsSet()
}

// HasAuthenticationCredentials returns a boolean if a field has been set.
func (o *StreamConnector) HasAuthenticationCredentials() bool {
	if o != nil && o.AuthenticationCredentials.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationCredentials gets a reference to the given NullableString and assigns it to the AuthenticationCredentials field.
func (o *StreamConnector) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials.Set(&v)
}

// SetAuthenticationCredentialsNil sets the value for AuthenticationCredentials to be an explicit nil
func (o *StreamConnector) SetAuthenticationCredentialsNil() {
	o.AuthenticationCredentials.Set(nil)
}

// UnsetAuthenticationCredentials ensures that no value is present for AuthenticationCredentials, not even an explicit nil
func (o *StreamConnector) UnsetAuthenticationCredentials() {
	o.AuthenticationCredentials.Unset()
}

// GetCa returns the Ca field value
func (o *StreamConnector) GetCa() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ca
}

// GetCaOk returns a tuple with the Ca field value
// and a boolean to check if the value has been set.
func (o *StreamConnector) GetCaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ca, true
}

// SetCa sets field value
func (o *StreamConnector) SetCa(v string) {
	o.Ca = v
}

// GetEndPoint returns the EndPoint field value
func (o *StreamConnector) GetEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value
// and a boolean to check if the value has been set.
func (o *StreamConnector) GetEndPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPoint, true
}

// SetEndPoint sets field value
func (o *StreamConnector) SetEndPoint(v string) {
	o.EndPoint = v
}

// GetLoginCredentials returns the LoginCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StreamConnector) GetLoginCredentials() string {
	if o == nil || utils.IsNil(o.LoginCredentials.Get()) {
		var ret string
		return ret
	}
	return *o.LoginCredentials.Get()
}

// GetLoginCredentialsOk returns a tuple with the LoginCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StreamConnector) GetLoginCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoginCredentials.Get(), o.LoginCredentials.IsSet()
}

// HasLoginCredentials returns a boolean if a field has been set.
func (o *StreamConnector) HasLoginCredentials() bool {
	if o != nil && o.LoginCredentials.IsSet() {
		return true
	}

	return false
}

// SetLoginCredentials gets a reference to the given NullableString and assigns it to the LoginCredentials field.
func (o *StreamConnector) SetLoginCredentials(v string) {
	o.LoginCredentials.Set(&v)
}

// SetLoginCredentialsNil sets the value for LoginCredentials to be an explicit nil
func (o *StreamConnector) SetLoginCredentialsNil() {
	o.LoginCredentials.Set(nil)
}

// UnsetLoginCredentials ensures that no value is present for LoginCredentials, not even an explicit nil
func (o *StreamConnector) UnsetLoginCredentials() {
	o.LoginCredentials.Unset()
}

// GetName returns the Name field value
func (o *StreamConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StreamConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StreamConnector) SetName(v string) {
	o.Name = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StreamConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StreamConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *StreamConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *StreamConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *StreamConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *StreamConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StreamConnector) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StreamConnector) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *StreamConnector) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *StreamConnector) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *StreamConnector) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *StreamConnector) UnsetQueue() {
	o.Queue.Unset()
}

// GetTemplate returns the Template field value
func (o *StreamConnector) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *StreamConnector) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *StreamConnector) SetTemplate(v string) {
	o.Template = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StreamConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StreamConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *StreamConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *StreamConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *StreamConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *StreamConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *StreamConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StreamConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StreamConnector) SetType(v string) {
	o.Type = v
}

func (o StreamConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationCredentials.IsSet() {
		toSerialize["authenticationCredentials"] = o.AuthenticationCredentials.Get()
	}
	toSerialize["ca"] = o.Ca
	toSerialize["endPoint"] = o.EndPoint
	if o.LoginCredentials.IsSet() {
		toSerialize["loginCredentials"] = o.LoginCredentials.Get()
	}
	toSerialize["name"] = o.Name
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	toSerialize["template"] = o.Template
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StreamConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ca",
		"endPoint",
		"name",
		"template",
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

	varStreamConnector := _StreamConnector{}

	err = json.Unmarshal(data, &varStreamConnector)

	if err != nil {
		return err
	}

	*o = StreamConnector(varStreamConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticationCredentials")
		delete(additionalProperties, "ca")
		delete(additionalProperties, "endPoint")
		delete(additionalProperties, "loginCredentials")
		delete(additionalProperties, "name")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "template")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStreamConnector struct {
	value *StreamConnector
	isSet bool
}

func (v NullableStreamConnector) Get() *StreamConnector {
	return v.value
}

func (v *NullableStreamConnector) Set(val *StreamConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamConnector(val *StreamConnector) *NullableStreamConnector {
	return &NullableStreamConnector{value: val, isSet: true}
}

func (v NullableStreamConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
