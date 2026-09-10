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

// checks if the EstProfile type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EstProfile{}

// EstProfile struct for EstProfile
type EstProfile struct {
	AuthorizationLevels CertificateProfileAuthorizationLevels `json:"authorizationLevels"`
	AuthorizationMode   string                                `json:"authorizationMode"`
	Ca                  string                                `json:"ca"`
	CertificateTemplate NullableCertificateTemplate           `json:"certificateTemplate,omitempty"`
	Constraints         NullableCertificateRequestConstraints `json:"constraints,omitempty"`
	CryptoPolicy        ManagedCertificateProfileCryptoPolicy `json:"cryptoPolicy"`
	CsrDataMapping      map[string]string                     `json:"csrDataMapping,omitempty"`
	Description         []LocalizedString                     `json:"description,omitempty"`
	DisplayName         []LocalizedString                     `json:"displayName,omitempty"`
	DnWhitelist         bool                                  `json:"dnWhitelist"`
	// Representation of a datasource execution flow
	DsFlow                        []DataSourceFlowEntry                 `json:"dsFlow,omitempty"`
	Enabled                       bool                                  `json:"enabled"`
	EnrollAuthorizedCas           []string                              `json:"enrollAuthorizedCas,omitempty"`
	GradingPolicies               []string                              `json:"gradingPolicies,omitempty"`
	MaxCertificatePerHolderPolicy NullableMaxCertificatePerHolderPolicy `json:"maxCertificatePerHolderPolicy,omitempty"`
	Module                        string                                `json:"module"`
	Name                          string                                `json:"name"`
	PasswordPolicy                utils.NullableString                  `json:"passwordPolicy,omitempty"`
	PkiConnector                  string                                `json:"pkiConnector"`
	RenewalAuthorizedCas          []string                              `json:"renewalAuthorizedCas,omitempty"`
	RenewalPeriod                 utils.NullableString                  `json:"renewalPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	RequestsPolicy                RequestsPolicy                        `json:"requestsPolicy"`
	SelfPermissions               CertificateProfileSelfPermissions     `json:"selfPermissions"`
	// Reference to a `Terms of service` object. If defined, it will be displayed on the enrollment workflow before starting certificate enrollment
	TermsOfService          *string                            `json:"termsOfService,omitempty"`
	ThirdPartyDiscoverySync utils.NullableBool                 `json:"thirdPartyDiscoverySync,omitempty"`
	Triggers                NullableCertificateProfileTriggers `json:"triggers,omitempty"`
	ValidationRuleset       NullableValidationRuleset          `json:"validationRuleset,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _EstProfile EstProfile

// NewEstProfile instantiates a new EstProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstProfile(authorizationLevels CertificateProfileAuthorizationLevels, authorizationMode string, ca string, cryptoPolicy ManagedCertificateProfileCryptoPolicy, dnWhitelist bool, enabled bool, module string, name string, pkiConnector string, requestsPolicy RequestsPolicy, selfPermissions CertificateProfileSelfPermissions) *EstProfile {
	this := EstProfile{}
	this.AuthorizationLevels = authorizationLevels
	this.AuthorizationMode = authorizationMode
	this.Ca = ca
	this.CryptoPolicy = cryptoPolicy
	this.DnWhitelist = dnWhitelist
	this.Enabled = enabled
	this.Module = module
	this.Name = name
	this.PkiConnector = pkiConnector
	this.RequestsPolicy = requestsPolicy
	this.SelfPermissions = selfPermissions
	var thirdPartyDiscoverySync bool = false
	this.ThirdPartyDiscoverySync = *utils.NewNullableBool(&thirdPartyDiscoverySync)
	return &this
}

// NewEstProfileWithDefaults instantiates a new EstProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstProfileWithDefaults() *EstProfile {
	this := EstProfile{}
	var thirdPartyDiscoverySync bool = false
	this.ThirdPartyDiscoverySync = *utils.NewNullableBool(&thirdPartyDiscoverySync)
	return &this
}

// GetAuthorizationLevels returns the AuthorizationLevels field value
func (o *EstProfile) GetAuthorizationLevels() CertificateProfileAuthorizationLevels {
	if o == nil {
		var ret CertificateProfileAuthorizationLevels
		return ret
	}

	return o.AuthorizationLevels
}

// GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field value
// and a boolean to check if the value has been set.
func (o *EstProfile) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationLevels, true
}

// SetAuthorizationLevels sets field value
func (o *EstProfile) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels) {
	o.AuthorizationLevels = v
}

// GetAuthorizationMode returns the AuthorizationMode field value
func (o *EstProfile) GetAuthorizationMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationMode
}

// GetAuthorizationModeOk returns a tuple with the AuthorizationMode field value
// and a boolean to check if the value has been set.
func (o *EstProfile) GetAuthorizationModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationMode, true
}

// SetAuthorizationMode sets field value
func (o *EstProfile) SetAuthorizationMode(v string) {
	o.AuthorizationMode = v
}

// GetCa returns the Ca field value
func (o *EstProfile) GetCa() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ca
}

// GetCaOk returns a tuple with the Ca field value
// and a boolean to check if the value has been set.
func (o *EstProfile) GetCaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ca, true
}

// SetCa sets field value
func (o *EstProfile) SetCa(v string) {
	o.Ca = v
}

// GetCertificateTemplate returns the CertificateTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetCertificateTemplate() CertificateTemplate {
	if o == nil || utils.IsNil(o.CertificateTemplate.Get()) {
		var ret CertificateTemplate
		return ret
	}
	return *o.CertificateTemplate.Get()
}

// GetCertificateTemplateOk returns a tuple with the CertificateTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetCertificateTemplateOk() (*CertificateTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateTemplate.Get(), o.CertificateTemplate.IsSet()
}

// HasCertificateTemplate returns a boolean if a field has been set.
func (o *EstProfile) HasCertificateTemplate() bool {
	if o != nil && o.CertificateTemplate.IsSet() {
		return true
	}

	return false
}

// SetCertificateTemplate gets a reference to the given NullableCertificateTemplate and assigns it to the CertificateTemplate field.
func (o *EstProfile) SetCertificateTemplate(v CertificateTemplate) {
	o.CertificateTemplate.Set(&v)
}

// SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil
func (o *EstProfile) SetCertificateTemplateNil() {
	o.CertificateTemplate.Set(nil)
}

// UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
func (o *EstProfile) UnsetCertificateTemplate() {
	o.CertificateTemplate.Unset()
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetConstraints() CertificateRequestConstraints {
	if o == nil || utils.IsNil(o.Constraints.Get()) {
		var ret CertificateRequestConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetConstraintsOk() (*CertificateRequestConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *EstProfile) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableCertificateRequestConstraints and assigns it to the Constraints field.
func (o *EstProfile) SetConstraints(v CertificateRequestConstraints) {
	o.Constraints.Set(&v)
}

// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *EstProfile) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *EstProfile) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetCryptoPolicy returns the CryptoPolicy field value
func (o *EstProfile) GetCryptoPolicy() ManagedCertificateProfileCryptoPolicy {
	if o == nil {
		var ret ManagedCertificateProfileCryptoPolicy
		return ret
	}

	return o.CryptoPolicy
}

// GetCryptoPolicyOk returns a tuple with the CryptoPolicy field value
// and a boolean to check if the value has been set.
func (o *EstProfile) GetCryptoPolicyOk() (*ManagedCertificateProfileCryptoPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoPolicy, true
}

// SetCryptoPolicy sets field value
func (o *EstProfile) SetCryptoPolicy(v ManagedCertificateProfileCryptoPolicy) {
	o.CryptoPolicy = v
}

// GetCsrDataMapping returns the CsrDataMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetCsrDataMapping() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.CsrDataMapping
}

// GetCsrDataMappingOk returns a tuple with the CsrDataMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetCsrDataMappingOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.CsrDataMapping) {
		return nil, false
	}
	return &o.CsrDataMapping, true
}

// HasCsrDataMapping returns a boolean if a field has been set.
func (o *EstProfile) HasCsrDataMapping() bool {
	if o != nil && !utils.IsNil(o.CsrDataMapping) {
		return true
	}

	return false
}

// SetCsrDataMapping gets a reference to the given map[string]string and assigns it to the CsrDataMapping field.
func (o *EstProfile) SetCsrDataMapping(v map[string]string) {
	o.CsrDataMapping = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EstProfile) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *EstProfile) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EstProfile) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *EstProfile) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetDnWhitelist returns the DnWhitelist field value
