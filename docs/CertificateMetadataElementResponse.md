# CertificateMetadataElementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Editable** | Pointer to **NullableBool** | Whether the metadata element is editable by the requester | [optional] 
**Metadata** | **string** | Technical metadata related to the certificate | 
**Value** | Pointer to **NullableString** | The value of the metadata element | [optional] 

## Methods

### NewCertificateMetadataElementResponse

`func NewCertificateMetadataElementResponse(metadata string, ) *CertificateMetadataElementResponse`

NewCertificateMetadataElementResponse instantiates a new CertificateMetadataElementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateMetadataElementResponseWithDefaults

`func NewCertificateMetadataElementResponseWithDefaults() *CertificateMetadataElementResponse`

NewCertificateMetadataElementResponseWithDefaults instantiates a new CertificateMetadataElementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditable

`func (o *CertificateMetadataElementResponse) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *CertificateMetadataElementResponse) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *CertificateMetadataElementResponse) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *CertificateMetadataElementResponse) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### SetEditableNil

`func (o *CertificateMetadataElementResponse) SetEditableNil(b bool)`

 SetEditableNil sets the value for Editable to be an explicit nil

### UnsetEditable
`func (o *CertificateMetadataElementResponse) UnsetEditable()`

UnsetEditable ensures that no value is present for Editable, not even an explicit nil
### GetMetadata

`func (o *CertificateMetadataElementResponse) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CertificateMetadataElementResponse) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CertificateMetadataElementResponse) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetValue

`func (o *CertificateMetadataElementResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CertificateMetadataElementResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CertificateMetadataElementResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CertificateMetadataElementResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CertificateMetadataElementResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CertificateMetadataElementResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


