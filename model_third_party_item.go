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

// checks if the ThirdPartyItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyItem{}

// ThirdPartyItem struct for ThirdPartyItem
type ThirdPartyItem struct {
	// The third party connector name on which this certificate is synchronized
	Connector string `json:"connector"`
	// The Id of this certificate on the third party
	Id string `json:"id"`
	// The fingerprint of this certificate on the third party
	Fingerprint NullableString `json:"fingerprint,omitempty"`
	// The date when the certificate was pushed to this third party
	PushDate NullableInt64 `json:"pushDate,omitempty"`
	// The date when the certificate was removed from this third party (in case of revocation)
	RemoveDate NullableInt64 `json:"removeDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThirdPartyItem ThirdPartyItem

// NewThirdPartyItem instantiates a new ThirdPartyItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyItem(connector string, id string) *ThirdPartyItem {
	this := ThirdPartyItem{}
	this.Connector = connector
	this.Id = id
	return &this
}

// NewThirdPartyItemWithDefaults instantiates a new ThirdPartyItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyItemWithDefaults() *ThirdPartyItem {
	this := ThirdPartyItem{}
	return &this
}

// GetConnector returns the Connector field value
func (o *ThirdPartyItem) GetConnector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyItem) GetConnectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connector, true
}

// SetConnector sets field value
func (o *ThirdPartyItem) SetConnector(v string) {
	o.Connector = v
}

// GetId returns the Id field value
func (o *ThirdPartyItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyItem) SetId(v string) {
	o.Id = v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyItem) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint.Get()) {
		var ret string
		return ret
	}
	return *o.Fingerprint.Get()
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyItem) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fingerprint.Get(), o.Fingerprint.IsSet()
}

// HasFingerprint returns a boolean if a field has been set.
func (o *ThirdPartyItem) HasFingerprint() bool {
	if o != nil && o.Fingerprint.IsSet() {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given NullableString and assigns it to the Fingerprint field.
func (o *ThirdPartyItem) SetFingerprint(v string) {
	o.Fingerprint.Set(&v)
}
// SetFingerprintNil sets the value for Fingerprint to be an explicit nil
func (o *ThirdPartyItem) SetFingerprintNil() {
	o.Fingerprint.Set(nil)
}

// UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
func (o *ThirdPartyItem) UnsetFingerprint() {
	o.Fingerprint.Unset()
}

// GetPushDate returns the PushDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyItem) GetPushDate() int64 {
	if o == nil || IsNil(o.PushDate.Get()) {
		var ret int64
		return ret
	}
	return *o.PushDate.Get()
}

// GetPushDateOk returns a tuple with the PushDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyItem) GetPushDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PushDate.Get(), o.PushDate.IsSet()
}

// HasPushDate returns a boolean if a field has been set.
func (o *ThirdPartyItem) HasPushDate() bool {
	if o != nil && o.PushDate.IsSet() {
		return true
	}

	return false
}

// SetPushDate gets a reference to the given NullableInt64 and assigns it to the PushDate field.
func (o *ThirdPartyItem) SetPushDate(v int64) {
	o.PushDate.Set(&v)
}
// SetPushDateNil sets the value for PushDate to be an explicit nil
func (o *ThirdPartyItem) SetPushDateNil() {
	o.PushDate.Set(nil)
}

// UnsetPushDate ensures that no value is present for PushDate, not even an explicit nil
func (o *ThirdPartyItem) UnsetPushDate() {
	o.PushDate.Unset()
}

// GetRemoveDate returns the RemoveDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyItem) GetRemoveDate() int64 {
	if o == nil || IsNil(o.RemoveDate.Get()) {
		var ret int64
		return ret
	}
	return *o.RemoveDate.Get()
}

// GetRemoveDateOk returns a tuple with the RemoveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyItem) GetRemoveDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoveDate.Get(), o.RemoveDate.IsSet()
}

// HasRemoveDate returns a boolean if a field has been set.
func (o *ThirdPartyItem) HasRemoveDate() bool {
	if o != nil && o.RemoveDate.IsSet() {
		return true
	}

	return false
}

// SetRemoveDate gets a reference to the given NullableInt64 and assigns it to the RemoveDate field.
func (o *ThirdPartyItem) SetRemoveDate(v int64) {
	o.RemoveDate.Set(&v)
}
// SetRemoveDateNil sets the value for RemoveDate to be an explicit nil
func (o *ThirdPartyItem) SetRemoveDateNil() {
	o.RemoveDate.Set(nil)
}

// UnsetRemoveDate ensures that no value is present for RemoveDate, not even an explicit nil
func (o *ThirdPartyItem) UnsetRemoveDate() {
	o.RemoveDate.Unset()
}

func (o ThirdPartyItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connector"] = o.Connector
	toSerialize["id"] = o.Id
	if o.Fingerprint.IsSet() {
		toSerialize["fingerprint"] = o.Fingerprint.Get()
	}
	if o.PushDate.IsSet() {
		toSerialize["pushDate"] = o.PushDate.Get()
	}
	if o.RemoveDate.IsSet() {
		toSerialize["removeDate"] = o.RemoveDate.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThirdPartyItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connector",
		"id",
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

	varThirdPartyItem := _ThirdPartyItem{}

	err = json.Unmarshal(data, &varThirdPartyItem)

	if err != nil {
		return err
	}

	*o = ThirdPartyItem(varThirdPartyItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connector")
		delete(additionalProperties, "id")
		delete(additionalProperties, "fingerprint")
		delete(additionalProperties, "pushDate")
		delete(additionalProperties, "removeDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThirdPartyItem struct {
	value *ThirdPartyItem
	isSet bool
}

func (v NullableThirdPartyItem) Get() *ThirdPartyItem {
	return v.value
}

func (v *NullableThirdPartyItem) Set(val *ThirdPartyItem) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyItem) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyItem(val *ThirdPartyItem) *NullableThirdPartyItem {
	return &NullableThirdPartyItem{value: val, isSet: true}
}

func (v NullableThirdPartyItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


