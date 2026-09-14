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

// checks if the ExternalAccountBindingResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExternalAccountBindingResponse{}

// ExternalAccountBindingResponse struct for ExternalAccountBindingResponse
type ExternalAccountBindingResponse struct {
	// Object internal ID
	Id              string   `json:"_id"`
	AllowedProfiles []string `json:"allowedProfiles"`
	// The date when the EAB was compromised
	CompromisedAt *int64 `json:"compromisedAt,omitempty"`
	// One of: `unspecified`, `keycompromise`, `cacompromise`, `affiliationchange`, `superseded`, `cessationofoperation`
	CompromissionReason     utils.NullableString            `json:"compromissionReason,omitempty"`
	CreatedAt               int64                           `json:"createdAt"`
	Description             *string                         `json:"description,omitempty"`
	EabPolicy               string                          `json:"eabPolicy"`
	EmailConstraint         *string                         `json:"emailConstraint,omitempty"`
	ExpirationDate          *int64                          `json:"expirationDate,omitempty"`
	IdentifierConstraint    *string                         `json:"identifierConstraint,omitempty"`
	MacKey                  string                          `json:"macKey"`
	MacKeyAlgorithm         ExternalAccountBindingAlgorithm `json:"macKeyAlgorithm"`
	MacKeyId                string                          `json:"macKeyId"`
	Name                    string                          `json:"name"`
	NumberOfKeyRegeneration int64                           `json:"numberOfKeyRegeneration"`
	Status                  ExternalAccountBindingStatus    `json:"status"`
	ValidationMethods       []AcmeAuthorizationType         `json:"validationMethods"`
	AdditionalProperties    map[string]interface{}
}

type _ExternalAccountBindingResponse ExternalAccountBindingResponse

// NewExternalAccountBindingResponse instantiates a new ExternalAccountBindingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAccountBindingResponse(id string, allowedProfiles []string, createdAt int64, eabPolicy string, macKey string, macKeyAlgorithm ExternalAccountBindingAlgorithm, macKeyId string, name string, numberOfKeyRegeneration int64, status ExternalAccountBindingStatus, validationMethods []AcmeAuthorizationType) *ExternalAccountBindingResponse {
	this := ExternalAccountBindingResponse{}
	this.Id = id
	this.AllowedProfiles = allowedProfiles
	this.CreatedAt = createdAt
	this.EabPolicy = eabPolicy
	this.MacKey = macKey
	this.MacKeyAlgorithm = macKeyAlgorithm
	this.MacKeyId = macKeyId
	this.Name = name
	this.NumberOfKeyRegeneration = numberOfKeyRegeneration
	this.Status = status
	this.ValidationMethods = validationMethods
	return &this
}

// NewExternalAccountBindingResponseWithDefaults instantiates a new ExternalAccountBindingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAccountBindingResponseWithDefaults() *ExternalAccountBindingResponse {
	this := ExternalAccountBindingResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ExternalAccountBindingResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalAccountBindingResponse) SetId(v string) {
	o.Id = v
}

// GetAllowedProfiles returns the AllowedProfiles field value
func (o *ExternalAccountBindingResponse) GetAllowedProfiles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedProfiles
}

// GetAllowedProfilesOk returns a tuple with the AllowedProfiles field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetAllowedProfilesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedProfiles, true
}

// SetAllowedProfiles sets field value
func (o *ExternalAccountBindingResponse) SetAllowedProfiles(v []string) {
	o.AllowedProfiles = v
}

// GetCompromisedAt returns the CompromisedAt field value if set, zero value otherwise.
func (o *ExternalAccountBindingResponse) GetCompromisedAt() int64 {
	if o == nil || utils.IsNil(o.CompromisedAt) {
		var ret int64
		return ret
	}
	return *o.CompromisedAt
}

// GetCompromisedAtOk returns a tuple with the CompromisedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetCompromisedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CompromisedAt) {
		return nil, false
	}
	return o.CompromisedAt, true
}

// HasCompromisedAt returns a boolean if a field has been set.
func (o *ExternalAccountBindingResponse) HasCompromisedAt() bool {
	if o != nil && !utils.IsNil(o.CompromisedAt) {
		return true
	}

	return false
}

// SetCompromisedAt gets a reference to the given int64 and assigns it to the CompromisedAt field.
func (o *ExternalAccountBindingResponse) SetCompromisedAt(v int64) {
	o.CompromisedAt = &v
}

// GetCompromissionReason returns the CompromissionReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalAccountBindingResponse) GetCompromissionReason() string {
	if o == nil || utils.IsNil(o.CompromissionReason.Get()) {
		var ret string
		return ret
	}
	return *o.CompromissionReason.Get()
}

// GetCompromissionReasonOk returns a tuple with the CompromissionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalAccountBindingResponse) GetCompromissionReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompromissionReason.Get(), o.CompromissionReason.IsSet()
}

