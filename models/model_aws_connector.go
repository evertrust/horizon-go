/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the AWSConnector type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AWSConnector{}

// AWSConnector struct for AWSConnector
type AWSConnector struct {
	// Name of the `password` [credentials](#tag/security.credentials) containing Access Key Id and Secret Access Key. If not defined, an account present in environment variables can be used.
	Credentials          utils.NullableString `json:"credentials,omitempty"`
	Name                 string               `json:"name"`
	Proxy                utils.NullableString `json:"proxy,omitempty"`
	Region               string               `json:"region"`
	RenewalPeriod        utils.NullableString `json:"renewalPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	ResourceGroupName    utils.NullableString `json:"resourceGroupName,omitempty"`
	RoleArn              utils.NullableString `json:"roleArn,omitempty"`
	TagKey               utils.NullableString `json:"tagKey,omitempty"`
	TagValue             utils.NullableString `json:"tagValue,omitempty"`
	ThrottleDuration     string               `json:"throttleDuration" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Timeout              utils.NullableString `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string               `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _AWSConnector AWSConnector

// NewAWSConnector instantiates a new AWSConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSConnector(name string, region string, throttleDuration string, type_ string) *AWSConnector {
	this := AWSConnector{}
	this.Name = name
	this.Region = region
	this.ThrottleDuration = throttleDuration
	this.Type = type_
	return &this
}

// NewAWSConnectorWithDefaults instantiates a new AWSConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSConnectorWithDefaults() *AWSConnector {
	this := AWSConnector{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSConnector) GetCredentials() string {
	if o == nil || utils.IsNil(o.Credentials.Get()) {
		var ret string
		return ret
	}
	return *o.Credentials.Get()
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSConnector) GetCredentialsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials.Get(), o.Credentials.IsSet()
}

// HasCredentials returns a boolean if a field has been set.
func (o *AWSConnector) HasCredentials() bool {
	if o != nil && o.Credentials.IsSet() {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given NullableString and assigns it to the Credentials field.
func (o *AWSConnector) SetCredentials(v string) {
	o.Credentials.Set(&v)
}

// SetCredentialsNil sets the value for Credentials to be an explicit nil
func (o *AWSConnector) SetCredentialsNil() {
	o.Credentials.Set(nil)
}

// UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
func (o *AWSConnector) UnsetCredentials() {
	o.Credentials.Unset()
}

// GetName returns the Name field value
func (o *AWSConnector) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AWSConnector) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AWSConnector) SetName(v string) {
	o.Name = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSConnector) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSConnector) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *AWSConnector) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *AWSConnector) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *AWSConnector) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *AWSConnector) UnsetProxy() {
	o.Proxy.Unset()
}

// GetRegion returns the Region field value
func (o *AWSConnector) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AWSConnector) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *AWSConnector) SetRegion(v string) {
	o.Region = v
}

// GetRenewalPeriod returns the RenewalPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSConnector) GetRenewalPeriod() string {
	if o == nil || utils.IsNil(o.RenewalPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RenewalPeriod.Get()
}

// GetRenewalPeriodOk returns a tuple with the RenewalPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSConnector) GetRenewalPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewalPeriod.Get(), o.RenewalPeriod.IsSet()
}

// HasRenewalPeriod returns a boolean if a field has been set.
func (o *AWSConnector) HasRenewalPeriod() bool {
	if o != nil && o.RenewalPeriod.IsSet() {
		return true
	}

	return false
}

// SetRenewalPeriod gets a reference to the given NullableString and assigns it to the RenewalPeriod field.
func (o *AWSConnector) SetRenewalPeriod(v string) {
	o.RenewalPeriod.Set(&v)
}

// SetRenewalPeriodNil sets the value for RenewalPeriod to be an explicit nil
func (o *AWSConnector) SetRenewalPeriodNil() {
	o.RenewalPeriod.Set(nil)
}

