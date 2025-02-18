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

// checks if the AcmeProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcmeProfileResponse{}

// AcmeProfileResponse struct for AcmeProfileResponse
type AcmeProfileResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	Module string `json:"module"`
	Name string `json:"name"`
	DisplayName []LocalizedString `json:"displayName,omitempty"`
	Description []LocalizedString `json:"description,omitempty"`
	Enabled bool `json:"enabled"`
	Timeout string `json:"timeout" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Meta NullableDirectoryMeta `json:"meta,omitempty"`
	Constraints NullableCertificateRequestConstraints `json:"constraints,omitempty"`
	AuthorizationMethods []string `json:"authorizationMethods,omitempty"`
	PkiConnector string `json:"pkiConnector"`
	Http01Port NullableInt64 `json:"http01Port,omitempty"`
	TlsAlpn01Port NullableInt64 `json:"tlsAlpn01Port,omitempty"`
	AuthorizeShortName bool `json:"authorizeShortName"`
	AuthorizeEmptyContact bool `json:"authorizeEmptyContact"`
	DefaultContacts []string `json:"defaultContacts,omitempty"`
	VerifyRetryCount int64 `json:"verifyRetryCount"`
	VerifyRetryDelay string `json:"verifyRetryDelay" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	RequireTermsOfService bool `json:"requireTermsOfService"`
	RenewalPeriod NullableString `json:"renewalPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	CsrDataMapping map[string]string `json:"csrDataMapping,omitempty"`
	MaxCertificatePerHolderPolicy NullableMaxCertificatePerHolderPolicy `json:"maxCertificatePerHolderPolicy,omitempty"`
	MaxDnsName NullableInt64 `json:"maxDnsName,omitempty"`
	Proxy NullableString `json:"proxy,omitempty"`
	AuthorizationLevels CertificateProfileAuthorizationLevels `json:"authorizationLevels"`
	Triggers NullableCertificateProfileTriggers `json:"triggers,omitempty"`
	RequestsPolicy RequestsPolicy `json:"requestsPolicy"`
	SelfPermissions CertificateProfileSelfPermissions `json:"selfPermissions"`
	CertificateTemplate NullableCertificateTemplate `json:"certificateTemplate,omitempty"`
	CryptoPolicy CertificateProfileCryptoPolicy `json:"cryptoPolicy"`
	GradingPolicies []string `json:"gradingPolicies,omitempty"`
	// Representation of a datasource execution flow
	DsFlow []DataSourceFlowEntry `json:"dsFlow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcmeProfileResponse AcmeProfileResponse

// NewAcmeProfileResponse instantiates a new AcmeProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcmeProfileResponse(id string, module string, name string, enabled bool, timeout string, pkiConnector string, authorizeShortName bool, authorizeEmptyContact bool, verifyRetryCount int64, verifyRetryDelay string, requireTermsOfService bool, authorizationLevels CertificateProfileAuthorizationLevels, requestsPolicy RequestsPolicy, selfPermissions CertificateProfileSelfPermissions, cryptoPolicy CertificateProfileCryptoPolicy) *AcmeProfileResponse {
	this := AcmeProfileResponse{}
	this.Id = id
	this.Module = module
	this.Name = name
	this.Enabled = enabled
	this.Timeout = timeout
	this.PkiConnector = pkiConnector
	this.AuthorizeShortName = authorizeShortName
	this.AuthorizeEmptyContact = authorizeEmptyContact
	this.VerifyRetryCount = verifyRetryCount
	this.VerifyRetryDelay = verifyRetryDelay
	this.RequireTermsOfService = requireTermsOfService
	this.AuthorizationLevels = authorizationLevels
	this.RequestsPolicy = requestsPolicy
	this.SelfPermissions = selfPermissions
	this.CryptoPolicy = cryptoPolicy
	return &this
}

