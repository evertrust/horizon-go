# DiscoveryEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**ActorId** | Pointer to **NullableString** | The identifier of the principal that was used when the event was raised | [optional] 
**Campaign** | **string** | The name of the discovery campaign concerned by the event | 
**CertificateId** | Pointer to **NullableString** | The ID of the certificate concerned by the event (in Horizon) | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ClientIp** | Pointer to **NullableString** | The IP of the machine where the Horizon client is running from | [optional] 
**ClientVersion** | Pointer to **NullableString** | The version of the discovery client that raised the event | [optional] 
**Code** | **string** | The code of the event to raise in the discovery events | 
**ErrorCode** | Pointer to **NullableString** | The error code of the event | [optional] 
**ErrorMessage** | Pointer to **NullableString** | The error message of the event | [optional] 
**Hostname** | Pointer to **NullableString** | The hostname concerned by the event | [optional] 
**Ip** | Pointer to **NullableString** | The IP address concerned by the event | [optional] 
**Port** | Pointer to **NullableInt64** | The TCP port concerned by the event | [optional] 
**RemoveAt** | Pointer to **NullableInt64** |  | [optional] 
**SessionId** | Pointer to **NullableString** | The ID of the discovery feed session | [optional] 
**Source** | Pointer to **NullableString** | The type of discovery that raised the event | [optional] 
**Status** | **string** | The type of event to raise | 
**Timestamp** | Pointer to **NullableInt64** | When did the event occur (Unix timestamp in milliseconds) | [optional] 

## Methods

### NewDiscoveryEventResponse

`func NewDiscoveryEventResponse(id string, campaign string, code string, status string, ) *DiscoveryEventResponse`

NewDiscoveryEventResponse instantiates a new DiscoveryEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryEventResponseWithDefaults

`func NewDiscoveryEventResponseWithDefaults() *DiscoveryEventResponse`

NewDiscoveryEventResponseWithDefaults instantiates a new DiscoveryEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscoveryEventResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscoveryEventResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscoveryEventResponse) SetId(v string)`

SetId sets Id field to given value.


### GetActorId

`func (o *DiscoveryEventResponse) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *DiscoveryEventResponse) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *DiscoveryEventResponse) SetActorId(v string)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *DiscoveryEventResponse) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### SetActorIdNil

`func (o *DiscoveryEventResponse) SetActorIdNil(b bool)`

 SetActorIdNil sets the value for ActorId to be an explicit nil

### UnsetActorId
`func (o *DiscoveryEventResponse) UnsetActorId()`

UnsetActorId ensures that no value is present for ActorId, not even an explicit nil
### GetCampaign

`func (o *DiscoveryEventResponse) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *DiscoveryEventResponse) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *DiscoveryEventResponse) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.


### GetCertificateId

`func (o *DiscoveryEventResponse) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *DiscoveryEventResponse) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *DiscoveryEventResponse) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *DiscoveryEventResponse) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *DiscoveryEventResponse) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *DiscoveryEventResponse) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetClientId

`func (o *DiscoveryEventResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DiscoveryEventResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DiscoveryEventResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *DiscoveryEventResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *DiscoveryEventResponse) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *DiscoveryEventResponse) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientIp

`func (o *DiscoveryEventResponse) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *DiscoveryEventResponse) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *DiscoveryEventResponse) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *DiscoveryEventResponse) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *DiscoveryEventResponse) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *DiscoveryEventResponse) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetClientVersion

`func (o *DiscoveryEventResponse) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *DiscoveryEventResponse) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *DiscoveryEventResponse) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *DiscoveryEventResponse) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.

### SetClientVersionNil

`func (o *DiscoveryEventResponse) SetClientVersionNil(b bool)`

 SetClientVersionNil sets the value for ClientVersion to be an explicit nil

### UnsetClientVersion
`func (o *DiscoveryEventResponse) UnsetClientVersion()`

UnsetClientVersion ensures that no value is present for ClientVersion, not even an explicit nil
### GetCode

`func (o *DiscoveryEventResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DiscoveryEventResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DiscoveryEventResponse) SetCode(v string)`

SetCode sets Code field to given value.


### GetErrorCode

`func (o *DiscoveryEventResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *DiscoveryEventResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *DiscoveryEventResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *DiscoveryEventResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *DiscoveryEventResponse) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *DiscoveryEventResponse) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *DiscoveryEventResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *DiscoveryEventResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *DiscoveryEventResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *DiscoveryEventResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *DiscoveryEventResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *DiscoveryEventResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetHostname

`func (o *DiscoveryEventResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DiscoveryEventResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DiscoveryEventResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DiscoveryEventResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *DiscoveryEventResponse) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *DiscoveryEventResponse) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetIp

`func (o *DiscoveryEventResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DiscoveryEventResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DiscoveryEventResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DiscoveryEventResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *DiscoveryEventResponse) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *DiscoveryEventResponse) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetPort

`func (o *DiscoveryEventResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DiscoveryEventResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DiscoveryEventResponse) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DiscoveryEventResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *DiscoveryEventResponse) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *DiscoveryEventResponse) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetRemoveAt

`func (o *DiscoveryEventResponse) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *DiscoveryEventResponse) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *DiscoveryEventResponse) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.

### HasRemoveAt

`func (o *DiscoveryEventResponse) HasRemoveAt() bool`

HasRemoveAt returns a boolean if a field has been set.

### SetRemoveAtNil

`func (o *DiscoveryEventResponse) SetRemoveAtNil(b bool)`

 SetRemoveAtNil sets the value for RemoveAt to be an explicit nil

### UnsetRemoveAt
`func (o *DiscoveryEventResponse) UnsetRemoveAt()`

UnsetRemoveAt ensures that no value is present for RemoveAt, not even an explicit nil
### GetSessionId

`func (o *DiscoveryEventResponse) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *DiscoveryEventResponse) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *DiscoveryEventResponse) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *DiscoveryEventResponse) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *DiscoveryEventResponse) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *DiscoveryEventResponse) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetSource

`func (o *DiscoveryEventResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DiscoveryEventResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DiscoveryEventResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DiscoveryEventResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *DiscoveryEventResponse) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *DiscoveryEventResponse) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetStatus

`func (o *DiscoveryEventResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiscoveryEventResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiscoveryEventResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *DiscoveryEventResponse) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DiscoveryEventResponse) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DiscoveryEventResponse) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DiscoveryEventResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DiscoveryEventResponse) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DiscoveryEventResponse) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


