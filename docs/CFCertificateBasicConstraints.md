# CFCertificateBasicConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCa** | **bool** | If the certificate is type of CA | 
**PathLen** | Pointer to **int64** | The path len of the certificate | [optional] 

## Methods

### NewCFCertificateBasicConstraints

`func NewCFCertificateBasicConstraints(isCa bool, ) *CFCertificateBasicConstraints`

NewCFCertificateBasicConstraints instantiates a new CFCertificateBasicConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFCertificateBasicConstraintsWithDefaults

`func NewCFCertificateBasicConstraintsWithDefaults() *CFCertificateBasicConstraints`

NewCFCertificateBasicConstraintsWithDefaults instantiates a new CFCertificateBasicConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCa

`func (o *CFCertificateBasicConstraints) GetIsCa() bool`

GetIsCa returns the IsCa field if non-nil, zero value otherwise.

### GetIsCaOk

`func (o *CFCertificateBasicConstraints) GetIsCaOk() (*bool, bool)`

GetIsCaOk returns a tuple with the IsCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCa

`func (o *CFCertificateBasicConstraints) SetIsCa(v bool)`

SetIsCa sets IsCa field to given value.


### GetPathLen

`func (o *CFCertificateBasicConstraints) GetPathLen() int64`

GetPathLen returns the PathLen field if non-nil, zero value otherwise.

### GetPathLenOk

`func (o *CFCertificateBasicConstraints) GetPathLenOk() (*int64, bool)`

GetPathLenOk returns a tuple with the PathLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathLen

`func (o *CFCertificateBasicConstraints) SetPathLen(v int64)`

SetPathLen sets PathLen field to given value.

### HasPathLen

`func (o *CFCertificateBasicConstraints) HasPathLen() bool`

HasPathLen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