func (o *EstProfile) GetDnWhitelist() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DnWhitelist
}

// GetDnWhitelistOk returns a tuple with the DnWhitelist field value
// and a boolean to check if the value has been set.
func (o *EstProfile) GetDnWhitelistOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnWhitelist, true
}

// SetDnWhitelist sets field value
func (o *EstProfile) SetDnWhitelist(v bool) {
	o.DnWhitelist = v
}

// GetDsFlow returns the DsFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetDsFlow() []DataSourceFlowEntry {
	if o == nil {
		var ret []DataSourceFlowEntry
		return ret
	}
	return o.DsFlow
}

// GetDsFlowOk returns a tuple with the DsFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetDsFlowOk() ([]DataSourceFlowEntry, bool) {
	if o == nil || utils.IsNil(o.DsFlow) {
		return nil, false
	}
	return o.DsFlow, true
}

// HasDsFlow returns a boolean if a field has been set.
func (o *EstProfile) HasDsFlow() bool {
	if o != nil && !utils.IsNil(o.DsFlow) {
		return true
	}

	return false
}

// SetDsFlow gets a reference to the given []DataSourceFlowEntry and assigns it to the DsFlow field.
func (o *EstProfile) SetDsFlow(v []DataSourceFlowEntry) {
	o.DsFlow = v
}

// GetEnabled returns the Enabled field value
func (o *EstProfile) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *EstProfile) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *EstProfile) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnrollAuthorizedCas returns the EnrollAuthorizedCas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetEnrollAuthorizedCas() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EnrollAuthorizedCas
}

// GetEnrollAuthorizedCasOk returns a tuple with the EnrollAuthorizedCas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetEnrollAuthorizedCasOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.EnrollAuthorizedCas) {
		return nil, false
	}
	return o.EnrollAuthorizedCas, true
}

// HasEnrollAuthorizedCas returns a boolean if a field has been set.
func (o *EstProfile) HasEnrollAuthorizedCas() bool {
	if o != nil && !utils.IsNil(o.EnrollAuthorizedCas) {
		return true
	}

	return false
}

// SetEnrollAuthorizedCas gets a reference to the given []string and assigns it to the EnrollAuthorizedCas field.
func (o *EstProfile) SetEnrollAuthorizedCas(v []string) {
	o.EnrollAuthorizedCas = v
}

// GetGradingPolicies returns the GradingPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetGradingPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GradingPolicies
}

// GetGradingPoliciesOk returns a tuple with the GradingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetGradingPoliciesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.GradingPolicies) {
		return nil, false
	}
	return o.GradingPolicies, true
}

// HasGradingPolicies returns a boolean if a field has been set.
func (o *EstProfile) HasGradingPolicies() bool {
	if o != nil && !utils.IsNil(o.GradingPolicies) {
		return true
	}

	return false
}

// SetGradingPolicies gets a reference to the given []string and assigns it to the GradingPolicies field.
func (o *EstProfile) SetGradingPolicies(v []string) {
	o.GradingPolicies = v
}

// GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy {
	if o == nil || utils.IsNil(o.MaxCertificatePerHolderPolicy.Get()) {
		var ret MaxCertificatePerHolderPolicy
		return ret
	}
	return *o.MaxCertificatePerHolderPolicy.Get()
}

// GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCertificatePerHolderPolicy.Get(), o.MaxCertificatePerHolderPolicy.IsSet()
}

// HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.
func (o *EstProfile) HasMaxCertificatePerHolderPolicy() bool {
	if o != nil && o.MaxCertificatePerHolderPolicy.IsSet() {
		return true
	}

	return false
}

// SetMaxCertificatePerHolderPolicy gets a reference to the given NullableMaxCertificatePerHolderPolicy and assigns it to the MaxCertificatePerHolderPolicy field.
func (o *EstProfile) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy) {
	o.MaxCertificatePerHolderPolicy.Set(&v)
}

// SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil
func (o *EstProfile) SetMaxCertificatePerHolderPolicyNil() {
	o.MaxCertificatePerHolderPolicy.Set(nil)
}

// UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
func (o *EstProfile) UnsetMaxCertificatePerHolderPolicy() {
	o.MaxCertificatePerHolderPolicy.Unset()
}

// GetModule returns the Module field value
func (o *EstProfile) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *EstProfile) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *EstProfile) SetModule(v string) {
	o.Module = v
}

// GetName returns the Name field value
func (o *EstProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EstProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EstProfile) SetName(v string) {
	o.Name = v
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetPasswordPolicy() string {
	if o == nil || utils.IsNil(o.PasswordPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordPolicy.Get()
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetPasswordPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordPolicy.Get(), o.PasswordPolicy.IsSet()
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *EstProfile) HasPasswordPolicy() bool {
	if o != nil && o.PasswordPolicy.IsSet() {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given NullableString and assigns it to the PasswordPolicy field.
func (o *EstProfile) SetPasswordPolicy(v string) {
	o.PasswordPolicy.Set(&v)
}

// SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil
func (o *EstProfile) SetPasswordPolicyNil() {
	o.PasswordPolicy.Set(nil)
}

// UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
func (o *EstProfile) UnsetPasswordPolicy() {
	o.PasswordPolicy.Unset()
}

// GetPkiConnector returns the PkiConnector field value
func (o *EstProfile) GetPkiConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PkiConnector
}

// GetPkiConnectorOk returns a tuple with the PkiConnector field value
// and a boolean to check if the value has been set.
func (o *EstProfile) GetPkiConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiConnector, true
}

// SetPkiConnector sets field value
func (o *EstProfile) SetPkiConnector(v string) {
	o.PkiConnector = v
}

// GetRenewalAuthorizedCas returns the RenewalAuthorizedCas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetRenewalAuthorizedCas() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RenewalAuthorizedCas
}

// GetRenewalAuthorizedCasOk returns a tuple with the RenewalAuthorizedCas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetRenewalAuthorizedCasOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.RenewalAuthorizedCas) {
		return nil, false
	}
	return o.RenewalAuthorizedCas, true
}

// HasRenewalAuthorizedCas returns a boolean if a field has been set.
func (o *EstProfile) HasRenewalAuthorizedCas() bool {
	if o != nil && !utils.IsNil(o.RenewalAuthorizedCas) {
		return true
	}

	return false
}

// SetRenewalAuthorizedCas gets a reference to the given []string and assigns it to the RenewalAuthorizedCas field.
func (o *EstProfile) SetRenewalAuthorizedCas(v []string) {
	o.RenewalAuthorizedCas = v
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetRenewalPeriod() string {
	if o == nil || utils.IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *EstProfile) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *EstProfile) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}

// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *EstProfile) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *EstProfile) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetRequestsPolicy returns the RequestsPolicy field value
func (o *EstProfile) GetRequestsPolicy() RequestsPolicy {
	if o == nil {
		var ret RequestsPolicy
		return ret
	}

	return o.RequestsPolicy
}

// GetRequestsPolicyOk returns a tuple with the RequestsPolicy field value
// and a boolean to check if the value has been set.
func (o *EstProfile) GetRequestsPolicyOk() (*RequestsPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestsPolicy, true
}

// SetRequestsPolicy sets field value
func (o *EstProfile) SetRequestsPolicy(v RequestsPolicy) {
	o.RequestsPolicy = v
}

// GetSelfPermissions returns the SelfPermissions field value
func (o *EstProfile) GetSelfPermissions() CertificateProfileSelfPermissions {
	if o == nil {
		var ret CertificateProfileSelfPermissions
		return ret
	}

	return o.SelfPermissions
}

// GetSelfPermissionsOk returns a tuple with the SelfPermissions field value
// and a boolean to check if the value has been set.
func (o *EstProfile) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfPermissions, true
}

// SetSelfPermissions sets field value
func (o *EstProfile) SetSelfPermissions(v CertificateProfileSelfPermissions) {
	o.SelfPermissions = v
}

// GetTermsOfService returns the TermsOfService field value if set, zero value otherwise.
func (o *EstProfile) GetTermsOfService() string {
	if o == nil || utils.IsNil(o.TermsOfService) {
		var ret string
		return ret
	}
	return *o.TermsOfService
}

