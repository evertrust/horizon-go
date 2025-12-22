# WebRARenewRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Csr** | Pointer to **NullableString** | The CSR used to renew the certificate, if in decentralized mode | [optional] 
**KeyType** | Pointer to **NullableString** | The key type of the certificate, if in centralized mode | [optional] 

## Methods

### NewWebRARenewRequestTemplate

`func NewWebRARenewRequestTemplate() *WebRARenewRequestTemplate`

NewWebRARenewRequestTemplate instantiates a new WebRARenewRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRARenewRequestTemplateWithDefaults

`func NewWebRARenewRequestTemplateWithDefaults() *WebRARenewRequestTemplate`

NewWebRARenewRequestTemplateWithDefaults instantiates a new WebRARenewRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsr

`func (o *WebRARenewRequestTemplate) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *WebRARenewRequestTemplate) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *WebRARenewRequestTemplate) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *WebRARenewRequestTemplate) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### SetCsrNil

`func (o *WebRARenewRequestTemplate) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *WebRARenewRequestTemplate) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil
### GetKeyType

`func (o *WebRARenewRequestTemplate) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *WebRARenewRequestTemplate) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *WebRARenewRequestTemplate) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *WebRARenewRequestTemplate) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *WebRARenewRequestTemplate) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *WebRARenewRequestTemplate) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