// HasCompromissionReason returns a boolean if a field has been set.
func (o *ExternalAccountBindingResponse) HasCompromissionReason() bool {
	if o != nil && o.CompromissionReason.IsSet() {
		return true
	}

	return false
}

// SetCompromissionReason gets a reference to the given NullableString and assigns it to the CompromissionReason field.
func (o *ExternalAccountBindingResponse) SetCompromissionReason(v string) {
	o.CompromissionReason.Set(&v)
}

// SetCompromissionReasonNil sets the value for CompromissionReason to be an explicit nil
func (o *ExternalAccountBindingResponse) SetCompromissionReasonNil() {
	o.CompromissionReason.Set(nil)
}

// UnsetCompromissionReason ensures that no value is present for CompromissionReason, not even an explicit nil
func (o *ExternalAccountBindingResponse) UnsetCompromissionReason() {
	o.CompromissionReason.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *ExternalAccountBindingResponse) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ExternalAccountBindingResponse) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExternalAccountBindingResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExternalAccountBindingResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExternalAccountBindingResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEabPolicy returns the EabPolicy field value
func (o *ExternalAccountBindingResponse) GetEabPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EabPolicy
}

// GetEabPolicyOk returns a tuple with the EabPolicy field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetEabPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EabPolicy, true
}

// SetEabPolicy sets field value
func (o *ExternalAccountBindingResponse) SetEabPolicy(v string) {
	o.EabPolicy = v
}

// GetEmailConstraint returns the EmailConstraint field value if set, zero value otherwise.
func (o *ExternalAccountBindingResponse) GetEmailConstraint() string {
	if o == nil || utils.IsNil(o.EmailConstraint) {
		var ret string
		return ret
	}
	return *o.EmailConstraint
}

// GetEmailConstraintOk returns a tuple with the EmailConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetEmailConstraintOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EmailConstraint) {
		return nil, false
	}
	return o.EmailConstraint, true
}

// HasEmailConstraint returns a boolean if a field has been set.
func (o *ExternalAccountBindingResponse) HasEmailConstraint() bool {
	if o != nil && !utils.IsNil(o.EmailConstraint) {
		return true
	}

	return false
}

// SetEmailConstraint gets a reference to the given string and assigns it to the EmailConstraint field.
func (o *ExternalAccountBindingResponse) SetEmailConstraint(v string) {
	o.EmailConstraint = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ExternalAccountBindingResponse) GetExpirationDate() int64 {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		var ret int64
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetExpirationDateOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ExternalAccountBindingResponse) HasExpirationDate() bool {
	if o != nil && !utils.IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given int64 and assigns it to the ExpirationDate field.
func (o *ExternalAccountBindingResponse) SetExpirationDate(v int64) {
	o.ExpirationDate = &v
}

// GetIdentifierConstraint returns the IdentifierConstraint field value if set, zero value otherwise.
func (o *ExternalAccountBindingResponse) GetIdentifierConstraint() string {
	if o == nil || utils.IsNil(o.IdentifierConstraint) {
		var ret string
		return ret
	}
	return *o.IdentifierConstraint
}

// GetIdentifierConstraintOk returns a tuple with the IdentifierConstraint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetIdentifierConstraintOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IdentifierConstraint) {
		return nil, false
	}
	return o.IdentifierConstraint, true
}

// HasIdentifierConstraint returns a boolean if a field has been set.
func (o *ExternalAccountBindingResponse) HasIdentifierConstraint() bool {
	if o != nil && !utils.IsNil(o.IdentifierConstraint) {
		return true
	}

	return false
}

// SetIdentifierConstraint gets a reference to the given string and assigns it to the IdentifierConstraint field.
func (o *ExternalAccountBindingResponse) SetIdentifierConstraint(v string) {
	o.IdentifierConstraint = &v
}

// GetMacKey returns the MacKey field value
func (o *ExternalAccountBindingResponse) GetMacKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacKey
}

// GetMacKeyOk returns a tuple with the MacKey field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetMacKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacKey, true
}

// SetMacKey sets field value
func (o *ExternalAccountBindingResponse) SetMacKey(v string) {
	o.MacKey = v
}

// GetMacKeyAlgorithm returns the MacKeyAlgorithm field value
func (o *ExternalAccountBindingResponse) GetMacKeyAlgorithm() ExternalAccountBindingAlgorithm {
	if o == nil {
		var ret ExternalAccountBindingAlgorithm
		return ret
	}

	return o.MacKeyAlgorithm
}

// GetMacKeyAlgorithmOk returns a tuple with the MacKeyAlgorithm field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetMacKeyAlgorithmOk() (*ExternalAccountBindingAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacKeyAlgorithm, true
}

// SetMacKeyAlgorithm sets field value
func (o *ExternalAccountBindingResponse) SetMacKeyAlgorithm(v ExternalAccountBindingAlgorithm) {
	o.MacKeyAlgorithm = v
}

// GetMacKeyId returns the MacKeyId field value
func (o *ExternalAccountBindingResponse) GetMacKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MacKeyId
}

