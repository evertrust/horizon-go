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

// checks if the AWSACMPCAConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AWSACMPCAConnectorResponse{}

// AWSACMPCAConnectorResponse struct for AWSACMPCAConnectorResponse
type AWSACMPCAConnectorResponse struct {
	// Object internal ID
	Id     string `json:"_id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Region string `json:"region"`
	CaArn  string `json:"caArn"`
	// Name of the `password` [credentials](#tag/security.credentials) containing Access Key Id and Secret Access Key. If not defined, an account present in environment variables can be used.
	AccessCredentials    utils.NullableString       `json:"accessCredentials,omitempty"`
	TemplateArn          utils.NullableString       `json:"templateArn,omitempty"`
	RoleArn              utils.NullableString       `json:"roleArn,omitempty"`
	ValidDays            utils.NullableString       `json:"validDays,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	RetryInterval        utils.NullableString       `json:"retryInterval,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	SigningHash          utils.NullableString       `json:"signingHash,omitempty"`
	CertificateUsage     utils.NullableString       `json:"certificateUsage,omitempty"`
	CaPolicyOid          utils.NullableString       `json:"caPolicyOid,omitempty"`
	Timeout              utils.NullableString       `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Proxy                utils.NullableString       `json:"proxy,omitempty"`
	Queue                utils.NullableString       `json:"queue,omitempty"`
	Status               NullablePKIConnectorStatus `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AWSACMPCAConnectorResponse AWSACMPCAConnectorResponse

// NewAWSACMPCAConnectorResponse instantiates a new AWSACMPCAConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSACMPCAConnectorResponse(id string, name string, type_ string, region string, caArn string) *AWSACMPCAConnectorResponse {
	this := AWSACMPCAConnectorResponse{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Region = region
	this.CaArn = caArn
	return &this
}

// NewAWSACMPCAConnectorResponseWithDefaults instantiates a new AWSACMPCAConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSACMPCAConnectorResponseWithDefaults() *AWSACMPCAConnectorResponse {
	this := AWSACMPCAConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AWSACMPCAConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AWSACMPCAConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AWSACMPCAConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AWSACMPCAConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AWSACMPCAConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AWSACMPCAConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *AWSACMPCAConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSACMPCAConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AWSACMPCAConnectorResponse) SetType(v string) {
	o.Type = v
}

// GetRegion returns the Region field value
func (o *AWSACMPCAConnectorResponse) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AWSACMPCAConnectorResponse) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AWSACMPCAConnectorResponse) SetRegion(v string) {
	o.Region = v
}

// GetCaArn returns the CaArn field value
func (o *AWSACMPCAConnectorResponse) GetCaArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaArn
}

// GetCaArnOk returns a tuple with the CaArn field value
// and a boolean to check if the value has been set.
func (o *AWSACMPCAConnectorResponse) GetCaArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaArn, true
}

// SetCaArn sets field value
func (o *AWSACMPCAConnectorResponse) SetCaArn(v string) {
	o.CaArn = v
}

// GetAccessCredentials returns the AccessCredentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetAccessCredentials() string {
	if o == nil || utils.IsNil(o.AccessCredentials.Get()) {
		var ret string
		return ret
	}
	return *o.AccessCredentials.Get()
}

// GetAccessCredentialsOk returns a tuple with the AccessCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetAccessCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessCredentials.Get(), o.AccessCredentials.IsSet()
}

// HasAccessCredentials returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasAccessCredentials() bool {
	if o != nil && o.AccessCredentials.IsSet() {
		return true
	}

	return false
}

// SetAccessCredentials gets a reference to the given NullableString and assigns it to the AccessCredentials field.
func (o *AWSACMPCAConnectorResponse) SetAccessCredentials(v string) {
	o.AccessCredentials.Set(&v)
}

// SetAccessCredentialsNil sets the value for AccessCredentials to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetAccessCredentialsNil() {
	o.AccessCredentials.Set(nil)
}

// UnsetAccessCredentials ensures that no value is present for AccessCredentials, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetAccessCredentials() {
	o.AccessCredentials.Unset()
}

// GetTemplateArn returns the TemplateArn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetTemplateArn() string {
	if o == nil || utils.IsNil(o.TemplateArn.Get()) {
		var ret string
		return ret
	}
	return *o.TemplateArn.Get()
}

// GetTemplateArnOk returns a tuple with the TemplateArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetTemplateArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateArn.Get(), o.TemplateArn.IsSet()
}

// HasTemplateArn returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasTemplateArn() bool {
	if o != nil && o.TemplateArn.IsSet() {
		return true
	}

	return false
}

// SetTemplateArn gets a reference to the given NullableString and assigns it to the TemplateArn field.
func (o *AWSACMPCAConnectorResponse) SetTemplateArn(v string) {
	o.TemplateArn.Set(&v)
}

