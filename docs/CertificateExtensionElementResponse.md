# CertificateExtensionElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputationRule** | Pointer to **NullableString** | Computation rule input will be evaluated and will override all other inputs | [optional] 
**Editable** | Pointer to **NullableBool** | Whether the extension element is editable by the requester | [optional] 
**Mandatory** | Pointer to **NullableBool** | Whether the extension element is mandatory to submit this request | [optional] 
**Regex** | Pointer to **NullableString** | The regular expression to validate the extension element | [optional] 
**Type** | **string** | The type of the extension element | 
**Value** | Pointer to **NullableString** | The value of the extension element | [optional] 

## Methods

### NewCertificateExtensionElementResponse

`func NewCertificateExtensionElementResponse(type_ string, ) *CertificateExtensionElementResponse`

NewCertificateExtensionElementResponse instantiates a new CertificateExtensionElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateExtensionElementResponseWithDefaults

`func NewCertificateExtensionElementResponseWithDefaults() *CertificateExtensionElementResponse`

NewCertificateExtensionElementResponseWithDefaults instantiates a new CertificateExtensionElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputationRule

`func (o *CertificateExtensionElementResponse) GetComputationRule() string`

GetComputationRule returns the ComputationRule field if non-nil, zero value otherwise.

### GetComputationRuleOk

`func (o *CertificateExtensionElementResponse) GetComputationRuleOk() (*string, bool)`

GetComputationRuleOk returns a tuple with the ComputationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputationRule

`func (o *CertificateExtensionElementResponse) SetComputationRule(v string)`

SetComputationRule sets ComputationRule field to given value.

### HasComputationRule

`func (o *CertificateExtensionElementResponse) HasComputationRule() bool`

HasComputationRule returns a boolean if a field has been set.

### SetComputationRuleNil

`func (o *CertificateExtensionElementResponse) SetComputationRuleNil(b bool)`

 SetComputationRuleNil sets the value for ComputationRule to be an explicit nil

### UnsetComputationRule
`func (o *CertificateExtensionElementResponse) UnsetComputationRule()`

UnsetComputationRule ensures that no value is present for ComputationRule, not even an explicit nil
### GetEditable

`func (o *CertificateExtensionElementResponse) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *CertificateExtensionElementResponse) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *CertificateExtensionElementResponse) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *CertificateExtensionElementResponse) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditableNil

`func (o *CertificateExtensionElementResponse) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *CertificateExtensionElementResponse) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetMandatory

`func (o *CertificateExtensionElementResponse) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *CertificateExtensionElementResponse) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *CertificateExtensionElementResponse) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *CertificateExtensionElementResponse) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.

### SetMandatoryNil

`func (o *CertificateExtensionElementResponse) SetMandatoryNil(b bool)`

 SetMandatoryNil sets the value for Mandatory to be an explicit nil

### UnsetMandatory
`func (o *CertificateExtensionElementResponse) UnsetMandatory()`

UnsetMandatory ensures that no value is present for Mandatory, not even an explicit nil
### GetRegex

`func (o *CertificateExtensionElementResponse) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *CertificateExtensionElementResponse) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *CertificateExtensionElementResponse) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *CertificateExtensionElementResponse) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *CertificateExtensionElementResponse) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *CertificateExtensionElementResponse) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetType

`func (o *CertificateExtensionElementResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateExtensionElementResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateExtensionElementResponse) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *CertificateExtensionElementResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CertificateExtensionElementResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CertificateExtensionElementResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CertificateExtensionElementResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CertificateExtensionElementResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CertificateExtensionElementResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


