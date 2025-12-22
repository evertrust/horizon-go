# TlsPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **int64** | The number of the port | 
**Version** | **string** | Protocol version used | 

## Methods

### NewTlsPort

`func NewTlsPort(port int64, version string, ) *TlsPort`

NewTlsPort instantiates a new TlsPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTlsPortWithDefaults

`func NewTlsPortWithDefaults() *TlsPort`

NewTlsPortWithDefaults instantiates a new TlsPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *TlsPort) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TlsPort) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TlsPort) SetPort(v int64)`

SetPort sets Port field to given value.


### GetVersion

`func (o *TlsPort) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TlsPort) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TlsPort) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


