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

// checks if the OTPKIConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OTPKIConnectorResponse{}

// OTPKIConnectorResponse struct for OTPKIConnectorResponse
type OTPKIConnectorResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Name of the `certificate` [credentials](#tag/security.credentials) to use to authenticate on the PKI
	AuthenticationCredentials string                     `json:"authenticationCredentials"`
	EmailMap                  utils.NullableString       `json:"emailMap,omitempty"`
	EndPoint                  string                     `json:"endPoint"`
	Name                      string                     `json:"name"`
	Profile                   string                     `json:"profile"`
	Proxy                     utils.NullableString       `json:"proxy,omitempty"`
	Queue                     utils.NullableString       `json:"queue,omitempty"`
	SanDnsMap                 utils.NullableString       `json:"sanDnsMap,omitempty"`
	SanEmailMap               utils.NullableString       `json:"sanEmailMap,omitempty"`
	Status                    NullablePKIConnectorStatus `json:"status,omitempty"`
	Timeout                   utils.NullableString       `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                      string                     `json:"type"`
	UidMap                    utils.NullableString       `json:"uidMap,omitempty"`
	Zone                      utils.NullableString       `json:"zone,omitempty"`
	// The name of the label where the zone value is stored on an enrolled certificate
	ZoneLabel            utils.NullableString `json:"zoneLabel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OTPKIConnectorResponse OTPKIConnectorResponse

// NewOTPKIConnectorResponse instantiates a new OTPKIConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTPKIConnectorResponse(id string, authenticationCredentials string, endPoint string, name string, profile string, type_ string) *OTPKIConnectorResponse {
	this := OTPKIConnectorResponse{}
	this.Id = id
	this.AuthenticationCredentials = authenticationCredentials
	this.EndPoint = endPoint
	this.Name = name
	this.Profile = profile
	this.Type = type_
	return &this
}

// NewOTPKIConnectorResponseWithDefaults instantiates a new OTPKIConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTPKIConnectorResponseWithDefaults() *OTPKIConnectorResponse {
	this := OTPKIConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OTPKIConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OTPKIConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OTPKIConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value
func (o *OTPKIConnectorResponse) GetAuthenticationCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationCredentials
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value
// and a boolean to check if the value has been set.
func (o *OTPKIConnectorResponse) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationCredentials, true
}

// SetAuthenticationCredentials sets field value
func (o *OTPKIConnectorResponse) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials = v
}

// GetEmailMap returns the EmailMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OTPKIConnectorResponse) GetEmailMap() string {
	if o == nil || utils.IsNil(o.EmailMap.Get()) {
		var ret string
		return ret
	}
	return *o.EmailMap.Get()
}

// GetEmailMapOk returns a tuple with the EmailMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OTPKIConnectorResponse) GetEmailMapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailMap.Get(), o.EmailMap.IsSet()
}

// HasEmailMap returns a boolean if a field has been set.
func (o *OTPKIConnectorResponse) HasEmailMap() bool {
	if o != nil && o.EmailMap.IsSet() {
		return true
	}

	return false
}

// SetEmailMap gets a reference to the given NullableString and assigns it to the EmailMap field.
func (o *OTPKIConnectorResponse) SetEmailMap(v string) {
	o.EmailMap.Set(&v)
}

// SetEmailMapNil sets the value for EmailMap to be an explicit nil
func (o *OTPKIConnectorResponse) SetEmailMapNil() {
	o.EmailMap.Set(nil)
}

// UnsetEmailMap ensures that no value is present for EmailMap, not even an explicit nil
func (o *OTPKIConnectorResponse) UnsetEmailMap() {
	o.EmailMap.Unset()
}

// GetEndPoint returns the EndPoint field value
func (o *OTPKIConnectorResponse) GetEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value
// and a boolean to check if the value has been set.
func (o *OTPKIConnectorResponse) GetEndPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPoint, true
}

// SetEndPoint sets field value
func (o *OTPKIConnectorResponse) SetEndPoint(v string) {
	o.EndPoint = v
}

// GetName returns the Name field value
func (o *OTPKIConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OTPKIConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OTPKIConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetProfile returns the Profile field value
func (o *OTPKIConnectorResponse) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *OTPKIConnectorResponse) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *OTPKIConnectorResponse) SetProfile(v string) {
	o.Profile = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OTPKIConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OTPKIConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *OTPKIConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *OTPKIConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *OTPKIConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *OTPKIConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OTPKIConnectorResponse) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OTPKIConnectorResponse) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *OTPKIConnectorResponse) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *OTPKIConnectorResponse) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *OTPKIConnectorResponse) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *OTPKIConnectorResponse) UnsetQueue() {
	o.Queue.Unset()
}

