/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the AcmeExternalProfile type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AcmeExternalProfile{}

// AcmeExternalProfile struct for AcmeExternalProfile
type AcmeExternalProfile struct {
	AcmeUrl              *string                               `json:"acmeUrl,omitempty"`
	AuthorizationLevels  CertificateProfileAuthorizationLevels `json:"authorizationLevels"`
	AuthorizationMethods []string                              `json:"authorizationMethods"`
	AuthorizedCas        []string                              `json:"authorizedCas"`
	CertificateTemplate  NullableCertificateTemplate           `json:"certificateTemplate,omitempty"`
	Constraints          NullableCertificateRequestConstraints `json:"constraints,omitempty"`
	CryptoPolicy         ManagedCertificateProfileCryptoPolicy `json:"cryptoPolicy"`
	Description          []LocalizedString                     `json:"description,omitempty"`
	DisplayName          []LocalizedString                     `json:"displayName,omitempty"`
	// Representation of a datasource execution flow
	DsFlow                        []DataSourceFlowEntry                 `json:"dsFlow,omitempty"`
	Enabled                       bool                                  `json:"enabled"`
	GradingPolicies               []string                              `json:"gradingPolicies,omitempty"`
	MaxCertificatePerHolderPolicy NullableMaxCertificatePerHolderPolicy `json:"maxCertificatePerHolderPolicy,omitempty"`
	Module                        string                                `json:"module"`
	Name                          string                                `json:"name"`
	PkiConnector                  string                                `json:"pkiConnector"`
	RenewalPeriod                 utils.NullableString                  `json:"renewalPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	RequestsPolicy                RequestsPolicy                        `json:"requestsPolicy"`
	RequireEAB                    bool                                  `json:"requireEAB"`
	SelfPermissions               CertificateProfileSelfPermissions     `json:"selfPermissions"`
	// Available from `2.8.2`
	ThirdPartyDiscoverySync utils.NullableBool                 `json:"thirdPartyDiscoverySync,omitempty"`
	Triggers                NullableCertificateProfileTriggers `json:"triggers,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _AcmeExternalProfile AcmeExternalProfile

// NewAcmeExternalProfile instantiates a new AcmeExternalProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcmeExternalProfile(authorizationLevels CertificateProfileAuthorizationLevels, authorizationMethods []string, authorizedCas []string, cryptoPolicy ManagedCertificateProfileCryptoPolicy, enabled bool, module string, name string, pkiConnector string, requestsPolicy RequestsPolicy, requireEAB bool, selfPermissions CertificateProfileSelfPermissions) *AcmeExternalProfile {
	this := AcmeExternalProfile{}
	this.AuthorizationLevels = authorizationLevels
	this.AuthorizationMethods = authorizationMethods
	this.AuthorizedCas = authorizedCas
	this.CryptoPolicy = cryptoPolicy
	this.Enabled = enabled
	this.Module = module
	this.Name = name
	this.PkiConnector = pkiConnector
	this.RequestsPolicy = requestsPolicy
	this.RequireEAB = requireEAB
	this.SelfPermissions = selfPermissions
	var thirdPartyDiscoverySync bool = false
	this.ThirdPartyDiscoverySync = *utils.NewNullableBool(&thirdPartyDiscoverySync)
	return &this
}

// NewAcmeExternalProfileWithDefaults instantiates a new AcmeExternalProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcmeExternalProfileWithDefaults() *AcmeExternalProfile {
	this := AcmeExternalProfile{}
	var thirdPartyDiscoverySync bool = false
	this.ThirdPartyDiscoverySync = *utils.NewNullableBool(&thirdPartyDiscoverySync)
	return &this
}

// GetAcmeUrl returns the AcmeUrl field value if set, zero value otherwise.
func (o *AcmeExternalProfile) GetAcmeUrl() string {
	if o == nil || utils.IsNil(o.AcmeUrl) {
		var ret string
		return ret
	}
	return *o.AcmeUrl
}

// GetAcmeUrlOk returns a tuple with the AcmeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfile) GetAcmeUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AcmeUrl) {
		return nil, false
	}
	return o.AcmeUrl, true
}

