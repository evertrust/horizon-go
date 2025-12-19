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

// checks if the RESTResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RESTResponse{}

// RESTResponse struct for RESTResponse
type RESTResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// The authentication type to use while making the REST call. Is linked to `credentials`.
	AuthenticationType string `json:"authenticationType"`
	// Name of the credentials to use for authentication
	Credentials utils.NullableString `json:"credentials,omitempty"`
	// The success HTTP codes for the request. If the return code is not in this list, the notification will be considered failed.
	ExpectedHttpCodes []int64 `json:"expectedHttpCodes"`
	// The headers of the request
	Headers []RESTHeader `json:"headers,omitempty"`
	// The HTTP method to use for the request
	Method string `json:"method"`
	// The body of the request. Can contain dynamic attributes.
	Payload utils.NullableString `json:"payload,omitempty"`
	// For UI purposes in order to format the body correctly
	PayloadType utils.NullableString `json:"payloadType,omitempty"`
	// Name of a Proxy to use while making the request
	Proxy utils.NullableString `json:"proxy,omitempty"`
	// Number of retries when the notification fails (depends on `expectedHttpCodes`)
	Retries interface{} `json:"retries,omitempty"`
	// Timeout for the HTTP request.
	Timeout string `json:"timeout" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type    string `json:"type"`
	// The URL to request
	Url string `json:"url"`
	// Event on which the notification runs. This MUST contain only one value.
	Events []string `json:"events"`
	// License usage at which the notification needs to run (between 0 and 100). Must be defined on `on_license_usage` event and must NOT be defined otherwise.
	LicenseUsagePercent utils.NullableInt64 `json:"licenseUsagePercent,omitempty"`
	// Name of the notification
	Name string `json:"name"`
	// Must be defined on `on_expire` event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed.
	RunOnRenewed utils.NullableBool `json:"runOnRenewed,omitempty"`
	// Time period at which the notification needs to run. Can only be defined on expiration and pending events.
	RunPeriod            utils.NullableString `json:"runPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	AdditionalProperties map[string]interface{}
}

type _RESTResponse RESTResponse

// NewRESTResponse instantiates a new RESTResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRESTResponse(id string, authenticationType string, expectedHttpCodes []int64, method string, timeout string, type_ string, url string, events []string, name string) *RESTResponse {
	this := RESTResponse{}
	this.Events = events
	this.Name = name
	this.Type = type_
	return &this
}

// NewRESTResponseWithDefaults instantiates a new RESTResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRESTResponseWithDefaults() *RESTResponse {
	this := RESTResponse{}
	return &this
}

// GetId returns the Id field value
func (o *RESTResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RESTResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RESTResponse) SetId(v string) {
	o.Id = v
}

// GetAuthenticationType returns the AuthenticationType field value
func (o *RESTResponse) GetAuthenticationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value
// and a boolean to check if the value has been set.
func (o *RESTResponse) GetAuthenticationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationType, true
}

// SetAuthenticationType sets field value
func (o *RESTResponse) SetAuthenticationType(v string) {
	o.AuthenticationType = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTResponse) GetCredentials() string {
	if o == nil || utils.IsNil(o.Credentials.Get()) {
		var ret string
		return ret
	}
	return *o.Credentials.Get()
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTResponse) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials.Get(), o.Credentials.IsSet()
}

// HasCredentials returns a boolean if a field has been set.
func (o *RESTResponse) HasCredentials() bool {
	if o != nil && o.Credentials.IsSet() {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given NullableString and assigns it to the Credentials field.
func (o *RESTResponse) SetCredentials(v string) {
	o.Credentials.Set(&v)
}

// SetCredentialsNil sets the value for Credentials to be an explicit nil
func (o *RESTResponse) SetCredentialsNil() {
	o.Credentials.Set(nil)
}

// UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
func (o *RESTResponse) UnsetCredentials() {
	o.Credentials.Unset()
}

// GetExpectedHttpCodes returns the ExpectedHttpCodes field value
func (o *RESTResponse) GetExpectedHttpCodes() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.ExpectedHttpCodes
}

// GetExpectedHttpCodesOk returns a tuple with the ExpectedHttpCodes field value
// and a boolean to check if the value has been set.
func (o *RESTResponse) GetExpectedHttpCodesOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpectedHttpCodes, true
}

// SetExpectedHttpCodes sets field value
func (o *RESTResponse) SetExpectedHttpCodes(v []int64) {
	o.ExpectedHttpCodes = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTResponse) GetHeaders() []RESTHeader {
	if o == nil {
		var ret []RESTHeader
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTResponse) GetHeadersOk() ([]RESTHeader, bool) {
	if o == nil || utils.IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *RESTResponse) HasHeaders() bool {
	if o != nil && !utils.IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []RESTHeader and assigns it to the Headers field.
func (o *RESTResponse) SetHeaders(v []RESTHeader) {
	o.Headers = v
}

// GetMethod returns the Method field value
func (o *RESTResponse) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *RESTResponse) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *RESTResponse) SetMethod(v string) {
	o.Method = v
}

