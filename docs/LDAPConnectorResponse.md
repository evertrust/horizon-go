# LDAPConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Type** | **string** |  | 
**Name** | **string** |  | 
**Hostname** | **string** |  | 
**Port** | Pointer to **NullableInt64** |  | [optional] 
**BaseDn** | **string** |  | 
**Filter** | Pointer to **NullableString** |  | [optional] 
**CertAttr** | Pointer to **NullableString** |  | [optional] 
**FollowReferrals** | Pointer to **NullableBool** |  | [optional] 
**UserIdentifierAttribute** | **string** |  | 
**CertificateAttribute** | **string** |  | 
**ThrottleDuration** | **string** |  | 
**ThrottleParallelism** | Pointer to **int64** |  | [optional] 
**Timeout** | Pointer to **NullableString** |  | [optional] 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Credentials** | **string** | Name of the &#x60;password&#x60; [credentials](#tag/api.security.credentials) containing login DN and password. | 
**MaxStoredCertificatePerHolder** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewLDAPConnectorResponse

`func NewLDAPConnectorResponse(id string, type_ string, name string, hostname string, baseDn string, userIdentifierAttribute string, certificateAttribute string, throttleDuration string, credentials string, ) *LDAPConnectorResponse`

NewLDAPConnectorResponse instantiates a new LDAPConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPConnectorResponseWithDefaults

`func NewLDAPConnectorResponseWithDefaults() *LDAPConnectorResponse`

NewLDAPConnectorResponseWithDefaults instantiates a new LDAPConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LDAPConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LDAPConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LDAPConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *LDAPConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LDAPConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LDAPConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *LDAPConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LDAPConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LDAPConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetHostname

`func (o *LDAPConnectorResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LDAPConnectorResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LDAPConnectorResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *LDAPConnectorResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LDAPConnectorResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LDAPConnectorResponse) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *LDAPConnectorResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *LDAPConnectorResponse) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *LDAPConnectorResponse) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetBaseDn

`func (o *LDAPConnectorResponse) GetBaseDn() string`

GetBaseDn returns the BaseDn field if non-nil, zero value otherwise.

### GetBaseDnOk

`func (o *LDAPConnectorResponse) GetBaseDnOk() (*string, bool)`

GetBaseDnOk returns a tuple with the BaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDn

`func (o *LDAPConnectorResponse) SetBaseDn(v string)`

SetBaseDn sets BaseDn field to given value.


### GetFilter

`func (o *LDAPConnectorResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LDAPConnectorResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LDAPConnectorResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LDAPConnectorResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *LDAPConnectorResponse) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *LDAPConnectorResponse) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetCertAttr

`func (o *LDAPConnectorResponse) GetCertAttr() string`

GetCertAttr returns the CertAttr field if non-nil, zero value otherwise.

### GetCertAttrOk

`func (o *LDAPConnectorResponse) GetCertAttrOk() (*string, bool)`

GetCertAttrOk returns a tuple with the CertAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAttr

`func (o *LDAPConnectorResponse) SetCertAttr(v string)`

SetCertAttr sets CertAttr field to given value.

### HasCertAttr

`func (o *LDAPConnectorResponse) HasCertAttr() bool`

HasCertAttr returns a boolean if a field has been set.

### SetCertAttrNil

`func (o *LDAPConnectorResponse) SetCertAttrNil(b bool)`

 SetCertAttrNil sets the value for CertAttr to be an explicit nil

### UnsetCertAttr
`func (o *LDAPConnectorResponse) UnsetCertAttr()`

UnsetCertAttr ensures that no value is present for CertAttr, not even an explicit nil
### GetFollowReferrals

`func (o *LDAPConnectorResponse) GetFollowReferrals() bool`

GetFollowReferrals returns the FollowReferrals field if non-nil, zero value otherwise.

### GetFollowReferralsOk

`func (o *LDAPConnectorResponse) GetFollowReferralsOk() (*bool, bool)`

GetFollowReferralsOk returns a tuple with the FollowReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowReferrals

`func (o *LDAPConnectorResponse) SetFollowReferrals(v bool)`

SetFollowReferrals sets FollowReferrals field to given value.

### HasFollowReferrals

`func (o *LDAPConnectorResponse) HasFollowReferrals() bool`

HasFollowReferrals returns a boolean if a field has been set.

### SetFollowReferralsNil

`func (o *LDAPConnectorResponse) SetFollowReferralsNil(b bool)`

 SetFollowReferralsNil sets the value for FollowReferrals to be an explicit nil

### UnsetFollowReferrals
`func (o *LDAPConnectorResponse) UnsetFollowReferrals()`

