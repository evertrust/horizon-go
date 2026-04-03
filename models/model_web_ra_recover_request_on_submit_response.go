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

// checks if the WebRARecoverRequestOnSubmitResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebRARecoverRequestOnSubmitResponse{}

// WebRARecoverRequestOnSubmitResponse struct for WebRARecoverRequestOnSubmitResponse
type WebRARecoverRequestOnSubmitResponse struct {
	// The certificate that was recovered.
	Certificate NullableCertificate `json:"certificate,omitempty"`
	// The module of the certificate recovered.
	Module Module `json:"module"`
	// The password to decrypt the PKCS12 file.
	Password NullableSecretString `json:"password,omitempty"`
	// The generated PKCS#12 for this request. This is only available after the request has been approved.
	Pkcs12 NullableSecretString `json:"pkcs12,omitempty"`
	// What this request will do. For a recovery request, this is always `recover`
	Workflow string `json:"workflow"`
	// Object internal ID
	Id string `json:"_id"`
	// The approver's principal identifier
	Approver utils.NullableString `json:"approver,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment utils.NullableString `json:"approverComment,omitempty"`
	// The request's contact email
	Contact utils.NullableString `json:"contact,omitempty"`
	// Certificate's Distinguished Name
	Dn *string `json:"dn,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun utils.NullableBool `json:"dryRun,omitempty"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount utils.NullableInt64 `json:"globalHolderIdCount,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId string `json:"holderId"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
	// The associated profile name
	Profile string `json:"profile"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount utils.NullableInt64 `json:"profileHolderIdCount,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The requester's principal identifier
	Requester utils.NullableString `json:"requester,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment utils.NullableString `json:"requesterComment,omitempty"`
	Status           RequestStatus        `json:"status"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team utils.NullableString `json:"team,omitempty"`
	// The result of the execution of triggers on this request
	TriggerResults       []TriggerResult `json:"triggerResults,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebRARecoverRequestOnSubmitResponse WebRARecoverRequestOnSubmitResponse

// NewWebRARecoverRequestOnSubmitResponse instantiates a new WebRARecoverRequestOnSubmitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebRARecoverRequestOnSubmitResponse(module Module, workflow string, id string, holderId string, lastModificationDate int64, profile string, registrationDate int64, removeAt int64, status RequestStatus) *WebRARecoverRequestOnSubmitResponse {
	this := WebRARecoverRequestOnSubmitResponse{}
	this.Id = id
	var dryRun bool = false
	this.DryRun = *utils.NewNullableBool(&dryRun)
	this.HolderId = holderId
	this.LastModificationDate = lastModificationDate
	this.Module = module
	this.Profile = profile
	this.RegistrationDate = registrationDate
	this.RemoveAt = removeAt
	this.Status = status
	this.Workflow = workflow
	return &this
}

// NewWebRARecoverRequestOnSubmitResponseWithDefaults instantiates a new WebRARecoverRequestOnSubmitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebRARecoverRequestOnSubmitResponseWithDefaults() *WebRARecoverRequestOnSubmitResponse {
	this := WebRARecoverRequestOnSubmitResponse{}
	var dryRun bool = false
	this.DryRun = *utils.NewNullableBool(&dryRun)
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetCertificate() Certificate {
	if o == nil || utils.IsNil(o.Certificate.Get()) {
		var ret Certificate
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetCertificateOk() (*Certificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableCertificate and assigns it to the Certificate field.
func (o *WebRARecoverRequestOnSubmitResponse) SetCertificate(v Certificate) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetModule returns the Module field value
func (o *WebRARecoverRequestOnSubmitResponse) GetModule() Module {
	if o == nil {
		var ret Module
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *WebRARecoverRequestOnSubmitResponse) GetModuleOk() (*Module, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *WebRARecoverRequestOnSubmitResponse) SetModule(v Module) {
	o.Module = v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetPassword() SecretString {
	if o == nil || utils.IsNil(o.Password.Get()) {
		var ret SecretString
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetPasswordOk() (*SecretString, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableSecretString and assigns it to the Password field.
func (o *WebRARecoverRequestOnSubmitResponse) SetPassword(v SecretString) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetPassword() {
	o.Password.Unset()
}

// GetPkcs12 returns the Pkcs12 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetPkcs12() SecretString {
	if o == nil || utils.IsNil(o.Pkcs12.Get()) {
		var ret SecretString
		return ret
	}
	return *o.Pkcs12.Get()
}

// GetPkcs12Ok returns a tuple with the Pkcs12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetPkcs12Ok() (*SecretString, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pkcs12.Get(), o.Pkcs12.IsSet()
}

// HasPkcs12 returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasPkcs12() bool {
	if o != nil && o.Pkcs12.IsSet() {
		return true
	}

	return false
}

// SetPkcs12 gets a reference to the given NullableSecretString and assigns it to the Pkcs12 field.
func (o *WebRARecoverRequestOnSubmitResponse) SetPkcs12(v SecretString) {
	o.Pkcs12.Set(&v)
}

// SetPkcs12Nil sets the value for Pkcs12 to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetPkcs12Nil() {
	o.Pkcs12.Set(nil)
}

// UnsetPkcs12 ensures that no value is present for Pkcs12, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetPkcs12() {
	o.Pkcs12.Unset()
}

// GetWorkflow returns the Workflow field value
func (o *WebRARecoverRequestOnSubmitResponse) GetWorkflow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
func (o *WebRARecoverRequestOnSubmitResponse) GetWorkflowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workflow, true
}

// SetWorkflow sets field value
func (o *WebRARecoverRequestOnSubmitResponse) SetWorkflow(v string) {
	o.Workflow = v
}

// GetId returns the Id field value
func (o *WebRARecoverRequestOnSubmitResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebRARecoverRequestOnSubmitResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebRARecoverRequestOnSubmitResponse) SetId(v string) {
	o.Id = v
}

// GetApprover returns the Approver field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetApprover() string {
	if o == nil || utils.IsNil(o.Approver.Get()) {
		var ret string
		return ret
	}
	return *o.Approver.Get()
}

// GetApproverOk returns a tuple with the Approver field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetApproverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approver.Get(), o.Approver.IsSet()
}

