# LicenseInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildTime** | **int64** |  | 
**Count** | **int64** | Current certificate count | 
**DcvCount** | **int64** | Current DCV count | 
**DcvLimit** | Pointer to **NullableInt64** | DCV license limit | [optional] 
**Expiration** | Pointer to **NullableInt64** |  | [optional] 
**IsValid** | **bool** |  | 
**Libraries** | [**[]LibraryInfoResponse**](LibraryInfoResponse.md) |  | 
**Limit** | Pointer to **NullableInt64** | Certificate license limit | [optional] 
**Modules** | [**[]ModuleLicenseInfoResponse**](ModuleLicenseInfoResponse.md) |  | 
**ReleaseChannel** | Pointer to **string** |  | [optional] 
**Version** | **string** |  | 

## Methods

### NewLicenseInfoResponse

`func NewLicenseInfoResponse(buildTime int64, count int64, dcvCount int64, isValid bool, libraries []LibraryInfoResponse, modules []ModuleLicenseInfoResponse, version string, ) *LicenseInfoResponse`

NewLicenseInfoResponse instantiates a new LicenseInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseInfoResponseWithDefaults

`func NewLicenseInfoResponseWithDefaults() *LicenseInfoResponse`

NewLicenseInfoResponseWithDefaults instantiates a new LicenseInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildTime

`func (o *LicenseInfoResponse) GetBuildTime() int64`

GetBuildTime returns the BuildTime field if non-nil, zero value otherwise.

### GetBuildTimeOk

`func (o *LicenseInfoResponse) GetBuildTimeOk() (*int64, bool)`

GetBuildTimeOk returns a tuple with the BuildTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildTime

`func (o *LicenseInfoResponse) SetBuildTime(v int64)`

SetBuildTime sets BuildTime field to given value.


### GetCount

`func (o *LicenseInfoResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LicenseInfoResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LicenseInfoResponse) SetCount(v int64)`

SetCount sets Count field to given value.


### GetDcvCount

`func (o *LicenseInfoResponse) GetDcvCount() int64`

GetDcvCount returns the DcvCount field if non-nil, zero value otherwise.

### GetDcvCountOk

`func (o *LicenseInfoResponse) GetDcvCountOk() (*int64, bool)`

GetDcvCountOk returns a tuple with the DcvCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcvCount

`func (o *LicenseInfoResponse) SetDcvCount(v int64)`

SetDcvCount sets DcvCount field to given value.


### GetDcvLimit

`func (o *LicenseInfoResponse) GetDcvLimit() int64`

GetDcvLimit returns the DcvLimit field if non-nil, zero value otherwise.

### GetDcvLimitOk

`func (o *LicenseInfoResponse) GetDcvLimitOk() (*int64, bool)`

GetDcvLimitOk returns a tuple with the DcvLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcvLimit

`func (o *LicenseInfoResponse) SetDcvLimit(v int64)`

SetDcvLimit sets DcvLimit field to given value.

### HasDcvLimit

`func (o *LicenseInfoResponse) HasDcvLimit() bool`

HasDcvLimit returns a boolean if a field has been set.

### SetDcvLimitNil

`func (o *LicenseInfoResponse) SetDcvLimitNil(b bool)`

 SetDcvLimitNil sets the value for DcvLimit to be an explicit nil

### UnsetDcvLimit
`func (o *LicenseInfoResponse) UnsetDcvLimit()`

UnsetDcvLimit ensures that no value is present for DcvLimit, not even an explicit nil
### GetExpiration

`func (o *LicenseInfoResponse) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *LicenseInfoResponse) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *LicenseInfoResponse) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *LicenseInfoResponse) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *LicenseInfoResponse) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *LicenseInfoResponse) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetIsValid

`func (o *LicenseInfoResponse) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *LicenseInfoResponse) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *LicenseInfoResponse) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.


### GetLibraries

`func (o *LicenseInfoResponse) GetLibraries() []LibraryInfoResponse`

GetLibraries returns the Libraries field if non-nil, zero value otherwise.

### GetLibrariesOk

`func (o *LicenseInfoResponse) GetLibrariesOk() (*[]LibraryInfoResponse, bool)`

GetLibrariesOk returns a tuple with the Libraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraries

`func (o *LicenseInfoResponse) SetLibraries(v []LibraryInfoResponse)`

SetLibraries sets Libraries field to given value.


### GetLimit

`func (o *LicenseInfoResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LicenseInfoResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LicenseInfoResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *LicenseInfoResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *LicenseInfoResponse) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *LicenseInfoResponse) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetModules

`func (o *LicenseInfoResponse) GetModules() []ModuleLicenseInfoResponse`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *LicenseInfoResponse) GetModulesOk() (*[]ModuleLicenseInfoResponse, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *LicenseInfoResponse) SetModules(v []ModuleLicenseInfoResponse)`

SetModules sets Modules field to given value.


### GetReleaseChannel

`func (o *LicenseInfoResponse) GetReleaseChannel() string`

GetReleaseChannel returns the ReleaseChannel field if non-nil, zero value otherwise.

### GetReleaseChannelOk

`func (o *LicenseInfoResponse) GetReleaseChannelOk() (*string, bool)`

GetReleaseChannelOk returns a tuple with the ReleaseChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseChannel

`func (o *LicenseInfoResponse) SetReleaseChannel(v string)`

SetReleaseChannel sets ReleaseChannel field to given value.

### HasReleaseChannel

`func (o *LicenseInfoResponse) HasReleaseChannel() bool`

HasReleaseChannel returns a boolean if a field has been set.

### GetVersion

`func (o *LicenseInfoResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LicenseInfoResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LicenseInfoResponse) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


