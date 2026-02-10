# StaticDomainDictionaryProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Domains** | [**[]StaticDomainDictionaryProviderDomainsInner**](StaticDomainDictionaryProviderDomainsInner.md) | The domain dictionaries | 

## Methods

### NewStaticDomainDictionaryProvider

`func NewStaticDomainDictionaryProvider(type_ string, domains []StaticDomainDictionaryProviderDomainsInner, ) *StaticDomainDictionaryProvider`

NewStaticDomainDictionaryProvider instantiates a new StaticDomainDictionaryProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticDomainDictionaryProviderWithDefaults

`func NewStaticDomainDictionaryProviderWithDefaults() *StaticDomainDictionaryProvider`

NewStaticDomainDictionaryProviderWithDefaults instantiates a new StaticDomainDictionaryProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StaticDomainDictionaryProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StaticDomainDictionaryProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StaticDomainDictionaryProvider) SetType(v string)`

SetType sets Type field to given value.


### GetDomains

`func (o *StaticDomainDictionaryProvider) GetDomains() []StaticDomainDictionaryProviderDomainsInner`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *StaticDomainDictionaryProvider) GetDomainsOk() (*[]StaticDomainDictionaryProviderDomainsInner, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *StaticDomainDictionaryProvider) SetDomains(v []StaticDomainDictionaryProviderDomainsInner)`

SetDomains sets Domains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


