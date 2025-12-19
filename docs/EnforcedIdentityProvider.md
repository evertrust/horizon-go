# EnforcedIdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the identity provider to be enforced | 
**Type** | **string** | The type of identity provider to be enforced | 

## Methods

### NewEnforcedIdentityProvider

`func NewEnforcedIdentityProvider(name string, type_ string, ) *EnforcedIdentityProvider`

NewEnforcedIdentityProvider instantiates a new EnforcedIdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnforcedIdentityProviderWithDefaults

`func NewEnforcedIdentityProviderWithDefaults() *EnforcedIdentityProvider`

NewEnforcedIdentityProviderWithDefaults instantiates a new EnforcedIdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnforcedIdentityProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnforcedIdentityProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnforcedIdentityProvider) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EnforcedIdentityProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnforcedIdentityProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnforcedIdentityProvider) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


