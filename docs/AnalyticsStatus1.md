# AnalyticsStatus1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The last discovery event id synchronized | [optional] [readonly] 
**Count** | **int64** | The number of discovery event synchronized | 
**Error** | Pointer to **NullableString** | If an error happened during the synchronization process | [optional] [readonly] 
**Ready** | **bool** | If the discovery event analytics is ready to use | 

## Methods

### NewAnalyticsStatus1

`func NewAnalyticsStatus1(count int64, ready bool, ) *AnalyticsStatus1`

NewAnalyticsStatus1 instantiates a new AnalyticsStatus1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsStatus1WithDefaults

`func NewAnalyticsStatus1WithDefaults() *AnalyticsStatus1`

NewAnalyticsStatus1WithDefaults instantiates a new AnalyticsStatus1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnalyticsStatus1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsStatus1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsStatus1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnalyticsStatus1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCount

`func (o *AnalyticsStatus1) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AnalyticsStatus1) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AnalyticsStatus1) SetCount(v int64)`

SetCount sets Count field to given value.


### GetError

`func (o *AnalyticsStatus1) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AnalyticsStatus1) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AnalyticsStatus1) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *AnalyticsStatus1) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *AnalyticsStatus1) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *AnalyticsStatus1) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetReady

`func (o *AnalyticsStatus1) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *AnalyticsStatus1) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *AnalyticsStatus1) SetReady(v bool)`

SetReady sets Ready field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


