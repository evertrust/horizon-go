/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using a JWKS service account  This method of authentication is designed for machine-to-machine clients (CI/CD pipelines, Kubernetes workloads, SaaS automation) that obtain a short-lived JWT from a third-party Identity Provider (e.g. GitHub CI, GitLab CI, Kubernetes).  It requires a service account to be declared in Horizon with: - a name, - one or more JWKS (static content or a JWKS URL) used to verify the JWT signature, - a set of validation rules applied to the JWT claims, - the roles and permissions granted on successful authentication.  The service account name is sent in the `X-API-SVA` header and the JWT in the `X-API-TOKEN` header:  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-SVA: my-service-account\" -H \"X-API-TOKEN: eyJhbGciOiJSUzI1NiIs...\" -H \"Accept: application/json\" ```  Unlike `API-ID`/`API-KEY` or X509 authentication, JWKS service account authentication does not create a `PLAY_SESSION` cookie: the JWT must be presented on every request.  Possible responses are:  | HTTP Response code | Additional information                                                                                                                      | |--------------------|---------------------------------------------------------------------------------------------------------------------------------------------| | 200                | The token was successfully authenticated                                                                                                    | | 401                | Authentication error, the precise cause is not exposed in the response body and is only recorded in the technical logs, not in audit events |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the JsonWebKey type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &JsonWebKey{}

// JsonWebKey struct for JsonWebKey
type JsonWebKey struct {
	// Algorithm
	Alg *string `json:"alg,omitempty"`
	// Exponent (for RSA)
	E *string `json:"e,omitempty"`
	// Symmetric key
	K *string `json:"k,omitempty"`
	// Key type
	Kty *string `json:"kty,omitempty"`
	// Modulus (for RSA)
	N *string `json:"n,omitempty"`
	// Key usage
	Use *string `json:"use,omitempty"`
	// X coordinate (for EC)
	X *string `json:"x,omitempty"`
	// Y coordinate (for EC)
	Y                    *string `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JsonWebKey JsonWebKey

// NewJsonWebKey instantiates a new JsonWebKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonWebKey() *JsonWebKey {
	this := JsonWebKey{}
	return &this
}

// NewJsonWebKeyWithDefaults instantiates a new JsonWebKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonWebKeyWithDefaults() *JsonWebKey {
	this := JsonWebKey{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *JsonWebKey) GetAlg() string {
	if o == nil || utils.IsNil(o.Alg) {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetAlgOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Alg) {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *JsonWebKey) HasAlg() bool {
	if o != nil && !utils.IsNil(o.Alg) {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *JsonWebKey) SetAlg(v string) {
	o.Alg = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *JsonWebKey) GetE() string {
	if o == nil || utils.IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetEOk() (*string, bool) {
	if o == nil || utils.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *JsonWebKey) HasE() bool {
	if o != nil && !utils.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *JsonWebKey) SetE(v string) {
	o.E = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *JsonWebKey) GetK() string {
	if o == nil || utils.IsNil(o.K) {
		var ret string
		return ret
	}
	return *o.K
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetKOk() (*string, bool) {
	if o == nil || utils.IsNil(o.K) {
		return nil, false
	}
	return o.K, true
}

// HasK returns a boolean if a field has been set.
func (o *JsonWebKey) HasK() bool {
	if o != nil && !utils.IsNil(o.K) {
		return true
	}

	return false
}

// SetK gets a reference to the given string and assigns it to the K field.
func (o *JsonWebKey) SetK(v string) {
	o.K = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *JsonWebKey) GetKty() string {
	if o == nil || utils.IsNil(o.Kty) {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetKtyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Kty) {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *JsonWebKey) HasKty() bool {
	if o != nil && !utils.IsNil(o.Kty) {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *JsonWebKey) SetKty(v string) {
	o.Kty = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *JsonWebKey) GetN() string {
	if o == nil || utils.IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetNOk() (*string, bool) {
	if o == nil || utils.IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *JsonWebKey) HasN() bool {
	if o != nil && !utils.IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *JsonWebKey) SetN(v string) {
	o.N = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *JsonWebKey) GetUse() string {
	if o == nil || utils.IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetUseOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *JsonWebKey) HasUse() bool {
	if o != nil && !utils.IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *JsonWebKey) SetUse(v string) {
	o.Use = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *JsonWebKey) GetX() string {
	if o == nil || utils.IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetXOk() (*string, bool) {
	if o == nil || utils.IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *JsonWebKey) HasX() bool {
	if o != nil && !utils.IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *JsonWebKey) SetX(v string) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *JsonWebKey) GetY() string {
	if o == nil || utils.IsNil(o.Y) {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonWebKey) GetYOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *JsonWebKey) HasY() bool {
	if o != nil && !utils.IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *JsonWebKey) SetY(v string) {
	o.Y = &v
}

func (o JsonWebKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JsonWebKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Alg) {
		toSerialize["alg"] = o.Alg
	}
	if !utils.IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !utils.IsNil(o.K) {
		toSerialize["k"] = o.K
	}
	if !utils.IsNil(o.Kty) {
		toSerialize["kty"] = o.Kty
	}
	if !utils.IsNil(o.N) {
		toSerialize["n"] = o.N
	}
	if !utils.IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}
	if !utils.IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !utils.IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JsonWebKey) UnmarshalJSON(data []byte) (err error) {
	varJsonWebKey := _JsonWebKey{}

	err = json.Unmarshal(data, &varJsonWebKey)

	if err != nil {
		return err
	}

	*o = JsonWebKey(varJsonWebKey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "e")
		delete(additionalProperties, "k")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "use")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJsonWebKey struct {
	value *JsonWebKey
	isSet bool
}

func (v NullableJsonWebKey) Get() *JsonWebKey {
	return v.value
}

func (v *NullableJsonWebKey) Set(val *JsonWebKey) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonWebKey) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonWebKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonWebKey(val *JsonWebKey) *NullableJsonWebKey {
	return &NullableJsonWebKey{value: val, isSet: true}
}

func (v NullableJsonWebKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonWebKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
