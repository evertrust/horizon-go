/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"encoding/json"
	"fmt"
)

// checks if the LicenseInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseInfoResponse{}

// LicenseInfoResponse struct for LicenseInfoResponse
type LicenseInfoResponse struct {
	IsValid bool `json:"isValid"`
	Expiration NullableInt64 `json:"expiration,omitempty"`
	Version string `json:"version"`
	BuildTime int64 `json:"buildTime"`
	Modules []ModuleLicenseInfo `json:"modules"`
	Libraries []LibraryInfo `json:"libraries"`
	AdditionalProperties map[string]interface{}
}

type _LicenseInfoResponse LicenseInfoResponse

// NewLicenseInfoResponse instantiates a new LicenseInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseInfoResponse(isValid bool, version string, buildTime int64, modules []ModuleLicenseInfo, libraries []LibraryInfo) *LicenseInfoResponse {
	this := LicenseInfoResponse{}
	this.IsValid = isValid
	this.Version = version
	this.BuildTime = buildTime
	this.Modules = modules
	this.Libraries = libraries
	return &this
}

// NewLicenseInfoResponseWithDefaults instantiates a new LicenseInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseInfoResponseWithDefaults() *LicenseInfoResponse {
	this := LicenseInfoResponse{}
	return &this
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

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LicenseInfoResponse) GetExpiration() int64 {
	if o == nil || IsNil(o.Expiration.Get()) {
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

// GetModules returns the Modules field value
func (o *LicenseInfoResponse) GetModules() []ModuleLicenseInfo {
	if o == nil {
		var ret []ModuleLicenseInfo
		return ret
	}

	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value
// and a boolean to check if the value has been set.
func (o *LicenseInfoResponse) GetModulesOk() ([]ModuleLicenseInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modules, true
}

// SetModules sets field value
func (o *LicenseInfoResponse) SetModules(v []ModuleLicenseInfo) {
	o.Modules = v
}

// GetLibraries returns the Libraries field value
func (o *LicenseInfoResponse) GetLibraries() []LibraryInfo {
	if o == nil {
		var ret []LibraryInfo
		return ret
	}

	return o.Libraries
}

// GetLibrariesOk returns a tuple with the Libraries field value
// and a boolean to check if the value has been set.
func (o *LicenseInfoResponse) GetLibrariesOk() ([]LibraryInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Libraries, true
}

// SetLibraries sets field value
func (o *LicenseInfoResponse) SetLibraries(v []LibraryInfo) {
	o.Libraries = v
}

func (o LicenseInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isValid"] = o.IsValid
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	toSerialize["version"] = o.Version
	toSerialize["buildTime"] = o.BuildTime
	toSerialize["modules"] = o.Modules
	toSerialize["libraries"] = o.Libraries

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
		"isValid",
		"version",
		"buildTime",
		"modules",
		"libraries",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
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
		delete(additionalProperties, "isValid")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "version")
		delete(additionalProperties, "buildTime")
		delete(additionalProperties, "modules")
		delete(additionalProperties, "libraries")
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


