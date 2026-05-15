# OidcIdentityProviderMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]OidcIdentityProviderMappingEntriesInner**](OidcIdentityProviderMappingEntriesInner.md) | Describe how each JWT claim corresponds to Horizon Roles and Teams | [optional] 
**Extraction** | Pointer to **string** | Computation Rule to extract the claims to map from the JWT claims | [optional] 

## Methods

### NewOidcIdentityProviderMapping

`func NewOidcIdentityProviderMapping() *OidcIdentityProviderMapping`

NewOidcIdentityProviderMapping instantiates a new OidcIdentityProviderMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcIdentityProviderMappingWithDefaults

`func NewOidcIdentityProviderMappingWithDefaults() *OidcIdentityProviderMapping`

NewOidcIdentityProviderMappingWithDefaults instantiates a new OidcIdentityProviderMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *OidcIdentityProviderMapping) GetEntries() []OidcIdentityProviderMappingEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *OidcIdentityProviderMapping) GetEntriesOk() (*[]OidcIdentityProviderMappingEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *OidcIdentityProviderMapping) SetEntries(v []OidcIdentityProviderMappingEntriesInner)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *OidcIdentityProviderMapping) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetExtraction

`func (o *OidcIdentityProviderMapping) GetExtraction() string`

GetExtraction returns the Extraction field if non-nil, zero value otherwise.

### GetExtractionOk

`func (o *OidcIdentityProviderMapping) GetExtractionOk() (*string, bool)`

GetExtractionOk returns a tuple with the Extraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraction

`func (o *OidcIdentityProviderMapping) SetExtraction(v string)`

SetExtraction sets Extraction field to given value.

### HasExtraction

`func (o *OidcIdentityProviderMapping) HasExtraction() bool`

HasExtraction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


