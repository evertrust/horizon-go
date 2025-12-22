# CertificateExtensionElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the extension element | 
**Value** | Pointer to **NullableString** | The value of the extension element | [optional] 

## Methods

### NewCertificateExtensionElement

`func NewCertificateExtensionElement(type_ string, ) *CertificateExtensionElement`

NewCertificateExtensionElement instantiates a new CertificateExtensionElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateExtensionElementWithDefaults

`func NewCertificateExtensionElementWithDefaults() *CertificateExtensionElement`

NewCertificateExtensionElementWithDefaults instantiates a new CertificateExtensionElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CertificateExtensionElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateExtensionElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateExtensionElement) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *CertificateExtensionElement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CertificateExtensionElement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CertificateExtensionElement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CertificateExtensionElement) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CertificateExtensionElement) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CertificateExtensionElement) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


