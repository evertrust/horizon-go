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

// checks if the DataSourceOutput type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DataSourceOutput{}

// DataSourceOutput struct for DataSourceOutput
type DataSourceOutput struct {
	// Key of the output
	Key string `json:"key"`
	// Determine if this attribute is multivalued
	Multi utils.NullableBool `json:"multi,omitempty"`
	// Determine if the attribute is selected on future fetches
	Selected             utils.NullableBool `json:"selected,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DataSourceOutput DataSourceOutput

// NewDataSourceOutput instantiates a new DataSourceOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceOutput(key string) *DataSourceOutput {
	this := DataSourceOutput{}
	this.Key = key
	var multi bool = false
	this.Multi = *utils.NewNullableBool(&multi)
	var selected bool = true
	this.Selected = *utils.NewNullableBool(&selected)
	return &this
}

// NewDataSourceOutputWithDefaults instantiates a new DataSourceOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceOutputWithDefaults() *DataSourceOutput {
	this := DataSourceOutput{}
	var multi bool = false
	this.Multi = *utils.NewNullableBool(&multi)
	var selected bool = true
	this.Selected = *utils.NewNullableBool(&selected)
	return &this
}

// GetKey returns the Key field value
func (o *DataSourceOutput) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DataSourceOutput) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DataSourceOutput) SetKey(v string) {
	o.Key = v
}

// GetMulti returns the Multi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourceOutput) GetMulti() bool {
	if o == nil || utils.IsNil(o.Multi.Get()) {
		var ret bool
		return ret
	}
	return *o.Multi.Get()
}

// GetMultiOk returns a tuple with the Multi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourceOutput) GetMultiOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Multi.Get(), o.Multi.IsSet()
}

// HasMulti returns a boolean if a field has been set.
func (o *DataSourceOutput) HasMulti() bool {
	if o != nil && o.Multi.IsSet() {
		return true
	}

	return false
}

// SetMulti gets a reference to the given NullableBool and assigns it to the Multi field.
func (o *DataSourceOutput) SetMulti(v bool) {
	o.Multi.Set(&v)
}

// SetMultiNil sets the value for Multi to be an explicit nil
func (o *DataSourceOutput) SetMultiNil() {
	o.Multi.Set(nil)
}

// UnsetMulti ensures that no value is present for Multi, not even an explicit nil
func (o *DataSourceOutput) UnsetMulti() {
	o.Multi.Unset()
}

// GetSelected returns the Selected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourceOutput) GetSelected() bool {
	if o == nil || utils.IsNil(o.Selected.Get()) {
		var ret bool
		return ret
	}
	return *o.Selected.Get()
}

// GetSelectedOk returns a tuple with the Selected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourceOutput) GetSelectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Selected.Get(), o.Selected.IsSet()
}

// HasSelected returns a boolean if a field has been set.
func (o *DataSourceOutput) HasSelected() bool {
	if o != nil && o.Selected.IsSet() {
		return true
	}

	return false
}

// SetSelected gets a reference to the given NullableBool and assigns it to the Selected field.
func (o *DataSourceOutput) SetSelected(v bool) {
	o.Selected.Set(&v)
}

// SetSelectedNil sets the value for Selected to be an explicit nil
func (o *DataSourceOutput) SetSelectedNil() {
	o.Selected.Set(nil)
}

// UnsetSelected ensures that no value is present for Selected, not even an explicit nil
func (o *DataSourceOutput) UnsetSelected() {
	o.Selected.Unset()
}

func (o DataSourceOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourceOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if o.Multi.IsSet() {
		toSerialize["multi"] = o.Multi.Get()
	}
	if o.Selected.IsSet() {
		toSerialize["selected"] = o.Selected.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DataSourceOutput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
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

	varDataSourceOutput := _DataSourceOutput{}

	err = json.Unmarshal(data, &varDataSourceOutput)

	if err != nil {
		return err
	}

	*o = DataSourceOutput(varDataSourceOutput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "multi")
		delete(additionalProperties, "selected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDataSourceOutput struct {
	value *DataSourceOutput
	isSet bool
}

func (v NullableDataSourceOutput) Get() *DataSourceOutput {
	return v.value
}

func (v *NullableDataSourceOutput) Set(val *DataSourceOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceOutput(val *DataSourceOutput) *NullableDataSourceOutput {
	return &NullableDataSourceOutput{value: val, isSet: true}
}

func (v NullableDataSourceOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
