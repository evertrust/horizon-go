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

// checks if the RawCredentialsResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RawCredentialsResponse{}

// RawCredentialsResponse struct for RawCredentialsResponse
type RawCredentialsResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// These credentials secret
	Secret SecretString `json:"secret"`
	Type   string       `json:"type"`
	// These credentials description
	Description utils.NullableString `json:"description,omitempty"`
	// The expiration date of these credentials
	Expires utils.NullableInt64 `json:"expires,omitempty"`
	// These credentials identifying name
	Name string `json:"name"`
	// On which configuration the credentials are usable
	Targets              []string             `json:"targets,omitempty"`
	Triggers             *CredentialsTriggers `json:"triggers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RawCredentialsResponse RawCredentialsResponse

// NewRawCredentialsResponse instantiates a new RawCredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRawCredentialsResponse(id string, secret SecretString, type_ string, name string) *RawCredentialsResponse {
	this := RawCredentialsResponse{}
	this.Id = id
	this.Secret = secret
	this.Type = type_
	this.Name = name
	return &this
}

// NewRawCredentialsResponseWithDefaults instantiates a new RawCredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRawCredentialsResponseWithDefaults() *RawCredentialsResponse {
	this := RawCredentialsResponse{}
	return &this
}

// GetId returns the Id field value
func (o *RawCredentialsResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RawCredentialsResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RawCredentialsResponse) SetId(v string) {
	o.Id = v
}

// GetSecret returns the Secret field value
func (o *RawCredentialsResponse) GetSecret() SecretString {
	if o == nil {
		var ret SecretString
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *RawCredentialsResponse) GetSecretOk() (*SecretString, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *RawCredentialsResponse) SetSecret(v SecretString) {
	o.Secret = v
}

// GetType returns the Type field value
func (o *RawCredentialsResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RawCredentialsResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RawCredentialsResponse) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RawCredentialsResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RawCredentialsResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *RawCredentialsResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *RawCredentialsResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *RawCredentialsResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *RawCredentialsResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RawCredentialsResponse) GetExpires() int64 {
	if o == nil || utils.IsNil(o.Expires.Get()) {
		var ret int64
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RawCredentialsResponse) GetExpiresOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *RawCredentialsResponse) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableInt64 and assigns it to the Expires field.
func (o *RawCredentialsResponse) SetExpires(v int64) {
	o.Expires.Set(&v)
}

// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *RawCredentialsResponse) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *RawCredentialsResponse) UnsetExpires() {
	o.Expires.Unset()
}

// GetName returns the Name field value
func (o *RawCredentialsResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RawCredentialsResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RawCredentialsResponse) SetName(v string) {
	o.Name = v
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *RawCredentialsResponse) GetTargets() []string {
	if o == nil || utils.IsNil(o.Targets) {
		var ret []string
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawCredentialsResponse) GetTargetsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *RawCredentialsResponse) HasTargets() bool {
	if o != nil && !utils.IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []string and assigns it to the Targets field.
func (o *RawCredentialsResponse) SetTargets(v []string) {
	o.Targets = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *RawCredentialsResponse) GetTriggers() CredentialsTriggers {
	if o == nil || utils.IsNil(o.Triggers) {
		var ret CredentialsTriggers
		return ret
	}
	return *o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RawCredentialsResponse) GetTriggersOk() (*CredentialsTriggers, bool) {
	if o == nil || utils.IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *RawCredentialsResponse) HasTriggers() bool {
	if o != nil && !utils.IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given CredentialsTriggers and assigns it to the Triggers field.
func (o *RawCredentialsResponse) SetTriggers(v CredentialsTriggers) {
	o.Triggers = &v
}

func (o RawCredentialsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RawCredentialsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["secret"] = o.Secret
	toSerialize["type"] = o.Type
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Expires.IsSet() {
		toSerialize["expires"] = o.Expires.Get()
	}
	toSerialize["name"] = o.Name
	if !utils.IsNil(o.Targets) {
		toSerialize["targets"] = o.Targets
	}
	if !utils.IsNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RawCredentialsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"secret",
		"type",
		"name",
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

	varRawCredentialsResponse := _RawCredentialsResponse{}

	err = json.Unmarshal(data, &varRawCredentialsResponse)

	if err != nil {
		return err
	}

	*o = RawCredentialsResponse(varRawCredentialsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "expires")
		delete(additionalProperties, "name")
		delete(additionalProperties, "targets")
		delete(additionalProperties, "triggers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRawCredentialsResponse struct {
	value *RawCredentialsResponse
	isSet bool
}

func (v NullableRawCredentialsResponse) Get() *RawCredentialsResponse {
	return v.value
}

func (v *NullableRawCredentialsResponse) Set(val *RawCredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRawCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRawCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRawCredentialsResponse(val *RawCredentialsResponse) *NullableRawCredentialsResponse {
	return &NullableRawCredentialsResponse{value: val, isSet: true}
}

func (v NullableRawCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRawCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