// GetSanDnsMap returns the SanDnsMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OTPKIConnectorResponse) GetSanDnsMap() string {
	if o == nil || utils.IsNil(o.SanDnsMap.Get()) {
		var ret string
		return ret
	}
	return *o.SanDnsMap.Get()
}

// GetSanDnsMapOk returns a tuple with the SanDnsMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OTPKIConnectorResponse) GetSanDnsMapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SanDnsMap.Get(), o.SanDnsMap.IsSet()
}

// HasSanDnsMap returns a boolean if a field has been set.
func (o *OTPKIConnectorResponse) HasSanDnsMap() bool {
	if o != nil && o.SanDnsMap.IsSet() {
		return true
	}

	return false
}

// SetSanDnsMap gets a reference to the given NullableString and assigns it to the SanDnsMap field.
func (o *OTPKIConnectorResponse) SetSanDnsMap(v string) {
	o.SanDnsMap.Set(&v)
}

// SetSanDnsMapNil sets the value for SanDnsMap to be an explicit nil
func (o *OTPKIConnectorResponse) SetSanDnsMapNil() {
	o.SanDnsMap.Set(nil)
}

// UnsetSanDnsMap ensures that no value is present for SanDnsMap, not even an explicit nil
func (o *OTPKIConnectorResponse) UnsetSanDnsMap() {
	o.SanDnsMap.Unset()
}

// GetSanEmailMap returns the SanEmailMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OTPKIConnectorResponse) GetSanEmailMap() string {
	if o == nil || utils.IsNil(o.SanEmailMap.Get()) {
		var ret string
		return ret
	}
	return *o.SanEmailMap.Get()
}

// GetSanEmailMapOk returns a tuple with the SanEmailMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OTPKIConnectorResponse) GetSanEmailMapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SanEmailMap.Get(), o.SanEmailMap.IsSet()
}

// HasSanEmailMap returns a boolean if a field has been set.
func (o *OTPKIConnectorResponse) HasSanEmailMap() bool {
	if o != nil && o.SanEmailMap.IsSet() {
		return true
	}

	return false
}

// SetSanEmailMap gets a reference to the given NullableString and assigns it to the SanEmailMap field.
func (o *OTPKIConnectorResponse) SetSanEmailMap(v string) {
	o.SanEmailMap.Set(&v)
}

// SetSanEmailMapNil sets the value for SanEmailMap to be an explicit nil
func (o *OTPKIConnectorResponse) SetSanEmailMapNil() {
	o.SanEmailMap.Set(nil)
}

// UnsetSanEmailMap ensures that no value is present for SanEmailMap, not even an explicit nil
func (o *OTPKIConnectorResponse) UnsetSanEmailMap() {
	o.SanEmailMap.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OTPKIConnectorResponse) GetStatus() PKIConnectorStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret PKIConnectorStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OTPKIConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *OTPKIConnectorResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePKIConnectorStatus and assigns it to the Status field.
func (o *OTPKIConnectorResponse) SetStatus(v PKIConnectorStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *OTPKIConnectorResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *OTPKIConnectorResponse) UnsetStatus() {
	o.Status.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OTPKIConnectorResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OTPKIConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *OTPKIConnectorResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *OTPKIConnectorResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *OTPKIConnectorResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *OTPKIConnectorResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *OTPKIConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OTPKIConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OTPKIConnectorResponse) SetType(v string) {
	o.Type = v
}

// GetUidMap returns the UidMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OTPKIConnectorResponse) GetUidMap() string {
	if o == nil || utils.IsNil(o.UidMap.Get()) {
		var ret string
		return ret
	}
	return *o.UidMap.Get()
}

// GetUidMapOk returns a tuple with the UidMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OTPKIConnectorResponse) GetUidMapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UidMap.Get(), o.UidMap.IsSet()
}

// HasUidMap returns a boolean if a field has been set.
func (o *OTPKIConnectorResponse) HasUidMap() bool {
	if o != nil && o.UidMap.IsSet() {
		return true
	}

	return false
}

// SetUidMap gets a reference to the given NullableString and assigns it to the UidMap field.
func (o *OTPKIConnectorResponse) SetUidMap(v string) {
	o.UidMap.Set(&v)
}

// SetUidMapNil sets the value for UidMap to be an explicit nil
func (o *OTPKIConnectorResponse) SetUidMapNil() {
	o.UidMap.Set(nil)
}

// UnsetUidMap ensures that no value is present for UidMap, not even an explicit nil
func (o *OTPKIConnectorResponse) UnsetUidMap() {
	o.UidMap.Unset()
}

