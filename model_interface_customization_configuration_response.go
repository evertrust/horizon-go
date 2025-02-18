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

// checks if the InterfaceCustomizationConfigurationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterfaceCustomizationConfigurationResponse{}

// InterfaceCustomizationConfigurationResponse struct for InterfaceCustomizationConfigurationResponse
type InterfaceCustomizationConfigurationResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	// The type of the configuration entry
	Type string `json:"type"`
	// A logo to display on the product, base64 encoded
	Logo NullableString `json:"logo,omitempty"`
	// The HTML color code for the left side of the banner gradient
	HeaderStart NullableString `json:"headerStart,omitempty"`
	// The HTML color code for the right side of the banner gradient
	HeaderEnd NullableString `json:"headerEnd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterfaceCustomizationConfigurationResponse InterfaceCustomizationConfigurationResponse

// NewInterfaceCustomizationConfigurationResponse instantiates a new InterfaceCustomizationConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceCustomizationConfigurationResponse(id string, type_ string) *InterfaceCustomizationConfigurationResponse {
	this := InterfaceCustomizationConfigurationResponse{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewInterfaceCustomizationConfigurationResponseWithDefaults instantiates a new InterfaceCustomizationConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceCustomizationConfigurationResponseWithDefaults() *InterfaceCustomizationConfigurationResponse {
	this := InterfaceCustomizationConfigurationResponse{}
	return &this
}

// GetId returns the Id field value
func (o *InterfaceCustomizationConfigurationResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InterfaceCustomizationConfigurationResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InterfaceCustomizationConfigurationResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *InterfaceCustomizationConfigurationResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InterfaceCustomizationConfigurationResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InterfaceCustomizationConfigurationResponse) SetType(v string) {
	o.Type = v
}

// GetLogo returns the Logo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceCustomizationConfigurationResponse) GetLogo() string {
	if o == nil || IsNil(o.Logo.Get()) {
		var ret string
		return ret
	}
	return *o.Logo.Get()
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceCustomizationConfigurationResponse) GetLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logo.Get(), o.Logo.IsSet()
}

// HasLogo returns a boolean if a field has been set.
func (o *InterfaceCustomizationConfigurationResponse) HasLogo() bool {
	if o != nil && o.Logo.IsSet() {
		return true
	}

	return false
}

// SetLogo gets a reference to the given NullableString and assigns it to the Logo field.
func (o *InterfaceCustomizationConfigurationResponse) SetLogo(v string) {
	o.Logo.Set(&v)
}
// SetLogoNil sets the value for Logo to be an explicit nil
func (o *InterfaceCustomizationConfigurationResponse) SetLogoNil() {
	o.Logo.Set(nil)
}

// UnsetLogo ensures that no value is present for Logo, not even an explicit nil
func (o *InterfaceCustomizationConfigurationResponse) UnsetLogo() {
	o.Logo.Unset()
}

// GetHeaderStart returns the HeaderStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceCustomizationConfigurationResponse) GetHeaderStart() string {
	if o == nil || IsNil(o.HeaderStart.Get()) {
		var ret string
		return ret
	}
	return *o.HeaderStart.Get()
}

// GetHeaderStartOk returns a tuple with the HeaderStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceCustomizationConfigurationResponse) GetHeaderStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderStart.Get(), o.HeaderStart.IsSet()
}

// HasHeaderStart returns a boolean if a field has been set.
func (o *InterfaceCustomizationConfigurationResponse) HasHeaderStart() bool {
	if o != nil && o.HeaderStart.IsSet() {
		return true
	}

	return false
}

// SetHeaderStart gets a reference to the given NullableString and assigns it to the HeaderStart field.
func (o *InterfaceCustomizationConfigurationResponse) SetHeaderStart(v string) {
	o.HeaderStart.Set(&v)
}
// SetHeaderStartNil sets the value for HeaderStart to be an explicit nil
func (o *InterfaceCustomizationConfigurationResponse) SetHeaderStartNil() {
	o.HeaderStart.Set(nil)
}

// UnsetHeaderStart ensures that no value is present for HeaderStart, not even an explicit nil
func (o *InterfaceCustomizationConfigurationResponse) UnsetHeaderStart() {
	o.HeaderStart.Unset()
}

// GetHeaderEnd returns the HeaderEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InterfaceCustomizationConfigurationResponse) GetHeaderEnd() string {
	if o == nil || IsNil(o.HeaderEnd.Get()) {
		var ret string
		return ret
	}
	return *o.HeaderEnd.Get()
}

// GetHeaderEndOk returns a tuple with the HeaderEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InterfaceCustomizationConfigurationResponse) GetHeaderEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderEnd.Get(), o.HeaderEnd.IsSet()
}

// HasHeaderEnd returns a boolean if a field has been set.
func (o *InterfaceCustomizationConfigurationResponse) HasHeaderEnd() bool {
	if o != nil && o.HeaderEnd.IsSet() {
		return true
	}

	return false
}

// SetHeaderEnd gets a reference to the given NullableString and assigns it to the HeaderEnd field.
func (o *InterfaceCustomizationConfigurationResponse) SetHeaderEnd(v string) {
	o.HeaderEnd.Set(&v)
}
// SetHeaderEndNil sets the value for HeaderEnd to be an explicit nil
func (o *InterfaceCustomizationConfigurationResponse) SetHeaderEndNil() {
	o.HeaderEnd.Set(nil)
}

// UnsetHeaderEnd ensures that no value is present for HeaderEnd, not even an explicit nil
func (o *InterfaceCustomizationConfigurationResponse) UnsetHeaderEnd() {
	o.HeaderEnd.Unset()
}

func (o InterfaceCustomizationConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceCustomizationConfigurationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["type"] = o.Type
	if o.Logo.IsSet() {
		toSerialize["logo"] = o.Logo.Get()
	}
	if o.HeaderStart.IsSet() {
		toSerialize["headerStart"] = o.HeaderStart.Get()
	}
	if o.HeaderEnd.IsSet() {
		toSerialize["headerEnd"] = o.HeaderEnd.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterfaceCustomizationConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"type",
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

	varInterfaceCustomizationConfigurationResponse := _InterfaceCustomizationConfigurationResponse{}

	err = json.Unmarshal(data, &varInterfaceCustomizationConfigurationResponse)

	if err != nil {
		return err
	}

	*o = InterfaceCustomizationConfigurationResponse(varInterfaceCustomizationConfigurationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "headerStart")
		delete(additionalProperties, "headerEnd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterfaceCustomizationConfigurationResponse struct {
	value *InterfaceCustomizationConfigurationResponse
	isSet bool
}

func (v NullableInterfaceCustomizationConfigurationResponse) Get() *InterfaceCustomizationConfigurationResponse {
	return v.value
}

func (v *NullableInterfaceCustomizationConfigurationResponse) Set(val *InterfaceCustomizationConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceCustomizationConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceCustomizationConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceCustomizationConfigurationResponse(val *InterfaceCustomizationConfigurationResponse) *NullableInterfaceCustomizationConfigurationResponse {
	return &NullableInterfaceCustomizationConfigurationResponse{value: val, isSet: true}
}

func (v NullableInterfaceCustomizationConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceCustomizationConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


