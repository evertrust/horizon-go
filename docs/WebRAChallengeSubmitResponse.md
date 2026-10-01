# WebRAChallengeSubmitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | The certificate that was enrolled (PEM) | 
**Pkcs12** | Pointer to **string** | The generated PKCS#12, only returned in centralized mode. It is encrypted with the submitted challenge as its password, in DER Base64 format. | [optional] 

## Methods

### NewWebRAChallengeSubmitResponse

`func NewWebRAChallengeSubmitResponse(certificate string, ) *WebRAChallengeSubmitResponse`

NewWebRAChallengeSubmitResponse instantiates a new WebRAChallengeSubmitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAChallengeSubmitResponseWithDefaults

`func NewWebRAChallengeSubmitResponseWithDefaults() *WebRAChallengeSubmitResponse`

NewWebRAChallengeSubmitResponseWithDefaults instantiates a new WebRAChallengeSubmitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *WebRAChallengeSubmitResponse) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *WebRAChallengeSubmitResponse) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *WebRAChallengeSubmitResponse) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetPkcs12

`func (o *WebRAChallengeSubmitResponse) GetPkcs12() string`

GetPkcs12 returns the Pkcs12 field if non-nil, zero value otherwise.

### GetPkcs12Ok

`func (o *WebRAChallengeSubmitResponse) GetPkcs12Ok() (*string, bool)`

GetPkcs12Ok returns a tuple with the Pkcs12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkcs12

`func (o *WebRAChallengeSubmitResponse) SetPkcs12(v string)`

SetPkcs12 sets Pkcs12 field to given value.

### HasPkcs12

`func (o *WebRAChallengeSubmitResponse) HasPkcs12() bool`

HasPkcs12 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


