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

// checks if the S3StorageBackendConfigResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &S3StorageBackendConfigResponse{}

// S3StorageBackendConfigResponse struct for S3StorageBackendConfigResponse
type S3StorageBackendConfigResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Name of the bucket to store items into
	Bucket string `json:"bucket"`
	// S3 Checksum mode
	ChecksumMode *string `json:"checksumMode,omitempty"`
	// Name of the `password` [credentials](#tag/security.credentials) containing AWS Secret Keys. If not defined, environment variables will be used
	Credentials *string `json:"credentials,omitempty"`
	// Simple description for this storage
	Description *string `json:"description,omitempty"`
	// Custom endpoint to use for S3
	Endpoint *string `json:"endpoint,omitempty"`
	// If enabled, force S3 path style requests
	ForcePathStyle bool `json:"forcePathStyle"`
	// The name for this storage
	Name           string `json:"name"`
	PartBufferSize string `json:"partBufferSize" validate:"regexp=(?i)^\\\\s*[+-]?(?:\\\\d+(?:\\\\.\\\\d*)?|\\\\.\\\\d+)\\\\s*(?:b|bytes?|k|kb|kib|m|mb|mib|g|gb|gib|t|tb|tib|p|pb|pib|e|eb|eib)?\\\\s*$"`
	// Reference to the proxy to use for connection to the S3
	Proxy *string `json:"proxy,omitempty"`
	// AWS Region for the S3 storage. If not defined, environment variables will be used
	Region *string `json:"region,omitempty"`
	// AWS Role ARN to impersonate
	RoleArn *string `json:"roleArn,omitempty"`
	// Timeout while connecting to the S3
	Timeout utils.NullableString `json:"timeout" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Type of storage
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _S3StorageBackendConfigResponse S3StorageBackendConfigResponse

// NewS3StorageBackendConfigResponse instantiates a new S3StorageBackendConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3StorageBackendConfigResponse(id string, bucket string, forcePathStyle bool, name string, partBufferSize string, timeout utils.NullableString, type_ string) *S3StorageBackendConfigResponse {
	this := S3StorageBackendConfigResponse{}
	this.Id = id
	this.Bucket = bucket
	this.ForcePathStyle = forcePathStyle
	this.Name = name
	this.PartBufferSize = partBufferSize
	this.Timeout = timeout
	this.Type = type_
	var checksumMode string = "when_required"
	this.ChecksumMode = &checksumMode
	return &this
}

// NewS3StorageBackendConfigResponseWithDefaults instantiates a new S3StorageBackendConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3StorageBackendConfigResponseWithDefaults() *S3StorageBackendConfigResponse {
	this := S3StorageBackendConfigResponse{}
	var checksumMode string = "when_required"
	this.ChecksumMode = &checksumMode
	var forcePathStyle bool = false
	this.ForcePathStyle = forcePathStyle
	var partBufferSize string = "9MB"
	this.PartBufferSize = partBufferSize
	return &this
}

// GetId returns the Id field value
func (o *S3StorageBackendConfigResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *S3StorageBackendConfigResponse) SetId(v string) {
	o.Id = v
}

// GetBucket returns the Bucket field value
func (o *S3StorageBackendConfigResponse) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value
func (o *S3StorageBackendConfigResponse) SetBucket(v string) {
	o.Bucket = v
}

// GetChecksumMode returns the ChecksumMode field value if set, zero value otherwise.
func (o *S3StorageBackendConfigResponse) GetChecksumMode() string {
	if o == nil || utils.IsNil(o.ChecksumMode) {
		var ret string
		return ret
	}
	return *o.ChecksumMode
}

// GetChecksumModeOk returns a tuple with the ChecksumMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetChecksumModeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ChecksumMode) {
		return nil, false
	}
	return o.ChecksumMode, true
}

// HasChecksumMode returns a boolean if a field has been set.
func (o *S3StorageBackendConfigResponse) HasChecksumMode() bool {
	if o != nil && !utils.IsNil(o.ChecksumMode) {
		return true
	}

	return false
}

// SetChecksumMode gets a reference to the given string and assigns it to the ChecksumMode field.
func (o *S3StorageBackendConfigResponse) SetChecksumMode(v string) {
	o.ChecksumMode = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *S3StorageBackendConfigResponse) GetCredentials() string {
	if o == nil || utils.IsNil(o.Credentials) {
		var ret string
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetCredentialsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *S3StorageBackendConfigResponse) HasCredentials() bool {
	if o != nil && !utils.IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given string and assigns it to the Credentials field.
func (o *S3StorageBackendConfigResponse) SetCredentials(v string) {
	o.Credentials = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *S3StorageBackendConfigResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *S3StorageBackendConfigResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *S3StorageBackendConfigResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *S3StorageBackendConfigResponse) GetEndpoint() string {
	if o == nil || utils.IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetEndpointOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *S3StorageBackendConfigResponse) HasEndpoint() bool {
	if o != nil && !utils.IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *S3StorageBackendConfigResponse) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetForcePathStyle returns the ForcePathStyle field value
func (o *S3StorageBackendConfigResponse) GetForcePathStyle() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForcePathStyle
}

// GetForcePathStyleOk returns a tuple with the ForcePathStyle field value
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetForcePathStyleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForcePathStyle, true
}

// SetForcePathStyle sets field value
func (o *S3StorageBackendConfigResponse) SetForcePathStyle(v bool) {
	o.ForcePathStyle = v
}

// GetName returns the Name field value
func (o *S3StorageBackendConfigResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *S3StorageBackendConfigResponse) SetName(v string) {
	o.Name = v
}

// GetPartBufferSize returns the PartBufferSize field value
func (o *S3StorageBackendConfigResponse) GetPartBufferSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartBufferSize
}

// GetPartBufferSizeOk returns a tuple with the PartBufferSize field value
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetPartBufferSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartBufferSize, true
}

// SetPartBufferSize sets field value
func (o *S3StorageBackendConfigResponse) SetPartBufferSize(v string) {
	o.PartBufferSize = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *S3StorageBackendConfigResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy) {
		var ret string
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetProxyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Proxy) {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *S3StorageBackendConfigResponse) HasProxy() bool {
	if o != nil && !utils.IsNil(o.Proxy) {
		return true
	}

	return false
}

// SetProxy gets a reference to the given string and assigns it to the Proxy field.
func (o *S3StorageBackendConfigResponse) SetProxy(v string) {
	o.Proxy = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *S3StorageBackendConfigResponse) GetRegion() string {
	if o == nil || utils.IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetRegionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *S3StorageBackendConfigResponse) HasRegion() bool {
	if o != nil && !utils.IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *S3StorageBackendConfigResponse) SetRegion(v string) {
	o.Region = &v
}

// GetRoleArn returns the RoleArn field value if set, zero value otherwise.
func (o *S3StorageBackendConfigResponse) GetRoleArn() string {
	if o == nil || utils.IsNil(o.RoleArn) {
		var ret string
		return ret
	}
	return *o.RoleArn
}

// GetRoleArnOk returns a tuple with the RoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetRoleArnOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RoleArn) {
		return nil, false
	}
	return o.RoleArn, true
}

// HasRoleArn returns a boolean if a field has been set.
func (o *S3StorageBackendConfigResponse) HasRoleArn() bool {
	if o != nil && !utils.IsNil(o.RoleArn) {
		return true
	}

	return false
}

// SetRoleArn gets a reference to the given string and assigns it to the RoleArn field.
func (o *S3StorageBackendConfigResponse) SetRoleArn(v string) {
	o.RoleArn = &v
}

// GetTimeout returns the Timeout field value
// If the value is explicit nil, the zero value for string will be returned
func (o *S3StorageBackendConfigResponse) GetTimeout() string {
	if o == nil || o.Timeout.Get() == nil {
		var ret string
		return ret
	}

	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *S3StorageBackendConfigResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// SetTimeout sets field value
func (o *S3StorageBackendConfigResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// GetType returns the Type field value
func (o *S3StorageBackendConfigResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *S3StorageBackendConfigResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *S3StorageBackendConfigResponse) SetType(v string) {
	o.Type = v
}

func (o S3StorageBackendConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o S3StorageBackendConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["bucket"] = o.Bucket
	if !utils.IsNil(o.ChecksumMode) {
		toSerialize["checksumMode"] = o.ChecksumMode
	}
	if !utils.IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	toSerialize["forcePathStyle"] = o.ForcePathStyle
	toSerialize["name"] = o.Name
	toSerialize["partBufferSize"] = o.PartBufferSize
	if !utils.IsNil(o.Proxy) {
		toSerialize["proxy"] = o.Proxy
	}
	if !utils.IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !utils.IsNil(o.RoleArn) {
		toSerialize["roleArn"] = o.RoleArn
	}
	toSerialize["timeout"] = o.Timeout.Get()
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *S3StorageBackendConfigResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"bucket",
		"forcePathStyle",
		"name",
		"partBufferSize",
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

	varS3StorageBackendConfigResponse := _S3StorageBackendConfigResponse{}

	err = json.Unmarshal(data, &varS3StorageBackendConfigResponse)

	if err != nil {
		return err
	}

	*o = S3StorageBackendConfigResponse(varS3StorageBackendConfigResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "bucket")
		delete(additionalProperties, "checksumMode")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "description")
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "forcePathStyle")
		delete(additionalProperties, "name")
		delete(additionalProperties, "partBufferSize")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "region")
		delete(additionalProperties, "roleArn")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableS3StorageBackendConfigResponse struct {
	value *S3StorageBackendConfigResponse
	isSet bool
}

func (v NullableS3StorageBackendConfigResponse) Get() *S3StorageBackendConfigResponse {
	return v.value
}

func (v *NullableS3StorageBackendConfigResponse) Set(val *S3StorageBackendConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableS3StorageBackendConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableS3StorageBackendConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3StorageBackendConfigResponse(val *S3StorageBackendConfigResponse) *NullableS3StorageBackendConfigResponse {
	return &NullableS3StorageBackendConfigResponse{value: val, isSet: true}
}

func (v NullableS3StorageBackendConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3StorageBackendConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
