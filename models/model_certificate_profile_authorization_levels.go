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

// checks if the CertificateProfileAuthorizationLevels type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CertificateProfileAuthorizationLevels{}

// CertificateProfileAuthorizationLevels struct for CertificateProfileAuthorizationLevels
type CertificateProfileAuthorizationLevels struct {
	ApproveEnroll        NullableAuthorizationLevel `json:"approveEnroll,omitempty"`
	ApproveMigrate       NullableAuthorizationLevel `json:"approveMigrate,omitempty"`
	ApproveRecover       NullableAuthorizationLevel `json:"approveRecover,omitempty"`
	ApproveRenew         NullableAuthorizationLevel `json:"approveRenew,omitempty"`
	ApproveRevoke        *AuthorizationLevel        `json:"approveRevoke,omitempty"`
	ApproveUpdate        AuthorizationLevel         `json:"approveUpdate"`
	AuditRequest         NullableAuthorizationLevel `json:"auditRequest,omitempty"`
	Enroll               NullableAuthorizationLevel `json:"enroll,omitempty"`
	EnrollApi            NullableAuthorizationLevel `json:"enrollApi,omitempty"`
	Migrate              NullableAuthorizationLevel `json:"migrate,omitempty"`
	Recover              NullableAuthorizationLevel `json:"recover,omitempty"`
	RecoverApi           NullableAuthorizationLevel `json:"recoverApi,omitempty"`
	Renew                NullableAuthorizationLevel `json:"renew,omitempty"`
	RenewApi             NullableAuthorizationLevel `json:"renewApi,omitempty"`
	RequestEnroll        NullableAuthorizationLevel `json:"requestEnroll,omitempty"`
	RequestMigrate       NullableAuthorizationLevel `json:"requestMigrate,omitempty"`
	RequestRecover       NullableAuthorizationLevel `json:"requestRecover,omitempty"`
	RequestRenew         NullableAuthorizationLevel `json:"requestRenew,omitempty"`
	RequestRevoke        *AuthorizationLevel        `json:"requestRevoke,omitempty"`
	RequestUpdate        AuthorizationLevel         `json:"requestUpdate"`
	Revoke               *AuthorizationLevel        `json:"revoke,omitempty"`
	Search               AuthorizationLevel         `json:"search"`
	Update               AuthorizationLevel         `json:"update"`
	AdditionalProperties map[string]interface{}
}

type _CertificateProfileAuthorizationLevels CertificateProfileAuthorizationLevels

// NewCertificateProfileAuthorizationLevels instantiates a new CertificateProfileAuthorizationLevels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateProfileAuthorizationLevels(approveUpdate AuthorizationLevel, requestUpdate AuthorizationLevel, search AuthorizationLevel, update AuthorizationLevel) *CertificateProfileAuthorizationLevels {
	this := CertificateProfileAuthorizationLevels{}
	this.ApproveUpdate = approveUpdate
	this.RequestUpdate = requestUpdate
	this.Search = search
	this.Update = update
	return &this
}

// NewCertificateProfileAuthorizationLevelsWithDefaults instantiates a new CertificateProfileAuthorizationLevels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateProfileAuthorizationLevelsWithDefaults() *CertificateProfileAuthorizationLevels {
	this := CertificateProfileAuthorizationLevels{}
	return &this
}

// GetApproveEnroll returns the ApproveEnroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetApproveEnroll() AuthorizationLevel {
	if o == nil || utils.IsNil(o.ApproveEnroll.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.ApproveEnroll.Get()
}

// GetApproveEnrollOk returns a tuple with the ApproveEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetApproveEnrollOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproveEnroll.Get(), o.ApproveEnroll.IsSet()
}

// HasApproveEnroll returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasApproveEnroll() bool {
	if o != nil && o.ApproveEnroll.IsSet() {
		return true
	}

	return false
}

// SetApproveEnroll gets a reference to the given NullableAuthorizationLevel and assigns it to the ApproveEnroll field.
func (o *CertificateProfileAuthorizationLevels) SetApproveEnroll(v AuthorizationLevel) {
	o.ApproveEnroll.Set(&v)
}

// SetApproveEnrollNil sets the value for ApproveEnroll to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetApproveEnrollNil() {
	o.ApproveEnroll.Set(nil)
}

