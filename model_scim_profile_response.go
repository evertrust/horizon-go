/*
    Horizon API

    ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

    API version: 2.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package horizon

import (
	"encoding/json"
	"fmt"
)

// checks if the ScimProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimProfileResponse{}

// ScimProfileResponse struct for ScimProfileResponse
type ScimProfileResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// The name of the Scim profile
	Name string `json:"name"`
	// The description of the Scim profile
	Description NullableString `json:"description,omitempty"`
	// The mail type corresponds to the mail coming from the scim provider that must be synchronised in horizon. By default, the mail type is \"work\".
	MailType NullableString `json:"mailType,omitempty"`
	// The mapping used to synchronize user and group between the scim provider and Horizon.
	Mappings []ScimProfileResponseMappingsInner `json:"mappings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ScimProfileResponse ScimProfileResponse

// NewScimProfileResponse instantiates a new ScimProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimProfileResponse(id string, name string) *ScimProfileResponse {
	this := ScimProfileResponse{}
	this.Id = id
	this.Name = name
	var mailType string = "work"
	this.MailType = *NewNullableString(&mailType)
	return &this
}

// NewScimProfileResponseWithDefaults instantiates a new ScimProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimProfileResponseWithDefaults() *ScimProfileResponse {
	this := ScimProfileResponse{}
	var mailType string = "work"
	this.MailType = *NewNullableString(&mailType)
	return &this
}

// GetId returns the Id field value
func (o *ScimProfileResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScimProfileResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScimProfileResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ScimProfileResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ScimProfileResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ScimProfileResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScimProfileResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScimProfileResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ScimProfileResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ScimProfileResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ScimProfileResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ScimProfileResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetMailType returns the MailType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScimProfileResponse) GetMailType() string {
	if o == nil || IsNil(o.MailType.Get()) {
		var ret string
		return ret
	}
	return *o.MailType.Get()
}

// GetMailTypeOk returns a tuple with the MailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScimProfileResponse) GetMailTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MailType.Get(), o.MailType.IsSet()
}

// HasMailType returns a boolean if a field has been set.
func (o *ScimProfileResponse) HasMailType() bool {
	if o != nil && o.MailType.IsSet() {
		return true
	}

	return false
}

// SetMailType gets a reference to the given NullableString and assigns it to the MailType field.
func (o *ScimProfileResponse) SetMailType(v string) {
	o.MailType.Set(&v)
}
// SetMailTypeNil sets the value for MailType to be an explicit nil
func (o *ScimProfileResponse) SetMailTypeNil() {
	o.MailType.Set(nil)
}

// UnsetMailType ensures that no value is present for MailType, not even an explicit nil
func (o *ScimProfileResponse) UnsetMailType() {
	o.MailType.Unset()
}

// GetMappings returns the Mappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScimProfileResponse) GetMappings() []ScimProfileResponseMappingsInner {
	if o == nil {
		var ret []ScimProfileResponseMappingsInner
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScimProfileResponse) GetMappingsOk() ([]ScimProfileResponseMappingsInner, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *ScimProfileResponse) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []ScimProfileResponseMappingsInner and assigns it to the Mappings field.
func (o *ScimProfileResponse) SetMappings(v []ScimProfileResponseMappingsInner) {
	o.Mappings = v
}

func (o ScimProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.MailType.IsSet() {
		toSerialize["mailType"] = o.MailType.Get()
	}
	if o.Mappings != nil {
		toSerialize["mappings"] = o.Mappings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScimProfileResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varScimProfileResponse := _ScimProfileResponse{}

	err = json.Unmarshal(data, &varScimProfileResponse)

	if err != nil {
		return err
	}

	*o = ScimProfileResponse(varScimProfileResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mailType")
		delete(additionalProperties, "mappings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScimProfileResponse struct {
	value *ScimProfileResponse
	isSet bool
}

func (v NullableScimProfileResponse) Get() *ScimProfileResponse {
	return v.value
}

func (v *NullableScimProfileResponse) Set(val *ScimProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScimProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScimProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimProfileResponse(val *ScimProfileResponse) *NullableScimProfileResponse {
	return &NullableScimProfileResponse{value: val, isSet: true}
}

func (v NullableScimProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


