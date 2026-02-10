/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the CertificateAuthorityRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CertificateAuthorityRequest{}

// CertificateAuthorityRequest struct for CertificateAuthorityRequest
type CertificateAuthorityRequest struct {
	Certificate                    string               `json:"certificate"`
	Name                           string               `json:"name"`
	SubjectKeyIdentifier           utils.NullableString `json:"subjectKeyIdentifier,omitempty"`
	ResponderUrl                   utils.NullableString `json:"responderUrl,omitempty"`
	CrlUrl                         utils.NullableString `json:"crlUrl,omitempty"`
	Refresh                        utils.NullableString `json:"refresh,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	TrustedForClientAuthentication bool                 `json:"trustedForClientAuthentication"`
	TrustedForServerAuthentication bool                 `json:"trustedForServerAuthentication"`
	OutdatedRevocationStatusPolicy string               `json:"outdatedRevocationStatusPolicy"`
	Timeout                        utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Proxy                          utils.NullableString `json:"proxy,omitempty"`
	CacheTimeToIdle                utils.NullableString `json:"cacheTimeToIdle,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Public                         bool                 `json:"public"`
	Downloadable                   *bool                `json:"downloadable,omitempty"`
	AdditionalProperties           map[string]interface{}
}

type _CertificateAuthorityRequest CertificateAuthorityRequest

// NewCertificateAuthorityRequest instantiates a new CertificateAuthorityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateAuthorityRequest(certificate string, name string, trustedForClientAuthentication bool, trustedForServerAuthentication bool, outdatedRevocationStatusPolicy string, public bool) *CertificateAuthorityRequest {
	this := CertificateAuthorityRequest{}
	this.Name = name
	this.TrustedForClientAuthentication = trustedForClientAuthentication
	this.TrustedForServerAuthentication = trustedForServerAuthentication
	this.OutdatedRevocationStatusPolicy = outdatedRevocationStatusPolicy
	this.Public = public
	return &this
}

// NewCertificateAuthorityRequestWithDefaults instantiates a new CertificateAuthorityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateAuthorityRequestWithDefaults() *CertificateAuthorityRequest {
	this := CertificateAuthorityRequest{}
	return &this
}

// GetCertificate returns the Certificate field value
func (o *CertificateAuthorityRequest) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityRequest) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *CertificateAuthorityRequest) SetCertificate(v string) {
	o.Certificate = v
}

// GetName returns the Name field value
func (o *CertificateAuthorityRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CertificateAuthorityRequest) SetName(v string) {
	o.Name = v
}

// GetSubjectKeyIdentifier returns the SubjectKeyIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAuthorityRequest) GetSubjectKeyIdentifier() string {
	if o == nil || utils.IsNil(o.SubjectKeyIdentifier.Get()) {
		var ret string
		return ret
	}
	return *o.SubjectKeyIdentifier.Get()
}

// GetSubjectKeyIdentifierOk returns a tuple with the SubjectKeyIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAuthorityRequest) GetSubjectKeyIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectKeyIdentifier.Get(), o.SubjectKeyIdentifier.IsSet()
}

// HasSubjectKeyIdentifier returns a boolean if a field has been set.
func (o *CertificateAuthorityRequest) HasSubjectKeyIdentifier() bool {
	if o != nil && o.SubjectKeyIdentifier.IsSet() {
		return true
	}

	return false
}

// SetSubjectKeyIdentifier gets a reference to the given NullableString and assigns it to the SubjectKeyIdentifier field.
func (o *CertificateAuthorityRequest) SetSubjectKeyIdentifier(v string) {
	o.SubjectKeyIdentifier.Set(&v)
}

// SetSubjectKeyIdentifierNil sets the value for SubjectKeyIdentifier to be an explicit nil
func (o *CertificateAuthorityRequest) SetSubjectKeyIdentifierNil() {
	o.SubjectKeyIdentifier.Set(nil)
}

