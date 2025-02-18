# DiscoveryEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The code of the event to raise in the discovery events | 
**Campaign** | **string** | The name of the discovery campaign concerned by the event | 
**SessionId** | Pointer to **NullableString** | The ID of the discovery feed session | [optional] 
**Status** | **string** | The type of event to raise | 
**ErrorCode** | Pointer to **NullableString** | The error code of the event | [optional] 
**ErrorMessage** | Pointer to **NullableString** | The error message of the event | [optional] 
**Timestamp** | Pointer to **NullableInt64** | When did the event occur (Unix timestamp in milliseconds) | [optional] 
**RemoveAt** | Pointer to **NullableInt64** |  | [optional] 
**ClientVersion** | Pointer to **NullableString** | The version of the discovery client that raised the event | [optional] 
**ClientIp** | Pointer to **NullableString** | The IP of the machine where the Horizon client is running from | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ActorId** | Pointer to **NullableString** | The identifier of the principal that was used when the event was raised | [optional] 
**CertificateId** | Pointer to **NullableString** | The ID of the certificate concerned by the event (in Horizon) | [optional] 
**Hostname** | Pointer to **NullableString** | The hostname concerned by the event | [optional] 
**Ip** | Pointer to **NullableString** | The IP address concerned by the event | [optional] 
**Port** | Pointer to **NullableInt64** | The TCP port concerned by the event | [optional] 
**Source** | Pointer to **NullableString** | The type of discovery that raised the event | [optional] 

## Methods

### NewDiscoveryEvent

`func NewDiscoveryEvent(code string, campaign string, status string, ) *DiscoveryEvent`

NewDiscoveryEvent instantiates a new DiscoveryEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryEventWithDefaults

`func NewDiscoveryEventWithDefaults() *DiscoveryEvent`

NewDiscoveryEventWithDefaults instantiates a new DiscoveryEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DiscoveryEvent) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DiscoveryEvent) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DiscoveryEvent) SetCode(v string)`

SetCode sets Code field to given value.


### GetCampaign

`func (o *DiscoveryEvent) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *DiscoveryEvent) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *DiscoveryEvent) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.


### GetSessionId

`func (o *DiscoveryEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *DiscoveryEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *DiscoveryEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *DiscoveryEvent) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *DiscoveryEvent) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *DiscoveryEvent) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetStatus

`func (o *DiscoveryEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiscoveryEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiscoveryEvent) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorCode

`func (o *DiscoveryEvent) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *DiscoveryEvent) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *DiscoveryEvent) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *DiscoveryEvent) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *DiscoveryEvent) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *DiscoveryEvent) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *DiscoveryEvent) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *DiscoveryEvent) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *DiscoveryEvent) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *DiscoveryEvent) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *DiscoveryEvent) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *DiscoveryEvent) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetTimestamp

`func (o *DiscoveryEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DiscoveryEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DiscoveryEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DiscoveryEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DiscoveryEvent) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DiscoveryEvent) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRemoveAt

`func (o *DiscoveryEvent) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *DiscoveryEvent) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *DiscoveryEvent) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.

### HasRemoveAt

`func (o *DiscoveryEvent) HasRemoveAt() bool`

HasRemoveAt returns a boolean if a field has been set.

### SetRemoveAtNil

`func (o *DiscoveryEvent) SetRemoveAtNil(b bool)`

 SetRemoveAtNil sets the value for RemoveAt to be an explicit nil

### UnsetRemoveAt
`func (o *DiscoveryEvent) UnsetRemoveAt()`

UnsetRemoveAt ensures that no value is present for RemoveAt, not even an explicit nil
### GetClientVersion

`func (o *DiscoveryEvent) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *DiscoveryEvent) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *DiscoveryEvent) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *DiscoveryEvent) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.

### SetClientVersionNil

`func (o *DiscoveryEvent) SetClientVersionNil(b bool)`

 SetClientVersionNil sets the value for ClientVersion to be an explicit nil

### UnsetClientVersion
`func (o *DiscoveryEvent) UnsetClientVersion()`

UnsetClientVersion ensures that no value is present for ClientVersion, not even an explicit nil
### GetClientIp

`func (o *DiscoveryEvent) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *DiscoveryEvent) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *DiscoveryEvent) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *DiscoveryEvent) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### SetClientIpNil

`func (o *DiscoveryEvent) SetClientIpNil(b bool)`

 SetClientIpNil sets the value for ClientIp to be an explicit nil

### UnsetClientIp
`func (o *DiscoveryEvent) UnsetClientIp()`

UnsetClientIp ensures that no value is present for ClientIp, not even an explicit nil
### GetClientId

`func (o *DiscoveryEvent) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DiscoveryEvent) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DiscoveryEvent) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *DiscoveryEvent) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *DiscoveryEvent) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *DiscoveryEvent) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetActorId

`func (o *DiscoveryEvent) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *DiscoveryEvent) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *DiscoveryEvent) SetActorId(v string)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *DiscoveryEvent) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### SetActorIdNil

`func (o *DiscoveryEvent) SetActorIdNil(b bool)`

 SetActorIdNil sets the value for ActorId to be an explicit nil

### UnsetActorId
`func (o *DiscoveryEvent) UnsetActorId()`

UnsetActorId ensures that no value is present for ActorId, not even an explicit nil
### GetCertificateId

`func (o *DiscoveryEvent) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *DiscoveryEvent) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *DiscoveryEvent) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *DiscoveryEvent) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### SetCertificateIdNil

`func (o *DiscoveryEvent) SetCertificateIdNil(b bool)`

 SetCertificateIdNil sets the value for CertificateId to be an explicit nil

### UnsetCertificateId
`func (o *DiscoveryEvent) UnsetCertificateId()`

UnsetCertificateId ensures that no value is present for CertificateId, not even an explicit nil
### GetHostname

`func (o *DiscoveryEvent) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *DiscoveryEvent) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *DiscoveryEvent) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *DiscoveryEvent) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *DiscoveryEvent) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *DiscoveryEvent) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetIp

`func (o *DiscoveryEvent) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DiscoveryEvent) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DiscoveryEvent) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DiscoveryEvent) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *DiscoveryEvent) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *DiscoveryEvent) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetPort

`func (o *DiscoveryEvent) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DiscoveryEvent) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DiscoveryEvent) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DiscoveryEvent) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *DiscoveryEvent) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *DiscoveryEvent) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetSource

`func (o *DiscoveryEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DiscoveryEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DiscoveryEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DiscoveryEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *DiscoveryEvent) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *DiscoveryEvent) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


