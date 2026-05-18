/*
   Horizon API

   ## Authentication  Most of the API calls that Horizon uses require you to be authenticated to the API. The first authentication can either be done through the use of an X509 certificate or using credentials of a local account, but every single API call afterward will need to bear the authentication information nonetheless. Regardless of the chosen authentication method, the authorization used must have sufficient permissions to perform the desired operation.  ### Authenticating using API-ID and API-KEY  This method of authentication requires you to send your Horizon local account credentials as HTTP headers. To check whether the credentials are correct, you can perform a *GET* request on `/api/v1/security/principals/self` and check for the response status : ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -H \"X-API-ID: administrator\" -H \"X-API-KEY: horizon\" -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Authenticating using an X509 certificate  This method of authentication requires to have a created authorization based on an X509 certificate that has the clientAuth EKU. It also requires you to have imported the CA that issued this certificate in Horizon and turning on the \"Trusted for client authentication\" switch on that CA. You must then present the certificate on the request you are performing.  To check for the authentication, you can perform a *GET* request on `/api/v1/security/principals/self` :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --cert horizon-login-dev-guide.pem --key horizon-login-dev-guide.key -H \"Accept: application/json\" ```  Possible responses are:  | HTTP Response code | Additional information                                                   | |--------------------|--------------------------------------------------------------------------| | 200                | The login information were correct                                       | | 401                | Authentication error, please refer to the response body for more details |  ### Handling next authentications using the Play Session  Once the first authentication is done, the API generates a cookie called \"PLAY_SESSION\". This cookie holds the authentication information that was used to make the first login (using either previously mentioned method). To save its value for later use, just append the _-c cookies.txt_ to either of the previous curl requests. Instead of using the credentials as headers or passing the certificate at each API call, you can use the cookie :  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self -b cookies.txt -H \"Accept: application/json\" ```  ### Handling CSRF Token    Our api are used by a frontend and require a CSRF protection. A CSRF token validation is needed when all of the following are true:  - The request method is not GET, HEAD or OPTIONS. - The request has one or more Cookie or Authorization headers.  Receiving the following response with valid credentials can mean that your request has failed the CSRF token validation:  ```json {     \"error\": \"SEC-AUTH-002\",     \"message\": \"Invalid credentials or principal does not exist\",     \"title\": \"Invalid credentials or principal does not exist\",     \"status\": 401 } ```  To avoid the CSRF token validation in api usage: - Authentication using API-ID and API-KEY headers should be prioritized as http basic authentication results in the creation of an Authorization header.  - Avoid the use of cookies as api usage does not require them.  If you cannot avoid those cases, the following procedure explains how to handle the CSRF token validation.   First you will have to retrieve a valid cookie CSRF token from the server.  ```shell  $ curl https://horizon.evertrust.fr/api/v1/security/principals/self --header 'X-API-ID:administrator' --header 'X-API-KEY:horizon' -c cookies.txt ```  Once done the file `cookies.txt` should have two entries: - A play session  - A CSRF token:  ```text localhost FALSE / FALSE 0 csrf-token 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d localhost FALSE / FALSE 1708942383 PLAY_SESSION eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7ImlkZW50aWZpZXIiOiJhZG1pbmlzdHJhdG9yIiwibmFtZSI6Ikhvcml6b24gQWRtaW5pc3RyYXRvciIsImlkcFR5cGUiOiJMb2NhbCIsImlkcE5hbWUiOiJsb2NhbCJ9LCJleHAiOjE3MDg5NDIzODMsIm5iZiI6MTcwODk0MTQ4MywiaWF0IjoxNzA4OTQxNDgzfQ.79xRjdGhaVv_5mM8bpkLgcL78QCEWu08zgthP_dt9Pc ```  To successfully authenticate to the server, both the csrf-token cookie and a `csrf-token` header containing the cookie content should be defined.  Sending a POST request using cookies without the `csrf-token` header will result in the forbidden html page:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'Content-Type: application/json' \\ -b cookies.txt \\ --data '{     \"name\": \"NEW_LABEL\",     \"displayName\" : [],     \"description\": [] }' ```  A valid authentication also copies the content in the `csrf-token` header:  ```shell curl --location 'localhost:9000/api/v1/certificate/labels' \\ --header 'X-API-ID: administrator' \\ --header 'X-API-KEY: evertrust' \\ --header 'csrf-token: 456aa18162e8736047dbd878617283aa361cd83e-1708941483170-da503a15304a666a96748f5d' \\ --header 'Content-Type: application/json' \\ --data '{     \"name\": \"NEW_LABEL\",     \"regex\": null,     \"displayName\" : [],     \"description\": [] }' ```

   API version: 2.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"

	"github.com/evertrust/horizon-go/v2/utils"
)

// checks if the AttachmentReportScheduledTask type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AttachmentReportScheduledTask{}

// AttachmentReportScheduledTask struct for AttachmentReportScheduledTask
type AttachmentReportScheduledTask struct {
	// Should the report be compressed using GZ
	CompressCsv *bool `json:"compressCsv,omitempty"`
	// Send an email with the CSV as an attachment
	ReportType           string               `json:"reportType"`
	Body                 utils.NullableString `json:"body,omitempty"`
	FileName             utils.NullableString `json:"fileName,omitempty"`
	From                 string               `json:"from"`
	HqlFields            []string             `json:"hqlFields,omitempty"`
	HqlQuery             utils.NullableString `json:"hqlQuery,omitempty"`
	HqlSortedBy          []SortElement        `json:"hqlSortedBy,omitempty"`
	HqlType              string               `json:"hqlType"`
	IsHtml               bool                 `json:"isHtml"`
	Recipients           []ReportRecipient    `json:"recipients"`
	Title                string               `json:"title"`
	Type                 string               `json:"type"`
	Cron                 string               `json:"cron"`
	Description          *string              `json:"description,omitempty"`
	Detail               utils.NullableString `json:"detail,omitempty"`
	Enabled              bool                 `json:"enabled"`
	ExecutionId          utils.NullableString `json:"executionId,omitempty"`
	Host                 utils.NullableString `json:"host,omitempty"`
	LastCompletionDate   utils.NullableInt64  `json:"lastCompletionDate,omitempty"`
	LastExecutionDate    utils.NullableInt64  `json:"lastExecutionDate,omitempty"`
	Name                 string               `json:"name"`
	Status               utils.NullableString `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AttachmentReportScheduledTask AttachmentReportScheduledTask

// NewAttachmentReportScheduledTask instantiates a new AttachmentReportScheduledTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentReportScheduledTask(reportType string, from string, hqlType string, isHtml bool, recipients []ReportRecipient, title string, type_ string, cron string, enabled bool, name string) *AttachmentReportScheduledTask {
	this := AttachmentReportScheduledTask{}
	this.From = from
	this.HqlType = hqlType
	this.IsHtml = isHtml
	this.Recipients = recipients
	this.ReportType = reportType
	this.Title = title
	this.Type = type_
	this.Cron = cron
	this.Enabled = enabled
	this.Name = name
	return &this
}

// NewAttachmentReportScheduledTaskWithDefaults instantiates a new AttachmentReportScheduledTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentReportScheduledTaskWithDefaults() *AttachmentReportScheduledTask {
	this := AttachmentReportScheduledTask{}
	return &this
}

// GetCompressCsv returns the CompressCsv field value if set, zero value otherwise.
func (o *AttachmentReportScheduledTask) GetCompressCsv() bool {
	if o == nil || utils.IsNil(o.CompressCsv) {
		var ret bool
		return ret
	}
	return *o.CompressCsv
}

// GetCompressCsvOk returns a tuple with the CompressCsv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetCompressCsvOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.CompressCsv) {
		return nil, false
	}
	return o.CompressCsv, true
}

// HasCompressCsv returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasCompressCsv() bool {
	if o != nil && !utils.IsNil(o.CompressCsv) {
		return true
	}

	return false
}

// SetCompressCsv gets a reference to the given bool and assigns it to the CompressCsv field.
func (o *AttachmentReportScheduledTask) SetCompressCsv(v bool) {
	o.CompressCsv = &v
}

// GetReportType returns the ReportType field value
func (o *AttachmentReportScheduledTask) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetReportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value
func (o *AttachmentReportScheduledTask) SetReportType(v string) {
	o.ReportType = v
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentReportScheduledTask) GetBody() string {
	if o == nil || utils.IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentReportScheduledTask) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *AttachmentReportScheduledTask) SetBody(v string) {
	o.Body.Set(&v)
}

// SetBodyNil sets the value for Body to be an explicit nil
func (o *AttachmentReportScheduledTask) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *AttachmentReportScheduledTask) UnsetBody() {
	o.Body.Unset()
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentReportScheduledTask) GetFileName() string {
	if o == nil || utils.IsNil(o.FileName.Get()) {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentReportScheduledTask) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *AttachmentReportScheduledTask) SetFileName(v string) {
	o.FileName.Set(&v)
}

// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *AttachmentReportScheduledTask) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *AttachmentReportScheduledTask) UnsetFileName() {
	o.FileName.Unset()
}

// GetFrom returns the From field value
func (o *AttachmentReportScheduledTask) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *AttachmentReportScheduledTask) SetFrom(v string) {
	o.From = v
}

// GetHqlFields returns the HqlFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentReportScheduledTask) GetHqlFields() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.HqlFields
}

// GetHqlFieldsOk returns a tuple with the HqlFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentReportScheduledTask) GetHqlFieldsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.HqlFields) {
		return nil, false
	}
	return o.HqlFields, true
}

// HasHqlFields returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasHqlFields() bool {
	if o != nil && !utils.IsNil(o.HqlFields) {
		return true
	}

	return false
}

// SetHqlFields gets a reference to the given []string and assigns it to the HqlFields field.
func (o *AttachmentReportScheduledTask) SetHqlFields(v []string) {
	o.HqlFields = v
}

// GetHqlQuery returns the HqlQuery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentReportScheduledTask) GetHqlQuery() string {
	if o == nil || utils.IsNil(o.HqlQuery.Get()) {
		var ret string
		return ret
	}
	return *o.HqlQuery.Get()
}

// GetHqlQueryOk returns a tuple with the HqlQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentReportScheduledTask) GetHqlQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HqlQuery.Get(), o.HqlQuery.IsSet()
}

// HasHqlQuery returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasHqlQuery() bool {
	if o != nil && o.HqlQuery.IsSet() {
		return true
	}

	return false
}

// SetHqlQuery gets a reference to the given NullableString and assigns it to the HqlQuery field.
func (o *AttachmentReportScheduledTask) SetHqlQuery(v string) {
	o.HqlQuery.Set(&v)
}

// SetHqlQueryNil sets the value for HqlQuery to be an explicit nil
func (o *AttachmentReportScheduledTask) SetHqlQueryNil() {
	o.HqlQuery.Set(nil)
}

// UnsetHqlQuery ensures that no value is present for HqlQuery, not even an explicit nil
func (o *AttachmentReportScheduledTask) UnsetHqlQuery() {
	o.HqlQuery.Unset()
}

// GetHqlSortedBy returns the HqlSortedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentReportScheduledTask) GetHqlSortedBy() []SortElement {
	if o == nil {
		var ret []SortElement
		return ret
	}
	return o.HqlSortedBy
}

// GetHqlSortedByOk returns a tuple with the HqlSortedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentReportScheduledTask) GetHqlSortedByOk() ([]SortElement, bool) {
	if o == nil || utils.IsNil(o.HqlSortedBy) {
		return nil, false
	}
	return o.HqlSortedBy, true
}

// HasHqlSortedBy returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasHqlSortedBy() bool {
	if o != nil && !utils.IsNil(o.HqlSortedBy) {
		return true
	}

	return false
}

// SetHqlSortedBy gets a reference to the given []SortElement and assigns it to the HqlSortedBy field.
func (o *AttachmentReportScheduledTask) SetHqlSortedBy(v []SortElement) {
	o.HqlSortedBy = v
}

// GetHqlType returns the HqlType field value
func (o *AttachmentReportScheduledTask) GetHqlType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HqlType
}

// GetHqlTypeOk returns a tuple with the HqlType field value
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetHqlTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HqlType, true
}

// SetHqlType sets field value
func (o *AttachmentReportScheduledTask) SetHqlType(v string) {
	o.HqlType = v
}

// GetIsHtml returns the IsHtml field value
func (o *AttachmentReportScheduledTask) GetIsHtml() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsHtml
}

// GetIsHtmlOk returns a tuple with the IsHtml field value
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetIsHtmlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsHtml, true
}

// SetIsHtml sets field value
func (o *AttachmentReportScheduledTask) SetIsHtml(v bool) {
	o.IsHtml = v
}

// GetRecipients returns the Recipients field value
func (o *AttachmentReportScheduledTask) GetRecipients() []ReportRecipient {
	if o == nil {
		var ret []ReportRecipient
		return ret
	}

	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetRecipientsOk() ([]ReportRecipient, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipients, true
}

// SetRecipients sets field value
func (o *AttachmentReportScheduledTask) SetRecipients(v []ReportRecipient) {
	o.Recipients = v
}

// GetTitle returns the Title field value
func (o *AttachmentReportScheduledTask) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AttachmentReportScheduledTask) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *AttachmentReportScheduledTask) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AttachmentReportScheduledTask) SetType(v string) {
	o.Type = v
}

// GetCron returns the Cron field value
func (o *AttachmentReportScheduledTask) GetCron() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cron
}

// GetCronOk returns a tuple with the Cron field value
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetCronOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cron, true
}

// SetCron sets field value
func (o *AttachmentReportScheduledTask) SetCron(v string) {
	o.Cron = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AttachmentReportScheduledTask) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AttachmentReportScheduledTask) SetDescription(v string) {
	o.Description = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentReportScheduledTask) GetDetail() string {
	if o == nil || utils.IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentReportScheduledTask) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *AttachmentReportScheduledTask) SetDetail(v string) {
	o.Detail.Set(&v)
}

// SetDetailNil sets the value for Detail to be an explicit nil
func (o *AttachmentReportScheduledTask) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *AttachmentReportScheduledTask) UnsetDetail() {
	o.Detail.Unset()
}

// GetEnabled returns the Enabled field value
func (o *AttachmentReportScheduledTask) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AttachmentReportScheduledTask) SetEnabled(v bool) {
	o.Enabled = v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentReportScheduledTask) GetExecutionId() string {
	if o == nil || utils.IsNil(o.ExecutionId.Get()) {
		var ret string
		return ret
	}
	return *o.ExecutionId.Get()
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentReportScheduledTask) GetExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecutionId.Get(), o.ExecutionId.IsSet()
}

// HasExecutionId returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasExecutionId() bool {
	if o != nil && o.ExecutionId.IsSet() {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given NullableString and assigns it to the ExecutionId field.
func (o *AttachmentReportScheduledTask) SetExecutionId(v string) {
	o.ExecutionId.Set(&v)
}

// SetExecutionIdNil sets the value for ExecutionId to be an explicit nil
func (o *AttachmentReportScheduledTask) SetExecutionIdNil() {
	o.ExecutionId.Set(nil)
}

// UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
func (o *AttachmentReportScheduledTask) UnsetExecutionId() {
	o.ExecutionId.Unset()
}

// GetHost returns the Host field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentReportScheduledTask) GetHost() string {
	if o == nil || utils.IsNil(o.Host.Get()) {
		var ret string
		return ret
	}
	return *o.Host.Get()
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentReportScheduledTask) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Host.Get(), o.Host.IsSet()
}

// HasHost returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasHost() bool {
	if o != nil && o.Host.IsSet() {
		return true
	}

	return false
}

// SetHost gets a reference to the given NullableString and assigns it to the Host field.
func (o *AttachmentReportScheduledTask) SetHost(v string) {
	o.Host.Set(&v)
}

// SetHostNil sets the value for Host to be an explicit nil
func (o *AttachmentReportScheduledTask) SetHostNil() {
	o.Host.Set(nil)
}

// UnsetHost ensures that no value is present for Host, not even an explicit nil
func (o *AttachmentReportScheduledTask) UnsetHost() {
	o.Host.Unset()
}

// GetLastCompletionDate returns the LastCompletionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentReportScheduledTask) GetLastCompletionDate() int64 {
	if o == nil || utils.IsNil(o.LastCompletionDate.Get()) {
		var ret int64
		return ret
	}
	return *o.LastCompletionDate.Get()
}

// GetLastCompletionDateOk returns a tuple with the LastCompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentReportScheduledTask) GetLastCompletionDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastCompletionDate.Get(), o.LastCompletionDate.IsSet()
}

// HasLastCompletionDate returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasLastCompletionDate() bool {
	if o != nil && o.LastCompletionDate.IsSet() {
		return true
	}

	return false
}

// SetLastCompletionDate gets a reference to the given NullableInt64 and assigns it to the LastCompletionDate field.
func (o *AttachmentReportScheduledTask) SetLastCompletionDate(v int64) {
	o.LastCompletionDate.Set(&v)
}

// SetLastCompletionDateNil sets the value for LastCompletionDate to be an explicit nil
func (o *AttachmentReportScheduledTask) SetLastCompletionDateNil() {
	o.LastCompletionDate.Set(nil)
}

// UnsetLastCompletionDate ensures that no value is present for LastCompletionDate, not even an explicit nil
func (o *AttachmentReportScheduledTask) UnsetLastCompletionDate() {
	o.LastCompletionDate.Unset()
}

// GetLastExecutionDate returns the LastExecutionDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentReportScheduledTask) GetLastExecutionDate() int64 {
	if o == nil || utils.IsNil(o.LastExecutionDate.Get()) {
		var ret int64
		return ret
	}
	return *o.LastExecutionDate.Get()
}

// GetLastExecutionDateOk returns a tuple with the LastExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentReportScheduledTask) GetLastExecutionDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastExecutionDate.Get(), o.LastExecutionDate.IsSet()
}

// HasLastExecutionDate returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasLastExecutionDate() bool {
	if o != nil && o.LastExecutionDate.IsSet() {
		return true
	}

	return false
}

// SetLastExecutionDate gets a reference to the given NullableInt64 and assigns it to the LastExecutionDate field.
func (o *AttachmentReportScheduledTask) SetLastExecutionDate(v int64) {
	o.LastExecutionDate.Set(&v)
}

// SetLastExecutionDateNil sets the value for LastExecutionDate to be an explicit nil
func (o *AttachmentReportScheduledTask) SetLastExecutionDateNil() {
	o.LastExecutionDate.Set(nil)
}

// UnsetLastExecutionDate ensures that no value is present for LastExecutionDate, not even an explicit nil
func (o *AttachmentReportScheduledTask) UnsetLastExecutionDate() {
	o.LastExecutionDate.Unset()
}

// GetName returns the Name field value
func (o *AttachmentReportScheduledTask) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AttachmentReportScheduledTask) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AttachmentReportScheduledTask) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentReportScheduledTask) GetStatus() string {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentReportScheduledTask) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *AttachmentReportScheduledTask) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *AttachmentReportScheduledTask) SetStatus(v string) {
	o.Status.Set(&v)
}

// SetStatusNil sets the value for Status to be an explicit nil
func (o *AttachmentReportScheduledTask) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *AttachmentReportScheduledTask) UnsetStatus() {
	o.Status.Unset()
}

func (o AttachmentReportScheduledTask) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentReportScheduledTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CompressCsv) {
		toSerialize["compressCsv"] = o.CompressCsv
	}
	toSerialize["reportType"] = o.ReportType
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.FileName.IsSet() {
		toSerialize["fileName"] = o.FileName.Get()
	}
	toSerialize["from"] = o.From
	if o.HqlFields != nil {
		toSerialize["hqlFields"] = o.HqlFields
	}
	if o.HqlQuery.IsSet() {
		toSerialize["hqlQuery"] = o.HqlQuery.Get()
	}
	if o.HqlSortedBy != nil {
		toSerialize["hqlSortedBy"] = o.HqlSortedBy
	}
	toSerialize["hqlType"] = o.HqlType
	toSerialize["isHtml"] = o.IsHtml
	toSerialize["recipients"] = o.Recipients
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type
	toSerialize["cron"] = o.Cron
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Detail.IsSet() {
		toSerialize["detail"] = o.Detail.Get()
	}
	toSerialize["enabled"] = o.Enabled
	if o.ExecutionId.IsSet() {
		toSerialize["executionId"] = o.ExecutionId.Get()
	}
	if o.Host.IsSet() {
		toSerialize["host"] = o.Host.Get()
	}
	if o.LastCompletionDate.IsSet() {
		toSerialize["lastCompletionDate"] = o.LastCompletionDate.Get()
	}
	if o.LastExecutionDate.IsSet() {
		toSerialize["lastExecutionDate"] = o.LastExecutionDate.Get()
	}
	toSerialize["name"] = o.Name
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AttachmentReportScheduledTask) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reportType",
		"from",
		"hqlType",
		"isHtml",
		"recipients",
		"title",
		"type",
		"cron",
		"enabled",
		"name",
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

	varAttachmentReportScheduledTask := _AttachmentReportScheduledTask{}

	err = json.Unmarshal(data, &varAttachmentReportScheduledTask)

	if err != nil {
		return err
	}

	*o = AttachmentReportScheduledTask(varAttachmentReportScheduledTask)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "compressCsv")
		delete(additionalProperties, "reportType")
		delete(additionalProperties, "body")
		delete(additionalProperties, "fileName")
		delete(additionalProperties, "from")
		delete(additionalProperties, "hqlFields")
		delete(additionalProperties, "hqlQuery")
		delete(additionalProperties, "hqlSortedBy")
		delete(additionalProperties, "hqlType")
		delete(additionalProperties, "isHtml")
		delete(additionalProperties, "recipients")
		delete(additionalProperties, "title")
		delete(additionalProperties, "type")
		delete(additionalProperties, "cron")
		delete(additionalProperties, "description")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "executionId")
		delete(additionalProperties, "host")
		delete(additionalProperties, "lastCompletionDate")
		delete(additionalProperties, "lastExecutionDate")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttachmentReportScheduledTask struct {
	value *AttachmentReportScheduledTask
	isSet bool
}

func (v NullableAttachmentReportScheduledTask) Get() *AttachmentReportScheduledTask {
	return v.value
}

func (v *NullableAttachmentReportScheduledTask) Set(val *AttachmentReportScheduledTask) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentReportScheduledTask) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentReportScheduledTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentReportScheduledTask(val *AttachmentReportScheduledTask) *NullableAttachmentReportScheduledTask {
	return &NullableAttachmentReportScheduledTask{value: val, isSet: true}
}

func (v NullableAttachmentReportScheduledTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentReportScheduledTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
