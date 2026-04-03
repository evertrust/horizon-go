# EstInitParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationMode** | Pointer to **string** | The authorization mode for EST. | [optional] 
**CsrInfoIgnored** | Pointer to **bool** | Indicates whether CSR info is ignored for EST. | [optional] 
**EnrollmentMode** | Pointer to **string** | The enrollment mode for EST. | [optional] 
**KeyType** | Pointer to **string** | The key type used for EST. | [optional] 
**Module** | **string** | The module of the initialization parameters. | 
**Profile** | **string** | The profile used for EST. | 

## Methods

### NewEstInitParameters

`func NewEstInitParameters(module string, profile string, ) *EstInitParameters`

NewEstInitParameters instantiates a new EstInitParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstInitParametersWithDefaults

`func NewEstInitParametersWithDefaults() *EstInitParameters`

NewEstInitParametersWithDefaults instantiates a new EstInitParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationMode

`func (o *EstInitParameters) GetAuthorizationMode() string`

GetAuthorizationMode returns the AuthorizationMode field if non-nil, zero value otherwise.

### GetAuthorizationModeOk

`func (o *EstInitParameters) GetAuthorizationModeOk() (*string, bool)`

GetAuthorizationModeOk returns a tuple with the AuthorizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMode

`func (o *EstInitParameters) SetAuthorizationMode(v string)`

SetAuthorizationMode sets AuthorizationMode field to given value.

### HasAuthorizationMode

`func (o *EstInitParameters) HasAuthorizationMode() bool`

HasAuthorizationMode returns a boolean if a field has been set.

### GetCsrInfoIgnored

`func (o *EstInitParameters) GetCsrInfoIgnored() bool`

GetCsrInfoIgnored returns the CsrInfoIgnored field if non-nil, zero value otherwise.

### GetCsrInfoIgnoredOk

`func (o *EstInitParameters) GetCsrInfoIgnoredOk() (*bool, bool)`

GetCsrInfoIgnoredOk returns a tuple with the CsrInfoIgnored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrInfoIgnored

`func (o *EstInitParameters) SetCsrInfoIgnored(v bool)`

SetCsrInfoIgnored sets CsrInfoIgnored field to given value.

### HasCsrInfoIgnored

`func (o *EstInitParameters) HasCsrInfoIgnored() bool`

HasCsrInfoIgnored returns a boolean if a field has been set.

### GetEnrollmentMode

`func (o *EstInitParameters) GetEnrollmentMode() string`

GetEnrollmentMode returns the EnrollmentMode field if non-nil, zero value otherwise.

### GetEnrollmentModeOk

`func (o *EstInitParameters) GetEnrollmentModeOk() (*string, bool)`

GetEnrollmentModeOk returns a tuple with the EnrollmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMode

`func (o *EstInitParameters) SetEnrollmentMode(v string)`

SetEnrollmentMode sets EnrollmentMode field to given value.

### HasEnrollmentMode

`func (o *EstInitParameters) HasEnrollmentMode() bool`

HasEnrollmentMode returns a boolean if a field has been set.

### GetKeyType

`func (o *EstInitParameters) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *EstInitParameters) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *EstInitParameters) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *EstInitParameters) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetModule

`func (o *EstInitParameters) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *EstInitParameters) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *EstInitParameters) SetModule(v string)`

SetModule sets Module field to given value.


### GetProfile

`func (o *EstInitParameters) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *EstInitParameters) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *EstInitParameters) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


