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

// checks if the Order type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Order{}

// Order struct for Order
type Order struct {
	// Object internal ID
	Id string `json:"_id"`
	// Object internal ID
	AccountId      string          `json:"accountId"`
	Authorizations []Authorization `json:"authorizations"`
	// Object internal ID
	Certificate          *string           `json:"certificate,omitempty"`
	ContactEmail         *string           `json:"contactEmail,omitempty"`
	Expires              *int64            `json:"expires,omitempty"`
	InitialIp            *string           `json:"initialIp,omitempty"`
	Label                []OrderLabelInner `json:"label,omitempty"`
	Metadata             map[string]string `json:"metadata"`
	NotAfter             *int64            `json:"notAfter,omitempty"`
	NotBefore            *int64            `json:"notBefore,omitempty"`
	Owner                *string           `json:"owner,omitempty"`
	Profile              string            `json:"profile"`
	RemoveAt             *int64            `json:"removeAt,omitempty"`
	Status               OrderStatus       `json:"status"`
	Team                 *string           `json:"team,omitempty"`
	Thumbprint           *string           `json:"thumbprint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Order Order

// NewOrder instantiates a new Order object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder(id string, accountId string, authorizations []Authorization, metadata map[string]string, profile string, status OrderStatus) *Order {
	this := Order{}
	this.Id = id
	this.AccountId = accountId
	this.Authorizations = authorizations
	this.Metadata = metadata
	this.Profile = profile
	this.Status = status
	return &this
}

// NewOrderWithDefaults instantiates a new Order object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderWithDefaults() *Order {
	this := Order{}
	return &this
}

// GetId returns the Id field value
func (o *Order) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Order) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Order) SetId(v string) {
	o.Id = v
}

// GetAccountId returns the AccountId field value
func (o *Order) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Order) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Order) SetAccountId(v string) {
	o.AccountId = v
}

// GetAuthorizations returns the Authorizations field value
func (o *Order) GetAuthorizations() []Authorization {
	if o == nil {
		var ret []Authorization
		return ret
	}

	return o.Authorizations
}

// GetAuthorizationsOk returns a tuple with the Authorizations field value
// and a boolean to check if the value has been set.
func (o *Order) GetAuthorizationsOk() ([]Authorization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authorizations, true
}

// SetAuthorizations sets field value
func (o *Order) SetAuthorizations(v []Authorization) {
	o.Authorizations = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *Order) GetCertificate() string {
	if o == nil || utils.IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetCertificateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *Order) HasCertificate() bool {
	if o != nil && !utils.IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *Order) SetCertificate(v string) {
	o.Certificate = &v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *Order) GetContactEmail() string {
	if o == nil || utils.IsNil(o.ContactEmail) {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetContactEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ContactEmail) {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *Order) HasContactEmail() bool {
	if o != nil && !utils.IsNil(o.ContactEmail) {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *Order) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *Order) GetExpires() int64 {
	if o == nil || utils.IsNil(o.Expires) {
		var ret int64
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetExpiresOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *Order) HasExpires() bool {
	if o != nil && !utils.IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given int64 and assigns it to the Expires field.
func (o *Order) SetExpires(v int64) {
	o.Expires = &v
}

// GetInitialIp returns the InitialIp field value if set, zero value otherwise.
func (o *Order) GetInitialIp() string {
	if o == nil || utils.IsNil(o.InitialIp) {
		var ret string
		return ret
	}
	return *o.InitialIp
}

// GetInitialIpOk returns a tuple with the InitialIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetInitialIpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InitialIp) {
		return nil, false
	}
	return o.InitialIp, true
}

// HasInitialIp returns a boolean if a field has been set.
func (o *Order) HasInitialIp() bool {
	if o != nil && !utils.IsNil(o.InitialIp) {
		return true
	}

	return false
}

// SetInitialIp gets a reference to the given string and assigns it to the InitialIp field.
func (o *Order) SetInitialIp(v string) {
	o.InitialIp = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Order) GetLabel() []OrderLabelInner {
	if o == nil || utils.IsNil(o.Label) {
		var ret []OrderLabelInner
		return ret
	}
	return o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetLabelOk() ([]OrderLabelInner, bool) {
	if o == nil || utils.IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Order) HasLabel() bool {
	if o != nil && !utils.IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given []OrderLabelInner and assigns it to the Label field.
func (o *Order) SetLabel(v []OrderLabelInner) {
	o.Label = v
}

// GetMetadata returns the Metadata field value
func (o *Order) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *Order) GetMetadataOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *Order) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *Order) GetNotAfter() int64 {
	if o == nil || utils.IsNil(o.NotAfter) {
		var ret int64
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetNotAfterOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.NotAfter) {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *Order) HasNotAfter() bool {
	if o != nil && !utils.IsNil(o.NotAfter) {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given int64 and assigns it to the NotAfter field.
func (o *Order) SetNotAfter(v int64) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *Order) GetNotBefore() int64 {
	if o == nil || utils.IsNil(o.NotBefore) {
		var ret int64
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetNotBeforeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.NotBefore) {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *Order) HasNotBefore() bool {
	if o != nil && !utils.IsNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given int64 and assigns it to the NotBefore field.
func (o *Order) SetNotBefore(v int64) {
	o.NotBefore = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *Order) GetOwner() string {
	if o == nil || utils.IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetOwnerOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *Order) HasOwner() bool {
	if o != nil && !utils.IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *Order) SetOwner(v string) {
	o.Owner = &v
}

// GetProfile returns the Profile field value
func (o *Order) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *Order) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *Order) SetProfile(v string) {
	o.Profile = v
}

// GetRemoveAt returns the RemoveAt field value if set, zero value otherwise.
func (o *Order) GetRemoveAt() int64 {
	if o == nil || utils.IsNil(o.RemoveAt) {
		var ret int64
		return ret
	}
	return *o.RemoveAt
}

// GetRemoveAtOk returns a tuple with the RemoveAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetRemoveAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.RemoveAt) {
		return nil, false
	}
	return o.RemoveAt, true
}

// HasRemoveAt returns a boolean if a field has been set.
func (o *Order) HasRemoveAt() bool {
	if o != nil && !utils.IsNil(o.RemoveAt) {
		return true
	}

	return false
}

// SetRemoveAt gets a reference to the given int64 and assigns it to the RemoveAt field.
func (o *Order) SetRemoveAt(v int64) {
	o.RemoveAt = &v
}

// GetStatus returns the Status field value
func (o *Order) GetStatus() OrderStatus {
	if o == nil {
		var ret OrderStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Order) GetStatusOk() (*OrderStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Order) SetStatus(v OrderStatus) {
	o.Status = v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *Order) GetTeam() string {
	if o == nil || utils.IsNil(o.Team) {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetTeamOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *Order) HasTeam() bool {
	if o != nil && !utils.IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *Order) SetTeam(v string) {
	o.Team = &v
}

// GetThumbprint returns the Thumbprint field value if set, zero value otherwise.
func (o *Order) GetThumbprint() string {
	if o == nil || utils.IsNil(o.Thumbprint) {
		var ret string
		return ret
	}
	return *o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetThumbprintOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Thumbprint) {
		return nil, false
	}
	return o.Thumbprint, true
}

// HasThumbprint returns a boolean if a field has been set.
func (o *Order) HasThumbprint() bool {
	if o != nil && !utils.IsNil(o.Thumbprint) {
		return true
	}

	return false
}

// SetThumbprint gets a reference to the given string and assigns it to the Thumbprint field.
func (o *Order) SetThumbprint(v string) {
	o.Thumbprint = &v
}

func (o Order) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Order) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["accountId"] = o.AccountId
	toSerialize["authorizations"] = o.Authorizations
	if !utils.IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !utils.IsNil(o.ContactEmail) {
		toSerialize["contactEmail"] = o.ContactEmail
	}
	if !utils.IsNil(o.Expires) {
		toSerialize["expires"] = o.Expires
	}
	if !utils.IsNil(o.InitialIp) {
		toSerialize["initialIp"] = o.InitialIp
	}
	if !utils.IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["metadata"] = o.Metadata
	if !utils.IsNil(o.NotAfter) {
		toSerialize["notAfter"] = o.NotAfter
	}
	if !utils.IsNil(o.NotBefore) {
		toSerialize["notBefore"] = o.NotBefore
	}
	if !utils.IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	toSerialize["profile"] = o.Profile
	if !utils.IsNil(o.RemoveAt) {
		toSerialize["removeAt"] = o.RemoveAt
	}
	toSerialize["status"] = o.Status
	if !utils.IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if !utils.IsNil(o.Thumbprint) {
		toSerialize["thumbprint"] = o.Thumbprint
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Order) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"accountId",
		"authorizations",
		"metadata",
		"profile",
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

	varOrder := _Order{}

	err = json.Unmarshal(data, &varOrder)

	if err != nil {
		return err
	}

	*o = Order(varOrder)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "authorizations")
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "contactEmail")
		delete(additionalProperties, "expires")
		delete(additionalProperties, "initialIp")
		delete(additionalProperties, "label")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "notAfter")
		delete(additionalProperties, "notBefore")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "removeAt")
		delete(additionalProperties, "status")
		delete(additionalProperties, "team")
		delete(additionalProperties, "thumbprint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrder struct {
	value *Order
	isSet bool
}

func (v NullableOrder) Get() *Order {
	return v.value
}

func (v *NullableOrder) Set(val *Order) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder(val *Order) *NullableOrder {
	return &NullableOrder{value: val, isSet: true}
}

func (v NullableOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
