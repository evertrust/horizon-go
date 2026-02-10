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

// checks if the WebRAImportRequestTemplateResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebRAImportRequestTemplateResponse{}

// WebRAImportRequestTemplateResponse struct for WebRAImportRequestTemplateResponse
type WebRAImportRequestTemplateResponse struct {
	// The contact email for this certificate
	ContactEmail NullableCertificateContactEmailElement `json:"contactEmail,omitempty"`
	// The host discovery data associated with the certificate (discovery metadata)
	DiscoveryData *HostDiscoveryData `json:"discoveryData,omitempty"`
	// Information about the discovery of this certificate
	DiscoveryInfo NullableDiscoveryInfo `json:"discoveryInfo,omitempty"`
	// The labels for this certificate
	Labels []RequestLabelElement `json:"labels,omitempty"`
	// The technical metadata for this certificate
	Metadata []CertificateMetadataElement `json:"metadata,omitempty"`
	// The owner for this certificate
	Owner NullableCertificateOwnerElement `json:"owner,omitempty"`
	// The team for this certificate
	Team NullableCertificateTeamElement `json:"team,omitempty"`
	// The third party data associated with the certificate
	ThirdPartyData       []ThirdPartyItem `json:"thirdPartyData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebRAImportRequestTemplateResponse WebRAImportRequestTemplateResponse

// NewWebRAImportRequestTemplateResponse instantiates a new WebRAImportRequestTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebRAImportRequestTemplateResponse() *WebRAImportRequestTemplateResponse {
	this := WebRAImportRequestTemplateResponse{}
	return &this
}

// NewWebRAImportRequestTemplateResponseWithDefaults instantiates a new WebRAImportRequestTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebRAImportRequestTemplateResponseWithDefaults() *WebRAImportRequestTemplateResponse {
	this := WebRAImportRequestTemplateResponse{}
	return &this
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestTemplateResponse) GetContactEmail() CertificateContactEmailElement {
	if o == nil || utils.IsNil(o.ContactEmail.Get()) {
		var ret CertificateContactEmailElement
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestTemplateResponse) GetContactEmailOk() (*CertificateContactEmailElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *WebRAImportRequestTemplateResponse) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableCertificateContactEmailElement and assigns it to the ContactEmail field.
func (o *WebRAImportRequestTemplateResponse) SetContactEmail(v CertificateContactEmailElement) {
	o.ContactEmail.Set(&v)
}

// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *WebRAImportRequestTemplateResponse) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *WebRAImportRequestTemplateResponse) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetDiscoveryData returns the DiscoveryData field value if set, zero value otherwise.
func (o *WebRAImportRequestTemplateResponse) GetDiscoveryData() HostDiscoveryData {
	if o == nil || utils.IsNil(o.DiscoveryData) {
		var ret HostDiscoveryData
		return ret
	}
	return *o.DiscoveryData
}

// GetDiscoveryDataOk returns a tuple with the DiscoveryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRAImportRequestTemplateResponse) GetDiscoveryDataOk() (*HostDiscoveryData, bool) {
	if o == nil || utils.IsNil(o.DiscoveryData) {
		return nil, false
	}
	return o.DiscoveryData, true
}

// HasDiscoveryData returns a boolean if a field has been set.
func (o *WebRAImportRequestTemplateResponse) HasDiscoveryData() bool {
	if o != nil && !utils.IsNil(o.DiscoveryData) {
		return true
	}

	return false
}

// SetDiscoveryData gets a reference to the given HostDiscoveryData and assigns it to the DiscoveryData field.
func (o *WebRAImportRequestTemplateResponse) SetDiscoveryData(v HostDiscoveryData) {
	o.DiscoveryData = &v
}

// GetDiscoveryInfo returns the DiscoveryInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestTemplateResponse) GetDiscoveryInfo() DiscoveryInfo {
	if o == nil || utils.IsNil(o.DiscoveryInfo.Get()) {
		var ret DiscoveryInfo
		return ret
	}
	return *o.DiscoveryInfo.Get()
}

// GetDiscoveryInfoOk returns a tuple with the DiscoveryInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestTemplateResponse) GetDiscoveryInfoOk() (*DiscoveryInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscoveryInfo.Get(), o.DiscoveryInfo.IsSet()
}

// HasDiscoveryInfo returns a boolean if a field has been set.
func (o *WebRAImportRequestTemplateResponse) HasDiscoveryInfo() bool {
	if o != nil && o.DiscoveryInfo.IsSet() {
		return true
	}

	return false
}

// SetDiscoveryInfo gets a reference to the given NullableDiscoveryInfo and assigns it to the DiscoveryInfo field.
func (o *WebRAImportRequestTemplateResponse) SetDiscoveryInfo(v DiscoveryInfo) {
	o.DiscoveryInfo.Set(&v)
}

// SetDiscoveryInfoNil sets the value for DiscoveryInfo to be an explicit nil
func (o *WebRAImportRequestTemplateResponse) SetDiscoveryInfoNil() {
	o.DiscoveryInfo.Set(nil)
}

// UnsetDiscoveryInfo ensures that no value is present for DiscoveryInfo, not even an explicit nil
func (o *WebRAImportRequestTemplateResponse) UnsetDiscoveryInfo() {
	o.DiscoveryInfo.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestTemplateResponse) GetLabels() []RequestLabelElement {
	if o == nil {
		var ret []RequestLabelElement
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestTemplateResponse) GetLabelsOk() ([]RequestLabelElement, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *WebRAImportRequestTemplateResponse) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []RequestLabelElement and assigns it to the Labels field.
func (o *WebRAImportRequestTemplateResponse) SetLabels(v []RequestLabelElement) {
	o.Labels = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestTemplateResponse) GetMetadata() []CertificateMetadataElement {
	if o == nil {
		var ret []CertificateMetadataElement
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestTemplateResponse) GetMetadataOk() ([]CertificateMetadataElement, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WebRAImportRequestTemplateResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []CertificateMetadataElement and assigns it to the Metadata field.
func (o *WebRAImportRequestTemplateResponse) SetMetadata(v []CertificateMetadataElement) {
	o.Metadata = v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestTemplateResponse) GetOwner() CertificateOwnerElement {
	if o == nil || utils.IsNil(o.Owner.Get()) {
		var ret CertificateOwnerElement
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestTemplateResponse) GetOwnerOk() (*CertificateOwnerElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *WebRAImportRequestTemplateResponse) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableCertificateOwnerElement and assigns it to the Owner field.
func (o *WebRAImportRequestTemplateResponse) SetOwner(v CertificateOwnerElement) {
	o.Owner.Set(&v)
}

// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *WebRAImportRequestTemplateResponse) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *WebRAImportRequestTemplateResponse) UnsetOwner() {
	o.Owner.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestTemplateResponse) GetTeam() CertificateTeamElement {
	if o == nil || utils.IsNil(o.Team.Get()) {
		var ret CertificateTeamElement
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestTemplateResponse) GetTeamOk() (*CertificateTeamElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *WebRAImportRequestTemplateResponse) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableCertificateTeamElement and assigns it to the Team field.
func (o *WebRAImportRequestTemplateResponse) SetTeam(v CertificateTeamElement) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *WebRAImportRequestTemplateResponse) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *WebRAImportRequestTemplateResponse) UnsetTeam() {
	o.Team.Unset()
}

// GetThirdPartyData returns the ThirdPartyData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAImportRequestTemplateResponse) GetThirdPartyData() []ThirdPartyItem {
	if o == nil {
		var ret []ThirdPartyItem
		return ret
	}
	return o.ThirdPartyData
}

// GetThirdPartyDataOk returns a tuple with the ThirdPartyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAImportRequestTemplateResponse) GetThirdPartyDataOk() ([]ThirdPartyItem, bool) {
	if o == nil || utils.IsNil(o.ThirdPartyData) {
		return nil, false
	}
	return o.ThirdPartyData, true
}

// HasThirdPartyData returns a boolean if a field has been set.
func (o *WebRAImportRequestTemplateResponse) HasThirdPartyData() bool {
	if o != nil && !utils.IsNil(o.ThirdPartyData) {
		return true
	}

	return false
}

// SetThirdPartyData gets a reference to the given []ThirdPartyItem and assigns it to the ThirdPartyData field.
func (o *WebRAImportRequestTemplateResponse) SetThirdPartyData(v []ThirdPartyItem) {
	o.ThirdPartyData = v
}

func (o WebRAImportRequestTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebRAImportRequestTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContactEmail.IsSet() {
		toSerialize["contactEmail"] = o.ContactEmail.Get()
	}
	if !utils.IsNil(o.DiscoveryData) {
		toSerialize["discoveryData"] = o.DiscoveryData
	}
	if o.DiscoveryInfo.IsSet() {
		toSerialize["discoveryInfo"] = o.DiscoveryInfo.Get()
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
	if o.ThirdPartyData != nil {
		toSerialize["thirdPartyData"] = o.ThirdPartyData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebRAImportRequestTemplateResponse) UnmarshalJSON(data []byte) (err error) {
	varWebRAImportRequestTemplateResponse := _WebRAImportRequestTemplateResponse{}

	err = json.Unmarshal(data, &varWebRAImportRequestTemplateResponse)

	if err != nil {
		return err
	}

	*o = WebRAImportRequestTemplateResponse(varWebRAImportRequestTemplateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "discoveryData")
		delete(additionalProperties, "discoveryInfo")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "team")
		delete(additionalProperties, "thirdPartyData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebRAImportRequestTemplateResponse struct {
	value *WebRAImportRequestTemplateResponse
	isSet bool
}

func (v NullableWebRAImportRequestTemplateResponse) Get() *WebRAImportRequestTemplateResponse {
	return v.value
}

func (v *NullableWebRAImportRequestTemplateResponse) Set(val *WebRAImportRequestTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebRAImportRequestTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebRAImportRequestTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebRAImportRequestTemplateResponse(val *WebRAImportRequestTemplateResponse) *NullableWebRAImportRequestTemplateResponse {
	return &NullableWebRAImportRequestTemplateResponse{value: val, isSet: true}
}

func (v NullableWebRAImportRequestTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebRAImportRequestTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
