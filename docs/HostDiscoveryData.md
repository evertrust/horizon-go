# HostDiscoveryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **NullableString** | The certificate&#39;s host ip | [optional] 
**Sources** | Pointer to **[]string** | Information on the type of discovery that discovered this certificate | [optional] 
**Hostnames** | Pointer to **[]string** | The certificate&#39;s host hostnames (netscan only) | [optional] 
**OperatingSystems** | Pointer to **[]string** | The certificate&#39;s host operating system (localscan only) | [optional] 
**Paths** | Pointer to **[]string** | The path to the certificate on the host machine (localscan only) | [optional] 
**Usages** | Pointer to **[]string** | The path of the configuration files that were used to find the certificates | [optional] 
**TlsPorts** | Pointer to [**[]TlsPort**](TlsPort.md) | The ports on which the certificate is exposed for https connexion | [optional] 

## Methods

### NewHostDiscoveryData

`func NewHostDiscoveryData() *HostDiscoveryData`

NewHostDiscoveryData instantiates a new HostDiscoveryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostDiscoveryDataWithDefaults

`func NewHostDiscoveryDataWithDefaults() *HostDiscoveryData`

NewHostDiscoveryDataWithDefaults instantiates a new HostDiscoveryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *HostDiscoveryData) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *HostDiscoveryData) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *HostDiscoveryData) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *HostDiscoveryData) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *HostDiscoveryData) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *HostDiscoveryData) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetSources

`func (o *HostDiscoveryData) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *HostDiscoveryData) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *HostDiscoveryData) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *HostDiscoveryData) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *HostDiscoveryData) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *HostDiscoveryData) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetHostnames

`func (o *HostDiscoveryData) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *HostDiscoveryData) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *HostDiscoveryData) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *HostDiscoveryData) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### SetHostnamesNil

`func (o *HostDiscoveryData) SetHostnamesNil(b bool)`

 SetHostnamesNil sets the value for Hostnames to be an explicit nil

### UnsetHostnames
`func (o *HostDiscoveryData) UnsetHostnames()`

UnsetHostnames ensures that no value is present for Hostnames, not even an explicit nil
### GetOperatingSystems

`func (o *HostDiscoveryData) GetOperatingSystems() []string`

GetOperatingSystems returns the OperatingSystems field if non-nil, zero value otherwise.

### GetOperatingSystemsOk

`func (o *HostDiscoveryData) GetOperatingSystemsOk() (*[]string, bool)`

GetOperatingSystemsOk returns a tuple with the OperatingSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystems

`func (o *HostDiscoveryData) SetOperatingSystems(v []string)`

SetOperatingSystems sets OperatingSystems field to given value.

### HasOperatingSystems

`func (o *HostDiscoveryData) HasOperatingSystems() bool`

HasOperatingSystems returns a boolean if a field has been set.

### SetOperatingSystemsNil

`func (o *HostDiscoveryData) SetOperatingSystemsNil(b bool)`

 SetOperatingSystemsNil sets the value for OperatingSystems to be an explicit nil

### UnsetOperatingSystems
`func (o *HostDiscoveryData) UnsetOperatingSystems()`

UnsetOperatingSystems ensures that no value is present for OperatingSystems, not even an explicit nil
### GetPaths

`func (o *HostDiscoveryData) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *HostDiscoveryData) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *HostDiscoveryData) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *HostDiscoveryData) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### SetPathsNil

`func (o *HostDiscoveryData) SetPathsNil(b bool)`

 SetPathsNil sets the value for Paths to be an explicit nil

### UnsetPaths
`func (o *HostDiscoveryData) UnsetPaths()`

UnsetPaths ensures that no value is present for Paths, not even an explicit nil
### GetUsages

`func (o *HostDiscoveryData) GetUsages() []string`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *HostDiscoveryData) GetUsagesOk() (*[]string, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *HostDiscoveryData) SetUsages(v []string)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *HostDiscoveryData) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### SetUsagesNil

`func (o *HostDiscoveryData) SetUsagesNil(b bool)`

 SetUsagesNil sets the value for Usages to be an explicit nil

### UnsetUsages
`func (o *HostDiscoveryData) UnsetUsages()`

UnsetUsages ensures that no value is present for Usages, not even an explicit nil
### GetTlsPorts

`func (o *HostDiscoveryData) GetTlsPorts() []TlsPort`

GetTlsPorts returns the TlsPorts field if non-nil, zero value otherwise.

### GetTlsPortsOk

`func (o *HostDiscoveryData) GetTlsPortsOk() (*[]TlsPort, bool)`

GetTlsPortsOk returns a tuple with the TlsPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsPorts

`func (o *HostDiscoveryData) SetTlsPorts(v []TlsPort)`

SetTlsPorts sets TlsPorts field to given value.

### HasTlsPorts

`func (o *HostDiscoveryData) HasTlsPorts() bool`

HasTlsPorts returns a boolean if a field has been set.

### SetTlsPortsNil

`func (o *HostDiscoveryData) SetTlsPortsNil(b bool)`

 SetTlsPortsNil sets the value for TlsPorts to be an explicit nil

### UnsetTlsPorts
`func (o *HostDiscoveryData) UnsetTlsPorts()`

UnsetTlsPorts ensures that no value is present for TlsPorts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


