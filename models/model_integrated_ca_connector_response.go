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

// checks if the IntegratedCAConnectorResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IntegratedCAConnectorResponse{}

// IntegratedCAConnectorResponse struct for IntegratedCAConnectorResponse
type IntegratedCAConnectorResponse struct {
	// Object internal ID
	Id                   string                     `json:"_id"`
	Name                 string                     `json:"name"`
	Type                 string                     `json:"type"`
	CaKey                NullableSecretString       `json:"caKey,omitempty"`
	CaCert               utils.NullableString       `json:"caCert,omitempty"`
	CrlPath              utils.NullableString       `json:"crlPath,omitempty"`
	CrlLifetime          utils.NullableString       `json:"crlLifetime,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	CertType             utils.NullableString       `json:"certType,omitempty"`
	SignAlg              utils.NullableString       `json:"signAlg,omitempty"`
	CrtLifetime          utils.NullableString       `json:"crtLifetime,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	CrtBackDate          utils.NullableString       `json:"crtBackDate,omitempty" validate:"regexp=^([0-9]+) *(ms|millisecond|milliseconds|s|second|seconds|m|minute|minutes|h|hour|hours|d|day|days)$"`
	CheckPop             utils.NullableBool         `json:"checkPop,omitempty"`
	Queue                utils.NullableString       `json:"queue,omitempty"`
	Status               NullablePKIConnectorStatus `json:"status,omitempty"`
	CryptoType           string                     `json:"cryptoType"`
	AdditionalProperties map[string]interface{}
}

type _IntegratedCAConnectorResponse IntegratedCAConnectorResponse

// NewIntegratedCAConnectorResponse instantiates a new IntegratedCAConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegratedCAConnectorResponse(id string, name string, type_ string, cryptoType string) *IntegratedCAConnectorResponse {
	this := IntegratedCAConnectorResponse{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.CryptoType = cryptoType
	return &this
}

// NewIntegratedCAConnectorResponseWithDefaults instantiates a new IntegratedCAConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegratedCAConnectorResponseWithDefaults() *IntegratedCAConnectorResponse {
	this := IntegratedCAConnectorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *IntegratedCAConnectorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IntegratedCAConnectorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IntegratedCAConnectorResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *IntegratedCAConnectorResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntegratedCAConnectorResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntegratedCAConnectorResponse) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *IntegratedCAConnectorResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntegratedCAConnectorResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IntegratedCAConnectorResponse) SetType(v string) {
	o.Type = v
}

// GetCaKey returns the CaKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegratedCAConnectorResponse) GetCaKey() SecretString {
	if o == nil || utils.IsNil(o.CaKey.Get()) {
		var ret SecretString
		return ret
	}
	return *o.CaKey.Get()
}

// GetCaKeyOk returns a tuple with the CaKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegratedCAConnectorResponse) GetCaKeyOk() (*SecretString, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaKey.Get(), o.CaKey.IsSet()
}

// HasCaKey returns a boolean if a field has been set.
func (o *IntegratedCAConnectorResponse) HasCaKey() bool {
	if o != nil && o.CaKey.IsSet() {
		return true
	}

	return false
}

// SetCaKey gets a reference to the given NullableSecretString and assigns it to the CaKey field.
func (o *IntegratedCAConnectorResponse) SetCaKey(v SecretString) {
	o.CaKey.Set(&v)
}

// SetCaKeyNil sets the value for CaKey to be an explicit nil
func (o *IntegratedCAConnectorResponse) SetCaKeyNil() {
	o.CaKey.Set(nil)
}

// UnsetCaKey ensures that no value is present for CaKey, not even an explicit nil
func (o *IntegratedCAConnectorResponse) UnsetCaKey() {
	o.CaKey.Unset()
}

// GetCaCert returns the CaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegratedCAConnectorResponse) GetCaCert() string {
	if o == nil || utils.IsNil(o.CaCert.Get()) {
		var ret string
		return ret
	}
	return *o.CaCert.Get()
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegratedCAConnectorResponse) GetCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCert.Get(), o.CaCert.IsSet()
}