// UnsetRenewalPeriod ensures that no value is present for RenewalPeriod, not even an explicit nil
func (o *AWSConnector) UnsetRenewalPeriod() {
	o.RenewalPeriod.Unset()
}

// GetResourceGroupName returns the ResourceGroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSConnector) GetResourceGroupName() string {
	if o == nil || utils.IsNil(o.ResourceGroupName.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceGroupName.Get()
}

// GetResourceGroupNameOk returns a tuple with the ResourceGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSConnector) GetResourceGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceGroupName.Get(), o.ResourceGroupName.IsSet()
}

// HasResourceGroupName returns a boolean if a field has been set.
func (o *AWSConnector) HasResourceGroupName() bool {
	if o != nil && o.ResourceGroupName.IsSet() {
		return true
	}

	return false
}

// SetResourceGroupName gets a reference to the given NullableString and assigns it to the ResourceGroupName field.
func (o *AWSConnector) SetResourceGroupName(v string) {
	o.ResourceGroupName.Set(&v)
}

// SetResourceGroupNameNil sets the value for ResourceGroupName to be an explicit nil
func (o *AWSConnector) SetResourceGroupNameNil() {
	o.ResourceGroupName.Set(nil)
}

// UnsetResourceGroupName ensures that no value is present for ResourceGroupName, not even an explicit nil
func (o *AWSConnector) UnsetResourceGroupName() {
	o.ResourceGroupName.Unset()
}

// GetRoleArn returns the RoleArn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSConnector) GetRoleArn() string {
	if o == nil || utils.IsNil(o.RoleArn.Get()) {
		var ret string
		return ret
	}
	return *o.RoleArn.Get()
}

// GetRoleArnOk returns a tuple with the RoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSConnector) GetRoleArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoleArn.Get(), o.RoleArn.IsSet()
}

// HasRoleArn returns a boolean if a field has been set.
func (o *AWSConnector) HasRoleArn() bool {
	if o != nil && o.RoleArn.IsSet() {
		return true
	}

	return false
}

// SetRoleArn gets a reference to the given NullableString and assigns it to the RoleArn field.
func (o *AWSConnector) SetRoleArn(v string) {
	o.RoleArn.Set(&v)
}

// SetRoleArnNil sets the value for RoleArn to be an explicit nil
func (o *AWSConnector) SetRoleArnNil() {
	o.RoleArn.Set(nil)
}

// UnsetRoleArn ensures that no value is present for RoleArn, not even an explicit nil
func (o *AWSConnector) UnsetRoleArn() {
	o.RoleArn.Unset()
}

// GetTagKey returns the TagKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSConnector) GetTagKey() string {
	if o == nil || utils.IsNil(o.TagKey.Get()) {
		var ret string
		return ret
	}
	return *o.TagKey.Get()
}

// GetTagKeyOk returns a tuple with the TagKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSConnector) GetTagKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagKey.Get(), o.TagKey.IsSet()
}

// HasTagKey returns a boolean if a field has been set.
func (o *AWSConnector) HasTagKey() bool {
	if o != nil && o.TagKey.IsSet() {
		return true
	}

	return false
}

// SetTagKey gets a reference to the given NullableString and assigns it to the TagKey field.
func (o *AWSConnector) SetTagKey(v string) {
	o.TagKey.Set(&v)
}

// SetTagKeyNil sets the value for TagKey to be an explicit nil
func (o *AWSConnector) SetTagKeyNil() {
	o.TagKey.Set(nil)
}

// UnsetTagKey ensures that no value is present for TagKey, not even an explicit nil
func (o *AWSConnector) UnsetTagKey() {
	o.TagKey.Unset()
}

// GetTagValue returns the TagValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSConnector) GetTagValue() string {
	if o == nil || utils.IsNil(o.TagValue.Get()) {
		var ret string
		return ret
	}
	return *o.TagValue.Get()
}

// GetTagValueOk returns a tuple with the TagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSConnector) GetTagValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagValue.Get(), o.TagValue.IsSet()
}

