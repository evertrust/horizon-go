# ScepInitParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationMode** | Pointer to **string** | The authorization mode for SCEP. | [optional] 
**CsrInfoIgnored** | Pointer to **bool** | Indicates whether CSR info is ignored for SCEP. | [optional] 
**KeyType** | Pointer to **string** | The key type used for SCEP. | [optional] 
**Module** | **string** | The module of the initialization parameters. | 
**Profile** | **string** | The profile used for SCEP. | 

## Methods

### NewScepInitParameters

`func NewScepInitParameters(module string, profile string, ) *ScepInitParameters`

NewScepInitParameters instantiates a new ScepInitParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScepInitParametersWithDefaults

`func NewScepInitParametersWithDefaults() *ScepInitParameters`

NewScepInitParametersWithDefaults instantiates a new ScepInitParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationMode

`func (o *ScepInitParameters) GetAuthorizationMode() string`

GetAuthorizationMode returns the AuthorizationMode field if non-nil, zero value otherwise.

### GetAuthorizationModeOk

`func (o *ScepInitParameters) GetAuthorizationModeOk() (*string, bool)`

GetAuthorizationModeOk returns a tuple with the AuthorizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMode

`func (o *ScepInitParameters) SetAuthorizationMode(v string)`

SetAuthorizationMode sets AuthorizationMode field to given value.

### HasAuthorizationMode

`func (o *ScepInitParameters) HasAuthorizationMode() bool`

HasAuthorizationMode returns a boolean if a field has been set.

### GetCsrInfoIgnored

`func (o *ScepInitParameters) GetCsrInfoIgnored() bool`

GetCsrInfoIgnored returns the CsrInfoIgnored field if non-nil, zero value otherwise.

### GetCsrInfoIgnoredOk

`func (o *ScepInitParameters) GetCsrInfoIgnoredOk() (*bool, bool)`

GetCsrInfoIgnoredOk returns a tuple with the CsrInfoIgnored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrInfoIgnored

`func (o *ScepInitParameters) SetCsrInfoIgnored(v bool)`

SetCsrInfoIgnored sets CsrInfoIgnored field to given value.

### HasCsrInfoIgnored

`func (o *ScepInitParameters) HasCsrInfoIgnored() bool`

HasCsrInfoIgnored returns a boolean if a field has been set.

### GetKeyType

`func (o *ScepInitParameters) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ScepInitParameters) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ScepInitParameters) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ScepInitParameters) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetModule

`func (o *ScepInitParameters) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *ScepInitParameters) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *ScepInitParameters) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *ScepInitParameters) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ScepInitParameters) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ScepInitParameters) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