// SetTemplateArnNil sets the value for TemplateArn to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetTemplateArnNil() {
	o.TemplateArn.Set(nil)
}

// UnsetTemplateArn ensures that no value is present for TemplateArn, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetTemplateArn() {
	o.TemplateArn.Unset()
}

// GetRoleArn returns the RoleArn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetRoleArn() string {
	if o == nil || utils.IsNil(o.RoleArn.Get()) {
		var ret string
		return ret
	}
	return *o.RoleArn.Get()
}

// GetRoleArnOk returns a tuple with the RoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetRoleArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleArn.Get(), o.RoleArn.IsSet()
}

// HasRoleArn returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasRoleArn() bool {
	if o != nil && o.RoleArn.IsSet() {
		return true
	}

	return false
}

// SetRoleArn gets a reference to the given NullableString and assigns it to the RoleArn field.
func (o *AWSACMPCAConnectorResponse) SetRoleArn(v string) {
	o.RoleArn.Set(&v)
}

// SetRoleArnNil sets the value for RoleArn to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetRoleArnNil() {
	o.RoleArn.Set(nil)
}

// UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetRoleArn() {
	o.RoleArn.Unset()
}

// GetValidDays returns the ValidDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetValidDays() string {
	if o == nil || utils.IsNil(o.ValidDays.Get()) {
		var ret string
		return ret
	}
	return *o.ValidDays.Get()
}

// GetValidDaysOk returns a tuple with the ValidDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetValidDaysOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidDays.Get(), o.ValidDays.IsSet()
}

// HasValidDays returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasValidDays() bool {
	if o != nil && o.ValidDays.IsSet() {
		return true
	}

	return false
}

// SetValidDays gets a reference to the given NullableString and assigns it to the ValidDays field.
func (o *AWSACMPCAConnectorResponse) SetValidDays(v string) {
	o.ValidDays.Set(&v)
}

// SetValidDaysNil sets the value for ValidDays to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetValidDaysNil() {
	o.ValidDays.Set(nil)
}

// UnsetValidDays ensures that no value is present for ValidDays, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetValidDays() {
	o.ValidDays.Unset()
}

// GetRetryInterval returns the RetryInterval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetRetryInterval() string {
	if o == nil || utils.IsNil(o.RetryInterval.Get()) {
		var ret string
		return ret
	}
	return *o.RetryInterval.Get()
}

// GetRetryIntervalOk returns a tuple with the RetryInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetRetryIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetryInterval.Get(), o.RetryInterval.IsSet()
}

// HasRetryInterval returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasRetryInterval() bool {
	if o != nil && o.RetryInterval.IsSet() {
		return true
	}

	return false
}

// SetRetryInterval gets a reference to the given NullableString and assigns it to the RetryInterval field.
func (o *AWSACMPCAConnectorResponse) SetRetryInterval(v string) {
	o.RetryInterval.Set(&v)
}

// SetRetryIntervalNil sets the value for RetryInterval to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetRetryIntervalNil() {
	o.RetryInterval.Set(nil)
}

// UnsetRetryInterval ensures that no value is present for RetryInterval, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetRetryInterval() {
	o.RetryInterval.Unset()
}

// GetSigningHash returns the SigningHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetSigningHash() string {
	if o == nil || utils.IsNil(o.SigningHash.Get()) {
		var ret string
		return ret
	}
	return *o.SigningHash.Get()
}

// GetSigningHashOk returns a tuple with the SigningHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetSigningHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SigningHash.Get(), o.SigningHash.IsSet()
}

// HasSigningHash returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasSigningHash() bool {
	if o != nil && o.SigningHash.IsSet() {
		return true
	}

	return false
}

// SetSigningHash gets a reference to the given NullableString and assigns it to the SigningHash field.
func (o *AWSACMPCAConnectorResponse) SetSigningHash(v string) {
	o.SigningHash.Set(&v)
}

// SetSigningHashNil sets the value for SigningHash to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetSigningHashNil() {
	o.SigningHash.Set(nil)
}

// UnsetSigningHash ensures that no value is present for SigningHash, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetSigningHash() {
	o.SigningHash.Unset()
}

// GetCertificateUsage returns the CertificateUsage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetCertificateUsage() string {
	if o == nil || utils.IsNil(o.CertificateUsage.Get()) {
		var ret string
		return ret
	}
	return *o.CertificateUsage.Get()
}

// GetCertificateUsageOk returns a tuple with the CertificateUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetCertificateUsageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateUsage.Get(), o.CertificateUsage.IsSet()
}

// HasCertificateUsage returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasCertificateUsage() bool {
	if o != nil && o.CertificateUsage.IsSet() {
		return true
	}

	return false
}

