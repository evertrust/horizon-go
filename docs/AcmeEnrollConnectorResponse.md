# AcmeEnrollConnectorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**Name** | **string** |  | 
**Type** | **string** |  | 
**EndPoint** | **string** | The directory url of the ACME endpoint | 
**Timeout** | **NullableString** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**Eab** | Pointer to **NullableString** | &#x60;password&#x60; credentials name to use for External Account Binding | [optional] 
**AccountKeyType** | **string** | The key type to use to generate the account key | 
**AccountEmail** | Pointer to **NullableString** | Email to associate with the account | [optional] 
**RotateAccount** | Pointer to **NullableBool** | If enable, regenerate the account (does not need to be specified on creation) | [optional] 
**Status** | Pointer to [**NullablePKIConnectorStatus**](PKIConnectorStatus.md) |  | [optional] 
**AccountUrl** | **string** | Url of the account on the ACME directory | 
**DomainDictionaryProvider** | Pointer to [**NullableDomainDictionaryProviders**](DomainDictionaryProviders.md) | The dictionary provider | [optional] 
**DnsChallengeProvider** | [**DnsChallengeProviders**](DnsChallengeProviders.md) | DNS Provider configuration to provision the DNS challenge. Available from &#x60;2.7.7&#x60; | 

## Methods

### NewAcmeEnrollConnectorResponse

`func NewAcmeEnrollConnectorResponse(id string, name string, type_ string, endPoint string, timeout NullableString, accountKeyType string, accountUrl string, dnsChallengeProvider DnsChallengeProviders, ) *AcmeEnrollConnectorResponse`

NewAcmeEnrollConnectorResponse instantiates a new AcmeEnrollConnectorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeEnrollConnectorResponseWithDefaults

`func NewAcmeEnrollConnectorResponseWithDefaults() *AcmeEnrollConnectorResponse`

NewAcmeEnrollConnectorResponseWithDefaults instantiates a new AcmeEnrollConnectorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AcmeEnrollConnectorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AcmeEnrollConnectorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AcmeEnrollConnectorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AcmeEnrollConnectorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcmeEnrollConnectorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcmeEnrollConnectorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AcmeEnrollConnectorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcmeEnrollConnectorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcmeEnrollConnectorResponse) SetType(v string)`

SetType sets Type field to given value.


### GetEndPoint

`func (o *AcmeEnrollConnectorResponse) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *AcmeEnrollConnectorResponse) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *AcmeEnrollConnectorResponse) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetTimeout

`func (o *AcmeEnrollConnectorResponse) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AcmeEnrollConnectorResponse) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AcmeEnrollConnectorResponse) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *AcmeEnrollConnectorResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *AcmeEnrollConnectorResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetProxy

`func (o *AcmeEnrollConnectorResponse) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AcmeEnrollConnectorResponse) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AcmeEnrollConnectorResponse) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AcmeEnrollConnectorResponse) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *AcmeEnrollConnectorResponse) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *AcmeEnrollConnectorResponse) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *AcmeEnrollConnectorResponse) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *AcmeEnrollConnectorResponse) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *AcmeEnrollConnectorResponse) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *AcmeEnrollConnectorResponse) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *AcmeEnrollConnectorResponse) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *AcmeEnrollConnectorResponse) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetEab

`func (o *AcmeEnrollConnectorResponse) GetEab() string`

GetEab returns the Eab field if non-nil, zero value otherwise.

### GetEabOk

`func (o *AcmeEnrollConnectorResponse) GetEabOk() (*string, bool)`

GetEabOk returns a tuple with the Eab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEab

`func (o *AcmeEnrollConnectorResponse) SetEab(v string)`

SetEab sets Eab field to given value.

### HasEab

`func (o *AcmeEnrollConnectorResponse) HasEab() bool`

HasEab returns a boolean if a field has been set.

### SetEabNil

`func (o *AcmeEnrollConnectorResponse) SetEabNil(b bool)`

 SetEabNil sets the value for Eab to be an explicit nil

### UnsetEab
`func (o *AcmeEnrollConnectorResponse) UnsetEab()`

UnsetEab ensures that no value is present for Eab, not even an explicit nil
### GetAccountKeyType

`func (o *AcmeEnrollConnectorResponse) GetAccountKeyType() string`

GetAccountKeyType returns the AccountKeyType field if non-nil, zero value otherwise.

### GetAccountKeyTypeOk

`func (o *AcmeEnrollConnectorResponse) GetAccountKeyTypeOk() (*string, bool)`

GetAccountKeyTypeOk returns a tuple with the AccountKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKeyType

`func (o *AcmeEnrollConnectorResponse) SetAccountKeyType(v string)`

SetAccountKeyType sets AccountKeyType field to given value.


### GetAccountEmail

`func (o *AcmeEnrollConnectorResponse) GetAccountEmail() string`

GetAccountEmail returns the AccountEmail field if non-nil, zero value otherwise.

### GetAccountEmailOk

`func (o *AcmeEnrollConnectorResponse) GetAccountEmailOk() (*string, bool)`

