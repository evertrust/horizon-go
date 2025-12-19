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

// checks if the OidcIdentityProviderResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OidcIdentityProviderResponse{}

// OidcIdentityProviderResponse struct for OidcIdentityProviderResponse
type OidcIdentityProviderResponse struct {
	// The internal ID of the Identity Provider
	Id string `json:"_id"`
	// Name of the `password` [credentials](#tag/security.credentials) containing the client ID  and secret to use to authenticate Horizon against the identity provider
	ClientCredentials string `json:"clientCredentials"`
	// The description of the identity provider
	Description []LocalizedString `json:"description,omitempty"`
	// The display name of the identity provider
	DisplayName []LocalizedString `json:"displayName,omitempty"`
	// The OpenID information that will be used as the user's email in Horizon
	EmailClaim string `json:"emailClaim"`
	// Whether the identity provider can be used to identify against Horizon
	Enabled bool `json:"enabled"`
	// Whether the identity provider can be selected on login to the Horizon UI
	EnabledOnUI bool `json:"enabledOnUI"`
	// The OpenID information that will be used as the user's identifier in Horizon
	IdentifierClaim string `json:"identifierClaim"`
	// The internal name of the identity provider
	Name string `json:"name"`
	// The OpenID information that will be used as the user's name in Horizon
	NameClaim string `json:"nameClaim"`
	// The URL of the identity provider OpenID callback
	ProviderMetadataUrl string `json:"providerMetadataUrl"`
	// The name of the proxy to use to reach the identity provider
	Proxy utils.NullableString `json:"proxy,omitempty"`
	// The scope where to retrieve the user data from
	Scope string `json:"scope"`
	// The timeout value to use when connecting to the identity provider (must be a valid finite duration)
	Timeout utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Trust AC coming from the system trust store or only trust AC imported in Horizon
	TrustSystemCAs bool `json:"trustSystemCAs"`
	// The type of Identity provider to register
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _OidcIdentityProviderResponse OidcIdentityProviderResponse

// NewOidcIdentityProviderResponse instantiates a new OidcIdentityProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcIdentityProviderResponse(id string, clientCredentials string, emailClaim string, enabled bool, enabledOnUI bool, identifierClaim string, name string, nameClaim string, providerMetadataUrl string, scope string, trustSystemCAs bool, type_ string) *OidcIdentityProviderResponse {
	this := OidcIdentityProviderResponse{}
	this.Id = id
	this.ClientCredentials = clientCredentials
	this.EmailClaim = emailClaim
	this.Enabled = enabled
	this.EnabledOnUI = enabledOnUI
	this.IdentifierClaim = identifierClaim
	this.Name = name
	this.NameClaim = nameClaim
	this.ProviderMetadataUrl = providerMetadataUrl
	this.Scope = scope
	this.TrustSystemCAs = trustSystemCAs
	this.Type = type_
	return &this
}

// NewOidcIdentityProviderResponseWithDefaults instantiates a new OidcIdentityProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcIdentityProviderResponseWithDefaults() *OidcIdentityProviderResponse {
	this := OidcIdentityProviderResponse{}
	var trustSystemCAs bool = true
	this.TrustSystemCAs = trustSystemCAs
	return &this
}

// GetId returns the Id field value
func (o *OidcIdentityProviderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OidcIdentityProviderResponse) SetId(v string) {
	o.Id = v
}

// GetClientCredentials returns the ClientCredentials field value
func (o *OidcIdentityProviderResponse) GetClientCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientCredentials
}

// GetClientCredentialsOk returns a tuple with the ClientCredentials field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetClientCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientCredentials, true
}

// SetClientCredentials sets field value
func (o *OidcIdentityProviderResponse) SetClientCredentials(v string) {
	o.ClientCredentials = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OidcIdentityProviderResponse) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OidcIdentityProviderResponse) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OidcIdentityProviderResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *OidcIdentityProviderResponse) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OidcIdentityProviderResponse) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OidcIdentityProviderResponse) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OidcIdentityProviderResponse) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *OidcIdentityProviderResponse) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetEmailClaim returns the EmailClaim field value
func (o *OidcIdentityProviderResponse) GetEmailClaim() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailClaim
}

// GetEmailClaimOk returns a tuple with the EmailClaim field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetEmailClaimOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailClaim, true
}

// SetEmailClaim sets field value
func (o *OidcIdentityProviderResponse) SetEmailClaim(v string) {
	o.EmailClaim = v
}

// GetEnabled returns the Enabled field value
func (o *OidcIdentityProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *OidcIdentityProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnabledOnUI returns the EnabledOnUI field value
func (o *OidcIdentityProviderResponse) GetEnabledOnUI() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnabledOnUI
}

// GetEnabledOnUIOk returns a tuple with the EnabledOnUI field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetEnabledOnUIOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnabledOnUI, true
}

