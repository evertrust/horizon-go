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

// checks if the ServiceAccountResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServiceAccountResponse{}

// ServiceAccountResponse struct for ServiceAccountResponse
type ServiceAccountResponse struct {
	// Object internal ID
	Id *string `json:"_id,omitempty"`
	// If true, this object was externally provisioned and cannot be edited
	Readonly bool `json:"readonly"`
	// Maximum duration in the future the JWT `iat` claim is allowed to be. Must be set together with `iatPastRestriction`.
	IatFutureRestriction *string `json:"iatFutureRestriction,omitempty"`
	// Maximum duration in the past the JWT `iat` claim is allowed to be. Must be set together with `iatFutureRestriction`.
	IatPastRestriction *string `json:"iatPastRestriction,omitempty"`
	// Template string used to compute the identifier of the principal authenticated by this service account.
	IdentifierMapping *string `json:"identifierMapping,omitempty"`
	// Allowed clock skew when validating JWT time-based claims.
	JwtAllowedClockSkew *string `json:"jwtAllowedClockSkew,omitempty"`
	// Internal name for the service account
	Name string `json:"name"`
	// List of permissions to apply for successfully validated JWTs
	Permissions []Permission `json:"permissions"`
	// List of roles to apply for successfully validated JWTs
	Roles       []string                  `json:"roles"`
	TrustConfig ServiceAccountTrustConfig `json:"trustConfig"`
	// List of rules to apply on top of signature verification for the incoming JWT to be trusted
	ValidationRules      []string `json:"validationRules"`
	AdditionalProperties map[string]interface{}
}

type _ServiceAccountResponse ServiceAccountResponse

// NewServiceAccountResponse instantiates a new ServiceAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountResponse(readonly bool, name string, permissions []Permission, roles []string, trustConfig ServiceAccountTrustConfig, validationRules []string) *ServiceAccountResponse {
	this := ServiceAccountResponse{}
	this.Readonly = readonly
	this.Name = name
	this.Permissions = permissions
	this.Roles = roles
	this.TrustConfig = trustConfig
	this.ValidationRules = validationRules
	return &this
}

// NewServiceAccountResponseWithDefaults instantiates a new ServiceAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountResponseWithDefaults() *ServiceAccountResponse {
	this := ServiceAccountResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceAccountResponse) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponse) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceAccountResponse) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceAccountResponse) SetId(v string) {
	o.Id = &v
}

// GetReadonly returns the Readonly field value
func (o *ServiceAccountResponse) GetReadonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponse) GetReadonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Readonly, true
}

// SetReadonly sets field value
func (o *ServiceAccountResponse) SetReadonly(v bool) {
	o.Readonly = v
}

// GetIatFutureRestriction returns the IatFutureRestriction field value if set, zero value otherwise.
func (o *ServiceAccountResponse) GetIatFutureRestriction() string {
	if o == nil || utils.IsNil(o.IatFutureRestriction) {
		var ret string
		return ret
	}
	return *o.IatFutureRestriction
}

// GetIatFutureRestrictionOk returns a tuple with the IatFutureRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponse) GetIatFutureRestrictionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IatFutureRestriction) {
		return nil, false
	}
	return o.IatFutureRestriction, true
}

// HasIatFutureRestriction returns a boolean if a field has been set.
func (o *ServiceAccountResponse) HasIatFutureRestriction() bool {
	if o != nil && !utils.IsNil(o.IatFutureRestriction) {
		return true
	}

	return false
}

// SetIatFutureRestriction gets a reference to the given string and assigns it to the IatFutureRestriction field.
func (o *ServiceAccountResponse) SetIatFutureRestriction(v string) {
	o.IatFutureRestriction = &v
}

// GetIatPastRestriction returns the IatPastRestriction field value if set, zero value otherwise.
func (o *ServiceAccountResponse) GetIatPastRestriction() string {
	if o == nil || utils.IsNil(o.IatPastRestriction) {
		var ret string
		return ret
	}
	return *o.IatPastRestriction
}

// GetIatPastRestrictionOk returns a tuple with the IatPastRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponse) GetIatPastRestrictionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IatPastRestriction) {
		return nil, false
	}
	return o.IatPastRestriction, true
}

// HasIatPastRestriction returns a boolean if a field has been set.
func (o *ServiceAccountResponse) HasIatPastRestriction() bool {
	if o != nil && !utils.IsNil(o.IatPastRestriction) {
		return true
	}

	return false
}

// SetIatPastRestriction gets a reference to the given string and assigns it to the IatPastRestriction field.
func (o *ServiceAccountResponse) SetIatPastRestriction(v string) {
	o.IatPastRestriction = &v
}

// GetIdentifierMapping returns the IdentifierMapping field value if set, zero value otherwise.
func (o *ServiceAccountResponse) GetIdentifierMapping() string {
	if o == nil || utils.IsNil(o.IdentifierMapping) {
		var ret string
		return ret
	}
	return *o.IdentifierMapping
}

// GetIdentifierMappingOk returns a tuple with the IdentifierMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponse) GetIdentifierMappingOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IdentifierMapping) {
		return nil, false
	}
	return o.IdentifierMapping, true
}

