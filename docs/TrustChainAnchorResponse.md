# TrustChainAnchorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | [**CFCertificateResponse**](CFCertificateResponse.md) |  | 
**Name** | **string** |  | 
**Subordinates** | Pointer to [**[]TrustChainAnchorResponse**](TrustChainAnchorResponse.md) |  | [optional] 

## Methods

### NewTrustChainAnchorResponse

`func NewTrustChainAnchorResponse(certificate CFCertificateResponse, name string, ) *TrustChainAnchorResponse`

NewTrustChainAnchorResponse instantiates a new TrustChainAnchorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustChainAnchorResponseWithDefaults

`func NewTrustChainAnchorResponseWithDefaults() *TrustChainAnchorResponse`

NewTrustChainAnchorResponseWithDefaults instantiates a new TrustChainAnchorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *TrustChainAnchorResponse) GetCertificate() CFCertificateResponse`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TrustChainAnchorResponse) GetCertificateOk() (*CFCertificateResponse, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TrustChainAnchorResponse) SetCertificate(v CFCertificateResponse)`

SetCertificate sets Certificate field to given value.


### GetName

`func (o *TrustChainAnchorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrustChainAnchorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrustChainAnchorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSubordinates

`func (o *TrustChainAnchorResponse) GetSubordinates() []TrustChainAnchorResponse`

GetSubordinates returns the Subordinates field if non-nil, zero value otherwise.

### GetSubordinatesOk

`func (o *TrustChainAnchorResponse) GetSubordinatesOk() (*[]TrustChainAnchorResponse, bool)`

GetSubordinatesOk returns a tuple with the Subordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubordinates

`func (o *TrustChainAnchorResponse) SetSubordinates(v []TrustChainAnchorResponse)`

SetSubordinates sets Subordinates field to given value.

### HasSubordinates

`func (o *TrustChainAnchorResponse) HasSubordinates() bool`

HasSubordinates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


