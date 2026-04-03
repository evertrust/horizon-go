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

// checks if the Chart type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Chart{}

// Chart struct for Chart
type Chart struct {
	// The colors of the chart
	Colors []string `json:"colors"`
	// The description of the chart
	Description utils.NullableString `json:"description,omitempty"`
	Direction   utils.NullableString `json:"direction,omitempty"`
	// The field that will be used to group data
	Fields []string `json:"fields"`
	// The height of the chart
	H utils.NullableInt64 `json:"h,omitempty"`
	// A condition to apply to the results of the aggregate. Only the aggregates results with more than 5 items in them can be kept for example
	Having NullableHaving `json:"having,omitempty"`
	// The index of the chart on the dashboard
	I utils.NullableString `json:"i,omitempty"`
	// The maximum number of results to display
	Limit utils.NullableInt64 `json:"limit,omitempty"`
	// The HCQL/HRQL query to build the chart from
	LocalQuery utils.NullableString `json:"localQuery,omitempty"`
	// Whether the logarithm scale is enabled or not
	Log bool `json:"log"`
	// How to sort the results in the chart (if applicable)
	SortOrder utils.NullableString `json:"sortOrder,omitempty"`
	// Title of the chart
	Title string `json:"title"`
	// The type of the chart
	Type string `json:"type"`
	// The width of the chart
	W utils.NullableInt64 `json:"w,omitempty"`
	// The horizontal position of the chart on the grid
	X utils.NullableInt64 `json:"x,omitempty"`
	// The vertical position of the chart on the grid
	Y                    utils.NullableInt64 `json:"y,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Chart Chart

// NewChart instantiates a new Chart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChart(colors []string, fields []string, log bool, title string, type_ string) *Chart {
	this := Chart{}
	this.Colors = colors
	this.Fields = fields
	this.Log = log
	this.Title = title
	this.Type = type_
	return &this
}

// NewChartWithDefaults instantiates a new Chart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartWithDefaults() *Chart {
	this := Chart{}
	return &this
}

// GetColors returns the Colors field value
func (o *Chart) GetColors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Colors
}

// GetColorsOk returns a tuple with the Colors field value
// and a boolean to check if the value has been set.
func (o *Chart) GetColorsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Colors, true
}

// SetColors sets field value
func (o *Chart) SetColors(v []string) {
	o.Colors = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chart) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chart) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Chart) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Chart) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Chart) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Chart) UnsetDescription() {
	o.Description.Unset()
}

// GetDirection returns the Direction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chart) GetDirection() string {
	if o == nil || utils.IsNil(o.Direction.Get()) {
		var ret string
		return ret
	}
	return *o.Direction.Get()
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chart) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Direction.Get(), o.Direction.IsSet()
}

// HasDirection returns a boolean if a field has been set.
func (o *Chart) HasDirection() bool {
	if o != nil && o.Direction.IsSet() {
		return true
	}

	return false
}

// SetDirection gets a reference to the given NullableString and assigns it to the Direction field.
func (o *Chart) SetDirection(v string) {
	o.Direction.Set(&v)
}

// SetDirectionNil sets the value for Direction to be an explicit nil
func (o *Chart) SetDirectionNil() {
	o.Direction.Set(nil)
}

// UnsetDirection ensures that no value is present for Direction, not even an explicit nil
func (o *Chart) UnsetDirection() {
	o.Direction.Unset()
}

// GetFields returns the Fields field value
func (o *Chart) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *Chart) GetFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *Chart) SetFields(v []string) {
	o.Fields = v
}

// GetH returns the H field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chart) GetH() int64 {
	if o == nil || utils.IsNil(o.H.Get()) {
		var ret int64
		return ret
	}
	return *o.H.Get()
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chart) GetHOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.H.Get(), o.H.IsSet()
}

// HasH returns a boolean if a field has been set.
func (o *Chart) HasH() bool {
	if o != nil && o.H.IsSet() {
		return true
	}

	return false
}

// SetH gets a reference to the given NullableInt64 and assigns it to the H field.
func (o *Chart) SetH(v int64) {
	o.H.Set(&v)
}

// SetHNil sets the value for H to be an explicit nil
func (o *Chart) SetHNil() {
	o.H.Set(nil)
}

// UnsetH ensures that no value is present for H, not even an explicit nil
func (o *Chart) UnsetH() {
	o.H.Unset()
}

// GetHaving returns the Having field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chart) GetHaving() Having {
	if o == nil || utils.IsNil(o.Having.Get()) {
		var ret Having
		return ret
	}
	return *o.Having.Get()
}

// GetHavingOk returns a tuple with the Having field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chart) GetHavingOk() (*Having, bool) {
	if o == nil {
		return nil, false
	}
	return o.Having.Get(), o.Having.IsSet()
}

// HasHaving returns a boolean if a field has been set.
func (o *Chart) HasHaving() bool {
	if o != nil && o.Having.IsSet() {
		return true
	}

	return false
}

// SetHaving gets a reference to the given NullableHaving and assigns it to the Having field.
func (o *Chart) SetHaving(v Having) {
	o.Having.Set(&v)
}

// SetHavingNil sets the value for Having to be an explicit nil
func (o *Chart) SetHavingNil() {
	o.Having.Set(nil)
}

// UnsetHaving ensures that no value is present for Having, not even an explicit nil
func (o *Chart) UnsetHaving() {
	o.Having.Unset()
}

// GetI returns the I field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chart) GetI() string {
	if o == nil || utils.IsNil(o.I.Get()) {
		var ret string
		return ret
	}
	return *o.I.Get()
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chart) GetIOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.I.Get(), o.I.IsSet()
}

// HasI returns a boolean if a field has been set.
func (o *Chart) HasI() bool {
	if o != nil && o.I.IsSet() {
		return true
	}

	return false
}

// SetI gets a reference to the given NullableString and assigns it to the I field.
func (o *Chart) SetI(v string) {
	o.I.Set(&v)
}

// SetINil sets the value for I to be an explicit nil
func (o *Chart) SetINil() {
	o.I.Set(nil)
}

// UnsetI ensures that no value is present for I, not even an explicit nil
func (o *Chart) UnsetI() {
	o.I.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chart) GetLimit() int64 {
	if o == nil || utils.IsNil(o.Limit.Get()) {
		var ret int64
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chart) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *Chart) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt64 and assigns it to the Limit field.
func (o *Chart) SetLimit(v int64) {
	o.Limit.Set(&v)
}

// SetLimitNil sets the value for Limit to be an explicit nil
func (o *Chart) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *Chart) UnsetLimit() {
	o.Limit.Unset()
}

// GetLocalQuery returns the LocalQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chart) GetLocalQuery() string {
	if o == nil || utils.IsNil(o.LocalQuery.Get()) {
		var ret string
		return ret
	}
	return *o.LocalQuery.Get()
}

// GetLocalQueryOk returns a tuple with the LocalQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chart) GetLocalQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalQuery.Get(), o.LocalQuery.IsSet()
}

// HasLocalQuery returns a boolean if a field has been set.
func (o *Chart) HasLocalQuery() bool {
	if o != nil && o.LocalQuery.IsSet() {
		return true
	}

	return false
}

// SetLocalQuery gets a reference to the given NullableString and assigns it to the LocalQuery field.
func (o *Chart) SetLocalQuery(v string) {
	o.LocalQuery.Set(&v)
}

// SetLocalQueryNil sets the value for LocalQuery to be an explicit nil
func (o *Chart) SetLocalQueryNil() {
	o.LocalQuery.Set(nil)
}

// UnsetLocalQuery ensures that no value is present for LocalQuery, not even an explicit nil
func (o *Chart) UnsetLocalQuery() {
	o.LocalQuery.Unset()
}

// GetLog returns the Log field value
func (o *Chart) GetLog() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Log
}

// GetLogOk returns a tuple with the Log field value
// and a boolean to check if the value has been set.
func (o *Chart) GetLogOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Log, true
}

// SetLog sets field value
func (o *Chart) SetLog(v bool) {
	o.Log = v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chart) GetSortOrder() string {
	if o == nil || utils.IsNil(o.SortOrder.Get()) {
		var ret string
		return ret
	}
	return *o.SortOrder.Get()
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chart) GetSortOrderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortOrder.Get(), o.SortOrder.IsSet()
}

// HasSortOrder returns a boolean if a field has been set.
func (o *Chart) HasSortOrder() bool {
	if o != nil && o.SortOrder.IsSet() {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given NullableString and assigns it to the SortOrder field.
func (o *Chart) SetSortOrder(v string) {
	o.SortOrder.Set(&v)
}

// SetSortOrderNil sets the value for SortOrder to be an explicit nil
func (o *Chart) SetSortOrderNil() {
	o.SortOrder.Set(nil)
}

// UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
func (o *Chart) UnsetSortOrder() {
	o.SortOrder.Unset()
}

// GetTitle returns the Title field value
func (o *Chart) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Chart) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Chart) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *Chart) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Chart) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Chart) SetType(v string) {
	o.Type = v
}

// GetW returns the W field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chart) GetW() int64 {
	if o == nil || utils.IsNil(o.W.Get()) {
		var ret int64
		return ret
	}
	return *o.W.Get()
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chart) GetWOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.W.Get(), o.W.IsSet()
}

// HasW returns a boolean if a field has been set.
func (o *Chart) HasW() bool {
	if o != nil && o.W.IsSet() {
		return true
	}

	return false
}

// SetW gets a reference to the given NullableInt64 and assigns it to the W field.
func (o *Chart) SetW(v int64) {
	o.W.Set(&v)
}

// SetWNil sets the value for W to be an explicit nil
func (o *Chart) SetWNil() {
	o.W.Set(nil)
}

// UnsetW ensures that no value is present for W, not even an explicit nil
func (o *Chart) UnsetW() {
	o.W.Unset()
}

// GetX returns the X field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chart) GetX() int64 {
	if o == nil || utils.IsNil(o.X.Get()) {
		var ret int64
		return ret
	}
	return *o.X.Get()
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chart) GetXOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.X.Get(), o.X.IsSet()
}

// HasX returns a boolean if a field has been set.
func (o *Chart) HasX() bool {
	if o != nil && o.X.IsSet() {
		return true
	}

	return false
}

// SetX gets a reference to the given NullableInt64 and assigns it to the X field.
func (o *Chart) SetX(v int64) {
	o.X.Set(&v)
}

// SetXNil sets the value for X to be an explicit nil
func (o *Chart) SetXNil() {
	o.X.Set(nil)
}

// UnsetX ensures that no value is present for X, not even an explicit nil
func (o *Chart) UnsetX() {
	o.X.Unset()
}

// GetY returns the Y field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Chart) GetY() int64 {
	if o == nil || utils.IsNil(o.Y.Get()) {
		var ret int64
		return ret
	}
	return *o.Y.Get()
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Chart) GetYOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Y.Get(), o.Y.IsSet()
}

// HasY returns a boolean if a field has been set.
func (o *Chart) HasY() bool {
	if o != nil && o.Y.IsSet() {
		return true
	}

	return false
}

// SetY gets a reference to the given NullableInt64 and assigns it to the Y field.
func (o *Chart) SetY(v int64) {
	o.Y.Set(&v)
}

// SetYNil sets the value for Y to be an explicit nil
func (o *Chart) SetYNil() {
	o.Y.Set(nil)
}

// UnsetY ensures that no value is present for Y, not even an explicit nil
func (o *Chart) UnsetY() {
	o.Y.Unset()
}

func (o Chart) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Chart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["colors"] = o.Colors
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Direction.IsSet() {
		toSerialize["direction"] = o.Direction.Get()
	}
	toSerialize["fields"] = o.Fields
	if o.H.IsSet() {
		toSerialize["h"] = o.H.Get()
	}
	if o.Having.IsSet() {
		toSerialize["having"] = o.Having.Get()
	}
	if o.I.IsSet() {
		toSerialize["i"] = o.I.Get()
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.LocalQuery.IsSet() {
		toSerialize["localQuery"] = o.LocalQuery.Get()
	}
	toSerialize["log"] = o.Log
	if o.SortOrder.IsSet() {
		toSerialize["sortOrder"] = o.SortOrder.Get()
	}
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type
	if o.W.IsSet() {
		toSerialize["w"] = o.W.Get()
	}
	if o.X.IsSet() {
		toSerialize["x"] = o.X.Get()
	}
	if o.Y.IsSet() {
		toSerialize["y"] = o.Y.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Chart) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"colors",
		"fields",
		"log",
		"title",
		"type",
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

	varChart := _Chart{}

	err = json.Unmarshal(data, &varChart)

	if err != nil {
		return err
	}

	*o = Chart(varChart)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "colors")
		delete(additionalProperties, "description")
		delete(additionalProperties, "direction")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "h")
		delete(additionalProperties, "having")
		delete(additionalProperties, "i")
		delete(additionalProperties, "limit")
		delete(additionalProperties, "localQuery")
		delete(additionalProperties, "log")
		delete(additionalProperties, "sortOrder")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "w")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChart struct {
	value *Chart
	isSet bool
}

func (v NullableChart) Get() *Chart {
	return v.value
}

func (v *NullableChart) Set(val *Chart) {
	v.value = val
	v.isSet = true
}

func (v NullableChart) IsSet() bool {
	return v.isSet
}

func (v *NullableChart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChart(val *Chart) *NullableChart {
	return &NullableChart{value: val, isSet: true}
}

func (v NullableChart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
