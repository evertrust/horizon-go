# SecretStoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | The PEM-encoded certificate | 
**Value** | **string** | The PEM-encoded private key of the certificate to store | 

## Methods

### NewSecretStoreRequest

`func NewSecretStoreRequest(certificate string, value string, ) *SecretStoreRequest`

NewSecretStoreRequest instantiates a new SecretStoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretStoreRequestWithDefaults

`func NewSecretStoreRequestWithDefaults() *SecretStoreRequest`

NewSecretStoreRequestWithDefaults instantiates a new SecretStoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *SecretStoreRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SecretStoreRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SecretStoreRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetValue

`func (o *SecretStoreRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SecretStoreRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SecretStoreRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


