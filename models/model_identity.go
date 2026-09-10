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

// checks if the Identity type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Identity{}

// Identity The principal's identity
type Identity struct {
	// The principal's certificate (in case of `X509` identity provider)
	Certificate utils.NullableString `json:"certificate,omitempty"`
	// The principal's e-mail
	Email utils.NullableString `json:"email,omitempty"`
	// The principal's identifier
	Identifier string `json:"identifier"`
	// The identity provider's name this principal is registered on
	IdentityProviderName utils.NullableString `json:"identityProviderName,omitempty"`
	// The identity provider's type this principal is registered on
	IdentityProviderType utils.NullableString `json:"identityProviderType,omitempty"`
	// The principal's name
	Name                 utils.NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Identity Identity

// NewIdentity instantiates a new Identity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentity(identifier string) *Identity {
	this := Identity{}
	this.Identifier = identifier
	return &this
}

// NewIdentityWithDefaults instantiates a new Identity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithDefaults() *Identity {
	this := Identity{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetCertificate() string {
	if o == nil || utils.IsNil(o.Certificate.Get()) {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *Identity) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *Identity) SetCertificate(v string) {
	o.Certificate.Set(&v)
}

// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *Identity) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *Identity) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetEmail() string {
	if o == nil || utils.IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *Identity) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *Identity) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *Identity) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *Identity) UnsetEmail() {
	o.Email.Unset()
}

// GetIdentifier returns the Identifier field value
func (o *Identity) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *Identity) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *Identity) SetIdentifier(v string) {
	o.Identifier = v
}

// GetIdentityProviderName returns the IdentityProviderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetIdentityProviderName() string {
	if o == nil || utils.IsNil(o.IdentityProviderName.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityProviderName.Get()
}

// GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetIdentityProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityProviderName.Get(), o.IdentityProviderName.IsSet()
}

// HasIdentityProviderName returns a boolean if a field has been set.
func (o *Identity) HasIdentityProviderName() bool {
	if o != nil && o.IdentityProviderName.IsSet() {
		return true
	}

	return false
}

// SetIdentityProviderName gets a reference to the given NullableString and assigns it to the IdentityProviderName field.
func (o *Identity) SetIdentityProviderName(v string) {
	o.IdentityProviderName.Set(&v)
}

// SetIdentityProviderNameNil sets the value for IdentityProviderName to be an explicit nil
func (o *Identity) SetIdentityProviderNameNil() {
	o.IdentityProviderName.Set(nil)
}

// UnsetIdentityProviderName ensures that no value is present for IdentityProviderName, not even an explicit nil
func (o *Identity) UnsetIdentityProviderName() {
	o.IdentityProviderName.Unset()
}

// GetIdentityProviderType returns the IdentityProviderType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetIdentityProviderType() string {
	if o == nil || utils.IsNil(o.IdentityProviderType.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityProviderType.Get()
}

// GetIdentityProviderTypeOk returns a tuple with the IdentityProviderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetIdentityProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityProviderType.Get(), o.IdentityProviderType.IsSet()
}

// HasIdentityProviderType returns a boolean if a field has been set.
func (o *Identity) HasIdentityProviderType() bool {
	if o != nil && o.IdentityProviderType.IsSet() {
		return true
	}

	return false
}

// SetIdentityProviderType gets a reference to the given NullableString and assigns it to the IdentityProviderType field.
func (o *Identity) SetIdentityProviderType(v string) {
	o.IdentityProviderType.Set(&v)
}

// SetIdentityProviderTypeNil sets the value for IdentityProviderType to be an explicit nil
func (o *Identity) SetIdentityProviderTypeNil() {
	o.IdentityProviderType.Set(nil)
}

// UnsetIdentityProviderType ensures that no value is present for IdentityProviderType, not even an explicit nil
func (o *Identity) UnsetIdentityProviderType() {
	o.IdentityProviderType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetName() string {
	if o == nil || utils.IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Identity) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Identity) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *Identity) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Identity) UnsetName() {
	o.Name.Unset()
}

func (o Identity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Identity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	toSerialize["identifier"] = o.Identifier
	if o.IdentityProviderName.IsSet() {
		toSerialize["identityProviderName"] = o.IdentityProviderName.Get()
	}
	if o.IdentityProviderType.IsSet() {
		toSerialize["identityProviderType"] = o.IdentityProviderType.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Identity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifier",
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

	varIdentity := _Identity{}

	err = json.Unmarshal(data, &varIdentity)

	if err != nil {
		return err
	}

	*o = Identity(varIdentity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "email")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "identityProviderName")
		delete(additionalProperties, "identityProviderType")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentity struct {
	value *Identity
	isSet bool
}

func (v NullableIdentity) Get() *Identity {
	return v.value
}

func (v *NullableIdentity) Set(val *Identity) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentity(val *Identity) *NullableIdentity {
	return &NullableIdentity{value: val, isSet: true}
}

func (v NullableIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