// GetPayload returns the Payload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTResponse) GetPayload() string {
	if o == nil || utils.IsNil(o.Payload.Get()) {
		var ret string
		return ret
	}
	return *o.Payload.Get()
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTResponse) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payload.Get(), o.Payload.IsSet()
}

// HasPayload returns a boolean if a field has been set.
func (o *RESTResponse) HasPayload() bool {
	if o != nil && o.Payload.IsSet() {
		return true
	}

	return false
}

// SetPayload gets a reference to the given NullableString and assigns it to the Payload field.
func (o *RESTResponse) SetPayload(v string) {
	o.Payload.Set(&v)
}

// SetPayloadNil sets the value for Payload to be an explicit nil
func (o *RESTResponse) SetPayloadNil() {
	o.Payload.Set(nil)
}

// UnsetPayload ensures that no value is present for Payload, not even an explicit nil
func (o *RESTResponse) UnsetPayload() {
	o.Payload.Unset()
}

// GetPayloadType returns the PayloadType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTResponse) GetPayloadType() string {
	if o == nil || utils.IsNil(o.PayloadType.Get()) {
		var ret string
		return ret
	}
	return *o.PayloadType.Get()
}

// GetPayloadTypeOk returns a tuple with the PayloadType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTResponse) GetPayloadTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayloadType.Get(), o.PayloadType.IsSet()
}

// HasPayloadType returns a boolean if a field has been set.
func (o *RESTResponse) HasPayloadType() bool {
	if o != nil && o.PayloadType.IsSet() {
		return true
	}

	return false
}

// SetPayloadType gets a reference to the given NullableString and assigns it to the PayloadType field.
func (o *RESTResponse) SetPayloadType(v string) {
	o.PayloadType.Set(&v)
}

// SetPayloadTypeNil sets the value for PayloadType to be an explicit nil
func (o *RESTResponse) SetPayloadTypeNil() {
	o.PayloadType.Set(nil)
}

// UnsetPayloadType ensures that no value is present for PayloadType, not even an explicit nil
func (o *RESTResponse) UnsetPayloadType() {
	o.PayloadType.Unset()
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *RESTResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *RESTResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *RESTResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *RESTResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetRetries returns the Retries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTResponse) GetRetries() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTResponse) GetRetriesOk() (*interface{}, bool) {
	if o == nil || utils.IsNil(o.Retries) {
		return nil, false
	}
	return &o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *RESTResponse) HasRetries() bool {
	if o != nil && !utils.IsNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given interface{} and assigns it to the Retries field.
func (o *RESTResponse) SetRetries(v interface{}) {
	o.Retries = v
}

// GetTimeout returns the Timeout field value
func (o *RESTResponse) GetTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *RESTResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *RESTResponse) SetTimeout(v string) {
	o.Timeout = v
}

// GetType returns the Type field value
func (o *RESTResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RESTResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RESTResponse) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *RESTResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RESTResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RESTResponse) SetUrl(v string) {
	o.Url = v
}

// GetEvents returns the Events field value
func (o *RESTResponse) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *RESTResponse) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *RESTResponse) SetEvents(v []string) {
	o.Events = v
}

// GetLicenseUsagePercent returns the LicenseUsagePercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTResponse) GetLicenseUsagePercent() int64 {
	if o == nil || utils.IsNil(o.LicenseUsagePercent.Get()) {
		var ret int64
		return ret
	}
	return *o.LicenseUsagePercent.Get()
}

// GetLicenseUsagePercentOk returns a tuple with the LicenseUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTResponse) GetLicenseUsagePercentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicenseUsagePercent.Get(), o.LicenseUsagePercent.IsSet()
}

// HasLicenseUsagePercent returns a boolean if a field has been set.
func (o *RESTResponse) HasLicenseUsagePercent() bool {
	if o != nil && o.LicenseUsagePercent.IsSet() {
		return true
	}

	return false
}

// SetLicenseUsagePercent gets a reference to the given NullableInt64 and assigns it to the LicenseUsagePercent field.
func (o *RESTResponse) SetLicenseUsagePercent(v int64) {
	o.LicenseUsagePercent.Set(&v)
}

// SetLicenseUsagePercentNil sets the value for LicenseUsagePercent to be an explicit nil
func (o *RESTResponse) SetLicenseUsagePercentNil() {
	o.LicenseUsagePercent.Set(nil)
}

// UnsetLicenseUsagePercent ensures that no value is present for LicenseUsagePercent, not even an explicit nil
func (o *RESTResponse) UnsetLicenseUsagePercent() {
	o.LicenseUsagePercent.Unset()
}

