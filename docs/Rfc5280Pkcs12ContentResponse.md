# Rfc5280Pkcs12ContentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | [**CFCertificateResponse**](CFCertificateResponse.md) |  | 
**PrivateKey** | **string** |  | 

## Methods

### NewRfc5280Pkcs12ContentResponse

`func NewRfc5280Pkcs12ContentResponse(certificate CFCertificateResponse, privateKey string, ) *Rfc5280Pkcs12ContentResponse`

NewRfc5280Pkcs12ContentResponse instantiates a new Rfc5280Pkcs12ContentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRfc5280Pkcs12ContentResponseWithDefaults

`func NewRfc5280Pkcs12ContentResponseWithDefaults() *Rfc5280Pkcs12ContentResponse`

NewRfc5280Pkcs12ContentResponseWithDefaults instantiates a new Rfc5280Pkcs12ContentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *Rfc5280Pkcs12ContentResponse) GetCertificate() CFCertificateResponse`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *Rfc5280Pkcs12ContentResponse) GetCertificateOk() (*CFCertificateResponse, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *Rfc5280Pkcs12ContentResponse) SetCertificate(v CFCertificateResponse)`

SetCertificate sets Certificate field to given value.


### GetPrivateKey

`func (o *Rfc5280Pkcs12ContentResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *Rfc5280Pkcs12ContentResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *Rfc5280Pkcs12ContentResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


