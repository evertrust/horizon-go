# FortiManagerConnectorResponseManagedDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adom** | **string** | Administrative domain the device belongs to | 
**Device** | **string** | Managed device name | 
**SynchronizeDevices** | Pointer to **bool** | Install the configuration onto the device after import | [optional] [default to false]
**Vdom** | **string** | Virtual domain on the managed device | 

## Methods

### NewFortiManagerConnectorResponseManagedDevice

`func NewFortiManagerConnectorResponseManagedDevice(adom string, device string, vdom string, ) *FortiManagerConnectorResponseManagedDevice`

NewFortiManagerConnectorResponseManagedDevice instantiates a new FortiManagerConnectorResponseManagedDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFortiManagerConnectorResponseManagedDeviceWithDefaults

`func NewFortiManagerConnectorResponseManagedDeviceWithDefaults() *FortiManagerConnectorResponseManagedDevice`

NewFortiManagerConnectorResponseManagedDeviceWithDefaults instantiates a new FortiManagerConnectorResponseManagedDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdom

`func (o *FortiManagerConnectorResponseManagedDevice) GetAdom() string`

GetAdom returns the Adom field if non-nil, zero value otherwise.

### GetAdomOk

`func (o *FortiManagerConnectorResponseManagedDevice) GetAdomOk() (*string, bool)`

GetAdomOk returns a tuple with the Adom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdom

`func (o *FortiManagerConnectorResponseManagedDevice) SetAdom(v string)`

SetAdom sets Adom field to given value.


### GetDevice

`func (o *FortiManagerConnectorResponseManagedDevice) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *FortiManagerConnectorResponseManagedDevice) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *FortiManagerConnectorResponseManagedDevice) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetSynchronizeDevices

`func (o *FortiManagerConnectorResponseManagedDevice) GetSynchronizeDevices() bool`

GetSynchronizeDevices returns the SynchronizeDevices field if non-nil, zero value otherwise.

### GetSynchronizeDevicesOk

`func (o *FortiManagerConnectorResponseManagedDevice) GetSynchronizeDevicesOk() (*bool, bool)`

GetSynchronizeDevicesOk returns a tuple with the SynchronizeDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizeDevices

`func (o *FortiManagerConnectorResponseManagedDevice) SetSynchronizeDevices(v bool)`

SetSynchronizeDevices sets SynchronizeDevices field to given value.

### HasSynchronizeDevices

`func (o *FortiManagerConnectorResponseManagedDevice) HasSynchronizeDevices() bool`

HasSynchronizeDevices returns a boolean if a field has been set.

### GetVdom

`func (o *FortiManagerConnectorResponseManagedDevice) GetVdom() string`

GetVdom returns the Vdom field if non-nil, zero value otherwise.

### GetVdomOk

`func (o *FortiManagerConnectorResponseManagedDevice) GetVdomOk() (*string, bool)`

GetVdomOk returns a tuple with the Vdom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdom

`func (o *FortiManagerConnectorResponseManagedDevice) SetVdom(v string)`

SetVdom sets Vdom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