// HasAcmeUrl returns a boolean if a field has been set.
func (o *AcmeExternalProfile) HasAcmeUrl() bool {
	if o != nil && !utils.IsNil(o.AcmeUrl) {
		return true
	}

	return false
}

// SetAcmeUrl gets a reference to the given string and assigns it to the AcmeUrl field.
func (o *AcmeExternalProfile) SetAcmeUrl(v string) {
	o.AcmeUrl = &v
}

// GetAuthorizationLevels returns the AuthorizationLevels field value
func (o *AcmeExternalProfile) GetAuthorizationLevels() CertificateProfileAuthorizationLevels {
	if o == nil {
		var ret CertificateProfileAuthorizationLevels
		return ret
	}

	return o.AuthorizationLevels
}

// GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfile) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationLevels, true
}

// SetAuthorizationLevels sets field value
func (o *AcmeExternalProfile) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels) {
	o.AuthorizationLevels = v
}

// GetAuthorizationMethods returns the AuthorizationMethods field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *AcmeExternalProfile) GetAuthorizationMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthorizationMethods
}

// GetAuthorizationMethodsOk returns a tuple with the AuthorizationMethods field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetAuthorizationMethodsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AuthorizationMethods) {
		return nil, false
	}
	return o.AuthorizationMethods, true
}

// SetAuthorizationMethods sets field value
func (o *AcmeExternalProfile) SetAuthorizationMethods(v []string) {
	o.AuthorizationMethods = v
}

// GetAuthorizedCas returns the AuthorizedCas field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *AcmeExternalProfile) GetAuthorizedCas() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthorizedCas
}

// GetAuthorizedCasOk returns a tuple with the AuthorizedCas field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetAuthorizedCasOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AuthorizedCas) {
		return nil, false
	}
	return o.AuthorizedCas, true
}

// SetAuthorizedCas sets field value
func (o *AcmeExternalProfile) SetAuthorizedCas(v []string) {
	o.AuthorizedCas = v
}

// GetCertificateTemplate returns the CertificateTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfile) GetCertificateTemplate() CertificateTemplate {
	if o == nil || utils.IsNil(o.CertificateTemplate.Get()) {
		var ret CertificateTemplate
		return ret
	}
	return *o.CertificateTemplate.Get()
}

// GetCertificateTemplateOk returns a tuple with the CertificateTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetCertificateTemplateOk() (*CertificateTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateTemplate.Get(), o.CertificateTemplate.IsSet()
}

// HasCertificateTemplate returns a boolean if a field has been set.
func (o *AcmeExternalProfile) HasCertificateTemplate() bool {
	if o != nil && o.CertificateTemplate.IsSet() {
		return true
	}

	return false
}

// SetCertificateTemplate gets a reference to the given NullableCertificateTemplate and assigns it to the CertificateTemplate field.
func (o *AcmeExternalProfile) SetCertificateTemplate(v CertificateTemplate) {
	o.CertificateTemplate.Set(&v)
}

// SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil
func (o *AcmeExternalProfile) SetCertificateTemplateNil() {
	o.CertificateTemplate.Set(nil)
}

// UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
func (o *AcmeExternalProfile) UnsetCertificateTemplate() {
	o.CertificateTemplate.Unset()
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfile) GetConstraints() CertificateRequestConstraints {
	if o == nil || utils.IsNil(o.Constraints.Get()) {
		var ret CertificateRequestConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetConstraintsOk() (*CertificateRequestConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *AcmeExternalProfile) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableCertificateRequestConstraints and assigns it to the Constraints field.
func (o *AcmeExternalProfile) SetConstraints(v CertificateRequestConstraints) {
	o.Constraints.Set(&v)
}

// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *AcmeExternalProfile) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *AcmeExternalProfile) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetCryptoPolicy returns the CryptoPolicy field value
func (o *AcmeExternalProfile) GetCryptoPolicy() ManagedCertificateProfileCryptoPolicy {
	if o == nil {
		var ret ManagedCertificateProfileCryptoPolicy
		return ret
	}

	return o.CryptoPolicy
}

// GetCryptoPolicyOk returns a tuple with the CryptoPolicy field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfile) GetCryptoPolicyOk() (*ManagedCertificateProfileCryptoPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoPolicy, true
}

// SetCryptoPolicy sets field value
func (o *AcmeExternalProfile) SetCryptoPolicy(v ManagedCertificateProfileCryptoPolicy) {
	o.CryptoPolicy = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfile) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AcmeExternalProfile) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *AcmeExternalProfile) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfile) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AcmeExternalProfile) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *AcmeExternalProfile) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetDsFlow returns the DsFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfile) GetDsFlow() []DataSourceFlowEntry {
	if o == nil {
		var ret []DataSourceFlowEntry
		return ret
	}
	return o.DsFlow
}

