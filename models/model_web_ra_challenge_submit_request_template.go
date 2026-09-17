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

// checks if the WebRAChallengeSubmitRequestTemplate type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebRAChallengeSubmitRequestTemplate{}

// WebRAChallengeSubmitRequestTemplate The user-data that will be used to generate the certificate.  The identity is either fixed by the profile's certificate template, or supplied here: - if the profile defines a certificate template, the identity is taken from the challenge and the `subject`, `sans` and `extensions` fields are rejected; - if the profile's certificate template is empty, the identity is taken from the `subject`, `sans` and `extensions` fields.
type WebRAChallengeSubmitRequestTemplate struct {
	// The certificate signing request to enroll, in decentralized mode. Mutually exclusive with `keyType`
	Csr *string `json:"csr,omitempty"`
	// List of extension elements that will be used to build the certificate's extensions. Only accepted if the profile's certificate template is empty
	Extensions []CertificateExtensionElement `json:"extensions,omitempty"`
	// The type of key that will be generated, in centralized mode. Mutually exclusive with `csr`
	KeyType *string `json:"keyType,omitempty" validate:"regexp=(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ec-brainpoolp256r1|ec-brainpoolp384r1|ec-brainpoolp512r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512)(\\\\\\\\+(rsa-2048|rsa-3072|rsa-4096|rsa-8192|ec-secp256r1|ec-secp384r1|ec-secp521r1|ec-brainpoolp256r1|ec-brainpoolp384r1|ec-brainpoolp512r1|ed-448|ed-25519|mldsa-44|mldsa-65|mldsa-87|slhdsa-sha2-128s|slhdsa-sha2-128f|slhdsa-sha2-192s|slhdsa-sha2-192f|slhdsa-sha2-256s|slhdsa-sha2-256f|slhdsa-sha2-128ssha256|slhdsa-sha2-128fsha256|slhdsa-sha2-192ssha512|slhdsa-sha2-192fsha512|slhdsa-sha2-256ssha512|slhdsa-sha2-256fsha512))?"`
	// List of metadata elements to set on the certificate. Only the `automation_policy` metadata may be set here, and only to a policy authorized on the profile
	Metadata []CertificateMetadataElement `json:"metadata,omitempty"`
	// List of SAN elements that will be used to build the certificate's Subject Alternative Name. Only accepted if the profile's certificate template is empty
	Sans []ListSANElement `json:"sans,omitempty"`
	// List of DN elements that will be used to build the certificate's Distinguished Name. Only accepted if the profile's certificate template is empty
	Subject              []IndexedDNElement `json:"subject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebRAChallengeSubmitRequestTemplate WebRAChallengeSubmitRequestTemplate

// NewWebRAChallengeSubmitRequestTemplate instantiates a new WebRAChallengeSubmitRequestTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebRAChallengeSubmitRequestTemplate() *WebRAChallengeSubmitRequestTemplate {
	this := WebRAChallengeSubmitRequestTemplate{}
	return &this
}

// NewWebRAChallengeSubmitRequestTemplateWithDefaults instantiates a new WebRAChallengeSubmitRequestTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebRAChallengeSubmitRequestTemplateWithDefaults() *WebRAChallengeSubmitRequestTemplate {
	this := WebRAChallengeSubmitRequestTemplate{}
	return &this
}

// GetCsr returns the Csr field value if set, zero value otherwise.
func (o *WebRAChallengeSubmitRequestTemplate) GetCsr() string {
	if o == nil || utils.IsNil(o.Csr) {
		var ret string
		return ret
	}
	return *o.Csr
}

// GetCsrOk returns a tuple with the Csr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRAChallengeSubmitRequestTemplate) GetCsrOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Csr) {
		return nil, false
	}
	return o.Csr, true
}

// HasCsr returns a boolean if a field has been set.
func (o *WebRAChallengeSubmitRequestTemplate) HasCsr() bool {
	if o != nil && !utils.IsNil(o.Csr) {
		return true
	}

	return false
}

// SetCsr gets a reference to the given string and assigns it to the Csr field.
func (o *WebRAChallengeSubmitRequestTemplate) SetCsr(v string) {
	o.Csr = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *WebRAChallengeSubmitRequestTemplate) GetExtensions() []CertificateExtensionElement {
	if o == nil || utils.IsNil(o.Extensions) {
		var ret []CertificateExtensionElement
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRAChallengeSubmitRequestTemplate) GetExtensionsOk() ([]CertificateExtensionElement, bool) {
	if o == nil || utils.IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *WebRAChallengeSubmitRequestTemplate) HasExtensions() bool {
	if o != nil && !utils.IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []CertificateExtensionElement and assigns it to the Extensions field.
func (o *WebRAChallengeSubmitRequestTemplate) SetExtensions(v []CertificateExtensionElement) {
	o.Extensions = v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *WebRAChallengeSubmitRequestTemplate) GetKeyType() string {
	if o == nil || utils.IsNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRAChallengeSubmitRequestTemplate) GetKeyTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *WebRAChallengeSubmitRequestTemplate) HasKeyType() bool {
	if o != nil && !utils.IsNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *WebRAChallengeSubmitRequestTemplate) SetKeyType(v string) {
	o.KeyType = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *WebRAChallengeSubmitRequestTemplate) GetMetadata() []CertificateMetadataElement {
	if o == nil || utils.IsNil(o.Metadata) {
		var ret []CertificateMetadataElement
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRAChallengeSubmitRequestTemplate) GetMetadataOk() ([]CertificateMetadataElement, bool) {
	if o == nil || utils.IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WebRAChallengeSubmitRequestTemplate) HasMetadata() bool {
	if o != nil && !utils.IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []CertificateMetadataElement and assigns it to the Metadata field.
func (o *WebRAChallengeSubmitRequestTemplate) SetMetadata(v []CertificateMetadataElement) {
	o.Metadata = v
}

// GetSans returns the Sans field value if set, zero value otherwise.
func (o *WebRAChallengeSubmitRequestTemplate) GetSans() []ListSANElement {
	if o == nil || utils.IsNil(o.Sans) {
		var ret []ListSANElement
		return ret
	}
	return o.Sans
}

// GetSansOk returns a tuple with the Sans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRAChallengeSubmitRequestTemplate) GetSansOk() ([]ListSANElement, bool) {
	if o == nil || utils.IsNil(o.Sans) {
		return nil, false
	}
	return o.Sans, true
}

// HasSans returns a boolean if a field has been set.
func (o *WebRAChallengeSubmitRequestTemplate) HasSans() bool {
	if o != nil && !utils.IsNil(o.Sans) {
		return true
	}

	return false
}

// SetSans gets a reference to the given []ListSANElement and assigns it to the Sans field.
func (o *WebRAChallengeSubmitRequestTemplate) SetSans(v []ListSANElement) {
	o.Sans = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *WebRAChallengeSubmitRequestTemplate) GetSubject() []IndexedDNElement {
	if o == nil || utils.IsNil(o.Subject) {
		var ret []IndexedDNElement
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebRAChallengeSubmitRequestTemplate) GetSubjectOk() ([]IndexedDNElement, bool) {
	if o == nil || utils.IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *WebRAChallengeSubmitRequestTemplate) HasSubject() bool {
	if o != nil && !utils.IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given []IndexedDNElement and assigns it to the Subject field.
func (o *WebRAChallengeSubmitRequestTemplate) SetSubject(v []IndexedDNElement) {
	o.Subject = v
}

func (o WebRAChallengeSubmitRequestTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebRAChallengeSubmitRequestTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Csr) {
		toSerialize["csr"] = o.Csr
	}
	if !utils.IsNil(o.Extensions) {
		toSerialize["extensions"] = o.Extensions
	}
	if !utils.IsNil(o.KeyType) {
		toSerialize["keyType"] = o.KeyType
	}
	if !utils.IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !utils.IsNil(o.Sans) {
		toSerialize["sans"] = o.Sans
	}
	if !utils.IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebRAChallengeSubmitRequestTemplate) UnmarshalJSON(data []byte) (err error) {
	varWebRAChallengeSubmitRequestTemplate := _WebRAChallengeSubmitRequestTemplate{}

	err = json.Unmarshal(data, &varWebRAChallengeSubmitRequestTemplate)

	if err != nil {
		return err
	}

	*o = WebRAChallengeSubmitRequestTemplate(varWebRAChallengeSubmitRequestTemplate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "csr")
		delete(additionalProperties, "extensions")
		delete(additionalProperties, "keyType")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "sans")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebRAChallengeSubmitRequestTemplate struct {
	value *WebRAChallengeSubmitRequestTemplate
	isSet bool
}

func (v NullableWebRAChallengeSubmitRequestTemplate) Get() *WebRAChallengeSubmitRequestTemplate {
	return v.value
}

func (v *NullableWebRAChallengeSubmitRequestTemplate) Set(val *WebRAChallengeSubmitRequestTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableWebRAChallengeSubmitRequestTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableWebRAChallengeSubmitRequestTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebRAChallengeSubmitRequestTemplate(val *WebRAChallengeSubmitRequestTemplate) *NullableWebRAChallengeSubmitRequestTemplate {
	return &NullableWebRAChallengeSubmitRequestTemplate{value: val, isSet: true}
}

func (v NullableWebRAChallengeSubmitRequestTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebRAChallengeSubmitRequestTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
