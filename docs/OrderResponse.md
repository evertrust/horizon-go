# OrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**AccountId** | **string** | Object internal ID | 
**Authorizations** | [**[]Authorization**](Authorization.md) |  | [default to []]
**Certificate** | Pointer to **NullableString** | Object internal ID | [optional] 
**ContactEmail** | Pointer to **NullableString** |  | [optional] 
**Expires** | Pointer to **NullableInt64** |  | [optional] 
**InitialIp** | Pointer to **NullableString** |  | [optional] 
**Label** | Pointer to [**[]OrderLabelInner**](OrderLabelInner.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] [default to {}]
**NotAfter** | Pointer to **NullableInt64** |  | [optional] 
**NotBefore** | Pointer to **NullableInt64** |  | [optional] 
**Owner** | Pointer to **NullableString** |  | [optional] 
**Profile** | **string** |  | 
**RemoveAt** | Pointer to **NullableInt64** |  | [optional] 
**Status** | [**OrderStatus**](OrderStatus.md) |  | 
**Team** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOrderResponse

`func NewOrderResponse(id string, accountId string, authorizations []Authorization, profile string, status OrderStatus, ) *OrderResponse`

NewOrderResponse instantiates a new OrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResponseWithDefaults

`func NewOrderResponseWithDefaults() *OrderResponse`

NewOrderResponseWithDefaults instantiates a new OrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *OrderResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OrderResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *OrderResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAuthorizations

`func (o *OrderResponse) GetAuthorizations() []Authorization`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *OrderResponse) GetAuthorizationsOk() (*[]Authorization, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *OrderResponse) SetAuthorizations(v []Authorization)`

SetAuthorizations sets Authorizations field to given value.


### GetCertificate

`func (o *OrderResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *OrderResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *OrderResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *OrderResponse) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *OrderResponse) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *OrderResponse) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetContactEmail

`func (o *OrderResponse) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *OrderResponse) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *OrderResponse) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *OrderResponse) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *OrderResponse) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *OrderResponse) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetExpires

`func (o *OrderResponse) GetExpires() int64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *OrderResponse) GetExpiresOk() (*int64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *OrderResponse) SetExpires(v int64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *OrderResponse) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *OrderResponse) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *OrderResponse) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetInitialIp

`func (o *OrderResponse) GetInitialIp() string`

GetInitialIp returns the InitialIp field if non-nil, zero value otherwise.

### GetInitialIpOk

`func (o *OrderResponse) GetInitialIpOk() (*string, bool)`

GetInitialIpOk returns a tuple with the InitialIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialIp

`func (o *OrderResponse) SetInitialIp(v string)`

SetInitialIp sets InitialIp field to given value.

### HasInitialIp

`func (o *OrderResponse) HasInitialIp() bool`

HasInitialIp returns a boolean if a field has been set.

### SetInitialIpNil

`func (o *OrderResponse) SetInitialIpNil(b bool)`

 SetInitialIpNil sets the value for InitialIp to be an explicit nil

### UnsetInitialIp
`func (o *OrderResponse) UnsetInitialIp()`

UnsetInitialIp ensures that no value is present for InitialIp, not even an explicit nil
### GetLabel

`func (o *OrderResponse) GetLabel() []OrderLabelInner`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *OrderResponse) GetLabelOk() (*[]OrderLabelInner, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *OrderResponse) SetLabel(v []OrderLabelInner)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *OrderResponse) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *OrderResponse) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *OrderResponse) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetMetadata

`func (o *OrderResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrderResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrderResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrderResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNotAfter

`func (o *OrderResponse) GetNotAfter() int64`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *OrderResponse) GetNotAfterOk() (*int64, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *OrderResponse) SetNotAfter(v int64)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *OrderResponse) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### SetNotAfterNil

`func (o *OrderResponse) SetNotAfterNil(b bool)`

 SetNotAfterNil sets the value for NotAfter to be an explicit nil

### UnsetNotAfter
`func (o *OrderResponse) UnsetNotAfter()`

UnsetNotAfter ensures that no value is present for NotAfter, not even an explicit nil
### GetNotBefore

`func (o *OrderResponse) GetNotBefore() int64`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *OrderResponse) GetNotBeforeOk() (*int64, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *OrderResponse) SetNotBefore(v int64)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *OrderResponse) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### SetNotBeforeNil

`func (o *OrderResponse) SetNotBeforeNil(b bool)`

 SetNotBeforeNil sets the value for NotBefore to be an explicit nil

### UnsetNotBefore
`func (o *OrderResponse) UnsetNotBefore()`

UnsetNotBefore ensures that no value is present for NotBefore, not even an explicit nil
### GetOwner

`func (o *OrderResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *OrderResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *OrderResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *OrderResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *OrderResponse) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *OrderResponse) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetProfile

`func (o *OrderResponse) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OrderResponse) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OrderResponse) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetRemoveAt

`func (o *OrderResponse) GetRemoveAt() int64`

GetRemoveAt returns the RemoveAt field if non-nil, zero value otherwise.

### GetRemoveAtOk

`func (o *OrderResponse) GetRemoveAtOk() (*int64, bool)`

GetRemoveAtOk returns a tuple with the RemoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAt

`func (o *OrderResponse) SetRemoveAt(v int64)`

SetRemoveAt sets RemoveAt field to given value.

### HasRemoveAt

`func (o *OrderResponse) HasRemoveAt() bool`

HasRemoveAt returns a boolean if a field has been set.

### SetRemoveAtNil

`func (o *OrderResponse) SetRemoveAtNil(b bool)`

 SetRemoveAtNil sets the value for RemoveAt to be an explicit nil

### UnsetRemoveAt
`func (o *OrderResponse) UnsetRemoveAt()`

UnsetRemoveAt ensures that no value is present for RemoveAt, not even an explicit nil
### GetStatus

`func (o *OrderResponse) GetStatus() OrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderResponse) GetStatusOk() (*OrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderResponse) SetStatus(v OrderStatus)`

SetStatus sets Status field to given value.


### GetTeam

`func (o *OrderResponse) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *OrderResponse) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *OrderResponse) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *OrderResponse) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *OrderResponse) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *OrderResponse) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetThumbprint

`func (o *OrderResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *OrderResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *OrderResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *OrderResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *OrderResponse) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *OrderResponse) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