// NewAcmeProfileResponseWithDefaults instantiates a new AcmeProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcmeProfileResponseWithDefaults() *AcmeProfileResponse {
	this := AcmeProfileResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AcmeProfileResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AcmeProfileResponse) SetId(v string) {
	o.Id = v
}

// GetModule returns the Module field value
func (o *AcmeProfileResponse) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *AcmeProfileResponse) SetModule(v string) {
	o.Module = v
}

// GetName returns the Name field value
func (o *AcmeProfileResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AcmeProfileResponse) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *AcmeProfileResponse) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *AcmeProfileResponse) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetEnabled returns the Enabled field value
func (o *AcmeProfileResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AcmeProfileResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetTimeout returns the Timeout field value
func (o *AcmeProfileResponse) GetTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *AcmeProfileResponse) SetTimeout(v string) {
	o.Timeout = v
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetMeta() DirectoryMeta {
	if o == nil || IsNil(o.Meta.Get()) {
		var ret DirectoryMeta
		return ret
	}
	return *o.Meta.Get()
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetMetaOk() (*DirectoryMeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Meta.Get(), o.Meta.IsSet()
}

// HasMeta returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasMeta() bool {
	if o != nil && o.Meta.IsSet() {
		return true
	}

	return false
}

// SetMeta gets a reference to the given NullableDirectoryMeta and assigns it to the Meta field.
func (o *AcmeProfileResponse) SetMeta(v DirectoryMeta) {
	o.Meta.Set(&v)
}
// SetMetaNil sets the value for Meta to be an explicit nil
func (o *AcmeProfileResponse) SetMetaNil() {
	o.Meta.Set(nil)
}

// UnsetMeta ensures that no value is present for Meta, not even an explicit nil
func (o *AcmeProfileResponse) UnsetMeta() {
	o.Meta.Unset()
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetConstraints() CertificateRequestConstraints {
	if o == nil || IsNil(o.Constraints.Get()) {
		var ret CertificateRequestConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetConstraintsOk() (*CertificateRequestConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableCertificateRequestConstraints and assigns it to the Constraints field.
func (o *AcmeProfileResponse) SetConstraints(v CertificateRequestConstraints) {
	o.Constraints.Set(&v)
}
// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *AcmeProfileResponse) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *AcmeProfileResponse) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetAuthorizationMethods returns the AuthorizationMethods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetAuthorizationMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AuthorizationMethods
}

// GetAuthorizationMethodsOk returns a tuple with the AuthorizationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetAuthorizationMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthorizationMethods) {
		return nil, false
	}
	return o.AuthorizationMethods, true
}

// HasAuthorizationMethods returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasAuthorizationMethods() bool {
	if o != nil && !IsNil(o.AuthorizationMethods) {
		return true
	}

	return false
}

// SetAuthorizationMethods gets a reference to the given []string and assigns it to the AuthorizationMethods field.
func (o *AcmeProfileResponse) SetAuthorizationMethods(v []string) {
	o.AuthorizationMethods = v
}

// GetPkiConnector returns the PkiConnector field value
func (o *AcmeProfileResponse) GetPkiConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PkiConnector
}

// GetPkiConnectorOk returns a tuple with the PkiConnector field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetPkiConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiConnector, true
}

// SetPkiConnector sets field value
func (o *AcmeProfileResponse) SetPkiConnector(v string) {
	o.PkiConnector = v
}

// GetHttp01Port returns the Http01Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetHttp01Port() int64 {
	if o == nil || IsNil(o.Http01Port.Get()) {
		var ret int64
		return ret
	}
	return *o.Http01Port.Get()
}

// GetHttp01PortOk returns a tuple with the Http01Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetHttp01PortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Http01Port.Get(), o.Http01Port.IsSet()
}

// HasHttp01Port returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasHttp01Port() bool {
	if o != nil && o.Http01Port.IsSet() {
		return true
	}

	return false
}

