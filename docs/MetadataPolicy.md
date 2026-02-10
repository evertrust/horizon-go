# MetadataPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | **string** |  | 
**EditableByRequester** | **bool** |  | 
**EditableByApprover** | **bool** |  | 

## Methods

### NewMetadataPolicy

`func NewMetadataPolicy(metadata string, editableByRequester bool, editableByApprover bool, ) *MetadataPolicy`

NewMetadataPolicy instantiates a new MetadataPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataPolicyWithDefaults

`func NewMetadataPolicyWithDefaults() *MetadataPolicy`

NewMetadataPolicyWithDefaults instantiates a new MetadataPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *MetadataPolicy) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetadataPolicy) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetadataPolicy) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetEditableByRequester

`func (o *MetadataPolicy) GetEditableByRequester() bool`

GetEditableByRequester returns the EditableByRequester field if non-nil, zero value otherwise.

### GetEditableByRequesterOk

`func (o *MetadataPolicy) GetEditableByRequesterOk() (*bool, bool)`

GetEditableByRequesterOk returns a tuple with the EditableByRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByRequester

`func (o *MetadataPolicy) SetEditableByRequester(v bool)`

SetEditableByRequester sets EditableByRequester field to given value.


### GetEditableByApprover

`func (o *MetadataPolicy) GetEditableByApprover() bool`

GetEditableByApprover returns the EditableByApprover field if non-nil, zero value otherwise.

### GetEditableByApproverOk

`func (o *MetadataPolicy) GetEditableByApproverOk() (*bool, bool)`

GetEditableByApproverOk returns a tuple with the EditableByApprover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableByApprover

`func (o *MetadataPolicy) SetEditableByApprover(v bool)`

SetEditableByApprover sets EditableByApprover field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


