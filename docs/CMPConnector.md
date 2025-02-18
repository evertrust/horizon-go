# CMPConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** |  | 
**EndPoint** | **string** |  | 
**Profile** | **string** |  | 
**IssuerCADN** | **string** |  | 
**IssuerCACert** | **string** |  | 
**SignerCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/api.security.credentials) to use to sign on the PKI | 
**EmailMap** | Pointer to **NullableString** |  | [optional] 
**SanDnsMap** | Pointer to **NullableString** |  | [optional] 
**CnMap** | Pointer to **NullableString** |  | [optional] 
**ProfileMap** | Pointer to **NullableString** |  | [optional] 
**IssuerMap** | Pointer to **NullableString** |  | [optional] 
**LegacyCMPStyle** | Pointer to **NullableBool** |  | [optional] 
**AuthenticationCredentials** | **string** | Name of the &#x60;certificate&#x60; [credentials](#tag/api.security.credentials) to use to authenticate on the PKI | 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCMPConnector

`func NewCMPConnector(name string, type_ string, endPoint string, profile string, issuerCADN string, issuerCACert string, signerCredentials string, authenticationCredentials string, ) *CMPConnector`

NewCMPConnector instantiates a new CMPConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCMPConnectorWithDefaults

`func NewCMPConnectorWithDefaults() *CMPConnector`

NewCMPConnectorWithDefaults instantiates a new CMPConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CMPConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CMPConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CMPConnector) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CMPConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CMPConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CMPConnector) SetType(v string)`

SetType sets Type field to given value.


### GetEndPoint

`func (o *CMPConnector) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *CMPConnector) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *CMPConnector) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetProfile

`func (o *CMPConnector) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CMPConnector) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CMPConnector) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetIssuerCADN

`func (o *CMPConnector) GetIssuerCADN() string`

GetIssuerCADN returns the IssuerCADN field if non-nil, zero value otherwise.

### GetIssuerCADNOk

`func (o *CMPConnector) GetIssuerCADNOk() (*string, bool)`

GetIssuerCADNOk returns a tuple with the IssuerCADN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCADN

`func (o *CMPConnector) SetIssuerCADN(v string)`

SetIssuerCADN sets IssuerCADN field to given value.


### GetIssuerCACert

`func (o *CMPConnector) GetIssuerCACert() string`

GetIssuerCACert returns the IssuerCACert field if non-nil, zero value otherwise.

### GetIssuerCACertOk

`func (o *CMPConnector) GetIssuerCACertOk() (*string, bool)`

GetIssuerCACertOk returns a tuple with the IssuerCACert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCACert

`func (o *CMPConnector) SetIssuerCACert(v string)`

SetIssuerCACert sets IssuerCACert field to given value.


### GetSignerCredentials

`func (o *CMPConnector) GetSignerCredentials() string`

GetSignerCredentials returns the SignerCredentials field if non-nil, zero value otherwise.

### GetSignerCredentialsOk

`func (o *CMPConnector) GetSignerCredentialsOk() (*string, bool)`

GetSignerCredentialsOk returns a tuple with the SignerCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerCredentials

`func (o *CMPConnector) SetSignerCredentials(v string)`

SetSignerCredentials sets SignerCredentials field to given value.


### GetEmailMap

`func (o *CMPConnector) GetEmailMap() string`

GetEmailMap returns the EmailMap field if non-nil, zero value otherwise.

### GetEmailMapOk

`func (o *CMPConnector) GetEmailMapOk() (*string, bool)`

GetEmailMapOk returns a tuple with the EmailMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailMap

`func (o *CMPConnector) SetEmailMap(v string)`

SetEmailMap sets EmailMap field to given value.

### HasEmailMap

`func (o *CMPConnector) HasEmailMap() bool`

HasEmailMap returns a boolean if a field has been set.

### SetEmailMapNil

`func (o *CMPConnector) SetEmailMapNil(b bool)`

 SetEmailMapNil sets the value for EmailMap to be an explicit nil

### UnsetEmailMap
`func (o *CMPConnector) UnsetEmailMap()`

UnsetEmailMap ensures that no value is present for EmailMap, not even an explicit nil
### GetSanDnsMap

`func (o *CMPConnector) GetSanDnsMap() string`

GetSanDnsMap returns the SanDnsMap field if non-nil, zero value otherwise.

### GetSanDnsMapOk

`func (o *CMPConnector) GetSanDnsMapOk() (*string, bool)`

GetSanDnsMapOk returns a tuple with the SanDnsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanDnsMap

`func (o *CMPConnector) SetSanDnsMap(v string)`

SetSanDnsMap sets SanDnsMap field to given value.

### HasSanDnsMap

`func (o *CMPConnector) HasSanDnsMap() bool`

HasSanDnsMap returns a boolean if a field has been set.

### SetSanDnsMapNil

`func (o *CMPConnector) SetSanDnsMapNil(b bool)`

 SetSanDnsMapNil sets the value for SanDnsMap to be an explicit nil

### UnsetSanDnsMap
`func (o *CMPConnector) UnsetSanDnsMap()`

UnsetSanDnsMap ensures that no value is present for SanDnsMap, not even an explicit nil
### GetCnMap

`func (o *CMPConnector) GetCnMap() string`

GetCnMap returns the CnMap field if non-nil, zero value otherwise.

### GetCnMapOk

`func (o *CMPConnector) GetCnMapOk() (*string, bool)`

GetCnMapOk returns a tuple with the CnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnMap

`func (o *CMPConnector) SetCnMap(v string)`

SetCnMap sets CnMap field to given value.

### HasCnMap

`func (o *CMPConnector) HasCnMap() bool`

HasCnMap returns a boolean if a field has been set.

### SetCnMapNil

`func (o *CMPConnector) SetCnMapNil(b bool)`

 SetCnMapNil sets the value for CnMap to be an explicit nil

### UnsetCnMap
`func (o *CMPConnector) UnsetCnMap()`

UnsetCnMap ensures that no value is present for CnMap, not even an explicit nil
### GetProfileMap

`func (o *CMPConnector) GetProfileMap() string`

GetProfileMap returns the ProfileMap field if non-nil, zero value otherwise.

### GetProfileMapOk

`func (o *CMPConnector) GetProfileMapOk() (*string, bool)`

GetProfileMapOk returns a tuple with the ProfileMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileMap

`func (o *CMPConnector) SetProfileMap(v string)`

SetProfileMap sets ProfileMap field to given value.

### HasProfileMap

`func (o *CMPConnector) HasProfileMap() bool`

HasProfileMap returns a boolean if a field has been set.

### SetProfileMapNil

`func (o *CMPConnector) SetProfileMapNil(b bool)`

 SetProfileMapNil sets the value for ProfileMap to be an explicit nil

### UnsetProfileMap
`func (o *CMPConnector) UnsetProfileMap()`

UnsetProfileMap ensures that no value is present for ProfileMap, not even an explicit nil
### GetIssuerMap

`func (o *CMPConnector) GetIssuerMap() string`

GetIssuerMap returns the IssuerMap field if non-nil, zero value otherwise.

### GetIssuerMapOk

`func (o *CMPConnector) GetIssuerMapOk() (*string, bool)`

GetIssuerMapOk returns a tuple with the IssuerMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerMap

`func (o *CMPConnector) SetIssuerMap(v string)`

SetIssuerMap sets IssuerMap field to given value.

### HasIssuerMap

`func (o *CMPConnector) HasIssuerMap() bool`

HasIssuerMap returns a boolean if a field has been set.

### SetIssuerMapNil

`func (o *CMPConnector) SetIssuerMapNil(b bool)`

 SetIssuerMapNil sets the value for IssuerMap to be an explicit nil

### UnsetIssuerMap
`func (o *CMPConnector) UnsetIssuerMap()`

UnsetIssuerMap ensures that no value is present for IssuerMap, not even an explicit nil
### GetLegacyCMPStyle

`func (o *CMPConnector) GetLegacyCMPStyle() bool`

GetLegacyCMPStyle returns the LegacyCMPStyle field if non-nil, zero value otherwise.

### GetLegacyCMPStyleOk

`func (o *CMPConnector) GetLegacyCMPStyleOk() (*bool, bool)`

GetLegacyCMPStyleOk returns a tuple with the LegacyCMPStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyCMPStyle

`func (o *CMPConnector) SetLegacyCMPStyle(v bool)`

SetLegacyCMPStyle sets LegacyCMPStyle field to given value.

### HasLegacyCMPStyle

`func (o *CMPConnector) HasLegacyCMPStyle() bool`

HasLegacyCMPStyle returns a boolean if a field has been set.

### SetLegacyCMPStyleNil

`func (o *CMPConnector) SetLegacyCMPStyleNil(b bool)`

 SetLegacyCMPStyleNil sets the value for LegacyCMPStyle to be an explicit nil

### UnsetLegacyCMPStyle
`func (o *CMPConnector) UnsetLegacyCMPStyle()`

UnsetLegacyCMPStyle ensures that no value is present for LegacyCMPStyle, not even an explicit nil
### GetAuthenticationCredentials

`func (o *CMPConnector) GetAuthenticationCredentials() string`

GetAuthenticationCredentials returns the AuthenticationCredentials field if non-nil, zero value otherwise.

### GetAuthenticationCredentialsOk

`func (o *CMPConnector) GetAuthenticationCredentialsOk() (*string, bool)`

GetAuthenticationCredentialsOk returns a tuple with the AuthenticationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationCredentials

`func (o *CMPConnector) SetAuthenticationCredentials(v string)`

SetAuthenticationCredentials sets AuthenticationCredentials field to given value.


### GetTimeout

`func (o *CMPConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CMPConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CMPConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CMPConnector) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *CMPConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *CMPConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *CMPConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *CMPConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *CMPConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *CMPConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *CMPConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *CMPConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *CMPConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *CMPConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *CMPConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *CMPConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *CMPConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *CMPConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