// SetHttp01Port gets a reference to the given NullableInt64 and assigns it to the Http01Port field.
func (o *AcmeProfileResponse) SetHttp01Port(v int64) {
	o.Http01Port.Set(&v)
}
// SetHttp01PortNil sets the value for Http01Port to be an explicit nil
func (o *AcmeProfileResponse) SetHttp01PortNil() {
	o.Http01Port.Set(nil)
}

// UnsetHttp01Port ensures that no value is present for Http01Port, not even an explicit nil
func (o *AcmeProfileResponse) UnsetHttp01Port() {
	o.Http01Port.Unset()
}

// GetTlsAlpn01Port returns the TlsAlpn01Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetTlsAlpn01Port() int64 {
	if o == nil || IsNil(o.TlsAlpn01Port.Get()) {
		var ret int64
		return ret
	}
	return *o.TlsAlpn01Port.Get()
}

// GetTlsAlpn01PortOk returns a tuple with the TlsAlpn01Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetTlsAlpn01PortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsAlpn01Port.Get(), o.TlsAlpn01Port.IsSet()
}

// HasTlsAlpn01Port returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasTlsAlpn01Port() bool {
	if o != nil && o.TlsAlpn01Port.IsSet() {
		return true
	}

	return false
}

// SetTlsAlpn01Port gets a reference to the given NullableInt64 and assigns it to the TlsAlpn01Port field.
func (o *AcmeProfileResponse) SetTlsAlpn01Port(v int64) {
	o.TlsAlpn01Port.Set(&v)
}
// SetTlsAlpn01PortNil sets the value for TlsAlpn01Port to be an explicit nil
func (o *AcmeProfileResponse) SetTlsAlpn01PortNil() {
	o.TlsAlpn01Port.Set(nil)
}

// UnsetTlsAlpn01Port ensures that no value is present for TlsAlpn01Port, not even an explicit nil
func (o *AcmeProfileResponse) UnsetTlsAlpn01Port() {
	o.TlsAlpn01Port.Unset()
}

// GetAuthorizeShortName returns the AuthorizeShortName field value
func (o *AcmeProfileResponse) GetAuthorizeShortName() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AuthorizeShortName
}

// GetAuthorizeShortNameOk returns a tuple with the AuthorizeShortName field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetAuthorizeShortNameOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizeShortName, true
}

// SetAuthorizeShortName sets field value
func (o *AcmeProfileResponse) SetAuthorizeShortName(v bool) {
	o.AuthorizeShortName = v
}

// GetAuthorizeEmptyContact returns the AuthorizeEmptyContact field value
func (o *AcmeProfileResponse) GetAuthorizeEmptyContact() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AuthorizeEmptyContact
}

// GetAuthorizeEmptyContactOk returns a tuple with the AuthorizeEmptyContact field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetAuthorizeEmptyContactOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizeEmptyContact, true
}

// SetAuthorizeEmptyContact sets field value
func (o *AcmeProfileResponse) SetAuthorizeEmptyContact(v bool) {
	o.AuthorizeEmptyContact = v
}

// GetDefaultContacts returns the DefaultContacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetDefaultContacts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DefaultContacts
}

// GetDefaultContactsOk returns a tuple with the DefaultContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetDefaultContactsOk() ([]string, bool) {
	if o == nil || IsNil(o.DefaultContacts) {
		return nil, false
	}
	return o.DefaultContacts, true
}

// HasDefaultContacts returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasDefaultContacts() bool {
	if o != nil && !IsNil(o.DefaultContacts) {
		return true
	}

	return false
}

// SetDefaultContacts gets a reference to the given []string and assigns it to the DefaultContacts field.
func (o *AcmeProfileResponse) SetDefaultContacts(v []string) {
	o.DefaultContacts = v
}

// GetVerifyRetryCount returns the VerifyRetryCount field value
func (o *AcmeProfileResponse) GetVerifyRetryCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VerifyRetryCount
}

