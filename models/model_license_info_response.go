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

// checks if the LicenseInfoResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LicenseInfoResponse{}

// LicenseInfoResponse struct for LicenseInfoResponse
type LicenseInfoResponse struct {
	BuildTime int64 `json:"buildTime"`
	// Current certificate count
	Count int64 `json:"count"`
	// Current DCV count
	DcvCount int64 `json:"dcvCount"`
	// DCV license limit
	DcvLimit   utils.NullableInt64   `json:"dcvLimit,omitempty"`
	Expiration utils.NullableInt64   `json:"expiration,omitempty"`
	IsValid    bool                  `json:"isValid"`
	Libraries  []LibraryInfoResponse `json:"libraries"`
	// Certificate license limit
	Limit                utils.NullableInt64         `json:"limit,omitempty"`
	Modules              []ModuleLicenseInfoResponse `json:"modules"`
	ReleaseChannel       *string                     `json:"releaseChannel,omitempty"`
	Version              string                      `json:"version"`
	AdditionalProperties map[string]interface{}
}

type _LicenseInfoResponse LicenseInfoResponse

// NewLicenseInfoResponse instantiates a new LicenseInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseInfoResponse(buildTime int64, count int64, dcvCount int64, isValid bool, libraries []LibraryInfoResponse, modules []ModuleLicenseInfoResponse, version string) *LicenseInfoResponse {
	this := LicenseInfoResponse{}
	this.BuildTime = buildTime
	this.Count = count
	this.DcvCount = dcvCount
	this.IsValid = isValid
	this.Libraries = libraries
	this.Modules = modules
	this.Version = version
	return &this
}

// NewLicenseInfoResponseWithDefaults instantiates a new LicenseInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseInfoResponseWithDefaults() *LicenseInfoResponse {
	this := LicenseInfoResponse{}
	return &this
}

// GetBuildTime returns the BuildTime field value
func (o *LicenseInfoResponse) GetBuildTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BuildTime
}

// GetBuildTimeOk returns a tuple with the BuildTime field value
// and a boolean to check if the value has been set.
func (o *LicenseInfoResponse) GetBuildTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildTime, true
}

// SetBuildTime sets field value
func (o *LicenseInfoResponse) SetBuildTime(v int64) {
	o.BuildTime = v
}

// GetCount returns the Count field value
func (o *LicenseInfoResponse) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *LicenseInfoResponse) GetCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *LicenseInfoResponse) SetCount(v int64) {
	o.Count = v
}

// GetDcvCount returns the DcvCount field value
func (o *LicenseInfoResponse) GetDcvCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DcvCount
}

// GetDcvCountOk returns a tuple with the DcvCount field value
// and a boolean to check if the value has been set.
func (o *LicenseInfoResponse) GetDcvCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcvCount, true
}

// SetDcvCount sets field value
func (o *LicenseInfoResponse) SetDcvCount(v int64) {
	o.DcvCount = v
}

// GetDcvLimit returns the DcvLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseInfoResponse) GetDcvLimit() int64 {
	if o == nil || utils.IsNil(o.DcvLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.DcvLimit.Get()
}

// GetDcvLimitOk returns a tuple with the DcvLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseInfoResponse) GetDcvLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DcvLimit.Get(), o.DcvLimit.IsSet()
}

// HasDcvLimit returns a boolean if a field has been set.
func (o *LicenseInfoResponse) HasDcvLimit() bool {
	if o != nil && o.DcvLimit.IsSet() {
		return true
	}

	return false
}

// SetDcvLimit gets a reference to the given NullableInt64 and assigns it to the DcvLimit field.
func (o *LicenseInfoResponse) SetDcvLimit(v int64) {
	o.DcvLimit.Set(&v)
}

// SetDcvLimitNil sets the value for DcvLimit to be an explicit nil
func (o *LicenseInfoResponse) SetDcvLimitNil() {
	o.DcvLimit.Set(nil)
}

