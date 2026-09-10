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

// checks if the DataSourceFlowTemplateEntryResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DataSourceFlowTemplateEntryResponse{}

// DataSourceFlowTemplateEntryResponse struct for DataSourceFlowTemplateEntryResponse
type DataSourceFlowTemplateEntryResponse struct {
	// Description of the datasource
	Description utils.NullableString `json:"description,omitempty"`
	// Display name of the datasource
	DisplayName []LocalizedString `json:"displayName,omitempty"`
	// List of inputs to use for this datasource
	Inputs []DataSourceInput `json:"inputs,omitempty"`
	// If true, the flow will stop with an error if this datasource does not return any result
	Mandatory *bool `json:"mandatory,omitempty"`
	// Name of the datasource
	Name *string `json:"name,omitempty"`
	// List of outputs for this datasource
	Outputs []DataSourceOutput `json:"outputs,omitempty"`
	// Stop the execution if this datasource's execution is successful
	StopOnSuccess        *bool           `json:"stopOnSuccess,omitempty"`
	Type                 *DataSourceType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DataSourceFlowTemplateEntryResponse DataSourceFlowTemplateEntryResponse

// NewDataSourceFlowTemplateEntryResponse instantiates a new DataSourceFlowTemplateEntryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceFlowTemplateEntryResponse() *DataSourceFlowTemplateEntryResponse {
	this := DataSourceFlowTemplateEntryResponse{}
	var mandatory bool = false
	this.Mandatory = &mandatory
	var stopOnSuccess bool = false
	this.StopOnSuccess = &stopOnSuccess
	return &this
}

// NewDataSourceFlowTemplateEntryResponseWithDefaults instantiates a new DataSourceFlowTemplateEntryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceFlowTemplateEntryResponseWithDefaults() *DataSourceFlowTemplateEntryResponse {
	this := DataSourceFlowTemplateEntryResponse{}
	var mandatory bool = false
	this.Mandatory = &mandatory
	var stopOnSuccess bool = false
	this.StopOnSuccess = &stopOnSuccess
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourceFlowTemplateEntryResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourceFlowTemplateEntryResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *DataSourceFlowTemplateEntryResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *DataSourceFlowTemplateEntryResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *DataSourceFlowTemplateEntryResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *DataSourceFlowTemplateEntryResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourceFlowTemplateEntryResponse) GetDisplayName() []LocalizedString {
	if o == nil {
		var ret []LocalizedString
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourceFlowTemplateEntryResponse) GetDisplayNameOk() ([]LocalizedString, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DataSourceFlowTemplateEntryResponse) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given []LocalizedString and assigns it to the DisplayName field.
func (o *DataSourceFlowTemplateEntryResponse) SetDisplayName(v []LocalizedString) {
	o.DisplayName = v
}

// GetInputs returns the Inputs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourceFlowTemplateEntryResponse) GetInputs() []DataSourceInput {
	if o == nil {
		var ret []DataSourceInput
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourceFlowTemplateEntryResponse) GetInputsOk() ([]DataSourceInput, bool) {
	if o == nil || utils.IsNil(o.Inputs) {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *DataSourceFlowTemplateEntryResponse) HasInputs() bool {
	if o != nil && !utils.IsNil(o.Inputs) {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []DataSourceInput and assigns it to the Inputs field.
func (o *DataSourceFlowTemplateEntryResponse) SetInputs(v []DataSourceInput) {
	o.Inputs = v
}

// GetMandatory returns the Mandatory field value if set, zero value otherwise.
func (o *DataSourceFlowTemplateEntryResponse) GetMandatory() bool {
	if o == nil || utils.IsNil(o.Mandatory) {
		var ret bool
		return ret
	}
	return *o.Mandatory
}

// GetMandatoryOk returns a tuple with the Mandatory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceFlowTemplateEntryResponse) GetMandatoryOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Mandatory) {
		return nil, false
	}
	return o.Mandatory, true
}

// HasMandatory returns a boolean if a field has been set.
func (o *DataSourceFlowTemplateEntryResponse) HasMandatory() bool {
	if o != nil && !utils.IsNil(o.Mandatory) {
		return true
	}

	return false
}

// SetMandatory gets a reference to the given bool and assigns it to the Mandatory field.
func (o *DataSourceFlowTemplateEntryResponse) SetMandatory(v bool) {
	o.Mandatory = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataSourceFlowTemplateEntryResponse) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceFlowTemplateEntryResponse) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataSourceFlowTemplateEntryResponse) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataSourceFlowTemplateEntryResponse) SetName(v string) {
	o.Name = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataSourceFlowTemplateEntryResponse) GetOutputs() []DataSourceOutput {
	if o == nil {
		var ret []DataSourceOutput
		return ret
	}
	return o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataSourceFlowTemplateEntryResponse) GetOutputsOk() ([]DataSourceOutput, bool) {
	if o == nil || utils.IsNil(o.Outputs) {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *DataSourceFlowTemplateEntryResponse) HasOutputs() bool {
	if o != nil && !utils.IsNil(o.Outputs) {
		return true
	}

	return false
}

// SetOutputs gets a reference to the given []DataSourceOutput and assigns it to the Outputs field.
func (o *DataSourceFlowTemplateEntryResponse) SetOutputs(v []DataSourceOutput) {
	o.Outputs = v
}

// GetStopOnSuccess returns the StopOnSuccess field value if set, zero value otherwise.
func (o *DataSourceFlowTemplateEntryResponse) GetStopOnSuccess() bool {
	if o == nil || utils.IsNil(o.StopOnSuccess) {
		var ret bool
		return ret
	}
	return *o.StopOnSuccess
}

// GetStopOnSuccessOk returns a tuple with the StopOnSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceFlowTemplateEntryResponse) GetStopOnSuccessOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.StopOnSuccess) {
		return nil, false
	}
	return o.StopOnSuccess, true
}

