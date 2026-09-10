# IntegratedCAConnectorAsyncParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateCacheDuration** | **string** | Duration for which certificates are kept in the connector&#39;s internal cache | 
**ResponseDelay** | **string** | Delay before the certificate is made available after an enroll request | 

## Methods

### NewIntegratedCAConnectorAsyncParams

`func NewIntegratedCAConnectorAsyncParams(certificateCacheDuration string, responseDelay string, ) *IntegratedCAConnectorAsyncParams`

NewIntegratedCAConnectorAsyncParams instantiates a new IntegratedCAConnectorAsyncParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegratedCAConnectorAsyncParamsWithDefaults

`func NewIntegratedCAConnectorAsyncParamsWithDefaults() *IntegratedCAConnectorAsyncParams`

NewIntegratedCAConnectorAsyncParamsWithDefaults instantiates a new IntegratedCAConnectorAsyncParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateCacheDuration

`func (o *IntegratedCAConnectorAsyncParams) GetCertificateCacheDuration() string`

GetCertificateCacheDuration returns the CertificateCacheDuration field if non-nil, zero value otherwise.

### GetCertificateCacheDurationOk

`func (o *IntegratedCAConnectorAsyncParams) GetCertificateCacheDurationOk() (*string, bool)`

GetCertificateCacheDurationOk returns a tuple with the CertificateCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCacheDuration

`func (o *IntegratedCAConnectorAsyncParams) SetCertificateCacheDuration(v string)`

SetCertificateCacheDuration sets CertificateCacheDuration field to given value.


### GetResponseDelay

`func (o *IntegratedCAConnectorAsyncParams) GetResponseDelay() string`

GetResponseDelay returns the ResponseDelay field if non-nil, zero value otherwise.

### GetResponseDelayOk

`func (o *IntegratedCAConnectorAsyncParams) GetResponseDelayOk() (*string, bool)`

GetResponseDelayOk returns a tuple with the ResponseDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDelay

`func (o *IntegratedCAConnectorAsyncParams) SetResponseDelay(v string)`

SetResponseDelay sets ResponseDelay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