// GetName returns the Name field value
func (o *RESTResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RESTResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RESTResponse) SetName(v string) {
	o.Name = v
}

// GetRunOnRenewed returns the RunOnRenewed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTResponse) GetRunOnRenewed() bool {
	if o == nil || utils.IsNil(o.RunOnRenewed.Get()) {
		var ret bool
		return ret
	}
	return *o.RunOnRenewed.Get()
}

// GetRunOnRenewedOk returns a tuple with the RunOnRenewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTResponse) GetRunOnRenewedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunOnRenewed.Get(), o.RunOnRenewed.IsSet()
}

// HasRunOnRenewed returns a boolean if a field has been set.
func (o *RESTResponse) HasRunOnRenewed() bool {
	if o != nil && o.RunOnRenewed.IsSet() {
		return true
	}

	return false
}

// SetRunOnRenewed gets a reference to the given NullableBool and assigns it to the RunOnRenewed field.
func (o *RESTResponse) SetRunOnRenewed(v bool) {
	o.RunOnRenewed.Set(&v)
}

// SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil
func (o *RESTResponse) SetRunOnRenewedNil() {
	o.RunOnRenewed.Set(nil)
}

// UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil
func (o *RESTResponse) UnsetRunOnRenewed() {
	o.RunOnRenewed.Unset()
}

// GetRunPeriod returns the RunPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RESTResponse) GetRunPeriod() string {
	if o == nil || utils.IsNil(o.RunPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RunPeriod.Get()
}

// GetRunPeriodOk returns a tuple with the RunPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RESTResponse) GetRunPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunPeriod.Get(), o.RunPeriod.IsSet()
}

// HasRunPeriod returns a boolean if a field has been set.
func (o *RESTResponse) HasRunPeriod() bool {
	if o != nil && o.RunPeriod.IsSet() {
		return true
	}

	return false
}

// SetRunPeriod gets a reference to the given NullableString and assigns it to the RunPeriod field.
func (o *RESTResponse) SetRunPeriod(v string) {
	o.RunPeriod.Set(&v)
}

// SetRunPeriodNil sets the value for RunPeriod to be an explicit nil
func (o *RESTResponse) SetRunPeriodNil() {
	o.RunPeriod.Set(nil)
}

// UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil
func (o *RESTResponse) UnsetRunPeriod() {
	o.RunPeriod.Unset()
}

func (o RESTResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RESTResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["authenticationType"] = o.AuthenticationType
	if o.Credentials.IsSet() {
		toSerialize["credentials"] = o.Credentials.Get()
	}
	toSerialize["expectedHttpCodes"] = o.ExpectedHttpCodes
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	toSerialize["method"] = o.Method
	if o.Payload.IsSet() {
		toSerialize["payload"] = o.Payload.Get()
	}
	if o.PayloadType.IsSet() {
		toSerialize["payloadType"] = o.PayloadType.Get()
	}
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}
	toSerialize["timeout"] = o.Timeout
	toSerialize["type"] = o.Type
	toSerialize["url"] = o.Url
	toSerialize["events"] = o.Events
	if o.LicenseUsagePercent.IsSet() {
		toSerialize["licenseUsagePercent"] = o.LicenseUsagePercent.Get()
	}
	toSerialize["name"] = o.Name
	if o.RunOnRenewed.IsSet() {
		toSerialize["runOnRenewed"] = o.RunOnRenewed.Get()
	}
	if o.RunPeriod.IsSet() {
		toSerialize["runPeriod"] = o.RunPeriod.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RESTResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"authenticationType",
		"expectedHttpCodes",
		"method",
		"timeout",
		"type",
		"url",
		"events",
		"name",
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

	varRESTResponse := _RESTResponse{}

	err = json.Unmarshal(data, &varRESTResponse)

	if err != nil {
		return err
	}

	*o = RESTResponse(varRESTResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "authenticationType")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "expectedHttpCodes")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "method")
		delete(additionalProperties, "payload")
		delete(additionalProperties, "payloadType")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "retries")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		delete(additionalProperties, "events")
		delete(additionalProperties, "licenseUsagePercent")
		delete(additionalProperties, "name")
		delete(additionalProperties, "runOnRenewed")
		delete(additionalProperties, "runPeriod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRESTResponse struct {
	value *RESTResponse
	isSet bool
}

func (v NullableRESTResponse) Get() *RESTResponse {
	return v.value
}

func (v *NullableRESTResponse) Set(val *RESTResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRESTResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRESTResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRESTResponse(val *RESTResponse) *NullableRESTResponse {
	return &NullableRESTResponse{value: val, isSet: true}
}

func (v NullableRESTResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRESTResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