// HasApprover returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasApprover() bool {
	if o != nil && o.Approver.IsSet() {
		return true
	}

	return false
}

// SetApprover gets a reference to the given NullableString and assigns it to the Approver field.
func (o *WebRARecoverRequestOnSubmitResponse) SetApprover(v string) {
	o.Approver.Set(&v)
}

// SetApproverNil sets the value for Approver to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetApproverNil() {
	o.Approver.Set(nil)
}

// UnsetApprover ensures that no value is present for Approver, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetApprover() {
	o.Approver.Unset()
}

// GetApproverComment returns the ApproverComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetApproverComment() string {
	if o == nil || utils.IsNil(o.ApproverComment.Get()) {
		var ret string
		return ret
	}
	return *o.ApproverComment.Get()
}

// GetApproverCommentOk returns a tuple with the ApproverComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetApproverCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproverComment.Get(), o.ApproverComment.IsSet()
}

// HasApproverComment returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasApproverComment() bool {
	if o != nil && o.ApproverComment.IsSet() {
		return true
	}

	return false
}

// SetApproverComment gets a reference to the given NullableString and assigns it to the ApproverComment field.
func (o *WebRARecoverRequestOnSubmitResponse) SetApproverComment(v string) {
	o.ApproverComment.Set(&v)
}

// SetApproverCommentNil sets the value for ApproverComment to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetApproverCommentNil() {
	o.ApproverComment.Set(nil)
}

// UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetApproverComment() {
	o.ApproverComment.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetContact() string {
	if o == nil || utils.IsNil(o.Contact.Get()) {
		var ret string
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableString and assigns it to the Contact field.
func (o *WebRARecoverRequestOnSubmitResponse) SetContact(v string) {
	o.Contact.Set(&v)
}

// SetContactNil sets the value for Contact to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetContact() {
	o.Contact.Unset()
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *WebRARecoverRequestOnSubmitResponse) GetDn() string {
	if o == nil || utils.IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRARecoverRequestOnSubmitResponse) GetDnOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasDn() bool {
	if o != nil && !utils.IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *WebRARecoverRequestOnSubmitResponse) SetDn(v string) {
	o.Dn = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetDryRun() bool {
	if o == nil || utils.IsNil(o.DryRun.Get()) {
		var ret bool
		return ret
	}
	return *o.DryRun.Get()
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DryRun.Get(), o.DryRun.IsSet()
}

// HasDryRun returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasDryRun() bool {
	if o != nil && o.DryRun.IsSet() {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given NullableBool and assigns it to the DryRun field.
func (o *WebRARecoverRequestOnSubmitResponse) SetDryRun(v bool) {
	o.DryRun.Set(&v)
}

// SetDryRunNil sets the value for DryRun to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetDryRunNil() {
	o.DryRun.Set(nil)
}

// UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetDryRun() {
	o.DryRun.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *WebRARecoverRequestOnSubmitResponse) GetExpirationDate() int64 {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		var ret int64
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRARecoverRequestOnSubmitResponse) GetExpirationDateOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasExpirationDate() bool {
	if o != nil && !utils.IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *WebRARecoverRequestOnSubmitResponse) SetExpirationDate(v int64) {
	o.ExpirationDate = &v
}

// GetGlobalHolderIdCount returns the GlobalHolderIdCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetGlobalHolderIdCount() int64 {
	if o == nil || utils.IsNil(o.GlobalHolderIdCount.Get()) {
		var ret int64
		return ret
	}
	return *o.GlobalHolderIdCount.Get()
}