// GetVerifyRetryCountOk returns a tuple with the VerifyRetryCount field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetVerifyRetryCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyRetryCount, true
}

// SetVerifyRetryCount sets field value
func (o *AcmeProfileResponse) SetVerifyRetryCount(v int64) {
	o.VerifyRetryCount = v
}

// GetVerifyRetryDelay returns the VerifyRetryDelay field value
func (o *AcmeProfileResponse) GetVerifyRetryDelay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerifyRetryDelay
}

// GetVerifyRetryDelayOk returns a tuple with the VerifyRetryDelay field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetVerifyRetryDelayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyRetryDelay, true
}

// SetVerifyRetryDelay sets field value
func (o *AcmeProfileResponse) SetVerifyRetryDelay(v string) {
	o.VerifyRetryDelay = v
}

// GetRequireTermsOfService returns the RequireTermsOfService field value
func (o *AcmeProfileResponse) GetRequireTermsOfService() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireTermsOfService
}

// GetRequireTermsOfServiceOk returns a tuple with the RequireTermsOfService field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetRequireTermsOfServiceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireTermsOfService, true
}

// SetRequireTermsOfService sets field value
func (o *AcmeProfileResponse) SetRequireTermsOfService(v bool) {
	o.RequireTermsOfService = v
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetRenewalPeriod() string {
	if o == nil || IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *AcmeProfileResponse) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}
// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *AcmeProfileResponse) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *AcmeProfileResponse) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetCsrDataMapping returns the CsrDataMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetCsrDataMapping() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.CsrDataMapping
}

// GetCsrDataMappingOk returns a tuple with the CsrDataMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetCsrDataMappingOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CsrDataMapping) {
		return nil, false
	}
	return &o.CsrDataMapping, true
}

// HasCsrDataMapping returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasCsrDataMapping() bool {
	if o != nil && !IsNil(o.CsrDataMapping) {
		return true
	}

	return false
}

// SetCsrDataMapping gets a reference to the given map[string]string and assigns it to the CsrDataMapping field.
func (o *AcmeProfileResponse) SetCsrDataMapping(v map[string]string) {
	o.CsrDataMapping = v
}

// GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy {
	if o == nil || IsNil(o.MaxCertificatePerHolderPolicy.Get()) {
		var ret MaxCertificatePerHolderPolicy
		return ret
	}
	return *o.MaxCertificatePerHolderPolicy.Get()
}

// GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCertificatePerHolderPolicy.Get(), o.MaxCertificatePerHolderPolicy.IsSet()
}

// HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasMaxCertificatePerHolderPolicy() bool {
	if o != nil && o.MaxCertificatePerHolderPolicy.IsSet() {
		return true
	}

	return false
}

// SetMaxCertificatePerHolderPolicy gets a reference to the given NullableMaxCertificatePerHolderPolicy and assigns it to the MaxCertificatePerHolderPolicy field.
func (o *AcmeProfileResponse) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy) {
	o.MaxCertificatePerHolderPolicy.Set(&v)
}
// SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil
func (o *AcmeProfileResponse) SetMaxCertificatePerHolderPolicyNil() {
	o.MaxCertificatePerHolderPolicy.Set(nil)
}

// UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
func (o *AcmeProfileResponse) UnsetMaxCertificatePerHolderPolicy() {
	o.MaxCertificatePerHolderPolicy.Unset()
}

// GetMaxDnsName returns the MaxDnsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetMaxDnsName() int64 {
	if o == nil || IsNil(o.MaxDnsName.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxDnsName.Get()
}

// GetMaxDnsNameOk returns a tuple with the MaxDnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetMaxDnsNameOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxDnsName.Get(), o.MaxDnsName.IsSet()
}

// HasMaxDnsName returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasMaxDnsName() bool {
	if o != nil && o.MaxDnsName.IsSet() {
		return true
	}

	return false
}

