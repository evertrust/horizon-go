# Pkcs12ContentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | [**CFCertificate**](CFCertificate.md) |  | 
**PrivateKey** | **string** |  | 

## Methods

### NewPkcs12ContentResponse

`func NewPkcs12ContentResponse(certificate CFCertificate, privateKey string, ) *Pkcs12ContentResponse`

NewPkcs12ContentResponse instantiates a new Pkcs12ContentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkcs12ContentResponseWithDefaults

`func NewPkcs12ContentResponseWithDefaults() *Pkcs12ContentResponse`

NewPkcs12ContentResponseWithDefaults instantiates a new Pkcs12ContentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *Pkcs12ContentResponse) GetCertificate() CFCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Pkcs12ContentResponse) GetCertificateOk() (*CFCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Pkcs12ContentResponse) SetCertificate(v CFCertificate)`

SetCertificate sets Certificate field to given value.


### GetPrivateKey

`func (o *Pkcs12ContentResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *Pkcs12ContentResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *Pkcs12ContentResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


