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

// checks if the GCPConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GCPConnectorResponse{}

// GCPConnectorResponse struct for GCPConnectorResponse
type GCPConnectorResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Identifier of the CA pool to issue from. The pool auto-selects an enabled certificate authority.
	CaPool string `json:"caPool"`
	// Validity applied to every certificate issued through this connector.
	CertificateLifetime string `json:"certificateLifetime" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Certificate template governing issuance policy. Accepts the template short name or its full resource path.
	CertificateTemplate *string `json:"certificateTemplate,omitempty"`
	// Name of the `raw` [credentials](#tag/security.credentials) holding the Google service account key (JSON). If not defined, Application Default Credentials are used (environment variable or workload identity).
	Credentials *string `json:"credentials,omitempty"`
	// Overrides the default Certificate Authority Service address and port (`privateca.googleapis.com:443`). If not set, the default service URL is used.
	Endpoint      *string                    `json:"endpoint,omitempty"`
	Impersonation *GCPConnectorImpersonation `json:"impersonation,omitempty"`
	// Google Cloud location (region) of the CA pool. Also used by the connector to derive the Certificate Authority Service endpoint.
	Location string `json:"location"`
	Name     string `json:"name"`
	// Identifier of the Google Cloud project hosting the CA pool. May differ from the project of the service account used to authenticate.
	ProjectId string `json:"projectId"`
	// Name of the proxy to use to connect to the GCP Api
	Proxy                *string                    `json:"proxy,omitempty"`
	Queue                utils.NullableString       `json:"queue,omitempty"`
	Status               NullablePKIConnectorStatus `json:"status,omitempty"`
	Timeout              *string                    `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string                     `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _GCPConnectorResponse GCPConnectorResponse

// NewGCPConnectorResponse instantiates a new GCPConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGCPConnectorResponse(id string, caPool string, certificateLifetime string, location string, name string, projectId string, type_ string) *GCPConnectorResponse {
	this := GCPConnectorResponse{}
	this.Id = id
	this.CaPool = caPool
	this.CertificateLifetime = certificateLifetime
	this.Location = location
	this.Name = name
	this.ProjectId = projectId
	this.Type = type_
	return &this
}

// NewGCPConnectorResponseWithDefaults instantiates a new GCPConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGCPConnectorResponseWithDefaults() *GCPConnectorResponse {
	this := GCPConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GCPConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GCPConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetCaPool returns the CaPool field value
func (o *GCPConnectorResponse) GetCaPool() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaPool
}

// GetCaPoolOk returns a tuple with the CaPool field value
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetCaPoolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaPool, true
}

// SetCaPool sets field value
func (o *GCPConnectorResponse) SetCaPool(v string) {
	o.CaPool = v
}

// GetCertificateLifetime returns the CertificateLifetime field value
func (o *GCPConnectorResponse) GetCertificateLifetime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CertificateLifetime
}

// GetCertificateLifetimeOk returns a tuple with the CertificateLifetime field value
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetCertificateLifetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateLifetime, true
}

// SetCertificateLifetime sets field value
func (o *GCPConnectorResponse) SetCertificateLifetime(v string) {
	o.CertificateLifetime = v
}

// GetCertificateTemplate returns the CertificateTemplate field value if set, zero value otherwise.
func (o *GCPConnectorResponse) GetCertificateTemplate() string {
	if o == nil || utils.IsNil(o.CertificateTemplate) {
		var ret string
		return ret
	}
	return *o.CertificateTemplate
}

// GetCertificateTemplateOk returns a tuple with the CertificateTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetCertificateTemplateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CertificateTemplate) {
		return nil, false
	}
	return o.CertificateTemplate, true
}

// HasCertificateTemplate returns a boolean if a field has been set.
func (o *GCPConnectorResponse) HasCertificateTemplate() bool {
	if o != nil && !utils.IsNil(o.CertificateTemplate) {
		return true
	}

	return false
}

// SetCertificateTemplate gets a reference to the given string and assigns it to the CertificateTemplate field.
func (o *GCPConnectorResponse) SetCertificateTemplate(v string) {
	o.CertificateTemplate = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *GCPConnectorResponse) GetCredentials() string {
	if o == nil || utils.IsNil(o.Credentials) {
		var ret string
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetCredentialsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *GCPConnectorResponse) HasCredentials() bool {
	if o != nil && !utils.IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given string and assigns it to the Credentials field.
func (o *GCPConnectorResponse) SetCredentials(v string) {
	o.Credentials = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *GCPConnectorResponse) GetEndpoint() string {
	if o == nil || utils.IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetEndpointOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *GCPConnectorResponse) HasEndpoint() bool {
	if o != nil && !utils.IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *GCPConnectorResponse) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetImpersonation returns the Impersonation field value if set, zero value otherwise.
func (o *GCPConnectorResponse) GetImpersonation() GCPConnectorImpersonation {
	if o == nil || utils.IsNil(o.Impersonation) {
		var ret GCPConnectorImpersonation
		return ret
	}
	return *o.Impersonation
}

// GetImpersonationOk returns a tuple with the Impersonation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetImpersonationOk() (*GCPConnectorImpersonation, bool) {
	if o == nil || utils.IsNil(o.Impersonation) {
		return nil, false
	}
	return o.Impersonation, true
}

// HasImpersonation returns a boolean if a field has been set.
func (o *GCPConnectorResponse) HasImpersonation() bool {
	if o != nil && !utils.IsNil(o.Impersonation) {
		return true
	}

	return false
}

// SetImpersonation gets a reference to the given GCPConnectorImpersonation and assigns it to the Impersonation field.
func (o *GCPConnectorResponse) SetImpersonation(v GCPConnectorImpersonation) {
	o.Impersonation = &v
}

// GetLocation returns the Location field value
func (o *GCPConnectorResponse) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *GCPConnectorResponse) SetLocation(v string) {
	o.Location = v
}

// GetName returns the Name field value
func (o *GCPConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GCPConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *GCPConnectorResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GCPConnectorResponse) SetProjectId(v string) {
	o.ProjectId = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *GCPConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy) {
		var ret string
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *GCPConnectorResponse) HasProxy() bool {
	if o != nil && !utils.IsNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given string and assigns it to the Proxy field.
func (o *GCPConnectorResponse) SetProxy(v string) {
	o.Proxy = &v
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GCPConnectorResponse) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GCPConnectorResponse) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *GCPConnectorResponse) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *GCPConnectorResponse) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *GCPConnectorResponse) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *GCPConnectorResponse) UnsetQueue() {
	o.Queue.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GCPConnectorResponse) GetStatus() PKIConnectorStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret PKIConnectorStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GCPConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *GCPConnectorResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePKIConnectorStatus and assigns it to the Status field.
func (o *GCPConnectorResponse) SetStatus(v PKIConnectorStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *GCPConnectorResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *GCPConnectorResponse) UnsetStatus() {
	o.Status.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *GCPConnectorResponse) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout) {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *GCPConnectorResponse) HasTimeout() bool {
	if o != nil && !utils.IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *GCPConnectorResponse) SetTimeout(v string) {
	o.Timeout = &v
}

// GetType returns the Type field value
func (o *GCPConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GCPConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GCPConnectorResponse) SetType(v string) {
	o.Type = v
}

func (o GCPConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GCPConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["caPool"] = o.CaPool
	toSerialize["certificateLifetime"] = o.CertificateLifetime
	if !utils.IsNil(o.CertificateTemplate) {
		toSerialize["certificateTemplate"] = o.CertificateTemplate
	}
	if !utils.IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !utils.IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !utils.IsNil(o.Impersonation) {
		toSerialize["impersonation"] = o.Impersonation
	}
	toSerialize["location"] = o.Location
	toSerialize["name"] = o.Name
	toSerialize["projectId"] = o.ProjectId
	if !utils.IsNil(o.Proxy) {
		toSerialize["proxy"] = o.Proxy
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if !utils.IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GCPConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"caPool",
		"certificateLifetime",
		"location",
		"name",
		"projectId",
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

	varGCPConnectorResponse := _GCPConnectorResponse{}

	err = json.Unmarshal(data, &varGCPConnectorResponse)

	if err != nil {
		return err
	}

	*o = GCPConnectorResponse(varGCPConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "caPool")
		delete(additionalProperties, "certificateLifetime")
		delete(additionalProperties, "certificateTemplate")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "impersonation")
		delete(additionalProperties, "location")
		delete(additionalProperties, "name")
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGCPConnectorResponse struct {
	value *GCPConnectorResponse
	isSet bool
}

func (v NullableGCPConnectorResponse) Get() *GCPConnectorResponse {
	return v.value
}

func (v *NullableGCPConnectorResponse) Set(val *GCPConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGCPConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGCPConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGCPConnectorResponse(val *GCPConnectorResponse) *NullableGCPConnectorResponse {
	return &NullableGCPConnectorResponse{value: val, isSet: true}
}

func (v NullableGCPConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGCPConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