// UnsetApproveEnroll ensures that no value is present for ApproveEnroll, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetApproveEnroll() {
	o.ApproveEnroll.Unset()
}

// GetApproveMigrate returns the ApproveMigrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetApproveMigrate() AuthorizationLevel {
	if o == nil || utils.IsNil(o.ApproveMigrate.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.ApproveMigrate.Get()
}

// GetApproveMigrateOk returns a tuple with the ApproveMigrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetApproveMigrateOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproveMigrate.Get(), o.ApproveMigrate.IsSet()
}

// HasApproveMigrate returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasApproveMigrate() bool {
	if o != nil && o.ApproveMigrate.IsSet() {
		return true
	}

	return false
}

// SetApproveMigrate gets a reference to the given NullableAuthorizationLevel and assigns it to the ApproveMigrate field.
func (o *CertificateProfileAuthorizationLevels) SetApproveMigrate(v AuthorizationLevel) {
	o.ApproveMigrate.Set(&v)
}

// SetApproveMigrateNil sets the value for ApproveMigrate to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetApproveMigrateNil() {
	o.ApproveMigrate.Set(nil)
}

// UnsetApproveMigrate ensures that no value is present for ApproveMigrate, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetApproveMigrate() {
	o.ApproveMigrate.Unset()
}

// GetApproveRecover returns the ApproveRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetApproveRecover() AuthorizationLevel {
	if o == nil || utils.IsNil(o.ApproveRecover.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.ApproveRecover.Get()
}

// GetApproveRecoverOk returns a tuple with the ApproveRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetApproveRecoverOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproveRecover.Get(), o.ApproveRecover.IsSet()
}

// HasApproveRecover returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasApproveRecover() bool {
	if o != nil && o.ApproveRecover.IsSet() {
		return true
	}

	return false
}

// SetApproveRecover gets a reference to the given NullableAuthorizationLevel and assigns it to the ApproveRecover field.
func (o *CertificateProfileAuthorizationLevels) SetApproveRecover(v AuthorizationLevel) {
	o.ApproveRecover.Set(&v)
}

// SetApproveRecoverNil sets the value for ApproveRecover to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetApproveRecoverNil() {
	o.ApproveRecover.Set(nil)
}

// UnsetApproveRecover ensures that no value is present for ApproveRecover, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetApproveRecover() {
	o.ApproveRecover.Unset()
}

// GetApproveRenew returns the ApproveRenew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetApproveRenew() AuthorizationLevel {
	if o == nil || utils.IsNil(o.ApproveRenew.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.ApproveRenew.Get()
}

// GetApproveRenewOk returns a tuple with the ApproveRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetApproveRenewOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproveRenew.Get(), o.ApproveRenew.IsSet()
}

// HasApproveRenew returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasApproveRenew() bool {
	if o != nil && o.ApproveRenew.IsSet() {
		return true
	}

	return false
}

// SetApproveRenew gets a reference to the given NullableAuthorizationLevel and assigns it to the ApproveRenew field.
func (o *CertificateProfileAuthorizationLevels) SetApproveRenew(v AuthorizationLevel) {
	o.ApproveRenew.Set(&v)
}

// SetApproveRenewNil sets the value for ApproveRenew to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetApproveRenewNil() {
	o.ApproveRenew.Set(nil)
}

// UnsetApproveRenew ensures that no value is present for ApproveRenew, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetApproveRenew() {
	o.ApproveRenew.Unset()
}

// GetApproveRevoke returns the ApproveRevoke field value if set, zero value otherwise.
func (o *CertificateProfileAuthorizationLevels) GetApproveRevoke() AuthorizationLevel {
	if o == nil || utils.IsNil(o.ApproveRevoke) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.ApproveRevoke
}

// GetApproveRevokeOk returns a tuple with the ApproveRevoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateProfileAuthorizationLevels) GetApproveRevokeOk() (*AuthorizationLevel, bool) {
	if o == nil || utils.IsNil(o.ApproveRevoke) {
		return nil, false
	}
	return o.ApproveRevoke, true
}

// HasApproveRevoke returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasApproveRevoke() bool {
	if o != nil && !utils.IsNil(o.ApproveRevoke) {
		return true
	}

	return false
}

// SetApproveRevoke gets a reference to the given AuthorizationLevel and assigns it to the ApproveRevoke field.
func (o *CertificateProfileAuthorizationLevels) SetApproveRevoke(v AuthorizationLevel) {
	o.ApproveRevoke = &v
}

