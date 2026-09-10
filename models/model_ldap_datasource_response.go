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

// checks if the LDAPDatasourceResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LDAPDatasourceResponse{}

// LDAPDatasourceResponse struct for LDAPDatasourceResponse
type LDAPDatasourceResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// List of attributes to fetch for this datasource
	Attributes []DataSourceOutput `json:"attributes,omitempty"`
	// LDAP Base DN
	BaseDn string `json:"baseDn"`
	// Name of the `password` [credentials](#tag/security.credentials) to use for LDAP Authentication
	Credentials string `json:"credentials"`
	// Description of the datasource
	Description *string `json:"description,omitempty"`
	// Disable hostname validation for the LDAP connection
	DisableHostnameValidation utils.NullableBool `json:"disableHostnameValidation,omitempty"`
	// The localized name of the datasource
	DisplayName []LocalizedString `json:"displayName,omitempty"`
	// LDAP Filter
	Filter string `json:"filter"`
	// Hostname of the LDAP server
	Hostname string `json:"hostname"`
	// Name of the datasource
	Name string `json:"name"`
	// Port on which to join the LDAP server
	Port utils.NullableInt64 `json:"port,omitempty"`
	// Name of the proxy to use to reach the LDAP server
	Proxy utils.NullableString `json:"proxy,omitempty"`
	// Use secure LDAP connection
	Secure bool `json:"secure"`
	// Timeout for the LDAP request
	Timeout string `json:"timeout" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Type of datasource
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _LDAPDatasourceResponse LDAPDatasourceResponse

// NewLDAPDatasourceResponse instantiates a new LDAPDatasourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLDAPDatasourceResponse(id string, baseDn string, credentials string, filter string, hostname string, name string, secure bool, timeout string, type_ string) *LDAPDatasourceResponse {
	this := LDAPDatasourceResponse{}
	this.Id = id
	this.BaseDn = baseDn
	this.Credentials = credentials
	this.Filter = filter
	this.Hostname = hostname
	this.Name = name
	this.Secure = secure
	this.Timeout = timeout
	this.Type = type_
	var disableHostnameValidation bool = false
	this.DisableHostnameValidation = *utils.NewNullableBool(&disableHostnameValidation)
	var port int64 = 389
	this.Port = *utils.NewNullableInt64(&port)
	return &this
}

// NewLDAPDatasourceResponseWithDefaults instantiates a new LDAPDatasourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLDAPDatasourceResponseWithDefaults() *LDAPDatasourceResponse {
	this := LDAPDatasourceResponse{}
	var disableHostnameValidation bool = false
	this.DisableHostnameValidation = *utils.NewNullableBool(&disableHostnameValidation)
	var port int64 = 389
	this.Port = *utils.NewNullableInt64(&port)
	return &this
}

// GetId returns the Id field value
func (o *LDAPDatasourceResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LDAPDatasourceResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LDAPDatasourceResponse) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPDatasourceResponse) GetAttributes() []DataSourceOutput {
	if o == nil {
		var ret []DataSourceOutput
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPDatasourceResponse) GetAttributesOk() ([]DataSourceOutput, bool) {
	if o == nil || utils.IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *LDAPDatasourceResponse) HasAttributes() bool {
	if o != nil && !utils.IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []DataSourceOutput and assigns it to the Attributes field.
func (o *LDAPDatasourceResponse) SetAttributes(v []DataSourceOutput) {
	o.Attributes = v
}

// GetBaseDn returns the BaseDn field value
func (o *LDAPDatasourceResponse) GetBaseDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseDn
}

// GetBaseDnOk returns a tuple with the BaseDn field value
// and a boolean to check if the value has been set.
func (o *LDAPDatasourceResponse) GetBaseDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseDn, true
}

// SetBaseDn sets field value
func (o *LDAPDatasourceResponse) SetBaseDn(v string) {
	o.BaseDn = v
}

// GetCredentials returns the Credentials field value
func (o *LDAPDatasourceResponse) GetCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *LDAPDatasourceResponse) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *LDAPDatasourceResponse) SetCredentials(v string) {
	o.Credentials = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LDAPDatasourceResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LDAPDatasourceResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LDAPDatasourceResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LDAPDatasourceResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDisableHostnameValidation returns the DisableHostnameValidation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPDatasourceResponse) GetDisableHostnameValidation() bool {
	if o == nil || utils.IsNil(o.DisableHostnameValidation.Get()) {
		var ret bool
		return ret
	}
	return *o.DisableHostnameValidation.Get()
}

// GetDisableHostnameValidationOk returns a tuple with the DisableHostnameValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPDatasourceResponse) GetDisableHostnameValidationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisableHostnameValidation.Get(), o.DisableHostnameValidation.IsSet()
}

// HasDisableHostnameValidation returns a boolean if a field has been set.
func (o *LDAPDatasourceResponse) HasDisableHostnameValidation() bool {
	if o != nil && o.DisableHostnameValidation.IsSet() {
		return true
	}

	return false
}

// SetDisableHostnameValidation gets a reference to the given NullableBool and assigns it to the DisableHostnameValidation field.
func (o *LDAPDatasourceResponse) SetDisableHostnameValidation(v bool) {
	o.DisableHostnameValidation.Set(&v)
}

// SetDisableHostnameValidationNil sets the value for DisableHostnameValidation to be an explicit nil
func (o *LDAPDatasourceResponse) SetDisableHostnameValidationNil() {
	o.DisableHostnameValidation.Set(nil)
}

// UnsetDisableHostnameValidation ensures that no value is present for DisableHostnameValidation, not even an explicit nil
func (o *LDAPDatasourceResponse) UnsetDisableHostnameValidation() {
	o.DisableHostnameValidation.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPDatasourceResponse) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPDatasourceResponse) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *LDAPDatasourceResponse) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *LDAPDatasourceResponse) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetFilter returns the Filter field value
func (o *LDAPDatasourceResponse) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *LDAPDatasourceResponse) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *LDAPDatasourceResponse) SetFilter(v string) {
	o.Filter = v
}

// GetHostname returns the Hostname field value
func (o *LDAPDatasourceResponse) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *LDAPDatasourceResponse) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *LDAPDatasourceResponse) SetHostname(v string) {
	o.Hostname = v
}

// GetName returns the Name field value
func (o *LDAPDatasourceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LDAPDatasourceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LDAPDatasourceResponse) SetName(v string) {
	o.Name = v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPDatasourceResponse) GetPort() int64 {
	if o == nil || utils.IsNil(o.Port.Get()) {
		var ret int64
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPDatasourceResponse) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *LDAPDatasourceResponse) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt64 and assigns it to the Port field.
func (o *LDAPDatasourceResponse) SetPort(v int64) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil
func (o *LDAPDatasourceResponse) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *LDAPDatasourceResponse) UnsetPort() {
	o.Port.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LDAPDatasourceResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LDAPDatasourceResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *LDAPDatasourceResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *LDAPDatasourceResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *LDAPDatasourceResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *LDAPDatasourceResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetSecure returns the Secure field value
func (o *LDAPDatasourceResponse) GetSecure() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Secure
}

// GetSecureOk returns a tuple with the Secure field value
// and a boolean to check if the value has been set.
func (o *LDAPDatasourceResponse) GetSecureOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secure, true
}

// SetSecure sets field value
func (o *LDAPDatasourceResponse) SetSecure(v bool) {
	o.Secure = v
}

// GetTimeout returns the Timeout field value
func (o *LDAPDatasourceResponse) GetTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *LDAPDatasourceResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *LDAPDatasourceResponse) SetTimeout(v string) {
	o.Timeout = v
}

// GetType returns the Type field value
func (o *LDAPDatasourceResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LDAPDatasourceResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LDAPDatasourceResponse) SetType(v string) {
	o.Type = v
}

func (o LDAPDatasourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LDAPDatasourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["baseDn"] = o.BaseDn
	toSerialize["credentials"] = o.Credentials
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.DisableHostnameValidation.IsSet() {
		toSerialize["disableHostnameValidation"] = o.DisableHostnameValidation.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["filter"] = o.Filter
	toSerialize["hostname"] = o.Hostname
	toSerialize["name"] = o.Name
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	toSerialize["secure"] = o.Secure
	toSerialize["timeout"] = o.Timeout
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LDAPDatasourceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"baseDn",
		"credentials",
		"filter",
		"hostname",
		"name",
		"secure",
		"timeout",
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

	varLDAPDatasourceResponse := _LDAPDatasourceResponse{}

	err = json.Unmarshal(data, &varLDAPDatasourceResponse)

	if err != nil {
		return err
	}

	*o = LDAPDatasourceResponse(varLDAPDatasourceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "baseDn")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "description")
		delete(additionalProperties, "disableHostnameValidation")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "name")
		delete(additionalProperties, "port")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "secure")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLDAPDatasourceResponse struct {
	value *LDAPDatasourceResponse
	isSet bool
}

func (v NullableLDAPDatasourceResponse) Get() *LDAPDatasourceResponse {
	return v.value
}

func (v *NullableLDAPDatasourceResponse) Set(val *LDAPDatasourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLDAPDatasourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLDAPDatasourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLDAPDatasourceResponse(val *LDAPDatasourceResponse) *NullableLDAPDatasourceResponse {
	return &NullableLDAPDatasourceResponse{value: val, isSet: true}
}

func (v NullableLDAPDatasourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLDAPDatasourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