// SetMaxDnsName gets a reference to the given NullableInt64 and assigns it to the MaxDnsName field.
func (o *AcmeProfileResponse) SetMaxDnsName(v int64) {
	o.MaxDnsName.Set(&v)
}
// SetMaxDnsNameNil sets the value for MaxDnsName to be an explicit nil
func (o *AcmeProfileResponse) SetMaxDnsNameNil() {
	o.MaxDnsName.Set(nil)
}

// UnsetMaxDnsName ensures that no value is present for MaxDnsName, not even an explicit nil
func (o *AcmeProfileResponse) UnsetMaxDnsName() {
	o.MaxDnsName.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetProxy() string {
	if o == nil || IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *AcmeProfileResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}
// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *AcmeProfileResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *AcmeProfileResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetAuthorizationLevels returns the AuthorizationLevels field value
func (o *AcmeProfileResponse) GetAuthorizationLevels() CertificateProfileAuthorizationLevels {
	if o == nil {
		var ret CertificateProfileAuthorizationLevels
		return ret
	}

	return o.AuthorizationLevels
}

// GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationLevels, true
}

// SetAuthorizationLevels sets field value
func (o *AcmeProfileResponse) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels) {
	o.AuthorizationLevels = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetTriggers() CertificateProfileTriggers {
	if o == nil || IsNil(o.Triggers.Get()) {
		var ret CertificateProfileTriggers
		return ret
	}
	return *o.Triggers.Get()
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetTriggersOk() (*CertificateProfileTriggers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers.Get(), o.Triggers.IsSet()
}

// HasTriggers returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasTriggers() bool {
	if o != nil && o.Triggers.IsSet() {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given NullableCertificateProfileTriggers and assigns it to the Triggers field.
func (o *AcmeProfileResponse) SetTriggers(v CertificateProfileTriggers) {
	o.Triggers.Set(&v)
}
// SetTriggersNil sets the value for Triggers to be an explicit nil
func (o *AcmeProfileResponse) SetTriggersNil() {
	o.Triggers.Set(nil)
}

// UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
func (o *AcmeProfileResponse) UnsetTriggers() {
	o.Triggers.Unset()
}

// GetRequestsPolicy returns the RequestsPolicy field value
func (o *AcmeProfileResponse) GetRequestsPolicy() RequestsPolicy {
	if o == nil {
		var ret RequestsPolicy
		return ret
	}

	return o.RequestsPolicy
}

// GetRequestsPolicyOk returns a tuple with the RequestsPolicy field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetRequestsPolicyOk() (*RequestsPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestsPolicy, true
}

// SetRequestsPolicy sets field value
func (o *AcmeProfileResponse) SetRequestsPolicy(v RequestsPolicy) {
	o.RequestsPolicy = v
}

// GetSelfPermissions returns the SelfPermissions field value
func (o *AcmeProfileResponse) GetSelfPermissions() CertificateProfileSelfPermissions {
	if o == nil {
		var ret CertificateProfileSelfPermissions
		return ret
	}

	return o.SelfPermissions
}

// GetSelfPermissionsOk returns a tuple with the SelfPermissions field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfPermissions, true
}

// SetSelfPermissions sets field value
func (o *AcmeProfileResponse) SetSelfPermissions(v CertificateProfileSelfPermissions) {
	o.SelfPermissions = v
}

// GetCertificateTemplate returns the CertificateTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetCertificateTemplate() CertificateTemplate {
	if o == nil || IsNil(o.CertificateTemplate.Get()) {
		var ret CertificateTemplate
		return ret
	}
	return *o.CertificateTemplate.Get()
}

// GetCertificateTemplateOk returns a tuple with the CertificateTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetCertificateTemplateOk() (*CertificateTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateTemplate.Get(), o.CertificateTemplate.IsSet()
}

// HasCertificateTemplate returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasCertificateTemplate() bool {
	if o != nil && o.CertificateTemplate.IsSet() {
		return true
	}

	return false
}

