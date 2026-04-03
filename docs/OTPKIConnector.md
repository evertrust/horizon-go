# OTPKIConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/security.credentials) to use to authenticate on the PKI | 
**EmailMap** | Pointer to **NullableString** |  | [optional] 
**EndPoint** | **string** |  | 
**Name** | **string** |  | 
**Profile** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**SanDnsMap** | Pointer to **NullableString** |  | [optional] 
**SanEmailMap** | Pointer to **NullableString** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**UidMap** | Pointer to **NullableString** |  | [optional] 
**Zone** | Pointer to **NullableString** |  | [optional] 
**ZoneLabel** | Pointer to **NullableString** | The name of the label where the zone value is stored on an enrolled certificate | [optional] 

## Methods

### NewOTPKIConnector

`func NewOTPKIConnector(authenticationCredentials string, endPoint string, name string, profile string, type_ string, ) *OTPKIConnector`

NewOTPKIConnector instantiates a new OTPKIConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTPKIConnectorWithDefaults

`func NewOTPKIConnectorWithDefaults() *OTPKIConnector`

NewOTPKIConnectorWithDefaults instantiates a new OTPKIConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationCredentials

`func (o *OTPKIConnector) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *OTPKIConnector) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *OTPKIConnector) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### GetEmailMap

`func (o *OTPKIConnector) GetEmailMap() string`

GetEmailMap returns the EmailMap field if non-nil, zero value otherwise.

### GetEmailMapOk

`func (o *OTPKIConnector) GetEmailMapOk() (*string, bool)`

GetEmailMapOk returns a tuple with the EmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMap

`func (o *OTPKIConnector) SetEmailMap(v string)`

SetEmailMap sets EmailMap field to given value.

### HasEmailMap

`func (o *OTPKIConnector) HasEmailMap() bool`

HasEmailMap returns a boolean if a field has been set.

### SetEmailMapNil

`func (o *OTPKIConnector) SetEmailMapNil(b bool)`

 SetEmailMapNil sets the value for EmailMap to be an explicit nil

### UnsetEmailMap
`func (o *OTPKIConnector) UnsetEmailMap()`

UnsetEmailMap ensures that no value is present for EmailMap, not even an explicit nil
### GetEndPoint

`func (o *OTPKIConnector) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *OTPKIConnector) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *OTPKIConnector) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetName

`func (o *OTPKIConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OTPKIConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OTPKIConnector) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *OTPKIConnector) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OTPKIConnector) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OTPKIConnector) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetProxy

`func (o *OTPKIConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *OTPKIConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *OTPKIConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *OTPKIConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *OTPKIConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *OTPKIConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *OTPKIConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *OTPKIConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *OTPKIConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *OTPKIConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *OTPKIConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *OTPKIConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetSanDnsMap

`func (o *OTPKIConnector) GetSanDnsMap() string`

GetSanDnsMap returns the SanDnsMap field if non-nil, zero value otherwise.

### GetSanDnsMapOk

`func (o *OTPKIConnector) GetSanDnsMapOk() (*string, bool)`

GetSanDnsMapOk returns a tuple with the SanDnsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanDnsMap

`func (o *OTPKIConnector) SetSanDnsMap(v string)`

SetSanDnsMap sets SanDnsMap field to given value.

### HasSanDnsMap

`func (o *OTPKIConnector) HasSanDnsMap() bool`

HasSanDnsMap returns a boolean if a field has been set.

### SetSanDnsMapNil

`func (o *OTPKIConnector) SetSanDnsMapNil(b bool)`

 SetSanDnsMapNil sets the value for SanDnsMap to be an explicit nil

### UnsetSanDnsMap
`func (o *OTPKIConnector) UnsetSanDnsMap()`

UnsetSanDnsMap ensures that no value is present for SanDnsMap, not even an explicit nil
### GetSanEmailMap

`func (o *OTPKIConnector) GetSanEmailMap() string`

GetSanEmailMap returns the SanEmailMap field if non-nil, zero value otherwise.

### GetSanEmailMapOk

`func (o *OTPKIConnector) GetSanEmailMapOk() (*string, bool)`

GetSanEmailMapOk returns a tuple with the SanEmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanEmailMap

`func (o *OTPKIConnector) SetSanEmailMap(v string)`

SetSanEmailMap sets SanEmailMap field to given value.

### HasSanEmailMap

`func (o *OTPKIConnector) HasSanEmailMap() bool`

HasSanEmailMap returns a boolean if a field has been set.

### SetSanEmailMapNil

`func (o *OTPKIConnector) SetSanEmailMapNil(b bool)`

 SetSanEmailMapNil sets the value for SanEmailMap to be an explicit nil

### UnsetSanEmailMap
`func (o *OTPKIConnector) UnsetSanEmailMap()`

UnsetSanEmailMap ensures that no value is present for SanEmailMap, not even an explicit nil
### GetTimeout

`func (o *OTPKIConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *OTPKIConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *OTPKIConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *OTPKIConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *OTPKIConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *OTPKIConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *OTPKIConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OTPKIConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OTPKIConnector) SetType(v string)`

SetType sets Type field to given value.


### GetUidMap

`func (o *OTPKIConnector) GetUidMap() string`

GetUidMap returns the UidMap field if non-nil, zero value otherwise.

### GetUidMapOk

`func (o *OTPKIConnector) GetUidMapOk() (*string, bool)`

GetUidMapOk returns a tuple with the UidMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidMap

`func (o *OTPKIConnector) SetUidMap(v string)`

SetUidMap sets UidMap field to given value.

### HasUidMap

`func (o *OTPKIConnector) HasUidMap() bool`

HasUidMap returns a boolean if a field has been set.

### SetUidMapNil

`func (o *OTPKIConnector) SetUidMapNil(b bool)`

 SetUidMapNil sets the value for UidMap to be an explicit nil

### UnsetUidMap
`func (o *OTPKIConnector) UnsetUidMap()`

UnsetUidMap ensures that no value is present for UidMap, not even an explicit nil
### GetZone

`func (o *OTPKIConnector) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *OTPKIConnector) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *OTPKIConnector) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *OTPKIConnector) HasZone() bool`

HasZone returns a boolean if a field has been set.

### SetZoneNil

`func (o *OTPKIConnector) SetZoneNil(b bool)`

 SetZoneNil sets the value for Zone to be an explicit nil

### UnsetZone
`func (o *OTPKIConnector) UnsetZone()`

UnsetZone ensures that no value is present for Zone, not even an explicit nil
### GetZoneLabel

`func (o *OTPKIConnector) GetZoneLabel() string`

GetZoneLabel returns the ZoneLabel field if non-nil, zero value otherwise.

### GetZoneLabelOk

`func (o *OTPKIConnector) GetZoneLabelOk() (*string, bool)`

GetZoneLabelOk returns a tuple with the ZoneLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneLabel

`func (o *OTPKIConnector) SetZoneLabel(v string)`

SetZoneLabel sets ZoneLabel field to given value.

### HasZoneLabel

`func (o *OTPKIConnector) HasZoneLabel() bool`

HasZoneLabel returns a boolean if a field has been set.

### SetZoneLabelNil

`func (o *OTPKIConnector) SetZoneLabelNil(b bool)`

 SetZoneLabelNil sets the value for ZoneLabel to be an explicit nil

### UnsetZoneLabel
`func (o *OTPKIConnector) UnsetZoneLabel()`

UnsetZoneLabel ensures that no value is present for ZoneLabel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


