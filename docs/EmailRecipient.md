# EmailRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of email recipient. Apart from the &#x60;static&#x60; recipient, all are deduced from the request&#39;s context. | 
**Email** | Pointer to **NullableString** | Mandatory for &#x60;static&#x60; recipient and ignored otherwise. The address to send the email to. | [optional] 
**Label** | Pointer to **NullableString** | Mandatory for &#x60;label&#x60; recipient and ignored otherwise. The label name to fetch the address from. | [optional] 

## Methods

### NewEmailRecipient

`func NewEmailRecipient(type_ string, ) *EmailRecipient`

NewEmailRecipient instantiates a new EmailRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailRecipientWithDefaults

`func NewEmailRecipientWithDefaults() *EmailRecipient`

NewEmailRecipientWithDefaults instantiates a new EmailRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EmailRecipient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailRecipient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailRecipient) SetType(v string)`

SetType sets Type field to given value.


### GetEmail

`func (o *EmailRecipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailRecipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailRecipient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmailRecipient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *EmailRecipient) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *EmailRecipient) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetLabel

`func (o *EmailRecipient) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EmailRecipient) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EmailRecipient) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EmailRecipient) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *EmailRecipient) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *EmailRecipient) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


