# PrincipalInfoSavedQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the query | 
**Query** | **string** | The saved HQL query | 
**Name** | **string** | Internal name of the saved request | 
**Description** | Pointer to **NullableString** | The saved request description | [optional] 

## Methods

### NewPrincipalInfoSavedQuery

`func NewPrincipalInfoSavedQuery(type_ string, query string, name string, ) *PrincipalInfoSavedQuery`

NewPrincipalInfoSavedQuery instantiates a new PrincipalInfoSavedQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalInfoSavedQueryWithDefaults

`func NewPrincipalInfoSavedQueryWithDefaults() *PrincipalInfoSavedQuery`

NewPrincipalInfoSavedQueryWithDefaults instantiates a new PrincipalInfoSavedQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PrincipalInfoSavedQuery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrincipalInfoSavedQuery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrincipalInfoSavedQuery) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *PrincipalInfoSavedQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *PrincipalInfoSavedQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *PrincipalInfoSavedQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetName

`func (o *PrincipalInfoSavedQuery) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalInfoSavedQuery) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalInfoSavedQuery) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PrincipalInfoSavedQuery) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrincipalInfoSavedQuery) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrincipalInfoSavedQuery) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrincipalInfoSavedQuery) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PrincipalInfoSavedQuery) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PrincipalInfoSavedQuery) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


