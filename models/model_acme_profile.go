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

// checks if the AcmeProfile type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AcmeProfile{}

// AcmeProfile struct for AcmeProfile
type AcmeProfile struct {
	AuthorizationLevels   CertificateProfileAuthorizationLevels `json:"authorizationLevels"`
	AuthorizationMethods  []string                              `json:"authorizationMethods,omitempty"`
	AuthorizeEmptyContact bool                                  `json:"authorizeEmptyContact"`
	AuthorizeShortName    bool                                  `json:"authorizeShortName"`
	CertificateTemplate   NullableCertificateTemplate           `json:"certificateTemplate,omitempty"`
	Constraints           NullableCertificateRequestConstraints `json:"constraints,omitempty"`
	CryptoPolicy          ManagedCertificateProfileCryptoPolicy `json:"cryptoPolicy"`
	CsrDataMapping        map[string]string                     `json:"csrDataMapping,omitempty"`
	DefaultContacts       []string                              `json:"defaultContacts,omitempty"`
	Description           []LocalizedString                     `json:"description,omitempty"`
	DisplayName           []LocalizedString                     `json:"displayName,omitempty"`
	// Representation of a datasource execution flow
	DsFlow                        []DataSourceFlowEntry                 `json:"dsFlow,omitempty"`
	Enabled                       bool                                  `json:"enabled"`
	GradingPolicies               []string                              `json:"gradingPolicies,omitempty"`
	Http01Port                    utils.NullableInt64                   `json:"http01Port,omitempty"`
	MaxCertificatePerHolderPolicy NullableMaxCertificatePerHolderPolicy `json:"maxCertificatePerHolderPolicy,omitempty"`
	MaxDnsName                    utils.NullableInt64                   `json:"maxDnsName,omitempty"`
	Meta                          NullableDirectoryMeta                 `json:"meta,omitempty"`
	Module                        string                                `json:"module"`
	Name                          string                                `json:"name"`
	PkiConnector                  string                                `json:"pkiConnector"`
	Proxy                         utils.NullableString                  `json:"proxy,omitempty"`
	RenewalPeriod                 utils.NullableString                  `json:"renewalPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	RequestsPolicy                RequestsPolicy                        `json:"requestsPolicy"`
	RequireTermsOfService         bool                                  `json:"requireTermsOfService"`
	SelfPermissions               CertificateProfileSelfPermissions     `json:"selfPermissions"`
	ThirdPartyDiscoverySync       utils.NullableBool                    `json:"thirdPartyDiscoverySync,omitempty"`
	Timeout                       string                                `json:"timeout" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	TlsAlpn01Port                 utils.NullableInt64                   `json:"tlsAlpn01Port,omitempty"`
	Triggers                      NullableCertificateProfileTriggers    `json:"triggers,omitempty"`
	VerifyRetryCount              int64                                 `json:"verifyRetryCount"`
	VerifyRetryDelay              string                                `json:"verifyRetryDelay" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	AdditionalProperties          map[string]interface{}
}

type _AcmeProfile AcmeProfile

// NewAcmeProfile instantiates a new AcmeProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcmeProfile(authorizationLevels CertificateProfileAuthorizationLevels, authorizeEmptyContact bool, authorizeShortName bool, cryptoPolicy ManagedCertificateProfileCryptoPolicy, enabled bool, module string, name string, pkiConnector string, requestsPolicy RequestsPolicy, requireTermsOfService bool, selfPermissions CertificateProfileSelfPermissions, timeout string, verifyRetryCount int64, verifyRetryDelay string) *AcmeProfile {
	this := AcmeProfile{}
	this.AuthorizationLevels = authorizationLevels
	this.AuthorizeEmptyContact = authorizeEmptyContact
	this.AuthorizeShortName = authorizeShortName
	this.CryptoPolicy = cryptoPolicy
	this.Enabled = enabled
	this.Module = module
	this.Name = name
	this.PkiConnector = pkiConnector
	this.RequestsPolicy = requestsPolicy
	this.RequireTermsOfService = requireTermsOfService
	this.SelfPermissions = selfPermissions
	this.Timeout = timeout
	this.VerifyRetryCount = verifyRetryCount
	this.VerifyRetryDelay = verifyRetryDelay
	var thirdPartyDiscoverySync bool = false
	this.ThirdPartyDiscoverySync = *utils.NewNullableBool(&thirdPartyDiscoverySync)
	return &this
}

// NewAcmeProfileWithDefaults instantiates a new AcmeProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcmeProfileWithDefaults() *AcmeProfile {
	this := AcmeProfile{}
	var thirdPartyDiscoverySync bool = false
	this.ThirdPartyDiscoverySync = *utils.NewNullableBool(&thirdPartyDiscoverySync)
	return &this
}

// GetAuthorizationLevels returns the AuthorizationLevels field value
func (o *AcmeProfile) GetAuthorizationLevels() CertificateProfileAuthorizationLevels {
	if o == nil {
		var ret CertificateProfileAuthorizationLevels
		return ret
	}

	return o.AuthorizationLevels
}

// GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationLevels, true
}

// SetAuthorizationLevels sets field value
func (o *AcmeProfile) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels) {
	o.AuthorizationLevels = v
}

// GetAuthorizationMethods returns the AuthorizationMethods field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetAuthorizationMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AuthorizationMethods
}

// GetAuthorizationMethodsOk returns a tuple with the AuthorizationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetAuthorizationMethodsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AuthorizationMethods) {
		return nil, false
	}
	return o.AuthorizationMethods, true
}

// HasAuthorizationMethods returns a boolean if a field has been set.
func (o *AcmeProfile) HasAuthorizationMethods() bool {
	if o != nil && !utils.IsNil(o.AuthorizationMethods) {
		return true
	}

	return false
}

// SetAuthorizationMethods gets a reference to the given []string and assigns it to the AuthorizationMethods field.
func (o *AcmeProfile) SetAuthorizationMethods(v []string) {
	o.AuthorizationMethods = v
}

// GetAuthorizeEmptyContact returns the AuthorizeEmptyContact field value
func (o *AcmeProfile) GetAuthorizeEmptyContact() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AuthorizeEmptyContact
}

// GetAuthorizeEmptyContactOk returns a tuple with the AuthorizeEmptyContact field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetAuthorizeEmptyContactOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizeEmptyContact, true
}

// SetAuthorizeEmptyContact sets field value
func (o *AcmeProfile) SetAuthorizeEmptyContact(v bool) {
	o.AuthorizeEmptyContact = v
}

// GetAuthorizeShortName returns the AuthorizeShortName field value
func (o *AcmeProfile) GetAuthorizeShortName() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AuthorizeShortName
}

// GetAuthorizeShortNameOk returns a tuple with the AuthorizeShortName field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetAuthorizeShortNameOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizeShortName, true
}

// SetAuthorizeShortName sets field value
func (o *AcmeProfile) SetAuthorizeShortName(v bool) {
	o.AuthorizeShortName = v
}

// GetCertificateTemplate returns the CertificateTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetCertificateTemplate() CertificateTemplate {
	if o == nil || utils.IsNil(o.CertificateTemplate.Get()) {
		var ret CertificateTemplate
		return ret
	}
	return *o.CertificateTemplate.Get()
}

// GetCertificateTemplateOk returns a tuple with the CertificateTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetCertificateTemplateOk() (*CertificateTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateTemplate.Get(), o.CertificateTemplate.IsSet()
}

// HasCertificateTemplate returns a boolean if a field has been set.
func (o *AcmeProfile) HasCertificateTemplate() bool {
	if o != nil && o.CertificateTemplate.IsSet() {
		return true
	}

	return false
}

// SetCertificateTemplate gets a reference to the given NullableCertificateTemplate and assigns it to the CertificateTemplate field.
func (o *AcmeProfile) SetCertificateTemplate(v CertificateTemplate) {
	o.CertificateTemplate.Set(&v)
}

// SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil
func (o *AcmeProfile) SetCertificateTemplateNil() {
	o.CertificateTemplate.Set(nil)
}

// UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
func (o *AcmeProfile) UnsetCertificateTemplate() {
	o.CertificateTemplate.Unset()
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetConstraints() CertificateRequestConstraints {
	if o == nil || utils.IsNil(o.Constraints.Get()) {
		var ret CertificateRequestConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetConstraintsOk() (*CertificateRequestConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *AcmeProfile) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableCertificateRequestConstraints and assigns it to the Constraints field.
func (o *AcmeProfile) SetConstraints(v CertificateRequestConstraints) {
	o.Constraints.Set(&v)
}

// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *AcmeProfile) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *AcmeProfile) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetCryptoPolicy returns the CryptoPolicy field value
func (o *AcmeProfile) GetCryptoPolicy() ManagedCertificateProfileCryptoPolicy {
	if o == nil {
		var ret ManagedCertificateProfileCryptoPolicy
		return ret
	}

	return o.CryptoPolicy
}

// GetCryptoPolicyOk returns a tuple with the CryptoPolicy field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetCryptoPolicyOk() (*ManagedCertificateProfileCryptoPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoPolicy, true
}

// SetCryptoPolicy sets field value
func (o *AcmeProfile) SetCryptoPolicy(v ManagedCertificateProfileCryptoPolicy) {
	o.CryptoPolicy = v
}

// GetCsrDataMapping returns the CsrDataMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetCsrDataMapping() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.CsrDataMapping
}

// GetCsrDataMappingOk returns a tuple with the CsrDataMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetCsrDataMappingOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.CsrDataMapping) {
		return nil, false
	}
	return &o.CsrDataMapping, true
}

// HasCsrDataMapping returns a boolean if a field has been set.
func (o *AcmeProfile) HasCsrDataMapping() bool {
	if o != nil && !utils.IsNil(o.CsrDataMapping) {
		return true
	}

	return false
}

// SetCsrDataMapping gets a reference to the given map[string]string and assigns it to the CsrDataMapping field.
func (o *AcmeProfile) SetCsrDataMapping(v map[string]string) {
	o.CsrDataMapping = v
}

// GetDefaultContacts returns the DefaultContacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetDefaultContacts() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DefaultContacts
}

// GetDefaultContactsOk returns a tuple with the DefaultContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetDefaultContactsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.DefaultContacts) {
		return nil, false
	}
	return o.DefaultContacts, true
}

// HasDefaultContacts returns a boolean if a field has been set.
func (o *AcmeProfile) HasDefaultContacts() bool {
	if o != nil && !utils.IsNil(o.DefaultContacts) {
		return true
	}

	return false
}

// SetDefaultContacts gets a reference to the given []string and assigns it to the DefaultContacts field.
func (o *AcmeProfile) SetDefaultContacts(v []string) {
	o.DefaultContacts = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AcmeProfile) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *AcmeProfile) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AcmeProfile) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *AcmeProfile) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetDsFlow returns the DsFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetDsFlow() []DataSourceFlowEntry {
	if o == nil {
		var ret []DataSourceFlowEntry
		return ret
	}
	return o.DsFlow
}

// GetDsFlowOk returns a tuple with the DsFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetDsFlowOk() ([]DataSourceFlowEntry, bool) {
	if o == nil || utils.IsNil(o.DsFlow) {
		return nil, false
	}
	return o.DsFlow, true
}

// HasDsFlow returns a boolean if a field has been set.
func (o *AcmeProfile) HasDsFlow() bool {
	if o != nil && !utils.IsNil(o.DsFlow) {
		return true
	}

	return false
}

// SetDsFlow gets a reference to the given []DataSourceFlowEntry and assigns it to the DsFlow field.
func (o *AcmeProfile) SetDsFlow(v []DataSourceFlowEntry) {
	o.DsFlow = v
}

// GetEnabled returns the Enabled field value
func (o *AcmeProfile) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AcmeProfile) SetEnabled(v bool) {
	o.Enabled = v
}

// GetGradingPolicies returns the GradingPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetGradingPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GradingPolicies
}

// GetGradingPoliciesOk returns a tuple with the GradingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetGradingPoliciesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.GradingPolicies) {
		return nil, false
	}
	return o.GradingPolicies, true
}

// HasGradingPolicies returns a boolean if a field has been set.
func (o *AcmeProfile) HasGradingPolicies() bool {
	if o != nil && !utils.IsNil(o.GradingPolicies) {
		return true
	}

	return false
}

// SetGradingPolicies gets a reference to the given []string and assigns it to the GradingPolicies field.
func (o *AcmeProfile) SetGradingPolicies(v []string) {
	o.GradingPolicies = v
}

// GetHttp01Port returns the Http01Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetHttp01Port() int64 {
	if o == nil || utils.IsNil(o.Http01Port.Get()) {
		var ret int64
		return ret
	}
	return *o.Http01Port.Get()
}

// GetHttp01PortOk returns a tuple with the Http01Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetHttp01PortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Http01Port.Get(), o.Http01Port.IsSet()
}

// HasHttp01Port returns a boolean if a field has been set.
func (o *AcmeProfile) HasHttp01Port() bool {
	if o != nil && o.Http01Port.IsSet() {
		return true
	}

	return false
}

// SetHttp01Port gets a reference to the given NullableInt64 and assigns it to the Http01Port field.
func (o *AcmeProfile) SetHttp01Port(v int64) {
	o.Http01Port.Set(&v)
}

// SetHttp01PortNil sets the value for Http01Port to be an explicit nil
func (o *AcmeProfile) SetHttp01PortNil() {
	o.Http01Port.Set(nil)
}

// UnsetHttp01Port ensures that no value is present for Http01Port, not even an explicit nil
func (o *AcmeProfile) UnsetHttp01Port() {
	o.Http01Port.Unset()
}

// GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy {
	if o == nil || utils.IsNil(o.MaxCertificatePerHolderPolicy.Get()) {
		var ret MaxCertificatePerHolderPolicy
		return ret
	}
	return *o.MaxCertificatePerHolderPolicy.Get()
}

// GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCertificatePerHolderPolicy.Get(), o.MaxCertificatePerHolderPolicy.IsSet()
}

// HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.
func (o *AcmeProfile) HasMaxCertificatePerHolderPolicy() bool {
	if o != nil && o.MaxCertificatePerHolderPolicy.IsSet() {
		return true
	}

	return false
}

// SetMaxCertificatePerHolderPolicy gets a reference to the given NullableMaxCertificatePerHolderPolicy and assigns it to the MaxCertificatePerHolderPolicy field.
func (o *AcmeProfile) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy) {
	o.MaxCertificatePerHolderPolicy.Set(&v)
}

// SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil
func (o *AcmeProfile) SetMaxCertificatePerHolderPolicyNil() {
	o.MaxCertificatePerHolderPolicy.Set(nil)
}

// UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
func (o *AcmeProfile) UnsetMaxCertificatePerHolderPolicy() {
	o.MaxCertificatePerHolderPolicy.Unset()
}

// GetMaxDnsName returns the MaxDnsName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetMaxDnsName() int64 {
	if o == nil || utils.IsNil(o.MaxDnsName.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxDnsName.Get()
}

// GetMaxDnsNameOk returns a tuple with the MaxDnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetMaxDnsNameOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxDnsName.Get(), o.MaxDnsName.IsSet()
}

// HasMaxDnsName returns a boolean if a field has been set.
func (o *AcmeProfile) HasMaxDnsName() bool {
	if o != nil && o.MaxDnsName.IsSet() {
		return true
	}

	return false
}

// SetMaxDnsName gets a reference to the given NullableInt64 and assigns it to the MaxDnsName field.
func (o *AcmeProfile) SetMaxDnsName(v int64) {
	o.MaxDnsName.Set(&v)
}

// SetMaxDnsNameNil sets the value for MaxDnsName to be an explicit nil
func (o *AcmeProfile) SetMaxDnsNameNil() {
	o.MaxDnsName.Set(nil)
}

// UnsetMaxDnsName ensures that no value is present for MaxDnsName, not even an explicit nil
func (o *AcmeProfile) UnsetMaxDnsName() {
	o.MaxDnsName.Unset()
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetMeta() DirectoryMeta {
	if o == nil || utils.IsNil(o.Meta.Get()) {
		var ret DirectoryMeta
		return ret
	}
	return *o.Meta.Get()
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetMetaOk() (*DirectoryMeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Meta.Get(), o.Meta.IsSet()
}

// HasMeta returns a boolean if a field has been set.
func (o *AcmeProfile) HasMeta() bool {
	if o != nil && o.Meta.IsSet() {
		return true
	}

	return false
}

// SetMeta gets a reference to the given NullableDirectoryMeta and assigns it to the Meta field.
func (o *AcmeProfile) SetMeta(v DirectoryMeta) {
	o.Meta.Set(&v)
}

// SetMetaNil sets the value for Meta to be an explicit nil
func (o *AcmeProfile) SetMetaNil() {
	o.Meta.Set(nil)
}

// UnsetMeta ensures that no value is present for Meta, not even an explicit nil
func (o *AcmeProfile) UnsetMeta() {
	o.Meta.Unset()
}

// GetModule returns the Module field value
func (o *AcmeProfile) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *AcmeProfile) SetModule(v string) {
	o.Module = v
}

// GetName returns the Name field value
func (o *AcmeProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AcmeProfile) SetName(v string) {
	o.Name = v
}

// GetPkiConnector returns the PkiConnector field value
func (o *AcmeProfile) GetPkiConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PkiConnector
}

// GetPkiConnectorOk returns a tuple with the PkiConnector field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetPkiConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiConnector, true
}

// SetPkiConnector sets field value
func (o *AcmeProfile) SetPkiConnector(v string) {
	o.PkiConnector = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *AcmeProfile) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *AcmeProfile) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *AcmeProfile) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *AcmeProfile) UnsetProxy() {
	o.Proxy.Unset()
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetRenewalPeriod() string {
	if o == nil || utils.IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *AcmeProfile) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *AcmeProfile) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}

// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *AcmeProfile) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *AcmeProfile) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetRequestsPolicy returns the RequestsPolicy field value
func (o *AcmeProfile) GetRequestsPolicy() RequestsPolicy {
	if o == nil {
		var ret RequestsPolicy
		return ret
	}

	return o.RequestsPolicy
}

// GetRequestsPolicyOk returns a tuple with the RequestsPolicy field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetRequestsPolicyOk() (*RequestsPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestsPolicy, true
}

// SetRequestsPolicy sets field value
func (o *AcmeProfile) SetRequestsPolicy(v RequestsPolicy) {
	o.RequestsPolicy = v
}

// GetRequireTermsOfService returns the RequireTermsOfService field value
func (o *AcmeProfile) GetRequireTermsOfService() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireTermsOfService
}

// GetRequireTermsOfServiceOk returns a tuple with the RequireTermsOfService field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetRequireTermsOfServiceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireTermsOfService, true
}

// SetRequireTermsOfService sets field value
func (o *AcmeProfile) SetRequireTermsOfService(v bool) {
	o.RequireTermsOfService = v
}

// GetSelfPermissions returns the SelfPermissions field value
func (o *AcmeProfile) GetSelfPermissions() CertificateProfileSelfPermissions {
	if o == nil {
		var ret CertificateProfileSelfPermissions
		return ret
	}

	return o.SelfPermissions
}

// GetSelfPermissionsOk returns a tuple with the SelfPermissions field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfPermissions, true
}

// SetSelfPermissions sets field value
func (o *AcmeProfile) SetSelfPermissions(v CertificateProfileSelfPermissions) {
	o.SelfPermissions = v
}

// GetThirdPartyDiscoverySync returns the ThirdPartyDiscoverySync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetThirdPartyDiscoverySync() bool {
	if o == nil || utils.IsNil(o.ThirdPartyDiscoverySync.Get()) {
		var ret bool
		return ret
	}
	return *o.ThirdPartyDiscoverySync.Get()
}

// GetThirdPartyDiscoverySyncOk returns a tuple with the ThirdPartyDiscoverySync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetThirdPartyDiscoverySyncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThirdPartyDiscoverySync.Get(), o.ThirdPartyDiscoverySync.IsSet()
}

// HasThirdPartyDiscoverySync returns a boolean if a field has been set.
func (o *AcmeProfile) HasThirdPartyDiscoverySync() bool {
	if o != nil && o.ThirdPartyDiscoverySync.IsSet() {
		return true
	}

	return false
}

// SetThirdPartyDiscoverySync gets a reference to the given NullableBool and assigns it to the ThirdPartyDiscoverySync field.
func (o *AcmeProfile) SetThirdPartyDiscoverySync(v bool) {
	o.ThirdPartyDiscoverySync.Set(&v)
}

// SetThirdPartyDiscoverySyncNil sets the value for ThirdPartyDiscoverySync to be an explicit nil
func (o *AcmeProfile) SetThirdPartyDiscoverySyncNil() {
	o.ThirdPartyDiscoverySync.Set(nil)
}

// UnsetThirdPartyDiscoverySync ensures that no value is present for ThirdPartyDiscoverySync, not even an explicit nil
func (o *AcmeProfile) UnsetThirdPartyDiscoverySync() {
	o.ThirdPartyDiscoverySync.Unset()
}

// GetTimeout returns the Timeout field value
func (o *AcmeProfile) GetTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *AcmeProfile) SetTimeout(v string) {
	o.Timeout = v
}

// GetTlsAlpn01Port returns the TlsAlpn01Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetTlsAlpn01Port() int64 {
	if o == nil || utils.IsNil(o.TlsAlpn01Port.Get()) {
		var ret int64
		return ret
	}
	return *o.TlsAlpn01Port.Get()
}

// GetTlsAlpn01PortOk returns a tuple with the TlsAlpn01Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetTlsAlpn01PortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsAlpn01Port.Get(), o.TlsAlpn01Port.IsSet()
}

// HasTlsAlpn01Port returns a boolean if a field has been set.
func (o *AcmeProfile) HasTlsAlpn01Port() bool {
	if o != nil && o.TlsAlpn01Port.IsSet() {
		return true
	}

	return false
}

// SetTlsAlpn01Port gets a reference to the given NullableInt64 and assigns it to the TlsAlpn01Port field.
func (o *AcmeProfile) SetTlsAlpn01Port(v int64) {
	o.TlsAlpn01Port.Set(&v)
}

// SetTlsAlpn01PortNil sets the value for TlsAlpn01Port to be an explicit nil
func (o *AcmeProfile) SetTlsAlpn01PortNil() {
	o.TlsAlpn01Port.Set(nil)
}

// UnsetTlsAlpn01Port ensures that no value is present for TlsAlpn01Port, not even an explicit nil
func (o *AcmeProfile) UnsetTlsAlpn01Port() {
	o.TlsAlpn01Port.Unset()
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeProfile) GetTriggers() CertificateProfileTriggers {
	if o == nil || utils.IsNil(o.Triggers.Get()) {
		var ret CertificateProfileTriggers
		return ret
	}
	return *o.Triggers.Get()
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeProfile) GetTriggersOk() (*CertificateProfileTriggers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers.Get(), o.Triggers.IsSet()
}

// HasTriggers returns a boolean if a field has been set.
func (o *AcmeProfile) HasTriggers() bool {
	if o != nil && o.Triggers.IsSet() {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given NullableCertificateProfileTriggers and assigns it to the Triggers field.
func (o *AcmeProfile) SetTriggers(v CertificateProfileTriggers) {
	o.Triggers.Set(&v)
}

// SetTriggersNil sets the value for Triggers to be an explicit nil
func (o *AcmeProfile) SetTriggersNil() {
	o.Triggers.Set(nil)
}

// UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
func (o *AcmeProfile) UnsetTriggers() {
	o.Triggers.Unset()
}

// GetVerifyRetryCount returns the VerifyRetryCount field value
func (o *AcmeProfile) GetVerifyRetryCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.VerifyRetryCount
}

// GetVerifyRetryCountOk returns a tuple with the VerifyRetryCount field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetVerifyRetryCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyRetryCount, true
}

// SetVerifyRetryCount sets field value
func (o *AcmeProfile) SetVerifyRetryCount(v int64) {
	o.VerifyRetryCount = v
}

// GetVerifyRetryDelay returns the VerifyRetryDelay field value
func (o *AcmeProfile) GetVerifyRetryDelay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerifyRetryDelay
}

// GetVerifyRetryDelayOk returns a tuple with the VerifyRetryDelay field value
// and a boolean to check if the value has been set.
func (o *AcmeProfile) GetVerifyRetryDelayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyRetryDelay, true
}

// SetVerifyRetryDelay sets field value
func (o *AcmeProfile) SetVerifyRetryDelay(v string) {
	o.VerifyRetryDelay = v
}

func (o AcmeProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcmeProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorizationLevels"] = o.AuthorizationLevels
	if o.AuthorizationMethods != nil {
		toSerialize["authorizationMethods"] = o.AuthorizationMethods
	}
	toSerialize["authorizeEmptyContact"] = o.AuthorizeEmptyContact
	toSerialize["authorizeShortName"] = o.AuthorizeShortName
	if o.CertificateTemplate.IsSet() {
		toSerialize["certificateTemplate"] = o.CertificateTemplate.Get()
	}
	if o.Constraints.IsSet() {
		toSerialize["constraints"] = o.Constraints.Get()
	}
	toSerialize["cryptoPolicy"] = o.CryptoPolicy
	if o.CsrDataMapping != nil {
		toSerialize["csrDataMapping"] = o.CsrDataMapping
	}
	if o.DefaultContacts != nil {
		toSerialize["defaultContacts"] = o.DefaultContacts
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.DsFlow != nil {
		toSerialize["dsFlow"] = o.DsFlow
	}
	toSerialize["enabled"] = o.Enabled
	if o.GradingPolicies != nil {
		toSerialize["gradingPolicies"] = o.GradingPolicies
	}
	if o.Http01Port.IsSet() {
		toSerialize["http01Port"] = o.Http01Port.Get()
	}
	if o.MaxCertificatePerHolderPolicy.IsSet() {
		toSerialize["maxCertificatePerHolderPolicy"] = o.MaxCertificatePerHolderPolicy.Get()
	}
	if o.MaxDnsName.IsSet() {
		toSerialize["maxDnsName"] = o.MaxDnsName.Get()
	}
	if o.Meta.IsSet() {
		toSerialize["meta"] = o.Meta.Get()
	}
	toSerialize["module"] = o.Module
	toSerialize["name"] = o.Name
	toSerialize["pkiConnector"] = o.PkiConnector
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.RenewalPeriod.IsSet() {
		toSerialize["renewalPeriod"] = o.RenewalPeriod.Get()
	}
	toSerialize["requestsPolicy"] = o.RequestsPolicy
	toSerialize["requireTermsOfService"] = o.RequireTermsOfService
	toSerialize["selfPermissions"] = o.SelfPermissions
	if o.ThirdPartyDiscoverySync.IsSet() {
		toSerialize["thirdPartyDiscoverySync"] = o.ThirdPartyDiscoverySync.Get()
	}
	toSerialize["timeout"] = o.Timeout
	if o.TlsAlpn01Port.IsSet() {
		toSerialize["tlsAlpn01Port"] = o.TlsAlpn01Port.Get()
	}
	if o.Triggers.IsSet() {
		toSerialize["triggers"] = o.Triggers.Get()
	}
	toSerialize["verifyRetryCount"] = o.VerifyRetryCount
	toSerialize["verifyRetryDelay"] = o.VerifyRetryDelay

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcmeProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorizationLevels",
		"authorizeEmptyContact",
		"authorizeShortName",
		"cryptoPolicy",
		"enabled",
		"module",
		"name",
		"pkiConnector",
		"requestsPolicy",
		"requireTermsOfService",
		"selfPermissions",
		"timeout",
		"verifyRetryCount",
		"verifyRetryDelay",
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

	varAcmeProfile := _AcmeProfile{}

	err = json.Unmarshal(data, &varAcmeProfile)

	if err != nil {
		return err
	}

	*o = AcmeProfile(varAcmeProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorizationLevels")
		delete(additionalProperties, "authorizationMethods")
		delete(additionalProperties, "authorizeEmptyContact")
		delete(additionalProperties, "authorizeShortName")
		delete(additionalProperties, "certificateTemplate")
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "cryptoPolicy")
		delete(additionalProperties, "csrDataMapping")
		delete(additionalProperties, "defaultContacts")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "dsFlow")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "gradingPolicies")
		delete(additionalProperties, "http01Port")
		delete(additionalProperties, "maxCertificatePerHolderPolicy")
		delete(additionalProperties, "maxDnsName")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "pkiConnector")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "renewalPeriod")
		delete(additionalProperties, "requestsPolicy")
		delete(additionalProperties, "requireTermsOfService")
		delete(additionalProperties, "selfPermissions")
		delete(additionalProperties, "thirdPartyDiscoverySync")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "tlsAlpn01Port")
		delete(additionalProperties, "triggers")
		delete(additionalProperties, "verifyRetryCount")
		delete(additionalProperties, "verifyRetryDelay")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcmeProfile struct {
	value *AcmeProfile
	isSet bool
}

func (v NullableAcmeProfile) Get() *AcmeProfile {
	return v.value
}

func (v *NullableAcmeProfile) Set(val *AcmeProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAcmeProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAcmeProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcmeProfile(val *AcmeProfile) *NullableAcmeProfile {
	return &NullableAcmeProfile{value: val, isSet: true}
}

func (v NullableAcmeProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcmeProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
