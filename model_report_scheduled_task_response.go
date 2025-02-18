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

// checks if the ReportScheduledTaskResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportScheduledTaskResponse{}

// ReportScheduledTaskResponse struct for ReportScheduledTaskResponse
type ReportScheduledTaskResponse struct {
	// Object internal ID
	Id string `json:"_id"`
	Type string `json:"type"`
	Cron string `json:"cron"`
	Host NullableString `json:"host,omitempty"`
	Status NullableString `json:"status,omitempty"`
	LastExecutionDate NullableInt64 `json:"lastExecutionDate,omitempty"`
	LastCompletionDate NullableInt64 `json:"lastCompletionDate,omitempty"`
	Detail NullableString `json:"detail,omitempty"`
	ExecutionId NullableString `json:"executionId,omitempty"`
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	FileName NullableString `json:"fileName,omitempty"`
	Recipients []ReportRecipient `json:"recipients"`
	From string `json:"from"`
	Title string `json:"title"`
	Body NullableString `json:"body,omitempty"`
	IsHtml bool `json:"isHtml"`
	HqlType string `json:"hqlType"`
	HqlQuery NullableString `json:"hqlQuery,omitempty"`
	HqlFields []string `json:"hqlFields,omitempty"`
	HqlSortedBy []SortElement `json:"hqlSortedBy,omitempty"`
	Description NullableString `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReportScheduledTaskResponse ReportScheduledTaskResponse

// NewReportScheduledTaskResponse instantiates a new ReportScheduledTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportScheduledTaskResponse(id string, type_ string, cron string, enabled bool, name string, recipients []ReportRecipient, from string, title string, isHtml bool, hqlType string) *ReportScheduledTaskResponse {
	this := ReportScheduledTaskResponse{}
	this.Id = id
	this.Type = type_
	this.Cron = cron
	this.Enabled = enabled
	this.Name = name
	this.Recipients = recipients
	this.From = from
	this.Title = title
	this.IsHtml = isHtml
	this.HqlType = hqlType
	return &this
}

// NewReportScheduledTaskResponseWithDefaults instantiates a new ReportScheduledTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportScheduledTaskResponseWithDefaults() *ReportScheduledTaskResponse {
	this := ReportScheduledTaskResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ReportScheduledTaskResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReportScheduledTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReportScheduledTaskResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ReportScheduledTaskResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReportScheduledTaskResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReportScheduledTaskResponse) SetType(v string) {
	o.Type = v
}

// GetCron returns the Cron field value
func (o *ReportScheduledTaskResponse) GetCron() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cron
}

// GetCronOk returns a tuple with the Cron field value
// and a boolean to check if the value has been set.
func (o *ReportScheduledTaskResponse) GetCronOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cron, true
}

// SetCron sets field value
func (o *ReportScheduledTaskResponse) SetCron(v string) {
	o.Cron = v
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetHost() string {
	if o == nil || IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *ReportScheduledTaskResponse) SetHost(v string) {
	o.Host.Set(&v)
}
// SetHostNil sets the value for Host to be an explicit nil
func (o *ReportScheduledTaskResponse) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *ReportScheduledTaskResponse) UnsetHost() {
	o.Host.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *ReportScheduledTaskResponse) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *ReportScheduledTaskResponse) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *ReportScheduledTaskResponse) UnsetStatus() {
	o.Status.Unset()
}

// GetLastExecutionDate returns the LastExecutionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetLastExecutionDate() int64 {
	if o == nil || IsNil(o.LastExecutionDate.Get()) {
		var ret int64
		return ret
	}
	return *o.LastExecutionDate.Get()
}

// GetLastExecutionDateOk returns a tuple with the LastExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetLastExecutionDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExecutionDate.Get(), o.LastExecutionDate.IsSet()
}

// HasLastExecutionDate returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasLastExecutionDate() bool {
	if o != nil && o.LastExecutionDate.IsSet() {
		return true
	}

	return false
}

// SetLastExecutionDate gets a reference to the given NullableInt64 and assigns it to the LastExecutionDate field.
func (o *ReportScheduledTaskResponse) SetLastExecutionDate(v int64) {
	o.LastExecutionDate.Set(&v)
}
// SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil
func (o *ReportScheduledTaskResponse) SetLastExecutionDateNil() {
	o.LastExecutionDate.Set(nil)
}

// UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
func (o *ReportScheduledTaskResponse) UnsetLastExecutionDate() {
	o.LastExecutionDate.Unset()
}

// GetLastCompletionDate returns the LastCompletionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetLastCompletionDate() int64 {
	if o == nil || IsNil(o.LastCompletionDate.Get()) {
		var ret int64
		return ret
	}
	return *o.LastCompletionDate.Get()
}

// GetLastCompletionDateOk returns a tuple with the LastCompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetLastCompletionDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastCompletionDate.Get(), o.LastCompletionDate.IsSet()
}

// HasLastCompletionDate returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasLastCompletionDate() bool {
	if o != nil && o.LastCompletionDate.IsSet() {
		return true
	}

	return false
}

// SetLastCompletionDate gets a reference to the given NullableInt64 and assigns it to the LastCompletionDate field.
func (o *ReportScheduledTaskResponse) SetLastCompletionDate(v int64) {
	o.LastCompletionDate.Set(&v)
}
// SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil
func (o *ReportScheduledTaskResponse) SetLastCompletionDateNil() {
	o.LastCompletionDate.Set(nil)
}

// UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
func (o *ReportScheduledTaskResponse) UnsetLastCompletionDate() {
	o.LastCompletionDate.Unset()
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetDetail() string {
	if o == nil || IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *ReportScheduledTaskResponse) SetDetail(v string) {
	o.Detail.Set(&v)
}
// SetDetailNil sets the value for Detail to be an explicit nil
func (o *ReportScheduledTaskResponse) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *ReportScheduledTaskResponse) UnsetDetail() {
	o.Detail.Unset()
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionId.Get()
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionId.Get(), o.ExecutionId.IsSet()
}

// HasExecutionId returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasExecutionId() bool {
	if o != nil && o.ExecutionId.IsSet() {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given NullableString and assigns it to the ExecutionId field.
func (o *ReportScheduledTaskResponse) SetExecutionId(v string) {
	o.ExecutionId.Set(&v)
}
// SetExecutionIdNil sets the value for ExecutionId to be an explicit nil
func (o *ReportScheduledTaskResponse) SetExecutionIdNil() {
	o.ExecutionId.Set(nil)
}

// UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
func (o *ReportScheduledTaskResponse) UnsetExecutionId() {
	o.ExecutionId.Unset()
}

// GetEnabled returns the Enabled field value
func (o *ReportScheduledTaskResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ReportScheduledTaskResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ReportScheduledTaskResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetName returns the Name field value
func (o *ReportScheduledTaskResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReportScheduledTaskResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReportScheduledTaskResponse) SetName(v string) {
	o.Name = v
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetFileName() string {
	if o == nil || IsNil(o.FileName.Get()) {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *ReportScheduledTaskResponse) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *ReportScheduledTaskResponse) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *ReportScheduledTaskResponse) UnsetFileName() {
	o.FileName.Unset()
}

// GetRecipients returns the Recipients field value
func (o *ReportScheduledTaskResponse) GetRecipients() []ReportRecipient {
	if o == nil {
		var ret []ReportRecipient
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ReportScheduledTaskResponse) GetRecipientsOk() ([]ReportRecipient, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *ReportScheduledTaskResponse) SetRecipients(v []ReportRecipient) {
	o.Recipients = v
}

// GetFrom returns the From field value
func (o *ReportScheduledTaskResponse) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ReportScheduledTaskResponse) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *ReportScheduledTaskResponse) SetFrom(v string) {
	o.From = v
}

// GetTitle returns the Title field value
func (o *ReportScheduledTaskResponse) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ReportScheduledTaskResponse) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ReportScheduledTaskResponse) SetTitle(v string) {
	o.Title = v
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *ReportScheduledTaskResponse) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *ReportScheduledTaskResponse) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *ReportScheduledTaskResponse) UnsetBody() {
	o.Body.Unset()
}

// GetIsHtml returns the IsHtml field value
func (o *ReportScheduledTaskResponse) GetIsHtml() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsHtml
}

// GetIsHtmlOk returns a tuple with the IsHtml field value
// and a boolean to check if the value has been set.
func (o *ReportScheduledTaskResponse) GetIsHtmlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsHtml, true
}

// SetIsHtml sets field value
func (o *ReportScheduledTaskResponse) SetIsHtml(v bool) {
	o.IsHtml = v
}

// GetHqlType returns the HqlType field value
func (o *ReportScheduledTaskResponse) GetHqlType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HqlType
}

// GetHqlTypeOk returns a tuple with the HqlType field value
// and a boolean to check if the value has been set.
func (o *ReportScheduledTaskResponse) GetHqlTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HqlType, true
}

// SetHqlType sets field value
func (o *ReportScheduledTaskResponse) SetHqlType(v string) {
	o.HqlType = v
}

// GetHqlQuery returns the HqlQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetHqlQuery() string {
	if o == nil || IsNil(o.HqlQuery.Get()) {
		var ret string
		return ret
	}
	return *o.HqlQuery.Get()
}

// GetHqlQueryOk returns a tuple with the HqlQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetHqlQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HqlQuery.Get(), o.HqlQuery.IsSet()
}

// HasHqlQuery returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasHqlQuery() bool {
	if o != nil && o.HqlQuery.IsSet() {
		return true
	}

	return false
}

// SetHqlQuery gets a reference to the given NullableString and assigns it to the HqlQuery field.
func (o *ReportScheduledTaskResponse) SetHqlQuery(v string) {
	o.HqlQuery.Set(&v)
}
// SetHqlQueryNil sets the value for HqlQuery to be an explicit nil
func (o *ReportScheduledTaskResponse) SetHqlQueryNil() {
	o.HqlQuery.Set(nil)
}

// UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
func (o *ReportScheduledTaskResponse) UnsetHqlQuery() {
	o.HqlQuery.Unset()
}

// GetHqlFields returns the HqlFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetHqlFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.HqlFields
}

// GetHqlFieldsOk returns a tuple with the HqlFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetHqlFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.HqlFields) {
		return nil, false
	}
	return o.HqlFields, true
}

// HasHqlFields returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasHqlFields() bool {
	if o != nil && !IsNil(o.HqlFields) {
		return true
	}

	return false
}

// SetHqlFields gets a reference to the given []string and assigns it to the HqlFields field.
func (o *ReportScheduledTaskResponse) SetHqlFields(v []string) {
	o.HqlFields = v
}

// GetHqlSortedBy returns the HqlSortedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetHqlSortedBy() []SortElement {
	if o == nil {
		var ret []SortElement
		return ret
	}
	return o.HqlSortedBy
}

// GetHqlSortedByOk returns a tuple with the HqlSortedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetHqlSortedByOk() ([]SortElement, bool) {
	if o == nil || IsNil(o.HqlSortedBy) {
		return nil, false
	}
	return o.HqlSortedBy, true
}

// HasHqlSortedBy returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasHqlSortedBy() bool {
	if o != nil && !IsNil(o.HqlSortedBy) {
		return true
	}

	return false
}

// SetHqlSortedBy gets a reference to the given []SortElement and assigns it to the HqlSortedBy field.
func (o *ReportScheduledTaskResponse) SetHqlSortedBy(v []SortElement) {
	o.HqlSortedBy = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduledTaskResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportScheduledTaskResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ReportScheduledTaskResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ReportScheduledTaskResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ReportScheduledTaskResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ReportScheduledTaskResponse) UnsetDescription() {
	o.Description.Unset()
}

func (o ReportScheduledTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportScheduledTaskResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["cron"] = o.Cron
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.LastExecutionDate.IsSet() {
		toSerialize["lastExecutionDate"] = o.LastExecutionDate.Get()
	}
	if o.LastCompletionDate.IsSet() {
		toSerialize["lastCompletionDate"] = o.LastCompletionDate.Get()
	}
	if o.Detail.IsSet() {
		toSerialize["detail"] = o.Detail.Get()
	}
	if o.ExecutionId.IsSet() {
		toSerialize["executionId"] = o.ExecutionId.Get()
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["name"] = o.Name
	if o.FileName.IsSet() {
		toSerialize["fileName"] = o.FileName.Get()
	}
	toSerialize["recipients"] = o.Recipients
	toSerialize["from"] = o.From
	toSerialize["title"] = o.Title
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	toSerialize["isHtml"] = o.IsHtml
	toSerialize["hqlType"] = o.HqlType
	if o.HqlQuery.IsSet() {
		toSerialize["hqlQuery"] = o.HqlQuery.Get()
	}
	if o.HqlFields != nil {
		toSerialize["hqlFields"] = o.HqlFields
	}
	if o.HqlSortedBy != nil {
		toSerialize["hqlSortedBy"] = o.HqlSortedBy
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReportScheduledTaskResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_id",
		"type",
		"cron",
		"enabled",
		"name",
		"recipients",
		"from",
		"title",
		"isHtml",
		"hqlType",
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

	varReportScheduledTaskResponse := _ReportScheduledTaskResponse{}

	err = json.Unmarshal(data, &varReportScheduledTaskResponse)

	if err != nil {
		return err
	}

	*o = ReportScheduledTaskResponse(varReportScheduledTaskResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "cron")
		delete(additionalProperties, "host")
		delete(additionalProperties, "status")
		delete(additionalProperties, "lastExecutionDate")
		delete(additionalProperties, "lastCompletionDate")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "executionId")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "name")
		delete(additionalProperties, "fileName")
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "from")
		delete(additionalProperties, "title")
		delete(additionalProperties, "body")
		delete(additionalProperties, "isHtml")
		delete(additionalProperties, "hqlType")
		delete(additionalProperties, "hqlQuery")
		delete(additionalProperties, "hqlFields")
		delete(additionalProperties, "hqlSortedBy")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReportScheduledTaskResponse struct {
	value *ReportScheduledTaskResponse
	isSet bool
}

func (v NullableReportScheduledTaskResponse) Get() *ReportScheduledTaskResponse {
	return v.value
}

func (v *NullableReportScheduledTaskResponse) Set(val *ReportScheduledTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReportScheduledTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReportScheduledTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportScheduledTaskResponse(val *ReportScheduledTaskResponse) *NullableReportScheduledTaskResponse {
	return &NullableReportScheduledTaskResponse{value: val, isSet: true}
}

func (v NullableReportScheduledTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportScheduledTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


