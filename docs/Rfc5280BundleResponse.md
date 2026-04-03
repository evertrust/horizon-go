# Rfc5280BundleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Value** | [**[]CFCertificateResponse**](CFCertificateResponse.md) |  | 

## Methods

### NewRfc5280BundleResponse

`func NewRfc5280BundleResponse(type_ string, value []CFCertificateResponse, ) *Rfc5280BundleResponse`

NewRfc5280BundleResponse instantiates a new Rfc5280BundleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRfc5280BundleResponseWithDefaults

`func NewRfc5280BundleResponseWithDefaults() *Rfc5280BundleResponse`

NewRfc5280BundleResponseWithDefaults instantiates a new Rfc5280BundleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Rfc5280BundleResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Rfc5280BundleResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Rfc5280BundleResponse) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *Rfc5280BundleResponse) GetValue() []CFCertificateResponse`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Rfc5280BundleResponse) GetValueOk() (*[]CFCertificateResponse, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Rfc5280BundleResponse) SetValue(v []CFCertificateResponse)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