// GetDsFlowOk returns a tuple with the DsFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetDsFlowOk() ([]DataSourceFlowEntry, bool) {
	if o == nil || utils.IsNil(o.DsFlow) {
		return nil, false
	}
	return o.DsFlow, true
}

// HasDsFlow returns a boolean if a field has been set.
func (o *AcmeExternalProfile) HasDsFlow() bool {
	if o != nil && !utils.IsNil(o.DsFlow) {
		return true
	}

	return false
}

// SetDsFlow gets a reference to the given []DataSourceFlowEntry and assigns it to the DsFlow field.
func (o *AcmeExternalProfile) SetDsFlow(v []DataSourceFlowEntry) {
	o.DsFlow = v
}

// GetEnabled returns the Enabled field value
func (o *AcmeExternalProfile) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfile) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AcmeExternalProfile) SetEnabled(v bool) {
	o.Enabled = v
}

// GetGradingPolicies returns the GradingPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfile) GetGradingPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GradingPolicies
}

// GetGradingPoliciesOk returns a tuple with the GradingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetGradingPoliciesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.GradingPolicies) {
		return nil, false
	}
	return o.GradingPolicies, true
}

// HasGradingPolicies returns a boolean if a field has been set.
func (o *AcmeExternalProfile) HasGradingPolicies() bool {
	if o != nil && !utils.IsNil(o.GradingPolicies) {
		return true
	}

	return false
}

// SetGradingPolicies gets a reference to the given []string and assigns it to the GradingPolicies field.
func (o *AcmeExternalProfile) SetGradingPolicies(v []string) {
	o.GradingPolicies = v
}

// GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfile) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy {
	if o == nil || utils.IsNil(o.MaxCertificatePerHolderPolicy.Get()) {
		var ret MaxCertificatePerHolderPolicy
		return ret
	}
	return *o.MaxCertificatePerHolderPolicy.Get()
}

// GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCertificatePerHolderPolicy.Get(), o.MaxCertificatePerHolderPolicy.IsSet()
}

// HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.
func (o *AcmeExternalProfile) HasMaxCertificatePerHolderPolicy() bool {
	if o != nil && o.MaxCertificatePerHolderPolicy.IsSet() {
		return true
	}

	return false
}

// SetMaxCertificatePerHolderPolicy gets a reference to the given NullableMaxCertificatePerHolderPolicy and assigns it to the MaxCertificatePerHolderPolicy field.
func (o *AcmeExternalProfile) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy) {
	o.MaxCertificatePerHolderPolicy.Set(&v)
}

// SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil
func (o *AcmeExternalProfile) SetMaxCertificatePerHolderPolicyNil() {
	o.MaxCertificatePerHolderPolicy.Set(nil)
}

// UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
func (o *AcmeExternalProfile) UnsetMaxCertificatePerHolderPolicy() {
	o.MaxCertificatePerHolderPolicy.Unset()
}

// GetModule returns the Module field value
func (o *AcmeExternalProfile) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfile) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *AcmeExternalProfile) SetModule(v string) {
	o.Module = v
}

// GetName returns the Name field value
func (o *AcmeExternalProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AcmeExternalProfile) SetName(v string) {
	o.Name = v
}

// GetPkiConnector returns the PkiConnector field value
func (o *AcmeExternalProfile) GetPkiConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PkiConnector
}

// GetPkiConnectorOk returns a tuple with the PkiConnector field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfile) GetPkiConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiConnector, true
}

