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

// checks if the MetaPKIConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MetaPKIConnector{}

// MetaPKIConnector struct for MetaPKIConnector
type MetaPKIConnector struct {
	// Name of the `certificate` [credentials](#tag/security.credentials) to use to authenticate on the PKI
	AuthenticationCredentials utils.NullableString `json:"authenticationCredentials,omitempty"`
	// MetaPKI base endpoint
	EndPoint string `json:"endPoint"`
	// Certificate authority of the endpoint
	EndPointIssuingCA    string               `json:"endPointIssuingCA"`
	FormPorteurName      utils.NullableString `json:"formPorteurName,omitempty"`
	Name                 string               `json:"name"`
	ProfilCle            utils.NullableString `json:"profilCle"`
	Profile              string               `json:"profile"`
	Proxy                utils.NullableString `json:"proxy,omitempty"`
	Queue                utils.NullableString `json:"queue,omitempty"`
	Timeout              utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string               `json:"type"`
	ValidDays            utils.NullableString `json:"validDays,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Workflow             utils.NullableString `json:"workflow"`
	AdditionalProperties map[string]interface{}
}

type _MetaPKIConnector MetaPKIConnector

// NewMetaPKIConnector instantiates a new MetaPKIConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaPKIConnector(endPoint string, endPointIssuingCA string, name string, profilCle utils.NullableString, profile string, type_ string, workflow utils.NullableString) *MetaPKIConnector {
	this := MetaPKIConnector{}
	this.EndPoint = endPoint
	this.EndPointIssuingCA = endPointIssuingCA
	this.Name = name
	this.ProfilCle = profilCle
	this.Profile = profile
	this.Type = type_
	this.Workflow = workflow
	return &this
}

// NewMetaPKIConnectorWithDefaults instantiates a new MetaPKIConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaPKIConnectorWithDefaults() *MetaPKIConnector {
	this := MetaPKIConnector{}
	return &this
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnector) GetAuthenticationCredentials() string {
	if o == nil || utils.IsNil(o.AuthenticationCredentials.Get()) {
		var ret string
		return ret
	}
	return *o.AuthenticationCredentials.Get()
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnector) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationCredentials.Get(), o.AuthenticationCredentials.IsSet()
}

// HasAuthenticationCredentials returns a boolean if a field has been set.
func (o *MetaPKIConnector) HasAuthenticationCredentials() bool {
	if o != nil && o.AuthenticationCredentials.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationCredentials gets a reference to the given NullableString and assigns it to the AuthenticationCredentials field.
func (o *MetaPKIConnector) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials.Set(&v)
}

// SetAuthenticationCredentialsNil sets the value for AuthenticationCredentials to be an explicit nil
func (o *MetaPKIConnector) SetAuthenticationCredentialsNil() {
	o.AuthenticationCredentials.Set(nil)
}

// UnsetAuthenticationCredentials ensures that no value is present for AuthenticationCredentials, not even an explicit nil
func (o *MetaPKIConnector) UnsetAuthenticationCredentials() {
	o.AuthenticationCredentials.Unset()
}

// GetEndPoint returns the EndPoint field value
func (o *MetaPKIConnector) GetEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value
// and a boolean to check if the value has been set.
func (o *MetaPKIConnector) GetEndPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPoint, true
}

// SetEndPoint sets field value
func (o *MetaPKIConnector) SetEndPoint(v string) {
	o.EndPoint = v
}

// GetEndPointIssuingCA returns the EndPointIssuingCA field value
func (o *MetaPKIConnector) GetEndPointIssuingCA() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPointIssuingCA
}

// GetEndPointIssuingCAOk returns a tuple with the EndPointIssuingCA field value
// and a boolean to check if the value has been set.
func (o *MetaPKIConnector) GetEndPointIssuingCAOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPointIssuingCA, true
}

// SetEndPointIssuingCA sets field value
func (o *MetaPKIConnector) SetEndPointIssuingCA(v string) {
	o.EndPointIssuingCA = v
}

// GetFormPorteurName returns the FormPorteurName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnector) GetFormPorteurName() string {
	if o == nil || utils.IsNil(o.FormPorteurName.Get()) {
		var ret string
		return ret
	}
	return *o.FormPorteurName.Get()
}

// GetFormPorteurNameOk returns a tuple with the FormPorteurName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnector) GetFormPorteurNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormPorteurName.Get(), o.FormPorteurName.IsSet()
}

// HasFormPorteurName returns a boolean if a field has been set.
func (o *MetaPKIConnector) HasFormPorteurName() bool {
	if o != nil && o.FormPorteurName.IsSet() {
		return true
	}

	return false
}

// SetFormPorteurName gets a reference to the given NullableString and assigns it to the FormPorteurName field.
func (o *MetaPKIConnector) SetFormPorteurName(v string) {
	o.FormPorteurName.Set(&v)
}

// SetFormPorteurNameNil sets the value for FormPorteurName to be an explicit nil
func (o *MetaPKIConnector) SetFormPorteurNameNil() {
	o.FormPorteurName.Set(nil)
}

// UnsetFormPorteurName ensures that no value is present for FormPorteurName, not even an explicit nil
func (o *MetaPKIConnector) UnsetFormPorteurName() {
	o.FormPorteurName.Unset()
}

// GetName returns the Name field value
func (o *MetaPKIConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetaPKIConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetaPKIConnector) SetName(v string) {
	o.Name = v
}

// GetProfilCle returns the ProfilCle field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MetaPKIConnector) GetProfilCle() string {
	if o == nil || o.ProfilCle.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProfilCle.Get()
}

// GetProfilCleOk returns a tuple with the ProfilCle field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnector) GetProfilCleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilCle.Get(), o.ProfilCle.IsSet()
}

// SetProfilCle sets field value
func (o *MetaPKIConnector) SetProfilCle(v string) {
	o.ProfilCle.Set(&v)
}

// GetProfile returns the Profile field value
func (o *MetaPKIConnector) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *MetaPKIConnector) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *MetaPKIConnector) SetProfile(v string) {
	o.Profile = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *MetaPKIConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *MetaPKIConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *MetaPKIConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *MetaPKIConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnector) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnector) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *MetaPKIConnector) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *MetaPKIConnector) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *MetaPKIConnector) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *MetaPKIConnector) UnsetQueue() {
	o.Queue.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *MetaPKIConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *MetaPKIConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *MetaPKIConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *MetaPKIConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *MetaPKIConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetaPKIConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetaPKIConnector) SetType(v string) {
	o.Type = v
}

// GetValidDays returns the ValidDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnector) GetValidDays() string {
	if o == nil || utils.IsNil(o.ValidDays.Get()) {
		var ret string
		return ret
	}
	return *o.ValidDays.Get()
}

// GetValidDaysOk returns a tuple with the ValidDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnector) GetValidDaysOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidDays.Get(), o.ValidDays.IsSet()
}

// HasValidDays returns a boolean if a field has been set.
func (o *MetaPKIConnector) HasValidDays() bool {
	if o != nil && o.ValidDays.IsSet() {
		return true
	}

	return false
}

// SetValidDays gets a reference to the given NullableString and assigns it to the ValidDays field.
func (o *MetaPKIConnector) SetValidDays(v string) {
	o.ValidDays.Set(&v)
}

// SetValidDaysNil sets the value for ValidDays to be an explicit nil
func (o *MetaPKIConnector) SetValidDaysNil() {
	o.ValidDays.Set(nil)
}

// UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil
func (o *MetaPKIConnector) UnsetValidDays() {
	o.ValidDays.Unset()
}

// GetWorkflow returns the Workflow field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MetaPKIConnector) GetWorkflow() string {
	if o == nil || o.Workflow.Get() == nil {
		var ret string
		return ret
	}

	return *o.Workflow.Get()
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnector) GetWorkflowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workflow.Get(), o.Workflow.IsSet()
}

// SetWorkflow sets field value
func (o *MetaPKIConnector) SetWorkflow(v string) {
	o.Workflow.Set(&v)
}

func (o MetaPKIConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaPKIConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationCredentials.IsSet() {
		toSerialize["authenticationCredentials"] = o.AuthenticationCredentials.Get()
	}
	toSerialize["endPoint"] = o.EndPoint
	toSerialize["endPointIssuingCA"] = o.EndPointIssuingCA
	if o.FormPorteurName.IsSet() {
		toSerialize["formPorteurName"] = o.FormPorteurName.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["profilCle"] = o.ProfilCle.Get()
	toSerialize["profile"] = o.Profile
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	toSerialize["type"] = o.Type
	if o.ValidDays.IsSet() {
		toSerialize["validDays"] = o.ValidDays.Get()
	}
	toSerialize["workflow"] = o.Workflow.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetaPKIConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"endPoint",
		"endPointIssuingCA",
		"name",
		"profilCle",
		"profile",
		"type",
		"workflow",
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

	varMetaPKIConnector := _MetaPKIConnector{}

	err = json.Unmarshal(data, &varMetaPKIConnector)

	if err != nil {
		return err
	}

	*o = MetaPKIConnector(varMetaPKIConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticationCredentials")
		delete(additionalProperties, "endPoint")
		delete(additionalProperties, "endPointIssuingCA")
		delete(additionalProperties, "formPorteurName")
		delete(additionalProperties, "name")
		delete(additionalProperties, "profilCle")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		delete(additionalProperties, "validDays")
		delete(additionalProperties, "workflow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetaPKIConnector struct {
	value *MetaPKIConnector
	isSet bool
}

func (v NullableMetaPKIConnector) Get() *MetaPKIConnector {
	return v.value
}

func (v *NullableMetaPKIConnector) Set(val *MetaPKIConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaPKIConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaPKIConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaPKIConnector(val *MetaPKIConnector) *NullableMetaPKIConnector {
	return &NullableMetaPKIConnector{value: val, isSet: true}
}

func (v NullableMetaPKIConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaPKIConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
