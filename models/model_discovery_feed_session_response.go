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

// checks if the DiscoveryFeedSessionResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DiscoveryFeedSessionResponse{}

// DiscoveryFeedSessionResponse struct for DiscoveryFeedSessionResponse
type DiscoveryFeedSessionResponse struct {
	// The name of the discovery campaign the feed session belongs to
	Campaign string `json:"campaign"`
	// The description of the discovery feed session
	Description utils.NullableString `json:"description,omitempty"`
	// Whether to generate an event on failure (defaults to the campaign setting)
	EventOnFailure utils.NullableBool `json:"eventOnFailure,omitempty"`
	// Whether to generate an event on success (defaults to the campaign setting)
	EventOnSuccess utils.NullableBool `json:"eventOnSuccess,omitempty"`
	// Whether to generate an event on warning (defaults to the campaign setting)
	EventOnWarning utils.NullableBool `json:"eventOnWarning,omitempty"`
	// The hosts on which the discovery campaign takes place
	Hosts []string `json:"hosts,omitempty"`
	// Object internal ID
	Id string `json:"id"`
	// The ports on which the discovery campaign takes place
	Ports                []int64 `json:"ports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveryFeedSessionResponse DiscoveryFeedSessionResponse

// NewDiscoveryFeedSessionResponse instantiates a new DiscoveryFeedSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryFeedSessionResponse(campaign string, id string) *DiscoveryFeedSessionResponse {
	this := DiscoveryFeedSessionResponse{}
	this.Campaign = campaign
	this.Id = id
	return &this
}

// NewDiscoveryFeedSessionResponseWithDefaults instantiates a new DiscoveryFeedSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryFeedSessionResponseWithDefaults() *DiscoveryFeedSessionResponse {
	this := DiscoveryFeedSessionResponse{}
	return &this
}

// GetCampaign returns the Campaign field value
func (o *DiscoveryFeedSessionResponse) GetCampaign() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value
// and a boolean to check if the value has been set.
func (o *DiscoveryFeedSessionResponse) GetCampaignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaign, true
}

// SetCampaign sets field value
func (o *DiscoveryFeedSessionResponse) SetCampaign(v string) {
	o.Campaign = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryFeedSessionResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryFeedSessionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DiscoveryFeedSessionResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DiscoveryFeedSessionResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DiscoveryFeedSessionResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DiscoveryFeedSessionResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetEventOnFailure returns the EventOnFailure field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryFeedSessionResponse) GetEventOnFailure() bool {
	if o == nil || utils.IsNil(o.EventOnFailure.Get()) {
		var ret bool
		return ret
	}
	return *o.EventOnFailure.Get()
}

// GetEventOnFailureOk returns a tuple with the EventOnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryFeedSessionResponse) GetEventOnFailureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventOnFailure.Get(), o.EventOnFailure.IsSet()
}

// HasEventOnFailure returns a boolean if a field has been set.
func (o *DiscoveryFeedSessionResponse) HasEventOnFailure() bool {
	if o != nil && o.EventOnFailure.IsSet() {
		return true
	}

	return false
}

// SetEventOnFailure gets a reference to the given NullableBool and assigns it to the EventOnFailure field.
func (o *DiscoveryFeedSessionResponse) SetEventOnFailure(v bool) {
	o.EventOnFailure.Set(&v)
}

// SetEventOnFailureNil sets the value for EventOnFailure to be an explicit nil
func (o *DiscoveryFeedSessionResponse) SetEventOnFailureNil() {
	o.EventOnFailure.Set(nil)
}

// UnsetEventOnFailure ensures that no value is present for EventOnFailure, not even an explicit nil
func (o *DiscoveryFeedSessionResponse) UnsetEventOnFailure() {
	o.EventOnFailure.Unset()
}

// GetEventOnSuccess returns the EventOnSuccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryFeedSessionResponse) GetEventOnSuccess() bool {
	if o == nil || utils.IsNil(o.EventOnSuccess.Get()) {
		var ret bool
		return ret
	}
	return *o.EventOnSuccess.Get()
}

// GetEventOnSuccessOk returns a tuple with the EventOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryFeedSessionResponse) GetEventOnSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventOnSuccess.Get(), o.EventOnSuccess.IsSet()
}

// HasEventOnSuccess returns a boolean if a field has been set.
func (o *DiscoveryFeedSessionResponse) HasEventOnSuccess() bool {
	if o != nil && o.EventOnSuccess.IsSet() {
		return true
	}

	return false
}

// SetEventOnSuccess gets a reference to the given NullableBool and assigns it to the EventOnSuccess field.
func (o *DiscoveryFeedSessionResponse) SetEventOnSuccess(v bool) {
	o.EventOnSuccess.Set(&v)
}

// SetEventOnSuccessNil sets the value for EventOnSuccess to be an explicit nil
func (o *DiscoveryFeedSessionResponse) SetEventOnSuccessNil() {
	o.EventOnSuccess.Set(nil)
}

// UnsetEventOnSuccess ensures that no value is present for EventOnSuccess, not even an explicit nil
func (o *DiscoveryFeedSessionResponse) UnsetEventOnSuccess() {
	o.EventOnSuccess.Unset()
}

// GetEventOnWarning returns the EventOnWarning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryFeedSessionResponse) GetEventOnWarning() bool {
	if o == nil || utils.IsNil(o.EventOnWarning.Get()) {
		var ret bool
		return ret
	}
	return *o.EventOnWarning.Get()
}

// GetEventOnWarningOk returns a tuple with the EventOnWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryFeedSessionResponse) GetEventOnWarningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventOnWarning.Get(), o.EventOnWarning.IsSet()
}

// HasEventOnWarning returns a boolean if a field has been set.
func (o *DiscoveryFeedSessionResponse) HasEventOnWarning() bool {
	if o != nil && o.EventOnWarning.IsSet() {
		return true
	}

	return false
}

// SetEventOnWarning gets a reference to the given NullableBool and assigns it to the EventOnWarning field.
func (o *DiscoveryFeedSessionResponse) SetEventOnWarning(v bool) {
	o.EventOnWarning.Set(&v)
}

// SetEventOnWarningNil sets the value for EventOnWarning to be an explicit nil
func (o *DiscoveryFeedSessionResponse) SetEventOnWarningNil() {
	o.EventOnWarning.Set(nil)
}

// UnsetEventOnWarning ensures that no value is present for EventOnWarning, not even an explicit nil
func (o *DiscoveryFeedSessionResponse) UnsetEventOnWarning() {
	o.EventOnWarning.Unset()
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryFeedSessionResponse) GetHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryFeedSessionResponse) GetHostsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *DiscoveryFeedSessionResponse) HasHosts() bool {
	if o != nil && !utils.IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *DiscoveryFeedSessionResponse) SetHosts(v []string) {
	o.Hosts = v
}

// GetId returns the Id field value
func (o *DiscoveryFeedSessionResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiscoveryFeedSessionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiscoveryFeedSessionResponse) SetId(v string) {
	o.Id = v
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryFeedSessionResponse) GetPorts() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryFeedSessionResponse) GetPortsOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *DiscoveryFeedSessionResponse) HasPorts() bool {
	if o != nil && !utils.IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []int64 and assigns it to the Ports field.
func (o *DiscoveryFeedSessionResponse) SetPorts(v []int64) {
	o.Ports = v
}

func (o DiscoveryFeedSessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryFeedSessionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaign"] = o.Campaign
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.EventOnFailure.IsSet() {
		toSerialize["eventOnFailure"] = o.EventOnFailure.Get()
	}
	if o.EventOnSuccess.IsSet() {
		toSerialize["eventOnSuccess"] = o.EventOnSuccess.Get()
	}
	if o.EventOnWarning.IsSet() {
		toSerialize["eventOnWarning"] = o.EventOnWarning.Get()
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	toSerialize["id"] = o.Id
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoveryFeedSessionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"campaign",
		"id",
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

	varDiscoveryFeedSessionResponse := _DiscoveryFeedSessionResponse{}

	err = json.Unmarshal(data, &varDiscoveryFeedSessionResponse)

	if err != nil {
		return err
	}

	*o = DiscoveryFeedSessionResponse(varDiscoveryFeedSessionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "campaign")
		delete(additionalProperties, "description")
		delete(additionalProperties, "eventOnFailure")
		delete(additionalProperties, "eventOnSuccess")
		delete(additionalProperties, "eventOnWarning")
		delete(additionalProperties, "hosts")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveryFeedSessionResponse struct {
	value *DiscoveryFeedSessionResponse
	isSet bool
}

func (v NullableDiscoveryFeedSessionResponse) Get() *DiscoveryFeedSessionResponse {
	return v.value
}

func (v *NullableDiscoveryFeedSessionResponse) Set(val *DiscoveryFeedSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryFeedSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryFeedSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryFeedSessionResponse(val *DiscoveryFeedSessionResponse) *NullableDiscoveryFeedSessionResponse {
	return &NullableDiscoveryFeedSessionResponse{value: val, isSet: true}
}

func (v NullableDiscoveryFeedSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryFeedSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
