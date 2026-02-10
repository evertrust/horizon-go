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

// checks if the MetaPKIConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MetaPKIConnectorResponse{}

// MetaPKIConnectorResponse struct for MetaPKIConnectorResponse
type MetaPKIConnectorResponse struct {
	// Object internal ID
	Id   string `json:"_id"`
	Name string `json:"name"`
	Type string `json:"type"`
	// MetaPKI base endpoint
	EndPoint string `json:"endPoint"`
	// Certificate authority of the endpoint
	EndPointIssuingCA string               `json:"endPointIssuingCA"`
	Profile           string               `json:"profile"`
	Workflow          utils.NullableString `json:"workflow"`
	ProfilCle         utils.NullableString `json:"profilCle"`
	ValidDays         utils.NullableString `json:"validDays,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	FormPorteurName   utils.NullableString `json:"formPorteurName,omitempty"`
	// Name of the `certificate` [credentials](#tag/security.credentials) to use to authenticate on the PKI
	AuthenticationCredentials utils.NullableString       `json:"authenticationCredentials,omitempty"`
	Timeout                   utils.NullableString       `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Proxy                     utils.NullableString       `json:"proxy,omitempty"`
	Queue                     utils.NullableString       `json:"queue,omitempty"`
	Status                    NullablePKIConnectorStatus `json:"status,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _MetaPKIConnectorResponse MetaPKIConnectorResponse

// NewMetaPKIConnectorResponse instantiates a new MetaPKIConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaPKIConnectorResponse(id string, name string, type_ string, endPoint string, endPointIssuingCA string, profile string, workflow utils.NullableString, profilCle utils.NullableString) *MetaPKIConnectorResponse {
	this := MetaPKIConnectorResponse{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.EndPoint = endPoint
	this.EndPointIssuingCA = endPointIssuingCA
	this.Profile = profile
	this.Workflow = workflow
	this.ProfilCle = profilCle
	return &this
}

// NewMetaPKIConnectorResponseWithDefaults instantiates a new MetaPKIConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaPKIConnectorResponseWithDefaults() *MetaPKIConnectorResponse {
	this := MetaPKIConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *MetaPKIConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MetaPKIConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MetaPKIConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *MetaPKIConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetaPKIConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetaPKIConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *MetaPKIConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetaPKIConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetaPKIConnectorResponse) SetType(v string) {
	o.Type = v
}

// GetEndPoint returns the EndPoint field value
func (o *MetaPKIConnectorResponse) GetEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value
// and a boolean to check if the value has been set.
func (o *MetaPKIConnectorResponse) GetEndPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPoint, true
}

// SetEndPoint sets field value
func (o *MetaPKIConnectorResponse) SetEndPoint(v string) {
	o.EndPoint = v
}

// GetEndPointIssuingCA returns the EndPointIssuingCA field value
func (o *MetaPKIConnectorResponse) GetEndPointIssuingCA() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPointIssuingCA
}

// GetEndPointIssuingCAOk returns a tuple with the EndPointIssuingCA field value
// and a boolean to check if the value has been set.
func (o *MetaPKIConnectorResponse) GetEndPointIssuingCAOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPointIssuingCA, true
}

// SetEndPointIssuingCA sets field value
func (o *MetaPKIConnectorResponse) SetEndPointIssuingCA(v string) {
	o.EndPointIssuingCA = v
}

// GetProfile returns the Profile field value
func (o *MetaPKIConnectorResponse) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *MetaPKIConnectorResponse) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *MetaPKIConnectorResponse) SetProfile(v string) {
	o.Profile = v
}

// GetWorkflow returns the Workflow field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MetaPKIConnectorResponse) GetWorkflow() string {
	if o == nil || o.Workflow.Get() == nil {
		var ret string
		return ret
	}

	return *o.Workflow.Get()
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnectorResponse) GetWorkflowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workflow.Get(), o.Workflow.IsSet()
}

// SetWorkflow sets field value
func (o *MetaPKIConnectorResponse) SetWorkflow(v string) {
	o.Workflow.Set(&v)
}

// GetProfilCle returns the ProfilCle field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MetaPKIConnectorResponse) GetProfilCle() string {
	if o == nil || o.ProfilCle.Get() == nil {
		var ret string
		return ret
	}

	return *o.ProfilCle.Get()
}

// GetProfilCleOk returns a tuple with the ProfilCle field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnectorResponse) GetProfilCleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfilCle.Get(), o.ProfilCle.IsSet()
}

// SetProfilCle sets field value
func (o *MetaPKIConnectorResponse) SetProfilCle(v string) {
	o.ProfilCle.Set(&v)
}

// GetValidDays returns the ValidDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnectorResponse) GetValidDays() string {
	if o == nil || utils.IsNil(o.ValidDays.Get()) {
		var ret string
		return ret
	}
	return *o.ValidDays.Get()
}

// GetValidDaysOk returns a tuple with the ValidDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnectorResponse) GetValidDaysOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidDays.Get(), o.ValidDays.IsSet()
}

// HasValidDays returns a boolean if a field has been set.
func (o *MetaPKIConnectorResponse) HasValidDays() bool {
	if o != nil && o.ValidDays.IsSet() {
		return true
	}

	return false
}

// SetValidDays gets a reference to the given NullableString and assigns it to the ValidDays field.
func (o *MetaPKIConnectorResponse) SetValidDays(v string) {
	o.ValidDays.Set(&v)
}

// SetValidDaysNil sets the value for ValidDays to be an explicit nil
func (o *MetaPKIConnectorResponse) SetValidDaysNil() {
	o.ValidDays.Set(nil)
}

// UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil
func (o *MetaPKIConnectorResponse) UnsetValidDays() {
	o.ValidDays.Unset()
}

// GetFormPorteurName returns the FormPorteurName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnectorResponse) GetFormPorteurName() string {
	if o == nil || utils.IsNil(o.FormPorteurName.Get()) {
		var ret string
		return ret
	}
	return *o.FormPorteurName.Get()
}

// GetFormPorteurNameOk returns a tuple with the FormPorteurName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnectorResponse) GetFormPorteurNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FormPorteurName.Get(), o.FormPorteurName.IsSet()
}

// HasFormPorteurName returns a boolean if a field has been set.
func (o *MetaPKIConnectorResponse) HasFormPorteurName() bool {
	if o != nil && o.FormPorteurName.IsSet() {
		return true
	}

	return false
}

// SetFormPorteurName gets a reference to the given NullableString and assigns it to the FormPorteurName field.
func (o *MetaPKIConnectorResponse) SetFormPorteurName(v string) {
	o.FormPorteurName.Set(&v)
}

// SetFormPorteurNameNil sets the value for FormPorteurName to be an explicit nil
func (o *MetaPKIConnectorResponse) SetFormPorteurNameNil() {
	o.FormPorteurName.Set(nil)
}

// UnsetFormPorteurName ensures that no value is present for FormPorteurName, not even an explicit nil
func (o *MetaPKIConnectorResponse) UnsetFormPorteurName() {
	o.FormPorteurName.Unset()
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnectorResponse) GetAuthenticationCredentials() string {
	if o == nil || utils.IsNil(o.AuthenticationCredentials.Get()) {
		var ret string
		return ret
	}
	return *o.AuthenticationCredentials.Get()
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnectorResponse) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthenticationCredentials.Get(), o.AuthenticationCredentials.IsSet()
}

// HasAuthenticationCredentials returns a boolean if a field has been set.
func (o *MetaPKIConnectorResponse) HasAuthenticationCredentials() bool {
	if o != nil && o.AuthenticationCredentials.IsSet() {
		return true
	}

	return false
}

// SetAuthenticationCredentials gets a reference to the given NullableString and assigns it to the AuthenticationCredentials field.
func (o *MetaPKIConnectorResponse) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials.Set(&v)
}

// SetAuthenticationCredentialsNil sets the value for AuthenticationCredentials to be an explicit nil
func (o *MetaPKIConnectorResponse) SetAuthenticationCredentialsNil() {
	o.AuthenticationCredentials.Set(nil)
}

// UnsetAuthenticationCredentials ensures that no value is present for AuthenticationCredentials, not even an explicit nil
func (o *MetaPKIConnectorResponse) UnsetAuthenticationCredentials() {
	o.AuthenticationCredentials.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnectorResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *MetaPKIConnectorResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *MetaPKIConnectorResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *MetaPKIConnectorResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *MetaPKIConnectorResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *MetaPKIConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *MetaPKIConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *MetaPKIConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *MetaPKIConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnectorResponse) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnectorResponse) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *MetaPKIConnectorResponse) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *MetaPKIConnectorResponse) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *MetaPKIConnectorResponse) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *MetaPKIConnectorResponse) UnsetQueue() {
	o.Queue.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetaPKIConnectorResponse) GetStatus() PKIConnectorStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret PKIConnectorStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetaPKIConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *MetaPKIConnectorResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePKIConnectorStatus and assigns it to the Status field.
func (o *MetaPKIConnectorResponse) SetStatus(v PKIConnectorStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *MetaPKIConnectorResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *MetaPKIConnectorResponse) UnsetStatus() {
	o.Status.Unset()
}

func (o MetaPKIConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetaPKIConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["endPoint"] = o.EndPoint
	toSerialize["endPointIssuingCA"] = o.EndPointIssuingCA
	toSerialize["profile"] = o.Profile
	toSerialize["workflow"] = o.Workflow.Get()
	toSerialize["profilCle"] = o.ProfilCle.Get()
	if o.ValidDays.IsSet() {
		toSerialize["validDays"] = o.ValidDays.Get()
	}
	if o.FormPorteurName.IsSet() {
		toSerialize["formPorteurName"] = o.FormPorteurName.Get()
	}
	if o.AuthenticationCredentials.IsSet() {
		toSerialize["authenticationCredentials"] = o.AuthenticationCredentials.Get()
	}
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

func (o *MetaPKIConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"name",
		"type",
		"endPoint",
		"endPointIssuingCA",
		"profile",
		"workflow",
		"profilCle",
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

	varMetaPKIConnectorResponse := _MetaPKIConnectorResponse{}

	err = json.Unmarshal(data, &varMetaPKIConnectorResponse)

	if err != nil {
		return err
	}

	*o = MetaPKIConnectorResponse(varMetaPKIConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "endPoint")
		delete(additionalProperties, "endPointIssuingCA")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "workflow")
		delete(additionalProperties, "profilCle")
		delete(additionalProperties, "validDays")
		delete(additionalProperties, "formPorteurName")
		delete(additionalProperties, "authenticationCredentials")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetaPKIConnectorResponse struct {
	value *MetaPKIConnectorResponse
	isSet bool
}

func (v NullableMetaPKIConnectorResponse) Get() *MetaPKIConnectorResponse {
	return v.value
}

func (v *NullableMetaPKIConnectorResponse) Set(val *MetaPKIConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaPKIConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaPKIConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaPKIConnectorResponse(val *MetaPKIConnectorResponse) *NullableMetaPKIConnectorResponse {
	return &NullableMetaPKIConnectorResponse{value: val, isSet: true}
}

func (v NullableMetaPKIConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaPKIConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
