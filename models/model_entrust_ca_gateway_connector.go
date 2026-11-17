/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in either the `X-API-TOKEN` header or the `Authorization` header using the `Bearer` scheme:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"Authorization: Bearer eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the EntrustCAGatewayConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EntrustCAGatewayConnector{}

// EntrustCAGatewayConnector struct for EntrustCAGatewayConnector
type EntrustCAGatewayConnector struct {
	// Name of the `certificate` [credentials](#tag/security.credentials) to use to authenticate on the PKI
	AuthenticationCredentials string `json:"authenticationCredentials"`
	// Identifier of the CA on the Entrust CA Gateway
	CaId string `json:"caId"`
	// Properties of the enrollment request, mapped to the name of the label whose value is sent
	CustomConnectorDataMapping *map[string]string `json:"customConnectorDataMapping,omitempty"`
	// Base URL of the Entrust CA Gateway, without the `/v1` API version
	EndPoint string `json:"endPoint"`
	Name     string `json:"name"`
	// Identifier of the CA Gateway profile used to enroll certificates
	ProfileId string  `json:"profileId"`
	Proxy     *string `json:"proxy,omitempty"`
	Queue     *string `json:"queue,omitempty"`
	Timeout   *string `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type      string  `json:"type"`
	// Overrides the validity of the issued certificates, as an ISO 8601 interval or duration accepted by the Entrust CA Gateway. The profile must allow the validity to be overridden. Accepted forms: - a duration only, starting when the certificate is issued: `P1Y3M10DT0H0M` - a start date and an expiry date: `2018-07-06T13:00Z/2019-07-06T13:00:00Z` - a start date and a duration: `2018-07-06T13:00Z/P1Y3M0DT0H0M` - a duration and an expiry date, the start date being computed from them: `P0Y3M0DT0H0M/2018-12-31T00:00Z`
	ValidityPeriod       *string `json:"validityPeriod,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntrustCAGatewayConnector EntrustCAGatewayConnector

// NewEntrustCAGatewayConnector instantiates a new EntrustCAGatewayConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntrustCAGatewayConnector(authenticationCredentials string, caId string, endPoint string, name string, profileId string, type_ string) *EntrustCAGatewayConnector {
	this := EntrustCAGatewayConnector{}
	this.AuthenticationCredentials = authenticationCredentials
	this.CaId = caId
	this.EndPoint = endPoint
	this.Name = name
	this.ProfileId = profileId
	this.Type = type_
	return &this
}

// NewEntrustCAGatewayConnectorWithDefaults instantiates a new EntrustCAGatewayConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntrustCAGatewayConnectorWithDefaults() *EntrustCAGatewayConnector {
	this := EntrustCAGatewayConnector{}
	return &this
}

// GetAuthenticationCredentials returns the AuthenticationCredentials field value
func (o *EntrustCAGatewayConnector) GetAuthenticationCredentials() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationCredentials
}

// GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field value
// and a boolean to check if the value has been set.
func (o *EntrustCAGatewayConnector) GetAuthenticationCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationCredentials, true
}

// SetAuthenticationCredentials sets field value
func (o *EntrustCAGatewayConnector) SetAuthenticationCredentials(v string) {
	o.AuthenticationCredentials = v
}

// GetCaId returns the CaId field value
func (o *EntrustCAGatewayConnector) GetCaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaId
}

// GetCaIdOk returns a tuple with the CaId field value
// and a boolean to check if the value has been set.
func (o *EntrustCAGatewayConnector) GetCaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaId, true
}

// SetCaId sets field value
func (o *EntrustCAGatewayConnector) SetCaId(v string) {
	o.CaId = v
}

// GetCustomConnectorDataMapping returns the CustomConnectorDataMapping field value if set, zero value otherwise.
func (o *EntrustCAGatewayConnector) GetCustomConnectorDataMapping() map[string]string {
	if o == nil || utils.IsNil(o.CustomConnectorDataMapping) {
		var ret map[string]string
		return ret
	}
	return *o.CustomConnectorDataMapping
}

// GetCustomConnectorDataMappingOk returns a tuple with the CustomConnectorDataMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrustCAGatewayConnector) GetCustomConnectorDataMappingOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.CustomConnectorDataMapping) {
		return nil, false
	}
	return o.CustomConnectorDataMapping, true
}

// HasCustomConnectorDataMapping returns a boolean if a field has been set.
func (o *EntrustCAGatewayConnector) HasCustomConnectorDataMapping() bool {
	if o != nil && !utils.IsNil(o.CustomConnectorDataMapping) {
		return true
	}

	return false
}

// SetCustomConnectorDataMapping gets a reference to the given map[string]string and assigns it to the CustomConnectorDataMapping field.
func (o *EntrustCAGatewayConnector) SetCustomConnectorDataMapping(v map[string]string) {
	o.CustomConnectorDataMapping = &v
}

// GetEndPoint returns the EndPoint field value
func (o *EntrustCAGatewayConnector) GetEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value
// and a boolean to check if the value has been set.
func (o *EntrustCAGatewayConnector) GetEndPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPoint, true
}

// SetEndPoint sets field value
func (o *EntrustCAGatewayConnector) SetEndPoint(v string) {
	o.EndPoint = v
}

