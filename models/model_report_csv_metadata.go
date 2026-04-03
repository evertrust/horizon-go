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

// checks if the ReportCSVMetadata type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReportCSVMetadata{}

// ReportCSVMetadata struct for ReportCSVMetadata
type ReportCSVMetadata struct {
	// Indicates when the report CSV will be deleted from database
	RemoveAt int64 `json:"removeAt"`
	// Id of the report that generated this CSV
	ReportId string `json:"reportId"`
	// Name of the report that generated this CSV
	ReportName string `json:"reportName"`
	// The unique identifier of the report CSV. With that UUID, it is possible to download the CSV
	Uuid                 string `json:"uuid" validate:"regexp=^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"`
	AdditionalProperties map[string]interface{}
}

type _ReportCSVMetadata ReportCSVMetadata

// NewReportCSVMetadata instantiates a new ReportCSVMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportCSVMetadata(removeAt int64, reportId string, reportName string, uuid string) *ReportCSVMetadata {
	this := ReportCSVMetadata{}
	this.RemoveAt = removeAt
	this.ReportId = reportId
	this.ReportName = reportName
	this.Uuid = uuid
	return &this
}

// NewReportCSVMetadataWithDefaults instantiates a new ReportCSVMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportCSVMetadataWithDefaults() *ReportCSVMetadata {
	this := ReportCSVMetadata{}
	return &this
}

// GetRemoveAt returns the RemoveAt field value
func (o *ReportCSVMetadata) GetRemoveAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RemoveAt
}

// GetRemoveAtOk returns a tuple with the RemoveAt field value
// and a boolean to check if the value has been set.
func (o *ReportCSVMetadata) GetRemoveAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoveAt, true
}

// SetRemoveAt sets field value
func (o *ReportCSVMetadata) SetRemoveAt(v int64) {
	o.RemoveAt = v
}

// GetReportId returns the ReportId field value
func (o *ReportCSVMetadata) GetReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value
// and a boolean to check if the value has been set.
func (o *ReportCSVMetadata) GetReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportId, true
}

// SetReportId sets field value
func (o *ReportCSVMetadata) SetReportId(v string) {
	o.ReportId = v
}

// GetReportName returns the ReportName field value
func (o *ReportCSVMetadata) GetReportName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportName
}

// GetReportNameOk returns a tuple with the ReportName field value
// and a boolean to check if the value has been set.
func (o *ReportCSVMetadata) GetReportNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportName, true
}

// SetReportName sets field value
func (o *ReportCSVMetadata) SetReportName(v string) {
	o.ReportName = v
}

// GetUuid returns the Uuid field value
func (o *ReportCSVMetadata) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ReportCSVMetadata) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ReportCSVMetadata) SetUuid(v string) {
	o.Uuid = v
}

func (o ReportCSVMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportCSVMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["removeAt"] = o.RemoveAt
	toSerialize["reportId"] = o.ReportId
	toSerialize["reportName"] = o.ReportName
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReportCSVMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"removeAt",
		"reportId",
		"reportName",
		"uuid",
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

	varReportCSVMetadata := _ReportCSVMetadata{}

	err = json.Unmarshal(data, &varReportCSVMetadata)

	if err != nil {
		return err
	}

	*o = ReportCSVMetadata(varReportCSVMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "removeAt")
		delete(additionalProperties, "reportId")
		delete(additionalProperties, "reportName")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReportCSVMetadata struct {
	value *ReportCSVMetadata
	isSet bool
}

func (v NullableReportCSVMetadata) Get() *ReportCSVMetadata {
	return v.value
}

func (v *NullableReportCSVMetadata) Set(val *ReportCSVMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableReportCSVMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableReportCSVMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportCSVMetadata(val *ReportCSVMetadata) *NullableReportCSVMetadata {
	return &NullableReportCSVMetadata{value: val, isSet: true}
}

func (v NullableReportCSVMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportCSVMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
