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

// checks if the WebRAUpdateRequestTemplateResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebRAUpdateRequestTemplateResponse{}

// WebRAUpdateRequestTemplateResponse struct for WebRAUpdateRequestTemplateResponse
type WebRAUpdateRequestTemplateResponse struct {
	// Information about the certificate's contact email and how to edit it
	ContactEmail NullableCertificateContactEmailElementResponse `json:"contactEmail,omitempty"`
	// Information about the certificate's labels and how to edit them
	Labels []RequestLabelElementResponse `json:"labels,omitempty"`
	// Information about the certificate's metadata and how to edit them
	Metadata []CertificateMetadataElementResponse `json:"metadata,omitempty"`
	// Information about the certificate's owner and how to edit it
	Owner NullableCertificateOwnerElementResponse `json:"owner,omitempty"`
	// Information about the certificate's team and how to edit it
	Team                 NullableCertificateTeamElementResponse `json:"team,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebRAUpdateRequestTemplateResponse WebRAUpdateRequestTemplateResponse

// NewWebRAUpdateRequestTemplateResponse instantiates a new WebRAUpdateRequestTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebRAUpdateRequestTemplateResponse() *WebRAUpdateRequestTemplateResponse {
	this := WebRAUpdateRequestTemplateResponse{}
	return &this
}

// NewWebRAUpdateRequestTemplateResponseWithDefaults instantiates a new WebRAUpdateRequestTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebRAUpdateRequestTemplateResponseWithDefaults() *WebRAUpdateRequestTemplateResponse {
	this := WebRAUpdateRequestTemplateResponse{}
	return &this
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestTemplateResponse) GetContactEmail() CertificateContactEmailElementResponse {
	if o == nil || utils.IsNil(o.ContactEmail.Get()) {
		var ret CertificateContactEmailElementResponse
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestTemplateResponse) GetContactEmailOk() (*CertificateContactEmailElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *WebRAUpdateRequestTemplateResponse) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableCertificateContactEmailElementResponse and assigns it to the ContactEmail field.
func (o *WebRAUpdateRequestTemplateResponse) SetContactEmail(v CertificateContactEmailElementResponse) {
	o.ContactEmail.Set(&v)
}

// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *WebRAUpdateRequestTemplateResponse) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *WebRAUpdateRequestTemplateResponse) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestTemplateResponse) GetLabels() []RequestLabelElementResponse {
	if o == nil {
		var ret []RequestLabelElementResponse
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestTemplateResponse) GetLabelsOk() ([]RequestLabelElementResponse, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *WebRAUpdateRequestTemplateResponse) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []RequestLabelElementResponse and assigns it to the Labels field.
func (o *WebRAUpdateRequestTemplateResponse) SetLabels(v []RequestLabelElementResponse) {
	o.Labels = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestTemplateResponse) GetMetadata() []CertificateMetadataElementResponse {
	if o == nil {
		var ret []CertificateMetadataElementResponse
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestTemplateResponse) GetMetadataOk() ([]CertificateMetadataElementResponse, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WebRAUpdateRequestTemplateResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []CertificateMetadataElementResponse and assigns it to the Metadata field.
func (o *WebRAUpdateRequestTemplateResponse) SetMetadata(v []CertificateMetadataElementResponse) {
	o.Metadata = v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestTemplateResponse) GetOwner() CertificateOwnerElementResponse {
	if o == nil || utils.IsNil(o.Owner.Get()) {
		var ret CertificateOwnerElementResponse
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestTemplateResponse) GetOwnerOk() (*CertificateOwnerElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *WebRAUpdateRequestTemplateResponse) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableCertificateOwnerElementResponse and assigns it to the Owner field.
func (o *WebRAUpdateRequestTemplateResponse) SetOwner(v CertificateOwnerElementResponse) {
	o.Owner.Set(&v)
}

// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *WebRAUpdateRequestTemplateResponse) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *WebRAUpdateRequestTemplateResponse) UnsetOwner() {
	o.Owner.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAUpdateRequestTemplateResponse) GetTeam() CertificateTeamElementResponse {
	if o == nil || utils.IsNil(o.Team.Get()) {
		var ret CertificateTeamElementResponse
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAUpdateRequestTemplateResponse) GetTeamOk() (*CertificateTeamElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *WebRAUpdateRequestTemplateResponse) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableCertificateTeamElementResponse and assigns it to the Team field.
func (o *WebRAUpdateRequestTemplateResponse) SetTeam(v CertificateTeamElementResponse) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *WebRAUpdateRequestTemplateResponse) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *WebRAUpdateRequestTemplateResponse) UnsetTeam() {
	o.Team.Unset()
}

func (o WebRAUpdateRequestTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebRAUpdateRequestTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContactEmail.IsSet() {
		toSerialize["contactEmail"] = o.ContactEmail.Get()
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
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebRAUpdateRequestTemplateResponse) UnmarshalJSON(data []byte) (err error) {
	varWebRAUpdateRequestTemplateResponse := _WebRAUpdateRequestTemplateResponse{}

	err = json.Unmarshal(data, &varWebRAUpdateRequestTemplateResponse)

	if err != nil {
		return err
	}

	*o = WebRAUpdateRequestTemplateResponse(varWebRAUpdateRequestTemplateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "team")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebRAUpdateRequestTemplateResponse struct {
	value *WebRAUpdateRequestTemplateResponse
	isSet bool
}

func (v NullableWebRAUpdateRequestTemplateResponse) Get() *WebRAUpdateRequestTemplateResponse {
	return v.value
}

func (v *NullableWebRAUpdateRequestTemplateResponse) Set(val *WebRAUpdateRequestTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebRAUpdateRequestTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebRAUpdateRequestTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebRAUpdateRequestTemplateResponse(val *WebRAUpdateRequestTemplateResponse) *NullableWebRAUpdateRequestTemplateResponse {
	return &NullableWebRAUpdateRequestTemplateResponse{value: val, isSet: true}
}

func (v NullableWebRAUpdateRequestTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebRAUpdateRequestTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
