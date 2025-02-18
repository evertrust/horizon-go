/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"encoding/json"
	"fmt"
)

// checks if the PrincipalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalResponse{}

// PrincipalResponse struct for PrincipalResponse
type PrincipalResponse struct {
	Identity Identity `json:"identity"`
	// The permissions of the principal
	Permissions []Permission `json:"permissions,omitempty"`
	// The roles of the principal
	Roles []Role `json:"roles,omitempty"`
	// The teams of the principal
	Teams []string `json:"teams,omitempty"`
	// The UI preferences of the principal
	Preferences NullablePrincipalInfoPreferences `json:"preferences,omitempty"`
	// The custom dashboards of the principal
	CustomDashboards []Dashboard `json:"customDashboards,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalResponse PrincipalResponse

// NewPrincipalResponse instantiates a new PrincipalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalResponse(identity Identity) *PrincipalResponse {
	this := PrincipalResponse{}
	this.Identity = identity
	return &this
}

// NewPrincipalResponseWithDefaults instantiates a new PrincipalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalResponseWithDefaults() *PrincipalResponse {
	this := PrincipalResponse{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *PrincipalResponse) GetIdentity() Identity {
	if o == nil {
		var ret Identity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *PrincipalResponse) GetIdentityOk() (*Identity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *PrincipalResponse) SetIdentity(v Identity) {
	o.Identity = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalResponse) GetPermissions() []Permission {
	if o == nil {
		var ret []Permission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalResponse) GetPermissionsOk() ([]Permission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PrincipalResponse) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []Permission and assigns it to the Permissions field.
func (o *PrincipalResponse) SetPermissions(v []Permission) {
	o.Permissions = v
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalResponse) GetRoles() []Role {
	if o == nil {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalResponse) GetRolesOk() ([]Role, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *PrincipalResponse) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *PrincipalResponse) SetRoles(v []Role) {
	o.Roles = v
}

// GetTeams returns the Teams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalResponse) GetTeams() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalResponse) GetTeamsOk() ([]string, bool) {
	if o == nil || IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *PrincipalResponse) HasTeams() bool {
	if o != nil && !IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *PrincipalResponse) SetTeams(v []string) {
	o.Teams = v
}

// GetPreferences returns the Preferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalResponse) GetPreferences() PrincipalInfoPreferences {
	if o == nil || IsNil(o.Preferences.Get()) {
		var ret PrincipalInfoPreferences
		return ret
	}
	return *o.Preferences.Get()
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalResponse) GetPreferencesOk() (*PrincipalInfoPreferences, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preferences.Get(), o.Preferences.IsSet()
}

// HasPreferences returns a boolean if a field has been set.
func (o *PrincipalResponse) HasPreferences() bool {
	if o != nil && o.Preferences.IsSet() {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given NullablePrincipalInfoPreferences and assigns it to the Preferences field.
func (o *PrincipalResponse) SetPreferences(v PrincipalInfoPreferences) {
	o.Preferences.Set(&v)
}
// SetPreferencesNil sets the value for Preferences to be an explicit nil
func (o *PrincipalResponse) SetPreferencesNil() {
	o.Preferences.Set(nil)
}

// UnsetPreferences ensures that no value is present for Preferences, not even an explicit nil
func (o *PrincipalResponse) UnsetPreferences() {
	o.Preferences.Unset()
}

// GetCustomDashboards returns the CustomDashboards field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalResponse) GetCustomDashboards() []Dashboard {
	if o == nil {
		var ret []Dashboard
		return ret
	}
	return o.CustomDashboards
}

// GetCustomDashboardsOk returns a tuple with the CustomDashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalResponse) GetCustomDashboardsOk() ([]Dashboard, bool) {
	if o == nil || IsNil(o.CustomDashboards) {
		return nil, false
	}
	return o.CustomDashboards, true
}

// HasCustomDashboards returns a boolean if a field has been set.
func (o *PrincipalResponse) HasCustomDashboards() bool {
	if o != nil && !IsNil(o.CustomDashboards) {
		return true
	}

	return false
}

// SetCustomDashboards gets a reference to the given []Dashboard and assigns it to the CustomDashboards field.
func (o *PrincipalResponse) SetCustomDashboards(v []Dashboard) {
	o.CustomDashboards = v
}

func (o PrincipalResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identity"] = o.Identity
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	if o.Preferences.IsSet() {
		toSerialize["preferences"] = o.Preferences.Get()
	}
	if o.CustomDashboards != nil {
		toSerialize["customDashboards"] = o.CustomDashboards
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identity",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPrincipalResponse := _PrincipalResponse{}

	err = json.Unmarshal(data, &varPrincipalResponse)

	if err != nil {
		return err
	}

	*o = PrincipalResponse(varPrincipalResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "teams")
		delete(additionalProperties, "preferences")
		delete(additionalProperties, "customDashboards")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalResponse struct {
	value *PrincipalResponse
	isSet bool
}

func (v NullablePrincipalResponse) Get() *PrincipalResponse {
	return v.value
}

func (v *NullablePrincipalResponse) Set(val *PrincipalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalResponse(val *PrincipalResponse) *NullablePrincipalResponse {
	return &NullablePrincipalResponse{value: val, isSet: true}
}

func (v NullablePrincipalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