// HasCaCert returns a boolean if a field has been set.
func (o *IntegratedCAConnectorResponse) HasCaCert() bool {
	if o != nil && o.CaCert.IsSet() {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given NullableString and assigns it to the CaCert field.
func (o *IntegratedCAConnectorResponse) SetCaCert(v string) {
	o.CaCert.Set(&v)
}

// SetCaCertNil sets the value for CaCert to be an explicit nil
func (o *IntegratedCAConnectorResponse) SetCaCertNil() {
	o.CaCert.Set(nil)
}

// UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
func (o *IntegratedCAConnectorResponse) UnsetCaCert() {
	o.CaCert.Unset()
}

// GetCrlPath returns the CrlPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegratedCAConnectorResponse) GetCrlPath() string {
	if o == nil || utils.IsNil(o.CrlPath.Get()) {
		var ret string
		return ret
	}
	return *o.CrlPath.Get()
}

// GetCrlPathOk returns a tuple with the CrlPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegratedCAConnectorResponse) GetCrlPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CrlPath.Get(), o.CrlPath.IsSet()
}

// HasCrlPath returns a boolean if a field has been set.
func (o *IntegratedCAConnectorResponse) HasCrlPath() bool {
	if o != nil && o.CrlPath.IsSet() {
		return true
	}

	return false
}

// SetCrlPath gets a reference to the given NullableString and assigns it to the CrlPath field.
func (o *IntegratedCAConnectorResponse) SetCrlPath(v string) {
	o.CrlPath.Set(&v)
}

// SetCrlPathNil sets the value for CrlPath to be an explicit nil
func (o *IntegratedCAConnectorResponse) SetCrlPathNil() {
	o.CrlPath.Set(nil)
}

// UnsetCrlPath ensures that no value is present for CrlPath, not even an explicit nil
func (o *IntegratedCAConnectorResponse) UnsetCrlPath() {
	o.CrlPath.Unset()
}

// GetCrlLifetime returns the CrlLifetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegratedCAConnectorResponse) GetCrlLifetime() string {
	if o == nil || utils.IsNil(o.CrlLifetime.Get()) {
		var ret string
		return ret
	}
	return *o.CrlLifetime.Get()
}

// GetCrlLifetimeOk returns a tuple with the CrlLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegratedCAConnectorResponse) GetCrlLifetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CrlLifetime.Get(), o.CrlLifetime.IsSet()
}

// HasCrlLifetime returns a boolean if a field has been set.
func (o *IntegratedCAConnectorResponse) HasCrlLifetime() bool {
	if o != nil && o.CrlLifetime.IsSet() {
		return true
	}

	return false
}

// SetCrlLifetime gets a reference to the given NullableString and assigns it to the CrlLifetime field.
func (o *IntegratedCAConnectorResponse) SetCrlLifetime(v string) {
	o.CrlLifetime.Set(&v)
}

// SetCrlLifetimeNil sets the value for CrlLifetime to be an explicit nil
func (o *IntegratedCAConnectorResponse) SetCrlLifetimeNil() {
	o.CrlLifetime.Set(nil)
}

// UnsetCrlLifetime ensures that no value is present for CrlLifetime, not even an explicit nil
func (o *IntegratedCAConnectorResponse) UnsetCrlLifetime() {
	o.CrlLifetime.Unset()
}

// GetCertType returns the CertType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegratedCAConnectorResponse) GetCertType() string {
	if o == nil || utils.IsNil(o.CertType.Get()) {
		var ret string
		return ret
	}
	return *o.CertType.Get()
}

// GetCertTypeOk returns a tuple with the CertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegratedCAConnectorResponse) GetCertTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CertType.Get(), o.CertType.IsSet()
}

// HasCertType returns a boolean if a field has been set.
func (o *IntegratedCAConnectorResponse) HasCertType() bool {
	if o != nil && o.CertType.IsSet() {
		return true
	}

	return false
}

// SetCertType gets a reference to the given NullableString and assigns it to the CertType field.
func (o *IntegratedCAConnectorResponse) SetCertType(v string) {
	o.CertType.Set(&v)
}

// SetCertTypeNil sets the value for CertType to be an explicit nil
func (o *IntegratedCAConnectorResponse) SetCertTypeNil() {
	o.CertType.Set(nil)
}

// UnsetCertType ensures that no value is present for CertType, not even an explicit nil
func (o *IntegratedCAConnectorResponse) UnsetCertType() {
	o.CertType.Unset()
}

// GetSignAlg returns the SignAlg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegratedCAConnectorResponse) GetSignAlg() string {
	if o == nil || utils.IsNil(o.SignAlg.Get()) {
		var ret string
		return ret
	}
	return *o.SignAlg.Get()
}

// GetSignAlgOk returns a tuple with the SignAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegratedCAConnectorResponse) GetSignAlgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SignAlg.Get(), o.SignAlg.IsSet()
}

