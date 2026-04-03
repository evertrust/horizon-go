# Rfc5280CRLResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | [**CFCrl**](CFCrl.md) |  | 

## Methods

### NewRfc5280CRLResponse

`func NewRfc5280CRLResponse(type_ string, value CFCrl, ) *Rfc5280CRLResponse`

NewRfc5280CRLResponse instantiates a new Rfc5280CRLResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRfc5280CRLResponseWithDefaults

`func NewRfc5280CRLResponseWithDefaults() *Rfc5280CRLResponse`

NewRfc5280CRLResponseWithDefaults instantiates a new Rfc5280CRLResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Rfc5280CRLResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Rfc5280CRLResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Rfc5280CRLResponse) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *Rfc5280CRLResponse) GetValue() CFCrl`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Rfc5280CRLResponse) GetValueOk() (*CFCrl, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Rfc5280CRLResponse) SetValue(v CFCrl)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


