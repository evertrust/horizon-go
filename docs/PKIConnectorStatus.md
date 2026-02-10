# PKIConnectorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastCheck** | **int64** | The date, in milliseconds since the epoch, of the last time the pki connector health check was ran. | 
**Status** | **string** | The status of the pki connector connection.  The &#39;unknown&#39; status means that the healthcheck is not available.  | 
**Message** | Pointer to **NullableString** | A meaningful message about the result of the health check (in case of error) | [optional] 

## Methods

### NewPKIConnectorStatus

`func NewPKIConnectorStatus(lastCheck int64, status string, ) *PKIConnectorStatus`

NewPKIConnectorStatus instantiates a new PKIConnectorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPKIConnectorStatusWithDefaults

`func NewPKIConnectorStatusWithDefaults() *PKIConnectorStatus`

NewPKIConnectorStatusWithDefaults instantiates a new PKIConnectorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastCheck

`func (o *PKIConnectorStatus) GetLastCheck() int64`

GetLastCheck returns the LastCheck field if non-nil, zero value otherwise.

### GetLastCheckOk

`func (o *PKIConnectorStatus) GetLastCheckOk() (*int64, bool)`

GetLastCheckOk returns a tuple with the LastCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheck

`func (o *PKIConnectorStatus) SetLastCheck(v int64)`

SetLastCheck sets LastCheck field to given value.


### GetStatus

`func (o *PKIConnectorStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PKIConnectorStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PKIConnectorStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *PKIConnectorStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PKIConnectorStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PKIConnectorStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PKIConnectorStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *PKIConnectorStatus) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *PKIConnectorStatus) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


