# CertificateMetadataElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | **string** | Technical metadata related to the certificate | 
**Value** | Pointer to **NullableString** | The value of the metadata element | [optional] 

## Methods

### NewCertificateMetadataElement

`func NewCertificateMetadataElement(metadata string, ) *CertificateMetadataElement`

NewCertificateMetadataElement instantiates a new CertificateMetadataElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateMetadataElementWithDefaults

`func NewCertificateMetadataElementWithDefaults() *CertificateMetadataElement`

NewCertificateMetadataElementWithDefaults instantiates a new CertificateMetadataElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *CertificateMetadataElement) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CertificateMetadataElement) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CertificateMetadataElement) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetValue

`func (o *CertificateMetadataElement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CertificateMetadataElement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CertificateMetadataElement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CertificateMetadataElement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CertificateMetadataElement) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CertificateMetadataElement) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


