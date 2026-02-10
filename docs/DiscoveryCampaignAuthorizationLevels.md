# DiscoveryCampaignAuthorizationLevels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feed** | [**AuthorizationLevel**](AuthorizationLevel.md) | The authorization level required to feed certificates into this discovery campaign | 
**Search** | [**AuthorizationLevel**](AuthorizationLevel.md) | The authorization level required to search the discovered certificates of this campaign | 

## Methods

### NewDiscoveryCampaignAuthorizationLevels

`func NewDiscoveryCampaignAuthorizationLevels(feed AuthorizationLevel, search AuthorizationLevel, ) *DiscoveryCampaignAuthorizationLevels`

NewDiscoveryCampaignAuthorizationLevels instantiates a new DiscoveryCampaignAuthorizationLevels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryCampaignAuthorizationLevelsWithDefaults

`func NewDiscoveryCampaignAuthorizationLevelsWithDefaults() *DiscoveryCampaignAuthorizationLevels`

NewDiscoveryCampaignAuthorizationLevelsWithDefaults instantiates a new DiscoveryCampaignAuthorizationLevels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeed

`func (o *DiscoveryCampaignAuthorizationLevels) GetFeed() AuthorizationLevel`

GetFeed returns the Feed field if non-nil, zero value otherwise.

### GetFeedOk

`func (o *DiscoveryCampaignAuthorizationLevels) GetFeedOk() (*AuthorizationLevel, bool)`

GetFeedOk returns a tuple with the Feed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeed

`func (o *DiscoveryCampaignAuthorizationLevels) SetFeed(v AuthorizationLevel)`

SetFeed sets Feed field to given value.


### GetSearch

`func (o *DiscoveryCampaignAuthorizationLevels) GetSearch() AuthorizationLevel`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *DiscoveryCampaignAuthorizationLevels) GetSearchOk() (*AuthorizationLevel, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *DiscoveryCampaignAuthorizationLevels) SetSearch(v AuthorizationLevel)`

SetSearch sets Search field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


