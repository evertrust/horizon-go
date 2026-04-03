/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the CFCertificationRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CFCertificationRequest{}

// CFCertificationRequest Certification request
type CFCertificationRequest struct {
	// Distinguished name
	Dn         string                `json:"dn"`
	DnElements []CFDistinguishedName `json:"dnElements"`
	// One of `rsa-2048`, `rsa-3072`, `rsa-4096`, `rsa-8192`, `ec-secp256r1`, `ec-secp384r1`, `ec-secp521r1`, `ed-448`, `ed-25519`, `mldsa-44`, `mldsa-65`, `mldsa-87`, `slhdsa-sha2-128s`, `slhdsa-sha2-128f`, `slhdsa-sha2-192s`, `slhdsa-sha2-192f`, `slhdsa-sha2-256s`, `slhdsa-sha2-256f`, `slhdsa-sha2-128ssha256`, `slhdsa-sha2-128fsha256`, `slhdsa-sha2-192ssha512`, `slhdsa-sha2-192fsha512`, `slhdsa-sha2-256ssha512`, `slhdsa-sha2-256fsha512` or `<primary key type>+<alternate key type>`
	KeyType              string                 `json:"keyType" validate:"regexp=(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512)(\\\\\\\\+(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512))?"`
	Pem                  string                 `json:"pem"`
	Sans                 []SubjectAlternateName `json:"sans,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CFCertificationRequest CFCertificationRequest

// NewCFCertificationRequest instantiates a new CFCertificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCFCertificationRequest(dn string, dnElements []CFDistinguishedName, keyType string, pem string) *CFCertificationRequest {
	this := CFCertificationRequest{}
	this.Dn = dn
	this.DnElements = dnElements
	this.KeyType = keyType
	this.Pem = pem
	return &this
}

// NewCFCertificationRequestWithDefaults instantiates a new CFCertificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCFCertificationRequestWithDefaults() *CFCertificationRequest {
	this := CFCertificationRequest{}
	return &this
}

// GetDn returns the Dn field value
func (o *CFCertificationRequest) GetDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dn
}

// GetDnOk returns a tuple with the Dn field value
// and a boolean to check if the value has been set.
func (o *CFCertificationRequest) GetDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dn, true
}

// SetDn sets field value
func (o *CFCertificationRequest) SetDn(v string) {
	o.Dn = v
}

// GetDnElements returns the DnElements field value
func (o *CFCertificationRequest) GetDnElements() []CFDistinguishedName {
	if o == nil {
		var ret []CFDistinguishedName
		return ret
	}

	return o.DnElements
}

// GetDnElementsOk returns a tuple with the DnElements field value
// and a boolean to check if the value has been set.
func (o *CFCertificationRequest) GetDnElementsOk() ([]CFDistinguishedName, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnElements, true
}

// SetDnElements sets field value
func (o *CFCertificationRequest) SetDnElements(v []CFDistinguishedName) {
	o.DnElements = v
}

// GetKeyType returns the KeyType field value
func (o *CFCertificationRequest) GetKeyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *CFCertificationRequest) GetKeyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *CFCertificationRequest) SetKeyType(v string) {
	o.KeyType = v
}

// GetPem returns the Pem field value
func (o *CFCertificationRequest) GetPem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pem
}

// GetPemOk returns a tuple with the Pem field value
// and a boolean to check if the value has been set.
func (o *CFCertificationRequest) GetPemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pem, true
}

// SetPem sets field value
func (o *CFCertificationRequest) SetPem(v string) {
	o.Pem = v
}

// GetSans returns the Sans field value if set, zero value otherwise.
func (o *CFCertificationRequest) GetSans() []SubjectAlternateName {
	if o == nil || utils.IsNil(o.Sans) {
		var ret []SubjectAlternateName
		return ret
	}
	return o.Sans
}

// GetSansOk returns a tuple with the Sans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CFCertificationRequest) GetSansOk() ([]SubjectAlternateName, bool) {
	if o == nil || utils.IsNil(o.Sans) {
		return nil, false
	}
	return o.Sans, true
}

// HasSans returns a boolean if a field has been set.
func (o *CFCertificationRequest) HasSans() bool {
	if o != nil && !utils.IsNil(o.Sans) {
		return true
	}

	return false
}

// SetSans gets a reference to the given []SubjectAlternateName and assigns it to the Sans field.
func (o *CFCertificationRequest) SetSans(v []SubjectAlternateName) {
	o.Sans = v
}

func (o CFCertificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CFCertificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dn"] = o.Dn
	toSerialize["dnElements"] = o.DnElements
	toSerialize["keyType"] = o.KeyType
	toSerialize["pem"] = o.Pem
	if !utils.IsNil(o.Sans) {
		toSerialize["sans"] = o.Sans
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CFCertificationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dn",
		"dnElements",
		"keyType",
		"pem",
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

	varCFCertificationRequest := _CFCertificationRequest{}

	err = json.Unmarshal(data, &varCFCertificationRequest)

	if err != nil {
		return err
	}

	*o = CFCertificationRequest(varCFCertificationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dn")
		delete(additionalProperties, "dnElements")
		delete(additionalProperties, "keyType")
		delete(additionalProperties, "pem")
		delete(additionalProperties, "sans")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCFCertificationRequest struct {
	value *CFCertificationRequest
	isSet bool
}

func (v NullableCFCertificationRequest) Get() *CFCertificationRequest {
	return v.value
}

func (v *NullableCFCertificationRequest) Set(val *CFCertificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCFCertificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCFCertificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCFCertificationRequest(val *CFCertificationRequest) *NullableCFCertificationRequest {
	return &NullableCFCertificationRequest{value: val, isSet: true}
}

func (v NullableCFCertificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCFCertificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
