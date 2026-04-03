# BootstrapAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Identifier of the account | 
**Password** | Pointer to **string** | Password of the account. If it was given in the input, it is not provided here | [optional] 

## Methods

### NewBootstrapAccount

`func NewBootstrapAccount(identifier string, ) *BootstrapAccount`

NewBootstrapAccount instantiates a new BootstrapAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootstrapAccountWithDefaults

`func NewBootstrapAccountWithDefaults() *BootstrapAccount`

NewBootstrapAccountWithDefaults instantiates a new BootstrapAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *BootstrapAccount) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BootstrapAccount) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BootstrapAccount) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetPassword

`func (o *BootstrapAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BootstrapAccount) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BootstrapAccount) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BootstrapAccount) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


