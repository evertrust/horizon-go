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

// checks if the Base type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Base{}

// Base struct for Base
type Base struct {
	// Name of the notification
	Name *string `json:"name,omitempty"`
	// The type of notification
	Type *string `json:"type,omitempty"`
	// Number of retries when the notification fails
	Retries NullableInt64 `json:"retries,omitempty"`
	// Time period at which the notification needs to run. Can only be defined on expiration and pending events.
	RunPeriod NullableString `json:"runPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// License usage at which the notification needs to run (between 0 and 100). Must be defined on `on_license_usage` event and must NOT be defined otherwise.
	LicenseUsagePercent NullableInt64 `json:"licenseUsagePercent,omitempty"`
	// Event on which the notification runs. This MUST contain only one value.
	Events []string `json:"events,omitempty"`
	// Must be defined on `on_expire` event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed.
	RunOnRenewed NullableBool `json:"runOnRenewed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Base Base

// NewBase instantiates a new Base object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBase() *Base {
	this := Base{}
	return &this
}

// NewBaseWithDefaults instantiates a new Base object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseWithDefaults() *Base {
	this := Base{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Base) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Base) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Base) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Base) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Base) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Base) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Base) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Base) SetType(v string) {
	o.Type = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Base) GetRetries() int64 {
	if o == nil || IsNil(o.Retries.Get()) {
		var ret int64
		return ret
	}
	return *o.Retries.Get()
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Base) GetRetriesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Retries.Get(), o.Retries.IsSet()
}

// HasRetries returns a boolean if a field has been set.
func (o *Base) HasRetries() bool {
	if o != nil && o.Retries.IsSet() {
		return true
	}

	return false
}

// SetRetries gets a reference to the given NullableInt64 and assigns it to the Retries field.
func (o *Base) SetRetries(v int64) {
	o.Retries.Set(&v)
}
// SetRetriesNil sets the value for Retries to be an explicit nil
func (o *Base) SetRetriesNil() {
	o.Retries.Set(nil)
}

// UnsetRetries ensures that no value is present for Retries, not even an explicit nil
func (o *Base) UnsetRetries() {
	o.Retries.Unset()
}

// GetRunPeriod returns the RunPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Base) GetRunPeriod() string {
	if o == nil || IsNil(o.RunPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RunPeriod.Get()
}

// GetRunPeriodOk returns a tuple with the RunPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Base) GetRunPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunPeriod.Get(), o.RunPeriod.IsSet()
}

// HasRunPeriod returns a boolean if a field has been set.
func (o *Base) HasRunPeriod() bool {
	if o != nil && o.RunPeriod.IsSet() {
		return true
	}

	return false
}

// SetRunPeriod gets a reference to the given NullableString and assigns it to the RunPeriod field.
func (o *Base) SetRunPeriod(v string) {
	o.RunPeriod.Set(&v)
}
// SetRunPeriodNil sets the value for RunPeriod to be an explicit nil
func (o *Base) SetRunPeriodNil() {
	o.RunPeriod.Set(nil)
}

// UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil
func (o *Base) UnsetRunPeriod() {
	o.RunPeriod.Unset()
}

// GetLicenseUsagePercent returns the LicenseUsagePercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Base) GetLicenseUsagePercent() int64 {
	if o == nil || IsNil(o.LicenseUsagePercent.Get()) {
		var ret int64
		return ret
	}
	return *o.LicenseUsagePercent.Get()
}

// GetLicenseUsagePercentOk returns a tuple with the LicenseUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Base) GetLicenseUsagePercentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicenseUsagePercent.Get(), o.LicenseUsagePercent.IsSet()
}

// HasLicenseUsagePercent returns a boolean if a field has been set.
func (o *Base) HasLicenseUsagePercent() bool {
	if o != nil && o.LicenseUsagePercent.IsSet() {
		return true
	}

	return false
}

// SetLicenseUsagePercent gets a reference to the given NullableInt64 and assigns it to the LicenseUsagePercent field.
func (o *Base) SetLicenseUsagePercent(v int64) {
	o.LicenseUsagePercent.Set(&v)
}
// SetLicenseUsagePercentNil sets the value for LicenseUsagePercent to be an explicit nil
func (o *Base) SetLicenseUsagePercentNil() {
	o.LicenseUsagePercent.Set(nil)
}

// UnsetLicenseUsagePercent ensures that no value is present for LicenseUsagePercent, not even an explicit nil
func (o *Base) UnsetLicenseUsagePercent() {
	o.LicenseUsagePercent.Unset()
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *Base) GetEvents() []string {
	if o == nil || IsNil(o.Events) {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Base) GetEventsOk() ([]string, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *Base) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *Base) SetEvents(v []string) {
	o.Events = v
}

// GetRunOnRenewed returns the RunOnRenewed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Base) GetRunOnRenewed() bool {
	if o == nil || IsNil(o.RunOnRenewed.Get()) {
		var ret bool
		return ret
	}
	return *o.RunOnRenewed.Get()
}

// GetRunOnRenewedOk returns a tuple with the RunOnRenewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Base) GetRunOnRenewedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunOnRenewed.Get(), o.RunOnRenewed.IsSet()
}

// HasRunOnRenewed returns a boolean if a field has been set.
func (o *Base) HasRunOnRenewed() bool {
	if o != nil && o.RunOnRenewed.IsSet() {
		return true
	}

	return false
}

// SetRunOnRenewed gets a reference to the given NullableBool and assigns it to the RunOnRenewed field.
func (o *Base) SetRunOnRenewed(v bool) {
	o.RunOnRenewed.Set(&v)
}
// SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil
func (o *Base) SetRunOnRenewedNil() {
	o.RunOnRenewed.Set(nil)
}

// UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil
func (o *Base) UnsetRunOnRenewed() {
	o.RunOnRenewed.Unset()
}

func (o Base) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Base) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Retries.IsSet() {
		toSerialize["retries"] = o.Retries.Get()
	}
	if o.RunPeriod.IsSet() {
		toSerialize["runPeriod"] = o.RunPeriod.Get()
	}
	if o.LicenseUsagePercent.IsSet() {
		toSerialize["licenseUsagePercent"] = o.LicenseUsagePercent.Get()
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if o.RunOnRenewed.IsSet() {
		toSerialize["runOnRenewed"] = o.RunOnRenewed.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Base) UnmarshalJSON(data []byte) (err error) {
	varBase := _Base{}

	err = json.Unmarshal(data, &varBase)

	if err != nil {
		return err
	}

	*o = Base(varBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "retries")
		delete(additionalProperties, "runPeriod")
		delete(additionalProperties, "licenseUsagePercent")
		delete(additionalProperties, "events")
		delete(additionalProperties, "runOnRenewed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBase struct {
	value *Base
	isSet bool
}

func (v NullableBase) Get() *Base {
	return v.value
}

func (v *NullableBase) Set(val *Base) {
	v.value = val
	v.isSet = true
}

func (v NullableBase) IsSet() bool {
	return v.isSet
}

func (v *NullableBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBase(val *Base) *NullableBase {
	return &NullableBase{value: val, isSet: true}
}

func (v NullableBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


