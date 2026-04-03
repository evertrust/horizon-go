/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the CertificateTemplate type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CertificateTemplate{}

// CertificateTemplate struct for CertificateTemplate
type CertificateTemplate struct {
	ContactEmailPolicy   NullableContactEmailPolicy `json:"contactEmailPolicy,omitempty"`
	Extensions           []ExtensionElement         `json:"extensions,omitempty"`
	Labels               []LabelElement             `json:"labels,omitempty"`
	MetadataPolicies     []MetadataPolicy           `json:"metadataPolicies,omitempty"`
	OwnerPolicy          NullableOwnerPolicy        `json:"ownerPolicy,omitempty"`
	Sans                 []SANElement               `json:"sans,omitempty"`
	Subject              []DNElement                `json:"subject,omitempty"`
	TeamPolicy           NullableTeamPolicy         `json:"teamPolicy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateTemplate CertificateTemplate

// NewCertificateTemplate instantiates a new CertificateTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateTemplate() *CertificateTemplate {
	this := CertificateTemplate{}
	return &this
}

// NewCertificateTemplateWithDefaults instantiates a new CertificateTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateTemplateWithDefaults() *CertificateTemplate {
	this := CertificateTemplate{}
	return &this
}

// GetContactEmailPolicy returns the ContactEmailPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateTemplate) GetContactEmailPolicy() ContactEmailPolicy {
	if o == nil || utils.IsNil(o.ContactEmailPolicy.Get()) {
		var ret ContactEmailPolicy
		return ret
	}
	return *o.ContactEmailPolicy.Get()
}

// GetContactEmailPolicyOk returns a tuple with the ContactEmailPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateTemplate) GetContactEmailPolicyOk() (*ContactEmailPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmailPolicy.Get(), o.ContactEmailPolicy.IsSet()
}

// HasContactEmailPolicy returns a boolean if a field has been set.
func (o *CertificateTemplate) HasContactEmailPolicy() bool {
	if o != nil && o.ContactEmailPolicy.IsSet() {
		return true
	}

	return false
}

// SetContactEmailPolicy gets a reference to the given NullableContactEmailPolicy and assigns it to the ContactEmailPolicy field.
func (o *CertificateTemplate) SetContactEmailPolicy(v ContactEmailPolicy) {
	o.ContactEmailPolicy.Set(&v)
}

// SetContactEmailPolicyNil sets the value for ContactEmailPolicy to be an explicit nil
func (o *CertificateTemplate) SetContactEmailPolicyNil() {
	o.ContactEmailPolicy.Set(nil)
}

// UnsetContactEmailPolicy ensures that no value is present for ContactEmailPolicy, not even an explicit nil
func (o *CertificateTemplate) UnsetContactEmailPolicy() {
	o.ContactEmailPolicy.Unset()
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateTemplate) GetExtensions() []ExtensionElement {
	if o == nil {
		var ret []ExtensionElement
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateTemplate) GetExtensionsOk() ([]ExtensionElement, bool) {
	if o == nil || utils.IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *CertificateTemplate) HasExtensions() bool {
	if o != nil && !utils.IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []ExtensionElement and assigns it to the Extensions field.
func (o *CertificateTemplate) SetExtensions(v []ExtensionElement) {
	o.Extensions = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateTemplate) GetLabels() []LabelElement {
	if o == nil {
		var ret []LabelElement
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateTemplate) GetLabelsOk() ([]LabelElement, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CertificateTemplate) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []LabelElement and assigns it to the Labels field.
func (o *CertificateTemplate) SetLabels(v []LabelElement) {
	o.Labels = v
}

// GetMetadataPolicies returns the MetadataPolicies field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateTemplate) GetMetadataPolicies() []MetadataPolicy {
	if o == nil {
		var ret []MetadataPolicy
		return ret
	}
	return o.MetadataPolicies
}

// GetMetadataPoliciesOk returns a tuple with the MetadataPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateTemplate) GetMetadataPoliciesOk() ([]MetadataPolicy, bool) {
	if o == nil || utils.IsNil(o.MetadataPolicies) {
		return nil, false
	}
	return o.MetadataPolicies, true
}

// HasMetadataPolicies returns a boolean if a field has been set.
func (o *CertificateTemplate) HasMetadataPolicies() bool {
	if o != nil && !utils.IsNil(o.MetadataPolicies) {
		return true
	}

	return false
}

// SetMetadataPolicies gets a reference to the given []MetadataPolicy and assigns it to the MetadataPolicies field.
func (o *CertificateTemplate) SetMetadataPolicies(v []MetadataPolicy) {
	o.MetadataPolicies = v
}

// GetOwnerPolicy returns the OwnerPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateTemplate) GetOwnerPolicy() OwnerPolicy {
	if o == nil || utils.IsNil(o.OwnerPolicy.Get()) {
		var ret OwnerPolicy
		return ret
	}
	return *o.OwnerPolicy.Get()
}

// GetOwnerPolicyOk returns a tuple with the OwnerPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateTemplate) GetOwnerPolicyOk() (*OwnerPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerPolicy.Get(), o.OwnerPolicy.IsSet()
}

// HasOwnerPolicy returns a boolean if a field has been set.
func (o *CertificateTemplate) HasOwnerPolicy() bool {
	if o != nil && o.OwnerPolicy.IsSet() {
		return true
	}

	return false
}

// SetOwnerPolicy gets a reference to the given NullableOwnerPolicy and assigns it to the OwnerPolicy field.
func (o *CertificateTemplate) SetOwnerPolicy(v OwnerPolicy) {
	o.OwnerPolicy.Set(&v)
}

// SetOwnerPolicyNil sets the value for OwnerPolicy to be an explicit nil
func (o *CertificateTemplate) SetOwnerPolicyNil() {
	o.OwnerPolicy.Set(nil)
}

// UnsetOwnerPolicy ensures that no value is present for OwnerPolicy, not even an explicit nil
func (o *CertificateTemplate) UnsetOwnerPolicy() {
	o.OwnerPolicy.Unset()
}

// GetSans returns the Sans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateTemplate) GetSans() []SANElement {
	if o == nil {
		var ret []SANElement
		return ret
	}
	return o.Sans
}

// GetSansOk returns a tuple with the Sans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateTemplate) GetSansOk() ([]SANElement, bool) {
	if o == nil || utils.IsNil(o.Sans) {
		return nil, false
	}
	return o.Sans, true
}

// HasSans returns a boolean if a field has been set.
func (o *CertificateTemplate) HasSans() bool {
	if o != nil && !utils.IsNil(o.Sans) {
		return true
	}

	return false
}

// SetSans gets a reference to the given []SANElement and assigns it to the Sans field.
func (o *CertificateTemplate) SetSans(v []SANElement) {
	o.Sans = v
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateTemplate) GetSubject() []DNElement {
	if o == nil {
		var ret []DNElement
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateTemplate) GetSubjectOk() ([]DNElement, bool) {
	if o == nil || utils.IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CertificateTemplate) HasSubject() bool {
	if o != nil && !utils.IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given []DNElement and assigns it to the Subject field.
func (o *CertificateTemplate) SetSubject(v []DNElement) {
	o.Subject = v
}

// GetTeamPolicy returns the TeamPolicy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertificateTemplate) GetTeamPolicy() TeamPolicy {
	if o == nil || utils.IsNil(o.TeamPolicy.Get()) {
		var ret TeamPolicy
		return ret
	}
	return *o.TeamPolicy.Get()
}

// GetTeamPolicyOk returns a tuple with the TeamPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertificateTemplate) GetTeamPolicyOk() (*TeamPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamPolicy.Get(), o.TeamPolicy.IsSet()
}

// HasTeamPolicy returns a boolean if a field has been set.
func (o *CertificateTemplate) HasTeamPolicy() bool {
	if o != nil && o.TeamPolicy.IsSet() {
		return true
	}

	return false
}

// SetTeamPolicy gets a reference to the given NullableTeamPolicy and assigns it to the TeamPolicy field.
func (o *CertificateTemplate) SetTeamPolicy(v TeamPolicy) {
	o.TeamPolicy.Set(&v)
}

// SetTeamPolicyNil sets the value for TeamPolicy to be an explicit nil
func (o *CertificateTemplate) SetTeamPolicyNil() {
	o.TeamPolicy.Set(nil)
}

// UnsetTeamPolicy ensures that no value is present for TeamPolicy, not even an explicit nil
func (o *CertificateTemplate) UnsetTeamPolicy() {
	o.TeamPolicy.Unset()
}

func (o CertificateTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContactEmailPolicy.IsSet() {
		toSerialize["contactEmailPolicy"] = o.ContactEmailPolicy.Get()
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.MetadataPolicies != nil {
		toSerialize["metadataPolicies"] = o.MetadataPolicies
	}
	if o.OwnerPolicy.IsSet() {
		toSerialize["ownerPolicy"] = o.OwnerPolicy.Get()
	}
	if o.Sans != nil {
		toSerialize["sans"] = o.Sans
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.TeamPolicy.IsSet() {
		toSerialize["teamPolicy"] = o.TeamPolicy.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateTemplate) UnmarshalJSON(data []byte) (err error) {
	varCertificateTemplate := _CertificateTemplate{}

	err = json.Unmarshal(data, &varCertificateTemplate)

	if err != nil {
		return err
	}

	*o = CertificateTemplate(varCertificateTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contactEmailPolicy")
		delete(additionalProperties, "extensions")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "metadataPolicies")
		delete(additionalProperties, "ownerPolicy")
		delete(additionalProperties, "sans")
		delete(additionalProperties, "subject")
		delete(additionalProperties, "teamPolicy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateTemplate struct {
	value *CertificateTemplate
	isSet bool
}

func (v NullableCertificateTemplate) Get() *CertificateTemplate {
	return v.value
}

func (v *NullableCertificateTemplate) Set(val *CertificateTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateTemplate(val *CertificateTemplate) *NullableCertificateTemplate {
	return &NullableCertificateTemplate{value: val, isSet: true}
}

func (v NullableCertificateTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