GetAccountEmailOk returns a tuple with the AccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEmail

`func (o *AcmeEnrollConnectorResponse) SetAccountEmail(v string)`

SetAccountEmail sets AccountEmail field to given value.

### HasAccountEmail

`func (o *AcmeEnrollConnectorResponse) HasAccountEmail() bool`

HasAccountEmail returns a boolean if a field has been set.

### SetAccountEmailNil

`func (o *AcmeEnrollConnectorResponse) SetAccountEmailNil(b bool)`

 SetAccountEmailNil sets the value for AccountEmail to be an explicit nil

### UnsetAccountEmail
`func (o *AcmeEnrollConnectorResponse) UnsetAccountEmail()`

UnsetAccountEmail ensures that no value is present for AccountEmail, not even an explicit nil
### GetRotateAccount

`func (o *AcmeEnrollConnectorResponse) GetRotateAccount() bool`

GetRotateAccount returns the RotateAccount field if non-nil, zero value otherwise.

### GetRotateAccountOk

`func (o *AcmeEnrollConnectorResponse) GetRotateAccountOk() (*bool, bool)`

GetRotateAccountOk returns a tuple with the RotateAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAccount

`func (o *AcmeEnrollConnectorResponse) SetRotateAccount(v bool)`

SetRotateAccount sets RotateAccount field to given value.

### HasRotateAccount

`func (o *AcmeEnrollConnectorResponse) HasRotateAccount() bool`

HasRotateAccount returns a boolean if a field has been set.

### SetRotateAccountNil

`func (o *AcmeEnrollConnectorResponse) SetRotateAccountNil(b bool)`

 SetRotateAccountNil sets the value for RotateAccount to be an explicit nil

### UnsetRotateAccount
`func (o *AcmeEnrollConnectorResponse) UnsetRotateAccount()`

UnsetRotateAccount ensures that no value is present for RotateAccount, not even an explicit nil
### GetStatus

`func (o *AcmeEnrollConnectorResponse) GetStatus() PKIConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AcmeEnrollConnectorResponse) GetStatusOk() (*PKIConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AcmeEnrollConnectorResponse) SetStatus(v PKIConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AcmeEnrollConnectorResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AcmeEnrollConnectorResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AcmeEnrollConnectorResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAccountUrl

`func (o *AcmeEnrollConnectorResponse) GetAccountUrl() string`

GetAccountUrl returns the AccountUrl field if non-nil, zero value otherwise.

### GetAccountUrlOk

`func (o *AcmeEnrollConnectorResponse) GetAccountUrlOk() (*string, bool)`

GetAccountUrlOk returns a tuple with the AccountUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUrl

`func (o *AcmeEnrollConnectorResponse) SetAccountUrl(v string)`

SetAccountUrl sets AccountUrl field to given value.


### GetDomainDictionaryProvider

`func (o *AcmeEnrollConnectorResponse) GetDomainDictionaryProvider() DomainDictionaryProviders`

GetDomainDictionaryProvider returns the DomainDictionaryProvider field if non-nil, zero value otherwise.

### GetDomainDictionaryProviderOk

`func (o *AcmeEnrollConnectorResponse) GetDomainDictionaryProviderOk() (*DomainDictionaryProviders, bool)`

GetDomainDictionaryProviderOk returns a tuple with the DomainDictionaryProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDictionaryProvider

`func (o *AcmeEnrollConnectorResponse) SetDomainDictionaryProvider(v DomainDictionaryProviders)`

SetDomainDictionaryProvider sets DomainDictionaryProvider field to given value.

### HasDomainDictionaryProvider

`func (o *AcmeEnrollConnectorResponse) HasDomainDictionaryProvider() bool`

HasDomainDictionaryProvider returns a boolean if a field has been set.

### SetDomainDictionaryProviderNil

`func (o *AcmeEnrollConnectorResponse) SetDomainDictionaryProviderNil(b bool)`

 SetDomainDictionaryProviderNil sets the value for DomainDictionaryProvider to be an explicit nil

### UnsetDomainDictionaryProvider
`func (o *AcmeEnrollConnectorResponse) UnsetDomainDictionaryProvider()`

UnsetDomainDictionaryProvider ensures that no value is present for DomainDictionaryProvider, not even an explicit nil
### GetDnsChallengeProvider

`func (o *AcmeEnrollConnectorResponse) GetDnsChallengeProvider() DnsChallengeProviders`

GetDnsChallengeProvider returns the DnsChallengeProvider field if non-nil, zero value otherwise.

### GetDnsChallengeProviderOk

`func (o *AcmeEnrollConnectorResponse) GetDnsChallengeProviderOk() (*DnsChallengeProviders, bool)`

GetDnsChallengeProviderOk returns a tuple with the DnsChallengeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsChallengeProvider

`func (o *AcmeEnrollConnectorResponse) SetDnsChallengeProvider(v DnsChallengeProviders)`

SetDnsChallengeProvider sets DnsChallengeProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


