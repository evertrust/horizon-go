/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the WebRAMigrateRequestTemplateResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebRAMigrateRequestTemplateResponse{}

// WebRAMigrateRequestTemplateResponse struct for WebRAMigrateRequestTemplateResponse
type WebRAMigrateRequestTemplateResponse struct {
	AutoRenew            *CertificateAutoRenewElementResponse           `json:"autoRenew,omitempty"`
	ContactEmail         NullableCertificateContactEmailElementResponse `json:"contactEmail,omitempty"`
	Labels               []RequestLabelElementResponse                  `json:"labels,omitempty"`
	Metadata             []CertificateMetadataElementResponse           `json:"metadata,omitempty"`
	Owner                NullableCertificateOwnerElementResponse        `json:"owner,omitempty"`
	Team                 NullableCertificateTeamElementResponse         `json:"team,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebRAMigrateRequestTemplateResponse WebRAMigrateRequestTemplateResponse

// NewWebRAMigrateRequestTemplateResponse instantiates a new WebRAMigrateRequestTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebRAMigrateRequestTemplateResponse() *WebRAMigrateRequestTemplateResponse {
	this := WebRAMigrateRequestTemplateResponse{}
	return &this
}

// NewWebRAMigrateRequestTemplateResponseWithDefaults instantiates a new WebRAMigrateRequestTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebRAMigrateRequestTemplateResponseWithDefaults() *WebRAMigrateRequestTemplateResponse {
	this := WebRAMigrateRequestTemplateResponse{}
	return &this
}

// GetAutoRenew returns the AutoRenew field value if set, zero value otherwise.
func (o *WebRAMigrateRequestTemplateResponse) GetAutoRenew() CertificateAutoRenewElementResponse {
	if o == nil || utils.IsNil(o.AutoRenew) {
		var ret CertificateAutoRenewElementResponse
		return ret
	}
	return *o.AutoRenew
}

// GetAutoRenewOk returns a tuple with the AutoRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRAMigrateRequestTemplateResponse) GetAutoRenewOk() (*CertificateAutoRenewElementResponse, bool) {
	if o == nil || utils.IsNil(o.AutoRenew) {
		return nil, false
	}
	return o.AutoRenew, true
}

// HasAutoRenew returns a boolean if a field has been set.
func (o *WebRAMigrateRequestTemplateResponse) HasAutoRenew() bool {
	if o != nil && !utils.IsNil(o.AutoRenew) {
		return true
	}

	return false
}

// SetAutoRenew gets a reference to the given CertificateAutoRenewElementResponse and assigns it to the AutoRenew field.
func (o *WebRAMigrateRequestTemplateResponse) SetAutoRenew(v CertificateAutoRenewElementResponse) {
	o.AutoRenew = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAMigrateRequestTemplateResponse) GetContactEmail() CertificateContactEmailElementResponse {
	if o == nil || utils.IsNil(o.ContactEmail.Get()) {
		var ret CertificateContactEmailElementResponse
		return ret
	}
	return *o.ContactEmail.Get()
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAMigrateRequestTemplateResponse) GetContactEmailOk() (*CertificateContactEmailElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactEmail.Get(), o.ContactEmail.IsSet()
}

// HasContactEmail returns a boolean if a field has been set.
func (o *WebRAMigrateRequestTemplateResponse) HasContactEmail() bool {
	if o != nil && o.ContactEmail.IsSet() {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given NullableCertificateContactEmailElementResponse and assigns it to the ContactEmail field.
func (o *WebRAMigrateRequestTemplateResponse) SetContactEmail(v CertificateContactEmailElementResponse) {
	o.ContactEmail.Set(&v)
}

// SetContactEmailNil sets the value for ContactEmail to be an explicit nil
func (o *WebRAMigrateRequestTemplateResponse) SetContactEmailNil() {
	o.ContactEmail.Set(nil)
}

// UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
func (o *WebRAMigrateRequestTemplateResponse) UnsetContactEmail() {
	o.ContactEmail.Unset()
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAMigrateRequestTemplateResponse) GetLabels() []RequestLabelElementResponse {
	if o == nil {
		var ret []RequestLabelElementResponse
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAMigrateRequestTemplateResponse) GetLabelsOk() ([]RequestLabelElementResponse, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *WebRAMigrateRequestTemplateResponse) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []RequestLabelElementResponse and assigns it to the Labels field.
func (o *WebRAMigrateRequestTemplateResponse) SetLabels(v []RequestLabelElementResponse) {
	o.Labels = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAMigrateRequestTemplateResponse) GetMetadata() []CertificateMetadataElementResponse {
	if o == nil {
		var ret []CertificateMetadataElementResponse
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAMigrateRequestTemplateResponse) GetMetadataOk() ([]CertificateMetadataElementResponse, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WebRAMigrateRequestTemplateResponse) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []CertificateMetadataElementResponse and assigns it to the Metadata field.
func (o *WebRAMigrateRequestTemplateResponse) SetMetadata(v []CertificateMetadataElementResponse) {
	o.Metadata = v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAMigrateRequestTemplateResponse) GetOwner() CertificateOwnerElementResponse {
	if o == nil || utils.IsNil(o.Owner.Get()) {
		var ret CertificateOwnerElementResponse
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAMigrateRequestTemplateResponse) GetOwnerOk() (*CertificateOwnerElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *WebRAMigrateRequestTemplateResponse) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableCertificateOwnerElementResponse and assigns it to the Owner field.
func (o *WebRAMigrateRequestTemplateResponse) SetOwner(v CertificateOwnerElementResponse) {
	o.Owner.Set(&v)
}

// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *WebRAMigrateRequestTemplateResponse) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *WebRAMigrateRequestTemplateResponse) UnsetOwner() {
	o.Owner.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebRAMigrateRequestTemplateResponse) GetTeam() CertificateTeamElementResponse {
	if o == nil || utils.IsNil(o.Team.Get()) {
		var ret CertificateTeamElementResponse
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebRAMigrateRequestTemplateResponse) GetTeamOk() (*CertificateTeamElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *WebRAMigrateRequestTemplateResponse) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableCertificateTeamElementResponse and assigns it to the Team field.
func (o *WebRAMigrateRequestTemplateResponse) SetTeam(v CertificateTeamElementResponse) {
	o.Team.Set(&v)
}

// SetTeamNil sets the value for Team to be an explicit nil
func (o *WebRAMigrateRequestTemplateResponse) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *WebRAMigrateRequestTemplateResponse) UnsetTeam() {
	o.Team.Unset()
}

func (o WebRAMigrateRequestTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebRAMigrateRequestTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AutoRenew) {
		toSerialize["autoRenew"] = o.AutoRenew
	}
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

func (o *WebRAMigrateRequestTemplateResponse) UnmarshalJSON(data []byte) (err error) {
	varWebRAMigrateRequestTemplateResponse := _WebRAMigrateRequestTemplateResponse{}

	err = json.Unmarshal(data, &varWebRAMigrateRequestTemplateResponse)

	if err != nil {
		return err
	}

	*o = WebRAMigrateRequestTemplateResponse(varWebRAMigrateRequestTemplateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "autoRenew")
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "team")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebRAMigrateRequestTemplateResponse struct {
	value *WebRAMigrateRequestTemplateResponse
	isSet bool
}

func (v NullableWebRAMigrateRequestTemplateResponse) Get() *WebRAMigrateRequestTemplateResponse {
	return v.value
}

func (v *NullableWebRAMigrateRequestTemplateResponse) Set(val *WebRAMigrateRequestTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebRAMigrateRequestTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebRAMigrateRequestTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebRAMigrateRequestTemplateResponse(val *WebRAMigrateRequestTemplateResponse) *NullableWebRAMigrateRequestTemplateResponse {
	return &NullableWebRAMigrateRequestTemplateResponse{value: val, isSet: true}
}

func (v NullableWebRAMigrateRequestTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebRAMigrateRequestTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