// SetCertificateTemplate gets a reference to the given NullableCertificateTemplate and assigns it to the CertificateTemplate field.
func (o *AcmeProfileResponse) SetCertificateTemplate(v CertificateTemplate) {
	o.CertificateTemplate.Set(&v)
}
// SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil
func (o *AcmeProfileResponse) SetCertificateTemplateNil() {
	o.CertificateTemplate.Set(nil)
}

// UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
func (o *AcmeProfileResponse) UnsetCertificateTemplate() {
	o.CertificateTemplate.Unset()
}

// GetCryptoPolicy returns the CryptoPolicy field value
func (o *AcmeProfileResponse) GetCryptoPolicy() CertificateProfileCryptoPolicy {
	if o == nil {
		var ret CertificateProfileCryptoPolicy
		return ret
	}

	return o.CryptoPolicy
}

// GetCryptoPolicyOk returns a tuple with the CryptoPolicy field value
// and a boolean to check if the value has been set.
func (o *AcmeProfileResponse) GetCryptoPolicyOk() (*CertificateProfileCryptoPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoPolicy, true
}

// SetCryptoPolicy sets field value
func (o *AcmeProfileResponse) SetCryptoPolicy(v CertificateProfileCryptoPolicy) {
	o.CryptoPolicy = v
}

// GetGradingPolicies returns the GradingPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetGradingPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GradingPolicies
}

// GetGradingPoliciesOk returns a tuple with the GradingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetGradingPoliciesOk() ([]string, bool) {
	if o == nil || IsNil(o.GradingPolicies) {
		return nil, false
	}
	return o.GradingPolicies, true
}

// HasGradingPolicies returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasGradingPolicies() bool {
	if o != nil && !IsNil(o.GradingPolicies) {
		return true
	}

	return false
}

// SetGradingPolicies gets a reference to the given []string and assigns it to the GradingPolicies field.
func (o *AcmeProfileResponse) SetGradingPolicies(v []string) {
	o.GradingPolicies = v
}

// GetDsFlow returns the DsFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfileResponse) GetDsFlow() []DataSourceFlowEntry {
	if o == nil {
		var ret []DataSourceFlowEntry
		return ret
	}
	return o.DsFlow
}

// GetDsFlowOk returns a tuple with the DsFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfileResponse) GetDsFlowOk() ([]DataSourceFlowEntry, bool) {
	if o == nil || IsNil(o.DsFlow) {
		return nil, false
	}
	return o.DsFlow, true
}

// HasDsFlow returns a boolean if a field has been set.
func (o *AcmeProfileResponse) HasDsFlow() bool {
	if o != nil && !IsNil(o.DsFlow) {
		return true
	}

	return false
}

// SetDsFlow gets a reference to the given []DataSourceFlowEntry and assigns it to the DsFlow field.
func (o *AcmeProfileResponse) SetDsFlow(v []DataSourceFlowEntry) {
	o.DsFlow = v
}

