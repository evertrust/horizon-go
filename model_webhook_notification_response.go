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

// checks if the WebhookNotificationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookNotificationResponse{}

// WebhookNotificationResponse struct for WebhookNotificationResponse
type WebhookNotificationResponse struct {
	Type string `json:"type"`
	// Number of retries when the notification fails (non 200 return code)
	Retries interface{} `json:"retries,omitempty"`
	WebhookTemplate WebhookTemplate `json:"webhookTemplate"`
	// Name of a Proxy to use while sending the webhook
	Proxy NullableString `json:"proxy,omitempty"`
	// Timeout for the webhook request
	Timeout *string `json:"timeout,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// Name of the notification
	Name string `json:"name"`
	// Time period at which the notification needs to run. Can only be defined on expiration and pending events.
	RunPeriod NullableString `json:"runPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	// License usage at which the notification needs to run (between 0 and 100). Must be defined on `on_license_usage` event and must NOT be defined otherwise.
	LicenseUsagePercent NullableInt64 `json:"licenseUsagePercent,omitempty"`
	// Event on which the notification runs. This MUST contain only one value.
	Events []string `json:"events"`
	// Must be defined on `on_expire` event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed.
	RunOnRenewed NullableBool `json:"runOnRenewed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebhookNotificationResponse WebhookNotificationResponse

// NewWebhookNotificationResponse instantiates a new WebhookNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookNotificationResponse(type_ string, webhookTemplate WebhookTemplate, name string, events []string) *WebhookNotificationResponse {
	this := WebhookNotificationResponse{}
	this.Name = name
	this.Type = type_
	this.Events = events
	return &this
}

// NewWebhookNotificationResponseWithDefaults instantiates a new WebhookNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookNotificationResponseWithDefaults() *WebhookNotificationResponse {
	this := WebhookNotificationResponse{}
	return &this
}

// GetType returns the Type field value
func (o *WebhookNotificationResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhookNotificationResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebhookNotificationResponse) SetType(v string) {
	o.Type = v
}

// GetRetries returns the Retries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookNotificationResponse) GetRetries() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookNotificationResponse) GetRetriesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Retries) {
		return nil, false
	}
	return &o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *WebhookNotificationResponse) HasRetries() bool {
	if o != nil && !IsNil(o.Retries) {
		return true
	}

	return false
}

// SetRetries gets a reference to the given interface{} and assigns it to the Retries field.
func (o *WebhookNotificationResponse) SetRetries(v interface{}) {
	o.Retries = v
}

// GetWebhookTemplate returns the WebhookTemplate field value
func (o *WebhookNotificationResponse) GetWebhookTemplate() WebhookTemplate {
	if o == nil {
		var ret WebhookTemplate
		return ret
	}

	return o.WebhookTemplate
}

// GetWebhookTemplateOk returns a tuple with the WebhookTemplate field value
// and a boolean to check if the value has been set.
func (o *WebhookNotificationResponse) GetWebhookTemplateOk() (*WebhookTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookTemplate, true
}

// SetWebhookTemplate sets field value
func (o *WebhookNotificationResponse) SetWebhookTemplate(v WebhookTemplate) {
	o.WebhookTemplate = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookNotificationResponse) GetProxy() string {
	if o == nil || IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookNotificationResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *WebhookNotificationResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *WebhookNotificationResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}
// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *WebhookNotificationResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *WebhookNotificationResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *WebhookNotificationResponse) GetTimeout() string {
	if o == nil || IsNil(o.Timeout) {
		var ret string
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookNotificationResponse) GetTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *WebhookNotificationResponse) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given string and assigns it to the Timeout field.
func (o *WebhookNotificationResponse) SetTimeout(v string) {
	o.Timeout = &v
}

// GetName returns the Name field value
func (o *WebhookNotificationResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebhookNotificationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebhookNotificationResponse) SetName(v string) {
	o.Name = v
}

// GetRunPeriod returns the RunPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookNotificationResponse) GetRunPeriod() string {
	if o == nil || IsNil(o.RunPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RunPeriod.Get()
}

// GetRunPeriodOk returns a tuple with the RunPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookNotificationResponse) GetRunPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunPeriod.Get(), o.RunPeriod.IsSet()
}

// HasRunPeriod returns a boolean if a field has been set.
func (o *WebhookNotificationResponse) HasRunPeriod() bool {
	if o != nil && o.RunPeriod.IsSet() {
		return true
	}

	return false
}

// SetRunPeriod gets a reference to the given NullableString and assigns it to the RunPeriod field.
func (o *WebhookNotificationResponse) SetRunPeriod(v string) {
	o.RunPeriod.Set(&v)
}
// SetRunPeriodNil sets the value for RunPeriod to be an explicit nil
func (o *WebhookNotificationResponse) SetRunPeriodNil() {
	o.RunPeriod.Set(nil)
}

// UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil
func (o *WebhookNotificationResponse) UnsetRunPeriod() {
	o.RunPeriod.Unset()
}

// GetLicenseUsagePercent returns the LicenseUsagePercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookNotificationResponse) GetLicenseUsagePercent() int64 {
	if o == nil || IsNil(o.LicenseUsagePercent.Get()) {
		var ret int64
		return ret
	}
	return *o.LicenseUsagePercent.Get()
}

// GetLicenseUsagePercentOk returns a tuple with the LicenseUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookNotificationResponse) GetLicenseUsagePercentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicenseUsagePercent.Get(), o.LicenseUsagePercent.IsSet()
}

// HasLicenseUsagePercent returns a boolean if a field has been set.
func (o *WebhookNotificationResponse) HasLicenseUsagePercent() bool {
	if o != nil && o.LicenseUsagePercent.IsSet() {
		return true
	}

	return false
}

// SetLicenseUsagePercent gets a reference to the given NullableInt64 and assigns it to the LicenseUsagePercent field.
func (o *WebhookNotificationResponse) SetLicenseUsagePercent(v int64) {
	o.LicenseUsagePercent.Set(&v)
}
// SetLicenseUsagePercentNil sets the value for LicenseUsagePercent to be an explicit nil
func (o *WebhookNotificationResponse) SetLicenseUsagePercentNil() {
	o.LicenseUsagePercent.Set(nil)
}

// UnsetLicenseUsagePercent ensures that no value is present for LicenseUsagePercent, not even an explicit nil
func (o *WebhookNotificationResponse) UnsetLicenseUsagePercent() {
	o.LicenseUsagePercent.Unset()
}

// GetEvents returns the Events field value
func (o *WebhookNotificationResponse) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *WebhookNotificationResponse) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *WebhookNotificationResponse) SetEvents(v []string) {
	o.Events = v
}

// GetRunOnRenewed returns the RunOnRenewed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookNotificationResponse) GetRunOnRenewed() bool {
	if o == nil || IsNil(o.RunOnRenewed.Get()) {
		var ret bool
		return ret
	}
	return *o.RunOnRenewed.Get()
}

// GetRunOnRenewedOk returns a tuple with the RunOnRenewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookNotificationResponse) GetRunOnRenewedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunOnRenewed.Get(), o.RunOnRenewed.IsSet()
}

// HasRunOnRenewed returns a boolean if a field has been set.
func (o *WebhookNotificationResponse) HasRunOnRenewed() bool {
	if o != nil && o.RunOnRenewed.IsSet() {
		return true
	}

	return false
}

// SetRunOnRenewed gets a reference to the given NullableBool and assigns it to the RunOnRenewed field.
func (o *WebhookNotificationResponse) SetRunOnRenewed(v bool) {
	o.RunOnRenewed.Set(&v)
}
// SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil
func (o *WebhookNotificationResponse) SetRunOnRenewedNil() {
	o.RunOnRenewed.Set(nil)
}

// UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil
func (o *WebhookNotificationResponse) UnsetRunOnRenewed() {
	o.RunOnRenewed.Unset()
}

func (o WebhookNotificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}
	toSerialize["webhookTemplate"] = o.WebhookTemplate
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	toSerialize["name"] = o.Name
	if o.RunPeriod.IsSet() {
		toSerialize["runPeriod"] = o.RunPeriod.Get()
	}
	if o.LicenseUsagePercent.IsSet() {
		toSerialize["licenseUsagePercent"] = o.LicenseUsagePercent.Get()
	}
	toSerialize["events"] = o.Events
	if o.RunOnRenewed.IsSet() {
		toSerialize["runOnRenewed"] = o.RunOnRenewed.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebhookNotificationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"webhookTemplate",
		"name",
		"events",
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

	varWebhookNotificationResponse := _WebhookNotificationResponse{}

	err = json.Unmarshal(data, &varWebhookNotificationResponse)

	if err != nil {
		return err
	}

	*o = WebhookNotificationResponse(varWebhookNotificationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "retries")
		delete(additionalProperties, "webhookTemplate")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "name")
		delete(additionalProperties, "runPeriod")
		delete(additionalProperties, "licenseUsagePercent")
		delete(additionalProperties, "events")
		delete(additionalProperties, "runOnRenewed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookNotificationResponse struct {
	value *WebhookNotificationResponse
	isSet bool
}

func (v NullableWebhookNotificationResponse) Get() *WebhookNotificationResponse {
	return v.value
}

func (v *NullableWebhookNotificationResponse) Set(val *WebhookNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookNotificationResponse(val *WebhookNotificationResponse) *NullableWebhookNotificationResponse {
	return &NullableWebhookNotificationResponse{value: val, isSet: true}
}

func (v NullableWebhookNotificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


