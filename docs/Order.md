# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AccountId** | **string** | Object internal ID | 
**Authorizations** | [**[]Authorization**](Authorization.md) |  | [default to []]
**Certificate** | Pointer to **string** | Object internal ID | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **int64** |  | [optional] 
**InitialIp** | Pointer to **string** |  | [optional] 
**Label** | Pointer to [**[]OrderLabelInner**](OrderLabelInner.md) |  | [optional] 
**Metadata** | **map[string]string** |  | [default to {}]
**NotAfter** | Pointer to **int64** |  | [optional] 
**NotBefore** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Profile** | **string** |  | 
**RemoveAt** | Pointer to **int64** |  | [optional] 
**Status** | [**OrderStatus**](OrderStatus.md) |  | 
**Team** | Pointer to **string** |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 

## Methods

### NewOrder

`func NewOrder(id string, accountId string, authorizations []Authorization, metadata map[string]string, profile string, status OrderStatus, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *Order) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Order) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Order) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAuthorizations

`func (o *Order) GetAuthorizations() []Authorization`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *Order) GetAuthorizationsOk() (*[]Authorization, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *Order) SetAuthorizations(v []Authorization)`

SetAuthorizations sets Authorizations field to given value.


### GetCertificate

`func (o *Order) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Order) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Order) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *Order) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetContactEmail

`func (o *Order) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Order) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Order) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *Order) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetExpires

`func (o *Order) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Order) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Order) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Order) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetInitialIp

`func (o *Order) GetInitialIp() string`

GetInitialIp returns the InitialIp field if non-nil, zero value otherwise.

### GetInitialIpOk

`func (o *Order) GetInitialIpOk() (*string, bool)`

GetInitialIpOk returns a tuple with the InitialIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialIp

`func (o *Order) SetInitialIp(v string)`

SetInitialIp sets InitialIp field to given value.

### HasInitialIp

`func (o *Order) HasInitialIp() bool`

HasInitialIp returns a boolean if a field has been set.

### GetLabel

`func (o *Order) GetLabel() []OrderLabelInner`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Order) GetLabelOk() (*[]OrderLabelInner, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Order) SetLabel(v []OrderLabelInner)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Order) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetadata

`func (o *Order) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Order) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Order) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### GetNotAfter

`func (o *Order) GetNotAfter() int64`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *Order) GetNotAfterOk() (*int64, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *Order) SetNotAfter(v int64)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *Order) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *Order) GetNotBefore() int64`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *Order) GetNotBeforeOk() (*int64, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *Order) SetNotBefore(v int64)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *Order) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetOwner

`func (o *Order) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Order) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Order) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Order) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetProfile

`func (o *Order) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Order) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Order) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetRemoveAt

`func (o *Order) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *Order) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *Order) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.

### HasRemoveAt

`func (o *Order) HasRemoveAt() bool`

HasRemoveAt returns a boolean if a field has been set.

### GetStatus

`func (o *Order) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.


### GetTeam

`func (o *Order) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *Order) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *Order) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *Order) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetThumbprint

`func (o *Order) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *Order) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *Order) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *Order) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


