# TermsOfService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | [**[]LocalizedString**](LocalizedString.md) | Localized text content displayed to the requester | 
**Description** | Pointer to **string** | Human-readable description for this Terms of Service entry | [optional] 
**Name** | **string** | Unique name for this Terms of Service entry | 

## Methods

### NewTermsOfService

`func NewTermsOfService(contents []LocalizedString, name string, ) *TermsOfService`

NewTermsOfService instantiates a new TermsOfService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermsOfServiceWithDefaults

`func NewTermsOfServiceWithDefaults() *TermsOfService`

NewTermsOfServiceWithDefaults instantiates a new TermsOfService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *TermsOfService) GetContents() []LocalizedString`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *TermsOfService) GetContentsOk() (*[]LocalizedString, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *TermsOfService) SetContents(v []LocalizedString)`

SetContents sets Contents field to given value.


### GetDescription

`func (o *TermsOfService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TermsOfService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TermsOfService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TermsOfService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *TermsOfService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TermsOfService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TermsOfService) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


