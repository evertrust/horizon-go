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

// checks if the EstProfileResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EstProfileResponse{}

// EstProfileResponse struct for EstProfileResponse
type EstProfileResponse struct {
	// Object internal ID
	Id                  string                                `json:"_id"`
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
	// Available from `2.8.2`
	ThirdPartyDiscoverySync utils.NullableBool                 `json:"thirdPartyDiscoverySync,omitempty"`
	Triggers                NullableCertificateProfileTriggers `json:"triggers,omitempty"`
	ValidationRuleset       NullableValidationRuleset          `json:"validationRuleset,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _EstProfileResponse EstProfileResponse

// NewEstProfileResponse instantiates a new EstProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstProfileResponse(id string, authorizationLevels CertificateProfileAuthorizationLevels, authorizationMode string, ca string, cryptoPolicy ManagedCertificateProfileCryptoPolicy, dnWhitelist bool, enabled bool, module string, name string, pkiConnector string, requestsPolicy RequestsPolicy, selfPermissions CertificateProfileSelfPermissions) *EstProfileResponse {
	this := EstProfileResponse{}
	this.Id = id
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

// NewEstProfileResponseWithDefaults instantiates a new EstProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstProfileResponseWithDefaults() *EstProfileResponse {
	this := EstProfileResponse{}
	var thirdPartyDiscoverySync bool = false
	this.ThirdPartyDiscoverySync = *utils.NewNullableBool(&thirdPartyDiscoverySync)
	return &this
}

// GetId returns the Id field value
func (o *EstProfileResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EstProfileResponse) SetId(v string) {
	o.Id = v
}

// GetAuthorizationLevels returns the AuthorizationLevels field value
func (o *EstProfileResponse) GetAuthorizationLevels() CertificateProfileAuthorizationLevels {
	if o == nil {
		var ret CertificateProfileAuthorizationLevels
		return ret
	}

	return o.AuthorizationLevels
}

// GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationLevels, true
}

// SetAuthorizationLevels sets field value
func (o *EstProfileResponse) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels) {
	o.AuthorizationLevels = v
}

// GetAuthorizationMode returns the AuthorizationMode field value
func (o *EstProfileResponse) GetAuthorizationMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthorizationMode
}

// GetAuthorizationModeOk returns a tuple with the AuthorizationMode field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetAuthorizationModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationMode, true
}

// SetAuthorizationMode sets field value
func (o *EstProfileResponse) SetAuthorizationMode(v string) {
	o.AuthorizationMode = v
}

// GetCa returns the Ca field value
func (o *EstProfileResponse) GetCa() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ca
}

// GetCaOk returns a tuple with the Ca field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetCaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ca, true
}

// SetCa sets field value
func (o *EstProfileResponse) SetCa(v string) {
	o.Ca = v
}

// GetCertificateTemplate returns the CertificateTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetCertificateTemplate() CertificateTemplate {
	if o == nil || utils.IsNil(o.CertificateTemplate.Get()) {
		var ret CertificateTemplate
		return ret
	}
	return *o.CertificateTemplate.Get()
}

// GetCertificateTemplateOk returns a tuple with the CertificateTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetCertificateTemplateOk() (*CertificateTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateTemplate.Get(), o.CertificateTemplate.IsSet()
}

// HasCertificateTemplate returns a boolean if a field has been set.
func (o *EstProfileResponse) HasCertificateTemplate() bool {
	if o != nil && o.CertificateTemplate.IsSet() {
		return true
	}

	return false
}

// SetCertificateTemplate gets a reference to the given NullableCertificateTemplate and assigns it to the CertificateTemplate field.
func (o *EstProfileResponse) SetCertificateTemplate(v CertificateTemplate) {
	o.CertificateTemplate.Set(&v)
}

// SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil
func (o *EstProfileResponse) SetCertificateTemplateNil() {
	o.CertificateTemplate.Set(nil)
}

// UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
func (o *EstProfileResponse) UnsetCertificateTemplate() {
	o.CertificateTemplate.Unset()
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetConstraints() CertificateRequestConstraints {
	if o == nil || utils.IsNil(o.Constraints.Get()) {
		var ret CertificateRequestConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetConstraintsOk() (*CertificateRequestConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *EstProfileResponse) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableCertificateRequestConstraints and assigns it to the Constraints field.
func (o *EstProfileResponse) SetConstraints(v CertificateRequestConstraints) {
	o.Constraints.Set(&v)
}

// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *EstProfileResponse) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *EstProfileResponse) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetCryptoPolicy returns the CryptoPolicy field value
func (o *EstProfileResponse) GetCryptoPolicy() ManagedCertificateProfileCryptoPolicy {
	if o == nil {
		var ret ManagedCertificateProfileCryptoPolicy
		return ret
	}

	return o.CryptoPolicy
}

// GetCryptoPolicyOk returns a tuple with the CryptoPolicy field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetCryptoPolicyOk() (*ManagedCertificateProfileCryptoPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoPolicy, true
}

// SetCryptoPolicy sets field value
func (o *EstProfileResponse) SetCryptoPolicy(v ManagedCertificateProfileCryptoPolicy) {
	o.CryptoPolicy = v
}

// GetCsrDataMapping returns the CsrDataMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetCsrDataMapping() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.CsrDataMapping
}

// GetCsrDataMappingOk returns a tuple with the CsrDataMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetCsrDataMappingOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.CsrDataMapping) {
		return nil, false
	}
	return &o.CsrDataMapping, true
}

// HasCsrDataMapping returns a boolean if a field has been set.
func (o *EstProfileResponse) HasCsrDataMapping() bool {
	if o != nil && !utils.IsNil(o.CsrDataMapping) {
		return true
	}

	return false
}

// SetCsrDataMapping gets a reference to the given map[string]string and assigns it to the CsrDataMapping field.
func (o *EstProfileResponse) SetCsrDataMapping(v map[string]string) {
	o.CsrDataMapping = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EstProfileResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *EstProfileResponse) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EstProfileResponse) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *EstProfileResponse) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetDnWhitelist returns the DnWhitelist field value
func (o *EstProfileResponse) GetDnWhitelist() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DnWhitelist
}

// GetDnWhitelistOk returns a tuple with the DnWhitelist field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetDnWhitelistOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnWhitelist, true
}

// SetDnWhitelist sets field value
func (o *EstProfileResponse) SetDnWhitelist(v bool) {
	o.DnWhitelist = v
}

// GetDsFlow returns the DsFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetDsFlow() []DataSourceFlowEntry {
	if o == nil {
		var ret []DataSourceFlowEntry
		return ret
	}
	return o.DsFlow
}

// GetDsFlowOk returns a tuple with the DsFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetDsFlowOk() ([]DataSourceFlowEntry, bool) {
	if o == nil || utils.IsNil(o.DsFlow) {
		return nil, false
	}
	return o.DsFlow, true
}

// HasDsFlow returns a boolean if a field has been set.
func (o *EstProfileResponse) HasDsFlow() bool {
	if o != nil && !utils.IsNil(o.DsFlow) {
		return true
	}

	return false
}

// SetDsFlow gets a reference to the given []DataSourceFlowEntry and assigns it to the DsFlow field.
func (o *EstProfileResponse) SetDsFlow(v []DataSourceFlowEntry) {
	o.DsFlow = v
}

// GetEnabled returns the Enabled field value
func (o *EstProfileResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *EstProfileResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnrollAuthorizedCas returns the EnrollAuthorizedCas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetEnrollAuthorizedCas() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EnrollAuthorizedCas
}

// GetEnrollAuthorizedCasOk returns a tuple with the EnrollAuthorizedCas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetEnrollAuthorizedCasOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.EnrollAuthorizedCas) {
		return nil, false
	}
	return o.EnrollAuthorizedCas, true
}

// HasEnrollAuthorizedCas returns a boolean if a field has been set.
func (o *EstProfileResponse) HasEnrollAuthorizedCas() bool {
	if o != nil && !utils.IsNil(o.EnrollAuthorizedCas) {
		return true
	}

	return false
}

// SetEnrollAuthorizedCas gets a reference to the given []string and assigns it to the EnrollAuthorizedCas field.
func (o *EstProfileResponse) SetEnrollAuthorizedCas(v []string) {
	o.EnrollAuthorizedCas = v
}

// GetGradingPolicies returns the GradingPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetGradingPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GradingPolicies
}

// GetGradingPoliciesOk returns a tuple with the GradingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetGradingPoliciesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.GradingPolicies) {
		return nil, false
	}
	return o.GradingPolicies, true
}

// HasGradingPolicies returns a boolean if a field has been set.
func (o *EstProfileResponse) HasGradingPolicies() bool {
	if o != nil && !utils.IsNil(o.GradingPolicies) {
		return true
	}

	return false
}

// SetGradingPolicies gets a reference to the given []string and assigns it to the GradingPolicies field.
func (o *EstProfileResponse) SetGradingPolicies(v []string) {
	o.GradingPolicies = v
}

// GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy {
	if o == nil || utils.IsNil(o.MaxCertificatePerHolderPolicy.Get()) {
		var ret MaxCertificatePerHolderPolicy
		return ret
	}
	return *o.MaxCertificatePerHolderPolicy.Get()
}

// GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCertificatePerHolderPolicy.Get(), o.MaxCertificatePerHolderPolicy.IsSet()
}

// HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.
func (o *EstProfileResponse) HasMaxCertificatePerHolderPolicy() bool {
	if o != nil && o.MaxCertificatePerHolderPolicy.IsSet() {
		return true
	}

	return false
}

// SetMaxCertificatePerHolderPolicy gets a reference to the given NullableMaxCertificatePerHolderPolicy and assigns it to the MaxCertificatePerHolderPolicy field.
func (o *EstProfileResponse) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy) {
	o.MaxCertificatePerHolderPolicy.Set(&v)
}

// SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil
func (o *EstProfileResponse) SetMaxCertificatePerHolderPolicyNil() {
	o.MaxCertificatePerHolderPolicy.Set(nil)
}

// UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
func (o *EstProfileResponse) UnsetMaxCertificatePerHolderPolicy() {
	o.MaxCertificatePerHolderPolicy.Unset()
}

// GetModule returns the Module field value
func (o *EstProfileResponse) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *EstProfileResponse) SetModule(v string) {
	o.Module = v
}

// GetName returns the Name field value
func (o *EstProfileResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EstProfileResponse) SetName(v string) {
	o.Name = v
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetPasswordPolicy() string {
	if o == nil || utils.IsNil(o.PasswordPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordPolicy.Get()
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetPasswordPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordPolicy.Get(), o.PasswordPolicy.IsSet()
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *EstProfileResponse) HasPasswordPolicy() bool {
	if o != nil && o.PasswordPolicy.IsSet() {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given NullableString and assigns it to the PasswordPolicy field.
func (o *EstProfileResponse) SetPasswordPolicy(v string) {
	o.PasswordPolicy.Set(&v)
}

// SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil
func (o *EstProfileResponse) SetPasswordPolicyNil() {
	o.PasswordPolicy.Set(nil)
}

// UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
func (o *EstProfileResponse) UnsetPasswordPolicy() {
	o.PasswordPolicy.Unset()
}

// GetPkiConnector returns the PkiConnector field value
func (o *EstProfileResponse) GetPkiConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PkiConnector
}

// GetPkiConnectorOk returns a tuple with the PkiConnector field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetPkiConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiConnector, true
}

// SetPkiConnector sets field value
func (o *EstProfileResponse) SetPkiConnector(v string) {
	o.PkiConnector = v
}

// GetRenewalAuthorizedCas returns the RenewalAuthorizedCas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetRenewalAuthorizedCas() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RenewalAuthorizedCas
}

// GetRenewalAuthorizedCasOk returns a tuple with the RenewalAuthorizedCas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetRenewalAuthorizedCasOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.RenewalAuthorizedCas) {
		return nil, false
	}
	return o.RenewalAuthorizedCas, true
}

// HasRenewalAuthorizedCas returns a boolean if a field has been set.
func (o *EstProfileResponse) HasRenewalAuthorizedCas() bool {
	if o != nil && !utils.IsNil(o.RenewalAuthorizedCas) {
		return true
	}

	return false
}

// SetRenewalAuthorizedCas gets a reference to the given []string and assigns it to the RenewalAuthorizedCas field.
func (o *EstProfileResponse) SetRenewalAuthorizedCas(v []string) {
	o.RenewalAuthorizedCas = v
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetRenewalPeriod() string {
	if o == nil || utils.IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *EstProfileResponse) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *EstProfileResponse) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}

// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *EstProfileResponse) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *EstProfileResponse) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetRequestsPolicy returns the RequestsPolicy field value
func (o *EstProfileResponse) GetRequestsPolicy() RequestsPolicy {
	if o == nil {
		var ret RequestsPolicy
		return ret
	}

	return o.RequestsPolicy
}

// GetRequestsPolicyOk returns a tuple with the RequestsPolicy field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetRequestsPolicyOk() (*RequestsPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestsPolicy, true
}

// SetRequestsPolicy sets field value
func (o *EstProfileResponse) SetRequestsPolicy(v RequestsPolicy) {
	o.RequestsPolicy = v
}

// GetSelfPermissions returns the SelfPermissions field value
func (o *EstProfileResponse) GetSelfPermissions() CertificateProfileSelfPermissions {
	if o == nil {
		var ret CertificateProfileSelfPermissions
		return ret
	}

	return o.SelfPermissions
}

// GetSelfPermissionsOk returns a tuple with the SelfPermissions field value
// and a boolean to check if the value has been set.
func (o *EstProfileResponse) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfPermissions, true
}

// SetSelfPermissions sets field value
func (o *EstProfileResponse) SetSelfPermissions(v CertificateProfileSelfPermissions) {
	o.SelfPermissions = v
}

// GetThirdPartyDiscoverySync returns the ThirdPartyDiscoverySync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetThirdPartyDiscoverySync() bool {
	if o == nil || utils.IsNil(o.ThirdPartyDiscoverySync.Get()) {
		var ret bool
		return ret
	}
	return *o.ThirdPartyDiscoverySync.Get()
}

// GetThirdPartyDiscoverySyncOk returns a tuple with the ThirdPartyDiscoverySync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetThirdPartyDiscoverySyncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThirdPartyDiscoverySync.Get(), o.ThirdPartyDiscoverySync.IsSet()
}

// HasThirdPartyDiscoverySync returns a boolean if a field has been set.
func (o *EstProfileResponse) HasThirdPartyDiscoverySync() bool {
	if o != nil && o.ThirdPartyDiscoverySync.IsSet() {
		return true
	}

	return false
}

// SetThirdPartyDiscoverySync gets a reference to the given NullableBool and assigns it to the ThirdPartyDiscoverySync field.
func (o *EstProfileResponse) SetThirdPartyDiscoverySync(v bool) {
	o.ThirdPartyDiscoverySync.Set(&v)
}

// SetThirdPartyDiscoverySyncNil sets the value for ThirdPartyDiscoverySync to be an explicit nil
func (o *EstProfileResponse) SetThirdPartyDiscoverySyncNil() {
	o.ThirdPartyDiscoverySync.Set(nil)
}

// UnsetThirdPartyDiscoverySync ensures that no value is present for ThirdPartyDiscoverySync, not even an explicit nil
func (o *EstProfileResponse) UnsetThirdPartyDiscoverySync() {
	o.ThirdPartyDiscoverySync.Unset()
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetTriggers() CertificateProfileTriggers {
	if o == nil || utils.IsNil(o.Triggers.Get()) {
		var ret CertificateProfileTriggers
		return ret
	}
	return *o.Triggers.Get()
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetTriggersOk() (*CertificateProfileTriggers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers.Get(), o.Triggers.IsSet()
}

// HasTriggers returns a boolean if a field has been set.
func (o *EstProfileResponse) HasTriggers() bool {
	if o != nil && o.Triggers.IsSet() {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given NullableCertificateProfileTriggers and assigns it to the Triggers field.
func (o *EstProfileResponse) SetTriggers(v CertificateProfileTriggers) {
	o.Triggers.Set(&v)
}

// SetTriggersNil sets the value for Triggers to be an explicit nil
func (o *EstProfileResponse) SetTriggersNil() {
	o.Triggers.Set(nil)
}

// UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
func (o *EstProfileResponse) UnsetTriggers() {
	o.Triggers.Unset()
}

// GetValidationRuleset returns the ValidationRuleset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EstProfileResponse) GetValidationRuleset() ValidationRuleset {
	if o == nil || utils.IsNil(o.ValidationRuleset.Get()) {
		var ret ValidationRuleset
		return ret
	}
	return *o.ValidationRuleset.Get()
}

// GetValidationRulesetOk returns a tuple with the ValidationRuleset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EstProfileResponse) GetValidationRulesetOk() (*ValidationRuleset, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationRuleset.Get(), o.ValidationRuleset.IsSet()
}

// HasValidationRuleset returns a boolean if a field has been set.
func (o *EstProfileResponse) HasValidationRuleset() bool {
	if o != nil && o.ValidationRuleset.IsSet() {
		return true
	}

	return false
}

// SetValidationRuleset gets a reference to the given NullableValidationRuleset and assigns it to the ValidationRuleset field.
func (o *EstProfileResponse) SetValidationRuleset(v ValidationRuleset) {
	o.ValidationRuleset.Set(&v)
}

// SetValidationRulesetNil sets the value for ValidationRuleset to be an explicit nil
func (o *EstProfileResponse) SetValidationRulesetNil() {
	o.ValidationRuleset.Set(nil)
}

// UnsetValidationRuleset ensures that no value is present for ValidationRuleset, not even an explicit nil
func (o *EstProfileResponse) UnsetValidationRuleset() {
	o.ValidationRuleset.Unset()
}

func (o EstProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
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

func (o *EstProfileResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
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

	varEstProfileResponse := _EstProfileResponse{}

	err = json.Unmarshal(data, &varEstProfileResponse)

	if err != nil {
		return err
	}

	*o = EstProfileResponse(varEstProfileResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
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
		delete(additionalProperties, "thirdPartyDiscoverySync")
		delete(additionalProperties, "triggers")
		delete(additionalProperties, "validationRuleset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEstProfileResponse struct {
	value *EstProfileResponse
	isSet bool
}

func (v NullableEstProfileResponse) Get() *EstProfileResponse {
	return v.value
}

func (v *NullableEstProfileResponse) Set(val *EstProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEstProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEstProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstProfileResponse(val *EstProfileResponse) *NullableEstProfileResponse {
	return &NullableEstProfileResponse{value: val, isSet: true}
}

func (v NullableEstProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
