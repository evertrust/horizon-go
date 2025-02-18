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

// checks if the WebRAUpdateRequestOnSubmitResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebRAUpdateRequestOnSubmitResponse{}

// WebRAUpdateRequestOnSubmitResponse struct for WebRAUpdateRequestOnSubmitResponse
type WebRAUpdateRequestOnSubmitResponse struct {
	// The module of the certificate updated.
	Module Module `json:"module"`
	// What this request will do. For an update request, this is always `update`
	Workflow string `json:"workflow"`
	// The user-data that will be used to generate the certificate
	Template WebRAUpdateRequestTemplate `json:"template"`
	// The certificate that was updated for this request. This is only available after the request has been approved
	Certificate NullableCertificate `json:"certificate,omitempty"`
	// Object internal ID
	Id string `json:"_id"`
	Status RequestStatus `json:"status"`
	// The associated profile name
	Profile string `json:"profile"`
	// Certificate's Distinguished Name
	Dn *string `json:"dn,omitempty"`
	// The requester's principal identifier
	Requester NullableString `json:"requester,omitempty"`
	// The team that will be assigned to this certificate. Teams are used to link certificates to people and to assign permissions to them
	Team NullableString `json:"team,omitempty"`
	// The approver's principal identifier
	Approver NullableString `json:"approver,omitempty"`
	// The request's contact email
	Contact NullableString `json:"contact,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment NullableString `json:"requesterComment,omitempty"`
	// Free-text field editable by the approver to provider more context on the request
	ApproverComment NullableString `json:"approverComment,omitempty"`
	// The date the request was created. This is set by the system
	RegistrationDate int64 `json:"registrationDate"`
	// The date the request was last modified. This is set by the system
	LastModificationDate int64 `json:"lastModificationDate"`
	// The date the request will expire. This is set by the system
	ExpirationDate *int64 `json:"expirationDate,omitempty"`
	// The date the requested will be deleted. This is set by the system
	RemoveAt int64 `json:"removeAt"`
	// The result of the execution of triggers on this request
	TriggerResults []TriggerResult `json:"triggerResults,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId string `json:"holderId"`
	// The number of certificates that are currently valid and have the same DN and SANs in the Horizon database
	GlobalHolderIdCount NullableInt64 `json:"globalHolderIdCount,omitempty"`
	// The number of certificates that are currently valid and have the same DN and SANs in the same enrollment profile
	ProfileHolderIdCount NullableInt64 `json:"profileHolderIdCount,omitempty"`
	// The labels set in this request
	Labels []LabelData `json:"labels,omitempty"`
	// The metadata set in this request
	Metadata []CertificateMetadata `json:"metadata,omitempty"`
	// If true, the request is validated, but will not result in an enrollment
	DryRun NullableBool `json:"dryRun,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebRAUpdateRequestOnSubmitResponse WebRAUpdateRequestOnSubmitResponse

// NewWebRAUpdateRequestOnSubmitResponse instantiates a new WebRAUpdateRequestOnSubmitResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebRAUpdateRequestOnSubmitResponse(module Module, workflow string, template WebRAUpdateRequestTemplate, id string, status RequestStatus, profile string, registrationDate int64, lastModificationDate int64, removeAt int64, holderId string) *WebRAUpdateRequestOnSubmitResponse {
	this := WebRAUpdateRequestOnSubmitResponse{}
	this.Id = id
	this.Module = module
	this.Workflow = workflow
	this.Status = status
	this.Profile = profile
	this.RegistrationDate = registrationDate
	this.LastModificationDate = lastModificationDate
	this.RemoveAt = removeAt
	this.HolderId = holderId
	var dryRun bool = false
	this.DryRun = *NewNullableBool(&dryRun)
	return &this
}

// NewWebRAUpdateRequestOnSubmitResponseWithDefaults instantiates a new WebRAUpdateRequestOnSubmitResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebRAUpdateRequestOnSubmitResponseWithDefaults() *WebRAUpdateRequestOnSubmitResponse {
	this := WebRAUpdateRequestOnSubmitResponse{}
	var dryRun bool = false
	this.DryRun = *NewNullableBool(&dryRun)
	return &this
}

// GetModule returns the Module field value
func (o *WebRAUpdateRequestOnSubmitResponse) GetModule() Module {
	if o == nil {
		var ret Module
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetModuleOk() (*Module, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *WebRAUpdateRequestOnSubmitResponse) SetModule(v Module) {
	o.Module = v
}

// GetWorkflow returns the Workflow field value
func (o *WebRAUpdateRequestOnSubmitResponse) GetWorkflow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetWorkflowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Workflow, true
}

// SetWorkflow sets field value
func (o *WebRAUpdateRequestOnSubmitResponse) SetWorkflow(v string) {
	o.Workflow = v
}

// GetTemplate returns the Template field value
func (o *WebRAUpdateRequestOnSubmitResponse) GetTemplate() WebRAUpdateRequestTemplate {
	if o == nil {
		var ret WebRAUpdateRequestTemplate
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetTemplateOk() (*WebRAUpdateRequestTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *WebRAUpdateRequestOnSubmitResponse) SetTemplate(v WebRAUpdateRequestTemplate) {
	o.Template = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetCertificate() Certificate {
	if o == nil || IsNil(o.Certificate.Get()) {
		var ret Certificate
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetCertificateOk() (*Certificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableCertificate and assigns it to the Certificate field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetCertificate(v Certificate) {
	o.Certificate.Set(&v)
}
// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetId returns the Id field value
func (o *WebRAUpdateRequestOnSubmitResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebRAUpdateRequestOnSubmitResponse) SetId(v string) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *WebRAUpdateRequestOnSubmitResponse) GetStatus() RequestStatus {
	if o == nil {
		var ret RequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetStatusOk() (*RequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WebRAUpdateRequestOnSubmitResponse) SetStatus(v RequestStatus) {
	o.Status = v
}

// GetProfile returns the Profile field value
func (o *WebRAUpdateRequestOnSubmitResponse) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *WebRAUpdateRequestOnSubmitResponse) SetProfile(v string) {
	o.Profile = v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *WebRAUpdateRequestOnSubmitResponse) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetDn(v string) {
	o.Dn = &v
}

// GetRequester returns the Requester field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetRequester() string {
	if o == nil || IsNil(o.Requester.Get()) {
		var ret string
		return ret
	}
	return *o.Requester.Get()
}

// GetRequesterOk returns a tuple with the Requester field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetRequesterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requester.Get(), o.Requester.IsSet()
}

// HasRequester returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasRequester() bool {
	if o != nil && o.Requester.IsSet() {
		return true
	}

	return false
}

// SetRequester gets a reference to the given NullableString and assigns it to the Requester field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetRequester(v string) {
	o.Requester.Set(&v)
}
// SetRequesterNil sets the value for Requester to be an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) SetRequesterNil() {
	o.Requester.Set(nil)
}

// UnsetRequester ensures that no value is present for Requester, not even an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) UnsetRequester() {
	o.Requester.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetTeam() string {
	if o == nil || IsNil(o.Team.Get()) {
		var ret string
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetTeamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableString and assigns it to the Team field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetTeam(v string) {
	o.Team.Set(&v)
}
// SetTeamNil sets the value for Team to be an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) UnsetTeam() {
	o.Team.Unset()
}

// GetApprover returns the Approver field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetApprover() string {
	if o == nil || IsNil(o.Approver.Get()) {
		var ret string
		return ret
	}
	return *o.Approver.Get()
}

// GetApproverOk returns a tuple with the Approver field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetApproverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approver.Get(), o.Approver.IsSet()
}

// HasApprover returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasApprover() bool {
	if o != nil && o.Approver.IsSet() {
		return true
	}

	return false
}

// SetApprover gets a reference to the given NullableString and assigns it to the Approver field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetApprover(v string) {
	o.Approver.Set(&v)
}
// SetApproverNil sets the value for Approver to be an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) SetApproverNil() {
	o.Approver.Set(nil)
}

// UnsetApprover ensures that no value is present for Approver, not even an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) UnsetApprover() {
	o.Approver.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetContact() string {
	if o == nil || IsNil(o.Contact.Get()) {
		var ret string
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableString and assigns it to the Contact field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetContact(v string) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) UnsetContact() {
	o.Contact.Unset()
}

// GetRequesterComment returns the RequesterComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetRequesterComment() string {
	if o == nil || IsNil(o.RequesterComment.Get()) {
		var ret string
		return ret
	}
	return *o.RequesterComment.Get()
}

// GetRequesterCommentOk returns a tuple with the RequesterComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetRequesterCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterComment.Get(), o.RequesterComment.IsSet()
}

// HasRequesterComment returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasRequesterComment() bool {
	if o != nil && o.RequesterComment.IsSet() {
		return true
	}

	return false
}

// SetRequesterComment gets a reference to the given NullableString and assigns it to the RequesterComment field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetRequesterComment(v string) {
	o.RequesterComment.Set(&v)
}
// SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) SetRequesterCommentNil() {
	o.RequesterComment.Set(nil)
}

// UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) UnsetRequesterComment() {
	o.RequesterComment.Unset()
}

// GetApproverComment returns the ApproverComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetApproverComment() string {
	if o == nil || IsNil(o.ApproverComment.Get()) {
		var ret string
		return ret
	}
	return *o.ApproverComment.Get()
}

// GetApproverCommentOk returns a tuple with the ApproverComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetApproverCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproverComment.Get(), o.ApproverComment.IsSet()
}

// HasApproverComment returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasApproverComment() bool {
	if o != nil && o.ApproverComment.IsSet() {
		return true
	}

	return false
}

// SetApproverComment gets a reference to the given NullableString and assigns it to the ApproverComment field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetApproverComment(v string) {
	o.ApproverComment.Set(&v)
}
// SetApproverCommentNil sets the value for ApproverComment to be an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) SetApproverCommentNil() {
	o.ApproverComment.Set(nil)
}

// UnsetApproverComment ensures that no value is present for ApproverComment, not even an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) UnsetApproverComment() {
	o.ApproverComment.Unset()
}

// GetRegistrationDate returns the RegistrationDate field value
func (o *WebRAUpdateRequestOnSubmitResponse) GetRegistrationDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RegistrationDate
}

// GetRegistrationDateOk returns a tuple with the RegistrationDate field value
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetRegistrationDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistrationDate, true
}

// SetRegistrationDate sets field value
func (o *WebRAUpdateRequestOnSubmitResponse) SetRegistrationDate(v int64) {
	o.RegistrationDate = v
}

// GetLastModificationDate returns the LastModificationDate field value
func (o *WebRAUpdateRequestOnSubmitResponse) GetLastModificationDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LastModificationDate
}

// GetLastModificationDateOk returns a tuple with the LastModificationDate field value
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetLastModificationDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModificationDate, true
}

// SetLastModificationDate sets field value
func (o *WebRAUpdateRequestOnSubmitResponse) SetLastModificationDate(v int64) {
	o.LastModificationDate = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *WebRAUpdateRequestOnSubmitResponse) GetExpirationDate() int64 {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret int64
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetExpirationDateOk() (*int64, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetExpirationDate(v int64) {
	o.ExpirationDate = &v
}

// GetRemoveAt returns the RemoveAt field value
func (o *WebRAUpdateRequestOnSubmitResponse) GetRemoveAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RemoveAt
}

// GetRemoveAtOk returns a tuple with the RemoveAt field value
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetRemoveAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoveAt, true
}

// SetRemoveAt sets field value
func (o *WebRAUpdateRequestOnSubmitResponse) SetRemoveAt(v int64) {
	o.RemoveAt = v
}

// GetTriggerResults returns the TriggerResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetTriggerResults() []TriggerResult {
	if o == nil {
		var ret []TriggerResult
		return ret
	}
	return o.TriggerResults
}

// GetTriggerResultsOk returns a tuple with the TriggerResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetTriggerResultsOk() ([]TriggerResult, bool) {
	if o == nil || IsNil(o.TriggerResults) {
		return nil, false
	}
	return o.TriggerResults, true
}

// HasTriggerResults returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasTriggerResults() bool {
	if o != nil && !IsNil(o.TriggerResults) {
		return true
	}

	return false
}

// SetTriggerResults gets a reference to the given []TriggerResult and assigns it to the TriggerResults field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetTriggerResults(v []TriggerResult) {
	o.TriggerResults = v
}

// GetHolderId returns the HolderId field value
func (o *WebRAUpdateRequestOnSubmitResponse) GetHolderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HolderId
}

// GetHolderIdOk returns a tuple with the HolderId field value
// and a boolean to check if the value has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) GetHolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HolderId, true
}

// SetHolderId sets field value
func (o *WebRAUpdateRequestOnSubmitResponse) SetHolderId(v string) {
	o.HolderId = v
}

// GetGlobalHolderIdCount returns the GlobalHolderIdCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetGlobalHolderIdCount() int64 {
	if o == nil || IsNil(o.GlobalHolderIdCount.Get()) {
		var ret int64
		return ret
	}
	return *o.GlobalHolderIdCount.Get()
}

// GetGlobalHolderIdCountOk returns a tuple with the GlobalHolderIdCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetGlobalHolderIdCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GlobalHolderIdCount.Get(), o.GlobalHolderIdCount.IsSet()
}

// HasGlobalHolderIdCount returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasGlobalHolderIdCount() bool {
	if o != nil && o.GlobalHolderIdCount.IsSet() {
		return true
	}

	return false
}

// SetGlobalHolderIdCount gets a reference to the given NullableInt64 and assigns it to the GlobalHolderIdCount field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetGlobalHolderIdCount(v int64) {
	o.GlobalHolderIdCount.Set(&v)
}
// SetGlobalHolderIdCountNil sets the value for GlobalHolderIdCount to be an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) SetGlobalHolderIdCountNil() {
	o.GlobalHolderIdCount.Set(nil)
}

// UnsetGlobalHolderIdCount ensures that no value is present for GlobalHolderIdCount, not even an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) UnsetGlobalHolderIdCount() {
	o.GlobalHolderIdCount.Unset()
}

// GetProfileHolderIdCount returns the ProfileHolderIdCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetProfileHolderIdCount() int64 {
	if o == nil || IsNil(o.ProfileHolderIdCount.Get()) {
		var ret int64
		return ret
	}
	return *o.ProfileHolderIdCount.Get()
}

// GetProfileHolderIdCountOk returns a tuple with the ProfileHolderIdCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetProfileHolderIdCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileHolderIdCount.Get(), o.ProfileHolderIdCount.IsSet()
}

// HasProfileHolderIdCount returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasProfileHolderIdCount() bool {
	if o != nil && o.ProfileHolderIdCount.IsSet() {
		return true
	}

	return false
}

// SetProfileHolderIdCount gets a reference to the given NullableInt64 and assigns it to the ProfileHolderIdCount field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetProfileHolderIdCount(v int64) {
	o.ProfileHolderIdCount.Set(&v)
}
// SetProfileHolderIdCountNil sets the value for ProfileHolderIdCount to be an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) SetProfileHolderIdCountNil() {
	o.ProfileHolderIdCount.Set(nil)
}

// UnsetProfileHolderIdCount ensures that no value is present for ProfileHolderIdCount, not even an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) UnsetProfileHolderIdCount() {
	o.ProfileHolderIdCount.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetLabels() []LabelData {
	if o == nil {
		var ret []LabelData
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetLabelsOk() ([]LabelData, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelData and assigns it to the Labels field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetLabels(v []LabelData) {
	o.Labels = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetMetadata() []CertificateMetadata {
	if o == nil {
		var ret []CertificateMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetMetadataOk() ([]CertificateMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []CertificateMetadata and assigns it to the Metadata field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetMetadata(v []CertificateMetadata) {
	o.Metadata = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestOnSubmitResponse) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun.Get()) {
		var ret bool
		return ret
	}
	return *o.DryRun.Get()
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestOnSubmitResponse) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DryRun.Get(), o.DryRun.IsSet()
}

// HasDryRun returns a boolean if a field has been set.
func (o *WebRAUpdateRequestOnSubmitResponse) HasDryRun() bool {
	if o != nil && o.DryRun.IsSet() {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given NullableBool and assigns it to the DryRun field.
func (o *WebRAUpdateRequestOnSubmitResponse) SetDryRun(v bool) {
	o.DryRun.Set(&v)
}
// SetDryRunNil sets the value for DryRun to be an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) SetDryRunNil() {
	o.DryRun.Set(nil)
}

// UnsetDryRun ensures that no value is present for DryRun, not even an explicit nil
func (o *WebRAUpdateRequestOnSubmitResponse) UnsetDryRun() {
	o.DryRun.Unset()
}

func (o WebRAUpdateRequestOnSubmitResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebRAUpdateRequestOnSubmitResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["module"] = o.Module
	toSerialize["workflow"] = o.Workflow
	toSerialize["template"] = o.Template
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	toSerialize["_id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["profile"] = o.Profile
	if !IsNil(o.Dn) {
		toSerialize["dn"] = o.Dn
	}
	if o.Requester.IsSet() {
		toSerialize["requester"] = o.Requester.Get()
	}
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}
	if o.Approver.IsSet() {
		toSerialize["approver"] = o.Approver.Get()
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.RequesterComment.IsSet() {
		toSerialize["requesterComment"] = o.RequesterComment.Get()
	}
	if o.ApproverComment.IsSet() {
		toSerialize["approverComment"] = o.ApproverComment.Get()
	}
	toSerialize["registrationDate"] = o.RegistrationDate
	toSerialize["lastModificationDate"] = o.LastModificationDate
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	toSerialize["removeAt"] = o.RemoveAt
	if o.TriggerResults != nil {
		toSerialize["triggerResults"] = o.TriggerResults
	}
	toSerialize["holderId"] = o.HolderId
	if o.GlobalHolderIdCount.IsSet() {
		toSerialize["globalHolderIdCount"] = o.GlobalHolderIdCount.Get()
	}
	if o.ProfileHolderIdCount.IsSet() {
		toSerialize["profileHolderIdCount"] = o.ProfileHolderIdCount.Get()
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.DryRun.IsSet() {
		toSerialize["dryRun"] = o.DryRun.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebRAUpdateRequestOnSubmitResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"module",
		"workflow",
		"template",
		"_id",
		"status",
		"profile",
		"registrationDate",
		"lastModificationDate",
		"removeAt",
		"holderId",
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

	varWebRAUpdateRequestOnSubmitResponse := _WebRAUpdateRequestOnSubmitResponse{}

	err = json.Unmarshal(data, &varWebRAUpdateRequestOnSubmitResponse)

	if err != nil {
		return err
	}

	*o = WebRAUpdateRequestOnSubmitResponse(varWebRAUpdateRequestOnSubmitResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "module")
		delete(additionalProperties, "workflow")
		delete(additionalProperties, "template")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "dn")
		delete(additionalProperties, "requester")
		delete(additionalProperties, "team")
		delete(additionalProperties, "approver")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "requesterComment")
		delete(additionalProperties, "approverComment")
		delete(additionalProperties, "registrationDate")
		delete(additionalProperties, "lastModificationDate")
		delete(additionalProperties, "expirationDate")
		delete(additionalProperties, "removeAt")
		delete(additionalProperties, "triggerResults")
		delete(additionalProperties, "holderId")
		delete(additionalProperties, "globalHolderIdCount")
		delete(additionalProperties, "profileHolderIdCount")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "dryRun")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebRAUpdateRequestOnSubmitResponse struct {
	value *WebRAUpdateRequestOnSubmitResponse
	isSet bool
}

func (v NullableWebRAUpdateRequestOnSubmitResponse) Get() *WebRAUpdateRequestOnSubmitResponse {
	return v.value
}

func (v *NullableWebRAUpdateRequestOnSubmitResponse) Set(val *WebRAUpdateRequestOnSubmitResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebRAUpdateRequestOnSubmitResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebRAUpdateRequestOnSubmitResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebRAUpdateRequestOnSubmitResponse(val *WebRAUpdateRequestOnSubmitResponse) *NullableWebRAUpdateRequestOnSubmitResponse {
	return &NullableWebRAUpdateRequestOnSubmitResponse{value: val, isSet: true}
}

func (v NullableWebRAUpdateRequestOnSubmitResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebRAUpdateRequestOnSubmitResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


