/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the AcmeEnrollConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AcmeEnrollConnectorResponse{}

// AcmeEnrollConnectorResponse Used to enroll certificate using the ACME protocol with DNS challenge on compatible public PKI
type AcmeEnrollConnectorResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// Email to associate with the account
	AccountEmail utils.NullableString `json:"accountEmail,omitempty"`
	// The key type to use to generate the account key
	AccountKeyType string `json:"accountKeyType" validate:"regexp=(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512)(\\\\\\\\+(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512))?"`
	// Url of the account on the ACME directory
	AccountUrl string `json:"accountUrl"`
	// DNS Provider configuration to provision the DNS challenge. Available from `2.7.7`
	DnsChallengeProvider DnsChallengeProviders `json:"dnsChallengeProvider"`
	// The dictionary provider
	DomainDictionaryProvider NullableDomainDictionaryProviders `json:"domainDictionaryProvider,omitempty"`
	// `password` credentials name to use for External Account Binding
	Eab utils.NullableString `json:"eab,omitempty"`
	// The directory url of the ACME endpoint
	EndPoint string               `json:"endPoint"`
	Name     string               `json:"name"`
	Proxy    utils.NullableString `json:"proxy,omitempty"`
	Queue    utils.NullableString `json:"queue,omitempty"`
	// If enable, regenerate the account (does not need to be specified on creation)
	RotateAccount        utils.NullableBool         `json:"rotateAccount,omitempty"`
	Status               NullablePKIConnectorStatus `json:"status,omitempty"`
	Timeout              utils.NullableString       `json:"timeout" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	Type                 string                     `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _AcmeEnrollConnectorResponse AcmeEnrollConnectorResponse

// NewAcmeEnrollConnectorResponse instantiates a new AcmeEnrollConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcmeEnrollConnectorResponse(id string, accountKeyType string, accountUrl string, dnsChallengeProvider DnsChallengeProviders, endPoint string, name string, timeout utils.NullableString, type_ string) *AcmeEnrollConnectorResponse {
	this := AcmeEnrollConnectorResponse{}
	this.Id = id
	this.AccountKeyType = accountKeyType
	this.AccountUrl = accountUrl
	this.DnsChallengeProvider = dnsChallengeProvider
	this.EndPoint = endPoint
	this.Name = name
	this.Timeout = timeout
	this.Type = type_
	return &this
}

// NewAcmeEnrollConnectorResponseWithDefaults instantiates a new AcmeEnrollConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcmeEnrollConnectorResponseWithDefaults() *AcmeEnrollConnectorResponse {
	this := AcmeEnrollConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AcmeEnrollConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AcmeEnrollConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AcmeEnrollConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetAccountEmail returns the AccountEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeEnrollConnectorResponse) GetAccountEmail() string {
	if o == nil || utils.IsNil(o.AccountEmail.Get()) {
		var ret string
		return ret
	}
	return *o.AccountEmail.Get()
}

// GetAccountEmailOk returns a tuple with the AccountEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeEnrollConnectorResponse) GetAccountEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountEmail.Get(), o.AccountEmail.IsSet()
}

// HasAccountEmail returns a boolean if a field has been set.
func (o *AcmeEnrollConnectorResponse) HasAccountEmail() bool {
	if o != nil && o.AccountEmail.IsSet() {
		return true
	}

	return false
}

// SetAccountEmail gets a reference to the given NullableString and assigns it to the AccountEmail field.
func (o *AcmeEnrollConnectorResponse) SetAccountEmail(v string) {
	o.AccountEmail.Set(&v)
}

// SetAccountEmailNil sets the value for AccountEmail to be an explicit nil
func (o *AcmeEnrollConnectorResponse) SetAccountEmailNil() {
	o.AccountEmail.Set(nil)
}

// UnsetAccountEmail ensures that no value is present for AccountEmail, not even an explicit nil
func (o *AcmeEnrollConnectorResponse) UnsetAccountEmail() {
	o.AccountEmail.Unset()
}