// UnsetDcvLimit ensures that no value is present for DcvLimit, not even an explicit nil
func (o *LicenseInfoResponse) UnsetDcvLimit() {
	o.DcvLimit.Unset()
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseInfoResponse) GetExpiration() int64 {
	if o == nil || utils.IsNil(o.Expiration.Get()) {
		var ret int64
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseInfoResponse) GetExpirationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *LicenseInfoResponse) HasExpiration() bool {
	if o != nil && o.Expiration.IsSet() {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given NullableInt64 and assigns it to the Expiration field.
func (o *LicenseInfoResponse) SetExpiration(v int64) {
	o.Expiration.Set(&v)
}

// SetExpirationNil sets the value for Expiration to be an explicit nil
func (o *LicenseInfoResponse) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
func (o *LicenseInfoResponse) UnsetExpiration() {
	o.Expiration.Unset()
}

// GetIsValid returns the IsValid field value
func (o *LicenseInfoResponse) GetIsValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value
// and a boolean to check if the value has been set.
func (o *LicenseInfoResponse) GetIsValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsValid, true
}

// SetIsValid sets field value
func (o *LicenseInfoResponse) SetIsValid(v bool) {
	o.IsValid = v
}

// GetLibraries returns the Libraries field value
func (o *LicenseInfoResponse) GetLibraries() []LibraryInfoResponse {
	if o == nil {
		var ret []LibraryInfoResponse
		return ret
	}

	return o.Libraries
}

// GetLibrariesOk returns a tuple with the Libraries field value
// and a boolean to check if the value has been set.
func (o *LicenseInfoResponse) GetLibrariesOk() ([]LibraryInfoResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Libraries, true
}

// SetLibraries sets field value
func (o *LicenseInfoResponse) SetLibraries(v []LibraryInfoResponse) {
	o.Libraries = v
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseInfoResponse) GetLimit() int64 {
	if o == nil || utils.IsNil(o.Limit.Get()) {
		var ret int64
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseInfoResponse) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *LicenseInfoResponse) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt64 and assigns it to the Limit field.
func (o *LicenseInfoResponse) SetLimit(v int64) {
	o.Limit.Set(&v)
}

// SetLimitNil sets the value for Limit to be an explicit nil
func (o *LicenseInfoResponse) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *LicenseInfoResponse) UnsetLimit() {
	o.Limit.Unset()
}

// GetModules returns the Modules field value
func (o *LicenseInfoResponse) GetModules() []ModuleLicenseInfoResponse {
	if o == nil {
		var ret []ModuleLicenseInfoResponse
		return ret
	}

	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value
// and a boolean to check if the value has been set.
func (o *LicenseInfoResponse) GetModulesOk() ([]ModuleLicenseInfoResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modules, true
}

// SetModules sets field value
func (o *LicenseInfoResponse) SetModules(v []ModuleLicenseInfoResponse) {
	o.Modules = v
}

// GetReleaseChannel returns the ReleaseChannel field value if set, zero value otherwise.
func (o *LicenseInfoResponse) GetReleaseChannel() string {
	if o == nil || utils.IsNil(o.ReleaseChannel) {
		var ret string
		return ret
	}
	return *o.ReleaseChannel
}

// GetReleaseChannelOk returns a tuple with the ReleaseChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseInfoResponse) GetReleaseChannelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReleaseChannel) {
		return nil, false
	}
	return o.ReleaseChannel, true
}

// HasReleaseChannel returns a boolean if a field has been set.
func (o *LicenseInfoResponse) HasReleaseChannel() bool {
	if o != nil && !utils.IsNil(o.ReleaseChannel) {
		return true
	}

	return false
}

// SetReleaseChannel gets a reference to the given string and assigns it to the ReleaseChannel field.
func (o *LicenseInfoResponse) SetReleaseChannel(v string) {
	o.ReleaseChannel = &v
}

// GetVersion returns the Version field value
func (o *LicenseInfoResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *LicenseInfoResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *LicenseInfoResponse) SetVersion(v string) {
	o.Version = v
}

func (o LicenseInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["buildTime"] = o.BuildTime
	toSerialize["count"] = o.Count
	toSerialize["dcvCount"] = o.DcvCount
	if o.DcvLimit.IsSet() {
		toSerialize["dcvLimit"] = o.DcvLimit.Get()
	}
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	toSerialize["isValid"] = o.IsValid
	toSerialize["libraries"] = o.Libraries
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	toSerialize["modules"] = o.Modules
	if !utils.IsNil(o.ReleaseChannel) {
		toSerialize["releaseChannel"] = o.ReleaseChannel
	}
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LicenseInfoResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"buildTime",
		"count",
		"dcvCount",
		"isValid",
		"libraries",
		"modules",
		"version",
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

	varLicenseInfoResponse := _LicenseInfoResponse{}

	err = json.Unmarshal(data, &varLicenseInfoResponse)

	if err != nil {
		return err
	}

	*o = LicenseInfoResponse(varLicenseInfoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "buildTime")
		delete(additionalProperties, "count")
		delete(additionalProperties, "dcvCount")
		delete(additionalProperties, "dcvLimit")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "isValid")
		delete(additionalProperties, "libraries")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "modules")
		delete(additionalProperties, "releaseChannel")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLicenseInfoResponse struct {
	value *LicenseInfoResponse
	isSet bool
}

func (v NullableLicenseInfoResponse) Get() *LicenseInfoResponse {
	return v.value
}

func (v *NullableLicenseInfoResponse) Set(val *LicenseInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseInfoResponse(val *LicenseInfoResponse) *NullableLicenseInfoResponse {
	return &NullableLicenseInfoResponse{value: val, isSet: true}
}

func (v NullableLicenseInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