// HasSignAlg returns a boolean if a field has been set.
func (o *IntegratedCAConnectorResponse) HasSignAlg() bool {
	if o != nil && o.SignAlg.IsSet() {
		return true
	}

	return false
}

// SetSignAlg gets a reference to the given NullableString and assigns it to the SignAlg field.
func (o *IntegratedCAConnectorResponse) SetSignAlg(v string) {
	o.SignAlg.Set(&v)
}

// SetSignAlgNil sets the value for SignAlg to be an explicit nil
func (o *IntegratedCAConnectorResponse) SetSignAlgNil() {
	o.SignAlg.Set(nil)
}

// UnsetSignAlg ensures that no value is present for SignAlg, not even an explicit nil
func (o *IntegratedCAConnectorResponse) UnsetSignAlg() {
	o.SignAlg.Unset()
}

// GetCrtLifetime returns the CrtLifetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegratedCAConnectorResponse) GetCrtLifetime() string {
	if o == nil || utils.IsNil(o.CrtLifetime.Get()) {
		var ret string
		return ret
	}
	return *o.CrtLifetime.Get()
}

// GetCrtLifetimeOk returns a tuple with the CrtLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegratedCAConnectorResponse) GetCrtLifetimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CrtLifetime.Get(), o.CrtLifetime.IsSet()
}

// HasCrtLifetime returns a boolean if a field has been set.
func (o *IntegratedCAConnectorResponse) HasCrtLifetime() bool {
	if o != nil && o.CrtLifetime.IsSet() {
		return true
	}

	return false
}

// SetCrtLifetime gets a reference to the given NullableString and assigns it to the CrtLifetime field.
func (o *IntegratedCAConnectorResponse) SetCrtLifetime(v string) {
	o.CrtLifetime.Set(&v)
}

// SetCrtLifetimeNil sets the value for CrtLifetime to be an explicit nil
func (o *IntegratedCAConnectorResponse) SetCrtLifetimeNil() {
	o.CrtLifetime.Set(nil)
}

// UnsetCrtLifetime ensures that no value is present for CrtLifetime, not even an explicit nil
func (o *IntegratedCAConnectorResponse) UnsetCrtLifetime() {
	o.CrtLifetime.Unset()
}

// GetCrtBackDate returns the CrtBackDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegratedCAConnectorResponse) GetCrtBackDate() string {
	if o == nil || utils.IsNil(o.CrtBackDate.Get()) {
		var ret string
		return ret
	}
	return *o.CrtBackDate.Get()
}

// GetCrtBackDateOk returns a tuple with the CrtBackDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegratedCAConnectorResponse) GetCrtBackDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CrtBackDate.Get(), o.CrtBackDate.IsSet()
}

// HasCrtBackDate returns a boolean if a field has been set.
func (o *IntegratedCAConnectorResponse) HasCrtBackDate() bool {
	if o != nil && o.CrtBackDate.IsSet() {
		return true
	}

	return false
}

// SetCrtBackDate gets a reference to the given NullableString and assigns it to the CrtBackDate field.
func (o *IntegratedCAConnectorResponse) SetCrtBackDate(v string) {
	o.CrtBackDate.Set(&v)
}

// SetCrtBackDateNil sets the value for CrtBackDate to be an explicit nil
func (o *IntegratedCAConnectorResponse) SetCrtBackDateNil() {
	o.CrtBackDate.Set(nil)
}

// UnsetCrtBackDate ensures that no value is present for CrtBackDate, not even an explicit nil
func (o *IntegratedCAConnectorResponse) UnsetCrtBackDate() {
	o.CrtBackDate.Unset()
}

// GetCheckPop returns the CheckPop field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegratedCAConnectorResponse) GetCheckPop() bool {
	if o == nil || utils.IsNil(o.CheckPop.Get()) {
		var ret bool
		return ret
	}
	return *o.CheckPop.Get()
}

// GetCheckPopOk returns a tuple with the CheckPop field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegratedCAConnectorResponse) GetCheckPopOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CheckPop.Get(), o.CheckPop.IsSet()
}

// HasCheckPop returns a boolean if a field has been set.
func (o *IntegratedCAConnectorResponse) HasCheckPop() bool {
	if o != nil && o.CheckPop.IsSet() {
		return true
	}

	return false
}

// SetCheckPop gets a reference to the given NullableBool and assigns it to the CheckPop field.
func (o *IntegratedCAConnectorResponse) SetCheckPop(v bool) {
	o.CheckPop.Set(&v)
}

