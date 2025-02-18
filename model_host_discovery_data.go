/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"encoding/json"
)

// checks if the HostDiscoveryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostDiscoveryData{}

// HostDiscoveryData struct for HostDiscoveryData
type HostDiscoveryData struct {
	// The certificate's host ip
	Ip NullableString `json:"ip,omitempty"`
	// Information on the type of discovery that discovered this certificate
	Sources []string `json:"sources,omitempty"`
	// The certificate's host hostnames (netscan only)
	Hostnames []string `json:"hostnames,omitempty"`
	// The certificate's host operating system (localscan only)
	OperatingSystems []string `json:"operatingSystems,omitempty"`
	// The path to the certificate on the host machine (localscan only)
	Paths []string `json:"paths,omitempty"`
	// The path of the configuration files that were used to find the certificates
	Usages []string `json:"usages,omitempty"`
	// The ports on which the certificate is exposed for https connexion
	TlsPorts []TlsPort `json:"tlsPorts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HostDiscoveryData HostDiscoveryData

// NewHostDiscoveryData instantiates a new HostDiscoveryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostDiscoveryData() *HostDiscoveryData {
	this := HostDiscoveryData{}
	return &this
}

// NewHostDiscoveryDataWithDefaults instantiates a new HostDiscoveryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostDiscoveryDataWithDefaults() *HostDiscoveryData {
	this := HostDiscoveryData{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostDiscoveryData) GetIp() string {
	if o == nil || IsNil(o.Ip.Get()) {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostDiscoveryData) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *HostDiscoveryData) HasIp() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *HostDiscoveryData) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *HostDiscoveryData) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *HostDiscoveryData) UnsetIp() {
	o.Ip.Unset()
}

// GetSources returns the Sources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostDiscoveryData) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostDiscoveryData) GetSourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *HostDiscoveryData) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *HostDiscoveryData) SetSources(v []string) {
	o.Sources = v
}

// GetHostnames returns the Hostnames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostDiscoveryData) GetHostnames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Hostnames
}

// GetHostnamesOk returns a tuple with the Hostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostDiscoveryData) GetHostnamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Hostnames) {
		return nil, false
	}
	return o.Hostnames, true
}

// HasHostnames returns a boolean if a field has been set.
func (o *HostDiscoveryData) HasHostnames() bool {
	if o != nil && !IsNil(o.Hostnames) {
		return true
	}

	return false
}

// SetHostnames gets a reference to the given []string and assigns it to the Hostnames field.
func (o *HostDiscoveryData) SetHostnames(v []string) {
	o.Hostnames = v
}

// GetOperatingSystems returns the OperatingSystems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostDiscoveryData) GetOperatingSystems() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperatingSystems
}

// GetOperatingSystemsOk returns a tuple with the OperatingSystems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostDiscoveryData) GetOperatingSystemsOk() ([]string, bool) {
	if o == nil || IsNil(o.OperatingSystems) {
		return nil, false
	}
	return o.OperatingSystems, true
}

// HasOperatingSystems returns a boolean if a field has been set.
func (o *HostDiscoveryData) HasOperatingSystems() bool {
	if o != nil && !IsNil(o.OperatingSystems) {
		return true
	}

	return false
}

// SetOperatingSystems gets a reference to the given []string and assigns it to the OperatingSystems field.
func (o *HostDiscoveryData) SetOperatingSystems(v []string) {
	o.OperatingSystems = v
}

// GetPaths returns the Paths field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostDiscoveryData) GetPaths() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostDiscoveryData) GetPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.Paths) {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *HostDiscoveryData) HasPaths() bool {
	if o != nil && !IsNil(o.Paths) {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *HostDiscoveryData) SetPaths(v []string) {
	o.Paths = v
}

// GetUsages returns the Usages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostDiscoveryData) GetUsages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostDiscoveryData) GetUsagesOk() ([]string, bool) {
	if o == nil || IsNil(o.Usages) {
		return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *HostDiscoveryData) HasUsages() bool {
	if o != nil && !IsNil(o.Usages) {
		return true
	}

	return false
}

// SetUsages gets a reference to the given []string and assigns it to the Usages field.
func (o *HostDiscoveryData) SetUsages(v []string) {
	o.Usages = v
}

// GetTlsPorts returns the TlsPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HostDiscoveryData) GetTlsPorts() []TlsPort {
	if o == nil {
		var ret []TlsPort
		return ret
	}
	return o.TlsPorts
}

// GetTlsPortsOk returns a tuple with the TlsPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HostDiscoveryData) GetTlsPortsOk() ([]TlsPort, bool) {
	if o == nil || IsNil(o.TlsPorts) {
		return nil, false
	}
	return o.TlsPorts, true
}

// HasTlsPorts returns a boolean if a field has been set.
func (o *HostDiscoveryData) HasTlsPorts() bool {
	if o != nil && !IsNil(o.TlsPorts) {
		return true
	}

	return false
}

// SetTlsPorts gets a reference to the given []TlsPort and assigns it to the TlsPorts field.
func (o *HostDiscoveryData) SetTlsPorts(v []TlsPort) {
	o.TlsPorts = v
}

func (o HostDiscoveryData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostDiscoveryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	if o.Hostnames != nil {
		toSerialize["hostnames"] = o.Hostnames
	}
	if o.OperatingSystems != nil {
		toSerialize["operatingSystems"] = o.OperatingSystems
	}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	if o.Usages != nil {
		toSerialize["usages"] = o.Usages
	}
	if o.TlsPorts != nil {
		toSerialize["tlsPorts"] = o.TlsPorts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HostDiscoveryData) UnmarshalJSON(data []byte) (err error) {
	varHostDiscoveryData := _HostDiscoveryData{}

	err = json.Unmarshal(data, &varHostDiscoveryData)

	if err != nil {
		return err
	}

	*o = HostDiscoveryData(varHostDiscoveryData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ip")
		delete(additionalProperties, "sources")
		delete(additionalProperties, "hostnames")
		delete(additionalProperties, "operatingSystems")
		delete(additionalProperties, "paths")
		delete(additionalProperties, "usages")
		delete(additionalProperties, "tlsPorts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHostDiscoveryData struct {
	value *HostDiscoveryData
	isSet bool
}

func (v NullableHostDiscoveryData) Get() *HostDiscoveryData {
	return v.value
}

func (v *NullableHostDiscoveryData) Set(val *HostDiscoveryData) {
	v.value = val
	v.isSet = true
}

func (v NullableHostDiscoveryData) IsSet() bool {
	return v.isSet
}

func (v *NullableHostDiscoveryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostDiscoveryData(val *HostDiscoveryData) *NullableHostDiscoveryData {
	return &NullableHostDiscoveryData{value: val, isSet: true}
}

func (v NullableHostDiscoveryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostDiscoveryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


