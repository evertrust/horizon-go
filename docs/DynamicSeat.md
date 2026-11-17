# DynamicSeat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessUnitId** | **string** | Identifier of the DigiCert ONE business unit the seats are created in. | 
**SeatType** | **string** | Type of the created seats. | 
**Type** | **string** |  | 

## Methods

### NewDynamicSeat

`func NewDynamicSeat(businessUnitId string, seatType string, type_ string, ) *DynamicSeat`

NewDynamicSeat instantiates a new DynamicSeat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSeatWithDefaults

`func NewDynamicSeatWithDefaults() *DynamicSeat`

NewDynamicSeatWithDefaults instantiates a new DynamicSeat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessUnitId

`func (o *DynamicSeat) GetBusinessUnitId() string`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *DynamicSeat) GetBusinessUnitIdOk() (*string, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *DynamicSeat) SetBusinessUnitId(v string)`

SetBusinessUnitId sets BusinessUnitId field to given value.


### GetSeatType

`func (o *DynamicSeat) GetSeatType() string`

GetSeatType returns the SeatType field if non-nil, zero value otherwise.

### GetSeatTypeOk

`func (o *DynamicSeat) GetSeatTypeOk() (*string, bool)`

GetSeatTypeOk returns a tuple with the SeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeatType

`func (o *DynamicSeat) SetSeatType(v string)`

SetSeatType sets SeatType field to given value.


### GetType

`func (o *DynamicSeat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DynamicSeat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DynamicSeat) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


