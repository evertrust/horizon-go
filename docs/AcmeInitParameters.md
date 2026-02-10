# AcmeInitParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** | The module of the initialization parameters. | 
**Profile** | **string** | The profile used for ACME. | 
**KeyType** | Pointer to **string** | The key type used for ACME. | [optional] 
**TlsAlpn01Port** | Pointer to **int64** | The TLS-ALPN-01 port for ACME. | [optional] 
**Http01Port** | Pointer to **int64** | The HTTP-01 port for ACME. | [optional] 

## Methods

### NewAcmeInitParameters

`func NewAcmeInitParameters(module string, profile string, ) *AcmeInitParameters`

NewAcmeInitParameters instantiates a new AcmeInitParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeInitParametersWithDefaults

`func NewAcmeInitParametersWithDefaults() *AcmeInitParameters`

NewAcmeInitParametersWithDefaults instantiates a new AcmeInitParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *AcmeInitParameters) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *AcmeInitParameters) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *AcmeInitParameters) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *AcmeInitParameters) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AcmeInitParameters) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AcmeInitParameters) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetKeyType

`func (o *AcmeInitParameters) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *AcmeInitParameters) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *AcmeInitParameters) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *AcmeInitParameters) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetTlsAlpn01Port

`func (o *AcmeInitParameters) GetTlsAlpn01Port() int64`

GetTlsAlpn01Port returns the TlsAlpn01Port field if non-nil, zero value otherwise.

### GetTlsAlpn01PortOk

`func (o *AcmeInitParameters) GetTlsAlpn01PortOk() (*int64, bool)`

GetTlsAlpn01PortOk returns a tuple with the TlsAlpn01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsAlpn01Port

`func (o *AcmeInitParameters) SetTlsAlpn01Port(v int64)`

SetTlsAlpn01Port sets TlsAlpn01Port field to given value.

### HasTlsAlpn01Port

`func (o *AcmeInitParameters) HasTlsAlpn01Port() bool`

HasTlsAlpn01Port returns a boolean if a field has been set.

### GetHttp01Port

`func (o *AcmeInitParameters) GetHttp01Port() int64`

GetHttp01Port returns the Http01Port field if non-nil, zero value otherwise.

### GetHttp01PortOk

`func (o *AcmeInitParameters) GetHttp01PortOk() (*int64, bool)`

GetHttp01PortOk returns a tuple with the Http01Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp01Port

`func (o *AcmeInitParameters) SetHttp01Port(v int64)`

SetHttp01Port sets Http01Port field to given value.

### HasHttp01Port

`func (o *AcmeInitParameters) HasHttp01Port() bool`

HasHttp01Port returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