// HasStopOnSuccess returns a boolean if a field has been set.
func (o *DataSourceFlowTemplateEntryResponse) HasStopOnSuccess() bool {
	if o != nil && !utils.IsNil(o.StopOnSuccess) {
		return true
	}

	return false
}

// SetStopOnSuccess gets a reference to the given bool and assigns it to the StopOnSuccess field.
func (o *DataSourceFlowTemplateEntryResponse) SetStopOnSuccess(v bool) {
	o.StopOnSuccess = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DataSourceFlowTemplateEntryResponse) GetType() DataSourceType {
	if o == nil || utils.IsNil(o.Type) {
		var ret DataSourceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceFlowTemplateEntryResponse) GetTypeOk() (*DataSourceType, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DataSourceFlowTemplateEntryResponse) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DataSourceType and assigns it to the Type field.
func (o *DataSourceFlowTemplateEntryResponse) SetType(v DataSourceType) {
	o.Type = &v
}

func (o DataSourceFlowTemplateEntryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSourceFlowTemplateEntryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Inputs != nil {
		toSerialize["inputs"] = o.Inputs
	}
	if !utils.IsNil(o.Mandatory) {
		toSerialize["mandatory"] = o.Mandatory
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Outputs != nil {
		toSerialize["outputs"] = o.Outputs
	}
	if !utils.IsNil(o.StopOnSuccess) {
		toSerialize["stopOnSuccess"] = o.StopOnSuccess
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DataSourceFlowTemplateEntryResponse) UnmarshalJSON(data []byte) (err error) {
	varDataSourceFlowTemplateEntryResponse := _DataSourceFlowTemplateEntryResponse{}

	err = json.Unmarshal(data, &varDataSourceFlowTemplateEntryResponse)

	if err != nil {
		return err
	}

	*o = DataSourceFlowTemplateEntryResponse(varDataSourceFlowTemplateEntryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "inputs")
		delete(additionalProperties, "mandatory")
		delete(additionalProperties, "name")
		delete(additionalProperties, "outputs")
		delete(additionalProperties, "stopOnSuccess")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDataSourceFlowTemplateEntryResponse struct {
	value *DataSourceFlowTemplateEntryResponse
	isSet bool
}

func (v NullableDataSourceFlowTemplateEntryResponse) Get() *DataSourceFlowTemplateEntryResponse {
	return v.value
}

func (v *NullableDataSourceFlowTemplateEntryResponse) Set(val *DataSourceFlowTemplateEntryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceFlowTemplateEntryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceFlowTemplateEntryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceFlowTemplateEntryResponse(val *DataSourceFlowTemplateEntryResponse) *NullableDataSourceFlowTemplateEntryResponse {
	return &NullableDataSourceFlowTemplateEntryResponse{value: val, isSet: true}
}

func (v NullableDataSourceFlowTemplateEntryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceFlowTemplateEntryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
