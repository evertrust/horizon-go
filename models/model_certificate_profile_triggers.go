/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the CertificateProfileTriggers type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CertificateProfileTriggers{}

// CertificateProfileTriggers struct for CertificateProfileTriggers
type CertificateProfileTriggers struct {
	OnEnroll             []string                                `json:"onEnroll,omitempty"`
	OnSubmitEnroll       []string                                `json:"onSubmitEnroll,omitempty"`
	OnCancelEnroll       []string                                `json:"onCancelEnroll,omitempty"`
	OnApproveEnroll      []string                                `json:"onApproveEnroll,omitempty"`
	OnDenyEnroll         []string                                `json:"onDenyEnroll,omitempty"`
	OnPendingEnroll      []CertificateProfileAsynchronousTrigger `json:"onPendingEnroll,omitempty"`
	OnRevoke             []string                                `json:"onRevoke,omitempty"`
	OnSubmitRevoke       []string                                `json:"onSubmitRevoke,omitempty"`
	OnCancelRevoke       []string                                `json:"onCancelRevoke,omitempty"`
	OnApproveRevoke      []string                                `json:"onApproveRevoke,omitempty"`
	OnDenyRevoke         []string                                `json:"onDenyRevoke,omitempty"`
	OnPendingRevoke      []CertificateProfileAsynchronousTrigger `json:"onPendingRevoke,omitempty"`
	OnUpdate             []string                                `json:"onUpdate,omitempty"`
	OnSubmitUpdate       []string                                `json:"onSubmitUpdate,omitempty"`
	OnCancelUpdate       []string                                `json:"onCancelUpdate,omitempty"`
	OnApproveUpdate      []string                                `json:"onApproveUpdate,omitempty"`
	OnDenyUpdate         []string                                `json:"onDenyUpdate,omitempty"`
	OnPendingUpdate      []CertificateProfileAsynchronousTrigger `json:"onPendingUpdate,omitempty"`
	OnRecover            []string                                `json:"onRecover,omitempty"`
	OnSubmitRecover      []string                                `json:"onSubmitRecover,omitempty"`
	OnCancelRecover      []string                                `json:"onCancelRecover,omitempty"`
	OnApproveRecover     []string                                `json:"onApproveRecover,omitempty"`
	OnDenyRecover        []string                                `json:"onDenyRecover,omitempty"`
	OnPendingRecover     []CertificateProfileAsynchronousTrigger `json:"onPendingRecover,omitempty"`
	OnMigrate            []string                                `json:"onMigrate,omitempty"`
	OnSubmitMigrate      []string                                `json:"onSubmitMigrate,omitempty"`
	OnCancelMigrate      []string                                `json:"onCancelMigrate,omitempty"`
	OnApproveMigrate     []string                                `json:"onApproveMigrate,omitempty"`
	OnDenyMigrate        []string                                `json:"onDenyMigrate,omitempty"`
	OnPendingMigrate     []CertificateProfileAsynchronousTrigger `json:"onPendingMigrate,omitempty"`
	OnExpire             []CertificateProfileAsynchronousTrigger `json:"onExpire,omitempty"`
	OnRenew              []string                                `json:"onRenew,omitempty"`
	OnSubmitRenew        []string                                `json:"onSubmitRenew,omitempty"`
	OnCancelRenew        []string                                `json:"onCancelRenew,omitempty"`
	OnApproveRenew       []string                                `json:"onApproveRenew,omitempty"`
	OnDenyRenew          []string                                `json:"onDenyRenew,omitempty"`
	OnPendingRenew       []CertificateProfileAsynchronousTrigger `json:"onPendingRenew,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateProfileTriggers CertificateProfileTriggers

// NewCertificateProfileTriggers instantiates a new CertificateProfileTriggers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateProfileTriggers() *CertificateProfileTriggers {
	this := CertificateProfileTriggers{}
	return &this
}

// NewCertificateProfileTriggersWithDefaults instantiates a new CertificateProfileTriggers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateProfileTriggersWithDefaults() *CertificateProfileTriggers {
	this := CertificateProfileTriggers{}
	return &this
}

// GetOnEnroll returns the OnEnroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnEnroll() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnEnroll
}

// GetOnEnrollOk returns a tuple with the OnEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnEnrollOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnEnroll) {
		return nil, false
	}
	return o.OnEnroll, true
}

// HasOnEnroll returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnEnroll() bool {
	if o != nil && !utils.IsNil(o.OnEnroll) {
		return true
	}

	return false
}

// SetOnEnroll gets a reference to the given []string and assigns it to the OnEnroll field.
func (o *CertificateProfileTriggers) SetOnEnroll(v []string) {
	o.OnEnroll = v
}

// GetOnSubmitEnroll returns the OnSubmitEnroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnSubmitEnroll() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnSubmitEnroll
}

// GetOnSubmitEnrollOk returns a tuple with the OnSubmitEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnSubmitEnrollOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnSubmitEnroll) {
		return nil, false
	}
	return o.OnSubmitEnroll, true
}

// HasOnSubmitEnroll returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnSubmitEnroll() bool {
	if o != nil && !utils.IsNil(o.OnSubmitEnroll) {
		return true
	}

	return false
}

// SetOnSubmitEnroll gets a reference to the given []string and assigns it to the OnSubmitEnroll field.
func (o *CertificateProfileTriggers) SetOnSubmitEnroll(v []string) {
	o.OnSubmitEnroll = v
}

// GetOnCancelEnroll returns the OnCancelEnroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnCancelEnroll() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnCancelEnroll
}

// GetOnCancelEnrollOk returns a tuple with the OnCancelEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnCancelEnrollOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnCancelEnroll) {
		return nil, false
	}
	return o.OnCancelEnroll, true
}

// HasOnCancelEnroll returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnCancelEnroll() bool {
	if o != nil && !utils.IsNil(o.OnCancelEnroll) {
		return true
	}

	return false
}

// SetOnCancelEnroll gets a reference to the given []string and assigns it to the OnCancelEnroll field.
func (o *CertificateProfileTriggers) SetOnCancelEnroll(v []string) {
	o.OnCancelEnroll = v
}

// GetOnApproveEnroll returns the OnApproveEnroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnApproveEnroll() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnApproveEnroll
}

// GetOnApproveEnrollOk returns a tuple with the OnApproveEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnApproveEnrollOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnApproveEnroll) {
		return nil, false
	}
	return o.OnApproveEnroll, true
}

// HasOnApproveEnroll returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnApproveEnroll() bool {
	if o != nil && !utils.IsNil(o.OnApproveEnroll) {
		return true
	}

	return false
}

// SetOnApproveEnroll gets a reference to the given []string and assigns it to the OnApproveEnroll field.
func (o *CertificateProfileTriggers) SetOnApproveEnroll(v []string) {
	o.OnApproveEnroll = v
}

// GetOnDenyEnroll returns the OnDenyEnroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnDenyEnroll() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnDenyEnroll
}

// GetOnDenyEnrollOk returns a tuple with the OnDenyEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnDenyEnrollOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnDenyEnroll) {
		return nil, false
	}
	return o.OnDenyEnroll, true
}

// HasOnDenyEnroll returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnDenyEnroll() bool {
	if o != nil && !utils.IsNil(o.OnDenyEnroll) {
		return true
	}

	return false
}

// SetOnDenyEnroll gets a reference to the given []string and assigns it to the OnDenyEnroll field.
func (o *CertificateProfileTriggers) SetOnDenyEnroll(v []string) {
	o.OnDenyEnroll = v
}

// GetOnPendingEnroll returns the OnPendingEnroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnPendingEnroll() []CertificateProfileAsynchronousTrigger {
	if o == nil {
		var ret []CertificateProfileAsynchronousTrigger
		return ret
	}
	return o.OnPendingEnroll
}

// GetOnPendingEnrollOk returns a tuple with the OnPendingEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnPendingEnrollOk() ([]CertificateProfileAsynchronousTrigger, bool) {
	if o == nil || utils.IsNil(o.OnPendingEnroll) {
		return nil, false
	}
	return o.OnPendingEnroll, true
}

// HasOnPendingEnroll returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnPendingEnroll() bool {
	if o != nil && !utils.IsNil(o.OnPendingEnroll) {
		return true
	}

	return false
}

// SetOnPendingEnroll gets a reference to the given []CertificateProfileAsynchronousTrigger and assigns it to the OnPendingEnroll field.
func (o *CertificateProfileTriggers) SetOnPendingEnroll(v []CertificateProfileAsynchronousTrigger) {
	o.OnPendingEnroll = v
}

// GetOnRevoke returns the OnRevoke field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnRevoke() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnRevoke
}

// GetOnRevokeOk returns a tuple with the OnRevoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnRevokeOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnRevoke) {
		return nil, false
	}
	return o.OnRevoke, true
}

// HasOnRevoke returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnRevoke() bool {
	if o != nil && !utils.IsNil(o.OnRevoke) {
		return true
	}

	return false
}

// SetOnRevoke gets a reference to the given []string and assigns it to the OnRevoke field.
func (o *CertificateProfileTriggers) SetOnRevoke(v []string) {
	o.OnRevoke = v
}

// GetOnSubmitRevoke returns the OnSubmitRevoke field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnSubmitRevoke() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnSubmitRevoke
}

// GetOnSubmitRevokeOk returns a tuple with the OnSubmitRevoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnSubmitRevokeOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnSubmitRevoke) {
		return nil, false
	}
	return o.OnSubmitRevoke, true
}

// HasOnSubmitRevoke returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnSubmitRevoke() bool {
	if o != nil && !utils.IsNil(o.OnSubmitRevoke) {
		return true
	}

	return false
}

// SetOnSubmitRevoke gets a reference to the given []string and assigns it to the OnSubmitRevoke field.
func (o *CertificateProfileTriggers) SetOnSubmitRevoke(v []string) {
	o.OnSubmitRevoke = v
}

// GetOnCancelRevoke returns the OnCancelRevoke field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnCancelRevoke() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnCancelRevoke
}

// GetOnCancelRevokeOk returns a tuple with the OnCancelRevoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnCancelRevokeOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnCancelRevoke) {
		return nil, false
	}
	return o.OnCancelRevoke, true
}

// HasOnCancelRevoke returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnCancelRevoke() bool {
	if o != nil && !utils.IsNil(o.OnCancelRevoke) {
		return true
	}

	return false
}

// SetOnCancelRevoke gets a reference to the given []string and assigns it to the OnCancelRevoke field.
func (o *CertificateProfileTriggers) SetOnCancelRevoke(v []string) {
	o.OnCancelRevoke = v
}

// GetOnApproveRevoke returns the OnApproveRevoke field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnApproveRevoke() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnApproveRevoke
}

// GetOnApproveRevokeOk returns a tuple with the OnApproveRevoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnApproveRevokeOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnApproveRevoke) {
		return nil, false
	}
	return o.OnApproveRevoke, true
}

// HasOnApproveRevoke returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnApproveRevoke() bool {
	if o != nil && !utils.IsNil(o.OnApproveRevoke) {
		return true
	}

	return false
}

// SetOnApproveRevoke gets a reference to the given []string and assigns it to the OnApproveRevoke field.
func (o *CertificateProfileTriggers) SetOnApproveRevoke(v []string) {
	o.OnApproveRevoke = v
}

// GetOnDenyRevoke returns the OnDenyRevoke field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnDenyRevoke() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnDenyRevoke
}

// GetOnDenyRevokeOk returns a tuple with the OnDenyRevoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnDenyRevokeOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnDenyRevoke) {
		return nil, false
	}
	return o.OnDenyRevoke, true
}

// HasOnDenyRevoke returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnDenyRevoke() bool {
	if o != nil && !utils.IsNil(o.OnDenyRevoke) {
		return true
	}

	return false
}

// SetOnDenyRevoke gets a reference to the given []string and assigns it to the OnDenyRevoke field.
func (o *CertificateProfileTriggers) SetOnDenyRevoke(v []string) {
	o.OnDenyRevoke = v
}

// GetOnPendingRevoke returns the OnPendingRevoke field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnPendingRevoke() []CertificateProfileAsynchronousTrigger {
	if o == nil {
		var ret []CertificateProfileAsynchronousTrigger
		return ret
	}
	return o.OnPendingRevoke
}

// GetOnPendingRevokeOk returns a tuple with the OnPendingRevoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnPendingRevokeOk() ([]CertificateProfileAsynchronousTrigger, bool) {
	if o == nil || utils.IsNil(o.OnPendingRevoke) {
		return nil, false
	}
	return o.OnPendingRevoke, true
}

// HasOnPendingRevoke returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnPendingRevoke() bool {
	if o != nil && !utils.IsNil(o.OnPendingRevoke) {
		return true
	}

	return false
}

// SetOnPendingRevoke gets a reference to the given []CertificateProfileAsynchronousTrigger and assigns it to the OnPendingRevoke field.
func (o *CertificateProfileTriggers) SetOnPendingRevoke(v []CertificateProfileAsynchronousTrigger) {
	o.OnPendingRevoke = v
}

// GetOnUpdate returns the OnUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnUpdate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnUpdate
}

// GetOnUpdateOk returns a tuple with the OnUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnUpdateOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnUpdate) {
		return nil, false
	}
	return o.OnUpdate, true
}

// HasOnUpdate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnUpdate() bool {
	if o != nil && !utils.IsNil(o.OnUpdate) {
		return true
	}

	return false
}

// SetOnUpdate gets a reference to the given []string and assigns it to the OnUpdate field.
func (o *CertificateProfileTriggers) SetOnUpdate(v []string) {
	o.OnUpdate = v
}

// GetOnSubmitUpdate returns the OnSubmitUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnSubmitUpdate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnSubmitUpdate
}

// GetOnSubmitUpdateOk returns a tuple with the OnSubmitUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnSubmitUpdateOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnSubmitUpdate) {
		return nil, false
	}
	return o.OnSubmitUpdate, true
}

// HasOnSubmitUpdate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnSubmitUpdate() bool {
	if o != nil && !utils.IsNil(o.OnSubmitUpdate) {
		return true
	}

	return false
}

// SetOnSubmitUpdate gets a reference to the given []string and assigns it to the OnSubmitUpdate field.
func (o *CertificateProfileTriggers) SetOnSubmitUpdate(v []string) {
	o.OnSubmitUpdate = v
}

// GetOnCancelUpdate returns the OnCancelUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnCancelUpdate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnCancelUpdate
}

// GetOnCancelUpdateOk returns a tuple with the OnCancelUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnCancelUpdateOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnCancelUpdate) {
		return nil, false
	}
	return o.OnCancelUpdate, true
}

// HasOnCancelUpdate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnCancelUpdate() bool {
	if o != nil && !utils.IsNil(o.OnCancelUpdate) {
		return true
	}

	return false
}

// SetOnCancelUpdate gets a reference to the given []string and assigns it to the OnCancelUpdate field.
func (o *CertificateProfileTriggers) SetOnCancelUpdate(v []string) {
	o.OnCancelUpdate = v
}

// GetOnApproveUpdate returns the OnApproveUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnApproveUpdate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnApproveUpdate
}

// GetOnApproveUpdateOk returns a tuple with the OnApproveUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnApproveUpdateOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnApproveUpdate) {
		return nil, false
	}
	return o.OnApproveUpdate, true
}

// HasOnApproveUpdate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnApproveUpdate() bool {
	if o != nil && !utils.IsNil(o.OnApproveUpdate) {
		return true
	}

	return false
}

// SetOnApproveUpdate gets a reference to the given []string and assigns it to the OnApproveUpdate field.
func (o *CertificateProfileTriggers) SetOnApproveUpdate(v []string) {
	o.OnApproveUpdate = v
}

// GetOnDenyUpdate returns the OnDenyUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnDenyUpdate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnDenyUpdate
}

// GetOnDenyUpdateOk returns a tuple with the OnDenyUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnDenyUpdateOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnDenyUpdate) {
		return nil, false
	}
	return o.OnDenyUpdate, true
}

// HasOnDenyUpdate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnDenyUpdate() bool {
	if o != nil && !utils.IsNil(o.OnDenyUpdate) {
		return true
	}

	return false
}

// SetOnDenyUpdate gets a reference to the given []string and assigns it to the OnDenyUpdate field.
func (o *CertificateProfileTriggers) SetOnDenyUpdate(v []string) {
	o.OnDenyUpdate = v
}

// GetOnPendingUpdate returns the OnPendingUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnPendingUpdate() []CertificateProfileAsynchronousTrigger {
	if o == nil {
		var ret []CertificateProfileAsynchronousTrigger
		return ret
	}
	return o.OnPendingUpdate
}

// GetOnPendingUpdateOk returns a tuple with the OnPendingUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnPendingUpdateOk() ([]CertificateProfileAsynchronousTrigger, bool) {
	if o == nil || utils.IsNil(o.OnPendingUpdate) {
		return nil, false
	}
	return o.OnPendingUpdate, true
}

// HasOnPendingUpdate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnPendingUpdate() bool {
	if o != nil && !utils.IsNil(o.OnPendingUpdate) {
		return true
	}

	return false
}

// SetOnPendingUpdate gets a reference to the given []CertificateProfileAsynchronousTrigger and assigns it to the OnPendingUpdate field.
func (o *CertificateProfileTriggers) SetOnPendingUpdate(v []CertificateProfileAsynchronousTrigger) {
	o.OnPendingUpdate = v
}

// GetOnRecover returns the OnRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnRecover() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnRecover
}

// GetOnRecoverOk returns a tuple with the OnRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnRecoverOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnRecover) {
		return nil, false
	}
	return o.OnRecover, true
}

// HasOnRecover returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnRecover() bool {
	if o != nil && !utils.IsNil(o.OnRecover) {
		return true
	}

	return false
}

// SetOnRecover gets a reference to the given []string and assigns it to the OnRecover field.
func (o *CertificateProfileTriggers) SetOnRecover(v []string) {
	o.OnRecover = v
}

// GetOnSubmitRecover returns the OnSubmitRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnSubmitRecover() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnSubmitRecover
}

// GetOnSubmitRecoverOk returns a tuple with the OnSubmitRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnSubmitRecoverOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnSubmitRecover) {
		return nil, false
	}
	return o.OnSubmitRecover, true
}

// HasOnSubmitRecover returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnSubmitRecover() bool {
	if o != nil && !utils.IsNil(o.OnSubmitRecover) {
		return true
	}

	return false
}

// SetOnSubmitRecover gets a reference to the given []string and assigns it to the OnSubmitRecover field.
func (o *CertificateProfileTriggers) SetOnSubmitRecover(v []string) {
	o.OnSubmitRecover = v
}

// GetOnCancelRecover returns the OnCancelRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnCancelRecover() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnCancelRecover
}

// GetOnCancelRecoverOk returns a tuple with the OnCancelRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnCancelRecoverOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnCancelRecover) {
		return nil, false
	}
	return o.OnCancelRecover, true
}

// HasOnCancelRecover returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnCancelRecover() bool {
	if o != nil && !utils.IsNil(o.OnCancelRecover) {
		return true
	}

	return false
}

// SetOnCancelRecover gets a reference to the given []string and assigns it to the OnCancelRecover field.
func (o *CertificateProfileTriggers) SetOnCancelRecover(v []string) {
	o.OnCancelRecover = v
}

// GetOnApproveRecover returns the OnApproveRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnApproveRecover() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnApproveRecover
}

// GetOnApproveRecoverOk returns a tuple with the OnApproveRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnApproveRecoverOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnApproveRecover) {
		return nil, false
	}
	return o.OnApproveRecover, true
}

// HasOnApproveRecover returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnApproveRecover() bool {
	if o != nil && !utils.IsNil(o.OnApproveRecover) {
		return true
	}

	return false
}

// SetOnApproveRecover gets a reference to the given []string and assigns it to the OnApproveRecover field.
func (o *CertificateProfileTriggers) SetOnApproveRecover(v []string) {
	o.OnApproveRecover = v
}

// GetOnDenyRecover returns the OnDenyRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnDenyRecover() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnDenyRecover
}

// GetOnDenyRecoverOk returns a tuple with the OnDenyRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnDenyRecoverOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnDenyRecover) {
		return nil, false
	}
	return o.OnDenyRecover, true
}

// HasOnDenyRecover returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnDenyRecover() bool {
	if o != nil && !utils.IsNil(o.OnDenyRecover) {
		return true
	}

	return false
}

// SetOnDenyRecover gets a reference to the given []string and assigns it to the OnDenyRecover field.
func (o *CertificateProfileTriggers) SetOnDenyRecover(v []string) {
	o.OnDenyRecover = v
}

// GetOnPendingRecover returns the OnPendingRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnPendingRecover() []CertificateProfileAsynchronousTrigger {
	if o == nil {
		var ret []CertificateProfileAsynchronousTrigger
		return ret
	}
	return o.OnPendingRecover
}

// GetOnPendingRecoverOk returns a tuple with the OnPendingRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnPendingRecoverOk() ([]CertificateProfileAsynchronousTrigger, bool) {
	if o == nil || utils.IsNil(o.OnPendingRecover) {
		return nil, false
	}
	return o.OnPendingRecover, true
}

// HasOnPendingRecover returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnPendingRecover() bool {
	if o != nil && !utils.IsNil(o.OnPendingRecover) {
		return true
	}

	return false
}

// SetOnPendingRecover gets a reference to the given []CertificateProfileAsynchronousTrigger and assigns it to the OnPendingRecover field.
func (o *CertificateProfileTriggers) SetOnPendingRecover(v []CertificateProfileAsynchronousTrigger) {
	o.OnPendingRecover = v
}

// GetOnMigrate returns the OnMigrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnMigrate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnMigrate
}

// GetOnMigrateOk returns a tuple with the OnMigrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnMigrateOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnMigrate) {
		return nil, false
	}
	return o.OnMigrate, true
}

// HasOnMigrate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnMigrate() bool {
	if o != nil && !utils.IsNil(o.OnMigrate) {
		return true
	}

	return false
}

// SetOnMigrate gets a reference to the given []string and assigns it to the OnMigrate field.
func (o *CertificateProfileTriggers) SetOnMigrate(v []string) {
	o.OnMigrate = v
}

// GetOnSubmitMigrate returns the OnSubmitMigrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnSubmitMigrate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnSubmitMigrate
}

// GetOnSubmitMigrateOk returns a tuple with the OnSubmitMigrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnSubmitMigrateOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnSubmitMigrate) {
		return nil, false
	}
	return o.OnSubmitMigrate, true
}

// HasOnSubmitMigrate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnSubmitMigrate() bool {
	if o != nil && !utils.IsNil(o.OnSubmitMigrate) {
		return true
	}

	return false
}

// SetOnSubmitMigrate gets a reference to the given []string and assigns it to the OnSubmitMigrate field.
func (o *CertificateProfileTriggers) SetOnSubmitMigrate(v []string) {
	o.OnSubmitMigrate = v
}

// GetOnCancelMigrate returns the OnCancelMigrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnCancelMigrate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnCancelMigrate
}

// GetOnCancelMigrateOk returns a tuple with the OnCancelMigrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnCancelMigrateOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnCancelMigrate) {
		return nil, false
	}
	return o.OnCancelMigrate, true
}

// HasOnCancelMigrate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnCancelMigrate() bool {
	if o != nil && !utils.IsNil(o.OnCancelMigrate) {
		return true
	}

	return false
}

// SetOnCancelMigrate gets a reference to the given []string and assigns it to the OnCancelMigrate field.
func (o *CertificateProfileTriggers) SetOnCancelMigrate(v []string) {
	o.OnCancelMigrate = v
}

// GetOnApproveMigrate returns the OnApproveMigrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnApproveMigrate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnApproveMigrate
}

// GetOnApproveMigrateOk returns a tuple with the OnApproveMigrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnApproveMigrateOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnApproveMigrate) {
		return nil, false
	}
	return o.OnApproveMigrate, true
}

// HasOnApproveMigrate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnApproveMigrate() bool {
	if o != nil && !utils.IsNil(o.OnApproveMigrate) {
		return true
	}

	return false
}

// SetOnApproveMigrate gets a reference to the given []string and assigns it to the OnApproveMigrate field.
func (o *CertificateProfileTriggers) SetOnApproveMigrate(v []string) {
	o.OnApproveMigrate = v
}

// GetOnDenyMigrate returns the OnDenyMigrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnDenyMigrate() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnDenyMigrate
}

// GetOnDenyMigrateOk returns a tuple with the OnDenyMigrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnDenyMigrateOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnDenyMigrate) {
		return nil, false
	}
	return o.OnDenyMigrate, true
}

// HasOnDenyMigrate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnDenyMigrate() bool {
	if o != nil && !utils.IsNil(o.OnDenyMigrate) {
		return true
	}

	return false
}

// SetOnDenyMigrate gets a reference to the given []string and assigns it to the OnDenyMigrate field.
func (o *CertificateProfileTriggers) SetOnDenyMigrate(v []string) {
	o.OnDenyMigrate = v
}

// GetOnPendingMigrate returns the OnPendingMigrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnPendingMigrate() []CertificateProfileAsynchronousTrigger {
	if o == nil {
		var ret []CertificateProfileAsynchronousTrigger
		return ret
	}
	return o.OnPendingMigrate
}

// GetOnPendingMigrateOk returns a tuple with the OnPendingMigrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnPendingMigrateOk() ([]CertificateProfileAsynchronousTrigger, bool) {
	if o == nil || utils.IsNil(o.OnPendingMigrate) {
		return nil, false
	}
	return o.OnPendingMigrate, true
}

// HasOnPendingMigrate returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnPendingMigrate() bool {
	if o != nil && !utils.IsNil(o.OnPendingMigrate) {
		return true
	}

	return false
}

// SetOnPendingMigrate gets a reference to the given []CertificateProfileAsynchronousTrigger and assigns it to the OnPendingMigrate field.
func (o *CertificateProfileTriggers) SetOnPendingMigrate(v []CertificateProfileAsynchronousTrigger) {
	o.OnPendingMigrate = v
}

// GetOnExpire returns the OnExpire field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnExpire() []CertificateProfileAsynchronousTrigger {
	if o == nil {
		var ret []CertificateProfileAsynchronousTrigger
		return ret
	}
	return o.OnExpire
}

// GetOnExpireOk returns a tuple with the OnExpire field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnExpireOk() ([]CertificateProfileAsynchronousTrigger, bool) {
	if o == nil || utils.IsNil(o.OnExpire) {
		return nil, false
	}
	return o.OnExpire, true
}

// HasOnExpire returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnExpire() bool {
	if o != nil && !utils.IsNil(o.OnExpire) {
		return true
	}

	return false
}

// SetOnExpire gets a reference to the given []CertificateProfileAsynchronousTrigger and assigns it to the OnExpire field.
func (o *CertificateProfileTriggers) SetOnExpire(v []CertificateProfileAsynchronousTrigger) {
	o.OnExpire = v
}

// GetOnRenew returns the OnRenew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnRenew() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnRenew
}

// GetOnRenewOk returns a tuple with the OnRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnRenewOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnRenew) {
		return nil, false
	}
	return o.OnRenew, true
}

// HasOnRenew returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnRenew() bool {
	if o != nil && !utils.IsNil(o.OnRenew) {
		return true
	}

	return false
}

// SetOnRenew gets a reference to the given []string and assigns it to the OnRenew field.
func (o *CertificateProfileTriggers) SetOnRenew(v []string) {
	o.OnRenew = v
}

// GetOnSubmitRenew returns the OnSubmitRenew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnSubmitRenew() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnSubmitRenew
}

// GetOnSubmitRenewOk returns a tuple with the OnSubmitRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnSubmitRenewOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnSubmitRenew) {
		return nil, false
	}
	return o.OnSubmitRenew, true
}

// HasOnSubmitRenew returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnSubmitRenew() bool {
	if o != nil && !utils.IsNil(o.OnSubmitRenew) {
		return true
	}

	return false
}

// SetOnSubmitRenew gets a reference to the given []string and assigns it to the OnSubmitRenew field.
func (o *CertificateProfileTriggers) SetOnSubmitRenew(v []string) {
	o.OnSubmitRenew = v
}

// GetOnCancelRenew returns the OnCancelRenew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnCancelRenew() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnCancelRenew
}

// GetOnCancelRenewOk returns a tuple with the OnCancelRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnCancelRenewOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnCancelRenew) {
		return nil, false
	}
	return o.OnCancelRenew, true
}

// HasOnCancelRenew returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnCancelRenew() bool {
	if o != nil && !utils.IsNil(o.OnCancelRenew) {
		return true
	}

	return false
}

// SetOnCancelRenew gets a reference to the given []string and assigns it to the OnCancelRenew field.
func (o *CertificateProfileTriggers) SetOnCancelRenew(v []string) {
	o.OnCancelRenew = v
}

// GetOnApproveRenew returns the OnApproveRenew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnApproveRenew() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnApproveRenew
}

// GetOnApproveRenewOk returns a tuple with the OnApproveRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnApproveRenewOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnApproveRenew) {
		return nil, false
	}
	return o.OnApproveRenew, true
}

// HasOnApproveRenew returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnApproveRenew() bool {
	if o != nil && !utils.IsNil(o.OnApproveRenew) {
		return true
	}

	return false
}

// SetOnApproveRenew gets a reference to the given []string and assigns it to the OnApproveRenew field.
func (o *CertificateProfileTriggers) SetOnApproveRenew(v []string) {
	o.OnApproveRenew = v
}

// GetOnDenyRenew returns the OnDenyRenew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnDenyRenew() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OnDenyRenew
}

// GetOnDenyRenewOk returns a tuple with the OnDenyRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnDenyRenewOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.OnDenyRenew) {
		return nil, false
	}
	return o.OnDenyRenew, true
}

// HasOnDenyRenew returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnDenyRenew() bool {
	if o != nil && !utils.IsNil(o.OnDenyRenew) {
		return true
	}

	return false
}

// SetOnDenyRenew gets a reference to the given []string and assigns it to the OnDenyRenew field.
func (o *CertificateProfileTriggers) SetOnDenyRenew(v []string) {
	o.OnDenyRenew = v
}

// GetOnPendingRenew returns the OnPendingRenew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileTriggers) GetOnPendingRenew() []CertificateProfileAsynchronousTrigger {
	if o == nil {
		var ret []CertificateProfileAsynchronousTrigger
		return ret
	}
	return o.OnPendingRenew
}

// GetOnPendingRenewOk returns a tuple with the OnPendingRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileTriggers) GetOnPendingRenewOk() ([]CertificateProfileAsynchronousTrigger, bool) {
	if o == nil || utils.IsNil(o.OnPendingRenew) {
		return nil, false
	}
	return o.OnPendingRenew, true
}

// HasOnPendingRenew returns a boolean if a field has been set.
func (o *CertificateProfileTriggers) HasOnPendingRenew() bool {
	if o != nil && !utils.IsNil(o.OnPendingRenew) {
		return true
	}

	return false
}

// SetOnPendingRenew gets a reference to the given []CertificateProfileAsynchronousTrigger and assigns it to the OnPendingRenew field.
func (o *CertificateProfileTriggers) SetOnPendingRenew(v []CertificateProfileAsynchronousTrigger) {
	o.OnPendingRenew = v
}

func (o CertificateProfileTriggers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateProfileTriggers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OnEnroll != nil {
		toSerialize["onEnroll"] = o.OnEnroll
	}
	if o.OnSubmitEnroll != nil {
		toSerialize["onSubmitEnroll"] = o.OnSubmitEnroll
	}
	if o.OnCancelEnroll != nil {
		toSerialize["onCancelEnroll"] = o.OnCancelEnroll
	}
	if o.OnApproveEnroll != nil {
		toSerialize["onApproveEnroll"] = o.OnApproveEnroll
	}
	if o.OnDenyEnroll != nil {
		toSerialize["onDenyEnroll"] = o.OnDenyEnroll
	}
	if o.OnPendingEnroll != nil {
		toSerialize["onPendingEnroll"] = o.OnPendingEnroll
	}
	if o.OnRevoke != nil {
		toSerialize["onRevoke"] = o.OnRevoke
	}
	if o.OnSubmitRevoke != nil {
		toSerialize["onSubmitRevoke"] = o.OnSubmitRevoke
	}
	if o.OnCancelRevoke != nil {
		toSerialize["onCancelRevoke"] = o.OnCancelRevoke
	}
	if o.OnApproveRevoke != nil {
		toSerialize["onApproveRevoke"] = o.OnApproveRevoke
	}
	if o.OnDenyRevoke != nil {
		toSerialize["onDenyRevoke"] = o.OnDenyRevoke
	}
	if o.OnPendingRevoke != nil {
		toSerialize["onPendingRevoke"] = o.OnPendingRevoke
	}
	if o.OnUpdate != nil {
		toSerialize["onUpdate"] = o.OnUpdate
	}
	if o.OnSubmitUpdate != nil {
		toSerialize["onSubmitUpdate"] = o.OnSubmitUpdate
	}
	if o.OnCancelUpdate != nil {
		toSerialize["onCancelUpdate"] = o.OnCancelUpdate
	}
	if o.OnApproveUpdate != nil {
		toSerialize["onApproveUpdate"] = o.OnApproveUpdate
	}
	if o.OnDenyUpdate != nil {
		toSerialize["onDenyUpdate"] = o.OnDenyUpdate
	}
	if o.OnPendingUpdate != nil {
		toSerialize["onPendingUpdate"] = o.OnPendingUpdate
	}
	if o.OnRecover != nil {
		toSerialize["onRecover"] = o.OnRecover
	}
	if o.OnSubmitRecover != nil {
		toSerialize["onSubmitRecover"] = o.OnSubmitRecover
	}
	if o.OnCancelRecover != nil {
		toSerialize["onCancelRecover"] = o.OnCancelRecover
	}
	if o.OnApproveRecover != nil {
		toSerialize["onApproveRecover"] = o.OnApproveRecover
	}
	if o.OnDenyRecover != nil {
		toSerialize["onDenyRecover"] = o.OnDenyRecover
	}
	if o.OnPendingRecover != nil {
		toSerialize["onPendingRecover"] = o.OnPendingRecover
	}
	if o.OnMigrate != nil {
		toSerialize["onMigrate"] = o.OnMigrate
	}
	if o.OnSubmitMigrate != nil {
		toSerialize["onSubmitMigrate"] = o.OnSubmitMigrate
	}
	if o.OnCancelMigrate != nil {
		toSerialize["onCancelMigrate"] = o.OnCancelMigrate
	}
	if o.OnApproveMigrate != nil {
		toSerialize["onApproveMigrate"] = o.OnApproveMigrate
	}
	if o.OnDenyMigrate != nil {
		toSerialize["onDenyMigrate"] = o.OnDenyMigrate
	}
	if o.OnPendingMigrate != nil {
		toSerialize["onPendingMigrate"] = o.OnPendingMigrate
	}
	if o.OnExpire != nil {
		toSerialize["onExpire"] = o.OnExpire
	}
	if o.OnRenew != nil {
		toSerialize["onRenew"] = o.OnRenew
	}
	if o.OnSubmitRenew != nil {
		toSerialize["onSubmitRenew"] = o.OnSubmitRenew
	}
	if o.OnCancelRenew != nil {
		toSerialize["onCancelRenew"] = o.OnCancelRenew
	}
	if o.OnApproveRenew != nil {
		toSerialize["onApproveRenew"] = o.OnApproveRenew
	}
	if o.OnDenyRenew != nil {
		toSerialize["onDenyRenew"] = o.OnDenyRenew
	}
	if o.OnPendingRenew != nil {
		toSerialize["onPendingRenew"] = o.OnPendingRenew
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateProfileTriggers) UnmarshalJSON(data []byte) (err error) {
	varCertificateProfileTriggers := _CertificateProfileTriggers{}

	err = json.Unmarshal(data, &varCertificateProfileTriggers)

	if err != nil {
		return err
	}

	*o = CertificateProfileTriggers(varCertificateProfileTriggers)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "onEnroll")
		delete(additionalProperties, "onSubmitEnroll")
		delete(additionalProperties, "onCancelEnroll")
		delete(additionalProperties, "onApproveEnroll")
		delete(additionalProperties, "onDenyEnroll")
		delete(additionalProperties, "onPendingEnroll")
		delete(additionalProperties, "onRevoke")
		delete(additionalProperties, "onSubmitRevoke")
		delete(additionalProperties, "onCancelRevoke")
		delete(additionalProperties, "onApproveRevoke")
		delete(additionalProperties, "onDenyRevoke")
		delete(additionalProperties, "onPendingRevoke")
		delete(additionalProperties, "onUpdate")
		delete(additionalProperties, "onSubmitUpdate")
		delete(additionalProperties, "onCancelUpdate")
		delete(additionalProperties, "onApproveUpdate")
		delete(additionalProperties, "onDenyUpdate")
		delete(additionalProperties, "onPendingUpdate")
		delete(additionalProperties, "onRecover")
		delete(additionalProperties, "onSubmitRecover")
		delete(additionalProperties, "onCancelRecover")
		delete(additionalProperties, "onApproveRecover")
		delete(additionalProperties, "onDenyRecover")
		delete(additionalProperties, "onPendingRecover")
		delete(additionalProperties, "onMigrate")
		delete(additionalProperties, "onSubmitMigrate")
		delete(additionalProperties, "onCancelMigrate")
		delete(additionalProperties, "onApproveMigrate")
		delete(additionalProperties, "onDenyMigrate")
		delete(additionalProperties, "onPendingMigrate")
		delete(additionalProperties, "onExpire")
		delete(additionalProperties, "onRenew")
		delete(additionalProperties, "onSubmitRenew")
		delete(additionalProperties, "onCancelRenew")
		delete(additionalProperties, "onApproveRenew")
		delete(additionalProperties, "onDenyRenew")
		delete(additionalProperties, "onPendingRenew")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateProfileTriggers struct {
	value *CertificateProfileTriggers
	isSet bool
}

func (v NullableCertificateProfileTriggers) Get() *CertificateProfileTriggers {
	return v.value
}

func (v *NullableCertificateProfileTriggers) Set(val *CertificateProfileTriggers) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateProfileTriggers) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateProfileTriggers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateProfileTriggers(val *CertificateProfileTriggers) *NullableCertificateProfileTriggers {
	return &NullableCertificateProfileTriggers{value: val, isSet: true}
}

func (v NullableCertificateProfileTriggers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateProfileTriggers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
