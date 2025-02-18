# LicenseInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsValid** | **bool** |  | 
**Expiration** | Pointer to **NullableInt64** |  | [optional] 
**Version** | **string** |  | 
**BuildTime** | **int64** |  | 
**Modules** | [**[]ModuleLicenseInfo**](ModuleLicenseInfo.md) |  | 
**Libraries** | [**[]LibraryInfo**](LibraryInfo.md) |  | 

## Methods

### NewLicenseInfoResponse

`func NewLicenseInfoResponse(isValid bool, version string, buildTime int64, modules []ModuleLicenseInfo, libraries []LibraryInfo, ) *LicenseInfoResponse`

NewLicenseInfoResponse instantiates a new LicenseInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseInfoResponseWithDefaults

`func NewLicenseInfoResponseWithDefaults() *LicenseInfoResponse`

NewLicenseInfoResponseWithDefaults instantiates a new LicenseInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetModules

`func (o *LicenseInfoResponse) GetModules() []ModuleLicenseInfo`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *LicenseInfoResponse) GetModulesOk() (*[]ModuleLicenseInfo, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *LicenseInfoResponse) SetModules(v []ModuleLicenseInfo)`

SetModules sets Modules field to given value.


### GetLibraries

`func (o *LicenseInfoResponse) GetLibraries() []LibraryInfo`

GetLibraries returns the Libraries field if non-nil, zero value otherwise.

### GetLibrariesOk

`func (o *LicenseInfoResponse) GetLibrariesOk() (*[]LibraryInfo, bool)`

GetLibrariesOk returns a tuple with the Libraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraries

`func (o *LicenseInfoResponse) SetLibraries(v []LibraryInfo)`

SetLibraries sets Libraries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