// SetPkiConnector sets field value
func (o *AcmeExternalProfile) SetPkiConnector(v string) {
	o.PkiConnector = v
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfile) GetRenewalPeriod() string {
	if o == nil || utils.IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *AcmeExternalProfile) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *AcmeExternalProfile) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}

// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *AcmeExternalProfile) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *AcmeExternalProfile) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetRequestsPolicy returns the RequestsPolicy field value
func (o *AcmeExternalProfile) GetRequestsPolicy() RequestsPolicy {
	if o == nil {
		var ret RequestsPolicy
		return ret
	}

	return o.RequestsPolicy
}

// GetRequestsPolicyOk returns a tuple with the RequestsPolicy field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfile) GetRequestsPolicyOk() (*RequestsPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestsPolicy, true
}

// SetRequestsPolicy sets field value
func (o *AcmeExternalProfile) SetRequestsPolicy(v RequestsPolicy) {
	o.RequestsPolicy = v
}

// GetRequireEAB returns the RequireEAB field value
func (o *AcmeExternalProfile) GetRequireEAB() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireEAB
}

// GetRequireEABOk returns a tuple with the RequireEAB field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfile) GetRequireEABOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireEAB, true
}

// SetRequireEAB sets field value
func (o *AcmeExternalProfile) SetRequireEAB(v bool) {
	o.RequireEAB = v
}

// GetSelfPermissions returns the SelfPermissions field value
func (o *AcmeExternalProfile) GetSelfPermissions() CertificateProfileSelfPermissions {
	if o == nil {
		var ret CertificateProfileSelfPermissions
		return ret
	}

	return o.SelfPermissions
}

// GetSelfPermissionsOk returns a tuple with the SelfPermissions field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfile) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfPermissions, true
}

// SetSelfPermissions sets field value
func (o *AcmeExternalProfile) SetSelfPermissions(v CertificateProfileSelfPermissions) {
	o.SelfPermissions = v
}

// GetThirdPartyDiscoverySync returns the ThirdPartyDiscoverySync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfile) GetThirdPartyDiscoverySync() bool {
	if o == nil || utils.IsNil(o.ThirdPartyDiscoverySync.Get()) {
		var ret bool
		return ret
	}
	return *o.ThirdPartyDiscoverySync.Get()
}

// GetThirdPartyDiscoverySyncOk returns a tuple with the ThirdPartyDiscoverySync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetThirdPartyDiscoverySyncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThirdPartyDiscoverySync.Get(), o.ThirdPartyDiscoverySync.IsSet()
}

// HasThirdPartyDiscoverySync returns a boolean if a field has been set.
func (o *AcmeExternalProfile) HasThirdPartyDiscoverySync() bool {
	if o != nil && o.ThirdPartyDiscoverySync.IsSet() {
		return true
	}

	return false
}

// SetThirdPartyDiscoverySync gets a reference to the given NullableBool and assigns it to the ThirdPartyDiscoverySync field.
func (o *AcmeExternalProfile) SetThirdPartyDiscoverySync(v bool) {
	o.ThirdPartyDiscoverySync.Set(&v)
}

// SetThirdPartyDiscoverySyncNil sets the value for ThirdPartyDiscoverySync to be an explicit nil
func (o *AcmeExternalProfile) SetThirdPartyDiscoverySyncNil() {
	o.ThirdPartyDiscoverySync.Set(nil)
}

// UnsetThirdPartyDiscoverySync ensures that no value is present for ThirdPartyDiscoverySync, not even an explicit nil
func (o *AcmeExternalProfile) UnsetThirdPartyDiscoverySync() {
	o.ThirdPartyDiscoverySync.Unset()
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfile) GetTriggers() CertificateProfileTriggers {
	if o == nil || utils.IsNil(o.Triggers.Get()) {
		var ret CertificateProfileTriggers
		return ret
	}
	return *o.Triggers.Get()
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfile) GetTriggersOk() (*CertificateProfileTriggers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers.Get(), o.Triggers.IsSet()
}