// SetCertificateUsage gets a reference to the given NullableString and assigns it to the CertificateUsage field.
func (o *AWSACMPCAConnectorResponse) SetCertificateUsage(v string) {
	o.CertificateUsage.Set(&v)
}

// SetCertificateUsageNil sets the value for CertificateUsage to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetCertificateUsageNil() {
	o.CertificateUsage.Set(nil)
}

// UnsetCertificateUsage ensures that no value is present for CertificateUsage, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetCertificateUsage() {
	o.CertificateUsage.Unset()
}

// GetCaPolicyOid returns the CaPolicyOid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetCaPolicyOid() string {
	if o == nil || utils.IsNil(o.CaPolicyOid.Get()) {
		var ret string
		return ret
	}
	return *o.CaPolicyOid.Get()
}

// GetCaPolicyOidOk returns a tuple with the CaPolicyOid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetCaPolicyOidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaPolicyOid.Get(), o.CaPolicyOid.IsSet()
}

// HasCaPolicyOid returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasCaPolicyOid() bool {
	if o != nil && o.CaPolicyOid.IsSet() {
		return true
	}

	return false
}

// SetCaPolicyOid gets a reference to the given NullableString and assigns it to the CaPolicyOid field.
func (o *AWSACMPCAConnectorResponse) SetCaPolicyOid(v string) {
	o.CaPolicyOid.Set(&v)
}

// SetCaPolicyOidNil sets the value for CaPolicyOid to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetCaPolicyOidNil() {
	o.CaPolicyOid.Set(nil)
}

// UnsetCaPolicyOid ensures that no value is present for CaPolicyOid, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetCaPolicyOid() {
	o.CaPolicyOid.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *AWSACMPCAConnectorResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *AWSACMPCAConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *AWSACMPCAConnectorResponse) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetQueue() {
	o.Queue.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSACMPCAConnectorResponse) GetStatus() PKIConnectorStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret PKIConnectorStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSACMPCAConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *AWSACMPCAConnectorResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePKIConnectorStatus and assigns it to the Status field.
func (o *AWSACMPCAConnectorResponse) SetStatus(v PKIConnectorStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *AWSACMPCAConnectorResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *AWSACMPCAConnectorResponse) UnsetStatus() {
	o.Status.Unset()
}

func (o AWSACMPCAConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AWSACMPCAConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["region"] = o.Region
	toSerialize["caArn"] = o.CaArn
	if o.AccessCredentials.IsSet() {
		toSerialize["accessCredentials"] = o.AccessCredentials.Get()
	}
	if o.TemplateArn.IsSet() {
		toSerialize["templateArn"] = o.TemplateArn.Get()
	}
	if o.RoleArn.IsSet() {
		toSerialize["roleArn"] = o.RoleArn.Get()
	}
	if o.ValidDays.IsSet() {
		toSerialize["validDays"] = o.ValidDays.Get()
	}
	if o.RetryInterval.IsSet() {
		toSerialize["retryInterval"] = o.RetryInterval.Get()
	}
	if o.SigningHash.IsSet() {
		toSerialize["signingHash"] = o.SigningHash.Get()
	}
	if o.CertificateUsage.IsSet() {
		toSerialize["certificateUsage"] = o.CertificateUsage.Get()
	}
	if o.CaPolicyOid.IsSet() {
		toSerialize["caPolicyOid"] = o.CaPolicyOid.Get()
	}
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AWSACMPCAConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"name",
		"type",
		"region",
		"caArn",
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

	varAWSACMPCAConnectorResponse := _AWSACMPCAConnectorResponse{}

	err = json.Unmarshal(data, &varAWSACMPCAConnectorResponse)

	if err != nil {
		return err
	}

	*o = AWSACMPCAConnectorResponse(varAWSACMPCAConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "region")
		delete(additionalProperties, "caArn")
		delete(additionalProperties, "accessCredentials")
		delete(additionalProperties, "templateArn")
		delete(additionalProperties, "roleArn")
		delete(additionalProperties, "validDays")
		delete(additionalProperties, "retryInterval")
		delete(additionalProperties, "signingHash")
		delete(additionalProperties, "certificateUsage")
		delete(additionalProperties, "caPolicyOid")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAWSACMPCAConnectorResponse struct {
	value *AWSACMPCAConnectorResponse
	isSet bool
}

func (v NullableAWSACMPCAConnectorResponse) Get() *AWSACMPCAConnectorResponse {
	return v.value
}

func (v *NullableAWSACMPCAConnectorResponse) Set(val *AWSACMPCAConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSACMPCAConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSACMPCAConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSACMPCAConnectorResponse(val *AWSACMPCAConnectorResponse) *NullableAWSACMPCAConnectorResponse {
	return &NullableAWSACMPCAConnectorResponse{value: val, isSet: true}
}

func (v NullableAWSACMPCAConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSACMPCAConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