UnsetFollowReferrals ensures that no value is present for FollowReferrals, not even an explicit nil
### GetUserIdentifierAttribute

`func (o *LDAPConnectorResponse) GetUserIdentifierAttribute() string`

GetUserIdentifierAttribute returns the UserIdentifierAttribute field if non-nil, zero value otherwise.

### GetUserIdentifierAttributeOk

`func (o *LDAPConnectorResponse) GetUserIdentifierAttributeOk() (*string, bool)`

GetUserIdentifierAttributeOk returns a tuple with the UserIdentifierAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentifierAttribute

`func (o *LDAPConnectorResponse) SetUserIdentifierAttribute(v string)`

SetUserIdentifierAttribute sets UserIdentifierAttribute field to given value.


### GetCertificateAttribute

`func (o *LDAPConnectorResponse) GetCertificateAttribute() string`

GetCertificateAttribute returns the CertificateAttribute field if non-nil, zero value otherwise.

### GetCertificateAttributeOk

`func (o *LDAPConnectorResponse) GetCertificateAttributeOk() (*string, bool)`

GetCertificateAttributeOk returns a tuple with the CertificateAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAttribute

`func (o *LDAPConnectorResponse) SetCertificateAttribute(v string)`

SetCertificateAttribute sets CertificateAttribute field to given value.


### GetThrottleDuration

`func (o *LDAPConnectorResponse) GetThrottleDuration() string`

GetThrottleDuration returns the ThrottleDuration field if non-nil, zero value otherwise.

### GetThrottleDurationOk

`func (o *LDAPConnectorResponse) GetThrottleDurationOk() (*string, bool)`

GetThrottleDurationOk returns a tuple with the ThrottleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDuration

`func (o *LDAPConnectorResponse) SetThrottleDuration(v string)`

SetThrottleDuration sets ThrottleDuration field to given value.


### GetThrottleParallelism

`func (o *LDAPConnectorResponse) GetThrottleParallelism() int64`

GetThrottleParallelism returns the ThrottleParallelism field if non-nil, zero value otherwise.

### GetThrottleParallelismOk

`func (o *LDAPConnectorResponse) GetThrottleParallelismOk() (*int64, bool)`

GetThrottleParallelismOk returns a tuple with the ThrottleParallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleParallelism

`func (o *LDAPConnectorResponse) SetThrottleParallelism(v int64)`

SetThrottleParallelism sets ThrottleParallelism field to given value.

### HasThrottleParallelism

`func (o *LDAPConnectorResponse) HasThrottleParallelism() bool`

HasThrottleParallelism returns a boolean if a field has been set.

### GetTimeout

`func (o *LDAPConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *LDAPConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *LDAPConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *LDAPConnectorResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *LDAPConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *LDAPConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *LDAPConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *LDAPConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *LDAPConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *LDAPConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *LDAPConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *LDAPConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetCredentials

`func (o *LDAPConnectorResponse) GetCredentials() string`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *LDAPConnectorResponse) GetCredentialsOk() (*string, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *LDAPConnectorResponse) SetCredentials(v string)`

SetCredentials sets Credentials field to given value.


### GetMaxStoredCertificatePerHolder

`func (o *LDAPConnectorResponse) GetMaxStoredCertificatePerHolder() int64`

GetMaxStoredCertificatePerHolder returns the MaxStoredCertificatePerHolder field if non-nil, zero value otherwise.

### GetMaxStoredCertificatePerHolderOk

`func (o *LDAPConnectorResponse) GetMaxStoredCertificatePerHolderOk() (*int64, bool)`

GetMaxStoredCertificatePerHolderOk returns a tuple with the MaxStoredCertificatePerHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStoredCertificatePerHolder

`func (o *LDAPConnectorResponse) SetMaxStoredCertificatePerHolder(v int64)`

SetMaxStoredCertificatePerHolder sets MaxStoredCertificatePerHolder field to given value.

### HasMaxStoredCertificatePerHolder

`func (o *LDAPConnectorResponse) HasMaxStoredCertificatePerHolder() bool`

HasMaxStoredCertificatePerHolder returns a boolean if a field has been set.

### SetMaxStoredCertificatePerHolderNil

`func (o *LDAPConnectorResponse) SetMaxStoredCertificatePerHolderNil(b bool)`

 SetMaxStoredCertificatePerHolderNil sets the value for MaxStoredCertificatePerHolder to be an explicit nil

### UnsetMaxStoredCertificatePerHolder
`func (o *LDAPConnectorResponse) UnsetMaxStoredCertificatePerHolder()`

UnsetMaxStoredCertificatePerHolder ensures that no value is present for MaxStoredCertificatePerHolder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