// SetCheckPopNil sets the value for CheckPop to be an explicit nil
func (o *IntegratedCAConnectorResponse) SetCheckPopNil() {
	o.CheckPop.Set(nil)
}

// UnsetCheckPop ensures that no value is present for CheckPop, not even an explicit nil
func (o *IntegratedCAConnectorResponse) UnsetCheckPop() {
	o.CheckPop.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegratedCAConnectorResponse) GetQueue() string {
	if o == nil || utils.IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegratedCAConnectorResponse) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *IntegratedCAConnectorResponse) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *IntegratedCAConnectorResponse) SetQueue(v string) {
	o.Queue.Set(&v)
}

// SetQueueNil sets the value for Queue to be an explicit nil
func (o *IntegratedCAConnectorResponse) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *IntegratedCAConnectorResponse) UnsetQueue() {
	o.Queue.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegratedCAConnectorResponse) GetStatus() PKIConnectorStatus {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret PKIConnectorStatus
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegratedCAConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *IntegratedCAConnectorResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullablePKIConnectorStatus and assigns it to the Status field.
func (o *IntegratedCAConnectorResponse) SetStatus(v PKIConnectorStatus) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *IntegratedCAConnectorResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *IntegratedCAConnectorResponse) UnsetStatus() {
	o.Status.Unset()
}

// GetCryptoType returns the CryptoType field value
func (o *IntegratedCAConnectorResponse) GetCryptoType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CryptoType
}

// GetCryptoTypeOk returns a tuple with the CryptoType field value
// and a boolean to check if the value has been set.
func (o *IntegratedCAConnectorResponse) GetCryptoTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CryptoType, true
}

// SetCryptoType sets field value
func (o *IntegratedCAConnectorResponse) SetCryptoType(v string) {
	o.CryptoType = v
}

func (o IntegratedCAConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegratedCAConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.CaKey.IsSet() {
		toSerialize["caKey"] = o.CaKey.Get()
	}
	if o.CaCert.IsSet() {
		toSerialize["caCert"] = o.CaCert.Get()
	}
	if o.CrlPath.IsSet() {
		toSerialize["crlPath"] = o.CrlPath.Get()
	}
	if o.CrlLifetime.IsSet() {
		toSerialize["crlLifetime"] = o.CrlLifetime.Get()
	}
	if o.CertType.IsSet() {
		toSerialize["certType"] = o.CertType.Get()
	}
	if o.SignAlg.IsSet() {
		toSerialize["signAlg"] = o.SignAlg.Get()
	}
	if o.CrtLifetime.IsSet() {
		toSerialize["crtLifetime"] = o.CrtLifetime.Get()
	}
	if o.CrtBackDate.IsSet() {
		toSerialize["crtBackDate"] = o.CrtBackDate.Get()
	}
	if o.CheckPop.IsSet() {
		toSerialize["checkPop"] = o.CheckPop.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	toSerialize["cryptoType"] = o.CryptoType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IntegratedCAConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"name",
		"type",
		"cryptoType",
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

	varIntegratedCAConnectorResponse := _IntegratedCAConnectorResponse{}

	err = json.Unmarshal(data, &varIntegratedCAConnectorResponse)

	if err != nil {
		return err
	}

	*o = IntegratedCAConnectorResponse(varIntegratedCAConnectorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "caKey")
		delete(additionalProperties, "caCert")
		delete(additionalProperties, "crlPath")
		delete(additionalProperties, "crlLifetime")
		delete(additionalProperties, "certType")
		delete(additionalProperties, "signAlg")
		delete(additionalProperties, "crtLifetime")
		delete(additionalProperties, "crtBackDate")
		delete(additionalProperties, "checkPop")
		delete(additionalProperties, "queue")
		delete(additionalProperties, "status")
		delete(additionalProperties, "cryptoType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIntegratedCAConnectorResponse struct {
	value *IntegratedCAConnectorResponse
	isSet bool
}

func (v NullableIntegratedCAConnectorResponse) Get() *IntegratedCAConnectorResponse {
	return v.value
}

func (v *NullableIntegratedCAConnectorResponse) Set(val *IntegratedCAConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegratedCAConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegratedCAConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegratedCAConnectorResponse(val *IntegratedCAConnectorResponse) *NullableIntegratedCAConnectorResponse {
	return &NullableIntegratedCAConnectorResponse{value: val, isSet: true}
}

func (v NullableIntegratedCAConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegratedCAConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
