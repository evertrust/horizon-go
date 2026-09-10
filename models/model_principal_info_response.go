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

// checks if the PrincipalInfoResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PrincipalInfoResponse{}

// PrincipalInfoResponse struct for PrincipalInfoResponse
type PrincipalInfoResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// The contact e-mail of the principal
	Contact utils.NullableString `json:"contact,omitempty"`
	// The creation date of the principal (UNIX Timestamp in milliseconds)
	CreationDate *int64 `json:"creationDate,omitempty"`
	// The custom dashboards of the principal. This is used by UI only. These values should not be manually set but should be copied on update
	CustomDashboards []Dashboard `json:"customDashboards,omitempty"`
	// If the principal is allowed to login horizon
	Enabled bool `json:"enabled"`
	// The identifier of the principal
	Identifier string `json:"identifier"`
	// The last authentication date of the principal (UNIX Timestamp in milliseconds)
	LastAuthentication utils.NullableInt64 `json:"lastAuthentication,omitempty"`
	// The last modification date of the principal (UNIX Timestamp in milliseconds)
	LastModification utils.NullableInt64 `json:"lastModification,omitempty"`
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

type _PrincipalInfoResponse PrincipalInfoResponse

// NewPrincipalInfoResponse instantiates a new PrincipalInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalInfoResponse(id string, enabled bool, identifier string) *PrincipalInfoResponse {
	this := PrincipalInfoResponse{}
	this.Id = id
	this.Enabled = enabled
	this.Identifier = identifier
	return &this
}

// NewPrincipalInfoResponseWithDefaults instantiates a new PrincipalInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalInfoResponseWithDefaults() *PrincipalInfoResponse {
	this := PrincipalInfoResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PrincipalInfoResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PrincipalInfoResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PrincipalInfoResponse) SetId(v string) {
	o.Id = v
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoResponse) GetContact() string {
	if o == nil || utils.IsNil(o.Contact.Get()) {
		var ret string
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoResponse) GetContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *PrincipalInfoResponse) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableString and assigns it to the Contact field.
func (o *PrincipalInfoResponse) SetContact(v string) {
	o.Contact.Set(&v)
}

// SetContactNil sets the value for Contact to be an explicit nil
func (o *PrincipalInfoResponse) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *PrincipalInfoResponse) UnsetContact() {
	o.Contact.Unset()
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *PrincipalInfoResponse) GetCreationDate() int64 {
	if o == nil || utils.IsNil(o.CreationDate) {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalInfoResponse) GetCreationDateOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *PrincipalInfoResponse) HasCreationDate() bool {
	if o != nil && !utils.IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *PrincipalInfoResponse) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetCustomDashboards returns the CustomDashboards field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoResponse) GetCustomDashboards() []Dashboard {
	if o == nil {
		var ret []Dashboard
		return ret
	}
	return o.CustomDashboards
}

// GetCustomDashboardsOk returns a tuple with the CustomDashboards field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoResponse) GetCustomDashboardsOk() ([]Dashboard, bool) {
	if o == nil || utils.IsNil(o.CustomDashboards) {
		return nil, false
	}
	return o.CustomDashboards, true
}

// HasCustomDashboards returns a boolean if a field has been set.
func (o *PrincipalInfoResponse) HasCustomDashboards() bool {
	if o != nil && !utils.IsNil(o.CustomDashboards) {
		return true
	}

	return false
}

// SetCustomDashboards gets a reference to the given []Dashboard and assigns it to the CustomDashboards field.
func (o *PrincipalInfoResponse) SetCustomDashboards(v []Dashboard) {
	o.CustomDashboards = v
}

// GetEnabled returns the Enabled field value
func (o *PrincipalInfoResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PrincipalInfoResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PrincipalInfoResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIdentifier returns the Identifier field value
func (o *PrincipalInfoResponse) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *PrincipalInfoResponse) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *PrincipalInfoResponse) SetIdentifier(v string) {
	o.Identifier = v
}

// GetLastAuthentication returns the LastAuthentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoResponse) GetLastAuthentication() int64 {
	if o == nil || utils.IsNil(o.LastAuthentication.Get()) {
		var ret int64
		return ret
	}
	return *o.LastAuthentication.Get()
}

// GetLastAuthenticationOk returns a tuple with the LastAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoResponse) GetLastAuthenticationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastAuthentication.Get(), o.LastAuthentication.IsSet()
}

// HasLastAuthentication returns a boolean if a field has been set.
func (o *PrincipalInfoResponse) HasLastAuthentication() bool {
	if o != nil && o.LastAuthentication.IsSet() {
		return true
	}

	return false
}

// SetLastAuthentication gets a reference to the given NullableInt64 and assigns it to the LastAuthentication field.
func (o *PrincipalInfoResponse) SetLastAuthentication(v int64) {
	o.LastAuthentication.Set(&v)
}

// SetLastAuthenticationNil sets the value for LastAuthentication to be an explicit nil
func (o *PrincipalInfoResponse) SetLastAuthenticationNil() {
	o.LastAuthentication.Set(nil)
}

// UnsetLastAuthentication ensures that no value is present for LastAuthentication, not even an explicit nil
func (o *PrincipalInfoResponse) UnsetLastAuthentication() {
	o.LastAuthentication.Unset()
}

// GetLastModification returns the LastModification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoResponse) GetLastModification() int64 {
	if o == nil || utils.IsNil(o.LastModification.Get()) {
		var ret int64
		return ret
	}
	return *o.LastModification.Get()
}

