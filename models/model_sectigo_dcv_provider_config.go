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

// checks if the SectigoDCVProviderConfig type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SectigoDCVProviderConfig{}

// SectigoDCVProviderConfig struct for SectigoDCVProviderConfig
type SectigoDCVProviderConfig struct {
	// Name of the login/password credentials configuration holding the SCM API client, with the OAuth client id as login and the client secret as password
	Credentials string `json:"credentials"`
	// DNS method used to validate a domain. It is a fallback: a domain that already holds a CNAME or TXT validation is re-validated with its own method, and this method applies only to a domain that has never been validated or whose existing validation uses a method that cannot be published over DNS.
	DcvMethod string `json:"dcvMethod"`
	// Sectigo Certificate Manager (SCM) API base URL
	Endpoint string `json:"endpoint"`
	// Unique name of the DCV provider configuration
	Name string `json:"name"`
	// OAuth token endpoint used to obtain a bearer token for the SCM API. Defaults to Sectigo's SSO realm.
	OauthTokenEndpoint *string `json:"oauthTokenEndpoint,omitempty"`
	// Restricts the domain listing to this Sectigo organization or department. When unset, every domain of the customer account is listed.
	OrganizationId *int64 `json:"organizationId,omitempty"`
	// Name of the HTTP proxy configuration to use
	Proxy *string `json:"proxy,omitempty"`
	// Request timeout
	Timeout utils.NullableString `json:"timeout" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Provider type discriminator
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _SectigoDCVProviderConfig SectigoDCVProviderConfig

// NewSectigoDCVProviderConfig instantiates a new SectigoDCVProviderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSectigoDCVProviderConfig(credentials string, dcvMethod string, endpoint string, name string, timeout utils.NullableString, type_ string) *SectigoDCVProviderConfig {
	this := SectigoDCVProviderConfig{}
	this.Credentials = credentials
	this.DcvMethod = dcvMethod
	this.Endpoint = endpoint
	this.Name = name
	this.Timeout = timeout
	this.Type = type_
	var oauthTokenEndpoint string = "https://auth.sso.sectigo.com/auth/realms/apiclients/protocol/openid-connect/token"
	this.OauthTokenEndpoint = &oauthTokenEndpoint
	return &this
}

// NewSectigoDCVProviderConfigWithDefaults instantiates a new SectigoDCVProviderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSectigoDCVProviderConfigWithDefaults() *SectigoDCVProviderConfig {
	this := SectigoDCVProviderConfig{}
	var oauthTokenEndpoint string = "https://auth.sso.sectigo.com/auth/realms/apiclients/protocol/openid-connect/token"
	this.OauthTokenEndpoint = &oauthTokenEndpoint
	return &this
}

// GetCredentials returns the Credentials field value
func (o *SectigoDCVProviderConfig) GetCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *SectigoDCVProviderConfig) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *SectigoDCVProviderConfig) SetCredentials(v string) {
	o.Credentials = v
}

// GetDcvMethod returns the DcvMethod field value
func (o *SectigoDCVProviderConfig) GetDcvMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DcvMethod
}

// GetDcvMethodOk returns a tuple with the DcvMethod field value
// and a boolean to check if the value has been set.
func (o *SectigoDCVProviderConfig) GetDcvMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcvMethod, true
}

// SetDcvMethod sets field value
func (o *SectigoDCVProviderConfig) SetDcvMethod(v string) {
	o.DcvMethod = v
}

// GetEndpoint returns the Endpoint field value
func (o *SectigoDCVProviderConfig) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *SectigoDCVProviderConfig) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *SectigoDCVProviderConfig) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetName returns the Name field value
func (o *SectigoDCVProviderConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SectigoDCVProviderConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SectigoDCVProviderConfig) SetName(v string) {
	o.Name = v
}

// GetOauthTokenEndpoint returns the OauthTokenEndpoint field value if set, zero value otherwise.
func (o *SectigoDCVProviderConfig) GetOauthTokenEndpoint() string {
	if o == nil || utils.IsNil(o.OauthTokenEndpoint) {
		var ret string
		return ret
	}
	return *o.OauthTokenEndpoint
}

// GetOauthTokenEndpointOk returns a tuple with the OauthTokenEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SectigoDCVProviderConfig) GetOauthTokenEndpointOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OauthTokenEndpoint) {
		return nil, false
	}
	return o.OauthTokenEndpoint, true
}

// HasOauthTokenEndpoint returns a boolean if a field has been set.
func (o *SectigoDCVProviderConfig) HasOauthTokenEndpoint() bool {
	if o != nil && !utils.IsNil(o.OauthTokenEndpoint) {
		return true
	}

	return false
}

// SetOauthTokenEndpoint gets a reference to the given string and assigns it to the OauthTokenEndpoint field.
func (o *SectigoDCVProviderConfig) SetOauthTokenEndpoint(v string) {
	o.OauthTokenEndpoint = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *SectigoDCVProviderConfig) GetOrganizationId() int64 {
	if o == nil || utils.IsNil(o.OrganizationId) {
		var ret int64
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SectigoDCVProviderConfig) GetOrganizationIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *SectigoDCVProviderConfig) HasOrganizationId() bool {
	if o != nil && !utils.IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int64 and assigns it to the OrganizationId field.
func (o *SectigoDCVProviderConfig) SetOrganizationId(v int64) {
	o.OrganizationId = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *SectigoDCVProviderConfig) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy) {
		var ret string
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SectigoDCVProviderConfig) GetProxyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *SectigoDCVProviderConfig) HasProxy() bool {
	if o != nil && !utils.IsNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given string and assigns it to the Proxy field.
func (o *SectigoDCVProviderConfig) SetProxy(v string) {
	o.Proxy = &v
}

// GetTimeout returns the Timeout field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SectigoDCVProviderConfig) GetTimeout() string {
	if o == nil || o.Timeout.Get() == nil {
		var ret string
		return ret
	}

	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SectigoDCVProviderConfig) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// SetTimeout sets field value
func (o *SectigoDCVProviderConfig) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// GetType returns the Type field value
func (o *SectigoDCVProviderConfig) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SectigoDCVProviderConfig) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SectigoDCVProviderConfig) SetType(v string) {
	o.Type = v
}

func (o SectigoDCVProviderConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SectigoDCVProviderConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.Credentials
	toSerialize["dcvMethod"] = o.DcvMethod
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["name"] = o.Name
	if !utils.IsNil(o.OauthTokenEndpoint) {
		toSerialize["oauthTokenEndpoint"] = o.OauthTokenEndpoint
	}
	if !utils.IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !utils.IsNil(o.Proxy) {
		toSerialize["proxy"] = o.Proxy
	}
	toSerialize["timeout"] = o.Timeout.Get()
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SectigoDCVProviderConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"credentials",
		"dcvMethod",
		"endpoint",
		"name",
		"timeout",
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

	varSectigoDCVProviderConfig := _SectigoDCVProviderConfig{}

	err = json.Unmarshal(data, &varSectigoDCVProviderConfig)

	if err != nil {
		return err
	}

	*o = SectigoDCVProviderConfig(varSectigoDCVProviderConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "dcvMethod")
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "name")
		delete(additionalProperties, "oauthTokenEndpoint")
		delete(additionalProperties, "organizationId")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSectigoDCVProviderConfig struct {
	value *SectigoDCVProviderConfig
	isSet bool
}

func (v NullableSectigoDCVProviderConfig) Get() *SectigoDCVProviderConfig {
	return v.value
}

func (v *NullableSectigoDCVProviderConfig) Set(val *SectigoDCVProviderConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSectigoDCVProviderConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSectigoDCVProviderConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSectigoDCVProviderConfig(val *SectigoDCVProviderConfig) *NullableSectigoDCVProviderConfig {
	return &NullableSectigoDCVProviderConfig{value: val, isSet: true}
}

func (v NullableSectigoDCVProviderConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSectigoDCVProviderConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
