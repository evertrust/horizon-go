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

// checks if the CertificateResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CertificateResponse{}

// CertificateResponse struct for CertificateResponse
type CertificateResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// The certificate's PEM-encoded content
	Certificate string `json:"certificate"`
	// The certificate's contact email. It will be used to send notifications about the certificate's expiration and revocation
	ContactEmail utils.NullableString `json:"contactEmail,omitempty"`
	// Whether the certificate's revocation status is synchronized with a CRL
	CrlSynchronized utils.NullableBool `json:"crlSynchronized,omitempty"`
	// If the certificate was discovered and is found to be issued by an existing trusted CA, this field will be set to true. If the certificate was discovered and is not found to be issued by an existing trusted CA, this field will be set to false. If the certificate was not discovered, this field will be null
	DiscoveredTrusted utils.NullableBool `json:"discoveredTrusted,omitempty"`
	// A list of metadata containing information on where the certificate was discovered
	DiscoveryData []HostDiscoveryData `json:"discoveryData,omitempty"`
	// A list of metadata containing information on how and when the certificate was discovered
	DiscoveryInfo []DiscoveryInfo `json:"discoveryInfo,omitempty"`
	// The certificate's Distinguished Name
	Dn string `json:"dn"`
	// Whether the certificate is escrowed
	Escrowed bool `json:"escrowed"`
	// The certificate's extensions
	Extensions []CertificateExtension `json:"extensions,omitempty"`
	// The certificate's grades for the enabled grading policies
	Grades []GradingPolicyResult `json:"grades,omitempty"`
	// The certificate's holder ID. This is a computed field that is used to count how many similar certificates are in use simultaneously by the same holder
	HolderId string `json:"holderId"`
	// The certificate's issuer Distinguished Name
	Issuer string `json:"issuer"`
	// The certificate's key type
	KeyType string `json:"keyType" validate:"regexp=(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ec-brainpoolp256r1|ec-brainpoolp384r1|ec-brainpoolp512r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512)(\\\\\\\\+(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ec-brainpoolp256r1|ec-brainpoolp384r1|ec-brainpoolp512r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512))?"`
	// The certificate's labels
	Labels []LabelData `json:"labels,omitempty"`
	// The certificate's technical metadata used internally
	Metadata []CertificateMetadata `json:"metadata"`
	// The certificate's module
	Module string `json:"module"`
	// The certificate's expiration date in milliseconds since the epoch
	NotAfter int64 `json:"notAfter"`
	// The certificate's start date in milliseconds since the epoch
	NotBefore int64 `json:"notBefore"`
	// The certificate's owner. This is a reference to a local identity identifier
	Owner utils.NullableString `json:"owner,omitempty"`
	// The certificate's profile
	Profile utils.NullableString `json:"profile,omitempty"`
	// The certificate's public key thumbprint
	PublicKeyThumbprint string `json:"publicKeyThumbprint"`
	// The certificate's revocation date in milliseconds since the epoch. This field is only present if the certificate is revoked
	RevocationDate utils.NullableInt64 `json:"revocationDate,omitempty"`
	// The certificate's revocation reason
	RevocationReason utils.NullableString `json:"revocationReason,omitempty"`
	// Whether the certificate is revoked
	Revoked bool `json:"revoked"`
	// Whether the certificate is self-signed
	SelfSigned bool `json:"selfSigned"`
	// The certificate's serial number
	Serial string `json:"serial"`
	// The certificate's signing algorithm
	SigningAlgorithm string `json:"signingAlgorithm"`
	// The certificate's Subject Alternate Names
	SubjectAlternateNames []SubjectAlternateName `json:"subjectAlternateNames"`
	// The certificate's team. This is a reference to a team identifier. It will be used to determine the certificate's permissions and send notifications
	Team utils.NullableString `json:"team,omitempty"`
	// The certificate's information about synchronization with Horizon supported third parties
	ThirdPartyData []ThirdPartyItem `json:"thirdPartyData,omitempty"`
	// The certificate's thumbprint
	Thumbprint string `json:"thumbprint"`
	// The result of the execution of triggers on this certificate
	TriggerResults       []TriggerResult `json:"triggerResults,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateResponse CertificateResponse

// NewCertificateResponse instantiates a new CertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateResponse(id string, certificate string, dn string, escrowed bool, holderId string, issuer string, keyType string, metadata []CertificateMetadata, module string, notAfter int64, notBefore int64, publicKeyThumbprint string, revoked bool, selfSigned bool, serial string, signingAlgorithm string, subjectAlternateNames []SubjectAlternateName, thumbprint string) *CertificateResponse {
	this := CertificateResponse{}
	this.Id = id
	this.Certificate = certificate
	this.Dn = dn
	this.Escrowed = escrowed
	this.HolderId = holderId
	this.Issuer = issuer
	this.KeyType = keyType
	this.Metadata = metadata
	this.Module = module
	this.NotAfter = notAfter
	this.NotBefore = notBefore
	this.PublicKeyThumbprint = publicKeyThumbprint
	this.Revoked = revoked
	this.SelfSigned = selfSigned
	this.Serial = serial
	this.SigningAlgorithm = signingAlgorithm
	this.SubjectAlternateNames = subjectAlternateNames
	this.Thumbprint = thumbprint
	return &this
}

// NewCertificateResponseWithDefaults instantiates a new CertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateResponseWithDefaults() *CertificateResponse {
	this := CertificateResponse{}
	return &this
}

// GetId returns the Id field value
func (o *CertificateResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CertificateResponse) SetId(v string) {
	o.Id = v
}

// GetCertificate returns the Certificate field value
func (o *CertificateResponse) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *CertificateResponse) SetCertificate(v string) {
	o.Certificate = v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetContactEmail() string {
	if o == nil || utils.IsNil(o.ContactEmail.Get()) {
		var ret string
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetContactEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *CertificateResponse) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableString and assigns it to the ContactEmail field.
func (o *CertificateResponse) SetContactEmail(v string) {
	o.ContactEmail.Set(&v)
}

// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *CertificateResponse) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *CertificateResponse) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetCrlSynchronized returns the CrlSynchronized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetCrlSynchronized() bool {
	if o == nil || utils.IsNil(o.CrlSynchronized.Get()) {
		var ret bool
		return ret
	}
	return *o.CrlSynchronized.Get()
}

// GetCrlSynchronizedOk returns a tuple with the CrlSynchronized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetCrlSynchronizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CrlSynchronized.Get(), o.CrlSynchronized.IsSet()
}

// HasCrlSynchronized returns a boolean if a field has been set.
func (o *CertificateResponse) HasCrlSynchronized() bool {
	if o != nil && o.CrlSynchronized.IsSet() {
		return true
	}

	return false
}

// SetCrlSynchronized gets a reference to the given NullableBool and assigns it to the CrlSynchronized field.
func (o *CertificateResponse) SetCrlSynchronized(v bool) {
	o.CrlSynchronized.Set(&v)
}

// SetCrlSynchronizedNil sets the value for CrlSynchronized to be an explicit nil
func (o *CertificateResponse) SetCrlSynchronizedNil() {
	o.CrlSynchronized.Set(nil)
}

// UnsetCrlSynchronized ensures that no value is present for CrlSynchronized, not even an explicit nil
func (o *CertificateResponse) UnsetCrlSynchronized() {
	o.CrlSynchronized.Unset()
}

// GetDiscoveredTrusted returns the DiscoveredTrusted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetDiscoveredTrusted() bool {
	if o == nil || utils.IsNil(o.DiscoveredTrusted.Get()) {
		var ret bool
		return ret
	}
	return *o.DiscoveredTrusted.Get()
}

// GetDiscoveredTrustedOk returns a tuple with the DiscoveredTrusted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetDiscoveredTrustedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscoveredTrusted.Get(), o.DiscoveredTrusted.IsSet()
}

// HasDiscoveredTrusted returns a boolean if a field has been set.
func (o *CertificateResponse) HasDiscoveredTrusted() bool {
	if o != nil && o.DiscoveredTrusted.IsSet() {
		return true
	}

	return false
}

// SetDiscoveredTrusted gets a reference to the given NullableBool and assigns it to the DiscoveredTrusted field.
func (o *CertificateResponse) SetDiscoveredTrusted(v bool) {
	o.DiscoveredTrusted.Set(&v)
}

// SetDiscoveredTrustedNil sets the value for DiscoveredTrusted to be an explicit nil
func (o *CertificateResponse) SetDiscoveredTrustedNil() {
	o.DiscoveredTrusted.Set(nil)
}

// UnsetDiscoveredTrusted ensures that no value is present for DiscoveredTrusted, not even an explicit nil
func (o *CertificateResponse) UnsetDiscoveredTrusted() {
	o.DiscoveredTrusted.Unset()
}

// GetDiscoveryData returns the DiscoveryData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetDiscoveryData() []HostDiscoveryData {
	if o == nil {
		var ret []HostDiscoveryData
		return ret
	}
	return o.DiscoveryData
}

// GetDiscoveryDataOk returns a tuple with the DiscoveryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetDiscoveryDataOk() ([]HostDiscoveryData, bool) {
	if o == nil || utils.IsNil(o.DiscoveryData) {
		return nil, false
	}
	return o.DiscoveryData, true
}

// HasDiscoveryData returns a boolean if a field has been set.
func (o *CertificateResponse) HasDiscoveryData() bool {
	if o != nil && !utils.IsNil(o.DiscoveryData) {
		return true
	}

	return false
}

// SetDiscoveryData gets a reference to the given []HostDiscoveryData and assigns it to the DiscoveryData field.
func (o *CertificateResponse) SetDiscoveryData(v []HostDiscoveryData) {
	o.DiscoveryData = v
}

// GetDiscoveryInfo returns the DiscoveryInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetDiscoveryInfo() []DiscoveryInfo {
	if o == nil {
		var ret []DiscoveryInfo
		return ret
	}
	return o.DiscoveryInfo
}

// GetDiscoveryInfoOk returns a tuple with the DiscoveryInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetDiscoveryInfoOk() ([]DiscoveryInfo, bool) {
	if o == nil || utils.IsNil(o.DiscoveryInfo) {
		return nil, false
	}
	return o.DiscoveryInfo, true
}

// HasDiscoveryInfo returns a boolean if a field has been set.
func (o *CertificateResponse) HasDiscoveryInfo() bool {
	if o != nil && !utils.IsNil(o.DiscoveryInfo) {
		return true
	}

	return false
}

// SetDiscoveryInfo gets a reference to the given []DiscoveryInfo and assigns it to the DiscoveryInfo field.
func (o *CertificateResponse) SetDiscoveryInfo(v []DiscoveryInfo) {
	o.DiscoveryInfo = v
}

// GetDn returns the Dn field value
func (o *CertificateResponse) GetDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dn
}

// GetDnOk returns a tuple with the Dn field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dn, true
}

// SetDn sets field value
func (o *CertificateResponse) SetDn(v string) {
	o.Dn = v
}

// GetEscrowed returns the Escrowed field value
func (o *CertificateResponse) GetEscrowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Escrowed
}

// GetEscrowedOk returns a tuple with the Escrowed field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetEscrowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Escrowed, true
}

// SetEscrowed sets field value
func (o *CertificateResponse) SetEscrowed(v bool) {
	o.Escrowed = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *CertificateResponse) GetExtensions() []CertificateExtension {
	if o == nil || utils.IsNil(o.Extensions) {
		var ret []CertificateExtension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetExtensionsOk() ([]CertificateExtension, bool) {
	if o == nil || utils.IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *CertificateResponse) HasExtensions() bool {
	if o != nil && !utils.IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []CertificateExtension and assigns it to the Extensions field.
func (o *CertificateResponse) SetExtensions(v []CertificateExtension) {
	o.Extensions = v
}

// GetGrades returns the Grades field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetGrades() []GradingPolicyResult {
	if o == nil {
		var ret []GradingPolicyResult
		return ret
	}
	return o.Grades
}

// GetGradesOk returns a tuple with the Grades field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetGradesOk() ([]GradingPolicyResult, bool) {
	if o == nil || utils.IsNil(o.Grades) {
		return nil, false
	}
	return o.Grades, true
}

// HasGrades returns a boolean if a field has been set.
func (o *CertificateResponse) HasGrades() bool {
	if o != nil && !utils.IsNil(o.Grades) {
		return true
	}

	return false
}

// SetGrades gets a reference to the given []GradingPolicyResult and assigns it to the Grades field.
func (o *CertificateResponse) SetGrades(v []GradingPolicyResult) {
	o.Grades = v
}

// GetHolderId returns the HolderId field value
func (o *CertificateResponse) GetHolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HolderId
}

// GetHolderIdOk returns a tuple with the HolderId field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetHolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HolderId, true
}

// SetHolderId sets field value
func (o *CertificateResponse) SetHolderId(v string) {
	o.HolderId = v
}

// GetIssuer returns the Issuer field value
func (o *CertificateResponse) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *CertificateResponse) SetIssuer(v string) {
	o.Issuer = v
}

// GetKeyType returns the KeyType field value
func (o *CertificateResponse) GetKeyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *CertificateResponse) SetKeyType(v string) {
	o.KeyType = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetLabels() []LabelData {
	if o == nil {
		var ret []LabelData
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetLabelsOk() ([]LabelData, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CertificateResponse) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelData and assigns it to the Labels field.
func (o *CertificateResponse) SetLabels(v []LabelData) {
	o.Labels = v
}

// GetMetadata returns the Metadata field value
func (o *CertificateResponse) GetMetadata() []CertificateMetadata {
	if o == nil {
		var ret []CertificateMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetMetadataOk() ([]CertificateMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *CertificateResponse) SetMetadata(v []CertificateMetadata) {
	o.Metadata = v
}

// GetModule returns the Module field value
func (o *CertificateResponse) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *CertificateResponse) SetModule(v string) {
	o.Module = v
}

// GetNotAfter returns the NotAfter field value
func (o *CertificateResponse) GetNotAfter() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetNotAfterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotAfter, true
}

// SetNotAfter sets field value
func (o *CertificateResponse) SetNotAfter(v int64) {
	o.NotAfter = v
}

// GetNotBefore returns the NotBefore field value
func (o *CertificateResponse) GetNotBefore() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetNotBeforeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotBefore, true
}

// SetNotBefore sets field value
func (o *CertificateResponse) SetNotBefore(v int64) {
	o.NotBefore = v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetOwner() string {
	if o == nil || utils.IsNil(o.Owner.Get()) {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *CertificateResponse) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *CertificateResponse) SetOwner(v string) {
	o.Owner.Set(&v)
}

// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *CertificateResponse) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *CertificateResponse) UnsetOwner() {
	o.Owner.Unset()
}

// GetProfile returns the Profile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetProfile() string {
	if o == nil || utils.IsNil(o.Profile.Get()) {
		var ret string
		return ret
	}
	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// HasProfile returns a boolean if a field has been set.
func (o *CertificateResponse) HasProfile() bool {
	if o != nil && o.Profile.IsSet() {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NullableString and assigns it to the Profile field.
func (o *CertificateResponse) SetProfile(v string) {
	o.Profile.Set(&v)
}

// SetProfileNil sets the value for Profile to be an explicit nil
func (o *CertificateResponse) SetProfileNil() {
	o.Profile.Set(nil)
}

// UnsetProfile ensures that no value is present for Profile, not even an explicit nil
func (o *CertificateResponse) UnsetProfile() {
	o.Profile.Unset()
}

// GetPublicKeyThumbprint returns the PublicKeyThumbprint field value
func (o *CertificateResponse) GetPublicKeyThumbprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKeyThumbprint
}

// GetPublicKeyThumbprintOk returns a tuple with the PublicKeyThumbprint field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetPublicKeyThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKeyThumbprint, true
}

// SetPublicKeyThumbprint sets field value
func (o *CertificateResponse) SetPublicKeyThumbprint(v string) {
	o.PublicKeyThumbprint = v
}

// GetRevocationDate returns the RevocationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetRevocationDate() int64 {
	if o == nil || utils.IsNil(o.RevocationDate.Get()) {
		var ret int64
		return ret
	}
	return *o.RevocationDate.Get()
}

// GetRevocationDateOk returns a tuple with the RevocationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetRevocationDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevocationDate.Get(), o.RevocationDate.IsSet()
}

// HasRevocationDate returns a boolean if a field has been set.
func (o *CertificateResponse) HasRevocationDate() bool {
	if o != nil && o.RevocationDate.IsSet() {
		return true
	}

	return false
}

// SetRevocationDate gets a reference to the given NullableInt64 and assigns it to the RevocationDate field.
func (o *CertificateResponse) SetRevocationDate(v int64) {
	o.RevocationDate.Set(&v)
}

// SetRevocationDateNil sets the value for RevocationDate to be an explicit nil
func (o *CertificateResponse) SetRevocationDateNil() {
	o.RevocationDate.Set(nil)
}

// UnsetRevocationDate ensures that no value is present for RevocationDate, not even an explicit nil
func (o *CertificateResponse) UnsetRevocationDate() {
	o.RevocationDate.Unset()
}

// GetRevocationReason returns the RevocationReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetRevocationReason() string {
	if o == nil || utils.IsNil(o.RevocationReason.Get()) {
		var ret string
		return ret
	}
	return *o.RevocationReason.Get()
}

// GetRevocationReasonOk returns a tuple with the RevocationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetRevocationReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RevocationReason.Get(), o.RevocationReason.IsSet()
}

// HasRevocationReason returns a boolean if a field has been set.
func (o *CertificateResponse) HasRevocationReason() bool {
	if o != nil && o.RevocationReason.IsSet() {
		return true
	}

	return false
}

// SetRevocationReason gets a reference to the given NullableString and assigns it to the RevocationReason field.
func (o *CertificateResponse) SetRevocationReason(v string) {
	o.RevocationReason.Set(&v)
}

// SetRevocationReasonNil sets the value for RevocationReason to be an explicit nil
func (o *CertificateResponse) SetRevocationReasonNil() {
	o.RevocationReason.Set(nil)
}

// UnsetRevocationReason ensures that no value is present for RevocationReason, not even an explicit nil
func (o *CertificateResponse) UnsetRevocationReason() {
	o.RevocationReason.Unset()
}

// GetRevoked returns the Revoked field value
func (o *CertificateResponse) GetRevoked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetRevokedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revoked, true
}

// SetRevoked sets field value
func (o *CertificateResponse) SetRevoked(v bool) {
	o.Revoked = v
}

// GetSelfSigned returns the SelfSigned field value
func (o *CertificateResponse) GetSelfSigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SelfSigned
}

// GetSelfSignedOk returns a tuple with the SelfSigned field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetSelfSignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfSigned, true
}

// SetSelfSigned sets field value
func (o *CertificateResponse) SetSelfSigned(v bool) {
	o.SelfSigned = v
}

// GetSerial returns the Serial field value
func (o *CertificateResponse) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *CertificateResponse) SetSerial(v string) {
	o.Serial = v
}

// GetSigningAlgorithm returns the SigningAlgorithm field value
func (o *CertificateResponse) GetSigningAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SigningAlgorithm
}

// GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetSigningAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningAlgorithm, true
}

// SetSigningAlgorithm sets field value
func (o *CertificateResponse) SetSigningAlgorithm(v string) {
	o.SigningAlgorithm = v
}

// GetSubjectAlternateNames returns the SubjectAlternateNames field value
func (o *CertificateResponse) GetSubjectAlternateNames() []SubjectAlternateName {
	if o == nil {
		var ret []SubjectAlternateName
		return ret
	}

	return o.SubjectAlternateNames
}

// GetSubjectAlternateNamesOk returns a tuple with the SubjectAlternateNames field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetSubjectAlternateNamesOk() ([]SubjectAlternateName, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectAlternateNames, true
}

// SetSubjectAlternateNames sets field value
func (o *CertificateResponse) SetSubjectAlternateNames(v []SubjectAlternateName) {
	o.SubjectAlternateNames = v
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetTeam() string {
	if o == nil || utils.IsNil(o.Team.Get()) {
		var ret string
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *CertificateResponse) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableString and assigns it to the Team field.
func (o *CertificateResponse) SetTeam(v string) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *CertificateResponse) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *CertificateResponse) UnsetTeam() {
	o.Team.Unset()
}

// GetThirdPartyData returns the ThirdPartyData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetThirdPartyData() []ThirdPartyItem {
	if o == nil {
		var ret []ThirdPartyItem
		return ret
	}
	return o.ThirdPartyData
}

// GetThirdPartyDataOk returns a tuple with the ThirdPartyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetThirdPartyDataOk() ([]ThirdPartyItem, bool) {
	if o == nil || utils.IsNil(o.ThirdPartyData) {
		return nil, false
	}
	return o.ThirdPartyData, true
}

// HasThirdPartyData returns a boolean if a field has been set.
func (o *CertificateResponse) HasThirdPartyData() bool {
	if o != nil && !utils.IsNil(o.ThirdPartyData) {
		return true
	}

	return false
}

// SetThirdPartyData gets a reference to the given []ThirdPartyItem and assigns it to the ThirdPartyData field.
func (o *CertificateResponse) SetThirdPartyData(v []ThirdPartyItem) {
	o.ThirdPartyData = v
}

// GetThumbprint returns the Thumbprint field value
func (o *CertificateResponse) GetThumbprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value
// and a boolean to check if the value has been set.
func (o *CertificateResponse) GetThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Thumbprint, true
}

// SetThumbprint sets field value
func (o *CertificateResponse) SetThumbprint(v string) {
	o.Thumbprint = v
}

// GetTriggerResults returns the TriggerResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateResponse) GetTriggerResults() []TriggerResult {
	if o == nil {
		var ret []TriggerResult
		return ret
	}
	return o.TriggerResults
}

// GetTriggerResultsOk returns a tuple with the TriggerResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateResponse) GetTriggerResultsOk() ([]TriggerResult, bool) {
	if o == nil || utils.IsNil(o.TriggerResults) {
		return nil, false
	}
	return o.TriggerResults, true
}

// HasTriggerResults returns a boolean if a field has been set.
func (o *CertificateResponse) HasTriggerResults() bool {
	if o != nil && !utils.IsNil(o.TriggerResults) {
		return true
	}

	return false
}

// SetTriggerResults gets a reference to the given []TriggerResult and assigns it to the TriggerResults field.
func (o *CertificateResponse) SetTriggerResults(v []TriggerResult) {
	o.TriggerResults = v
}

func (o CertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["certificate"] = o.Certificate
	if o.ContactEmail.IsSet() {
		toSerialize["contactEmail"] = o.ContactEmail.Get()
	}
	if o.CrlSynchronized.IsSet() {
		toSerialize["crlSynchronized"] = o.CrlSynchronized.Get()
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
	toSerialize["dn"] = o.Dn
	toSerialize["escrowed"] = o.Escrowed
	if !utils.IsNil(o.Extensions) {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Grades != nil {
		toSerialize["grades"] = o.Grades
	}
	toSerialize["holderId"] = o.HolderId
	toSerialize["issuer"] = o.Issuer
	toSerialize["keyType"] = o.KeyType
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["metadata"] = o.Metadata
	toSerialize["module"] = o.Module
	toSerialize["notAfter"] = o.NotAfter
	toSerialize["notBefore"] = o.NotBefore
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.Profile.IsSet() {
		toSerialize["profile"] = o.Profile.Get()
	}
	toSerialize["publicKeyThumbprint"] = o.PublicKeyThumbprint
	if o.RevocationDate.IsSet() {
		toSerialize["revocationDate"] = o.RevocationDate.Get()
	}
	if o.RevocationReason.IsSet() {
		toSerialize["revocationReason"] = o.RevocationReason.Get()
	}
	toSerialize["revoked"] = o.Revoked
	toSerialize["selfSigned"] = o.SelfSigned
	toSerialize["serial"] = o.Serial
	toSerialize["signingAlgorithm"] = o.SigningAlgorithm
	toSerialize["subjectAlternateNames"] = o.SubjectAlternateNames
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}
	if o.ThirdPartyData != nil {
		toSerialize["thirdPartyData"] = o.ThirdPartyData
	}
	toSerialize["thumbprint"] = o.Thumbprint
	if o.TriggerResults != nil {
		toSerialize["triggerResults"] = o.TriggerResults
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"certificate",
		"dn",
		"escrowed",
		"holderId",
		"issuer",
		"keyType",
		"metadata",
		"module",
		"notAfter",
		"notBefore",
		"publicKeyThumbprint",
		"revoked",
		"selfSigned",
		"serial",
		"signingAlgorithm",
		"subjectAlternateNames",
		"thumbprint",
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

	varCertificateResponse := _CertificateResponse{}

	err = json.Unmarshal(data, &varCertificateResponse)

	if err != nil {
		return err
	}

	*o = CertificateResponse(varCertificateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "crlSynchronized")
		delete(additionalProperties, "discoveredTrusted")
		delete(additionalProperties, "discoveryData")
		delete(additionalProperties, "discoveryInfo")
		delete(additionalProperties, "dn")
		delete(additionalProperties, "escrowed")
		delete(additionalProperties, "extensions")
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
		delete(additionalProperties, "profile")
		delete(additionalProperties, "publicKeyThumbprint")
		delete(additionalProperties, "revocationDate")
		delete(additionalProperties, "revocationReason")
		delete(additionalProperties, "revoked")
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

type NullableCertificateResponse struct {
	value *CertificateResponse
	isSet bool
}

func (v NullableCertificateResponse) Get() *CertificateResponse {
	return v.value
}

func (v *NullableCertificateResponse) Set(val *CertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateResponse(val *CertificateResponse) *NullableCertificateResponse {
	return &NullableCertificateResponse{value: val, isSet: true}
}

func (v NullableCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