// GetName returns the Name field value
func (o *EntrustCAGatewayConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntrustCAGatewayConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntrustCAGatewayConnector) SetName(v string) {
	o.Name = v
}

// GetProfileId returns the ProfileId field value
func (o *EntrustCAGatewayConnector) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *EntrustCAGatewayConnector) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *EntrustCAGatewayConnector) SetProfileId(v string) {
	o.ProfileId = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *EntrustCAGatewayConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy) {
		var ret string
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrustCAGatewayConnector) GetProxyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *EntrustCAGatewayConnector) HasProxy() bool {
	if o != nil && !utils.IsNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given string and assigns it to the Proxy field.
func (o *EntrustCAGatewayConnector) SetProxy(v string) {
	o.Proxy = &v
}

// GetQueue returns the Queue field value if set, zero value otherwise.
func (o *EntrustCAGatewayConnector) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue) {
		var ret string
		return ret
	}
	return *o.Queue
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrustCAGatewayConnector) GetQueueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Queue) {
		return nil, false
	}
	return o.Queue, true
}

// HasQueue returns a boolean if a field has been set.
func (o *EntrustCAGatewayConnector) HasQueue() bool {
	if o != nil && !utils.IsNil(o.Queue) {
		return true
	}

	return false
}

// SetQueue gets a reference to the given string and assigns it to the Queue field.
func (o *EntrustCAGatewayConnector) SetQueue(v string) {
	o.Queue = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *EntrustCAGatewayConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout) {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrustCAGatewayConnector) GetTimeoutOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *EntrustCAGatewayConnector) HasTimeout() bool {
	if o != nil && !utils.IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *EntrustCAGatewayConnector) SetTimeout(v string) {
	o.Timeout = &v
}

// GetType returns the Type field value
func (o *EntrustCAGatewayConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntrustCAGatewayConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EntrustCAGatewayConnector) SetType(v string) {
	o.Type = v
}

// GetValidityPeriod returns the ValidityPeriod field value if set, zero value otherwise.
func (o *EntrustCAGatewayConnector) GetValidityPeriod() string {
	if o == nil || utils.IsNil(o.ValidityPeriod) {
		var ret string
		return ret
	}
	return *o.ValidityPeriod
}

// GetValidityPeriodOk returns a tuple with the ValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntrustCAGatewayConnector) GetValidityPeriodOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ValidityPeriod) {
		return nil, false
	}
	return o.ValidityPeriod, true
}

// HasValidityPeriod returns a boolean if a field has been set.
func (o *EntrustCAGatewayConnector) HasValidityPeriod() bool {
	if o != nil && !utils.IsNil(o.ValidityPeriod) {
		return true
	}

	return false
}

// SetValidityPeriod gets a reference to the given string and assigns it to the ValidityPeriod field.
func (o *EntrustCAGatewayConnector) SetValidityPeriod(v string) {
	o.ValidityPeriod = &v
}

func (o EntrustCAGatewayConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntrustCAGatewayConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authenticationCredentials"] = o.AuthenticationCredentials
	toSerialize["caId"] = o.CaId
	if !utils.IsNil(o.CustomConnectorDataMapping) {
		toSerialize["customConnectorDataMapping"] = o.CustomConnectorDataMapping
	}
	toSerialize["endPoint"] = o.EndPoint
	toSerialize["name"] = o.Name
	toSerialize["profileId"] = o.ProfileId
	if !utils.IsNil(o.Proxy) {
		toSerialize["proxy"] = o.Proxy
	}
	if !utils.IsNil(o.Queue) {
		toSerialize["queue"] = o.Queue
	}
	if !utils.IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	toSerialize["type"] = o.Type
	if !utils.IsNil(o.ValidityPeriod) {
		toSerialize["validityPeriod"] = o.ValidityPeriod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntrustCAGatewayConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authenticationCredentials",
		"caId",
		"endPoint",
		"name",
		"profileId",
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

	varEntrustCAGatewayConnector := _EntrustCAGatewayConnector{}

	err = json.Unmarshal(data, &varEntrustCAGatewayConnector)

	if err != nil {
		return err
	}

	*o = EntrustCAGatewayConnector(varEntrustCAGatewayConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticationCredentials")
		delete(additionalProperties, "caId")
		delete(additionalProperties, "customConnectorDataMapping")
		delete(additionalProperties, "endPoint")
		delete(additionalProperties, "name")
		delete(additionalProperties, "profileId")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		delete(additionalProperties, "validityPeriod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntrustCAGatewayConnector struct {
	value *EntrustCAGatewayConnector
	isSet bool
}

func (v NullableEntrustCAGatewayConnector) Get() *EntrustCAGatewayConnector {
	return v.value
}

func (v *NullableEntrustCAGatewayConnector) Set(val *EntrustCAGatewayConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableEntrustCAGatewayConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableEntrustCAGatewayConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntrustCAGatewayConnector(val *EntrustCAGatewayConnector) *NullableEntrustCAGatewayConnector {
	return &NullableEntrustCAGatewayConnector{value: val, isSet: true}
}

func (v NullableEntrustCAGatewayConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntrustCAGatewayConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
