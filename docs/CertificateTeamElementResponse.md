# CertificateTeamElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorized** | Pointer to **[]string** | The list of authorized teams | [optional] 
**ComputationRule** | Pointer to **NullableString** | Computation rule input will be evaluated and will override all other inputs | [optional] 
**Description** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The description of the team element | [optional] 
**Editable** | Pointer to **NullableBool** | Whether the team element is editable by the requester | [optional] 
**Mandatory** | Pointer to **NullableBool** | Whether the team element is mandatory to submit this request | [optional] 
**Value** | Pointer to **NullableString** | The value of the team element. This should be a team identifier | [optional] 

## Methods

### NewCertificateTeamElementResponse

`func NewCertificateTeamElementResponse() *CertificateTeamElementResponse`

NewCertificateTeamElementResponse instantiates a new CertificateTeamElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateTeamElementResponseWithDefaults

`func NewCertificateTeamElementResponseWithDefaults() *CertificateTeamElementResponse`

NewCertificateTeamElementResponseWithDefaults instantiates a new CertificateTeamElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorized

`func (o *CertificateTeamElementResponse) GetAuthorized() []string`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *CertificateTeamElementResponse) GetAuthorizedOk() (*[]string, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *CertificateTeamElementResponse) SetAuthorized(v []string)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *CertificateTeamElementResponse) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### SetAuthorizedNil

`func (o *CertificateTeamElementResponse) SetAuthorizedNil(b bool)`

 SetAuthorizedNil sets the value for Authorized to be an explicit nil

### UnsetAuthorized
`func (o *CertificateTeamElementResponse) UnsetAuthorized()`

UnsetAuthorized ensures that no value is present for Authorized, not even an explicit nil
### GetComputationRule

`func (o *CertificateTeamElementResponse) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *CertificateTeamElementResponse) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *CertificateTeamElementResponse) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *CertificateTeamElementResponse) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *CertificateTeamElementResponse) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *CertificateTeamElementResponse) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetDescription

`func (o *CertificateTeamElementResponse) GetDescription() []LocalizedString`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CertificateTeamElementResponse) GetDescriptionOk() (*[]LocalizedString, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CertificateTeamElementResponse) SetDescription(v []LocalizedString)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CertificateTeamElementResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CertificateTeamElementResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CertificateTeamElementResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEditable

`func (o *CertificateTeamElementResponse) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *CertificateTeamElementResponse) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *CertificateTeamElementResponse) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *CertificateTeamElementResponse) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditableNil

`func (o *CertificateTeamElementResponse) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *CertificateTeamElementResponse) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetMandatory

`func (o *CertificateTeamElementResponse) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *CertificateTeamElementResponse) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *CertificateTeamElementResponse) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *CertificateTeamElementResponse) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### SetMandatoryNil

`func (o *CertificateTeamElementResponse) SetMandatoryNil(b bool)`

 SetMandatoryNil sets the value for Mandatory to be an explicit nil

### UnsetMandatory
`func (o *CertificateTeamElementResponse) UnsetMandatory()`

UnsetMandatory ensures that no value is present for Mandatory, not even an explicit nil
### GetValue

`func (o *CertificateTeamElementResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CertificateTeamElementResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CertificateTeamElementResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CertificateTeamElementResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CertificateTeamElementResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CertificateTeamElementResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


