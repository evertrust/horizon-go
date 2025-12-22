# PrincipalInfoSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | Pointer to **NullableString** | The contact e-mail of the principal to search for | [optional] 
**Identifier** | Pointer to **NullableString** | The identifier of the principal to search for | [optional] 

## Methods

### NewPrincipalInfoSearchRequest

`func NewPrincipalInfoSearchRequest() *PrincipalInfoSearchRequest`

NewPrincipalInfoSearchRequest instantiates a new PrincipalInfoSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalInfoSearchRequestWithDefaults

`func NewPrincipalInfoSearchRequestWithDefaults() *PrincipalInfoSearchRequest`

NewPrincipalInfoSearchRequestWithDefaults instantiates a new PrincipalInfoSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *PrincipalInfoSearchRequest) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PrincipalInfoSearchRequest) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PrincipalInfoSearchRequest) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PrincipalInfoSearchRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *PrincipalInfoSearchRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *PrincipalInfoSearchRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetIdentifier

`func (o *PrincipalInfoSearchRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PrincipalInfoSearchRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PrincipalInfoSearchRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PrincipalInfoSearchRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *PrincipalInfoSearchRequest) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *PrincipalInfoSearchRequest) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