// GetApproveUpdate returns the ApproveUpdate field value
func (o *CertificateProfileAuthorizationLevels) GetApproveUpdate() AuthorizationLevel {
	if o == nil {
		var ret AuthorizationLevel
		return ret
	}

	return o.ApproveUpdate
}

// GetApproveUpdateOk returns a tuple with the ApproveUpdate field value
// and a boolean to check if the value has been set.
func (o *CertificateProfileAuthorizationLevels) GetApproveUpdateOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproveUpdate, true
}

// SetApproveUpdate sets field value
func (o *CertificateProfileAuthorizationLevels) SetApproveUpdate(v AuthorizationLevel) {
	o.ApproveUpdate = v
}

// GetAuditRequest returns the AuditRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetAuditRequest() AuthorizationLevel {
	if o == nil || utils.IsNil(o.AuditRequest.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.AuditRequest.Get()
}

// GetAuditRequestOk returns a tuple with the AuditRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetAuditRequestOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuditRequest.Get(), o.AuditRequest.IsSet()
}

// HasAuditRequest returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasAuditRequest() bool {
	if o != nil && o.AuditRequest.IsSet() {
		return true
	}

	return false
}

// SetAuditRequest gets a reference to the given NullableAuthorizationLevel and assigns it to the AuditRequest field.
func (o *CertificateProfileAuthorizationLevels) SetAuditRequest(v AuthorizationLevel) {
	o.AuditRequest.Set(&v)
}

// SetAuditRequestNil sets the value for AuditRequest to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetAuditRequestNil() {
	o.AuditRequest.Set(nil)
}

// UnsetAuditRequest ensures that no value is present for AuditRequest, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetAuditRequest() {
	o.AuditRequest.Unset()
}

// GetEnroll returns the Enroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetEnroll() AuthorizationLevel {
	if o == nil || utils.IsNil(o.Enroll.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.Enroll.Get()
}

// GetEnrollOk returns a tuple with the Enroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetEnrollOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enroll.Get(), o.Enroll.IsSet()
}

// HasEnroll returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasEnroll() bool {
	if o != nil && o.Enroll.IsSet() {
		return true
	}

	return false
}

// SetEnroll gets a reference to the given NullableAuthorizationLevel and assigns it to the Enroll field.
func (o *CertificateProfileAuthorizationLevels) SetEnroll(v AuthorizationLevel) {
	o.Enroll.Set(&v)
}

// SetEnrollNil sets the value for Enroll to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetEnrollNil() {
	o.Enroll.Set(nil)
}

// UnsetEnroll ensures that no value is present for Enroll, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetEnroll() {
	o.Enroll.Unset()
}

// GetEnrollApi returns the EnrollApi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetEnrollApi() AuthorizationLevel {
	if o == nil || utils.IsNil(o.EnrollApi.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.EnrollApi.Get()
}

// GetEnrollApiOk returns a tuple with the EnrollApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetEnrollApiOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnrollApi.Get(), o.EnrollApi.IsSet()
}

// HasEnrollApi returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasEnrollApi() bool {
	if o != nil && o.EnrollApi.IsSet() {
		return true
	}

	return false
}

// SetEnrollApi gets a reference to the given NullableAuthorizationLevel and assigns it to the EnrollApi field.
func (o *CertificateProfileAuthorizationLevels) SetEnrollApi(v AuthorizationLevel) {
	o.EnrollApi.Set(&v)
}

// SetEnrollApiNil sets the value for EnrollApi to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetEnrollApiNil() {
	o.EnrollApi.Set(nil)
}

// UnsetEnrollApi ensures that no value is present for EnrollApi, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetEnrollApi() {
	o.EnrollApi.Unset()
}

// GetMigrate returns the Migrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetMigrate() AuthorizationLevel {
	if o == nil || utils.IsNil(o.Migrate.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.Migrate.Get()
}

// GetMigrateOk returns a tuple with the Migrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetMigrateOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Migrate.Get(), o.Migrate.IsSet()
}

// HasMigrate returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasMigrate() bool {
	if o != nil && o.Migrate.IsSet() {
		return true
	}

	return false
}