// GetAccountKeyType returns the AccountKeyType field value
func (o *AcmeEnrollConnectorResponse) GetAccountKeyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountKeyType
}

// GetAccountKeyTypeOk returns a tuple with the AccountKeyType field value
// and a boolean to check if the value has been set.
func (o *AcmeEnrollConnectorResponse) GetAccountKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountKeyType, true
}

// SetAccountKeyType sets field value
func (o *AcmeEnrollConnectorResponse) SetAccountKeyType(v string) {
	o.AccountKeyType = v
}

// GetAccountUrl returns the AccountUrl field value
func (o *AcmeEnrollConnectorResponse) GetAccountUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountUrl
}

// GetAccountUrlOk returns a tuple with the AccountUrl field value
// and a boolean to check if the value has been set.
func (o *AcmeEnrollConnectorResponse) GetAccountUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountUrl, true
}

// SetAccountUrl sets field value
func (o *AcmeEnrollConnectorResponse) SetAccountUrl(v string) {
	o.AccountUrl = v
}

// GetDnsChallengeProvider returns the DnsChallengeProvider field value
func (o *AcmeEnrollConnectorResponse) GetDnsChallengeProvider() DnsChallengeProviders {
	if o == nil {
		var ret DnsChallengeProviders
		return ret
	}

	return o.DnsChallengeProvider
}

// GetDnsChallengeProviderOk returns a tuple with the DnsChallengeProvider field value
// and a boolean to check if the value has been set.
func (o *AcmeEnrollConnectorResponse) GetDnsChallengeProviderOk() (*DnsChallengeProviders, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsChallengeProvider, true
}

// SetDnsChallengeProvider sets field value
func (o *AcmeEnrollConnectorResponse) SetDnsChallengeProvider(v DnsChallengeProviders) {
	o.DnsChallengeProvider = v
}

// GetDomainDictionaryProvider returns the DomainDictionaryProvider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeEnrollConnectorResponse) GetDomainDictionaryProvider() DomainDictionaryProviders {
	if o == nil || utils.IsNil(o.DomainDictionaryProvider.Get()) {
		var ret DomainDictionaryProviders
		return ret
	}
	return *o.DomainDictionaryProvider.Get()
}

// GetDomainDictionaryProviderOk returns a tuple with the DomainDictionaryProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeEnrollConnectorResponse) GetDomainDictionaryProviderOk() (*DomainDictionaryProviders, bool) {
	if o == nil {
		return nil, false
	}
	return o.DomainDictionaryProvider.Get(), o.DomainDictionaryProvider.IsSet()
}

// HasDomainDictionaryProvider returns a boolean if a field has been set.
func (o *AcmeEnrollConnectorResponse) HasDomainDictionaryProvider() bool {
	if o != nil && o.DomainDictionaryProvider.IsSet() {
		return true
	}

	return false
}

// SetDomainDictionaryProvider gets a reference to the given NullableDomainDictionaryProviders and assigns it to the DomainDictionaryProvider field.
func (o *AcmeEnrollConnectorResponse) SetDomainDictionaryProvider(v DomainDictionaryProviders) {
	o.DomainDictionaryProvider.Set(&v)
}

// SetDomainDictionaryProviderNil sets the value for DomainDictionaryProvider to be an explicit nil
func (o *AcmeEnrollConnectorResponse) SetDomainDictionaryProviderNil() {
	o.DomainDictionaryProvider.Set(nil)
}

// UnsetDomainDictionaryProvider ensures that no value is present for DomainDictionaryProvider, not even an explicit nil
func (o *AcmeEnrollConnectorResponse) UnsetDomainDictionaryProvider() {
	o.DomainDictionaryProvider.Unset()
}

// GetEab returns the Eab field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeEnrollConnectorResponse) GetEab() string {
	if o == nil || utils.IsNil(o.Eab.Get()) {
		var ret string
		return ret
	}
	return *o.Eab.Get()
}

