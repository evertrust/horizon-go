/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the OrderResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OrderResponse{}

// OrderResponse struct for OrderResponse
type OrderResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Object internal ID
	AccountId      string          `json:"accountId"`
	Authorizations []Authorization `json:"authorizations"`
	// Object internal ID
	Certificate          utils.NullableString `json:"certificate,omitempty"`
	ContactEmail         utils.NullableString `json:"contactEmail,omitempty"`
	Expires              utils.NullableInt64  `json:"expires,omitempty"`
	InitialIp            utils.NullableString `json:"initialIp,omitempty"`
	Label                []OrderLabelInner    `json:"label,omitempty"`
	Metadata             *map[string]string   `json:"metadata,omitempty"`
	NotAfter             utils.NullableInt64  `json:"notAfter,omitempty"`
	NotBefore            utils.NullableInt64  `json:"notBefore,omitempty"`
	Owner                utils.NullableString `json:"owner,omitempty"`
	Profile              string               `json:"profile"`
	RemoveAt             utils.NullableInt64  `json:"removeAt,omitempty"`
	Status               OrderStatus          `json:"status"`
	Team                 utils.NullableString `json:"team,omitempty"`
	Thumbprint           utils.NullableString `json:"thumbprint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderResponse OrderResponse

// NewOrderResponse instantiates a new OrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderResponse(id string, accountId string, authorizations []Authorization, profile string, status OrderStatus) *OrderResponse {
	this := OrderResponse{}
	this.Id = id
	this.AccountId = accountId
	this.Authorizations = authorizations
	this.Profile = profile
	this.Status = status
	return &this
}

// NewOrderResponseWithDefaults instantiates a new OrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderResponseWithDefaults() *OrderResponse {
	this := OrderResponse{}
	return &this
}

// GetId returns the Id field value
func (o *OrderResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderResponse) SetId(v string) {
	o.Id = v
}

// GetAccountId returns the AccountId field value
func (o *OrderResponse) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *OrderResponse) SetAccountId(v string) {
	o.AccountId = v
}

// GetAuthorizations returns the Authorizations field value
func (o *OrderResponse) GetAuthorizations() []Authorization {
	if o == nil {
		var ret []Authorization
		return ret
	}

	return o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetAuthorizationsOk() ([]Authorization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authorizations, true
}

// SetAuthorizations sets field value
func (o *OrderResponse) SetAuthorizations(v []Authorization) {
	o.Authorizations = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponse) GetCertificate() string {
	if o == nil || utils.IsNil(o.Certificate.Get()) {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponse) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *OrderResponse) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *OrderResponse) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *OrderResponse) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *OrderResponse) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponse) GetContactEmail() string {
	if o == nil || utils.IsNil(o.ContactEmail.Get()) {
		var ret string
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponse) GetContactEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *OrderResponse) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableString and assigns it to the ContactEmail field.
func (o *OrderResponse) SetContactEmail(v string) {
	o.ContactEmail.Set(&v)
}

// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *OrderResponse) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *OrderResponse) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponse) GetExpires() int64 {
	if o == nil || utils.IsNil(o.Expires.Get()) {
		var ret int64
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponse) GetExpiresOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *OrderResponse) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableInt64 and assigns it to the Expires field.
func (o *OrderResponse) SetExpires(v int64) {
	o.Expires.Set(&v)
}

// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *OrderResponse) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *OrderResponse) UnsetExpires() {
	o.Expires.Unset()
}

// GetInitialIp returns the InitialIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponse) GetInitialIp() string {
	if o == nil || utils.IsNil(o.InitialIp.Get()) {
		var ret string
		return ret
	}
	return *o.InitialIp.Get()
}

// GetInitialIpOk returns a tuple with the InitialIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponse) GetInitialIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitialIp.Get(), o.InitialIp.IsSet()
}

// HasInitialIp returns a boolean if a field has been set.
func (o *OrderResponse) HasInitialIp() bool {
	if o != nil && o.InitialIp.IsSet() {
		return true
	}

	return false
}

// SetInitialIp gets a reference to the given NullableString and assigns it to the InitialIp field.
func (o *OrderResponse) SetInitialIp(v string) {
	o.InitialIp.Set(&v)
}

// SetInitialIpNil sets the value for InitialIp to be an explicit nil
func (o *OrderResponse) SetInitialIpNil() {
	o.InitialIp.Set(nil)
}

// UnsetInitialIp ensures that no value is present for InitialIp, not even an explicit nil
func (o *OrderResponse) UnsetInitialIp() {
	o.InitialIp.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponse) GetLabel() []OrderLabelInner {
	if o == nil {
		var ret []OrderLabelInner
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponse) GetLabelOk() ([]OrderLabelInner, bool) {
	if o == nil || utils.IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *OrderResponse) HasLabel() bool {
	if o != nil && !utils.IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given []OrderLabelInner and assigns it to the Label field.
func (o *OrderResponse) SetLabel(v []OrderLabelInner) {
	o.Label = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OrderResponse) GetMetadata() map[string]string {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OrderResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *OrderResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponse) GetNotAfter() int64 {
	if o == nil || utils.IsNil(o.NotAfter.Get()) {
		var ret int64
		return ret
	}
	return *o.NotAfter.Get()
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponse) GetNotAfterOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotAfter.Get(), o.NotAfter.IsSet()
}

// HasNotAfter returns a boolean if a field has been set.
func (o *OrderResponse) HasNotAfter() bool {
	if o != nil && o.NotAfter.IsSet() {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given NullableInt64 and assigns it to the NotAfter field.
func (o *OrderResponse) SetNotAfter(v int64) {
	o.NotAfter.Set(&v)
}

// SetNotAfterNil sets the value for NotAfter to be an explicit nil
func (o *OrderResponse) SetNotAfterNil() {
	o.NotAfter.Set(nil)
}

// UnsetNotAfter ensures that no value is present for NotAfter, not even an explicit nil
func (o *OrderResponse) UnsetNotAfter() {
	o.NotAfter.Unset()
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponse) GetNotBefore() int64 {
	if o == nil || utils.IsNil(o.NotBefore.Get()) {
		var ret int64
		return ret
	}
	return *o.NotBefore.Get()
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponse) GetNotBeforeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotBefore.Get(), o.NotBefore.IsSet()
}

// HasNotBefore returns a boolean if a field has been set.
func (o *OrderResponse) HasNotBefore() bool {
	if o != nil && o.NotBefore.IsSet() {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given NullableInt64 and assigns it to the NotBefore field.
func (o *OrderResponse) SetNotBefore(v int64) {
	o.NotBefore.Set(&v)
}

// SetNotBeforeNil sets the value for NotBefore to be an explicit nil
func (o *OrderResponse) SetNotBeforeNil() {
	o.NotBefore.Set(nil)
}

// UnsetNotBefore ensures that no value is present for NotBefore, not even an explicit nil
func (o *OrderResponse) UnsetNotBefore() {
	o.NotBefore.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponse) GetOwner() string {
	if o == nil || utils.IsNil(o.Owner.Get()) {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponse) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *OrderResponse) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *OrderResponse) SetOwner(v string) {
	o.Owner.Set(&v)
}

// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *OrderResponse) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *OrderResponse) UnsetOwner() {
	o.Owner.Unset()
}

// GetProfile returns the Profile field value
func (o *OrderResponse) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *OrderResponse) SetProfile(v string) {
	o.Profile = v
}

// GetRemoveAt returns the RemoveAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponse) GetRemoveAt() int64 {
	if o == nil || utils.IsNil(o.RemoveAt.Get()) {
		var ret int64
		return ret
	}
	return *o.RemoveAt.Get()
}

// GetRemoveAtOk returns a tuple with the RemoveAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponse) GetRemoveAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveAt.Get(), o.RemoveAt.IsSet()
}

// HasRemoveAt returns a boolean if a field has been set.
func (o *OrderResponse) HasRemoveAt() bool {
	if o != nil && o.RemoveAt.IsSet() {
		return true
	}

	return false
}

// SetRemoveAt gets a reference to the given NullableInt64 and assigns it to the RemoveAt field.
func (o *OrderResponse) SetRemoveAt(v int64) {
	o.RemoveAt.Set(&v)
}

// SetRemoveAtNil sets the value for RemoveAt to be an explicit nil
func (o *OrderResponse) SetRemoveAtNil() {
	o.RemoveAt.Set(nil)
}

// UnsetRemoveAt ensures that no value is present for RemoveAt, not even an explicit nil
func (o *OrderResponse) UnsetRemoveAt() {
	o.RemoveAt.Unset()
}

// GetStatus returns the Status field value
func (o *OrderResponse) GetStatus() OrderStatus {
	if o == nil {
		var ret OrderStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrderResponse) GetStatusOk() (*OrderStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrderResponse) SetStatus(v OrderStatus) {
	o.Status = v
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponse) GetTeam() string {
	if o == nil || utils.IsNil(o.Team.Get()) {
		var ret string
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponse) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *OrderResponse) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableString and assigns it to the Team field.
func (o *OrderResponse) SetTeam(v string) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *OrderResponse) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *OrderResponse) UnsetTeam() {
	o.Team.Unset()
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderResponse) GetThumbprint() string {
	if o == nil || utils.IsNil(o.Thumbprint.Get()) {
		var ret string
		return ret
	}
	return *o.Thumbprint.Get()
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderResponse) GetThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thumbprint.Get(), o.Thumbprint.IsSet()
}

// HasThumbprint returns a boolean if a field has been set.
func (o *OrderResponse) HasThumbprint() bool {
	if o != nil && o.Thumbprint.IsSet() {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given NullableString and assigns it to the Thumbprint field.
func (o *OrderResponse) SetThumbprint(v string) {
	o.Thumbprint.Set(&v)
}

// SetThumbprintNil sets the value for Thumbprint to be an explicit nil
func (o *OrderResponse) SetThumbprintNil() {
	o.Thumbprint.Set(nil)
}

// UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
func (o *OrderResponse) UnsetThumbprint() {
	o.Thumbprint.Unset()
}

func (o OrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["accountId"] = o.AccountId
	toSerialize["authorizations"] = o.Authorizations
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.ContactEmail.IsSet() {
		toSerialize["contactEmail"] = o.ContactEmail.Get()
	}
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	if o.InitialIp.IsSet() {
		toSerialize["initialIp"] = o.InitialIp.Get()
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
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
	toSerialize["profile"] = o.Profile
	if o.RemoveAt.IsSet() {
		toSerialize["removeAt"] = o.RemoveAt.Get()
	}
	toSerialize["status"] = o.Status
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}
	if o.Thumbprint.IsSet() {
		toSerialize["thumbprint"] = o.Thumbprint.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"accountId",
		"authorizations",
		"profile",
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

	varOrderResponse := _OrderResponse{}

	err = json.Unmarshal(data, &varOrderResponse)

	if err != nil {
		return err
	}

	*o = OrderResponse(varOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "authorizations")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "expires")
		delete(additionalProperties, "initialIp")
		delete(additionalProperties, "label")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "notAfter")
		delete(additionalProperties, "notBefore")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "removeAt")
		delete(additionalProperties, "status")
		delete(additionalProperties, "team")
		delete(additionalProperties, "thumbprint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderResponse struct {
	value *OrderResponse
	isSet bool
}

func (v NullableOrderResponse) Get() *OrderResponse {
	return v.value
}

func (v *NullableOrderResponse) Set(val *OrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderResponse(val *OrderResponse) *NullableOrderResponse {
	return &NullableOrderResponse{value: val, isSet: true}
}

func (v NullableOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
