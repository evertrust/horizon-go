# AcmeExternalInitParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Module** | **string** | The module of the initialization parameters. | 
**Profile** | **string** | The profile used for ACME external. | 
**KeyType** | Pointer to **string** | The key type used for ACME external. | [optional] 
**AcmeUrl** | Pointer to **string** | The ACME URL for the external ACME server. | [optional] 
**AllowedAuthorizationMethod** | Pointer to **[]string** | The allowed authorization methods for ACME external. | [optional] 
**RequireEAB** | Pointer to **bool** | Indicates whether EAB is required for ACME external. | [optional] 

## Methods

### NewAcmeExternalInitParameters

`func NewAcmeExternalInitParameters(module string, profile string, ) *AcmeExternalInitParameters`

NewAcmeExternalInitParameters instantiates a new AcmeExternalInitParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcmeExternalInitParametersWithDefaults

`func NewAcmeExternalInitParametersWithDefaults() *AcmeExternalInitParameters`

NewAcmeExternalInitParametersWithDefaults instantiates a new AcmeExternalInitParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModule

`func (o *AcmeExternalInitParameters) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *AcmeExternalInitParameters) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *AcmeExternalInitParameters) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *AcmeExternalInitParameters) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AcmeExternalInitParameters) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AcmeExternalInitParameters) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetKeyType

`func (o *AcmeExternalInitParameters) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *AcmeExternalInitParameters) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *AcmeExternalInitParameters) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *AcmeExternalInitParameters) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetAcmeUrl

`func (o *AcmeExternalInitParameters) GetAcmeUrl() string`

GetAcmeUrl returns the AcmeUrl field if non-nil, zero value otherwise.

### GetAcmeUrlOk

`func (o *AcmeExternalInitParameters) GetAcmeUrlOk() (*string, bool)`

GetAcmeUrlOk returns a tuple with the AcmeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcmeUrl

`func (o *AcmeExternalInitParameters) SetAcmeUrl(v string)`

SetAcmeUrl sets AcmeUrl field to given value.

### HasAcmeUrl

`func (o *AcmeExternalInitParameters) HasAcmeUrl() bool`

HasAcmeUrl returns a boolean if a field has been set.

### GetAllowedAuthorizationMethod

`func (o *AcmeExternalInitParameters) GetAllowedAuthorizationMethod() []string`

GetAllowedAuthorizationMethod returns the AllowedAuthorizationMethod field if non-nil, zero value otherwise.

### GetAllowedAuthorizationMethodOk

`func (o *AcmeExternalInitParameters) GetAllowedAuthorizationMethodOk() (*[]string, bool)`

GetAllowedAuthorizationMethodOk returns a tuple with the AllowedAuthorizationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthorizationMethod

`func (o *AcmeExternalInitParameters) SetAllowedAuthorizationMethod(v []string)`

SetAllowedAuthorizationMethod sets AllowedAuthorizationMethod field to given value.

### HasAllowedAuthorizationMethod

`func (o *AcmeExternalInitParameters) HasAllowedAuthorizationMethod() bool`

HasAllowedAuthorizationMethod returns a boolean if a field has been set.

### GetRequireEAB

`func (o *AcmeExternalInitParameters) GetRequireEAB() bool`

GetRequireEAB returns the RequireEAB field if non-nil, zero value otherwise.

### GetRequireEABOk

`func (o *AcmeExternalInitParameters) GetRequireEABOk() (*bool, bool)`

GetRequireEABOk returns a tuple with the RequireEAB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEAB

`func (o *AcmeExternalInitParameters) SetRequireEAB(v bool)`

SetRequireEAB sets RequireEAB field to given value.

### HasRequireEAB

`func (o *AcmeExternalInitParameters) HasRequireEAB() bool`

HasRequireEAB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