// HasIdentifierMapping returns a boolean if a field has been set.
func (o *ServiceAccountResponse) HasIdentifierMapping() bool {
	if o != nil && !utils.IsNil(o.IdentifierMapping) {
		return true
	}

	return false
}

// SetIdentifierMapping gets a reference to the given string and assigns it to the IdentifierMapping field.
func (o *ServiceAccountResponse) SetIdentifierMapping(v string) {
	o.IdentifierMapping = &v
}

// GetJwtAllowedClockSkew returns the JwtAllowedClockSkew field value if set, zero value otherwise.
func (o *ServiceAccountResponse) GetJwtAllowedClockSkew() string {
	if o == nil || utils.IsNil(o.JwtAllowedClockSkew) {
		var ret string
		return ret
	}
	return *o.JwtAllowedClockSkew
}

// GetJwtAllowedClockSkewOk returns a tuple with the JwtAllowedClockSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponse) GetJwtAllowedClockSkewOk() (*string, bool) {
	if o == nil || utils.IsNil(o.JwtAllowedClockSkew) {
		return nil, false
	}
	return o.JwtAllowedClockSkew, true
}

// HasJwtAllowedClockSkew returns a boolean if a field has been set.
func (o *ServiceAccountResponse) HasJwtAllowedClockSkew() bool {
	if o != nil && !utils.IsNil(o.JwtAllowedClockSkew) {
		return true
	}

	return false
}

// SetJwtAllowedClockSkew gets a reference to the given string and assigns it to the JwtAllowedClockSkew field.
func (o *ServiceAccountResponse) SetJwtAllowedClockSkew(v string) {
	o.JwtAllowedClockSkew = &v
}

// GetName returns the Name field value
func (o *ServiceAccountResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceAccountResponse) SetName(v string) {
	o.Name = v
}

// GetPermissions returns the Permissions field value
func (o *ServiceAccountResponse) GetPermissions() []Permission {
	if o == nil {
		var ret []Permission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponse) GetPermissionsOk() ([]Permission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *ServiceAccountResponse) SetPermissions(v []Permission) {
	o.Permissions = v
}

// GetRoles returns the Roles field value
func (o *ServiceAccountResponse) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponse) GetRolesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *ServiceAccountResponse) SetRoles(v []string) {
	o.Roles = v
}

// GetTrustConfig returns the TrustConfig field value
func (o *ServiceAccountResponse) GetTrustConfig() ServiceAccountTrustConfig {
	if o == nil {
		var ret ServiceAccountTrustConfig
		return ret
	}

	return o.TrustConfig
}

// GetTrustConfigOk returns a tuple with the TrustConfig field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponse) GetTrustConfigOk() (*ServiceAccountTrustConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustConfig, true
}

// SetTrustConfig sets field value
func (o *ServiceAccountResponse) SetTrustConfig(v ServiceAccountTrustConfig) {
	o.TrustConfig = v
}

// GetValidationRules returns the ValidationRules field value
func (o *ServiceAccountResponse) GetValidationRules() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ValidationRules
}

// GetValidationRulesOk returns a tuple with the ValidationRules field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountResponse) GetValidationRulesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidationRules, true
}

// SetValidationRules sets field value
func (o *ServiceAccountResponse) SetValidationRules(v []string) {
	o.ValidationRules = v
}

func (o ServiceAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	toSerialize["readonly"] = o.Readonly
	if !utils.IsNil(o.IatFutureRestriction) {
		toSerialize["iatFutureRestriction"] = o.IatFutureRestriction
	}
	if !utils.IsNil(o.IatPastRestriction) {
		toSerialize["iatPastRestriction"] = o.IatPastRestriction
	}
	if !utils.IsNil(o.IdentifierMapping) {
		toSerialize["identifierMapping"] = o.IdentifierMapping
	}
	if !utils.IsNil(o.JwtAllowedClockSkew) {
		toSerialize["jwtAllowedClockSkew"] = o.JwtAllowedClockSkew
	}
	toSerialize["name"] = o.Name
	toSerialize["permissions"] = o.Permissions
	toSerialize["roles"] = o.Roles
	toSerialize["trustConfig"] = o.TrustConfig
	toSerialize["validationRules"] = o.ValidationRules

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceAccountResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"readonly",
		"name",
		"permissions",
		"roles",
		"trustConfig",
		"validationRules",
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

	varServiceAccountResponse := _ServiceAccountResponse{}

	err = json.Unmarshal(data, &varServiceAccountResponse)

	if err != nil {
		return err
	}

	*o = ServiceAccountResponse(varServiceAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "readonly")
		delete(additionalProperties, "iatFutureRestriction")
		delete(additionalProperties, "iatPastRestriction")
		delete(additionalProperties, "identifierMapping")
		delete(additionalProperties, "jwtAllowedClockSkew")
		delete(additionalProperties, "name")
		delete(additionalProperties, "permissions")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "trustConfig")
		delete(additionalProperties, "validationRules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceAccountResponse struct {
	value *ServiceAccountResponse
	isSet bool
}

func (v NullableServiceAccountResponse) Get() *ServiceAccountResponse {
	return v.value
}

func (v *NullableServiceAccountResponse) Set(val *ServiceAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountResponse(val *ServiceAccountResponse) *NullableServiceAccountResponse {
	return &NullableServiceAccountResponse{value: val, isSet: true}
}

func (v NullableServiceAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
