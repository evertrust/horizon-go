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

// checks if the JamfProfileResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &JamfProfileResponse{}

// JamfProfileResponse struct for JamfProfileResponse
type JamfProfileResponse struct {
	// Object internal ID
	Id                  string                                `json:"_id"`
	AuthorizationLevels CertificateProfileAuthorizationLevels `json:"authorizationLevels"`
	Caps                []string                              `json:"caps"`
	CertificateTemplate NullableCertificateTemplate           `json:"certificateTemplate,omitempty"`
	Constraints         NullableCertificateRequestConstraints `json:"constraints,omitempty"`
	CryptoPolicy        ManagedCertificateProfileCryptoPolicy `json:"cryptoPolicy"`
	CsrDataMapping      map[string]string                     `json:"csrDataMapping,omitempty"`
	Description         []LocalizedString                     `json:"description,omitempty"`
	DeviceIdField       utils.NullableString                  `json:"deviceIdField,omitempty"`
	DisplayName         []LocalizedString                     `json:"displayName,omitempty"`
	// Representation of a datasource execution flow
	DsFlow                        []DataSourceFlowEntry                 `json:"dsFlow,omitempty"`
	Enabled                       bool                                  `json:"enabled"`
	EncryptionAlgorithm           string                                `json:"encryptionAlgorithm"`
	GradingPolicies               []string                              `json:"gradingPolicies,omitempty"`
	MaxCertificatePerHolderPolicy NullableMaxCertificatePerHolderPolicy `json:"maxCertificatePerHolderPolicy,omitempty"`
	Mode                          string                                `json:"mode"`
	Module                        string                                `json:"module"`
	Name                          string                                `json:"name"`
	PasswordPolicy                utils.NullableString                  `json:"passwordPolicy,omitempty"`
	PkiConnector                  string                                `json:"pkiConnector"`
	PostPKIOperation              utils.NullableBool                    `json:"postPKIOperation,omitempty"`
	RenewalPeriod                 utils.NullableString                  `json:"renewalPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	RequestsPolicy                RequestsPolicy                        `json:"requestsPolicy"`
	ScepRA                        string                                `json:"scepRA"`
	SelfPermissions               CertificateProfileSelfPermissions     `json:"selfPermissions"`
	ThirdPartyConnector           string                                `json:"thirdPartyConnector"`
	ThirdPartyDiscoverySync       utils.NullableBool                    `json:"thirdPartyDiscoverySync,omitempty"`
	Triggers                      NullableCertificateProfileTriggers    `json:"triggers,omitempty"`
	AdditionalProperties          map[string]interface{}
}

type _JamfProfileResponse JamfProfileResponse

// NewJamfProfileResponse instantiates a new JamfProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJamfProfileResponse(id string, authorizationLevels CertificateProfileAuthorizationLevels, caps []string, cryptoPolicy ManagedCertificateProfileCryptoPolicy, enabled bool, encryptionAlgorithm string, mode string, module string, name string, pkiConnector string, requestsPolicy RequestsPolicy, scepRA string, selfPermissions CertificateProfileSelfPermissions, thirdPartyConnector string) *JamfProfileResponse {
	this := JamfProfileResponse{}
	this.Id = id
	this.AuthorizationLevels = authorizationLevels
	this.Caps = caps
	this.CryptoPolicy = cryptoPolicy
	this.Enabled = enabled
	this.EncryptionAlgorithm = encryptionAlgorithm
	this.Mode = mode
	this.Module = module
	this.Name = name
	this.PkiConnector = pkiConnector
	this.RequestsPolicy = requestsPolicy
	this.ScepRA = scepRA
	this.SelfPermissions = selfPermissions
	this.ThirdPartyConnector = thirdPartyConnector
	var thirdPartyDiscoverySync bool = false
	this.ThirdPartyDiscoverySync = *utils.NewNullableBool(&thirdPartyDiscoverySync)
	return &this
}

// NewJamfProfileResponseWithDefaults instantiates a new JamfProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJamfProfileResponseWithDefaults() *JamfProfileResponse {
	this := JamfProfileResponse{}
	var thirdPartyDiscoverySync bool = false
	this.ThirdPartyDiscoverySync = *utils.NewNullableBool(&thirdPartyDiscoverySync)
	return &this
}

// GetId returns the Id field value
func (o *JamfProfileResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JamfProfileResponse) SetId(v string) {
	o.Id = v
}

// GetAuthorizationLevels returns the AuthorizationLevels field value
func (o *JamfProfileResponse) GetAuthorizationLevels() CertificateProfileAuthorizationLevels {
	if o == nil {
		var ret CertificateProfileAuthorizationLevels
		return ret
	}

	return o.AuthorizationLevels
}

// GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationLevels, true
}

// SetAuthorizationLevels sets field value
func (o *JamfProfileResponse) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels) {
	o.AuthorizationLevels = v
}

// GetCaps returns the Caps field value
func (o *JamfProfileResponse) GetCaps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Caps
}

// GetCapsOk returns a tuple with the Caps field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetCapsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Caps, true
}

// SetCaps sets field value
func (o *JamfProfileResponse) SetCaps(v []string) {
	o.Caps = v
}

// GetCertificateTemplate returns the CertificateTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetCertificateTemplate() CertificateTemplate {
	if o == nil || utils.IsNil(o.CertificateTemplate.Get()) {
		var ret CertificateTemplate
		return ret
	}
	return *o.CertificateTemplate.Get()
}

// GetCertificateTemplateOk returns a tuple with the CertificateTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetCertificateTemplateOk() (*CertificateTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateTemplate.Get(), o.CertificateTemplate.IsSet()
}

// HasCertificateTemplate returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasCertificateTemplate() bool {
	if o != nil && o.CertificateTemplate.IsSet() {
		return true
	}

	return false
}

// SetCertificateTemplate gets a reference to the given NullableCertificateTemplate and assigns it to the CertificateTemplate field.
func (o *JamfProfileResponse) SetCertificateTemplate(v CertificateTemplate) {
	o.CertificateTemplate.Set(&v)
}

// SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil
func (o *JamfProfileResponse) SetCertificateTemplateNil() {
	o.CertificateTemplate.Set(nil)
}

// UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
func (o *JamfProfileResponse) UnsetCertificateTemplate() {
	o.CertificateTemplate.Unset()
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetConstraints() CertificateRequestConstraints {
	if o == nil || utils.IsNil(o.Constraints.Get()) {
		var ret CertificateRequestConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetConstraintsOk() (*CertificateRequestConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableCertificateRequestConstraints and assigns it to the Constraints field.
func (o *JamfProfileResponse) SetConstraints(v CertificateRequestConstraints) {
	o.Constraints.Set(&v)
}

// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *JamfProfileResponse) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *JamfProfileResponse) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetCryptoPolicy returns the CryptoPolicy field value
func (o *JamfProfileResponse) GetCryptoPolicy() ManagedCertificateProfileCryptoPolicy {
	if o == nil {
		var ret ManagedCertificateProfileCryptoPolicy
		return ret
	}

	return o.CryptoPolicy
}

// GetCryptoPolicyOk returns a tuple with the CryptoPolicy field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetCryptoPolicyOk() (*ManagedCertificateProfileCryptoPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoPolicy, true
}

// SetCryptoPolicy sets field value
func (o *JamfProfileResponse) SetCryptoPolicy(v ManagedCertificateProfileCryptoPolicy) {
	o.CryptoPolicy = v
}

// GetCsrDataMapping returns the CsrDataMapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetCsrDataMapping() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.CsrDataMapping
}

// GetCsrDataMappingOk returns a tuple with the CsrDataMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetCsrDataMappingOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.CsrDataMapping) {
		return nil, false
	}
	return &o.CsrDataMapping, true
}

// HasCsrDataMapping returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasCsrDataMapping() bool {
	if o != nil && !utils.IsNil(o.CsrDataMapping) {
		return true
	}

	return false
}

// SetCsrDataMapping gets a reference to the given map[string]string and assigns it to the CsrDataMapping field.
func (o *JamfProfileResponse) SetCsrDataMapping(v map[string]string) {
	o.CsrDataMapping = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *JamfProfileResponse) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetDeviceIdField returns the DeviceIdField field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetDeviceIdField() string {
	if o == nil || utils.IsNil(o.DeviceIdField.Get()) {
		var ret string
		return ret
	}
	return *o.DeviceIdField.Get()
}

// GetDeviceIdFieldOk returns a tuple with the DeviceIdField field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetDeviceIdFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceIdField.Get(), o.DeviceIdField.IsSet()
}

// HasDeviceIdField returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasDeviceIdField() bool {
	if o != nil && o.DeviceIdField.IsSet() {
		return true
	}

	return false
}

// SetDeviceIdField gets a reference to the given NullableString and assigns it to the DeviceIdField field.
func (o *JamfProfileResponse) SetDeviceIdField(v string) {
	o.DeviceIdField.Set(&v)
}

// SetDeviceIdFieldNil sets the value for DeviceIdField to be an explicit nil
func (o *JamfProfileResponse) SetDeviceIdFieldNil() {
	o.DeviceIdField.Set(nil)
}

// UnsetDeviceIdField ensures that no value is present for DeviceIdField, not even an explicit nil
func (o *JamfProfileResponse) UnsetDeviceIdField() {
	o.DeviceIdField.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *JamfProfileResponse) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetDsFlow returns the DsFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetDsFlow() []DataSourceFlowEntry {
	if o == nil {
		var ret []DataSourceFlowEntry
		return ret
	}
	return o.DsFlow
}

// GetDsFlowOk returns a tuple with the DsFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetDsFlowOk() ([]DataSourceFlowEntry, bool) {
	if o == nil || utils.IsNil(o.DsFlow) {
		return nil, false
	}
	return o.DsFlow, true
}

// HasDsFlow returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasDsFlow() bool {
	if o != nil && !utils.IsNil(o.DsFlow) {
		return true
	}

	return false
}

// SetDsFlow gets a reference to the given []DataSourceFlowEntry and assigns it to the DsFlow field.
func (o *JamfProfileResponse) SetDsFlow(v []DataSourceFlowEntry) {
	o.DsFlow = v
}

// GetEnabled returns the Enabled field value
func (o *JamfProfileResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *JamfProfileResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value
func (o *JamfProfileResponse) GetEncryptionAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetEncryptionAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptionAlgorithm, true
}

// SetEncryptionAlgorithm sets field value
func (o *JamfProfileResponse) SetEncryptionAlgorithm(v string) {
	o.EncryptionAlgorithm = v
}

// GetGradingPolicies returns the GradingPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetGradingPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GradingPolicies
}

// GetGradingPoliciesOk returns a tuple with the GradingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetGradingPoliciesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.GradingPolicies) {
		return nil, false
	}
	return o.GradingPolicies, true
}

// HasGradingPolicies returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasGradingPolicies() bool {
	if o != nil && !utils.IsNil(o.GradingPolicies) {
		return true
	}

	return false
}

// SetGradingPolicies gets a reference to the given []string and assigns it to the GradingPolicies field.
func (o *JamfProfileResponse) SetGradingPolicies(v []string) {
	o.GradingPolicies = v
}

// GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy {
	if o == nil || utils.IsNil(o.MaxCertificatePerHolderPolicy.Get()) {
		var ret MaxCertificatePerHolderPolicy
		return ret
	}
	return *o.MaxCertificatePerHolderPolicy.Get()
}

// GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCertificatePerHolderPolicy.Get(), o.MaxCertificatePerHolderPolicy.IsSet()
}

// HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasMaxCertificatePerHolderPolicy() bool {
	if o != nil && o.MaxCertificatePerHolderPolicy.IsSet() {
		return true
	}

	return false
}

// SetMaxCertificatePerHolderPolicy gets a reference to the given NullableMaxCertificatePerHolderPolicy and assigns it to the MaxCertificatePerHolderPolicy field.
func (o *JamfProfileResponse) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy) {
	o.MaxCertificatePerHolderPolicy.Set(&v)
}

// SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil
func (o *JamfProfileResponse) SetMaxCertificatePerHolderPolicyNil() {
	o.MaxCertificatePerHolderPolicy.Set(nil)
}

// UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
func (o *JamfProfileResponse) UnsetMaxCertificatePerHolderPolicy() {
	o.MaxCertificatePerHolderPolicy.Unset()
}

// GetMode returns the Mode field value
func (o *JamfProfileResponse) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *JamfProfileResponse) SetMode(v string) {
	o.Mode = v
}

// GetModule returns the Module field value
func (o *JamfProfileResponse) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *JamfProfileResponse) SetModule(v string) {
	o.Module = v
}

// GetName returns the Name field value
func (o *JamfProfileResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JamfProfileResponse) SetName(v string) {
	o.Name = v
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetPasswordPolicy() string {
	if o == nil || utils.IsNil(o.PasswordPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordPolicy.Get()
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetPasswordPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordPolicy.Get(), o.PasswordPolicy.IsSet()
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasPasswordPolicy() bool {
	if o != nil && o.PasswordPolicy.IsSet() {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given NullableString and assigns it to the PasswordPolicy field.
func (o *JamfProfileResponse) SetPasswordPolicy(v string) {
	o.PasswordPolicy.Set(&v)
}

// SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil
func (o *JamfProfileResponse) SetPasswordPolicyNil() {
	o.PasswordPolicy.Set(nil)
}

// UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
func (o *JamfProfileResponse) UnsetPasswordPolicy() {
	o.PasswordPolicy.Unset()
}

// GetPkiConnector returns the PkiConnector field value
func (o *JamfProfileResponse) GetPkiConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PkiConnector
}

// GetPkiConnectorOk returns a tuple with the PkiConnector field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetPkiConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiConnector, true
}

// SetPkiConnector sets field value
func (o *JamfProfileResponse) SetPkiConnector(v string) {
	o.PkiConnector = v
}

// GetPostPKIOperation returns the PostPKIOperation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetPostPKIOperation() bool {
	if o == nil || utils.IsNil(o.PostPKIOperation.Get()) {
		var ret bool
		return ret
	}
	return *o.PostPKIOperation.Get()
}

// GetPostPKIOperationOk returns a tuple with the PostPKIOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetPostPKIOperationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostPKIOperation.Get(), o.PostPKIOperation.IsSet()
}

// HasPostPKIOperation returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasPostPKIOperation() bool {
	if o != nil && o.PostPKIOperation.IsSet() {
		return true
	}

	return false
}

// SetPostPKIOperation gets a reference to the given NullableBool and assigns it to the PostPKIOperation field.
func (o *JamfProfileResponse) SetPostPKIOperation(v bool) {
	o.PostPKIOperation.Set(&v)
}

// SetPostPKIOperationNil sets the value for PostPKIOperation to be an explicit nil
func (o *JamfProfileResponse) SetPostPKIOperationNil() {
	o.PostPKIOperation.Set(nil)
}

// UnsetPostPKIOperation ensures that no value is present for PostPKIOperation, not even an explicit nil
func (o *JamfProfileResponse) UnsetPostPKIOperation() {
	o.PostPKIOperation.Unset()
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetRenewalPeriod() string {
	if o == nil || utils.IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *JamfProfileResponse) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}

// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *JamfProfileResponse) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *JamfProfileResponse) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetRequestsPolicy returns the RequestsPolicy field value
func (o *JamfProfileResponse) GetRequestsPolicy() RequestsPolicy {
	if o == nil {
		var ret RequestsPolicy
		return ret
	}

	return o.RequestsPolicy
}

// GetRequestsPolicyOk returns a tuple with the RequestsPolicy field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetRequestsPolicyOk() (*RequestsPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestsPolicy, true
}

// SetRequestsPolicy sets field value
func (o *JamfProfileResponse) SetRequestsPolicy(v RequestsPolicy) {
	o.RequestsPolicy = v
}

// GetScepRA returns the ScepRA field value
func (o *JamfProfileResponse) GetScepRA() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScepRA
}

// GetScepRAOk returns a tuple with the ScepRA field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetScepRAOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScepRA, true
}

// SetScepRA sets field value
func (o *JamfProfileResponse) SetScepRA(v string) {
	o.ScepRA = v
}

// GetSelfPermissions returns the SelfPermissions field value
func (o *JamfProfileResponse) GetSelfPermissions() CertificateProfileSelfPermissions {
	if o == nil {
		var ret CertificateProfileSelfPermissions
		return ret
	}

	return o.SelfPermissions
}

// GetSelfPermissionsOk returns a tuple with the SelfPermissions field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfPermissions, true
}

// SetSelfPermissions sets field value
func (o *JamfProfileResponse) SetSelfPermissions(v CertificateProfileSelfPermissions) {
	o.SelfPermissions = v
}

// GetThirdPartyConnector returns the ThirdPartyConnector field value
func (o *JamfProfileResponse) GetThirdPartyConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThirdPartyConnector
}

// GetThirdPartyConnectorOk returns a tuple with the ThirdPartyConnector field value
// and a boolean to check if the value has been set.
func (o *JamfProfileResponse) GetThirdPartyConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThirdPartyConnector, true
}

// SetThirdPartyConnector sets field value
func (o *JamfProfileResponse) SetThirdPartyConnector(v string) {
	o.ThirdPartyConnector = v
}

// GetThirdPartyDiscoverySync returns the ThirdPartyDiscoverySync field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetThirdPartyDiscoverySync() bool {
	if o == nil || utils.IsNil(o.ThirdPartyDiscoverySync.Get()) {
		var ret bool
		return ret
	}
	return *o.ThirdPartyDiscoverySync.Get()
}

// GetThirdPartyDiscoverySyncOk returns a tuple with the ThirdPartyDiscoverySync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetThirdPartyDiscoverySyncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThirdPartyDiscoverySync.Get(), o.ThirdPartyDiscoverySync.IsSet()
}

// HasThirdPartyDiscoverySync returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasThirdPartyDiscoverySync() bool {
	if o != nil && o.ThirdPartyDiscoverySync.IsSet() {
		return true
	}

	return false
}

// SetThirdPartyDiscoverySync gets a reference to the given NullableBool and assigns it to the ThirdPartyDiscoverySync field.
func (o *JamfProfileResponse) SetThirdPartyDiscoverySync(v bool) {
	o.ThirdPartyDiscoverySync.Set(&v)
}

// SetThirdPartyDiscoverySyncNil sets the value for ThirdPartyDiscoverySync to be an explicit nil
func (o *JamfProfileResponse) SetThirdPartyDiscoverySyncNil() {
	o.ThirdPartyDiscoverySync.Set(nil)
}

// UnsetThirdPartyDiscoverySync ensures that no value is present for ThirdPartyDiscoverySync, not even an explicit nil
func (o *JamfProfileResponse) UnsetThirdPartyDiscoverySync() {
	o.ThirdPartyDiscoverySync.Unset()
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JamfProfileResponse) GetTriggers() CertificateProfileTriggers {
	if o == nil || utils.IsNil(o.Triggers.Get()) {
		var ret CertificateProfileTriggers
		return ret
	}
	return *o.Triggers.Get()
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JamfProfileResponse) GetTriggersOk() (*CertificateProfileTriggers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers.Get(), o.Triggers.IsSet()
}

// HasTriggers returns a boolean if a field has been set.
func (o *JamfProfileResponse) HasTriggers() bool {
	if o != nil && o.Triggers.IsSet() {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given NullableCertificateProfileTriggers and assigns it to the Triggers field.
func (o *JamfProfileResponse) SetTriggers(v CertificateProfileTriggers) {
	o.Triggers.Set(&v)
}

// SetTriggersNil sets the value for Triggers to be an explicit nil
func (o *JamfProfileResponse) SetTriggersNil() {
	o.Triggers.Set(nil)
}

// UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
func (o *JamfProfileResponse) UnsetTriggers() {
	o.Triggers.Unset()
}

func (o JamfProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JamfProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["authorizationLevels"] = o.AuthorizationLevels
	toSerialize["caps"] = o.Caps
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
	if o.DeviceIdField.IsSet() {
		toSerialize["deviceIdField"] = o.DeviceIdField.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.DsFlow != nil {
		toSerialize["dsFlow"] = o.DsFlow
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["encryptionAlgorithm"] = o.EncryptionAlgorithm
	if o.GradingPolicies != nil {
		toSerialize["gradingPolicies"] = o.GradingPolicies
	}
	if o.MaxCertificatePerHolderPolicy.IsSet() {
		toSerialize["maxCertificatePerHolderPolicy"] = o.MaxCertificatePerHolderPolicy.Get()
	}
	toSerialize["mode"] = o.Mode
	toSerialize["module"] = o.Module
	toSerialize["name"] = o.Name
	if o.PasswordPolicy.IsSet() {
		toSerialize["passwordPolicy"] = o.PasswordPolicy.Get()
	}
	toSerialize["pkiConnector"] = o.PkiConnector
	if o.PostPKIOperation.IsSet() {
		toSerialize["postPKIOperation"] = o.PostPKIOperation.Get()
	}
	if o.RenewalPeriod.IsSet() {
		toSerialize["renewalPeriod"] = o.RenewalPeriod.Get()
	}
	toSerialize["requestsPolicy"] = o.RequestsPolicy
	toSerialize["scepRA"] = o.ScepRA
	toSerialize["selfPermissions"] = o.SelfPermissions
	toSerialize["thirdPartyConnector"] = o.ThirdPartyConnector
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

func (o *JamfProfileResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"authorizationLevels",
		"caps",
		"cryptoPolicy",
		"enabled",
		"encryptionAlgorithm",
		"mode",
		"module",
		"name",
		"pkiConnector",
		"requestsPolicy",
		"scepRA",
		"selfPermissions",
		"thirdPartyConnector",
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

	varJamfProfileResponse := _JamfProfileResponse{}

	err = json.Unmarshal(data, &varJamfProfileResponse)

	if err != nil {
		return err
	}

	*o = JamfProfileResponse(varJamfProfileResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "authorizationLevels")
		delete(additionalProperties, "caps")
		delete(additionalProperties, "certificateTemplate")
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "cryptoPolicy")
		delete(additionalProperties, "csrDataMapping")
		delete(additionalProperties, "description")
		delete(additionalProperties, "deviceIdField")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "dsFlow")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "encryptionAlgorithm")
		delete(additionalProperties, "gradingPolicies")
		delete(additionalProperties, "maxCertificatePerHolderPolicy")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "passwordPolicy")
		delete(additionalProperties, "pkiConnector")
		delete(additionalProperties, "postPKIOperation")
		delete(additionalProperties, "renewalPeriod")
		delete(additionalProperties, "requestsPolicy")
		delete(additionalProperties, "scepRA")
		delete(additionalProperties, "selfPermissions")
		delete(additionalProperties, "thirdPartyConnector")
		delete(additionalProperties, "thirdPartyDiscoverySync")
		delete(additionalProperties, "triggers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJamfProfileResponse struct {
	value *JamfProfileResponse
	isSet bool
}

func (v NullableJamfProfileResponse) Get() *JamfProfileResponse {
	return v.value
}

func (v *NullableJamfProfileResponse) Set(val *JamfProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJamfProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJamfProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJamfProfileResponse(val *JamfProfileResponse) *NullableJamfProfileResponse {
	return &NullableJamfProfileResponse{value: val, isSet: true}
}

func (v NullableJamfProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJamfProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