func (o AcmeProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcmeProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["module"] = o.Module
	toSerialize["name"] = o.Name
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["timeout"] = o.Timeout
	if o.Meta.IsSet() {
		toSerialize["meta"] = o.Meta.Get()
	}
	if o.Constraints.IsSet() {
		toSerialize["constraints"] = o.Constraints.Get()
	}
	if o.AuthorizationMethods != nil {
		toSerialize["authorizationMethods"] = o.AuthorizationMethods
	}
	toSerialize["pkiConnector"] = o.PkiConnector
	if o.Http01Port.IsSet() {
		toSerialize["http01Port"] = o.Http01Port.Get()
	}
	if o.TlsAlpn01Port.IsSet() {
		toSerialize["tlsAlpn01Port"] = o.TlsAlpn01Port.Get()
	}
	toSerialize["authorizeShortName"] = o.AuthorizeShortName
	toSerialize["authorizeEmptyContact"] = o.AuthorizeEmptyContact
	if o.DefaultContacts != nil {
		toSerialize["defaultContacts"] = o.DefaultContacts
	}
	toSerialize["verifyRetryCount"] = o.VerifyRetryCount
	toSerialize["verifyRetryDelay"] = o.VerifyRetryDelay
	toSerialize["requireTermsOfService"] = o.RequireTermsOfService
	if o.RenewalPeriod.IsSet() {
		toSerialize["renewalPeriod"] = o.RenewalPeriod.Get()
	}
	if o.CsrDataMapping != nil {
		toSerialize["csrDataMapping"] = o.CsrDataMapping
	}
	if o.MaxCertificatePerHolderPolicy.IsSet() {
		toSerialize["maxCertificatePerHolderPolicy"] = o.MaxCertificatePerHolderPolicy.Get()
	}
	if o.MaxDnsName.IsSet() {
		toSerialize["maxDnsName"] = o.MaxDnsName.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	toSerialize["authorizationLevels"] = o.AuthorizationLevels
	if o.Triggers.IsSet() {
		toSerialize["triggers"] = o.Triggers.Get()
	}
	toSerialize["requestsPolicy"] = o.RequestsPolicy
	toSerialize["selfPermissions"] = o.SelfPermissions
	if o.CertificateTemplate.IsSet() {
		toSerialize["certificateTemplate"] = o.CertificateTemplate.Get()
	}
	toSerialize["cryptoPolicy"] = o.CryptoPolicy
	if o.GradingPolicies != nil {
		toSerialize["gradingPolicies"] = o.GradingPolicies
	}
	if o.DsFlow != nil {
		toSerialize["dsFlow"] = o.DsFlow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcmeProfileResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"module",
		"name",
		"enabled",
		"timeout",
		"pkiConnector",
		"authorizeShortName",
		"authorizeEmptyContact",
		"verifyRetryCount",
		"verifyRetryDelay",
		"requireTermsOfService",
		"authorizationLevels",
		"requestsPolicy",
		"selfPermissions",
		"cryptoPolicy",
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

	varAcmeProfileResponse := _AcmeProfileResponse{}

	err = json.Unmarshal(data, &varAcmeProfileResponse)

	if err != nil {
		return err
	}

	*o = AcmeProfileResponse(varAcmeProfileResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "authorizationMethods")
		delete(additionalProperties, "pkiConnector")
		delete(additionalProperties, "http01Port")
		delete(additionalProperties, "tlsAlpn01Port")
		delete(additionalProperties, "authorizeShortName")
		delete(additionalProperties, "authorizeEmptyContact")
		delete(additionalProperties, "defaultContacts")
		delete(additionalProperties, "verifyRetryCount")
		delete(additionalProperties, "verifyRetryDelay")
		delete(additionalProperties, "requireTermsOfService")
		delete(additionalProperties, "renewalPeriod")
		delete(additionalProperties, "csrDataMapping")
		delete(additionalProperties, "maxCertificatePerHolderPolicy")
		delete(additionalProperties, "maxDnsName")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "authorizationLevels")
		delete(additionalProperties, "triggers")
		delete(additionalProperties, "requestsPolicy")
		delete(additionalProperties, "selfPermissions")
		delete(additionalProperties, "certificateTemplate")
		delete(additionalProperties, "cryptoPolicy")
		delete(additionalProperties, "gradingPolicies")
		delete(additionalProperties, "dsFlow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcmeProfileResponse struct {
	value *AcmeProfileResponse
	isSet bool
}

func (v NullableAcmeProfileResponse) Get() *AcmeProfileResponse {
	return v.value
}

func (v *NullableAcmeProfileResponse) Set(val *AcmeProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcmeProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcmeProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcmeProfileResponse(val *AcmeProfileResponse) *NullableAcmeProfileResponse {
	return &NullableAcmeProfileResponse{value: val, isSet: true}
}

func (v NullableAcmeProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcmeProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


