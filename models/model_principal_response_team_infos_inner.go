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

// checks if the PrincipalResponseTeamInfosInner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PrincipalResponseTeamInfosInner{}

// PrincipalResponseTeamInfosInner The team information
type PrincipalResponseTeamInfosInner struct {
	Description []LocalizedStringResponse `json:"description,omitempty"`
	DisplayName []LocalizedStringResponse `json:"displayName,omitempty"`
	// `true` if this team is externally managed (SCIM,...)
	ExternallyManaged *bool `json:"externallyManaged,omitempty"`
	// `true` if the principal is a manager of this team
	Manager              *bool   `json:"manager,omitempty"`
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalResponseTeamInfosInner PrincipalResponseTeamInfosInner

// NewPrincipalResponseTeamInfosInner instantiates a new PrincipalResponseTeamInfosInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalResponseTeamInfosInner() *PrincipalResponseTeamInfosInner {
	this := PrincipalResponseTeamInfosInner{}
	return &this
}

// NewPrincipalResponseTeamInfosInnerWithDefaults instantiates a new PrincipalResponseTeamInfosInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalResponseTeamInfosInnerWithDefaults() *PrincipalResponseTeamInfosInner {
	this := PrincipalResponseTeamInfosInner{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PrincipalResponseTeamInfosInner) GetDescription() []LocalizedStringResponse {
	if o == nil || utils.IsNil(o.Description) {
		var ret []LocalizedStringResponse
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalResponseTeamInfosInner) GetDescriptionOk() ([]LocalizedStringResponse, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PrincipalResponseTeamInfosInner) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedStringResponse and assigns it to the Description field.
func (o *PrincipalResponseTeamInfosInner) SetDescription(v []LocalizedStringResponse) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PrincipalResponseTeamInfosInner) GetDisplayName() []LocalizedStringResponse {
	if o == nil || utils.IsNil(o.DisplayName) {
		var ret []LocalizedStringResponse
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalResponseTeamInfosInner) GetDisplayNameOk() ([]LocalizedStringResponse, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PrincipalResponseTeamInfosInner) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedStringResponse and assigns it to the DisplayName field.
func (o *PrincipalResponseTeamInfosInner) SetDisplayName(v []LocalizedStringResponse) {
	o.DisplayName = v
}

// GetExternallyManaged returns the ExternallyManaged field value if set, zero value otherwise.
func (o *PrincipalResponseTeamInfosInner) GetExternallyManaged() bool {
	if o == nil || utils.IsNil(o.ExternallyManaged) {
		var ret bool
		return ret
	}
	return *o.ExternallyManaged
}

// GetExternallyManagedOk returns a tuple with the ExternallyManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalResponseTeamInfosInner) GetExternallyManagedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ExternallyManaged) {
		return nil, false
	}
	return o.ExternallyManaged, true
}

// HasExternallyManaged returns a boolean if a field has been set.
func (o *PrincipalResponseTeamInfosInner) HasExternallyManaged() bool {
	if o != nil && !utils.IsNil(o.ExternallyManaged) {
		return true
	}

	return false
}

// SetExternallyManaged gets a reference to the given bool and assigns it to the ExternallyManaged field.
func (o *PrincipalResponseTeamInfosInner) SetExternallyManaged(v bool) {
	o.ExternallyManaged = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *PrincipalResponseTeamInfosInner) GetManager() bool {
	if o == nil || utils.IsNil(o.Manager) {
		var ret bool
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalResponseTeamInfosInner) GetManagerOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *PrincipalResponseTeamInfosInner) HasManager() bool {
	if o != nil && !utils.IsNil(o.Manager) {
		return true
	}

	return false
}

// SetManager gets a reference to the given bool and assigns it to the Manager field.
func (o *PrincipalResponseTeamInfosInner) SetManager(v bool) {
	o.Manager = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PrincipalResponseTeamInfosInner) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalResponseTeamInfosInner) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PrincipalResponseTeamInfosInner) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PrincipalResponseTeamInfosInner) SetName(v string) {
	o.Name = &v
}

func (o PrincipalResponseTeamInfosInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalResponseTeamInfosInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !utils.IsNil(o.ExternallyManaged) {
		toSerialize["externallyManaged"] = o.ExternallyManaged
	}
	if !utils.IsNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalResponseTeamInfosInner) UnmarshalJSON(data []byte) (err error) {
	varPrincipalResponseTeamInfosInner := _PrincipalResponseTeamInfosInner{}

	err = json.Unmarshal(data, &varPrincipalResponseTeamInfosInner)

	if err != nil {
		return err
	}

	*o = PrincipalResponseTeamInfosInner(varPrincipalResponseTeamInfosInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "externallyManaged")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalResponseTeamInfosInner struct {
	value *PrincipalResponseTeamInfosInner
	isSet bool
}

func (v NullablePrincipalResponseTeamInfosInner) Get() *PrincipalResponseTeamInfosInner {
	return v.value
}

func (v *NullablePrincipalResponseTeamInfosInner) Set(val *PrincipalResponseTeamInfosInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalResponseTeamInfosInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalResponseTeamInfosInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalResponseTeamInfosInner(val *PrincipalResponseTeamInfosInner) *NullablePrincipalResponseTeamInfosInner {
	return &NullablePrincipalResponseTeamInfosInner{value: val, isSet: true}
}

func (v NullablePrincipalResponseTeamInfosInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalResponseTeamInfosInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
