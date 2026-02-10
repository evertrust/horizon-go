# DNSDatasource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of datasource | 
**Name** | **string** | Name of the datasource | 
**DisplayName** | Pointer to [**[]LocalizedString**](LocalizedString.md) | The localized name of the datasource | [optional] 
**Description** | Pointer to **string** | Description of the datasource | [optional] 
**Host** | Pointer to **NullableString** | Ip of the DNS server. If empty, Horizon Server DNS is used | [optional] 
**Port** | Pointer to **NullableInt64** | Port on which to join the DNS server | [optional] [default to 53]
**Timeout** | Pointer to **NullableString** | Timeout for the DNS request | [optional] [default to "10 seconds"]
**RecordTypes** | Pointer to **[]string** | Type of DNS records to fetch. All available record types are fetched if null | [optional] 
**Lookup** | **string** | Host to lookup | 

## Methods

### NewDNSDatasource

`func NewDNSDatasource(type_ string, name string, lookup string, ) *DNSDatasource`

NewDNSDatasource instantiates a new DNSDatasource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSDatasourceWithDefaults

`func NewDNSDatasourceWithDefaults() *DNSDatasource`

NewDNSDatasourceWithDefaults instantiates a new DNSDatasource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DNSDatasource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSDatasource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSDatasource) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *DNSDatasource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSDatasource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSDatasource) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *DNSDatasource) GetDisplayName() []LocalizedString`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DNSDatasource) GetDisplayNameOk() (*[]LocalizedString, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DNSDatasource) SetDisplayName(v []LocalizedString)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DNSDatasource) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *DNSDatasource) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *DNSDatasource) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *DNSDatasource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DNSDatasource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DNSDatasource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DNSDatasource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHost

`func (o *DNSDatasource) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DNSDatasource) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DNSDatasource) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DNSDatasource) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *DNSDatasource) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *DNSDatasource) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *DNSDatasource) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DNSDatasource) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DNSDatasource) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DNSDatasource) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *DNSDatasource) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *DNSDatasource) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetTimeout

`func (o *DNSDatasource) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DNSDatasource) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DNSDatasource) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DNSDatasource) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *DNSDatasource) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *DNSDatasource) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetRecordTypes

`func (o *DNSDatasource) GetRecordTypes() []string`

GetRecordTypes returns the RecordTypes field if non-nil, zero value otherwise.

### GetRecordTypesOk

`func (o *DNSDatasource) GetRecordTypesOk() (*[]string, bool)`

GetRecordTypesOk returns a tuple with the RecordTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordTypes

`func (o *DNSDatasource) SetRecordTypes(v []string)`

SetRecordTypes sets RecordTypes field to given value.

### HasRecordTypes

`func (o *DNSDatasource) HasRecordTypes() bool`

HasRecordTypes returns a boolean if a field has been set.

### SetRecordTypesNil

`func (o *DNSDatasource) SetRecordTypesNil(b bool)`

 SetRecordTypesNil sets the value for RecordTypes to be an explicit nil

### UnsetRecordTypes
`func (o *DNSDatasource) UnsetRecordTypes()`

UnsetRecordTypes ensures that no value is present for RecordTypes, not even an explicit nil
### GetLookup

`func (o *DNSDatasource) GetLookup() string`

GetLookup returns the Lookup field if non-nil, zero value otherwise.

### GetLookupOk

`func (o *DNSDatasource) GetLookupOk() (*string, bool)`

GetLookupOk returns a tuple with the Lookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookup

`func (o *DNSDatasource) SetLookup(v string)`

SetLookup sets Lookup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


