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

// checks if the AcmeExternalProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcmeExternalProfileResponse{}

// AcmeExternalProfileResponse struct for AcmeExternalProfileResponse
type AcmeExternalProfileResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	Module string `json:"module"`
	Name string `json:"name"`
	DisplayName []LocalizedString `json:"displayName,omitempty"`
	Description []LocalizedString `json:"description,omitempty"`
	Enabled bool `json:"enabled"`
	Constraints NullableCertificateRequestConstraints `json:"constraints,omitempty"`
	AuthorizationMethods []string `json:"authorizationMethods"`
	PkiConnector string `json:"pkiConnector"`
	AcmeUrl *string `json:"acmeUrl,omitempty"`
	RequireEAB bool `json:"requireEAB"`
	MaxCertificatePerHolderPolicy NullableMaxCertificatePerHolderPolicy `json:"maxCertificatePerHolderPolicy,omitempty"`
	AuthorizedCas []string `json:"authorizedCas"`
	RenewalPeriod NullableString `json:"renewalPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
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

type _AcmeExternalProfileResponse AcmeExternalProfileResponse

// NewAcmeExternalProfileResponse instantiates a new AcmeExternalProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcmeExternalProfileResponse(id string, module string, name string, enabled bool, authorizationMethods []string, pkiConnector string, requireEAB bool, authorizedCas []string, authorizationLevels CertificateProfileAuthorizationLevels, requestsPolicy RequestsPolicy, selfPermissions CertificateProfileSelfPermissions, cryptoPolicy CertificateProfileCryptoPolicy) *AcmeExternalProfileResponse {
	this := AcmeExternalProfileResponse{}
	this.Id = id
	this.Module = module
	this.Name = name
	this.Enabled = enabled
	this.AuthorizationMethods = authorizationMethods
	this.PkiConnector = pkiConnector
	this.RequireEAB = requireEAB
	this.AuthorizedCas = authorizedCas
	this.AuthorizationLevels = authorizationLevels
	this.RequestsPolicy = requestsPolicy
	this.SelfPermissions = selfPermissions
	this.CryptoPolicy = cryptoPolicy
	return &this
}

// NewAcmeExternalProfileResponseWithDefaults instantiates a new AcmeExternalProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcmeExternalProfileResponseWithDefaults() *AcmeExternalProfileResponse {
	this := AcmeExternalProfileResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AcmeExternalProfileResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfileResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AcmeExternalProfileResponse) SetId(v string) {
	o.Id = v
}

// GetModule returns the Module field value
func (o *AcmeExternalProfileResponse) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfileResponse) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *AcmeExternalProfileResponse) SetModule(v string) {
	o.Module = v
}

// GetName returns the Name field value
func (o *AcmeExternalProfileResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfileResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AcmeExternalProfileResponse) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfileResponse) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfileResponse) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AcmeExternalProfileResponse) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *AcmeExternalProfileResponse) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfileResponse) GetDescription() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfileResponse) GetDescriptionOk() ([]LocalizedString, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AcmeExternalProfileResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []LocalizedString and assigns it to the Description field.
func (o *AcmeExternalProfileResponse) SetDescription(v []LocalizedString) {
	o.Description = v
}

// GetEnabled returns the Enabled field value
func (o *AcmeExternalProfileResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfileResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AcmeExternalProfileResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfileResponse) GetConstraints() CertificateRequestConstraints {
	if o == nil || IsNil(o.Constraints.Get()) {
		var ret CertificateRequestConstraints
		return ret
	}
	return *o.Constraints.Get()
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfileResponse) GetConstraintsOk() (*CertificateRequestConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints.Get(), o.Constraints.IsSet()
}

// HasConstraints returns a boolean if a field has been set.
func (o *AcmeExternalProfileResponse) HasConstraints() bool {
	if o != nil && o.Constraints.IsSet() {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given NullableCertificateRequestConstraints and assigns it to the Constraints field.
func (o *AcmeExternalProfileResponse) SetConstraints(v CertificateRequestConstraints) {
	o.Constraints.Set(&v)
}
// SetConstraintsNil sets the value for Constraints to be an explicit nil
func (o *AcmeExternalProfileResponse) SetConstraintsNil() {
	o.Constraints.Set(nil)
}

// UnsetConstraints ensures that no value is present for Constraints, not even an explicit nil
func (o *AcmeExternalProfileResponse) UnsetConstraints() {
	o.Constraints.Unset()
}

// GetAuthorizationMethods returns the AuthorizationMethods field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *AcmeExternalProfileResponse) GetAuthorizationMethods() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthorizationMethods
}

// GetAuthorizationMethodsOk returns a tuple with the AuthorizationMethods field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfileResponse) GetAuthorizationMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthorizationMethods) {
		return nil, false
	}
	return o.AuthorizationMethods, true
}

// SetAuthorizationMethods sets field value
func (o *AcmeExternalProfileResponse) SetAuthorizationMethods(v []string) {
	o.AuthorizationMethods = v
}

// GetPkiConnector returns the PkiConnector field value
func (o *AcmeExternalProfileResponse) GetPkiConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PkiConnector
}

// GetPkiConnectorOk returns a tuple with the PkiConnector field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfileResponse) GetPkiConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiConnector, true
}

// SetPkiConnector sets field value
func (o *AcmeExternalProfileResponse) SetPkiConnector(v string) {
	o.PkiConnector = v
}

// GetAcmeUrl returns the AcmeUrl field value if set, zero value otherwise.
func (o *AcmeExternalProfileResponse) GetAcmeUrl() string {
	if o == nil || IsNil(o.AcmeUrl) {
		var ret string
		return ret
	}
	return *o.AcmeUrl
}

// GetAcmeUrlOk returns a tuple with the AcmeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfileResponse) GetAcmeUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AcmeUrl) {
		return nil, false
	}
	return o.AcmeUrl, true
}

// HasAcmeUrl returns a boolean if a field has been set.
func (o *AcmeExternalProfileResponse) HasAcmeUrl() bool {
	if o != nil && !IsNil(o.AcmeUrl) {
		return true
	}

	return false
}

// SetAcmeUrl gets a reference to the given string and assigns it to the AcmeUrl field.
func (o *AcmeExternalProfileResponse) SetAcmeUrl(v string) {
	o.AcmeUrl = &v
}

// GetRequireEAB returns the RequireEAB field value
func (o *AcmeExternalProfileResponse) GetRequireEAB() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireEAB
}

// GetRequireEABOk returns a tuple with the RequireEAB field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfileResponse) GetRequireEABOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireEAB, true
}

// SetRequireEAB sets field value
func (o *AcmeExternalProfileResponse) SetRequireEAB(v bool) {
	o.RequireEAB = v
}

// GetMaxCertificatePerHolderPolicy returns the MaxCertificatePerHolderPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfileResponse) GetMaxCertificatePerHolderPolicy() MaxCertificatePerHolderPolicy {
	if o == nil || IsNil(o.MaxCertificatePerHolderPolicy.Get()) {
		var ret MaxCertificatePerHolderPolicy
		return ret
	}
	return *o.MaxCertificatePerHolderPolicy.Get()
}

// GetMaxCertificatePerHolderPolicyOk returns a tuple with the MaxCertificatePerHolderPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfileResponse) GetMaxCertificatePerHolderPolicyOk() (*MaxCertificatePerHolderPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxCertificatePerHolderPolicy.Get(), o.MaxCertificatePerHolderPolicy.IsSet()
}

// HasMaxCertificatePerHolderPolicy returns a boolean if a field has been set.
func (o *AcmeExternalProfileResponse) HasMaxCertificatePerHolderPolicy() bool {
	if o != nil && o.MaxCertificatePerHolderPolicy.IsSet() {
		return true
	}

	return false
}

// SetMaxCertificatePerHolderPolicy gets a reference to the given NullableMaxCertificatePerHolderPolicy and assigns it to the MaxCertificatePerHolderPolicy field.
func (o *AcmeExternalProfileResponse) SetMaxCertificatePerHolderPolicy(v MaxCertificatePerHolderPolicy) {
	o.MaxCertificatePerHolderPolicy.Set(&v)
}
// SetMaxCertificatePerHolderPolicyNil sets the value for MaxCertificatePerHolderPolicy to be an explicit nil
func (o *AcmeExternalProfileResponse) SetMaxCertificatePerHolderPolicyNil() {
	o.MaxCertificatePerHolderPolicy.Set(nil)
}

// UnsetMaxCertificatePerHolderPolicy ensures that no value is present for MaxCertificatePerHolderPolicy, not even an explicit nil
func (o *AcmeExternalProfileResponse) UnsetMaxCertificatePerHolderPolicy() {
	o.MaxCertificatePerHolderPolicy.Unset()
}

// GetAuthorizedCas returns the AuthorizedCas field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *AcmeExternalProfileResponse) GetAuthorizedCas() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthorizedCas
}

// GetAuthorizedCasOk returns a tuple with the AuthorizedCas field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfileResponse) GetAuthorizedCasOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthorizedCas) {
		return nil, false
	}
	return o.AuthorizedCas, true
}

// SetAuthorizedCas sets field value
func (o *AcmeExternalProfileResponse) SetAuthorizedCas(v []string) {
	o.AuthorizedCas = v
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfileResponse) GetRenewalPeriod() string {
	if o == nil || IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfileResponse) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *AcmeExternalProfileResponse) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *AcmeExternalProfileResponse) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}
// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *AcmeExternalProfileResponse) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *AcmeExternalProfileResponse) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetAuthorizationLevels returns the AuthorizationLevels field value
func (o *AcmeExternalProfileResponse) GetAuthorizationLevels() CertificateProfileAuthorizationLevels {
	if o == nil {
		var ret CertificateProfileAuthorizationLevels
		return ret
	}

	return o.AuthorizationLevels
}

// GetAuthorizationLevelsOk returns a tuple with the AuthorizationLevels field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfileResponse) GetAuthorizationLevelsOk() (*CertificateProfileAuthorizationLevels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorizationLevels, true
}

// SetAuthorizationLevels sets field value
func (o *AcmeExternalProfileResponse) SetAuthorizationLevels(v CertificateProfileAuthorizationLevels) {
	o.AuthorizationLevels = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfileResponse) GetTriggers() CertificateProfileTriggers {
	if o == nil || IsNil(o.Triggers.Get()) {
		var ret CertificateProfileTriggers
		return ret
	}
	return *o.Triggers.Get()
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfileResponse) GetTriggersOk() (*CertificateProfileTriggers, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers.Get(), o.Triggers.IsSet()
}

// HasTriggers returns a boolean if a field has been set.
func (o *AcmeExternalProfileResponse) HasTriggers() bool {
	if o != nil && o.Triggers.IsSet() {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given NullableCertificateProfileTriggers and assigns it to the Triggers field.
func (o *AcmeExternalProfileResponse) SetTriggers(v CertificateProfileTriggers) {
	o.Triggers.Set(&v)
}
// SetTriggersNil sets the value for Triggers to be an explicit nil
func (o *AcmeExternalProfileResponse) SetTriggersNil() {
	o.Triggers.Set(nil)
}

// UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
func (o *AcmeExternalProfileResponse) UnsetTriggers() {
	o.Triggers.Unset()
}

// GetRequestsPolicy returns the RequestsPolicy field value
func (o *AcmeExternalProfileResponse) GetRequestsPolicy() RequestsPolicy {
	if o == nil {
		var ret RequestsPolicy
		return ret
	}

	return o.RequestsPolicy
}

// GetRequestsPolicyOk returns a tuple with the RequestsPolicy field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfileResponse) GetRequestsPolicyOk() (*RequestsPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestsPolicy, true
}

// SetRequestsPolicy sets field value
func (o *AcmeExternalProfileResponse) SetRequestsPolicy(v RequestsPolicy) {
	o.RequestsPolicy = v
}

// GetSelfPermissions returns the SelfPermissions field value
func (o *AcmeExternalProfileResponse) GetSelfPermissions() CertificateProfileSelfPermissions {
	if o == nil {
		var ret CertificateProfileSelfPermissions
		return ret
	}

	return o.SelfPermissions
}

// GetSelfPermissionsOk returns a tuple with the SelfPermissions field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfileResponse) GetSelfPermissionsOk() (*CertificateProfileSelfPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfPermissions, true
}

// SetSelfPermissions sets field value
func (o *AcmeExternalProfileResponse) SetSelfPermissions(v CertificateProfileSelfPermissions) {
	o.SelfPermissions = v
}

// GetCertificateTemplate returns the CertificateTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfileResponse) GetCertificateTemplate() CertificateTemplate {
	if o == nil || IsNil(o.CertificateTemplate.Get()) {
		var ret CertificateTemplate
		return ret
	}
	return *o.CertificateTemplate.Get()
}

// GetCertificateTemplateOk returns a tuple with the CertificateTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfileResponse) GetCertificateTemplateOk() (*CertificateTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateTemplate.Get(), o.CertificateTemplate.IsSet()
}

// HasCertificateTemplate returns a boolean if a field has been set.
func (o *AcmeExternalProfileResponse) HasCertificateTemplate() bool {
	if o != nil && o.CertificateTemplate.IsSet() {
		return true
	}

	return false
}

// SetCertificateTemplate gets a reference to the given NullableCertificateTemplate and assigns it to the CertificateTemplate field.
func (o *AcmeExternalProfileResponse) SetCertificateTemplate(v CertificateTemplate) {
	o.CertificateTemplate.Set(&v)
}
// SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil
func (o *AcmeExternalProfileResponse) SetCertificateTemplateNil() {
	o.CertificateTemplate.Set(nil)
}

// UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
func (o *AcmeExternalProfileResponse) UnsetCertificateTemplate() {
	o.CertificateTemplate.Unset()
}

// GetCryptoPolicy returns the CryptoPolicy field value
func (o *AcmeExternalProfileResponse) GetCryptoPolicy() CertificateProfileCryptoPolicy {
	if o == nil {
		var ret CertificateProfileCryptoPolicy
		return ret
	}

	return o.CryptoPolicy
}

// GetCryptoPolicyOk returns a tuple with the CryptoPolicy field value
// and a boolean to check if the value has been set.
func (o *AcmeExternalProfileResponse) GetCryptoPolicyOk() (*CertificateProfileCryptoPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoPolicy, true
}

// SetCryptoPolicy sets field value
func (o *AcmeExternalProfileResponse) SetCryptoPolicy(v CertificateProfileCryptoPolicy) {
	o.CryptoPolicy = v
}

// GetGradingPolicies returns the GradingPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfileResponse) GetGradingPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GradingPolicies
}

// GetGradingPoliciesOk returns a tuple with the GradingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfileResponse) GetGradingPoliciesOk() ([]string, bool) {
	if o == nil || IsNil(o.GradingPolicies) {
		return nil, false
	}
	return o.GradingPolicies, true
}

// HasGradingPolicies returns a boolean if a field has been set.
func (o *AcmeExternalProfileResponse) HasGradingPolicies() bool {
	if o != nil && !IsNil(o.GradingPolicies) {
		return true
	}

	return false
}

// SetGradingPolicies gets a reference to the given []string and assigns it to the GradingPolicies field.
func (o *AcmeExternalProfileResponse) SetGradingPolicies(v []string) {
	o.GradingPolicies = v
}

// GetDsFlow returns the DsFlow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeExternalProfileResponse) GetDsFlow() []DataSourceFlowEntry {
	if o == nil {
		var ret []DataSourceFlowEntry
		return ret
	}
	return o.DsFlow
}

// GetDsFlowOk returns a tuple with the DsFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeExternalProfileResponse) GetDsFlowOk() ([]DataSourceFlowEntry, bool) {
	if o == nil || IsNil(o.DsFlow) {
		return nil, false
	}
	return o.DsFlow, true
}

// HasDsFlow returns a boolean if a field has been set.
func (o *AcmeExternalProfileResponse) HasDsFlow() bool {
	if o != nil && !IsNil(o.DsFlow) {
		return true
	}

	return false
}

// SetDsFlow gets a reference to the given []DataSourceFlowEntry and assigns it to the DsFlow field.
func (o *AcmeExternalProfileResponse) SetDsFlow(v []DataSourceFlowEntry) {
	o.DsFlow = v
}

func (o AcmeExternalProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcmeExternalProfileResponse) ToMap() (map[string]interface{}, error) {
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
	if o.Constraints.IsSet() {
		toSerialize["constraints"] = o.Constraints.Get()
	}
	if o.AuthorizationMethods != nil {
		toSerialize["authorizationMethods"] = o.AuthorizationMethods
	}
	toSerialize["pkiConnector"] = o.PkiConnector
	if !IsNil(o.AcmeUrl) {
		toSerialize["acmeUrl"] = o.AcmeUrl
	}
	toSerialize["requireEAB"] = o.RequireEAB
	if o.MaxCertificatePerHolderPolicy.IsSet() {
		toSerialize["maxCertificatePerHolderPolicy"] = o.MaxCertificatePerHolderPolicy.Get()
	}
	if o.AuthorizedCas != nil {
		toSerialize["authorizedCas"] = o.AuthorizedCas
	}
	if o.RenewalPeriod.IsSet() {
		toSerialize["renewalPeriod"] = o.RenewalPeriod.Get()
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

func (o *AcmeExternalProfileResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"module",
		"name",
		"enabled",
		"authorizationMethods",
		"pkiConnector",
		"requireEAB",
		"authorizedCas",
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

	varAcmeExternalProfileResponse := _AcmeExternalProfileResponse{}

	err = json.Unmarshal(data, &varAcmeExternalProfileResponse)

	if err != nil {
		return err
	}

	*o = AcmeExternalProfileResponse(varAcmeExternalProfileResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "module")
		delete(additionalProperties, "name")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "description")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "constraints")
		delete(additionalProperties, "authorizationMethods")
		delete(additionalProperties, "pkiConnector")
		delete(additionalProperties, "acmeUrl")
		delete(additionalProperties, "requireEAB")
		delete(additionalProperties, "maxCertificatePerHolderPolicy")
		delete(additionalProperties, "authorizedCas")
		delete(additionalProperties, "renewalPeriod")
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

type NullableAcmeExternalProfileResponse struct {
	value *AcmeExternalProfileResponse
	isSet bool
}

func (v NullableAcmeExternalProfileResponse) Get() *AcmeExternalProfileResponse {
	return v.value
}

func (v *NullableAcmeExternalProfileResponse) Set(val *AcmeExternalProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcmeExternalProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcmeExternalProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcmeExternalProfileResponse(val *AcmeExternalProfileResponse) *NullableAcmeExternalProfileResponse {
	return &NullableAcmeExternalProfileResponse{value: val, isSet: true}
}

func (v NullableAcmeExternalProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcmeExternalProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