// SetMigrate gets a reference to the given NullableAuthorizationLevel and assigns it to the Migrate field.
func (o *CertificateProfileAuthorizationLevels) SetMigrate(v AuthorizationLevel) {
	o.Migrate.Set(&v)
}

// SetMigrateNil sets the value for Migrate to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetMigrateNil() {
	o.Migrate.Set(nil)
}

// UnsetMigrate ensures that no value is present for Migrate, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetMigrate() {
	o.Migrate.Unset()
}

// GetRecover returns the Recover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetRecover() AuthorizationLevel {
	if o == nil || utils.IsNil(o.Recover.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.Recover.Get()
}

// GetRecoverOk returns a tuple with the Recover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetRecoverOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recover.Get(), o.Recover.IsSet()
}

// HasRecover returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasRecover() bool {
	if o != nil && o.Recover.IsSet() {
		return true
	}

	return false
}

// SetRecover gets a reference to the given NullableAuthorizationLevel and assigns it to the Recover field.
func (o *CertificateProfileAuthorizationLevels) SetRecover(v AuthorizationLevel) {
	o.Recover.Set(&v)
}

// SetRecoverNil sets the value for Recover to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetRecoverNil() {
	o.Recover.Set(nil)
}

// UnsetRecover ensures that no value is present for Recover, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetRecover() {
	o.Recover.Unset()
}

// GetRecoverApi returns the RecoverApi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetRecoverApi() AuthorizationLevel {
	if o == nil || utils.IsNil(o.RecoverApi.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.RecoverApi.Get()
}

// GetRecoverApiOk returns a tuple with the RecoverApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetRecoverApiOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoverApi.Get(), o.RecoverApi.IsSet()
}

// HasRecoverApi returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasRecoverApi() bool {
	if o != nil && o.RecoverApi.IsSet() {
		return true
	}

	return false
}

// SetRecoverApi gets a reference to the given NullableAuthorizationLevel and assigns it to the RecoverApi field.
func (o *CertificateProfileAuthorizationLevels) SetRecoverApi(v AuthorizationLevel) {
	o.RecoverApi.Set(&v)
}

// SetRecoverApiNil sets the value for RecoverApi to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetRecoverApiNil() {
	o.RecoverApi.Set(nil)
}

// UnsetRecoverApi ensures that no value is present for RecoverApi, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetRecoverApi() {
	o.RecoverApi.Unset()
}

// GetRenew returns the Renew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetRenew() AuthorizationLevel {
	if o == nil || utils.IsNil(o.Renew.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.Renew.Get()
}

// GetRenewOk returns a tuple with the Renew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetRenewOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Renew.Get(), o.Renew.IsSet()
}

// HasRenew returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasRenew() bool {
	if o != nil && o.Renew.IsSet() {
		return true
	}

	return false
}

// SetRenew gets a reference to the given NullableAuthorizationLevel and assigns it to the Renew field.
func (o *CertificateProfileAuthorizationLevels) SetRenew(v AuthorizationLevel) {
	o.Renew.Set(&v)
}

// SetRenewNil sets the value for Renew to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetRenewNil() {
	o.Renew.Set(nil)
}

// UnsetRenew ensures that no value is present for Renew, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetRenew() {
	o.Renew.Unset()
}

// GetRenewApi returns the RenewApi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetRenewApi() AuthorizationLevel {
	if o == nil || utils.IsNil(o.RenewApi.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.RenewApi.Get()
}

// GetRenewApiOk returns a tuple with the RenewApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetRenewApiOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.RenewApi.Get(), o.RenewApi.IsSet()
}

// HasRenewApi returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasRenewApi() bool {
	if o != nil && o.RenewApi.IsSet() {
		return true
	}

	return false
}

// SetRenewApi gets a reference to the given NullableAuthorizationLevel and assigns it to the RenewApi field.
func (o *CertificateProfileAuthorizationLevels) SetRenewApi(v AuthorizationLevel) {
	o.RenewApi.Set(&v)
}

// SetRenewApiNil sets the value for RenewApi to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetRenewApiNil() {
	o.RenewApi.Set(nil)
}

// UnsetRenewApi ensures that no value is present for RenewApi, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetRenewApi() {
	o.RenewApi.Unset()
}

// GetRequestEnroll returns the RequestEnroll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetRequestEnroll() AuthorizationLevel {
	if o == nil || utils.IsNil(o.RequestEnroll.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.RequestEnroll.Get()
}

