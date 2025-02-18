# PrincipalInfoSearchResultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | The identifier of the principal matching the search | 
**Contact** | Pointer to **NullableString** | The mail of the principal matching the search | [optional] 
**ProviderType** | **string** | The type of the identity provider on which this user is registered | 
**ProviderName** | **string** | The name of the identity provider on which this user is registered | 

## Methods

### NewPrincipalInfoSearchResultResponse

`func NewPrincipalInfoSearchResultResponse(identifier string, providerType string, providerName string, ) *PrincipalInfoSearchResultResponse`

NewPrincipalInfoSearchResultResponse instantiates a new PrincipalInfoSearchResultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalInfoSearchResultResponseWithDefaults

`func NewPrincipalInfoSearchResultResponseWithDefaults() *PrincipalInfoSearchResultResponse`

NewPrincipalInfoSearchResultResponseWithDefaults instantiates a new PrincipalInfoSearchResultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *PrincipalInfoSearchResultResponse) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PrincipalInfoSearchResultResponse) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PrincipalInfoSearchResultResponse) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetContact

`func (o *PrincipalInfoSearchResultResponse) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PrincipalInfoSearchResultResponse) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PrincipalInfoSearchResultResponse) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PrincipalInfoSearchResultResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *PrincipalInfoSearchResultResponse) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *PrincipalInfoSearchResultResponse) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetProviderType

`func (o *PrincipalInfoSearchResultResponse) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *PrincipalInfoSearchResultResponse) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *PrincipalInfoSearchResultResponse) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetProviderName

`func (o *PrincipalInfoSearchResultResponse) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *PrincipalInfoSearchResultResponse) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *PrincipalInfoSearchResultResponse) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


