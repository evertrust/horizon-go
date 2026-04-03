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

// checks if the DiscoveryCampaignResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DiscoveryCampaignResponse{}

// DiscoveryCampaignResponse struct for DiscoveryCampaignResponse
type DiscoveryCampaignResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// The authorization levels of the discovery campaign
	AuthorizationLevels DiscoveryCampaignAuthorizationLevels `json:"authorizationLevels"`
	// The description of the discovery campaign
	Description utils.NullableString `json:"description,omitempty"`
	// Whether the discovery campaign is enabled, i.e. whether it can be fed
	Enabled bool `json:"enabled"`
	// Whether to log a Horizon event in case of failure
	EventOnFailure bool `json:"eventOnFailure"`
	// Whether to log a Horizon event in case of success
	EventOnSuccess bool `json:"eventOnSuccess"`
	// Whether to log a Horizon event in case of warning
	EventOnWarning bool `json:"eventOnWarning"`
	// The grading policies to apply to grade the discovered certificates on this campaign
	GradingPolicies []string `json:"gradingPolicies,omitempty"`
	// The hosts to be scanned by the discovery campaign
	Hosts []string `json:"hosts,omitempty"`
	// The name of the discovery campaign
	Name string `json:"name"`
	// The ports to be scanned by the discovery campaign
	Ports                []string `json:"ports,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveryCampaignResponse DiscoveryCampaignResponse

// NewDiscoveryCampaignResponse instantiates a new DiscoveryCampaignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryCampaignResponse(id string, authorizationLevels DiscoveryCampaignAuthorizationLevels, enabled bool, eventOnFailure bool, eventOnSuccess bool, eventOnWarning bool, name string) *DiscoveryCampaignResponse {
	this := DiscoveryCampaignResponse{}
	this.Id = id
	this.AuthorizationLevels = authorizationLevels
	this.Enabled = enabled
	this.EventOnFailure = eventOnFailure
	this.EventOnSuccess = eventOnSuccess
	this.EventOnWarning = eventOnWarning
	this.Name = name
	return &this
}

// NewDiscoveryCampaignResponseWithDefaults instantiates a new DiscoveryCampaignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryCampaignResponseWithDefaults() *DiscoveryCampaignResponse {
	this := DiscoveryCampaignResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DiscoveryCampaignResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiscoveryCampaignResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiscoveryCampaignResponse) SetId(v string) {
	o.Id = v
}

// GetAuthorizationLevels returns the AuthorizationLevels field value
func (o *DiscoveryCampaignResponse) GetAuthorizationLevels() DiscoveryCampaignAuthorizationLevels {
	if o == nil {
		var ret DiscoveryCampaignAuthorizationLevels
		return ret
	}

	return o.AuthorizationLevels
}

// GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field value
// and a boolean to check if the value has been set.
func (o *DiscoveryCampaignResponse) GetAuthorizationLevelsOk() (*DiscoveryCampaignAuthorizationLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationLevels, true
}

// SetAuthorizationLevels sets field value
func (o *DiscoveryCampaignResponse) SetAuthorizationLevels(v DiscoveryCampaignAuthorizationLevels) {
	o.AuthorizationLevels = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryCampaignResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryCampaignResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DiscoveryCampaignResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DiscoveryCampaignResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DiscoveryCampaignResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DiscoveryCampaignResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetEnabled returns the Enabled field value
func (o *DiscoveryCampaignResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DiscoveryCampaignResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DiscoveryCampaignResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEventOnFailure returns the EventOnFailure field value
func (o *DiscoveryCampaignResponse) GetEventOnFailure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EventOnFailure
}

// GetEventOnFailureOk returns a tuple with the EventOnFailure field value
// and a boolean to check if the value has been set.
func (o *DiscoveryCampaignResponse) GetEventOnFailureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventOnFailure, true
}

// SetEventOnFailure sets field value
func (o *DiscoveryCampaignResponse) SetEventOnFailure(v bool) {
	o.EventOnFailure = v
}

// GetEventOnSuccess returns the EventOnSuccess field value
func (o *DiscoveryCampaignResponse) GetEventOnSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EventOnSuccess
}

// GetEventOnSuccessOk returns a tuple with the EventOnSuccess field value
// and a boolean to check if the value has been set.
func (o *DiscoveryCampaignResponse) GetEventOnSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventOnSuccess, true
}

// SetEventOnSuccess sets field value
func (o *DiscoveryCampaignResponse) SetEventOnSuccess(v bool) {
	o.EventOnSuccess = v
}

// GetEventOnWarning returns the EventOnWarning field value
func (o *DiscoveryCampaignResponse) GetEventOnWarning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EventOnWarning
}

// GetEventOnWarningOk returns a tuple with the EventOnWarning field value
// and a boolean to check if the value has been set.
func (o *DiscoveryCampaignResponse) GetEventOnWarningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventOnWarning, true
}

// SetEventOnWarning sets field value
func (o *DiscoveryCampaignResponse) SetEventOnWarning(v bool) {
	o.EventOnWarning = v
}

// GetGradingPolicies returns the GradingPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryCampaignResponse) GetGradingPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GradingPolicies
}

// GetGradingPoliciesOk returns a tuple with the GradingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryCampaignResponse) GetGradingPoliciesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.GradingPolicies) {
		return nil, false
	}
	return o.GradingPolicies, true
}

// HasGradingPolicies returns a boolean if a field has been set.
func (o *DiscoveryCampaignResponse) HasGradingPolicies() bool {
	if o != nil && !utils.IsNil(o.GradingPolicies) {
		return true
	}

	return false
}

// SetGradingPolicies gets a reference to the given []string and assigns it to the GradingPolicies field.
func (o *DiscoveryCampaignResponse) SetGradingPolicies(v []string) {
	o.GradingPolicies = v
}

// GetHosts returns the Hosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryCampaignResponse) GetHosts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryCampaignResponse) GetHostsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *DiscoveryCampaignResponse) HasHosts() bool {
	if o != nil && !utils.IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *DiscoveryCampaignResponse) SetHosts(v []string) {
	o.Hosts = v
}

// GetName returns the Name field value
func (o *DiscoveryCampaignResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DiscoveryCampaignResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DiscoveryCampaignResponse) SetName(v string) {
	o.Name = v
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryCampaignResponse) GetPorts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryCampaignResponse) GetPortsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *DiscoveryCampaignResponse) HasPorts() bool {
	if o != nil && !utils.IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []string and assigns it to the Ports field.
func (o *DiscoveryCampaignResponse) SetPorts(v []string) {
	o.Ports = v
}

func (o DiscoveryCampaignResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryCampaignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["authorizationLevels"] = o.AuthorizationLevels
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["eventOnFailure"] = o.EventOnFailure
	toSerialize["eventOnSuccess"] = o.EventOnSuccess
	toSerialize["eventOnWarning"] = o.EventOnWarning
	if o.GradingPolicies != nil {
		toSerialize["gradingPolicies"] = o.GradingPolicies
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	toSerialize["name"] = o.Name
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoveryCampaignResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"authorizationLevels",
		"enabled",
		"eventOnFailure",
		"eventOnSuccess",
		"eventOnWarning",
		"name",
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

	varDiscoveryCampaignResponse := _DiscoveryCampaignResponse{}

	err = json.Unmarshal(data, &varDiscoveryCampaignResponse)

	if err != nil {
		return err
	}

	*o = DiscoveryCampaignResponse(varDiscoveryCampaignResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "authorizationLevels")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "eventOnFailure")
		delete(additionalProperties, "eventOnSuccess")
		delete(additionalProperties, "eventOnWarning")
		delete(additionalProperties, "gradingPolicies")
		delete(additionalProperties, "hosts")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveryCampaignResponse struct {
	value *DiscoveryCampaignResponse
	isSet bool
}

func (v NullableDiscoveryCampaignResponse) Get() *DiscoveryCampaignResponse {
	return v.value
}

func (v *NullableDiscoveryCampaignResponse) Set(val *DiscoveryCampaignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryCampaignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryCampaignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryCampaignResponse(val *DiscoveryCampaignResponse) *NullableDiscoveryCampaignResponse {
	return &NullableDiscoveryCampaignResponse{value: val, isSet: true}
}

func (v NullableDiscoveryCampaignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryCampaignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
