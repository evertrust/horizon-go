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

// checks if the DiscoveryFeed type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DiscoveryFeed{}

// DiscoveryFeed struct for DiscoveryFeed
type DiscoveryFeed struct {
	// The name of the discovery campaign to feed into
	Campaign string `json:"campaign"`
	// The PEM-encoded certificate to feed the discovery campaign with
	Certificate string `json:"certificate"`
	// The code of the event to raise in the discovery events
	Code utils.NullableString `json:"code,omitempty"`
	// The host discovery data to feed the discovery campaign with (discovery metadata)
	HostDiscoveryData HostDiscoveryData `json:"hostDiscoveryData"`
	// The list of certificate metadata to feed the discovery campaign with
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
	// The PEM-encoded private key to feed the discovery campaign with
	PrivateKey utils.NullableString `json:"privateKey,omitempty"`
	// The ID of the previously opened discovery feed session
	SessionId            utils.NullableString `json:"sessionId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiscoveryFeed DiscoveryFeed

// NewDiscoveryFeed instantiates a new DiscoveryFeed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryFeed(campaign string, certificate string, hostDiscoveryData HostDiscoveryData) *DiscoveryFeed {
	this := DiscoveryFeed{}
	this.Campaign = campaign
	this.Certificate = certificate
	this.HostDiscoveryData = hostDiscoveryData
	return &this
}

// NewDiscoveryFeedWithDefaults instantiates a new DiscoveryFeed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryFeedWithDefaults() *DiscoveryFeed {
	this := DiscoveryFeed{}
	return &this
}

// GetCampaign returns the Campaign field value
func (o *DiscoveryFeed) GetCampaign() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value
// and a boolean to check if the value has been set.
func (o *DiscoveryFeed) GetCampaignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaign, true
}

// SetCampaign sets field value
func (o *DiscoveryFeed) SetCampaign(v string) {
	o.Campaign = v
}

// GetCertificate returns the Certificate field value
func (o *DiscoveryFeed) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *DiscoveryFeed) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *DiscoveryFeed) SetCertificate(v string) {
	o.Certificate = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryFeed) GetCode() string {
	if o == nil || utils.IsNil(o.Code.Get()) {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryFeed) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *DiscoveryFeed) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *DiscoveryFeed) SetCode(v string) {
	o.Code.Set(&v)
}

// SetCodeNil sets the value for Code to be an explicit nil
func (o *DiscoveryFeed) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *DiscoveryFeed) UnsetCode() {
	o.Code.Unset()
}

// GetHostDiscoveryData returns the HostDiscoveryData field value
func (o *DiscoveryFeed) GetHostDiscoveryData() HostDiscoveryData {
	if o == nil {
		var ret HostDiscoveryData
		return ret
	}

	return o.HostDiscoveryData
}

// GetHostDiscoveryDataOk returns a tuple with the HostDiscoveryData field value
// and a boolean to check if the value has been set.
func (o *DiscoveryFeed) GetHostDiscoveryDataOk() (*HostDiscoveryData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostDiscoveryData, true
}

// SetHostDiscoveryData sets field value
func (o *DiscoveryFeed) SetHostDiscoveryData(v HostDiscoveryData) {
	o.HostDiscoveryData = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryFeed) GetMetadata() []CertificateMetadata {
	if o == nil {
		var ret []CertificateMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryFeed) GetMetadataOk() ([]CertificateMetadata, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DiscoveryFeed) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []CertificateMetadata and assigns it to the Metadata field.
func (o *DiscoveryFeed) SetMetadata(v []CertificateMetadata) {
	o.Metadata = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryFeed) GetPrivateKey() string {
	if o == nil || utils.IsNil(o.PrivateKey.Get()) {
		var ret string
		return ret
	}
	return *o.PrivateKey.Get()
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryFeed) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateKey.Get(), o.PrivateKey.IsSet()
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *DiscoveryFeed) HasPrivateKey() bool {
	if o != nil && o.PrivateKey.IsSet() {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given NullableString and assigns it to the PrivateKey field.
func (o *DiscoveryFeed) SetPrivateKey(v string) {
	o.PrivateKey.Set(&v)
}

// SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil
func (o *DiscoveryFeed) SetPrivateKeyNil() {
	o.PrivateKey.Set(nil)
}

// UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
func (o *DiscoveryFeed) UnsetPrivateKey() {
	o.PrivateKey.Unset()
}

// GetSessionId returns the SessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DiscoveryFeed) GetSessionId() string {
	if o == nil || utils.IsNil(o.SessionId.Get()) {
		var ret string
		return ret
	}
	return *o.SessionId.Get()
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DiscoveryFeed) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionId.Get(), o.SessionId.IsSet()
}

// HasSessionId returns a boolean if a field has been set.
func (o *DiscoveryFeed) HasSessionId() bool {
	if o != nil && o.SessionId.IsSet() {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given NullableString and assigns it to the SessionId field.
func (o *DiscoveryFeed) SetSessionId(v string) {
	o.SessionId.Set(&v)
}

// SetSessionIdNil sets the value for SessionId to be an explicit nil
func (o *DiscoveryFeed) SetSessionIdNil() {
	o.SessionId.Set(nil)
}

// UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
func (o *DiscoveryFeed) UnsetSessionId() {
	o.SessionId.Unset()
}

func (o DiscoveryFeed) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryFeed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaign"] = o.Campaign
	toSerialize["certificate"] = o.Certificate
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	toSerialize["hostDiscoveryData"] = o.HostDiscoveryData
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.PrivateKey.IsSet() {
		toSerialize["privateKey"] = o.PrivateKey.Get()
	}
	if o.SessionId.IsSet() {
		toSerialize["sessionId"] = o.SessionId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiscoveryFeed) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"campaign",
		"certificate",
		"hostDiscoveryData",
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

	varDiscoveryFeed := _DiscoveryFeed{}

	err = json.Unmarshal(data, &varDiscoveryFeed)

	if err != nil {
		return err
	}

	*o = DiscoveryFeed(varDiscoveryFeed)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "campaign")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "code")
		delete(additionalProperties, "hostDiscoveryData")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "privateKey")
		delete(additionalProperties, "sessionId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiscoveryFeed struct {
	value *DiscoveryFeed
	isSet bool
}

func (v NullableDiscoveryFeed) Get() *DiscoveryFeed {
	return v.value
}

func (v *NullableDiscoveryFeed) Set(val *DiscoveryFeed) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryFeed) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryFeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryFeed(val *DiscoveryFeed) *NullableDiscoveryFeed {
	return &NullableDiscoveryFeed{value: val, isSet: true}
}

func (v NullableDiscoveryFeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryFeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
