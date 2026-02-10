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

// checks if the CMPConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CMPConnector{}

// CMPConnector struct for CMPConnector
type CMPConnector struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	EndPoint     string `json:"endPoint"`
	Profile      string `json:"profile"`
	IssuerCADN   string `json:"issuerCADN"`
	IssuerCACert string `json:"issuerCACert"`
	// Name of the `certificate` [credentials](#tag/security.credentials) to use to sign on the PKI
	SignerCredentials string               `json:"signerCredentials"`
	EmailMap          utils.NullableString `json:"emailMap,omitempty"`
	SanDnsMap         utils.NullableString `json:"sanDnsMap,omitempty"`
	CnMap             utils.NullableString `json:"cnMap,omitempty"`
	ProfileMap        utils.NullableString `json:"profileMap,omitempty"`
	IssuerMap         utils.NullableString `json:"issuerMap,omitempty"`
	LegacyCMPStyle    utils.NullableBool   `json:"legacyCMPStyle,omitempty"`
	// Name of the `certificate` [credentials](#tag/security.credentials) to use to authenticate on the PKI
	AuthenticationCredentials string               `json:"authenticationCredentials"`
	Timeout                   utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Proxy                     utils.NullableString `json:"proxy,omitempty"`
	Queue                     utils.NullableString `json:"queue,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _CMPConnector CMPConnector

// NewCMPConnector instantiates a new CMPConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCMPConnector(name string, type_ string, endPoint string, profile string, issuerCADN string, issuerCACert string, signerCredentials string, authenticationCredentials string) *CMPConnector {
	this := CMPConnector{}
	this.Name = name
	this.Type = type_
	this.EndPoint = endPoint
	this.Profile = profile
	this.IssuerCADN = issuerCADN
	this.IssuerCACert = issuerCACert
	this.SignerCredentials = signerCredentials
	this.AuthenticationCredentials = authenticationCredentials
	return &this
}

// NewCMPConnectorWithDefaults instantiates a new CMPConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCMPConnectorWithDefaults() *CMPConnector {
	this := CMPConnector{}
	return &this
}

// GetName returns the Name field value
func (o *CMPConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CMPConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CMPConnector) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CMPConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CMPConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CMPConnector) SetType(v string) {
	o.Type = v
}

// GetEndPoint returns the EndPoint field value
func (o *CMPConnector) GetEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value
// and a boolean to check if the value has been set.
func (o *CMPConnector) GetEndPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPoint, true
}

// SetEndPoint sets field value
func (o *CMPConnector) SetEndPoint(v string) {
	o.EndPoint = v
}

// GetProfile returns the Profile field value
func (o *CMPConnector) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *CMPConnector) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *CMPConnector) SetProfile(v string) {
	o.Profile = v
}

// GetIssuerCADN returns the IssuerCADN field value
func (o *CMPConnector) GetIssuerCADN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerCADN
}

// GetIssuerCADNOk returns a tuple with the IssuerCADN field value
// and a boolean to check if the value has been set.
func (o *CMPConnector) GetIssuerCADNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerCADN, true
}

// SetIssuerCADN sets field value
func (o *CMPConnector) SetIssuerCADN(v string) {
	o.IssuerCADN = v
}

// GetIssuerCACert returns the IssuerCACert field value
func (o *CMPConnector) GetIssuerCACert() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerCACert
}

// GetIssuerCACertOk returns a tuple with the IssuerCACert field value
// and a boolean to check if the value has been set.
func (o *CMPConnector) GetIssuerCACertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerCACert, true
}

// SetIssuerCACert sets field value
func (o *CMPConnector) SetIssuerCACert(v string) {
	o.IssuerCACert = v
}

// GetSignerCredentials returns the SignerCredentials field value
func (o *CMPConnector) GetSignerCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignerCredentials
}

// GetSignerCredentialsOk returns a tuple with the SignerCredentials field value
// and a boolean to check if the value has been set.
func (o *CMPConnector) GetSignerCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignerCredentials, true
}

// SetSignerCredentials sets field value
func (o *CMPConnector) SetSignerCredentials(v string) {
	o.SignerCredentials = v
}

// GetEmailMap returns the EmailMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CMPConnector) GetEmailMap() string {
	if o == nil || utils.IsNil(o.EmailMap.Get()) {
		var ret string
		return ret
	}
	return *o.EmailMap.Get()
}

// GetEmailMapOk returns a tuple with the EmailMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CMPConnector) GetEmailMapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailMap.Get(), o.EmailMap.IsSet()
}

// HasEmailMap returns a boolean if a field has been set.
func (o *CMPConnector) HasEmailMap() bool {
	if o != nil && o.EmailMap.IsSet() {
		return true
	}

	return false
}

// SetEmailMap gets a reference to the given NullableString and assigns it to the EmailMap field.
func (o *CMPConnector) SetEmailMap(v string) {
	o.EmailMap.Set(&v)
}

// SetEmailMapNil sets the value for EmailMap to be an explicit nil
func (o *CMPConnector) SetEmailMapNil() {
	o.EmailMap.Set(nil)
}

// UnsetEmailMap ensures that no value is present for EmailMap, not even an explicit nil
func (o *CMPConnector) UnsetEmailMap() {
	o.EmailMap.Unset()
}

// GetSanDnsMap returns the SanDnsMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CMPConnector) GetSanDnsMap() string {
	if o == nil || utils.IsNil(o.SanDnsMap.Get()) {
		var ret string
		return ret
	}
	return *o.SanDnsMap.Get()
}

// GetSanDnsMapOk returns a tuple with the SanDnsMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CMPConnector) GetSanDnsMapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SanDnsMap.Get(), o.SanDnsMap.IsSet()
}

// HasSanDnsMap returns a boolean if a field has been set.
func (o *CMPConnector) HasSanDnsMap() bool {
	if o != nil && o.SanDnsMap.IsSet() {
		return true
	}

	return false
}

// SetSanDnsMap gets a reference to the given NullableString and assigns it to the SanDnsMap field.
func (o *CMPConnector) SetSanDnsMap(v string) {
	o.SanDnsMap.Set(&v)
}

// SetSanDnsMapNil sets the value for SanDnsMap to be an explicit nil
func (o *CMPConnector) SetSanDnsMapNil() {
	o.SanDnsMap.Set(nil)
}

// UnsetSanDnsMap ensures that no value is present for SanDnsMap, not even an explicit nil
func (o *CMPConnector) UnsetSanDnsMap() {
	o.SanDnsMap.Unset()
}

// GetCnMap returns the CnMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CMPConnector) GetCnMap() string {
	if o == nil || utils.IsNil(o.CnMap.Get()) {
		var ret string
		return ret
	}
	return *o.CnMap.Get()
}

// GetCnMapOk returns a tuple with the CnMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CMPConnector) GetCnMapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CnMap.Get(), o.CnMap.IsSet()
}

// HasCnMap returns a boolean if a field has been set.
func (o *CMPConnector) HasCnMap() bool {
	if o != nil && o.CnMap.IsSet() {
		return true
	}

	return false
}

// SetCnMap gets a reference to the given NullableString and assigns it to the CnMap field.
func (o *CMPConnector) SetCnMap(v string) {
	o.CnMap.Set(&v)
}

// SetCnMapNil sets the value for CnMap to be an explicit nil
func (o *CMPConnector) SetCnMapNil() {
	o.CnMap.Set(nil)
}

// UnsetCnMap ensures that no value is present for CnMap, not even an explicit nil
func (o *CMPConnector) UnsetCnMap() {
	o.CnMap.Unset()
}

// GetProfileMap returns the ProfileMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CMPConnector) GetProfileMap() string {
	if o == nil || utils.IsNil(o.ProfileMap.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileMap.Get()
}

// GetProfileMapOk returns a tuple with the ProfileMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CMPConnector) GetProfileMapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileMap.Get(), o.ProfileMap.IsSet()
}

// HasProfileMap returns a boolean if a field has been set.
func (o *CMPConnector) HasProfileMap() bool {
	if o != nil && o.ProfileMap.IsSet() {
		return true
	}

	return false
}

// SetProfileMap gets a reference to the given NullableString and assigns it to the ProfileMap field.
func (o *CMPConnector) SetProfileMap(v string) {
	o.ProfileMap.Set(&v)
}

// SetProfileMapNil sets the value for ProfileMap to be an explicit nil
func (o *CMPConnector) SetProfileMapNil() {
	o.ProfileMap.Set(nil)
}

// UnsetProfileMap ensures that no value is present for ProfileMap, not even an explicit nil
func (o *CMPConnector) UnsetProfileMap() {
	o.ProfileMap.Unset()
}

// GetIssuerMap returns the IssuerMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CMPConnector) GetIssuerMap() string {
	if o == nil || utils.IsNil(o.IssuerMap.Get()) {
		var ret string
		return ret
	}
	return *o.IssuerMap.Get()
}

// GetIssuerMapOk returns a tuple with the IssuerMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CMPConnector) GetIssuerMapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuerMap.Get(), o.IssuerMap.IsSet()
}

// HasIssuerMap returns a boolean if a field has been set.
func (o *CMPConnector) HasIssuerMap() bool {
	if o != nil && o.IssuerMap.IsSet() {
		return true
	}

	return false
}

// SetIssuerMap gets a reference to the given NullableString and assigns it to the IssuerMap field.
func (o *CMPConnector) SetIssuerMap(v string) {
	o.IssuerMap.Set(&v)
}

// SetIssuerMapNil sets the value for IssuerMap to be an explicit nil
func (o *CMPConnector) SetIssuerMapNil() {
	o.IssuerMap.Set(nil)
}

// UnsetIssuerMap ensures that no value is present for IssuerMap, not even an explicit nil
func (o *CMPConnector) UnsetIssuerMap() {
	o.IssuerMap.Unset()
}

// GetLegacyCMPStyle returns the LegacyCMPStyle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CMPConnector) GetLegacyCMPStyle() bool {
	if o == nil || utils.IsNil(o.LegacyCMPStyle.Get()) {
		var ret bool
		return ret
	}
	return *o.LegacyCMPStyle.Get()
}

// GetLegacyCMPStyleOk returns a tuple with the LegacyCMPStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CMPConnector) GetLegacyCMPStyleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LegacyCMPStyle.Get(), o.LegacyCMPStyle.IsSet()
}

// HasLegacyCMPStyle returns a boolean if a field has been set.
func (o *CMPConnector) HasLegacyCMPStyle() bool {
	if o != nil && o.LegacyCMPStyle.IsSet() {
		return true
	}

	return false
}

// SetLegacyCMPStyle gets a reference to the given NullableBool and assigns it to the LegacyCMPStyle field.
func (o *CMPConnector) SetLegacyCMPStyle(v bool) {
	o.LegacyCMPStyle.Set(&v)
}

// SetLegacyCMPStyleNil sets the value for LegacyCMPStyle to be an explicit nil
func (o *CMPConnector) SetLegacyCMPStyleNil() {
	o.LegacyCMPStyle.Set(nil)
}

// UnsetLegacyCMPStyle ensures that no value is present for LegacyCMPStyle, not even an explicit nil
func (o *CMPConnector) UnsetLegacyCMPStyle() {
	o.LegacyCMPStyle.Unset()
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value
func (o *CMPConnector) GetAuthenticationCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationCredentials
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value
// and a boolean to check if the value has been set.
func (o *CMPConnector) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationCredentials, true
}

// SetAuthenticationCredentials sets field value
func (o *CMPConnector) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CMPConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CMPConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *CMPConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *CMPConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *CMPConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *CMPConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CMPConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CMPConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *CMPConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *CMPConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *CMPConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *CMPConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CMPConnector) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CMPConnector) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *CMPConnector) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *CMPConnector) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *CMPConnector) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *CMPConnector) UnsetQueue() {
	o.Queue.Unset()
}

func (o CMPConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CMPConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["endPoint"] = o.EndPoint
	toSerialize["profile"] = o.Profile
	toSerialize["issuerCADN"] = o.IssuerCADN
	toSerialize["issuerCACert"] = o.IssuerCACert
	toSerialize["signerCredentials"] = o.SignerCredentials
	if o.EmailMap.IsSet() {
		toSerialize["emailMap"] = o.EmailMap.Get()
	}
	if o.SanDnsMap.IsSet() {
		toSerialize["sanDnsMap"] = o.SanDnsMap.Get()
	}
	if o.CnMap.IsSet() {
		toSerialize["cnMap"] = o.CnMap.Get()
	}
	if o.ProfileMap.IsSet() {
		toSerialize["profileMap"] = o.ProfileMap.Get()
	}
	if o.IssuerMap.IsSet() {
		toSerialize["issuerMap"] = o.IssuerMap.Get()
	}
	if o.LegacyCMPStyle.IsSet() {
		toSerialize["legacyCMPStyle"] = o.LegacyCMPStyle.Get()
	}
	toSerialize["authenticationCredentials"] = o.AuthenticationCredentials
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CMPConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"type",
		"endPoint",
		"profile",
		"issuerCADN",
		"issuerCACert",
		"signerCredentials",
		"authenticationCredentials",
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

	varCMPConnector := _CMPConnector{}

	err = json.Unmarshal(data, &varCMPConnector)

	if err != nil {
		return err
	}

	*o = CMPConnector(varCMPConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "endPoint")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "issuerCADN")
		delete(additionalProperties, "issuerCACert")
		delete(additionalProperties, "signerCredentials")
		delete(additionalProperties, "emailMap")
		delete(additionalProperties, "sanDnsMap")
		delete(additionalProperties, "cnMap")
		delete(additionalProperties, "profileMap")
		delete(additionalProperties, "issuerMap")
		delete(additionalProperties, "legacyCMPStyle")
		delete(additionalProperties, "authenticationCredentials")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCMPConnector struct {
	value *CMPConnector
	isSet bool
}

func (v NullableCMPConnector) Get() *CMPConnector {
	return v.value
}

func (v *NullableCMPConnector) Set(val *CMPConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableCMPConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableCMPConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCMPConnector(val *CMPConnector) *NullableCMPConnector {
	return &NullableCMPConnector{value: val, isSet: true}
}

func (v NullableCMPConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCMPConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
