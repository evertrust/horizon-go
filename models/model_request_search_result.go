/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the RequestSearchResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RequestSearchResult{}

// RequestSearchResult struct for RequestSearchResult
type RequestSearchResult struct {
	// The ID of the request
	Id string `json:"_id"`
	// The certificate associated with the request
	Certificate NullableCertificate `json:"certificate,omitempty"`
	// The id of the certificate in the request
	CertificateId utils.NullableString `json:"certificateId,omitempty"`
	// Associated certificate's Distinguished Name
	Dn *string `json:"dn,omitempty"`
	// The computed holderID for this request. This is set by the system based on DN and SANs
	HolderId *string `json:"holderId,omitempty"`
	Module   *Module `json:"module,omitempty"`
	// The permissions of the principal on this request.
	Permissions RequestPermissions `json:"permissions"`
	// Any profile configured for a protocol in Horizon
	Profile *string `json:"profile,omitempty"`
	// Free-text field editable by the requester to provider more context on the request
	RequesterComment     utils.NullableString `json:"requesterComment,omitempty"`
	Status               *RequestStatus       `json:"status,omitempty"`
	Workflow             *Workflow            `json:"workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestSearchResult RequestSearchResult

// NewRequestSearchResult instantiates a new RequestSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSearchResult(id string, permissions RequestPermissions) *RequestSearchResult {
	this := RequestSearchResult{}
	this.Id = id
	this.Permissions = permissions
	return &this
}

// NewRequestSearchResultWithDefaults instantiates a new RequestSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSearchResultWithDefaults() *RequestSearchResult {
	this := RequestSearchResult{}
	return &this
}

// GetId returns the Id field value
func (o *RequestSearchResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RequestSearchResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RequestSearchResult) SetId(v string) {
	o.Id = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSearchResult) GetCertificate() Certificate {
	if o == nil || utils.IsNil(o.Certificate.Get()) {
		var ret Certificate
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSearchResult) GetCertificateOk() (*Certificate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *RequestSearchResult) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableCertificate and assigns it to the Certificate field.
func (o *RequestSearchResult) SetCertificate(v Certificate) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *RequestSearchResult) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *RequestSearchResult) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetCertificateId returns the CertificateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSearchResult) GetCertificateId() string {
	if o == nil || utils.IsNil(o.CertificateId.Get()) {
		var ret string
		return ret
	}
	return *o.CertificateId.Get()
}

// GetCertificateIdOk returns a tuple with the CertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSearchResult) GetCertificateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertificateId.Get(), o.CertificateId.IsSet()
}

// HasCertificateId returns a boolean if a field has been set.
func (o *RequestSearchResult) HasCertificateId() bool {
	if o != nil && o.CertificateId.IsSet() {
		return true
	}

	return false
}

// SetCertificateId gets a reference to the given NullableString and assigns it to the CertificateId field.
func (o *RequestSearchResult) SetCertificateId(v string) {
	o.CertificateId.Set(&v)
}

// SetCertificateIdNil sets the value for CertificateId to be an explicit nil
func (o *RequestSearchResult) SetCertificateIdNil() {
	o.CertificateId.Set(nil)
}

// UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
func (o *RequestSearchResult) UnsetCertificateId() {
	o.CertificateId.Unset()
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *RequestSearchResult) GetDn() string {
	if o == nil || utils.IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSearchResult) GetDnOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *RequestSearchResult) HasDn() bool {
	if o != nil && !utils.IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *RequestSearchResult) SetDn(v string) {
	o.Dn = &v
}

// GetHolderId returns the HolderId field value if set, zero value otherwise.
func (o *RequestSearchResult) GetHolderId() string {
	if o == nil || utils.IsNil(o.HolderId) {
		var ret string
		return ret
	}
	return *o.HolderId
}

// GetHolderIdOk returns a tuple with the HolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSearchResult) GetHolderIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.HolderId) {
		return nil, false
	}
	return o.HolderId, true
}

// HasHolderId returns a boolean if a field has been set.
func (o *RequestSearchResult) HasHolderId() bool {
	if o != nil && !utils.IsNil(o.HolderId) {
		return true
	}

	return false
}

// SetHolderId gets a reference to the given string and assigns it to the HolderId field.
func (o *RequestSearchResult) SetHolderId(v string) {
	o.HolderId = &v
}

// GetModule returns the Module field value if set, zero value otherwise.
func (o *RequestSearchResult) GetModule() Module {
	if o == nil || utils.IsNil(o.Module) {
		var ret Module
		return ret
	}
	return *o.Module
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSearchResult) GetModuleOk() (*Module, bool) {
	if o == nil || utils.IsNil(o.Module) {
		return nil, false
	}
	return o.Module, true
}

// HasModule returns a boolean if a field has been set.
func (o *RequestSearchResult) HasModule() bool {
	if o != nil && !utils.IsNil(o.Module) {
		return true
	}

	return false
}

// SetModule gets a reference to the given Module and assigns it to the Module field.
func (o *RequestSearchResult) SetModule(v Module) {
	o.Module = &v
}

// GetPermissions returns the Permissions field value
func (o *RequestSearchResult) GetPermissions() RequestPermissions {
	if o == nil {
		var ret RequestPermissions
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *RequestSearchResult) GetPermissionsOk() (*RequestPermissions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Permissions, true
}

// SetPermissions sets field value
func (o *RequestSearchResult) SetPermissions(v RequestPermissions) {
	o.Permissions = v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *RequestSearchResult) GetProfile() string {
	if o == nil || utils.IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSearchResult) GetProfileOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *RequestSearchResult) HasProfile() bool {
	if o != nil && !utils.IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *RequestSearchResult) SetProfile(v string) {
	o.Profile = &v
}

// GetRequesterComment returns the RequesterComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestSearchResult) GetRequesterComment() string {
	if o == nil || utils.IsNil(o.RequesterComment.Get()) {
		var ret string
		return ret
	}
	return *o.RequesterComment.Get()
}

// GetRequesterCommentOk returns a tuple with the RequesterComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestSearchResult) GetRequesterCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterComment.Get(), o.RequesterComment.IsSet()
}

// HasRequesterComment returns a boolean if a field has been set.
func (o *RequestSearchResult) HasRequesterComment() bool {
	if o != nil && o.RequesterComment.IsSet() {
		return true
	}

	return false
}

// SetRequesterComment gets a reference to the given NullableString and assigns it to the RequesterComment field.
func (o *RequestSearchResult) SetRequesterComment(v string) {
	o.RequesterComment.Set(&v)
}

// SetRequesterCommentNil sets the value for RequesterComment to be an explicit nil
func (o *RequestSearchResult) SetRequesterCommentNil() {
	o.RequesterComment.Set(nil)
}

// UnsetRequesterComment ensures that no value is present for RequesterComment, not even an explicit nil
func (o *RequestSearchResult) UnsetRequesterComment() {
	o.RequesterComment.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RequestSearchResult) GetStatus() RequestStatus {
	if o == nil || utils.IsNil(o.Status) {
		var ret RequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSearchResult) GetStatusOk() (*RequestStatus, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RequestSearchResult) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given RequestStatus and assigns it to the Status field.
func (o *RequestSearchResult) SetStatus(v RequestStatus) {
	o.Status = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *RequestSearchResult) GetWorkflow() Workflow {
	if o == nil || utils.IsNil(o.Workflow) {
		var ret Workflow
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSearchResult) GetWorkflowOk() (*Workflow, bool) {
	if o == nil || utils.IsNil(o.Workflow) {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *RequestSearchResult) HasWorkflow() bool {
	if o != nil && !utils.IsNil(o.Workflow) {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given Workflow and assigns it to the Workflow field.
func (o *RequestSearchResult) SetWorkflow(v Workflow) {
	o.Workflow = &v
}

func (o RequestSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.CertificateId.IsSet() {
		toSerialize["certificateId"] = o.CertificateId.Get()
	}
	if !utils.IsNil(o.Dn) {
		toSerialize["dn"] = o.Dn
	}
	if !utils.IsNil(o.HolderId) {
		toSerialize["holderId"] = o.HolderId
	}
	if !utils.IsNil(o.Module) {
		toSerialize["module"] = o.Module
	}
	toSerialize["permissions"] = o.Permissions
	if !utils.IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if o.RequesterComment.IsSet() {
		toSerialize["requesterComment"] = o.RequesterComment.Get()
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.Workflow) {
		toSerialize["workflow"] = o.Workflow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestSearchResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"permissions",
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

	varRequestSearchResult := _RequestSearchResult{}

	err = json.Unmarshal(data, &varRequestSearchResult)

	if err != nil {
		return err
	}

	*o = RequestSearchResult(varRequestSearchResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "certificateId")
		delete(additionalProperties, "dn")
		delete(additionalProperties, "holderId")
		delete(additionalProperties, "module")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "requesterComment")
		delete(additionalProperties, "status")
		delete(additionalProperties, "workflow")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestSearchResult struct {
	value *RequestSearchResult
	isSet bool
}

func (v NullableRequestSearchResult) Get() *RequestSearchResult {
	return v.value
}

func (v *NullableRequestSearchResult) Set(val *RequestSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSearchResult(val *RequestSearchResult) *NullableRequestSearchResult {
	return &NullableRequestSearchResult{value: val, isSet: true}
}

func (v NullableRequestSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
