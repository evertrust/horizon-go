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

// checks if the DiscoveryEventResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DiscoveryEventResponse{}

// DiscoveryEventResponse struct for DiscoveryEventResponse
type DiscoveryEventResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// The identifier of the principal that was used when the event was raised
	ActorId utils.NullableString `json:"actorId,omitempty"`
	// The name of the discovery campaign concerned by the event
	Campaign string `json:"campaign"`
	// The ID of the certificate concerned by the event (in Horizon)
	CertificateId utils.NullableString `json:"certificateId,omitempty"`
	ClientId      utils.NullableString `json:"clientId,omitempty"`
	// The IP of the machine where the Horizon client is running from
	ClientIp utils.NullableString `json:"clientIp,omitempty"`
	// The version of the discovery client that raised the event
	ClientVersion utils.NullableString `json:"clientVersion,omitempty"`
	// The code of the event to raise in the discovery events
	Code string `json:"code"`
	// The error code of the event
	ErrorCode utils.NullableString `json:"errorCode,omitempty"`
	// The error message of the event
	ErrorMessage utils.NullableString `json:"errorMessage,omitempty"`
	// The hostname concerned by the event
	Hostname utils.NullableString `json:"hostname,omitempty"`
	// The IP address concerned by the event
	Ip utils.NullableString `json:"ip,omitempty"`
	// The TCP port concerned by the event
	Port     utils.NullableInt64 `json:"port,omitempty"`
	RemoveAt utils.NullableInt64 `json:"removeAt,omitempty"`
	// The ID of the discovery feed session
	SessionId utils.NullableString `json:"sessionId,omitempty"`
	// The type of discovery that raised the event
	Source utils.NullableString `json:"source,omitempty"`
	// The type of event to raise
	Status string `json:"status"`
	// When did the event occur (Unix timestamp in milliseconds)
	Timestamp            utils.NullableInt64 `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveryEventResponse DiscoveryEventResponse

// NewDiscoveryEventResponse instantiates a new DiscoveryEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryEventResponse(id string, campaign string, code string, status string) *DiscoveryEventResponse {
	this := DiscoveryEventResponse{}
	this.Id = id
	this.Campaign = campaign
	this.Code = code
	this.Status = status
	return &this
}

// NewDiscoveryEventResponseWithDefaults instantiates a new DiscoveryEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryEventResponseWithDefaults() *DiscoveryEventResponse {
	this := DiscoveryEventResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DiscoveryEventResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiscoveryEventResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiscoveryEventResponse) SetId(v string) {
	o.Id = v
}

// GetActorId returns the ActorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetActorId() string {
	if o == nil || utils.IsNil(o.ActorId.Get()) {
		var ret string
		return ret
	}
	return *o.ActorId.Get()
}

// GetActorIdOk returns a tuple with the ActorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetActorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActorId.Get(), o.ActorId.IsSet()
}

// HasActorId returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasActorId() bool {
	if o != nil && o.ActorId.IsSet() {
		return true
	}

	return false
}

// SetActorId gets a reference to the given NullableString and assigns it to the ActorId field.
func (o *DiscoveryEventResponse) SetActorId(v string) {
	o.ActorId.Set(&v)
}

// SetActorIdNil sets the value for ActorId to be an explicit nil
func (o *DiscoveryEventResponse) SetActorIdNil() {
	o.ActorId.Set(nil)
}

// UnsetActorId ensures that no value is present for ActorId, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetActorId() {
	o.ActorId.Unset()
}

// GetCampaign returns the Campaign field value
func (o *DiscoveryEventResponse) GetCampaign() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value
// and a boolean to check if the value has been set.
func (o *DiscoveryEventResponse) GetCampaignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaign, true
}

// SetCampaign sets field value
func (o *DiscoveryEventResponse) SetCampaign(v string) {
	o.Campaign = v
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetCertificateId() string {
	if o == nil || utils.IsNil(o.CertificateId.Get()) {
		var ret string
		return ret
	}
	return *o.CertificateId.Get()
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetCertificateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateId.Get(), o.CertificateId.IsSet()
}

// HasCertificateId returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasCertificateId() bool {
	if o != nil && o.CertificateId.IsSet() {
		return true
	}

	return false
}

// SetCertificateId gets a reference to the given NullableString and assigns it to the CertificateId field.
func (o *DiscoveryEventResponse) SetCertificateId(v string) {
	o.CertificateId.Set(&v)
}

// SetCertificateIdNil sets the value for CertificateId to be an explicit nil
func (o *DiscoveryEventResponse) SetCertificateIdNil() {
	o.CertificateId.Set(nil)
}

// UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetCertificateId() {
	o.CertificateId.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetClientId() string {
	if o == nil || utils.IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *DiscoveryEventResponse) SetClientId(v string) {
	o.ClientId.Set(&v)
}

// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *DiscoveryEventResponse) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetClientId() {
	o.ClientId.Unset()
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetClientIp() string {
	if o == nil || utils.IsNil(o.ClientIp.Get()) {
		var ret string
		return ret
	}
	return *o.ClientIp.Get()
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetClientIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientIp.Get(), o.ClientIp.IsSet()
}

// HasClientIp returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasClientIp() bool {
	if o != nil && o.ClientIp.IsSet() {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given NullableString and assigns it to the ClientIp field.
func (o *DiscoveryEventResponse) SetClientIp(v string) {
	o.ClientIp.Set(&v)
}

// SetClientIpNil sets the value for ClientIp to be an explicit nil
func (o *DiscoveryEventResponse) SetClientIpNil() {
	o.ClientIp.Set(nil)
}

// UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetClientIp() {
	o.ClientIp.Unset()
}

// GetClientVersion returns the ClientVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetClientVersion() string {
	if o == nil || utils.IsNil(o.ClientVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ClientVersion.Get()
}

// GetClientVersionOk returns a tuple with the ClientVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetClientVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientVersion.Get(), o.ClientVersion.IsSet()
}

// HasClientVersion returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasClientVersion() bool {
	if o != nil && o.ClientVersion.IsSet() {
		return true
	}

	return false
}

// SetClientVersion gets a reference to the given NullableString and assigns it to the ClientVersion field.
func (o *DiscoveryEventResponse) SetClientVersion(v string) {
	o.ClientVersion.Set(&v)
}

// SetClientVersionNil sets the value for ClientVersion to be an explicit nil
func (o *DiscoveryEventResponse) SetClientVersionNil() {
	o.ClientVersion.Set(nil)
}

// UnsetClientVersion ensures that no value is present for ClientVersion, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetClientVersion() {
	o.ClientVersion.Unset()
}

// GetCode returns the Code field value
func (o *DiscoveryEventResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DiscoveryEventResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DiscoveryEventResponse) SetCode(v string) {
	o.Code = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetErrorCode() string {
	if o == nil || utils.IsNil(o.ErrorCode.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableString and assigns it to the ErrorCode field.
func (o *DiscoveryEventResponse) SetErrorCode(v string) {
	o.ErrorCode.Set(&v)
}

// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *DiscoveryEventResponse) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetErrorMessage() string {
	if o == nil || utils.IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *DiscoveryEventResponse) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *DiscoveryEventResponse) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetHostname returns the Hostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetHostname() string {
	if o == nil || utils.IsNil(o.Hostname.Get()) {
		var ret string
		return ret
	}
	return *o.Hostname.Get()
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hostname.Get(), o.Hostname.IsSet()
}

// HasHostname returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasHostname() bool {
	if o != nil && o.Hostname.IsSet() {
		return true
	}

	return false
}

// SetHostname gets a reference to the given NullableString and assigns it to the Hostname field.
func (o *DiscoveryEventResponse) SetHostname(v string) {
	o.Hostname.Set(&v)
}

// SetHostnameNil sets the value for Hostname to be an explicit nil
func (o *DiscoveryEventResponse) SetHostnameNil() {
	o.Hostname.Set(nil)
}

// UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetHostname() {
	o.Hostname.Unset()
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetIp() string {
	if o == nil || utils.IsNil(o.Ip.Get()) {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *DiscoveryEventResponse) SetIp(v string) {
	o.Ip.Set(&v)
}

// SetIpNil sets the value for Ip to be an explicit nil
func (o *DiscoveryEventResponse) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetIp() {
	o.Ip.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetPort() int64 {
	if o == nil || utils.IsNil(o.Port.Get()) {
		var ret int64
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetPortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt64 and assigns it to the Port field.
func (o *DiscoveryEventResponse) SetPort(v int64) {
	o.Port.Set(&v)
}

// SetPortNil sets the value for Port to be an explicit nil
func (o *DiscoveryEventResponse) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetPort() {
	o.Port.Unset()
}

// GetRemoveAt returns the RemoveAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetRemoveAt() int64 {
	if o == nil || utils.IsNil(o.RemoveAt.Get()) {
		var ret int64
		return ret
	}
	return *o.RemoveAt.Get()
}

// GetRemoveAtOk returns a tuple with the RemoveAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetRemoveAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveAt.Get(), o.RemoveAt.IsSet()
}

// HasRemoveAt returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasRemoveAt() bool {
	if o != nil && o.RemoveAt.IsSet() {
		return true
	}

	return false
}

// SetRemoveAt gets a reference to the given NullableInt64 and assigns it to the RemoveAt field.
func (o *DiscoveryEventResponse) SetRemoveAt(v int64) {
	o.RemoveAt.Set(&v)
}

// SetRemoveAtNil sets the value for RemoveAt to be an explicit nil
func (o *DiscoveryEventResponse) SetRemoveAtNil() {
	o.RemoveAt.Set(nil)
}

// UnsetRemoveAt ensures that no value is present for RemoveAt, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetRemoveAt() {
	o.RemoveAt.Unset()
}

// GetSessionId returns the SessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetSessionId() string {
	if o == nil || utils.IsNil(o.SessionId.Get()) {
		var ret string
		return ret
	}
	return *o.SessionId.Get()
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionId.Get(), o.SessionId.IsSet()
}

// HasSessionId returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasSessionId() bool {
	if o != nil && o.SessionId.IsSet() {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given NullableString and assigns it to the SessionId field.
func (o *DiscoveryEventResponse) SetSessionId(v string) {
	o.SessionId.Set(&v)
}

// SetSessionIdNil sets the value for SessionId to be an explicit nil
func (o *DiscoveryEventResponse) SetSessionIdNil() {
	o.SessionId.Set(nil)
}

// UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetSessionId() {
	o.SessionId.Unset()
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetSource() string {
	if o == nil || utils.IsNil(o.Source.Get()) {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableString and assigns it to the Source field.
func (o *DiscoveryEventResponse) SetSource(v string) {
	o.Source.Set(&v)
}

// SetSourceNil sets the value for Source to be an explicit nil
func (o *DiscoveryEventResponse) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetSource() {
	o.Source.Unset()
}

// GetStatus returns the Status field value
func (o *DiscoveryEventResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DiscoveryEventResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DiscoveryEventResponse) SetStatus(v string) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryEventResponse) GetTimestamp() int64 {
	if o == nil || utils.IsNil(o.Timestamp.Get()) {
		var ret int64
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryEventResponse) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DiscoveryEventResponse) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableInt64 and assigns it to the Timestamp field.
func (o *DiscoveryEventResponse) SetTimestamp(v int64) {
	o.Timestamp.Set(&v)
}

// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *DiscoveryEventResponse) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *DiscoveryEventResponse) UnsetTimestamp() {
	o.Timestamp.Unset()
}

func (o DiscoveryEventResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryEventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.ActorId.IsSet() {
		toSerialize["actorId"] = o.ActorId.Get()
	}
	toSerialize["campaign"] = o.Campaign
	if o.CertificateId.IsSet() {
		toSerialize["certificateId"] = o.CertificateId.Get()
	}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	if o.ClientIp.IsSet() {
		toSerialize["clientIp"] = o.ClientIp.Get()
	}
	if o.ClientVersion.IsSet() {
		toSerialize["clientVersion"] = o.ClientVersion.Get()
	}
	toSerialize["code"] = o.Code
	if o.ErrorCode.IsSet() {
		toSerialize["errorCode"] = o.ErrorCode.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.Hostname.IsSet() {
		toSerialize["hostname"] = o.Hostname.Get()
	}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if o.RemoveAt.IsSet() {
		toSerialize["removeAt"] = o.RemoveAt.Get()
	}
	if o.SessionId.IsSet() {
		toSerialize["sessionId"] = o.SessionId.Get()
	}
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	toSerialize["status"] = o.Status
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoveryEventResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"campaign",
		"code",
		"status",
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

	varDiscoveryEventResponse := _DiscoveryEventResponse{}

	err = json.Unmarshal(data, &varDiscoveryEventResponse)

	if err != nil {
		return err
	}

	*o = DiscoveryEventResponse(varDiscoveryEventResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "actorId")
		delete(additionalProperties, "campaign")
		delete(additionalProperties, "certificateId")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "clientIp")
		delete(additionalProperties, "clientVersion")
		delete(additionalProperties, "code")
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "hostname")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "port")
		delete(additionalProperties, "removeAt")
		delete(additionalProperties, "sessionId")
		delete(additionalProperties, "source")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveryEventResponse struct {
	value *DiscoveryEventResponse
	isSet bool
}

func (v NullableDiscoveryEventResponse) Get() *DiscoveryEventResponse {
	return v.value
}

func (v *NullableDiscoveryEventResponse) Set(val *DiscoveryEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryEventResponse(val *DiscoveryEventResponse) *NullableDiscoveryEventResponse {
	return &NullableDiscoveryEventResponse{value: val, isSet: true}
}

func (v NullableDiscoveryEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
