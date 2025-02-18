# WcceForestMappingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Forest** | **string** |  | 
**TemplateMappings** | [**[]WcceTemplateMapping**](WcceTemplateMapping.md) |  | 

## Methods

### NewWcceForestMappingResponse

`func NewWcceForestMappingResponse(id string, forest string, templateMappings []WcceTemplateMapping, ) *WcceForestMappingResponse`

NewWcceForestMappingResponse instantiates a new WcceForestMappingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWcceForestMappingResponseWithDefaults

`func NewWcceForestMappingResponseWithDefaults() *WcceForestMappingResponse`

NewWcceForestMappingResponseWithDefaults instantiates a new WcceForestMappingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WcceForestMappingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WcceForestMappingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WcceForestMappingResponse) SetId(v string)`

SetId sets Id field to given value.


### GetForest

`func (o *WcceForestMappingResponse) GetForest() string`

GetForest returns the Forest field if non-nil, zero value otherwise.

### GetForestOk

`func (o *WcceForestMappingResponse) GetForestOk() (*string, bool)`

GetForestOk returns a tuple with the Forest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForest

`func (o *WcceForestMappingResponse) SetForest(v string)`

SetForest sets Forest field to given value.


### GetTemplateMappings

`func (o *WcceForestMappingResponse) GetTemplateMappings() []WcceTemplateMapping`

GetTemplateMappings returns the TemplateMappings field if non-nil, zero value otherwise.

### GetTemplateMappingsOk

`func (o *WcceForestMappingResponse) GetTemplateMappingsOk() (*[]WcceTemplateMapping, bool)`

GetTemplateMappingsOk returns a tuple with the TemplateMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateMappings

`func (o *WcceForestMappingResponse) SetTemplateMappings(v []WcceTemplateMapping)`

SetTemplateMappings sets TemplateMappings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