// GetMacKeyIdOk returns a tuple with the MacKeyId field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetMacKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MacKeyId, true
}

// SetMacKeyId sets field value
func (o *ExternalAccountBindingResponse) SetMacKeyId(v string) {
	o.MacKeyId = v
}

// GetName returns the Name field value
func (o *ExternalAccountBindingResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalAccountBindingResponse) SetName(v string) {
	o.Name = v
}

// GetNumberOfKeyRegeneration returns the NumberOfKeyRegeneration field value
func (o *ExternalAccountBindingResponse) GetNumberOfKeyRegeneration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumberOfKeyRegeneration
}

// GetNumberOfKeyRegenerationOk returns a tuple with the NumberOfKeyRegeneration field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetNumberOfKeyRegenerationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfKeyRegeneration, true
}

// SetNumberOfKeyRegeneration sets field value
func (o *ExternalAccountBindingResponse) SetNumberOfKeyRegeneration(v int64) {
	o.NumberOfKeyRegeneration = v
}

// GetStatus returns the Status field value
func (o *ExternalAccountBindingResponse) GetStatus() ExternalAccountBindingStatus {
	if o == nil {
		var ret ExternalAccountBindingStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetStatusOk() (*ExternalAccountBindingStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExternalAccountBindingResponse) SetStatus(v ExternalAccountBindingStatus) {
	o.Status = v
}

// GetValidationMethods returns the ValidationMethods field value
func (o *ExternalAccountBindingResponse) GetValidationMethods() []AcmeAuthorizationType {
	if o == nil {
		var ret []AcmeAuthorizationType
		return ret
	}

	return o.ValidationMethods
}

// GetValidationMethodsOk returns a tuple with the ValidationMethods field value
// and a boolean to check if the value has been set.
func (o *ExternalAccountBindingResponse) GetValidationMethodsOk() ([]AcmeAuthorizationType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationMethods, true
}

// SetValidationMethods sets field value
func (o *ExternalAccountBindingResponse) SetValidationMethods(v []AcmeAuthorizationType) {
	o.ValidationMethods = v
}

func (o ExternalAccountBindingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalAccountBindingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["allowedProfiles"] = o.AllowedProfiles
	if !utils.IsNil(o.CompromisedAt) {
		toSerialize["compromisedAt"] = o.CompromisedAt
	}
	if o.CompromissionReason.IsSet() {
		toSerialize["compromissionReason"] = o.CompromissionReason.Get()
	}
	toSerialize["createdAt"] = o.CreatedAt
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["eabPolicy"] = o.EabPolicy
	if !utils.IsNil(o.EmailConstraint) {
		toSerialize["emailConstraint"] = o.EmailConstraint
	}
	if !utils.IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !utils.IsNil(o.IdentifierConstraint) {
		toSerialize["identifierConstraint"] = o.IdentifierConstraint
	}
	toSerialize["macKey"] = o.MacKey
	toSerialize["macKeyAlgorithm"] = o.MacKeyAlgorithm
	toSerialize["macKeyId"] = o.MacKeyId
	toSerialize["name"] = o.Name
	toSerialize["numberOfKeyRegeneration"] = o.NumberOfKeyRegeneration
	toSerialize["status"] = o.Status
	toSerialize["validationMethods"] = o.ValidationMethods

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExternalAccountBindingResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"allowedProfiles",
		"createdAt",
		"eabPolicy",
		"macKey",
		"macKeyAlgorithm",
		"macKeyId",
		"name",
		"numberOfKeyRegeneration",
		"status",
		"validationMethods",
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

	varExternalAccountBindingResponse := _ExternalAccountBindingResponse{}

	err = json.Unmarshal(data, &varExternalAccountBindingResponse)

	if err != nil {
		return err
	}

	*o = ExternalAccountBindingResponse(varExternalAccountBindingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "allowedProfiles")
		delete(additionalProperties, "compromisedAt")
		delete(additionalProperties, "compromissionReason")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "description")
		delete(additionalProperties, "eabPolicy")
		delete(additionalProperties, "emailConstraint")
		delete(additionalProperties, "expirationDate")
		delete(additionalProperties, "identifierConstraint")
		delete(additionalProperties, "macKey")
		delete(additionalProperties, "macKeyAlgorithm")
		delete(additionalProperties, "macKeyId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "numberOfKeyRegeneration")
		delete(additionalProperties, "status")
		delete(additionalProperties, "validationMethods")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExternalAccountBindingResponse struct {
	value *ExternalAccountBindingResponse
	isSet bool
}

func (v NullableExternalAccountBindingResponse) Get() *ExternalAccountBindingResponse {
	return v.value
}

func (v *NullableExternalAccountBindingResponse) Set(val *ExternalAccountBindingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAccountBindingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAccountBindingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAccountBindingResponse(val *ExternalAccountBindingResponse) *NullableExternalAccountBindingResponse {
	return &NullableExternalAccountBindingResponse{value: val, isSet: true}
}

func (v NullableExternalAccountBindingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAccountBindingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