// GetEabOk returns a tuple with the Eab field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeEnrollConnectorResponse) GetEabOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Eab.Get(), o.Eab.IsSet()
}

// HasEab returns a boolean if a field has been set.
func (o *AcmeEnrollConnectorResponse) HasEab() bool {
	if o != nil && o.Eab.IsSet() {
		return true
	}

	return false
}

// SetEab gets a reference to the given NullableString and assigns it to the Eab field.
func (o *AcmeEnrollConnectorResponse) SetEab(v string) {
	o.Eab.Set(&v)
}

// SetEabNil sets the value for Eab to be an explicit nil
func (o *AcmeEnrollConnectorResponse) SetEabNil() {
	o.Eab.Set(nil)
}

// UnsetEab ensures that no value is present for Eab, not even an explicit nil
func (o *AcmeEnrollConnectorResponse) UnsetEab() {
	o.Eab.Unset()
}

// GetEndPoint returns the EndPoint field value
func (o *AcmeEnrollConnectorResponse) GetEndPoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value
// and a boolean to check if the value has been set.
func (o *AcmeEnrollConnectorResponse) GetEndPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPoint, true
}

// SetEndPoint sets field value
func (o *AcmeEnrollConnectorResponse) SetEndPoint(v string) {
	o.EndPoint = v
}

// GetName returns the Name field value
func (o *AcmeEnrollConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AcmeEnrollConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AcmeEnrollConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeEnrollConnectorResponse) GetProxy() string {
	if o == nil || utils.IsNil(o.Proxy.Get()) {
		var ret string
		return ret
	}
	return *o.Proxy.Get()
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeEnrollConnectorResponse) GetProxyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proxy.Get(), o.Proxy.IsSet()
}

// HasProxy returns a boolean if a field has been set.
func (o *AcmeEnrollConnectorResponse) HasProxy() bool {
	if o != nil && o.Proxy.IsSet() {
		return true
	}

	return false
}

// SetProxy gets a reference to the given NullableString and assigns it to the Proxy field.
func (o *AcmeEnrollConnectorResponse) SetProxy(v string) {
	o.Proxy.Set(&v)
}

// SetProxyNil sets the value for Proxy to be an explicit nil
func (o *AcmeEnrollConnectorResponse) SetProxyNil() {
	o.Proxy.Set(nil)
}

// UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
func (o *AcmeEnrollConnectorResponse) UnsetProxy() {
	o.Proxy.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeEnrollConnectorResponse) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeEnrollConnectorResponse) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *AcmeEnrollConnectorResponse) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *AcmeEnrollConnectorResponse) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *AcmeEnrollConnectorResponse) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *AcmeEnrollConnectorResponse) UnsetQueue() {
	o.Queue.Unset()
}

// GetRotateAccount returns the RotateAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeEnrollConnectorResponse) GetRotateAccount() bool {
	if o == nil || utils.IsNil(o.RotateAccount.Get()) {
		var ret bool
		return ret
	}
	return *o.RotateAccount.Get()
}

// GetRotateAccountOk returns a tuple with the RotateAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeEnrollConnectorResponse) GetRotateAccountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RotateAccount.Get(), o.RotateAccount.IsSet()
}

// HasRotateAccount returns a boolean if a field has been set.
func (o *AcmeEnrollConnectorResponse) HasRotateAccount() bool {
	if o != nil && o.RotateAccount.IsSet() {
		return true
	}

	return false
}

// SetRotateAccount gets a reference to the given NullableBool and assigns it to the RotateAccount field.
func (o *AcmeEnrollConnectorResponse) SetRotateAccount(v bool) {
	o.RotateAccount.Set(&v)
}

// SetRotateAccountNil sets the value for RotateAccount to be an explicit nil
func (o *AcmeEnrollConnectorResponse) SetRotateAccountNil() {
	o.RotateAccount.Set(nil)
}

