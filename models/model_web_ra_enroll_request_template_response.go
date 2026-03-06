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

// checks if the WebRAEnrollRequestTemplateResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebRAEnrollRequestTemplateResponse{}

// WebRAEnrollRequestTemplateResponse struct for WebRAEnrollRequestTemplateResponse
type WebRAEnrollRequestTemplateResponse struct {
	// Describes how certificates will be enrolled on this profile
	Capabilities NullableManagedCertificateProfileCryptoPolicy `json:"capabilities,omitempty"`
	// Information about the certificate's contact email and how to edit it
	ContactEmail NullableCertificateContactEmailElementResponse `json:"contactEmail,omitempty"`
	// Information about the certificate's extensions and how to edit them
	Extensions []CertificateExtensionElementResponse `json:"extensions,omitempty"`
	// List of labels used internally to tag and group certificates
	Labels []RequestLabelElementResponse `json:"labels,omitempty"`
	// The technical metadata for this certificate
	Metadata []CertificateMetadataElementResponse `json:"metadata,omitempty"`
	// Information about the certificate's owner and how to edit it
	Owner NullableCertificateOwnerElementResponse `json:"owner,omitempty"`
	// The password policy that will be used to generate the certificate's PKCS#12 password
	PasswordPolicy NullablePasswordPolicy `json:"passwordPolicy,omitempty"`
	// List of SAN elements that will be used to build the certificate's Subject Alternative Name
	Sans []ListSANElementResponse `json:"sans,omitempty"`
	// List of DN elements that will be used to build the certificate's Distinguished Name
	Subject []IndexedDNElementResponse `json:"subject,omitempty"`
	// Information about the certificate's team and how to edit it
	Team                 NullableCertificateTeamElementResponse `json:"team,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebRAEnrollRequestTemplateResponse WebRAEnrollRequestTemplateResponse

// NewWebRAEnrollRequestTemplateResponse instantiates a new WebRAEnrollRequestTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebRAEnrollRequestTemplateResponse() *WebRAEnrollRequestTemplateResponse {
	this := WebRAEnrollRequestTemplateResponse{}
	return &this
}

// NewWebRAEnrollRequestTemplateResponseWithDefaults instantiates a new WebRAEnrollRequestTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebRAEnrollRequestTemplateResponseWithDefaults() *WebRAEnrollRequestTemplateResponse {
	this := WebRAEnrollRequestTemplateResponse{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplateResponse) GetCapabilities() ManagedCertificateProfileCryptoPolicy {
	if o == nil || utils.IsNil(o.Capabilities.Get()) {
		var ret ManagedCertificateProfileCryptoPolicy
		return ret
	}
	return *o.Capabilities.Get()
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplateResponse) GetCapabilitiesOk() (*ManagedCertificateProfileCryptoPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities.Get(), o.Capabilities.IsSet()
}

// HasCapabilities returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplateResponse) HasCapabilities() bool {
	if o != nil && o.Capabilities.IsSet() {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given NullableManagedCertificateProfileCryptoPolicy and assigns it to the Capabilities field.
func (o *WebRAEnrollRequestTemplateResponse) SetCapabilities(v ManagedCertificateProfileCryptoPolicy) {
	o.Capabilities.Set(&v)
}

// SetCapabilitiesNil sets the value for Capabilities to be an explicit nil
func (o *WebRAEnrollRequestTemplateResponse) SetCapabilitiesNil() {
	o.Capabilities.Set(nil)
}

// UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
func (o *WebRAEnrollRequestTemplateResponse) UnsetCapabilities() {
	o.Capabilities.Unset()
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplateResponse) GetContactEmail() CertificateContactEmailElementResponse {
	if o == nil || utils.IsNil(o.ContactEmail.Get()) {
		var ret CertificateContactEmailElementResponse
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplateResponse) GetContactEmailOk() (*CertificateContactEmailElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplateResponse) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableCertificateContactEmailElementResponse and assigns it to the ContactEmail field.
func (o *WebRAEnrollRequestTemplateResponse) SetContactEmail(v CertificateContactEmailElementResponse) {
	o.ContactEmail.Set(&v)
}

// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *WebRAEnrollRequestTemplateResponse) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *WebRAEnrollRequestTemplateResponse) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplateResponse) GetExtensions() []CertificateExtensionElementResponse {
	if o == nil {
		var ret []CertificateExtensionElementResponse
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplateResponse) GetExtensionsOk() ([]CertificateExtensionElementResponse, bool) {
	if o == nil || utils.IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplateResponse) HasExtensions() bool {
	if o != nil && !utils.IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []CertificateExtensionElementResponse and assigns it to the Extensions field.
func (o *WebRAEnrollRequestTemplateResponse) SetExtensions(v []CertificateExtensionElementResponse) {
	o.Extensions = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplateResponse) GetLabels() []RequestLabelElementResponse {
	if o == nil {
		var ret []RequestLabelElementResponse
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplateResponse) GetLabelsOk() ([]RequestLabelElementResponse, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplateResponse) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []RequestLabelElementResponse and assigns it to the Labels field.
func (o *WebRAEnrollRequestTemplateResponse) SetLabels(v []RequestLabelElementResponse) {
	o.Labels = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplateResponse) GetMetadata() []CertificateMetadataElementResponse {
	if o == nil {
		var ret []CertificateMetadataElementResponse
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplateResponse) GetMetadataOk() ([]CertificateMetadataElementResponse, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplateResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []CertificateMetadataElementResponse and assigns it to the Metadata field.
func (o *WebRAEnrollRequestTemplateResponse) SetMetadata(v []CertificateMetadataElementResponse) {
	o.Metadata = v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplateResponse) GetOwner() CertificateOwnerElementResponse {
	if o == nil || utils.IsNil(o.Owner.Get()) {
		var ret CertificateOwnerElementResponse
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplateResponse) GetOwnerOk() (*CertificateOwnerElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplateResponse) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableCertificateOwnerElementResponse and assigns it to the Owner field.
func (o *WebRAEnrollRequestTemplateResponse) SetOwner(v CertificateOwnerElementResponse) {
	o.Owner.Set(&v)
}

// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *WebRAEnrollRequestTemplateResponse) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *WebRAEnrollRequestTemplateResponse) UnsetOwner() {
	o.Owner.Unset()
}

// GetPasswordPolicy returns the PasswordPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplateResponse) GetPasswordPolicy() PasswordPolicy {
	if o == nil || utils.IsNil(o.PasswordPolicy.Get()) {
		var ret PasswordPolicy
		return ret
	}
	return *o.PasswordPolicy.Get()
}

// GetPasswordPolicyOk returns a tuple with the PasswordPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplateResponse) GetPasswordPolicyOk() (*PasswordPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordPolicy.Get(), o.PasswordPolicy.IsSet()
}

// HasPasswordPolicy returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplateResponse) HasPasswordPolicy() bool {
	if o != nil && o.PasswordPolicy.IsSet() {
		return true
	}

	return false
}

// SetPasswordPolicy gets a reference to the given NullablePasswordPolicy and assigns it to the PasswordPolicy field.
func (o *WebRAEnrollRequestTemplateResponse) SetPasswordPolicy(v PasswordPolicy) {
	o.PasswordPolicy.Set(&v)
}

// SetPasswordPolicyNil sets the value for PasswordPolicy to be an explicit nil
func (o *WebRAEnrollRequestTemplateResponse) SetPasswordPolicyNil() {
	o.PasswordPolicy.Set(nil)
}

// UnsetPasswordPolicy ensures that no value is present for PasswordPolicy, not even an explicit nil
func (o *WebRAEnrollRequestTemplateResponse) UnsetPasswordPolicy() {
	o.PasswordPolicy.Unset()
}

// GetSans returns the Sans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplateResponse) GetSans() []ListSANElementResponse {
	if o == nil {
		var ret []ListSANElementResponse
		return ret
	}
	return o.Sans
}

// GetSansOk returns a tuple with the Sans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplateResponse) GetSansOk() ([]ListSANElementResponse, bool) {
	if o == nil || utils.IsNil(o.Sans) {
		return nil, false
	}
	return o.Sans, true
}

// HasSans returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplateResponse) HasSans() bool {
	if o != nil && !utils.IsNil(o.Sans) {
		return true
	}

	return false
}

// SetSans gets a reference to the given []ListSANElementResponse and assigns it to the Sans field.
func (o *WebRAEnrollRequestTemplateResponse) SetSans(v []ListSANElementResponse) {
	o.Sans = v
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplateResponse) GetSubject() []IndexedDNElementResponse {
	if o == nil {
		var ret []IndexedDNElementResponse
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplateResponse) GetSubjectOk() ([]IndexedDNElementResponse, bool) {
	if o == nil || utils.IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplateResponse) HasSubject() bool {
	if o != nil && !utils.IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given []IndexedDNElementResponse and assigns it to the Subject field.
func (o *WebRAEnrollRequestTemplateResponse) SetSubject(v []IndexedDNElementResponse) {
	o.Subject = v
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplateResponse) GetTeam() CertificateTeamElementResponse {
	if o == nil || utils.IsNil(o.Team.Get()) {
		var ret CertificateTeamElementResponse
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplateResponse) GetTeamOk() (*CertificateTeamElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplateResponse) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableCertificateTeamElementResponse and assigns it to the Team field.
func (o *WebRAEnrollRequestTemplateResponse) SetTeam(v CertificateTeamElementResponse) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *WebRAEnrollRequestTemplateResponse) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *WebRAEnrollRequestTemplateResponse) UnsetTeam() {
	o.Team.Unset()
}

func (o WebRAEnrollRequestTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebRAEnrollRequestTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Capabilities.IsSet() {
		toSerialize["capabilities"] = o.Capabilities.Get()
	}
	if o.ContactEmail.IsSet() {
		toSerialize["contactEmail"] = o.ContactEmail.Get()
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.PasswordPolicy.IsSet() {
		toSerialize["passwordPolicy"] = o.PasswordPolicy.Get()
	}
	if o.Sans != nil {
		toSerialize["sans"] = o.Sans
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebRAEnrollRequestTemplateResponse) UnmarshalJSON(data []byte) (err error) {
	varWebRAEnrollRequestTemplateResponse := _WebRAEnrollRequestTemplateResponse{}

	err = json.Unmarshal(data, &varWebRAEnrollRequestTemplateResponse)

	if err != nil {
		return err
	}

	*o = WebRAEnrollRequestTemplateResponse(varWebRAEnrollRequestTemplateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capabilities")
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "extensions")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "passwordPolicy")
		delete(additionalProperties, "sans")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "team")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebRAEnrollRequestTemplateResponse struct {
	value *WebRAEnrollRequestTemplateResponse
	isSet bool
}

func (v NullableWebRAEnrollRequestTemplateResponse) Get() *WebRAEnrollRequestTemplateResponse {
	return v.value
}

func (v *NullableWebRAEnrollRequestTemplateResponse) Set(val *WebRAEnrollRequestTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebRAEnrollRequestTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebRAEnrollRequestTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebRAEnrollRequestTemplateResponse(val *WebRAEnrollRequestTemplateResponse) *NullableWebRAEnrollRequestTemplateResponse {
	return &NullableWebRAEnrollRequestTemplateResponse{value: val, isSet: true}
}

func (v NullableWebRAEnrollRequestTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebRAEnrollRequestTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
