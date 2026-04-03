/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the EmailNotificationResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EmailNotificationResponse{}

// EmailNotificationResponse struct for EmailNotificationResponse
type EmailNotificationResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Attach the certificate in DER format if available
	AttachDerCertificate utils.NullableBool `json:"attachDerCertificate,omitempty"`
	// Attach the certificate and its trust chain (bundle) in PEM format if available
	AttachPemBundle utils.NullableBool `json:"attachPemBundle,omitempty"`
	// Attach the certificate in PEM format if available
	AttachPemCertificate utils.NullableBool `json:"attachPemCertificate,omitempty"`
	// Attach the certificate in PKCS7 format if available
	AttachPkcs7 utils.NullableBool `json:"attachPkcs7,omitempty"`
	// Attach the certificate and its trust chain (bundle) in PKCS7 format if available
	AttachPkcs7Bundle utils.NullableBool `json:"attachPkcs7Bundle,omitempty"`
	// Attach the certificate in PKCS#12 format if available
	AttachPkcs12  utils.NullableBool `json:"attachPkcs12,omitempty"`
	EmailTemplate EmailTemplate      `json:"emailTemplate"`
	// On events triggering an enrollment, select if mail is sent: - **Always**: set the value to `null`  - **Only when a PKCS#12 is available in the request**: set the value to `true`  - **Only when a PKCS#12 is not in the request**: set the value to `false`
	IfPkcs12 utils.NullableBool `json:"ifPkcs12,omitempty"`
	Type     string             `json:"type"`
	// Event on which the notification runs. This MUST contain only one value.
	Events []string `json:"events"`
	// License usage at which the notification needs to run (between 0 and 100). Must be defined on `on_license_usage` event and must NOT be defined otherwise.
	LicenseUsagePercent utils.NullableInt64 `json:"licenseUsagePercent,omitempty"`
	// Name of the notification
	Name string `json:"name"`
	// Number of retries when the notification fails
	Retries utils.NullableInt64 `json:"retries,omitempty"`
	// Must be defined on `on_expire` event and must NOT be defined otherwise. If true, the notification runs even if the certificate was renewed.
	RunOnRenewed utils.NullableBool `json:"runOnRenewed,omitempty"`
	// Time period at which the notification needs to run. Can only be defined on expiration and pending events.
	RunPeriod            utils.NullableString `json:"runPeriod,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	AdditionalProperties map[string]interface{}
}

type _EmailNotificationResponse EmailNotificationResponse

// NewEmailNotificationResponse instantiates a new EmailNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotificationResponse(id string, emailTemplate EmailTemplate, type_ string, events []string, name string) *EmailNotificationResponse {
	this := EmailNotificationResponse{}
	this.Events = events
	this.Name = name
	this.Type = type_
	return &this
}

// NewEmailNotificationResponseWithDefaults instantiates a new EmailNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationResponseWithDefaults() *EmailNotificationResponse {
	this := EmailNotificationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *EmailNotificationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EmailNotificationResponse) SetId(v string) {
	o.Id = v
}

// GetAttachDerCertificate returns the AttachDerCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailNotificationResponse) GetAttachDerCertificate() bool {
	if o == nil || utils.IsNil(o.AttachDerCertificate.Get()) {
		var ret bool
		return ret
	}
	return *o.AttachDerCertificate.Get()
}

// GetAttachDerCertificateOk returns a tuple with the AttachDerCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailNotificationResponse) GetAttachDerCertificateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachDerCertificate.Get(), o.AttachDerCertificate.IsSet()
}

// HasAttachDerCertificate returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasAttachDerCertificate() bool {
	if o != nil && o.AttachDerCertificate.IsSet() {
		return true
	}

	return false
}

// SetAttachDerCertificate gets a reference to the given NullableBool and assigns it to the AttachDerCertificate field.
func (o *EmailNotificationResponse) SetAttachDerCertificate(v bool) {
	o.AttachDerCertificate.Set(&v)
}

// SetAttachDerCertificateNil sets the value for AttachDerCertificate to be an explicit nil
func (o *EmailNotificationResponse) SetAttachDerCertificateNil() {
	o.AttachDerCertificate.Set(nil)
}

// UnsetAttachDerCertificate ensures that no value is present for AttachDerCertificate, not even an explicit nil
func (o *EmailNotificationResponse) UnsetAttachDerCertificate() {
	o.AttachDerCertificate.Unset()
}

// GetAttachPemBundle returns the AttachPemBundle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailNotificationResponse) GetAttachPemBundle() bool {
	if o == nil || utils.IsNil(o.AttachPemBundle.Get()) {
		var ret bool
		return ret
	}
	return *o.AttachPemBundle.Get()
}

// GetAttachPemBundleOk returns a tuple with the AttachPemBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailNotificationResponse) GetAttachPemBundleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachPemBundle.Get(), o.AttachPemBundle.IsSet()
}

// HasAttachPemBundle returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasAttachPemBundle() bool {
	if o != nil && o.AttachPemBundle.IsSet() {
		return true
	}

	return false
}

// SetAttachPemBundle gets a reference to the given NullableBool and assigns it to the AttachPemBundle field.
func (o *EmailNotificationResponse) SetAttachPemBundle(v bool) {
	o.AttachPemBundle.Set(&v)
}

// SetAttachPemBundleNil sets the value for AttachPemBundle to be an explicit nil
func (o *EmailNotificationResponse) SetAttachPemBundleNil() {
	o.AttachPemBundle.Set(nil)
}

// UnsetAttachPemBundle ensures that no value is present for AttachPemBundle, not even an explicit nil
func (o *EmailNotificationResponse) UnsetAttachPemBundle() {
	o.AttachPemBundle.Unset()
}

// GetAttachPemCertificate returns the AttachPemCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailNotificationResponse) GetAttachPemCertificate() bool {
	if o == nil || utils.IsNil(o.AttachPemCertificate.Get()) {
		var ret bool
		return ret
	}
	return *o.AttachPemCertificate.Get()
}

// GetAttachPemCertificateOk returns a tuple with the AttachPemCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailNotificationResponse) GetAttachPemCertificateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachPemCertificate.Get(), o.AttachPemCertificate.IsSet()
}

// HasAttachPemCertificate returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasAttachPemCertificate() bool {
	if o != nil && o.AttachPemCertificate.IsSet() {
		return true
	}

	return false
}

// SetAttachPemCertificate gets a reference to the given NullableBool and assigns it to the AttachPemCertificate field.
func (o *EmailNotificationResponse) SetAttachPemCertificate(v bool) {
	o.AttachPemCertificate.Set(&v)
}

// SetAttachPemCertificateNil sets the value for AttachPemCertificate to be an explicit nil
func (o *EmailNotificationResponse) SetAttachPemCertificateNil() {
	o.AttachPemCertificate.Set(nil)
}

// UnsetAttachPemCertificate ensures that no value is present for AttachPemCertificate, not even an explicit nil
func (o *EmailNotificationResponse) UnsetAttachPemCertificate() {
	o.AttachPemCertificate.Unset()
}

// GetAttachPkcs7 returns the AttachPkcs7 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailNotificationResponse) GetAttachPkcs7() bool {
	if o == nil || utils.IsNil(o.AttachPkcs7.Get()) {
		var ret bool
		return ret
	}
	return *o.AttachPkcs7.Get()
}

// GetAttachPkcs7Ok returns a tuple with the AttachPkcs7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailNotificationResponse) GetAttachPkcs7Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachPkcs7.Get(), o.AttachPkcs7.IsSet()
}

// HasAttachPkcs7 returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasAttachPkcs7() bool {
	if o != nil && o.AttachPkcs7.IsSet() {
		return true
	}

	return false
}

// SetAttachPkcs7 gets a reference to the given NullableBool and assigns it to the AttachPkcs7 field.
func (o *EmailNotificationResponse) SetAttachPkcs7(v bool) {
	o.AttachPkcs7.Set(&v)
}

// SetAttachPkcs7Nil sets the value for AttachPkcs7 to be an explicit nil
func (o *EmailNotificationResponse) SetAttachPkcs7Nil() {
	o.AttachPkcs7.Set(nil)
}

// UnsetAttachPkcs7 ensures that no value is present for AttachPkcs7, not even an explicit nil
func (o *EmailNotificationResponse) UnsetAttachPkcs7() {
	o.AttachPkcs7.Unset()
}

// GetAttachPkcs7Bundle returns the AttachPkcs7Bundle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailNotificationResponse) GetAttachPkcs7Bundle() bool {
	if o == nil || utils.IsNil(o.AttachPkcs7Bundle.Get()) {
		var ret bool
		return ret
	}
	return *o.AttachPkcs7Bundle.Get()
}

// GetAttachPkcs7BundleOk returns a tuple with the AttachPkcs7Bundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailNotificationResponse) GetAttachPkcs7BundleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachPkcs7Bundle.Get(), o.AttachPkcs7Bundle.IsSet()
}

// HasAttachPkcs7Bundle returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasAttachPkcs7Bundle() bool {
	if o != nil && o.AttachPkcs7Bundle.IsSet() {
		return true
	}

	return false
}

// SetAttachPkcs7Bundle gets a reference to the given NullableBool and assigns it to the AttachPkcs7Bundle field.
func (o *EmailNotificationResponse) SetAttachPkcs7Bundle(v bool) {
	o.AttachPkcs7Bundle.Set(&v)
}

// SetAttachPkcs7BundleNil sets the value for AttachPkcs7Bundle to be an explicit nil
func (o *EmailNotificationResponse) SetAttachPkcs7BundleNil() {
	o.AttachPkcs7Bundle.Set(nil)
}

// UnsetAttachPkcs7Bundle ensures that no value is present for AttachPkcs7Bundle, not even an explicit nil
func (o *EmailNotificationResponse) UnsetAttachPkcs7Bundle() {
	o.AttachPkcs7Bundle.Unset()
}

// GetAttachPkcs12 returns the AttachPkcs12 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailNotificationResponse) GetAttachPkcs12() bool {
	if o == nil || utils.IsNil(o.AttachPkcs12.Get()) {
		var ret bool
		return ret
	}
	return *o.AttachPkcs12.Get()
}

// GetAttachPkcs12Ok returns a tuple with the AttachPkcs12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailNotificationResponse) GetAttachPkcs12Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachPkcs12.Get(), o.AttachPkcs12.IsSet()
}

// HasAttachPkcs12 returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasAttachPkcs12() bool {
	if o != nil && o.AttachPkcs12.IsSet() {
		return true
	}

	return false
}

// SetAttachPkcs12 gets a reference to the given NullableBool and assigns it to the AttachPkcs12 field.
func (o *EmailNotificationResponse) SetAttachPkcs12(v bool) {
	o.AttachPkcs12.Set(&v)
}

// SetAttachPkcs12Nil sets the value for AttachPkcs12 to be an explicit nil
func (o *EmailNotificationResponse) SetAttachPkcs12Nil() {
	o.AttachPkcs12.Set(nil)
}

// UnsetAttachPkcs12 ensures that no value is present for AttachPkcs12, not even an explicit nil
func (o *EmailNotificationResponse) UnsetAttachPkcs12() {
	o.AttachPkcs12.Unset()
}

// GetEmailTemplate returns the EmailTemplate field value
func (o *EmailNotificationResponse) GetEmailTemplate() EmailTemplate {
	if o == nil {
		var ret EmailTemplate
		return ret
	}

	return o.EmailTemplate
}

// GetEmailTemplateOk returns a tuple with the EmailTemplate field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationResponse) GetEmailTemplateOk() (*EmailTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailTemplate, true
}

// SetEmailTemplate sets field value
func (o *EmailNotificationResponse) SetEmailTemplate(v EmailTemplate) {
	o.EmailTemplate = v
}

// GetIfPkcs12 returns the IfPkcs12 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailNotificationResponse) GetIfPkcs12() bool {
	if o == nil || utils.IsNil(o.IfPkcs12.Get()) {
		var ret bool
		return ret
	}
	return *o.IfPkcs12.Get()
}

// GetIfPkcs12Ok returns a tuple with the IfPkcs12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailNotificationResponse) GetIfPkcs12Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IfPkcs12.Get(), o.IfPkcs12.IsSet()
}

// HasIfPkcs12 returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasIfPkcs12() bool {
	if o != nil && o.IfPkcs12.IsSet() {
		return true
	}

	return false
}

// SetIfPkcs12 gets a reference to the given NullableBool and assigns it to the IfPkcs12 field.
func (o *EmailNotificationResponse) SetIfPkcs12(v bool) {
	o.IfPkcs12.Set(&v)
}

// SetIfPkcs12Nil sets the value for IfPkcs12 to be an explicit nil
func (o *EmailNotificationResponse) SetIfPkcs12Nil() {
	o.IfPkcs12.Set(nil)
}

// UnsetIfPkcs12 ensures that no value is present for IfPkcs12, not even an explicit nil
func (o *EmailNotificationResponse) UnsetIfPkcs12() {
	o.IfPkcs12.Unset()
}

// GetType returns the Type field value
func (o *EmailNotificationResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EmailNotificationResponse) SetType(v string) {
	o.Type = v
}

// GetEvents returns the Events field value
func (o *EmailNotificationResponse) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationResponse) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *EmailNotificationResponse) SetEvents(v []string) {
	o.Events = v
}

// GetLicenseUsagePercent returns the LicenseUsagePercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailNotificationResponse) GetLicenseUsagePercent() int64 {
	if o == nil || utils.IsNil(o.LicenseUsagePercent.Get()) {
		var ret int64
		return ret
	}
	return *o.LicenseUsagePercent.Get()
}

// GetLicenseUsagePercentOk returns a tuple with the LicenseUsagePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailNotificationResponse) GetLicenseUsagePercentOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LicenseUsagePercent.Get(), o.LicenseUsagePercent.IsSet()
}

// HasLicenseUsagePercent returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasLicenseUsagePercent() bool {
	if o != nil && o.LicenseUsagePercent.IsSet() {
		return true
	}

	return false
}

// SetLicenseUsagePercent gets a reference to the given NullableInt64 and assigns it to the LicenseUsagePercent field.
func (o *EmailNotificationResponse) SetLicenseUsagePercent(v int64) {
	o.LicenseUsagePercent.Set(&v)
}

// SetLicenseUsagePercentNil sets the value for LicenseUsagePercent to be an explicit nil
func (o *EmailNotificationResponse) SetLicenseUsagePercentNil() {
	o.LicenseUsagePercent.Set(nil)
}

// UnsetLicenseUsagePercent ensures that no value is present for LicenseUsagePercent, not even an explicit nil
func (o *EmailNotificationResponse) UnsetLicenseUsagePercent() {
	o.LicenseUsagePercent.Unset()
}

// GetName returns the Name field value
func (o *EmailNotificationResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EmailNotificationResponse) SetName(v string) {
	o.Name = v
}

// GetRetries returns the Retries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailNotificationResponse) GetRetries() int64 {
	if o == nil || utils.IsNil(o.Retries.Get()) {
		var ret int64
		return ret
	}
	return *o.Retries.Get()
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailNotificationResponse) GetRetriesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Retries.Get(), o.Retries.IsSet()
}

// HasRetries returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasRetries() bool {
	if o != nil && o.Retries.IsSet() {
		return true
	}

	return false
}

// SetRetries gets a reference to the given NullableInt64 and assigns it to the Retries field.
func (o *EmailNotificationResponse) SetRetries(v int64) {
	o.Retries.Set(&v)
}

// SetRetriesNil sets the value for Retries to be an explicit nil
func (o *EmailNotificationResponse) SetRetriesNil() {
	o.Retries.Set(nil)
}

// UnsetRetries ensures that no value is present for Retries, not even an explicit nil
func (o *EmailNotificationResponse) UnsetRetries() {
	o.Retries.Unset()
}

// GetRunOnRenewed returns the RunOnRenewed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailNotificationResponse) GetRunOnRenewed() bool {
	if o == nil || utils.IsNil(o.RunOnRenewed.Get()) {
		var ret bool
		return ret
	}
	return *o.RunOnRenewed.Get()
}

// GetRunOnRenewedOk returns a tuple with the RunOnRenewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailNotificationResponse) GetRunOnRenewedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunOnRenewed.Get(), o.RunOnRenewed.IsSet()
}

// HasRunOnRenewed returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasRunOnRenewed() bool {
	if o != nil && o.RunOnRenewed.IsSet() {
		return true
	}

	return false
}

// SetRunOnRenewed gets a reference to the given NullableBool and assigns it to the RunOnRenewed field.
func (o *EmailNotificationResponse) SetRunOnRenewed(v bool) {
	o.RunOnRenewed.Set(&v)
}

// SetRunOnRenewedNil sets the value for RunOnRenewed to be an explicit nil
func (o *EmailNotificationResponse) SetRunOnRenewedNil() {
	o.RunOnRenewed.Set(nil)
}

// UnsetRunOnRenewed ensures that no value is present for RunOnRenewed, not even an explicit nil
func (o *EmailNotificationResponse) UnsetRunOnRenewed() {
	o.RunOnRenewed.Unset()
}

// GetRunPeriod returns the RunPeriod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailNotificationResponse) GetRunPeriod() string {
	if o == nil || utils.IsNil(o.RunPeriod.Get()) {
		var ret string
		return ret
	}
	return *o.RunPeriod.Get()
}

// GetRunPeriodOk returns a tuple with the RunPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailNotificationResponse) GetRunPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunPeriod.Get(), o.RunPeriod.IsSet()
}

// HasRunPeriod returns a boolean if a field has been set.
func (o *EmailNotificationResponse) HasRunPeriod() bool {
	if o != nil && o.RunPeriod.IsSet() {
		return true
	}

	return false
}

// SetRunPeriod gets a reference to the given NullableString and assigns it to the RunPeriod field.
func (o *EmailNotificationResponse) SetRunPeriod(v string) {
	o.RunPeriod.Set(&v)
}

// SetRunPeriodNil sets the value for RunPeriod to be an explicit nil
func (o *EmailNotificationResponse) SetRunPeriodNil() {
	o.RunPeriod.Set(nil)
}

// UnsetRunPeriod ensures that no value is present for RunPeriod, not even an explicit nil
func (o *EmailNotificationResponse) UnsetRunPeriod() {
	o.RunPeriod.Unset()
}

func (o EmailNotificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.AttachDerCertificate.IsSet() {
		toSerialize["attachDerCertificate"] = o.AttachDerCertificate.Get()
	}
	if o.AttachPemBundle.IsSet() {
		toSerialize["attachPemBundle"] = o.AttachPemBundle.Get()
	}
	if o.AttachPemCertificate.IsSet() {
		toSerialize["attachPemCertificate"] = o.AttachPemCertificate.Get()
	}
	if o.AttachPkcs7.IsSet() {
		toSerialize["attachPkcs7"] = o.AttachPkcs7.Get()
	}
	if o.AttachPkcs7Bundle.IsSet() {
		toSerialize["attachPkcs7Bundle"] = o.AttachPkcs7Bundle.Get()
	}
	if o.AttachPkcs12.IsSet() {
		toSerialize["attachPkcs12"] = o.AttachPkcs12.Get()
	}
	toSerialize["emailTemplate"] = o.EmailTemplate
	if o.IfPkcs12.IsSet() {
		toSerialize["ifPkcs12"] = o.IfPkcs12.Get()
	}
	toSerialize["type"] = o.Type
	toSerialize["events"] = o.Events
	if o.LicenseUsagePercent.IsSet() {
		toSerialize["licenseUsagePercent"] = o.LicenseUsagePercent.Get()
	}
	toSerialize["name"] = o.Name
	if o.Retries.IsSet() {
		toSerialize["retries"] = o.Retries.Get()
	}
	if o.RunOnRenewed.IsSet() {
		toSerialize["runOnRenewed"] = o.RunOnRenewed.Get()
	}
	if o.RunPeriod.IsSet() {
		toSerialize["runPeriod"] = o.RunPeriod.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmailNotificationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"emailTemplate",
		"type",
		"events",
		"name",
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

	varEmailNotificationResponse := _EmailNotificationResponse{}

	err = json.Unmarshal(data, &varEmailNotificationResponse)

	if err != nil {
		return err
	}

	*o = EmailNotificationResponse(varEmailNotificationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "attachDerCertificate")
		delete(additionalProperties, "attachPemBundle")
		delete(additionalProperties, "attachPemCertificate")
		delete(additionalProperties, "attachPkcs7")
		delete(additionalProperties, "attachPkcs7Bundle")
		delete(additionalProperties, "attachPkcs12")
		delete(additionalProperties, "emailTemplate")
		delete(additionalProperties, "ifPkcs12")
		delete(additionalProperties, "type")
		delete(additionalProperties, "events")
		delete(additionalProperties, "licenseUsagePercent")
		delete(additionalProperties, "name")
		delete(additionalProperties, "retries")
		delete(additionalProperties, "runOnRenewed")
		delete(additionalProperties, "runPeriod")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmailNotificationResponse struct {
	value *EmailNotificationResponse
	isSet bool
}

func (v NullableEmailNotificationResponse) Get() *EmailNotificationResponse {
	return v.value
}

func (v *NullableEmailNotificationResponse) Set(val *EmailNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotificationResponse(val *EmailNotificationResponse) *NullableEmailNotificationResponse {
	return &NullableEmailNotificationResponse{value: val, isSet: true}
}

func (v NullableEmailNotificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
