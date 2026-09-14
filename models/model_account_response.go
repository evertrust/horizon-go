/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the AccountResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AccountResponse{}

// AccountResponse struct for AccountResponse
type AccountResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// The date when the account was compromised
	CompromisedAt *int64 `json:"compromisedAt,omitempty"`
	// One of: `unspecified`, `keycompromise`, `cacompromise`, `affiliationchange`, `superseded`, `cessationofoperation`
	CompromissionReason  utils.NullableString `json:"compromissionReason,omitempty"`
	Contact              []string             `json:"contact,omitempty"`
	CreatedAt            int64                `json:"createdAt"`
	EabName              *string              `json:"eabName,omitempty"`
	ExpirationDate       *int64               `json:"expirationDate,omitempty"`
	InitialIp            *string              `json:"initialIp,omitempty"`
	Jwk                  JsonWebKey           `json:"jwk"`
	KeyThumbprint        string               `json:"keyThumbprint"`
	Status               AccountStatus        `json:"status"`
	TermsOfServiceAgreed *bool                `json:"termsOfServiceAgreed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountResponse AccountResponse

// NewAccountResponse instantiates a new AccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountResponse(id string, createdAt int64, jwk JsonWebKey, keyThumbprint string, status AccountStatus) *AccountResponse {
	this := AccountResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Jwk = jwk
	this.KeyThumbprint = keyThumbprint
	this.Status = status
	return &this
}

// NewAccountResponseWithDefaults instantiates a new AccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountResponseWithDefaults() *AccountResponse {
	this := AccountResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AccountResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountResponse) SetId(v string) {
	o.Id = v
}

// GetCompromisedAt returns the CompromisedAt field value if set, zero value otherwise.
func (o *AccountResponse) GetCompromisedAt() int64 {
	if o == nil || utils.IsNil(o.CompromisedAt) {
		var ret int64
		return ret
	}
	return *o.CompromisedAt
}

// GetCompromisedAtOk returns a tuple with the CompromisedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetCompromisedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CompromisedAt) {
		return nil, false
	}
	return o.CompromisedAt, true
}

// HasCompromisedAt returns a boolean if a field has been set.
func (o *AccountResponse) HasCompromisedAt() bool {
	if o != nil && !utils.IsNil(o.CompromisedAt) {
		return true
	}

	return false
}

// SetCompromisedAt gets a reference to the given int64 and assigns it to the CompromisedAt field.
func (o *AccountResponse) SetCompromisedAt(v int64) {
	o.CompromisedAt = &v
}

// GetCompromissionReason returns the CompromissionReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountResponse) GetCompromissionReason() string {
	if o == nil || utils.IsNil(o.CompromissionReason.Get()) {
		var ret string
		return ret
	}
	return *o.CompromissionReason.Get()
}

// GetCompromissionReasonOk returns a tuple with the CompromissionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountResponse) GetCompromissionReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompromissionReason.Get(), o.CompromissionReason.IsSet()
}

// HasCompromissionReason returns a boolean if a field has been set.
func (o *AccountResponse) HasCompromissionReason() bool {
	if o != nil && o.CompromissionReason.IsSet() {
		return true
	}

	return false
}

// SetCompromissionReason gets a reference to the given NullableString and assigns it to the CompromissionReason field.
func (o *AccountResponse) SetCompromissionReason(v string) {
	o.CompromissionReason.Set(&v)
}

// SetCompromissionReasonNil sets the value for CompromissionReason to be an explicit nil
func (o *AccountResponse) SetCompromissionReasonNil() {
	o.CompromissionReason.Set(nil)
}

// UnsetCompromissionReason ensures that no value is present for CompromissionReason, not even an explicit nil
func (o *AccountResponse) UnsetCompromissionReason() {
	o.CompromissionReason.Unset()
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *AccountResponse) GetContact() []string {
	if o == nil || utils.IsNil(o.Contact) {
		var ret []string
		return ret
	}
	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetContactOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *AccountResponse) HasContact() bool {
	if o != nil && !utils.IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given []string and assigns it to the Contact field.
func (o *AccountResponse) SetContact(v []string) {
	o.Contact = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AccountResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AccountResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetEabName returns the EabName field value if set, zero value otherwise.
func (o *AccountResponse) GetEabName() string {
	if o == nil || utils.IsNil(o.EabName) {
		var ret string
		return ret
	}
	return *o.EabName
}

// GetEabNameOk returns a tuple with the EabName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetEabNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EabName) {
		return nil, false
	}
	return o.EabName, true
}

// HasEabName returns a boolean if a field has been set.
func (o *AccountResponse) HasEabName() bool {
	if o != nil && !utils.IsNil(o.EabName) {
		return true
	}

	return false
}

// SetEabName gets a reference to the given string and assigns it to the EabName field.
func (o *AccountResponse) SetEabName(v string) {
	o.EabName = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *AccountResponse) GetExpirationDate() int64 {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		var ret int64
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetExpirationDateOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *AccountResponse) HasExpirationDate() bool {
	if o != nil && !utils.IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *AccountResponse) SetExpirationDate(v int64) {
	o.ExpirationDate = &v
}

// GetInitialIp returns the InitialIp field value if set, zero value otherwise.
func (o *AccountResponse) GetInitialIp() string {
	if o == nil || utils.IsNil(o.InitialIp) {
		var ret string
		return ret
	}
	return *o.InitialIp
}

// GetInitialIpOk returns a tuple with the InitialIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetInitialIpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InitialIp) {
		return nil, false
	}
	return o.InitialIp, true
}

// HasInitialIp returns a boolean if a field has been set.
func (o *AccountResponse) HasInitialIp() bool {
	if o != nil && !utils.IsNil(o.InitialIp) {
		return true
	}

	return false
}

// SetInitialIp gets a reference to the given string and assigns it to the InitialIp field.
func (o *AccountResponse) SetInitialIp(v string) {
	o.InitialIp = &v
}

// GetJwk returns the Jwk field value
func (o *AccountResponse) GetJwk() JsonWebKey {
	if o == nil {
		var ret JsonWebKey
		return ret
	}

	return o.Jwk
}

// GetJwkOk returns a tuple with the Jwk field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetJwkOk() (*JsonWebKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jwk, true
}

// SetJwk sets field value
func (o *AccountResponse) SetJwk(v JsonWebKey) {
	o.Jwk = v
}

// GetKeyThumbprint returns the KeyThumbprint field value
func (o *AccountResponse) GetKeyThumbprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyThumbprint
}

// GetKeyThumbprintOk returns a tuple with the KeyThumbprint field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetKeyThumbprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyThumbprint, true
}

// SetKeyThumbprint sets field value
func (o *AccountResponse) SetKeyThumbprint(v string) {
	o.KeyThumbprint = v
}

// GetStatus returns the Status field value
func (o *AccountResponse) GetStatus() AccountStatus {
	if o == nil {
		var ret AccountStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetStatusOk() (*AccountStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccountResponse) SetStatus(v AccountStatus) {
	o.Status = v
}

// GetTermsOfServiceAgreed returns the TermsOfServiceAgreed field value if set, zero value otherwise.
func (o *AccountResponse) GetTermsOfServiceAgreed() bool {
	if o == nil || utils.IsNil(o.TermsOfServiceAgreed) {
		var ret bool
		return ret
	}
	return *o.TermsOfServiceAgreed
}

// GetTermsOfServiceAgreedOk returns a tuple with the TermsOfServiceAgreed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountResponse) GetTermsOfServiceAgreedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.TermsOfServiceAgreed) {
		return nil, false
	}
	return o.TermsOfServiceAgreed, true
}

// HasTermsOfServiceAgreed returns a boolean if a field has been set.
func (o *AccountResponse) HasTermsOfServiceAgreed() bool {
	if o != nil && !utils.IsNil(o.TermsOfServiceAgreed) {
		return true
	}

	return false
}

// SetTermsOfServiceAgreed gets a reference to the given bool and assigns it to the TermsOfServiceAgreed field.
func (o *AccountResponse) SetTermsOfServiceAgreed(v bool) {
	o.TermsOfServiceAgreed = &v
}

func (o AccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if !utils.IsNil(o.CompromisedAt) {
		toSerialize["compromisedAt"] = o.CompromisedAt
	}
	if o.CompromissionReason.IsSet() {
		toSerialize["compromissionReason"] = o.CompromissionReason.Get()
	}
	if !utils.IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	toSerialize["createdAt"] = o.CreatedAt
	if !utils.IsNil(o.EabName) {
		toSerialize["eabName"] = o.EabName
	}
	if !utils.IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !utils.IsNil(o.InitialIp) {
		toSerialize["initialIp"] = o.InitialIp
	}
	toSerialize["jwk"] = o.Jwk
	toSerialize["keyThumbprint"] = o.KeyThumbprint
	toSerialize["status"] = o.Status
	if !utils.IsNil(o.TermsOfServiceAgreed) {
		toSerialize["termsOfServiceAgreed"] = o.TermsOfServiceAgreed
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"createdAt",
		"jwk",
		"keyThumbprint",
		"status",
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

	varAccountResponse := _AccountResponse{}

	err = json.Unmarshal(data, &varAccountResponse)

	if err != nil {
		return err
	}

	*o = AccountResponse(varAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "compromisedAt")
		delete(additionalProperties, "compromissionReason")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "eabName")
		delete(additionalProperties, "expirationDate")
		delete(additionalProperties, "initialIp")
		delete(additionalProperties, "jwk")
		delete(additionalProperties, "keyThumbprint")
		delete(additionalProperties, "status")
		delete(additionalProperties, "termsOfServiceAgreed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountResponse struct {
	value *AccountResponse
	isSet bool
}

func (v NullableAccountResponse) Get() *AccountResponse {
	return v.value
}

func (v *NullableAccountResponse) Set(val *AccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountResponse(val *AccountResponse) *NullableAccountResponse {
	return &NullableAccountResponse{value: val, isSet: true}
}

func (v NullableAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