// UnsetRotateAccount ensures that no value is present for RotateAccount, not even an explicit nil
func (o *AcmeEnrollConnectorResponse) UnsetRotateAccount() {
	o.RotateAccount.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AcmeEnrollConnectorResponse) GetStatus() PKIConnectorStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret PKIConnectorStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeEnrollConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *AcmeEnrollConnectorResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePKIConnectorStatus and assigns it to the Status field.
func (o *AcmeEnrollConnectorResponse) SetStatus(v PKIConnectorStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *AcmeEnrollConnectorResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *AcmeEnrollConnectorResponse) UnsetStatus() {
	o.Status.Unset()
}

// GetTimeout returns the Timeout field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AcmeEnrollConnectorResponse) GetTimeout() string {
	if o == nil || o.Timeout.Get() == nil {
		var ret string
		return ret
	}

	return *o.Timeout.Get()
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AcmeEnrollConnectorResponse) GetTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeout.Get(), o.Timeout.IsSet()
}

// SetTimeout sets field value
func (o *AcmeEnrollConnectorResponse) SetTimeout(v string) {
	o.Timeout.Set(&v)
}

// GetType returns the Type field value
func (o *AcmeEnrollConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AcmeEnrollConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AcmeEnrollConnectorResponse) SetType(v string) {
	o.Type = v
}

func (o AcmeEnrollConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcmeEnrollConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	if o.AccountEmail.IsSet() {
		toSerialize["accountEmail"] = o.AccountEmail.Get()
	}
	toSerialize["accountKeyType"] = o.AccountKeyType
	toSerialize["accountUrl"] = o.AccountUrl
	toSerialize["dnsChallengeProvider"] = o.DnsChallengeProvider
	if o.DomainDictionaryProvider.IsSet() {
		toSerialize["domainDictionaryProvider"] = o.DomainDictionaryProvider.Get()
	}
	if o.Eab.IsSet() {
		toSerialize["eab"] = o.Eab.Get()
	}
	toSerialize["endPoint"] = o.EndPoint
	toSerialize["name"] = o.Name
	if o.Proxy.IsSet() {
		toSerialize["proxy"] = o.Proxy.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.RotateAccount.IsSet() {
		toSerialize["rotateAccount"] = o.RotateAccount.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	toSerialize["timeout"] = o.Timeout.Get()
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcmeEnrollConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"accountKeyType",
		"accountUrl",
		"dnsChallengeProvider",
		"endPoint",
		"name",
		"timeout",
		"type",
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

	varAcmeEnrollConnectorResponse := _AcmeEnrollConnectorResponse{}

	err = json.Unmarshal(data, &varAcmeEnrollConnectorResponse)

	if err != nil {
		return err
	}

	*o = AcmeEnrollConnectorResponse(varAcmeEnrollConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "accountEmail")
		delete(additionalProperties, "accountKeyType")
		delete(additionalProperties, "accountUrl")
		delete(additionalProperties, "dnsChallengeProvider")
		delete(additionalProperties, "domainDictionaryProvider")
		delete(additionalProperties, "eab")
		delete(additionalProperties, "endPoint")
		delete(additionalProperties, "name")
		delete(additionalProperties, "proxy")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "rotateAccount")
		delete(additionalProperties, "status")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcmeEnrollConnectorResponse struct {
	value *AcmeEnrollConnectorResponse
	isSet bool
}

func (v NullableAcmeEnrollConnectorResponse) Get() *AcmeEnrollConnectorResponse {
	return v.value
}

func (v *NullableAcmeEnrollConnectorResponse) Set(val *AcmeEnrollConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcmeEnrollConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcmeEnrollConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcmeEnrollConnectorResponse(val *AcmeEnrollConnectorResponse) *NullableAcmeEnrollConnectorResponse {
	return &NullableAcmeEnrollConnectorResponse{value: val, isSet: true}
}

func (v NullableAcmeEnrollConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcmeEnrollConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
