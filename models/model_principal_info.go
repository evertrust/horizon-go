/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the PrincipalInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PrincipalInfo{}

// PrincipalInfo struct for PrincipalInfo
type PrincipalInfo struct {
	// The contact e-mail of the principal
	Contact utils.NullableString `json:"contact,omitempty"`
	// The custom dashboards of the principal. This is used by UI only. These values should not be manually set but should be copied on update
	CustomDashboards []Dashboard `json:"customDashboards,omitempty"`
	// If the principal is allowed to login horizon
	Enabled bool `json:"enabled"`
	// The identifier of the principal
	Identifier string `json:"identifier"`
	// The permissions of the principal
	Permissions []Permission `json:"permissions,omitempty"`
	// The UI preferences of the principal. This is used by UI only. These values should not be manually set but should be copied on update
	Preferences NullablePrincipalInfoPreferences `json:"preferences,omitempty"`
	// The roles of the principal
	Roles []string `json:"roles,omitempty"`
	// The saved HQL queries of the principal. This is used by UI only. These values should not be manually set but should be copied on update
	SavedQueries []PrincipalInfoSavedQuery `json:"savedQueries,omitempty"`
	// The teams of the principal
	Teams                []string `json:"teams,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalInfo PrincipalInfo

// NewPrincipalInfo instantiates a new PrincipalInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalInfo(enabled bool, identifier string) *PrincipalInfo {
	this := PrincipalInfo{}
	this.Enabled = enabled
	this.Identifier = identifier
	return &this
}

// NewPrincipalInfoWithDefaults instantiates a new PrincipalInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalInfoWithDefaults() *PrincipalInfo {
	this := PrincipalInfo{}
	return &this
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfo) GetContact() string {
	if o == nil || utils.IsNil(o.Contact.Get()) {
		var ret string
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfo) GetContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *PrincipalInfo) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableString and assigns it to the Contact field.
func (o *PrincipalInfo) SetContact(v string) {
	o.Contact.Set(&v)
}

// SetContactNil sets the value for Contact to be an explicit nil
func (o *PrincipalInfo) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *PrincipalInfo) UnsetContact() {
	o.Contact.Unset()
}

// GetCustomDashboards returns the CustomDashboards field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfo) GetCustomDashboards() []Dashboard {
	if o == nil {
		var ret []Dashboard
		return ret
	}
	return o.CustomDashboards
}

// GetCustomDashboardsOk returns a tuple with the CustomDashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfo) GetCustomDashboardsOk() ([]Dashboard, bool) {
	if o == nil || utils.IsNil(o.CustomDashboards) {
		return nil, false
	}
	return o.CustomDashboards, true
}

// HasCustomDashboards returns a boolean if a field has been set.
func (o *PrincipalInfo) HasCustomDashboards() bool {
	if o != nil && !utils.IsNil(o.CustomDashboards) {
		return true
	}

	return false
}

// SetCustomDashboards gets a reference to the given []Dashboard and assigns it to the CustomDashboards field.
func (o *PrincipalInfo) SetCustomDashboards(v []Dashboard) {
	o.CustomDashboards = v
}

// GetEnabled returns the Enabled field value
func (o *PrincipalInfo) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PrincipalInfo) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PrincipalInfo) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIdentifier returns the Identifier field value
func (o *PrincipalInfo) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *PrincipalInfo) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *PrincipalInfo) SetIdentifier(v string) {
	o.Identifier = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfo) GetPermissions() []Permission {
	if o == nil {
		var ret []Permission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfo) GetPermissionsOk() ([]Permission, bool) {
	if o == nil || utils.IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PrincipalInfo) HasPermissions() bool {
	if o != nil && !utils.IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []Permission and assigns it to the Permissions field.
func (o *PrincipalInfo) SetPermissions(v []Permission) {
	o.Permissions = v
}

// GetPreferences returns the Preferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfo) GetPreferences() PrincipalInfoPreferences {
	if o == nil || utils.IsNil(o.Preferences.Get()) {
		var ret PrincipalInfoPreferences
		return ret
	}
	return *o.Preferences.Get()
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfo) GetPreferencesOk() (*PrincipalInfoPreferences, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preferences.Get(), o.Preferences.IsSet()
}

// HasPreferences returns a boolean if a field has been set.
func (o *PrincipalInfo) HasPreferences() bool {
	if o != nil && o.Preferences.IsSet() {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given NullablePrincipalInfoPreferences and assigns it to the Preferences field.
func (o *PrincipalInfo) SetPreferences(v PrincipalInfoPreferences) {
	o.Preferences.Set(&v)
}

// SetPreferencesNil sets the value for Preferences to be an explicit nil
func (o *PrincipalInfo) SetPreferencesNil() {
	o.Preferences.Set(nil)
}

// UnsetPreferences ensures that no value is present for Preferences, not even an explicit nil
func (o *PrincipalInfo) UnsetPreferences() {
	o.Preferences.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfo) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfo) GetRolesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *PrincipalInfo) HasRoles() bool {
	if o != nil && !utils.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *PrincipalInfo) SetRoles(v []string) {
	o.Roles = v
}

// GetSavedQueries returns the SavedQueries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfo) GetSavedQueries() []PrincipalInfoSavedQuery {
	if o == nil {
		var ret []PrincipalInfoSavedQuery
		return ret
	}
	return o.SavedQueries
}

// GetSavedQueriesOk returns a tuple with the SavedQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfo) GetSavedQueriesOk() ([]PrincipalInfoSavedQuery, bool) {
	if o == nil || utils.IsNil(o.SavedQueries) {
		return nil, false
	}
	return o.SavedQueries, true
}

// HasSavedQueries returns a boolean if a field has been set.
func (o *PrincipalInfo) HasSavedQueries() bool {
	if o != nil && !utils.IsNil(o.SavedQueries) {
		return true
	}

	return false
}

// SetSavedQueries gets a reference to the given []PrincipalInfoSavedQuery and assigns it to the SavedQueries field.
func (o *PrincipalInfo) SetSavedQueries(v []PrincipalInfoSavedQuery) {
	o.SavedQueries = v
}

// GetTeams returns the Teams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfo) GetTeams() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfo) GetTeamsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *PrincipalInfo) HasTeams() bool {
	if o != nil && !utils.IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *PrincipalInfo) SetTeams(v []string) {
	o.Teams = v
}

func (o PrincipalInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.CustomDashboards != nil {
		toSerialize["customDashboards"] = o.CustomDashboards
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["identifier"] = o.Identifier
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.Preferences.IsSet() {
		toSerialize["preferences"] = o.Preferences.Get()
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.SavedQueries != nil {
		toSerialize["savedQueries"] = o.SavedQueries
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"identifier",
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

	varPrincipalInfo := _PrincipalInfo{}

	err = json.Unmarshal(data, &varPrincipalInfo)

	if err != nil {
		return err
	}

	*o = PrincipalInfo(varPrincipalInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contact")
		delete(additionalProperties, "customDashboards")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "preferences")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "savedQueries")
		delete(additionalProperties, "teams")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalInfo struct {
	value *PrincipalInfo
	isSet bool
}

func (v NullablePrincipalInfo) Get() *PrincipalInfo {
	return v.value
}

func (v *NullablePrincipalInfo) Set(val *PrincipalInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalInfo(val *PrincipalInfo) *NullablePrincipalInfo {
	return &NullablePrincipalInfo{value: val, isSet: true}
}

func (v NullablePrincipalInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
