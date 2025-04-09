# CertificateFindRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal id of the certificate to find | 
**Pem** | **string** | PEM encoded certificate | 

## Methods

### NewCertificateFindRequest

`func NewCertificateFindRequest(id string, pem string, ) *CertificateFindRequest`

NewCertificateFindRequest instantiates a new CertificateFindRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateFindRequestWithDefaults

`func NewCertificateFindRequestWithDefaults() *CertificateFindRequest`

NewCertificateFindRequestWithDefaults instantiates a new CertificateFindRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateFindRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateFindRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateFindRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPem

`func (o *CertificateFindRequest) GetPem() string`

GetPem returns the Pem field if non-nil, zero value otherwise.

### GetPemOk

`func (o *CertificateFindRequest) GetPemOk() (*string, bool)`

GetPemOk returns a tuple with the Pem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPem

`func (o *CertificateFindRequest) SetPem(v string)`

SetPem sets Pem field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