// UnsetSubjectKeyIdentifier ensures that no value is present for SubjectKeyIdentifier, not even an explicit nil
func (o *CertificateAuthorityRequest) UnsetSubjectKeyIdentifier() {
	o.SubjectKeyIdentifier.Unset()
}

// GetResponderUrl returns the ResponderUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAuthorityRequest) GetResponderUrl() string {
	if o == nil || utils.IsNil(o.ResponderUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ResponderUrl.Get()
}

// GetResponderUrlOk returns a tuple with the ResponderUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAuthorityRequest) GetResponderUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResponderUrl.Get(), o.ResponderUrl.IsSet()
}

// HasResponderUrl returns a boolean if a field has been set.
func (o *CertificateAuthorityRequest) HasResponderUrl() bool {
	if o != nil && o.ResponderUrl.IsSet() {
		return true
	}

	return false
}

// SetResponderUrl gets a reference to the given NullableString and assigns it to the ResponderUrl field.
func (o *CertificateAuthorityRequest) SetResponderUrl(v string) {
	o.ResponderUrl.Set(&v)
}

// SetResponderUrlNil sets the value for ResponderUrl to be an explicit nil
func (o *CertificateAuthorityRequest) SetResponderUrlNil() {
	o.ResponderUrl.Set(nil)
}

// UnsetResponderUrl ensures that no value is present for ResponderUrl, not even an explicit nil
func (o *CertificateAuthorityRequest) UnsetResponderUrl() {
	o.ResponderUrl.Unset()
}

// GetCrlUrl returns the CrlUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAuthorityRequest) GetCrlUrl() string {
	if o == nil || utils.IsNil(o.CrlUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CrlUrl.Get()
}

// GetCrlUrlOk returns a tuple with the CrlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAuthorityRequest) GetCrlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CrlUrl.Get(), o.CrlUrl.IsSet()
}

// HasCrlUrl returns a boolean if a field has been set.
func (o *CertificateAuthorityRequest) HasCrlUrl() bool {
	if o != nil && o.CrlUrl.IsSet() {
		return true
	}

	return false
}

// SetCrlUrl gets a reference to the given NullableString and assigns it to the CrlUrl field.
func (o *CertificateAuthorityRequest) SetCrlUrl(v string) {
	o.CrlUrl.Set(&v)
}

// SetCrlUrlNil sets the value for CrlUrl to be an explicit nil
func (o *CertificateAuthorityRequest) SetCrlUrlNil() {
	o.CrlUrl.Set(nil)
}

// UnsetCrlUrl ensures that no value is present for CrlUrl, not even an explicit nil
func (o *CertificateAuthorityRequest) UnsetCrlUrl() {
	o.CrlUrl.Unset()
}

// GetRefresh returns the Refresh field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAuthorityRequest) GetRefresh() string {
	if o == nil || utils.IsNil(o.Refresh.Get()) {
		var ret string
		return ret
	}
	return *o.Refresh.Get()
}

// GetRefreshOk returns a tuple with the Refresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAuthorityRequest) GetRefreshOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Refresh.Get(), o.Refresh.IsSet()
}

// HasRefresh returns a boolean if a field has been set.
func (o *CertificateAuthorityRequest) HasRefresh() bool {
	if o != nil && o.Refresh.IsSet() {
		return true
	}

	return false
}

// SetRefresh gets a reference to the given NullableString and assigns it to the Refresh field.
func (o *CertificateAuthorityRequest) SetRefresh(v string) {
	o.Refresh.Set(&v)
}

// SetRefreshNil sets the value for Refresh to be an explicit nil
func (o *CertificateAuthorityRequest) SetRefreshNil() {
	o.Refresh.Set(nil)
}

// UnsetRefresh ensures that no value is present for Refresh, not even an explicit nil
func (o *CertificateAuthorityRequest) UnsetRefresh() {
	o.Refresh.Unset()
}