// GetGlobalHolderIdCountOk returns a tuple with the GlobalHolderIdCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetGlobalHolderIdCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalHolderIdCount.Get(), o.GlobalHolderIdCount.IsSet()
}

// HasGlobalHolderIdCount returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasGlobalHolderIdCount() bool {
	if o != nil && o.GlobalHolderIdCount.IsSet() {
		return true
	}

	return false
}

// SetGlobalHolderIdCount gets a reference to the given NullableInt64 and assigns it to the GlobalHolderIdCount field.
func (o *WebRARecoverRequestOnSubmitResponse) SetGlobalHolderIdCount(v int64) {
	o.GlobalHolderIdCount.Set(&v)
}

// SetGlobalHolderIdCountNil sets the value for GlobalHolderIdCount to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetGlobalHolderIdCountNil() {
	o.GlobalHolderIdCount.Set(nil)
}

// UnsetGlobalHolderIdCount ensures that no value is present for GlobalHolderIdCount, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetGlobalHolderIdCount() {
	o.GlobalHolderIdCount.Unset()
}

// GetHolderId returns the HolderId field value
func (o *WebRARecoverRequestOnSubmitResponse) GetHolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HolderId
}

// GetHolderIdOk returns a tuple with the HolderId field value
// and a boolean to check if the value has been set.
func (o *WebRARecoverRequestOnSubmitResponse) GetHolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HolderId, true
}

// SetHolderId sets field value
func (o *WebRARecoverRequestOnSubmitResponse) SetHolderId(v string) {
	o.HolderId = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetLabels() []LabelData {
	if o == nil {
		var ret []LabelData
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetLabelsOk() ([]LabelData, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelData and assigns it to the Labels field.
func (o *WebRARecoverRequestOnSubmitResponse) SetLabels(v []LabelData) {
	o.Labels = v
}

// GetLastModificationDate returns the LastModificationDate field value
func (o *WebRARecoverRequestOnSubmitResponse) GetLastModificationDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LastModificationDate
}

// GetLastModificationDateOk returns a tuple with the LastModificationDate field value
// and a boolean to check if the value has been set.
func (o *WebRARecoverRequestOnSubmitResponse) GetLastModificationDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModificationDate, true
}

// SetLastModificationDate sets field value
func (o *WebRARecoverRequestOnSubmitResponse) SetLastModificationDate(v int64) {
	o.LastModificationDate = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetMetadata() []CertificateMetadata {
	if o == nil {
		var ret []CertificateMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetMetadataOk() ([]CertificateMetadata, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []CertificateMetadata and assigns it to the Metadata field.
func (o *WebRARecoverRequestOnSubmitResponse) SetMetadata(v []CertificateMetadata) {
	o.Metadata = v
}

// GetProfile returns the Profile field value
func (o *WebRARecoverRequestOnSubmitResponse) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *WebRARecoverRequestOnSubmitResponse) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *WebRARecoverRequestOnSubmitResponse) SetProfile(v string) {
	o.Profile = v
}

// GetProfileHolderIdCount returns the ProfileHolderIdCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetProfileHolderIdCount() int64 {
	if o == nil || utils.IsNil(o.ProfileHolderIdCount.Get()) {
		var ret int64
		return ret
	}
	return *o.ProfileHolderIdCount.Get()
}

// GetProfileHolderIdCountOk returns a tuple with the ProfileHolderIdCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetProfileHolderIdCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileHolderIdCount.Get(), o.ProfileHolderIdCount.IsSet()
}

// HasProfileHolderIdCount returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasProfileHolderIdCount() bool {
	if o != nil && o.ProfileHolderIdCount.IsSet() {
		return true
	}

	return false
}

// SetProfileHolderIdCount gets a reference to the given NullableInt64 and assigns it to the ProfileHolderIdCount field.
func (o *WebRARecoverRequestOnSubmitResponse) SetProfileHolderIdCount(v int64) {
	o.ProfileHolderIdCount.Set(&v)
}

// SetProfileHolderIdCountNil sets the value for ProfileHolderIdCount to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetProfileHolderIdCountNil() {
	o.ProfileHolderIdCount.Set(nil)
}

