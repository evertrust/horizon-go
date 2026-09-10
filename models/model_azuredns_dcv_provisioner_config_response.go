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

// checks if the AzurednsDCVProvisionerConfigResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AzurednsDCVProvisionerConfigResponse{}

// AzurednsDCVProvisionerConfigResponse struct for AzurednsDCVProvisionerConfigResponse
type AzurednsDCVProvisionerConfigResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Azure AD authority host URL override
	AuthorityHost utils.NullableString `json:"authorityHost,omitempty"`
	// Name of the credentials configuration holding the Azure client ID and secret
	Credentials utils.NullableString `json:"credentials,omitempty"`
	// DNS zone used for CNAME delegation when provisioning DCV challenges
	DelegationZone utils.NullableString `json:"delegationZone,omitempty"`
	// Azure DNS API endpoint URL override
	Endpoint utils.NullableString `json:"endpoint,omitempty"`
	// Unique name of the DCV provisioner configuration
	Name string `json:"name"`
	// Name of the HTTP proxy configuration to use
	Proxy utils.NullableString `json:"proxy,omitempty"`
	// Azure resource group containing the DNS zones
	ResourceGroupName string `json:"resourceGroupName"`
	// Azure subscription ID
	SubscriptionId string `json:"subscriptionId"`
	// Azure Active Directory tenant ID
	TenantId string `json:"tenantId"`
	// Request timeout
	Timeout utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// DNS record cache duration
	Ttl utils.NullableString `json:"ttl" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Provisioner type discriminator
	Type string `json:"type"`
	// A set of Azure DNS zone ID with regex
	ZoneIdMappings       []ZoneIdMappings `json:"zoneIdMappings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AzurednsDCVProvisionerConfigResponse AzurednsDCVProvisionerConfigResponse

// NewAzurednsDCVProvisionerConfigResponse instantiates a new AzurednsDCVProvisionerConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzurednsDCVProvisionerConfigResponse(id string, name string, resourceGroupName string, subscriptionId string, tenantId string, ttl utils.NullableString, type_ string) *AzurednsDCVProvisionerConfigResponse {
	this := AzurednsDCVProvisionerConfigResponse{}
	this.Id = id
	this.Name = name
	this.ResourceGroupName = resourceGroupName
	this.SubscriptionId = subscriptionId
	this.TenantId = tenantId
	this.Ttl = ttl
	this.Type = type_
	return &this
}

// NewAzurednsDCVProvisionerConfigResponseWithDefaults instantiates a new AzurednsDCVProvisionerConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzurednsDCVProvisionerConfigResponseWithDefaults() *AzurednsDCVProvisionerConfigResponse {
	this := AzurednsDCVProvisionerConfigResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AzurednsDCVProvisionerConfigResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AzurednsDCVProvisionerConfigResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AzurednsDCVProvisionerConfigResponse) SetId(v string) {
	o.Id = v
}

// GetAuthorityHost returns the AuthorityHost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzurednsDCVProvisionerConfigResponse) GetAuthorityHost() string {
	if o == nil || utils.IsNil(o.AuthorityHost.Get()) {
		var ret string
		return ret
	}
	return *o.AuthorityHost.Get()
}

// GetAuthorityHostOk returns a tuple with the AuthorityHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzurednsDCVProvisionerConfigResponse) GetAuthorityHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorityHost.Get(), o.AuthorityHost.IsSet()
}

// HasAuthorityHost returns a boolean if a field has been set.
func (o *AzurednsDCVProvisionerConfigResponse) HasAuthorityHost() bool {
	if o != nil && o.AuthorityHost.IsSet() {
		return true
	}

	return false
}

// SetAuthorityHost gets a reference to the given NullableString and assigns it to the AuthorityHost field.
func (o *AzurednsDCVProvisionerConfigResponse) SetAuthorityHost(v string) {
	o.AuthorityHost.Set(&v)
}

// SetAuthorityHostNil sets the value for AuthorityHost to be an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) SetAuthorityHostNil() {
	o.AuthorityHost.Set(nil)
}

// UnsetAuthorityHost ensures that no value is present for AuthorityHost, not even an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) UnsetAuthorityHost() {
	o.AuthorityHost.Unset()
}

// GetCredentials returns the Credentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzurednsDCVProvisionerConfigResponse) GetCredentials() string {
	if o == nil || utils.IsNil(o.Credentials.Get()) {
		var ret string
		return ret
	}
	return *o.Credentials.Get()
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzurednsDCVProvisionerConfigResponse) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials.Get(), o.Credentials.IsSet()
}

