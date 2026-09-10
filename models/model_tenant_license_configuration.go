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

// checks if the TenantLicenseConfiguration type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TenantLicenseConfiguration{}

// TenantLicenseConfiguration License configuration for a tenant. At least one of certificate or dcv must be provided.
type TenantLicenseConfiguration struct {
	Certificate *TenantCertificateLicenseConfiguration `json:"certificate,omitempty"`
	Dcv         *TenantDCVLicenseConfiguration         `json:"dcv,omitempty"`
	// Custom license expiration date for this tenant
	Expiration           *int64 `json:"expiration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TenantLicenseConfiguration TenantLicenseConfiguration

// NewTenantLicenseConfiguration instantiates a new TenantLicenseConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantLicenseConfiguration() *TenantLicenseConfiguration {
	this := TenantLicenseConfiguration{}
	return &this
}

// NewTenantLicenseConfigurationWithDefaults instantiates a new TenantLicenseConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantLicenseConfigurationWithDefaults() *TenantLicenseConfiguration {
	this := TenantLicenseConfiguration{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *TenantLicenseConfiguration) GetCertificate() TenantCertificateLicenseConfiguration {
	if o == nil || utils.IsNil(o.Certificate) {
		var ret TenantCertificateLicenseConfiguration
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantLicenseConfiguration) GetCertificateOk() (*TenantCertificateLicenseConfiguration, bool) {
	if o == nil || utils.IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *TenantLicenseConfiguration) HasCertificate() bool {
	if o != nil && !utils.IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given TenantCertificateLicenseConfiguration and assigns it to the Certificate field.
func (o *TenantLicenseConfiguration) SetCertificate(v TenantCertificateLicenseConfiguration) {
	o.Certificate = &v
}

// GetDcv returns the Dcv field value if set, zero value otherwise.
func (o *TenantLicenseConfiguration) GetDcv() TenantDCVLicenseConfiguration {
	if o == nil || utils.IsNil(o.Dcv) {
		var ret TenantDCVLicenseConfiguration
		return ret
	}
	return *o.Dcv
}

// GetDcvOk returns a tuple with the Dcv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantLicenseConfiguration) GetDcvOk() (*TenantDCVLicenseConfiguration, bool) {
	if o == nil || utils.IsNil(o.Dcv) {
		return nil, false
	}
	return o.Dcv, true
}

// HasDcv returns a boolean if a field has been set.
func (o *TenantLicenseConfiguration) HasDcv() bool {
	if o != nil && !utils.IsNil(o.Dcv) {
		return true
	}

	return false
}

// SetDcv gets a reference to the given TenantDCVLicenseConfiguration and assigns it to the Dcv field.
func (o *TenantLicenseConfiguration) SetDcv(v TenantDCVLicenseConfiguration) {
	o.Dcv = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *TenantLicenseConfiguration) GetExpiration() int64 {
	if o == nil || utils.IsNil(o.Expiration) {
		var ret int64
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantLicenseConfiguration) GetExpirationOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *TenantLicenseConfiguration) HasExpiration() bool {
	if o != nil && !utils.IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given int64 and assigns it to the Expiration field.
func (o *TenantLicenseConfiguration) SetExpiration(v int64) {
	o.Expiration = &v
}

func (o TenantLicenseConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantLicenseConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !utils.IsNil(o.Dcv) {
		toSerialize["dcv"] = o.Dcv
	}
	if !utils.IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TenantLicenseConfiguration) UnmarshalJSON(data []byte) (err error) {
	varTenantLicenseConfiguration := _TenantLicenseConfiguration{}

	err = json.Unmarshal(data, &varTenantLicenseConfiguration)

	if err != nil {
		return err
	}

	*o = TenantLicenseConfiguration(varTenantLicenseConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "certificate")
		delete(additionalProperties, "dcv")
		delete(additionalProperties, "expiration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTenantLicenseConfiguration struct {
	value *TenantLicenseConfiguration
	isSet bool
}

func (v NullableTenantLicenseConfiguration) Get() *TenantLicenseConfiguration {
	return v.value
}

func (v *NullableTenantLicenseConfiguration) Set(val *TenantLicenseConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantLicenseConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantLicenseConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantLicenseConfiguration(val *TenantLicenseConfiguration) *NullableTenantLicenseConfiguration {
	return &NullableTenantLicenseConfiguration{value: val, isSet: true}
}

func (v NullableTenantLicenseConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantLicenseConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