// GetTermsOfServiceOk returns a tuple with the TermsOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstProfile) GetTermsOfServiceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TermsOfService) {
		return nil, false
	}
	return o.TermsOfService, true
}

// HasTermsOfService returns a boolean if a field has been set.
func (o *EstProfile) HasTermsOfService() bool {
	if o != nil && !utils.IsNil(o.TermsOfService) {
		return true
	}

	return false
}

// SetTermsOfService gets a reference to the given string and assigns it to the TermsOfService field.
func (o *EstProfile) SetTermsOfService(v string) {
	o.TermsOfService = &v
}

// GetThirdPartyDiscoverySync returns the ThirdPartyDiscoverySync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetThirdPartyDiscoverySync() bool {
	if o == nil || utils.IsNil(o.ThirdPartyDiscoverySync.Get()) {
		var ret bool
		return ret
	}
	return *o.ThirdPartyDiscoverySync.Get()
}

// GetThirdPartyDiscoverySyncOk returns a tuple with the ThirdPartyDiscoverySync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetThirdPartyDiscoverySyncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThirdPartyDiscoverySync.Get(), o.ThirdPartyDiscoverySync.IsSet()
}

// HasThirdPartyDiscoverySync returns a boolean if a field has been set.
func (o *EstProfile) HasThirdPartyDiscoverySync() bool {
	if o != nil && o.ThirdPartyDiscoverySync.IsSet() {
		return true
	}

	return false
}

// SetThirdPartyDiscoverySync gets a reference to the given NullableBool and assigns it to the ThirdPartyDiscoverySync field.
func (o *EstProfile) SetThirdPartyDiscoverySync(v bool) {
	o.ThirdPartyDiscoverySync.Set(&v)
}

// SetThirdPartyDiscoverySyncNil sets the value for ThirdPartyDiscoverySync to be an explicit nil
func (o *EstProfile) SetThirdPartyDiscoverySyncNil() {
	o.ThirdPartyDiscoverySync.Set(nil)
}