// HasCredentials returns a boolean if a field has been set.
func (o *AzurednsDCVProvisionerConfigResponse) HasCredentials() bool {
	if o != nil && o.Credentials.IsSet() {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given NullableString and assigns it to the Credentials field.
func (o *AzurednsDCVProvisionerConfigResponse) SetCredentials(v string) {
	o.Credentials.Set(&v)
}

// SetCredentialsNil sets the value for Credentials to be an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) SetCredentialsNil() {
	o.Credentials.Set(nil)
}

// UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) UnsetCredentials() {
	o.Credentials.Unset()
}

// GetDelegationZone returns the DelegationZone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzurednsDCVProvisionerConfigResponse) GetDelegationZone() string {
	if o == nil || utils.IsNil(o.DelegationZone.Get()) {
		var ret string
		return ret
	}
	return *o.DelegationZone.Get()
}

// GetDelegationZoneOk returns a tuple with the DelegationZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzurednsDCVProvisionerConfigResponse) GetDelegationZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DelegationZone.Get(), o.DelegationZone.IsSet()
}

// HasDelegationZone returns a boolean if a field has been set.
func (o *AzurednsDCVProvisionerConfigResponse) HasDelegationZone() bool {
	if o != nil && o.DelegationZone.IsSet() {
		return true
	}

	return false
}

// SetDelegationZone gets a reference to the given NullableString and assigns it to the DelegationZone field.
func (o *AzurednsDCVProvisionerConfigResponse) SetDelegationZone(v string) {
	o.DelegationZone.Set(&v)
}

// SetDelegationZoneNil sets the value for DelegationZone to be an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) SetDelegationZoneNil() {
	o.DelegationZone.Set(nil)
}

// UnsetDelegationZone ensures that no value is present for DelegationZone, not even an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) UnsetDelegationZone() {
	o.DelegationZone.Unset()
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzurednsDCVProvisionerConfigResponse) GetEndpoint() string {
	if o == nil || utils.IsNil(o.Endpoint.Get()) {
		var ret string
		return ret
	}
	return *o.Endpoint.Get()
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzurednsDCVProvisionerConfigResponse) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Endpoint.Get(), o.Endpoint.IsSet()
}

// HasEndpoint returns a boolean if a field has been set.
func (o *AzurednsDCVProvisionerConfigResponse) HasEndpoint() bool {
	if o != nil && o.Endpoint.IsSet() {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given NullableString and assigns it to the Endpoint field.
func (o *AzurednsDCVProvisionerConfigResponse) SetEndpoint(v string) {
	o.Endpoint.Set(&v)
}

// SetEndpointNil sets the value for Endpoint to be an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) SetEndpointNil() {
	o.Endpoint.Set(nil)
}

// UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) UnsetEndpoint() {
	o.Endpoint.Unset()
}

// GetName returns the Name field value
func (o *AzurednsDCVProvisionerConfigResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AzurednsDCVProvisionerConfigResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AzurednsDCVProvisionerConfigResponse) SetName(v string) {
	o.Name = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzurednsDCVProvisionerConfigResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzurednsDCVProvisionerConfigResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *AzurednsDCVProvisionerConfigResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *AzurednsDCVProvisionerConfigResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetResourceGroupName returns the ResourceGroupName field value
func (o *AzurednsDCVProvisionerConfigResponse) GetResourceGroupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceGroupName
}

// GetResourceGroupNameOk returns a tuple with the ResourceGroupName field value
// and a boolean to check if the value has been set.
func (o *AzurednsDCVProvisionerConfigResponse) GetResourceGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceGroupName, true
}

// SetResourceGroupName sets field value
func (o *AzurednsDCVProvisionerConfigResponse) SetResourceGroupName(v string) {
	o.ResourceGroupName = v
}

// GetSubscriptionId returns the SubscriptionId field value
func (o *AzurednsDCVProvisionerConfigResponse) GetSubscriptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value
// and a boolean to check if the value has been set.
func (o *AzurednsDCVProvisionerConfigResponse) GetSubscriptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionId, true
}

// SetSubscriptionId sets field value
func (o *AzurednsDCVProvisionerConfigResponse) SetSubscriptionId(v string) {
	o.SubscriptionId = v
}

// GetTenantId returns the TenantId field value
func (o *AzurednsDCVProvisionerConfigResponse) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *AzurednsDCVProvisionerConfigResponse) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *AzurednsDCVProvisionerConfigResponse) SetTenantId(v string) {
	o.TenantId = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AzurednsDCVProvisionerConfigResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzurednsDCVProvisionerConfigResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *AzurednsDCVProvisionerConfigResponse) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *AzurednsDCVProvisionerConfigResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *AzurednsDCVProvisionerConfigResponse) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetTtl returns the Ttl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AzurednsDCVProvisionerConfigResponse) GetTtl() string {
	if o == nil || o.Ttl.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ttl.Get()
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AzurednsDCVProvisionerConfigResponse) GetTtlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ttl.Get(), o.Ttl.IsSet()
}

