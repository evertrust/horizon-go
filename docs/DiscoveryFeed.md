# DiscoveryFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaign** | **string** | The name of the discovery campaign to feed into | 
**SessionId** | Pointer to **NullableString** | The ID of the previously opened discovery feed session | [optional] 
**Code** | Pointer to **NullableString** | The code of the event to raise in the discovery events | [optional] 
**Certificate** | **string** | The PEM-encoded certificate to feed the discovery campaign with | 
**HostDiscoveryData** | [**HostDiscoveryData**](HostDiscoveryData.md) | The host discovery data to feed the discovery campaign with (discovery metadata) | 
**Metadata** | Pointer to [**[]CertificateMetadata**](CertificateMetadata.md) | The list of certificate metadata to feed the discovery campaign with | [optional] 
**PrivateKey** | Pointer to **NullableString** | The PEM-encoded private key to feed the discovery campaign with | [optional] 

## Methods

### NewDiscoveryFeed

`func NewDiscoveryFeed(campaign string, certificate string, hostDiscoveryData HostDiscoveryData, ) *DiscoveryFeed`

NewDiscoveryFeed instantiates a new DiscoveryFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryFeedWithDefaults

`func NewDiscoveryFeedWithDefaults() *DiscoveryFeed`

NewDiscoveryFeedWithDefaults instantiates a new DiscoveryFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaign

`func (o *DiscoveryFeed) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *DiscoveryFeed) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *DiscoveryFeed) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.


### GetSessionId

`func (o *DiscoveryFeed) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *DiscoveryFeed) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *DiscoveryFeed) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *DiscoveryFeed) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *DiscoveryFeed) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *DiscoveryFeed) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetCode

`func (o *DiscoveryFeed) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DiscoveryFeed) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DiscoveryFeed) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DiscoveryFeed) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *DiscoveryFeed) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *DiscoveryFeed) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCertificate

`func (o *DiscoveryFeed) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *DiscoveryFeed) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *DiscoveryFeed) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetHostDiscoveryData

`func (o *DiscoveryFeed) GetHostDiscoveryData() HostDiscoveryData`

GetHostDiscoveryData returns the HostDiscoveryData field if non-nil, zero value otherwise.

### GetHostDiscoveryDataOk

`func (o *DiscoveryFeed) GetHostDiscoveryDataOk() (*HostDiscoveryData, bool)`

GetHostDiscoveryDataOk returns a tuple with the HostDiscoveryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDiscoveryData

`func (o *DiscoveryFeed) SetHostDiscoveryData(v HostDiscoveryData)`

SetHostDiscoveryData sets HostDiscoveryData field to given value.


### GetMetadata

`func (o *DiscoveryFeed) GetMetadata() []CertificateMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DiscoveryFeed) GetMetadataOk() (*[]CertificateMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DiscoveryFeed) SetMetadata(v []CertificateMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DiscoveryFeed) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *DiscoveryFeed) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *DiscoveryFeed) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPrivateKey

`func (o *DiscoveryFeed) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *DiscoveryFeed) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *DiscoveryFeed) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *DiscoveryFeed) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *DiscoveryFeed) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *DiscoveryFeed) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