// GetLastModificationOk returns a tuple with the LastModification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoResponse) GetLastModificationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastModification.Get(), o.LastModification.IsSet()
}

// HasLastModification returns a boolean if a field has been set.
func (o *PrincipalInfoResponse) HasLastModification() bool {
	if o != nil && o.LastModification.IsSet() {
		return true
	}

	return false
}

// SetLastModification gets a reference to the given NullableInt64 and assigns it to the LastModification field.
func (o *PrincipalInfoResponse) SetLastModification(v int64) {
	o.LastModification.Set(&v)
}

// SetLastModificationNil sets the value for LastModification to be an explicit nil
func (o *PrincipalInfoResponse) SetLastModificationNil() {
	o.LastModification.Set(nil)
}

// UnsetLastModification ensures that no value is present for LastModification, not even an explicit nil
func (o *PrincipalInfoResponse) UnsetLastModification() {
	o.LastModification.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoResponse) GetPermissions() []Permission {
	if o == nil {
		var ret []Permission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoResponse) GetPermissionsOk() ([]Permission, bool) {
	if o == nil || utils.IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PrincipalInfoResponse) HasPermissions() bool {
	if o != nil && !utils.IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []Permission and assigns it to the Permissions field.
func (o *PrincipalInfoResponse) SetPermissions(v []Permission) {
	o.Permissions = v
}

// GetPreferences returns the Preferences field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoResponse) GetPreferences() PrincipalInfoPreferences {
	if o == nil || utils.IsNil(o.Preferences.Get()) {
		var ret PrincipalInfoPreferences
		return ret
	}
	return *o.Preferences.Get()
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoResponse) GetPreferencesOk() (*PrincipalInfoPreferences, bool) {
	if o == nil {
		return nil, false
	}
	return o.Preferences.Get(), o.Preferences.IsSet()
}

// HasPreferences returns a boolean if a field has been set.
func (o *PrincipalInfoResponse) HasPreferences() bool {
	if o != nil && o.Preferences.IsSet() {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given NullablePrincipalInfoPreferences and assigns it to the Preferences field.
func (o *PrincipalInfoResponse) SetPreferences(v PrincipalInfoPreferences) {
	o.Preferences.Set(&v)
}

// SetPreferencesNil sets the value for Preferences to be an explicit nil
func (o *PrincipalInfoResponse) SetPreferencesNil() {
	o.Preferences.Set(nil)
}

// UnsetPreferences ensures that no value is present for Preferences, not even an explicit nil
func (o *PrincipalInfoResponse) UnsetPreferences() {
	o.Preferences.Unset()
}

// GetRoles returns the Roles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoResponse) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoResponse) GetRolesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *PrincipalInfoResponse) HasRoles() bool {
	if o != nil && !utils.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *PrincipalInfoResponse) SetRoles(v []string) {
	o.Roles = v
}

// GetSavedQueries returns the SavedQueries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoResponse) GetSavedQueries() []PrincipalInfoSavedQuery {
	if o == nil {
		var ret []PrincipalInfoSavedQuery
		return ret
	}
	return o.SavedQueries
}

// GetSavedQueriesOk returns a tuple with the SavedQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoResponse) GetSavedQueriesOk() ([]PrincipalInfoSavedQuery, bool) {
	if o == nil || utils.IsNil(o.SavedQueries) {
		return nil, false
	}
	return o.SavedQueries, true
}

// HasSavedQueries returns a boolean if a field has been set.
func (o *PrincipalInfoResponse) HasSavedQueries() bool {
	if o != nil && !utils.IsNil(o.SavedQueries) {
		return true
	}

	return false
}

// SetSavedQueries gets a reference to the given []PrincipalInfoSavedQuery and assigns it to the SavedQueries field.
func (o *PrincipalInfoResponse) SetSavedQueries(v []PrincipalInfoSavedQuery) {
	o.SavedQueries = v
}

// GetTeams returns the Teams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PrincipalInfoResponse) GetTeams() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PrincipalInfoResponse) GetTeamsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *PrincipalInfoResponse) HasTeams() bool {
	if o != nil && !utils.IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *PrincipalInfoResponse) SetTeams(v []string) {
	o.Teams = v
}

func (o PrincipalInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if !utils.IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.CustomDashboards != nil {
		toSerialize["customDashboards"] = o.CustomDashboards
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["identifier"] = o.Identifier
	if o.LastAuthentication.IsSet() {
		toSerialize["lastAuthentication"] = o.LastAuthentication.Get()
	}
	if o.LastModification.IsSet() {
		toSerialize["lastModification"] = o.LastModification.Get()
	}
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

func (o *PrincipalInfoResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
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

	varPrincipalInfoResponse := _PrincipalInfoResponse{}

	err = json.Unmarshal(data, &varPrincipalInfoResponse)

	if err != nil {
		return err
	}

	*o = PrincipalInfoResponse(varPrincipalInfoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "creationDate")
		delete(additionalProperties, "customDashboards")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "lastAuthentication")
		delete(additionalProperties, "lastModification")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "preferences")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "savedQueries")
		delete(additionalProperties, "teams")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalInfoResponse struct {
	value *PrincipalInfoResponse
	isSet bool
}

func (v NullablePrincipalInfoResponse) Get() *PrincipalInfoResponse {
	return v.value
}

func (v *NullablePrincipalInfoResponse) Set(val *PrincipalInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalInfoResponse(val *PrincipalInfoResponse) *NullablePrincipalInfoResponse {
	return &NullablePrincipalInfoResponse{value: val, isSet: true}
}

func (v NullablePrincipalInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
