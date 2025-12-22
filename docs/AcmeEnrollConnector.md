# AcmeEnrollConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountEmail** | Pointer to **NullableString** | Email to associate with the account | [optional] 
**AccountKeyType** | **string** | The key type to use to generate the account key | 
**DnsChallengeProvider** | [**DnsChallengeProviders**](DnsChallengeProviders.md) | DNS Provider configuration to provision the DNS challenge | 
**DomainDictionaryProvider** | Pointer to [**NullableDomainDictionaryProviders**](DomainDictionaryProviders.md) | The dictionary provider | [optional] 
**Eab** | Pointer to **NullableString** | &#x60;password&#x60; credentials name to use for External Account Binding | [optional] 
**EndPoint** | **string** | The directory url of the ACME endpoint | 
**Name** | **string** |  | 
**Proxy** | Pointer to **NullableString** |  | [optional] 
**Queue** | Pointer to **NullableString** |  | [optional] 
**RotateAccount** | Pointer to **NullableBool** | If enabled, regenerate the account (does not need to be specified on creation) | [optional] 
**Timeout** | **NullableString** |  | 
**Type** | **string** |  | 

## Methods

### NewAcmeEnrollConnector

`func NewAcmeEnrollConnector(accountKeyType string, dnsChallengeProvider DnsChallengeProviders, endPoint string, name string, timeout NullableString, type_ string, ) *AcmeEnrollConnector`

NewAcmeEnrollConnector instantiates a new AcmeEnrollConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeEnrollConnectorWithDefaults

`func NewAcmeEnrollConnectorWithDefaults() *AcmeEnrollConnector`

NewAcmeEnrollConnectorWithDefaults instantiates a new AcmeEnrollConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountEmail

`func (o *AcmeEnrollConnector) GetAccountEmail() string`

GetAccountEmail returns the AccountEmail field if non-nil, zero value otherwise.

### GetAccountEmailOk

`func (o *AcmeEnrollConnector) GetAccountEmailOk() (*string, bool)`

GetAccountEmailOk returns a tuple with the AccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEmail

`func (o *AcmeEnrollConnector) SetAccountEmail(v string)`

SetAccountEmail sets AccountEmail field to given value.

### HasAccountEmail

`func (o *AcmeEnrollConnector) HasAccountEmail() bool`

HasAccountEmail returns a boolean if a field has been set.

### SetAccountEmailNil

`func (o *AcmeEnrollConnector) SetAccountEmailNil(b bool)`

 SetAccountEmailNil sets the value for AccountEmail to be an explicit nil

### UnsetAccountEmail
`func (o *AcmeEnrollConnector) UnsetAccountEmail()`

UnsetAccountEmail ensures that no value is present for AccountEmail, not even an explicit nil
### GetAccountKeyType

`func (o *AcmeEnrollConnector) GetAccountKeyType() string`

GetAccountKeyType returns the AccountKeyType field if non-nil, zero value otherwise.

### GetAccountKeyTypeOk

`func (o *AcmeEnrollConnector) GetAccountKeyTypeOk() (*string, bool)`

GetAccountKeyTypeOk returns a tuple with the AccountKeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKeyType

`func (o *AcmeEnrollConnector) SetAccountKeyType(v string)`

SetAccountKeyType sets AccountKeyType field to given value.


### GetDnsChallengeProvider

`func (o *AcmeEnrollConnector) GetDnsChallengeProvider() DnsChallengeProviders`

GetDnsChallengeProvider returns the DnsChallengeProvider field if non-nil, zero value otherwise.

### GetDnsChallengeProviderOk

`func (o *AcmeEnrollConnector) GetDnsChallengeProviderOk() (*DnsChallengeProviders, bool)`

GetDnsChallengeProviderOk returns a tuple with the DnsChallengeProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsChallengeProvider

`func (o *AcmeEnrollConnector) SetDnsChallengeProvider(v DnsChallengeProviders)`

SetDnsChallengeProvider sets DnsChallengeProvider field to given value.


### GetDomainDictionaryProvider

`func (o *AcmeEnrollConnector) GetDomainDictionaryProvider() DomainDictionaryProviders`

GetDomainDictionaryProvider returns the DomainDictionaryProvider field if non-nil, zero value otherwise.

### GetDomainDictionaryProviderOk

`func (o *AcmeEnrollConnector) GetDomainDictionaryProviderOk() (*DomainDictionaryProviders, bool)`

GetDomainDictionaryProviderOk returns a tuple with the DomainDictionaryProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDictionaryProvider

`func (o *AcmeEnrollConnector) SetDomainDictionaryProvider(v DomainDictionaryProviders)`

SetDomainDictionaryProvider sets DomainDictionaryProvider field to given value.

### HasDomainDictionaryProvider

`func (o *AcmeEnrollConnector) HasDomainDictionaryProvider() bool`

HasDomainDictionaryProvider returns a boolean if a field has been set.

### SetDomainDictionaryProviderNil

`func (o *AcmeEnrollConnector) SetDomainDictionaryProviderNil(b bool)`

 SetDomainDictionaryProviderNil sets the value for DomainDictionaryProvider to be an explicit nil

### UnsetDomainDictionaryProvider
`func (o *AcmeEnrollConnector) UnsetDomainDictionaryProvider()`

UnsetDomainDictionaryProvider ensures that no value is present for DomainDictionaryProvider, not even an explicit nil
### GetEab

`func (o *AcmeEnrollConnector) GetEab() string`

GetEab returns the Eab field if non-nil, zero value otherwise.

### GetEabOk

`func (o *AcmeEnrollConnector) GetEabOk() (*string, bool)`

GetEabOk returns a tuple with the Eab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEab

`func (o *AcmeEnrollConnector) SetEab(v string)`

SetEab sets Eab field to given value.

### HasEab

`func (o *AcmeEnrollConnector) HasEab() bool`

HasEab returns a boolean if a field has been set.

### SetEabNil

`func (o *AcmeEnrollConnector) SetEabNil(b bool)`

 SetEabNil sets the value for Eab to be an explicit nil

### UnsetEab
`func (o *AcmeEnrollConnector) UnsetEab()`

UnsetEab ensures that no value is present for Eab, not even an explicit nil
### GetEndPoint

`func (o *AcmeEnrollConnector) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *AcmeEnrollConnector) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *AcmeEnrollConnector) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.


### GetName

`func (o *AcmeEnrollConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AcmeEnrollConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AcmeEnrollConnector) SetName(v string)`

SetName sets Name field to given value.


### GetProxy

`func (o *AcmeEnrollConnector) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *AcmeEnrollConnector) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *AcmeEnrollConnector) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *AcmeEnrollConnector) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### SetProxyNil

`func (o *AcmeEnrollConnector) SetProxyNil(b bool)`

 SetProxyNil sets the value for Proxy to be an explicit nil

### UnsetProxy
`func (o *AcmeEnrollConnector) UnsetProxy()`

UnsetProxy ensures that no value is present for Proxy, not even an explicit nil
### GetQueue

`func (o *AcmeEnrollConnector) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *AcmeEnrollConnector) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *AcmeEnrollConnector) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *AcmeEnrollConnector) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### SetQueueNil

`func (o *AcmeEnrollConnector) SetQueueNil(b bool)`

 SetQueueNil sets the value for Queue to be an explicit nil

### UnsetQueue
`func (o *AcmeEnrollConnector) UnsetQueue()`

UnsetQueue ensures that no value is present for Queue, not even an explicit nil
### GetRotateAccount

`func (o *AcmeEnrollConnector) GetRotateAccount() bool`

GetRotateAccount returns the RotateAccount field if non-nil, zero value otherwise.

### GetRotateAccountOk

`func (o *AcmeEnrollConnector) GetRotateAccountOk() (*bool, bool)`

GetRotateAccountOk returns a tuple with the RotateAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAccount

`func (o *AcmeEnrollConnector) SetRotateAccount(v bool)`

SetRotateAccount sets RotateAccount field to given value.

### HasRotateAccount

`func (o *AcmeEnrollConnector) HasRotateAccount() bool`

HasRotateAccount returns a boolean if a field has been set.

### SetRotateAccountNil

`func (o *AcmeEnrollConnector) SetRotateAccountNil(b bool)`

 SetRotateAccountNil sets the value for RotateAccount to be an explicit nil

### UnsetRotateAccount
`func (o *AcmeEnrollConnector) UnsetRotateAccount()`

UnsetRotateAccount ensures that no value is present for RotateAccount, not even an explicit nil
### GetTimeout

`func (o *AcmeEnrollConnector) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AcmeEnrollConnector) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AcmeEnrollConnector) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.


### SetTimeoutNil

`func (o *AcmeEnrollConnector) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *AcmeEnrollConnector) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetType

`func (o *AcmeEnrollConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AcmeEnrollConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AcmeEnrollConnector) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