// UnsetThirdPartyDiscoverySync ensures that no value is present for ThirdPartyDiscoverySync, not even an explicit nil
func (o *EstProfile) UnsetThirdPartyDiscoverySync() {
	o.ThirdPartyDiscoverySync.Unset()
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetTriggers() CertificateProfileTriggers {
	if o == nil || utils.IsNil(o.Triggers.Get()) {
		var ret CertificateProfileTriggers
		return ret
	}
	return *o.Triggers.Get()
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetTriggersOk() (*CertificateProfileTriggers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers.Get(), o.Triggers.IsSet()
}

// HasTriggers returns a boolean if a field has been set.
func (o *EstProfile) HasTriggers() bool {
	if o != nil && o.Triggers.IsSet() {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given NullableCertificateProfileTriggers and assigns it to the Triggers field.
func (o *EstProfile) SetTriggers(v CertificateProfileTriggers) {
	o.Triggers.Set(&v)
}

// SetTriggersNil sets the value for Triggers to be an explicit nil
func (o *EstProfile) SetTriggersNil() {
	o.Triggers.Set(nil)
}

// UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
func (o *EstProfile) UnsetTriggers() {
	o.Triggers.Unset()
}

// GetValidationRuleset returns the ValidationRuleset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfile) GetValidationRuleset() ValidationRuleset {
	if o == nil || utils.IsNil(o.ValidationRuleset.Get()) {
		var ret ValidationRuleset
		return ret
	}
	return *o.ValidationRuleset.Get()
}

// GetValidationRulesetOk returns a tuple with the ValidationRuleset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfile) GetValidationRulesetOk() (*ValidationRuleset, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationRuleset.Get(), o.ValidationRuleset.IsSet()
}

// HasValidationRuleset returns a boolean if a field has been set.
func (o *EstProfile) HasValidationRuleset() bool {
	if o != nil && o.ValidationRuleset.IsSet() {
		return true
	}

	return false
}

// SetValidationRuleset gets a reference to the given NullableValidationRuleset and assigns it to the ValidationRuleset field.
func (o *EstProfile) SetValidationRuleset(v ValidationRuleset) {
	o.ValidationRuleset.Set(&v)
}

// SetValidationRulesetNil sets the value for ValidationRuleset to be an explicit nil
func (o *EstProfile) SetValidationRulesetNil() {
	o.ValidationRuleset.Set(nil)
}

// UnsetValidationRuleset ensures that no value is present for ValidationRuleset, not even an explicit nil
func (o *EstProfile) UnsetValidationRuleset() {
	o.ValidationRuleset.Unset()
}

func (o EstProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authorizationLevels"] = o.AuthorizationLevels
	toSerialize["authorizationMode"] = o.AuthorizationMode
	toSerialize["ca"] = o.Ca
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["dnWhitelist"] = o.DnWhitelist
	if o.DsFlow != nil {
		toSerialize["dsFlow"] = o.DsFlow
	}
	toSerialize["enabled"] = o.Enabled
	if o.EnrollAuthorizedCas != nil {
		toSerialize["enrollAuthorizedCas"] = o.EnrollAuthorizedCas
	}
	if o.GradingPolicies != nil {
		toSerialize["gradingPolicies"] = o.GradingPolicies
	}
	if o.MaxCertificatePerHolderPolicy.IsSet() {
		toSerialize["maxCertificatePerHolderPolicy"] = o.MaxCertificatePerHolderPolicy.Get()
	}
	toSerialize["module"] = o.Module
	toSerialize["name"] = o.Name
	if o.PasswordPolicy.IsSet() {
		toSerialize["passwordPolicy"] = o.PasswordPolicy.Get()
	}
	toSerialize["pkiConnector"] = o.PkiConnector
	if o.RenewalAuthorizedCas != nil {
		toSerialize["renewalAuthorizedCas"] = o.RenewalAuthorizedCas
	}
	if o.RenewalPeriod.IsSet() {
		toSerialize["renewalPeriod"] = o.RenewalPeriod.Get()
	}
	toSerialize["requestsPolicy"] = o.RequestsPolicy
	toSerialize["selfPermissions"] = o.SelfPermissions
	if !utils.IsNil(o.TermsOfService) {
		toSerialize["termsOfService"] = o.TermsOfService
	}
	if o.ThirdPartyDiscoverySync.IsSet() {
		toSerialize["thirdPartyDiscoverySync"] = o.ThirdPartyDiscoverySync.Get()
	}
	if o.Triggers.IsSet() {
		toSerialize["triggers"] = o.Triggers.Get()
	}
	if o.ValidationRuleset.IsSet() {
		toSerialize["validationRuleset"] = o.ValidationRuleset.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EstProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorizationLevels",
		"authorizationMode",
		"ca",
		"cryptoPolicy",
		"dnWhitelist",
		"enabled",
		"module",
		"name",
		"pkiConnector",
		"requestsPolicy",
		"selfPermissions",
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

	varEstProfile := _EstProfile{}

	err = json.Unmarshal(data, &varEstProfile)

	if err != nil {
		return err
	}

	*o = EstProfile(varEstProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authorizationLevels")
		delete(additionalProperties, "authorizationMode")
		delete(additionalProperties, "ca")
		delete(additionalProperties, "certificateTemplate")
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "cryptoPolicy")
		delete(additionalProperties, "csrDataMapping")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "dnWhitelist")
		delete(additionalProperties, "dsFlow")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "enrollAuthorizedCas")
		delete(additionalProperties, "gradingPolicies")
		delete(additionalProperties, "maxCertificatePerHolderPolicy")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "passwordPolicy")
		delete(additionalProperties, "pkiConnector")
		delete(additionalProperties, "renewalAuthorizedCas")
		delete(additionalProperties, "renewalPeriod")
		delete(additionalProperties, "requestsPolicy")
		delete(additionalProperties, "selfPermissions")
		delete(additionalProperties, "termsOfService")
		delete(additionalProperties, "thirdPartyDiscoverySync")
		delete(additionalProperties, "triggers")
		delete(additionalProperties, "validationRuleset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEstProfile struct {
	value *EstProfile
	isSet bool
}

func (v NullableEstProfile) Get() *EstProfile {
	return v.value
}

func (v *NullableEstProfile) Set(val *EstProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableEstProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableEstProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstProfile(val *EstProfile) *NullableEstProfile {
	return &NullableEstProfile{value: val, isSet: true}
}

func (v NullableEstProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
