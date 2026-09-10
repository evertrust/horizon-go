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

// checks if the RESTDatasourceResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RESTDatasourceResponse{}

// RESTDatasourceResponse struct for RESTDatasourceResponse
type RESTDatasourceResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// List of attributes to fetch for this datasource
	Attributes []DataSourceOutput `json:"attributes,omitempty"`
	// The authentication type to use while making the REST call. Is linked to `credentials`.
	AuthenticationType string `json:"authenticationType"`
	// Name of the [credentials](#tag/security.credentials) to use for authentication
	Credentials *string `json:"credentials,omitempty"`
	// Description of the datasource
	Description *string `json:"description,omitempty"`
	// The localized name of the datasource
	DisplayName []LocalizedString `json:"displayName,omitempty"`
	// The success HTTP codes for the request. If the return code is not in this list, the request will be considered failed.
	ExpectedHttpCodes []int64 `json:"expectedHttpCodes"`
	// The headers of the request
	Headers []Header `json:"headers,omitempty"`
	// The HTTP method to use for the request
	Method string `json:"method"`
	// Name of the datasource
	Name string `json:"name"`
	// HTTP response codes that indicate a \"not found\" result. Must not overlap with `expectedHttpCodes`. When received, the datasource will return a `not_found` status with no results. Combined with the `mandatory` parameter on a datasource flow entry, this can be used to stop a flow when a required resource is not found.
	NotFoundHttpCodes []int64 `json:"notFoundHttpCodes,omitempty"`
	// The body of the request
	Payload utils.NullableString `json:"payload,omitempty"`
	// For UI purposes in order to format the body correctly
	PayloadType utils.NullableString `json:"payloadType,omitempty"`
	// Name of a Proxy to use while making the request
	Proxy utils.NullableString `json:"proxy,omitempty"`
	// Timeout for the HTTP request.
	Timeout string `json:"timeout" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Type of datasource
	Type string `json:"type"`
	// The URL to request
	Url                  string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _RESTDatasourceResponse RESTDatasourceResponse

// NewRESTDatasourceResponse instantiates a new RESTDatasourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRESTDatasourceResponse(id string, authenticationType string, expectedHttpCodes []int64, method string, name string, timeout string, type_ string, url string) *RESTDatasourceResponse {
	this := RESTDatasourceResponse{}
	this.Id = id
	this.AuthenticationType = authenticationType
	this.ExpectedHttpCodes = expectedHttpCodes
	this.Method = method
	this.Name = name
	this.Timeout = timeout
	this.Type = type_
	this.Url = url
	return &this
}

// NewRESTDatasourceResponseWithDefaults instantiates a new RESTDatasourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRESTDatasourceResponseWithDefaults() *RESTDatasourceResponse {
	this := RESTDatasourceResponse{}
	return &this
}

// GetId returns the Id field value
func (o *RESTDatasourceResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RESTDatasourceResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RESTDatasourceResponse) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDatasourceResponse) GetAttributes() []DataSourceOutput {
	if o == nil {
		var ret []DataSourceOutput
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDatasourceResponse) GetAttributesOk() ([]DataSourceOutput, bool) {
	if o == nil || utils.IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RESTDatasourceResponse) HasAttributes() bool {
	if o != nil && !utils.IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []DataSourceOutput and assigns it to the Attributes field.
func (o *RESTDatasourceResponse) SetAttributes(v []DataSourceOutput) {
	o.Attributes = v
}

// GetAuthenticationType returns the AuthenticationType field value
func (o *RESTDatasourceResponse) GetAuthenticationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value
// and a boolean to check if the value has been set.
func (o *RESTDatasourceResponse) GetAuthenticationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationType, true
}

// SetAuthenticationType sets field value
func (o *RESTDatasourceResponse) SetAuthenticationType(v string) {
	o.AuthenticationType = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *RESTDatasourceResponse) GetCredentials() string {
	if o == nil || utils.IsNil(o.Credentials) {
		var ret string
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTDatasourceResponse) GetCredentialsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *RESTDatasourceResponse) HasCredentials() bool {
	if o != nil && !utils.IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given string and assigns it to the Credentials field.
func (o *RESTDatasourceResponse) SetCredentials(v string) {
	o.Credentials = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RESTDatasourceResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTDatasourceResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RESTDatasourceResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RESTDatasourceResponse) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDatasourceResponse) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDatasourceResponse) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RESTDatasourceResponse) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *RESTDatasourceResponse) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetExpectedHttpCodes returns the ExpectedHttpCodes field value
func (o *RESTDatasourceResponse) GetExpectedHttpCodes() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.ExpectedHttpCodes
}

// GetExpectedHttpCodesOk returns a tuple with the ExpectedHttpCodes field value
// and a boolean to check if the value has been set.
func (o *RESTDatasourceResponse) GetExpectedHttpCodesOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedHttpCodes, true
}

// SetExpectedHttpCodes sets field value
func (o *RESTDatasourceResponse) SetExpectedHttpCodes(v []int64) {
	o.ExpectedHttpCodes = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDatasourceResponse) GetHeaders() []Header {
	if o == nil {
		var ret []Header
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDatasourceResponse) GetHeadersOk() ([]Header, bool) {
	if o == nil || utils.IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *RESTDatasourceResponse) HasHeaders() bool {
	if o != nil && !utils.IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []Header and assigns it to the Headers field.
func (o *RESTDatasourceResponse) SetHeaders(v []Header) {
	o.Headers = v
}

// GetMethod returns the Method field value
func (o *RESTDatasourceResponse) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *RESTDatasourceResponse) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *RESTDatasourceResponse) SetMethod(v string) {
	o.Method = v
}

// GetName returns the Name field value
func (o *RESTDatasourceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RESTDatasourceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RESTDatasourceResponse) SetName(v string) {
	o.Name = v
}

// GetNotFoundHttpCodes returns the NotFoundHttpCodes field value if set, zero value otherwise.
func (o *RESTDatasourceResponse) GetNotFoundHttpCodes() []int64 {
	if o == nil || utils.IsNil(o.NotFoundHttpCodes) {
		var ret []int64
		return ret
	}
	return o.NotFoundHttpCodes
}

// GetNotFoundHttpCodesOk returns a tuple with the NotFoundHttpCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RESTDatasourceResponse) GetNotFoundHttpCodesOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.NotFoundHttpCodes) {
		return nil, false
	}
	return o.NotFoundHttpCodes, true
}

// HasNotFoundHttpCodes returns a boolean if a field has been set.
func (o *RESTDatasourceResponse) HasNotFoundHttpCodes() bool {
	if o != nil && !utils.IsNil(o.NotFoundHttpCodes) {
		return true
	}

	return false
}

// SetNotFoundHttpCodes gets a reference to the given []int64 and assigns it to the NotFoundHttpCodes field.
func (o *RESTDatasourceResponse) SetNotFoundHttpCodes(v []int64) {
	o.NotFoundHttpCodes = v
}

// GetPayload returns the Payload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDatasourceResponse) GetPayload() string {
	if o == nil || utils.IsNil(o.Payload.Get()) {
		var ret string
		return ret
	}
	return *o.Payload.Get()
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDatasourceResponse) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payload.Get(), o.Payload.IsSet()
}

// HasPayload returns a boolean if a field has been set.
func (o *RESTDatasourceResponse) HasPayload() bool {
	if o != nil && o.Payload.IsSet() {
		return true
	}

	return false
}

// SetPayload gets a reference to the given NullableString and assigns it to the Payload field.
func (o *RESTDatasourceResponse) SetPayload(v string) {
	o.Payload.Set(&v)
}

// SetPayloadNil sets the value for Payload to be an explicit nil
func (o *RESTDatasourceResponse) SetPayloadNil() {
	o.Payload.Set(nil)
}

// UnsetPayload ensures that no value is present for Payload, not even an explicit nil
func (o *RESTDatasourceResponse) UnsetPayload() {
	o.Payload.Unset()
}

// GetPayloadType returns the PayloadType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDatasourceResponse) GetPayloadType() string {
	if o == nil || utils.IsNil(o.PayloadType.Get()) {
		var ret string
		return ret
	}
	return *o.PayloadType.Get()
}

// GetPayloadTypeOk returns a tuple with the PayloadType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDatasourceResponse) GetPayloadTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayloadType.Get(), o.PayloadType.IsSet()
}

// HasPayloadType returns a boolean if a field has been set.
func (o *RESTDatasourceResponse) HasPayloadType() bool {
	if o != nil && o.PayloadType.IsSet() {
		return true
	}

	return false
}

// SetPayloadType gets a reference to the given NullableString and assigns it to the PayloadType field.
func (o *RESTDatasourceResponse) SetPayloadType(v string) {
	o.PayloadType.Set(&v)
}

// SetPayloadTypeNil sets the value for PayloadType to be an explicit nil
func (o *RESTDatasourceResponse) SetPayloadTypeNil() {
	o.PayloadType.Set(nil)
}

// UnsetPayloadType ensures that no value is present for PayloadType, not even an explicit nil
func (o *RESTDatasourceResponse) UnsetPayloadType() {
	o.PayloadType.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTDatasourceResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTDatasourceResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *RESTDatasourceResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *RESTDatasourceResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *RESTDatasourceResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *RESTDatasourceResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetTimeout returns the Timeout field value
func (o *RESTDatasourceResponse) GetTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *RESTDatasourceResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *RESTDatasourceResponse) SetTimeout(v string) {
	o.Timeout = v
}

// GetType returns the Type field value
func (o *RESTDatasourceResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RESTDatasourceResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RESTDatasourceResponse) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *RESTDatasourceResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RESTDatasourceResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RESTDatasourceResponse) SetUrl(v string) {
	o.Url = v
}

func (o RESTDatasourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RESTDatasourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["authenticationType"] = o.AuthenticationType
	if !utils.IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["expectedHttpCodes"] = o.ExpectedHttpCodes
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	toSerialize["method"] = o.Method
	toSerialize["name"] = o.Name
	if !utils.IsNil(o.NotFoundHttpCodes) {
		toSerialize["notFoundHttpCodes"] = o.NotFoundHttpCodes
	}
	if o.Payload.IsSet() {
		toSerialize["payload"] = o.Payload.Get()
	}
	if o.PayloadType.IsSet() {
		toSerialize["payloadType"] = o.PayloadType.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	toSerialize["timeout"] = o.Timeout
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RESTDatasourceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"authenticationType",
		"expectedHttpCodes",
		"method",
		"name",
		"timeout",
		"type",
		"url",
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

	varRESTDatasourceResponse := _RESTDatasourceResponse{}

	err = json.Unmarshal(data, &varRESTDatasourceResponse)

	if err != nil {
		return err
	}

	*o = RESTDatasourceResponse(varRESTDatasourceResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "authenticationType")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "expectedHttpCodes")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "method")
		delete(additionalProperties, "name")
		delete(additionalProperties, "notFoundHttpCodes")
		delete(additionalProperties, "payload")
		delete(additionalProperties, "payloadType")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRESTDatasourceResponse struct {
	value *RESTDatasourceResponse
	isSet bool
}

func (v NullableRESTDatasourceResponse) Get() *RESTDatasourceResponse {
	return v.value
}

func (v *NullableRESTDatasourceResponse) Set(val *RESTDatasourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRESTDatasourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRESTDatasourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRESTDatasourceResponse(val *RESTDatasourceResponse) *NullableRESTDatasourceResponse {
	return &NullableRESTDatasourceResponse{value: val, isSet: true}
}

func (v NullableRESTDatasourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRESTDatasourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
