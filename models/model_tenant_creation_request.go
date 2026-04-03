/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the TenantCreationRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TenantCreationRequest{}

// TenantCreationRequest struct for TenantCreationRequest
type TenantCreationRequest struct {
	// Password to use for the administrator account of this tenant
	AdministratorPassword *string `json:"administratorPassword,omitempty"`
	// A simple description for this tenant
	Description *string `json:"description,omitempty"`
	// Custom license expiration date for this tenant
	LicenseExpiration *int64 `json:"licenseExpiration,omitempty"`
	// License limit for this tenant
	LicenseLimit int64 `json:"licenseLimit"`
	// The tenant internal name
	Name                 string `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _TenantCreationRequest TenantCreationRequest

// NewTenantCreationRequest instantiates a new TenantCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantCreationRequest(licenseLimit int64, name string) *TenantCreationRequest {
	this := TenantCreationRequest{}
	this.LicenseLimit = licenseLimit
	this.Name = name
	return &this
}

// NewTenantCreationRequestWithDefaults instantiates a new TenantCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantCreationRequestWithDefaults() *TenantCreationRequest {
	this := TenantCreationRequest{}
	return &this
}

// GetAdministratorPassword returns the AdministratorPassword field value if set, zero value otherwise.
func (o *TenantCreationRequest) GetAdministratorPassword() string {
	if o == nil || utils.IsNil(o.AdministratorPassword) {
		var ret string
		return ret
	}
	return *o.AdministratorPassword
}

// GetAdministratorPasswordOk returns a tuple with the AdministratorPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreationRequest) GetAdministratorPasswordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AdministratorPassword) {
		return nil, false
	}
	return o.AdministratorPassword, true
}

// HasAdministratorPassword returns a boolean if a field has been set.
func (o *TenantCreationRequest) HasAdministratorPassword() bool {
	if o != nil && !utils.IsNil(o.AdministratorPassword) {
		return true
	}

	return false
}

// SetAdministratorPassword gets a reference to the given string and assigns it to the AdministratorPassword field.
func (o *TenantCreationRequest) SetAdministratorPassword(v string) {
	o.AdministratorPassword = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TenantCreationRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TenantCreationRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TenantCreationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLicenseExpiration returns the LicenseExpiration field value if set, zero value otherwise.
func (o *TenantCreationRequest) GetLicenseExpiration() int64 {
	if o == nil || utils.IsNil(o.LicenseExpiration) {
		var ret int64
		return ret
	}
	return *o.LicenseExpiration
}

// GetLicenseExpirationOk returns a tuple with the LicenseExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantCreationRequest) GetLicenseExpirationOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.LicenseExpiration) {
		return nil, false
	}
	return o.LicenseExpiration, true
}

// HasLicenseExpiration returns a boolean if a field has been set.
func (o *TenantCreationRequest) HasLicenseExpiration() bool {
	if o != nil && !utils.IsNil(o.LicenseExpiration) {
		return true
	}

	return false
}

// SetLicenseExpiration gets a reference to the given int64 and assigns it to the LicenseExpiration field.
func (o *TenantCreationRequest) SetLicenseExpiration(v int64) {
	o.LicenseExpiration = &v
}

// GetLicenseLimit returns the LicenseLimit field value
func (o *TenantCreationRequest) GetLicenseLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LicenseLimit
}

// GetLicenseLimitOk returns a tuple with the LicenseLimit field value
// and a boolean to check if the value has been set.
func (o *TenantCreationRequest) GetLicenseLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseLimit, true
}

// SetLicenseLimit sets field value
func (o *TenantCreationRequest) SetLicenseLimit(v int64) {
	o.LicenseLimit = v
}

// GetName returns the Name field value
func (o *TenantCreationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TenantCreationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TenantCreationRequest) SetName(v string) {
	o.Name = v
}

func (o TenantCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantCreationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AdministratorPassword) {
		toSerialize["administratorPassword"] = o.AdministratorPassword
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.LicenseExpiration) {
		toSerialize["licenseExpiration"] = o.LicenseExpiration
	}
	toSerialize["licenseLimit"] = o.LicenseLimit
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TenantCreationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"licenseLimit",
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

	varTenantCreationRequest := _TenantCreationRequest{}

	err = json.Unmarshal(data, &varTenantCreationRequest)

	if err != nil {
		return err
	}

	*o = TenantCreationRequest(varTenantCreationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "administratorPassword")
		delete(additionalProperties, "description")
		delete(additionalProperties, "licenseExpiration")
		delete(additionalProperties, "licenseLimit")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTenantCreationRequest struct {
	value *TenantCreationRequest
	isSet bool
}

func (v NullableTenantCreationRequest) Get() *TenantCreationRequest {
	return v.value
}

func (v *NullableTenantCreationRequest) Set(val *TenantCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantCreationRequest(val *TenantCreationRequest) *NullableTenantCreationRequest {
	return &NullableTenantCreationRequest{value: val, isSet: true}
}

func (v NullableTenantCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
