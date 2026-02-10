# ReportRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Email** | Pointer to **NullableString** |  | [optional] 
**Team** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewReportRecipient

`func NewReportRecipient(type_ string, ) *ReportRecipient`

NewReportRecipient instantiates a new ReportRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportRecipientWithDefaults

`func NewReportRecipientWithDefaults() *ReportRecipient`

NewReportRecipientWithDefaults instantiates a new ReportRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ReportRecipient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportRecipient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportRecipient) SetType(v string)`

SetType sets Type field to given value.


### GetEmail

`func (o *ReportRecipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ReportRecipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ReportRecipient) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ReportRecipient) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ReportRecipient) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ReportRecipient) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetTeam

`func (o *ReportRecipient) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ReportRecipient) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ReportRecipient) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ReportRecipient) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ReportRecipient) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ReportRecipient) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


