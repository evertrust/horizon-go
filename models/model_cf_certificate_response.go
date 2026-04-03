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

// checks if the CFCertificateResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CFCertificateResponse{}

// CFCertificateResponse struct for CFCertificateResponse
type CFCertificateResponse struct {
	Aias *CFCertificateResponseAias `json:"aias,omitempty"`
	// The certificate AKI
	AuthorityKeyIdentifier *string                                `json:"authorityKeyIdentifier,omitempty"`
	BasicConstraints       *CFCertificateResponseBasicConstraints `json:"basicConstraints,omitempty"`
	// The thumbprint of the certificate using SHAOne algorithm
	CertificateSHAOneThumbprint string `json:"certificateSHAOneThumbprint"`
	// The certificate's thumbprint
	CertificateThumbprint string `json:"certificateThumbprint"`
	// The certificate's CRLDP if any
	Crldps []string `json:"crldps,omitempty"`
	// The certificate's Distinguished Name
	Dn         string                `json:"dn"`
	DnElements []CFDistinguishedName `json:"dnElements"`
	// The certificate extended key's usage
	ExtendedKeyUsages []string `json:"extendedKeyUsages"`
	// The certificate's extensions
	Extensions []CertificateExtension `json:"extensions,omitempty"`
	// If the extended key usage are critical
	IsExtendedKeyUsagesCritical bool `json:"isExtendedKeyUsagesCritical"`
	// If the key usage of the certificate are critical
	IsKeyUsagesCritical bool `json:"isKeyUsagesCritical"`
	// The certificate's issuer Distinguished Name
	IssuerDn string `json:"issuerDn"`
	// The certificate's key type
	KeyType string `json:"keyType" validate:"regexp=(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512)(\\\\\\\\+(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512))?"`
	// The certificate key's usage
	KeyUsages []string `json:"keyUsages"`
	// The certificate's expiration date in milliseconds since the epoch
	NotAfter int64 `json:"notAfter"`
	// The certificate's start date in milliseconds since the epoch
	NotBefore int64 `json:"notBefore"`
	// The certificate's PEM-encoded content
	Pem      string                               `json:"pem"`
	Policies []CFCertificateResponsePoliciesInner `json:"policies,omitempty"`
	// The certificate's public key thumbprint
	PublicKeyThumbprint string `json:"publicKeyThumbprint"`
	// The certificate's SAN
	Sans []SubjectAlternateName `json:"sans,omitempty"`
	// Whether the certificate is self-signed
	SelfSigned bool `json:"selfSigned"`
	// The certificate's serial number
	Serial *string `json:"serial,omitempty"`
	// The certificate's signing algorithm
	SigningAlgorithm      string                                            `json:"signingAlgorithm"`
	SubjectKeyIdentifier  string                                            `json:"subjectKeyIdentifier"`
	UnsupportedExtensions []CFCertificateResponseUnsupportedExtensionsInner `json:"unsupportedExtensions,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _CFCertificateResponse CFCertificateResponse

// NewCFCertificateResponse instantiates a new CFCertificateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFCertificateResponse(certificateSHAOneThumbprint string, certificateThumbprint string, dn string, dnElements []CFDistinguishedName, extendedKeyUsages []string, isExtendedKeyUsagesCritical bool, isKeyUsagesCritical bool, issuerDn string, keyType string, keyUsages []string, notAfter int64, notBefore int64, pem string, publicKeyThumbprint string, selfSigned bool, signingAlgorithm string, subjectKeyIdentifier string) *CFCertificateResponse {
	this := CFCertificateResponse{}
	this.CertificateSHAOneThumbprint = certificateSHAOneThumbprint
	this.CertificateThumbprint = certificateThumbprint
	this.Dn = dn
	this.DnElements = dnElements
	this.ExtendedKeyUsages = extendedKeyUsages
	this.IsExtendedKeyUsagesCritical = isExtendedKeyUsagesCritical
	this.IsKeyUsagesCritical = isKeyUsagesCritical
	this.IssuerDn = issuerDn
	this.KeyType = keyType
	this.KeyUsages = keyUsages
	this.NotAfter = notAfter
	this.NotBefore = notBefore
	this.Pem = pem
	this.PublicKeyThumbprint = publicKeyThumbprint
	this.SelfSigned = selfSigned
	this.SigningAlgorithm = signingAlgorithm
	this.SubjectKeyIdentifier = subjectKeyIdentifier
	return &this
}

// NewCFCertificateResponseWithDefaults instantiates a new CFCertificateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFCertificateResponseWithDefaults() *CFCertificateResponse {
	this := CFCertificateResponse{}
	return &this
}

// GetAias returns the Aias field value if set, zero value otherwise.
func (o *CFCertificateResponse) GetAias() CFCertificateResponseAias {
	if o == nil || utils.IsNil(o.Aias) {
		var ret CFCertificateResponseAias
		return ret
	}
	return *o.Aias
}

// GetAiasOk returns a tuple with the Aias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetAiasOk() (*CFCertificateResponseAias, bool) {
	if o == nil || utils.IsNil(o.Aias) {
		return nil, false
	}
	return o.Aias, true
}

// HasAias returns a boolean if a field has been set.
func (o *CFCertificateResponse) HasAias() bool {
	if o != nil && !utils.IsNil(o.Aias) {
		return true
	}

	return false
}

// SetAias gets a reference to the given CFCertificateResponseAias and assigns it to the Aias field.
func (o *CFCertificateResponse) SetAias(v CFCertificateResponseAias) {
	o.Aias = &v
}

// GetAuthorityKeyIdentifier returns the AuthorityKeyIdentifier field value if set, zero value otherwise.
func (o *CFCertificateResponse) GetAuthorityKeyIdentifier() string {
	if o == nil || utils.IsNil(o.AuthorityKeyIdentifier) {
		var ret string
		return ret
	}
	return *o.AuthorityKeyIdentifier
}

// GetAuthorityKeyIdentifierOk returns a tuple with the AuthorityKeyIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetAuthorityKeyIdentifierOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AuthorityKeyIdentifier) {
		return nil, false
	}
	return o.AuthorityKeyIdentifier, true
}

// HasAuthorityKeyIdentifier returns a boolean if a field has been set.
func (o *CFCertificateResponse) HasAuthorityKeyIdentifier() bool {
	if o != nil && !utils.IsNil(o.AuthorityKeyIdentifier) {
		return true
	}

	return false
}

// SetAuthorityKeyIdentifier gets a reference to the given string and assigns it to the AuthorityKeyIdentifier field.
func (o *CFCertificateResponse) SetAuthorityKeyIdentifier(v string) {
	o.AuthorityKeyIdentifier = &v
}

// GetBasicConstraints returns the BasicConstraints field value if set, zero value otherwise.
func (o *CFCertificateResponse) GetBasicConstraints() CFCertificateResponseBasicConstraints {
	if o == nil || utils.IsNil(o.BasicConstraints) {
		var ret CFCertificateResponseBasicConstraints
		return ret
	}
	return *o.BasicConstraints
}

// GetBasicConstraintsOk returns a tuple with the BasicConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetBasicConstraintsOk() (*CFCertificateResponseBasicConstraints, bool) {
	if o == nil || utils.IsNil(o.BasicConstraints) {
		return nil, false
	}
	return o.BasicConstraints, true
}

// HasBasicConstraints returns a boolean if a field has been set.
func (o *CFCertificateResponse) HasBasicConstraints() bool {
	if o != nil && !utils.IsNil(o.BasicConstraints) {
		return true
	}

	return false
}

// SetBasicConstraints gets a reference to the given CFCertificateResponseBasicConstraints and assigns it to the BasicConstraints field.
func (o *CFCertificateResponse) SetBasicConstraints(v CFCertificateResponseBasicConstraints) {
	o.BasicConstraints = &v
}

// GetCertificateSHAOneThumbprint returns the CertificateSHAOneThumbprint field value
func (o *CFCertificateResponse) GetCertificateSHAOneThumbprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateSHAOneThumbprint
}

// GetCertificateSHAOneThumbprintOk returns a tuple with the CertificateSHAOneThumbprint field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetCertificateSHAOneThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateSHAOneThumbprint, true
}

// SetCertificateSHAOneThumbprint sets field value
func (o *CFCertificateResponse) SetCertificateSHAOneThumbprint(v string) {
	o.CertificateSHAOneThumbprint = v
}

// GetCertificateThumbprint returns the CertificateThumbprint field value
func (o *CFCertificateResponse) GetCertificateThumbprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateThumbprint
}

// GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetCertificateThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateThumbprint, true
}

// SetCertificateThumbprint sets field value
func (o *CFCertificateResponse) SetCertificateThumbprint(v string) {
	o.CertificateThumbprint = v
}

// GetCrldps returns the Crldps field value if set, zero value otherwise.
func (o *CFCertificateResponse) GetCrldps() []string {
	if o == nil || utils.IsNil(o.Crldps) {
		var ret []string
		return ret
	}
	return o.Crldps
}

// GetCrldpsOk returns a tuple with the Crldps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetCrldpsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Crldps) {
		return nil, false
	}
	return o.Crldps, true
}

// HasCrldps returns a boolean if a field has been set.
func (o *CFCertificateResponse) HasCrldps() bool {
	if o != nil && !utils.IsNil(o.Crldps) {
		return true
	}

	return false
}

// SetCrldps gets a reference to the given []string and assigns it to the Crldps field.
func (o *CFCertificateResponse) SetCrldps(v []string) {
	o.Crldps = v
}

// GetDn returns the Dn field value
func (o *CFCertificateResponse) GetDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dn
}

// GetDnOk returns a tuple with the Dn field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dn, true
}

// SetDn sets field value
func (o *CFCertificateResponse) SetDn(v string) {
	o.Dn = v
}

// GetDnElements returns the DnElements field value
func (o *CFCertificateResponse) GetDnElements() []CFDistinguishedName {
	if o == nil {
		var ret []CFDistinguishedName
		return ret
	}

	return o.DnElements
}

// GetDnElementsOk returns a tuple with the DnElements field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetDnElementsOk() ([]CFDistinguishedName, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnElements, true
}

// SetDnElements sets field value
func (o *CFCertificateResponse) SetDnElements(v []CFDistinguishedName) {
	o.DnElements = v
}

// GetExtendedKeyUsages returns the ExtendedKeyUsages field value
func (o *CFCertificateResponse) GetExtendedKeyUsages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ExtendedKeyUsages
}

// GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetExtendedKeyUsagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtendedKeyUsages, true
}

// SetExtendedKeyUsages sets field value
func (o *CFCertificateResponse) SetExtendedKeyUsages(v []string) {
	o.ExtendedKeyUsages = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CFCertificateResponse) GetExtensions() []CertificateExtension {
	if o == nil {
		var ret []CertificateExtension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CFCertificateResponse) GetExtensionsOk() ([]CertificateExtension, bool) {
	if o == nil || utils.IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *CFCertificateResponse) HasExtensions() bool {
	if o != nil && !utils.IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []CertificateExtension and assigns it to the Extensions field.
func (o *CFCertificateResponse) SetExtensions(v []CertificateExtension) {
	o.Extensions = v
}

// GetIsExtendedKeyUsagesCritical returns the IsExtendedKeyUsagesCritical field value
func (o *CFCertificateResponse) GetIsExtendedKeyUsagesCritical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExtendedKeyUsagesCritical
}

// GetIsExtendedKeyUsagesCriticalOk returns a tuple with the IsExtendedKeyUsagesCritical field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetIsExtendedKeyUsagesCriticalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExtendedKeyUsagesCritical, true
}

// SetIsExtendedKeyUsagesCritical sets field value
func (o *CFCertificateResponse) SetIsExtendedKeyUsagesCritical(v bool) {
	o.IsExtendedKeyUsagesCritical = v
}

// GetIsKeyUsagesCritical returns the IsKeyUsagesCritical field value
func (o *CFCertificateResponse) GetIsKeyUsagesCritical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsKeyUsagesCritical
}

// GetIsKeyUsagesCriticalOk returns a tuple with the IsKeyUsagesCritical field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetIsKeyUsagesCriticalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsKeyUsagesCritical, true
}

// SetIsKeyUsagesCritical sets field value
func (o *CFCertificateResponse) SetIsKeyUsagesCritical(v bool) {
	o.IsKeyUsagesCritical = v
}

// GetIssuerDn returns the IssuerDn field value
func (o *CFCertificateResponse) GetIssuerDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerDn
}

// GetIssuerDnOk returns a tuple with the IssuerDn field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetIssuerDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerDn, true
}

// SetIssuerDn sets field value
func (o *CFCertificateResponse) SetIssuerDn(v string) {
	o.IssuerDn = v
}

// GetKeyType returns the KeyType field value
func (o *CFCertificateResponse) GetKeyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *CFCertificateResponse) SetKeyType(v string) {
	o.KeyType = v
}

// GetKeyUsages returns the KeyUsages field value
func (o *CFCertificateResponse) GetKeyUsages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.KeyUsages
}

// GetKeyUsagesOk returns a tuple with the KeyUsages field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetKeyUsagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyUsages, true
}

// SetKeyUsages sets field value
func (o *CFCertificateResponse) SetKeyUsages(v []string) {
	o.KeyUsages = v
}

// GetNotAfter returns the NotAfter field value
func (o *CFCertificateResponse) GetNotAfter() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetNotAfterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotAfter, true
}

// SetNotAfter sets field value
func (o *CFCertificateResponse) SetNotAfter(v int64) {
	o.NotAfter = v
}

// GetNotBefore returns the NotBefore field value
func (o *CFCertificateResponse) GetNotBefore() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetNotBeforeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotBefore, true
}

// SetNotBefore sets field value
func (o *CFCertificateResponse) SetNotBefore(v int64) {
	o.NotBefore = v
}

// GetPem returns the Pem field value
func (o *CFCertificateResponse) GetPem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pem
}

// GetPemOk returns a tuple with the Pem field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetPemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pem, true
}

// SetPem sets field value
func (o *CFCertificateResponse) SetPem(v string) {
	o.Pem = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *CFCertificateResponse) GetPolicies() []CFCertificateResponsePoliciesInner {
	if o == nil || utils.IsNil(o.Policies) {
		var ret []CFCertificateResponsePoliciesInner
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetPoliciesOk() ([]CFCertificateResponsePoliciesInner, bool) {
	if o == nil || utils.IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *CFCertificateResponse) HasPolicies() bool {
	if o != nil && !utils.IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []CFCertificateResponsePoliciesInner and assigns it to the Policies field.
func (o *CFCertificateResponse) SetPolicies(v []CFCertificateResponsePoliciesInner) {
	o.Policies = v
}

// GetPublicKeyThumbprint returns the PublicKeyThumbprint field value
func (o *CFCertificateResponse) GetPublicKeyThumbprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKeyThumbprint
}

// GetPublicKeyThumbprintOk returns a tuple with the PublicKeyThumbprint field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetPublicKeyThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKeyThumbprint, true
}

// SetPublicKeyThumbprint sets field value
func (o *CFCertificateResponse) SetPublicKeyThumbprint(v string) {
	o.PublicKeyThumbprint = v
}

// GetSans returns the Sans field value if set, zero value otherwise.
func (o *CFCertificateResponse) GetSans() []SubjectAlternateName {
	if o == nil || utils.IsNil(o.Sans) {
		var ret []SubjectAlternateName
		return ret
	}
	return o.Sans
}

// GetSansOk returns a tuple with the Sans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetSansOk() ([]SubjectAlternateName, bool) {
	if o == nil || utils.IsNil(o.Sans) {
		return nil, false
	}
	return o.Sans, true
}

// HasSans returns a boolean if a field has been set.
func (o *CFCertificateResponse) HasSans() bool {
	if o != nil && !utils.IsNil(o.Sans) {
		return true
	}

	return false
}

// SetSans gets a reference to the given []SubjectAlternateName and assigns it to the Sans field.
func (o *CFCertificateResponse) SetSans(v []SubjectAlternateName) {
	o.Sans = v
}

// GetSelfSigned returns the SelfSigned field value
func (o *CFCertificateResponse) GetSelfSigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SelfSigned
}

// GetSelfSignedOk returns a tuple with the SelfSigned field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetSelfSignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfSigned, true
}

// SetSelfSigned sets field value
func (o *CFCertificateResponse) SetSelfSigned(v bool) {
	o.SelfSigned = v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *CFCertificateResponse) GetSerial() string {
	if o == nil || utils.IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetSerialOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *CFCertificateResponse) HasSerial() bool {
	if o != nil && !utils.IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *CFCertificateResponse) SetSerial(v string) {
	o.Serial = &v
}

// GetSigningAlgorithm returns the SigningAlgorithm field value
func (o *CFCertificateResponse) GetSigningAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SigningAlgorithm
}

// GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetSigningAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningAlgorithm, true
}

// SetSigningAlgorithm sets field value
func (o *CFCertificateResponse) SetSigningAlgorithm(v string) {
	o.SigningAlgorithm = v
}

// GetSubjectKeyIdentifier returns the SubjectKeyIdentifier field value
func (o *CFCertificateResponse) GetSubjectKeyIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectKeyIdentifier
}

// GetSubjectKeyIdentifierOk returns a tuple with the SubjectKeyIdentifier field value
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetSubjectKeyIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectKeyIdentifier, true
}

// SetSubjectKeyIdentifier sets field value
func (o *CFCertificateResponse) SetSubjectKeyIdentifier(v string) {
	o.SubjectKeyIdentifier = v
}

// GetUnsupportedExtensions returns the UnsupportedExtensions field value if set, zero value otherwise.
func (o *CFCertificateResponse) GetUnsupportedExtensions() []CFCertificateResponseUnsupportedExtensionsInner {
	if o == nil || utils.IsNil(o.UnsupportedExtensions) {
		var ret []CFCertificateResponseUnsupportedExtensionsInner
		return ret
	}
	return o.UnsupportedExtensions
}

// GetUnsupportedExtensionsOk returns a tuple with the UnsupportedExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificateResponse) GetUnsupportedExtensionsOk() ([]CFCertificateResponseUnsupportedExtensionsInner, bool) {
	if o == nil || utils.IsNil(o.UnsupportedExtensions) {
		return nil, false
	}
	return o.UnsupportedExtensions, true
}

// HasUnsupportedExtensions returns a boolean if a field has been set.
func (o *CFCertificateResponse) HasUnsupportedExtensions() bool {
	if o != nil && !utils.IsNil(o.UnsupportedExtensions) {
		return true
	}

	return false
}

// SetUnsupportedExtensions gets a reference to the given []CFCertificateResponseUnsupportedExtensionsInner and assigns it to the UnsupportedExtensions field.
func (o *CFCertificateResponse) SetUnsupportedExtensions(v []CFCertificateResponseUnsupportedExtensionsInner) {
	o.UnsupportedExtensions = v
}

func (o CFCertificateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CFCertificateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Aias) {
		toSerialize["aias"] = o.Aias
	}
	if !utils.IsNil(o.AuthorityKeyIdentifier) {
		toSerialize["authorityKeyIdentifier"] = o.AuthorityKeyIdentifier
	}
	if !utils.IsNil(o.BasicConstraints) {
		toSerialize["basicConstraints"] = o.BasicConstraints
	}
	toSerialize["certificateSHAOneThumbprint"] = o.CertificateSHAOneThumbprint
	toSerialize["certificateThumbprint"] = o.CertificateThumbprint
	if !utils.IsNil(o.Crldps) {
		toSerialize["crldps"] = o.Crldps
	}
	toSerialize["dn"] = o.Dn
	toSerialize["dnElements"] = o.DnElements
	toSerialize["extendedKeyUsages"] = o.ExtendedKeyUsages
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	toSerialize["isExtendedKeyUsagesCritical"] = o.IsExtendedKeyUsagesCritical
	toSerialize["isKeyUsagesCritical"] = o.IsKeyUsagesCritical
	toSerialize["issuerDn"] = o.IssuerDn
	toSerialize["keyType"] = o.KeyType
	toSerialize["keyUsages"] = o.KeyUsages
	toSerialize["notAfter"] = o.NotAfter
	toSerialize["notBefore"] = o.NotBefore
	toSerialize["pem"] = o.Pem
	if !utils.IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	toSerialize["publicKeyThumbprint"] = o.PublicKeyThumbprint
	if !utils.IsNil(o.Sans) {
		toSerialize["sans"] = o.Sans
	}
	toSerialize["selfSigned"] = o.SelfSigned
	if !utils.IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	toSerialize["signingAlgorithm"] = o.SigningAlgorithm
	toSerialize["subjectKeyIdentifier"] = o.SubjectKeyIdentifier
	if !utils.IsNil(o.UnsupportedExtensions) {
		toSerialize["unsupportedExtensions"] = o.UnsupportedExtensions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CFCertificateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"certificateSHAOneThumbprint",
		"certificateThumbprint",
		"dn",
		"dnElements",
		"extendedKeyUsages",
		"isExtendedKeyUsagesCritical",
		"isKeyUsagesCritical",
		"issuerDn",
		"keyType",
		"keyUsages",
		"notAfter",
		"notBefore",
		"pem",
		"publicKeyThumbprint",
		"selfSigned",
		"signingAlgorithm",
		"subjectKeyIdentifier",
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

	varCFCertificateResponse := _CFCertificateResponse{}

	err = json.Unmarshal(data, &varCFCertificateResponse)

	if err != nil {
		return err
	}

	*o = CFCertificateResponse(varCFCertificateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aias")
		delete(additionalProperties, "authorityKeyIdentifier")
		delete(additionalProperties, "basicConstraints")
		delete(additionalProperties, "certificateSHAOneThumbprint")
		delete(additionalProperties, "certificateThumbprint")
		delete(additionalProperties, "crldps")
		delete(additionalProperties, "dn")
		delete(additionalProperties, "dnElements")
		delete(additionalProperties, "extendedKeyUsages")
		delete(additionalProperties, "extensions")
		delete(additionalProperties, "isExtendedKeyUsagesCritical")
		delete(additionalProperties, "isKeyUsagesCritical")
		delete(additionalProperties, "issuerDn")
		delete(additionalProperties, "keyType")
		delete(additionalProperties, "keyUsages")
		delete(additionalProperties, "notAfter")
		delete(additionalProperties, "notBefore")
		delete(additionalProperties, "pem")
		delete(additionalProperties, "policies")
		delete(additionalProperties, "publicKeyThumbprint")
		delete(additionalProperties, "sans")
		delete(additionalProperties, "selfSigned")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "signingAlgorithm")
		delete(additionalProperties, "subjectKeyIdentifier")
		delete(additionalProperties, "unsupportedExtensions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCFCertificateResponse struct {
	value *CFCertificateResponse
	isSet bool
}

func (v NullableCFCertificateResponse) Get() *CFCertificateResponse {
	return v.value
}

func (v *NullableCFCertificateResponse) Set(val *CFCertificateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCFCertificateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCFCertificateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFCertificateResponse(val *CFCertificateResponse) *NullableCFCertificateResponse {
	return &NullableCFCertificateResponse{value: val, isSet: true}
}

func (v NullableCFCertificateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFCertificateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
