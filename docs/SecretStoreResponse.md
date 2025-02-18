# SecretStoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | [**CFCertificate**](CFCertificate.md) |  | 

## Methods

### NewSecretStoreResponse

`func NewSecretStoreResponse(certificate CFCertificate, ) *SecretStoreResponse`

NewSecretStoreResponse instantiates a new SecretStoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStoreResponseWithDefaults

`func NewSecretStoreResponseWithDefaults() *SecretStoreResponse`

NewSecretStoreResponseWithDefaults instantiates a new SecretStoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *SecretStoreResponse) GetCertificate() CFCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SecretStoreResponse) GetCertificateOk() (*CFCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SecretStoreResponse) SetCertificate(v CFCertificate)`

SetCertificate sets Certificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


