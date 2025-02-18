# SubjectAlternateName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SanType** | **string** | The type of the SAN | 
**Value** | **string** | The value of the SAN | 

## Methods

### NewSubjectAlternateName

`func NewSubjectAlternateName(sanType string, value string, ) *SubjectAlternateName`

NewSubjectAlternateName instantiates a new SubjectAlternateName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectAlternateNameWithDefaults

`func NewSubjectAlternateNameWithDefaults() *SubjectAlternateName`

NewSubjectAlternateNameWithDefaults instantiates a new SubjectAlternateName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSanType

`func (o *SubjectAlternateName) GetSanType() string`

GetSanType returns the SanType field if non-nil, zero value otherwise.

### GetSanTypeOk

`func (o *SubjectAlternateName) GetSanTypeOk() (*string, bool)`

GetSanTypeOk returns a tuple with the SanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanType

`func (o *SubjectAlternateName) SetSanType(v string)`

SetSanType sets SanType field to given value.


### GetValue

`func (o *SubjectAlternateName) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SubjectAlternateName) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SubjectAlternateName) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