// UnsetProfileHolderIdCount ensures that no value is present for ProfileHolderIdCount, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetProfileHolderIdCount() {
	o.ProfileHolderIdCount.Unset()
}

// GetRegistrationDate returns the RegistrationDate field value
func (o *WebRARecoverRequestOnSubmitResponse) GetRegistrationDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RegistrationDate
}

// GetRegistrationDateOk returns a tuple with the RegistrationDate field value
// and a boolean to check if the value has been set.
func (o *WebRARecoverRequestOnSubmitResponse) GetRegistrationDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistrationDate, true
}

// SetRegistrationDate sets field value
func (o *WebRARecoverRequestOnSubmitResponse) SetRegistrationDate(v int64) {
	o.RegistrationDate = v
}

// GetRemoveAt returns the RemoveAt field value
func (o *WebRARecoverRequestOnSubmitResponse) GetRemoveAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RemoveAt
}

// GetRemoveAtOk returns a tuple with the RemoveAt field value
// and a boolean to check if the value has been set.
func (o *WebRARecoverRequestOnSubmitResponse) GetRemoveAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoveAt, true
}

// SetRemoveAt sets field value
func (o *WebRARecoverRequestOnSubmitResponse) SetRemoveAt(v int64) {
	o.RemoveAt = v
}

// GetRequester returns the Requester field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetRequester() string {
	if o == nil || utils.IsNil(o.Requester.Get()) {
		var ret string
		return ret
	}
	return *o.Requester.Get()
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetRequesterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requester.Get(), o.Requester.IsSet()
}

// HasRequester returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasRequester() bool {
	if o != nil && o.Requester.IsSet() {
		return true
	}

	return false
}

// SetRequester gets a reference to the given NullableString and assigns it to the Requester field.
func (o *WebRARecoverRequestOnSubmitResponse) SetRequester(v string) {
	o.Requester.Set(&v)
}

// SetRequesterNil sets the value for Requester to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetRequesterNil() {
	o.Requester.Set(nil)
}

// UnsetRequester ensures that no value is present for Requester, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetRequester() {
	o.Requester.Unset()
}

// GetRequesterComment returns the RequesterComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetRequesterComment() string {
	if o == nil || utils.IsNil(o.RequesterComment.Get()) {
		var ret string
		return ret
	}
	return *o.RequesterComment.Get()
}

// GetRequesterCommentOk returns a tuple with the RequesterComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetRequesterCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterComment.Get(), o.RequesterComment.IsSet()
}

// HasRequesterComment returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasRequesterComment() bool {
	if o != nil && o.RequesterComment.IsSet() {
		return true
	}

	return false
}

// SetRequesterComment gets a reference to the given NullableString and assigns it to the RequesterComment field.
func (o *WebRARecoverRequestOnSubmitResponse) SetRequesterComment(v string) {
	o.RequesterComment.Set(&v)
}

// SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetRequesterCommentNil() {
	o.RequesterComment.Set(nil)
}

// UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetRequesterComment() {
	o.RequesterComment.Unset()
}

