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

// checks if the WebRAEnrollRequestTemplate type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebRAEnrollRequestTemplate{}

// WebRAEnrollRequestTemplate struct for WebRAEnrollRequestTemplate
type WebRAEnrollRequestTemplate struct {
	// If decentralized enrollment is enabled, this field will contain the CSR that will be used to generate the certificate
	Csr utils.NullableString `json:"csr,omitempty"`
	// The type of key that will be used to generate the certificate, if in centralized mode
	KeyType utils.NullableString `json:"keyType,omitempty" validate:"regexp=(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512)(\\\\\\\\+(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512))?"`
	// Information about the certificate's contact email and how to edit it
	ContactEmail NullableCertificateContactEmailElement `json:"contactEmail,omitempty"`
	// Information about the certificate's extensions and how to edit them
	Extensions []CertificateExtensionElement `json:"extensions,omitempty"`
	// List of labels used internally to tag and group certificates
	Labels []RequestLabelElement `json:"labels,omitempty"`
	// The technical metadata for this certificate
	Metadata []CertificateMetadataElement `json:"metadata,omitempty"`
	// Information about the certificate's owner and how to edit it
	Owner NullableCertificateOwnerElement `json:"owner,omitempty"`
	// List of SAN elements that will be used to build the certificate's Subject Alternative Name
	Sans []ListSANElement `json:"sans,omitempty"`
	// List of DN elements that will be used to build the certificate's Distinguished Name
	Subject []IndexedDNElement `json:"subject,omitempty"`
	// Information about the certificate's team and how to edit it
	Team                 NullableCertificateTeamElement `json:"team,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebRAEnrollRequestTemplate WebRAEnrollRequestTemplate

// NewWebRAEnrollRequestTemplate instantiates a new WebRAEnrollRequestTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebRAEnrollRequestTemplate() *WebRAEnrollRequestTemplate {
	this := WebRAEnrollRequestTemplate{}
	return &this
}

// NewWebRAEnrollRequestTemplateWithDefaults instantiates a new WebRAEnrollRequestTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebRAEnrollRequestTemplateWithDefaults() *WebRAEnrollRequestTemplate {
	this := WebRAEnrollRequestTemplate{}
	return &this
}

// GetCsr returns the Csr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplate) GetCsr() string {
	if o == nil || utils.IsNil(o.Csr.Get()) {
		var ret string
		return ret
	}
	return *o.Csr.Get()
}

// GetCsrOk returns a tuple with the Csr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplate) GetCsrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Csr.Get(), o.Csr.IsSet()
}

// HasCsr returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplate) HasCsr() bool {
	if o != nil && o.Csr.IsSet() {
		return true
	}

	return false
}

// SetCsr gets a reference to the given NullableString and assigns it to the Csr field.
func (o *WebRAEnrollRequestTemplate) SetCsr(v string) {
	o.Csr.Set(&v)
}

// SetCsrNil sets the value for Csr to be an explicit nil
func (o *WebRAEnrollRequestTemplate) SetCsrNil() {
	o.Csr.Set(nil)
}

// UnsetCsr ensures that no value is present for Csr, not even an explicit nil
func (o *WebRAEnrollRequestTemplate) UnsetCsr() {
	o.Csr.Unset()
}

// GetKeyType returns the KeyType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplate) GetKeyType() string {
	if o == nil || utils.IsNil(o.KeyType.Get()) {
		var ret string
		return ret
	}
	return *o.KeyType.Get()
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplate) GetKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyType.Get(), o.KeyType.IsSet()
}

// HasKeyType returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplate) HasKeyType() bool {
	if o != nil && o.KeyType.IsSet() {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given NullableString and assigns it to the KeyType field.
func (o *WebRAEnrollRequestTemplate) SetKeyType(v string) {
	o.KeyType.Set(&v)
}

// SetKeyTypeNil sets the value for KeyType to be an explicit nil
func (o *WebRAEnrollRequestTemplate) SetKeyTypeNil() {
	o.KeyType.Set(nil)
}

// UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
func (o *WebRAEnrollRequestTemplate) UnsetKeyType() {
	o.KeyType.Unset()
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplate) GetContactEmail() CertificateContactEmailElement {
	if o == nil || utils.IsNil(o.ContactEmail.Get()) {
		var ret CertificateContactEmailElement
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplate) GetContactEmailOk() (*CertificateContactEmailElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplate) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableCertificateContactEmailElement and assigns it to the ContactEmail field.
func (o *WebRAEnrollRequestTemplate) SetContactEmail(v CertificateContactEmailElement) {
	o.ContactEmail.Set(&v)
}

// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *WebRAEnrollRequestTemplate) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *WebRAEnrollRequestTemplate) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplate) GetExtensions() []CertificateExtensionElement {
	if o == nil {
		var ret []CertificateExtensionElement
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplate) GetExtensionsOk() ([]CertificateExtensionElement, bool) {
	if o == nil || utils.IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplate) HasExtensions() bool {
	if o != nil && !utils.IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []CertificateExtensionElement and assigns it to the Extensions field.
func (o *WebRAEnrollRequestTemplate) SetExtensions(v []CertificateExtensionElement) {
	o.Extensions = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplate) GetLabels() []RequestLabelElement {
	if o == nil {
		var ret []RequestLabelElement
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplate) GetLabelsOk() ([]RequestLabelElement, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplate) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []RequestLabelElement and assigns it to the Labels field.
func (o *WebRAEnrollRequestTemplate) SetLabels(v []RequestLabelElement) {
	o.Labels = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplate) GetMetadata() []CertificateMetadataElement {
	if o == nil {
		var ret []CertificateMetadataElement
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplate) GetMetadataOk() ([]CertificateMetadataElement, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplate) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []CertificateMetadataElement and assigns it to the Metadata field.
func (o *WebRAEnrollRequestTemplate) SetMetadata(v []CertificateMetadataElement) {
	o.Metadata = v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplate) GetOwner() CertificateOwnerElement {
	if o == nil || utils.IsNil(o.Owner.Get()) {
		var ret CertificateOwnerElement
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplate) GetOwnerOk() (*CertificateOwnerElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplate) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableCertificateOwnerElement and assigns it to the Owner field.
func (o *WebRAEnrollRequestTemplate) SetOwner(v CertificateOwnerElement) {
	o.Owner.Set(&v)
}

// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *WebRAEnrollRequestTemplate) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *WebRAEnrollRequestTemplate) UnsetOwner() {
	o.Owner.Unset()
}

// GetSans returns the Sans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplate) GetSans() []ListSANElement {
	if o == nil {
		var ret []ListSANElement
		return ret
	}
	return o.Sans
}

// GetSansOk returns a tuple with the Sans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplate) GetSansOk() ([]ListSANElement, bool) {
	if o == nil || utils.IsNil(o.Sans) {
		return nil, false
	}
	return o.Sans, true
}

// HasSans returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplate) HasSans() bool {
	if o != nil && !utils.IsNil(o.Sans) {
		return true
	}

	return false
}

// SetSans gets a reference to the given []ListSANElement and assigns it to the Sans field.
func (o *WebRAEnrollRequestTemplate) SetSans(v []ListSANElement) {
	o.Sans = v
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplate) GetSubject() []IndexedDNElement {
	if o == nil {
		var ret []IndexedDNElement
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplate) GetSubjectOk() ([]IndexedDNElement, bool) {
	if o == nil || utils.IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplate) HasSubject() bool {
	if o != nil && !utils.IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given []IndexedDNElement and assigns it to the Subject field.
func (o *WebRAEnrollRequestTemplate) SetSubject(v []IndexedDNElement) {
	o.Subject = v
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAEnrollRequestTemplate) GetTeam() CertificateTeamElement {
	if o == nil || utils.IsNil(o.Team.Get()) {
		var ret CertificateTeamElement
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAEnrollRequestTemplate) GetTeamOk() (*CertificateTeamElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *WebRAEnrollRequestTemplate) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableCertificateTeamElement and assigns it to the Team field.
func (o *WebRAEnrollRequestTemplate) SetTeam(v CertificateTeamElement) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *WebRAEnrollRequestTemplate) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *WebRAEnrollRequestTemplate) UnsetTeam() {
	o.Team.Unset()
}

func (o WebRAEnrollRequestTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebRAEnrollRequestTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Csr.IsSet() {
		toSerialize["csr"] = o.Csr.Get()
	}
	if o.KeyType.IsSet() {
		toSerialize["keyType"] = o.KeyType.Get()
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

func (o *WebRAEnrollRequestTemplate) UnmarshalJSON(data []byte) (err error) {
	varWebRAEnrollRequestTemplate := _WebRAEnrollRequestTemplate{}

	err = json.Unmarshal(data, &varWebRAEnrollRequestTemplate)

	if err != nil {
		return err
	}

	*o = WebRAEnrollRequestTemplate(varWebRAEnrollRequestTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "csr")
		delete(additionalProperties, "keyType")
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "extensions")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "sans")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "team")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebRAEnrollRequestTemplate struct {
	value *WebRAEnrollRequestTemplate
	isSet bool
}

func (v NullableWebRAEnrollRequestTemplate) Get() *WebRAEnrollRequestTemplate {
	return v.value
}

func (v *NullableWebRAEnrollRequestTemplate) Set(val *WebRAEnrollRequestTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableWebRAEnrollRequestTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableWebRAEnrollRequestTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebRAEnrollRequestTemplate(val *WebRAEnrollRequestTemplate) *NullableWebRAEnrollRequestTemplate {
	return &NullableWebRAEnrollRequestTemplate{value: val, isSet: true}
}

func (v NullableWebRAEnrollRequestTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebRAEnrollRequestTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
