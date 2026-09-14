# AccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Object internal ID | 
**CompromisedAt** | Pointer to **int64** | The date when the account was compromised | [optional] 
**CompromissionReason** | Pointer to **NullableString** | One of: &#x60;unspecified&#x60;, &#x60;keycompromise&#x60;, &#x60;cacompromise&#x60;, &#x60;affiliationchange&#x60;, &#x60;superseded&#x60;, &#x60;cessationofoperation&#x60; | [optional] 
**Contact** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | **int64** |  | 
**EabName** | Pointer to **string** |  | [optional] 
**ExpirationDate** | Pointer to **int64** |  | [optional] 
**InitialIp** | Pointer to **string** |  | [optional] 
**Jwk** | [**JsonWebKey**](JsonWebKey.md) |  | 
**KeyThumbprint** | **string** |  | 
**Status** | [**AccountStatus**](AccountStatus.md) |  | 
**TermsOfServiceAgreed** | Pointer to **bool** |  | [optional] 

## Methods

### NewAccountResponse

`func NewAccountResponse(id string, createdAt int64, jwk JsonWebKey, keyThumbprint string, status AccountStatus, ) *AccountResponse`

NewAccountResponse instantiates a new AccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountResponseWithDefaults

`func NewAccountResponseWithDefaults() *AccountResponse`

NewAccountResponseWithDefaults instantiates a new AccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCompromisedAt

`func (o *AccountResponse) GetCompromisedAt() int64`

GetCompromisedAt returns the CompromisedAt field if non-nil, zero value otherwise.

### GetCompromisedAtOk

`func (o *AccountResponse) GetCompromisedAtOk() (*int64, bool)`

GetCompromisedAtOk returns a tuple with the CompromisedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromisedAt

`func (o *AccountResponse) SetCompromisedAt(v int64)`

SetCompromisedAt sets CompromisedAt field to given value.

### HasCompromisedAt

`func (o *AccountResponse) HasCompromisedAt() bool`

HasCompromisedAt returns a boolean if a field has been set.

### GetCompromissionReason

`func (o *AccountResponse) GetCompromissionReason() string`

GetCompromissionReason returns the CompromissionReason field if non-nil, zero value otherwise.

### GetCompromissionReasonOk

`func (o *AccountResponse) GetCompromissionReasonOk() (*string, bool)`

GetCompromissionReasonOk returns a tuple with the CompromissionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompromissionReason

`func (o *AccountResponse) SetCompromissionReason(v string)`

SetCompromissionReason sets CompromissionReason field to given value.

### HasCompromissionReason

`func (o *AccountResponse) HasCompromissionReason() bool`

HasCompromissionReason returns a boolean if a field has been set.

### SetCompromissionReasonNil

`func (o *AccountResponse) SetCompromissionReasonNil(b bool)`

 SetCompromissionReasonNil sets the value for CompromissionReason to be an explicit nil

### UnsetCompromissionReason
`func (o *AccountResponse) UnsetCompromissionReason()`

UnsetCompromissionReason ensures that no value is present for CompromissionReason, not even an explicit nil
### GetContact

`func (o *AccountResponse) GetContact() []string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *AccountResponse) GetContactOk() (*[]string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *AccountResponse) SetContact(v []string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *AccountResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountResponse) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountResponse) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountResponse) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetEabName

`func (o *AccountResponse) GetEabName() string`

GetEabName returns the EabName field if non-nil, zero value otherwise.

### GetEabNameOk

`func (o *AccountResponse) GetEabNameOk() (*string, bool)`

GetEabNameOk returns a tuple with the EabName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEabName

`func (o *AccountResponse) SetEabName(v string)`

SetEabName sets EabName field to given value.

### HasEabName

`func (o *AccountResponse) HasEabName() bool`

HasEabName returns a boolean if a field has been set.

### GetExpirationDate

`func (o *AccountResponse) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *AccountResponse) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *AccountResponse) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *AccountResponse) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetInitialIp

`func (o *AccountResponse) GetInitialIp() string`

GetInitialIp returns the InitialIp field if non-nil, zero value otherwise.

### GetInitialIpOk

`func (o *AccountResponse) GetInitialIpOk() (*string, bool)`

GetInitialIpOk returns a tuple with the InitialIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialIp

`func (o *AccountResponse) SetInitialIp(v string)`

SetInitialIp sets InitialIp field to given value.

### HasInitialIp

`func (o *AccountResponse) HasInitialIp() bool`

HasInitialIp returns a boolean if a field has been set.

### GetJwk

`func (o *AccountResponse) GetJwk() JsonWebKey`

GetJwk returns the Jwk field if non-nil, zero value otherwise.

### GetJwkOk

`func (o *AccountResponse) GetJwkOk() (*JsonWebKey, bool)`

GetJwkOk returns a tuple with the Jwk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwk

`func (o *AccountResponse) SetJwk(v JsonWebKey)`

SetJwk sets Jwk field to given value.


### GetKeyThumbprint

`func (o *AccountResponse) GetKeyThumbprint() string`

GetKeyThumbprint returns the KeyThumbprint field if non-nil, zero value otherwise.

### GetKeyThumbprintOk

`func (o *AccountResponse) GetKeyThumbprintOk() (*string, bool)`

GetKeyThumbprintOk returns a tuple with the KeyThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyThumbprint

`func (o *AccountResponse) SetKeyThumbprint(v string)`

SetKeyThumbprint sets KeyThumbprint field to given value.


### GetStatus

`func (o *AccountResponse) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountResponse) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountResponse) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.


### GetTermsOfServiceAgreed

`func (o *AccountResponse) GetTermsOfServiceAgreed() bool`

GetTermsOfServiceAgreed returns the TermsOfServiceAgreed field if non-nil, zero value otherwise.

### GetTermsOfServiceAgreedOk

`func (o *AccountResponse) GetTermsOfServiceAgreedOk() (*bool, bool)`

GetTermsOfServiceAgreedOk returns a tuple with the TermsOfServiceAgreed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceAgreed

`func (o *AccountResponse) SetTermsOfServiceAgreed(v bool)`

SetTermsOfServiceAgreed sets TermsOfServiceAgreed field to given value.

### HasTermsOfServiceAgreed

`func (o *AccountResponse) HasTermsOfServiceAgreed() bool`

HasTermsOfServiceAgreed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