// GetStatus returns the Status field value
func (o *WebRARecoverRequestOnSubmitResponse) GetStatus() RequestStatus {
	if o == nil {
		var ret RequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WebRARecoverRequestOnSubmitResponse) GetStatusOk() (*RequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WebRARecoverRequestOnSubmitResponse) SetStatus(v RequestStatus) {
	o.Status = v
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetTeam() string {
	if o == nil || utils.IsNil(o.Team.Get()) {
		var ret string
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableString and assigns it to the Team field.
func (o *WebRARecoverRequestOnSubmitResponse) SetTeam(v string) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *WebRARecoverRequestOnSubmitResponse) UnsetTeam() {
	o.Team.Unset()
}

// GetTriggerResults returns the TriggerResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRARecoverRequestOnSubmitResponse) GetTriggerResults() []TriggerResult {
	if o == nil {
		var ret []TriggerResult
		return ret
	}
	return o.TriggerResults
}

// GetTriggerResultsOk returns a tuple with the TriggerResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRARecoverRequestOnSubmitResponse) GetTriggerResultsOk() ([]TriggerResult, bool) {
	if o == nil || utils.IsNil(o.TriggerResults) {
		return nil, false
	}
	return o.TriggerResults, true
}

// HasTriggerResults returns a boolean if a field has been set.
func (o *WebRARecoverRequestOnSubmitResponse) HasTriggerResults() bool {
	if o != nil && !utils.IsNil(o.TriggerResults) {
		return true
	}

	return false
}

// SetTriggerResults gets a reference to the given []TriggerResult and assigns it to the TriggerResults field.
func (o *WebRARecoverRequestOnSubmitResponse) SetTriggerResults(v []TriggerResult) {
	o.TriggerResults = v
}

func (o WebRARecoverRequestOnSubmitResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebRARecoverRequestOnSubmitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	toSerialize["module"] = o.Module
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.Pkcs12.IsSet() {
		toSerialize["pkcs12"] = o.Pkcs12.Get()
	}
	toSerialize["workflow"] = o.Workflow
	toSerialize["_id"] = o.Id
	if o.Approver.IsSet() {
		toSerialize["approver"] = o.Approver.Get()
	}
	if o.ApproverComment.IsSet() {
		toSerialize["approverComment"] = o.ApproverComment.Get()
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if !utils.IsNil(o.Dn) {
		toSerialize["dn"] = o.Dn
	}
	if o.DryRun.IsSet() {
		toSerialize["dryRun"] = o.DryRun.Get()
	}
	if !utils.IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if o.GlobalHolderIdCount.IsSet() {
		toSerialize["globalHolderIdCount"] = o.GlobalHolderIdCount.Get()
	}
	toSerialize["holderId"] = o.HolderId
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["lastModificationDate"] = o.LastModificationDate
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["profile"] = o.Profile
	if o.ProfileHolderIdCount.IsSet() {
		toSerialize["profileHolderIdCount"] = o.ProfileHolderIdCount.Get()
	}
	toSerialize["registrationDate"] = o.RegistrationDate
	toSerialize["removeAt"] = o.RemoveAt
	if o.Requester.IsSet() {
		toSerialize["requester"] = o.Requester.Get()
	}
	if o.RequesterComment.IsSet() {
		toSerialize["requesterComment"] = o.RequesterComment.Get()
	}
	toSerialize["status"] = o.Status
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}
	if o.TriggerResults != nil {
		toSerialize["triggerResults"] = o.TriggerResults
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebRARecoverRequestOnSubmitResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"module",
		"workflow",
		"_id",
		"holderId",
		"lastModificationDate",
		"profile",
		"registrationDate",
		"removeAt",
		"status",
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

	varWebRARecoverRequestOnSubmitResponse := _WebRARecoverRequestOnSubmitResponse{}

	err = json.Unmarshal(data, &varWebRARecoverRequestOnSubmitResponse)

	if err != nil {
		return err
	}

	*o = WebRARecoverRequestOnSubmitResponse(varWebRARecoverRequestOnSubmitResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "module")
		delete(additionalProperties, "password")
		delete(additionalProperties, "pkcs12")
		delete(additionalProperties, "workflow")
		delete(additionalProperties, "_id")
		delete(additionalProperties, "approver")
		delete(additionalProperties, "approverComment")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "dn")
		delete(additionalProperties, "dryRun")
		delete(additionalProperties, "expirationDate")
		delete(additionalProperties, "globalHolderIdCount")
		delete(additionalProperties, "holderId")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "lastModificationDate")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "profileHolderIdCount")
		delete(additionalProperties, "registrationDate")
		delete(additionalProperties, "removeAt")
		delete(additionalProperties, "requester")
		delete(additionalProperties, "requesterComment")
		delete(additionalProperties, "status")
		delete(additionalProperties, "team")
		delete(additionalProperties, "triggerResults")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebRARecoverRequestOnSubmitResponse struct {
	value *WebRARecoverRequestOnSubmitResponse
	isSet bool
}

func (v NullableWebRARecoverRequestOnSubmitResponse) Get() *WebRARecoverRequestOnSubmitResponse {
	return v.value
}

func (v *NullableWebRARecoverRequestOnSubmitResponse) Set(val *WebRARecoverRequestOnSubmitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebRARecoverRequestOnSubmitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebRARecoverRequestOnSubmitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebRARecoverRequestOnSubmitResponse(val *WebRARecoverRequestOnSubmitResponse) *NullableWebRARecoverRequestOnSubmitResponse {
	return &NullableWebRARecoverRequestOnSubmitResponse{value: val, isSet: true}
}

func (v NullableWebRARecoverRequestOnSubmitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebRARecoverRequestOnSubmitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
