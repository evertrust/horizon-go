/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the CertificateSearchResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CertificateSearchResult{}

// CertificateSearchResult struct for CertificateSearchResult
type CertificateSearchResult struct {
	Id                utils.NullableString  `json:"_id,omitempty"`
	Certificate       utils.NullableString  `json:"certificate,omitempty"`
	ContactEmail      utils.NullableString  `json:"contactEmail,omitempty"`
	DiscoveredTrusted utils.NullableBool    `json:"discoveredTrusted,omitempty"`
	DiscoveryData     []HostDiscoveryData   `json:"discoveryData,omitempty"`
	DiscoveryInfo     []DiscoveryInfo       `json:"discoveryInfo,omitempty"`
	Dn                utils.NullableString  `json:"dn,omitempty"`
	Grades            []GradingPolicyResult `json:"grades,omitempty"`
	HolderId          utils.NullableString  `json:"holderId,omitempty"`
	Issuer            utils.NullableString  `json:"issuer,omitempty"`
	// One of `rsa-2048`, `rsa-3072`, `rsa-4096`, `rsa-8192`, `ec-secp256r1`, `ec-secp384r1`, `ec-secp521r1`, `ed-448`, `ed-25519`, `mldsa-44`, `mldsa-65`, `mldsa-87`, `slhdsa-sha2-128s`, `slhdsa-sha2-128f`, `slhdsa-sha2-192s`, `slhdsa-sha2-192f`, `slhdsa-sha2-256s`, `slhdsa-sha2-256f`, `slhdsa-sha2-128ssha256`, `slhdsa-sha2-128fsha256`, `slhdsa-sha2-192ssha512`, `slhdsa-sha2-192fsha512`, `slhdsa-sha2-256ssha512`, `slhdsa-sha2-256fsha512` or `<primary key type>+<alternate key type>`
	KeyType             utils.NullableString           `json:"keyType,omitempty" validate:"regexp=(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512)(\\\\\\\\+(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512))?"`
	Labels              []LabelData                    `json:"labels,omitempty"`
	Metadata            []CertificateMetadata          `json:"metadata,omitempty"`
	Module              utils.NullableString           `json:"module,omitempty"`
	NotAfter            utils.NullableInt64            `json:"notAfter,omitempty"`
	NotBefore           utils.NullableInt64            `json:"notBefore,omitempty"`
	Owner               utils.NullableString           `json:"owner,omitempty"`
	Permissions         NullableCertificatePermissions `json:"permissions,omitempty"`
	PrivateKey          NullableEscrowedPrivateKey     `json:"privateKey,omitempty"`
	Profile             utils.NullableString           `json:"profile,omitempty"`
	PublicKeyThumbprint utils.NullableString           `json:"publicKeyThumbprint,omitempty"`
	RevocationDate      utils.NullableInt64            `json:"revocationDate,omitempty"`
	// One of: `unspecified`, `keycompromise`, `cacompromise`, `affiliationchange`, `superseded`, `cessationofoperation`
	RevocationReason      utils.NullableString   `json:"revocationReason,omitempty"`
	SelfSigned            utils.NullableBool     `json:"selfSigned,omitempty"`
	Serial                utils.NullableString   `json:"serial,omitempty"`
	SigningAlgorithm      utils.NullableString   `json:"signingAlgorithm,omitempty"`
	SubjectAlternateNames []SubjectAlternateName `json:"subjectAlternateNames,omitempty"`
	Team                  utils.NullableString   `json:"team,omitempty"`
	ThirdPartyData        []ThirdPartyItem       `json:"thirdPartyData,omitempty"`
	Thumbprint            utils.NullableString   `json:"thumbprint,omitempty"`
	TriggerResults        []TriggerResult        `json:"triggerResults,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _CertificateSearchResult CertificateSearchResult

// NewCertificateSearchResult instantiates a new CertificateSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateSearchResult() *CertificateSearchResult {
	this := CertificateSearchResult{}
	return &this
}

// NewCertificateSearchResultWithDefaults instantiates a new CertificateSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateSearchResultWithDefaults() *CertificateSearchResult {
	this := CertificateSearchResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetId() string {
	if o == nil || utils.IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *CertificateSearchResult) SetId(v string) {
	o.Id.Set(&v)
}

// SetIdNil sets the value for Id to be an explicit nil
func (o *CertificateSearchResult) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CertificateSearchResult) UnsetId() {
	o.Id.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetCertificate() string {
	if o == nil || utils.IsNil(o.Certificate.Get()) {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *CertificateSearchResult) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *CertificateSearchResult) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *CertificateSearchResult) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetContactEmail() string {
	if o == nil || utils.IsNil(o.ContactEmail.Get()) {
		var ret string
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetContactEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableString and assigns it to the ContactEmail field.
func (o *CertificateSearchResult) SetContactEmail(v string) {
	o.ContactEmail.Set(&v)
}

// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *CertificateSearchResult) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *CertificateSearchResult) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetDiscoveredTrusted returns the DiscoveredTrusted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetDiscoveredTrusted() bool {
	if o == nil || utils.IsNil(o.DiscoveredTrusted.Get()) {
		var ret bool
		return ret
	}
	return *o.DiscoveredTrusted.Get()
}

// GetDiscoveredTrustedOk returns a tuple with the DiscoveredTrusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetDiscoveredTrustedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscoveredTrusted.Get(), o.DiscoveredTrusted.IsSet()
}

// HasDiscoveredTrusted returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasDiscoveredTrusted() bool {
	if o != nil && o.DiscoveredTrusted.IsSet() {
		return true
	}

	return false
}

// SetDiscoveredTrusted gets a reference to the given NullableBool and assigns it to the DiscoveredTrusted field.
func (o *CertificateSearchResult) SetDiscoveredTrusted(v bool) {
	o.DiscoveredTrusted.Set(&v)
}

// SetDiscoveredTrustedNil sets the value for DiscoveredTrusted to be an explicit nil
func (o *CertificateSearchResult) SetDiscoveredTrustedNil() {
	o.DiscoveredTrusted.Set(nil)
}

// UnsetDiscoveredTrusted ensures that no value is present for DiscoveredTrusted, not even an explicit nil
func (o *CertificateSearchResult) UnsetDiscoveredTrusted() {
	o.DiscoveredTrusted.Unset()
}

// GetDiscoveryData returns the DiscoveryData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetDiscoveryData() []HostDiscoveryData {
	if o == nil {
		var ret []HostDiscoveryData
		return ret
	}
	return o.DiscoveryData
}

// GetDiscoveryDataOk returns a tuple with the DiscoveryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetDiscoveryDataOk() ([]HostDiscoveryData, bool) {
	if o == nil || utils.IsNil(o.DiscoveryData) {
		return nil, false
	}
	return o.DiscoveryData, true
}

// HasDiscoveryData returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasDiscoveryData() bool {
	if o != nil && !utils.IsNil(o.DiscoveryData) {
		return true
	}

	return false
}

// SetDiscoveryData gets a reference to the given []HostDiscoveryData and assigns it to the DiscoveryData field.
func (o *CertificateSearchResult) SetDiscoveryData(v []HostDiscoveryData) {
	o.DiscoveryData = v
}

// GetDiscoveryInfo returns the DiscoveryInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetDiscoveryInfo() []DiscoveryInfo {
	if o == nil {
		var ret []DiscoveryInfo
		return ret
	}
	return o.DiscoveryInfo
}

// GetDiscoveryInfoOk returns a tuple with the DiscoveryInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetDiscoveryInfoOk() ([]DiscoveryInfo, bool) {
	if o == nil || utils.IsNil(o.DiscoveryInfo) {
		return nil, false
	}
	return o.DiscoveryInfo, true
}

// HasDiscoveryInfo returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasDiscoveryInfo() bool {
	if o != nil && !utils.IsNil(o.DiscoveryInfo) {
		return true
	}

	return false
}

// SetDiscoveryInfo gets a reference to the given []DiscoveryInfo and assigns it to the DiscoveryInfo field.
func (o *CertificateSearchResult) SetDiscoveryInfo(v []DiscoveryInfo) {
	o.DiscoveryInfo = v
}

// GetDn returns the Dn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetDn() string {
	if o == nil || utils.IsNil(o.Dn.Get()) {
		var ret string
		return ret
	}
	return *o.Dn.Get()
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dn.Get(), o.Dn.IsSet()
}

// HasDn returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasDn() bool {
	if o != nil && o.Dn.IsSet() {
		return true
	}

	return false
}

// SetDn gets a reference to the given NullableString and assigns it to the Dn field.
func (o *CertificateSearchResult) SetDn(v string) {
	o.Dn.Set(&v)
}

// SetDnNil sets the value for Dn to be an explicit nil
func (o *CertificateSearchResult) SetDnNil() {
	o.Dn.Set(nil)
}

// UnsetDn ensures that no value is present for Dn, not even an explicit nil
func (o *CertificateSearchResult) UnsetDn() {
	o.Dn.Unset()
}

// GetGrades returns the Grades field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetGrades() []GradingPolicyResult {
	if o == nil {
		var ret []GradingPolicyResult
		return ret
	}
	return o.Grades
}

// GetGradesOk returns a tuple with the Grades field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetGradesOk() ([]GradingPolicyResult, bool) {
	if o == nil || utils.IsNil(o.Grades) {
		return nil, false
	}
	return o.Grades, true
}

// HasGrades returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasGrades() bool {
	if o != nil && !utils.IsNil(o.Grades) {
		return true
	}

	return false
}

// SetGrades gets a reference to the given []GradingPolicyResult and assigns it to the Grades field.
func (o *CertificateSearchResult) SetGrades(v []GradingPolicyResult) {
	o.Grades = v
}

// GetHolderId returns the HolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetHolderId() string {
	if o == nil || utils.IsNil(o.HolderId.Get()) {
		var ret string
		return ret
	}
	return *o.HolderId.Get()
}

// GetHolderIdOk returns a tuple with the HolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetHolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HolderId.Get(), o.HolderId.IsSet()
}

// HasHolderId returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasHolderId() bool {
	if o != nil && o.HolderId.IsSet() {
		return true
	}

	return false
}

// SetHolderId gets a reference to the given NullableString and assigns it to the HolderId field.
func (o *CertificateSearchResult) SetHolderId(v string) {
	o.HolderId.Set(&v)
}

// SetHolderIdNil sets the value for HolderId to be an explicit nil
func (o *CertificateSearchResult) SetHolderIdNil() {
	o.HolderId.Set(nil)
}

// UnsetHolderId ensures that no value is present for HolderId, not even an explicit nil
func (o *CertificateSearchResult) UnsetHolderId() {
	o.HolderId.Unset()
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetIssuer() string {
	if o == nil || utils.IsNil(o.Issuer.Get()) {
		var ret string
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullableString and assigns it to the Issuer field.
func (o *CertificateSearchResult) SetIssuer(v string) {
	o.Issuer.Set(&v)
}

// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *CertificateSearchResult) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *CertificateSearchResult) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetKeyType returns the KeyType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetKeyType() string {
	if o == nil || utils.IsNil(o.KeyType.Get()) {
		var ret string
		return ret
	}
	return *o.KeyType.Get()
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyType.Get(), o.KeyType.IsSet()
}

// HasKeyType returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasKeyType() bool {
	if o != nil && o.KeyType.IsSet() {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given NullableString and assigns it to the KeyType field.
func (o *CertificateSearchResult) SetKeyType(v string) {
	o.KeyType.Set(&v)
}

// SetKeyTypeNil sets the value for KeyType to be an explicit nil
func (o *CertificateSearchResult) SetKeyTypeNil() {
	o.KeyType.Set(nil)
}

// UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
func (o *CertificateSearchResult) UnsetKeyType() {
	o.KeyType.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetLabels() []LabelData {
	if o == nil {
		var ret []LabelData
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetLabelsOk() ([]LabelData, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelData and assigns it to the Labels field.
func (o *CertificateSearchResult) SetLabels(v []LabelData) {
	o.Labels = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetMetadata() []CertificateMetadata {
	if o == nil {
		var ret []CertificateMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetMetadataOk() ([]CertificateMetadata, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []CertificateMetadata and assigns it to the Metadata field.
func (o *CertificateSearchResult) SetMetadata(v []CertificateMetadata) {
	o.Metadata = v
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetModule() string {
	if o == nil || utils.IsNil(o.Module.Get()) {
		var ret string
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableString and assigns it to the Module field.
func (o *CertificateSearchResult) SetModule(v string) {
	o.Module.Set(&v)
}

// SetModuleNil sets the value for Module to be an explicit nil
func (o *CertificateSearchResult) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *CertificateSearchResult) UnsetModule() {
	o.Module.Unset()
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetNotAfter() int64 {
	if o == nil || utils.IsNil(o.NotAfter.Get()) {
		var ret int64
		return ret
	}
	return *o.NotAfter.Get()
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetNotAfterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotAfter.Get(), o.NotAfter.IsSet()
}

// HasNotAfter returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasNotAfter() bool {
	if o != nil && o.NotAfter.IsSet() {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given NullableInt64 and assigns it to the NotAfter field.
func (o *CertificateSearchResult) SetNotAfter(v int64) {
	o.NotAfter.Set(&v)
}

// SetNotAfterNil sets the value for NotAfter to be an explicit nil
func (o *CertificateSearchResult) SetNotAfterNil() {
	o.NotAfter.Set(nil)
}

// UnsetNotAfter ensures that no value is present for NotAfter, not even an explicit nil
func (o *CertificateSearchResult) UnsetNotAfter() {
	o.NotAfter.Unset()
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetNotBefore() int64 {
	if o == nil || utils.IsNil(o.NotBefore.Get()) {
		var ret int64
		return ret
	}
	return *o.NotBefore.Get()
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetNotBeforeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotBefore.Get(), o.NotBefore.IsSet()
}

// HasNotBefore returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasNotBefore() bool {
	if o != nil && o.NotBefore.IsSet() {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given NullableInt64 and assigns it to the NotBefore field.
func (o *CertificateSearchResult) SetNotBefore(v int64) {
	o.NotBefore.Set(&v)
}

// SetNotBeforeNil sets the value for NotBefore to be an explicit nil
func (o *CertificateSearchResult) SetNotBeforeNil() {
	o.NotBefore.Set(nil)
}

// UnsetNotBefore ensures that no value is present for NotBefore, not even an explicit nil
func (o *CertificateSearchResult) UnsetNotBefore() {
	o.NotBefore.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetOwner() string {
	if o == nil || utils.IsNil(o.Owner.Get()) {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *CertificateSearchResult) SetOwner(v string) {
	o.Owner.Set(&v)
}

// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *CertificateSearchResult) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *CertificateSearchResult) UnsetOwner() {
	o.Owner.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetPermissions() CertificatePermissions {
	if o == nil || utils.IsNil(o.Permissions.Get()) {
		var ret CertificatePermissions
		return ret
	}
	return *o.Permissions.Get()
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetPermissionsOk() (*CertificatePermissions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions.Get(), o.Permissions.IsSet()
}

// HasPermissions returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasPermissions() bool {
	if o != nil && o.Permissions.IsSet() {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given NullableCertificatePermissions and assigns it to the Permissions field.
func (o *CertificateSearchResult) SetPermissions(v CertificatePermissions) {
	o.Permissions.Set(&v)
}

// SetPermissionsNil sets the value for Permissions to be an explicit nil
func (o *CertificateSearchResult) SetPermissionsNil() {
	o.Permissions.Set(nil)
}

// UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
func (o *CertificateSearchResult) UnsetPermissions() {
	o.Permissions.Unset()
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetPrivateKey() EscrowedPrivateKey {
	if o == nil || utils.IsNil(o.PrivateKey.Get()) {
		var ret EscrowedPrivateKey
		return ret
	}
	return *o.PrivateKey.Get()
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetPrivateKeyOk() (*EscrowedPrivateKey, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKey.Get(), o.PrivateKey.IsSet()
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasPrivateKey() bool {
	if o != nil && o.PrivateKey.IsSet() {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given NullableEscrowedPrivateKey and assigns it to the PrivateKey field.
func (o *CertificateSearchResult) SetPrivateKey(v EscrowedPrivateKey) {
	o.PrivateKey.Set(&v)
}

// SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil
func (o *CertificateSearchResult) SetPrivateKeyNil() {
	o.PrivateKey.Set(nil)
}

// UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
func (o *CertificateSearchResult) UnsetPrivateKey() {
	o.PrivateKey.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetProfile() string {
	if o == nil || utils.IsNil(o.Profile.Get()) {
		var ret string
		return ret
	}
	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// HasProfile returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasProfile() bool {
	if o != nil && o.Profile.IsSet() {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NullableString and assigns it to the Profile field.
func (o *CertificateSearchResult) SetProfile(v string) {
	o.Profile.Set(&v)
}

// SetProfileNil sets the value for Profile to be an explicit nil
func (o *CertificateSearchResult) SetProfileNil() {
	o.Profile.Set(nil)
}

// UnsetProfile ensures that no value is present for Profile, not even an explicit nil
func (o *CertificateSearchResult) UnsetProfile() {
	o.Profile.Unset()
}

// GetPublicKeyThumbprint returns the PublicKeyThumbprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetPublicKeyThumbprint() string {
	if o == nil || utils.IsNil(o.PublicKeyThumbprint.Get()) {
		var ret string
		return ret
	}
	return *o.PublicKeyThumbprint.Get()
}

// GetPublicKeyThumbprintOk returns a tuple with the PublicKeyThumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetPublicKeyThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKeyThumbprint.Get(), o.PublicKeyThumbprint.IsSet()
}

// HasPublicKeyThumbprint returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasPublicKeyThumbprint() bool {
	if o != nil && o.PublicKeyThumbprint.IsSet() {
		return true
	}

	return false
}

// SetPublicKeyThumbprint gets a reference to the given NullableString and assigns it to the PublicKeyThumbprint field.
func (o *CertificateSearchResult) SetPublicKeyThumbprint(v string) {
	o.PublicKeyThumbprint.Set(&v)
}

// SetPublicKeyThumbprintNil sets the value for PublicKeyThumbprint to be an explicit nil
func (o *CertificateSearchResult) SetPublicKeyThumbprintNil() {
	o.PublicKeyThumbprint.Set(nil)
}

// UnsetPublicKeyThumbprint ensures that no value is present for PublicKeyThumbprint, not even an explicit nil
func (o *CertificateSearchResult) UnsetPublicKeyThumbprint() {
	o.PublicKeyThumbprint.Unset()
}

// GetRevocationDate returns the RevocationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetRevocationDate() int64 {
	if o == nil || utils.IsNil(o.RevocationDate.Get()) {
		var ret int64
		return ret
	}
	return *o.RevocationDate.Get()
}

// GetRevocationDateOk returns a tuple with the RevocationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetRevocationDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevocationDate.Get(), o.RevocationDate.IsSet()
}

// HasRevocationDate returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasRevocationDate() bool {
	if o != nil && o.RevocationDate.IsSet() {
		return true
	}

	return false
}

// SetRevocationDate gets a reference to the given NullableInt64 and assigns it to the RevocationDate field.
func (o *CertificateSearchResult) SetRevocationDate(v int64) {
	o.RevocationDate.Set(&v)
}

// SetRevocationDateNil sets the value for RevocationDate to be an explicit nil
func (o *CertificateSearchResult) SetRevocationDateNil() {
	o.RevocationDate.Set(nil)
}

// UnsetRevocationDate ensures that no value is present for RevocationDate, not even an explicit nil
func (o *CertificateSearchResult) UnsetRevocationDate() {
	o.RevocationDate.Unset()
}

// GetRevocationReason returns the RevocationReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetRevocationReason() string {
	if o == nil || utils.IsNil(o.RevocationReason.Get()) {
		var ret string
		return ret
	}
	return *o.RevocationReason.Get()
}

// GetRevocationReasonOk returns a tuple with the RevocationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetRevocationReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevocationReason.Get(), o.RevocationReason.IsSet()
}

// HasRevocationReason returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasRevocationReason() bool {
	if o != nil && o.RevocationReason.IsSet() {
		return true
	}

	return false
}

// SetRevocationReason gets a reference to the given NullableString and assigns it to the RevocationReason field.
func (o *CertificateSearchResult) SetRevocationReason(v string) {
	o.RevocationReason.Set(&v)
}

// SetRevocationReasonNil sets the value for RevocationReason to be an explicit nil
func (o *CertificateSearchResult) SetRevocationReasonNil() {
	o.RevocationReason.Set(nil)
}

// UnsetRevocationReason ensures that no value is present for RevocationReason, not even an explicit nil
func (o *CertificateSearchResult) UnsetRevocationReason() {
	o.RevocationReason.Unset()
}

// GetSelfSigned returns the SelfSigned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetSelfSigned() bool {
	if o == nil || utils.IsNil(o.SelfSigned.Get()) {
		var ret bool
		return ret
	}
	return *o.SelfSigned.Get()
}

// GetSelfSignedOk returns a tuple with the SelfSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetSelfSignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelfSigned.Get(), o.SelfSigned.IsSet()
}

// HasSelfSigned returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasSelfSigned() bool {
	if o != nil && o.SelfSigned.IsSet() {
		return true
	}

	return false
}

// SetSelfSigned gets a reference to the given NullableBool and assigns it to the SelfSigned field.
func (o *CertificateSearchResult) SetSelfSigned(v bool) {
	o.SelfSigned.Set(&v)
}

// SetSelfSignedNil sets the value for SelfSigned to be an explicit nil
func (o *CertificateSearchResult) SetSelfSignedNil() {
	o.SelfSigned.Set(nil)
}

// UnsetSelfSigned ensures that no value is present for SelfSigned, not even an explicit nil
func (o *CertificateSearchResult) UnsetSelfSigned() {
	o.SelfSigned.Unset()
}

// GetSerial returns the Serial field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetSerial() string {
	if o == nil || utils.IsNil(o.Serial.Get()) {
		var ret string
		return ret
	}
	return *o.Serial.Get()
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Serial.Get(), o.Serial.IsSet()
}

// HasSerial returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasSerial() bool {
	if o != nil && o.Serial.IsSet() {
		return true
	}

	return false
}

// SetSerial gets a reference to the given NullableString and assigns it to the Serial field.
func (o *CertificateSearchResult) SetSerial(v string) {
	o.Serial.Set(&v)
}

// SetSerialNil sets the value for Serial to be an explicit nil
func (o *CertificateSearchResult) SetSerialNil() {
	o.Serial.Set(nil)
}

// UnsetSerial ensures that no value is present for Serial, not even an explicit nil
func (o *CertificateSearchResult) UnsetSerial() {
	o.Serial.Unset()
}

// GetSigningAlgorithm returns the SigningAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetSigningAlgorithm() string {
	if o == nil || utils.IsNil(o.SigningAlgorithm.Get()) {
		var ret string
		return ret
	}
	return *o.SigningAlgorithm.Get()
}

// GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetSigningAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningAlgorithm.Get(), o.SigningAlgorithm.IsSet()
}

// HasSigningAlgorithm returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasSigningAlgorithm() bool {
	if o != nil && o.SigningAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetSigningAlgorithm gets a reference to the given NullableString and assigns it to the SigningAlgorithm field.
func (o *CertificateSearchResult) SetSigningAlgorithm(v string) {
	o.SigningAlgorithm.Set(&v)
}

// SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil
func (o *CertificateSearchResult) SetSigningAlgorithmNil() {
	o.SigningAlgorithm.Set(nil)
}

// UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
func (o *CertificateSearchResult) UnsetSigningAlgorithm() {
	o.SigningAlgorithm.Unset()
}

// GetSubjectAlternateNames returns the SubjectAlternateNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetSubjectAlternateNames() []SubjectAlternateName {
	if o == nil {
		var ret []SubjectAlternateName
		return ret
	}
	return o.SubjectAlternateNames
}

// GetSubjectAlternateNamesOk returns a tuple with the SubjectAlternateNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetSubjectAlternateNamesOk() ([]SubjectAlternateName, bool) {
	if o == nil || utils.IsNil(o.SubjectAlternateNames) {
		return nil, false
	}
	return o.SubjectAlternateNames, true
}

// HasSubjectAlternateNames returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasSubjectAlternateNames() bool {
	if o != nil && !utils.IsNil(o.SubjectAlternateNames) {
		return true
	}

	return false
}

// SetSubjectAlternateNames gets a reference to the given []SubjectAlternateName and assigns it to the SubjectAlternateNames field.
func (o *CertificateSearchResult) SetSubjectAlternateNames(v []SubjectAlternateName) {
	o.SubjectAlternateNames = v
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetTeam() string {
	if o == nil || utils.IsNil(o.Team.Get()) {
		var ret string
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableString and assigns it to the Team field.
func (o *CertificateSearchResult) SetTeam(v string) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *CertificateSearchResult) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *CertificateSearchResult) UnsetTeam() {
	o.Team.Unset()
}

// GetThirdPartyData returns the ThirdPartyData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetThirdPartyData() []ThirdPartyItem {
	if o == nil {
		var ret []ThirdPartyItem
		return ret
	}
	return o.ThirdPartyData
}

// GetThirdPartyDataOk returns a tuple with the ThirdPartyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetThirdPartyDataOk() ([]ThirdPartyItem, bool) {
	if o == nil || utils.IsNil(o.ThirdPartyData) {
		return nil, false
	}
	return o.ThirdPartyData, true
}

// HasThirdPartyData returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasThirdPartyData() bool {
	if o != nil && !utils.IsNil(o.ThirdPartyData) {
		return true
	}

	return false
}

// SetThirdPartyData gets a reference to the given []ThirdPartyItem and assigns it to the ThirdPartyData field.
func (o *CertificateSearchResult) SetThirdPartyData(v []ThirdPartyItem) {
	o.ThirdPartyData = v
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetThumbprint() string {
	if o == nil || utils.IsNil(o.Thumbprint.Get()) {
		var ret string
		return ret
	}
	return *o.Thumbprint.Get()
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thumbprint.Get(), o.Thumbprint.IsSet()
}

// HasThumbprint returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasThumbprint() bool {
	if o != nil && o.Thumbprint.IsSet() {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given NullableString and assigns it to the Thumbprint field.
func (o *CertificateSearchResult) SetThumbprint(v string) {
	o.Thumbprint.Set(&v)
}

// SetThumbprintNil sets the value for Thumbprint to be an explicit nil
func (o *CertificateSearchResult) SetThumbprintNil() {
	o.Thumbprint.Set(nil)
}

// UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
func (o *CertificateSearchResult) UnsetThumbprint() {
	o.Thumbprint.Unset()
}

// GetTriggerResults returns the TriggerResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateSearchResult) GetTriggerResults() []TriggerResult {
	if o == nil {
		var ret []TriggerResult
		return ret
	}
	return o.TriggerResults
}

// GetTriggerResultsOk returns a tuple with the TriggerResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateSearchResult) GetTriggerResultsOk() ([]TriggerResult, bool) {
	if o == nil || utils.IsNil(o.TriggerResults) {
		return nil, false
	}
	return o.TriggerResults, true
}

// HasTriggerResults returns a boolean if a field has been set.
func (o *CertificateSearchResult) HasTriggerResults() bool {
	if o != nil && !utils.IsNil(o.TriggerResults) {
		return true
	}

	return false
}

// SetTriggerResults gets a reference to the given []TriggerResult and assigns it to the TriggerResults field.
func (o *CertificateSearchResult) SetTriggerResults(v []TriggerResult) {
	o.TriggerResults = v
}

func (o CertificateSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["_id"] = o.Id.Get()
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.ContactEmail.IsSet() {
		toSerialize["contactEmail"] = o.ContactEmail.Get()
	}
	if o.DiscoveredTrusted.IsSet() {
		toSerialize["discoveredTrusted"] = o.DiscoveredTrusted.Get()
	}
	if o.DiscoveryData != nil {
		toSerialize["discoveryData"] = o.DiscoveryData
	}
	if o.DiscoveryInfo != nil {
		toSerialize["discoveryInfo"] = o.DiscoveryInfo
	}
	if o.Dn.IsSet() {
		toSerialize["dn"] = o.Dn.Get()
	}
	if o.Grades != nil {
		toSerialize["grades"] = o.Grades
	}
	if o.HolderId.IsSet() {
		toSerialize["holderId"] = o.HolderId.Get()
	}
	if o.Issuer.IsSet() {
		toSerialize["issuer"] = o.Issuer.Get()
	}
	if o.KeyType.IsSet() {
		toSerialize["keyType"] = o.KeyType.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	if o.NotAfter.IsSet() {
		toSerialize["notAfter"] = o.NotAfter.Get()
	}
	if o.NotBefore.IsSet() {
		toSerialize["notBefore"] = o.NotBefore.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.Permissions.IsSet() {
		toSerialize["permissions"] = o.Permissions.Get()
	}
	if o.PrivateKey.IsSet() {
		toSerialize["privateKey"] = o.PrivateKey.Get()
	}
	if o.Profile.IsSet() {
		toSerialize["profile"] = o.Profile.Get()
	}
	if o.PublicKeyThumbprint.IsSet() {
		toSerialize["publicKeyThumbprint"] = o.PublicKeyThumbprint.Get()
	}
	if o.RevocationDate.IsSet() {
		toSerialize["revocationDate"] = o.RevocationDate.Get()
	}
	if o.RevocationReason.IsSet() {
		toSerialize["revocationReason"] = o.RevocationReason.Get()
	}
	if o.SelfSigned.IsSet() {
		toSerialize["selfSigned"] = o.SelfSigned.Get()
	}
	if o.Serial.IsSet() {
		toSerialize["serial"] = o.Serial.Get()
	}
	if o.SigningAlgorithm.IsSet() {
		toSerialize["signingAlgorithm"] = o.SigningAlgorithm.Get()
	}
	if o.SubjectAlternateNames != nil {
		toSerialize["subjectAlternateNames"] = o.SubjectAlternateNames
	}
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}
	if o.ThirdPartyData != nil {
		toSerialize["thirdPartyData"] = o.ThirdPartyData
	}
	if o.Thumbprint.IsSet() {
		toSerialize["thumbprint"] = o.Thumbprint.Get()
	}
	if o.TriggerResults != nil {
		toSerialize["triggerResults"] = o.TriggerResults
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateSearchResult) UnmarshalJSON(data []byte) (err error) {
	varCertificateSearchResult := _CertificateSearchResult{}

	err = json.Unmarshal(data, &varCertificateSearchResult)

	if err != nil {
		return err
	}

	*o = CertificateSearchResult(varCertificateSearchResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "discoveredTrusted")
		delete(additionalProperties, "discoveryData")
		delete(additionalProperties, "discoveryInfo")
		delete(additionalProperties, "dn")
		delete(additionalProperties, "grades")
		delete(additionalProperties, "holderId")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "keyType")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "module")
		delete(additionalProperties, "notAfter")
		delete(additionalProperties, "notBefore")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "privateKey")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "publicKeyThumbprint")
		delete(additionalProperties, "revocationDate")
		delete(additionalProperties, "revocationReason")
		delete(additionalProperties, "selfSigned")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "signingAlgorithm")
		delete(additionalProperties, "subjectAlternateNames")
		delete(additionalProperties, "team")
		delete(additionalProperties, "thirdPartyData")
		delete(additionalProperties, "thumbprint")
		delete(additionalProperties, "triggerResults")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateSearchResult struct {
	value *CertificateSearchResult
	isSet bool
}

func (v NullableCertificateSearchResult) Get() *CertificateSearchResult {
	return v.value
}

func (v *NullableCertificateSearchResult) Set(val *CertificateSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateSearchResult(val *CertificateSearchResult) *NullableCertificateSearchResult {
	return &NullableCertificateSearchResult{value: val, isSet: true}
}

func (v NullableCertificateSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
