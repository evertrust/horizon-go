# WebraInitParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationMode** | Pointer to **string** | The authorization mode for WebRA. | [optional] 
**EnrollmentMode** | Pointer to **string** | The enrollment mode for WebRA. | [optional] 
**KeyType** | Pointer to **string** | The key type used for WebRA. | [optional] 
**Module** | **string** | The module of the initialization parameters. | 
**PasswordPolicy** | Pointer to [**PasswordPolicy**](PasswordPolicy.md) | The password policy for WebRA. | [optional] 
**Profile** | **string** | The profile used for WebRA. | 

## Methods

### NewWebraInitParameters

`func NewWebraInitParameters(module string, profile string, ) *WebraInitParameters`

NewWebraInitParameters instantiates a new WebraInitParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebraInitParametersWithDefaults

`func NewWebraInitParametersWithDefaults() *WebraInitParameters`

NewWebraInitParametersWithDefaults instantiates a new WebraInitParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationMode

`func (o *WebraInitParameters) GetAuthorizationMode() string`

GetAuthorizationMode returns the AuthorizationMode field if non-nil, zero value otherwise.

### GetAuthorizationModeOk

`func (o *WebraInitParameters) GetAuthorizationModeOk() (*string, bool)`

GetAuthorizationModeOk returns a tuple with the AuthorizationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationMode

`func (o *WebraInitParameters) SetAuthorizationMode(v string)`

SetAuthorizationMode sets AuthorizationMode field to given value.

### HasAuthorizationMode

`func (o *WebraInitParameters) HasAuthorizationMode() bool`

HasAuthorizationMode returns a boolean if a field has been set.

### GetEnrollmentMode

`func (o *WebraInitParameters) GetEnrollmentMode() string`

GetEnrollmentMode returns the EnrollmentMode field if non-nil, zero value otherwise.

### GetEnrollmentModeOk

`func (o *WebraInitParameters) GetEnrollmentModeOk() (*string, bool)`

GetEnrollmentModeOk returns a tuple with the EnrollmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentMode

`func (o *WebraInitParameters) SetEnrollmentMode(v string)`

SetEnrollmentMode sets EnrollmentMode field to given value.

### HasEnrollmentMode

`func (o *WebraInitParameters) HasEnrollmentMode() bool`

HasEnrollmentMode returns a boolean if a field has been set.

### GetKeyType

`func (o *WebraInitParameters) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *WebraInitParameters) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *WebraInitParameters) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *WebraInitParameters) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetModule

`func (o *WebraInitParameters) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *WebraInitParameters) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *WebraInitParameters) SetModule(v string)`

SetModule sets Module field to given value.


### GetPasswordPolicy

`func (o *WebraInitParameters) GetPasswordPolicy() PasswordPolicy`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *WebraInitParameters) GetPasswordPolicyOk() (*PasswordPolicy, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *WebraInitParameters) SetPasswordPolicy(v PasswordPolicy)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *WebraInitParameters) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetProfile

`func (o *WebraInitParameters) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *WebraInitParameters) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *WebraInitParameters) SetProfile(v string)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