// HasTriggers returns a boolean if a field has been set.
func (o *AcmeExternalProfile) HasTriggers() bool {
	if o != nil && o.Triggers.IsSet() {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given NullableCertificateProfileTriggers and assigns it to the Triggers field.
func (o *AcmeExternalProfile) SetTriggers(v CertificateProfileTriggers) {
	o.Triggers.Set(&v)
}

// SetTriggersNil sets the value for Triggers to be an explicit nil
func (o *AcmeExternalProfile) SetTriggersNil() {
	o.Triggers.Set(nil)
}

// UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
func (o *AcmeExternalProfile) UnsetTriggers() {
	o.Triggers.Unset()
}

func (o AcmeExternalProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcmeExternalProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AcmeUrl) {
		toSerialize["acmeUrl"] = o.AcmeUrl
	}
	toSerialize["authorizationLevels"] = o.AuthorizationLevels
	if o.AuthorizationMethods != nil {
		toSerialize["authorizationMethods"] = o.AuthorizationMethods
	}
	if o.AuthorizedCas != nil {
		toSerialize["authorizedCas"] = o.AuthorizedCas
	}
	if o.CertificateTemplate.IsSet() {
		toSerialize["certificateTemplate"] = o.CertificateTemplate.Get()
	}
	if o.Constraints.IsSet() {
		toSerialize["constraints"] = o.Constraints.Get()
	}
	toSerialize["cryptoPolicy"] = o.CryptoPolicy
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
	if o.MaxCertificatePerHolderPolicy.IsSet() {
		toSerialize["maxCertificatePerHolderPolicy"] = o.MaxCertificatePerHolderPolicy.Get()
	}
	toSerialize["module"] = o.Module
	toSerialize["name"] = o.Name
	toSerialize["pkiConnector"] = o.PkiConnector
	if o.RenewalPeriod.IsSet() {
		toSerialize["renewalPeriod"] = o.RenewalPeriod.Get()
	}
	toSerialize["requestsPolicy"] = o.RequestsPolicy
	toSerialize["requireEAB"] = o.RequireEAB
	toSerialize["selfPermissions"] = o.SelfPermissions
	if o.ThirdPartyDiscoverySync.IsSet() {
		toSerialize["thirdPartyDiscoverySync"] = o.ThirdPartyDiscoverySync.Get()
	}
	if o.Triggers.IsSet() {
		toSerialize["triggers"] = o.Triggers.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcmeExternalProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authorizationLevels",
		"authorizationMethods",
		"authorizedCas",
		"cryptoPolicy",
		"enabled",
		"module",
		"name",
		"pkiConnector",
		"requestsPolicy",
		"requireEAB",
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

	varAcmeExternalProfile := _AcmeExternalProfile{}

	err = json.Unmarshal(data, &varAcmeExternalProfile)

	if err != nil {
		return err
	}

	*o = AcmeExternalProfile(varAcmeExternalProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "acmeUrl")
		delete(additionalProperties, "authorizationLevels")
		delete(additionalProperties, "authorizationMethods")
		delete(additionalProperties, "authorizedCas")
		delete(additionalProperties, "certificateTemplate")
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "cryptoPolicy")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "dsFlow")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "gradingPolicies")
		delete(additionalProperties, "maxCertificatePerHolderPolicy")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "pkiConnector")
		delete(additionalProperties, "renewalPeriod")
		delete(additionalProperties, "requestsPolicy")
		delete(additionalProperties, "requireEAB")
		delete(additionalProperties, "selfPermissions")
		delete(additionalProperties, "thirdPartyDiscoverySync")
		delete(additionalProperties, "triggers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcmeExternalProfile struct {
	value *AcmeExternalProfile
	isSet bool
}

func (v NullableAcmeExternalProfile) Get() *AcmeExternalProfile {
	return v.value
}

func (v *NullableAcmeExternalProfile) Set(val *AcmeExternalProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAcmeExternalProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAcmeExternalProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcmeExternalProfile(val *AcmeExternalProfile) *NullableAcmeExternalProfile {
	return &NullableAcmeExternalProfile{value: val, isSet: true}
}

func (v NullableAcmeExternalProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcmeExternalProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
