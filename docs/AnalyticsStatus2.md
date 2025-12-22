# AnalyticsStatus2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The last event id synchronized | [optional] [readonly] 
**Count** | **int64** | The number of event synchronized | 
**Error** | Pointer to **NullableString** | If an error happened during the synchronization process | [optional] [readonly] 
**Ready** | **bool** | If the event analytics is ready to use | 

## Methods

### NewAnalyticsStatus2

`func NewAnalyticsStatus2(count int64, ready bool, ) *AnalyticsStatus2`

NewAnalyticsStatus2 instantiates a new AnalyticsStatus2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsStatus2WithDefaults

`func NewAnalyticsStatus2WithDefaults() *AnalyticsStatus2`

NewAnalyticsStatus2WithDefaults instantiates a new AnalyticsStatus2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnalyticsStatus2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsStatus2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsStatus2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnalyticsStatus2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCount

`func (o *AnalyticsStatus2) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AnalyticsStatus2) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AnalyticsStatus2) SetCount(v int64)`

SetCount sets Count field to given value.


### GetError

`func (o *AnalyticsStatus2) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AnalyticsStatus2) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AnalyticsStatus2) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AnalyticsStatus2) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AnalyticsStatus2) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AnalyticsStatus2) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetReady

`func (o *AnalyticsStatus2) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *AnalyticsStatus2) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *AnalyticsStatus2) SetReady(v bool)`

SetReady sets Ready field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