// HasTagValue returns a boolean if a field has been set.
func (o *AWSConnector) HasTagValue() bool {
	if o != nil && o.TagValue.IsSet() {
		return true
	}

	return false
}

// SetTagValue gets a reference to the given NullableString and assigns it to the TagValue field.
func (o *AWSConnector) SetTagValue(v string) {
	o.TagValue.Set(&v)
}

// SetTagValueNil sets the value for TagValue to be an explicit nil
func (o *AWSConnector) SetTagValueNil() {
	o.TagValue.Set(nil)
}

// UnsetTagValue ensures that no value is present for TagValue, not even an explicit nil
func (o *AWSConnector) UnsetTagValue() {
	o.TagValue.Unset()
}

// GetThrottleDuration returns the ThrottleDuration field value
func (o *AWSConnector) GetThrottleDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThrottleDuration
}

// GetThrottleDurationOk returns a tuple with the ThrottleDuration field value
// and a boolean to check if the value has been set.
func (o *AWSConnector) GetThrottleDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThrottleDuration, true
}

// SetThrottleDuration sets field value
func (o *AWSConnector) SetThrottleDuration(v string) {
	o.ThrottleDuration = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSConnector) GetTimeout() string {
	if o == nil || utils.IsNil(o.Timeout.Get()) {
		var ret string
		return ret
	}
	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AWSConnector) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// HasTimeout returns a boolean if a field has been set.
func (o *AWSConnector) HasTimeout() bool {
	if o != nil && o.Timeout.IsSet() {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given NullableString and assigns it to the Timeout field.
func (o *AWSConnector) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// SetTimeoutNil sets the value for Timeout to be an explicit nil
func (o *AWSConnector) SetTimeoutNil() {
	o.Timeout.Set(nil)
}

// UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
func (o *AWSConnector) UnsetTimeout() {
	o.Timeout.Unset()
}

// GetType returns the Type field value
func (o *AWSConnector) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSConnector) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AWSConnector) SetType(v string) {
	o.Type = v
}

func (o AWSConnector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AWSConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Credentials.IsSet() {
		toSerialize["credentials"] = o.Credentials.Get()
	}
	toSerialize["name"] = o.Name
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	toSerialize["region"] = o.Region
	if o.RenewalPeriod.IsSet() {
		toSerialize["renewalPeriod"] = o.RenewalPeriod.Get()
	}
	if o.ResourceGroupName.IsSet() {
		toSerialize["resourceGroupName"] = o.ResourceGroupName.Get()
	}
	if o.RoleArn.IsSet() {
		toSerialize["roleArn"] = o.RoleArn.Get()
	}
	if o.TagKey.IsSet() {
		toSerialize["tagKey"] = o.TagKey.Get()
	}
	if o.TagValue.IsSet() {
		toSerialize["tagValue"] = o.TagValue.Get()
	}
	toSerialize["throttleDuration"] = o.ThrottleDuration
	if o.Timeout.IsSet() {
		toSerialize["timeout"] = o.Timeout.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AWSConnector) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"region",
		"throttleDuration",
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

	varAWSConnector := _AWSConnector{}

	err = json.Unmarshal(data, &varAWSConnector)

	if err != nil {
		return err
	}

	*o = AWSConnector(varAWSConnector)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "name")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "region")
		delete(additionalProperties, "renewalPeriod")
		delete(additionalProperties, "resourceGroupName")
		delete(additionalProperties, "roleArn")
		delete(additionalProperties, "tagKey")
		delete(additionalProperties, "tagValue")
		delete(additionalProperties, "throttleDuration")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAWSConnector struct {
	value *AWSConnector
	isSet bool
}

func (v NullableAWSConnector) Get() *AWSConnector {
	return v.value
}

func (v *NullableAWSConnector) Set(val *AWSConnector) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSConnector(val *AWSConnector) *NullableAWSConnector {
	return &NullableAWSConnector{value: val, isSet: true}
}

func (v NullableAWSConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
