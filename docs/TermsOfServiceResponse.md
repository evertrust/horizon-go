# TermsOfServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Contents** | [**[]LocalizedString**](LocalizedString.md) | Localized text content displayed to the requester | 
**Description** | Pointer to **string** | Human-readable description for this Terms of Service entry | [optional] 
**Name** | **string** | Unique name for this Terms of Service entry | 

## Methods

### NewTermsOfServiceResponse

`func NewTermsOfServiceResponse(id string, contents []LocalizedString, name string, ) *TermsOfServiceResponse`

NewTermsOfServiceResponse instantiates a new TermsOfServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTermsOfServiceResponseWithDefaults

`func NewTermsOfServiceResponseWithDefaults() *TermsOfServiceResponse`

NewTermsOfServiceResponseWithDefaults instantiates a new TermsOfServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TermsOfServiceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TermsOfServiceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TermsOfServiceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetContents

`func (o *TermsOfServiceResponse) GetContents() []LocalizedString`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *TermsOfServiceResponse) GetContentsOk() (*[]LocalizedString, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *TermsOfServiceResponse) SetContents(v []LocalizedString)`

SetContents sets Contents field to given value.


### GetDescription

`func (o *TermsOfServiceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TermsOfServiceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TermsOfServiceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TermsOfServiceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *TermsOfServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TermsOfServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TermsOfServiceResponse) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


