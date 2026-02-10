# TrustChainAnchor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Certificate** | [**CFCertificate**](CFCertificate.md) |  | 
**Subordinates** | Pointer to [**[]TrustChainAnchor**](TrustChainAnchor.md) |  | [optional] 

## Methods

### NewTrustChainAnchor

`func NewTrustChainAnchor(name string, certificate CFCertificate, ) *TrustChainAnchor`

NewTrustChainAnchor instantiates a new TrustChainAnchor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustChainAnchorWithDefaults

`func NewTrustChainAnchorWithDefaults() *TrustChainAnchor`

NewTrustChainAnchorWithDefaults instantiates a new TrustChainAnchor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TrustChainAnchor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustChainAnchor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustChainAnchor) SetName(v string)`

SetName sets Name field to given value.


### GetCertificate

`func (o *TrustChainAnchor) GetCertificate() CFCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TrustChainAnchor) GetCertificateOk() (*CFCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TrustChainAnchor) SetCertificate(v CFCertificate)`

SetCertificate sets Certificate field to given value.


### GetSubordinates

`func (o *TrustChainAnchor) GetSubordinates() []TrustChainAnchor`

GetSubordinates returns the Subordinates field if non-nil, zero value otherwise.

### GetSubordinatesOk

`func (o *TrustChainAnchor) GetSubordinatesOk() (*[]TrustChainAnchor, bool)`

GetSubordinatesOk returns a tuple with the Subordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubordinates

`func (o *TrustChainAnchor) SetSubordinates(v []TrustChainAnchor)`

SetSubordinates sets Subordinates field to given value.

### HasSubordinates

`func (o *TrustChainAnchor) HasSubordinates() bool`

HasSubordinates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


