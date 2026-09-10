# ZoneIdMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regex** | **string** | Regex for mapping | 
**ZoneId** | **string** | Cloudflare DNS zone ID where TXT records will be provisioned with regex | 

## Methods

### NewZoneIdMappings

`func NewZoneIdMappings(regex string, zoneId string, ) *ZoneIdMappings`

NewZoneIdMappings instantiates a new ZoneIdMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneIdMappingsWithDefaults

`func NewZoneIdMappingsWithDefaults() *ZoneIdMappings`

NewZoneIdMappingsWithDefaults instantiates a new ZoneIdMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegex

`func (o *ZoneIdMappings) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *ZoneIdMappings) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *ZoneIdMappings) SetRegex(v string)`

SetRegex sets Regex field to given value.


### GetZoneId

`func (o *ZoneIdMappings) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ZoneIdMappings) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ZoneIdMappings) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