// GetRequestEnrollOk returns a tuple with the RequestEnroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetRequestEnrollOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestEnroll.Get(), o.RequestEnroll.IsSet()
}

// HasRequestEnroll returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasRequestEnroll() bool {
	if o != nil && o.RequestEnroll.IsSet() {
		return true
	}

	return false
}

// SetRequestEnroll gets a reference to the given NullableAuthorizationLevel and assigns it to the RequestEnroll field.
func (o *CertificateProfileAuthorizationLevels) SetRequestEnroll(v AuthorizationLevel) {
	o.RequestEnroll.Set(&v)
}

// SetRequestEnrollNil sets the value for RequestEnroll to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetRequestEnrollNil() {
	o.RequestEnroll.Set(nil)
}

// UnsetRequestEnroll ensures that no value is present for RequestEnroll, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetRequestEnroll() {
	o.RequestEnroll.Unset()
}

// GetRequestMigrate returns the RequestMigrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetRequestMigrate() AuthorizationLevel {
	if o == nil || utils.IsNil(o.RequestMigrate.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.RequestMigrate.Get()
}

// GetRequestMigrateOk returns a tuple with the RequestMigrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetRequestMigrateOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestMigrate.Get(), o.RequestMigrate.IsSet()
}

// HasRequestMigrate returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasRequestMigrate() bool {
	if o != nil && o.RequestMigrate.IsSet() {
		return true
	}

	return false
}

// SetRequestMigrate gets a reference to the given NullableAuthorizationLevel and assigns it to the RequestMigrate field.
func (o *CertificateProfileAuthorizationLevels) SetRequestMigrate(v AuthorizationLevel) {
	o.RequestMigrate.Set(&v)
}

// SetRequestMigrateNil sets the value for RequestMigrate to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetRequestMigrateNil() {
	o.RequestMigrate.Set(nil)
}

// UnsetRequestMigrate ensures that no value is present for RequestMigrate, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetRequestMigrate() {
	o.RequestMigrate.Unset()
}

// GetRequestRecover returns the RequestRecover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetRequestRecover() AuthorizationLevel {
	if o == nil || utils.IsNil(o.RequestRecover.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.RequestRecover.Get()
}

// GetRequestRecoverOk returns a tuple with the RequestRecover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetRequestRecoverOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestRecover.Get(), o.RequestRecover.IsSet()
}

// HasRequestRecover returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasRequestRecover() bool {
	if o != nil && o.RequestRecover.IsSet() {
		return true
	}

	return false
}

// SetRequestRecover gets a reference to the given NullableAuthorizationLevel and assigns it to the RequestRecover field.
func (o *CertificateProfileAuthorizationLevels) SetRequestRecover(v AuthorizationLevel) {
	o.RequestRecover.Set(&v)
}

// SetRequestRecoverNil sets the value for RequestRecover to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetRequestRecoverNil() {
	o.RequestRecover.Set(nil)
}

// UnsetRequestRecover ensures that no value is present for RequestRecover, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetRequestRecover() {
	o.RequestRecover.Unset()
}

// GetRequestRenew returns the RequestRenew field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateProfileAuthorizationLevels) GetRequestRenew() AuthorizationLevel {
	if o == nil || utils.IsNil(o.RequestRenew.Get()) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.RequestRenew.Get()
}

// GetRequestRenewOk returns a tuple with the RequestRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateProfileAuthorizationLevels) GetRequestRenewOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestRenew.Get(), o.RequestRenew.IsSet()
}

// HasRequestRenew returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasRequestRenew() bool {
	if o != nil && o.RequestRenew.IsSet() {
		return true
	}

	return false
}

// SetRequestRenew gets a reference to the given NullableAuthorizationLevel and assigns it to the RequestRenew field.
func (o *CertificateProfileAuthorizationLevels) SetRequestRenew(v AuthorizationLevel) {
	o.RequestRenew.Set(&v)
}

// SetRequestRenewNil sets the value for RequestRenew to be an explicit nil
func (o *CertificateProfileAuthorizationLevels) SetRequestRenewNil() {
	o.RequestRenew.Set(nil)
}

// UnsetRequestRenew ensures that no value is present for RequestRenew, not even an explicit nil
func (o *CertificateProfileAuthorizationLevels) UnsetRequestRenew() {
	o.RequestRenew.Unset()
}