// GetTrustedForClientAuthentication returns the TrustedForClientAuthentication field value
func (o *CertificateAuthorityRequest) GetTrustedForClientAuthentication() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TrustedForClientAuthentication
}

// GetTrustedForClientAuthenticationOk returns a tuple with the TrustedForClientAuthentication field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityRequest) GetTrustedForClientAuthenticationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustedForClientAuthentication, true
}

// SetTrustedForClientAuthentication sets field value
func (o *CertificateAuthorityRequest) SetTrustedForClientAuthentication(v bool) {
	o.TrustedForClientAuthentication = v
}

// GetTrustedForServerAuthentication returns the TrustedForServerAuthentication field value
func (o *CertificateAuthorityRequest) GetTrustedForServerAuthentication() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TrustedForServerAuthentication
}

// GetTrustedForServerAuthenticationOk returns a tuple with the TrustedForServerAuthentication field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityRequest) GetTrustedForServerAuthenticationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustedForServerAuthentication, true
}

// SetTrustedForServerAuthentication sets field value
func (o *CertificateAuthorityRequest) SetTrustedForServerAuthentication(v bool) {
	o.TrustedForServerAuthentication = v
}

// GetOutdatedRevocationStatusPolicy returns the OutdatedRevocationStatusPolicy field value
func (o *CertificateAuthorityRequest) GetOutdatedRevocationStatusPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutdatedRevocationStatusPolicy
}

// GetOutdatedRevocationStatusPolicyOk returns a tuple with the OutdatedRevocationStatusPolicy field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityRequest) GetOutdatedRevocationStatusPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutdatedRevocationStatusPolicy, true
}

// SetOutdatedRevocationStatusPolicy sets field value
func (o *CertificateAuthorityRequest) SetOutdatedRevocationStatusPolicy(v string) {
	o.OutdatedRevocationStatusPolicy = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAuthorityRequest) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAuthorityRequest) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *CertificateAuthorityRequest) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *CertificateAuthorityRequest) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *CertificateAuthorityRequest) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *CertificateAuthorityRequest) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAuthorityRequest) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAuthorityRequest) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *CertificateAuthorityRequest) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *CertificateAuthorityRequest) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *CertificateAuthorityRequest) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *CertificateAuthorityRequest) UnsetProxy() {
	o.Proxy.Unset()
}

// GetCacheTimeToIdle returns the CacheTimeToIdle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateAuthorityRequest) GetCacheTimeToIdle() string {
	if o == nil || utils.IsNil(o.CacheTimeToIdle.Get()) {
		var ret string
		return ret
	}
	return *o.CacheTimeToIdle.Get()
}

// GetCacheTimeToIdleOk returns a tuple with the CacheTimeToIdle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateAuthorityRequest) GetCacheTimeToIdleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CacheTimeToIdle.Get(), o.CacheTimeToIdle.IsSet()
}

// HasCacheTimeToIdle returns a boolean if a field has been set.
func (o *CertificateAuthorityRequest) HasCacheTimeToIdle() bool {
	if o != nil && o.CacheTimeToIdle.IsSet() {
		return true
	}

	return false
}

// SetCacheTimeToIdle gets a reference to the given NullableString and assigns it to the CacheTimeToIdle field.
func (o *CertificateAuthorityRequest) SetCacheTimeToIdle(v string) {
	o.CacheTimeToIdle.Set(&v)
}

// SetCacheTimeToIdleNil sets the value for CacheTimeToIdle to be an explicit nil
func (o *CertificateAuthorityRequest) SetCacheTimeToIdleNil() {
	o.CacheTimeToIdle.Set(nil)
}

// UnsetCacheTimeToIdle ensures that no value is present for CacheTimeToIdle, not even an explicit nil
func (o *CertificateAuthorityRequest) UnsetCacheTimeToIdle() {
	o.CacheTimeToIdle.Unset()
}