// GetZone returns the Zone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OTPKIConnectorResponse) GetZone() string {
	if o == nil || utils.IsNil(o.Zone.Get()) {
		var ret string
		return ret
	}
	return *o.Zone.Get()
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OTPKIConnectorResponse) GetZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Zone.Get(), o.Zone.IsSet()
}

// HasZone returns a boolean if a field has been set.
func (o *OTPKIConnectorResponse) HasZone() bool {
	if o != nil && o.Zone.IsSet() {
		return true
	}

	return false
}

// SetZone gets a reference to the given NullableString and assigns it to the Zone field.
func (o *OTPKIConnectorResponse) SetZone(v string) {
	o.Zone.Set(&v)
}

// SetZoneNil sets the value for Zone to be an explicit nil
func (o *OTPKIConnectorResponse) SetZoneNil() {
	o.Zone.Set(nil)
}

// UnsetZone ensures that no value is present for Zone, not even an explicit nil
func (o *OTPKIConnectorResponse) UnsetZone() {
	o.Zone.Unset()
}

// GetZoneLabel returns the ZoneLabel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OTPKIConnectorResponse) GetZoneLabel() string {
	if o == nil || utils.IsNil(o.ZoneLabel.Get()) {
		var ret string
		return ret
	}
	return *o.ZoneLabel.Get()
}

// GetZoneLabelOk returns a tuple with the ZoneLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OTPKIConnectorResponse) GetZoneLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZoneLabel.Get(), o.ZoneLabel.IsSet()
}

// HasZoneLabel returns a boolean if a field has been set.
func (o *OTPKIConnectorResponse) HasZoneLabel() bool {
	if o != nil && o.ZoneLabel.IsSet() {
		return true
	}

	return false
}

// SetZoneLabel gets a reference to the given NullableString and assigns it to the ZoneLabel field.
func (o *OTPKIConnectorResponse) SetZoneLabel(v string) {
	o.ZoneLabel.Set(&v)
}

// SetZoneLabelNil sets the value for ZoneLabel to be an explicit nil
func (o *OTPKIConnectorResponse) SetZoneLabelNil() {
	o.ZoneLabel.Set(nil)
}

// UnsetZoneLabel ensures that no value is present for ZoneLabel, not even an explicit nil
func (o *OTPKIConnectorResponse) UnsetZoneLabel() {
	o.ZoneLabel.Unset()
}

func (o OTPKIConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTPKIConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["authenticationCredentials"] = o.AuthenticationCredentials
	if o.EmailMap.IsSet() {
		toSerialize["emailMap"] = o.EmailMap.Get()
	}
	toSerialize["endPoint"] = o.EndPoint
	toSerialize["name"] = o.Name
	toSerialize["profile"] = o.Profile
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.SanDnsMap.IsSet() {
		toSerialize["sanDnsMap"] = o.SanDnsMap.Get()
	}
	if o.SanEmailMap.IsSet() {
		toSerialize["sanEmailMap"] = o.SanEmailMap.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	toSerialize["type"] = o.Type
	if o.UidMap.IsSet() {
		toSerialize["uidMap"] = o.UidMap.Get()
	}
	if o.Zone.IsSet() {
		toSerialize["zone"] = o.Zone.Get()
	}
	if o.ZoneLabel.IsSet() {
		toSerialize["zoneLabel"] = o.ZoneLabel.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OTPKIConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"authenticationCredentials",
		"endPoint",
		"name",
		"profile",
		"type",
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

	varOTPKIConnectorResponse := _OTPKIConnectorResponse{}

	err = json.Unmarshal(data, &varOTPKIConnectorResponse)

	if err != nil {
		return err
	}

	*o = OTPKIConnectorResponse(varOTPKIConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "authenticationCredentials")
		delete(additionalProperties, "emailMap")
		delete(additionalProperties, "endPoint")
		delete(additionalProperties, "name")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "sanDnsMap")
		delete(additionalProperties, "sanEmailMap")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		delete(additionalProperties, "uidMap")
		delete(additionalProperties, "zone")
		delete(additionalProperties, "zoneLabel")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOTPKIConnectorResponse struct {
	value *OTPKIConnectorResponse
	isSet bool
}

func (v NullableOTPKIConnectorResponse) Get() *OTPKIConnectorResponse {
	return v.value
}

func (v *NullableOTPKIConnectorResponse) Set(val *OTPKIConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOTPKIConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOTPKIConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTPKIConnectorResponse(val *OTPKIConnectorResponse) *NullableOTPKIConnectorResponse {
	return &NullableOTPKIConnectorResponse{value: val, isSet: true}
}

func (v NullableOTPKIConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTPKIConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
