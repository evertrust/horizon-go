# Identifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Identifier type. Known values are &#x60;dns&#x60; and &#x60;ip&#x60;, but any string is accepted.  | 
**Value** | **string** |  | 

## Methods

### NewIdentifier

`func NewIdentifier(type_ string, value string, ) *Identifier`

NewIdentifier instantiates a new Identifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifierWithDefaults

`func NewIdentifierWithDefaults() *Identifier`

NewIdentifierWithDefaults instantiates a new Identifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Identifier) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Identifier) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Identifier) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *Identifier) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Identifier) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Identifier) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