// SetEnabledOnUI sets field value
func (o *OidcIdentityProviderResponse) SetEnabledOnUI(v bool) {
	o.EnabledOnUI = v
}

// GetIdentifierClaim returns the IdentifierClaim field value
func (o *OidcIdentityProviderResponse) GetIdentifierClaim() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentifierClaim
}

// GetIdentifierClaimOk returns a tuple with the IdentifierClaim field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetIdentifierClaimOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentifierClaim, true
}

// SetIdentifierClaim sets field value
func (o *OidcIdentityProviderResponse) SetIdentifierClaim(v string) {
	o.IdentifierClaim = v
}

// GetName returns the Name field value
func (o *OidcIdentityProviderResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OidcIdentityProviderResponse) SetName(v string) {
	o.Name = v
}

// GetNameClaim returns the NameClaim field value
func (o *OidcIdentityProviderResponse) GetNameClaim() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameClaim
}

// GetNameClaimOk returns a tuple with the NameClaim field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetNameClaimOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameClaim, true
}

// SetNameClaim sets field value
func (o *OidcIdentityProviderResponse) SetNameClaim(v string) {
	o.NameClaim = v
}

// GetProviderMetadataUrl returns the ProviderMetadataUrl field value
func (o *OidcIdentityProviderResponse) GetProviderMetadataUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderMetadataUrl
}

// GetProviderMetadataUrlOk returns a tuple with the ProviderMetadataUrl field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetProviderMetadataUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderMetadataUrl, true
}

// SetProviderMetadataUrl sets field value
func (o *OidcIdentityProviderResponse) SetProviderMetadataUrl(v string) {
	o.ProviderMetadataUrl = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OidcIdentityProviderResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OidcIdentityProviderResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *OidcIdentityProviderResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *OidcIdentityProviderResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *OidcIdentityProviderResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *OidcIdentityProviderResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetScope returns the Scope field value
func (o *OidcIdentityProviderResponse) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *OidcIdentityProviderResponse) SetScope(v string) {
	o.Scope = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OidcIdentityProviderResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OidcIdentityProviderResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *OidcIdentityProviderResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *OidcIdentityProviderResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *OidcIdentityProviderResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *OidcIdentityProviderResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetTrustSystemCAs returns the TrustSystemCAs field value
func (o *OidcIdentityProviderResponse) GetTrustSystemCAs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TrustSystemCAs
}

// GetTrustSystemCAsOk returns a tuple with the TrustSystemCAs field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetTrustSystemCAsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustSystemCAs, true
}

// SetTrustSystemCAs sets field value
func (o *OidcIdentityProviderResponse) SetTrustSystemCAs(v bool) {
	o.TrustSystemCAs = v
}

// GetType returns the Type field value
func (o *OidcIdentityProviderResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OidcIdentityProviderResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OidcIdentityProviderResponse) SetType(v string) {
	o.Type = v
}

func (o OidcIdentityProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OidcIdentityProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["clientCredentials"] = o.ClientCredentials
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["emailClaim"] = o.EmailClaim
	toSerialize["enabled"] = o.Enabled
	toSerialize["enabledOnUI"] = o.EnabledOnUI
	toSerialize["identifierClaim"] = o.IdentifierClaim
	toSerialize["name"] = o.Name
	toSerialize["nameClaim"] = o.NameClaim
	toSerialize["providerMetadataUrl"] = o.ProviderMetadataUrl
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	toSerialize["scope"] = o.Scope
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	toSerialize["trustSystemCAs"] = o.TrustSystemCAs
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OidcIdentityProviderResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"clientCredentials",
		"emailClaim",
		"enabled",
		"enabledOnUI",
		"identifierClaim",
		"name",
		"nameClaim",
		"providerMetadataUrl",
		"scope",
		"trustSystemCAs",
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

	varOidcIdentityProviderResponse := _OidcIdentityProviderResponse{}

	err = json.Unmarshal(data, &varOidcIdentityProviderResponse)

	if err != nil {
		return err
	}

	*o = OidcIdentityProviderResponse(varOidcIdentityProviderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "clientCredentials")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "emailClaim")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "enabledOnUI")
		delete(additionalProperties, "identifierClaim")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nameClaim")
		delete(additionalProperties, "providerMetadataUrl")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "trustSystemCAs")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOidcIdentityProviderResponse struct {
	value *OidcIdentityProviderResponse
	isSet bool
}

func (v NullableOidcIdentityProviderResponse) Get() *OidcIdentityProviderResponse {
	return v.value
}

func (v *NullableOidcIdentityProviderResponse) Set(val *OidcIdentityProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcIdentityProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcIdentityProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcIdentityProviderResponse(val *OidcIdentityProviderResponse) *NullableOidcIdentityProviderResponse {
	return &NullableOidcIdentityProviderResponse{value: val, isSet: true}
}

func (v NullableOidcIdentityProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcIdentityProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
