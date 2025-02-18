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

// checks if the CFCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CFCertificate{}

// CFCertificate struct for CFCertificate
type CFCertificate struct {
	// The certificate's Distinguished Name
	Dn string `json:"dn"`
	DnElements []CFDistinguishedName `json:"dnElements"`
	// The certificate's issuer Distinguished Name
	IssuerDn string `json:"issuerDn"`
	// The certificate's serial number
	Serial *string `json:"serial,omitempty"`
	// The certificate's start date in milliseconds since the epoch
	NotBefore int64 `json:"notBefore"`
	// The certificate's expiration date in milliseconds since the epoch
	NotAfter int64 `json:"notAfter"`
	// The certificate's key type
	KeyType string `json:"keyType" validate:"regexp=(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87)(\\\\\\\\+(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87))?"`
	// The certificate's signing algorithm
	SigningAlgorithm string `json:"signingAlgorithm"`
	// The certificate's PEM-encoded content
	Pem string `json:"pem"`
	SubjectKeyIdentifier string `json:"subjectKeyIdentifier"`
	// The certificate's thumbprint
	CertificateThumbprint string `json:"certificateThumbprint"`
	// The thumbprint of the certificate using SHAOne algorithm
	CertificateSHAOneThumbprint string `json:"certificateSHAOneThumbprint"`
	// The certificate's public key thumbprint
	PublicKeyThumbprint string `json:"publicKeyThumbprint"`
	// The certificate key's usage
	KeyUsages []string `json:"keyUsages"`
	// If the key usage of the certificate are critical
	IsKeyUsagesCritical bool `json:"isKeyUsagesCritical"`
	// The certificate extended key's usage
	ExtendedKeyUsages []string `json:"extendedKeyUsages"`
	// If the extended key usage are critical
	IsExtendedKeyUsagesCritical bool `json:"isExtendedKeyUsagesCritical"`
	// Whether the certificate is self-signed
	SelfSigned bool `json:"selfSigned"`
	// The certificate's SAN
	Sans []SubjectAlternateName `json:"sans,omitempty"`
	BasicConstraints CFCertificateBasicConstraints `json:"basicConstraints"`
	// The certificate's extensions
	Extensions []CertificateExtension `json:"extensions,omitempty"`
	// The certificate's CRLDP if any
	Crldps []string `json:"crldps,omitempty"`
	Aias *CFCertificateAias `json:"aias,omitempty"`
	Policies []CFCertificatePoliciesInner `json:"policies,omitempty"`
	// The certificate AKI
	AuthorityKeyIdentifier *string `json:"authorityKeyIdentifier,omitempty"`
	UnsupportedExtensions []CFCertificateUnsupportedExtensionsInner `json:"unsupportedExtensions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CFCertificate CFCertificate

// NewCFCertificate instantiates a new CFCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFCertificate(dn string, dnElements []CFDistinguishedName, issuerDn string, notBefore int64, notAfter int64, keyType string, signingAlgorithm string, pem string, subjectKeyIdentifier string, certificateThumbprint string, certificateSHAOneThumbprint string, publicKeyThumbprint string, keyUsages []string, isKeyUsagesCritical bool, extendedKeyUsages []string, isExtendedKeyUsagesCritical bool, selfSigned bool, basicConstraints CFCertificateBasicConstraints) *CFCertificate {
	this := CFCertificate{}
	this.Dn = dn
	this.DnElements = dnElements
	this.IssuerDn = issuerDn
	this.NotBefore = notBefore
	this.NotAfter = notAfter
	this.KeyType = keyType
	this.SigningAlgorithm = signingAlgorithm
	this.Pem = pem
	this.SubjectKeyIdentifier = subjectKeyIdentifier
	this.CertificateThumbprint = certificateThumbprint
	this.CertificateSHAOneThumbprint = certificateSHAOneThumbprint
	this.PublicKeyThumbprint = publicKeyThumbprint
	this.KeyUsages = keyUsages
	this.IsKeyUsagesCritical = isKeyUsagesCritical
	this.ExtendedKeyUsages = extendedKeyUsages
	this.IsExtendedKeyUsagesCritical = isExtendedKeyUsagesCritical
	this.SelfSigned = selfSigned
	this.BasicConstraints = basicConstraints
	return &this
}

// NewCFCertificateWithDefaults instantiates a new CFCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFCertificateWithDefaults() *CFCertificate {
	this := CFCertificate{}
	return &this
}

// GetDn returns the Dn field value
func (o *CFCertificate) GetDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dn
}

// GetDnOk returns a tuple with the Dn field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dn, true
}

// SetDn sets field value
func (o *CFCertificate) SetDn(v string) {
	o.Dn = v
}

// GetDnElements returns the DnElements field value
func (o *CFCertificate) GetDnElements() []CFDistinguishedName {
	if o == nil {
		var ret []CFDistinguishedName
		return ret
	}

	return o.DnElements
}

// GetDnElementsOk returns a tuple with the DnElements field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetDnElementsOk() ([]CFDistinguishedName, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnElements, true
}

// SetDnElements sets field value
func (o *CFCertificate) SetDnElements(v []CFDistinguishedName) {
	o.DnElements = v
}

// GetIssuerDn returns the IssuerDn field value
func (o *CFCertificate) GetIssuerDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerDn
}

// GetIssuerDnOk returns a tuple with the IssuerDn field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetIssuerDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerDn, true
}

// SetIssuerDn sets field value
func (o *CFCertificate) SetIssuerDn(v string) {
	o.IssuerDn = v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *CFCertificate) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *CFCertificate) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *CFCertificate) SetSerial(v string) {
	o.Serial = &v
}

// GetNotBefore returns the NotBefore field value
func (o *CFCertificate) GetNotBefore() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetNotBeforeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotBefore, true
}

// SetNotBefore sets field value
func (o *CFCertificate) SetNotBefore(v int64) {
	o.NotBefore = v
}

// GetNotAfter returns the NotAfter field value
func (o *CFCertificate) GetNotAfter() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetNotAfterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotAfter, true
}

// SetNotAfter sets field value
func (o *CFCertificate) SetNotAfter(v int64) {
	o.NotAfter = v
}

// GetKeyType returns the KeyType field value
func (o *CFCertificate) GetKeyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *CFCertificate) SetKeyType(v string) {
	o.KeyType = v
}

// GetSigningAlgorithm returns the SigningAlgorithm field value
func (o *CFCertificate) GetSigningAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SigningAlgorithm
}

// GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetSigningAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SigningAlgorithm, true
}

// SetSigningAlgorithm sets field value
func (o *CFCertificate) SetSigningAlgorithm(v string) {
	o.SigningAlgorithm = v
}

// GetPem returns the Pem field value
func (o *CFCertificate) GetPem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pem
}

// GetPemOk returns a tuple with the Pem field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetPemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pem, true
}

// SetPem sets field value
func (o *CFCertificate) SetPem(v string) {
	o.Pem = v
}

// GetSubjectKeyIdentifier returns the SubjectKeyIdentifier field value
func (o *CFCertificate) GetSubjectKeyIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectKeyIdentifier
}

// GetSubjectKeyIdentifierOk returns a tuple with the SubjectKeyIdentifier field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetSubjectKeyIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectKeyIdentifier, true
}

// SetSubjectKeyIdentifier sets field value
func (o *CFCertificate) SetSubjectKeyIdentifier(v string) {
	o.SubjectKeyIdentifier = v
}

// GetCertificateThumbprint returns the CertificateThumbprint field value
func (o *CFCertificate) GetCertificateThumbprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateThumbprint
}

// GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetCertificateThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateThumbprint, true
}

// SetCertificateThumbprint sets field value
func (o *CFCertificate) SetCertificateThumbprint(v string) {
	o.CertificateThumbprint = v
}

// GetCertificateSHAOneThumbprint returns the CertificateSHAOneThumbprint field value
func (o *CFCertificate) GetCertificateSHAOneThumbprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateSHAOneThumbprint
}

// GetCertificateSHAOneThumbprintOk returns a tuple with the CertificateSHAOneThumbprint field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetCertificateSHAOneThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateSHAOneThumbprint, true
}

// SetCertificateSHAOneThumbprint sets field value
func (o *CFCertificate) SetCertificateSHAOneThumbprint(v string) {
	o.CertificateSHAOneThumbprint = v
}

// GetPublicKeyThumbprint returns the PublicKeyThumbprint field value
func (o *CFCertificate) GetPublicKeyThumbprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKeyThumbprint
}

// GetPublicKeyThumbprintOk returns a tuple with the PublicKeyThumbprint field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetPublicKeyThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKeyThumbprint, true
}

// SetPublicKeyThumbprint sets field value
func (o *CFCertificate) SetPublicKeyThumbprint(v string) {
	o.PublicKeyThumbprint = v
}

// GetKeyUsages returns the KeyUsages field value
func (o *CFCertificate) GetKeyUsages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.KeyUsages
}

// GetKeyUsagesOk returns a tuple with the KeyUsages field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetKeyUsagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyUsages, true
}

// SetKeyUsages sets field value
func (o *CFCertificate) SetKeyUsages(v []string) {
	o.KeyUsages = v
}

// GetIsKeyUsagesCritical returns the IsKeyUsagesCritical field value
func (o *CFCertificate) GetIsKeyUsagesCritical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsKeyUsagesCritical
}

// GetIsKeyUsagesCriticalOk returns a tuple with the IsKeyUsagesCritical field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetIsKeyUsagesCriticalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsKeyUsagesCritical, true
}

// SetIsKeyUsagesCritical sets field value
func (o *CFCertificate) SetIsKeyUsagesCritical(v bool) {
	o.IsKeyUsagesCritical = v
}

// GetExtendedKeyUsages returns the ExtendedKeyUsages field value
func (o *CFCertificate) GetExtendedKeyUsages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ExtendedKeyUsages
}

// GetExtendedKeyUsagesOk returns a tuple with the ExtendedKeyUsages field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetExtendedKeyUsagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtendedKeyUsages, true
}

// SetExtendedKeyUsages sets field value
func (o *CFCertificate) SetExtendedKeyUsages(v []string) {
	o.ExtendedKeyUsages = v
}

// GetIsExtendedKeyUsagesCritical returns the IsExtendedKeyUsagesCritical field value
func (o *CFCertificate) GetIsExtendedKeyUsagesCritical() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsExtendedKeyUsagesCritical
}

// GetIsExtendedKeyUsagesCriticalOk returns a tuple with the IsExtendedKeyUsagesCritical field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetIsExtendedKeyUsagesCriticalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsExtendedKeyUsagesCritical, true
}

// SetIsExtendedKeyUsagesCritical sets field value
func (o *CFCertificate) SetIsExtendedKeyUsagesCritical(v bool) {
	o.IsExtendedKeyUsagesCritical = v
}

// GetSelfSigned returns the SelfSigned field value
func (o *CFCertificate) GetSelfSigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SelfSigned
}

// GetSelfSignedOk returns a tuple with the SelfSigned field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetSelfSignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelfSigned, true
}

// SetSelfSigned sets field value
func (o *CFCertificate) SetSelfSigned(v bool) {
	o.SelfSigned = v
}

// GetSans returns the Sans field value if set, zero value otherwise.
func (o *CFCertificate) GetSans() []SubjectAlternateName {
	if o == nil || IsNil(o.Sans) {
		var ret []SubjectAlternateName
		return ret
	}
	return o.Sans
}

// GetSansOk returns a tuple with the Sans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetSansOk() ([]SubjectAlternateName, bool) {
	if o == nil || IsNil(o.Sans) {
		return nil, false
	}
	return o.Sans, true
}

// HasSans returns a boolean if a field has been set.
func (o *CFCertificate) HasSans() bool {
	if o != nil && !IsNil(o.Sans) {
		return true
	}

	return false
}

// SetSans gets a reference to the given []SubjectAlternateName and assigns it to the Sans field.
func (o *CFCertificate) SetSans(v []SubjectAlternateName) {
	o.Sans = v
}

// GetBasicConstraints returns the BasicConstraints field value
func (o *CFCertificate) GetBasicConstraints() CFCertificateBasicConstraints {
	if o == nil {
		var ret CFCertificateBasicConstraints
		return ret
	}

	return o.BasicConstraints
}

// GetBasicConstraintsOk returns a tuple with the BasicConstraints field value
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetBasicConstraintsOk() (*CFCertificateBasicConstraints, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasicConstraints, true
}

// SetBasicConstraints sets field value
func (o *CFCertificate) SetBasicConstraints(v CFCertificateBasicConstraints) {
	o.BasicConstraints = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CFCertificate) GetExtensions() []CertificateExtension {
	if o == nil {
		var ret []CertificateExtension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CFCertificate) GetExtensionsOk() ([]CertificateExtension, bool) {
	if o == nil || IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *CFCertificate) HasExtensions() bool {
	if o != nil && !IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []CertificateExtension and assigns it to the Extensions field.
func (o *CFCertificate) SetExtensions(v []CertificateExtension) {
	o.Extensions = v
}

// GetCrldps returns the Crldps field value if set, zero value otherwise.
func (o *CFCertificate) GetCrldps() []string {
	if o == nil || IsNil(o.Crldps) {
		var ret []string
		return ret
	}
	return o.Crldps
}

// GetCrldpsOk returns a tuple with the Crldps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetCrldpsOk() ([]string, bool) {
	if o == nil || IsNil(o.Crldps) {
		return nil, false
	}
	return o.Crldps, true
}

// HasCrldps returns a boolean if a field has been set.
func (o *CFCertificate) HasCrldps() bool {
	if o != nil && !IsNil(o.Crldps) {
		return true
	}

	return false
}

// SetCrldps gets a reference to the given []string and assigns it to the Crldps field.
func (o *CFCertificate) SetCrldps(v []string) {
	o.Crldps = v
}

// GetAias returns the Aias field value if set, zero value otherwise.
func (o *CFCertificate) GetAias() CFCertificateAias {
	if o == nil || IsNil(o.Aias) {
		var ret CFCertificateAias
		return ret
	}
	return *o.Aias
}

// GetAiasOk returns a tuple with the Aias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetAiasOk() (*CFCertificateAias, bool) {
	if o == nil || IsNil(o.Aias) {
		return nil, false
	}
	return o.Aias, true
}

// HasAias returns a boolean if a field has been set.
func (o *CFCertificate) HasAias() bool {
	if o != nil && !IsNil(o.Aias) {
		return true
	}

	return false
}

// SetAias gets a reference to the given CFCertificateAias and assigns it to the Aias field.
func (o *CFCertificate) SetAias(v CFCertificateAias) {
	o.Aias = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *CFCertificate) GetPolicies() []CFCertificatePoliciesInner {
	if o == nil || IsNil(o.Policies) {
		var ret []CFCertificatePoliciesInner
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetPoliciesOk() ([]CFCertificatePoliciesInner, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *CFCertificate) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []CFCertificatePoliciesInner and assigns it to the Policies field.
func (o *CFCertificate) SetPolicies(v []CFCertificatePoliciesInner) {
	o.Policies = v
}

// GetAuthorityKeyIdentifier returns the AuthorityKeyIdentifier field value if set, zero value otherwise.
func (o *CFCertificate) GetAuthorityKeyIdentifier() string {
	if o == nil || IsNil(o.AuthorityKeyIdentifier) {
		var ret string
		return ret
	}
	return *o.AuthorityKeyIdentifier
}

// GetAuthorityKeyIdentifierOk returns a tuple with the AuthorityKeyIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetAuthorityKeyIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorityKeyIdentifier) {
		return nil, false
	}
	return o.AuthorityKeyIdentifier, true
}

// HasAuthorityKeyIdentifier returns a boolean if a field has been set.
func (o *CFCertificate) HasAuthorityKeyIdentifier() bool {
	if o != nil && !IsNil(o.AuthorityKeyIdentifier) {
		return true
	}

	return false
}

// SetAuthorityKeyIdentifier gets a reference to the given string and assigns it to the AuthorityKeyIdentifier field.
func (o *CFCertificate) SetAuthorityKeyIdentifier(v string) {
	o.AuthorityKeyIdentifier = &v
}

// GetUnsupportedExtensions returns the UnsupportedExtensions field value if set, zero value otherwise.
func (o *CFCertificate) GetUnsupportedExtensions() []CFCertificateUnsupportedExtensionsInner {
	if o == nil || IsNil(o.UnsupportedExtensions) {
		var ret []CFCertificateUnsupportedExtensionsInner
		return ret
	}
	return o.UnsupportedExtensions
}

// GetUnsupportedExtensionsOk returns a tuple with the UnsupportedExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificate) GetUnsupportedExtensionsOk() ([]CFCertificateUnsupportedExtensionsInner, bool) {
	if o == nil || IsNil(o.UnsupportedExtensions) {
		return nil, false
	}
	return o.UnsupportedExtensions, true
}

// HasUnsupportedExtensions returns a boolean if a field has been set.
func (o *CFCertificate) HasUnsupportedExtensions() bool {
	if o != nil && !IsNil(o.UnsupportedExtensions) {
		return true
	}

	return false
}

// SetUnsupportedExtensions gets a reference to the given []CFCertificateUnsupportedExtensionsInner and assigns it to the UnsupportedExtensions field.
func (o *CFCertificate) SetUnsupportedExtensions(v []CFCertificateUnsupportedExtensionsInner) {
	o.UnsupportedExtensions = v
}

func (o CFCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CFCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dn"] = o.Dn
	toSerialize["dnElements"] = o.DnElements
	toSerialize["issuerDn"] = o.IssuerDn
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	toSerialize["notBefore"] = o.NotBefore
	toSerialize["notAfter"] = o.NotAfter
	toSerialize["keyType"] = o.KeyType
	toSerialize["signingAlgorithm"] = o.SigningAlgorithm
	toSerialize["pem"] = o.Pem
	toSerialize["subjectKeyIdentifier"] = o.SubjectKeyIdentifier
	toSerialize["certificateThumbprint"] = o.CertificateThumbprint
	toSerialize["certificateSHAOneThumbprint"] = o.CertificateSHAOneThumbprint
	toSerialize["publicKeyThumbprint"] = o.PublicKeyThumbprint
	toSerialize["keyUsages"] = o.KeyUsages
	toSerialize["isKeyUsagesCritical"] = o.IsKeyUsagesCritical
	toSerialize["extendedKeyUsages"] = o.ExtendedKeyUsages
	toSerialize["isExtendedKeyUsagesCritical"] = o.IsExtendedKeyUsagesCritical
	toSerialize["selfSigned"] = o.SelfSigned
	if !IsNil(o.Sans) {
		toSerialize["sans"] = o.Sans
	}
	toSerialize["basicConstraints"] = o.BasicConstraints
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if !IsNil(o.Crldps) {
		toSerialize["crldps"] = o.Crldps
	}
	if !IsNil(o.Aias) {
		toSerialize["aias"] = o.Aias
	}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	if !IsNil(o.AuthorityKeyIdentifier) {
		toSerialize["authorityKeyIdentifier"] = o.AuthorityKeyIdentifier
	}
	if !IsNil(o.UnsupportedExtensions) {
		toSerialize["unsupportedExtensions"] = o.UnsupportedExtensions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CFCertificate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dn",
		"dnElements",
		"issuerDn",
		"notBefore",
		"notAfter",
		"keyType",
		"signingAlgorithm",
		"pem",
		"subjectKeyIdentifier",
		"certificateThumbprint",
		"certificateSHAOneThumbprint",
		"publicKeyThumbprint",
		"keyUsages",
		"isKeyUsagesCritical",
		"extendedKeyUsages",
		"isExtendedKeyUsagesCritical",
		"selfSigned",
		"basicConstraints",
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

	varCFCertificate := _CFCertificate{}

	err = json.Unmarshal(data, &varCFCertificate)

	if err != nil {
		return err
	}

	*o = CFCertificate(varCFCertificate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dn")
		delete(additionalProperties, "dnElements")
		delete(additionalProperties, "issuerDn")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "notBefore")
		delete(additionalProperties, "notAfter")
		delete(additionalProperties, "keyType")
		delete(additionalProperties, "signingAlgorithm")
		delete(additionalProperties, "pem")
		delete(additionalProperties, "subjectKeyIdentifier")
		delete(additionalProperties, "certificateThumbprint")
		delete(additionalProperties, "certificateSHAOneThumbprint")
		delete(additionalProperties, "publicKeyThumbprint")
		delete(additionalProperties, "keyUsages")
		delete(additionalProperties, "isKeyUsagesCritical")
		delete(additionalProperties, "extendedKeyUsages")
		delete(additionalProperties, "isExtendedKeyUsagesCritical")
		delete(additionalProperties, "selfSigned")
		delete(additionalProperties, "sans")
		delete(additionalProperties, "basicConstraints")
		delete(additionalProperties, "extensions")
		delete(additionalProperties, "crldps")
		delete(additionalProperties, "aias")
		delete(additionalProperties, "policies")
		delete(additionalProperties, "authorityKeyIdentifier")
		delete(additionalProperties, "unsupportedExtensions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCFCertificate struct {
	value *CFCertificate
	isSet bool
}

func (v NullableCFCertificate) Get() *CFCertificate {
	return v.value
}

func (v *NullableCFCertificate) Set(val *CFCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCFCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCFCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFCertificate(val *CFCertificate) *NullableCFCertificate {
	return &NullableCFCertificate{value: val, isSet: true}
}

func (v NullableCFCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


