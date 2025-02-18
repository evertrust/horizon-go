# DiscoveryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | **string** | The discovery campaign&#39;s name | 
**LastDiscoveryDate** | **int64** | When this certificate was discovered for the last time | 
**Identifier** | Pointer to **NullableString** | Identifier of the user that discovered this certificate | [optional] 

## Methods

### NewDiscoveryInfo

`func NewDiscoveryInfo(campaign string, lastDiscoveryDate int64, ) *DiscoveryInfo`

NewDiscoveryInfo instantiates a new DiscoveryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryInfoWithDefaults

`func NewDiscoveryInfoWithDefaults() *DiscoveryInfo`

NewDiscoveryInfoWithDefaults instantiates a new DiscoveryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaign

`func (o *DiscoveryInfo) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *DiscoveryInfo) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *DiscoveryInfo) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.


### GetLastDiscoveryDate

`func (o *DiscoveryInfo) GetLastDiscoveryDate() int64`

GetLastDiscoveryDate returns the LastDiscoveryDate field if non-nil, zero value otherwise.

### GetLastDiscoveryDateOk

`func (o *DiscoveryInfo) GetLastDiscoveryDateOk() (*int64, bool)`

GetLastDiscoveryDateOk returns a tuple with the LastDiscoveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDiscoveryDate

`func (o *DiscoveryInfo) SetLastDiscoveryDate(v int64)`

SetLastDiscoveryDate sets LastDiscoveryDate field to given value.


### GetIdentifier

`func (o *DiscoveryInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DiscoveryInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DiscoveryInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *DiscoveryInfo) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *DiscoveryInfo) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *DiscoveryInfo) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


