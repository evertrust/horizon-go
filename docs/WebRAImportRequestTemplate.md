# WebRAImportRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to [**NullableCertificateContactEmailElement**](CertificateContactEmailElement.md) | The contact email for this certificate | [optional] 
**DiscoveryData** | Pointer to [**HostDiscoveryData**](HostDiscoveryData.md) | The host discovery data associated with the certificate (discovery metadata) | [optional] 
**DiscoveryInfo** | Pointer to [**NullableDiscoveryInfo**](DiscoveryInfo.md) | Information about the discovery of this certificate | [optional] 
**Labels** | Pointer to [**[]RequestLabelElement**](RequestLabelElement.md) | The labels for this certificate | [optional] 
**Metadata** | Pointer to [**[]CertificateMetadataElement**](CertificateMetadataElement.md) | The technical metadata for this certificate | [optional] 
**Owner** | Pointer to [**NullableCertificateOwnerElement**](CertificateOwnerElement.md) | The owner for this certificate | [optional] 
**PrivateKey** | Pointer to **NullableString** | The PEM-encoded private key associated with the certificate. Mandatory if target profile has escrow enabled, forbidden otherwise | [optional] 
**Team** | Pointer to [**NullableCertificateTeamElement**](CertificateTeamElement.md) | The team for this certificate | [optional] 
**ThirdPartyData** | Pointer to [**[]ThirdPartyItem**](ThirdPartyItem.md) | The third party data associated with the certificate | [optional] 

## Methods

### NewWebRAImportRequestTemplate

`func NewWebRAImportRequestTemplate() *WebRAImportRequestTemplate`

NewWebRAImportRequestTemplate instantiates a new WebRAImportRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebRAImportRequestTemplateWithDefaults

`func NewWebRAImportRequestTemplateWithDefaults() *WebRAImportRequestTemplate`

NewWebRAImportRequestTemplateWithDefaults instantiates a new WebRAImportRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *WebRAImportRequestTemplate) GetContactEmail() CertificateContactEmailElement`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *WebRAImportRequestTemplate) GetContactEmailOk() (*CertificateContactEmailElement, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *WebRAImportRequestTemplate) SetContactEmail(v CertificateContactEmailElement)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *WebRAImportRequestTemplate) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *WebRAImportRequestTemplate) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *WebRAImportRequestTemplate) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetDiscoveryData

`func (o *WebRAImportRequestTemplate) GetDiscoveryData() HostDiscoveryData`

GetDiscoveryData returns the DiscoveryData field if non-nil, zero value otherwise.

### GetDiscoveryDataOk

`func (o *WebRAImportRequestTemplate) GetDiscoveryDataOk() (*HostDiscoveryData, bool)`

GetDiscoveryDataOk returns a tuple with the DiscoveryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryData

`func (o *WebRAImportRequestTemplate) SetDiscoveryData(v HostDiscoveryData)`

SetDiscoveryData sets DiscoveryData field to given value.

### HasDiscoveryData

`func (o *WebRAImportRequestTemplate) HasDiscoveryData() bool`

HasDiscoveryData returns a boolean if a field has been set.

### GetDiscoveryInfo

`func (o *WebRAImportRequestTemplate) GetDiscoveryInfo() DiscoveryInfo`

GetDiscoveryInfo returns the DiscoveryInfo field if non-nil, zero value otherwise.

### GetDiscoveryInfoOk

`func (o *WebRAImportRequestTemplate) GetDiscoveryInfoOk() (*DiscoveryInfo, bool)`

GetDiscoveryInfoOk returns a tuple with the DiscoveryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryInfo

`func (o *WebRAImportRequestTemplate) SetDiscoveryInfo(v DiscoveryInfo)`

SetDiscoveryInfo sets DiscoveryInfo field to given value.

### HasDiscoveryInfo

`func (o *WebRAImportRequestTemplate) HasDiscoveryInfo() bool`

HasDiscoveryInfo returns a boolean if a field has been set.

### SetDiscoveryInfoNil

`func (o *WebRAImportRequestTemplate) SetDiscoveryInfoNil(b bool)`

 SetDiscoveryInfoNil sets the value for DiscoveryInfo to be an explicit nil

### UnsetDiscoveryInfo
`func (o *WebRAImportRequestTemplate) UnsetDiscoveryInfo()`

UnsetDiscoveryInfo ensures that no value is present for DiscoveryInfo, not even an explicit nil
### GetLabels

`func (o *WebRAImportRequestTemplate) GetLabels() []RequestLabelElement`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *WebRAImportRequestTemplate) GetLabelsOk() (*[]RequestLabelElement, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *WebRAImportRequestTemplate) SetLabels(v []RequestLabelElement)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *WebRAImportRequestTemplate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *WebRAImportRequestTemplate) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *WebRAImportRequestTemplate) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMetadata

`func (o *WebRAImportRequestTemplate) GetMetadata() []CertificateMetadataElement`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebRAImportRequestTemplate) GetMetadataOk() (*[]CertificateMetadataElement, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebRAImportRequestTemplate) SetMetadata(v []CertificateMetadataElement)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebRAImportRequestTemplate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *WebRAImportRequestTemplate) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *WebRAImportRequestTemplate) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOwner

`func (o *WebRAImportRequestTemplate) GetOwner() CertificateOwnerElement`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *WebRAImportRequestTemplate) GetOwnerOk() (*CertificateOwnerElement, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *WebRAImportRequestTemplate) SetOwner(v CertificateOwnerElement)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *WebRAImportRequestTemplate) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *WebRAImportRequestTemplate) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *WebRAImportRequestTemplate) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetPrivateKey

`func (o *WebRAImportRequestTemplate) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *WebRAImportRequestTemplate) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *WebRAImportRequestTemplate) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *WebRAImportRequestTemplate) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *WebRAImportRequestTemplate) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *WebRAImportRequestTemplate) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetTeam

`func (o *WebRAImportRequestTemplate) GetTeam() CertificateTeamElement`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *WebRAImportRequestTemplate) GetTeamOk() (*CertificateTeamElement, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *WebRAImportRequestTemplate) SetTeam(v CertificateTeamElement)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *WebRAImportRequestTemplate) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *WebRAImportRequestTemplate) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *WebRAImportRequestTemplate) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetThirdPartyData

`func (o *WebRAImportRequestTemplate) GetThirdPartyData() []ThirdPartyItem`

GetThirdPartyData returns the ThirdPartyData field if non-nil, zero value otherwise.

### GetThirdPartyDataOk

`func (o *WebRAImportRequestTemplate) GetThirdPartyDataOk() (*[]ThirdPartyItem, bool)`

GetThirdPartyDataOk returns a tuple with the ThirdPartyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyData

`func (o *WebRAImportRequestTemplate) SetThirdPartyData(v []ThirdPartyItem)`

SetThirdPartyData sets ThirdPartyData field to given value.

### HasThirdPartyData

`func (o *WebRAImportRequestTemplate) HasThirdPartyData() bool`

HasThirdPartyData returns a boolean if a field has been set.

### SetThirdPartyDataNil

`func (o *WebRAImportRequestTemplate) SetThirdPartyDataNil(b bool)`

 SetThirdPartyDataNil sets the value for ThirdPartyData to be an explicit nil

### UnsetThirdPartyData
`func (o *WebRAImportRequestTemplate) UnsetThirdPartyData()`

UnsetThirdPartyData ensures that no value is present for ThirdPartyData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


