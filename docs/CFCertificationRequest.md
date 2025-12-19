# CFCertificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dn** | **string** | Distinguished name | 
**DnElements** | [**[]CFDistinguishedName**](CFDistinguishedName.md) |  | 
**KeyType** | **string** | One of &#x60;rsa-2048&#x60;, &#x60;rsa-3072&#x60;, &#x60;rsa-4096&#x60;, &#x60;rsa-8192&#x60;, &#x60;ec-secp256r1&#x60;, &#x60;ec-secp384r1&#x60;, &#x60;ec-secp521r1&#x60;, &#x60;ed-448&#x60;, &#x60;ed-25519&#x60;, &#x60;mldsa-44&#x60;, &#x60;mldsa-65&#x60;, &#x60;mldsa-87&#x60;, &#x60;slhdsa-sha2-128s&#x60;, &#x60;slhdsa-sha2-128f&#x60;, &#x60;slhdsa-sha2-192s&#x60;, &#x60;slhdsa-sha2-192f&#x60;, &#x60;slhdsa-sha2-256s&#x60;, &#x60;slhdsa-sha2-256f&#x60;, &#x60;slhdsa-sha2-128ssha256&#x60;, &#x60;slhdsa-sha2-128fsha256&#x60;, &#x60;slhdsa-sha2-192ssha512&#x60;, &#x60;slhdsa-sha2-192fsha512&#x60;, &#x60;slhdsa-sha2-256ssha512&#x60;, &#x60;slhdsa-sha2-256fsha512&#x60; or &#x60;&lt;primary key type&gt;+&lt;alternate key type&gt;&#x60; | 
**Pem** | **string** |  | 
**Sans** | Pointer to [**[]SubjectAlternateName**](SubjectAlternateName.md) |  | [optional] 

## Methods

### NewCFCertificationRequest

`func NewCFCertificationRequest(dn string, dnElements []CFDistinguishedName, keyType string, pem string, ) *CFCertificationRequest`

NewCFCertificationRequest instantiates a new CFCertificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCFCertificationRequestWithDefaults

`func NewCFCertificationRequestWithDefaults() *CFCertificationRequest`

NewCFCertificationRequestWithDefaults instantiates a new CFCertificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDn

`func (o *CFCertificationRequest) GetDn() string`

GetDn returns the Dn field if non-nil, zero value otherwise.

### GetDnOk

`func (o *CFCertificationRequest) GetDnOk() (*string, bool)`

GetDnOk returns a tuple with the Dn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDn

`func (o *CFCertificationRequest) SetDn(v string)`

SetDn sets Dn field to given value.


### GetDnElements

`func (o *CFCertificationRequest) GetDnElements() []CFDistinguishedName`

GetDnElements returns the DnElements field if non-nil, zero value otherwise.

### GetDnElementsOk

`func (o *CFCertificationRequest) GetDnElementsOk() (*[]CFDistinguishedName, bool)`

GetDnElementsOk returns a tuple with the DnElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnElements

`func (o *CFCertificationRequest) SetDnElements(v []CFDistinguishedName)`

SetDnElements sets DnElements field to given value.


### GetKeyType

`func (o *CFCertificationRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CFCertificationRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CFCertificationRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetPem

`func (o *CFCertificationRequest) GetPem() string`

GetPem returns the Pem field if non-nil, zero value otherwise.

### GetPemOk

`func (o *CFCertificationRequest) GetPemOk() (*string, bool)`

GetPemOk returns a tuple with the Pem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPem

`func (o *CFCertificationRequest) SetPem(v string)`

SetPem sets Pem field to given value.


### GetSans

`func (o *CFCertificationRequest) GetSans() []SubjectAlternateName`

GetSans returns the Sans field if non-nil, zero value otherwise.

### GetSansOk

`func (o *CFCertificationRequest) GetSansOk() (*[]SubjectAlternateName, bool)`

GetSansOk returns a tuple with the Sans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSans

`func (o *CFCertificationRequest) SetSans(v []SubjectAlternateName)`

SetSans sets Sans field to given value.

### HasSans

`func (o *CFCertificationRequest) HasSans() bool`

HasSans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


