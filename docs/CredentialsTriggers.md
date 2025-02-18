# CredentialsTriggers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnCredentialsExpiration** | Pointer to **[]string** | The notifications to be sent when the credentials expire | [optional] 

## Methods

### NewCredentialsTriggers

`func NewCredentialsTriggers() *CredentialsTriggers`

NewCredentialsTriggers instantiates a new CredentialsTriggers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsTriggersWithDefaults

`func NewCredentialsTriggersWithDefaults() *CredentialsTriggers`

NewCredentialsTriggersWithDefaults instantiates a new CredentialsTriggers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnCredentialsExpiration

`func (o *CredentialsTriggers) GetOnCredentialsExpiration() []string`

GetOnCredentialsExpiration returns the OnCredentialsExpiration field if non-nil, zero value otherwise.

### GetOnCredentialsExpirationOk

`func (o *CredentialsTriggers) GetOnCredentialsExpirationOk() (*[]string, bool)`

GetOnCredentialsExpirationOk returns a tuple with the OnCredentialsExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCredentialsExpiration

`func (o *CredentialsTriggers) SetOnCredentialsExpiration(v []string)`

SetOnCredentialsExpiration sets OnCredentialsExpiration field to given value.

### HasOnCredentialsExpiration

`func (o *CredentialsTriggers) HasOnCredentialsExpiration() bool`

HasOnCredentialsExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


