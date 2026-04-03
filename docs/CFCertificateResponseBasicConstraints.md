# CFCertificateResponseBasicConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCa** | **bool** | If the certificate is type of CA | 
**PathLen** | Pointer to **int64** | The path len of the certificate | [optional] 

## Methods

### NewCFCertificateResponseBasicConstraints

`func NewCFCertificateResponseBasicConstraints(isCa bool, ) *CFCertificateResponseBasicConstraints`

NewCFCertificateResponseBasicConstraints instantiates a new CFCertificateResponseBasicConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFCertificateResponseBasicConstraintsWithDefaults

`func NewCFCertificateResponseBasicConstraintsWithDefaults() *CFCertificateResponseBasicConstraints`

NewCFCertificateResponseBasicConstraintsWithDefaults instantiates a new CFCertificateResponseBasicConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCa

`func (o *CFCertificateResponseBasicConstraints) GetIsCa() bool`

GetIsCa returns the IsCa field if non-nil, zero value otherwise.

### GetIsCaOk

`func (o *CFCertificateResponseBasicConstraints) GetIsCaOk() (*bool, bool)`

GetIsCaOk returns a tuple with the IsCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCa

`func (o *CFCertificateResponseBasicConstraints) SetIsCa(v bool)`

SetIsCa sets IsCa field to given value.


### GetPathLen

`func (o *CFCertificateResponseBasicConstraints) GetPathLen() int64`

GetPathLen returns the PathLen field if non-nil, zero value otherwise.

### GetPathLenOk

`func (o *CFCertificateResponseBasicConstraints) GetPathLenOk() (*int64, bool)`

GetPathLenOk returns a tuple with the PathLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathLen

`func (o *CFCertificateResponseBasicConstraints) SetPathLen(v int64)`

SetPathLen sets PathLen field to given value.

### HasPathLen

`func (o *CFCertificateResponseBasicConstraints) HasPathLen() bool`

HasPathLen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