// SetTtl sets field value
func (o *AzurednsDCVProvisionerConfigResponse) SetTtl(v string) {
	o.Ttl.Set(&v)
}

// GetType returns the Type field value
func (o *AzurednsDCVProvisionerConfigResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AzurednsDCVProvisionerConfigResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AzurednsDCVProvisionerConfigResponse) SetType(v string) {
	o.Type = v
}

// GetZoneIdMappings returns the ZoneIdMappings field value if set, zero value otherwise.
func (o *AzurednsDCVProvisionerConfigResponse) GetZoneIdMappings() []ZoneIdMappings {
	if o == nil || utils.IsNil(o.ZoneIdMappings) {
		var ret []ZoneIdMappings
		return ret
	}
	return o.ZoneIdMappings
}

// GetZoneIdMappingsOk returns a tuple with the ZoneIdMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurednsDCVProvisionerConfigResponse) GetZoneIdMappingsOk() ([]ZoneIdMappings, bool) {
	if o == nil || utils.IsNil(o.ZoneIdMappings) {
		return nil, false
	}
	return o.ZoneIdMappings, true
}

// HasZoneIdMappings returns a boolean if a field has been set.
func (o *AzurednsDCVProvisionerConfigResponse) HasZoneIdMappings() bool {
	if o != nil && !utils.IsNil(o.ZoneIdMappings) {
		return true
	}

	return false
}

// SetZoneIdMappings gets a reference to the given []ZoneIdMappings and assigns it to the ZoneIdMappings field.
func (o *AzurednsDCVProvisionerConfigResponse) SetZoneIdMappings(v []ZoneIdMappings) {
	o.ZoneIdMappings = v
}

func (o AzurednsDCVProvisionerConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzurednsDCVProvisionerConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.AuthorityHost.IsSet() {
		toSerialize["authorityHost"] = o.AuthorityHost.Get()
	}
	if o.Credentials.IsSet() {
		toSerialize["credentials"] = o.Credentials.Get()
	}
	if o.DelegationZone.IsSet() {
		toSerialize["delegationZone"] = o.DelegationZone.Get()
	}
	if o.Endpoint.IsSet() {
		toSerialize["endpoint"] = o.Endpoint.Get()
	}
	toSerialize["name"] = o.Name
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	toSerialize["resourceGroupName"] = o.ResourceGroupName
	toSerialize["subscriptionId"] = o.SubscriptionId
	toSerialize["tenantId"] = o.TenantId
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	toSerialize["ttl"] = o.Ttl.Get()
	toSerialize["type"] = o.Type
	if !utils.IsNil(o.ZoneIdMappings) {
		toSerialize["zoneIdMappings"] = o.ZoneIdMappings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AzurednsDCVProvisionerConfigResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"name",
		"resourceGroupName",
		"subscriptionId",
		"tenantId",
		"ttl",
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

	varAzurednsDCVProvisionerConfigResponse := _AzurednsDCVProvisionerConfigResponse{}

	err = json.Unmarshal(data, &varAzurednsDCVProvisionerConfigResponse)

	if err != nil {
		return err
	}

	*o = AzurednsDCVProvisionerConfigResponse(varAzurednsDCVProvisionerConfigResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "authorityHost")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "delegationZone")
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "name")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "resourceGroupName")
		delete(additionalProperties, "subscriptionId")
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "ttl")
		delete(additionalProperties, "type")
		delete(additionalProperties, "zoneIdMappings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAzurednsDCVProvisionerConfigResponse struct {
	value *AzurednsDCVProvisionerConfigResponse
	isSet bool
}

func (v NullableAzurednsDCVProvisionerConfigResponse) Get() *AzurednsDCVProvisionerConfigResponse {
	return v.value
}

func (v *NullableAzurednsDCVProvisionerConfigResponse) Set(val *AzurednsDCVProvisionerConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAzurednsDCVProvisionerConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAzurednsDCVProvisionerConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzurednsDCVProvisionerConfigResponse(val *AzurednsDCVProvisionerConfigResponse) *NullableAzurednsDCVProvisionerConfigResponse {
	return &NullableAzurednsDCVProvisionerConfigResponse{value: val, isSet: true}
}

func (v NullableAzurednsDCVProvisionerConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzurednsDCVProvisionerConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