// GetRequestRevoke returns the RequestRevoke field value if set, zero value otherwise.
func (o *CertificateProfileAuthorizationLevels) GetRequestRevoke() AuthorizationLevel {
	if o == nil || utils.IsNil(o.RequestRevoke) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.RequestRevoke
}

// GetRequestRevokeOk returns a tuple with the RequestRevoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateProfileAuthorizationLevels) GetRequestRevokeOk() (*AuthorizationLevel, bool) {
	if o == nil || utils.IsNil(o.RequestRevoke) {
		return nil, false
	}
	return o.RequestRevoke, true
}

// HasRequestRevoke returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasRequestRevoke() bool {
	if o != nil && !utils.IsNil(o.RequestRevoke) {
		return true
	}

	return false
}

// SetRequestRevoke gets a reference to the given AuthorizationLevel and assigns it to the RequestRevoke field.
func (o *CertificateProfileAuthorizationLevels) SetRequestRevoke(v AuthorizationLevel) {
	o.RequestRevoke = &v
}

// GetRequestUpdate returns the RequestUpdate field value
func (o *CertificateProfileAuthorizationLevels) GetRequestUpdate() AuthorizationLevel {
	if o == nil {
		var ret AuthorizationLevel
		return ret
	}

	return o.RequestUpdate
}

// GetRequestUpdateOk returns a tuple with the RequestUpdate field value
// and a boolean to check if the value has been set.
func (o *CertificateProfileAuthorizationLevels) GetRequestUpdateOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestUpdate, true
}

// SetRequestUpdate sets field value
func (o *CertificateProfileAuthorizationLevels) SetRequestUpdate(v AuthorizationLevel) {
	o.RequestUpdate = v
}

// GetRevoke returns the Revoke field value if set, zero value otherwise.
func (o *CertificateProfileAuthorizationLevels) GetRevoke() AuthorizationLevel {
	if o == nil || utils.IsNil(o.Revoke) {
		var ret AuthorizationLevel
		return ret
	}
	return *o.Revoke
}

// GetRevokeOk returns a tuple with the Revoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateProfileAuthorizationLevels) GetRevokeOk() (*AuthorizationLevel, bool) {
	if o == nil || utils.IsNil(o.Revoke) {
		return nil, false
	}
	return o.Revoke, true
}

// HasRevoke returns a boolean if a field has been set.
func (o *CertificateProfileAuthorizationLevels) HasRevoke() bool {
	if o != nil && !utils.IsNil(o.Revoke) {
		return true
	}

	return false
}

// SetRevoke gets a reference to the given AuthorizationLevel and assigns it to the Revoke field.
func (o *CertificateProfileAuthorizationLevels) SetRevoke(v AuthorizationLevel) {
	o.Revoke = &v
}

// GetSearch returns the Search field value
func (o *CertificateProfileAuthorizationLevels) GetSearch() AuthorizationLevel {
	if o == nil {
		var ret AuthorizationLevel
		return ret
	}

	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *CertificateProfileAuthorizationLevels) GetSearchOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value
func (o *CertificateProfileAuthorizationLevels) SetSearch(v AuthorizationLevel) {
	o.Search = v
}

// GetUpdate returns the Update field value
func (o *CertificateProfileAuthorizationLevels) GetUpdate() AuthorizationLevel {
	if o == nil {
		var ret AuthorizationLevel
		return ret
	}

	return o.Update
}

// GetUpdateOk returns a tuple with the Update field value
// and a boolean to check if the value has been set.
func (o *CertificateProfileAuthorizationLevels) GetUpdateOk() (*AuthorizationLevel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Update, true
}

// SetUpdate sets field value
func (o *CertificateProfileAuthorizationLevels) SetUpdate(v AuthorizationLevel) {
	o.Update = v
}

