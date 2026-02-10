# StaticDomainDictionaryProviderDomainsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Top level domain name | 
**Dictionary** | [**[]MapEntry**](MapEntry.md) |  | 

## Methods

### NewStaticDomainDictionaryProviderDomainsInner

`func NewStaticDomainDictionaryProviderDomainsInner(domain string, dictionary []MapEntry, ) *StaticDomainDictionaryProviderDomainsInner`

NewStaticDomainDictionaryProviderDomainsInner instantiates a new StaticDomainDictionaryProviderDomainsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticDomainDictionaryProviderDomainsInnerWithDefaults

`func NewStaticDomainDictionaryProviderDomainsInnerWithDefaults() *StaticDomainDictionaryProviderDomainsInner`

NewStaticDomainDictionaryProviderDomainsInnerWithDefaults instantiates a new StaticDomainDictionaryProviderDomainsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *StaticDomainDictionaryProviderDomainsInner) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *StaticDomainDictionaryProviderDomainsInner) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *StaticDomainDictionaryProviderDomainsInner) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDictionary

`func (o *StaticDomainDictionaryProviderDomainsInner) GetDictionary() []MapEntry`

GetDictionary returns the Dictionary field if non-nil, zero value otherwise.

### GetDictionaryOk

`func (o *StaticDomainDictionaryProviderDomainsInner) GetDictionaryOk() (*[]MapEntry, bool)`

GetDictionaryOk returns a tuple with the Dictionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionary

`func (o *StaticDomainDictionaryProviderDomainsInner) SetDictionary(v []MapEntry)`

SetDictionary sets Dictionary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