// GetPublic returns the Public field value
func (o *CertificateAuthorityRequest) GetPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityRequest) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *CertificateAuthorityRequest) SetPublic(v bool) {
	o.Public = v
}

// GetDownloadable returns the Downloadable field value if set, zero value otherwise.
func (o *CertificateAuthorityRequest) GetDownloadable() bool {
	if o == nil || utils.IsNil(o.Downloadable) {
		var ret bool
		return ret
	}
	return *o.Downloadable
}

// GetDownloadableOk returns a tuple with the Downloadable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateAuthorityRequest) GetDownloadableOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Downloadable) {
		return nil, false
	}
	return o.Downloadable, true
}

// HasDownloadable returns a boolean if a field has been set.
func (o *CertificateAuthorityRequest) HasDownloadable() bool {
	if o != nil && !utils.IsNil(o.Downloadable) {
		return true
	}

	return false
}

// SetDownloadable gets a reference to the given bool and assigns it to the Downloadable field.
func (o *CertificateAuthorityRequest) SetDownloadable(v bool) {
	o.Downloadable = &v
}

func (o CertificateAuthorityRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateAuthorityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["certificate"] = o.Certificate
	toSerialize["name"] = o.Name
	if o.SubjectKeyIdentifier.IsSet() {
		toSerialize["subjectKeyIdentifier"] = o.SubjectKeyIdentifier.Get()
	}
	if o.ResponderUrl.IsSet() {
		toSerialize["responderUrl"] = o.ResponderUrl.Get()
	}
	if o.CrlUrl.IsSet() {
		toSerialize["crlUrl"] = o.CrlUrl.Get()
	}
	if o.Refresh.IsSet() {
		toSerialize["refresh"] = o.Refresh.Get()
	}
	toSerialize["trustedForClientAuthentication"] = o.TrustedForClientAuthentication
	toSerialize["trustedForServerAuthentication"] = o.TrustedForServerAuthentication
	toSerialize["outdatedRevocationStatusPolicy"] = o.OutdatedRevocationStatusPolicy
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.CacheTimeToIdle.IsSet() {
		toSerialize["cacheTimeToIdle"] = o.CacheTimeToIdle.Get()
	}
	toSerialize["public"] = o.Public
	if !utils.IsNil(o.Downloadable) {
		toSerialize["downloadable"] = o.Downloadable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateAuthorityRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"certificate",
		"name",
		"trustedForClientAuthentication",
		"trustedForServerAuthentication",
		"outdatedRevocationStatusPolicy",
		"public",
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

	varCertificateAuthorityRequest := _CertificateAuthorityRequest{}

	err = json.Unmarshal(data, &varCertificateAuthorityRequest)

	if err != nil {
		return err
	}

	*o = CertificateAuthorityRequest(varCertificateAuthorityRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "name")
		delete(additionalProperties, "subjectKeyIdentifier")
		delete(additionalProperties, "responderUrl")
		delete(additionalProperties, "crlUrl")
		delete(additionalProperties, "refresh")
		delete(additionalProperties, "trustedForClientAuthentication")
		delete(additionalProperties, "trustedForServerAuthentication")
		delete(additionalProperties, "outdatedRevocationStatusPolicy")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "cacheTimeToIdle")
		delete(additionalProperties, "public")
		delete(additionalProperties, "downloadable")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateAuthorityRequest struct {
	value *CertificateAuthorityRequest
	isSet bool
}

func (v NullableCertificateAuthorityRequest) Get() *CertificateAuthorityRequest {
	return v.value
}

func (v *NullableCertificateAuthorityRequest) Set(val *CertificateAuthorityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateAuthorityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateAuthorityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateAuthorityRequest(val *CertificateAuthorityRequest) *NullableCertificateAuthorityRequest {
	return &NullableCertificateAuthorityRequest{value: val, isSet: true}
}

func (v NullableCertificateAuthorityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateAuthorityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