func (o CertificateProfileAuthorizationLevels) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateProfileAuthorizationLevels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApproveEnroll.IsSet() {
		toSerialize["approveEnroll"] = o.ApproveEnroll.Get()
	}
	if o.ApproveMigrate.IsSet() {
		toSerialize["approveMigrate"] = o.ApproveMigrate.Get()
	}
	if o.ApproveRecover.IsSet() {
		toSerialize["approveRecover"] = o.ApproveRecover.Get()
	}
	if o.ApproveRenew.IsSet() {
		toSerialize["approveRenew"] = o.ApproveRenew.Get()
	}
	if !utils.IsNil(o.ApproveRevoke) {
		toSerialize["approveRevoke"] = o.ApproveRevoke
	}
	toSerialize["approveUpdate"] = o.ApproveUpdate
	if o.AuditRequest.IsSet() {
		toSerialize["auditRequest"] = o.AuditRequest.Get()
	}
	if o.Enroll.IsSet() {
		toSerialize["enroll"] = o.Enroll.Get()
	}
	if o.EnrollApi.IsSet() {
		toSerialize["enrollApi"] = o.EnrollApi.Get()
	}
	if o.Migrate.IsSet() {
		toSerialize["migrate"] = o.Migrate.Get()
	}
	if o.Recover.IsSet() {
		toSerialize["recover"] = o.Recover.Get()
	}
	if o.RecoverApi.IsSet() {
		toSerialize["recoverApi"] = o.RecoverApi.Get()
	}
	if o.Renew.IsSet() {
		toSerialize["renew"] = o.Renew.Get()
	}
	if o.RenewApi.IsSet() {
		toSerialize["renewApi"] = o.RenewApi.Get()
	}
	if o.RequestEnroll.IsSet() {
		toSerialize["requestEnroll"] = o.RequestEnroll.Get()
	}
	if o.RequestMigrate.IsSet() {
		toSerialize["requestMigrate"] = o.RequestMigrate.Get()
	}
	if o.RequestRecover.IsSet() {
		toSerialize["requestRecover"] = o.RequestRecover.Get()
	}
	if o.RequestRenew.IsSet() {
		toSerialize["requestRenew"] = o.RequestRenew.Get()
	}
	if !utils.IsNil(o.RequestRevoke) {
		toSerialize["requestRevoke"] = o.RequestRevoke
	}
	toSerialize["requestUpdate"] = o.RequestUpdate
	if !utils.IsNil(o.Revoke) {
		toSerialize["revoke"] = o.Revoke
	}
	toSerialize["search"] = o.Search
	toSerialize["update"] = o.Update

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateProfileAuthorizationLevels) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approveUpdate",
		"requestUpdate",
		"search",
		"update",
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

	varCertificateProfileAuthorizationLevels := _CertificateProfileAuthorizationLevels{}

	err = json.Unmarshal(data, &varCertificateProfileAuthorizationLevels)

	if err != nil {
		return err
	}

	*o = CertificateProfileAuthorizationLevels(varCertificateProfileAuthorizationLevels)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "approveEnroll")
		delete(additionalProperties, "approveMigrate")
		delete(additionalProperties, "approveRecover")
		delete(additionalProperties, "approveRenew")
		delete(additionalProperties, "approveRevoke")
		delete(additionalProperties, "approveUpdate")
		delete(additionalProperties, "auditRequest")
		delete(additionalProperties, "enroll")
		delete(additionalProperties, "enrollApi")
		delete(additionalProperties, "migrate")
		delete(additionalProperties, "recover")
		delete(additionalProperties, "recoverApi")
		delete(additionalProperties, "renew")
		delete(additionalProperties, "renewApi")
		delete(additionalProperties, "requestEnroll")
		delete(additionalProperties, "requestMigrate")
		delete(additionalProperties, "requestRecover")
		delete(additionalProperties, "requestRenew")
		delete(additionalProperties, "requestRevoke")
		delete(additionalProperties, "requestUpdate")
		delete(additionalProperties, "revoke")
		delete(additionalProperties, "search")
		delete(additionalProperties, "update")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateProfileAuthorizationLevels struct {
	value *CertificateProfileAuthorizationLevels
	isSet bool
}

func (v NullableCertificateProfileAuthorizationLevels) Get() *CertificateProfileAuthorizationLevels {
	return v.value
}

func (v *NullableCertificateProfileAuthorizationLevels) Set(val *CertificateProfileAuthorizationLevels) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateProfileAuthorizationLevels) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateProfileAuthorizationLevels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateProfileAuthorizationLevels(val *CertificateProfileAuthorizationLevels) *NullableCertificateProfileAuthorizationLevels {
	return &NullableCertificateProfileAuthorizationLevels{value: val, isSet: true}
}

func (v NullableCertificateProfileAuthorizationLevels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateProfileAuthorizationLevels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
