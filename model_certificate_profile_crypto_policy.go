/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"encoding/json"
)

// checks if the CertificateProfileCryptoPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateProfileCryptoPolicy{}

// CertificateProfileCryptoPolicy struct for CertificateProfileCryptoPolicy
type CertificateProfileCryptoPolicy struct {
	// Whether this profile supports centralized enrollment
	Centralized NullableBool `json:"centralized,omitempty"`
	// Whether this profile supports decentralized enrollment
	Decentralized NullableBool `json:"decentralized,omitempty"`
	// Default key type used for centralized enrollment
	DefaultKeyType NullableString `json:"defaultKeyType,omitempty" validate:"regexp=(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87)(\\\\\\\\+(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87))?"`
	// List of authorized key types for enrollment
	AuthorizedKeyTypes []string `json:"authorizedKeyTypes,omitempty"`
	// If both centralized and decentralized enrollment are supported, this is the preferred mode
	PreferredEnrollmentMode NullableString `json:"preferredEnrollmentMode,omitempty"`
	// Whether this profile will escrow the certificate private keys
	Escrow NullableBool `json:"escrow,omitempty"`
	// Password policy for the P12 file
	P12passwordPolicy NullableString `json:"p12passwordPolicy,omitempty"`
	// Whether the user will be required to input their PKCS#12 password upon enrollment
	P12passwordMode NullableString `json:"p12passwordMode,omitempty"`
	// Encryption type for the P12 file
	P12storeEncryptionType NullableString `json:"p12storeEncryptionType,omitempty"`
	// Whether the PKCS#12 password will be displayed to the user upon enrollment
	ShowP12PasswordOnEnroll NullableBool `json:"showP12PasswordOnEnroll,omitempty"`
	// Whether the PKCS#12 file will be displayed to the user upon enrollment
	ShowP12OnEnroll NullableBool `json:"showP12OnEnroll,omitempty"`
	// Whether the PKCS#12 password will be displayed to the user upon recovery
	ShowP12PasswordOnRecover NullableBool `json:"showP12PasswordOnRecover,omitempty"`
	// Whether the PKCS#12 file will be displayed to the user upon recovery
	ShowP12OnRecover NullableBool `json:"showP12OnRecover,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateProfileCryptoPolicy CertificateProfileCryptoPolicy

// NewCertificateProfileCryptoPolicy instantiates a new CertificateProfileCryptoPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateProfileCryptoPolicy() *CertificateProfileCryptoPolicy {
	this := CertificateProfileCryptoPolicy{}
	var centralized bool = false
	this.Centralized = *NewNullableBool(&centralized)
	var decentralized bool = false
	this.Decentralized = *NewNullableBool(&decentralized)
	var escrow bool = false
	this.Escrow = *NewNullableBool(&escrow)
	return &this
}

// NewCertificateProfileCryptoPolicyWithDefaults instantiates a new CertificateProfileCryptoPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateProfileCryptoPolicyWithDefaults() *CertificateProfileCryptoPolicy {
	this := CertificateProfileCryptoPolicy{}
	var centralized bool = false
	this.Centralized = *NewNullableBool(&centralized)
	var decentralized bool = false
	this.Decentralized = *NewNullableBool(&decentralized)
	var escrow bool = false
	this.Escrow = *NewNullableBool(&escrow)
	return &this
}

// GetCentralized returns the Centralized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetCentralized() bool {
	if o == nil || IsNil(o.Centralized.Get()) {
		var ret bool
		return ret
	}
	return *o.Centralized.Get()
}

// GetCentralizedOk returns a tuple with the Centralized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetCentralizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Centralized.Get(), o.Centralized.IsSet()
}

// HasCentralized returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasCentralized() bool {
	if o != nil && o.Centralized.IsSet() {
		return true
	}

	return false
}

// SetCentralized gets a reference to the given NullableBool and assigns it to the Centralized field.
func (o *CertificateProfileCryptoPolicy) SetCentralized(v bool) {
	o.Centralized.Set(&v)
}
// SetCentralizedNil sets the value for Centralized to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetCentralizedNil() {
	o.Centralized.Set(nil)
}

// UnsetCentralized ensures that no value is present for Centralized, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetCentralized() {
	o.Centralized.Unset()
}

// GetDecentralized returns the Decentralized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetDecentralized() bool {
	if o == nil || IsNil(o.Decentralized.Get()) {
		var ret bool
		return ret
	}
	return *o.Decentralized.Get()
}

// GetDecentralizedOk returns a tuple with the Decentralized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetDecentralizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Decentralized.Get(), o.Decentralized.IsSet()
}

// HasDecentralized returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasDecentralized() bool {
	if o != nil && o.Decentralized.IsSet() {
		return true
	}

	return false
}

// SetDecentralized gets a reference to the given NullableBool and assigns it to the Decentralized field.
func (o *CertificateProfileCryptoPolicy) SetDecentralized(v bool) {
	o.Decentralized.Set(&v)
}
// SetDecentralizedNil sets the value for Decentralized to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetDecentralizedNil() {
	o.Decentralized.Set(nil)
}

// UnsetDecentralized ensures that no value is present for Decentralized, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetDecentralized() {
	o.Decentralized.Unset()
}

// GetDefaultKeyType returns the DefaultKeyType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetDefaultKeyType() string {
	if o == nil || IsNil(o.DefaultKeyType.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultKeyType.Get()
}

// GetDefaultKeyTypeOk returns a tuple with the DefaultKeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetDefaultKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultKeyType.Get(), o.DefaultKeyType.IsSet()
}

// HasDefaultKeyType returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasDefaultKeyType() bool {
	if o != nil && o.DefaultKeyType.IsSet() {
		return true
	}

	return false
}

// SetDefaultKeyType gets a reference to the given NullableString and assigns it to the DefaultKeyType field.
func (o *CertificateProfileCryptoPolicy) SetDefaultKeyType(v string) {
	o.DefaultKeyType.Set(&v)
}
// SetDefaultKeyTypeNil sets the value for DefaultKeyType to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetDefaultKeyTypeNil() {
	o.DefaultKeyType.Set(nil)
}

// UnsetDefaultKeyType ensures that no value is present for DefaultKeyType, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetDefaultKeyType() {
	o.DefaultKeyType.Unset()
}

// GetAuthorizedKeyTypes returns the AuthorizedKeyTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetAuthorizedKeyTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AuthorizedKeyTypes
}

// GetAuthorizedKeyTypesOk returns a tuple with the AuthorizedKeyTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetAuthorizedKeyTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthorizedKeyTypes) {
		return nil, false
	}
	return o.AuthorizedKeyTypes, true
}

// HasAuthorizedKeyTypes returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasAuthorizedKeyTypes() bool {
	if o != nil && !IsNil(o.AuthorizedKeyTypes) {
		return true
	}

	return false
}

// SetAuthorizedKeyTypes gets a reference to the given []string and assigns it to the AuthorizedKeyTypes field.
func (o *CertificateProfileCryptoPolicy) SetAuthorizedKeyTypes(v []string) {
	o.AuthorizedKeyTypes = v
}

// GetPreferredEnrollmentMode returns the PreferredEnrollmentMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetPreferredEnrollmentMode() string {
	if o == nil || IsNil(o.PreferredEnrollmentMode.Get()) {
		var ret string
		return ret
	}
	return *o.PreferredEnrollmentMode.Get()
}

// GetPreferredEnrollmentModeOk returns a tuple with the PreferredEnrollmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetPreferredEnrollmentModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferredEnrollmentMode.Get(), o.PreferredEnrollmentMode.IsSet()
}

// HasPreferredEnrollmentMode returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasPreferredEnrollmentMode() bool {
	if o != nil && o.PreferredEnrollmentMode.IsSet() {
		return true
	}

	return false
}

// SetPreferredEnrollmentMode gets a reference to the given NullableString and assigns it to the PreferredEnrollmentMode field.
func (o *CertificateProfileCryptoPolicy) SetPreferredEnrollmentMode(v string) {
	o.PreferredEnrollmentMode.Set(&v)
}
// SetPreferredEnrollmentModeNil sets the value for PreferredEnrollmentMode to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetPreferredEnrollmentModeNil() {
	o.PreferredEnrollmentMode.Set(nil)
}

// UnsetPreferredEnrollmentMode ensures that no value is present for PreferredEnrollmentMode, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetPreferredEnrollmentMode() {
	o.PreferredEnrollmentMode.Unset()
}

// GetEscrow returns the Escrow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetEscrow() bool {
	if o == nil || IsNil(o.Escrow.Get()) {
		var ret bool
		return ret
	}
	return *o.Escrow.Get()
}

// GetEscrowOk returns a tuple with the Escrow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetEscrowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Escrow.Get(), o.Escrow.IsSet()
}

// HasEscrow returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasEscrow() bool {
	if o != nil && o.Escrow.IsSet() {
		return true
	}

	return false
}

// SetEscrow gets a reference to the given NullableBool and assigns it to the Escrow field.
func (o *CertificateProfileCryptoPolicy) SetEscrow(v bool) {
	o.Escrow.Set(&v)
}
// SetEscrowNil sets the value for Escrow to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetEscrowNil() {
	o.Escrow.Set(nil)
}

// UnsetEscrow ensures that no value is present for Escrow, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetEscrow() {
	o.Escrow.Unset()
}

// GetP12passwordPolicy returns the P12passwordPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetP12passwordPolicy() string {
	if o == nil || IsNil(o.P12passwordPolicy.Get()) {
		var ret string
		return ret
	}
	return *o.P12passwordPolicy.Get()
}

// GetP12passwordPolicyOk returns a tuple with the P12passwordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetP12passwordPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.P12passwordPolicy.Get(), o.P12passwordPolicy.IsSet()
}

// HasP12passwordPolicy returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasP12passwordPolicy() bool {
	if o != nil && o.P12passwordPolicy.IsSet() {
		return true
	}

	return false
}

// SetP12passwordPolicy gets a reference to the given NullableString and assigns it to the P12passwordPolicy field.
func (o *CertificateProfileCryptoPolicy) SetP12passwordPolicy(v string) {
	o.P12passwordPolicy.Set(&v)
}
// SetP12passwordPolicyNil sets the value for P12passwordPolicy to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetP12passwordPolicyNil() {
	o.P12passwordPolicy.Set(nil)
}

// UnsetP12passwordPolicy ensures that no value is present for P12passwordPolicy, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetP12passwordPolicy() {
	o.P12passwordPolicy.Unset()
}

// GetP12passwordMode returns the P12passwordMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetP12passwordMode() string {
	if o == nil || IsNil(o.P12passwordMode.Get()) {
		var ret string
		return ret
	}
	return *o.P12passwordMode.Get()
}

// GetP12passwordModeOk returns a tuple with the P12passwordMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetP12passwordModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.P12passwordMode.Get(), o.P12passwordMode.IsSet()
}

// HasP12passwordMode returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasP12passwordMode() bool {
	if o != nil && o.P12passwordMode.IsSet() {
		return true
	}

	return false
}

// SetP12passwordMode gets a reference to the given NullableString and assigns it to the P12passwordMode field.
func (o *CertificateProfileCryptoPolicy) SetP12passwordMode(v string) {
	o.P12passwordMode.Set(&v)
}
// SetP12passwordModeNil sets the value for P12passwordMode to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetP12passwordModeNil() {
	o.P12passwordMode.Set(nil)
}

// UnsetP12passwordMode ensures that no value is present for P12passwordMode, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetP12passwordMode() {
	o.P12passwordMode.Unset()
}

// GetP12storeEncryptionType returns the P12storeEncryptionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetP12storeEncryptionType() string {
	if o == nil || IsNil(o.P12storeEncryptionType.Get()) {
		var ret string
		return ret
	}
	return *o.P12storeEncryptionType.Get()
}

// GetP12storeEncryptionTypeOk returns a tuple with the P12storeEncryptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetP12storeEncryptionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.P12storeEncryptionType.Get(), o.P12storeEncryptionType.IsSet()
}

// HasP12storeEncryptionType returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasP12storeEncryptionType() bool {
	if o != nil && o.P12storeEncryptionType.IsSet() {
		return true
	}

	return false
}

// SetP12storeEncryptionType gets a reference to the given NullableString and assigns it to the P12storeEncryptionType field.
func (o *CertificateProfileCryptoPolicy) SetP12storeEncryptionType(v string) {
	o.P12storeEncryptionType.Set(&v)
}
// SetP12storeEncryptionTypeNil sets the value for P12storeEncryptionType to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetP12storeEncryptionTypeNil() {
	o.P12storeEncryptionType.Set(nil)
}

// UnsetP12storeEncryptionType ensures that no value is present for P12storeEncryptionType, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetP12storeEncryptionType() {
	o.P12storeEncryptionType.Unset()
}

// GetShowP12PasswordOnEnroll returns the ShowP12PasswordOnEnroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetShowP12PasswordOnEnroll() bool {
	if o == nil || IsNil(o.ShowP12PasswordOnEnroll.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowP12PasswordOnEnroll.Get()
}

// GetShowP12PasswordOnEnrollOk returns a tuple with the ShowP12PasswordOnEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetShowP12PasswordOnEnrollOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowP12PasswordOnEnroll.Get(), o.ShowP12PasswordOnEnroll.IsSet()
}

// HasShowP12PasswordOnEnroll returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasShowP12PasswordOnEnroll() bool {
	if o != nil && o.ShowP12PasswordOnEnroll.IsSet() {
		return true
	}

	return false
}

// SetShowP12PasswordOnEnroll gets a reference to the given NullableBool and assigns it to the ShowP12PasswordOnEnroll field.
func (o *CertificateProfileCryptoPolicy) SetShowP12PasswordOnEnroll(v bool) {
	o.ShowP12PasswordOnEnroll.Set(&v)
}
// SetShowP12PasswordOnEnrollNil sets the value for ShowP12PasswordOnEnroll to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetShowP12PasswordOnEnrollNil() {
	o.ShowP12PasswordOnEnroll.Set(nil)
}

// UnsetShowP12PasswordOnEnroll ensures that no value is present for ShowP12PasswordOnEnroll, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetShowP12PasswordOnEnroll() {
	o.ShowP12PasswordOnEnroll.Unset()
}

// GetShowP12OnEnroll returns the ShowP12OnEnroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetShowP12OnEnroll() bool {
	if o == nil || IsNil(o.ShowP12OnEnroll.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowP12OnEnroll.Get()
}

// GetShowP12OnEnrollOk returns a tuple with the ShowP12OnEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetShowP12OnEnrollOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowP12OnEnroll.Get(), o.ShowP12OnEnroll.IsSet()
}

// HasShowP12OnEnroll returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasShowP12OnEnroll() bool {
	if o != nil && o.ShowP12OnEnroll.IsSet() {
		return true
	}

	return false
}

// SetShowP12OnEnroll gets a reference to the given NullableBool and assigns it to the ShowP12OnEnroll field.
func (o *CertificateProfileCryptoPolicy) SetShowP12OnEnroll(v bool) {
	o.ShowP12OnEnroll.Set(&v)
}
// SetShowP12OnEnrollNil sets the value for ShowP12OnEnroll to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetShowP12OnEnrollNil() {
	o.ShowP12OnEnroll.Set(nil)
}

// UnsetShowP12OnEnroll ensures that no value is present for ShowP12OnEnroll, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetShowP12OnEnroll() {
	o.ShowP12OnEnroll.Unset()
}

// GetShowP12PasswordOnRecover returns the ShowP12PasswordOnRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetShowP12PasswordOnRecover() bool {
	if o == nil || IsNil(o.ShowP12PasswordOnRecover.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowP12PasswordOnRecover.Get()
}

// GetShowP12PasswordOnRecoverOk returns a tuple with the ShowP12PasswordOnRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetShowP12PasswordOnRecoverOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowP12PasswordOnRecover.Get(), o.ShowP12PasswordOnRecover.IsSet()
}

// HasShowP12PasswordOnRecover returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasShowP12PasswordOnRecover() bool {
	if o != nil && o.ShowP12PasswordOnRecover.IsSet() {
		return true
	}

	return false
}

// SetShowP12PasswordOnRecover gets a reference to the given NullableBool and assigns it to the ShowP12PasswordOnRecover field.
func (o *CertificateProfileCryptoPolicy) SetShowP12PasswordOnRecover(v bool) {
	o.ShowP12PasswordOnRecover.Set(&v)
}
// SetShowP12PasswordOnRecoverNil sets the value for ShowP12PasswordOnRecover to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetShowP12PasswordOnRecoverNil() {
	o.ShowP12PasswordOnRecover.Set(nil)
}

// UnsetShowP12PasswordOnRecover ensures that no value is present for ShowP12PasswordOnRecover, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetShowP12PasswordOnRecover() {
	o.ShowP12PasswordOnRecover.Unset()
}

// GetShowP12OnRecover returns the ShowP12OnRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileCryptoPolicy) GetShowP12OnRecover() bool {
	if o == nil || IsNil(o.ShowP12OnRecover.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowP12OnRecover.Get()
}

// GetShowP12OnRecoverOk returns a tuple with the ShowP12OnRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileCryptoPolicy) GetShowP12OnRecoverOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowP12OnRecover.Get(), o.ShowP12OnRecover.IsSet()
}

// HasShowP12OnRecover returns a boolean if a field has been set.
func (o *CertificateProfileCryptoPolicy) HasShowP12OnRecover() bool {
	if o != nil && o.ShowP12OnRecover.IsSet() {
		return true
	}

	return false
}

// SetShowP12OnRecover gets a reference to the given NullableBool and assigns it to the ShowP12OnRecover field.
func (o *CertificateProfileCryptoPolicy) SetShowP12OnRecover(v bool) {
	o.ShowP12OnRecover.Set(&v)
}
// SetShowP12OnRecoverNil sets the value for ShowP12OnRecover to be an explicit nil
func (o *CertificateProfileCryptoPolicy) SetShowP12OnRecoverNil() {
	o.ShowP12OnRecover.Set(nil)
}

// UnsetShowP12OnRecover ensures that no value is present for ShowP12OnRecover, not even an explicit nil
func (o *CertificateProfileCryptoPolicy) UnsetShowP12OnRecover() {
	o.ShowP12OnRecover.Unset()
}

func (o CertificateProfileCryptoPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateProfileCryptoPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Centralized.IsSet() {
		toSerialize["centralized"] = o.Centralized.Get()
	}
	if o.Decentralized.IsSet() {
		toSerialize["decentralized"] = o.Decentralized.Get()
	}
	if o.DefaultKeyType.IsSet() {
		toSerialize["defaultKeyType"] = o.DefaultKeyType.Get()
	}
	if o.AuthorizedKeyTypes != nil {
		toSerialize["authorizedKeyTypes"] = o.AuthorizedKeyTypes
	}
	if o.PreferredEnrollmentMode.IsSet() {
		toSerialize["preferredEnrollmentMode"] = o.PreferredEnrollmentMode.Get()
	}
	if o.Escrow.IsSet() {
		toSerialize["escrow"] = o.Escrow.Get()
	}
	if o.P12passwordPolicy.IsSet() {
		toSerialize["p12passwordPolicy"] = o.P12passwordPolicy.Get()
	}
	if o.P12passwordMode.IsSet() {
		toSerialize["p12passwordMode"] = o.P12passwordMode.Get()
	}
	if o.P12storeEncryptionType.IsSet() {
		toSerialize["p12storeEncryptionType"] = o.P12storeEncryptionType.Get()
	}
	if o.ShowP12PasswordOnEnroll.IsSet() {
		toSerialize["showP12PasswordOnEnroll"] = o.ShowP12PasswordOnEnroll.Get()
	}
	if o.ShowP12OnEnroll.IsSet() {
		toSerialize["showP12OnEnroll"] = o.ShowP12OnEnroll.Get()
	}
	if o.ShowP12PasswordOnRecover.IsSet() {
		toSerialize["showP12PasswordOnRecover"] = o.ShowP12PasswordOnRecover.Get()
	}
	if o.ShowP12OnRecover.IsSet() {
		toSerialize["showP12OnRecover"] = o.ShowP12OnRecover.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateProfileCryptoPolicy) UnmarshalJSON(data []byte) (err error) {
	varCertificateProfileCryptoPolicy := _CertificateProfileCryptoPolicy{}

	err = json.Unmarshal(data, &varCertificateProfileCryptoPolicy)

	if err != nil {
		return err
	}

	*o = CertificateProfileCryptoPolicy(varCertificateProfileCryptoPolicy)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "centralized")
		delete(additionalProperties, "decentralized")
		delete(additionalProperties, "defaultKeyType")
		delete(additionalProperties, "authorizedKeyTypes")
		delete(additionalProperties, "preferredEnrollmentMode")
		delete(additionalProperties, "escrow")
		delete(additionalProperties, "p12passwordPolicy")
		delete(additionalProperties, "p12passwordMode")
		delete(additionalProperties, "p12storeEncryptionType")
		delete(additionalProperties, "showP12PasswordOnEnroll")
		delete(additionalProperties, "showP12OnEnroll")
		delete(additionalProperties, "showP12PasswordOnRecover")
		delete(additionalProperties, "showP12OnRecover")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateProfileCryptoPolicy struct {
	value *CertificateProfileCryptoPolicy
	isSet bool
}

func (v NullableCertificateProfileCryptoPolicy) Get() *CertificateProfileCryptoPolicy {
	return v.value
}

func (v *NullableCertificateProfileCryptoPolicy) Set(val *CertificateProfileCryptoPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateProfileCryptoPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateProfileCryptoPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateProfileCryptoPolicy(val *CertificateProfileCryptoPolicy) *NullableCertificateProfileCryptoPolicy {
	return &NullableCertificateProfileCryptoPolicy{value: val, isSet: true}
}

func (v NullableCertificateProfileCryptoPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateProfileCryptoPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


