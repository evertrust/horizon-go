# AnalyticsStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | The number of certificate synchronized | 
**Error** | Pointer to **NullableString** | If an error happened during the synchronization process | [optional] [readonly] 
**MaxLastModification** | Pointer to **NullableInt64** | The last modification date synchronized | [optional] [readonly] 
**Ready** | **bool** | If the certificate analytics is ready to use | 

## Methods

### NewAnalyticsStatus

`func NewAnalyticsStatus(count int64, ready bool, ) *AnalyticsStatus`

NewAnalyticsStatus instantiates a new AnalyticsStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsStatusWithDefaults

`func NewAnalyticsStatusWithDefaults() *AnalyticsStatus`

NewAnalyticsStatusWithDefaults instantiates a new AnalyticsStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AnalyticsStatus) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AnalyticsStatus) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AnalyticsStatus) SetCount(v int64)`

SetCount sets Count field to given value.


### GetError

`func (o *AnalyticsStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AnalyticsStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AnalyticsStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AnalyticsStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AnalyticsStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AnalyticsStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetMaxLastModification

`func (o *AnalyticsStatus) GetMaxLastModification() int64`

GetMaxLastModification returns the MaxLastModification field if non-nil, zero value otherwise.

### GetMaxLastModificationOk

`func (o *AnalyticsStatus) GetMaxLastModificationOk() (*int64, bool)`

GetMaxLastModificationOk returns a tuple with the MaxLastModification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLastModification

`func (o *AnalyticsStatus) SetMaxLastModification(v int64)`

SetMaxLastModification sets MaxLastModification field to given value.

### HasMaxLastModification

`func (o *AnalyticsStatus) HasMaxLastModification() bool`

HasMaxLastModification returns a boolean if a field has been set.

### SetMaxLastModificationNil

`func (o *AnalyticsStatus) SetMaxLastModificationNil(b bool)`

 SetMaxLastModificationNil sets the value for MaxLastModification to be an explicit nil

### UnsetMaxLastModification
`func (o *AnalyticsStatus) UnsetMaxLastModification()`

UnsetMaxLastModification ensures that no value is present for MaxLastModification, not even an explicit nil
### GetReady

`func (o *AnalyticsStatus) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *AnalyticsStatus) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *AnalyticsStatus) SetReady(v bool)`

SetReady sets Ready field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


