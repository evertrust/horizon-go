# EmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **NullableString** | The body of the e-mail | [optional] 
**From** | **string** | The sender name of the e-mail | 
**IsHtml** | **bool** | Whether the e-mail contains HTML code | 
**Title** | **string** | The title of the e-mail | 
**To** | [**[]EmailRecipient**](EmailRecipient.md) | The recipient(s) of the e-mail | 

## Methods

### NewEmailTemplate

`func NewEmailTemplate(from string, isHtml bool, title string, to []EmailRecipient, ) *EmailTemplate`

NewEmailTemplate instantiates a new EmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTemplateWithDefaults

`func NewEmailTemplateWithDefaults() *EmailTemplate`

NewEmailTemplateWithDefaults instantiates a new EmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *EmailTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailTemplate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *EmailTemplate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *EmailTemplate) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *EmailTemplate) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetFrom

`func (o *EmailTemplate) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailTemplate) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailTemplate) SetFrom(v string)`

SetFrom sets From field to given value.


### GetIsHtml

`func (o *EmailTemplate) GetIsHtml() bool`

GetIsHtml returns the IsHtml field if non-nil, zero value otherwise.

### GetIsHtmlOk

`func (o *EmailTemplate) GetIsHtmlOk() (*bool, bool)`

GetIsHtmlOk returns a tuple with the IsHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHtml

`func (o *EmailTemplate) SetIsHtml(v bool)`

SetIsHtml sets IsHtml field to given value.


### GetTitle

`func (o *EmailTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EmailTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EmailTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetTo

`func (o *EmailTemplate) GetTo() []EmailRecipient`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailTemplate) GetToOk() (*[]EmailRecipient, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailTemplate) SetTo(v []EmailRecipient)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


