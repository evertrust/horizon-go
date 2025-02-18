/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"encoding/json"
)

// checks if the ScepEnrollRequestTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScepEnrollRequestTemplate{}

// ScepEnrollRequestTemplate struct for ScepEnrollRequestTemplate
type ScepEnrollRequestTemplate struct {
	// List of DN elements that will be used to build the certificate's Distinguished Name
	Subject []IndexedDNElement `json:"subject,omitempty"`
	// List of SAN elements that will be used to build the certificate's Subject Alternative Name
	Sans []ListSANElement `json:"sans,omitempty"`
	// Information about the certificate's extensions and how to edit them
	Extensions []CertificateExtensionElement `json:"extensions,omitempty"`
	// List of labels used internally to tag and group certificates
	Labels []RequestLabelElement `json:"labels,omitempty"`
	// Information about the certificate's contact email and how to edit it
	ContactEmail NullableCertificateContactEmailElement `json:"contactEmail,omitempty"`
	// Information about the certificate's owner and how to edit it
	Owner NullableCertificateOwnerElement `json:"owner,omitempty"`
	// Information about the certificate's team and how to edit it
	Team NullableCertificateTeamElement `json:"team,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScepEnrollRequestTemplate ScepEnrollRequestTemplate

// NewScepEnrollRequestTemplate instantiates a new ScepEnrollRequestTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScepEnrollRequestTemplate() *ScepEnrollRequestTemplate {
	this := ScepEnrollRequestTemplate{}
	return &this
}

// NewScepEnrollRequestTemplateWithDefaults instantiates a new ScepEnrollRequestTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScepEnrollRequestTemplateWithDefaults() *ScepEnrollRequestTemplate {
	this := ScepEnrollRequestTemplate{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScepEnrollRequestTemplate) GetSubject() []IndexedDNElement {
	if o == nil {
		var ret []IndexedDNElement
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScepEnrollRequestTemplate) GetSubjectOk() ([]IndexedDNElement, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *ScepEnrollRequestTemplate) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given []IndexedDNElement and assigns it to the Subject field.
func (o *ScepEnrollRequestTemplate) SetSubject(v []IndexedDNElement) {
	o.Subject = v
}

// GetSans returns the Sans field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScepEnrollRequestTemplate) GetSans() []ListSANElement {
	if o == nil {
		var ret []ListSANElement
		return ret
	}
	return o.Sans
}

// GetSansOk returns a tuple with the Sans field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScepEnrollRequestTemplate) GetSansOk() ([]ListSANElement, bool) {
	if o == nil || IsNil(o.Sans) {
		return nil, false
	}
	return o.Sans, true
}

// HasSans returns a boolean if a field has been set.
func (o *ScepEnrollRequestTemplate) HasSans() bool {
	if o != nil && !IsNil(o.Sans) {
		return true
	}

	return false
}

// SetSans gets a reference to the given []ListSANElement and assigns it to the Sans field.
func (o *ScepEnrollRequestTemplate) SetSans(v []ListSANElement) {
	o.Sans = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScepEnrollRequestTemplate) GetExtensions() []CertificateExtensionElement {
	if o == nil {
		var ret []CertificateExtensionElement
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScepEnrollRequestTemplate) GetExtensionsOk() ([]CertificateExtensionElement, bool) {
	if o == nil || IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *ScepEnrollRequestTemplate) HasExtensions() bool {
	if o != nil && !IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []CertificateExtensionElement and assigns it to the Extensions field.
func (o *ScepEnrollRequestTemplate) SetExtensions(v []CertificateExtensionElement) {
	o.Extensions = v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScepEnrollRequestTemplate) GetLabels() []RequestLabelElement {
	if o == nil {
		var ret []RequestLabelElement
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScepEnrollRequestTemplate) GetLabelsOk() ([]RequestLabelElement, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ScepEnrollRequestTemplate) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []RequestLabelElement and assigns it to the Labels field.
func (o *ScepEnrollRequestTemplate) SetLabels(v []RequestLabelElement) {
	o.Labels = v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScepEnrollRequestTemplate) GetContactEmail() CertificateContactEmailElement {
	if o == nil || IsNil(o.ContactEmail.Get()) {
		var ret CertificateContactEmailElement
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScepEnrollRequestTemplate) GetContactEmailOk() (*CertificateContactEmailElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *ScepEnrollRequestTemplate) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableCertificateContactEmailElement and assigns it to the ContactEmail field.
func (o *ScepEnrollRequestTemplate) SetContactEmail(v CertificateContactEmailElement) {
	o.ContactEmail.Set(&v)
}
// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *ScepEnrollRequestTemplate) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *ScepEnrollRequestTemplate) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScepEnrollRequestTemplate) GetOwner() CertificateOwnerElement {
	if o == nil || IsNil(o.Owner.Get()) {
		var ret CertificateOwnerElement
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScepEnrollRequestTemplate) GetOwnerOk() (*CertificateOwnerElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *ScepEnrollRequestTemplate) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableCertificateOwnerElement and assigns it to the Owner field.
func (o *ScepEnrollRequestTemplate) SetOwner(v CertificateOwnerElement) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *ScepEnrollRequestTemplate) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *ScepEnrollRequestTemplate) UnsetOwner() {
	o.Owner.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScepEnrollRequestTemplate) GetTeam() CertificateTeamElement {
	if o == nil || IsNil(o.Team.Get()) {
		var ret CertificateTeamElement
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScepEnrollRequestTemplate) GetTeamOk() (*CertificateTeamElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *ScepEnrollRequestTemplate) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableCertificateTeamElement and assigns it to the Team field.
func (o *ScepEnrollRequestTemplate) SetTeam(v CertificateTeamElement) {
	o.Team.Set(&v)
}
// SetTeamNil sets the value for Team to be an explicit nil
func (o *ScepEnrollRequestTemplate) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *ScepEnrollRequestTemplate) UnsetTeam() {
	o.Team.Unset()
}

func (o ScepEnrollRequestTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScepEnrollRequestTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Sans != nil {
		toSerialize["sans"] = o.Sans
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.ContactEmail.IsSet() {
		toSerialize["contactEmail"] = o.ContactEmail.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScepEnrollRequestTemplate) UnmarshalJSON(data []byte) (err error) {
	varScepEnrollRequestTemplate := _ScepEnrollRequestTemplate{}

	err = json.Unmarshal(data, &varScepEnrollRequestTemplate)

	if err != nil {
		return err
	}

	*o = ScepEnrollRequestTemplate(varScepEnrollRequestTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subject")
		delete(additionalProperties, "sans")
		delete(additionalProperties, "extensions")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "team")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScepEnrollRequestTemplate struct {
	value *ScepEnrollRequestTemplate
	isSet bool
}

func (v NullableScepEnrollRequestTemplate) Get() *ScepEnrollRequestTemplate {
	return v.value
}

func (v *NullableScepEnrollRequestTemplate) Set(val *ScepEnrollRequestTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableScepEnrollRequestTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableScepEnrollRequestTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScepEnrollRequestTemplate(val *ScepEnrollRequestTemplate) *NullableScepEnrollRequestTemplate {
	return &NullableScepEnrollRequestTemplate{value: val, isSet: true}
}

func (v NullableScepEnrollRequestTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScepEnrollRequestTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


