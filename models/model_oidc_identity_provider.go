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

// checks if the OidcIdentityProvider type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OidcIdentityProvider{}

// OidcIdentityProvider struct for OidcIdentityProvider
type OidcIdentityProvider struct {
	// The internal name of the identity provider
	Name string `json:"name"`
	// The display name of the identity provider
	DisplayName []LocalizedString `json:"displayName,omitempty"`
	// The description of the identity provider
	Description []LocalizedString `json:"description,omitempty"`
	// The type of Identity provider to register
	Type string `json:"type"`
	// Whether the identity provider can be used to identify against Horizon
	Enabled bool `json:"enabled"`
	// Whether the identity provider can be selected on login to the Horizon UI
	EnabledOnUI bool `json:"enabledOnUI"`
	// The name of the proxy to use to reach the identity provider
	Proxy utils.NullableString `json:"proxy,omitempty"`
	// The timeout value to use when connecting to the identity provider (must be a valid finite duration)
	Timeout utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// The URL of the identity provider OpenID callback
	ProviderMetadataUrl string `json:"providerMetadataUrl"`
	// Name of the `password` [credentials](#tag/security.credentials) containing the client ID  and secret to use to authenticate Horizon against the identity provider
	ClientCredentials string `json:"clientCredentials"`
	// The scope where to retrieve the user data from
	Scope string `json:"scope"`
	// Trust AC coming from the system trust store or only trust AC imported in Horizon
	TrustSystemCAs bool `json:"trustSystemCAs"`
	// The OpenID information that will be used as the user's identifier in Horizon
	IdentifierClaim *string `json:"identifierClaim,omitempty"`
	// The OpenID information that will be used as the user's email in Horizon
	EmailClaim *string `json:"emailClaim,omitempty"`
	// The OpenID information that will be used as the user's name in Horizon
	NameClaim            *string `json:"nameClaim,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OidcIdentityProvider OidcIdentityProvider

// NewOidcIdentityProvider instantiates a new OidcIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcIdentityProvider(name string, type_ string, enabled bool, enabledOnUI bool, providerMetadataUrl string, clientCredentials string, scope string, trustSystemCAs bool) *OidcIdentityProvider {
	this := OidcIdentityProvider{}
	this.Name = name
	this.Type = type_
	this.Enabled = enabled
	this.EnabledOnUI = enabledOnUI
	this.ProviderMetadataUrl = providerMetadataUrl
	this.ClientCredentials = clientCredentials
	this.Scope = scope
	this.TrustSystemCAs = trustSystemCAs
	var identifierClaim string = "{{email}}"
	this.IdentifierClaim = &identifierClaim
	var emailClaim string = "{{email}}"
	this.EmailClaim = &emailClaim
	var nameClaim string = "{{name}}"
	this.NameClaim = &nameClaim
	return &this
}

// NewOidcIdentityProviderWithDefaults instantiates a new OidcIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcIdentityProviderWithDefaults() *OidcIdentityProvider {
	this := OidcIdentityProvider{}
	var trustSystemCAs bool = true
	this.TrustSystemCAs = trustSystemCAs
	var identifierClaim string = "{{email}}"
	this.IdentifierClaim = &identifierClaim
	var emailClaim string = "{{email}}"
	this.EmailClaim = &emailClaim
	var nameClaim string = "{{name}}"
	this.NameClaim = &nameClaim
	return &this
}

// GetName returns the Name field value
func (o *OidcIdentityProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OidcIdentityProvider) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OidcIdentityProvider) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OidcIdentityProvider) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OidcIdentityProvider) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *OidcIdentityProvider) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OidcIdentityProvider) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OidcIdentityProvider) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OidcIdentityProvider) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *OidcIdentityProvider) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetType returns the Type field value
func (o *OidcIdentityProvider) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProvider) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OidcIdentityProvider) SetType(v string) {
	o.Type = v
}

// GetEnabled returns the Enabled field value
func (o *OidcIdentityProvider) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProvider) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *OidcIdentityProvider) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnabledOnUI returns the EnabledOnUI field value
func (o *OidcIdentityProvider) GetEnabledOnUI() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnabledOnUI
}

// GetEnabledOnUIOk returns a tuple with the EnabledOnUI field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProvider) GetEnabledOnUIOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnabledOnUI, true
}

// SetEnabledOnUI sets field value
func (o *OidcIdentityProvider) SetEnabledOnUI(v bool) {
	o.EnabledOnUI = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OidcIdentityProvider) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OidcIdentityProvider) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *OidcIdentityProvider) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *OidcIdentityProvider) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *OidcIdentityProvider) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *OidcIdentityProvider) UnsetProxy() {
	o.Proxy.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OidcIdentityProvider) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OidcIdentityProvider) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *OidcIdentityProvider) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *OidcIdentityProvider) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *OidcIdentityProvider) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *OidcIdentityProvider) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetProviderMetadataUrl returns the ProviderMetadataUrl field value
func (o *OidcIdentityProvider) GetProviderMetadataUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderMetadataUrl
}

// GetProviderMetadataUrlOk returns a tuple with the ProviderMetadataUrl field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProvider) GetProviderMetadataUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderMetadataUrl, true
}

// SetProviderMetadataUrl sets field value
func (o *OidcIdentityProvider) SetProviderMetadataUrl(v string) {
	o.ProviderMetadataUrl = v
}

// GetClientCredentials returns the ClientCredentials field value
func (o *OidcIdentityProvider) GetClientCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientCredentials
}

// GetClientCredentialsOk returns a tuple with the ClientCredentials field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProvider) GetClientCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientCredentials, true
}

// SetClientCredentials sets field value
func (o *OidcIdentityProvider) SetClientCredentials(v string) {
	o.ClientCredentials = v
}

// GetScope returns the Scope field value
func (o *OidcIdentityProvider) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProvider) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *OidcIdentityProvider) SetScope(v string) {
	o.Scope = v
}

// GetTrustSystemCAs returns the TrustSystemCAs field value
func (o *OidcIdentityProvider) GetTrustSystemCAs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TrustSystemCAs
}

// GetTrustSystemCAsOk returns a tuple with the TrustSystemCAs field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProvider) GetTrustSystemCAsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustSystemCAs, true
}

// SetTrustSystemCAs sets field value
func (o *OidcIdentityProvider) SetTrustSystemCAs(v bool) {
	o.TrustSystemCAs = v
}

// GetIdentifierClaim returns the IdentifierClaim field value if set, zero value otherwise.
func (o *OidcIdentityProvider) GetIdentifierClaim() string {
	if o == nil || utils.IsNil(o.IdentifierClaim) {
		var ret string
		return ret
	}
	return *o.IdentifierClaim
}

// GetIdentifierClaimOk returns a tuple with the IdentifierClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcIdentityProvider) GetIdentifierClaimOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IdentifierClaim) {
		return nil, false
	}
	return o.IdentifierClaim, true
}

// HasIdentifierClaim returns a boolean if a field has been set.
func (o *OidcIdentityProvider) HasIdentifierClaim() bool {
	if o != nil && !utils.IsNil(o.IdentifierClaim) {
		return true
	}

	return false
}

// SetIdentifierClaim gets a reference to the given string and assigns it to the IdentifierClaim field.
func (o *OidcIdentityProvider) SetIdentifierClaim(v string) {
	o.IdentifierClaim = &v
}

// GetEmailClaim returns the EmailClaim field value if set, zero value otherwise.
func (o *OidcIdentityProvider) GetEmailClaim() string {
	if o == nil || utils.IsNil(o.EmailClaim) {
		var ret string
		return ret
	}
	return *o.EmailClaim
}

// GetEmailClaimOk returns a tuple with the EmailClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcIdentityProvider) GetEmailClaimOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EmailClaim) {
		return nil, false
	}
	return o.EmailClaim, true
}

// HasEmailClaim returns a boolean if a field has been set.
func (o *OidcIdentityProvider) HasEmailClaim() bool {
	if o != nil && !utils.IsNil(o.EmailClaim) {
		return true
	}

	return false
}

// SetEmailClaim gets a reference to the given string and assigns it to the EmailClaim field.
func (o *OidcIdentityProvider) SetEmailClaim(v string) {
	o.EmailClaim = &v
}

// GetNameClaim returns the NameClaim field value if set, zero value otherwise.
func (o *OidcIdentityProvider) GetNameClaim() string {
	if o == nil || utils.IsNil(o.NameClaim) {
		var ret string
		return ret
	}
	return *o.NameClaim
}

// GetNameClaimOk returns a tuple with the NameClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcIdentityProvider) GetNameClaimOk() (*string, bool) {
	if o == nil || utils.IsNil(o.NameClaim) {
		return nil, false
	}
	return o.NameClaim, true
}

// HasNameClaim returns a boolean if a field has been set.
func (o *OidcIdentityProvider) HasNameClaim() bool {
	if o != nil && !utils.IsNil(o.NameClaim) {
		return true
	}

	return false
}

// SetNameClaim gets a reference to the given string and assigns it to the NameClaim field.
func (o *OidcIdentityProvider) SetNameClaim(v string) {
	o.NameClaim = &v
}

func (o OidcIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OidcIdentityProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	toSerialize["enabled"] = o.Enabled
	toSerialize["enabledOnUI"] = o.EnabledOnUI
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	toSerialize["providerMetadataUrl"] = o.ProviderMetadataUrl
	toSerialize["clientCredentials"] = o.ClientCredentials
	toSerialize["scope"] = o.Scope
	toSerialize["trustSystemCAs"] = o.TrustSystemCAs
	if !utils.IsNil(o.IdentifierClaim) {
		toSerialize["identifierClaim"] = o.IdentifierClaim
	}
	if !utils.IsNil(o.EmailClaim) {
		toSerialize["emailClaim"] = o.EmailClaim
	}
	if !utils.IsNil(o.NameClaim) {
		toSerialize["nameClaim"] = o.NameClaim
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OidcIdentityProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"enabled",
		"enabledOnUI",
		"providerMetadataUrl",
		"clientCredentials",
		"scope",
		"trustSystemCAs",
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

	varOidcIdentityProvider := _OidcIdentityProvider{}

	err = json.Unmarshal(data, &varOidcIdentityProvider)

	if err != nil {
		return err
	}

	*o = OidcIdentityProvider(varOidcIdentityProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "enabledOnUI")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "providerMetadataUrl")
		delete(additionalProperties, "clientCredentials")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "trustSystemCAs")
		delete(additionalProperties, "identifierClaim")
		delete(additionalProperties, "emailClaim")
		delete(additionalProperties, "nameClaim")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOidcIdentityProvider struct {
	value *OidcIdentityProvider
	isSet bool
}

func (v NullableOidcIdentityProvider) Get() *OidcIdentityProvider {
	return v.value
}

func (v *NullableOidcIdentityProvider) Set(val *OidcIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcIdentityProvider(val *OidcIdentityProvider) *NullableOidcIdentityProvider {
	return &NullableOidcIdentityProvider{value: val, isSet: true}
}

func (v NullableOidcIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
