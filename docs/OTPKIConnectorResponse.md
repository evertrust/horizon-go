# OTPKIConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**EndPoint** | **string** |  | 
**Profile** | **string** |  | 
**EmailMap** | Pointer to **NullableString** |  | [optional] 
**SanDnsMap** | Pointer to **NullableString** |  | [optional] 
**SanEmailMap** | Pointer to **NullableString** |  | [optional] 
**UidMap** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**ZoneLabel** | Pointer to **NullableString** | The name of the label where the zone value is stored on an enrolled certificate | [optional] 
**AuthenticationCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 

## Methods

### NewOTPKIConnectorResponse

`func NewOTPKIConnectorResponse(id string, name string, type_ string, endPoint string, profile string, authenticationCredentials string, ) *OTPKIConnectorResponse`

NewOTPKIConnectorResponse instantiates a new OTPKIConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTPKIConnectorResponseWithDefaults

`func NewOTPKIConnectorResponseWithDefaults() *OTPKIConnectorResponse`

NewOTPKIConnectorResponseWithDefaults instantiates a new OTPKIConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OTPKIConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OTPKIConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OTPKIConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OTPKIConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OTPKIConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OTPKIConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *OTPKIConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OTPKIConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OTPKIConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetEndPoint

`func (o *OTPKIConnectorResponse) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *OTPKIConnectorResponse) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *OTPKIConnectorResponse) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetProfile

`func (o *OTPKIConnectorResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OTPKIConnectorResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OTPKIConnectorResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetEmailMap

`func (o *OTPKIConnectorResponse) GetEmailMap() string`

GetEmailMap returns the EmailMap field if non-nil, zero value otherwise.

### GetEmailMapOk

`func (o *OTPKIConnectorResponse) GetEmailMapOk() (*string, bool)`

GetEmailMapOk returns a tuple with the EmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMap

`func (o *OTPKIConnectorResponse) SetEmailMap(v string)`

SetEmailMap sets EmailMap field to given value.

### HasEmailMap

`func (o *OTPKIConnectorResponse) HasEmailMap() bool`

HasEmailMap returns a boolean if a field has been set.

### SetEmailMapNil

`func (o *OTPKIConnectorResponse) SetEmailMapNil(b bool)`

 SetEmailMapNil sets the value for EmailMap to be an explicit nil

### UnsetEmailMap
`func (o *OTPKIConnectorResponse) UnsetEmailMap()`

UnsetEmailMap ensures that no value is present for EmailMap, not even an explicit nil
### GetSanDnsMap

`func (o *OTPKIConnectorResponse) GetSanDnsMap() string`

GetSanDnsMap returns the SanDnsMap field if non-nil, zero value otherwise.

### GetSanDnsMapOk

`func (o *OTPKIConnectorResponse) GetSanDnsMapOk() (*string, bool)`

GetSanDnsMapOk returns a tuple with the SanDnsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanDnsMap

`func (o *OTPKIConnectorResponse) SetSanDnsMap(v string)`

SetSanDnsMap sets SanDnsMap field to given value.

### HasSanDnsMap

`func (o *OTPKIConnectorResponse) HasSanDnsMap() bool`

HasSanDnsMap returns a boolean if a field has been set.

### SetSanDnsMapNil

`func (o *OTPKIConnectorResponse) SetSanDnsMapNil(b bool)`

 SetSanDnsMapNil sets the value for SanDnsMap to be an explicit nil

### UnsetSanDnsMap
`func (o *OTPKIConnectorResponse) UnsetSanDnsMap()`

UnsetSanDnsMap ensures that no value is present for SanDnsMap, not even an explicit nil
### GetSanEmailMap

`func (o *OTPKIConnectorResponse) GetSanEmailMap() string`

GetSanEmailMap returns the SanEmailMap field if non-nil, zero value otherwise.

### GetSanEmailMapOk

`func (o *OTPKIConnectorResponse) GetSanEmailMapOk() (*string, bool)`

GetSanEmailMapOk returns a tuple with the SanEmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanEmailMap

`func (o *OTPKIConnectorResponse) SetSanEmailMap(v string)`

SetSanEmailMap sets SanEmailMap field to given value.

### HasSanEmailMap

`func (o *OTPKIConnectorResponse) HasSanEmailMap() bool`

HasSanEmailMap returns a boolean if a field has been set.

### SetSanEmailMapNil

`func (o *OTPKIConnectorResponse) SetSanEmailMapNil(b bool)`

 SetSanEmailMapNil sets the value for SanEmailMap to be an explicit nil

### UnsetSanEmailMap
`func (o *OTPKIConnectorResponse) UnsetSanEmailMap()`

UnsetSanEmailMap ensures that no value is present for SanEmailMap, not even an explicit nil
### GetUidMap

`func (o *OTPKIConnectorResponse) GetUidMap() string`

GetUidMap returns the UidMap field if non-nil, zero value otherwise.

### GetUidMapOk

`func (o *OTPKIConnectorResponse) GetUidMapOk() (*string, bool)`

GetUidMapOk returns a tuple with the UidMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidMap

`func (o *OTPKIConnectorResponse) SetUidMap(v string)`

SetUidMap sets UidMap field to given value.

### HasUidMap

`func (o *OTPKIConnectorResponse) HasUidMap() bool`

HasUidMap returns a boolean if a field has been set.

### SetUidMapNil

`func (o *OTPKIConnectorResponse) SetUidMapNil(b bool)`

 SetUidMapNil sets the value for UidMap to be an explicit nil

### UnsetUidMap
`func (o *OTPKIConnectorResponse) UnsetUidMap()`

UnsetUidMap ensures that no value is present for UidMap, not even an explicit nil
### GetZone

`func (o *OTPKIConnectorResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *OTPKIConnectorResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *OTPKIConnectorResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *OTPKIConnectorResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *OTPKIConnectorResponse) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *OTPKIConnectorResponse) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZoneLabel

`func (o *OTPKIConnectorResponse) GetZoneLabel() string`

GetZoneLabel returns the ZoneLabel field if non-nil, zero value otherwise.

### GetZoneLabelOk

`func (o *OTPKIConnectorResponse) GetZoneLabelOk() (*string, bool)`

GetZoneLabelOk returns a tuple with the ZoneLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneLabel

`func (o *OTPKIConnectorResponse) SetZoneLabel(v string)`

SetZoneLabel sets ZoneLabel field to given value.

### HasZoneLabel

`func (o *OTPKIConnectorResponse) HasZoneLabel() bool`

HasZoneLabel returns a boolean if a field has been set.

### SetZoneLabelNil

`func (o *OTPKIConnectorResponse) SetZoneLabelNil(b bool)`

 SetZoneLabelNil sets the value for ZoneLabel to be an explicit nil

### UnsetZoneLabel
`func (o *OTPKIConnectorResponse) UnsetZoneLabel()`

UnsetZoneLabel ensures that no value is present for ZoneLabel, not even an explicit nil
### GetAuthenticationCredentials

`func (o *OTPKIConnectorResponse) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *OTPKIConnectorResponse) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *OTPKIConnectorResponse) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### GetTimeout

`func (o *OTPKIConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *OTPKIConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *OTPKIConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *OTPKIConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *OTPKIConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *OTPKIConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *OTPKIConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *OTPKIConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *OTPKIConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *OTPKIConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *OTPKIConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *OTPKIConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *OTPKIConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *OTPKIConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *OTPKIConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *OTPKIConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *OTPKIConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *OTPKIConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetStatus

`func (o *OTPKIConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OTPKIConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OTPKIConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OTPKIConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *OTPKIConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *OTPKIConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


